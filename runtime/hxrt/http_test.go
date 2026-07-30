package hxrt

import (
	"errors"
	"io"
	"mime"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func httpTestString(value string) *string {
	return &value
}

// HttpResponse and these helpers keep older request-builder assertions concise
// inside this test package. Production code has no completed-response carrier:
// the helper deliberately drains the new live exchange only for test inspection.
type HttpResponse struct {
	exchange *HttpExchange
	body     *ByteView
	err      *string
}

func HttpRequestExecute(request *HttpRequest) *HttpResponse {
	exchange := HttpRequestStartExchange(request)
	response := &HttpResponse{exchange: exchange, body: &ByteView{raw: []byte{}}, err: HttpExchangeError(exchange)}
	if response.err != nil {
		HttpExchangeCancel(exchange)
		return response
	}
	var body []byte
	for {
		result := HttpExchangeReadResponseChunk(exchange, 1024)
		body = append(body, byteViewRaw(HttpReadResultBody(result))...)
		if err := HttpReadResultError(result); err != nil {
			response.err = err
			break
		}
		if HttpReadResultEOF(result) {
			break
		}
	}
	response.body = &ByteView{raw: body}
	if response.err == nil {
		HttpExchangeClose(exchange)
	} else {
		HttpExchangeCancel(exchange)
	}
	return response
}

func HttpResponseError(response *HttpResponse) *string {
	if response == nil {
		return httpTestString("Invalid URL")
	}
	return response.err
}

func HttpResponseStatus(response *HttpResponse) int {
	if response == nil {
		return 0
	}
	return HttpExchangeStatus(response.exchange)
}

func HttpResponseBody(response *HttpResponse) *ByteView {
	if response == nil || response.body == nil {
		return &ByteView{raw: []byte{}}
	}
	return response.body
}

func HttpResponseHeaderCount(response *HttpResponse) int {
	if response == nil {
		return 0
	}
	return HttpExchangeHeaderCount(response.exchange)
}

func HttpResponseHeaderName(response *HttpResponse, index int) *string {
	if response == nil {
		return httpTestString("")
	}
	return HttpExchangeHeaderName(response.exchange, index)
}

func HttpResponseHeaderValueCount(response *HttpResponse, index int) int {
	if response == nil {
		return 0
	}
	return HttpExchangeHeaderValueCount(response.exchange, index)
}

func HttpResponseHeaderValue(response *HttpResponse, headerIndex int, valueIndex int) *string {
	if response == nil {
		return httpTestString("")
	}
	return HttpExchangeHeaderValue(response.exchange, headerIndex, valueIndex)
}

type httpBytesThenErrorBody struct {
	closed atomic.Int32
	read   bool
}

func (body *httpBytesThenErrorBody) Read(destination []byte) (int, error) {
	if body.read {
		return 0, io.EOF
	}
	body.read = true
	return copy(destination, []byte("part")), io.ErrUnexpectedEOF
}

func (body *httpBytesThenErrorBody) Close() error {
	body.closed.Add(1)
	return nil
}

type httpBlockingBody struct {
	closed    chan struct{}
	started   chan struct{}
	closeOnce atomic.Bool
	startOnce atomic.Bool
}

func (body *httpBlockingBody) Read(_ []byte) (int, error) {
	if body.startOnce.CompareAndSwap(false, true) {
		close(body.started)
	}
	<-body.closed
	return 0, errors.New("body canceled")
}

func (body *httpBlockingBody) Close() error {
	if body.closeOnce.CompareAndSwap(false, true) {
		close(body.closed)
	}
	return nil
}

type httpBoundedBody struct {
	remaining int
	maxRead   int
	closed    atomic.Int32
}

func (body *httpBoundedBody) Read(destination []byte) (int, error) {
	if len(destination) > body.maxRead {
		body.maxRead = len(destination)
	}
	if body.remaining == 0 {
		return 0, io.EOF
	}
	count := len(destination)
	if count > body.remaining {
		count = body.remaining
	}
	for index := 0; index < count; index++ {
		destination[index] = byte(index)
	}
	body.remaining -= count
	return count, nil
}

func (body *httpBoundedBody) Close() error {
	body.closed.Add(1)
	return nil
}

type httpInvalidCountBody struct {
	count int
}

func (body *httpInvalidCountBody) Read(_ []byte) (int, error) {
	return body.count, nil
}

func (body *httpInvalidCountBody) Close() error {
	return nil
}

func TestHttpRequestTypedGetAndResponseBoundary(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if got := strings.Join(request.URL.Query()["q"], ","); got != "first,last" {
			t.Errorf("query q values = %q, want first,last", got)
		}
		if got := request.Header.Get("X-Test"); got != "typed" {
			t.Errorf("X-Test = %q, want typed", got)
		}
		response.Header().Add("X-Multi", "one")
		response.Header().Add("X-Multi", "two")
		response.Header().Set("A-First", "yes")
		response.WriteHeader(http.StatusCreated)
		_, _ = response.Write([]byte("response-body"))
	}))
	defer server.Close()

	request := HttpRequestNew(httpTestString(server.URL), false, nil, 1)
	HttpRequestAddParameter(request, httpTestString("q"), httpTestString("first"), httpTestString("q"), httpTestString("first"))
	HttpRequestAddParameter(request, httpTestString("q"), httpTestString("last"), httpTestString("q"), httpTestString("last"))
	HttpRequestAddHeader(request, httpTestString("X-Test"), httpTestString("typed"))
	response := HttpRequestExecute(request)

	if err := HttpResponseError(response); err != nil {
		t.Fatalf("HttpResponseError = %q, want nil", *err)
	}
	if got := HttpResponseStatus(response); got != http.StatusCreated {
		t.Fatalf("HttpResponseStatus = %d, want %d", got, http.StatusCreated)
	}
	if got := string(byteViewRaw(HttpResponseBody(response))); got != "response-body" {
		t.Fatalf("HttpResponseBody = %q, want response-body", got)
	}

	names := make([]string, 0, HttpResponseHeaderCount(response))
	multiValues := make([]string, 0, 2)
	for index := 0; index < HttpResponseHeaderCount(response); index++ {
		name := *HttpResponseHeaderName(response, index)
		names = append(names, name)
		if strings.EqualFold(name, "X-Multi") {
			for valueIndex := 0; valueIndex < HttpResponseHeaderValueCount(response, index); valueIndex++ {
				multiValues = append(multiValues, *HttpResponseHeaderValue(response, index, valueIndex))
			}
		}
	}
	if !sort.StringsAreSorted(names) {
		t.Fatalf("response header names = %v, want deterministic lexical order", names)
	}
	if got := strings.Join(multiValues, ","); got != "one,two" {
		t.Fatalf("X-Multi values = %q, want one,two", got)
	}
}

