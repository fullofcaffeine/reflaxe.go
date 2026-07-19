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
		hx_post_234 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_234
		if hxrt.StringEqualStringPtr(func(hx_obj_240 map[string]any) *string {
			hx_field_241 := hx_obj_240["name"]
			if hx_field_241 == nil {
				var hx_zero_242 *string
				return hx_zero_242
			}
			return hx_field_241.(*string)
		}(func(hx_value_238 any) map[string]any {
			if hx_value_238 == nil {
				var hx_zero_239 map[string]any
				return hx_zero_239
			}
			return hx_value_238.(map[string]any)
		}(self.headers.Get(i))), name) {
			hx_array_target_236 := self.headers
			hx_array_index_237 := i
			hx_obj_235 := map[string]any{}
			hx_obj_235["name"] = name
			hx_obj_235["value"] = value
			hx_array_target_236.Set(hx_array_index_237, hx_obj_235)
			return
		}
	}
	hx_arr_243 := self.headers
	hx_obj_244 := map[string]any{}
	hx_obj_244["name"] = name
	hx_obj_244["value"] = value
	hx_arr_243.Push(hx_obj_244)
}

func (self *haxe__http__HttpBase) addHeader(header *string, value *string) {
	hx_arr_245 := self.headers
	hx_obj_246 := map[string]any{}
	hx_obj_246["name"] = header
	hx_obj_246["value"] = value
	hx_arr_245.Push(hx_obj_246)
}

func (self *haxe__http__HttpBase) setParameter(name *string, value *string) {
	_g := 0
	_g1 := self.params.Len()
	for _g < _g1 {
		hx_post_247 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_247
		if hxrt.StringEqualStringPtr(func(hx_obj_253 map[string]any) *string {
			hx_field_254 := hx_obj_253["name"]
			if hx_field_254 == nil {
				var hx_zero_255 *string
				return hx_zero_255
			}
			return hx_field_254.(*string)
		}(func(hx_value_251 any) map[string]any {
			if hx_value_251 == nil {
				var hx_zero_252 map[string]any
				return hx_zero_252
			}
			return hx_value_251.(map[string]any)
		}(self.params.Get(i))), name) {
			hx_array_target_249 := self.params
			hx_array_index_250 := i
			hx_obj_248 := map[string]any{}
			hx_obj_248["name"] = name
			hx_obj_248["value"] = value
			hx_array_target_249.Set(hx_array_index_250, hx_obj_248)
			return
		}
	}
	hx_arr_256 := self.params
	hx_obj_257 := map[string]any{}
	hx_obj_257["name"] = name
	hx_obj_257["value"] = value
	hx_arr_256.Push(hx_obj_257)
}

func (self *haxe__http__HttpBase) addParameter(name *string, value *string) {
	hx_arr_258 := self.params
	hx_obj_259 := map[string]any{}
	hx_obj_259["name"] = name
	hx_obj_259["value"] = value
	hx_arr_258.Push(hx_obj_259)
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
		hx_obj_260 := map[string]any{}
		hx_obj_260["fileName"] = hxrt.StringFromLiteral("haxe/http/HttpBase.hx")
		hx_obj_260["lineNumber"] = 106
		hx_obj_260["className"] = hxrt.StringFromLiteral("haxe.http.HttpBase")
		hx_obj_260["methodName"] = hxrt.StringFromLiteral("request")
		return hx_obj_260
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
			func(hx_fn func(*string), hx_arg_0 *string) {
				if hx_fn == nil {
					hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
					return
				}
				hx_fn(hx_arg_0)
			}(self.onData, s)
		}
	}
	func(hx_fn func(*haxe__io__Bytes), hx_arg_0 *haxe__io__Bytes) {
		if hx_fn == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
			return
		}
		hx_fn(hx_arg_0)
	}(self.onBytes, data)
}

func (self *haxe__http__HttpBase) get_responseData() *string {
	if hxrt.StringEqualStringPtr(self.responseAsString, nil) && (self.responseBytes != nil) {
		self.responseAsString = self.responseBytes.__hx_this.getString(0, self.responseBytes.length, haxe__io__Encoding_UTF8)
	}
	return self.responseAsString
}
