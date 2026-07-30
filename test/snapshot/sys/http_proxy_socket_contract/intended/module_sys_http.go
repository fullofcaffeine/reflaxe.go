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
	pumpUpload(exchange *hxrt.HttpExchange, upload map[string]any) map[string]any
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
		}(upload))
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
	if (proxy != nil) && !hxrt.StringEqualStringPtr(func(hx_obj_115 map[string]any) *string {
		hx_field_116 := hx_obj_115["host"]
		if hx_field_116 == nil {
			var hx_zero_117 *string
			return hx_zero_117
		}
		return hx_field_116.(*string)
	}(proxy), nil) {
		var hx_if_98 *string
		if func(hx_obj_89 map[string]any) map[string]any {
			hx_field_90 := hx_obj_89["auth"]
			if hx_field_90 == nil {
				var hx_zero_91 map[string]any
				return hx_zero_91
			}
			return hx_field_90.(map[string]any)
		}(proxy) == nil {
			hx_if_98 = nil
		} else {
			hx_if_98 = func(hx_obj_95 map[string]any) *string {
				hx_field_96 := hx_obj_95["user"]
				if hx_field_96 == nil {
					var hx_zero_97 *string
					return hx_zero_97
				}
				return hx_field_96.(*string)
			}(func(hx_obj_92 map[string]any) map[string]any {
				hx_field_93 := hx_obj_92["auth"]
				if hx_field_93 == nil {
					var hx_zero_94 map[string]any
					return hx_zero_94
				}
				return hx_field_93.(map[string]any)
			}(proxy))
		}
		user := hx_if_98
		var hx_if_108 *string
		if func(hx_obj_99 map[string]any) map[string]any {
			hx_field_100 := hx_obj_99["auth"]
			if hx_field_100 == nil {
				var hx_zero_101 map[string]any
				return hx_zero_101
			}
			return hx_field_100.(map[string]any)
		}(proxy) == nil {
			hx_if_108 = nil
		} else {
			hx_if_108 = func(hx_obj_105 map[string]any) *string {
				hx_field_106 := hx_obj_105["pass"]
				if hx_field_106 == nil {
					var hx_zero_107 *string
					return hx_zero_107
				}
				return hx_field_106.(*string)
			}(func(hx_obj_102 map[string]any) map[string]any {
				hx_field_103 := hx_obj_102["auth"]
				if hx_field_103 == nil {
					var hx_zero_104 map[string]any
					return hx_zero_104
				}
				return hx_field_103.(map[string]any)
			}(proxy))
		}
		pass := hx_if_108
		hxrt.HttpRequestSetProxy(request, func(hx_obj_109 map[string]any) *string {
			hx_field_110 := hx_obj_109["host"]
			if hx_field_110 == nil {
				var hx_zero_111 *string
				return hx_zero_111
			}
			return hx_field_110.(*string)
		}(proxy), func(hx_obj_112 map[string]any) int {
			hx_field_113 := hx_obj_112["port"]
			if hx_field_113 == nil {
				var hx_zero_114 int
				return hx_zero_114
			}
			return hx_field_113.(int)
		}(proxy), user, pass)
	}
	if sock != nil {
		hxrt.HttpRequestSetSocket(request, sock.handle)
	}
	exchange := hxrt.HttpRequestStartExchange(request)
	var hx_if_118 map[string]any
	if upload == nil {
		hx_if_118 = nil
	} else {
		hx_if_118 = self.__hx_this.pumpUpload(exchange, upload)
	}
	uploadResult := hx_if_118
	hxrt.HttpExchangeAwaitResponse(exchange)
	var hx_if_122 *string
	if uploadResult == nil {
		hx_if_122 = nil
	} else {
		hx_if_122 = func(hx_obj_119 map[string]any) *string {
			hx_field_120 := hx_obj_119["sourceError"]
			if hx_field_120 == nil {
				var hx_zero_121 *string
				return hx_zero_121
			}
			return hx_field_120.(*string)
		}(uploadResult)
	}
	sourceError := hx_if_122
	var hx_if_126 *string
	if uploadResult == nil {
		hx_if_126 = nil
	} else {
		hx_if_126 = func(hx_obj_123 map[string]any) *string {
			hx_field_124 := hx_obj_123["sinkError"]
			if hx_field_124 == nil {
				var hx_zero_125 *string
				return hx_zero_125
			}
			return hx_field_124.(*string)
		}(uploadResult)
	}
	sinkError := hx_if_126
	errorMessage := sourceError
	completed := false
	if hxrt.StringEqualStringPtr(errorMessage, nil) {
		nativeError := hxrt.HttpExchangeError(exchange)
		if !hxrt.StringEqualStringPtr(nativeError, nil) {
			errorMessage = nativeError
		} else {
			if (hxrt.HttpExchangeStatus(exchange) == 0) && !hxrt.StringEqualStringPtr(sinkError, nil) {
				errorMessage = sinkError
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
				}, func(hx_caught_127 any) {
					error := hxrt.ExceptionCaught(hx_caught_127)
					errorMessage = hxrt.ExceptionMessage(error)
				})
			}
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

