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
	hx_obj_45 := map[string]any{}
	hx_obj_45["param"] = argname
	hx_obj_45["filename"] = filename
	hx_obj_45["io"] = file
	hx_obj_45["size"] = size
	hx_obj_45["mimeType"] = mimeType
	self.file = hx_obj_45
}

func (self *sys__Http) customRequest(post bool, api *haxe__io__Output, sock *sys__net__Socket, method *string) {
	self.__hx_this.requestWith((post || (self.file != nil)), api, sock, method, false)
}

func (self *sys__Http) getResponseHeaderValues(key *string) *hxrt.Array {
	values := func(hx_value_46 any) *hxrt.Array {
		if hx_value_46 == nil {
			var hx_zero_47 *hxrt.Array
			return hx_zero_47
		}
		return hx_value_46.(*hxrt.Array)
	}(self.responseHeadersSameKey.__hx_this.get(key))
	if values == nil {
		normalized := hxrt.StringToLowerCaseStringPtr(key)
		if !hxrt.StringEqualStringPtr(normalized, key) {
			values = func(hx_value_48 any) *hxrt.Array {
				if hx_value_48 == nil {
					var hx_zero_49 *hxrt.Array
					return hx_zero_49
				}
				return hx_value_48.(*hxrt.Array)
			}(self.responseHeadersSameKey.__hx_this.get(normalized))
		}
	}
	if values != nil {
		return values
	}
	var this1 haxe__IMap = self.responseHeaders
	value := func(hx_value_50 any) *string {
		if hx_value_50 == nil {
			var hx_zero_51 *string
			return hx_zero_51
		}
		return hx_value_50.(*string)
	}(this1.(*haxe__ds__StringMap).__hx_this.get(key))
	if hxrt.StringEqualStringPtr(value, nil) {
		normalized_1 := hxrt.StringToLowerCaseStringPtr(key)
		if !hxrt.StringEqualStringPtr(normalized_1, key) {
			var this1_1 haxe__IMap = self.responseHeaders
			value = func(hx_value_52 any) *string {
				if hx_value_52 == nil {
					var hx_zero_53 *string
					return hx_zero_53
				}
				return hx_value_52.(*string)
			}(this1_1.(*haxe__ds__StringMap).__hx_this.get(normalized_1))
		}
	}
	var hx_if_54 *hxrt.Array
	if hxrt.StringEqualStringPtr(value, nil) {
		hx_if_54 = nil
	} else {
		hx_if_54 = hxrt.NewArray(value)
	}
	return hx_if_54
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
		parameter := func(hx_value_55 any) map[string]any {
			if hx_value_55 == nil {
				var hx_zero_56 map[string]any
				return hx_zero_56
			}
			return hx_value_55.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		hxrt.HttpRequestAddParameter(request, func(hx_obj_57 map[string]any) *string {
			hx_field_58 := hx_obj_57["name"]
			if hx_field_58 == nil {
				var hx_zero_59 *string
				return hx_zero_59
			}
			return hx_field_58.(*string)
		}(parameter), func(hx_obj_60 map[string]any) *string {
			hx_field_61 := hx_obj_60["value"]
			if hx_field_61 == nil {
				var hx_zero_62 *string
				return hx_zero_62
			}
			return hx_field_61.(*string)
		}(parameter))
	}
	_g_1 := 0
	_g1_1 := self.headers
	for _g_1 < _g1_1.Len() {
		header := func(hx_value_63 any) map[string]any {
			if hx_value_63 == nil {
				var hx_zero_64 map[string]any
				return hx_zero_64
			}
			return hx_value_63.(map[string]any)
		}(_g1_1.Get(_g_1))
		_g_1 = int(int32((_g_1 + 1)))
		hxrt.HttpRequestAddHeader(request, func(hx_obj_65 map[string]any) *string {
			hx_field_66 := hx_obj_65["name"]
			if hx_field_66 == nil {
				var hx_zero_67 *string
				return hx_zero_67
			}
			return hx_field_66.(*string)
		}(header), func(hx_obj_68 map[string]any) *string {
			hx_field_69 := hx_obj_68["value"]
			if hx_field_69 == nil {
				var hx_zero_70 *string
				return hx_zero_70
			}
			return hx_field_69.(*string)
		}(header))
	}
	var uploadError *string = nil
	upload := self.file
	if upload != nil {
		hxrt.HttpRequestSetMultipartUpload(request, func(hx_obj_71 map[string]any) *string {
			hx_field_72 := hx_obj_71["param"]
			if hx_field_72 == nil {
				var hx_zero_73 *string
				return hx_zero_73
			}
			return hx_field_72.(*string)
		}(upload), func(hx_obj_74 map[string]any) *string {
			hx_field_75 := hx_obj_74["filename"]
			if hx_field_75 == nil {
				var hx_zero_76 *string
				return hx_zero_76
			}
			return hx_field_75.(*string)
		}(upload), func(hx_obj_77 map[string]any) *string {
			hx_field_78 := hx_obj_77["mimeType"]
			if hx_field_78 == nil {
				var hx_zero_79 *string
				return hx_zero_79
			}
			return hx_field_78.(*string)
		}(upload), func(hx_obj_80 map[string]any) int {
			hx_field_81 := hx_obj_80["size"]
			if hx_field_81 == nil {
				var hx_zero_82 int
				return hx_zero_82
			}
			return hx_field_81.(int)
		}(upload), func(requested int) *hxrt.ByteView {
			var result *hxrt.ByteView = nil
			if requested > 0 {
				chunk := haxe__io__Bytes_alloc(requested)
				hxrt.TryCatch(func() {
					count := func(hx_obj_85 map[string]any) *haxe__io__Input {
						hx_field_86 := hx_obj_85["io"]
						if hx_field_86 == nil {
							var hx_zero_87 *haxe__io__Input
							return hx_zero_87
						}
						return hx_field_86.(*haxe__io__Input)
					}(upload).__hx_this.readBytes(chunk, 0, requested)
					if count > 0 {
						if count < requested {
							chunk = chunk.__hx_this.sub(0, count)
						}
						result = chunk.__hx_this.__hx_nativeView()
					}
				}, func(hx_caught_83 any) {
					switch hx_typed_84 := hx_caught_83.(type) {
					case *haxe__io__Eof:
						hx_tmp := hx_typed_84
						_ = hx_tmp
					default:
						error := hxrt.ExceptionCaught(hx_caught_83)
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
	if (proxy != nil) && !hxrt.StringEqualStringPtr(func(hx_obj_114 map[string]any) *string {
		hx_field_115 := hx_obj_114["host"]
		if hx_field_115 == nil {
			var hx_zero_116 *string
			return hx_zero_116
		}
		return hx_field_115.(*string)
	}(proxy), nil) {
		var hx_if_97 *string
		if func(hx_obj_88 map[string]any) map[string]any {
			hx_field_89 := hx_obj_88["auth"]
			if hx_field_89 == nil {
				var hx_zero_90 map[string]any
				return hx_zero_90
			}
			return hx_field_89.(map[string]any)
		}(proxy) == nil {
			hx_if_97 = nil
		} else {
			hx_if_97 = func(hx_obj_94 map[string]any) *string {
				hx_field_95 := hx_obj_94["user"]
				if hx_field_95 == nil {
					var hx_zero_96 *string
					return hx_zero_96
				}
				return hx_field_95.(*string)
			}(func(hx_obj_91 map[string]any) map[string]any {
				hx_field_92 := hx_obj_91["auth"]
				if hx_field_92 == nil {
					var hx_zero_93 map[string]any
					return hx_zero_93
				}
				return hx_field_92.(map[string]any)
			}(proxy))
		}
		user := hx_if_97
		var hx_if_107 *string
		if func(hx_obj_98 map[string]any) map[string]any {
			hx_field_99 := hx_obj_98["auth"]
			if hx_field_99 == nil {
				var hx_zero_100 map[string]any
				return hx_zero_100
			}
			return hx_field_99.(map[string]any)
		}(proxy) == nil {
			hx_if_107 = nil
		} else {
			hx_if_107 = func(hx_obj_104 map[string]any) *string {
				hx_field_105 := hx_obj_104["pass"]
				if hx_field_105 == nil {
					var hx_zero_106 *string
					return hx_zero_106
				}
				return hx_field_105.(*string)
			}(func(hx_obj_101 map[string]any) map[string]any {
				hx_field_102 := hx_obj_101["auth"]
				if hx_field_102 == nil {
					var hx_zero_103 map[string]any
					return hx_zero_103
				}
				return hx_field_102.(map[string]any)
			}(proxy))
		}
		pass := hx_if_107
		hxrt.HttpRequestSetProxy(request, func(hx_obj_108 map[string]any) *string {
			hx_field_109 := hx_obj_108["host"]
			if hx_field_109 == nil {
				var hx_zero_110 *string
				return hx_zero_110
			}
			return hx_field_109.(*string)
		}(proxy), func(hx_obj_111 map[string]any) int {
			hx_field_112 := hx_obj_111["port"]
			if hx_field_112 == nil {
				var hx_zero_113 int
				return hx_zero_113
			}
			return hx_field_112.(int)
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
			encoded = hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("multipart file="), func(hx_obj_117 map[string]any) *string {
				hx_field_118 := hx_obj_117["filename"]
				if hx_field_118 == nil {
					var hx_zero_119 *string
					return hx_zero_119
				}
				return hx_field_118.(*string)
			}(self.file)), hxrt.StringFromLiteral(";mime=")), func(hx_obj_120 map[string]any) *string {
				hx_field_121 := hx_obj_120["mimeType"]
				if hx_field_121 == nil {
					var hx_zero_122 *string
					return hx_zero_122
				}
				return hx_field_121.(*string)
			}(self.file)), hxrt.StringFromLiteral(";size=")), func(hx_obj_123 map[string]any) int {
				hx_field_124 := hx_obj_123["size"]
				if hx_field_124 == nil {
					var hx_zero_125 int
					return hx_zero_125
				}
				return hx_field_124.(int)
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
		hx_post_126 := _g
		_g = int(int32((_g + 1)))
		headerIndex := hx_post_126
		name := hxrt.StdString(hxrt.HttpResponseHeaderName(response, headerIndex))
		normalized := hxrt.StringToLowerCaseStringPtr(name)
		valueCount := hxrt.HttpResponseHeaderValueCount(response, headerIndex)
		values := hxrt.NewArray()
		_g_1 := 0
		_g1_1 := valueCount
		for _g_1 < _g1_1 {
			hx_post_127 := _g_1
			_g_1 = int(int32((_g_1 + 1)))
			valueIndex := hx_post_127
			values.Push(hxrt.StdString(hxrt.HttpResponseHeaderValue(response, headerIndex, valueIndex)))
		}
		if values.Len() == 0 {
			continue
		}
		last := func(hx_value_129 any) *string {
			if hx_value_129 == nil {
				var hx_zero_130 *string
				return hx_zero_130
			}
			return hx_value_129.(*string)
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
		header := func(hx_value_131 any) map[string]any {
			if hx_value_131 == nil {
				var hx_zero_132 map[string]any
				return hx_zero_132
			}
			return hx_value_131.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(hxrt.StringToLowerCaseStringPtr(func(hx_obj_133 map[string]any) *string {
			hx_field_134 := hx_obj_133["name"]
			if hx_field_134 == nil {
				var hx_zero_135 *string
				return hx_zero_135
			}
			return hx_field_134.(*string)
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
		parameter := func(hx_value_136 any) map[string]any {
			if hx_value_136 == nil {
				var hx_zero_137 map[string]any
				return hx_zero_137
			}
			return hx_value_136.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		byName.__hx_this.set(func(hx_obj_138 map[string]any) *string {
			hx_field_139 := hx_obj_138["name"]
			if hx_field_139 == nil {
				var hx_zero_140 *string
				return hx_zero_140
			}
			return hx_field_139.(*string)
		}(parameter), func(hx_obj_141 map[string]any) *string {
			hx_field_142 := hx_obj_141["value"]
			if hx_field_142 == nil {
				var hx_zero_143 *string
				return hx_zero_143
			}
			return hx_field_142.(*string)
		}(parameter))
	}
	_g_1 := hxrt.NewArray()
	name := func(hx_value_144 any) map[string]any {
		if hx_value_144 == nil {
			var hx_zero_145 map[string]any
			return hx_zero_145
		}
		return hx_value_144.(map[string]any)
	}(byName.__hx_this.keys())
	for func(hx_obj_146 map[string]any) func() bool {
		hx_field_147 := hx_obj_146["hasNext"]
		if hx_field_147 == nil {
			var hx_zero_148 func() bool
			return hx_zero_148
		}
		return hx_field_147.(func() bool)
	}(name)() {
		name_1 := func(hx_obj_149 map[string]any) func() *string {
			hx_field_150 := hx_obj_149["next"]
			if hx_field_150 == nil {
				var hx_zero_151 func() *string
				return hx_zero_151
			}
			return hx_field_150.(func() *string)
		}(name)()
		_g_1.Push(name_1)
	}
	names := _g_1
	encoded := hxrt.NewArray()
	emitted := New_haxe__ds__StringMap()
	_g_2 := 0
	for _g_2 < names.Len() {
		hx_tmp := func(hx_value_153 any) *string {
			if hx_value_153 == nil {
				var hx_zero_154 *string
				return hx_zero_154
			}
			return hx_value_153.(*string)
		}(names.Get(_g_2))
		_ = hx_tmp
		_g_2 = int(int32((_g_2 + 1)))
		next := -1
		_g_3 := 0
		_g1_1 := names.Len()
		for _g_3 < _g1_1 {
			hx_post_155 := _g_3
			_g_3 = int(int32((_g_3 + 1)))
			index := hx_post_155
			if !func(hx_value_158 any) bool {
				if hx_value_158 == nil {
					var hx_zero_159 bool
					return hx_zero_159
				}
				return hx_value_158.(bool)
			}(emitted.__hx_this.exists(func(hx_value_156 any) *string {
				if hx_value_156 == nil {
					var hx_zero_157 *string
					return hx_zero_157
				}
				return hx_value_156.(*string)
			}(names.Get(index)))) && ((next < 0) || (Reflect_compare(func(hx_value_160 any) *string {
				if hx_value_160 == nil {
					var hx_zero_161 *string
					return hx_zero_161
				}
				return hx_value_160.(*string)
			}(names.Get(index)), func(hx_value_162 any) *string {
				if hx_value_162 == nil {
					var hx_zero_163 *string
					return hx_zero_163
				}
				return hx_value_162.(*string)
			}(names.Get(next))) < 0)) {
				next = index
			}
		}
		if next >= 0 {
			name_2 := func(hx_value_164 any) *string {
				if hx_value_164 == nil {
					var hx_zero_165 *string
					return hx_zero_165
				}
				return hx_value_164.(*string)
			}(names.Get(next))
			emitted.__hx_this.set(name_2, true)
			encoded.Push(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(StringTools_urlEncode(name_2), hxrt.StringFromLiteral("=")), StringTools_urlEncode(func(hx_value_167 any) *string {
				if hx_value_167 == nil {
					var hx_zero_168 *string
					return hx_zero_168
				}
				return hx_value_167.(*string)
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
		hx_post_169 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_169
		if func() int {
			var c any = hxrt.StringCharCodeAtAnyStringPtr(value, index)
			var hx_if_170 int
			if c == nil {
				hx_if_170 = -1
			} else {
				hx_if_170 = c.(int)
			}
			return hx_if_170
		}() == 44 {
			return index
		}
	}
	return -1
}

func sys__Http_hxrt_proxyDescriptor() *string {
	proxy := sys__Http_PROXY
	if (proxy == nil) || hxrt.StringEqualStringPtr(func(hx_obj_171 map[string]any) *string {
		hx_field_172 := hx_obj_171["host"]
		if hx_field_172 == nil {
			var hx_zero_173 *string
			return hx_zero_173
		}
		return hx_field_172.(*string)
	}(proxy), nil) {
		return hxrt.StringFromLiteral("null")
	}
	var hx_if_183 *string
	if func(hx_obj_174 map[string]any) map[string]any {
		hx_field_175 := hx_obj_174["auth"]
		if hx_field_175 == nil {
			var hx_zero_176 map[string]any
			return hx_zero_176
		}
		return hx_field_175.(map[string]any)
	}(proxy) == nil {
		hx_if_183 = nil
	} else {
		hx_if_183 = func(hx_obj_180 map[string]any) *string {
			hx_field_181 := hx_obj_180["user"]
			if hx_field_181 == nil {
				var hx_zero_182 *string
				return hx_zero_182
			}
			return hx_field_181.(*string)
		}(func(hx_obj_177 map[string]any) map[string]any {
			hx_field_178 := hx_obj_177["auth"]
			if hx_field_178 == nil {
				var hx_zero_179 map[string]any
				return hx_zero_179
			}
			return hx_field_178.(map[string]any)
		}(proxy))
	}
	user := hx_if_183
	var hx_if_193 *string
	if func(hx_obj_184 map[string]any) map[string]any {
		hx_field_185 := hx_obj_184["auth"]
		if hx_field_185 == nil {
			var hx_zero_186 map[string]any
			return hx_zero_186
		}
		return hx_field_185.(map[string]any)
	}(proxy) == nil {
		hx_if_193 = nil
	} else {
		hx_if_193 = func(hx_obj_190 map[string]any) *string {
			hx_field_191 := hx_obj_190["pass"]
			if hx_field_191 == nil {
				var hx_zero_192 *string
				return hx_zero_192
			}
			return hx_field_191.(*string)
		}(func(hx_obj_187 map[string]any) map[string]any {
			hx_field_188 := hx_obj_187["auth"]
			if hx_field_188 == nil {
				var hx_zero_189 map[string]any
				return hx_zero_189
			}
			return hx_field_188.(map[string]any)
		}(proxy))
	}
	pass := hx_if_193
	return hxrt.StdString(hxrt.HttpProxyDescriptor(func(hx_obj_194 map[string]any) *string {
		hx_field_195 := hx_obj_194["host"]
		if hx_field_195 == nil {
			var hx_zero_196 *string
			return hx_zero_196
		}
		return hx_field_195.(*string)
	}(proxy), func(hx_obj_197 map[string]any) int {
		hx_field_198 := hx_obj_197["port"]
		if hx_field_198 == nil {
			var hx_zero_199 int
			return hx_zero_199
		}
		return hx_field_198.(int)
	}(proxy), user, pass))
}

func sys__Http_normalizedMethod(method *string) *string {
	if hxrt.StringEqualStringPtr(method, nil) {
		return nil
	}
	normalized := hxrt.StringToUpperCaseStringPtr(method)
	var hx_if_200 *string
	if hxrt.StringEqualStringPtr(normalized, hxrt.StringFromLiteral("")) || hxrt.StringEqualStringPtr(normalized, hxrt.StringFromLiteral("NULL")) {
		hx_if_200 = nil
	} else {
		hx_if_200 = normalized
	}
	return hx_if_200
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
