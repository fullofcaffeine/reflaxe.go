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
	recordResponseHeaders(exchange *hxrt.HttpExchange)
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
	_gthis := self
	output := New_haxe__io__BytesOutput()
	previousOnError := self.onError
	failed := false
	self.onError = func(message *string) {
		_gthis.responseBytes = output.__hx_this.getBytes()
		_gthis.responseAsString = nil
		failed = true
		_gthis.onError = previousOnError
		func(hx_fn func(*string), hx_arg_0 *string) {
			if hx_fn == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
				return
			}
			hx_fn(hx_arg_0)
		}(_gthis.onError, message)
	}
	usePost := (((((post != nil) && (post.(bool) == true)) || (self.postBytes != nil)) || !hxrt.StringEqualStringPtr(self.postData, nil)) || (self.file != nil))
	self.__hx_this.customRequest(usePost, output.haxe__io__Output, nil, nil)
	if !failed {
		self.__hx_this.success(output.__hx_this.getBytes())
	}
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
	self.__hx_this.requestWith((post || (self.file != nil)), api, sock, method)
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
	exchange := hxrt.HttpRequestStartExchange(request)
	errorMessage := uploadError
	completed := false
	if hxrt.StringEqualStringPtr(errorMessage, nil) {
		nativeError := hxrt.HttpExchangeError(exchange)
		if !hxrt.StringEqualStringPtr(nativeError, nil) {
			errorMessage = nativeError
		} else {
			hxrt.TryCatch(func() {
				self.__hx_this.recordResponseHeaders(exchange)
				status := hxrt.HttpExchangeStatus(exchange)
				func(hx_fn func(int), hx_arg_0 int) {
					if hx_fn == nil {
						hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
						return
					}
					hx_fn(hx_arg_0)
				}(self.onStatus, status)
				contentLength := hxrt.HttpExchangeContentLength(exchange)
				if contentLength == -2 {
					hxrt.Throw(hxrt.StringFromLiteral("Content-Length exceeds Haxe Int range"))
				}
				if contentLength >= 0 {
					api.__hx_this.prepare(contentLength)
				}
				for true {
					read := hxrt.HttpExchangeReadResponseChunk(exchange, 1024)
					payload := haxe__io__Bytes___hx_fromNativeView(hxrt.HttpReadResultBody(read))
					if payload.length > 0 {
						api.__hx_this.writeBytes(payload, 0, payload.length)
					}
					readError := hxrt.HttpReadResultError(read)
					if !hxrt.StringEqualStringPtr(readError, nil) {
						hxrt.Throw(hxrt.StringFromLiteral("Transfer aborted"))
					}
					if hxrt.HttpReadResultEOF(read) {
						break
					}
				}
				if status >= 400 {
					hxrt.Throw(hxrt.StringConcatAny(hxrt.StringFromLiteral("Http Error #"), status))
				}
				api.__hx_this.close()
				completed = true
			}, func(hx_caught_122 any) {
				error := hxrt.ExceptionCaught(hx_caught_122)
				errorMessage = hxrt.ExceptionMessage(error)
			})
		}
	}
	if completed {
		hxrt.HttpExchangeClose(exchange)
	} else {
		hxrt.HttpExchangeCancel(exchange)
	}
	if !hxrt.StringEqualStringPtr(errorMessage, nil) {
		func(hx_fn func(*string), hx_arg_0 *string) {
			if hx_fn == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
				return
			}
			hx_fn(hx_arg_0)
		}(self.onError, errorMessage)
	}
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
			encoded = hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("multipart file="), func(hx_obj_124 map[string]any) *string {
				hx_field_125 := hx_obj_124["filename"]
				if hx_field_125 == nil {
					var hx_zero_126 *string
					return hx_zero_126
				}
				return hx_field_125.(*string)
			}(self.file)), hxrt.StringFromLiteral(";mime=")), func(hx_obj_127 map[string]any) *string {
				hx_field_128 := hx_obj_127["mimeType"]
				if hx_field_128 == nil {
					var hx_zero_129 *string
					return hx_zero_129
				}
				return hx_field_128.(*string)
			}(self.file)), hxrt.StringFromLiteral(";size=")), func(hx_obj_130 map[string]any) int {
				hx_field_131 := hx_obj_130["size"]
				if hx_field_131 == nil {
					var hx_zero_132 int
					return hx_zero_132
				}
				return hx_field_131.(int)
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
	var this1 haxe__IMap = self.responseHeaders
	this1.(*haxe__ds__StringMap).__hx_this.set(hxrt.StringFromLiteral("content-type"), mediaType)
	var this1_1 haxe__IMap = self.responseHeaders
	this1_1.(*haxe__ds__StringMap).__hx_this.set(hxrt.StringFromLiteral("Content-Type"), mediaType)
	var errorMessage *string = nil
	hxrt.TryCatch(func() {
		func(hx_fn func(int), hx_arg_0 int) {
			if hx_fn == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
				return
			}
			hx_fn(hx_arg_0)
		}(self.onStatus, 200)
		api.__hx_this.prepare(payload.length)
		if payload.length > 0 {
			api.__hx_this.writeBytes(payload, 0, payload.length)
		}
		api.__hx_this.close()
	}, func(hx_caught_133 any) {
		error := hxrt.ExceptionCaught(hx_caught_133)
		errorMessage = hxrt.ExceptionMessage(error)
	})
	if !hxrt.StringEqualStringPtr(errorMessage, nil) {
		func(hx_fn func(*string), hx_arg_0 *string) {
			if hx_fn == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
				return
			}
			hx_fn(hx_arg_0)
		}(self.onError, errorMessage)
	}
}

