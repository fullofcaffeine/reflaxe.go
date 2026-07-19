package main

import "snapshot/hxrt"

type I_sys__Http interface {
	setHeader(name *string, value *string)
	addHeader(header *string, value *string)
	setParameter(name *string, value *string)
	addParameter(name *string, value *string)
	setPostData(data *string)
	setPostBytes(data *haxe__io__Bytes)
	request(post any)
	hasOnData() bool
	success(data *haxe__io__Bytes)
	get_responseData() *string
	fileTransfert(argname *string, filename *string, file *haxe__io__Input, size int, mimeType *string)
	fileTransfer(argname *string, filename *string, file *haxe__io__Input, size int, mimeType *string)
	customRequest(post bool, api *haxe__io__Output, sock *sys__net__Socket, method *string)
	getResponseHeaderValues(key *string) *hxrt.Array
	requestWith(post bool, api *haxe__io__Output, sock *sys__net__Socket, method *string)
	handleDataRequest(post bool, api *haxe__io__Output, method *string)
	recordResponseHeaders(response *hxrt.HttpResponse)
	resetResponseHeaders()
	hasHeader(name *string) bool
	buildMultipartBody(upload map[string]any) *string
	encodedParameters() *string
}

type sys__Http struct {
	*haxe__http__HttpBase
	__hx_this              I_sys__Http
	noShutdown             bool
	cnxTimeout             float64
	responseHeaders        *haxe__ds__StringMap
	responseHeadersSameKey *haxe__ds__StringMap
	file                   map[string]any
}

func New_sys__Http(url *string) *sys__Http {
	self := &sys__Http{}
	self.haxe__http__HttpBase = New_haxe__http__HttpBase(url)
	self.haxe__http__HttpBase.__hx_this = self
	self.__hx_this = self
	self.noShutdown = false
	self.cnxTimeout = 10
	self.__hx_this.resetResponseHeaders()
	return self
}

func (self *sys__Http) request(post any) {
	usePost := (((((post != nil) && (post.(bool) == true)) || (self.postBytes != nil)) || !hxrt.StringEqualStringPtr(self.postData, nil)) || (self.file != nil))
	self.__hx_this.requestWith(usePost, nil, nil, nil)
}

func (self *sys__Http) fileTransfert(argname *string, filename *string, file *haxe__io__Input, size int, mimeType *string) {
	self.__hx_this.fileTransfer(argname, filename, file, size, mimeType)
}

func (self *sys__Http) fileTransfer(argname *string, filename *string, file *haxe__io__Input, size int, mimeType *string) {
	hx_obj_39 := map[string]any{}
	hx_obj_39["param"] = argname
	hx_obj_39["filename"] = filename
	hx_obj_39["io"] = file
	hx_obj_39["size"] = size
	hx_obj_39["mimeType"] = mimeType
	self.file = hx_obj_39
}

func (self *sys__Http) customRequest(post bool, api *haxe__io__Output, sock *sys__net__Socket, method *string) {
	self.__hx_this.requestWith((post || (self.file != nil)), api, sock, method)
}

func (self *sys__Http) getResponseHeaderValues(key *string) *hxrt.Array {
	values := func(hx_value_40 any) *hxrt.Array {
		if hx_value_40 == nil {
			var hx_zero_41 *hxrt.Array
			return hx_zero_41
		}
		return hx_value_40.(*hxrt.Array)
	}(self.responseHeadersSameKey.__hx_this.get(key))
	if values == nil {
		normalized := hxrt.StringToLowerCaseStringPtr(key)
		if !hxrt.StringEqualStringPtr(normalized, key) {
			values = func(hx_value_42 any) *hxrt.Array {
				if hx_value_42 == nil {
					var hx_zero_43 *hxrt.Array
					return hx_zero_43
				}
				return hx_value_42.(*hxrt.Array)
			}(self.responseHeadersSameKey.__hx_this.get(normalized))
		}
	}
	if values != nil {
		return values
	}
	var this1 haxe__IMap = self.responseHeaders
	value := func(hx_value_44 any) *string {
		if hx_value_44 == nil {
			var hx_zero_45 *string
			return hx_zero_45
		}
		return hx_value_44.(*string)
	}(this1.(*haxe__ds__StringMap).__hx_this.get(key))
	if hxrt.StringEqualStringPtr(value, nil) {
		normalized_1 := hxrt.StringToLowerCaseStringPtr(key)
		if !hxrt.StringEqualStringPtr(normalized_1, key) {
			var this1_1 haxe__IMap = self.responseHeaders
			value = func(hx_value_46 any) *string {
				if hx_value_46 == nil {
					var hx_zero_47 *string
					return hx_zero_47
				}
				return hx_value_46.(*string)
			}(this1_1.(*haxe__ds__StringMap).__hx_this.get(normalized_1))
		}
	}
	var hx_if_48 *hxrt.Array
	if hxrt.StringEqualStringPtr(value, nil) {
		hx_if_48 = nil
	} else {
		hx_if_48 = hxrt.NewArray(value)
	}
	return hx_if_48
}

