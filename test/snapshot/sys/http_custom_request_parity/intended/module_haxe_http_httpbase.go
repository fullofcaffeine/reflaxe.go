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
	headers          []map[string]any
	params           []map[string]any
	emptyOnData      func(*string)
}

func New_haxe__http__HttpBase(url *string) *haxe__http__HttpBase {
	self := &haxe__http__HttpBase{}
	self.__hx_this = self
	self.url = url
	self.headers = []map[string]any{}
	self.params = []map[string]any{}
	self.emptyOnData = self.onData
	return self
}

func (self *haxe__http__HttpBase) setHeader(name *string, value *string) {
	_g := 0
	_g1 := len(self.headers)
	for _g < _g1 {
		hx_post_14 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_14
		if hxrt.StringEqualStringPtr(func(hx_obj_16 map[string]any) *string {
			hx_field_17 := hx_obj_16["name"]
			if hx_field_17 == nil {
				var hx_zero_18 *string
				return hx_zero_18
			}
			return hx_field_17.(*string)
		}(self.headers[i]), name) {
			hx_obj_15 := map[string]any{}
			hx_obj_15["name"] = name
			hx_obj_15["value"] = value
			self.headers[i] = hx_obj_15
			return
		}
	}
	hx_arr_19 := self.headers
	hx_arr_19 = append(hx_arr_19, func() map[string]any {
		hx_obj_20 := map[string]any{}
		hx_obj_20["name"] = name
		hx_obj_20["value"] = value
		return hx_obj_20
	}())
	self.headers = hx_arr_19
}

func (self *haxe__http__HttpBase) addHeader(header *string, value *string) {
	hx_arr_21 := self.headers
	hx_arr_21 = append(hx_arr_21, func() map[string]any {
		hx_obj_22 := map[string]any{}
		hx_obj_22["name"] = header
		hx_obj_22["value"] = value
		return hx_obj_22
	}())
	self.headers = hx_arr_21
}

func (self *haxe__http__HttpBase) setParameter(name *string, value *string) {
	_g := 0
	_g1 := len(self.params)
	for _g < _g1 {
		hx_post_23 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_23
		if hxrt.StringEqualStringPtr(func(hx_obj_25 map[string]any) *string {
			hx_field_26 := hx_obj_25["name"]
			if hx_field_26 == nil {
				var hx_zero_27 *string
				return hx_zero_27
			}
			return hx_field_26.(*string)
		}(self.params[i]), name) {
			hx_obj_24 := map[string]any{}
			hx_obj_24["name"] = name
			hx_obj_24["value"] = value
			self.params[i] = hx_obj_24
			return
		}
	}
	hx_arr_28 := self.params
	hx_arr_28 = append(hx_arr_28, func() map[string]any {
		hx_obj_29 := map[string]any{}
		hx_obj_29["name"] = name
		hx_obj_29["value"] = value
		return hx_obj_29
	}())
	self.params = hx_arr_28
}

func (self *haxe__http__HttpBase) addParameter(name *string, value *string) {
	hx_arr_30 := self.params
	hx_arr_30 = append(hx_arr_30, func() map[string]any {
		hx_obj_31 := map[string]any{}
		hx_obj_31["name"] = name
		hx_obj_31["value"] = value
		return hx_obj_31
	}())
	self.params = hx_arr_30
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
		hx_obj_32 := map[string]any{}
		hx_obj_32["fileName"] = hxrt.StringFromLiteral("haxe/http/HttpBase.hx")
		hx_obj_32["lineNumber"] = 106
		hx_obj_32["className"] = hxrt.StringFromLiteral("haxe.http.HttpBase")
		hx_obj_32["methodName"] = hxrt.StringFromLiteral("request")
		return hx_obj_32
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