func TestHttpRequestPostFormPreservesOriginalQueryContract(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if got := request.URL.Query().Get("base"); got != "from-url" {
			t.Errorf("URL query base = %q, want from-url", got)
		}
		rawBody, err := io.ReadAll(request.Body)
		if err != nil {
			t.Errorf("read request body: %v", err)
		}
		form, err := url.ParseQuery(string(rawBody))
		if err != nil {
			t.Errorf("parse form body: %v", err)
		}
		if got := form.Get("base"); got != "" {
			t.Errorf("form base = %q, want URL query to remain outside the body", got)
		}
		if got := form.Get("field"); got != "value" {
			t.Errorf("form field = %q, want value", got)
		}
		if got := request.Header.Get("Content-Type"); got != "application/x-www-form-urlencoded" {
			t.Errorf("Content-Type = %q", got)
		}
		response.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	request := HttpRequestNew(httpTestString(server.URL+"?base=from-url"), true, nil, 1)
	HttpRequestAddParameter(request, httpTestString("field"), httpTestString("value"), httpTestString("field"), httpTestString("value"))
	response := HttpRequestExecute(request)
	if err := HttpResponseError(response); err != nil {
		t.Fatalf("HttpResponseError = %q, want nil", *err)
	}
}

func TestHttpRequestMultipartStreamsDeclaredInputWithPartialChunks(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if err := request.ParseMultipartForm(1024); err != nil {
			t.Errorf("ParseMultipartForm: %v", err)
			response.WriteHeader(http.StatusBadRequest)
			return
		}
		if got := request.FormValue("note"); got != "hello" {
			t.Errorf("note = %q, want hello", got)
		}
		file, header, err := request.FormFile("asset")
		if err != nil {
			t.Errorf("FormFile(asset): %v", err)
			response.WriteHeader(http.StatusBadRequest)
			return
		}
		defer file.Close()
		payload, err := io.ReadAll(file)
		if err != nil {
			t.Errorf("read multipart file: %v", err)
		}
		if header.Filename != "demo.txt" {
			t.Errorf("filename = %q, want demo.txt", header.Filename)
		}
		if got := header.Header.Get("Content-Type"); got != "text/plain" {
			t.Errorf("Content-Type = %q, want text/plain", got)
		}
		if got := string(payload); got != "payload" {
			t.Errorf("payload = %q, want payload", got)
		}
		response.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	payload := []byte("payload")
	offset := 0
	readCalls := 0
	request := HttpRequestNew(httpTestString(server.URL), true, nil, 1)
	HttpRequestAddParameter(request, httpTestString("note"), httpTestString("hello"), httpTestString("note"), httpTestString("hello"))
	HttpRequestSetMultipartUpload(
		request,
		httpTestString("asset"),
		httpTestString("demo.txt"),
		httpTestString("text/plain"),
		len(payload),
		func(limit int) *ByteView {
			readCalls++
			if offset >= len(payload) {
				return nil
			}
			count := len(payload) - offset
			if count > 2 {
				count = 2
			}
			if count > limit {
				count = limit
			}
			chunk := append([]byte(nil), payload[offset:offset+count]...)
			offset += count
			return &ByteView{raw: chunk}
		},
	)
	response := HttpRequestExecute(request)

	if err := HttpResponseError(response); err != nil {
		t.Fatalf("HttpResponseError = %q, want nil", *err)
	}
	if got := HttpResponseStatus(response); got != http.StatusNoContent {
		t.Fatalf("HttpResponseStatus = %d, want %d", got, http.StatusNoContent)
	}
	if readCalls != 4 {
		t.Fatalf("multipart read calls = %d, want 4 partial reads", readCalls)
	}
	if request.body != nil {
		t.Fatal("multipart payload was retained in the buffered request body")
	}
}