func (self *sys__Http) requestWith(post bool, api *haxe__io__Output, sock *sys__net__Socket, method *string) {
	self.responseAsString = nil
	self.responseBytes = nil
	self.__hx_this.resetResponseHeaders()
	if StringTools_startsWith(self.url, hxrt.StringFromLiteral("data:")) {
		self.__hx_this.handleDataRequest(post, api, method)
		return
	}
	request := hxrt.HttpRequestNew(self.url, post, method, self.cnxTimeout)
	_g := 0
	_g1 := self.params
	for _g < _g1.Len() {
		parameter := func(hx_value_49 any) map[string]any {
			if hx_value_49 == nil {
				var hx_zero_50 map[string]any
				return hx_zero_50
			}
			return hx_value_49.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		hxrt.HttpRequestAddParameter(request, func(hx_obj_51 map[string]any) *string {
			hx_field_52 := hx_obj_51["name"]
			if hx_field_52 == nil {
				var hx_zero_53 *string
				return hx_zero_53
			}
			return hx_field_52.(*string)
		}(parameter), func(hx_obj_54 map[string]any) *string {
			hx_field_55 := hx_obj_54["value"]
			if hx_field_55 == nil {
				var hx_zero_56 *string
				return hx_zero_56
			}
			return hx_field_55.(*string)
		}(parameter))
	}
	_g_1 := 0
	_g1_1 := self.headers
	for _g_1 < _g1_1.Len() {
		header := func(hx_value_57 any) map[string]any {
			if hx_value_57 == nil {
				var hx_zero_58 map[string]any
				return hx_zero_58
			}
			return hx_value_57.(map[string]any)
		}(_g1_1.Get(_g_1))
		_g_1 = int(int32((_g_1 + 1)))
		hxrt.HttpRequestAddHeader(request, func(hx_obj_59 map[string]any) *string {
			hx_field_60 := hx_obj_59["name"]
			if hx_field_60 == nil {
				var hx_zero_61 *string
				return hx_zero_61
			}
			return hx_field_60.(*string)
		}(header), func(hx_obj_62 map[string]any) *string {
			hx_field_63 := hx_obj_62["value"]
			if hx_field_63 == nil {
				var hx_zero_64 *string
				return hx_zero_64
			}
			return hx_field_63.(*string)
		}(header))
	}
	if self.file != nil {
		hxrt.HttpRequestSetBodyString(request, self.__hx_this.buildMultipartBody(self.file))
		if !self.__hx_this.hasHeader(hxrt.StringFromLiteral("Content-Type")) {
			hxrt.HttpRequestAddHeader(request, hxrt.StringFromLiteral("Content-Type"), hxrt.StringFromLiteral("multipart/form-data; boundary=hxrt-go-boundary"))
		}
	} else {
		if self.postBytes != nil {
			hxrt.HttpRequestSetBodyView(request, self.postBytes.__hx_this.__hx_nativeView())
		} else {
			if !hxrt.StringEqualStringPtr(self.postData, nil) {
				hxrt.HttpRequestSetBodyString(request, self.postData)
			}
		}
	}
	proxy := sys__Http_PROXY
	if (proxy != nil) && !hxrt.StringEqualStringPtr(func(hx_obj_91 map[string]any) *string {
		hx_field_92 := hx_obj_91["host"]
		if hx_field_92 == nil {
			var hx_zero_93 *string
			return hx_zero_93
		}
		return hx_field_92.(*string)
	}(proxy), nil) {
		var hx_if_74 *string
		if func(hx_obj_65 map[string]any) map[string]any {
			hx_field_66 := hx_obj_65["auth"]
			if hx_field_66 == nil {
				var hx_zero_67 map[string]any
				return hx_zero_67
			}
			return hx_field_66.(map[string]any)
		}(proxy) == nil {
			hx_if_74 = nil
		} else {
			hx_if_74 = func(hx_obj_71 map[string]any) *string {
				hx_field_72 := hx_obj_71["user"]
				if hx_field_72 == nil {
					var hx_zero_73 *string
					return hx_zero_73
				}
				return hx_field_72.(*string)
			}(func(hx_obj_68 map[string]any) map[string]any {
				hx_field_69 := hx_obj_68["auth"]
				if hx_field_69 == nil {
					var hx_zero_70 map[string]any
					return hx_zero_70
				}
				return hx_field_69.(map[string]any)
			}(proxy))
		}
		user := hx_if_74
		var hx_if_84 *string
		if func(hx_obj_75 map[string]any) map[string]any {
			hx_field_76 := hx_obj_75["auth"]
			if hx_field_76 == nil {
				var hx_zero_77 map[string]any
				return hx_zero_77
			}
			return hx_field_76.(map[string]any)
		}(proxy) == nil {
			hx_if_84 = nil
		} else {
			hx_if_84 = func(hx_obj_81 map[string]any) *string {
				hx_field_82 := hx_obj_81["pass"]
				if hx_field_82 == nil {
					var hx_zero_83 *string
					return hx_zero_83
				}
				return hx_field_82.(*string)
			}(func(hx_obj_78 map[string]any) map[string]any {
				hx_field_79 := hx_obj_78["auth"]
				if hx_field_79 == nil {
					var hx_zero_80 map[string]any
					return hx_zero_80
				}
				return hx_field_79.(map[string]any)
			}(proxy))
		}
		pass := hx_if_84
		hxrt.HttpRequestSetProxy(request, func(hx_obj_85 map[string]any) *string {
			hx_field_86 := hx_obj_85["host"]
			if hx_field_86 == nil {
				var hx_zero_87 *string
				return hx_zero_87
			}
			return hx_field_86.(*string)
		}(proxy), func(hx_obj_88 map[string]any) int {
			hx_field_89 := hx_obj_88["port"]
			if hx_field_89 == nil {
				var hx_zero_90 int
				return hx_zero_90
			}
			return hx_field_89.(int)
		}(proxy), user, pass)
	}
	if sock != nil {
		hxrt.HttpRequestSetSocket(request, sock.handle)
	}
	response := hxrt.HttpRequestExecute(request)
	status := hxrt.HttpResponseStatus(response)
	nativeError := hxrt.HttpResponseError(response)
	if (status == 0) && !hxrt.StringEqualStringPtr(nativeError, nil) {
		func(hx_fn func(*string), hx_arg_0 *string) {
			if hx_fn == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
				return
			}
			hx_fn(hx_arg_0)
		}(self.onError, nativeError)
		return
	}
	self.__hx_this.recordResponseHeaders(response)
	func(hx_fn func(int), hx_arg_0 int) {
		if hx_fn == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
			return
		}
		hx_fn(hx_arg_0)
	}(self.onStatus, status)
	if !hxrt.StringEqualStringPtr(nativeError, nil) {
		func(hx_fn func(*string), hx_arg_0 *string) {
			if hx_fn == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
				return
			}
			hx_fn(hx_arg_0)
		}(self.onError, nativeError)
		return
	}
	payload := haxe__io__Bytes___hx_fromNativeView(hxrt.HttpResponseBody(response))
	self.responseBytes = payload
	self.responseAsString = payload.__hx_this.toString()
	sys__Http_capture(api, payload)
	if status >= 400 {
		func(hx_fn func(*string), hx_arg_0 *string) {
			if hx_fn == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
				return
			}
			hx_fn(hx_arg_0)
		}(self.onError, hxrt.StringConcatAny(hxrt.StringFromLiteral("Http Error #"), status))
		return
	}
	func(hx_fn func(*string), hx_arg_0 *string) {
		if hx_fn == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
			return
		}
		hx_fn(hx_arg_0)
	}(self.onData, self.responseAsString)
	func(hx_fn func(*haxe__io__Bytes), hx_arg_0 *haxe__io__Bytes) {
		if hx_fn == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
			return
		}
		hx_fn(hx_arg_0)
	}(self.onBytes, payload)
}

