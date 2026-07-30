package hxrt

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"mime"
	"mime/multipart"
	"net"
	"net/http"
	"net/http/httptrace"
	"net/textproto"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// HttpRequest is the opaque native request builder used by staged sys.Http.
//
// What: Retains URL, method, parameters, headers, body, proxy, timeout, and an
// optional typed socket until synchronous execution.
// Why: These are Go transport inputs, but none of them should expose or own a
// generated Haxe Http object or its callback policy.
// How: Typed builder functions populate this value; HttpRequestStartExchange
// consumes it once and returns a live representation-neutral HttpExchange.
type HttpRequest struct {
	rawURL    string
	method    string
	hasMethod bool
	post      bool
	timeout   float64
	params    []httpRequestParameter
	headers   []httpNameValue
	body      []byte
	hasBody   bool
	upload    *httpMultipartUpload
	proxyURL  *url.URL
	socket    *SocketHandle
}

// httpNameValue keeps one request-header occurrence.
//
// What: Stores each staged header as its own typed pair.
// Why: Repeated values are part of HttpBase's source contract and a Go map
// would make the native boundary silently collapse or reorder that authority.
// How: Header application walks this slice once and uses http.Header.Add for
// ordinary fields while routing Go-special fields through explicit policy.
type httpNameValue struct {
	name  string
	value string
}

// httpRequestParameter keeps one parameter in its two required spellings.
//
// What: Carries the raw source value and the staged Haxe percent-encoded value.
// Why: Multipart needs raw text, while query/form bytes must use Haxe's encoder
// without asking Go to parse, sort, collapse, or re-encode the collection.
// How: The native builder consumes these entries in source order.
type httpRequestParameter struct {
	name         string
	value        string
	encodedName  string
	encodedValue string
}

// httpMultipartUpload is the native framing description for one staged upload.
//
// What: Retains only scalar multipart metadata and the declared byte count.
// Why: Generated Haxe must read its Input on the public caller, never from a Go
// transport goroutine, while native framing still needs these fixed values.
// How: StartExchange creates a native pipe-backed body and exposes its typed
// upload sink for source-driven bounded writes.
type httpMultipartUpload struct {
	parameter string
	filename  string
	mimeType  string
	size      int
}

// httpMultipartPipeBody is the net/http-owned request-body side of an upload.
//
// What: Emits deterministic prefix bytes, caller-pumped pipe bytes, then the
// closing delimiter while implementing the concurrent Close contract.
// Why: A no-op request-body closer cannot unblock a staged upload write after
// timeout, cancellation, server close, or an early response.
// How: CloseWithError on the pipe reader releases a blocked native sink write;
// the read phase itself remains single-consumer transport state.
type httpMultipartPipeBody struct {
	prefix    *bytes.Reader
	pipe      *io.PipeReader
	tail      *bytes.Reader
	phase     int
	closeOnce sync.Once
}

// HttpUploadSink is the opaque native destination pumped by staged sys.Http.
//
// What: Owns the pipe writer and exact declared-size progress for one upload.
// Why: Haxe Input calls must stay on the synchronous source caller, while each
// native write must still be cancellable by transport resource closure.
// How: One staged writer serializes bounded immutable chunks; abort may close
// the writer concurrently to release a blocked call.
type HttpUploadSink struct {
	writer   *io.PipeWriter
	expected int64
	timeout  time.Duration
	timed    bool

	writeMu sync.Mutex
	stateMu sync.Mutex
	written int64
	closed  bool
	err     error
}

// HttpExchange is the opaque live response owned by staged sys.Http.
//
// What: Retains response headers, a live bounded body reader, cancellation, and
// the one-use native transport until staged source completes or aborts it.
// Why: Returning one fully buffered body delays callbacks, loses partial bytes
// on a later read failure, and makes retained memory proportional to the body.
// How: Staged source reads immutable chunks through typed accessors and closes
// or cancels this handle exactly once; cleanup is idempotent for error paths.
type HttpExchange struct {
	status      int
	headerNames []string
	headers     http.Header
	err         *string
	response    *http.Response
	transport   *http.Transport
	socket      *SocketHandle
	upload      *HttpUploadSink
	requestBody io.Closer
	connection  net.Conn
	cancel      context.CancelFunc
	timeout     time.Duration
	timed       bool
	ready       chan struct{}
	readyOnce   sync.Once
	cleanupOnce sync.Once
	stateMu     sync.Mutex
	closed      bool
}

// HttpReadResult preserves one body read and its terminal state together.
//
// What: Carries at most the requested immutable bytes plus EOF or a native read
// failure from the same call.
// Why: Go readers may legally return useful bytes and an error together; using
// only the error discards Haxe-visible response progress.
// How: Staged source writes body first, then interprets error or EOF.
type HttpReadResult struct {
	body *ByteView
	err  *string
	eof  bool
}