func TestHttpRequestMultipartEarlyEOFAbortsTheExchangeAndReleasesTheServer(t *testing.T) {
	handlerDone := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		defer close(handlerDone)
		_, _ = io.ReadAll(request.Body)
		response.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	sent := false
	request := HttpRequestNew(httpTestString(server.URL), true, nil, 1)
	HttpRequestSetMultipartUpload(
		request,
		httpTestString("asset"),
		httpTestString("short.txt"),
		httpTestString("text/plain"),
		7,
		func(_ int) *ByteView {
			if sent {
				return nil
			}
			sent = true
			return &ByteView{raw: []byte("two")}
		},
	)

	response := HttpRequestExecute(request)
	if err := HttpResponseError(response); err == nil {
		t.Fatal("HttpResponseError = nil, want an early-upload EOF")
	}
	select {
	case <-handlerDone:
	case <-time.After(time.Second):
		t.Fatal("multipart upload failure left the server-side request blocked")
	}
}

func TestHttpRequestPreservesRawQueryParameterAndHeaderMultiplicity(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		wantTarget := "/capture?base=one%20two&base=two%2Bthree&a=one%20space&b=x%2By&a=second"
		if request.RequestURI != wantTarget {
			t.Errorf("request target = %q, want %q", request.RequestURI, wantTarget)
		}
		if got := strings.Join(request.Header.Values("X-Repeat"), ","); got != "one,two" {
			t.Errorf("X-Repeat values = %q, want one,two", got)
		}
		if request.Host != "example.test" {
			t.Errorf("Host = %q, want example.test", request.Host)
		}
		if !request.Close {
			t.Error("Connection: close did not set the native request close policy")
		}
		response.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	request := HttpRequestNew(
		httpTestString(server.URL+"/capture?base=one%20two&base=two%2Bthree"),
		false,
		nil,
		1,
	)
	HttpRequestAddParameter(
		request,
		httpTestString("a"),
		httpTestString("one space"),
		httpTestString("a"),
		httpTestString("one%20space"),
	)
	HttpRequestAddParameter(
		request,
		httpTestString("b"),
		httpTestString("x+y"),
		httpTestString("b"),
		httpTestString("x%2By"),
	)
	HttpRequestAddParameter(
		request,
		httpTestString("a"),
		httpTestString("second"),
		httpTestString("a"),
		httpTestString("second"),
	)
	HttpRequestAddHeader(request, httpTestString("X-Repeat"), httpTestString("one"))
	HttpRequestAddHeader(request, httpTestString("X-Repeat"), httpTestString("two"))
	HttpRequestAddHeader(request, httpTestString("Host"), httpTestString("example.test"))
	HttpRequestAddHeader(request, httpTestString("Connection"), httpTestString("close"))

	result := HttpRequestExecute(request)
	if err := HttpResponseError(result); err != nil {
		t.Fatalf("HttpResponseError = %q, want nil", *err)
	}
}