func (self *sys__Http) handleDataRequest(post bool, api *haxe__io__Output, method *string) {
	encoded := hxrt.StringSubstrStringPtr(self.url, hxrt.StringLengthStringPtr(hxrt.StringFromLiteral("data:")), 0, false)
	mediaType := hxrt.StringFromLiteral("text/plain")
	comma := sys__Http_firstComma(encoded)
	if comma >= 0 {
		if comma > 0 {
			mediaType = hxrt.StringSubstrStringPtr(encoded, 0, comma, true)
		}
		encoded = hxrt.StringSubstrStringPtr(encoded, int(int32((hxrt.Int32Wrap(comma) + hxrt.Int32Wrap(1)))), 0, false)
	}
	if post {
		if self.file != nil {
			encoded = hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("multipart file="), func(hx_obj_94 map[string]any) *string {
				hx_field_95 := hx_obj_94["filename"]
				if hx_field_95 == nil {
					var hx_zero_96 *string
					return hx_zero_96
				}
				return hx_field_95.(*string)
			}(self.file)), hxrt.StringFromLiteral(";mime=")), func(hx_obj_97 map[string]any) *string {
				hx_field_98 := hx_obj_97["mimeType"]
				if hx_field_98 == nil {
					var hx_zero_99 *string
					return hx_zero_99
				}
				return hx_field_98.(*string)
			}(self.file)), hxrt.StringFromLiteral(";size=")), func(hx_obj_100 map[string]any) int {
				hx_field_101 := hx_obj_100["size"]
				if hx_field_101 == nil {
					var hx_zero_102 int
					return hx_zero_102
				}
				return hx_field_101.(int)
			}(self.file))
		} else {
			if self.postBytes != nil {
				encoded = self.postBytes.__hx_this.toString()
			} else {
				if !hxrt.StringEqualStringPtr(self.postData, nil) {
					encoded = self.postData
				} else {
					encoded = self.__hx_this.encodedParameters()
				}
			}
		}
	}
	payloadText := StringTools_urlDecode(encoded)
	normalizedMethod := sys__Http_normalizedMethod(method)
	if !hxrt.StringEqualStringPtr(normalizedMethod, nil) {
		payloadText = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(normalizedMethod, hxrt.StringFromLiteral(" ")), payloadText)
	}
	payload := haxe__io__Bytes_ofString(payloadText, nil)
	self.responseBytes = payload
	self.responseAsString = payloadText
	var this1 haxe__IMap = self.responseHeaders
	this1.(*haxe__ds__StringMap).__hx_this.set(hxrt.StringFromLiteral("content-type"), mediaType)
	var this1_1 haxe__IMap = self.responseHeaders
	this1_1.(*haxe__ds__StringMap).__hx_this.set(hxrt.StringFromLiteral("Content-Type"), mediaType)
	sys__Http_capture(api, payload)
	func(hx_fn func(int), hx_arg_0 int) {
		if hx_fn == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
			return
		}
		hx_fn(hx_arg_0)
	}(self.onStatus, 200)
	func(hx_fn func(*string), hx_arg_0 *string) {
		if hx_fn == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
			return
		}
		hx_fn(hx_arg_0)
	}(self.onData, payloadText)
	func(hx_fn func(*haxe__io__Bytes), hx_arg_0 *haxe__io__Bytes) {
		if hx_fn == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
			return
		}
		hx_fn(hx_arg_0)
	}(self.onBytes, payload)
}

