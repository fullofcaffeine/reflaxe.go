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
	pumpUpload(exchange *hxrt.HttpExchange, upload hxrt.Anonymous) hxrt.Anonymous
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
	file                   hxrt.Anonymous
}

func (self *sys__Http) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_sys__Http(url *string) *sys__Http {
	return hx_init_sys__Http(hx_alloc_sys__Http(nil), url)
}

func hx_alloc_sys__Http(owner any) *sys__Http {
	self := &sys__Http{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_sys__Http)
	}
	self.haxe__http__HttpBase = hx_alloc_haxe__http__HttpBase(self.__hx_this)
	return self
}

func hx_init_sys__Http(self *sys__Http, url *string) *sys__Http {
	hx_init_haxe__http__HttpBase(self.haxe__http__HttpBase, url)
	self.noShutdown = false
	self.cnxTimeout = 10
	self.__hx_this.resetResponseHeaders()
	return self
}

func (self *sys__Http) request(post any) {
	_gthis := self
	output := New_haxe__io__BytesOutput()
	previousOnError := hxrt.ReadDynamicFunction(self.onError)
	failed := false
	self.onError = hxrt.NewFunction(func(message *string) {
		_gthis.responseBytes = output.__hx_this.getBytes()
		_gthis.responseAsString = nil
		failed = true
		_gthis.onError = previousOnError
		func() {
			var hx_callable_1 *hxrt.Function[func(*string)] = _gthis.onError
			var hx_argument_2 *string = message
			hxrt.FunctionCall(hx_callable_1)(hx_argument_2)
		}()
	})
	usePost := (((((post != nil) && (post.(bool) == true)) || (self.postBytes != nil)) || !hxrt.StringEqualStringPtr(self.postData, nil)) || (self.file != nil))
	self.__hx_this.customRequest(usePost, func(hx_value *haxe__io__BytesOutput) *haxe__io__Output {
		if hx_value == nil {
			return nil
		}
		return hx_value.haxe__io__Output
	}(output), nil, nil)
	if !failed {
		self.__hx_this.success(output.__hx_this.getBytes())
	}
}

func (self *sys__Http) fileTransfert(argname *string, filename *string, file *haxe__io__Input, size int, mimeType *string) {
	self.__hx_this.fileTransfer(argname, filename, file, size, mimeType)
}

func (self *sys__Http) fileTransfer(argname *string, filename *string, file *haxe__io__Input, size int, mimeType *string) {
	hx_obj_3 := map[string]any{}
	hx_obj_3["param"] = argname
	hx_obj_3["filename"] = filename
	hx_obj_3["io"] = file
	hx_obj_3["size"] = size
	hx_obj_3["mimeType"] = mimeType
	self.file = hxrt.AnonymousFromMap(hx_obj_3)
}

func (self *sys__Http) customRequest(post bool, api *haxe__io__Output, sock *sys__net__Socket, method *string) {
	self.__hx_this.requestWith((post || (self.file != nil)), api, sock, method)
}

