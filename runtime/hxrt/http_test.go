package hxrt

import (
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sort"
	"strings"
	"testing"
	"time"
)

func httpTestString(value string) *string {
	return &value
}

func TestHttpRequestTypedGetAndResponseBoundary(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if got := request.URL.Query().Get("q"); got != "last" {
			t.Errorf("query q = %q, want last", got)
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
	HttpRequestAddParameter(request, httpTestString("q"), httpTestString("first"))
	HttpRequestAddParameter(request, httpTestString("q"), httpTestString("last"))
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
		if got := form.Get("base"); got != "from-url" {
			t.Errorf("form base = %q, want from-url", got)
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
	HttpRequestAddParameter(request, httpTestString("field"), httpTestString("value"))
	response := HttpRequestExecute(request)
	if err := HttpResponseError(response); err != nil {
		t.Fatalf("HttpResponseError = %q, want nil", *err)
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
		server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
			time.Sleep(100 * time.Millisecond)
			_, _ = response.Write([]byte("late"))
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