func TestHttpRequestPostFormKeepsOrderedParametersSeparateFromURLQuery(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if request.RequestURI != "/submit?base=from%20url&base=second" {
			t.Errorf("request target = %q", request.RequestURI)
		}
		rawBody, err := io.ReadAll(request.Body)
		if err != nil {
			t.Errorf("read request body: %v", err)
		}
		if got := string(rawBody); got != "field=one%20space&field=two&other=x%2By" {
			t.Errorf("form body = %q", got)
		}
		response.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	request := HttpRequestNew(httpTestString(server.URL+"/submit?base=from%20url&base=second"), true, nil, 1)
	HttpRequestAddParameter(
		request,
		httpTestString("field"),
		httpTestString("one space"),
		httpTestString("field"),
		httpTestString("one%20space"),
	)
	HttpRequestAddParameter(
		request,
		httpTestString("field"),
		httpTestString("two"),
		httpTestString("field"),
		httpTestString("two"),
	)
	HttpRequestAddParameter(
		request,
		httpTestString("other"),
		httpTestString("x+y"),
		httpTestString("other"),
		httpTestString("x%2By"),
	)

	result := HttpRequestExecute(request)
	if err := HttpResponseError(result); err != nil {
		t.Fatalf("HttpResponseError = %q, want nil", *err)
	}
}

