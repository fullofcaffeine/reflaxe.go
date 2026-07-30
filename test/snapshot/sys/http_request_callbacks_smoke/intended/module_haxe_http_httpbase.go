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
		hx_post_202 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_202
		if hxrt.StringEqualStringPtr(func(hx_obj_208 map[string]any) *string {
			hx_field_209 := hx_obj_208["name"]
			if hx_field_209 == nil {
				var hx_zero_210 *string
				return hx_zero_210
			}
			return hx_field_209.(*string)
		}(func(hx_value_206 any) map[string]any {
			if hx_value_206 == nil {
				var hx_zero_207 map[string]any
				return hx_zero_207
			}
			return hx_value_206.(map[string]any)
		}(self.headers.Get(i))), name) {
			hx_array_target_204 := self.headers
			hx_array_index_205 := i
			hx_obj_203 := map[string]any{}
			hx_obj_203["name"] = name
			hx_obj_203["value"] = value
			hx_array_target_204.Set(hx_array_index_205, hx_obj_203)
			return
		}
	}
	hx_arr_211 := self.headers
	hx_obj_212 := map[string]any{}
	hx_obj_212["name"] = name
	hx_obj_212["value"] = value
	hx_arr_211.Push(hx_obj_212)
}

func (self *haxe__http__HttpBase) addHeader(header *string, value *string) {
	hx_arr_213 := self.headers
	hx_obj_214 := map[string]any{}
	hx_obj_214["name"] = header
	hx_obj_214["value"] = value
	hx_arr_213.Push(hx_obj_214)
}

func (self *haxe__http__HttpBase) setParameter(name *string, value *string) {
	_g := 0
	_g1 := self.params.Len()
	for _g < _g1 {
		hx_post_215 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_215
		if hxrt.StringEqualStringPtr(func(hx_obj_221 map[string]any) *string {
			hx_field_222 := hx_obj_221["name"]
			if hx_field_222 == nil {
				var hx_zero_223 *string
				return hx_zero_223
			}
			return hx_field_222.(*string)
		}(func(hx_value_219 any) map[string]any {
			if hx_value_219 == nil {
				var hx_zero_220 map[string]any
				return hx_zero_220
			}
			return hx_value_219.(map[string]any)
		}(self.params.Get(i))), name) {
			hx_array_target_217 := self.params
			hx_array_index_218 := i
			hx_obj_216 := map[string]any{}
			hx_obj_216["name"] = name
			hx_obj_216["value"] = value
			hx_array_target_217.Set(hx_array_index_218, hx_obj_216)
			return
		}
	}
	hx_arr_224 := self.params
	hx_obj_225 := map[string]any{}
	hx_obj_225["name"] = name
	hx_obj_225["value"] = value
	hx_arr_224.Push(hx_obj_225)
}

func (self *haxe__http__HttpBase) addParameter(name *string, value *string) {
	hx_arr_226 := self.params
	hx_obj_227 := map[string]any{}
	hx_obj_227["name"] = name
	hx_obj_227["value"] = value
	hx_arr_226.Push(hx_obj_227)
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
		hx_obj_228 := map[string]any{}
		hx_obj_228["fileName"] = hxrt.StringFromLiteral("haxe/http/HttpBase.hx")
		hx_obj_228["lineNumber"] = 106
		hx_obj_228["className"] = hxrt.StringFromLiteral("haxe.http.HttpBase")
		hx_obj_228["methodName"] = hxrt.StringFromLiteral("request")
		return hx_obj_228
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