func (self *sys__Http) recordResponseHeaders(response *hxrt.HttpResponse) {
	count := hxrt.HttpResponseHeaderCount(response)
	_g := 0
	_g1 := count
	for _g < _g1 {
		hx_post_103 := _g
		_g = int(int32((_g + 1)))
		headerIndex := hx_post_103
		name := hxrt.StdString(hxrt.HttpResponseHeaderName(response, headerIndex))
		normalized := hxrt.StringToLowerCaseStringPtr(name)
		valueCount := hxrt.HttpResponseHeaderValueCount(response, headerIndex)
		values := hxrt.NewArray()
		_g_1 := 0
		_g1_1 := valueCount
		for _g_1 < _g1_1 {
			hx_post_104 := _g_1
			_g_1 = int(int32((_g_1 + 1)))
			valueIndex := hx_post_104
			values.Push(hxrt.StdString(hxrt.HttpResponseHeaderValue(response, headerIndex, valueIndex)))
		}
		if values.Len() == 0 {
			continue
		}
		last := func(hx_value_106 any) *string {
			if hx_value_106 == nil {
				var hx_zero_107 *string
				return hx_zero_107
			}
			return hx_value_106.(*string)
		}(values.Get(int(int32((hxrt.Int32Wrap(values.Len()) - hxrt.Int32Wrap(1))))))
		var this1 haxe__IMap = self.responseHeaders
		this1.(*haxe__ds__StringMap).__hx_this.set(name, last)
		if !hxrt.StringEqualStringPtr(normalized, name) {
			var this1_1 haxe__IMap = self.responseHeaders
			this1_1.(*haxe__ds__StringMap).__hx_this.set(normalized, last)
		}
		if values.Len() > 1 {
			self.responseHeadersSameKey.__hx_this.set(name, values)
			if !hxrt.StringEqualStringPtr(normalized, name) {
				self.responseHeadersSameKey.__hx_this.set(normalized, values)
			}
		}
	}
}

