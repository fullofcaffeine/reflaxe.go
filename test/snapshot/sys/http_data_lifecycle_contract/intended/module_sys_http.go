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
	hx_obj_1 := map[string]any{}
	hx_obj_1["param"] = argname
	hx_obj_1["filename"] = filename
	hx_obj_1["io"] = file
	hx_obj_1["size"] = size
	hx_obj_1["mimeType"] = mimeType
	self.file = hx_obj_1
}

func (self *sys__Http) customRequest(post bool, api *haxe__io__Output, sock *sys__net__Socket, method *string) {
	self.__hx_this.requestWith((post || (self.file != nil)), api, sock, method)
}

func (self *sys__Http) getResponseHeaderValues(key *string) *hxrt.Array {
	values := func(hx_value_2 any) *hxrt.Array {
		if hx_value_2 == nil {
			var hx_zero_3 *hxrt.Array
			return hx_zero_3
		}
		return hx_value_2.(*hxrt.Array)
	}(self.responseHeadersSameKey.__hx_this.get(key))
	if values == nil {
		normalized := hxrt.StringToLowerCaseStringPtr(key)
		if !hxrt.StringEqualStringPtr(normalized, key) {
			values = func(hx_value_4 any) *hxrt.Array {
				if hx_value_4 == nil {
					var hx_zero_5 *hxrt.Array
					return hx_zero_5
				}
				return hx_value_4.(*hxrt.Array)
			}(self.responseHeadersSameKey.__hx_this.get(normalized))
		}
	}
	if values != nil {
		return values
	}
	var this1 haxe__IMap = self.responseHeaders
	value := func(hx_value_6 any) *string {
		if hx_value_6 == nil {
			var hx_zero_7 *string
			return hx_zero_7
		}
		return hx_value_6.(*string)
	}(this1.(*haxe__ds__StringMap).__hx_this.get(key))
	if hxrt.StringEqualStringPtr(value, nil) {
		normalized_1 := hxrt.StringToLowerCaseStringPtr(key)
		if !hxrt.StringEqualStringPtr(normalized_1, key) {
			var this1_1 haxe__IMap = self.responseHeaders
			value = func(hx_value_8 any) *string {
				if hx_value_8 == nil {
					var hx_zero_9 *string
					return hx_zero_9
				}
				return hx_value_8.(*string)
			}(this1_1.(*haxe__ds__StringMap).__hx_this.get(normalized_1))
		}
	}
	var hx_if_10 *hxrt.Array
	if hxrt.StringEqualStringPtr(value, nil) {
		hx_if_10 = nil
	} else {
		hx_if_10 = hxrt.NewArray(value)
	}
	return hx_if_10
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
		parameter := func(hx_value_11 any) map[string]any {
			if hx_value_11 == nil {
				var hx_zero_12 map[string]any
				return hx_zero_12
			}
			return hx_value_11.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		hxrt.HttpRequestAddParameter(request, func(hx_obj_13 map[string]any) *string {
			hx_field_14 := hx_obj_13["name"]
			if hx_field_14 == nil {
				var hx_zero_15 *string
				return hx_zero_15
			}
			return hx_field_14.(*string)
		}(parameter), func(hx_obj_16 map[string]any) *string {
			hx_field_17 := hx_obj_16["value"]
			if hx_field_17 == nil {
				var hx_zero_18 *string
				return hx_zero_18
			}
			return hx_field_17.(*string)
		}(parameter), StringTools_urlEncode(func(hx_obj_19 map[string]any) *string {
			hx_field_20 := hx_obj_19["name"]
			if hx_field_20 == nil {
				var hx_zero_21 *string
				return hx_zero_21
			}
			return hx_field_20.(*string)
		}(parameter)), StringTools_urlEncode(func(hx_obj_22 map[string]any) *string {
			hx_field_23 := hx_obj_22["value"]
			if hx_field_23 == nil {
				var hx_zero_24 *string
				return hx_zero_24
			}
			return hx_field_23.(*string)
		}(parameter)))
	}
	_g_1 := 0
	_g1_1 := self.headers
	for _g_1 < _g1_1.Len() {
		header := func(hx_value_25 any) map[string]any {
			if hx_value_25 == nil {
				var hx_zero_26 map[string]any
				return hx_zero_26
			}
			return hx_value_25.(map[string]any)
		}(_g1_1.Get(_g_1))
		_g_1 = int(int32((_g_1 + 1)))
		hxrt.HttpRequestAddHeader(request, func(hx_obj_27 map[string]any) *string {
			hx_field_28 := hx_obj_27["name"]
			if hx_field_28 == nil {
				var hx_zero_29 *string
				return hx_zero_29
			}
			return hx_field_28.(*string)
		}(header), func(hx_obj_30 map[string]any) *string {
			hx_field_31 := hx_obj_30["value"]
			if hx_field_31 == nil {
				var hx_zero_32 *string
				return hx_zero_32
			}
			return hx_field_31.(*string)
		}(header))
	}
	upload := self.file
	if upload != nil {
		hxrt.HttpRequestSetMultipartUpload(request, func(hx_obj_33 map[string]any) *string {
			hx_field_34 := hx_obj_33["param"]
			if hx_field_34 == nil {
				var hx_zero_35 *string
				return hx_zero_35
			}
			return hx_field_34.(*string)
		}(upload), func(hx_obj_36 map[string]any) *string {
			hx_field_37 := hx_obj_36["filename"]
			if hx_field_37 == nil {
				var hx_zero_38 *string
				return hx_zero_38
			}
			return hx_field_37.(*string)
		}(upload), func(hx_obj_39 map[string]any) *string {
			hx_field_40 := hx_obj_39["mimeType"]
			if hx_field_40 == nil {
				var hx_zero_41 *string
				return hx_zero_41
			}
			return hx_field_40.(*string)
		}(upload), func(hx_obj_42 map[string]any) int {
			hx_field_43 := hx_obj_42["size"]
			if hx_field_43 == nil {
				var hx_zero_44 int
				return hx_zero_44
			}
			return hx_field_43.(int)
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
	if (proxy != nil) && !hxrt.StringEqualStringPtr(func(hx_obj_71 map[string]any) *string {
		hx_field_72 := hx_obj_71["host"]
		if hx_field_72 == nil {
			var hx_zero_73 *string
			return hx_zero_73
		}
		return hx_field_72.(*string)
	}(proxy), nil) {
		var hx_if_54 *string
		if func(hx_obj_45 map[string]any) map[string]any {
			hx_field_46 := hx_obj_45["auth"]
			if hx_field_46 == nil {
				var hx_zero_47 map[string]any
				return hx_zero_47
			}
			return hx_field_46.(map[string]any)
		}(proxy) == nil {
			hx_if_54 = nil
		} else {
			hx_if_54 = func(hx_obj_51 map[string]any) *string {
				hx_field_52 := hx_obj_51["user"]
				if hx_field_52 == nil {
					var hx_zero_53 *string
					return hx_zero_53
				}
				return hx_field_52.(*string)
			}(func(hx_obj_48 map[string]any) map[string]any {
				hx_field_49 := hx_obj_48["auth"]
				if hx_field_49 == nil {
					var hx_zero_50 map[string]any
					return hx_zero_50
				}
				return hx_field_49.(map[string]any)
			}(proxy))
		}
		user := hx_if_54
		var hx_if_64 *string
		if func(hx_obj_55 map[string]any) map[string]any {
			hx_field_56 := hx_obj_55["auth"]
			if hx_field_56 == nil {
				var hx_zero_57 map[string]any
				return hx_zero_57
			}
			return hx_field_56.(map[string]any)
		}(proxy) == nil {
			hx_if_64 = nil
		} else {
			hx_if_64 = func(hx_obj_61 map[string]any) *string {
				hx_field_62 := hx_obj_61["pass"]
				if hx_field_62 == nil {
					var hx_zero_63 *string
					return hx_zero_63
				}
				return hx_field_62.(*string)
			}(func(hx_obj_58 map[string]any) map[string]any {
				hx_field_59 := hx_obj_58["auth"]
				if hx_field_59 == nil {
					var hx_zero_60 map[string]any
					return hx_zero_60
				}
				return hx_field_59.(map[string]any)
			}(proxy))
		}
		pass := hx_if_64
		hxrt.HttpRequestSetProxy(request, func(hx_obj_65 map[string]any) *string {
			hx_field_66 := hx_obj_65["host"]
			if hx_field_66 == nil {
				var hx_zero_67 *string
				return hx_zero_67
			}
			return hx_field_66.(*string)
		}(proxy), func(hx_obj_68 map[string]any) int {
			hx_field_69 := hx_obj_68["port"]
			if hx_field_69 == nil {
				var hx_zero_70 int
				return hx_zero_70
			}
			return hx_field_69.(int)
		}(proxy), user, pass)
	}
	if sock != nil {
		hxrt.HttpRequestSetSocket(request, sock.handle)
	}
	exchange := hxrt.HttpRequestStartExchange(request)
	var hx_if_74 map[string]any
	if upload == nil {
		hx_if_74 = nil
	} else {
		hx_if_74 = self.__hx_this.pumpUpload(exchange, upload)
	}
	uploadResult := hx_if_74
	hxrt.HttpExchangeAwaitResponse(exchange)
	var hx_if_78 *string
	if uploadResult == nil {
		hx_if_78 = nil
	} else {
		hx_if_78 = func(hx_obj_75 map[string]any) *string {
			hx_field_76 := hx_obj_75["sourceError"]
			if hx_field_76 == nil {
				var hx_zero_77 *string
				return hx_zero_77
			}
			return hx_field_76.(*string)
		}(uploadResult)
	}
	sourceError := hx_if_78
	var hx_if_82 *string
	if uploadResult == nil {
		hx_if_82 = nil
	} else {
		hx_if_82 = func(hx_obj_79 map[string]any) *string {
			hx_field_80 := hx_obj_79["sinkError"]
			if hx_field_80 == nil {
				var hx_zero_81 *string
				return hx_zero_81
			}
			return hx_field_80.(*string)
		}(uploadResult)
	}
	sinkError := hx_if_82
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
				}, func(hx_caught_83 any) {
					error := hxrt.ExceptionCaught(hx_caught_83)
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
		hx_obj_85 := map[string]any{}
		hx_obj_85["sourceError"] = nil
		hx_obj_85["sinkError"] = hxrt.StringFromLiteral("HTTP upload sink is unavailable")
		return hx_obj_85
	}
	remaining := func(hx_obj_86 map[string]any) int {
		hx_field_87 := hx_obj_86["size"]
		if hx_field_87 == nil {
			var hx_zero_88 int
			return hx_zero_88
		}
		return hx_field_87.(int)
	}(upload)
	var sourceError *string = nil
	var sinkError *string = nil
	for remaining > 0 {
		var hx_if_89 int
		if remaining > 32768 {
			hx_if_89 = 32768
		} else {
			hx_if_89 = remaining
		}
		requested := hx_if_89
		chunk := haxe__io__Bytes_alloc(requested)
		count := 0
		hxrt.TryCatch(func() {
			count = func(hx_obj_92 map[string]any) *haxe__io__Input {
				hx_field_93 := hx_obj_92["io"]
				if hx_field_93 == nil {
					var hx_zero_94 *haxe__io__Input
					return hx_zero_94
				}
				return hx_field_93.(*haxe__io__Input)
			}(upload).__hx_this.readBytes(chunk, 0, requested)
		}, func(hx_caught_90 any) {
			switch hx_typed_91 := hx_caught_90.(type) {
			case *haxe__io__Eof:
				hx_tmp := hx_typed_91
				_ = hx_tmp
				sourceError = hxrt.StringFromLiteral("Transfer aborted")
			default:
				error := hxrt.ExceptionCaught(hx_caught_90)
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
	hx_obj_95 := map[string]any{}
	hx_obj_95["sourceError"] = sourceError
	hx_obj_95["sinkError"] = sinkError
	return hx_obj_95
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
		encoded = hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("multipart file="), func(hx_obj_96 map[string]any) *string {
			hx_field_97 := hx_obj_96["filename"]
			if hx_field_97 == nil {
				var hx_zero_98 *string
				return hx_zero_98
			}
			return hx_field_97.(*string)
		}(self.file)), hxrt.StringFromLiteral(";mime=")), func(hx_obj_99 map[string]any) *string {
			hx_field_100 := hx_obj_99["mimeType"]
			if hx_field_100 == nil {
				var hx_zero_101 *string
				return hx_zero_101
			}
			return hx_field_100.(*string)
		}(self.file)), hxrt.StringFromLiteral(";size=")), func(hx_obj_102 map[string]any) int {
			hx_field_103 := hx_obj_102["size"]
			if hx_field_103 == nil {
				var hx_zero_104 int
				return hx_zero_104
			}
			return hx_field_103.(int)
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
	}, func(hx_caught_105 any) {
		error := hxrt.ExceptionCaught(hx_caught_105)
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
		hx_post_107 := _g
		_g = int(int32((_g + 1)))
		headerIndex := hx_post_107
		name := hxrt.StdString(hxrt.HttpExchangeHeaderName(exchange, headerIndex))
		normalized := hxrt.StringToLowerCaseStringPtr(name)
		valueCount := hxrt.HttpExchangeHeaderValueCount(exchange, headerIndex)
		values := hxrt.NewArray()
		_g_1 := 0
		_g1_1 := valueCount
		for _g_1 < _g1_1 {
			hx_post_108 := _g_1
			_g_1 = int(int32((_g_1 + 1)))
			valueIndex := hx_post_108
			values.Push(hxrt.StdString(hxrt.HttpExchangeHeaderValue(exchange, headerIndex, valueIndex)))
		}
		if values.Len() == 0 {
			continue
		}
		last := func(hx_value_110 any) *string {
			if hx_value_110 == nil {
				var hx_zero_111 *string
				return hx_zero_111
			}
			return hx_value_110.(*string)
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
		parameter := func(hx_value_112 any) map[string]any {
			if hx_value_112 == nil {
				var hx_zero_113 map[string]any
				return hx_zero_113
			}
			return hx_value_112.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		encoded.Push(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(StringTools_urlEncode(func(hx_obj_115 map[string]any) *string {
			hx_field_116 := hx_obj_115["name"]
			if hx_field_116 == nil {
				var hx_zero_117 *string
				return hx_zero_117
			}
			return hx_field_116.(*string)
		}(parameter)), hxrt.StringFromLiteral("=")), StringTools_urlEncode(func(hx_obj_118 map[string]any) *string {
			hx_field_119 := hx_obj_118["value"]
			if hx_field_119 == nil {
				var hx_zero_120 *string
				return hx_zero_120
			}
			return hx_field_119.(*string)
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
		hx_post_121 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_121
		if func() int {
			var c any = hxrt.StringCharCodeAtAnyStringPtr(value, index)
			var hx_if_122 int
			if c == nil {
				hx_if_122 = -1
			} else {
				hx_if_122 = c.(int)
			}
			return hx_if_122
		}() == 44 {
			return index
		}
	}
	return -1
}

func sys__Http_hxrt_proxyDescriptor() *string {
	proxy := sys__Http_PROXY
	if (proxy == nil) || hxrt.StringEqualStringPtr(func(hx_obj_123 map[string]any) *string {
		hx_field_124 := hx_obj_123["host"]
		if hx_field_124 == nil {
			var hx_zero_125 *string
			return hx_zero_125
		}
		return hx_field_124.(*string)
	}(proxy), nil) {
		return hxrt.StringFromLiteral("null")
	}
	var hx_if_135 *string
	if func(hx_obj_126 map[string]any) map[string]any {
		hx_field_127 := hx_obj_126["auth"]
		if hx_field_127 == nil {
			var hx_zero_128 map[string]any
			return hx_zero_128
		}
		return hx_field_127.(map[string]any)
	}(proxy) == nil {
		hx_if_135 = nil
	} else {
		hx_if_135 = func(hx_obj_132 map[string]any) *string {
			hx_field_133 := hx_obj_132["user"]
			if hx_field_133 == nil {
				var hx_zero_134 *string
				return hx_zero_134
			}
			return hx_field_133.(*string)
		}(func(hx_obj_129 map[string]any) map[string]any {
			hx_field_130 := hx_obj_129["auth"]
			if hx_field_130 == nil {
				var hx_zero_131 map[string]any
				return hx_zero_131
			}
			return hx_field_130.(map[string]any)
		}(proxy))
	}
	user := hx_if_135
	var hx_if_145 *string
	if func(hx_obj_136 map[string]any) map[string]any {
		hx_field_137 := hx_obj_136["auth"]
		if hx_field_137 == nil {
			var hx_zero_138 map[string]any
			return hx_zero_138
		}
		return hx_field_137.(map[string]any)
	}(proxy) == nil {
		hx_if_145 = nil
	} else {
		hx_if_145 = func(hx_obj_142 map[string]any) *string {
			hx_field_143 := hx_obj_142["pass"]
			if hx_field_143 == nil {
				var hx_zero_144 *string
				return hx_zero_144
			}
			return hx_field_143.(*string)
		}(func(hx_obj_139 map[string]any) map[string]any {
			hx_field_140 := hx_obj_139["auth"]
			if hx_field_140 == nil {
				var hx_zero_141 map[string]any
				return hx_zero_141
			}
			return hx_field_140.(map[string]any)
		}(proxy))
	}
	pass := hx_if_145
	return hxrt.StdString(hxrt.HttpProxyDescriptor(func(hx_obj_146 map[string]any) *string {
		hx_field_147 := hx_obj_146["host"]
		if hx_field_147 == nil {
			var hx_zero_148 *string
			return hx_zero_148
		}
		return hx_field_147.(*string)
	}(proxy), func(hx_obj_149 map[string]any) int {
		hx_field_150 := hx_obj_149["port"]
		if hx_field_150 == nil {
			var hx_zero_151 int
			return hx_zero_151
		}
		return hx_field_150.(int)
	}(proxy), user, pass))
}

func sys__Http_hxrt_statusError(status int) *string {
	var hx_if_152 *string
	if (status < 200) || (status >= 400) {
		hx_if_152 = hxrt.StringConcatAny(hxrt.StringFromLiteral("Http Error #"), status)
	} else {
		hx_if_152 = nil
	}
	return hx_if_152
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
