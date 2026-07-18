package main

import (
	"bytes"
	"io"
	"net"
	"net/http"
	"net/url"
	"snapshot/hxrt"
	"strings"
	"time"
)

func main() {
	http := New_sys__Http(hxrt.StringFromLiteral("data:text/plain,hello%20leaf"))
	sink := New_haxe__io__BytesBuffer()
	http.customRequest(false, sink)
	values := http.getResponseHeaderValues(hxrt.StringFromLiteral("Content-Type"))
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("headers="), func() *string {
		var hx_if_12 *string
		if (values != nil) && (values.Len() > 0) {
			hx_if_12 = hxrt.StdString(func(hx_value_10 any) *string {
				if hx_value_10 == nil {
					var hx_zero_11 *string
					return hx_zero_11
				}
				return hx_value_10.(*string)
			}(values.Get(0)))
		} else {
			hx_if_12 = hxrt.StringFromLiteral("null")
		}
		return hx_if_12
	}()))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("direct="), sys__Http_requestUrl(hxrt.StringFromLiteral("data:text/plain,direct%20ok"))))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("proxy0="), sys__Http_hxrt_proxyDescriptor()))
	hxrt.Println(v_2)
	sys__Http_PROXY = func() map[string]any {
		hx_obj_13 := map[string]any{}
		hx_obj_13["host"] = hxrt.StringFromLiteral("proxy.local")
		hx_obj_13["port"] = 3128
		hx_obj_14 := map[string]any{}
		hx_obj_14["user"] = hxrt.StringFromLiteral("scott")
		hx_obj_14["pass"] = hxrt.StringFromLiteral("tiger")
		hx_obj_13["auth"] = hx_obj_14
		return hx_obj_13
	}()
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("proxy1="), sys__Http_hxrt_proxyDescriptor()))
	hxrt.Println(v_3)
}

type hxrt__http__Pair struct {
	name  *string
	value *string
}

type hxrt__http__FileUpload struct {
	param    *string
	filename *string
	size     int
	mimeType *string
	fileRef  any
}

var sys__Http_PROXY any = nil

type sys__Http struct {
	url                    *string
	responseAsString       *string
	responseBytes          *haxe__io__Bytes
	postData               *string
	postBytes              *haxe__io__Bytes
	headers                []hxrt__http__Pair
	params                 []hxrt__http__Pair
	onData                 func(*string)
	onBytes                func(*haxe__io__Bytes)
	onError                func(*string)
	onStatus               func(int)
	noShutdown             bool
	cnxTimeout             float64
	responseHeaders        *haxe__ds__StringMap
	responseHeadersSameKey map[string][]*string
	fileUpload             *hxrt__http__FileUpload
}

func New_sys__Http(url *string) *sys__Http {
	self := &sys__Http{url: url, headers: []hxrt__http__Pair{}, params: []hxrt__http__Pair{}, cnxTimeout: 10, responseHeaders: New_haxe__ds__StringMap(), responseHeadersSameKey: map[string][]*string{}}
	self.onData = func(data *string) {}
	self.onBytes = func(data *haxe__io__Bytes) {}
	self.onError = func(msg *string) {}
	self.onStatus = func(status int) {}
	return self
}

func (self *sys__Http) setHeader(name *string, value *string) {
	if self == nil {
		return
	}
	for i := 0; i < len(self.headers); i++ {
		if *hxrt.StdString(self.headers[i].name) == *hxrt.StdString(name) {
			self.headers[i] = hxrt__http__Pair{name: name, value: value}
			return
		}
	}
	self.headers = append(self.headers, hxrt__http__Pair{name: name, value: value})
}

func (self *sys__Http) addHeader(header *string, value *string) {
	if self == nil {
		return
	}
	self.headers = append(self.headers, hxrt__http__Pair{name: header, value: value})
}

func (self *sys__Http) setParameter(name *string, value *string) {
	if self == nil {
		return
	}
	for i := 0; i < len(self.params); i++ {
		if *hxrt.StdString(self.params[i].name) == *hxrt.StdString(name) {
			self.params[i] = hxrt__http__Pair{name: name, value: value}
			return
		}
	}
	self.params = append(self.params, hxrt__http__Pair{name: name, value: value})
}

func (self *sys__Http) addParameter(name *string, value *string) {
	if self == nil {
		return
	}
	self.params = append(self.params, hxrt__http__Pair{name: name, value: value})
}

func (self *sys__Http) setPostData(data *string) {
	if self == nil {
		return
	}
	self.postData = data
	self.postBytes = nil
}

