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
	if (proxy != nil) && !hxrt.StringEqualStringPtr(func(hx_obj_97 map[string]any) *string {
		hx_field_98 := hx_obj_97["host"]
		if hx_field_98 == nil {
			var hx_zero_99 *string
			return hx_zero_99
		}
		return hx_field_98.(*string)
	}(proxy), nil) {
		var hx_if_80 *string
		if func(hx_obj_71 map[string]any) map[string]any {
			hx_field_72 := hx_obj_71["auth"]
			if hx_field_72 == nil {
				var hx_zero_73 map[string]any
				return hx_zero_73
			}
			return hx_field_72.(map[string]any)
		}(proxy) == nil {
			hx_if_80 = nil
		} else {
			hx_if_80 = func(hx_obj_77 map[string]any) *string {
				hx_field_78 := hx_obj_77["user"]
				if hx_field_78 == nil {
					var hx_zero_79 *string
					return hx_zero_79
				}
				return hx_field_78.(*string)
			}(func(hx_obj_74 map[string]any) map[string]any {
				hx_field_75 := hx_obj_74["auth"]
				if hx_field_75 == nil {
					var hx_zero_76 map[string]any
					return hx_zero_76
				}
				return hx_field_75.(map[string]any)
			}(proxy))
		}
		user := hx_if_80
		var hx_if_90 *string
		if func(hx_obj_81 map[string]any) map[string]any {
			hx_field_82 := hx_obj_81["auth"]
			if hx_field_82 == nil {
				var hx_zero_83 map[string]any
				return hx_zero_83
			}
			return hx_field_82.(map[string]any)
		}(proxy) == nil {
			hx_if_90 = nil
		} else {
			hx_if_90 = func(hx_obj_87 map[string]any) *string {
				hx_field_88 := hx_obj_87["pass"]
				if hx_field_88 == nil {
					var hx_zero_89 *string
					return hx_zero_89
				}
				return hx_field_88.(*string)
			}(func(hx_obj_84 map[string]any) map[string]any {
				hx_field_85 := hx_obj_84["auth"]
				if hx_field_85 == nil {
					var hx_zero_86 map[string]any
					return hx_zero_86
				}
				return hx_field_85.(map[string]any)
			}(proxy))
		}
		pass := hx_if_90
		hxrt.HttpRequestSetProxy(request, func(hx_obj_91 map[string]any) *string {
			hx_field_92 := hx_obj_91["host"]
			if hx_field_92 == nil {
				var hx_zero_93 *string
				return hx_zero_93
			}
			return hx_field_92.(*string)
		}(proxy), func(hx_obj_94 map[string]any) int {
			hx_field_95 := hx_obj_94["port"]
			if hx_field_95 == nil {
				var hx_zero_96 int
				return hx_zero_96
			}
			return hx_field_95.(int)
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
			encoded = hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("multipart file="), func(hx_obj_100 map[string]any) *string {
				hx_field_101 := hx_obj_100["filename"]
				if hx_field_101 == nil {
					var hx_zero_102 *string
					return hx_zero_102
				}
				return hx_field_101.(*string)
			}(self.file)), hxrt.StringFromLiteral(";mime=")), func(hx_obj_103 map[string]any) *string {
				hx_field_104 := hx_obj_103["mimeType"]
				if hx_field_104 == nil {
					var hx_zero_105 *string
					return hx_zero_105
				}
				return hx_field_104.(*string)
			}(self.file)), hxrt.StringFromLiteral(";size=")), func(hx_obj_106 map[string]any) int {
				hx_field_107 := hx_obj_106["size"]
				if hx_field_107 == nil {
					var hx_zero_108 int
					return hx_zero_108
				}
				return hx_field_107.(int)
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
		hx_post_109 := _g
		_g = int(int32((_g + 1)))
		headerIndex := hx_post_109
		name := hxrt.StdString(hxrt.HttpResponseHeaderName(response, headerIndex))
		normalized := hxrt.StringToLowerCaseStringPtr(name)
		valueCount := hxrt.HttpResponseHeaderValueCount(response, headerIndex)
		values := hxrt.NewArray()
		_g_1 := 0
		_g1_1 := valueCount
		for _g_1 < _g1_1 {
			hx_post_110 := _g_1
			_g_1 = int(int32((_g_1 + 1)))
			valueIndex := hx_post_110
			values.Push(hxrt.StdString(hxrt.HttpResponseHeaderValue(response, headerIndex, valueIndex)))
		}
		if values.Len() == 0 {
			continue
		}
		last := func(hx_value_112 any) *string {
			if hx_value_112 == nil {
				var hx_zero_113 *string
				return hx_zero_113
			}
			return hx_value_112.(*string)
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
		header := func(hx_value_114 any) map[string]any {
			if hx_value_114 == nil {
				var hx_zero_115 map[string]any
				return hx_zero_115
			}
			return hx_value_114.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(hxrt.StringToLowerCaseStringPtr(func(hx_obj_116 map[string]any) *string {
			hx_field_117 := hx_obj_116["name"]
			if hx_field_117 == nil {
				var hx_zero_118 *string
				return hx_zero_118
			}
			return hx_field_117.(*string)
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
		parameter := func(hx_value_119 any) map[string]any {
			if hx_value_119 == nil {
				var hx_zero_120 map[string]any
				return hx_zero_120
			}
			return hx_value_119.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("--hxrt-go-boundary\r\n"))
		x := hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Content-Disposition: form-data; name=\""), func(hx_obj_121 map[string]any) *string {
			hx_field_122 := hx_obj_121["name"]
			if hx_field_122 == nil {
				var hx_zero_123 *string
				return hx_zero_123
			}
			return hx_field_122.(*string)
		}(parameter)), hxrt.StringFromLiteral("\"\r\n\r\n"))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		x_1 := func(hx_obj_124 map[string]any) *string {
			hx_field_125 := hx_obj_124["value"]
			if hx_field_125 == nil {
				var hx_zero_126 *string
				return hx_zero_126
			}
			return hx_field_125.(*string)
		}(parameter)
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x_1))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("\r\n"))
	}
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("--hxrt-go-boundary\r\n"))
	x_2 := hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Content-Disposition: form-data; name=\""), func(hx_obj_127 map[string]any) *string {
		hx_field_128 := hx_obj_127["param"]
		if hx_field_128 == nil {
			var hx_zero_129 *string
			return hx_zero_129
		}
		return hx_field_128.(*string)
	}(upload)), hxrt.StringFromLiteral("\"; filename=\"")), func(hx_obj_130 map[string]any) *string {
		hx_field_131 := hx_obj_130["filename"]
		if hx_field_131 == nil {
			var hx_zero_132 *string
			return hx_zero_132
		}
		return hx_field_131.(*string)
	}(upload)), hxrt.StringFromLiteral("\"\r\n"))
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x_2))
	x_3 := hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Content-Type: "), func(hx_obj_133 map[string]any) *string {
		hx_field_134 := hx_obj_133["mimeType"]
		if hx_field_134 == nil {
			var hx_zero_135 *string
			return hx_zero_135
		}
		return hx_field_134.(*string)
	}(upload)), hxrt.StringFromLiteral("\r\n\r\n"))
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x_3))
	x_4 := hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("[uploaded-bytes="), func(hx_obj_136 map[string]any) int {
		hx_field_137 := hx_obj_136["size"]
		if hx_field_137 == nil {
			var hx_zero_138 int
			return hx_zero_138
		}
		return hx_field_137.(int)
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
		parameter := func(hx_value_139 any) map[string]any {
			if hx_value_139 == nil {
				var hx_zero_140 map[string]any
				return hx_zero_140
			}
			return hx_value_139.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		byName.__hx_this.set(func(hx_obj_141 map[string]any) *string {
			hx_field_142 := hx_obj_141["name"]
			if hx_field_142 == nil {
				var hx_zero_143 *string
				return hx_zero_143
			}
			return hx_field_142.(*string)
		}(parameter), func(hx_obj_144 map[string]any) *string {
			hx_field_145 := hx_obj_144["value"]
			if hx_field_145 == nil {
				var hx_zero_146 *string
				return hx_zero_146
			}
			return hx_field_145.(*string)
		}(parameter))
	}
	_g_1 := hxrt.NewArray()
	name := func(hx_value_147 any) map[string]any {
		if hx_value_147 == nil {
			var hx_zero_148 map[string]any
			return hx_zero_148
		}
		return hx_value_147.(map[string]any)
	}(byName.__hx_this.keys())
	for func(hx_obj_149 map[string]any) func() bool {
		hx_field_150 := hx_obj_149["hasNext"]
		if hx_field_150 == nil {
			var hx_zero_151 func() bool
			return hx_zero_151
		}
		return hx_field_150.(func() bool)
	}(name)() {
		name_1 := func(hx_obj_152 map[string]any) func() *string {
			hx_field_153 := hx_obj_152["next"]
			if hx_field_153 == nil {
				var hx_zero_154 func() *string
				return hx_zero_154
			}
			return hx_field_153.(func() *string)
		}(name)()
		_g_1.Push(name_1)
	}
	names := _g_1
	encoded := hxrt.NewArray()
	emitted := New_haxe__ds__StringMap()
	_g_2 := 0
	for _g_2 < names.Len() {
		hx_tmp := func(hx_value_156 any) *string {
			if hx_value_156 == nil {
				var hx_zero_157 *string
				return hx_zero_157
			}
			return hx_value_156.(*string)
		}(names.Get(_g_2))
		_ = hx_tmp
		_g_2 = int(int32((_g_2 + 1)))
		next := -1
		_g_3 := 0
		_g1_1 := names.Len()
		for _g_3 < _g1_1 {
			hx_post_158 := _g_3
			_g_3 = int(int32((_g_3 + 1)))
			index := hx_post_158
			if !func(hx_value_161 any) bool {
				if hx_value_161 == nil {
					var hx_zero_162 bool
					return hx_zero_162
				}
				return hx_value_161.(bool)
			}(emitted.__hx_this.exists(func(hx_value_159 any) *string {
				if hx_value_159 == nil {
					var hx_zero_160 *string
					return hx_zero_160
				}
				return hx_value_159.(*string)
			}(names.Get(index)))) && ((next < 0) || (Reflect_compare(func(hx_value_163 any) *string {
				if hx_value_163 == nil {
					var hx_zero_164 *string
					return hx_zero_164
				}
				return hx_value_163.(*string)
			}(names.Get(index)), func(hx_value_165 any) *string {
				if hx_value_165 == nil {
					var hx_zero_166 *string
					return hx_zero_166
				}
				return hx_value_165.(*string)
			}(names.Get(next))) < 0)) {
				next = index
			}
		}
		if next >= 0 {
			name_2 := func(hx_value_167 any) *string {
				if hx_value_167 == nil {
					var hx_zero_168 *string
					return hx_zero_168
				}
				return hx_value_167.(*string)
			}(names.Get(next))
			emitted.__hx_this.set(name_2, true)
			encoded.Push(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(StringTools_urlEncode(name_2), hxrt.StringFromLiteral("=")), StringTools_urlEncode(func(hx_value_170 any) *string {
				if hx_value_170 == nil {
					var hx_zero_171 *string
					return hx_zero_171
				}
				return hx_value_170.(*string)
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
		hx_post_172 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_172
		if func() int {
			var c any = hxrt.StringCharCodeAtAnyStringPtr(value, index)
			var hx_if_173 int
			if c == nil {
				hx_if_173 = -1
			} else {
				hx_if_173 = c.(int)
			}
			return hx_if_173
		}() == 44 {
			return index
		}
	}
	return -1
}

func sys__Http_hxrt_proxyDescriptor() *string {
	proxy := sys__Http_PROXY
	if (proxy == nil) || hxrt.StringEqualStringPtr(func(hx_obj_174 map[string]any) *string {
		hx_field_175 := hx_obj_174["host"]
		if hx_field_175 == nil {
			var hx_zero_176 *string
			return hx_zero_176
		}
		return hx_field_175.(*string)
	}(proxy), nil) {
		return hxrt.StringFromLiteral("null")
	}
	var hx_if_186 *string
	if func(hx_obj_177 map[string]any) map[string]any {
		hx_field_178 := hx_obj_177["auth"]
		if hx_field_178 == nil {
			var hx_zero_179 map[string]any
			return hx_zero_179
		}
		return hx_field_178.(map[string]any)
	}(proxy) == nil {
		hx_if_186 = nil
	} else {
		hx_if_186 = func(hx_obj_183 map[string]any) *string {
			hx_field_184 := hx_obj_183["user"]
			if hx_field_184 == nil {
				var hx_zero_185 *string
				return hx_zero_185
			}
			return hx_field_184.(*string)
		}(func(hx_obj_180 map[string]any) map[string]any {
			hx_field_181 := hx_obj_180["auth"]
			if hx_field_181 == nil {
				var hx_zero_182 map[string]any
				return hx_zero_182
			}
			return hx_field_181.(map[string]any)
		}(proxy))
	}
	user := hx_if_186
	var hx_if_196 *string
	if func(hx_obj_187 map[string]any) map[string]any {
		hx_field_188 := hx_obj_187["auth"]
		if hx_field_188 == nil {
			var hx_zero_189 map[string]any
			return hx_zero_189
		}
		return hx_field_188.(map[string]any)
	}(proxy) == nil {
		hx_if_196 = nil
	} else {
		hx_if_196 = func(hx_obj_193 map[string]any) *string {
			hx_field_194 := hx_obj_193["pass"]
			if hx_field_194 == nil {
				var hx_zero_195 *string
				return hx_zero_195
			}
			return hx_field_194.(*string)
		}(func(hx_obj_190 map[string]any) map[string]any {
			hx_field_191 := hx_obj_190["auth"]
			if hx_field_191 == nil {
				var hx_zero_192 map[string]any
				return hx_zero_192
			}
			return hx_field_191.(map[string]any)
		}(proxy))
	}
	pass := hx_if_196
	return hxrt.StdString(hxrt.HttpProxyDescriptor(func(hx_obj_197 map[string]any) *string {
		hx_field_198 := hx_obj_197["host"]
		if hx_field_198 == nil {
			var hx_zero_199 *string
			return hx_zero_199
		}
		return hx_field_198.(*string)
	}(proxy), func(hx_obj_200 map[string]any) int {
		hx_field_201 := hx_obj_200["port"]
		if hx_field_201 == nil {
			var hx_zero_202 int
			return hx_zero_202
		}
		return hx_field_201.(int)
	}(proxy), user, pass))
}

func sys__Http_normalizedMethod(method *string) *string {
	if hxrt.StringEqualStringPtr(method, nil) {
		return nil
	}
	normalized := hxrt.StringToUpperCaseStringPtr(method)
	var hx_if_203 *string
	if hxrt.StringEqualStringPtr(normalized, hxrt.StringFromLiteral("")) || hxrt.StringEqualStringPtr(normalized, hxrt.StringFromLiteral("NULL")) {
		hx_if_203 = nil
	} else {
		hx_if_203 = normalized
	}
	return hx_if_203
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
