package main

import (
	"bytes"
	"io"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"snapshot/hxrt"
	"strings"
	"time"
)

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

func main() {
	http := New_sys__Http(hxrt.StringFromLiteral("data:text/plain,hello%20from%20haxe.go"))
	dataLog := hxrt.StringFromLiteral("")
	statusLog := -1
	byteCount := -1
	errLog := hxrt.StringFromLiteral("")
	http.onData = func(data *string) {
		dataLog = data
	}
	http.onStatus = func(status int) {
		statusLog = status
	}
	http.onBytes = func(bytes *haxe__io__Bytes) {
		byteCount = bytes.length
	}
	http.onError = func(msg *string) {
		errLog = msg
	}
	http.setHeader(hxrt.StringFromLiteral("X-Test"), hxrt.StringFromLiteral("one"))
	http.setHeader(hxrt.StringFromLiteral("X-Test"), hxrt.StringFromLiteral("two"))
	http.addHeader(hxrt.StringFromLiteral("X-Trace"), hxrt.StringFromLiteral("ok"))
	http.setParameter(hxrt.StringFromLiteral("a"), hxrt.StringFromLiteral("1"))
	http.addParameter(hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("2"))
	http.request()
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("data="), dataLog)))
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("bytes="), byteCount)))
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("status="), statusLog)))
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("response="), http.get_responseData()))
	hxrt.Println(v)
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("error="), errLog)))
	post := New_sys__Http(hxrt.StringFromLiteral("data:text/plain,ignored"))
	post.setPostData(hxrt.StringFromLiteral("post-body"))
	postData := hxrt.StringFromLiteral("")
	post.onData = func(data *string) {
		postData = data
	}
	post.request(true)
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("post="), postData)))
	bad := New_sys__Http(hxrt.StringFromLiteral("://bad"))
	badErr := hxrt.StringFromLiteral("")
	bad.onError = func(msg *string) {
		badErr = msg
	}
	bad.request()
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("bad="), badErr)))
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("direct="), sys__Http_requestUrl(hxrt.StringFromLiteral("data:text/plain,direct%20ok"))))
	hxrt.Println(v_1)
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

type Std struct {
}

type Type struct {
}

type Reflect struct {
}

func Reflect_compare(a any, b any) int {
	toFloat := func(value any) (float64, bool) {
		switch v := value.(type) {
		case int:
			return float64(v), true
		case int8:
			return float64(v), true
		case int16:
			return float64(v), true
		case int32:
			return float64(v), true
		case int64:
			return float64(v), true
		case uint:
			return float64(v), true
		case uint8:
			return float64(v), true
		case uint16:
			return float64(v), true
		case uint32:
			return float64(v), true
		case uint64:
			return float64(v), true
		case float32:
			return float64(v), true
		case float64:
			return v, true
		default:
			return 0, false
		}
	}
	if af, ok := toFloat(a); ok {
		if bf, okB := toFloat(b); okB {
			if af < bf {
				return -1
			}
			if af > bf {
				return 1
			}
			return 0
		}
	}
	aStr := *hxrt.StdString(a)
	bStr := *hxrt.StdString(b)
	if aStr < bStr {
		return -1
	}
	if aStr > bStr {
		return 1
	}
	return 0
}

func Reflect_compareMethods(a any, b any) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)
	if !av.IsValid() || !bv.IsValid() {
		return !av.IsValid() && !bv.IsValid()
	}
	if av.Kind() == reflect.Func && bv.Kind() == reflect.Func {
		if av.IsNil() || bv.IsNil() {
			return av.IsNil() && bv.IsNil()
		}
		return av.Pointer() == bv.Pointer()
	}
	return reflect.DeepEqual(a, b)
}