func (self *sys__Http) getResponseHeaderValues(key *string) *hxrt.Array {
	values := func(hx_value_4 any) *hxrt.Array {
		if hx_value_4 == nil {
			var hx_zero_5 *hxrt.Array
			return hx_zero_5
		}
		return hx_value_4.(*hxrt.Array)
	}(self.responseHeadersSameKey.__hx_this.get(key))
	if values == nil {
		normalized := hxrt.StringToLowerCaseStringPtr(key)
		if !hxrt.StringEqualStringPtr(normalized, key) {
			values = func(hx_value_6 any) *hxrt.Array {
				if hx_value_6 == nil {
					var hx_zero_7 *hxrt.Array
					return hx_zero_7
				}
				return hx_value_6.(*hxrt.Array)
			}(self.responseHeadersSameKey.__hx_this.get(normalized))
		}
	}
	if values != nil {
		return values
	}
	var this1 haxe__IMap = self.responseHeaders
	value := func(hx_value_8 any) *string {
		if hx_value_8 == nil {
			var hx_zero_9 *string
			return hx_zero_9
		}
		return hx_value_8.(*string)
	}(this1.(*haxe__ds__StringMap).__hx_this.get(key))
	if hxrt.StringEqualStringPtr(value, nil) {
		normalized_1 := hxrt.StringToLowerCaseStringPtr(key)
		if !hxrt.StringEqualStringPtr(normalized_1, key) {
			var this1_1 haxe__IMap = self.responseHeaders
			value = func(hx_value_10 any) *string {
				if hx_value_10 == nil {
					var hx_zero_11 *string
					return hx_zero_11
				}
				return hx_value_10.(*string)
			}(this1_1.(*haxe__ds__StringMap).__hx_this.get(normalized_1))
		}
	}
	var hx_if_12 *hxrt.Array
	if hxrt.StringEqualStringPtr(value, nil) {
		hx_if_12 = nil
	} else {
		hx_if_12 = hxrt.NewArray(value)
	}
	return hx_if_12
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
		parameter := hx_structural_recover(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		hxrt.HttpRequestAddParameter(request, func(hx_value_13 any) *string {
			if hx_value_13 == nil {
				var hx_zero_14 *string
				return hx_zero_14
			}
			return hx_value_13.(*string)
		}(hxrt.AnonymousGet(parameter, "name")), func(hx_value_15 any) *string {
			if hx_value_15 == nil {
				var hx_zero_16 *string
				return hx_zero_16
			}
			return hx_value_15.(*string)
		}(hxrt.AnonymousGet(parameter, "value")), StringTools_urlEncode(func(hx_value_17 any) *string {
			if hx_value_17 == nil {
				var hx_zero_18 *string
				return hx_zero_18
			}
			return hx_value_17.(*string)
		}(hxrt.AnonymousGet(parameter, "name"))), StringTools_urlEncode(func(hx_value_19 any) *string {
			if hx_value_19 == nil {
				var hx_zero_20 *string
				return hx_zero_20
			}
			return hx_value_19.(*string)
		}(hxrt.AnonymousGet(parameter, "value"))))
	}
	_g_1 := 0
	_g1_1 := self.headers
	for _g_1 < _g1_1.Len() {
		header := hx_structural_recover(_g1_1.Get(_g_1))
		_g_1 = int(int32((_g_1 + 1)))
		hxrt.HttpRequestAddHeader(request, func(hx_value_21 any) *string {
			if hx_value_21 == nil {
				var hx_zero_22 *string
				return hx_zero_22
			}
			return hx_value_21.(*string)
		}(hxrt.AnonymousGet(header, "name")), func(hx_value_23 any) *string {
			if hx_value_23 == nil {
				var hx_zero_24 *string
				return hx_zero_24
			}
			return hx_value_23.(*string)
		}(hxrt.AnonymousGet(header, "value")))
	}
	upload := self.file
	if upload != nil {
		hxrt.HttpRequestSetMultipartUpload(request, func(hx_value_25 any) *string {
			if hx_value_25 == nil {
				var hx_zero_26 *string
				return hx_zero_26
			}
			return hx_value_25.(*string)
		}(hxrt.AnonymousGet(upload, "param")), func(hx_value_27 any) *string {
			if hx_value_27 == nil {
				var hx_zero_28 *string
				return hx_zero_28
			}
			return hx_value_27.(*string)
		}(hxrt.AnonymousGet(upload, "filename")), func(hx_value_29 any) *string {
			if hx_value_29 == nil {
				var hx_zero_30 *string
				return hx_zero_30
			}
			return hx_value_29.(*string)
		}(hxrt.AnonymousGet(upload, "mimeType")), func(hx_value_31 any) int {
			if hx_value_31 == nil {
				var hx_zero_32 int
				return hx_zero_32
			}
			return hx_value_31.(int)
		}(hxrt.AnonymousGet(upload, "size")))
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
	if (proxy != nil) && !hxrt.StringEqualStringPtr(func(hx_value_43 any) *string {
		if hx_value_43 == nil {
			var hx_zero_44 *string
			return hx_zero_44
		}
		return hx_value_43.(*string)
	}(hxrt.AnonymousGet(proxy, "host")), nil) {
		var hx_if_35 *string
		if hx_structural_recover(hxrt.AnonymousGet(proxy, "auth")) == nil {
			hx_if_35 = nil
		} else {
			hx_if_35 = func(hx_value_33 any) *string {
				if hx_value_33 == nil {
					var hx_zero_34 *string
					return hx_zero_34
				}
				return hx_value_33.(*string)
			}(hxrt.AnonymousGet(hx_structural_recover(hxrt.AnonymousGet(proxy, "auth")), "user"))
		}
		user := hx_if_35
		var hx_if_38 *string
		if hx_structural_recover(hxrt.AnonymousGet(proxy, "auth")) == nil {
			hx_if_38 = nil
		} else {
			hx_if_38 = func(hx_value_36 any) *string {
				if hx_value_36 == nil {
					var hx_zero_37 *string
					return hx_zero_37
				}
				return hx_value_36.(*string)
			}(hxrt.AnonymousGet(hx_structural_recover(hxrt.AnonymousGet(proxy, "auth")), "pass"))
		}
		pass := hx_if_38
		hxrt.HttpRequestSetProxy(request, func(hx_value_39 any) *string {
			if hx_value_39 == nil {
				var hx_zero_40 *string
				return hx_zero_40
			}
			return hx_value_39.(*string)
		}(hxrt.AnonymousGet(proxy, "host")), func(hx_value_41 any) int {
			if hx_value_41 == nil {
				var hx_zero_42 int
				return hx_zero_42
			}
			return hx_value_41.(int)
		}(hxrt.AnonymousGet(proxy, "port")), user, pass)
	}
	if sock != nil {
		hxrt.HttpRequestSetSocket(request, sock.handle)
	}
	exchange := hxrt.HttpRequestStartExchange(request)
	var hx_if_45 hxrt.Anonymous
	if upload == nil {
		hx_if_45 = nil
	} else {
		hx_if_45 = self.__hx_this.pumpUpload(exchange, upload)
	}
	uploadResult := hx_if_45
	hxrt.HttpExchangeAwaitResponse(exchange)
	var hx_if_48 *string
	if uploadResult == nil {
		hx_if_48 = nil
	} else {
		hx_if_48 = func(hx_value_46 any) *string {
			if hx_value_46 == nil {
				var hx_zero_47 *string
				return hx_zero_47
			}
			return hx_value_46.(*string)
		}(hxrt.AnonymousGet(uploadResult, "sourceError"))
	}
	sourceError := hx_if_48
	var hx_if_51 *string
	if uploadResult == nil {
		hx_if_51 = nil
	} else {
		hx_if_51 = func(hx_value_49 any) *string {
			if hx_value_49 == nil {
				var hx_zero_50 *string
				return hx_zero_50
			}
			return hx_value_49.(*string)
		}(hxrt.AnonymousGet(uploadResult, "sinkError"))
	}
	sinkError := hx_if_51
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
					func() {
						var hx_callable_55 *hxrt.Function[func(int)] = self.onStatus
						var hx_argument_56 int = status
						hxrt.FunctionCall(hx_callable_55)(hx_argument_56)
					}()
					contentLength := hxrt.HttpExchangeContentLength(exchange)
					if contentLength == -2 {
						hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("Content-Length exceeds Haxe Int range")))
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
							hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("Transfer aborted")))
						}
						if hxrt.HttpReadResultEOF(read) {
							break
						}
					}
					statusError := sys__Http_hxrt_statusError(status)
					if !hxrt.StringEqualStringPtr(statusError, nil) {
						hxrt.Throw(haxe__Exception_thrown(statusError))
					}
					api.__hx_this.close()
					completed = true
				}, func(hx_caught_52 any) {
					hx_exception_53 := haxe__Exception_caught(hx_caught_52)
					hx_payload_54 := haxe__Exception_catchValue(hx_caught_52)
					_ = hx_payload_54
					var error *haxe__Exception = hx_recover_class_2(hx_exception_53)
					errorMessage = error.__hx_this.get_message()
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
		func() {
			var hx_callable_57 *hxrt.Function[func(*string)] = self.onError
			var hx_argument_58 *string = errorMessage
			hxrt.FunctionCall(hx_callable_57)(hx_argument_58)
		}()
	}
}

