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
	exchange := hxrt.HttpRequestStartExchange(request)
	var hx_if_117 map[string]any
	if upload == nil {
		hx_if_117 = nil
	} else {
		hx_if_117 = self.__hx_this.pumpUpload(exchange, upload)
	}
	uploadResult := hx_if_117
	hxrt.HttpExchangeAwaitResponse(exchange)
	var hx_if_121 *string
	if uploadResult == nil {
		hx_if_121 = nil
	} else {
		hx_if_121 = func(hx_obj_118 map[string]any) *string {
			hx_field_119 := hx_obj_118["sourceError"]
			if hx_field_119 == nil {
				var hx_zero_120 *string
				return hx_zero_120
			}
			return hx_field_119.(*string)
		}(uploadResult)
	}
	sourceError := hx_if_121
	var hx_if_125 *string
	if uploadResult == nil {
		hx_if_125 = nil
	} else {
		hx_if_125 = func(hx_obj_122 map[string]any) *string {
			hx_field_123 := hx_obj_122["sinkError"]
			if hx_field_123 == nil {
				var hx_zero_124 *string
				return hx_zero_124
			}
			return hx_field_123.(*string)
		}(uploadResult)
	}
	sinkError := hx_if_125
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
				}, func(hx_caught_126 any) {
					error := hxrt.ExceptionCaught(hx_caught_126)
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
		hx_obj_128 := map[string]any{}
		hx_obj_128["sourceError"] = nil
		hx_obj_128["sinkError"] = hxrt.StringFromLiteral("HTTP upload sink is unavailable")
		return hx_obj_128
	}
	remaining := func(hx_obj_129 map[string]any) int {
		hx_field_130 := hx_obj_129["size"]
		if hx_field_130 == nil {
			var hx_zero_131 int
			return hx_zero_131
		}
		return hx_field_130.(int)
	}(upload)
	var sourceError *string = nil
	var sinkError *string = nil
	for remaining > 0 {
		var hx_if_132 int
		if remaining > 32768 {
			hx_if_132 = 32768
		} else {
			hx_if_132 = remaining
		}
		requested := hx_if_132
		chunk := haxe__io__Bytes_alloc(requested)
		count := 0
		hxrt.TryCatch(func() {
			count = func(hx_obj_135 map[string]any) *haxe__io__Input {
				hx_field_136 := hx_obj_135["io"]
				if hx_field_136 == nil {
					var hx_zero_137 *haxe__io__Input
					return hx_zero_137
				}
				return hx_field_136.(*haxe__io__Input)
			}(upload).__hx_this.readBytes(chunk, 0, requested)
		}, func(hx_caught_133 any) {
			switch hx_typed_134 := hx_caught_133.(type) {
			case *haxe__io__Eof:
				hx_tmp := hx_typed_134
				_ = hx_tmp
				sourceError = hxrt.StringFromLiteral("Transfer aborted")
			default:
				error := hxrt.ExceptionCaught(hx_caught_133)
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
	hx_obj_138 := map[string]any{}
	hx_obj_138["sourceError"] = sourceError
	hx_obj_138["sinkError"] = sinkError
	return hx_obj_138
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
			encoded = hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("multipart file="), func(hx_obj_139 map[string]any) *string {
				hx_field_140 := hx_obj_139["filename"]
				if hx_field_140 == nil {
					var hx_zero_141 *string
					return hx_zero_141
				}
				return hx_field_140.(*string)
			}(self.file)), hxrt.StringFromLiteral(";mime=")), func(hx_obj_142 map[string]any) *string {
				hx_field_143 := hx_obj_142["mimeType"]
				if hx_field_143 == nil {
					var hx_zero_144 *string
					return hx_zero_144
				}
				return hx_field_143.(*string)
			}(self.file)), hxrt.StringFromLiteral(";size=")), func(hx_obj_145 map[string]any) int {
				hx_field_146 := hx_obj_145["size"]
				if hx_field_146 == nil {
					var hx_zero_147 int
					return hx_zero_147
				}
				return hx_field_146.(int)
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
	}, func(hx_caught_148 any) {
		error := hxrt.ExceptionCaught(hx_caught_148)
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
		hx_post_150 := _g
		_g = int(int32((_g + 1)))
		headerIndex := hx_post_150
		name := hxrt.StdString(hxrt.HttpExchangeHeaderName(exchange, headerIndex))
		normalized := hxrt.StringToLowerCaseStringPtr(name)
		valueCount := hxrt.HttpExchangeHeaderValueCount(exchange, headerIndex)
		values := hxrt.NewArray()
		_g_1 := 0
		_g1_1 := valueCount
		for _g_1 < _g1_1 {
			hx_post_151 := _g_1
			_g_1 = int(int32((_g_1 + 1)))
			valueIndex := hx_post_151
			values.Push(hxrt.StdString(hxrt.HttpExchangeHeaderValue(exchange, headerIndex, valueIndex)))
		}
		if values.Len() == 0 {
			continue
		}
		last := func(hx_value_153 any) *string {
			if hx_value_153 == nil {
				var hx_zero_154 *string
				return hx_zero_154
			}
			return hx_value_153.(*string)
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
		parameter := func(hx_value_155 any) map[string]any {
			if hx_value_155 == nil {
				var hx_zero_156 map[string]any
				return hx_zero_156
			}
			return hx_value_155.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		encoded.Push(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(StringTools_urlEncode(func(hx_obj_158 map[string]any) *string {
			hx_field_159 := hx_obj_158["name"]
			if hx_field_159 == nil {
				var hx_zero_160 *string
				return hx_zero_160
			}
			return hx_field_159.(*string)
		}(parameter)), hxrt.StringFromLiteral("=")), StringTools_urlEncode(func(hx_obj_161 map[string]any) *string {
			hx_field_162 := hx_obj_161["value"]
			if hx_field_162 == nil {
				var hx_zero_163 *string
				return hx_zero_163
			}
			return hx_field_162.(*string)
		}(parameter))))
	}
	return hxrt.StringJoinAny(encoded.Values(), hxrt.StringFromLiteral("&"))
}

