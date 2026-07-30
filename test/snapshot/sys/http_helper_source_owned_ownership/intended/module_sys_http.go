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
		}(parameter), StringTools_urlEncode(func(hx_obj_62 map[string]any) *string {
			hx_field_63 := hx_obj_62["name"]
			if hx_field_63 == nil {
				var hx_zero_64 *string
				return hx_zero_64
			}
			return hx_field_63.(*string)
		}(parameter)), StringTools_urlEncode(func(hx_obj_65 map[string]any) *string {
			hx_field_66 := hx_obj_65["value"]
			if hx_field_66 == nil {
				var hx_zero_67 *string
				return hx_zero_67
			}
			return hx_field_66.(*string)
		}(parameter)))
	}
	_g_1 := 0
	_g1_1 := self.headers
	for _g_1 < _g1_1.Len() {
		header := func(hx_value_68 any) map[string]any {
			if hx_value_68 == nil {
				var hx_zero_69 map[string]any
				return hx_zero_69
			}
			return hx_value_68.(map[string]any)
		}(_g1_1.Get(_g_1))
		_g_1 = int(int32((_g_1 + 1)))
		hxrt.HttpRequestAddHeader(request, func(hx_obj_70 map[string]any) *string {
			hx_field_71 := hx_obj_70["name"]
			if hx_field_71 == nil {
				var hx_zero_72 *string
				return hx_zero_72
			}
			return hx_field_71.(*string)
		}(header), func(hx_obj_73 map[string]any) *string {
			hx_field_74 := hx_obj_73["value"]
			if hx_field_74 == nil {
				var hx_zero_75 *string
				return hx_zero_75
			}
			return hx_field_74.(*string)
		}(header))
	}
	var uploadError *string = nil
	upload := self.file
	if upload != nil {
		hxrt.HttpRequestSetMultipartUpload(request, func(hx_obj_76 map[string]any) *string {
			hx_field_77 := hx_obj_76["param"]
			if hx_field_77 == nil {
				var hx_zero_78 *string
				return hx_zero_78
			}
			return hx_field_77.(*string)
		}(upload), func(hx_obj_79 map[string]any) *string {
			hx_field_80 := hx_obj_79["filename"]
			if hx_field_80 == nil {
				var hx_zero_81 *string
				return hx_zero_81
			}
			return hx_field_80.(*string)
		}(upload), func(hx_obj_82 map[string]any) *string {
			hx_field_83 := hx_obj_82["mimeType"]
			if hx_field_83 == nil {
				var hx_zero_84 *string
				return hx_zero_84
			}
			return hx_field_83.(*string)
		}(upload), func(hx_obj_85 map[string]any) int {
			hx_field_86 := hx_obj_85["size"]
			if hx_field_86 == nil {
				var hx_zero_87 int
				return hx_zero_87
			}
			return hx_field_86.(int)
		}(upload), func(requested int) *hxrt.ByteView {
			var result *hxrt.ByteView = nil
			if requested > 0 {
				chunk := haxe__io__Bytes_alloc(requested)
				hxrt.TryCatch(func() {
					count := func(hx_obj_90 map[string]any) *haxe__io__Input {
						hx_field_91 := hx_obj_90["io"]
						if hx_field_91 == nil {
							var hx_zero_92 *haxe__io__Input
							return hx_zero_92
						}
						return hx_field_91.(*haxe__io__Input)
					}(upload).__hx_this.readBytes(chunk, 0, requested)
					if count > 0 {
						if count < requested {
							chunk = chunk.__hx_this.sub(0, count)
						}
						result = chunk.__hx_this.__hx_nativeView()
					}
				}, func(hx_caught_88 any) {
					switch hx_typed_89 := hx_caught_88.(type) {
					case *haxe__io__Eof:
						hx_tmp := hx_typed_89
						_ = hx_tmp
					default:
						error := hxrt.ExceptionCaught(hx_caught_88)
						uploadError = hxrt.ExceptionMessage(error)
					}
				})
			}
			return result
		})
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
	if (proxy != nil) && !hxrt.StringEqualStringPtr(func(hx_obj_119 map[string]any) *string {
		hx_field_120 := hx_obj_119["host"]
		if hx_field_120 == nil {
			var hx_zero_121 *string
			return hx_zero_121
		}
		return hx_field_120.(*string)
	}(proxy), nil) {
		var hx_if_102 *string
		if func(hx_obj_93 map[string]any) map[string]any {
			hx_field_94 := hx_obj_93["auth"]
			if hx_field_94 == nil {
				var hx_zero_95 map[string]any
				return hx_zero_95
			}
			return hx_field_94.(map[string]any)
		}(proxy) == nil {
			hx_if_102 = nil
		} else {
			hx_if_102 = func(hx_obj_99 map[string]any) *string {
				hx_field_100 := hx_obj_99["user"]
				if hx_field_100 == nil {
					var hx_zero_101 *string
					return hx_zero_101
				}
				return hx_field_100.(*string)
			}(func(hx_obj_96 map[string]any) map[string]any {
				hx_field_97 := hx_obj_96["auth"]
				if hx_field_97 == nil {
					var hx_zero_98 map[string]any
					return hx_zero_98
				}
				return hx_field_97.(map[string]any)
			}(proxy))
		}
		user := hx_if_102
		var hx_if_112 *string
		if func(hx_obj_103 map[string]any) map[string]any {
			hx_field_104 := hx_obj_103["auth"]
			if hx_field_104 == nil {
				var hx_zero_105 map[string]any
				return hx_zero_105
			}
			return hx_field_104.(map[string]any)
		}(proxy) == nil {
			hx_if_112 = nil
		} else {
			hx_if_112 = func(hx_obj_109 map[string]any) *string {
				hx_field_110 := hx_obj_109["pass"]
				if hx_field_110 == nil {
					var hx_zero_111 *string
					return hx_zero_111
				}
				return hx_field_110.(*string)
			}(func(hx_obj_106 map[string]any) map[string]any {
				hx_field_107 := hx_obj_106["auth"]
				if hx_field_107 == nil {
					var hx_zero_108 map[string]any
					return hx_zero_108
				}
				return hx_field_107.(map[string]any)
			}(proxy))
		}
		pass := hx_if_112
		hxrt.HttpRequestSetProxy(request, func(hx_obj_113 map[string]any) *string {
			hx_field_114 := hx_obj_113["host"]
			if hx_field_114 == nil {
				var hx_zero_115 *string
				return hx_zero_115
			}
			return hx_field_114.(*string)
		}(proxy), func(hx_obj_116 map[string]any) int {
			hx_field_117 := hx_obj_116["port"]
			if hx_field_117 == nil {
				var hx_zero_118 int
				return hx_zero_118
			}
			return hx_field_117.(int)
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
			encoded = hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("multipart file="), func(hx_obj_122 map[string]any) *string {
				hx_field_123 := hx_obj_122["filename"]
				if hx_field_123 == nil {
					var hx_zero_124 *string
					return hx_zero_124
				}
				return hx_field_123.(*string)
			}(self.file)), hxrt.StringFromLiteral(";mime=")), func(hx_obj_125 map[string]any) *string {
				hx_field_126 := hx_obj_125["mimeType"]
				if hx_field_126 == nil {
					var hx_zero_127 *string
					return hx_zero_127
				}
				return hx_field_126.(*string)
			}(self.file)), hxrt.StringFromLiteral(";size=")), func(hx_obj_128 map[string]any) int {
				hx_field_129 := hx_obj_128["size"]
				if hx_field_129 == nil {
					var hx_zero_130 int
					return hx_zero_130
				}
				return hx_field_129.(int)
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
		hx_post_131 := _g
		_g = int(int32((_g + 1)))
		headerIndex := hx_post_131
		name := hxrt.StdString(hxrt.HttpResponseHeaderName(response, headerIndex))
		normalized := hxrt.StringToLowerCaseStringPtr(name)
		valueCount := hxrt.HttpResponseHeaderValueCount(response, headerIndex)
		values := hxrt.NewArray()
		_g_1 := 0
		_g1_1 := valueCount
		for _g_1 < _g1_1 {
			hx_post_132 := _g_1
			_g_1 = int(int32((_g_1 + 1)))
			valueIndex := hx_post_132
			values.Push(hxrt.StdString(hxrt.HttpResponseHeaderValue(response, headerIndex, valueIndex)))
		}
		if values.Len() == 0 {
			continue
		}
		last := func(hx_value_134 any) *string {
			if hx_value_134 == nil {
				var hx_zero_135 *string
				return hx_zero_135
			}
			return hx_value_134.(*string)
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

func (self *sys__Http) encodedParameters() *string {
	encoded := hxrt.NewArray()
	_g := 0
	_g1 := self.params
	for _g < _g1.Len() {
		parameter := func(hx_value_136 any) map[string]any {
			if hx_value_136 == nil {
				var hx_zero_137 map[string]any
				return hx_zero_137
			}
			return hx_value_136.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		encoded.Push(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(StringTools_urlEncode(func(hx_obj_139 map[string]any) *string {
			hx_field_140 := hx_obj_139["name"]
			if hx_field_140 == nil {
				var hx_zero_141 *string
				return hx_zero_141
			}
			return hx_field_140.(*string)
		}(parameter)), hxrt.StringFromLiteral("=")), StringTools_urlEncode(func(hx_obj_142 map[string]any) *string {
			hx_field_143 := hx_obj_142["value"]
			if hx_field_143 == nil {
				var hx_zero_144 *string
				return hx_zero_144
			}
			return hx_field_143.(*string)
		}(parameter))))
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
		hx_post_145 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_145
		if func() int {
			var c any = hxrt.StringCharCodeAtAnyStringPtr(value, index)
			var hx_if_146 int
			if c == nil {
				hx_if_146 = -1
			} else {
				hx_if_146 = c.(int)
			}
			return hx_if_146
		}() == 44 {
			return index
		}
	}
	return -1
}

func sys__Http_hxrt_proxyDescriptor() *string {
	proxy := sys__Http_PROXY
	if (proxy == nil) || hxrt.StringEqualStringPtr(func(hx_obj_147 map[string]any) *string {
		hx_field_148 := hx_obj_147["host"]
		if hx_field_148 == nil {
			var hx_zero_149 *string
			return hx_zero_149
		}
		return hx_field_148.(*string)
	}(proxy), nil) {
		return hxrt.StringFromLiteral("null")
	}
	var hx_if_159 *string
	if func(hx_obj_150 map[string]any) map[string]any {
		hx_field_151 := hx_obj_150["auth"]
		if hx_field_151 == nil {
			var hx_zero_152 map[string]any
			return hx_zero_152
		}
		return hx_field_151.(map[string]any)
	}(proxy) == nil {
		hx_if_159 = nil
	} else {
		hx_if_159 = func(hx_obj_156 map[string]any) *string {
			hx_field_157 := hx_obj_156["user"]
			if hx_field_157 == nil {
				var hx_zero_158 *string
				return hx_zero_158
			}
			return hx_field_157.(*string)
		}(func(hx_obj_153 map[string]any) map[string]any {
			hx_field_154 := hx_obj_153["auth"]
			if hx_field_154 == nil {
				var hx_zero_155 map[string]any
				return hx_zero_155
			}
			return hx_field_154.(map[string]any)
		}(proxy))
	}
	user := hx_if_159
	var hx_if_169 *string
	if func(hx_obj_160 map[string]any) map[string]any {
		hx_field_161 := hx_obj_160["auth"]
		if hx_field_161 == nil {
			var hx_zero_162 map[string]any
			return hx_zero_162
		}
		return hx_field_161.(map[string]any)
	}(proxy) == nil {
		hx_if_169 = nil
	} else {
		hx_if_169 = func(hx_obj_166 map[string]any) *string {
			hx_field_167 := hx_obj_166["pass"]
			if hx_field_167 == nil {
				var hx_zero_168 *string
				return hx_zero_168
			}
			return hx_field_167.(*string)
		}(func(hx_obj_163 map[string]any) map[string]any {
			hx_field_164 := hx_obj_163["auth"]
			if hx_field_164 == nil {
				var hx_zero_165 map[string]any
				return hx_zero_165
			}
			return hx_field_164.(map[string]any)
		}(proxy))
	}
	pass := hx_if_169
	return hxrt.StdString(hxrt.HttpProxyDescriptor(func(hx_obj_170 map[string]any) *string {
		hx_field_171 := hx_obj_170["host"]
		if hx_field_171 == nil {
			var hx_zero_172 *string
			return hx_zero_172
		}
		return hx_field_171.(*string)
	}(proxy), func(hx_obj_173 map[string]any) int {
		hx_field_174 := hx_obj_173["port"]
		if hx_field_174 == nil {
			var hx_zero_175 int
			return hx_zero_175
		}
		return hx_field_174.(int)
	}(proxy), user, pass))
}

func sys__Http_normalizedMethod(method *string) *string {
	if hxrt.StringEqualStringPtr(method, nil) {
		return nil
	}
	normalized := hxrt.StringToUpperCaseStringPtr(method)
	var hx_if_176 *string
	if hxrt.StringEqualStringPtr(normalized, hxrt.StringFromLiteral("")) || hxrt.StringEqualStringPtr(normalized, hxrt.StringFromLiteral("NULL")) {
		hx_if_176 = nil
	} else {
		hx_if_176 = normalized
	}
	return hx_if_176
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