func (self *sys__Http) setPostBytes(data *haxe__io__Bytes) {
	if self == nil {
		return
	}
	self.postBytes = data
	self.postData = nil
}

func (self *sys__Http) fileTransfer(argname *string, filename *string, file any, size int, mimeType ...*string) {
	if self == nil {
		return
	}
	resolvedMime := hxrt.StringFromLiteral("application/octet-stream")
	if len(mimeType) > 0 {
		if mimeType[0] != nil {
			resolvedMime = mimeType[0]
		}
	}
	self.fileUpload = &hxrt__http__FileUpload{param: argname, filename: filename, size: size, mimeType: resolvedMime, fileRef: file}
}

func (self *sys__Http) fileTransfert(argname *string, filename *string, file any, size int, mimeType ...*string) {
	self.fileTransfer(argname, filename, file, size, mimeType...)
}

func (self *sys__Http) getResponseHeaderValues(key *string) *hxrt.Array {
	return sys__GoHttpHelpers_getResponseHeaderValues(self, key)
}

func (self *sys__Http) get_responseData() *string {
	if self == nil {
		return hxrt.StringFromLiteral("")
	}
	if self.responseAsString == nil && self.responseBytes != nil {
		self.responseAsString = self.responseBytes.toString()
	}
	return self.responseAsString
}

func (self *sys__Http) customRequest(post bool, api any, rest ...any) {
	var socketOverride any = nil
	var methodOverride *string = nil
	if len(rest) >= 1 {
		switch candidate := rest[0].(type) {
		case string:
			if len(rest) == 1 {
				methodOverride = hxrt.StringFromLiteral(candidate)
			}
		case *string:
			if len(rest) == 1 {
				methodOverride = candidate
			}
		default:
			socketOverride = candidate
		}
	}
	if len(rest) >= 2 {
		switch candidate := rest[1].(type) {
		case *string:
			methodOverride = candidate
		case string:
			methodOverride = hxrt.StringFromLiteral(candidate)
		}
	}
	self.hxrt__http__requestWith(post, methodOverride, api, socketOverride)
}

func (self *sys__Http) request(post ...bool) {
	if self == nil {
		return
	}
	isPost := false
	if len(post) > 0 {
		isPost = post[0]
	}
	if self.postData != nil || self.postBytes != nil || self.fileUpload != nil {
		isPost = true
	}
	self.hxrt__http__requestWith(isPost, nil, nil, nil)
}

