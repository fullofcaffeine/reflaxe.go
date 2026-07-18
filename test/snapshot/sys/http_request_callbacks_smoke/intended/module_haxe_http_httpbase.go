package main

import "snapshot/hxrt"

type I_haxe__http__HttpBase interface {
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
}

type haxe__http__HttpBase struct {
	__hx_this        I_haxe__http__HttpBase
	url              *string
	responseData     *string
	responseBytes    *haxe__io__Bytes
	responseAsString *string
	postData         *string
	postBytes        *haxe__io__Bytes
	headers          *hxrt.Array
	params           *hxrt.Array
	emptyOnData      func(*string)
	onData           func(*string)
	onBytes          func(*haxe__io__Bytes)
	onError          func(*string)
	onStatus         func(int)
}

func New_haxe__http__HttpBase(url *string) *haxe__http__HttpBase {
	self := &haxe__http__HttpBase{}
	self.__hx_this = self
	self.onData = func(data *string) {
	}
	self.onBytes = func(data *haxe__io__Bytes) {
	}
	self.onError = func(msg *string) {
	}
	self.onStatus = func(status int) {
	}
	self.url = url
	self.headers = hxrt.NewArray()
	self.params = hxrt.NewArray()
	self.emptyOnData = self.onData
	return self
}

func (self *haxe__http__HttpBase) setHeader(name *string, value *string) {
	_g := 0
	_g1 := self.headers.Len()
	for _g < _g1 {
		hx_post_228 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_228
		if hxrt.StringEqualStringPtr(func(hx_obj_234 map[string]any) *string {
			hx_field_235 := hx_obj_234["name"]
			if hx_field_235 == nil {
				var hx_zero_236 *string
				return hx_zero_236
			}
			return hx_field_235.(*string)
		}(func(hx_value_232 any) map[string]any {
			if hx_value_232 == nil {
				var hx_zero_233 map[string]any
				return hx_zero_233
			}
			return hx_value_232.(map[string]any)
		}(self.headers.Get(i))), name) {
			hx_array_target_230 := self.headers
			hx_array_index_231 := i
			hx_obj_229 := map[string]any{}
			hx_obj_229["name"] = name
			hx_obj_229["value"] = value
			hx_array_target_230.Set(hx_array_index_231, hx_obj_229)
			return
		}
	}
	hx_arr_237 := self.headers
	hx_obj_238 := map[string]any{}
	hx_obj_238["name"] = name
	hx_obj_238["value"] = value
	hx_arr_237.Push(hx_obj_238)
}

func (self *haxe__http__HttpBase) addHeader(header *string, value *string) {
	hx_arr_239 := self.headers
	hx_obj_240 := map[string]any{}
	hx_obj_240["name"] = header
	hx_obj_240["value"] = value
	hx_arr_239.Push(hx_obj_240)
}

func (self *haxe__http__HttpBase) setParameter(name *string, value *string) {
	_g := 0
	_g1 := self.params.Len()
	for _g < _g1 {
		hx_post_241 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_241
		if hxrt.StringEqualStringPtr(func(hx_obj_247 map[string]any) *string {
			hx_field_248 := hx_obj_247["name"]
			if hx_field_248 == nil {
				var hx_zero_249 *string
				return hx_zero_249
			}
			return hx_field_248.(*string)
		}(func(hx_value_245 any) map[string]any {
			if hx_value_245 == nil {
				var hx_zero_246 map[string]any
				return hx_zero_246
			}
			return hx_value_245.(map[string]any)
		}(self.params.Get(i))), name) {
			hx_array_target_243 := self.params
			hx_array_index_244 := i
			hx_obj_242 := map[string]any{}
			hx_obj_242["name"] = name
			hx_obj_242["value"] = value
			hx_array_target_243.Set(hx_array_index_244, hx_obj_242)
			return
		}
	}
	hx_arr_250 := self.params
	hx_obj_251 := map[string]any{}
	hx_obj_251["name"] = name
	hx_obj_251["value"] = value
	hx_arr_250.Push(hx_obj_251)
}

func (self *haxe__http__HttpBase) addParameter(name *string, value *string) {
	hx_arr_252 := self.params
	hx_obj_253 := map[string]any{}
	hx_obj_253["name"] = name
	hx_obj_253["value"] = value
	hx_arr_252.Push(hx_obj_253)
}

func (self *haxe__http__HttpBase) setPostData(data *string) {
	self.postData = data
	self.postBytes = nil
}

func (self *haxe__http__HttpBase) setPostBytes(data *haxe__io__Bytes) {
	self.postBytes = data
	self.postData = nil
}

func (self *haxe__http__HttpBase) request(post any) {
	hxrt.Throw(New_haxe__exceptions__NotImplementedException(nil, nil, func() map[string]any {
		hx_obj_254 := map[string]any{}
		hx_obj_254["fileName"] = hxrt.StringFromLiteral("haxe/http/HttpBase.hx")
		hx_obj_254["lineNumber"] = 106
		hx_obj_254["className"] = hxrt.StringFromLiteral("haxe.http.HttpBase")
		hx_obj_254["methodName"] = hxrt.StringFromLiteral("request")
		return hx_obj_254
	}()))
}

func (self *haxe__http__HttpBase) hasOnData() bool {
	return !func() bool {
		var f1 any = any(self.onData)
		var f2 any = any(self.emptyOnData)
		return hxrt.ReflectCompareMethods(f1, f2)
	}()
}

func (self *haxe__http__HttpBase) success(data *haxe__io__Bytes) {
	self.responseBytes = data
	self.responseAsString = nil
	if self.__hx_this.hasOnData() {
		s := self.__hx_this.get_responseData()
		if !hxrt.StringEqualStringPtr(s, nil) {
			self.onData(s)
		}
	}
	self.onBytes(data)
}

func (self *haxe__http__HttpBase) get_responseData() *string {
	if hxrt.StringEqualStringPtr(self.responseAsString, nil) && (self.responseBytes != nil) {
		self.responseAsString = self.responseBytes.__hx_this.getString(0, self.responseBytes.length, haxe__io__Encoding_UTF8)
	}
	return self.responseAsString
}