func (self *sys__Http) pumpUpload(exchange *hxrt.HttpExchange, upload hxrt.Anonymous) hxrt.Anonymous {
	sink := hxrt.HttpExchangeUploadSink(exchange)
	if sink == nil {
		hx_obj_59 := map[string]any{}
		hx_obj_59["sourceError"] = nil
		hx_obj_59["sinkError"] = hxrt.StringFromLiteral("HTTP upload sink is unavailable")
		return hxrt.AnonymousFromMap(hx_obj_59)
	}
	remaining := func(hx_value_60 any) int {
		if hx_value_60 == nil {
			var hx_zero_61 int
			return hx_zero_61
		}
		return hx_value_60.(int)
	}(hxrt.AnonymousGet(upload, "size"))
	var sourceError *string = nil
	var sinkError *string = nil
	for remaining > 0 {
		var hx_if_62 int
		if remaining > 32768 {
			hx_if_62 = 32768
		} else {
			hx_if_62 = remaining
		}
		requested := hx_if_62
		chunk := haxe__io__Bytes_alloc(requested)
		count := 0
		hxrt.TryCatch(func() {
			count = hx_recover_class_4(hxrt.AnonymousGet(upload, "io")).__hx_this.readBytes(chunk, 0, requested)
		}, func(hx_caught_63 any) {
			hx_exception_64 := haxe__Exception_caught(hx_caught_63)
			hx_payload_65 := haxe__Exception_catchValue(hx_caught_63)
			if hx_is_nominal_0(hx_payload_65) {
				var hx_tmp *haxe__io__Eof = hx_recover_class_3(hx_payload_65)
				_ = hx_tmp
				sourceError = hxrt.StringFromLiteral("Transfer aborted")
			} else {
				var error *haxe__Exception = hx_recover_class_2(hx_exception_64)
				sourceError = error.__hx_this.get_message()
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
		remaining = int((hxrt.Int32Wrap(remaining) - hxrt.Int32Wrap(count)))
	}
	if !hxrt.StringEqualStringPtr(sourceError, nil) {
		hxrt.HttpUploadSinkAbort(sink, sourceError)
	} else {
		if hxrt.StringEqualStringPtr(sinkError, nil) {
			sinkError = hxrt.HttpUploadSinkFinish(sink)
		}
	}
	hx_obj_66 := map[string]any{}
	hx_obj_66["sourceError"] = sourceError
	hx_obj_66["sinkError"] = sinkError
	return hxrt.AnonymousFromMap(hx_obj_66)
}

func (self *sys__Http) handleDataRequest(post bool, api *haxe__io__Output, method *string) {
	if !hxrt.StringEqualStringPtr(method, nil) && hxrt.StringEqualStringPtr(method, hxrt.StringFromLiteral("")) {
		func() {
			var hx_callable_67 *hxrt.Function[func(*string)] = self.onError
			var hx_argument_68 *string = hxrt.StringFromLiteral("HTTP method must not be empty")
			hxrt.FunctionCall(hx_callable_67)(hx_argument_68)
		}()
		return
	}
	encoded := hxrt.StringSubstrStringPtr(self.url, hxrt.StringLengthStringPtr(hxrt.StringFromLiteral("data:")), 0, false)
	mediaType := hxrt.StringFromLiteral("text/plain")
	comma := sys__Http_firstComma(encoded)
	if comma >= 0 {
		if comma > 0 {
			mediaType = hxrt.StringSubstrStringPtr(encoded, 0, comma, true)
		}
		encoded = hxrt.StringSubstrStringPtr(encoded, int((hxrt.Int32Wrap(comma) + hxrt.Int32Wrap(1))), 0, false)
	}
	if self.file != nil {
		encoded = hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("multipart file="), func(hx_value_69 any) *string {
			if hx_value_69 == nil {
				var hx_zero_70 *string
				return hx_zero_70
			}
			return hx_value_69.(*string)
		}(hxrt.AnonymousGet(self.file, "filename"))), hxrt.StringFromLiteral(";mime=")), func(hx_value_71 any) *string {
			if hx_value_71 == nil {
				var hx_zero_72 *string
				return hx_zero_72
			}
			return hx_value_71.(*string)
		}(hxrt.AnonymousGet(self.file, "mimeType"))), hxrt.StringFromLiteral(";size=")), func(hx_value_73 any) int {
			if hx_value_73 == nil {
				var hx_zero_74 int
				return hx_zero_74
			}
			return hx_value_73.(int)
		}(hxrt.AnonymousGet(self.file, "size")))
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
		func() {
			var hx_callable_78 *hxrt.Function[func(int)] = self.onStatus
			var hx_argument_79 int = 200
			hxrt.FunctionCall(hx_callable_78)(hx_argument_79)
		}()
		api.__hx_this.prepare(payload.length)
		if payload.length > 0 {
			api.__hx_this.writeBytes(payload, 0, payload.length)
		}
		api.__hx_this.close()
	}, func(hx_caught_75 any) {
		hx_exception_76 := haxe__Exception_caught(hx_caught_75)
		hx_payload_77 := haxe__Exception_catchValue(hx_caught_75)
		_ = hx_payload_77
		var error *haxe__Exception = hx_recover_class_2(hx_exception_76)
		errorMessage = error.__hx_this.get_message()
	})
	if !hxrt.StringEqualStringPtr(errorMessage, nil) {
		func() {
			var hx_callable_80 *hxrt.Function[func(*string)] = self.onError
			var hx_argument_81 *string = errorMessage
			hxrt.FunctionCall(hx_callable_80)(hx_argument_81)
		}()
	}
}

