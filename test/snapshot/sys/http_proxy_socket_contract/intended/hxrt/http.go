package hxrt

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net"
	"net/http"
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
	rawURL   string
	method   string
	post     bool
	timeout  float64
	params   []httpRequestParameter
	headers  []httpNameValue
	body     []byte
	hasBody  bool
	upload   *httpMultipartUpload
	proxyURL *url.URL
	socket   *SocketHandle
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

// httpMultipartUpload is the native transport description for one staged file upload.
//
// What: Retains scalar multipart metadata, the declared byte count, and a typed
// callback that yields the next bounded immutable byte view.
// Why: The Haxe Input owns source-visible reading semantics, while net/http must
// pull request bytes without buffering the complete file or observing a
// generated Input layout.
// How: HttpRequestStartExchange wraps this description in an io.Reader that
// emits the deterministic multipart prefix, exactly size callback bytes, and
// the closing boundary.
type httpMultipartUpload struct {
	parameter string
	filename  string
	mimeType  string
	size      int
	read      func(int) *ByteView
}

type httpMultipartBody struct {
	prefix    *bytes.Reader
	upload    *httpMultipartUpload
	remaining int
	tail      *bytes.Reader
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
	cancel      context.CancelFunc
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
// What: Selects GET or POST unless an explicit non-empty method is supplied.
// Why: Method normalization is a Go net/http representation detail; the staged
// caller still decides which source-visible request mode applies.
// How: Uppercase the explicit token and initialize ordered native entry lists.
func HttpRequestNew(rawURL *string, post bool, method *string, timeout float64) *HttpRequest {
	selectedMethod := http.MethodGet
	if post {
		selectedMethod = http.MethodPost
	}
	if method != nil {
		candidate := strings.ToUpper(*StdString(method))
		if candidate != "" && candidate != "NULL" {
			selectedMethod = candidate
		}
	}
	return &HttpRequest{
		rawURL:  *StdString(rawURL),
		method:  selectedMethod,
		post:    post,
		timeout: timeout,
		params:  make([]httpRequestParameter, 0),
		headers: make([]httpNameValue, 0),
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
// What: Records the public field metadata, declared byte count, and typed chunk
// callback selected by staged sys.Http.
// Why: fileTransfer must send the caller's Input bytes, but retaining the whole
// payload in HttpRequest would make upload memory proportional to file size.
// How: Defer callback reads until net/http consumes the body; every callback is
// bounded by the destination buffer and remaining declared byte count.
func HttpRequestSetMultipartUpload(request *HttpRequest, parameter *string, filename *string, mimeType *string, size int,
	read func(int) *ByteView) {
	if request == nil {
		return
	}
	request.post = true
	request.upload = &httpMultipartUpload{
		parameter: *StdString(parameter),
		filename:  *StdString(filename),
		mimeType:  *StdString(mimeType),
		size:      size,
		read:      read,
	}
}

// HttpRequestSetProxy configures one explicit HTTP proxy.
func HttpRequestSetProxy(request *HttpRequest, host *string, port int, user *string, pass *string) {
	if request == nil {
		return
	}
	request.proxyURL = httpProxyURL(host, port, user, pass)
}

// HttpRequestSetSocket supplies the typed socket consumed by customRequest.
func HttpRequestSetSocket(request *HttpRequest, socket *SocketHandle) {
	if request == nil {
		return
	}
	request.socket = socket
}

// HttpRequestStartExchange starts one synchronous Go HTTP exchange.
//
// What: Applies the request builder and returns as soon as response headers or
// a startup transport failure are available.
// Why: net/http resources and dialing are native capabilities, but body
// buffering would hide Haxe-visible callback timing and partial progress.
// How: Validate and consume the builder, create a one-use client, retain its
// response body behind HttpExchange, and leave body reads and final cleanup to
// explicit typed capabilities.
func HttpRequestStartExchange(request *HttpRequest) *HttpExchange {
	if request == nil {
		return httpErrorExchange("Invalid URL")
	}
	parsedURL, err := url.Parse(request.rawURL)
	if err != nil || parsedURL == nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return httpRequestErrorExchange(request, "Invalid URL")
	}

	var bodyReader io.Reader
	var contentLength int64
	var multipartContentType string
	if request.upload != nil {
		multipartBody, length, contentType, buildErr := newHttpMultipartBody(request)
		if buildErr != nil {
			return httpRequestErrorExchange(request, buildErr.Error())
		}
		bodyReader = multipartBody
		contentLength = length
		multipartContentType = contentType
	} else if request.post {
		if request.hasBody {
			bodyReader = bytes.NewReader(request.body)
		} else {
			bodyReader = strings.NewReader(httpEncodedParameters(request.params))
		}
	} else {
		encodedParameters := httpEncodedParameters(request.params)
		if encodedParameters != "" {
			if parsedURL.RawQuery == "" {
				parsedURL.RawQuery = encodedParameters
			} else {
				parsedURL.RawQuery += "&" + encodedParameters
			}
		}
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

	transport := &http.Transport{}
	if request.proxyURL != nil {
		transport.Proxy = http.ProxyURL(request.proxyURL)
	}
	if request.socket != nil {
		nativeRequest.Close = true
		transport.DisableKeepAlives = true
		socketConsumed := false
		transport.DialContext = func(ctx context.Context, network string, address string) (net.Conn, error) {
			if socketConsumed {
				return nil, io.EOF
			}
			socketConsumed = true
			connection := request.socket.snapshotConn()
			if connection == nil {
				dialer := &net.Dialer{}
				connection, err = dialer.DialContext(ctx, network, address)
				if err != nil {
					return nil, err
				}
				request.socket.installConn(connection)
			}
			return connection, nil
		}
	}

	timeout := time.Duration(request.timeout * float64(time.Second))
	if timeout <= 0 {
		timeout = 10 * time.Second
	}
	ctx, cancel := context.WithCancel(nativeRequest.Context())
	nativeRequest = nativeRequest.WithContext(ctx)
	client := &http.Client{Transport: transport, Timeout: timeout}
	nativeResponse, err := client.Do(nativeRequest)
	if err != nil {
		cancel()
		transport.CloseIdleConnections()
		if request.socket != nil {
			_ = request.socket.close()
		}
		return httpErrorExchange(err.Error())
	}

	headerNames := make([]string, 0, len(nativeResponse.Header))
	for name := range nativeResponse.Header {
		headerNames = append(headerNames, name)
	}
	sort.Strings(headerNames)
	headers := nativeResponse.Header.Clone()
	return &HttpExchange{
		status:      nativeResponse.StatusCode,
		headerNames: headerNames,
		headers:     headers,
		response:    nativeResponse,
		transport:   transport,
		socket:      request.socket,
		cancel:      cancel,
	}
}

func (body *httpMultipartBody) Read(destination []byte) (int, error) {
	if len(destination) == 0 {
		return 0, nil
	}
	if body.prefix.Len() > 0 {
		return body.prefix.Read(destination)
	}
	if body.remaining > 0 {
		if body.upload.read == nil {
			return 0, errors.New("multipart upload reader is nil")
		}
		limit := len(destination)
		if limit > body.remaining {
			limit = body.remaining
		}
		view := body.upload.read(limit)
		if view == nil {
			return 0, io.ErrUnexpectedEOF
		}
		raw := view.raw
		if len(raw) == 0 {
			return 0, io.ErrNoProgress
		}
		if len(raw) > limit {
			return 0, fmt.Errorf("multipart upload reader returned %d bytes, limit %d", len(raw), limit)
		}
		count := copy(destination, raw)
		body.remaining -= count
		return count, nil
	}
	if body.tail.Len() > 0 {
		return body.tail.Read(destination)
	}
	return 0, io.EOF
}

// newHttpMultipartBody builds one atomic multipart framing plan.
//
// What: Derives the prefix, bounded upload reader, closing delimiter, exact
// length, and Content-Type from one newly generated boundary.
// Why: A public reusable boundary can collide with file bytes, while separately
// assembled body and header values can disagree on the delimiter.
// How: multipart.Writer generates the boundary; this function validates all
// header-bearing metadata before returning any body that can reach the network.
func newHttpMultipartBody(request *HttpRequest) (*httpMultipartBody, int64, string, error) {
	upload := request.upload
	if upload == nil {
		return nil, 0, "", errors.New("multipart upload is missing")
	}
	if upload.size < 0 {
		return nil, 0, "", errors.New("multipart upload size must be non-negative")
	}
	if err := validateHttpMultipartToken("multipart parameter", upload.parameter); err != nil {
		return nil, 0, "", err
	}
	if err := validateHttpMultipartToken("multipart filename", upload.filename); err != nil {
		return nil, 0, "", err
	}
	mediaType, _, err := mime.ParseMediaType(upload.mimeType)
	if err != nil || mediaType == "" {
		return nil, 0, "", errors.New("multipart media type is invalid")
	}
	if err := validateHttpMultipartToken("multipart media type", upload.mimeType); err != nil {
		return nil, 0, "", err
	}

	var prefix bytes.Buffer
	writer := multipart.NewWriter(&prefix)
	boundary := writer.Boundary()
	contentType := writer.FormDataContentType()
	for _, parameter := range request.params {
		if err := validateHttpMultipartToken("multipart field name", parameter.name); err != nil {
			return nil, 0, "", err
		}
		if err := writer.WriteField(parameter.name, parameter.value); err != nil {
			return nil, 0, "", err
		}
	}
	header := make(textproto.MIMEHeader)
	header.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
		httpMultipartQuote(upload.parameter), httpMultipartQuote(upload.filename)))
	header.Set("Content-Type", upload.mimeType)
	if _, err := writer.CreatePart(header); err != nil {
		return nil, 0, "", err
	}
	tailBytes := []byte("\r\n--" + boundary + "--\r\n")
	length := int64(prefix.Len()) + int64(upload.size) + int64(len(tailBytes))
	return &httpMultipartBody{
		prefix:    bytes.NewReader(prefix.Bytes()),
		upload:    upload,
		remaining: upload.size,
		tail:      bytes.NewReader(tailBytes),
	}, length, contentType, nil
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
	exchange.stateMu.Unlock()
	if closed {
		return httpReadError("HTTP exchange is closed")
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
// What: Marks the exchange closed and releases context, body, transport, and
// optional socket resources.
// Why: Staged callbacks can fail at several points, and none may leak or perform
// duplicate native cleanup.
// How: sync.Once owns the terminal transition; Body.Close unblocks a live read.
func (exchange *HttpExchange) cleanup() {
	exchange.cleanupOnce.Do(func() {
		exchange.stateMu.Lock()
		exchange.closed = true
		exchange.stateMu.Unlock()
		if exchange.cancel != nil {
			exchange.cancel()
		}
		if exchange.response != nil && exchange.response.Body != nil {
			_ = exchange.response.Body.Close()
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
	return &HttpExchange{err: StringFromLiteral(message), headers: make(http.Header)}
}

func httpReadError(message string) *HttpReadResult {
	return &HttpReadResult{body: &ByteView{raw: []byte{}}, err: StringFromLiteral(message)}
}
