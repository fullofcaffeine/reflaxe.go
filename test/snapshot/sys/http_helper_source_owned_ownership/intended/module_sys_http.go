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
	requestWith(post bool, api *haxe__io__Output, sock *sys__net__Socket, method *string, deliverSuccessCallbacks bool)
	handleDataRequest(post bool, api *haxe__io__Output, method *string, deliverSuccessCallbacks bool)
	recordResponseHeaders(response *hxrt.HttpResponse)
	resetResponseHeaders()
	hasHeader(name *string) bool
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
	self.__hx_this.requestWith(usePost, nil, nil, nil, true)
}

func (self *sys__Http) fileTransfert(argname *string, filename *string, file *haxe__io__Input, size int, mimeType *string) {
	self.__hx_this.fileTransfer(argname, filename, file, size, mimeType)
}

func (self *sys__Http) fileTransfer(argname *string, filename *string, file *haxe__io__Input, size int, mimeType *string) {
	hx_obj_44 := map[string]any{}
	hx_obj_44["param"] = argname
	hx_obj_44["filename"] = filename
	hx_obj_44["io"] = file
	hx_obj_44["size"] = size
	hx_obj_44["mimeType"] = mimeType
	self.file = hx_obj_44
}

func (self *sys__Http) customRequest(post bool, api *haxe__io__Output, sock *sys__net__Socket, method *string) {
	self.__hx_this.requestWith((post || (self.file != nil)), api, sock, method, false)
}

func (self *sys__Http) getResponseHeaderValues(key *string) *hxrt.Array {
	values := func(hx_value_45 any) *hxrt.Array {
		if hx_value_45 == nil {
			var hx_zero_46 *hxrt.Array
			return hx_zero_46
		}
		return hx_value_45.(*hxrt.Array)
	}(self.responseHeadersSameKey.__hx_this.get(key))
	if values == nil {
		normalized := hxrt.StringToLowerCaseStringPtr(key)
		if !hxrt.StringEqualStringPtr(normalized, key) {
			values = func(hx_value_47 any) *hxrt.Array {
				if hx_value_47 == nil {
					var hx_zero_48 *hxrt.Array
					return hx_zero_48
				}
				return hx_value_47.(*hxrt.Array)
			}(self.responseHeadersSameKey.__hx_this.get(normalized))
		}
	}
	if values != nil {
		return values
	}
	var this1 haxe__IMap = self.responseHeaders
	value := func(hx_value_49 any) *string {
		if hx_value_49 == nil {
			var hx_zero_50 *string
			return hx_zero_50
		}
		return hx_value_49.(*string)
	}(this1.(*haxe__ds__StringMap).__hx_this.get(key))
	if hxrt.StringEqualStringPtr(value, nil) {
		normalized_1 := hxrt.StringToLowerCaseStringPtr(key)
		if !hxrt.StringEqualStringPtr(normalized_1, key) {
			var this1_1 haxe__IMap = self.responseHeaders
			value = func(hx_value_51 any) *string {
				if hx_value_51 == nil {
					var hx_zero_52 *string
					return hx_zero_52
				}
				return hx_value_51.(*string)
			}(this1_1.(*haxe__ds__StringMap).__hx_this.get(normalized_1))
		}
	}
	var hx_if_53 *hxrt.Array
	if hxrt.StringEqualStringPtr(value, nil) {
		hx_if_53 = nil
	} else {
		hx_if_53 = hxrt.NewArray(value)
	}
	return hx_if_53
}

