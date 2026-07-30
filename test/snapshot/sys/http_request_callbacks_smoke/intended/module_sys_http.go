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
		}(parameter), StringTools_urlEncode(func(hx_obj_57 map[string]any) *string {
			hx_field_58 := hx_obj_57["name"]
			if hx_field_58 == nil {
				var hx_zero_59 *string
				return hx_zero_59
			}
			return hx_field_58.(*string)
		}(parameter)), StringTools_urlEncode(func(hx_obj_60 map[string]any) *string {
			hx_field_61 := hx_obj_60["value"]
			if hx_field_61 == nil {
				var hx_zero_62 *string
				return hx_zero_62
			}
			return hx_field_61.(*string)
		}(parameter)))
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
	if (proxy != nil) && !hxrt.StringEqualStringPtr(func(hx_obj_109 map[string]any) *string {
		hx_field_110 := hx_obj_109["host"]
		if hx_field_110 == nil {
			var hx_zero_111 *string
			return hx_zero_111
		}
		return hx_field_110.(*string)
	}(proxy), nil) {
		var hx_if_92 *string
		if func(hx_obj_83 map[string]any) map[string]any {
			hx_field_84 := hx_obj_83["auth"]
			if hx_field_84 == nil {
				var hx_zero_85 map[string]any
				return hx_zero_85
			}
			return hx_field_84.(map[string]any)
		}(proxy) == nil {
			hx_if_92 = nil
		} else {
			hx_if_92 = func(hx_obj_89 map[string]any) *string {
				hx_field_90 := hx_obj_89["user"]
				if hx_field_90 == nil {
					var hx_zero_91 *string
					return hx_zero_91
				}
				return hx_field_90.(*string)
			}(func(hx_obj_86 map[string]any) map[string]any {
				hx_field_87 := hx_obj_86["auth"]
				if hx_field_87 == nil {
					var hx_zero_88 map[string]any
					return hx_zero_88
				}
				return hx_field_87.(map[string]any)
			}(proxy))
		}
		user := hx_if_92
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
				hx_field_100 := hx_obj_99["pass"]
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
		pass := hx_if_102
		hxrt.HttpRequestSetProxy(request, func(hx_obj_103 map[string]any) *string {
			hx_field_104 := hx_obj_103["host"]
			if hx_field_104 == nil {
				var hx_zero_105 *string
				return hx_zero_105
			}
			return hx_field_104.(*string)
		}(proxy), func(hx_obj_106 map[string]any) int {
			hx_field_107 := hx_obj_106["port"]
			if hx_field_107 == nil {
				var hx_zero_108 int
				return hx_zero_108
			}
			return hx_field_107.(int)
		}(proxy), user, pass)
	}
	if sock != nil {
		hxrt.HttpRequestSetSocket(request, sock.handle)
	}
	exchange := hxrt.HttpRequestStartExchange(request)
	var hx_if_112 map[string]any
	if upload == nil {
		hx_if_112 = nil
	} else {
		hx_if_112 = self.__hx_this.pumpUpload(exchange, upload)
	}
	uploadResult := hx_if_112
	hxrt.HttpExchangeAwaitResponse(exchange)
	var hx_if_116 *string
	if uploadResult == nil {
		hx_if_116 = nil
	} else {
		hx_if_116 = func(hx_obj_113 map[string]any) *string {
			hx_field_114 := hx_obj_113["sourceError"]
			if hx_field_114 == nil {
				var hx_zero_115 *string
				return hx_zero_115
			}
			return hx_field_114.(*string)
		}(uploadResult)
	}
	sourceError := hx_if_116
	var hx_if_120 *string
	if uploadResult == nil {
		hx_if_120 = nil
	} else {
		hx_if_120 = func(hx_obj_117 map[string]any) *string {
			hx_field_118 := hx_obj_117["sinkError"]
			if hx_field_118 == nil {
				var hx_zero_119 *string
				return hx_zero_119
			}
			return hx_field_118.(*string)
		}(uploadResult)
	}
	sinkError := hx_if_120
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
					statusError := sys__Http_hxrt_statusError(status)
					if !hxrt.StringEqualStringPtr(statusError, nil) {
						hxrt.Throw(statusError)
					}
					api.__hx_this.close()
					completed = true
				}, func(hx_caught_121 any) {
					error := hxrt.ExceptionCaught(hx_caught_121)
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
		hx_obj_123 := map[string]any{}
		hx_obj_123["sourceError"] = nil
		hx_obj_123["sinkError"] = hxrt.StringFromLiteral("HTTP upload sink is unavailable")
		return hx_obj_123
	}
	remaining := func(hx_obj_124 map[string]any) int {
		hx_field_125 := hx_obj_124["size"]
		if hx_field_125 == nil {
			var hx_zero_126 int
			return hx_zero_126
		}
		return hx_field_125.(int)
	}(upload)
	var sourceError *string = nil
	var sinkError *string = nil
	for remaining > 0 {
		var hx_if_127 int
		if remaining > 32768 {
			hx_if_127 = 32768
		} else {
			hx_if_127 = remaining
		}
		requested := hx_if_127
		chunk := haxe__io__Bytes_alloc(requested)
		count := 0
		hxrt.TryCatch(func() {
			count = func(hx_obj_130 map[string]any) *haxe__io__Input {
				hx_field_131 := hx_obj_130["io"]
				if hx_field_131 == nil {
					var hx_zero_132 *haxe__io__Input
					return hx_zero_132
				}
				return hx_field_131.(*haxe__io__Input)
			}(upload).__hx_this.readBytes(chunk, 0, requested)
		}, func(hx_caught_128 any) {
			switch hx_typed_129 := hx_caught_128.(type) {
			case *haxe__io__Eof:
				hx_tmp := hx_typed_129
				_ = hx_tmp
				sourceError = hxrt.StringFromLiteral("Transfer aborted")
			default:
				error := hxrt.ExceptionCaught(hx_caught_128)
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
	hx_obj_133 := map[string]any{}
	hx_obj_133["sourceError"] = sourceError
	hx_obj_133["sinkError"] = sinkError
	return hx_obj_133
}

func (self *sys__Http) handleDataRequest(post bool, api *haxe__io__Output, method *string) {
	if !hxrt.StringEqualStringPtr(method, nil) && hxrt.StringEqualStringPtr(method, hxrt.StringFromLiteral("")) {
		func(hx_fn func(*string), hx_arg_0 *string) {
			if hx_fn == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
				return
			}
			hx_fn(hx_arg_0)
		}(self.onError, hxrt.StringFromLiteral("HTTP method must not be empty"))
		return
	}
	encoded := hxrt.StringSubstrStringPtr(self.url, hxrt.StringLengthStringPtr(hxrt.StringFromLiteral("data:")), 0, false)
	mediaType := hxrt.StringFromLiteral("text/plain")
	comma := sys__Http_firstComma(encoded)
	if comma >= 0 {
		if comma > 0 {
			mediaType = hxrt.StringSubstrStringPtr(encoded, 0, comma, true)
		}
		encoded = hxrt.StringSubstrStringPtr(encoded, int(int32((hxrt.Int32Wrap(comma) + hxrt.Int32Wrap(1)))), 0, false)
	}
	if self.file != nil {
		encoded = hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("multipart file="), func(hx_obj_134 map[string]any) *string {
			hx_field_135 := hx_obj_134["filename"]
			if hx_field_135 == nil {
				var hx_zero_136 *string
				return hx_zero_136
			}
			return hx_field_135.(*string)
		}(self.file)), hxrt.StringFromLiteral(";mime=")), func(hx_obj_137 map[string]any) *string {
			hx_field_138 := hx_obj_137["mimeType"]
			if hx_field_138 == nil {
				var hx_zero_139 *string
				return hx_zero_139
			}
			return hx_field_138.(*string)
		}(self.file)), hxrt.StringFromLiteral(";size=")), func(hx_obj_140 map[string]any) int {
			hx_field_141 := hx_obj_140["size"]
			if hx_field_141 == nil {
				var hx_zero_142 int
				return hx_zero_142
			}
			return hx_field_141.(int)
		}(self.file))
	} else {
		if self.postBytes != nil {
			encoded = self.postBytes.__hx_this.toString()
		} else {
			if !hxrt.StringEqualStringPtr(self.postData, nil) {
				encoded = self.postData
			} else {
				if post {
					encoded = self.__hx_this.encodedParameters()
				}
			}
		}
	}
	payloadText := StringTools_urlDecode(encoded)
	explicitMethod := sys__Http_explicitMethod(method)
	if !hxrt.StringEqualStringPtr(explicitMethod, nil) {
		payloadText = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(explicitMethod, hxrt.StringFromLiteral(" ")), payloadText)
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
	}, func(hx_caught_143 any) {
		error := hxrt.ExceptionCaught(hx_caught_143)
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
		hx_post_145 := _g
		_g = int(int32((_g + 1)))
		headerIndex := hx_post_145
		name := hxrt.StdString(hxrt.HttpExchangeHeaderName(exchange, headerIndex))
		normalized := hxrt.StringToLowerCaseStringPtr(name)
		valueCount := hxrt.HttpExchangeHeaderValueCount(exchange, headerIndex)
		values := hxrt.NewArray()
		_g_1 := 0
		_g1_1 := valueCount
		for _g_1 < _g1_1 {
			hx_post_146 := _g_1
			_g_1 = int(int32((_g_1 + 1)))
			valueIndex := hx_post_146
			values.Push(hxrt.StdString(hxrt.HttpExchangeHeaderValue(exchange, headerIndex, valueIndex)))
		}
		if values.Len() == 0 {
			continue
		}
		last := func(hx_value_148 any) *string {
			if hx_value_148 == nil {
				var hx_zero_149 *string
				return hx_zero_149
			}
			return hx_value_148.(*string)
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
		parameter := func(hx_value_150 any) map[string]any {
			if hx_value_150 == nil {
				var hx_zero_151 map[string]any
				return hx_zero_151
			}
			return hx_value_150.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		encoded.Push(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(StringTools_urlEncode(func(hx_obj_153 map[string]any) *string {
			hx_field_154 := hx_obj_153["name"]
			if hx_field_154 == nil {
				var hx_zero_155 *string
				return hx_zero_155
			}
			return hx_field_154.(*string)
		}(parameter)), hxrt.StringFromLiteral("=")), StringTools_urlEncode(func(hx_obj_156 map[string]any) *string {
			hx_field_157 := hx_obj_156["value"]
			if hx_field_157 == nil {
				var hx_zero_158 *string
				return hx_zero_158
			}
			return hx_field_157.(*string)
		}(parameter))))
	}
	return hxrt.StringJoinAny(encoded.Values(), hxrt.StringFromLiteral("&"))
}

