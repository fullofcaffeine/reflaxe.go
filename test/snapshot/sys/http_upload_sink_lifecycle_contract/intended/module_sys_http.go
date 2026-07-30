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
	hx_obj_163 := map[string]any{}
	hx_obj_163["param"] = argname
	hx_obj_163["filename"] = filename
	hx_obj_163["io"] = file
	hx_obj_163["size"] = size
	hx_obj_163["mimeType"] = mimeType
	self.file = hx_obj_163
}

func (self *sys__Http) customRequest(post bool, api *haxe__io__Output, sock *sys__net__Socket, method *string) {
	self.__hx_this.requestWith((post || (self.file != nil)), api, sock, method)
}

func (self *sys__Http) getResponseHeaderValues(key *string) *hxrt.Array {
	values := func(hx_value_164 any) *hxrt.Array {
		if hx_value_164 == nil {
			var hx_zero_165 *hxrt.Array
			return hx_zero_165
		}
		return hx_value_164.(*hxrt.Array)
	}(self.responseHeadersSameKey.__hx_this.get(key))
	if values == nil {
		normalized := hxrt.StringToLowerCaseStringPtr(key)
		if !hxrt.StringEqualStringPtr(normalized, key) {
			values = func(hx_value_166 any) *hxrt.Array {
				if hx_value_166 == nil {
					var hx_zero_167 *hxrt.Array
					return hx_zero_167
				}
				return hx_value_166.(*hxrt.Array)
			}(self.responseHeadersSameKey.__hx_this.get(normalized))
		}
	}
	if values != nil {
		return values
	}
	var this1 haxe__IMap = self.responseHeaders
	value := func(hx_value_168 any) *string {
		if hx_value_168 == nil {
			var hx_zero_169 *string
			return hx_zero_169
		}
		return hx_value_168.(*string)
	}(this1.(*haxe__ds__StringMap).__hx_this.get(key))
	if hxrt.StringEqualStringPtr(value, nil) {
		normalized_1 := hxrt.StringToLowerCaseStringPtr(key)
		if !hxrt.StringEqualStringPtr(normalized_1, key) {
			var this1_1 haxe__IMap = self.responseHeaders
			value = func(hx_value_170 any) *string {
				if hx_value_170 == nil {
					var hx_zero_171 *string
					return hx_zero_171
				}
				return hx_value_170.(*string)
			}(this1_1.(*haxe__ds__StringMap).__hx_this.get(normalized_1))
		}
	}
	var hx_if_172 *hxrt.Array
	if hxrt.StringEqualStringPtr(value, nil) {
		hx_if_172 = nil
	} else {
		hx_if_172 = hxrt.NewArray(value)
	}
	return hx_if_172
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
		parameter := func(hx_value_173 any) map[string]any {
			if hx_value_173 == nil {
				var hx_zero_174 map[string]any
				return hx_zero_174
			}
			return hx_value_173.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		hxrt.HttpRequestAddParameter(request, func(hx_obj_175 map[string]any) *string {
			hx_field_176 := hx_obj_175["name"]
			if hx_field_176 == nil {
				var hx_zero_177 *string
				return hx_zero_177
			}
			return hx_field_176.(*string)
		}(parameter), func(hx_obj_178 map[string]any) *string {
			hx_field_179 := hx_obj_178["value"]
			if hx_field_179 == nil {
				var hx_zero_180 *string
				return hx_zero_180
			}
			return hx_field_179.(*string)
		}(parameter), StringTools_urlEncode(func(hx_obj_181 map[string]any) *string {
			hx_field_182 := hx_obj_181["name"]
			if hx_field_182 == nil {
				var hx_zero_183 *string
				return hx_zero_183
			}
			return hx_field_182.(*string)
		}(parameter)), StringTools_urlEncode(func(hx_obj_184 map[string]any) *string {
			hx_field_185 := hx_obj_184["value"]
			if hx_field_185 == nil {
				var hx_zero_186 *string
				return hx_zero_186
			}
			return hx_field_185.(*string)
		}(parameter)))
	}
	_g_1 := 0
	_g1_1 := self.headers
	for _g_1 < _g1_1.Len() {
		header := func(hx_value_187 any) map[string]any {
			if hx_value_187 == nil {
				var hx_zero_188 map[string]any
				return hx_zero_188
			}
			return hx_value_187.(map[string]any)
		}(_g1_1.Get(_g_1))
		_g_1 = int(int32((_g_1 + 1)))
		hxrt.HttpRequestAddHeader(request, func(hx_obj_189 map[string]any) *string {
			hx_field_190 := hx_obj_189["name"]
			if hx_field_190 == nil {
				var hx_zero_191 *string
				return hx_zero_191
			}
			return hx_field_190.(*string)
		}(header), func(hx_obj_192 map[string]any) *string {
			hx_field_193 := hx_obj_192["value"]
			if hx_field_193 == nil {
				var hx_zero_194 *string
				return hx_zero_194
			}
			return hx_field_193.(*string)
		}(header))
	}
	upload := self.file
	if upload != nil {
		hxrt.HttpRequestSetMultipartUpload(request, func(hx_obj_195 map[string]any) *string {
			hx_field_196 := hx_obj_195["param"]
			if hx_field_196 == nil {
				var hx_zero_197 *string
				return hx_zero_197
			}
			return hx_field_196.(*string)
		}(upload), func(hx_obj_198 map[string]any) *string {
			hx_field_199 := hx_obj_198["filename"]
			if hx_field_199 == nil {
				var hx_zero_200 *string
				return hx_zero_200
			}
			return hx_field_199.(*string)
		}(upload), func(hx_obj_201 map[string]any) *string {
			hx_field_202 := hx_obj_201["mimeType"]
			if hx_field_202 == nil {
				var hx_zero_203 *string
				return hx_zero_203
			}
			return hx_field_202.(*string)
		}(upload), func(hx_obj_204 map[string]any) int {
			hx_field_205 := hx_obj_204["size"]
			if hx_field_205 == nil {
				var hx_zero_206 int
				return hx_zero_206
			}
			return hx_field_205.(int)
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
	if (proxy != nil) && !hxrt.StringEqualStringPtr(func(hx_obj_233 map[string]any) *string {
		hx_field_234 := hx_obj_233["host"]
		if hx_field_234 == nil {
			var hx_zero_235 *string
			return hx_zero_235
		}
		return hx_field_234.(*string)
	}(proxy), nil) {
		var hx_if_216 *string
		if func(hx_obj_207 map[string]any) map[string]any {
			hx_field_208 := hx_obj_207["auth"]
			if hx_field_208 == nil {
				var hx_zero_209 map[string]any
				return hx_zero_209
			}
			return hx_field_208.(map[string]any)
		}(proxy) == nil {
			hx_if_216 = nil
		} else {
			hx_if_216 = func(hx_obj_213 map[string]any) *string {
				hx_field_214 := hx_obj_213["user"]
				if hx_field_214 == nil {
					var hx_zero_215 *string
					return hx_zero_215
				}
				return hx_field_214.(*string)
			}(func(hx_obj_210 map[string]any) map[string]any {
				hx_field_211 := hx_obj_210["auth"]
				if hx_field_211 == nil {
					var hx_zero_212 map[string]any
					return hx_zero_212
				}
				return hx_field_211.(map[string]any)
			}(proxy))
		}
		user := hx_if_216
		var hx_if_226 *string
		if func(hx_obj_217 map[string]any) map[string]any {
			hx_field_218 := hx_obj_217["auth"]
			if hx_field_218 == nil {
				var hx_zero_219 map[string]any
				return hx_zero_219
			}
			return hx_field_218.(map[string]any)
		}(proxy) == nil {
			hx_if_226 = nil
		} else {
			hx_if_226 = func(hx_obj_223 map[string]any) *string {
				hx_field_224 := hx_obj_223["pass"]
				if hx_field_224 == nil {
					var hx_zero_225 *string
					return hx_zero_225
				}
				return hx_field_224.(*string)
			}(func(hx_obj_220 map[string]any) map[string]any {
				hx_field_221 := hx_obj_220["auth"]
				if hx_field_221 == nil {
					var hx_zero_222 map[string]any
					return hx_zero_222
				}
				return hx_field_221.(map[string]any)
			}(proxy))
		}
		pass := hx_if_226
		hxrt.HttpRequestSetProxy(request, func(hx_obj_227 map[string]any) *string {
			hx_field_228 := hx_obj_227["host"]
			if hx_field_228 == nil {
				var hx_zero_229 *string
				return hx_zero_229
			}
			return hx_field_228.(*string)
		}(proxy), func(hx_obj_230 map[string]any) int {
			hx_field_231 := hx_obj_230["port"]
			if hx_field_231 == nil {
				var hx_zero_232 int
				return hx_zero_232
			}
			return hx_field_231.(int)
		}(proxy), user, pass)
	}
	if sock != nil {
		hxrt.HttpRequestSetSocket(request, sock.handle)
	}
	exchange := hxrt.HttpRequestStartExchange(request)
	var hx_if_236 map[string]any
	if upload == nil {
		hx_if_236 = nil
	} else {
		hx_if_236 = self.__hx_this.pumpUpload(exchange, upload)
	}
	uploadResult := hx_if_236
	hxrt.HttpExchangeAwaitResponse(exchange)
	var hx_if_240 *string
	if uploadResult == nil {
		hx_if_240 = nil
	} else {
		hx_if_240 = func(hx_obj_237 map[string]any) *string {
			hx_field_238 := hx_obj_237["sourceError"]
			if hx_field_238 == nil {
				var hx_zero_239 *string
				return hx_zero_239
			}
			return hx_field_238.(*string)
		}(uploadResult)
	}
	sourceError := hx_if_240
	var hx_if_244 *string
	if uploadResult == nil {
		hx_if_244 = nil
	} else {
		hx_if_244 = func(hx_obj_241 map[string]any) *string {
			hx_field_242 := hx_obj_241["sinkError"]
			if hx_field_242 == nil {
				var hx_zero_243 *string
				return hx_zero_243
			}
			return hx_field_242.(*string)
		}(uploadResult)
	}
	sinkError := hx_if_244
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
				}, func(hx_caught_245 any) {
					error := hxrt.ExceptionCaught(hx_caught_245)
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
		hx_obj_247 := map[string]any{}
		hx_obj_247["sourceError"] = nil
		hx_obj_247["sinkError"] = hxrt.StringFromLiteral("HTTP upload sink is unavailable")
		return hx_obj_247
	}
	remaining := func(hx_obj_248 map[string]any) int {
		hx_field_249 := hx_obj_248["size"]
		if hx_field_249 == nil {
			var hx_zero_250 int
			return hx_zero_250
		}
		return hx_field_249.(int)
	}(upload)
	var sourceError *string = nil
	var sinkError *string = nil
	for remaining > 0 {
		var hx_if_251 int
		if remaining > 32768 {
			hx_if_251 = 32768
		} else {
			hx_if_251 = remaining
		}
		requested := hx_if_251
		chunk := haxe__io__Bytes_alloc(requested)
		count := 0
		hxrt.TryCatch(func() {
			count = func(hx_obj_254 map[string]any) *haxe__io__Input {
				hx_field_255 := hx_obj_254["io"]
				if hx_field_255 == nil {
					var hx_zero_256 *haxe__io__Input
					return hx_zero_256
				}
				return hx_field_255.(*haxe__io__Input)
			}(upload).__hx_this.readBytes(chunk, 0, requested)
		}, func(hx_caught_252 any) {
			switch hx_typed_253 := hx_caught_252.(type) {
			case *haxe__io__Eof:
				hx_tmp := hx_typed_253
				_ = hx_tmp
				sourceError = hxrt.StringFromLiteral("Transfer aborted")
			default:
				error := hxrt.ExceptionCaught(hx_caught_252)
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
	hx_obj_257 := map[string]any{}
	hx_obj_257["sourceError"] = sourceError
	hx_obj_257["sinkError"] = sinkError
	return hx_obj_257
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
		encoded = hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("multipart file="), func(hx_obj_258 map[string]any) *string {
			hx_field_259 := hx_obj_258["filename"]
			if hx_field_259 == nil {
				var hx_zero_260 *string
				return hx_zero_260
			}
			return hx_field_259.(*string)
		}(self.file)), hxrt.StringFromLiteral(";mime=")), func(hx_obj_261 map[string]any) *string {
			hx_field_262 := hx_obj_261["mimeType"]
			if hx_field_262 == nil {
				var hx_zero_263 *string
				return hx_zero_263
			}
			return hx_field_262.(*string)
		}(self.file)), hxrt.StringFromLiteral(";size=")), func(hx_obj_264 map[string]any) int {
			hx_field_265 := hx_obj_264["size"]
			if hx_field_265 == nil {
				var hx_zero_266 int
				return hx_zero_266
			}
			return hx_field_265.(int)
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
	}, func(hx_caught_267 any) {
		error := hxrt.ExceptionCaught(hx_caught_267)
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
		hx_post_269 := _g
		_g = int(int32((_g + 1)))
		headerIndex := hx_post_269
		name := hxrt.StdString(hxrt.HttpExchangeHeaderName(exchange, headerIndex))
		normalized := hxrt.StringToLowerCaseStringPtr(name)
		valueCount := hxrt.HttpExchangeHeaderValueCount(exchange, headerIndex)
		values := hxrt.NewArray()
		_g_1 := 0
		_g1_1 := valueCount
		for _g_1 < _g1_1 {
			hx_post_270 := _g_1
			_g_1 = int(int32((_g_1 + 1)))
			valueIndex := hx_post_270
			values.Push(hxrt.StdString(hxrt.HttpExchangeHeaderValue(exchange, headerIndex, valueIndex)))
		}
		if values.Len() == 0 {
			continue
		}
		last := func(hx_value_272 any) *string {
			if hx_value_272 == nil {
				var hx_zero_273 *string
				return hx_zero_273
			}
			return hx_value_272.(*string)
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
		parameter := func(hx_value_274 any) map[string]any {
			if hx_value_274 == nil {
				var hx_zero_275 map[string]any
				return hx_zero_275
			}
			return hx_value_274.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		encoded.Push(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(StringTools_urlEncode(func(hx_obj_277 map[string]any) *string {
			hx_field_278 := hx_obj_277["name"]
			if hx_field_278 == nil {
				var hx_zero_279 *string
				return hx_zero_279
			}
			return hx_field_278.(*string)
		}(parameter)), hxrt.StringFromLiteral("=")), StringTools_urlEncode(func(hx_obj_280 map[string]any) *string {
			hx_field_281 := hx_obj_280["value"]
			if hx_field_281 == nil {
				var hx_zero_282 *string
				return hx_zero_282
			}
			return hx_field_281.(*string)
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
		hx_post_283 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_283
		if func() int {
			var c any = hxrt.StringCharCodeAtAnyStringPtr(value, index)
			var hx_if_284 int
			if c == nil {
				hx_if_284 = -1
			} else {
				hx_if_284 = c.(int)
			}
			return hx_if_284
		}() == 44 {
			return index
		}
	}
	return -1
}

func sys__Http_hxrt_proxyDescriptor() *string {
	proxy := sys__Http_PROXY
	if (proxy == nil) || hxrt.StringEqualStringPtr(func(hx_obj_285 map[string]any) *string {
		hx_field_286 := hx_obj_285["host"]
		if hx_field_286 == nil {
			var hx_zero_287 *string
			return hx_zero_287
		}
		return hx_field_286.(*string)
	}(proxy), nil) {
		return hxrt.StringFromLiteral("null")
	}
	var hx_if_297 *string
	if func(hx_obj_288 map[string]any) map[string]any {
		hx_field_289 := hx_obj_288["auth"]
		if hx_field_289 == nil {
			var hx_zero_290 map[string]any
			return hx_zero_290
		}
		return hx_field_289.(map[string]any)
	}(proxy) == nil {
		hx_if_297 = nil
	} else {
		hx_if_297 = func(hx_obj_294 map[string]any) *string {
			hx_field_295 := hx_obj_294["user"]
			if hx_field_295 == nil {
				var hx_zero_296 *string
				return hx_zero_296
			}
			return hx_field_295.(*string)
		}(func(hx_obj_291 map[string]any) map[string]any {
			hx_field_292 := hx_obj_291["auth"]
			if hx_field_292 == nil {
				var hx_zero_293 map[string]any
				return hx_zero_293
			}
			return hx_field_292.(map[string]any)
		}(proxy))
	}
	user := hx_if_297
	var hx_if_307 *string
	if func(hx_obj_298 map[string]any) map[string]any {
		hx_field_299 := hx_obj_298["auth"]
		if hx_field_299 == nil {
			var hx_zero_300 map[string]any
			return hx_zero_300
		}
		return hx_field_299.(map[string]any)
	}(proxy) == nil {
		hx_if_307 = nil
	} else {
		hx_if_307 = func(hx_obj_304 map[string]any) *string {
			hx_field_305 := hx_obj_304["pass"]
			if hx_field_305 == nil {
				var hx_zero_306 *string
				return hx_zero_306
			}
			return hx_field_305.(*string)
		}(func(hx_obj_301 map[string]any) map[string]any {
			hx_field_302 := hx_obj_301["auth"]
			if hx_field_302 == nil {
				var hx_zero_303 map[string]any
				return hx_zero_303
			}
			return hx_field_302.(map[string]any)
		}(proxy))
	}
	pass := hx_if_307
	return hxrt.StdString(hxrt.HttpProxyDescriptor(func(hx_obj_308 map[string]any) *string {
		hx_field_309 := hx_obj_308["host"]
		if hx_field_309 == nil {
			var hx_zero_310 *string
			return hx_zero_310
		}
		return hx_field_309.(*string)
	}(proxy), func(hx_obj_311 map[string]any) int {
		hx_field_312 := hx_obj_311["port"]
		if hx_field_312 == nil {
			var hx_zero_313 int
			return hx_zero_313
		}
		return hx_field_312.(int)
	}(proxy), user, pass))
}

func sys__Http_hxrt_statusError(status int) *string {
	var hx_if_314 *string
	if (status < 200) || (status >= 400) {
		hx_if_314 = hxrt.StringConcatAny(hxrt.StringFromLiteral("Http Error #"), status)
	} else {
		hx_if_314 = nil
	}
	return hx_if_314
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