var sys__Http_PROXY map[string]any = nil

func sys__Http_firstComma(value *string) int {
	_g := 0
	_g1 := hxrt.StringLengthStringPtr(value)
	for _g < _g1 {
		hx_post_164 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_164
		if func() int {
			var c any = hxrt.StringCharCodeAtAnyStringPtr(value, index)
			var hx_if_165 int
			if c == nil {
				hx_if_165 = -1
			} else {
				hx_if_165 = c.(int)
			}
			return hx_if_165
		}() == 44 {
			return index
		}
	}
	return -1
}

func sys__Http_hxrt_proxyDescriptor() *string {
	proxy := sys__Http_PROXY
	if (proxy == nil) || hxrt.StringEqualStringPtr(func(hx_obj_166 map[string]any) *string {
		hx_field_167 := hx_obj_166["host"]
		if hx_field_167 == nil {
			var hx_zero_168 *string
			return hx_zero_168
		}
		return hx_field_167.(*string)
	}(proxy), nil) {
		return hxrt.StringFromLiteral("null")
	}
	var hx_if_178 *string
	if func(hx_obj_169 map[string]any) map[string]any {
		hx_field_170 := hx_obj_169["auth"]
		if hx_field_170 == nil {
			var hx_zero_171 map[string]any
			return hx_zero_171
		}
		return hx_field_170.(map[string]any)
	}(proxy) == nil {
		hx_if_178 = nil
	} else {
		hx_if_178 = func(hx_obj_175 map[string]any) *string {
			hx_field_176 := hx_obj_175["user"]
			if hx_field_176 == nil {
				var hx_zero_177 *string
				return hx_zero_177
			}
			return hx_field_176.(*string)
		}(func(hx_obj_172 map[string]any) map[string]any {
			hx_field_173 := hx_obj_172["auth"]
			if hx_field_173 == nil {
				var hx_zero_174 map[string]any
				return hx_zero_174
			}
			return hx_field_173.(map[string]any)
		}(proxy))
	}
	user := hx_if_178
	var hx_if_188 *string
	if func(hx_obj_179 map[string]any) map[string]any {
		hx_field_180 := hx_obj_179["auth"]
		if hx_field_180 == nil {
			var hx_zero_181 map[string]any
			return hx_zero_181
		}
		return hx_field_180.(map[string]any)
	}(proxy) == nil {
		hx_if_188 = nil
	} else {
		hx_if_188 = func(hx_obj_185 map[string]any) *string {
			hx_field_186 := hx_obj_185["pass"]
			if hx_field_186 == nil {
				var hx_zero_187 *string
				return hx_zero_187
			}
			return hx_field_186.(*string)
		}(func(hx_obj_182 map[string]any) map[string]any {
			hx_field_183 := hx_obj_182["auth"]
			if hx_field_183 == nil {
				var hx_zero_184 map[string]any
				return hx_zero_184
			}
			return hx_field_183.(map[string]any)
		}(proxy))
	}
	pass := hx_if_188
	return hxrt.StdString(hxrt.HttpProxyDescriptor(func(hx_obj_189 map[string]any) *string {
		hx_field_190 := hx_obj_189["host"]
		if hx_field_190 == nil {
			var hx_zero_191 *string
			return hx_zero_191
		}
		return hx_field_190.(*string)
	}(proxy), func(hx_obj_192 map[string]any) int {
		hx_field_193 := hx_obj_192["port"]
		if hx_field_193 == nil {
			var hx_zero_194 int
			return hx_zero_194
		}
		return hx_field_193.(int)
	}(proxy), user, pass))
}

func sys__Http_normalizedMethod(method *string) *string {
	if hxrt.StringEqualStringPtr(method, nil) {
		return nil
	}
	normalized := hxrt.StringToUpperCaseStringPtr(method)
	var hx_if_195 *string
	if hxrt.StringEqualStringPtr(normalized, hxrt.StringFromLiteral("")) || hxrt.StringEqualStringPtr(normalized, hxrt.StringFromLiteral("NULL")) {
		hx_if_195 = nil
	} else {
		hx_if_195 = normalized
	}
	return hx_if_195
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