// HttpRequestNew creates one typed native request builder.
//
// What: Selects GET or POST unless an explicit method token is supplied.
// Why: Haxe customRequest writes the caller's token exactly; changing its case
// changes observable wire behavior and can select a different server handler.
// How: Retain both the exact string and whether it was explicitly supplied so
// an empty token becomes a validation error instead of Go's implicit GET.
func HttpRequestNew(rawURL *string, post bool, method *string, timeout float64) *HttpRequest {
	selectedMethod := http.MethodGet
	if post {
		selectedMethod = http.MethodPost
	}
	hasMethod := method != nil
	if method != nil {
		selectedMethod = *StdString(method)
	}
	return &HttpRequest{
		rawURL:    *StdString(rawURL),
		method:    selectedMethod,
		hasMethod: hasMethod,
		post:      post,
		timeout:   timeout,
		params:    make([]httpRequestParameter, 0),
		headers:   make([]httpNameValue, 0),
	}
}

// HttpRequestAddParameter appends one source-ordered query/form/multipart field.
//
// What: Retains the raw name/value for multipart and the staged Haxe percent
// encoding for query/form serialization.
// Why: HttpBase add/set semantics preserve ordered multiplicity, and URL
// encoding is Haxe-visible library policy rather than a Go url.Values policy.
// How: Staged sys.Http passes both representations; native execution never
// reparses, sorts, collapses, or re-encodes them.
func HttpRequestAddParameter(request *HttpRequest, name *string, value *string, encodedName *string, encodedValue *string) {
	if request == nil {
		return
	}
	request.params = append(request.params, httpRequestParameter{
		name:         *StdString(name),
		value:        *StdString(value),
		encodedName:  *StdString(encodedName),
		encodedValue: *StdString(encodedValue),
	})
}

// HttpRequestAddHeader appends one source-ordered native header entry.
// Public replacement semantics remain in haxe.http.HttpBase before entries
// cross this boundary.
func HttpRequestAddHeader(request *HttpRequest, name *string, value *string) {
	if request == nil {
		return
	}
	request.headers = append(request.headers, httpNameValue{
		name:  *StdString(name),
		value: *StdString(value),
	})
}

// HttpRequestSetBodyString installs an explicit UTF-8 request body.
func HttpRequestSetBodyString(request *HttpRequest, value *string) {
	if request == nil {
		return
	}
	request.body = []byte(*StdString(value))
	request.hasBody = true
}

// HttpRequestSetBodyView installs an explicit immutable native byte view.
func HttpRequestSetBodyView(request *HttpRequest, value *ByteView) {
	if request == nil {
		return
	}
	if value == nil {
		request.body = []byte{}
	} else {
		request.body = value.raw
	}
	request.hasBody = true
}

// HttpRequestSetMultipartUpload installs one bounded streaming file body.
//
// What: Records the public field metadata and declared byte count.
// Why: fileTransfer must send the caller's Input bytes, but retaining the whole
// payload or a generated callback in HttpRequest would violate ownership.
// How: StartExchange derives one native pipe; staged source later pumps its
// Input through the typed sink on the original synchronous caller.
func HttpRequestSetMultipartUpload(request *HttpRequest, parameter *string, filename *string, mimeType *string, size int) {
	if request == nil {
		return
	}
	request.post = true
	request.upload = &httpMultipartUpload{
		parameter: *StdString(parameter),
		filename:  *StdString(filename),
		mimeType:  *StdString(mimeType),
		size:      size,
	}
}

// HttpRequestSetProxy configures native absolute-target HTTP proxying and HTTPS
// CONNECT negotiation from the staged scalar PROXY authority.
func HttpRequestSetProxy(request *HttpRequest, host *string, port int, user *string, pass *string) {
	if request == nil {
		return
	}
	request.proxyURL = httpProxyURL(host, port, user, pass)
}

// HttpRequestSetSocket supplies the typed socket consumed by customRequest.
//
// What: Retains one caller-owned socket for a plain-HTTP exchange.
// Why: Socket identity is typed, but the shared handle cannot currently retain
// sys.ssl.Socket verification, CA, hostname, or client-certificate policy.
// How: StartExchange consumes and closes it for HTTP and rejects HTTPS before
// transport rather than accidentally applying TLS twice or not at all.
func HttpRequestSetSocket(request *HttpRequest, socket *SocketHandle) {
	if request == nil {
		return
	}
	request.socket = socket
}

var errHTTPProgressTimeout = errors.New("HTTP progress timeout")