func Reflect_field(obj any, field *string) any {
	if obj == nil {
		return nil
	}
	key := *hxrt.StdString(field)
	if metadataValue, ok := hxrt_typeClassMetadataField(obj, key); ok {
		return metadataValue
	}
	switch value := obj.(type) {
	case map[string]any:
		return value[key]
	case map[any]any:
		return value[key]
	case *map[string]any:
		if value == nil {
			return nil
		}
		return (*value)[key]
	case *map[any]any:
		if value == nil {
			return nil
		}
		return (*value)[key]
	}
	rv := reflect.ValueOf(obj)
	if !rv.IsValid() {
		return nil
	}
	if rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return nil
		}
		rv = rv.Elem()
	}
	if rv.Kind() == reflect.Struct {
		if fieldValue := rv.FieldByName(key); fieldValue.IsValid() && fieldValue.CanInterface() {
			return fieldValue.Interface()
		}
	}
	method := reflect.ValueOf(obj).MethodByName(key)
	if method.IsValid() {
		return method.Interface()
	}
	return nil
}

func Reflect_hasField(obj any, field *string) bool {
	if obj == nil {
		return false
	}
	key := *hxrt.StdString(field)
	if _, ok := hxrt_typeClassMetadataField(obj, key); ok {
		return true
	}
	switch value := obj.(type) {
	case map[string]any:
		_, ok := value[key]
		return ok
	case map[any]any:
		_, ok := value[key]
		return ok
	case *map[string]any:
		if value == nil {
			return false
		}
		_, ok := (*value)[key]
		return ok
	case *map[any]any:
		if value == nil {
			return false
		}
		_, ok := (*value)[key]
		return ok
	}
	rv := reflect.ValueOf(obj)
	if !rv.IsValid() {
		return false
	}
	if rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return false
		}
		rv = rv.Elem()
	}
	if rv.Kind() == reflect.Struct {
		if rv.FieldByName(key).IsValid() {
			return true
		}
	}
	return reflect.ValueOf(obj).MethodByName(key).IsValid()
}

func Reflect_setField(obj any, field *string, value any) {
	if obj == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		return
	}
	key := *hxrt.StdString(field)
	switch target := obj.(type) {
	case map[string]any:
		target[key] = value
		return
	case map[any]any:
		target[key] = value
		return
	case *map[string]any:
		if target == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
			return
		}
		(*target)[key] = value
		return
	case *map[any]any:
		if target == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
			return
		}
		(*target)[key] = value
		return
	}
	rv := reflect.ValueOf(obj)
	if !rv.IsValid() || rv.Kind() != reflect.Pointer {
		return
	}
	if rv.IsNil() {
		hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		return
	}
	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return
	}
	fieldValue := rv.FieldByName(key)
	if !fieldValue.IsValid() || !fieldValue.CanSet() {
		return
	}
	if value == nil {
		fieldValue.Set(reflect.Zero(fieldValue.Type()))
		return
	}
	incoming := reflect.ValueOf(value)
	if incoming.Type().AssignableTo(fieldValue.Type()) {
		fieldValue.Set(incoming)
		return
	}
	if incoming.Type().ConvertibleTo(fieldValue.Type()) {
		fieldValue.Set(incoming.Convert(fieldValue.Type()))
		return
	}
	if fieldValue.Kind() == reflect.Interface {
		fieldValue.Set(incoming)
	}
}

type haxe__ds__Option struct {
	tag    int
	params []any
}

var haxe__ds__Option_None *haxe__ds__Option = &haxe__ds__Option{tag: 1, params: []any{}}

func haxe__ds__Option_Some(value any) *haxe__ds__Option {
	return &haxe__ds__Option{tag: 0, params: []any{value}}
}

type ValueType struct {
	tag    int
	params []any
}

var ValueType_TNull *ValueType = &ValueType{tag: 0, params: []any{}}

var ValueType_TInt *ValueType = &ValueType{tag: 1, params: []any{}}

var ValueType_TFloat *ValueType = &ValueType{tag: 2, params: []any{}}

var ValueType_TBool *ValueType = &ValueType{tag: 3, params: []any{}}

var ValueType_TObject *ValueType = &ValueType{tag: 4, params: []any{}}

var ValueType_TFunction *ValueType = &ValueType{tag: 5, params: []any{}}

var ValueType_TUnknown *ValueType = &ValueType{tag: 8, params: []any{}}

func ValueType_TClass(c any) *ValueType {
	return &ValueType{tag: 6, params: []any{c}}
}

func ValueType_TEnum(e any) *ValueType {
	return &ValueType{tag: 7, params: []any{e}}
}