func (self *sys__Http) hxrt__http__requestWith(post bool, methodOverride *string, api any, sock any) {
	self.responseAsString = nil
	self.responseBytes = nil
	self.responseHeaders = New_haxe__ds__StringMap()
	self.responseHeadersSameKey = map[string][]*string{}
	rawUrl := *hxrt.StdString(self.url)
	parsedURL, err := url.Parse(rawUrl)
	if err != nil || parsedURL == nil {
		if self.onError != nil {
			self.onError(hxrt.StringFromLiteral("Invalid URL"))
		}
		return
	}
	query := parsedURL.Query()
	for _, param := range self.params {
		query.Set(*hxrt.StdString(param.name), *hxrt.StdString(param.value))
	}
	var bodyReader io.Reader = nil
	var contentTypeOverride *string = nil
	if post {
		if self.fileUpload != nil {
			multipartPayload := ""
			for _, param := range self.params {
				multipartPayload += "--hxrt-go-boundary\r\n"
				multipartPayload += "Content-Disposition: form-data; name=\"" + *hxrt.StdString(param.name) + "\"\r\n\r\n"
				multipartPayload += *hxrt.StdString(param.value) + "\r\n"
			}
			multipartPayload += "--hxrt-go-boundary\r\n"
			multipartPayload += "Content-Disposition: form-data; name=\"" + *hxrt.StdString(self.fileUpload.param) + "\"; filename=\"" + *hxrt.StdString(self.fileUpload.filename) + "\"\r\n"
			multipartPayload += "Content-Type: " + *hxrt.StdString(self.fileUpload.mimeType) + "\r\n\r\n"
			multipartPayload += "[uploaded-bytes=" + *hxrt.StdString(self.fileUpload.size) + "]\r\n"
			multipartPayload += "--hxrt-go-boundary--\r\n"
			bodyReader = strings.NewReader(multipartPayload)
			contentTypeOverride = hxrt.StringFromLiteral("multipart/form-data; boundary=hxrt-go-boundary")
		} else if self.postBytes != nil {
			rawBody := make([]byte, len(self.postBytes.b))
			for i := 0; i < len(self.postBytes.b); i++ {
				rawBody[i] = byte(self.postBytes.b[i])
			}
			bodyReader = bytes.NewReader(rawBody)
		} else if self.postData != nil {
			bodyReader = strings.NewReader(*hxrt.StdString(self.postData))
		} else {
			encoded := query.Encode()
			bodyReader = strings.NewReader(encoded)
			hasContentType := false
			for _, header := range self.headers {
				if strings.EqualFold(*hxrt.StdString(header.name), "Content-Type") {
					hasContentType = true
					break
				}
			}
			if !hasContentType {
				contentTypeOverride = hxrt.StringFromLiteral("application/x-www-form-urlencoded")
			}
		}
	} else {
		parsedURL.RawQuery = query.Encode()
	}
	if parsedURL.Scheme == "data" {
		payload := parsedURL.Opaque
		mediaType := "text/plain"
		commaIndex := strings.Index(payload, ",")
		if commaIndex >= 0 {
			if commaIndex > 0 {
				mediaType = payload[:commaIndex]
			}
			payload = payload[commaIndex+1:]
		}
		if post {
			if self.fileUpload != nil {
				payload = "multipart file=" + *hxrt.StdString(self.fileUpload.filename) + ";mime=" + *hxrt.StdString(self.fileUpload.mimeType) + ";size=" + *hxrt.StdString(self.fileUpload.size)
			} else if bodyReader != nil {
				rawBody, readErr := io.ReadAll(bodyReader)
				if readErr == nil {
					payload = string(rawBody)
				}
			}
		}
		decoded, decodeErr := url.QueryUnescape(payload)
		if decodeErr == nil {
			payload = decoded
		}
		if methodOverride != nil {
			methodToken := strings.ToUpper(*hxrt.StdString(methodOverride))
			if methodToken != "" && methodToken != "NULL" {
				payload = methodToken + " " + payload
			}
		}
		rawPayload := []byte(payload)
		intPayload := make([]int, len(rawPayload))
		for i := 0; i < len(rawPayload); i++ {
			intPayload[i] = int(rawPayload[i])
		}
		self.responseBytes = haxe__io__Bytes_ofData(intPayload)
		self.responseAsString = hxrt.StringFromLiteral(payload)
		self.responseHeaders = New_haxe__ds__StringMap()
		self.responseHeaders.set(hxrt.StringFromLiteral("content-type"), hxrt.StringFromLiteral(mediaType))
		self.responseHeaders.set(hxrt.StringFromLiteral("Content-Type"), hxrt.StringFromLiteral(mediaType))
		self.responseHeadersSameKey = map[string][]*string{}
		hxrt__http__captureApi(api, self.responseBytes)
		if self.onStatus != nil {
			self.onStatus(200)
		}
		if self.onData != nil {
			self.onData(self.responseAsString)
		}
		if self.onBytes != nil {
			self.onBytes(self.responseBytes)
		}
		return
	}
	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		if self.onError != nil {
			self.onError(hxrt.StringFromLiteral("Invalid URL"))
		}
		return
	}
	method := "GET"
	if post {
		method = "POST"
	}
	if methodOverride != nil {
		methodToken := strings.ToUpper(*hxrt.StdString(methodOverride))
		if methodToken != "" && methodToken != "NULL" {
			method = methodToken
		}
	}
	request, err := http.NewRequest(method, parsedURL.String(), bodyReader)
	if err != nil {
		if self.onError != nil {
			self.onError(hxrt.StringFromLiteral(err.Error()))
		}
		return
	}
	for _, header := range self.headers {
		request.Header.Set(*hxrt.StdString(header.name), *hxrt.StdString(header.value))
	}
	if contentTypeOverride != nil && request.Header.Get("Content-Type") == "" {
		request.Header.Set("Content-Type", *hxrt.StdString(contentTypeOverride))
	}
	transport := &http.Transport{}
	proxyURL := hxrt__http__proxyURL()
	if proxyURL != nil {
		transport.Proxy = http.ProxyURL(proxyURL)
	}
	var socketAdapter interface {
		hxrt__socket_conn() net.Conn
		hxrt__socket_setConn(net.Conn)
		close()
	}
	if candidate, ok := sock.(interface {
		hxrt__socket_conn() net.Conn
		hxrt__socket_setConn(net.Conn)
		close()
	}); ok {
		socketAdapter = candidate
		transport.DisableKeepAlives = true
		request.Close = true
		socketConsumed := false
		transport.Dial = func(network string, addr string) (net.Conn, error) {
			if socketConsumed {
				return nil, io.EOF
			}
			socketConsumed = true
			conn := socketAdapter.hxrt__socket_conn()
			if conn == nil {
				dialConn, dialErr := net.Dial(network, addr)
				if dialErr != nil {
					return nil, dialErr
				}
				socketAdapter.hxrt__socket_setConn(dialConn)
				conn = dialConn
			}
			return conn, nil
		}
		defer socketAdapter.close()
	}
	timeout := time.Duration(self.cnxTimeout * float64(time.Second))
	if timeout <= 0 {
		timeout = 10 * time.Second
	}
	client := &http.Client{Transport: transport, Timeout: timeout}
	response, err := client.Do(request)
	if err != nil {
		if self.onError != nil {
			self.onError(hxrt.StringFromLiteral(err.Error()))
		}
		return
	}
	defer response.Body.Close()
	self.responseHeaders = New_haxe__ds__StringMap()
	self.responseHeadersSameKey = map[string][]*string{}
	for name, values := range response.Header {
		if len(values) == 0 {
			continue
		}
		lowerKey := strings.ToLower(name)
		lastValue := hxrt.StringFromLiteral(values[len(values)-1])
		self.responseHeaders.set(hxrt.StringFromLiteral(name), lastValue)
		if lowerKey != name {
			self.responseHeaders.set(hxrt.StringFromLiteral(lowerKey), lastValue)
		}
		if len(values) > 1 {
			allValues := make([]*string, 0, len(values))
			for _, rawValue := range values {
				allValues = append(allValues, hxrt.StringFromLiteral(rawValue))
			}
			self.responseHeadersSameKey[name] = allValues
			if lowerKey != name {
				self.responseHeadersSameKey[lowerKey] = allValues
			}
		}
	}
	if self.onStatus != nil {
		self.onStatus(response.StatusCode)
	}
	rawPayload, err := io.ReadAll(response.Body)
	if err != nil {
		if self.onError != nil {
			self.onError(hxrt.StringFromLiteral(err.Error()))
		}
		return
	}
	intPayload := make([]int, len(rawPayload))
	for i := 0; i < len(rawPayload); i++ {
		intPayload[i] = int(rawPayload[i])
	}
	self.responseBytes = haxe__io__Bytes_ofData(intPayload)
	self.responseAsString = hxrt.StringFromLiteral(string(rawPayload))
	hxrt__http__captureApi(api, self.responseBytes)
	if response.StatusCode >= 400 {
		if self.onError != nil {
			self.onError(hxrt.StringConcatAny(hxrt.StringFromLiteral("Http Error #"), response.StatusCode))
		}
		return
	}
	if self.onData != nil {
		self.onData(self.responseAsString)
	}
	if self.onBytes != nil {
		self.onBytes(self.responseBytes)
	}
}