// httpProgressTimeout converts the public Float policy into a native deadline.
//
// What: Distinguishes unlimited, immediate, and positive progress budgets.
// Why: time.Duration uses zero as "disabled", while sys.Http defines zero as
// an immediate deadline and negative values as no native deadline.
// How: Reject non-finite values, clamp oversized positive budgets, and round a
// positive sub-nanosecond value up so it cannot silently become unlimited.
func httpProgressTimeout(timeout float64) (time.Duration, bool, error) {
	if math.IsNaN(timeout) || math.IsInf(timeout, 0) {
		return 0, false, errors.New("HTTP timeout must be finite")
	}
	if timeout < 0 {
		return 0, false, nil
	}
	if timeout == 0 {
		return 0, true, errHTTPProgressTimeout
	}
	nanoseconds := timeout * float64(time.Second)
	if nanoseconds >= float64(math.MaxInt64) {
		return time.Duration(math.MaxInt64), true, nil
	}
	duration := time.Duration(nanoseconds)
	if duration <= 0 {
		duration = time.Nanosecond
	}
	return duration, true, nil
}

// httpReturnRedirectResponse keeps the first received 3xx response live.
//
// What: Stops net/http before it dials a redirect destination.
// Why: Haxe sys.Http exposes the received redirect status, headers, and body.
// How: ErrUseLastResponse asks Client.Do to return that response without
// converting the policy decision into a transport error.
func httpReturnRedirectResponse(_ *http.Request, _ []*http.Request) error {
	return http.ErrUseLastResponse
}

// HttpRequestStartExchange starts one Go HTTP exchange.
//
// What: Applies the request builder and returns as soon as response headers or
// a startup transport failure are available, except that multipart requests
// return earlier with a caller-driven upload sink.
// Why: net/http resources and dialing are native capabilities, but body
// buffering would hide response progress and transport-goroutine callbacks
// into generated Haxe would violate source execution ownership.
// How: Validate and consume the builder, create a one-use client, retain its
// resources behind HttpExchange, and run only multipart transport work on a
// native goroutine while the staged caller pumps its Input through the pipe.
func HttpRequestStartExchange(request *HttpRequest) *HttpExchange {
	if request == nil {
		return httpErrorExchange("Invalid URL")
	}
	parsedURL, err := url.Parse(request.rawURL)
	if err != nil || parsedURL == nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return httpRequestErrorExchange(request, "Invalid URL")
	}
	scheme := strings.ToLower(parsedURL.Scheme)
	if scheme != "http" && scheme != "https" {
		return httpRequestErrorExchange(request, "Unsupported HTTP URL scheme")
	}
	if request.hasMethod && request.method == "" {
		return httpRequestErrorExchange(request, "HTTP method must not be empty")
	}
	if request.socket != nil && scheme == "https" {
		return httpRequestErrorExchange(request, "HTTPS custom sockets are not supported")
	}
	timeout, timed, timeoutErr := httpProgressTimeout(request.timeout)
	if timeoutErr != nil {
		return httpRequestErrorExchange(request, timeoutErr.Error())
	}

	var bodyReader io.Reader
	var contentLength int64
	var multipartContentType string
	var requestBody io.Closer
	var uploadSink *HttpUploadSink
	if request.upload != nil {
		multipartBody, sink, length, contentType, buildErr := newHttpMultipartBody(request)
		if buildErr != nil {
			return httpRequestErrorExchange(request, buildErr.Error())
		}
		bodyReader = multipartBody
		requestBody = multipartBody
		uploadSink = sink
		uploadSink.timeout = timeout
		uploadSink.timed = timed
		contentLength = length
		multipartContentType = contentType
	} else if request.hasBody {
		bodyReader = bytes.NewReader(request.body)
		if !request.post {
			httpAppendEncodedParameters(parsedURL, request.params)
		}
	} else if request.post {
		bodyReader = strings.NewReader(httpEncodedParameters(request.params))
	} else {
		httpAppendEncodedParameters(parsedURL, request.params)
	}

	nativeRequest, err := http.NewRequest(request.method, parsedURL.String(), bodyReader)
	if err != nil {
		return httpRequestErrorExchange(request, err.Error())
	}
	if request.upload != nil {
		nativeRequest.ContentLength = contentLength
	}
	if err := applyHttpRequestHeaders(nativeRequest, request, multipartContentType); err != nil {
		return httpRequestErrorExchange(request, err.Error())
	}

	dialer := &net.Dialer{}
	transport := &http.Transport{
		DisableCompression:    true,
		ResponseHeaderTimeout: timeout,
		TLSHandshakeTimeout:   timeout,
	}
	if timed {
		dialer.Timeout = timeout
		transport.DialContext = dialer.DialContext
	}
	if request.proxyURL != nil {
		transport.Proxy = http.ProxyURL(request.proxyURL)
	}
	if request.socket != nil {
		nativeRequest.Close = true
		transport.DisableKeepAlives = true
		var socketDialMu sync.Mutex
		socketConsumed := false
		transport.DialContext = func(ctx context.Context, network string, address string) (net.Conn, error) {
			socketDialMu.Lock()
			defer socketDialMu.Unlock()
			if socketConsumed {
				return nil, io.EOF
			}
			socketConsumed = true
			connection := request.socket.snapshotConn()
			if connection == nil {
				connection, err = dialer.DialContext(ctx, network, address)
				if err != nil {
					return nil, err
				}
				request.socket.installConn(connection)
			}
			return connection, nil
		}
	}

	exchange := &HttpExchange{
		headers:     make(http.Header),
		transport:   transport,
		socket:      request.socket,
		upload:      uploadSink,
		requestBody: requestBody,
		timeout:     timeout,
		timed:       timed,
		ready:       make(chan struct{}),
	}
	trace := &httptrace.ClientTrace{
		GotConn: func(info httptrace.GotConnInfo) {
			_ = info.Conn.SetDeadline(time.Time{})
			exchange.stateMu.Lock()
			exchange.connection = info.Conn
			exchange.stateMu.Unlock()
		},
	}
	tracedContext := httptrace.WithClientTrace(nativeRequest.Context(), trace)
	ctx, cancel := context.WithCancel(tracedContext)
	nativeRequest = nativeRequest.WithContext(ctx)
	exchange.cancel = cancel
	client := &http.Client{
		Transport:     transport,
		CheckRedirect: httpReturnRedirectResponse,
	}

	if uploadSink != nil {
		go exchange.run(client, nativeRequest)
	} else {
		exchange.run(client, nativeRequest)
	}
	return exchange
}