func hxrt_typeCallAny(callable any, args []any) (any, bool) {
	result := any(nil)
	ok := false
	defer func() {
		if recover() != nil {
			result = nil
			ok = false
		}
	}()
	if callable == nil {
		return nil, false
	}
	fn := reflect.ValueOf(callable)
	if !fn.IsValid() || fn.Kind() != reflect.Func {
		return nil, false
	}
	fnType := fn.Type()
	if fnType.NumIn() != len(args) {
		return nil, false
	}
	in := make([]reflect.Value, len(args))
	for i := 0; i < len(args); i++ {
		paramType := fnType.In(i)
		arg := args[i]
		if arg == nil {
			in[i] = reflect.Zero(paramType)
			continue
		}
		v := reflect.ValueOf(arg)
		if v.IsValid() && v.Type().AssignableTo(paramType) {
			in[i] = v
			continue
		}
		if v.IsValid() && v.Type().ConvertibleTo(paramType) {
			in[i] = v.Convert(paramType)
			continue
		}
		if paramType.Kind() == reflect.Interface && v.IsValid() {
			in[i] = v
			continue
		}
		return nil, false
	}
	out := fn.Call(in)
	if len(out) == 0 {
		return nil, true
	}
	first := out[0]
	if !first.IsValid() {
		return nil, true
	}
	result = first.Interface()
	ok = true
	return result, ok
}

func hxrt_typeArrayValues(value *hxrt.Array) []any {
	if value == nil {
		return []any{}
	}
	return value.Values()
}

func hxrt_typeResolvedClassName(value any) (string, bool) {
	switch current := value.(type) {
	case *hxrt__TypeClassValue:
		if current == nil || current.name == nil {
			return "", false
		}
		return *current.name, true
	case hxrt__TypeClassValue:
		if current.name == nil {
			return "", false
		}
		return *current.name, true
	case string:
		return current, true
	case *string:
		if current == nil {
			return "", false
		}
		return *current, true
	default:
		return "", false
	}
}

func hxrt_typeResolvedEnumName(value any) (string, bool) {
	switch current := value.(type) {
	case *hxrt__TypeEnumValue:
		if current == nil || current.name == nil {
			return "", false
		}
		return *current.name, true
	case hxrt__TypeEnumValue:
		if current.name == nil {
			return "", false
		}
		return *current.name, true
	case string:
		return current, true
	case *string:
		if current == nil {
			return "", false
		}
		return *current, true
	default:
		return "", false
	}
}

func hxrt_typeCreateClassInstance(className string, args []any) (any, bool) {
	switch className {
	case "Lambda":
		return nil, false
	case "Main":
		return nil, false
	case "StringBuf":
		return nil, false
	case "haxe.Int64Helper":
		return nil, false
	case "haxe._Int32.Int32_Impl_":
		return nil, false
	case "haxe._Int64.Int64_Impl_":
		return nil, false
	case "haxe._Int64.___Int64":
		return hxrt_typeCallAny(New_haxe___Int64_____Int64, args)
	case "haxe.ds.StringMap":
		return hxrt_typeCallAny(New_haxe__ds__StringMap, args)
	case "haxe.exceptions.NotImplementedException":
		return hxrt_typeCallAny(New_haxe__exceptions__NotImplementedException, args)
	case "haxe.exceptions.PosException":
		return hxrt_typeCallAny(New_haxe__exceptions__PosException, args)
	case "haxe.http.HttpBase":
		return hxrt_typeCallAny(New_haxe__http__HttpBase, args)
	case "haxe.io.Bytes":
		return hxrt_typeCallAny(New_haxe__io__Bytes, args)
	case "haxe.io.BytesBuffer":
		return hxrt_typeCallAny(New_haxe__io__BytesBuffer, args)
	case "haxe.io.FPHelper":
		return nil, false
	case "haxe.iterators.MapKeyValueIterator":
		return hxrt_typeCallAny(New_haxe__iterators__MapKeyValueIterator, args)
	case "sys.GoHttpHelpers":
		return nil, false
	default:
		return nil, false
	}
}

