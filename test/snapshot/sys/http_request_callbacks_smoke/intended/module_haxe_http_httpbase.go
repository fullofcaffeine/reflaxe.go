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
		hx_post_14 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_14
		if hxrt.StringEqualStringPtr(func(hx_obj_20 map[string]any) *string {
			hx_field_21 := hx_obj_20["name"]
			if hx_field_21 == nil {
				var hx_zero_22 *string
				return hx_zero_22
			}
			return hx_field_21.(*string)
		}(func(hx_value_18 any) map[string]any {
			if hx_value_18 == nil {
				var hx_zero_19 map[string]any
				return hx_zero_19
			}
			return hx_value_18.(map[string]any)
		}(self.headers.Get(i))), name) {
			hx_array_target_16 := self.headers
			hx_array_index_17 := i
			hx_obj_15 := map[string]any{}
			hx_obj_15["name"] = name
			hx_obj_15["value"] = value
			hx_array_target_16.Set(hx_array_index_17, hx_obj_15)
			return
		}
	}
	hx_arr_23 := self.headers
	hx_obj_24 := map[string]any{}
	hx_obj_24["name"] = name
	hx_obj_24["value"] = value
	hx_arr_23.Push(hx_obj_24)
}

func (self *haxe__http__HttpBase) addHeader(header *string, value *string) {
	hx_arr_25 := self.headers
	hx_obj_26 := map[string]any{}
	hx_obj_26["name"] = header
	hx_obj_26["value"] = value
	hx_arr_25.Push(hx_obj_26)
}

func (self *haxe__http__HttpBase) setParameter(name *string, value *string) {
	_g := 0
	_g1 := self.params.Len()
	for _g < _g1 {
		hx_post_27 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_27
		if hxrt.StringEqualStringPtr(func(hx_obj_33 map[string]any) *string {
			hx_field_34 := hx_obj_33["name"]
			if hx_field_34 == nil {
				var hx_zero_35 *string
				return hx_zero_35
			}
			return hx_field_34.(*string)
		}(func(hx_value_31 any) map[string]any {
			if hx_value_31 == nil {
				var hx_zero_32 map[string]any
				return hx_zero_32
			}
			return hx_value_31.(map[string]any)
		}(self.params.Get(i))), name) {
			hx_array_target_29 := self.params
			hx_array_index_30 := i
			hx_obj_28 := map[string]any{}
			hx_obj_28["name"] = name
			hx_obj_28["value"] = value
			hx_array_target_29.Set(hx_array_index_30, hx_obj_28)
			return
		}
	}
	hx_arr_36 := self.params
	hx_obj_37 := map[string]any{}
	hx_obj_37["name"] = name
	hx_obj_37["value"] = value
	hx_arr_36.Push(hx_obj_37)
}

func (self *haxe__http__HttpBase) addParameter(name *string, value *string) {
	hx_arr_38 := self.params
	hx_obj_39 := map[string]any{}
	hx_obj_39["name"] = name
	hx_obj_39["value"] = value
	hx_arr_38.Push(hx_obj_39)
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
		hx_obj_40 := map[string]any{}
		hx_obj_40["fileName"] = hxrt.StringFromLiteral("haxe/http/HttpBase.hx")
		hx_obj_40["lineNumber"] = 106
		hx_obj_40["className"] = hxrt.StringFromLiteral("haxe.http.HttpBase")
		hx_obj_40["methodName"] = hxrt.StringFromLiteral("request")
		return hx_obj_40
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