func (self *sys__Http) pumpUpload(exchange *hxrt.HttpExchange, upload map[string]any) map[string]any {
	sink := hxrt.HttpExchangeUploadSink(exchange)
	if sink == nil {
		hx_obj_129 := map[string]any{}
		hx_obj_129["sourceError"] = nil
		hx_obj_129["sinkError"] = hxrt.StringFromLiteral("HTTP upload sink is unavailable")
		return hx_obj_129
	}
	remaining := func(hx_obj_130 map[string]any) int {
		hx_field_131 := hx_obj_130["size"]
		if hx_field_131 == nil {
			var hx_zero_132 int
			return hx_zero_132
		}
		return hx_field_131.(int)
	}(upload)
	var sourceError *string = nil
	var sinkError *string = nil
	for remaining > 0 {
		var hx_if_133 int
		if remaining > 32768 {
			hx_if_133 = 32768
		} else {
			hx_if_133 = remaining
		}
		requested := hx_if_133
		chunk := haxe__io__Bytes_alloc(requested)
		count := 0
		hxrt.TryCatch(func() {
			count = func(hx_obj_136 map[string]any) *haxe__io__Input {
				hx_field_137 := hx_obj_136["io"]
				if hx_field_137 == nil {
					var hx_zero_138 *haxe__io__Input
					return hx_zero_138
				}
				return hx_field_137.(*haxe__io__Input)
			}(upload).__hx_this.readBytes(chunk, 0, requested)
		}, func(hx_caught_134 any) {
			switch hx_typed_135 := hx_caught_134.(type) {
			case *haxe__io__Eof:
				hx_tmp := hx_typed_135
				_ = hx_tmp
				sourceError = hxrt.StringFromLiteral("Transfer aborted")
			default:
				error := hxrt.ExceptionCaught(hx_caught_134)
				sourceError = hxrt.ExceptionMessage(error)
			}
		})
		if !hxrt.StringEqualStringPtr(sourceError, nil) {
			break
		}
		if count <= 0 {
			sourceError = hxrt.StringFromLiteral("multipart upload made no progress")
			break
		}
		if count > requested {
			sourceError = hxrt.StringFromLiteral("multipart upload exceeded the requested chunk size")
			break
		}
		if count < requested {
			chunk = chunk.__hx_this.sub(0, count)
		}
		sinkError = hxrt.HttpUploadSinkWriteChunk(sink, chunk.__hx_this.__hx_nativeView())
		if !hxrt.StringEqualStringPtr(sinkError, nil) {
			break
		}
		remaining = int(int32((hxrt.Int32Wrap(remaining) - hxrt.Int32Wrap(count))))
	}
	if !hxrt.StringEqualStringPtr(sourceError, nil) {
		hxrt.HttpUploadSinkAbort(sink, sourceError)
	} else {
		if hxrt.StringEqualStringPtr(sinkError, nil) {
			sinkError = hxrt.HttpUploadSinkFinish(sink)
		}
	}
	hx_obj_139 := map[string]any{}
	hx_obj_139["sourceError"] = sourceError
	hx_obj_139["sinkError"] = sinkError
	return hx_obj_139
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
			encoded = hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("multipart file="), func(hx_obj_140 map[string]any) *string {
				hx_field_141 := hx_obj_140["filename"]
				if hx_field_141 == nil {
					var hx_zero_142 *string
					return hx_zero_142
				}
				return hx_field_141.(*string)
			}(self.file)), hxrt.StringFromLiteral(";mime=")), func(hx_obj_143 map[string]any) *string {
				hx_field_144 := hx_obj_143["mimeType"]
				if hx_field_144 == nil {
					var hx_zero_145 *string
					return hx_zero_145
				}
				return hx_field_144.(*string)
			}(self.file)), hxrt.StringFromLiteral(";size=")), func(hx_obj_146 map[string]any) int {
				hx_field_147 := hx_obj_146["size"]
				if hx_field_147 == nil {
					var hx_zero_148 int
					return hx_zero_148
				}
				return hx_field_147.(int)
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
	}, func(hx_caught_149 any) {
		error := hxrt.ExceptionCaught(hx_caught_149)
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
		hx_post_151 := _g
		_g = int(int32((_g + 1)))
		headerIndex := hx_post_151
		name := hxrt.StdString(hxrt.HttpExchangeHeaderName(exchange, headerIndex))
		normalized := hxrt.StringToLowerCaseStringPtr(name)
		valueCount := hxrt.HttpExchangeHeaderValueCount(exchange, headerIndex)
		values := hxrt.NewArray()
		_g_1 := 0
		_g1_1 := valueCount
		for _g_1 < _g1_1 {
			hx_post_152 := _g_1
			_g_1 = int(int32((_g_1 + 1)))
			valueIndex := hx_post_152
			values.Push(hxrt.StdString(hxrt.HttpExchangeHeaderValue(exchange, headerIndex, valueIndex)))
		}
		if values.Len() == 0 {
			continue
		}
		last := func(hx_value_154 any) *string {
			if hx_value_154 == nil {
				var hx_zero_155 *string
				return hx_zero_155
			}
			return hx_value_154.(*string)
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
		parameter := func(hx_value_156 any) map[string]any {
			if hx_value_156 == nil {
				var hx_zero_157 map[string]any
				return hx_zero_157
			}
			return hx_value_156.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		encoded.Push(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(StringTools_urlEncode(func(hx_obj_159 map[string]any) *string {
			hx_field_160 := hx_obj_159["name"]
			if hx_field_160 == nil {
				var hx_zero_161 *string
				return hx_zero_161
			}
			return hx_field_160.(*string)
		}(parameter)), hxrt.StringFromLiteral("=")), StringTools_urlEncode(func(hx_obj_162 map[string]any) *string {
			hx_field_163 := hx_obj_162["value"]
			if hx_field_163 == nil {
				var hx_zero_164 *string
				return hx_zero_164
			}
			return hx_field_163.(*string)
		}(parameter))))
	}
	return hxrt.StringJoinAny(encoded.Values(), hxrt.StringFromLiteral("&"))
}