func hxrt_typeCreateClassEmptyInstance(className string) (any, bool) {
	switch className {
	case "haxe._Int64.___Int64":
		return &haxe___Int64_____Int64{}, true
	case "haxe.ds.StringMap":
		return &haxe__ds__StringMap{}, true
	case "haxe.exceptions.NotImplementedException":
		return &haxe__exceptions__NotImplementedException{}, true
	case "haxe.exceptions.PosException":
		return &haxe__exceptions__PosException{}, true
	case "haxe.http.HttpBase":
		return &haxe__http__HttpBase{}, true
	case "haxe.io.Bytes":
		return &haxe__io__Bytes{}, true
	case "haxe.io.BytesBuffer":
		return &haxe__io__BytesBuffer{}, true
	case "haxe.iterators.MapKeyValueIterator":
		return &haxe__iterators__MapKeyValueIterator{}, true
	default:
		return nil, false
	}
}

func hxrt_typeCreateEnumInstance(enumName string, constructorName string, constructorIndex int, useIndex bool, args []any) (any, bool) {
	switch enumName {
	case "ValueType":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TNull, true
			case 1:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TInt, true
			case 2:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TFloat, true
			case 3:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TBool, true
			case 4:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TObject, true
			case 5:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TFunction, true
			case 6:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(ValueType_TClass, args)
			case 7:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(ValueType_TEnum, args)
			case 8:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TUnknown, true
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "TNull":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TNull, true
		case "TInt":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TInt, true
		case "TFloat":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TFloat, true
		case "TBool":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TBool, true
		case "TObject":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TObject, true
		case "TFunction":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TFunction, true
		case "TClass":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(ValueType_TClass, args)
		case "TEnum":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(ValueType_TEnum, args)
		case "TUnknown":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TUnknown, true
		default:
			return nil, false
		}
	case "haxe.io.Encoding":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__io__Encoding_UTF8, true
			case 1:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__io__Encoding_RawNative, true
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "UTF8":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__io__Encoding_UTF8, true
		case "RawNative":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__io__Encoding_RawNative, true
		default:
			return nil, false
		}
	case "haxe.io.Error":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__io__Error_Blocked, true
			case 1:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__io__Error_Overflow, true
			case 2:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__io__Error_OutsideBounds, true
			case 3:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__io__Error_Custom, args)
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "Blocked":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__io__Error_Blocked, true
		case "Overflow":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__io__Error_Overflow, true
		case "OutsideBounds":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__io__Error_OutsideBounds, true
		case "Custom":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__io__Error_Custom, args)
		default:
			return nil, false
		}
	default:
		return nil, false
	}
}

func Type_getClass(o any) any {
	if hxrt.AnyEqualsNull(o) {
		return nil
	}
	switch value := o.(type) {
	case *hxrt__TypeClassValue:
		if value == nil {
			return nil
		}
		return value
	case hxrt__TypeClassValue:
		copyValue := value
		return &copyValue
	case *hxrt.Array:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Array")}
	case *haxe___Int64_____Int64:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe._Int64.___Int64")}
	case *haxe__ds__StringMap:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ds.StringMap")}
	case *haxe__exceptions__NotImplementedException:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.exceptions.NotImplementedException")}
	case *haxe__exceptions__PosException:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.exceptions.PosException")}
	case *haxe__http__HttpBase:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.http.HttpBase")}
	case *haxe__io__Bytes:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.io.Bytes")}
	case *haxe__io__BytesBuffer:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.io.BytesBuffer")}
	case *haxe__iterators__MapKeyValueIterator:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.iterators.MapKeyValueIterator")}
	default:
		return nil
	}
}

func Type_getEnum(o any) any {
	if hxrt.AnyEqualsNull(o) {
		return nil
	}
	switch value := o.(type) {
	case *hxrt__TypeEnumValue:
		if value == nil {
			return nil
		}
		return value
	case hxrt__TypeEnumValue:
		copyValue := value
		return &copyValue
	case *ValueType:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("ValueType")}
	case *haxe__io__Encoding:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("haxe.io.Encoding")}
	case *haxe__io__Error:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("haxe.io.Error")}
	default:
		return nil
	}
}

