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
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("proxy0="), sys__Http_hxrt_proxyDescriptor()))
	hx_obj_1 := map[string]any{}
	hx_obj_1["host"] = hxrt.StringFromLiteral("proxy.local")
	hx_obj_1["port"] = 3128
	hx_obj_2 := map[string]any{}
	hx_obj_2["user"] = hxrt.StringFromLiteral("scott")
	hx_obj_2["pass"] = hxrt.StringFromLiteral("tiger")
	hx_obj_1["auth"] = hx_obj_2
	sys__Http_PROXY = hx_obj_1
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("proxy1="), sys__Http_hxrt_proxyDescriptor()))
	hx_obj_3 := map[string]any{}
	hx_obj_3["host"] = hxrt.StringFromLiteral("proxy.local")
	hx_obj_3["port"] = 80
	hx_obj_4 := map[string]any{}
	hx_obj_4["user"] = hxrt.StringFromLiteral("scott")
	hx_obj_4["pass"] = nil
	hx_obj_3["auth"] = hx_obj_4
	sys__Http_PROXY = hx_obj_3
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("proxy2="), sys__Http_hxrt_proxyDescriptor()))
	hx_obj_5 := map[string]any{}
	hx_obj_5["host"] = hxrt.StringFromLiteral("proxy.local:9000")
	hx_obj_5["port"] = 3128
	hx_obj_5["auth"] = nil
	sys__Http_PROXY = hx_obj_5
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("proxy3="), sys__Http_hxrt_proxyDescriptor()))
	hx_obj_6 := map[string]any{}
	hx_obj_6["host"] = nil
	hx_obj_6["port"] = 3128
	hx_obj_6["auth"] = nil
	sys__Http_PROXY = hx_obj_6
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("proxy4="), sys__Http_hxrt_proxyDescriptor()))
	sys__Http_PROXY = nil
	http := New_sys__Http(hxrt.StringFromLiteral("data:text/plain,body"))
	_ = http
	sink := New_haxe__io__BytesBuffer()
	_ = sink
	http.customRequest(false, sink, func() map[string]any {
		hx_obj_7 := map[string]any{}
		hx_obj_7["marker"] = hxrt.StringFromLiteral("sock")
		return hx_obj_7
	}(), hxrt.StringFromLiteral("PATCH"))
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("methodSock="), sink.getBytes().toString()))
}

type haxe__io__Encoding struct {
}

type haxe__io__Input struct {
}

type haxe__io__Output struct {
}

type haxe__io__Bytes struct {
	b             []int
	length        int
	__hx_raw      []byte
	__hx_rawValid bool
}

type haxe__io__BytesBuffer struct {
	b []int
}

func New_haxe__io__Input() *haxe__io__Input {
	return &haxe__io__Input{}
}

func New_haxe__io__Output() *haxe__io__Output {
	return &haxe__io__Output{}
}

func New_haxe__io__Bytes(length int, b []int) *haxe__io__Bytes {
	if b == nil {
		b = make([]int, length)
	}
	return &haxe__io__Bytes{b: b, length: len(b)}
}

func haxe__io__Bytes_alloc(length int) *haxe__io__Bytes {
	return &haxe__io__Bytes{b: make([]int, length), length: length}
}

func haxe__io__Bytes_ofString(value *string, encoding ...*haxe__io__Encoding) *haxe__io__Bytes {
	raw := []byte(*hxrt.StdString(value))
	converted := make([]int, len(raw))
	for i := 0; i < len(raw); i++ {
		converted[i] = int(raw[i])
	}
	return &haxe__io__Bytes{b: converted, length: len(converted), __hx_raw: raw, __hx_rawValid: true}
}

func (self *haxe__io__Bytes) toString() *string {
	if self == nil {
		return hxrt.StringFromLiteral("")
	}
	return hxrt.BytesToString(self.b)
}

func (self *haxe__io__Bytes) get(pos int) int {
	return self.b[pos]
}

func (self *haxe__io__Bytes) set(pos int, value int) {
	self.b[pos] = value
	self.__hx_rawValid = false
}

func New_haxe__io__BytesBuffer() *haxe__io__BytesBuffer {
	return &haxe__io__BytesBuffer{b: []int{}}
}

func (self *haxe__io__BytesBuffer) addByte(value int) {
	self.b = append(self.b, (value & 255))
}

func (self *haxe__io__BytesBuffer) add(src *haxe__io__Bytes) {
	if src == nil {
		return
	}
	self.b = append(self.b, src.b...)
}

func (self *haxe__io__BytesBuffer) addString(value *string, encoding ...*haxe__io__Encoding) {
	self.add(haxe__io__Bytes_ofString(value))
}

func (self *haxe__io__BytesBuffer) getBytes() *haxe__io__Bytes {
	copied := hxrt.BytesClone(self.b)
	return &haxe__io__Bytes{b: copied, length: len(copied)}
}

func (self *haxe__io__BytesBuffer) get_length() int {
	return len(self.b)
}

type haxe__ds__IntMap struct {
	h map[int]any
}

type haxe__ds__StringMap struct {
	h map[string]any
}

type haxe__ds__ObjectMap struct {
	h map[any]any
}