// run owns net/http execution without ever entering generated Haxe.
func (exchange *HttpExchange) run(client *http.Client, request *http.Request) {
	nativeResponse, err := client.Do(request)
	if exchange.requestBody != nil {
		_ = exchange.requestBody.Close()
	}
	if err != nil && nativeResponse != nil && nativeResponse.Body != nil {
		_ = nativeResponse.Body.Close()
	}

	exchange.stateMu.Lock()
	if err != nil {
		exchange.err = StringFromLiteral(err.Error())
	} else {
		headerNames := make([]string, 0, len(nativeResponse.Header))
		for name := range nativeResponse.Header {
			headerNames = append(headerNames, name)
		}
		sort.Strings(headerNames)
		exchange.status = nativeResponse.StatusCode
		exchange.headerNames = headerNames
		exchange.headers = nativeResponse.Header.Clone()
		exchange.response = nativeResponse
	}
	exchange.stateMu.Unlock()
	exchange.signalReady()
}

var errHTTPUploadBodyClosed = errors.New("HTTP upload body closed")

type httpUploadWriteResult struct {
	count int
	err   error
}

func (body *httpMultipartPipeBody) Read(destination []byte) (int, error) {
	if len(destination) == 0 {
		return 0, nil
	}
	for {
		switch body.phase {
		case 0:
			count, err := body.prefix.Read(destination)
			if errors.Is(err, io.EOF) {
				body.phase = 1
				if count > 0 {
					return count, nil
				}
				continue
			}
			return count, err
		case 1:
			count, err := body.pipe.Read(destination)
			if errors.Is(err, io.EOF) {
				body.phase = 2
				if count > 0 {
					return count, nil
				}
				continue
			}
			return count, err
		case 2:
			count, err := body.tail.Read(destination)
			if errors.Is(err, io.EOF) {
				body.phase = 3
				if count > 0 {
					return count, nil
				}
				continue
			}
			return count, err
		default:
			return 0, io.EOF
		}
	}
}

func (body *httpMultipartPipeBody) Close() error {
	body.closeOnce.Do(func() {
		_ = body.pipe.CloseWithError(errHTTPUploadBodyClosed)
	})
	return nil
}

// newHttpMultipartBody builds one atomic multipart framing plan.
//
// What: Derives the prefix, bounded upload reader, closing delimiter, exact
// length, and Content-Type from one newly generated boundary.
// Why: A public reusable boundary can collide with file bytes, while separately
// assembled body and header values can disagree on the delimiter.
// How: multipart.Writer generates the boundary; this function validates all
// header-bearing metadata before returning any body that can reach the network.
func newHttpMultipartBody(request *HttpRequest) (*httpMultipartPipeBody, *HttpUploadSink, int64, string, error) {
	upload := request.upload
	if upload == nil {
		return nil, nil, 0, "", errors.New("multipart upload is missing")
	}
	if upload.size < 0 {
		return nil, nil, 0, "", errors.New("multipart upload size must be non-negative")
	}
	if err := validateHttpMultipartToken("multipart parameter", upload.parameter); err != nil {
		return nil, nil, 0, "", err
	}
	if err := validateHttpMultipartToken("multipart filename", upload.filename); err != nil {
		return nil, nil, 0, "", err
	}
	mediaType, _, err := mime.ParseMediaType(upload.mimeType)
	if err != nil || mediaType == "" {
		return nil, nil, 0, "", errors.New("multipart media type is invalid")
	}
	if err := validateHttpMultipartToken("multipart media type", upload.mimeType); err != nil {
		return nil, nil, 0, "", err
	}

	var prefix bytes.Buffer
	writer := multipart.NewWriter(&prefix)
	boundary := writer.Boundary()
	contentType := writer.FormDataContentType()
	for _, parameter := range request.params {
		if err := validateHttpMultipartToken("multipart field name", parameter.name); err != nil {
			return nil, nil, 0, "", err
		}
		if err := writer.WriteField(parameter.name, parameter.value); err != nil {
			return nil, nil, 0, "", err
		}
	}
	header := make(textproto.MIMEHeader)
	header.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
		httpMultipartQuote(upload.parameter), httpMultipartQuote(upload.filename)))
	header.Set("Content-Type", upload.mimeType)
	if _, err := writer.CreatePart(header); err != nil {
		return nil, nil, 0, "", err
	}
	tailBytes := []byte("\r\n--" + boundary + "--\r\n")
	length := int64(prefix.Len()) + int64(upload.size) + int64(len(tailBytes))
	pipeReader, pipeWriter := io.Pipe()
	body := &httpMultipartPipeBody{
		prefix: bytes.NewReader(prefix.Bytes()),
		pipe:   pipeReader,
		tail:   bytes.NewReader(tailBytes),
	}
	sink := &HttpUploadSink{
		writer:   pipeWriter,
		expected: int64(upload.size),
	}
	return body, sink, length, contentType, nil
}