func (self *sys__Http) recordResponseHeaders(exchange *hxrt.HttpExchange) {
	count := hxrt.HttpExchangeHeaderCount(exchange)
	_g := 0
	_g1 := count
	for _g < _g1 {
		hx_post_135 := _g
		_g = int(int32((_g + 1)))
		headerIndex := hx_post_135
		name := hxrt.StdString(hxrt.HttpExchangeHeaderName(exchange, headerIndex))
		normalized := hxrt.StringToLowerCaseStringPtr(name)
		valueCount := hxrt.HttpExchangeHeaderValueCount(exchange, headerIndex)
		values := hxrt.NewArray()
		_g_1 := 0
		_g1_1 := valueCount
		for _g_1 < _g1_1 {
			hx_post_136 := _g_1
			_g_1 = int(int32((_g_1 + 1)))
			valueIndex := hx_post_136
			values.Push(hxrt.StdString(hxrt.HttpExchangeHeaderValue(exchange, headerIndex, valueIndex)))
		}
		if values.Len() == 0 {
			continue
		}
		last := func(hx_value_138 any) *string {
			if hx_value_138 == nil {
				var hx_zero_139 *string
				return hx_zero_139
			}
			return hx_value_138.(*string)
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
		parameter := func(hx_value_140 any) map[string]any {
			if hx_value_140 == nil {
				var hx_zero_141 map[string]any
				return hx_zero_141
			}
			return hx_value_140.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		encoded.Push(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(StringTools_urlEncode(func(hx_obj_143 map[string]any) *string {
			hx_field_144 := hx_obj_143["name"]
			if hx_field_144 == nil {
				var hx_zero_145 *string
				return hx_zero_145
			}
			return hx_field_144.(*string)
		}(parameter)), hxrt.StringFromLiteral("=")), StringTools_urlEncode(func(hx_obj_146 map[string]any) *string {
			hx_field_147 := hx_obj_146["value"]
			if hx_field_147 == nil {
				var hx_zero_148 *string
				return hx_zero_148
			}
			return hx_field_147.(*string)
		}(parameter))))
	}
	return hxrt.StringJoinAny(encoded.Values(), hxrt.StringFromLiteral("&"))
}

var sys__Http_PROXY map[string]any = nil

func sys__Http_firstComma(value *string) int {
	_g := 0
	_g1 := hxrt.StringLengthStringPtr(value)
	for _g < _g1 {
		hx_post_149 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_149
		if func() int {
			var c any = hxrt.StringCharCodeAtAnyStringPtr(value, index)
			var hx_if_150 int
			if c == nil {
				hx_if_150 = -1
			} else {
				hx_if_150 = c.(int)
			}
			return hx_if_150
		}() == 44 {
			return index
		}
	}
	return -1
}

func sys__Http_hxrt_proxyDescriptor() *string {
	proxy := sys__Http_PROXY
	if (proxy == nil) || hxrt.StringEqualStringPtr(func(hx_obj_151 map[string]any) *string {
		hx_field_152 := hx_obj_151["host"]
		if hx_field_152 == nil {
			var hx_zero_153 *string
			return hx_zero_153
		}
		return hx_field_152.(*string)
	}(proxy), nil) {
		return hxrt.StringFromLiteral("null")
	}
	var hx_if_163 *string
	if func(hx_obj_154 map[string]any) map[string]any {
		hx_field_155 := hx_obj_154["auth"]
		if hx_field_155 == nil {
			var hx_zero_156 map[string]any
			return hx_zero_156
		}
		return hx_field_155.(map[string]any)
	}(proxy) == nil {
		hx_if_163 = nil
	} else {
		hx_if_163 = func(hx_obj_160 map[string]any) *string {
			hx_field_161 := hx_obj_160["user"]
			if hx_field_161 == nil {
				var hx_zero_162 *string
				return hx_zero_162
			}
			return hx_field_161.(*string)
		}(func(hx_obj_157 map[string]any) map[string]any {
			hx_field_158 := hx_obj_157["auth"]
			if hx_field_158 == nil {
				var hx_zero_159 map[string]any
				return hx_zero_159
			}
			return hx_field_158.(map[string]any)
		}(proxy))
	}
	user := hx_if_163
	var hx_if_173 *string
	if func(hx_obj_164 map[string]any) map[string]any {
		hx_field_165 := hx_obj_164["auth"]
		if hx_field_165 == nil {
			var hx_zero_166 map[string]any
			return hx_zero_166
		}
		return hx_field_165.(map[string]any)
	}(proxy) == nil {
		hx_if_173 = nil
	} else {
		hx_if_173 = func(hx_obj_170 map[string]any) *string {
			hx_field_171 := hx_obj_170["pass"]
			if hx_field_171 == nil {
				var hx_zero_172 *string
				return hx_zero_172
			}
			return hx_field_171.(*string)
		}(func(hx_obj_167 map[string]any) map[string]any {
			hx_field_168 := hx_obj_167["auth"]
			if hx_field_168 == nil {
				var hx_zero_169 map[string]any
				return hx_zero_169
			}
			return hx_field_168.(map[string]any)
		}(proxy))
	}
	pass := hx_if_173
	return hxrt.StdString(hxrt.HttpProxyDescriptor(func(hx_obj_174 map[string]any) *string {
		hx_field_175 := hx_obj_174["host"]
		if hx_field_175 == nil {
			var hx_zero_176 *string
			return hx_zero_176
		}
		return hx_field_175.(*string)
	}(proxy), func(hx_obj_177 map[string]any) int {
		hx_field_178 := hx_obj_177["port"]
		if hx_field_178 == nil {
			var hx_zero_179 int
			return hx_zero_179
		}
		return hx_field_178.(int)
	}(proxy), user, pass))
}

func sys__Http_normalizedMethod(method *string) *string {
	if hxrt.StringEqualStringPtr(method, nil) {
		return nil
	}
	normalized := hxrt.StringToUpperCaseStringPtr(method)
	var hx_if_180 *string
	if hxrt.StringEqualStringPtr(normalized, hxrt.StringFromLiteral("")) || hxrt.StringEqualStringPtr(normalized, hxrt.StringFromLiteral("NULL")) {
		hx_if_180 = nil
	} else {
		hx_if_180 = normalized
	}
	return hx_if_180
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