func Type_getSuperClass(c any) any {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return nil
	}
	switch className {
	case "Lambda":
		return nil
	case "Main":
		return nil
	case "StringBuf":
		return nil
	case "haxe.Int64Helper":
		return nil
	case "haxe._Int32.Int32_Impl_":
		return nil
	case "haxe._Int64.Int64_Impl_":
		return nil
	case "haxe._Int64.___Int64":
		return nil
	case "haxe.ds.StringMap":
		return nil
	case "haxe.exceptions.NotImplementedException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.exceptions.PosException")}
	case "haxe.exceptions.PosException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.Exception")}
	case "haxe.http.HttpBase":
		return nil
	case "haxe.io.Bytes":
		return nil
	case "haxe.io.BytesBuffer":
		return nil
	case "haxe.io.FPHelper":
		return nil
	case "haxe.iterators.MapKeyValueIterator":
		return nil
	case "sys.GoHttpHelpers":
		return nil
	default:
		return nil
	}
}

func Type_getClassName(c any) *string {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return nil
	}
	return hxrt.StringFromLiteral(className)
}

func Type_getClassFields(c any) *hxrt.Array {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return hxrt.NewArray()
	}
	switch className {
	case "Lambda":
		return hxrt.NewArray(hxrt.StringFromLiteral("exists"))
	case "Main":
		return hxrt.NewArray(hxrt.StringFromLiteral("main"))
	case "StringBuf":
		return hxrt.NewArray()
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe._Int32.Int32_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.Int64_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.___Int64":
		return hxrt.NewArray()
	case "haxe.ds.StringMap":
		return hxrt.NewArray()
	case "haxe.exceptions.NotImplementedException":
		return hxrt.NewArray()
	case "haxe.exceptions.PosException":
		return hxrt.NewArray()
	case "haxe.http.HttpBase":
		return hxrt.NewArray()
	case "haxe.io.Bytes":
		return hxrt.NewArray(hxrt.StringFromLiteral("__hx_fromNativeView"), hxrt.StringFromLiteral("alloc"), hxrt.StringFromLiteral("fastGet"), hxrt.StringFromLiteral("ofData"), hxrt.StringFromLiteral("ofHex"), hxrt.StringFromLiteral("ofString"), hxrt.StringFromLiteral("rawNativeUsesUtf16LE"))
	case "haxe.io.BytesBuffer":
		return hxrt.NewArray()
	case "haxe.io.FPHelper":
		return hxrt.NewArray(hxrt.StringFromLiteral("doubleToI64"), hxrt.StringFromLiteral("floatToI32"), hxrt.StringFromLiteral("i32ToFloat"), hxrt.StringFromLiteral("i64ToDouble"))
	case "haxe.iterators.MapKeyValueIterator":
		return hxrt.NewArray()
	case "sys.GoHttpHelpers":
		return hxrt.NewArray(hxrt.StringFromLiteral("captureApi"), hxrt.StringFromLiteral("getResponseHeaderValues"))
	default:
		return hxrt.NewArray()
	}
}

