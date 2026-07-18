package hxrt

import (
	"bytes"
	"context"
	"io"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

// HttpRequest is the opaque native request builder used by staged sys.Http.
//
// What: Retains URL, method, parameters, headers, body, proxy, timeout, and an
// optional typed socket until synchronous execution.
// Why: These are Go transport inputs, but none of them should expose or own a
// generated Haxe Http object or its callback policy.
// How: Typed builder functions populate this value; HttpRequestExecute consumes
// it once and returns a representation-neutral HttpResponse.
type HttpRequest struct {
	rawURL   string
	method   string
	post     bool
	timeout  float64
	params   url.Values
	headers  http.Header
	body     []byte
	hasBody  bool
	proxyURL *url.URL
	socket   *SocketHandle
}

// HttpResponse is the opaque native result inspected by staged sys.Http.
//
// What: Carries a status, immutable body view, normalized native headers, or a
// transport error string.
// Why: Status classification, public maps, callback order, and Haxe exceptions
// remain source policy and therefore must not be decided in this carrier.
// How: Keep native results only; indexed accessors below cross them into staged
// source without maps or generated layouts at the boundary.
type HttpResponse struct {
	status      int
	body        *ByteView
	headerNames []string
	headers     http.Header
	err         *string
}

// HttpRequestNew creates one typed native request builder.
//
// What: Selects GET or POST unless an explicit non-empty method is supplied.
// Why: Method normalization is a Go net/http representation detail; the staged
// caller still decides which source-visible request mode applies.
// How: Uppercase the explicit token and initialize deterministic native maps.
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
		params:  make(url.Values),
		headers: make(http.Header),
	}
}

// HttpRequestAddParameter records the last value for one query/form key.
// This mirrors the previous Go url.Values.Set behavior while staged source owns
// the public setParameter/addParameter collection semantics.
func HttpRequestAddParameter(request *HttpRequest, name *string, value *string) {
	if request == nil {
		return
	}
	request.params.Set(*StdString(name), *StdString(value))
}

// HttpRequestAddHeader records the last value for one native header key.
// Public header ordering and replacement remain in haxe.http.HttpBase.
func HttpRequestAddHeader(request *HttpRequest, name *string, value *string) {
	if request == nil {
		return
	}
	request.headers.Set(*StdString(name), *StdString(value))
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

// HttpRequestExecute performs one synchronous Go HTTP exchange.
//
// What: Applies parameters/body/header/proxy/socket inputs and returns status,
// headers, and response bytes.
// Why: net/http resources and dialing are native capabilities; callback timing
// and HTTP-status error policy are deliberately absent here.
// How: Validate the URL, create a one-use transport/client, fully read and close
// the body, and convert expected native failures into an explicit result string.
func HttpRequestExecute(request *HttpRequest) *HttpResponse {
	if request == nil {
		return httpErrorResponse("Invalid URL")
	}
	parsedURL, err := url.Parse(request.rawURL)
	if err != nil || parsedURL == nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return httpErrorResponse("Invalid URL")
	}

	var bodyReader io.Reader
	if request.post {
		if request.hasBody {
			bodyReader = bytes.NewReader(request.body)
		} else {
			form := parsedURL.Query()
			for name, values := range request.params {
				if len(values) > 0 {
					form.Set(name, values[len(values)-1])
				}
			}
			bodyReader = strings.NewReader(form.Encode())
			if request.headers.Get("Content-Type") == "" {
				request.headers.Set("Content-Type", "application/x-www-form-urlencoded")
			}
		}
	} else {
		query := parsedURL.Query()
		for name, values := range request.params {
			if len(values) > 0 {
				query.Set(name, values[len(values)-1])
			}
		}
		parsedURL.RawQuery = query.Encode()
	}

	nativeRequest, err := http.NewRequest(request.method, parsedURL.String(), bodyReader)
	if err != nil {
		return httpErrorResponse(err.Error())
	}
	for name, values := range request.headers {
		for _, value := range values {
			nativeRequest.Header.Set(name, value)
		}
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
		defer request.socket.close() //nolint:errcheck // response errors remain the public signal
	}

	timeout := time.Duration(request.timeout * float64(time.Second))
	if timeout <= 0 {
		timeout = 10 * time.Second
	}
	client := &http.Client{Transport: transport, Timeout: timeout}
	defer transport.CloseIdleConnections()
	nativeResponse, err := client.Do(nativeRequest)
	if err != nil {
		return httpErrorResponse(err.Error())
	}
	defer nativeResponse.Body.Close() //nolint:errcheck // read failure below is more actionable

	headerNames := make([]string, 0, len(nativeResponse.Header))
	for name := range nativeResponse.Header {
		headerNames = append(headerNames, name)
	}
	sort.Strings(headerNames)
	headers := nativeResponse.Header.Clone()
	rawBody, err := io.ReadAll(nativeResponse.Body)
	if err != nil {
		return &HttpResponse{
			status:      nativeResponse.StatusCode,
			body:        &ByteView{raw: []byte{}},
			headerNames: headerNames,
			headers:     headers,
			err:         StringFromLiteral(err.Error()),
		}
	}
	return &HttpResponse{
		status:      nativeResponse.StatusCode,
		body:        &ByteView{raw: rawBody},
		headerNames: headerNames,
		headers:     headers,
	}
}

// HttpResponseError returns nil on a completed exchange or its native failure.
func HttpResponseError(response *HttpResponse) *string {
	if response == nil {
		return StringFromLiteral("Invalid URL")
	}
	return response.err
}

// HttpResponseStatus returns the native status without classifying it as success.
func HttpResponseStatus(response *HttpResponse) int {
	if response == nil {
		return 0
	}
	return response.status
}

// HttpResponseBody returns an immutable empty view when no body is available.
func HttpResponseBody(response *HttpResponse) *ByteView {
	if response == nil || response.body == nil {
		return &ByteView{raw: []byte{}}
	}
	return response.body
}

// HttpResponseHeaderCount returns the number of distinct native header keys.
func HttpResponseHeaderCount(response *HttpResponse) int {
	if response == nil {
		return 0
	}
	return len(response.headerNames)
}

// HttpResponseHeaderName returns one native header name or an empty string.
func HttpResponseHeaderName(response *HttpResponse, index int) *string {
	if response == nil || index < 0 || index >= len(response.headerNames) {
		return StringFromLiteral("")
	}
	return StringFromLiteral(response.headerNames[index])
}

// HttpResponseHeaderValueCount returns the value count for one indexed key.
func HttpResponseHeaderValueCount(response *HttpResponse, index int) int {
	if response == nil || index < 0 || index >= len(response.headerNames) {
		return 0
	}
	return len(response.headers.Values(response.headerNames[index]))
}

// HttpResponseHeaderValue returns one indexed value or an empty string.
func HttpResponseHeaderValue(response *HttpResponse, headerIndex int, valueIndex int) *string {
	if response == nil || headerIndex < 0 || headerIndex >= len(response.headerNames) {
		return StringFromLiteral("")
	}
	values := response.headers.Values(response.headerNames[headerIndex])
	if valueIndex < 0 || valueIndex >= len(values) {
		return StringFromLiteral("")
	}
	return StringFromLiteral(values[valueIndex])
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

func httpErrorResponse(message string) *HttpResponse {
	return &HttpResponse{err: StringFromLiteral(message), body: &ByteView{raw: []byte{}}, headers: make(http.Header)}
}