func TestHttpMultipartBoundaryMetadataAndLengthAreAtomic(t *testing.T) {
	type observedRequest struct {
		boundary string
		length   int64
		bodySize int
		values   []string
		payload  string
	}
	observed := make(chan observedRequest, 3)
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		mediaType, parameters, err := mime.ParseMediaType(request.Header.Get("Content-Type"))
		if err != nil {
			t.Errorf("parse Content-Type: %v", err)
		}
		if mediaType != "multipart/form-data" {
			t.Errorf("media type = %q", mediaType)
		}
		rawBody, err := io.ReadAll(request.Body)
		if err != nil {
			t.Errorf("read multipart body: %v", err)
		}
		if int64(len(rawBody)) != request.ContentLength {
			t.Errorf("body bytes = %d, Content-Length = %d", len(rawBody), request.ContentLength)
		}
		request.Body = io.NopCloser(strings.NewReader(string(rawBody)))
		if err := request.ParseMultipartForm(1024); err != nil {
			t.Errorf("ParseMultipartForm: %v", err)
		}
		file, _, err := request.FormFile("asset")
		if err != nil {
			t.Errorf("FormFile(asset): %v", err)
		}
		payload := ""
		if file != nil {
			value, readErr := io.ReadAll(file)
			if readErr != nil {
				t.Errorf("read file: %v", readErr)
			}
			payload = string(value)
			_ = file.Close()
		}
		observed <- observedRequest{
			boundary: parameters["boundary"],
			length:   request.ContentLength,
			bodySize: len(rawBody),
			values:   request.MultipartForm.Value["note"],
			payload:  payload,
		}
		response.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	run := func(payload string) {
		request := HttpRequestNew(httpTestString(server.URL), true, nil, 1)
		HttpRequestAddParameter(request, httpTestString("note"), httpTestString("first"), httpTestString("note"), httpTestString("first"))
		HttpRequestAddParameter(request, httpTestString("note"), httpTestString("second"), httpTestString("note"), httpTestString("second"))
		offset := 0
		HttpRequestSetMultipartUpload(
			request,
			httpTestString("asset"),
			httpTestString("demo.txt"),
			httpTestString("text/plain"),
			len(payload),
			func(limit int) *ByteView {
				if offset >= len(payload) {
					return nil
				}
				count := len(payload) - offset
				if count > limit {
					count = limit
				}
				chunk := &ByteView{raw: []byte(payload[offset : offset+count])}
				offset += count
				return chunk
			},
		)
		result := HttpRequestExecute(request)
		if err := HttpResponseError(result); err != nil {
			t.Fatalf("HttpResponseError = %q, want nil", *err)
		}
	}

	run("")
	run("x")
	run("prefix\r\n--hxrt-go-boundary\r\nsuffix")
	first := <-observed
	second := <-observed
	third := <-observed
	if first.boundary == "" || second.boundary == "" || third.boundary == "" ||
		first.boundary == second.boundary || first.boundary == third.boundary || second.boundary == third.boundary {
		t.Fatalf("boundaries = %q, %q, and %q, want distinct non-empty values", first.boundary, second.boundary, third.boundary)
	}
	for _, item := range []observedRequest{first, second, third} {
		if item.length != int64(item.bodySize) {
			t.Errorf("declared length = %d, body size = %d", item.length, item.bodySize)
		}
		if strings.Join(item.values, ",") != "first,second" {
			t.Errorf("multipart note values = %v", item.values)
		}
	}
	if second.payload != "x" {
		t.Errorf("one-byte multipart payload = %q", second.payload)
	}
	if third.payload != "prefix\r\n--hxrt-go-boundary\r\nsuffix" {
		t.Errorf("multipart payload = %q", third.payload)
	}
}

func TestHttpMultipartRejectsConflictingContentTypeAndHostileMetadataBeforeNetworkIO(t *testing.T) {
	var requests atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		requests.Add(1)
		response.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	cases := []struct {
		name      string
		parameter string
		filename  string
		mediaType string
		header    string
	}{
		{
			name:      "conflicting boundary",
			parameter: "asset",
			filename:  "demo.txt",
			mediaType: "text/plain",
			header:    "multipart/form-data; boundary=caller-boundary",
		},
		{name: "parameter newline", parameter: "asset\r\nX-Evil: yes", filename: "demo.txt", mediaType: "text/plain"},
		{name: "filename nul", parameter: "asset", filename: "demo\x00.txt", mediaType: "text/plain"},
		{name: "invalid media type", parameter: "asset", filename: "demo.txt", mediaType: "text/plain\r\nX-Evil: yes"},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			request := HttpRequestNew(httpTestString(server.URL), true, nil, 1)
			if testCase.header != "" {
				HttpRequestAddHeader(request, httpTestString("Content-Type"), httpTestString(testCase.header))
			}
			HttpRequestSetMultipartUpload(
				request,
				httpTestString(testCase.parameter),
				httpTestString(testCase.filename),
				httpTestString(testCase.mediaType),
				0,
				func(_ int) *ByteView { return nil },
			)
			result := HttpRequestExecute(request)
			if err := HttpResponseError(result); err == nil {
				t.Fatal("HttpResponseError = nil, want request validation failure")
			}
		})
	}
	if got := requests.Load(); got != 0 {
		t.Fatalf("server received %d requests, want validation before network I/O", got)
	}
}