func (self *sys__Http) resetResponseHeaders() {
	values := New_haxe__ds__StringMap()
	self.responseHeaders = values
	self.responseHeadersSameKey = New_haxe__ds__StringMap()
}

func (self *sys__Http) hasHeader(name *string) bool {
	_g := 0
	_g1 := self.headers
	for _g < _g1.Len() {
		header := func(hx_value_108 any) map[string]any {
			if hx_value_108 == nil {
				var hx_zero_109 map[string]any
				return hx_zero_109
			}
			return hx_value_108.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(hxrt.StringToLowerCaseStringPtr(func(hx_obj_110 map[string]any) *string {
			hx_field_111 := hx_obj_110["name"]
			if hx_field_111 == nil {
				var hx_zero_112 *string
				return hx_zero_112
			}
			return hx_field_111.(*string)
		}(header)), hxrt.StringToLowerCaseStringPtr(name)) {
			return true
		}
	}
	return false
}

func (self *sys__Http) buildMultipartBody(upload map[string]any) *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	_g := 0
	_g1 := self.params
	for _g < _g1.Len() {
		parameter := func(hx_value_113 any) map[string]any {
			if hx_value_113 == nil {
				var hx_zero_114 map[string]any
				return hx_zero_114
			}
			return hx_value_113.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("--hxrt-go-boundary\r\n"))
		x := hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Content-Disposition: form-data; name=\""), func(hx_obj_115 map[string]any) *string {
			hx_field_116 := hx_obj_115["name"]
			if hx_field_116 == nil {
				var hx_zero_117 *string
				return hx_zero_117
			}
			return hx_field_116.(*string)
		}(parameter)), hxrt.StringFromLiteral("\"\r\n\r\n"))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		x_1 := func(hx_obj_118 map[string]any) *string {
			hx_field_119 := hx_obj_118["value"]
			if hx_field_119 == nil {
				var hx_zero_120 *string
				return hx_zero_120
			}
			return hx_field_119.(*string)
		}(parameter)
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x_1))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("\r\n"))
	}
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("--hxrt-go-boundary\r\n"))
	x_2 := hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Content-Disposition: form-data; name=\""), func(hx_obj_121 map[string]any) *string {
		hx_field_122 := hx_obj_121["param"]
		if hx_field_122 == nil {
			var hx_zero_123 *string
			return hx_zero_123
		}
		return hx_field_122.(*string)
	}(upload)), hxrt.StringFromLiteral("\"; filename=\"")), func(hx_obj_124 map[string]any) *string {
		hx_field_125 := hx_obj_124["filename"]
		if hx_field_125 == nil {
			var hx_zero_126 *string
			return hx_zero_126
		}
		return hx_field_125.(*string)
	}(upload)), hxrt.StringFromLiteral("\"\r\n"))
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x_2))
	x_3 := hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Content-Type: "), func(hx_obj_127 map[string]any) *string {
		hx_field_128 := hx_obj_127["mimeType"]
		if hx_field_128 == nil {
			var hx_zero_129 *string
			return hx_zero_129
		}
		return hx_field_128.(*string)
	}(upload)), hxrt.StringFromLiteral("\r\n\r\n"))
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x_3))
	x_4 := hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("[uploaded-bytes="), func(hx_obj_130 map[string]any) int {
		hx_field_131 := hx_obj_130["size"]
		if hx_field_131 == nil {
			var hx_zero_132 int
			return hx_zero_132
		}
		return hx_field_131.(int)
	}(upload)), hxrt.StringFromLiteral("]\r\n"))
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x_4))
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("--hxrt-go-boundary--\r\n"))
	return out_b
}