func hxrt__http__captureApi(api any, payload *haxe__io__Bytes) {
	sys__GoHttpHelpers_captureApi(api, payload)
}

func hxrt__http__proxyURL() *url.URL {
	if sys__Http_PROXY == nil {
		return nil
	}
	config, ok := sys__Http_PROXY.(map[string]any)
	if !ok {
		return nil
	}
	host := *hxrt.StdString(config["host"])
	if host == "" {
		return nil
	}
	if host == "null" {
		return nil
	}
	port := *hxrt.StdString(config["port"])
	hostPort := host
	if port != "" && port != "null" && !strings.Contains(hostPort, ":") {
		hostPort = hostPort + ":" + port
	}
	proxyURL, err := url.Parse("http://" + hostPort)
	if err != nil {
		return nil
	}
	if authValue, ok := config["auth"]; ok {
		if authMap, ok := authValue.(map[string]any); ok {
			user := *hxrt.StdString(authMap["user"])
			pass := *hxrt.StdString(authMap["pass"])
			if user != "" && user != "null" {
				if pass == "null" {
					pass = ""
				}
				proxyURL.User = url.UserPassword(user, pass)
			}
		}
	}
	return proxyURL
}

func sys__Http_hxrt_proxyDescriptor() *string {
	proxyURL := hxrt__http__proxyURL()
	if proxyURL == nil {
		return hxrt.StringFromLiteral("null")
	}
	return hxrt.StringFromLiteral(proxyURL.String())
}

func sys__Http_requestUrl(url *string) *string {
	self := New_sys__Http(url)
	result := hxrt.StringFromLiteral("")
	self.onData = func(data *string) { result = data }
	self.onError = func(msg *string) { result = msg }
	self.request()
	return result
}