func (self *sys__Http) recordResponseHeaders(exchange *hxrt.HttpExchange) {
	count := hxrt.HttpExchangeHeaderCount(exchange)
	_g := 0
	_g1 := count
	for _g < _g1 {
		hx_post_82 := _g
		_g = int(int32((_g + 1)))
		headerIndex := hx_post_82
		name := hxrt.StdString(hxrt.HttpExchangeHeaderName(exchange, headerIndex))
		normalized := hxrt.StringToLowerCaseStringPtr(name)
		valueCount := hxrt.HttpExchangeHeaderValueCount(exchange, headerIndex)
		values := hxrt.NewArray()
		_g_1 := 0
		_g1_1 := valueCount
		for _g_1 < _g1_1 {
			hx_post_83 := _g_1
			_g_1 = int(int32((_g_1 + 1)))
			valueIndex := hx_post_83
			values.Push(hxrt.StdString(hxrt.HttpExchangeHeaderValue(exchange, headerIndex, valueIndex)))
		}
		if values.Len() == 0 {
			continue
		}
		last := func(hx_value_85 any) *string {
			if hx_value_85 == nil {
				var hx_zero_86 *string
				return hx_zero_86
			}
			return hx_value_85.(*string)
		}(values.Get(int((hxrt.Int32Wrap(values.Len()) - hxrt.Int32Wrap(1)))))
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
		parameter := hx_structural_recover(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		encoded.Push(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(StringTools_urlEncode(func(hx_value_88 any) *string {
			if hx_value_88 == nil {
				var hx_zero_89 *string
				return hx_zero_89
			}
			return hx_value_88.(*string)
		}(hxrt.AnonymousGet(parameter, "name"))), hxrt.StringFromLiteral("=")), StringTools_urlEncode(func(hx_value_90 any) *string {
			if hx_value_90 == nil {
				var hx_zero_91 *string
				return hx_zero_91
			}
			return hx_value_90.(*string)
		}(hxrt.AnonymousGet(parameter, "value")))))
	}
	return hxrt.StringJoinAny(encoded.Values(), hxrt.StringFromLiteral("&"))
}