var sys__Http_PROXY map[string]any = nil

func sys__Http_firstComma(value *string) int {
	_g := 0
	_g1 := hxrt.StringLengthStringPtr(value)
	for _g < _g1 {
		hx_post_165 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_165
		if func() int {
			var c any = hxrt.StringCharCodeAtAnyStringPtr(value, index)
			var hx_if_166 int
			if c == nil {
				hx_if_166 = -1
			} else {
				hx_if_166 = c.(int)
			}
			return hx_if_166
		}() == 44 {
			return index
		}
	}
	return -1
}

func sys__Http_hxrt_proxyDescriptor() *string {
	proxy := sys__Http_PROXY
	if (proxy == nil) || hxrt.StringEqualStringPtr(func(hx_obj_167 map[string]any) *string {
		hx_field_168 := hx_obj_167["host"]
		if hx_field_168 == nil {
			var hx_zero_169 *string
			return hx_zero_169
		}
		return hx_field_168.(*string)
	}(proxy), nil) {
		return hxrt.StringFromLiteral("null")
	}
	var hx_if_179 *string
	if func(hx_obj_170 map[string]any) map[string]any {
		hx_field_171 := hx_obj_170["auth"]
		if hx_field_171 == nil {
			var hx_zero_172 map[string]any
			return hx_zero_172
		}
		return hx_field_171.(map[string]any)
	}(proxy) == nil {
		hx_if_179 = nil
	} else {
		hx_if_179 = func(hx_obj_176 map[string]any) *string {
			hx_field_177 := hx_obj_176["user"]
			if hx_field_177 == nil {
				var hx_zero_178 *string
				return hx_zero_178
			}
			return hx_field_177.(*string)
		}(func(hx_obj_173 map[string]any) map[string]any {
			hx_field_174 := hx_obj_173["auth"]
			if hx_field_174 == nil {
				var hx_zero_175 map[string]any
				return hx_zero_175
			}
			return hx_field_174.(map[string]any)
		}(proxy))
	}
	user := hx_if_179
	var hx_if_189 *string
	if func(hx_obj_180 map[string]any) map[string]any {
		hx_field_181 := hx_obj_180["auth"]
		if hx_field_181 == nil {
			var hx_zero_182 map[string]any
			return hx_zero_182
		}
		return hx_field_181.(map[string]any)
	}(proxy) == nil {
		hx_if_189 = nil
	} else {
		hx_if_189 = func(hx_obj_186 map[string]any) *string {
			hx_field_187 := hx_obj_186["pass"]
			if hx_field_187 == nil {
				var hx_zero_188 *string
				return hx_zero_188
			}
			return hx_field_187.(*string)
		}(func(hx_obj_183 map[string]any) map[string]any {
			hx_field_184 := hx_obj_183["auth"]
			if hx_field_184 == nil {
				var hx_zero_185 map[string]any
				return hx_zero_185
			}
			return hx_field_184.(map[string]any)
		}(proxy))
	}
	pass := hx_if_189
	return hxrt.StdString(hxrt.HttpProxyDescriptor(func(hx_obj_190 map[string]any) *string {
		hx_field_191 := hx_obj_190["host"]
		if hx_field_191 == nil {
			var hx_zero_192 *string
			return hx_zero_192
		}
		return hx_field_191.(*string)
	}(proxy), func(hx_obj_193 map[string]any) int {
		hx_field_194 := hx_obj_193["port"]
		if hx_field_194 == nil {
			var hx_zero_195 int
			return hx_zero_195
		}
		return hx_field_194.(int)
	}(proxy), user, pass))
}

func sys__Http_normalizedMethod(method *string) *string {
	if hxrt.StringEqualStringPtr(method, nil) {
		return nil
	}
	normalized := hxrt.StringToUpperCaseStringPtr(method)
	var hx_if_196 *string
	if hxrt.StringEqualStringPtr(normalized, hxrt.StringFromLiteral("")) || hxrt.StringEqualStringPtr(normalized, hxrt.StringFromLiteral("NULL")) {
		hx_if_196 = nil
	} else {
		hx_if_196 = normalized
	}
	return hx_if_196
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