func (self *sys__Http) encodedParameters() *string {
	byName := New_haxe__ds__StringMap()
	_g := 0
	_g1 := self.params
	for _g < _g1.Len() {
		parameter := func(hx_value_133 any) map[string]any {
			if hx_value_133 == nil {
				var hx_zero_134 map[string]any
				return hx_zero_134
			}
			return hx_value_133.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		byName.__hx_this.set(func(hx_obj_135 map[string]any) *string {
			hx_field_136 := hx_obj_135["name"]
			if hx_field_136 == nil {
				var hx_zero_137 *string
				return hx_zero_137
			}
			return hx_field_136.(*string)
		}(parameter), func(hx_obj_138 map[string]any) *string {
			hx_field_139 := hx_obj_138["value"]
			if hx_field_139 == nil {
				var hx_zero_140 *string
				return hx_zero_140
			}
			return hx_field_139.(*string)
		}(parameter))
	}
	_g_1 := hxrt.NewArray()
	name := func(hx_value_141 any) map[string]any {
		if hx_value_141 == nil {
			var hx_zero_142 map[string]any
			return hx_zero_142
		}
		return hx_value_141.(map[string]any)
	}(byName.__hx_this.keys())
	for func(hx_obj_143 map[string]any) func() bool {
		hx_field_144 := hx_obj_143["hasNext"]
		if hx_field_144 == nil {
			var hx_zero_145 func() bool
			return hx_zero_145
		}
		return hx_field_144.(func() bool)
	}(name)() {
		name_1 := func(hx_obj_146 map[string]any) func() *string {
			hx_field_147 := hx_obj_146["next"]
			if hx_field_147 == nil {
				var hx_zero_148 func() *string
				return hx_zero_148
			}
			return hx_field_147.(func() *string)
		}(name)()
		_g_1.Push(name_1)
	}
	names := _g_1
	encoded := hxrt.NewArray()
	emitted := New_haxe__ds__StringMap()
	_g_2 := 0
	for _g_2 < names.Len() {
		hx_tmp := func(hx_value_150 any) *string {
			if hx_value_150 == nil {
				var hx_zero_151 *string
				return hx_zero_151
			}
			return hx_value_150.(*string)
		}(names.Get(_g_2))
		_ = hx_tmp
		_g_2 = int(int32((_g_2 + 1)))
		next := -1
		_g_3 := 0
		_g1_1 := names.Len()
		for _g_3 < _g1_1 {
			hx_post_152 := _g_3
			_g_3 = int(int32((_g_3 + 1)))
			index := hx_post_152
			if !func(hx_value_155 any) bool {
				if hx_value_155 == nil {
					var hx_zero_156 bool
					return hx_zero_156
				}
				return hx_value_155.(bool)
			}(emitted.__hx_this.exists(func(hx_value_153 any) *string {
				if hx_value_153 == nil {
					var hx_zero_154 *string
					return hx_zero_154
				}
				return hx_value_153.(*string)
			}(names.Get(index)))) && ((next < 0) || (Reflect_compare(func(hx_value_157 any) *string {
				if hx_value_157 == nil {
					var hx_zero_158 *string
					return hx_zero_158
				}
				return hx_value_157.(*string)
			}(names.Get(index)), func(hx_value_159 any) *string {
				if hx_value_159 == nil {
					var hx_zero_160 *string
					return hx_zero_160
				}
				return hx_value_159.(*string)
			}(names.Get(next))) < 0)) {
				next = index
			}
		}
		if next >= 0 {
			name_2 := func(hx_value_161 any) *string {
				if hx_value_161 == nil {
					var hx_zero_162 *string
					return hx_zero_162
				}
				return hx_value_161.(*string)
			}(names.Get(next))
			emitted.__hx_this.set(name_2, true)
			encoded.Push(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(StringTools_urlEncode(name_2), hxrt.StringFromLiteral("=")), StringTools_urlEncode(func(hx_value_164 any) *string {
				if hx_value_164 == nil {
					var hx_zero_165 *string
					return hx_zero_165
				}
				return hx_value_164.(*string)
			}(byName.__hx_this.get(name_2)))))
		}
	}
	return hxrt.StringJoinAny(encoded.Values(), hxrt.StringFromLiteral("&"))
}

var sys__Http_PROXY map[string]any = nil

func sys__Http_capture(api *haxe__io__Output, payload *haxe__io__Bytes) {
	if api != nil {
		api.__hx_this.writeFullBytes(payload, 0, payload.length)
	}
}

