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
	recordResponseHeaders(response *hxrt.HttpResponse)
	resetResponseHeaders()
	hasHeader(name *string) bool
	buildMultipartBody(upload map[string]any) *string
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
	self.__hx_this.requestWith(usePost, nil, nil, nil)
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
		}(parameter))
	}
	_g_1 := 0
	_g1_1 := self.headers
	for _g_1 < _g1_1.Len() {
		header := func(hx_value_62 any) map[string]any {
			if hx_value_62 == nil {
				var hx_zero_63 map[string]any
				return hx_zero_63
			}
			return hx_value_62.(map[string]any)
		}(_g1_1.Get(_g_1))
		_g_1 = int(int32((_g_1 + 1)))
		hxrt.HttpRequestAddHeader(request, func(hx_obj_64 map[string]any) *string {
			hx_field_65 := hx_obj_64["name"]
			if hx_field_65 == nil {
				var hx_zero_66 *string
				return hx_zero_66
			}
			return hx_field_65.(*string)
		}(header), func(hx_obj_67 map[string]any) *string {
			hx_field_68 := hx_obj_67["value"]
			if hx_field_68 == nil {
				var hx_zero_69 *string
				return hx_zero_69
			}
			return hx_field_68.(*string)
		}(header))
	}
	if self.file != nil {
		hxrt.HttpRequestSetBodyString(request, self.__hx_this.buildMultipartBody(self.file))
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
	if (proxy != nil) && !hxrt.StringEqualStringPtr(func(hx_obj_96 map[string]any) *string {
		hx_field_97 := hx_obj_96["host"]
		if hx_field_97 == nil {
			var hx_zero_98 *string
			return hx_zero_98
		}
		return hx_field_97.(*string)
	}(proxy), nil) {
		var hx_if_79 *string
		if func(hx_obj_70 map[string]any) map[string]any {
			hx_field_71 := hx_obj_70["auth"]
			if hx_field_71 == nil {
				var hx_zero_72 map[string]any
				return hx_zero_72
			}
			return hx_field_71.(map[string]any)
		}(proxy) == nil {
			hx_if_79 = nil
		} else {
			hx_if_79 = func(hx_obj_76 map[string]any) *string {
				hx_field_77 := hx_obj_76["user"]
				if hx_field_77 == nil {
					var hx_zero_78 *string
					return hx_zero_78
				}
				return hx_field_77.(*string)
			}(func(hx_obj_73 map[string]any) map[string]any {
				hx_field_74 := hx_obj_73["auth"]
				if hx_field_74 == nil {
					var hx_zero_75 map[string]any
					return hx_zero_75
				}
				return hx_field_74.(map[string]any)
			}(proxy))
		}
		user := hx_if_79
		var hx_if_89 *string
		if func(hx_obj_80 map[string]any) map[string]any {
			hx_field_81 := hx_obj_80["auth"]
			if hx_field_81 == nil {
				var hx_zero_82 map[string]any
				return hx_zero_82
			}
			return hx_field_81.(map[string]any)
		}(proxy) == nil {
			hx_if_89 = nil
		} else {
			hx_if_89 = func(hx_obj_86 map[string]any) *string {
				hx_field_87 := hx_obj_86["pass"]
				if hx_field_87 == nil {
					var hx_zero_88 *string
					return hx_zero_88
				}
				return hx_field_87.(*string)
			}(func(hx_obj_83 map[string]any) map[string]any {
				hx_field_84 := hx_obj_83["auth"]
				if hx_field_84 == nil {
					var hx_zero_85 map[string]any
					return hx_zero_85
				}
				return hx_field_84.(map[string]any)
			}(proxy))
		}
		pass := hx_if_89
		hxrt.HttpRequestSetProxy(request, func(hx_obj_90 map[string]any) *string {
			hx_field_91 := hx_obj_90["host"]
			if hx_field_91 == nil {
				var hx_zero_92 *string
				return hx_zero_92
			}
			return hx_field_91.(*string)
		}(proxy), func(hx_obj_93 map[string]any) int {
			hx_field_94 := hx_obj_93["port"]
			if hx_field_94 == nil {
				var hx_zero_95 int
				return hx_zero_95
			}
			return hx_field_94.(int)
		}(proxy), user, pass)
	}
	if sock != nil {
		hxrt.HttpRequestSetSocket(request, sock.handle)
	}
	response := hxrt.HttpRequestExecute(request)
	status := hxrt.HttpResponseStatus(response)
	nativeError := hxrt.HttpResponseError(response)
	if (status == 0) && !hxrt.StringEqualStringPtr(nativeError, nil) {
		self.onError(nativeError)
		return
	}
	self.__hx_this.recordResponseHeaders(response)
	self.onStatus(status)
	if !hxrt.StringEqualStringPtr(nativeError, nil) {
		self.onError(nativeError)
		return
	}
	payload := haxe__io__Bytes___hx_fromNativeView(hxrt.HttpResponseBody(response))
	self.responseBytes = payload
	self.responseAsString = payload.__hx_this.toString()
	sys__Http_capture(api, payload)
	if status >= 400 {
		self.onError(hxrt.StringConcatAny(hxrt.StringFromLiteral("Http Error #"), status))
		return
	}
	self.onData(self.responseAsString)
	self.onBytes(payload)
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
			encoded = hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("multipart file="), func(hx_obj_99 map[string]any) *string {
				hx_field_100 := hx_obj_99["filename"]
				if hx_field_100 == nil {
					var hx_zero_101 *string
					return hx_zero_101
				}
				return hx_field_100.(*string)
			}(self.file)), hxrt.StringFromLiteral(";mime=")), func(hx_obj_102 map[string]any) *string {
				hx_field_103 := hx_obj_102["mimeType"]
				if hx_field_103 == nil {
					var hx_zero_104 *string
					return hx_zero_104
				}
				return hx_field_103.(*string)
			}(self.file)), hxrt.StringFromLiteral(";size=")), func(hx_obj_105 map[string]any) int {
				hx_field_106 := hx_obj_105["size"]
				if hx_field_106 == nil {
					var hx_zero_107 int
					return hx_zero_107
				}
				return hx_field_106.(int)
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
	self.responseBytes = payload
	self.responseAsString = payloadText
	var this1 haxe__IMap = self.responseHeaders
	this1.(*haxe__ds__StringMap).__hx_this.set(hxrt.StringFromLiteral("content-type"), mediaType)
	var this1_1 haxe__IMap = self.responseHeaders
	this1_1.(*haxe__ds__StringMap).__hx_this.set(hxrt.StringFromLiteral("Content-Type"), mediaType)
	sys__Http_capture(api, payload)
	self.onStatus(200)
	self.onData(payloadText)
	self.onBytes(payload)
}

