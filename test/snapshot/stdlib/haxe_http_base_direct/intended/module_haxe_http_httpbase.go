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
		hx_post_3 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_3
		if hxrt.StringEqualStringPtr(func(hx_obj_9 map[string]any) *string {
			hx_field_10 := hx_obj_9["name"]
			if hx_field_10 == nil {
				var hx_zero_11 *string
				return hx_zero_11
			}
			return hx_field_10.(*string)
		}(func(hx_value_7 any) map[string]any {
			if hx_value_7 == nil {
				var hx_zero_8 map[string]any
				return hx_zero_8
			}
			return hx_value_7.(map[string]any)
		}(self.headers.Get(i))), name) {
			hx_array_target_5 := self.headers
			hx_array_index_6 := i
			hx_obj_4 := map[string]any{}
			hx_obj_4["name"] = name
			hx_obj_4["value"] = value
			hx_array_target_5.Set(hx_array_index_6, hx_obj_4)
			return
		}
	}
	hx_arr_12 := self.headers
	hx_obj_13 := map[string]any{}
	hx_obj_13["name"] = name
	hx_obj_13["value"] = value
	hx_arr_12.Push(hx_obj_13)
}

func (self *haxe__http__HttpBase) addHeader(header *string, value *string) {
	hx_arr_14 := self.headers
	hx_obj_15 := map[string]any{}
	hx_obj_15["name"] = header
	hx_obj_15["value"] = value
	hx_arr_14.Push(hx_obj_15)
}

func (self *haxe__http__HttpBase) setParameter(name *string, value *string) {
	_g := 0
	_g1 := self.params.Len()
	for _g < _g1 {
		hx_post_16 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_16
		if hxrt.StringEqualStringPtr(func(hx_obj_22 map[string]any) *string {
			hx_field_23 := hx_obj_22["name"]
			if hx_field_23 == nil {
				var hx_zero_24 *string
				return hx_zero_24
			}
			return hx_field_23.(*string)
		}(func(hx_value_20 any) map[string]any {
			if hx_value_20 == nil {
				var hx_zero_21 map[string]any
				return hx_zero_21
			}
			return hx_value_20.(map[string]any)
		}(self.params.Get(i))), name) {
			hx_array_target_18 := self.params
			hx_array_index_19 := i
			hx_obj_17 := map[string]any{}
			hx_obj_17["name"] = name
			hx_obj_17["value"] = value
			hx_array_target_18.Set(hx_array_index_19, hx_obj_17)
			return
		}
	}
	hx_arr_25 := self.params
	hx_obj_26 := map[string]any{}
	hx_obj_26["name"] = name
	hx_obj_26["value"] = value
	hx_arr_25.Push(hx_obj_26)
}

func (self *haxe__http__HttpBase) addParameter(name *string, value *string) {
	hx_arr_27 := self.params
	hx_obj_28 := map[string]any{}
	hx_obj_28["name"] = name
	hx_obj_28["value"] = value
	hx_arr_27.Push(hx_obj_28)
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
		hx_obj_29 := map[string]any{}
		hx_obj_29["fileName"] = hxrt.StringFromLiteral("haxe/http/HttpBase.hx")
		hx_obj_29["lineNumber"] = 106
		hx_obj_29["className"] = hxrt.StringFromLiteral("haxe.http.HttpBase")
		hx_obj_29["methodName"] = hxrt.StringFromLiteral("request")
		return hx_obj_29
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