func (self *sys__Http) requestWith(post bool, api *haxe__io__Output, sock *sys__net__Socket, method *string, deliverSuccessCallbacks bool) {
	self.responseAsString = nil
	self.responseBytes = nil
	self.__hx_this.resetResponseHeaders()
	if StringTools_startsWith(self.url, hxrt.StringFromLiteral("data:")) {
		self.__hx_this.handleDataRequest(post, api, method, deliverSuccessCallbacks)
		return
	}
	request := hxrt.HttpRequestNew(self.url, post, method, self.cnxTimeout)
	_g := 0
	_g1 := self.params
	for _g < _g1.Len() {
		parameter := func(hx_value_54 any) map[string]any {
			if hx_value_54 == nil {
				var hx_zero_55 map[string]any
				return hx_zero_55
			}
			return hx_value_54.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		hxrt.HttpRequestAddParameter(request, func(hx_obj_56 map[string]any) *string {
			hx_field_57 := hx_obj_56["name"]
			if hx_field_57 == nil {
				var hx_zero_58 *string
				return hx_zero_58
			}
			return hx_field_57.(*string)
		}(parameter), func(hx_obj_59 map[string]any) *string {
			hx_field_60 := hx_obj_59["value"]
			if hx_field_60 == nil {
				var hx_zero_61 *string
				return hx_zero_61
			}
			return hx_field_60.(*string)
		}(parameter))
	}
	_g_1 := 0
	_g1_1 := self.headers
	for _g_1 < _g1_1.Len() {
		header := func(hx_value_62 any) map[string]any {
			if hx_value_62 == nil {
				var hx_zero_63 map[string]any
				return hx_zero_63
			}
			return hx_value_62.(map[string]any)
		}(_g1_1.Get(_g_1))
		_g_1 = int(int32((_g_1 + 1)))
		hxrt.HttpRequestAddHeader(request, func(hx_obj_64 map[string]any) *string {
			hx_field_65 := hx_obj_64["name"]
			if hx_field_65 == nil {
				var hx_zero_66 *string
				return hx_zero_66
			}
			return hx_field_65.(*string)
		}(header), func(hx_obj_67 map[string]any) *string {
			hx_field_68 := hx_obj_67["value"]
			if hx_field_68 == nil {
				var hx_zero_69 *string
				return hx_zero_69
			}
			return hx_field_68.(*string)
		}(header))
	}
	var uploadError *string = nil
	upload := self.file
	if upload != nil {
		hxrt.HttpRequestSetMultipartUpload(request, func(hx_obj_70 map[string]any) *string {
			hx_field_71 := hx_obj_70["param"]
			if hx_field_71 == nil {
				var hx_zero_72 *string
				return hx_zero_72
			}
			return hx_field_71.(*string)
		}(upload), func(hx_obj_73 map[string]any) *string {
			hx_field_74 := hx_obj_73["filename"]
			if hx_field_74 == nil {
				var hx_zero_75 *string
				return hx_zero_75
			}
			return hx_field_74.(*string)
		}(upload), func(hx_obj_76 map[string]any) *string {
			hx_field_77 := hx_obj_76["mimeType"]
			if hx_field_77 == nil {
				var hx_zero_78 *string
				return hx_zero_78
			}
			return hx_field_77.(*string)
		}(upload), func(hx_obj_79 map[string]any) int {
			hx_field_80 := hx_obj_79["size"]
			if hx_field_80 == nil {
				var hx_zero_81 int
				return hx_zero_81
			}
			return hx_field_80.(int)
		}(upload), func(requested int) *hxrt.ByteView {
			var result *hxrt.ByteView = nil
			if requested > 0 {
				chunk := haxe__io__Bytes_alloc(requested)
				hxrt.TryCatch(func() {
					count := func(hx_obj_84 map[string]any) *haxe__io__Input {
						hx_field_85 := hx_obj_84["io"]
						if hx_field_85 == nil {
							var hx_zero_86 *haxe__io__Input
							return hx_zero_86
						}
						return hx_field_85.(*haxe__io__Input)
					}(upload).__hx_this.readBytes(chunk, 0, requested)
					if count > 0 {
						if count < requested {
							chunk = chunk.__hx_this.sub(0, count)
						}
						result = chunk.__hx_this.__hx_nativeView()
					}
				}, func(hx_caught_82 any) {
					switch hx_typed_83 := hx_caught_82.(type) {
					case *haxe__io__Eof:
						hx_tmp := hx_typed_83
						_ = hx_tmp
					default:
						error := hxrt.ExceptionCaught(hx_caught_82)
						uploadError = hxrt.ExceptionMessage(error)
					}
				})
			}
			return result
		})
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
	if (proxy != nil) && !hxrt.StringEqualStringPtr(func(hx_obj_113 map[string]any) *string {
		hx_field_114 := hx_obj_113["host"]
		if hx_field_114 == nil {
			var hx_zero_115 *string
			return hx_zero_115
		}
		return hx_field_114.(*string)
	}(proxy), nil) {
		var hx_if_96 *string
		if func(hx_obj_87 map[string]any) map[string]any {
			hx_field_88 := hx_obj_87["auth"]
			if hx_field_88 == nil {
				var hx_zero_89 map[string]any
				return hx_zero_89
			}
			return hx_field_88.(map[string]any)
		}(proxy) == nil {
			hx_if_96 = nil
		} else {
			hx_if_96 = func(hx_obj_93 map[string]any) *string {
				hx_field_94 := hx_obj_93["user"]
				if hx_field_94 == nil {
					var hx_zero_95 *string
					return hx_zero_95
				}
				return hx_field_94.(*string)
			}(func(hx_obj_90 map[string]any) map[string]any {
				hx_field_91 := hx_obj_90["auth"]
				if hx_field_91 == nil {
					var hx_zero_92 map[string]any
					return hx_zero_92
				}
				return hx_field_91.(map[string]any)
			}(proxy))
		}
		user := hx_if_96
		var hx_if_106 *string
		if func(hx_obj_97 map[string]any) map[string]any {
			hx_field_98 := hx_obj_97["auth"]
			if hx_field_98 == nil {
				var hx_zero_99 map[string]any
				return hx_zero_99
			}
			return hx_field_98.(map[string]any)
		}(proxy) == nil {
			hx_if_106 = nil
		} else {
			hx_if_106 = func(hx_obj_103 map[string]any) *string {
				hx_field_104 := hx_obj_103["pass"]
				if hx_field_104 == nil {
					var hx_zero_105 *string
					return hx_zero_105
				}
				return hx_field_104.(*string)
			}(func(hx_obj_100 map[string]any) map[string]any {
				hx_field_101 := hx_obj_100["auth"]
				if hx_field_101 == nil {
					var hx_zero_102 map[string]any
					return hx_zero_102
				}
				return hx_field_101.(map[string]any)
			}(proxy))
		}
		pass := hx_if_106
		hxrt.HttpRequestSetProxy(request, func(hx_obj_107 map[string]any) *string {
			hx_field_108 := hx_obj_107["host"]
			if hx_field_108 == nil {
				var hx_zero_109 *string
				return hx_zero_109
			}
			return hx_field_108.(*string)
		}(proxy), func(hx_obj_110 map[string]any) int {
			hx_field_111 := hx_obj_110["port"]
			if hx_field_111 == nil {
				var hx_zero_112 int
				return hx_zero_112
			}
			return hx_field_111.(int)
		}(proxy), user, pass)
	}
	if sock != nil {
		hxrt.HttpRequestSetSocket(request, sock.handle)
	}
	response := hxrt.HttpRequestExecute(request)
	if !hxrt.StringEqualStringPtr(uploadError, nil) {
		func(hx_fn func(*string), hx_arg_0 *string) {
			if hx_fn == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
				return
			}
			hx_fn(hx_arg_0)
		}(self.onError, uploadError)
		return
	}
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
	payloadText := payload.__hx_this.toString()
	if deliverSuccessCallbacks {
		self.responseBytes = payload
		self.responseAsString = payloadText
	}
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
	if api != nil {
		api.__hx_this.close()
	}
	if deliverSuccessCallbacks {
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
}

func (self *sys__Http) handleDataRequest(post bool, api *haxe__io__Output, method *string, deliverSuccessCallbacks bool) {
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
			encoded = hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("multipart file="), func(hx_obj_116 map[string]any) *string {
				hx_field_117 := hx_obj_116["filename"]
				if hx_field_117 == nil {
					var hx_zero_118 *string
					return hx_zero_118
				}
				return hx_field_117.(*string)
			}(self.file)), hxrt.StringFromLiteral(";mime=")), func(hx_obj_119 map[string]any) *string {
				hx_field_120 := hx_obj_119["mimeType"]
				if hx_field_120 == nil {
					var hx_zero_121 *string
					return hx_zero_121
				}
				return hx_field_120.(*string)
			}(self.file)), hxrt.StringFromLiteral(";size=")), func(hx_obj_122 map[string]any) int {
				hx_field_123 := hx_obj_122["size"]
				if hx_field_123 == nil {
					var hx_zero_124 int
					return hx_zero_124
				}
				return hx_field_123.(int)
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
	if deliverSuccessCallbacks {
		self.responseBytes = payload
		self.responseAsString = payloadText
	}
	var this1 haxe__IMap = self.responseHeaders
	this1.(*haxe__ds__StringMap).__hx_this.set(hxrt.StringFromLiteral("content-type"), mediaType)
	var this1_1 haxe__IMap = self.responseHeaders
	this1_1.(*haxe__ds__StringMap).__hx_this.set(hxrt.StringFromLiteral("Content-Type"), mediaType)
	sys__Http_capture(api, payload)
	if api != nil {
		api.__hx_this.close()
	}
	func(hx_fn func(int), hx_arg_0 int) {
		if hx_fn == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
			return
		}
		hx_fn(hx_arg_0)
	}(self.onStatus, 200)
	if deliverSuccessCallbacks {
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
}