func TestHttpRequestRejectsUnsupportedSpecialHeadersBeforeNetworkIO(t *testing.T) {
	var requests atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		requests.Add(1)
		response.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	for _, name := range []string{"Transfer-Encoding", "Trailer", "Upgrade"} {
		t.Run(name, func(t *testing.T) {
			request := HttpRequestNew(httpTestString(server.URL), false, nil, 1)
			HttpRequestAddHeader(request, httpTestString(name), httpTestString("unsupported"))
			result := HttpRequestExecute(request)
			if err := HttpResponseError(result); err == nil {
				t.Fatal("HttpResponseError = nil, want special-header validation error")
			}
		})
	}
	t.Run("unsupported connection token", func(t *testing.T) {
		request := HttpRequestNew(httpTestString(server.URL), false, nil, 1)
		HttpRequestAddHeader(request, httpTestString("Connection"), httpTestString("keep-alive"))
		result := HttpRequestExecute(request)
		if err := HttpResponseError(result); err == nil {
			t.Fatal("HttpResponseError = nil, want Connection validation error")
		}
	})
	t.Run("mismatched content length", func(t *testing.T) {
		request := HttpRequestNew(httpTestString(server.URL), false, nil, 1)
		HttpRequestAddHeader(request, httpTestString("Content-Length"), httpTestString("7"))
		result := HttpRequestExecute(request)
		if err := HttpResponseError(result); err == nil {
			t.Fatal("HttpResponseError = nil, want Content-Length validation error")
		}
	})
	if got := requests.Load(); got != 0 {
		t.Fatalf("server received %d requests, want validation before network I/O", got)
	}
}

func TestHttpResponseRetainsStatusAndHeadersWhenBodyReadFails(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Length", "20")
		response.Header().Set("X-Started", "yes")
		response.WriteHeader(http.StatusAccepted)
		_, _ = response.Write([]byte("short"))
	}))
	defer server.Close()

	response := HttpRequestExecute(HttpRequestNew(httpTestString(server.URL), false, nil, 1))
	if err := HttpResponseError(response); err == nil {
		t.Fatal("HttpResponseError = nil, want truncated-body failure")
	}
	if got := HttpResponseStatus(response); got != http.StatusAccepted {
		t.Fatalf("HttpResponseStatus = %d, want %d so staged callbacks retain their order", got, http.StatusAccepted)
	}
	if HttpResponseHeaderCount(response) == 0 {
		t.Fatal("response headers were discarded after the body read failure")
	}
}

func TestHttpRequestTimeoutAndIdleConnectionCleanupAreBounded(t *testing.T) {
	t.Run("timeout", func(t *testing.T) {
		cancelled := make(chan struct{})
		server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
			select {
			case <-request.Context().Done():
				close(cancelled)
			case <-time.After(time.Second):
				t.Error("server did not observe request cancellation")
			}
		}))
		defer server.Close()

		started := time.Now()
		response := HttpRequestExecute(HttpRequestNew(httpTestString(server.URL), false, nil, 0.02))
		if err := HttpResponseError(response); err == nil {
			t.Fatal("HttpResponseError = nil, want a bounded client timeout")
		}
		if elapsed := time.Since(started); elapsed > time.Second {
			t.Fatalf("timed request took %s, want a bounded failure", elapsed)
		}
		select {
		case <-cancelled:
		case <-time.After(time.Second):
			t.Fatal("client timeout did not cancel the server request context")
		}
	})

	t.Run("idle connection", func(t *testing.T) {
		closed := make(chan struct{}, 1)
		server := httptest.NewUnstartedServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
			_, _ = response.Write([]byte("done"))
		}))
		server.Config.ConnState = func(connection net.Conn, state http.ConnState) {
			if state == http.StateClosed {
				select {
				case closed <- struct{}{}:
				default:
				}
			}
		}
		server.Start()
		defer server.Close()

		response := HttpRequestExecute(HttpRequestNew(httpTestString(server.URL), false, nil, 1))
		if err := HttpResponseError(response); err != nil {
			t.Fatalf("HttpResponseError = %q, want nil", *err)
		}
		select {
		case <-closed:
		case <-time.After(time.Second):
			t.Fatal("one-use HTTP transport retained an idle connection")
		}
	})
}