var sys__Http_PROXY map[string]any = nil

func sys__Http_explicitMethod(method *string) *string {
	if hxrt.StringEqualStringPtr(method, nil) || hxrt.StringEqualStringPtr(method, hxrt.StringFromLiteral("")) {
		return nil
	}
	return method
}

func sys__Http_firstComma(value *string) int {
	_g := 0
	_g1 := hxrt.StringLengthStringPtr(value)
	for _g < _g1 {
		hx_post_159 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_159
		if func() int {
			var c any = hxrt.StringCharCodeAtAnyStringPtr(value, index)
			var hx_if_160 int
			if c == nil {
				hx_if_160 = -1
			} else {
				hx_if_160 = c.(int)
			}
			return hx_if_160
		}() == 44 {
			return index
		}
	}
	return -1
}

func sys__Http_hxrt_proxyDescriptor() *string {
	proxy := sys__Http_PROXY
	if (proxy == nil) || hxrt.StringEqualStringPtr(func(hx_obj_161 map[string]any) *string {
		hx_field_162 := hx_obj_161["host"]
		if hx_field_162 == nil {
			var hx_zero_163 *string
			return hx_zero_163
		}
		return hx_field_162.(*string)
	}(proxy), nil) {
		return hxrt.StringFromLiteral("null")
	}
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
			hx_field_171 := hx_obj_170["user"]
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
	user := hx_if_173
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
			hx_field_181 := hx_obj_180["pass"]
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
	pass := hx_if_183
	return hxrt.StdString(hxrt.HttpProxyDescriptor(func(hx_obj_184 map[string]any) *string {
		hx_field_185 := hx_obj_184["host"]
		if hx_field_185 == nil {
			var hx_zero_186 *string
			return hx_zero_186
		}
		return hx_field_185.(*string)
	}(proxy), func(hx_obj_187 map[string]any) int {
		hx_field_188 := hx_obj_187["port"]
		if hx_field_188 == nil {
			var hx_zero_189 int
			return hx_zero_189
		}
		return hx_field_188.(int)
	}(proxy), user, pass))
}

func sys__Http_hxrt_statusError(status int) *string {
	var hx_if_190 *string
	if (status < 200) || (status >= 400) {
		hx_if_190 = hxrt.StringConcatAny(hxrt.StringFromLiteral("Http Error #"), status)
	} else {
		hx_if_190 = nil
	}
	return hx_if_190
}

func sys__Http_requestUrl(url *string) *string {
	request := New_sys__Http(url)
	result := hxrt.StringFromLiteral("")
	request.onData = func(data *string) {
		result = data
	}
	request.onError = func(message *string) {
		hxrt.Throw(message)
	}
	request.__hx_this.request(false)
	return result
}