func Type_getInstanceFields(c any) *hxrt.Array {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return hxrt.NewArray()
	}
	switch className {
	case "Lambda":
		return hxrt.NewArray()
	case "Main":
		return hxrt.NewArray()
	case "StringBuf":
		return hxrt.NewArray()
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe._Int32.Int32_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.Int64_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.___Int64":
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	case "haxe.ds.StringMap":
		return hxrt.NewArray(hxrt.StringFromLiteral("clear"), hxrt.StringFromLiteral("copy"), hxrt.StringFromLiteral("copyIMap"), hxrt.StringFromLiteral("exists"), hxrt.StringFromLiteral("existsIMap"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("getIMap"), hxrt.StringFromLiteral("h"), hxrt.StringFromLiteral("iterator"), hxrt.StringFromLiteral("keyValueIterator"), hxrt.StringFromLiteral("keys"), hxrt.StringFromLiteral("remove"), hxrt.StringFromLiteral("removeIMap"), hxrt.StringFromLiteral("set"), hxrt.StringFromLiteral("setIMap"), hxrt.StringFromLiteral("toString"))
	case "haxe.exceptions.NotImplementedException":
		return hxrt.NewArray(hxrt.StringFromLiteral("details"), hxrt.StringFromLiteral("get_message"), hxrt.StringFromLiteral("get_native"), hxrt.StringFromLiteral("get_previous"), hxrt.StringFromLiteral("get_stack"), hxrt.StringFromLiteral("message"), hxrt.StringFromLiteral("native"), hxrt.StringFromLiteral("posInfos"), hxrt.StringFromLiteral("previous"), hxrt.StringFromLiteral("stack"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("unwrap"))
	case "haxe.exceptions.PosException":
		return hxrt.NewArray(hxrt.StringFromLiteral("details"), hxrt.StringFromLiteral("get_message"), hxrt.StringFromLiteral("get_native"), hxrt.StringFromLiteral("get_previous"), hxrt.StringFromLiteral("get_stack"), hxrt.StringFromLiteral("message"), hxrt.StringFromLiteral("native"), hxrt.StringFromLiteral("posInfos"), hxrt.StringFromLiteral("previous"), hxrt.StringFromLiteral("stack"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("unwrap"))
	case "haxe.http.HttpBase":
		return hxrt.NewArray(hxrt.StringFromLiteral("addHeader"), hxrt.StringFromLiteral("addParameter"), hxrt.StringFromLiteral("emptyOnData"), hxrt.StringFromLiteral("get_responseData"), hxrt.StringFromLiteral("hasOnData"), hxrt.StringFromLiteral("headers"), hxrt.StringFromLiteral("onBytes"), hxrt.StringFromLiteral("onData"), hxrt.StringFromLiteral("onError"), hxrt.StringFromLiteral("onStatus"), hxrt.StringFromLiteral("params"), hxrt.StringFromLiteral("postBytes"), hxrt.StringFromLiteral("postData"), hxrt.StringFromLiteral("request"), hxrt.StringFromLiteral("responseAsString"), hxrt.StringFromLiteral("responseBytes"), hxrt.StringFromLiteral("responseData"), hxrt.StringFromLiteral("setHeader"), hxrt.StringFromLiteral("setParameter"), hxrt.StringFromLiteral("setPostBytes"), hxrt.StringFromLiteral("setPostData"), hxrt.StringFromLiteral("success"), hxrt.StringFromLiteral("url"))
	case "haxe.io.Bytes":
		return hxrt.NewArray(hxrt.StringFromLiteral("__hx_dataExposed"), hxrt.StringFromLiteral("__hx_nativeView"), hxrt.StringFromLiteral("__hx_raw"), hxrt.StringFromLiteral("__hx_rawValid"), hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("blit"), hxrt.StringFromLiteral("compare"), hxrt.StringFromLiteral("fill"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("getData"), hxrt.StringFromLiteral("getDouble"), hxrt.StringFromLiteral("getFloat"), hxrt.StringFromLiteral("getInt32"), hxrt.StringFromLiteral("getInt64"), hxrt.StringFromLiteral("getString"), hxrt.StringFromLiteral("getUInt16"), hxrt.StringFromLiteral("length"), hxrt.StringFromLiteral("readString"), hxrt.StringFromLiteral("set"), hxrt.StringFromLiteral("setDouble"), hxrt.StringFromLiteral("setFloat"), hxrt.StringFromLiteral("setInt32"), hxrt.StringFromLiteral("setInt64"), hxrt.StringFromLiteral("setUInt16"), hxrt.StringFromLiteral("sub"), hxrt.StringFromLiteral("toHex"), hxrt.StringFromLiteral("toString"))
	case "haxe.io.BytesBuffer":
		return hxrt.NewArray(hxrt.StringFromLiteral("add"), hxrt.StringFromLiteral("addByte"), hxrt.StringFromLiteral("addBytes"), hxrt.StringFromLiteral("addDouble"), hxrt.StringFromLiteral("addFloat"), hxrt.StringFromLiteral("addInt32"), hxrt.StringFromLiteral("addInt64"), hxrt.StringFromLiteral("addString"), hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("getBytes"), hxrt.StringFromLiteral("get_length"), hxrt.StringFromLiteral("length"))
	case "haxe.io.FPHelper":
		return hxrt.NewArray()
	case "haxe.iterators.MapKeyValueIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("keys"), hxrt.StringFromLiteral("map"), hxrt.StringFromLiteral("next"))
	case "sys.GoHttpHelpers":
		return hxrt.NewArray()
	default:
		return hxrt.NewArray()
	}
}

func Type_getEnumName(e any) *string {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	return hxrt.StringFromLiteral(enumName)
}

func Type_resolveClass(name *string) any {
	if name == nil {
		return nil
	}
	rawName := *hxrt.StdString(name)
	switch rawName {
	case "Lambda":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "Main":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "StringBuf":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Int64Helper":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int32.Int32_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int64.Int64_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int64.___Int64":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds.StringMap":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.exceptions.NotImplementedException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.exceptions.PosException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.http.HttpBase":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.Bytes":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.BytesBuffer":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.FPHelper":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.iterators.MapKeyValueIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.GoHttpHelpers":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	default:
		return nil
	}
}

func Type_resolveEnum(name *string) any {
	if name == nil {
		return nil
	}
	rawName := *hxrt.StdString(name)
	switch rawName {
	case "ValueType":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.Encoding":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.Error":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	default:
		return nil
	}
}

func Type_createInstance(cl any, args *hxrt.Array) any {
	className, ok := hxrt_typeResolvedClassName(cl)
	if !ok {
		return nil
	}
	instance, ok := hxrt_typeCreateClassInstance(className, hxrt_typeArrayValues(args))
	if !ok {
		return nil
	}
	return instance
}

func Type_createEmptyInstance(cl any) any {
	className, ok := hxrt_typeResolvedClassName(cl)
	if !ok {
		return nil
	}
	instance, ok := hxrt_typeCreateClassEmptyInstance(className)
	if !ok {
		return nil
	}
	return instance
}

func Type_createEnum(e any, constr *string, params *hxrt.Array) any {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	constructorName := ""
	if constr != nil {
		constructorName = *hxrt.StdString(constr)
	}
	enumValue, ok := hxrt_typeCreateEnumInstance(enumName, constructorName, 0, false, hxrt_typeArrayValues(params))
	if !ok {
		return nil
	}
	return enumValue
}

func Type_createEnumIndex(e any, index int, params *hxrt.Array) any {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	enumValue, ok := hxrt_typeCreateEnumInstance(enumName, "", index, true, hxrt_typeArrayValues(params))
	if !ok {
		return nil
	}
	return enumValue
}

func Type_enumConstructor(e any) *string {
	if hxrt.AnyEqualsNull(e) {
		return nil
	}
	switch value := e.(type) {
	case *ValueType:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("TNull")
		case 1:
			return hxrt.StringFromLiteral("TInt")
		case 2:
			return hxrt.StringFromLiteral("TFloat")
		case 3:
			return hxrt.StringFromLiteral("TBool")
		case 4:
			return hxrt.StringFromLiteral("TObject")
		case 5:
			return hxrt.StringFromLiteral("TFunction")
		case 6:
			return hxrt.StringFromLiteral("TClass")
		case 7:
			return hxrt.StringFromLiteral("TEnum")
		case 8:
			return hxrt.StringFromLiteral("TUnknown")
		default:
			return nil
		}
	case *haxe__io__Encoding:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("UTF8")
		case 1:
			return hxrt.StringFromLiteral("RawNative")
		default:
			return nil
		}
	case *haxe__io__Error:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("Blocked")
		case 1:
			return hxrt.StringFromLiteral("Overflow")
		case 2:
			return hxrt.StringFromLiteral("OutsideBounds")
		case 3:
			return hxrt.StringFromLiteral("Custom")
		default:
			return nil
		}
	default:
		return nil
	}
}

