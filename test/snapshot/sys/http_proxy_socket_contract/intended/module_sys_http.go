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
	hx_obj_45 := map[string]any{}
	hx_obj_45["param"] = argname
	hx_obj_45["filename"] = filename
	hx_obj_45["io"] = file
	hx_obj_45["size"] = size
	hx_obj_45["mimeType"] = mimeType
	self.file = hx_obj_45
}

func (self *sys__Http) customRequest(post bool, api *haxe__io__Output, sock *sys__net__Socket, method *string) {
	self.__hx_this.requestWith((post || (self.file != nil)), api, sock, method)
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
		}(parameter), StringTools_urlEncode(func(hx_obj_63 map[string]any) *string {
			hx_field_64 := hx_obj_63["name"]
			if hx_field_64 == nil {
				var hx_zero_65 *string
				return hx_zero_65
			}
			return hx_field_64.(*string)
		}(parameter)), StringTools_urlEncode(func(hx_obj_66 map[string]any) *string {
			hx_field_67 := hx_obj_66["value"]
			if hx_field_67 == nil {
				var hx_zero_68 *string
				return hx_zero_68
			}
			return hx_field_67.(*string)
		}(parameter)))
	}
	_g_1 := 0
	_g1_1 := self.headers
	for _g_1 < _g1_1.Len() {
		header := func(hx_value_69 any) map[string]any {
			if hx_value_69 == nil {
				var hx_zero_70 map[string]any
				return hx_zero_70
			}
			return hx_value_69.(map[string]any)
		}(_g1_1.Get(_g_1))
		_g_1 = int(int32((_g_1 + 1)))
		hxrt.HttpRequestAddHeader(request, func(hx_obj_71 map[string]any) *string {
			hx_field_72 := hx_obj_71["name"]
			if hx_field_72 == nil {
				var hx_zero_73 *string
				return hx_zero_73
			}
			return hx_field_72.(*string)
		}(header), func(hx_obj_74 map[string]any) *string {
			hx_field_75 := hx_obj_74["value"]
			if hx_field_75 == nil {
				var hx_zero_76 *string
				return hx_zero_76
			}
			return hx_field_75.(*string)
		}(header))
	}
	var uploadError *string = nil
	upload := self.file
	if upload != nil {
		hxrt.HttpRequestSetMultipartUpload(request, func(hx_obj_77 map[string]any) *string {
			hx_field_78 := hx_obj_77["param"]
			if hx_field_78 == nil {
				var hx_zero_79 *string
				return hx_zero_79
			}
			return hx_field_78.(*string)
		}(upload), func(hx_obj_80 map[string]any) *string {
			hx_field_81 := hx_obj_80["filename"]
			if hx_field_81 == nil {
				var hx_zero_82 *string
				return hx_zero_82
			}
			return hx_field_81.(*string)
		}(upload), func(hx_obj_83 map[string]any) *string {
			hx_field_84 := hx_obj_83["mimeType"]
			if hx_field_84 == nil {
				var hx_zero_85 *string
				return hx_zero_85
			}
			return hx_field_84.(*string)
		}(upload), func(hx_obj_86 map[string]any) int {
			hx_field_87 := hx_obj_86["size"]
			if hx_field_87 == nil {
				var hx_zero_88 int
				return hx_zero_88
			}
			return hx_field_87.(int)
		}(upload), func(requested int) *hxrt.ByteView {
			var result *hxrt.ByteView = nil
			if requested > 0 {
				chunk := haxe__io__Bytes_alloc(requested)
				hxrt.TryCatch(func() {
					count := func(hx_obj_91 map[string]any) *haxe__io__Input {
						hx_field_92 := hx_obj_91["io"]
						if hx_field_92 == nil {
							var hx_zero_93 *haxe__io__Input
							return hx_zero_93
						}
						return hx_field_92.(*haxe__io__Input)
					}(upload).__hx_this.readBytes(chunk, 0, requested)
					if count > 0 {
						if count < requested {
							chunk = chunk.__hx_this.sub(0, count)
						}
						result = chunk.__hx_this.__hx_nativeView()
					}
				}, func(hx_caught_89 any) {
					switch hx_typed_90 := hx_caught_89.(type) {
					case *haxe__io__Eof:
						hx_tmp := hx_typed_90
						_ = hx_tmp
					default:
						error := hxrt.ExceptionCaught(hx_caught_89)
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
	if (proxy != nil) && !hxrt.StringEqualStringPtr(func(hx_obj_120 map[string]any) *string {
		hx_field_121 := hx_obj_120["host"]
		if hx_field_121 == nil {
			var hx_zero_122 *string
			return hx_zero_122
		}
		return hx_field_121.(*string)
	}(proxy), nil) {
		var hx_if_103 *string
		if func(hx_obj_94 map[string]any) map[string]any {
			hx_field_95 := hx_obj_94["auth"]
			if hx_field_95 == nil {
				var hx_zero_96 map[string]any
				return hx_zero_96
			}
			return hx_field_95.(map[string]any)
		}(proxy) == nil {
			hx_if_103 = nil
		} else {
			hx_if_103 = func(hx_obj_100 map[string]any) *string {
				hx_field_101 := hx_obj_100["user"]
				if hx_field_101 == nil {
					var hx_zero_102 *string
					return hx_zero_102
				}
				return hx_field_101.(*string)
			}(func(hx_obj_97 map[string]any) map[string]any {
				hx_field_98 := hx_obj_97["auth"]
				if hx_field_98 == nil {
					var hx_zero_99 map[string]any
					return hx_zero_99
				}
				return hx_field_98.(map[string]any)
			}(proxy))
		}
		user := hx_if_103
		var hx_if_113 *string
		if func(hx_obj_104 map[string]any) map[string]any {
			hx_field_105 := hx_obj_104["auth"]
			if hx_field_105 == nil {
				var hx_zero_106 map[string]any
				return hx_zero_106
			}
			return hx_field_105.(map[string]any)
		}(proxy) == nil {
			hx_if_113 = nil
		} else {
			hx_if_113 = func(hx_obj_110 map[string]any) *string {
				hx_field_111 := hx_obj_110["pass"]
				if hx_field_111 == nil {
					var hx_zero_112 *string
					return hx_zero_112
				}
				return hx_field_111.(*string)
			}(func(hx_obj_107 map[string]any) map[string]any {
				hx_field_108 := hx_obj_107["auth"]
				if hx_field_108 == nil {
					var hx_zero_109 map[string]any
					return hx_zero_109
				}
				return hx_field_108.(map[string]any)
			}(proxy))
		}
		pass := hx_if_113
		hxrt.HttpRequestSetProxy(request, func(hx_obj_114 map[string]any) *string {
			hx_field_115 := hx_obj_114["host"]
			if hx_field_115 == nil {
				var hx_zero_116 *string
				return hx_zero_116
			}
			return hx_field_115.(*string)
		}(proxy), func(hx_obj_117 map[string]any) int {
			hx_field_118 := hx_obj_117["port"]
			if hx_field_118 == nil {
				var hx_zero_119 int
				return hx_zero_119
			}
			return hx_field_118.(int)
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
			}, func(hx_caught_123 any) {
				error := hxrt.ExceptionCaught(hx_caught_123)
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
			encoded = hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("multipart file="), func(hx_obj_125 map[string]any) *string {
				hx_field_126 := hx_obj_125["filename"]
				if hx_field_126 == nil {
					var hx_zero_127 *string
					return hx_zero_127
				}
				return hx_field_126.(*string)
			}(self.file)), hxrt.StringFromLiteral(";mime=")), func(hx_obj_128 map[string]any) *string {
				hx_field_129 := hx_obj_128["mimeType"]
				if hx_field_129 == nil {
					var hx_zero_130 *string
					return hx_zero_130
				}
				return hx_field_129.(*string)
			}(self.file)), hxrt.StringFromLiteral(";size=")), func(hx_obj_131 map[string]any) int {
				hx_field_132 := hx_obj_131["size"]
				if hx_field_132 == nil {
					var hx_zero_133 int
					return hx_zero_133
				}
				return hx_field_132.(int)
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
	}, func(hx_caught_134 any) {
		error := hxrt.ExceptionCaught(hx_caught_134)
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
		hx_post_136 := _g
		_g = int(int32((_g + 1)))
		headerIndex := hx_post_136
		name := hxrt.StdString(hxrt.HttpExchangeHeaderName(exchange, headerIndex))
		normalized := hxrt.StringToLowerCaseStringPtr(name)
		valueCount := hxrt.HttpExchangeHeaderValueCount(exchange, headerIndex)
		values := hxrt.NewArray()
		_g_1 := 0
		_g1_1 := valueCount
		for _g_1 < _g1_1 {
			hx_post_137 := _g_1
			_g_1 = int(int32((_g_1 + 1)))
			valueIndex := hx_post_137
			values.Push(hxrt.StdString(hxrt.HttpExchangeHeaderValue(exchange, headerIndex, valueIndex)))
		}
		if values.Len() == 0 {
			continue
		}
		last := func(hx_value_139 any) *string {
			if hx_value_139 == nil {
				var hx_zero_140 *string
				return hx_zero_140
			}
			return hx_value_139.(*string)
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
		parameter := func(hx_value_141 any) map[string]any {
			if hx_value_141 == nil {
				var hx_zero_142 map[string]any
				return hx_zero_142
			}
			return hx_value_141.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		encoded.Push(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(StringTools_urlEncode(func(hx_obj_144 map[string]any) *string {
			hx_field_145 := hx_obj_144["name"]
			if hx_field_145 == nil {
				var hx_zero_146 *string
				return hx_zero_146
			}
			return hx_field_145.(*string)
		}(parameter)), hxrt.StringFromLiteral("=")), StringTools_urlEncode(func(hx_obj_147 map[string]any) *string {
			hx_field_148 := hx_obj_147["value"]
			if hx_field_148 == nil {
				var hx_zero_149 *string
				return hx_zero_149
			}
			return hx_field_148.(*string)
		}(parameter))))
	}
	return hxrt.StringJoinAny(encoded.Values(), hxrt.StringFromLiteral("&"))
}

var sys__Http_PROXY map[string]any = nil

func sys__Http_firstComma(value *string) int {
	_g := 0
	_g1 := hxrt.StringLengthStringPtr(value)
	for _g < _g1 {
		hx_post_150 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_150
		if func() int {
			var c any = hxrt.StringCharCodeAtAnyStringPtr(value, index)
			var hx_if_151 int
			if c == nil {
				hx_if_151 = -1
			} else {
				hx_if_151 = c.(int)
			}
			return hx_if_151
		}() == 44 {
			return index
		}
	}
	return -1
}

func sys__Http_hxrt_proxyDescriptor() *string {
	proxy := sys__Http_PROXY
	if (proxy == nil) || hxrt.StringEqualStringPtr(func(hx_obj_152 map[string]any) *string {
		hx_field_153 := hx_obj_152["host"]
		if hx_field_153 == nil {
			var hx_zero_154 *string
			return hx_zero_154
		}
		return hx_field_153.(*string)
	}(proxy), nil) {
		return hxrt.StringFromLiteral("null")
	}
	var hx_if_164 *string
	if func(hx_obj_155 map[string]any) map[string]any {
		hx_field_156 := hx_obj_155["auth"]
		if hx_field_156 == nil {
			var hx_zero_157 map[string]any
			return hx_zero_157
		}
		return hx_field_156.(map[string]any)
	}(proxy) == nil {
		hx_if_164 = nil
	} else {
		hx_if_164 = func(hx_obj_161 map[string]any) *string {
			hx_field_162 := hx_obj_161["user"]
			if hx_field_162 == nil {
				var hx_zero_163 *string
				return hx_zero_163
			}
			return hx_field_162.(*string)
		}(func(hx_obj_158 map[string]any) map[string]any {
			hx_field_159 := hx_obj_158["auth"]
			if hx_field_159 == nil {
				var hx_zero_160 map[string]any
				return hx_zero_160
			}
			return hx_field_159.(map[string]any)
		}(proxy))
	}
	user := hx_if_164
	var hx_if_174 *string
	if func(hx_obj_165 map[string]any) map[string]any {
		hx_field_166 := hx_obj_165["auth"]
		if hx_field_166 == nil {
			var hx_zero_167 map[string]any
			return hx_zero_167
		}
		return hx_field_166.(map[string]any)
	}(proxy) == nil {
		hx_if_174 = nil
	} else {
		hx_if_174 = func(hx_obj_171 map[string]any) *string {
			hx_field_172 := hx_obj_171["pass"]
			if hx_field_172 == nil {
				var hx_zero_173 *string
				return hx_zero_173
			}
			return hx_field_172.(*string)
		}(func(hx_obj_168 map[string]any) map[string]any {
			hx_field_169 := hx_obj_168["auth"]
			if hx_field_169 == nil {
				var hx_zero_170 map[string]any
				return hx_zero_170
			}
			return hx_field_169.(map[string]any)
		}(proxy))
	}
	pass := hx_if_174
	return hxrt.StdString(hxrt.HttpProxyDescriptor(func(hx_obj_175 map[string]any) *string {
		hx_field_176 := hx_obj_175["host"]
		if hx_field_176 == nil {
			var hx_zero_177 *string
			return hx_zero_177
		}
		return hx_field_176.(*string)
	}(proxy), func(hx_obj_178 map[string]any) int {
		hx_field_179 := hx_obj_178["port"]
		if hx_field_179 == nil {
			var hx_zero_180 int
			return hx_zero_180
		}
		return hx_field_179.(int)
	}(proxy), user, pass))
}

func sys__Http_normalizedMethod(method *string) *string {
	if hxrt.StringEqualStringPtr(method, nil) {
		return nil
	}
	normalized := hxrt.StringToUpperCaseStringPtr(method)
	var hx_if_181 *string
	if hxrt.StringEqualStringPtr(normalized, hxrt.StringFromLiteral("")) || hxrt.StringEqualStringPtr(normalized, hxrt.StringFromLiteral("NULL")) {
		hx_if_181 = nil
	} else {
		hx_if_181 = normalized
	}
	return hx_if_181
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
