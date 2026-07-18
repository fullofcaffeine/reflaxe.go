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
		hx_post_233 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_233
		if hxrt.StringEqualStringPtr(func(hx_obj_239 map[string]any) *string {
			hx_field_240 := hx_obj_239["name"]
			if hx_field_240 == nil {
				var hx_zero_241 *string
				return hx_zero_241
			}
			return hx_field_240.(*string)
		}(func(hx_value_237 any) map[string]any {
			if hx_value_237 == nil {
				var hx_zero_238 map[string]any
				return hx_zero_238
			}
			return hx_value_237.(map[string]any)
		}(self.headers.Get(i))), name) {
			hx_array_target_235 := self.headers
			hx_array_index_236 := i
			hx_obj_234 := map[string]any{}
			hx_obj_234["name"] = name
			hx_obj_234["value"] = value
			hx_array_target_235.Set(hx_array_index_236, hx_obj_234)
			return
		}
	}
	hx_arr_242 := self.headers
	hx_obj_243 := map[string]any{}
	hx_obj_243["name"] = name
	hx_obj_243["value"] = value
	hx_arr_242.Push(hx_obj_243)
}

func (self *haxe__http__HttpBase) addHeader(header *string, value *string) {
	hx_arr_244 := self.headers
	hx_obj_245 := map[string]any{}
	hx_obj_245["name"] = header
	hx_obj_245["value"] = value
	hx_arr_244.Push(hx_obj_245)
}

func (self *haxe__http__HttpBase) setParameter(name *string, value *string) {
	_g := 0
	_g1 := self.params.Len()
	for _g < _g1 {
		hx_post_246 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_246
		if hxrt.StringEqualStringPtr(func(hx_obj_252 map[string]any) *string {
			hx_field_253 := hx_obj_252["name"]
			if hx_field_253 == nil {
				var hx_zero_254 *string
				return hx_zero_254
			}
			return hx_field_253.(*string)
		}(func(hx_value_250 any) map[string]any {
			if hx_value_250 == nil {
				var hx_zero_251 map[string]any
				return hx_zero_251
			}
			return hx_value_250.(map[string]any)
		}(self.params.Get(i))), name) {
			hx_array_target_248 := self.params
			hx_array_index_249 := i
			hx_obj_247 := map[string]any{}
			hx_obj_247["name"] = name
			hx_obj_247["value"] = value
			hx_array_target_248.Set(hx_array_index_249, hx_obj_247)
			return
		}
	}
	hx_arr_255 := self.params
	hx_obj_256 := map[string]any{}
	hx_obj_256["name"] = name
	hx_obj_256["value"] = value
	hx_arr_255.Push(hx_obj_256)
}

func (self *haxe__http__HttpBase) addParameter(name *string, value *string) {
	hx_arr_257 := self.params
	hx_obj_258 := map[string]any{}
	hx_obj_258["name"] = name
	hx_obj_258["value"] = value
	hx_arr_257.Push(hx_obj_258)
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
		hx_obj_259 := map[string]any{}
		hx_obj_259["fileName"] = hxrt.StringFromLiteral("haxe/http/HttpBase.hx")
		hx_obj_259["lineNumber"] = 106
		hx_obj_259["className"] = hxrt.StringFromLiteral("haxe.http.HttpBase")
		hx_obj_259["methodName"] = hxrt.StringFromLiteral("request")
		return hx_obj_259
	}()))
}

func (self *haxe__http__HttpBase) hasOnData() bool {
	return !Reflect_compareMethods(self.onData, self.emptyOnData)
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