func Type_enumIndex(e any) int {
	if hxrt.AnyEqualsNull(e) {
		return -1
	}
	switch value := e.(type) {
	case *ValueType:
		if value == nil {
			return -1
		}
		return value.tag
	case *haxe__io__Encoding:
		if value == nil {
			return -1
		}
		return value.tag
	case *haxe__io__Error:
		if value == nil {
			return -1
		}
		return value.tag
	default:
		return -1
	}
}

func Type_getEnumConstructs(e any) *hxrt.Array {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return hxrt.NewArray()
	}
	switch enumName {
	case "ValueType":
		return hxrt.NewArray(hxrt.StringFromLiteral("TNull"), hxrt.StringFromLiteral("TInt"), hxrt.StringFromLiteral("TFloat"), hxrt.StringFromLiteral("TBool"), hxrt.StringFromLiteral("TObject"), hxrt.StringFromLiteral("TFunction"), hxrt.StringFromLiteral("TClass"), hxrt.StringFromLiteral("TEnum"), hxrt.StringFromLiteral("TUnknown"))
	case "haxe.io.Encoding":
		return hxrt.NewArray(hxrt.StringFromLiteral("UTF8"), hxrt.StringFromLiteral("RawNative"))
	case "haxe.io.Error":
		return hxrt.NewArray(hxrt.StringFromLiteral("Blocked"), hxrt.StringFromLiteral("Overflow"), hxrt.StringFromLiteral("OutsideBounds"), hxrt.StringFromLiteral("Custom"))
	default:
		return hxrt.NewArray()
	}
}

