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
	onData(data *string)
	onBytes(data *haxe__io__Bytes)
	onError(msg *string)
	onStatus(status int)
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
}

func New_haxe__http__HttpBase(url *string) *haxe__http__HttpBase {
	self := &haxe__http__HttpBase{}
	self.__hx_this = self
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
		hx_post_20 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_20
		if hxrt.StringEqualStringPtr(func(hx_obj_26 map[string]any) *string {
			hx_field_27 := hx_obj_26["name"]
			if hx_field_27 == nil {
				var hx_zero_28 *string
				return hx_zero_28
			}
			return hx_field_27.(*string)
		}(func(hx_value_24 any) map[string]any {
			if hx_value_24 == nil {
				var hx_zero_25 map[string]any
				return hx_zero_25
			}
			return hx_value_24.(map[string]any)
		}(self.headers.Get(i))), name) {
			hx_array_target_22 := self.headers
			hx_array_index_23 := i
			hx_obj_21 := map[string]any{}
			hx_obj_21["name"] = name
			hx_obj_21["value"] = value
			hx_array_target_22.Set(hx_array_index_23, hx_obj_21)
			return
		}
	}
	hx_arr_29 := self.headers
	hx_obj_30 := map[string]any{}
	hx_obj_30["name"] = name
	hx_obj_30["value"] = value
	hx_arr_29.Push(hx_obj_30)
}

func (self *haxe__http__HttpBase) addHeader(header *string, value *string) {
	hx_arr_31 := self.headers
	hx_obj_32 := map[string]any{}
	hx_obj_32["name"] = header
	hx_obj_32["value"] = value
	hx_arr_31.Push(hx_obj_32)
}

func (self *haxe__http__HttpBase) setParameter(name *string, value *string) {
	_g := 0
	_g1 := self.params.Len()
	for _g < _g1 {
		hx_post_33 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_33
		if hxrt.StringEqualStringPtr(func(hx_obj_39 map[string]any) *string {
			hx_field_40 := hx_obj_39["name"]
			if hx_field_40 == nil {
				var hx_zero_41 *string
				return hx_zero_41
			}
			return hx_field_40.(*string)
		}(func(hx_value_37 any) map[string]any {
			if hx_value_37 == nil {
				var hx_zero_38 map[string]any
				return hx_zero_38
			}
			return hx_value_37.(map[string]any)
		}(self.params.Get(i))), name) {
			hx_array_target_35 := self.params
			hx_array_index_36 := i
			hx_obj_34 := map[string]any{}
			hx_obj_34["name"] = name
			hx_obj_34["value"] = value
			hx_array_target_35.Set(hx_array_index_36, hx_obj_34)
			return
		}
	}
	hx_arr_42 := self.params
	hx_obj_43 := map[string]any{}
	hx_obj_43["name"] = name
	hx_obj_43["value"] = value
	hx_arr_42.Push(hx_obj_43)
}

func (self *haxe__http__HttpBase) addParameter(name *string, value *string) {
	hx_arr_44 := self.params
	hx_obj_45 := map[string]any{}
	hx_obj_45["name"] = name
	hx_obj_45["value"] = value
	hx_arr_44.Push(hx_obj_45)
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
		hx_obj_46 := map[string]any{}
		hx_obj_46["fileName"] = hxrt.StringFromLiteral("haxe/http/HttpBase.hx")
		hx_obj_46["lineNumber"] = 106
		hx_obj_46["className"] = hxrt.StringFromLiteral("haxe.http.HttpBase")
		hx_obj_46["methodName"] = hxrt.StringFromLiteral("request")
		return hx_obj_46
	}()))
}

func (self *haxe__http__HttpBase) onData(data *string) {
}

func (self *haxe__http__HttpBase) onBytes(data *haxe__io__Bytes) {
}

func (self *haxe__http__HttpBase) onError(msg *string) {
}

func (self *haxe__http__HttpBase) onStatus(status int) {
}

func (self *haxe__http__HttpBase) hasOnData() bool {
	return !Reflect_compareMethods(self.onData, self.emptyOnData)
}

func (self *haxe__http__HttpBase) success(data *haxe__io__Bytes) {
	self.responseBytes = data
	self.responseAsString = nil
	if self.hasOnData() {
		s := self.get_responseData()
		if !hxrt.StringEqualStringPtr(s, nil) {
			self.onData(s)
		}
	}
	self.onBytes(data)
}

func (self *haxe__http__HttpBase) get_responseData() *string {
	if hxrt.StringEqualStringPtr(self.responseAsString, nil) && (self.responseBytes != nil) {
		self.responseAsString = self.responseBytes.getString(0, self.responseBytes.length, haxe__io__Encoding_UTF8)
	}
	return self.responseAsString
}
