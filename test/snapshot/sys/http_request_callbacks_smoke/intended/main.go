package main

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"snapshot/hxrt"
	"strings"
)

func main() {
	http := New_sys__Http(hxrt.StringFromLiteral("data:text/plain,hello%20from%20haxe.go"))
	_ = http
	dataLog := hxrt.StringFromLiteral("")
	_ = dataLog
	statusLog := -1
	_ = statusLog
	byteCount := -1
	_ = byteCount
	errLog := hxrt.StringFromLiteral("")
	_ = errLog
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
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("data="), dataLog))
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("bytes="), byteCount))
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("status="), statusLog))
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("response="), http.get_responseData()))
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("error="), errLog))
	post := New_sys__Http(hxrt.StringFromLiteral("data:text/plain,ignored"))
	_ = post
	post.setPostData(hxrt.StringFromLiteral("post-body"))
	postData := hxrt.StringFromLiteral("")
	_ = postData
	post.onData = func(data *string) {
		postData = data
	}
	post.request(true)
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("post="), postData))
	bad := New_sys__Http(hxrt.StringFromLiteral("://bad"))
	_ = bad
	badErr := hxrt.StringFromLiteral("")
	_ = badErr
	bad.onError = func(msg *string) {
		badErr = msg
	}
	bad.request()
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("bad="), badErr))
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("direct="), sys__Http_requestUrl(hxrt.StringFromLiteral("data:text/plain,direct%20ok"))))
}

type haxe__io__Encoding struct {
}

type haxe__io__Input struct {
}

type haxe__io__Output struct {
}

type haxe__io__Bytes struct {
	b      []int
	length int
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
	raw := hxrt.BytesFromString(value)
	return &haxe__io__Bytes{b: raw, length: len(raw)}
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
}

func New_haxe__io__BytesBuffer() *haxe__io__BytesBuffer {
	return &haxe__io__BytesBuffer{b: []int{}}
}

func (self *haxe__io__BytesBuffer) addByte(value int) {
	self.b = append(self.b, value)
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

type hxrt__http__Pair struct {
	name  *string
	value *string
}

type sys__Http struct {
	url              *string
	responseAsString *string
	responseBytes    *haxe__io__Bytes
	postData         *string
	postBytes        *haxe__io__Bytes
	headers          []hxrt__http__Pair
	params           []hxrt__http__Pair
	onData           func(*string)
	onBytes          func(*haxe__io__Bytes)
	onError          func(*string)
	onStatus         func(int)
}

func New_sys__Http(url *string) *sys__Http {
	self := &sys__Http{url: url, headers: []hxrt__http__Pair{}, params: []hxrt__http__Pair{}}
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

func (self *sys__Http) get_responseData() *string {
	if self == nil {
		return hxrt.StringFromLiteral("")
	}
	if self.responseAsString == nil && self.responseBytes != nil {
		self.responseAsString = self.responseBytes.toString()
	}
	return self.responseAsString
}

func (self *sys__Http) request(post ...bool) {
	if self == nil {
		return
	}
	isPost := false
	if len(post) > 0 {
		isPost = post[0]
	}
	if self.postData != nil || self.postBytes != nil {
		isPost = true
	}
	rawUrl := *hxrt.StdString(self.url)
	parsedURL, err := url.Parse(rawUrl)
	if err != nil || parsedURL == nil {
		if self.onError != nil {
			self.onError(hxrt.StringFromLiteral("Invalid URL"))
		}
		return
	}
	if parsedURL.Scheme == "data" {
		payload := parsedURL.Opaque
		if isPost {
			if self.postBytes != nil {
				rawBody := make([]byte, len(self.postBytes.b))
				for i := 0; i < len(self.postBytes.b); i++ {
					rawBody[i] = byte(self.postBytes.b[i])
				}
				payload = string(rawBody)
			} else if self.postData != nil {
				payload = *hxrt.StdString(self.postData)
			}
		}
		commaIndex := strings.Index(payload, ",")
		if commaIndex >= 0 {
			payload = payload[commaIndex+1:]
		}
		decoded, decodeErr := url.QueryUnescape(payload)
		if decodeErr == nil {
			payload = decoded
		}
		rawPayload := []byte(payload)
		intPayload := make([]int, len(rawPayload))
		for i := 0; i < len(rawPayload); i++ {
			intPayload[i] = int(rawPayload[i])
		}
		self.responseBytes = &haxe__io__Bytes{b: intPayload, length: len(intPayload)}
		self.responseAsString = hxrt.StringFromLiteral(payload)
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
	query := parsedURL.Query()
	for _, param := range self.params {
		query.Set(*hxrt.StdString(param.name), *hxrt.StdString(param.value))
	}
	var bodyReader io.Reader = nil
	if isPost {
		if self.postBytes != nil {
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
				self.headers = append(self.headers, hxrt__http__Pair{name: hxrt.StringFromLiteral("Content-Type"), value: hxrt.StringFromLiteral("application/x-www-form-urlencoded")})
			}
		}
	} else {
		parsedURL.RawQuery = query.Encode()
	}
	method := "GET"
	if isPost {
		method = "POST"
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
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		if self.onError != nil {
			self.onError(hxrt.StringFromLiteral(err.Error()))
		}
		return
	}
	defer response.Body.Close()
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
	if response.StatusCode >= 400 {
		if self.onError != nil {
			self.onError(hxrt.StringFromLiteral(response.Status))
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

func sys__Http_requestUrl(url *string) *string {
	self := New_sys__Http(url)
	result := hxrt.StringFromLiteral("")
	self.onData = func(data *string) { result = data }
	self.onError = func(msg *string) { result = msg }
	self.request()
	return result
}