func TestHttpRequestCustomMethodBodyProxyAndSocketLifecycle(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPatch {
			t.Errorf("method = %q, want PATCH", request.Method)
		}
		body, err := io.ReadAll(request.Body)
		if err != nil {
			t.Errorf("read request body: %v", err)
		}
		if got := string(body); got != "payload" {
			t.Errorf("body = %q, want payload", got)
		}
		_, _ = response.Write([]byte("ok"))
	}))
	defer server.Close()

	connection, err := net.Dial("tcp", server.Listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	socket := SocketNewTCP()
	socket.installConn(connection)
	request := HttpRequestNew(httpTestString(server.URL), true, httpTestString("patch"), 1)
	HttpRequestSetBodyString(request, httpTestString("payload"))
	HttpRequestSetSocket(request, socket)
	response := HttpRequestExecute(request)
	if err := HttpResponseError(response); err != nil {
		t.Fatalf("HttpResponseError = %q, want nil", *err)
	}
	if socket.snapshotConn() != nil {
		t.Fatal("custom-request socket remained attached after synchronous execution")
	}

	if got := *HttpProxyDescriptor(httpTestString("proxy.local"), 3128, httpTestString("user"), httpTestString("pass")); got != "http://user:pass@proxy.local:3128" {
		t.Fatalf("HttpProxyDescriptor = %q", got)
	}
}

func TestHttpExchangePublishesHeadersAndFirstChunkBeforeBodyCompletion(t *testing.T) {
	firstChunkSent := make(chan struct{})
	releaseBody := make(chan struct{}, 1)
	defer func() {
		select {
		case releaseBody <- struct{}{}:
		default:
		}
	}()
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Length", "11")
		response.Header().Set("X-Stream", "ready")
		response.WriteHeader(http.StatusCreated)
		response.(http.Flusher).Flush()
		_, _ = response.Write([]byte("hello"))
		response.(http.Flusher).Flush()
		close(firstChunkSent)
		<-releaseBody
		_, _ = response.Write([]byte(" world"))
	}))
	defer server.Close()

	exchangeReady := make(chan *HttpExchange, 1)
	go func() {
		request := HttpRequestNew(httpTestString(server.URL), false, nil, 1)
		exchangeReady <- HttpRequestStartExchange(request)
	}()

	var exchange *HttpExchange
	select {
	case exchange = <-exchangeReady:
	case <-time.After(time.Second):
		t.Fatal("exchange did not publish response headers before body completion")
	}
	if err := HttpExchangeError(exchange); err != nil {
		t.Fatalf("HttpExchangeError = %q, want nil", *err)
	}
	if got := HttpExchangeStatus(exchange); got != http.StatusCreated {
		t.Fatalf("HttpExchangeStatus = %d, want %d", got, http.StatusCreated)
	}
	if got := HttpExchangeContentLength(exchange); got != 11 {
		t.Fatalf("HttpExchangeContentLength = %d, want 11", got)
	}
	select {
	case <-firstChunkSent:
	case <-time.After(time.Second):
		t.Fatal("server did not flush the first body chunk")
	}

	first := HttpExchangeReadResponseChunk(exchange, 5)
	if got := string(byteViewRaw(HttpReadResultBody(first))); got != "hello" {
		t.Fatalf("first streamed body chunk = %q, want hello", got)
	}
	if err := HttpReadResultError(first); err != nil {
		t.Fatalf("first chunk error = %q, want nil", *err)
	}
	if HttpReadResultEOF(first) {
		t.Fatal("first chunk reported EOF before the server released the tail")
	}

	releaseBody <- struct{}{}
	var tail strings.Builder
	for {
		result := HttpExchangeReadResponseChunk(exchange, 3)
		tail.Write(byteViewRaw(HttpReadResultBody(result)))
		if err := HttpReadResultError(result); err != nil {
			t.Fatalf("tail read error = %q, want nil", *err)
		}
		if HttpReadResultEOF(result) {
			break
		}
	}
	if got := tail.String(); got != " world" {
		t.Fatalf("streamed body tail = %q, want world", got)
	}
	HttpExchangeClose(exchange)
	HttpExchangeClose(exchange)
}