func Type_enumParameters(e any) *hxrt.Array {
	if hxrt.AnyEqualsNull(e) {
		return hxrt.NewArray()
	}
	switch value := e.(type) {
	case *ValueType:
		if value == nil || value.params == nil {
			return hxrt.NewArray()
		}
		return hxrt.NewArray(value.params...)
	case *haxe__io__Encoding:
		if value == nil || value.params == nil {
			return hxrt.NewArray()
		}
		return hxrt.NewArray(value.params...)
	case *haxe__io__Error:
		if value == nil || value.params == nil {
			return hxrt.NewArray()
		}
		return hxrt.NewArray(value.params...)
	default:
		return hxrt.NewArray()
	}
}

func Type_allEnums(e any) *hxrt.Array {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return hxrt.NewArray()
	}
	switch enumName {
	case "ValueType":
		return hxrt.NewArray(ValueType_TNull, ValueType_TInt, ValueType_TFloat, ValueType_TBool, ValueType_TObject, ValueType_TFunction, ValueType_TUnknown)
	case "haxe.io.Encoding":
		return hxrt.NewArray(haxe__io__Encoding_UTF8, haxe__io__Encoding_RawNative)
	case "haxe.io.Error":
		return hxrt.NewArray(haxe__io__Error_Blocked, haxe__io__Error_Overflow, haxe__io__Error_OutsideBounds)
	default:
		return hxrt.NewArray()
	}
}

func Type_typeof(v any) *ValueType {
	if hxrt.AnyEqualsNull(v) {
		return ValueType_TNull
	}
	switch v.(type) {
	case *hxrt__TypeClassValue, hxrt__TypeClassValue, *hxrt__TypeEnumValue, hxrt__TypeEnumValue:
		return ValueType_TObject
	}
	if enumValue := Type_getEnum(v); enumValue != nil {
		return ValueType_TEnum(enumValue)
	}
	if classValue := Type_getClass(v); classValue != nil {
		return ValueType_TClass(classValue)
	}
	switch v.(type) {
	case bool:
		return ValueType_TBool
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr:
		return ValueType_TInt
	case float32, float64:
		return ValueType_TFloat
	case string, *string:
		return ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("String")})
	case *hxrt.Array:
		return ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Array")})
	}
	ref := reflect.ValueOf(v)
	if !ref.IsValid() {
		return ValueType_TNull
	}
	switch ref.Kind() {
	case reflect.Func:
		return ValueType_TFunction
	case reflect.Slice, reflect.Array:
		return ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Array")})
	case reflect.Map, reflect.Struct, reflect.Interface, reflect.Pointer:
		return ValueType_TObject
	default:
		return ValueType_TUnknown
	}
}

func Type_enumEq(a any, b any) bool {
	return reflect.DeepEqual(a, b)
}

func hxrt_typeClassMetadataField(value any, key string) (any, bool) {
	className, ok := hxrt_typeResolvedClassName(value)
	if !ok {
		return nil, false
	}
	switch className {
	default:
		return nil, false
	}
}