func httpMultipartQuote(value string) string {
	return strings.NewReplacer("\\", "\\\\", `"`, `\"`).Replace(value)
}

// httpEncodedParameters joins the staged Haxe spellings without normalization.
//
// What: Produces the query/form bytes in source order with repeated entries.
// Why: url.Values would sort names, collapse source intent through Set, and
// replace existing raw-query spelling with Go's encoding policy.
// How: Join the already encoded name/value pairs with the required delimiters.
func httpEncodedParameters(parameters []httpRequestParameter) string {
	encoded := make([]string, 0, len(parameters))
	for _, parameter := range parameters {
		encoded = append(encoded, parameter.encodedName+"="+parameter.encodedValue)
	}
	return strings.Join(encoded, "&")
}

// httpAppendEncodedParameters preserves the caller's existing raw query.
//
// What: Appends staged query fields without reparsing the URL.
// Why: Parsing and re-encoding can reorder, collapse, or rewrite caller-owned
// escapes, including when an explicit body is sent with post == false.
// How: Join the pre-encoded ordered parameters and add one delimiter only when
// both the original query and new fields are present.
func httpAppendEncodedParameters(target *url.URL, parameters []httpRequestParameter) {
	if target == nil {
		return
	}
	encoded := httpEncodedParameters(parameters)
	if encoded == "" {
		return
	}
	if target.RawQuery == "" {
		target.RawQuery = encoded
	} else {
		target.RawQuery += "&" + encoded
	}
}