func (self *sys__Http) recordResponseHeaders(response *hxrt.HttpResponse) {
	count := hxrt.HttpResponseHeaderCount(response)
	_g := 0
	_g1 := count
	for _g < _g1 {
		hx_post_125 := _g
		_g = int(int32((_g + 1)))
		headerIndex := hx_post_125
		name := hxrt.StdString(hxrt.HttpResponseHeaderName(response, headerIndex))
		normalized := hxrt.StringToLowerCaseStringPtr(name)
		valueCount := hxrt.HttpResponseHeaderValueCount(response, headerIndex)
		values := hxrt.NewArray()
		_g_1 := 0
		_g1_1 := valueCount
		for _g_1 < _g1_1 {
			hx_post_126 := _g_1
			_g_1 = int(int32((_g_1 + 1)))
			valueIndex := hx_post_126
			values.Push(hxrt.StdString(hxrt.HttpResponseHeaderValue(response, headerIndex, valueIndex)))
		}
		if values.Len() == 0 {
			continue
		}
		last := func(hx_value_128 any) *string {
			if hx_value_128 == nil {
				var hx_zero_129 *string
				return hx_zero_129
			}
			return hx_value_128.(*string)
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
		header := func(hx_value_130 any) map[string]any {
			if hx_value_130 == nil {
				var hx_zero_131 map[string]any
				return hx_zero_131
			}
			return hx_value_130.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(hxrt.StringToLowerCaseStringPtr(func(hx_obj_132 map[string]any) *string {
			hx_field_133 := hx_obj_132["name"]
			if hx_field_133 == nil {
				var hx_zero_134 *string
				return hx_zero_134
			}
			return hx_field_133.(*string)
		}(header)), hxrt.StringToLowerCaseStringPtr(name)) {
			return true
		}
	}
	return false
}

func (self *sys__Http) encodedParameters() *string {
	byName := New_haxe__ds__StringMap()
	_g := 0
	_g1 := self.params
	for _g < _g1.Len() {
		parameter := func(hx_value_135 any) map[string]any {
			if hx_value_135 == nil {
				var hx_zero_136 map[string]any
				return hx_zero_136
			}
			return hx_value_135.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		byName.__hx_this.set(func(hx_obj_137 map[string]any) *string {
			hx_field_138 := hx_obj_137["name"]
			if hx_field_138 == nil {
				var hx_zero_139 *string
				return hx_zero_139
			}
			return hx_field_138.(*string)
		}(parameter), func(hx_obj_140 map[string]any) *string {
			hx_field_141 := hx_obj_140["value"]
			if hx_field_141 == nil {
				var hx_zero_142 *string
				return hx_zero_142
			}
			return hx_field_141.(*string)
		}(parameter))
	}
	_g_1 := hxrt.NewArray()
	name := func(hx_value_143 any) map[string]any {
		if hx_value_143 == nil {
			var hx_zero_144 map[string]any
			return hx_zero_144
		}
		return hx_value_143.(map[string]any)
	}(byName.__hx_this.keys())
	for func(hx_obj_145 map[string]any) func() bool {
		hx_field_146 := hx_obj_145["hasNext"]
		if hx_field_146 == nil {
			var hx_zero_147 func() bool
			return hx_zero_147
		}
		return hx_field_146.(func() bool)
	}(name)() {
		name_1 := func(hx_obj_148 map[string]any) func() *string {
			hx_field_149 := hx_obj_148["next"]
			if hx_field_149 == nil {
				var hx_zero_150 func() *string
				return hx_zero_150
			}
			return hx_field_149.(func() *string)
		}(name)()
		_g_1.Push(name_1)
	}
	names := _g_1
	encoded := hxrt.NewArray()
	emitted := New_haxe__ds__StringMap()
	_g_2 := 0
	for _g_2 < names.Len() {
		hx_tmp := func(hx_value_152 any) *string {
			if hx_value_152 == nil {
				var hx_zero_153 *string
				return hx_zero_153
			}
			return hx_value_152.(*string)
		}(names.Get(_g_2))
		_ = hx_tmp
		_g_2 = int(int32((_g_2 + 1)))
		next := -1
		_g_3 := 0
		_g1_1 := names.Len()
		for _g_3 < _g1_1 {
			hx_post_154 := _g_3
			_g_3 = int(int32((_g_3 + 1)))
			index := hx_post_154
			if !func(hx_value_157 any) bool {
				if hx_value_157 == nil {
					var hx_zero_158 bool
					return hx_zero_158
				}
				return hx_value_157.(bool)
			}(emitted.__hx_this.exists(func(hx_value_155 any) *string {
				if hx_value_155 == nil {
					var hx_zero_156 *string
					return hx_zero_156
				}
				return hx_value_155.(*string)
			}(names.Get(index)))) && ((next < 0) || (Reflect_compare(func(hx_value_159 any) *string {
				if hx_value_159 == nil {
					var hx_zero_160 *string
					return hx_zero_160
				}
				return hx_value_159.(*string)
			}(names.Get(index)), func(hx_value_161 any) *string {
				if hx_value_161 == nil {
					var hx_zero_162 *string
					return hx_zero_162
				}
				return hx_value_161.(*string)
			}(names.Get(next))) < 0)) {
				next = index
			}
		}
		if next >= 0 {
			name_2 := func(hx_value_163 any) *string {
				if hx_value_163 == nil {
					var hx_zero_164 *string
					return hx_zero_164
				}
				return hx_value_163.(*string)
			}(names.Get(next))
			emitted.__hx_this.set(name_2, true)
			encoded.Push(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(StringTools_urlEncode(name_2), hxrt.StringFromLiteral("=")), StringTools_urlEncode(func(hx_value_166 any) *string {
				if hx_value_166 == nil {
					var hx_zero_167 *string
					return hx_zero_167
				}
				return hx_value_166.(*string)
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
		hx_post_168 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_168
		if func() int {
			var c any = hxrt.StringCharCodeAtAnyStringPtr(value, index)
			var hx_if_169 int
			if c == nil {
				hx_if_169 = -1
			} else {
				hx_if_169 = c.(int)
			}
			return hx_if_169
		}() == 44 {
			return index
		}
	}
	return -1
}

func sys__Http_hxrt_proxyDescriptor() *string {
	proxy := sys__Http_PROXY
	if (proxy == nil) || hxrt.StringEqualStringPtr(func(hx_obj_170 map[string]any) *string {
		hx_field_171 := hx_obj_170["host"]
		if hx_field_171 == nil {
			var hx_zero_172 *string
			return hx_zero_172
		}
		return hx_field_171.(*string)
	}(proxy), nil) {
		return hxrt.StringFromLiteral("null")
	}
	var hx_if_182 *string
	if func(hx_obj_173 map[string]any) map[string]any {
		hx_field_174 := hx_obj_173["auth"]
		if hx_field_174 == nil {
			var hx_zero_175 map[string]any
			return hx_zero_175
		}
		return hx_field_174.(map[string]any)
	}(proxy) == nil {
		hx_if_182 = nil
	} else {
		hx_if_182 = func(hx_obj_179 map[string]any) *string {
			hx_field_180 := hx_obj_179["user"]
			if hx_field_180 == nil {
				var hx_zero_181 *string
				return hx_zero_181
			}
			return hx_field_180.(*string)
		}(func(hx_obj_176 map[string]any) map[string]any {
			hx_field_177 := hx_obj_176["auth"]
			if hx_field_177 == nil {
				var hx_zero_178 map[string]any
				return hx_zero_178
			}
			return hx_field_177.(map[string]any)
		}(proxy))
	}
	user := hx_if_182
	var hx_if_192 *string
	if func(hx_obj_183 map[string]any) map[string]any {
		hx_field_184 := hx_obj_183["auth"]
		if hx_field_184 == nil {
			var hx_zero_185 map[string]any
			return hx_zero_185
		}
		return hx_field_184.(map[string]any)
	}(proxy) == nil {
		hx_if_192 = nil
	} else {
		hx_if_192 = func(hx_obj_189 map[string]any) *string {
			hx_field_190 := hx_obj_189["pass"]
			if hx_field_190 == nil {
				var hx_zero_191 *string
				return hx_zero_191
			}
			return hx_field_190.(*string)
		}(func(hx_obj_186 map[string]any) map[string]any {
			hx_field_187 := hx_obj_186["auth"]
			if hx_field_187 == nil {
				var hx_zero_188 map[string]any
				return hx_zero_188
			}
			return hx_field_187.(map[string]any)
		}(proxy))
	}
	pass := hx_if_192
	return hxrt.StdString(hxrt.HttpProxyDescriptor(func(hx_obj_193 map[string]any) *string {
		hx_field_194 := hx_obj_193["host"]
		if hx_field_194 == nil {
			var hx_zero_195 *string
			return hx_zero_195
		}
		return hx_field_194.(*string)
	}(proxy), func(hx_obj_196 map[string]any) int {
		hx_field_197 := hx_obj_196["port"]
		if hx_field_197 == nil {
			var hx_zero_198 int
			return hx_zero_198
		}
		return hx_field_197.(int)
	}(proxy), user, pass))
}

func sys__Http_normalizedMethod(method *string) *string {
	if hxrt.StringEqualStringPtr(method, nil) {
		return nil
	}
	normalized := hxrt.StringToUpperCaseStringPtr(method)
	var hx_if_199 *string
	if hxrt.StringEqualStringPtr(normalized, hxrt.StringFromLiteral("")) || hxrt.StringEqualStringPtr(normalized, hxrt.StringFromLiteral("NULL")) {
		hx_if_199 = nil
	} else {
		hx_if_199 = normalized
	}
	return hx_if_199
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