func (self *sys__Http) recordResponseHeaders(response *hxrt.HttpResponse) {
	count := hxrt.HttpResponseHeaderCount(response)
	_g := 0
	_g1 := count
	for _g < _g1 {
		hx_post_108 := _g
		_g = int(int32((_g + 1)))
		headerIndex := hx_post_108
		name := hxrt.StdString(hxrt.HttpResponseHeaderName(response, headerIndex))
		normalized := hxrt.StringToLowerCaseStringPtr(name)
		valueCount := hxrt.HttpResponseHeaderValueCount(response, headerIndex)
		values := hxrt.NewArray()
		_g_1 := 0
		_g1_1 := valueCount
		for _g_1 < _g1_1 {
			hx_post_109 := _g_1
			_g_1 = int(int32((_g_1 + 1)))
			valueIndex := hx_post_109
			values.Push(hxrt.StdString(hxrt.HttpResponseHeaderValue(response, headerIndex, valueIndex)))
		}
		if values.Len() == 0 {
			continue
		}
		last := func(hx_value_111 any) *string {
			if hx_value_111 == nil {
				var hx_zero_112 *string
				return hx_zero_112
			}
			return hx_value_111.(*string)
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
		header := func(hx_value_113 any) map[string]any {
			if hx_value_113 == nil {
				var hx_zero_114 map[string]any
				return hx_zero_114
			}
			return hx_value_113.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(hxrt.StringToLowerCaseStringPtr(func(hx_obj_115 map[string]any) *string {
			hx_field_116 := hx_obj_115["name"]
			if hx_field_116 == nil {
				var hx_zero_117 *string
				return hx_zero_117
			}
			return hx_field_116.(*string)
		}(header)), hxrt.StringToLowerCaseStringPtr(name)) {
			return true
		}
	}
	return false
}

func (self *sys__Http) buildMultipartBody(upload map[string]any) *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	_g := 0
	_g1 := self.params
	for _g < _g1.Len() {
		parameter := func(hx_value_118 any) map[string]any {
			if hx_value_118 == nil {
				var hx_zero_119 map[string]any
				return hx_zero_119
			}
			return hx_value_118.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("--hxrt-go-boundary\r\n"))
		x := hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Content-Disposition: form-data; name=\""), func(hx_obj_120 map[string]any) *string {
			hx_field_121 := hx_obj_120["name"]
			if hx_field_121 == nil {
				var hx_zero_122 *string
				return hx_zero_122
			}
			return hx_field_121.(*string)
		}(parameter)), hxrt.StringFromLiteral("\"\r\n\r\n"))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		x_1 := func(hx_obj_123 map[string]any) *string {
			hx_field_124 := hx_obj_123["value"]
			if hx_field_124 == nil {
				var hx_zero_125 *string
				return hx_zero_125
			}
			return hx_field_124.(*string)
		}(parameter)
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x_1))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("\r\n"))
	}
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("--hxrt-go-boundary\r\n"))
	x_2 := hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Content-Disposition: form-data; name=\""), func(hx_obj_126 map[string]any) *string {
		hx_field_127 := hx_obj_126["param"]
		if hx_field_127 == nil {
			var hx_zero_128 *string
			return hx_zero_128
		}
		return hx_field_127.(*string)
	}(upload)), hxrt.StringFromLiteral("\"; filename=\"")), func(hx_obj_129 map[string]any) *string {
		hx_field_130 := hx_obj_129["filename"]
		if hx_field_130 == nil {
			var hx_zero_131 *string
			return hx_zero_131
		}
		return hx_field_130.(*string)
	}(upload)), hxrt.StringFromLiteral("\"\r\n"))
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x_2))
	x_3 := hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Content-Type: "), func(hx_obj_132 map[string]any) *string {
		hx_field_133 := hx_obj_132["mimeType"]
		if hx_field_133 == nil {
			var hx_zero_134 *string
			return hx_zero_134
		}
		return hx_field_133.(*string)
	}(upload)), hxrt.StringFromLiteral("\r\n\r\n"))
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x_3))
	x_4 := hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("[uploaded-bytes="), func(hx_obj_135 map[string]any) int {
		hx_field_136 := hx_obj_135["size"]
		if hx_field_136 == nil {
			var hx_zero_137 int
			return hx_zero_137
		}
		return hx_field_136.(int)
	}(upload)), hxrt.StringFromLiteral("]\r\n"))
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x_4))
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("--hxrt-go-boundary--\r\n"))
	return out_b
}