var sys__Http_PROXY hxrt.Anonymous = nil

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
		hx_post_92 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_92
		if func() int {
			var c any = hxrt.StringCharCodeAtAnyStringPtr(value, index)
			var hx_if_93 int
			if c == nil {
				hx_if_93 = -1
			} else {
				hx_if_93 = c.(int)
			}
			return hx_if_93
		}() == 44 {
			return index
		}
	}
	return -1
}

func sys__Http_hxrt_proxyDescriptor() *string {
	proxy := sys__Http_PROXY
	if (proxy == nil) || hxrt.StringEqualStringPtr(func(hx_value_94 any) *string {
		if hx_value_94 == nil {
			var hx_zero_95 *string
			return hx_zero_95
		}
		return hx_value_94.(*string)
	}(hxrt.AnonymousGet(proxy, "host")), nil) {
		return hxrt.StringFromLiteral("null")
	}
	var hx_if_98 *string
	if hx_structural_recover(hxrt.AnonymousGet(proxy, "auth")) == nil {
		hx_if_98 = nil
	} else {
		hx_if_98 = func(hx_value_96 any) *string {
			if hx_value_96 == nil {
				var hx_zero_97 *string
				return hx_zero_97
			}
			return hx_value_96.(*string)
		}(hxrt.AnonymousGet(hx_structural_recover(hxrt.AnonymousGet(proxy, "auth")), "user"))
	}
	user := hx_if_98
	var hx_if_101 *string
	if hx_structural_recover(hxrt.AnonymousGet(proxy, "auth")) == nil {
		hx_if_101 = nil
	} else {
		hx_if_101 = func(hx_value_99 any) *string {
			if hx_value_99 == nil {
				var hx_zero_100 *string
				return hx_zero_100
			}
			return hx_value_99.(*string)
		}(hxrt.AnonymousGet(hx_structural_recover(hxrt.AnonymousGet(proxy, "auth")), "pass"))
	}
	pass := hx_if_101
	return hxrt.StdString(hxrt.HttpProxyDescriptor(func(hx_value_102 any) *string {
		if hx_value_102 == nil {
			var hx_zero_103 *string
			return hx_zero_103
		}
		return hx_value_102.(*string)
	}(hxrt.AnonymousGet(proxy, "host")), func(hx_value_104 any) int {
		if hx_value_104 == nil {
			var hx_zero_105 int
			return hx_zero_105
		}
		return hx_value_104.(int)
	}(hxrt.AnonymousGet(proxy, "port")), user, pass))
}

func sys__Http_hxrt_statusError(status int) *string {
	var hx_if_106 *string
	if (status < 200) || (status >= 400) {
		hx_if_106 = hxrt.StringConcatAny(hxrt.StringFromLiteral("Http Error #"), status)
	} else {
		hx_if_106 = nil
	}
	return hx_if_106
}

func sys__Http_requestUrl(url *string) *string {
	request := New_sys__Http(url)
	result := hxrt.StringFromLiteral("")
	request.onData = hxrt.NewFunction(func(data *string) {
		result = data
	})
	request.onError = hxrt.NewFunction(func(message *string) {
		hxrt.Throw(haxe__Exception_thrown(message))
	})
	request.__hx_this.request(false)
	return result
}