func TestHttpExchangeReadPreservesBytesAlongsideTerminalFailure(t *testing.T) {
	body := &httpBytesThenErrorBody{}
	exchange := &HttpExchange{
		response: &http.Response{
			Body:          body,
			ContentLength: 8,
		},
	}

	result := HttpExchangeReadResponseChunk(exchange, 4)
	if got := string(byteViewRaw(HttpReadResultBody(result))); got != "part" {
		t.Fatalf("partial body = %q, want part", got)
	}
	if err := HttpReadResultError(result); err == nil {
		t.Fatal("terminal read error = nil, want unexpected EOF")
	}
	if HttpReadResultEOF(result) {
		t.Fatal("unexpected EOF was misclassified as successful completion")
	}
	HttpExchangeCancel(exchange)
	if got := body.closed.Load(); got != 1 {
		t.Fatalf("response body close count = %d, want 1", got)
	}
}

func TestHttpExchangeReadRejectsInvalidNativeByteCounts(t *testing.T) {
	for _, count := range []int{-1, 5} {
		t.Run(strconv.Itoa(count), func(t *testing.T) {
			exchange := &HttpExchange{
				response: &http.Response{Body: &httpInvalidCountBody{count: count}},
			}
			result := HttpExchangeReadResponseChunk(exchange, 4)
			if err := HttpReadResultError(result); err == nil {
				t.Fatalf("invalid native byte count %d produced no error", count)
			}
			if got := len(byteViewRaw(HttpReadResultBody(result))); got != 0 {
				t.Fatalf("invalid native byte count %d exposed %d bytes, want 0", count, got)
			}
			HttpExchangeCancel(exchange)
		})
	}
}

func TestHttpExchangeContentLengthProtectsHaxeIntRange(t *testing.T) {
	unknown := &HttpExchange{response: &http.Response{ContentLength: -1}}
	if got := HttpExchangeContentLength(unknown); got != -1 {
		t.Fatalf("unknown content length = %d, want -1", got)
	}
	oversized := &HttpExchange{response: &http.Response{ContentLength: int64(1 << 31)}}
	if got := HttpExchangeContentLength(oversized); got != -2 {
		t.Fatalf("oversized content length = %d, want -2", got)
	}
}

func TestHttpExchangeCancelUnblocksReadAndIsIdempotent(t *testing.T) {
	body := &httpBlockingBody{closed: make(chan struct{}), started: make(chan struct{})}
	exchange := &HttpExchange{
		response: &http.Response{Body: body},
	}
	readDone := make(chan *HttpReadResult, 1)
	go func() {
		readDone <- HttpExchangeReadResponseChunk(exchange, 32)
	}()

	select {
	case <-body.started:
	case <-time.After(time.Second):
		t.Fatal("response-body read did not start")
	}
	HttpExchangeCancel(exchange)
	HttpExchangeCancel(exchange)
	select {
	case result := <-readDone:
		if err := HttpReadResultError(result); err == nil {
			t.Fatal("canceled read error = nil, want cancellation signal")
		}
	case <-time.After(time.Second):
		t.Fatal("cancel did not unblock the response-body read")
	}
}

func TestHttpExchangeRetainsOnlyBoundedReadChunks(t *testing.T) {
	const bodySize = 5 * 1024 * 1024
	body := &httpBoundedBody{remaining: bodySize}
	exchange := &HttpExchange{
		response: &http.Response{
			Body:          body,
			ContentLength: bodySize,
		},
	}

	total := 0
	for {
		result := HttpExchangeReadResponseChunk(exchange, 1024)
		chunk := byteViewRaw(HttpReadResultBody(result))
		if len(chunk) > 1024 {
			t.Fatalf("chunk length = %d, want at most 1024", len(chunk))
		}
		total += len(chunk)
		if err := HttpReadResultError(result); err != nil {
			t.Fatalf("bounded read error = %q, want nil", *err)
		}
		if HttpReadResultEOF(result) {
			break
		}
	}
	if total != bodySize {
		t.Fatalf("streamed bytes = %d, want %d", total, bodySize)
	}
	if body.maxRead > 1024 {
		t.Fatalf("largest native read buffer = %d, want at most 1024", body.maxRead)
	}
	HttpExchangeClose(exchange)
	if got := body.closed.Load(); got != 1 {
		t.Fatalf("response body close count = %d, want 1", got)
	}
}
