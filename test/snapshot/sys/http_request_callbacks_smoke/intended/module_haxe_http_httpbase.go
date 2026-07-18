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
		hx_post_24 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_24
		if hxrt.StringEqualStringPtr(func(hx_obj_30 map[string]any) *string {
			hx_field_31 := hx_obj_30["name"]
			if hx_field_31 == nil {
				var hx_zero_32 *string
				return hx_zero_32
			}
			return hx_field_31.(*string)
		}(func(hx_value_28 any) map[string]any {
			if hx_value_28 == nil {
				var hx_zero_29 map[string]any
				return hx_zero_29
			}
			return hx_value_28.(map[string]any)
		}(self.headers.Get(i))), name) {
			hx_array_target_26 := self.headers
			hx_array_index_27 := i
			hx_obj_25 := map[string]any{}
			hx_obj_25["name"] = name
			hx_obj_25["value"] = value
			hx_array_target_26.Set(hx_array_index_27, hx_obj_25)
			return
		}
	}
	hx_arr_33 := self.headers
	hx_obj_34 := map[string]any{}
	hx_obj_34["name"] = name
	hx_obj_34["value"] = value
	hx_arr_33.Push(hx_obj_34)
}

func (self *haxe__http__HttpBase) addHeader(header *string, value *string) {
	hx_arr_35 := self.headers
	hx_obj_36 := map[string]any{}
	hx_obj_36["name"] = header
	hx_obj_36["value"] = value
	hx_arr_35.Push(hx_obj_36)
}

func (self *haxe__http__HttpBase) setParameter(name *string, value *string) {
	_g := 0
	_g1 := self.params.Len()
	for _g < _g1 {
		hx_post_37 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_37
		if hxrt.StringEqualStringPtr(func(hx_obj_43 map[string]any) *string {
			hx_field_44 := hx_obj_43["name"]
			if hx_field_44 == nil {
				var hx_zero_45 *string
				return hx_zero_45
			}
			return hx_field_44.(*string)
		}(func(hx_value_41 any) map[string]any {
			if hx_value_41 == nil {
				var hx_zero_42 map[string]any
				return hx_zero_42
			}
			return hx_value_41.(map[string]any)
		}(self.params.Get(i))), name) {
			hx_array_target_39 := self.params
			hx_array_index_40 := i
			hx_obj_38 := map[string]any{}
			hx_obj_38["name"] = name
			hx_obj_38["value"] = value
			hx_array_target_39.Set(hx_array_index_40, hx_obj_38)
			return
		}
	}
	hx_arr_46 := self.params
	hx_obj_47 := map[string]any{}
	hx_obj_47["name"] = name
	hx_obj_47["value"] = value
	hx_arr_46.Push(hx_obj_47)
}

func (self *haxe__http__HttpBase) addParameter(name *string, value *string) {
	hx_arr_48 := self.params
	hx_obj_49 := map[string]any{}
	hx_obj_49["name"] = name
	hx_obj_49["value"] = value
	hx_arr_48.Push(hx_obj_49)
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
		hx_obj_50 := map[string]any{}
		hx_obj_50["fileName"] = hxrt.StringFromLiteral("haxe/http/HttpBase.hx")
		hx_obj_50["lineNumber"] = 106
		hx_obj_50["className"] = hxrt.StringFromLiteral("haxe.http.HttpBase")
		hx_obj_50["methodName"] = hxrt.StringFromLiteral("request")
		return hx_obj_50
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
	if self.__hx_this.hasOnData() {
		s := self.__hx_this.get_responseData()
		if !hxrt.StringEqualStringPtr(s, nil) {
			self.__hx_this.onData(s)
		}
	}
	self.__hx_this.onBytes(data)
}

func (self *haxe__http__HttpBase) get_responseData() *string {
	if hxrt.StringEqualStringPtr(self.responseAsString, nil) && (self.responseBytes != nil) {
		self.responseAsString = self.responseBytes.__hx_this.getString(0, self.responseBytes.length, haxe__io__Encoding_UTF8)
	}
	return self.responseAsString
}