type haxe__ds__EnumValueMap struct {
	h map[any]any
}

type haxe__ds__List struct {
	items  []any
	length int
}

func New_haxe__ds__IntMap() *haxe__ds__IntMap {
	return &haxe__ds__IntMap{h: map[int]any{}}
}

func (self *haxe__ds__IntMap) set(key int, value any) {
	self.h[key] = value
}

func (self *haxe__ds__IntMap) get(key int) any {
	value := self.h[key]
	return value
}

func (self *haxe__ds__IntMap) exists(key int) bool {
	_, ok := self.h[key]
	return ok
}

func (self *haxe__ds__IntMap) remove(key int) bool {
	_, ok := self.h[key]
	delete(self.h, key)
	return ok
}

func New_haxe__ds__StringMap() *haxe__ds__StringMap {
	return &haxe__ds__StringMap{h: map[string]any{}}
}

func (self *haxe__ds__StringMap) set(key *string, value any) {
	self.h[*hxrt.StdString(key)] = value
}

func (self *haxe__ds__StringMap) get(key *string) any {
	value := self.h[*hxrt.StdString(key)]
	return value
}

func (self *haxe__ds__StringMap) exists(key *string) bool {
	_, ok := self.h[*hxrt.StdString(key)]
	return ok
}

func (self *haxe__ds__StringMap) remove(key *string) bool {
	_, ok := self.h[*hxrt.StdString(key)]
	delete(self.h, *hxrt.StdString(key))
	return ok
}

func New_haxe__ds__ObjectMap() *haxe__ds__ObjectMap {
	return &haxe__ds__ObjectMap{h: map[any]any{}}
}

func (self *haxe__ds__ObjectMap) set(key any, value any) {
	self.h[key] = value
}

func (self *haxe__ds__ObjectMap) get(key any) any {
	return self.h[key]
}

func (self *haxe__ds__ObjectMap) exists(key any) bool {
	_, ok := self.h[key]
	return ok
}

func (self *haxe__ds__ObjectMap) remove(key any) bool {
	_, ok := self.h[key]
	delete(self.h, key)
	return ok
}

func New_haxe__ds__EnumValueMap() *haxe__ds__EnumValueMap {
	return &haxe__ds__EnumValueMap{h: map[any]any{}}
}

func (self *haxe__ds__EnumValueMap) set(key any, value any) {
	self.h[key] = value
}

func (self *haxe__ds__EnumValueMap) get(key any) any {
	return self.h[key]
}

func (self *haxe__ds__EnumValueMap) exists(key any) bool {
	_, ok := self.h[key]
	return ok
}

func (self *haxe__ds__EnumValueMap) remove(key any) bool {
	_, ok := self.h[key]
	delete(self.h, key)
	return ok
}

func New_haxe__ds__List() *haxe__ds__List {
	return &haxe__ds__List{items: []any{}, length: 0}
}

func (self *haxe__ds__List) add(item any) {
	self.items = append(self.items, item)
	self.length = len(self.items)
}

func (self *haxe__ds__List) push(item any) {
	self.items = append([]any{item}, self.items...)
	self.length = len(self.items)
}

func (self *haxe__ds__List) pop() any {
	if len(self.items) == 0 {
		return nil
	}
	head := self.items[0]
	self.items = self.items[1:]
	self.length = len(self.items)
	return head
}

func (self *haxe__ds__List) first() any {
	if len(self.items) == 0 {
		return nil
	}
	return self.items[0]
}

func (self *haxe__ds__List) last() any {
	size := len(self.items)
	if size == 0 {
		return nil
	}
	return self.items[(size - 1)]
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

func (self *sys__Http) getResponseHeaderValues(key *string) []*string {
	if self == nil {
		return nil
	}
	rawKey := *hxrt.StdString(key)
	normalized := strings.ToLower(rawKey)
	if self.responseHeadersSameKey != nil {
		if values, ok := self.responseHeadersSameKey[rawKey]; ok {
			return values
		}
		if values, ok := self.responseHeadersSameKey[normalized]; ok {
			return values
		}
	}
	if self.responseHeaders == nil {
		return nil
	}
	single := self.responseHeaders.get(hxrt.StringFromLiteral(rawKey))
	if single == nil && rawKey != normalized {
		single = self.responseHeaders.get(hxrt.StringFromLiteral(normalized))
	}
	if single == nil {
		return nil
	}
	return []*string{hxrt.StdString(single)}
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
		self.responseBytes = &haxe__io__Bytes{b: intPayload, length: len(intPayload)}
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
	self.responseBytes = &haxe__io__Bytes{b: intPayload, length: len(intPayload)}
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
	if api == nil || payload == nil {
		return
	}
	switch out := api.(type) {
	case *haxe__io__BytesBuffer:
		out.add(payload)
	case interface{ add(*haxe__io__Bytes) }:
		out.add(payload)
	case interface {
		writeBytes(*haxe__io__Bytes, int, int) int
	}:
		out.writeBytes(payload, 0, payload.length)
	case interface {
		writeFullBytes(*haxe__io__Bytes, int, int)
	}:
		out.writeFullBytes(payload, 0, payload.length)
	case interface{ writeString(*string) }:
		out.writeString(payload.toString())
	}
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