// applyHttpRequestHeaders installs ordinary and Go-special request fields.
//
// What: Preserves repeated ordinary header values and translates Host,
// Content-Length, Connection, and multipart Content-Type into net/http fields.
// Why: Those names are not ordinary map entries in Go, and silently accepting
// unsupported framing controls would make the staged request differ on the wire.
// How: Validate the ordered entries before dialing, add ordinary values, route
// supported special cases, and reject unimplemented framing controls.
func applyHttpRequestHeaders(nativeRequest *http.Request, request *HttpRequest, multipartContentType string) error {
	hostSet := false
	contentLengthSet := false
	contentTypeSet := false
	for _, header := range request.headers {
		name := strings.TrimSpace(header.name)
		lowerName := strings.ToLower(name)
		switch lowerName {
		case "host":
			if hostSet {
				return errors.New("multiple Host headers are not supported")
			}
			if header.value == "" {
				return errors.New("Host header must not be empty")
			}
			nativeRequest.Host = header.value
			hostSet = true
		case "content-length":
			if contentLengthSet {
				return errors.New("multiple Content-Length headers are not supported")
			}
			declared, err := strconv.ParseInt(strings.TrimSpace(header.value), 10, 64)
			if err != nil || declared < 0 {
				return errors.New("Content-Length header is invalid")
			}
			if nativeRequest.ContentLength >= 0 && nativeRequest.ContentLength != declared {
				return fmt.Errorf("Content-Length header %d does not match request body length %d", declared, nativeRequest.ContentLength)
			}
			nativeRequest.ContentLength = declared
			contentLengthSet = true
		case "connection":
			if !strings.EqualFold(strings.TrimSpace(header.value), "close") {
				return errors.New("only Connection: close is supported")
			}
			nativeRequest.Close = true
		case "transfer-encoding", "trailer", "upgrade":
			return fmt.Errorf("%s request header is not supported", name)
		case "content-type":
			if multipartContentType == "" {
				nativeRequest.Header.Add(name, header.value)
				contentTypeSet = true
				continue
			}
			mediaType, parameters, err := mime.ParseMediaType(header.value)
			if err != nil || !strings.EqualFold(mediaType, "multipart/form-data") {
				return errors.New("multipart Content-Type header is invalid")
			}
			if callerBoundary := parameters["boundary"]; callerBoundary != "" {
				_, generatedParameters, _ := mime.ParseMediaType(multipartContentType)
				if callerBoundary != generatedParameters["boundary"] {
					return errors.New("multipart Content-Type boundary conflicts with generated body boundary")
				}
			}
			contentTypeSet = true
		default:
			nativeRequest.Header.Add(name, header.value)
		}
	}
	if multipartContentType != "" {
		nativeRequest.Header.Set("Content-Type", multipartContentType)
	} else if request.post && !request.hasBody && !contentTypeSet {
		nativeRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	return nil
}

// validateHttpMultipartToken rejects metadata that could create new part headers.
//
// What: Forbids CR, LF, and NUL in multipart header-bearing text.
// Why: Quoting backslashes and quotes alone does not prevent header injection.
// How: Return a deterministic validation error before the request is dialed.
func validateHttpMultipartToken(label string, value string) error {
	if strings.ContainsAny(value, "\r\n\x00") {
		return fmt.Errorf("%s contains a forbidden control character", label)
	}
	return nil
}

// HttpExchangeUploadSink returns the caller-driven sink for a multipart request.
func HttpExchangeUploadSink(exchange *HttpExchange) *HttpUploadSink {
	if exchange == nil {
		return nil
	}
	return exchange.upload
}

// HttpUploadSinkWriteChunk copies one bounded immutable source chunk.
//
// What: Writes one non-empty view without exceeding the declared upload size.
// Why: Source reads remain in staged Haxe, but native cancellation must unblock
// a write that net/http no longer needs after an early terminal event.
// How: Serialize the one staged writer while abort may close the pipe
// concurrently; retain partial native progress before returning its error.
func HttpUploadSinkWriteChunk(sink *HttpUploadSink, chunk *ByteView) *string {
	if sink == nil {
		return StringFromLiteral("HTTP upload sink is unavailable")
	}
	var raw []byte
	if chunk != nil {
		raw = chunk.raw
	}
	if len(raw) == 0 {
		return StringFromLiteral("multipart upload made no progress")
	}

	sink.writeMu.Lock()
	defer sink.writeMu.Unlock()

	sink.stateMu.Lock()
	if sink.closed {
		err := sink.err
		sink.stateMu.Unlock()
		if err == nil {
			err = errHTTPUploadBodyClosed
		}
		return StringFromLiteral(err.Error())
	}
	if int64(len(raw)) > sink.expected-sink.written {
		written := sink.written
		expected := sink.expected
		sink.stateMu.Unlock()
		return StringFromLiteral(fmt.Sprintf(
			"multipart upload exceeds declared size: wrote %d, chunk %d, expected %d",
			written,
			len(raw),
			expected,
		))
	}
	writer := sink.writer
	sink.stateMu.Unlock()

	count := 0
	var err error
	if sink.timed {
		completed := make(chan httpUploadWriteResult, 1)
		go func() {
			written, writeErr := writer.Write(raw)
			completed <- httpUploadWriteResult{count: written, err: writeErr}
		}()
		timer := time.NewTimer(sink.timeout)
		select {
		case result := <-completed:
			if !timer.Stop() {
				<-timer.C
			}
			count = result.count
			err = result.err
		case <-timer.C:
			sink.abort(errHTTPProgressTimeout)
			result := <-completed
			count = result.count
			err = errHTTPProgressTimeout
		}
	} else {
		count, err = writer.Write(raw)
	}
	if err == nil && count != len(raw) {
		err = io.ErrShortWrite
	}
	sink.stateMu.Lock()
	sink.written += int64(count)
	if err != nil && sink.err == nil {
		sink.err = err
	}
	sink.stateMu.Unlock()
	if err != nil {
		return StringFromLiteral(err.Error())
	}
	return nil
}

// HttpUploadSinkFinish validates exact size and releases the multipart tail.
func HttpUploadSinkFinish(sink *HttpUploadSink) *string {
	if sink == nil {
		return StringFromLiteral("HTTP upload sink is unavailable")
	}
	sink.writeMu.Lock()
	defer sink.writeMu.Unlock()

	sink.stateMu.Lock()
	if sink.closed {
		err := sink.err
		sink.stateMu.Unlock()
		if err == nil {
			return nil
		}
		return StringFromLiteral(err.Error())
	}
	if sink.written != sink.expected {
		err := fmt.Errorf(
			"multipart upload ended after %d bytes; expected %d",
			sink.written,
			sink.expected,
		)
		sink.closed = true
		sink.err = err
		writer := sink.writer
		sink.stateMu.Unlock()
		_ = writer.CloseWithError(err)
		return StringFromLiteral(err.Error())
	}
	sink.closed = true
	writer := sink.writer
	sink.stateMu.Unlock()
	if err := writer.Close(); err != nil {
		return StringFromLiteral(err.Error())
	}
	return nil
}

// HttpUploadSinkAbort reports a source failure and releases native transport.
func HttpUploadSinkAbort(sink *HttpUploadSink, message *string) {
	if sink == nil {
		return
	}
	text := "HTTP upload canceled"
	if message != nil && *StdString(message) != "" {
		text = *StdString(message)
	}
	sink.abort(errors.New(text))
}

func (sink *HttpUploadSink) abort(err error) {
	sink.stateMu.Lock()
	if sink.closed {
		sink.stateMu.Unlock()
		return
	}
	sink.closed = true
	if sink.err == nil {
		sink.err = err
	}
	writer := sink.writer
	sink.stateMu.Unlock()
	_ = writer.CloseWithError(err)
}

// HttpExchangeAwaitResponse waits only for native headers or terminal failure.
func HttpExchangeAwaitResponse(exchange *HttpExchange) {
	if exchange != nil {
		exchange.awaitResponse()
	}
}

func (exchange *HttpExchange) awaitResponse() {
	if exchange.ready != nil {
		<-exchange.ready
	}
}

func (exchange *HttpExchange) signalReady() {
	exchange.readyOnce.Do(func() {
		if exchange.ready != nil {
			close(exchange.ready)
		}
	})
}

// HttpExchangeError returns only a startup/transport failure before headers.
func HttpExchangeError(exchange *HttpExchange) *string {
	if exchange == nil {
		return StringFromLiteral("Invalid URL")
	}
	return exchange.err
}

// HttpExchangeStatus returns the native status without classifying it as success.
func HttpExchangeStatus(exchange *HttpExchange) int {
	if exchange == nil {
		return 0
	}
	return exchange.status
}

// HttpExchangeContentLength returns the declared size, -1 when unknown, or -2
// when the value cannot fit the Haxe Int accepted by Output.prepare.
func HttpExchangeContentLength(exchange *HttpExchange) int {
	if exchange == nil || exchange.response == nil || exchange.response.ContentLength < 0 {
		return -1
	}
	const maxHaxeInt = int64(1<<31 - 1)
	if exchange.response.ContentLength > maxHaxeInt {
		return -2
	}
	return int(exchange.response.ContentLength)
}

// HttpExchangeHeaderCount returns the number of distinct native header keys.
func HttpExchangeHeaderCount(exchange *HttpExchange) int {
	if exchange == nil {
		return 0
	}
	return len(exchange.headerNames)
}

// HttpExchangeHeaderName returns one native header name or an empty string.
func HttpExchangeHeaderName(exchange *HttpExchange, index int) *string {
	if exchange == nil || index < 0 || index >= len(exchange.headerNames) {
		return StringFromLiteral("")
	}
	return StringFromLiteral(exchange.headerNames[index])
}

// HttpExchangeHeaderValueCount returns the value count for one indexed key.
func HttpExchangeHeaderValueCount(exchange *HttpExchange, index int) int {
	if exchange == nil || index < 0 || index >= len(exchange.headerNames) {
		return 0
	}
	return len(exchange.headers.Values(exchange.headerNames[index]))
}

// HttpExchangeHeaderValue returns one indexed value or an empty string.
func HttpExchangeHeaderValue(exchange *HttpExchange, headerIndex int, valueIndex int) *string {
	if exchange == nil || headerIndex < 0 || headerIndex >= len(exchange.headerNames) {
		return StringFromLiteral("")
	}
	values := exchange.headers.Values(exchange.headerNames[headerIndex])
	if valueIndex < 0 || valueIndex >= len(values) {
		return StringFromLiteral("")
	}
	return StringFromLiteral(values[valueIndex])
}

// HttpExchangeReadResponseChunk performs one bounded native body read.
//
// What: Returns at most maxBytes while preserving bytes and a same-call error.
// Why: A Go Reader may return both, and Haxe must observe the bytes first.
// How: Allocate only the requested buffer, retain no aggregate body, and encode
// clean io.EOF separately from transfer failures.
func HttpExchangeReadResponseChunk(exchange *HttpExchange, maxBytes int) *HttpReadResult {
	if exchange == nil || exchange.response == nil || exchange.response.Body == nil {
		return httpReadError("HTTP response body is unavailable")
	}
	if maxBytes <= 0 {
		return httpReadError("HTTP response chunk size must be positive")
	}
	exchange.stateMu.Lock()
	closed := exchange.closed
	connection := exchange.connection
	timeout := exchange.timeout
	timed := exchange.timed
	exchange.stateMu.Unlock()
	if closed {
		return httpReadError("HTTP exchange is closed")
	}
	if exchange.response.Body == http.NoBody {
		return &HttpReadResult{body: &ByteView{raw: []byte{}}, eof: true}
	}
	if connection != nil {
		deadline := time.Time{}
		if timed {
			deadline = time.Now().Add(timeout)
		}
		if err := connection.SetReadDeadline(deadline); err != nil {
			return httpReadError(err.Error())
		}
	}

	buffer := make([]byte, maxBytes)
	count, err := exchange.response.Body.Read(buffer)
	if count < 0 || count > len(buffer) {
		return httpReadError(fmt.Sprintf("HTTP response body returned invalid byte count %d", count))
	}
	result := &HttpReadResult{body: &ByteView{raw: buffer[:count:count]}}
	if errors.Is(err, io.EOF) {
		result.eof = true
		return result
	}
	if err != nil {
		result.err = StringFromLiteral(err.Error())
		return result
	}
	if count == 0 {
		result.err = StringFromLiteral(io.ErrNoProgress.Error())
	}
	return result
}

// HttpReadResultBody returns the immutable progress from one native read.
func HttpReadResultBody(result *HttpReadResult) *ByteView {
	if result == nil || result.body == nil {
		return &ByteView{raw: []byte{}}
	}
	return result.body
}

// HttpReadResultError returns a terminal transfer failure, when present.
func HttpReadResultError(result *HttpReadResult) *string {
	if result == nil {
		return StringFromLiteral("HTTP response read result is unavailable")
	}
	return result.err
}

// HttpReadResultEOF reports only clean response-body completion.
func HttpReadResultEOF(result *HttpReadResult) bool {
	return result != nil && result.eof
}

// HttpExchangeClose releases a fully consumed exchange exactly once.
func HttpExchangeClose(exchange *HttpExchange) {
	if exchange != nil {
		exchange.cleanup()
	}
}

// HttpExchangeCancel aborts an incomplete exchange and unblocks native reads.
func HttpExchangeCancel(exchange *HttpExchange) {
	if exchange != nil {
		exchange.cleanup()
	}
}

// cleanup converges every successful, failed, or repeated terminal path.
//
// What: Marks the exchange closed and releases upload, context, request body,
// response body, transport, and optional socket resources.
// Why: Staged callbacks can fail at several points, and none may leak or perform
// duplicate native cleanup.
// How: sync.Once owns the terminal transition; cancel and both body closers
// unblock pending upload or response work before cleanup waits for readiness.
func (exchange *HttpExchange) cleanup() {
	exchange.cleanupOnce.Do(func() {
		exchange.stateMu.Lock()
		exchange.closed = true
		exchange.stateMu.Unlock()
		if exchange.upload != nil {
			exchange.upload.abort(errors.New("HTTP exchange canceled"))
		}
		if exchange.cancel != nil {
			exchange.cancel()
		}
		if exchange.requestBody != nil {
			_ = exchange.requestBody.Close()
		}
		exchange.awaitResponse()
		exchange.stateMu.Lock()
		response := exchange.response
		exchange.stateMu.Unlock()
		if response != nil && response.Body != nil {
			_ = response.Body.Close()
		}
		if exchange.transport != nil {
			exchange.transport.CloseIdleConnections()
		}
		if exchange.socket != nil {
			_ = exchange.socket.close()
		}
	})
}

// HttpProxyDescriptor formats the same proxy used by native execution.
func HttpProxyDescriptor(host *string, port int, user *string, pass *string) *string {
	proxyURL := httpProxyURL(host, port, user, pass)
	if proxyURL == nil {
		return StringFromLiteral("null")
	}
	return StringFromLiteral(proxyURL.String())
}

func httpProxyURL(host *string, port int, user *string, pass *string) *url.URL {
	if host == nil {
		return nil
	}
	rawHost := *StdString(host)
	if rawHost == "" || rawHost == "null" {
		return nil
	}
	hostPort := rawHost
	if !strings.Contains(hostPort, ":") {
		hostPort += ":" + strconv.Itoa(port)
	}
	proxyURL, err := url.Parse("http://" + hostPort)
	if err != nil {
		return nil
	}
	if user != nil {
		username := *StdString(user)
		if username != "" && username != "null" {
			password := ""
			if pass != nil && *StdString(pass) != "null" {
				password = *StdString(pass)
			}
			proxyURL.User = url.UserPassword(username, password)
		}
	}
	return proxyURL
}

func httpRequestErrorExchange(request *HttpRequest, message string) *HttpExchange {
	if request != nil && request.socket != nil {
		_ = request.socket.close()
	}
	return httpErrorExchange(message)
}

func httpErrorExchange(message string) *HttpExchange {
	ready := make(chan struct{})
	close(ready)
	return &HttpExchange{
		err:     StringFromLiteral(message),
		headers: make(http.Header),
		ready:   ready,
	}
}

func httpReadError(message string) *HttpReadResult {
	return &HttpReadResult{body: &ByteView{raw: []byte{}}, err: StringFromLiteral(message)}
}