func (self *sys__Http) encodedParameters() *string {
	byName := New_haxe__ds__StringMap()
	_g := 0
	_g1 := self.params
	for _g < _g1.Len() {
		parameter := func(hx_value_138 any) map[string]any {
			if hx_value_138 == nil {
				var hx_zero_139 map[string]any
				return hx_zero_139
			}
			return hx_value_138.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		byName.__hx_this.set(func(hx_obj_140 map[string]any) *string {
			hx_field_141 := hx_obj_140["name"]
			if hx_field_141 == nil {
				var hx_zero_142 *string
				return hx_zero_142
			}
			return hx_field_141.(*string)
		}(parameter), func(hx_obj_143 map[string]any) *string {
			hx_field_144 := hx_obj_143["value"]
			if hx_field_144 == nil {
				var hx_zero_145 *string
				return hx_zero_145
			}
			return hx_field_144.(*string)
		}(parameter))
	}
	_g_1 := hxrt.NewArray()
	name := func(hx_value_146 any) map[string]any {
		if hx_value_146 == nil {
			var hx_zero_147 map[string]any
			return hx_zero_147
		}
		return hx_value_146.(map[string]any)
	}(byName.__hx_this.keys())
	for func(hx_obj_148 map[string]any) func() bool {
		hx_field_149 := hx_obj_148["hasNext"]
		if hx_field_149 == nil {
			var hx_zero_150 func() bool
			return hx_zero_150
		}
		return hx_field_149.(func() bool)
	}(name)() {
		name_1 := func(hx_obj_151 map[string]any) func() *string {
			hx_field_152 := hx_obj_151["next"]
			if hx_field_152 == nil {
				var hx_zero_153 func() *string
				return hx_zero_153
			}
			return hx_field_152.(func() *string)
		}(name)()
		_g_1.Push(name_1)
	}
	names := _g_1
	encoded := hxrt.NewArray()
	emitted := New_haxe__ds__StringMap()
	_g_2 := 0
	for _g_2 < names.Len() {
		hx_tmp := func(hx_value_155 any) *string {
			if hx_value_155 == nil {
				var hx_zero_156 *string
				return hx_zero_156
			}
			return hx_value_155.(*string)
		}(names.Get(_g_2))
		_ = hx_tmp
		_g_2 = int(int32((_g_2 + 1)))
		next := -1
		_g_3 := 0
		_g1_1 := names.Len()
		for _g_3 < _g1_1 {
			hx_post_157 := _g_3
			_g_3 = int(int32((_g_3 + 1)))
			index := hx_post_157
			if !func(hx_value_160 any) bool {
				if hx_value_160 == nil {
					var hx_zero_161 bool
					return hx_zero_161
				}
				return hx_value_160.(bool)
			}(emitted.__hx_this.exists(func(hx_value_158 any) *string {
				if hx_value_158 == nil {
					var hx_zero_159 *string
					return hx_zero_159
				}
				return hx_value_158.(*string)
			}(names.Get(index)))) && ((next < 0) || (Reflect_compare(func(hx_value_162 any) *string {
				if hx_value_162 == nil {
					var hx_zero_163 *string
					return hx_zero_163
				}
				return hx_value_162.(*string)
			}(names.Get(index)), func(hx_value_164 any) *string {
				if hx_value_164 == nil {
					var hx_zero_165 *string
					return hx_zero_165
				}
				return hx_value_164.(*string)
			}(names.Get(next))) < 0)) {
				next = index
			}
		}
		if next >= 0 {
			name_2 := func(hx_value_166 any) *string {
				if hx_value_166 == nil {
					var hx_zero_167 *string
					return hx_zero_167
				}
				return hx_value_166.(*string)
			}(names.Get(next))
			emitted.__hx_this.set(name_2, true)
			encoded.Push(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(StringTools_urlEncode(name_2), hxrt.StringFromLiteral("=")), StringTools_urlEncode(func(hx_value_169 any) *string {
				if hx_value_169 == nil {
					var hx_zero_170 *string
					return hx_zero_170
				}
				return hx_value_169.(*string)
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
		hx_post_171 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_171
		if func() int {
			var c any = hxrt.StringCharCodeAtAnyStringPtr(value, index)
			var hx_if_172 int
			if c == nil {
				hx_if_172 = -1
			} else {
				hx_if_172 = c.(int)
			}
			return hx_if_172
		}() == 44 {
			return index
		}
	}
	return -1
}

func sys__Http_hxrt_proxyDescriptor() *string {
	proxy := sys__Http_PROXY
	if (proxy == nil) || hxrt.StringEqualStringPtr(func(hx_obj_173 map[string]any) *string {
		hx_field_174 := hx_obj_173["host"]
		if hx_field_174 == nil {
			var hx_zero_175 *string
			return hx_zero_175
		}
		return hx_field_174.(*string)
	}(proxy), nil) {
		return hxrt.StringFromLiteral("null")
	}
	var hx_if_185 *string
	if func(hx_obj_176 map[string]any) map[string]any {
		hx_field_177 := hx_obj_176["auth"]
		if hx_field_177 == nil {
			var hx_zero_178 map[string]any
			return hx_zero_178
		}
		return hx_field_177.(map[string]any)
	}(proxy) == nil {
		hx_if_185 = nil
	} else {
		hx_if_185 = func(hx_obj_182 map[string]any) *string {
			hx_field_183 := hx_obj_182["user"]
			if hx_field_183 == nil {
				var hx_zero_184 *string
				return hx_zero_184
			}
			return hx_field_183.(*string)
		}(func(hx_obj_179 map[string]any) map[string]any {
			hx_field_180 := hx_obj_179["auth"]
			if hx_field_180 == nil {
				var hx_zero_181 map[string]any
				return hx_zero_181
			}
			return hx_field_180.(map[string]any)
		}(proxy))
	}
	user := hx_if_185
	var hx_if_195 *string
	if func(hx_obj_186 map[string]any) map[string]any {
		hx_field_187 := hx_obj_186["auth"]
		if hx_field_187 == nil {
			var hx_zero_188 map[string]any
			return hx_zero_188
		}
		return hx_field_187.(map[string]any)
	}(proxy) == nil {
		hx_if_195 = nil
	} else {
		hx_if_195 = func(hx_obj_192 map[string]any) *string {
			hx_field_193 := hx_obj_192["pass"]
			if hx_field_193 == nil {
				var hx_zero_194 *string
				return hx_zero_194
			}
			return hx_field_193.(*string)
		}(func(hx_obj_189 map[string]any) map[string]any {
			hx_field_190 := hx_obj_189["auth"]
			if hx_field_190 == nil {
				var hx_zero_191 map[string]any
				return hx_zero_191
			}
			return hx_field_190.(map[string]any)
		}(proxy))
	}
	pass := hx_if_195
	return hxrt.StdString(hxrt.HttpProxyDescriptor(func(hx_obj_196 map[string]any) *string {
		hx_field_197 := hx_obj_196["host"]
		if hx_field_197 == nil {
			var hx_zero_198 *string
			return hx_zero_198
		}
		return hx_field_197.(*string)
	}(proxy), func(hx_obj_199 map[string]any) int {
		hx_field_200 := hx_obj_199["port"]
		if hx_field_200 == nil {
			var hx_zero_201 int
			return hx_zero_201
		}
		return hx_field_200.(int)
	}(proxy), user, pass))
}

func sys__Http_normalizedMethod(method *string) *string {
	if hxrt.StringEqualStringPtr(method, nil) {
		return nil
	}
	normalized := hxrt.StringToUpperCaseStringPtr(method)
	var hx_if_202 *string
	if hxrt.StringEqualStringPtr(normalized, hxrt.StringFromLiteral("")) || hxrt.StringEqualStringPtr(normalized, hxrt.StringFromLiteral("NULL")) {
		hx_if_202 = nil
	} else {
		hx_if_202 = normalized
	}
	return hx_if_202
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