func sys__Http_firstComma(value *string) int {
	_g := 0
	_g1 := hxrt.StringLengthStringPtr(value)
	for _g < _g1 {
		hx_post_166 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_166
		if func() int {
			var c any = hxrt.StringCharCodeAtAnyStringPtr(value, index)
			var hx_if_167 int
			if c == nil {
				hx_if_167 = -1
			} else {
				hx_if_167 = c.(int)
			}
			return hx_if_167
		}() == 44 {
			return index
		}
	}
	return -1
}

func sys__Http_hxrt_proxyDescriptor() *string {
	proxy := sys__Http_PROXY
	if (proxy == nil) || hxrt.StringEqualStringPtr(func(hx_obj_168 map[string]any) *string {
		hx_field_169 := hx_obj_168["host"]
		if hx_field_169 == nil {
			var hx_zero_170 *string
			return hx_zero_170
		}
		return hx_field_169.(*string)
	}(proxy), nil) {
		return hxrt.StringFromLiteral("null")
	}
	var hx_if_180 *string
	if func(hx_obj_171 map[string]any) map[string]any {
		hx_field_172 := hx_obj_171["auth"]
		if hx_field_172 == nil {
			var hx_zero_173 map[string]any
			return hx_zero_173
		}
		return hx_field_172.(map[string]any)
	}(proxy) == nil {
		hx_if_180 = nil
	} else {
		hx_if_180 = func(hx_obj_177 map[string]any) *string {
			hx_field_178 := hx_obj_177["user"]
			if hx_field_178 == nil {
				var hx_zero_179 *string
				return hx_zero_179
			}
			return hx_field_178.(*string)
		}(func(hx_obj_174 map[string]any) map[string]any {
			hx_field_175 := hx_obj_174["auth"]
			if hx_field_175 == nil {
				var hx_zero_176 map[string]any
				return hx_zero_176
			}
			return hx_field_175.(map[string]any)
		}(proxy))
	}
	user := hx_if_180
	var hx_if_190 *string
	if func(hx_obj_181 map[string]any) map[string]any {
		hx_field_182 := hx_obj_181["auth"]
		if hx_field_182 == nil {
			var hx_zero_183 map[string]any
			return hx_zero_183
		}
		return hx_field_182.(map[string]any)
	}(proxy) == nil {
		hx_if_190 = nil
	} else {
		hx_if_190 = func(hx_obj_187 map[string]any) *string {
			hx_field_188 := hx_obj_187["pass"]
			if hx_field_188 == nil {
				var hx_zero_189 *string
				return hx_zero_189
			}
			return hx_field_188.(*string)
		}(func(hx_obj_184 map[string]any) map[string]any {
			hx_field_185 := hx_obj_184["auth"]
			if hx_field_185 == nil {
				var hx_zero_186 map[string]any
				return hx_zero_186
			}
			return hx_field_185.(map[string]any)
		}(proxy))
	}
	pass := hx_if_190
	return hxrt.StdString(hxrt.HttpProxyDescriptor(func(hx_obj_191 map[string]any) *string {
		hx_field_192 := hx_obj_191["host"]
		if hx_field_192 == nil {
			var hx_zero_193 *string
			return hx_zero_193
		}
		return hx_field_192.(*string)
	}(proxy), func(hx_obj_194 map[string]any) int {
		hx_field_195 := hx_obj_194["port"]
		if hx_field_195 == nil {
			var hx_zero_196 int
			return hx_zero_196
		}
		return hx_field_195.(int)
	}(proxy), user, pass))
}

func sys__Http_normalizedMethod(method *string) *string {
	if hxrt.StringEqualStringPtr(method, nil) {
		return nil
	}
	normalized := hxrt.StringToUpperCaseStringPtr(method)
	var hx_if_197 *string
	if hxrt.StringEqualStringPtr(normalized, hxrt.StringFromLiteral("")) || hxrt.StringEqualStringPtr(normalized, hxrt.StringFromLiteral("NULL")) {
		hx_if_197 = nil
	} else {
		hx_if_197 = normalized
	}
	return hx_if_197
}

func sys__Http_requestUrl(url *string) *string {
	request := New_sys__Http(url)
	result := hxrt.StringFromLiteral("")
	request.onData = func(data *string) {
		result = data
	}
	request.onError = func(message *string) {
		result = message
	}
	request.__hx_this.request(false)
	return result
}
