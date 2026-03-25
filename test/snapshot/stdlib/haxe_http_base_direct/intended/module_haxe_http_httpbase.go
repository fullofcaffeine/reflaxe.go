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
		hx_post_3 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_3
		if hxrt.StringEqualStringPtr(func(hx_obj_4 map[string]any) *string {
			hx_field_5 := hx_obj_4["name"]
			if hx_field_5 == nil {
				var hx_zero_6 *string
				return hx_zero_6
			}
			return hx_field_5.(*string)
		}(self.headers[i]), name) {
			hx_obj_7 := map[string]any{}
			hx_obj_7["name"] = name
			hx_obj_7["value"] = value
			self.headers[i] = hx_obj_7
			return
		}
	}
	hx_arr_8 := self.headers
	hx_arr_8 = append(hx_arr_8, func() map[string]any {
		hx_obj_9 := map[string]any{}
		hx_obj_9["name"] = name
		hx_obj_9["value"] = value
		return hx_obj_9
	}())
	self.headers = hx_arr_8
}

func (self *haxe__http__HttpBase) addHeader(header *string, value *string) {
	hx_arr_10 := self.headers
	hx_arr_10 = append(hx_arr_10, func() map[string]any {
		hx_obj_11 := map[string]any{}
		hx_obj_11["name"] = header
		hx_obj_11["value"] = value
		return hx_obj_11
	}())
	self.headers = hx_arr_10
}

func (self *haxe__http__HttpBase) setParameter(name *string, value *string) {
	_g := 0
	_g1 := len(self.params)
	for _g < _g1 {
		hx_post_12 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_12
		if hxrt.StringEqualStringPtr(func(hx_obj_13 map[string]any) *string {
			hx_field_14 := hx_obj_13["name"]
			if hx_field_14 == nil {
				var hx_zero_15 *string
				return hx_zero_15
			}
			return hx_field_14.(*string)
		}(self.params[i]), name) {
			hx_obj_16 := map[string]any{}
			hx_obj_16["name"] = name
			hx_obj_16["value"] = value
			self.params[i] = hx_obj_16
			return
		}
	}
	hx_arr_17 := self.params
	hx_arr_17 = append(hx_arr_17, func() map[string]any {
		hx_obj_18 := map[string]any{}
		hx_obj_18["name"] = name
		hx_obj_18["value"] = value
		return hx_obj_18
	}())
	self.params = hx_arr_17
}

func (self *haxe__http__HttpBase) addParameter(name *string, value *string) {
	hx_arr_19 := self.params
	hx_arr_19 = append(hx_arr_19, func() map[string]any {
		hx_obj_20 := map[string]any{}
		hx_obj_20["name"] = name
		hx_obj_20["value"] = value
		return hx_obj_20
	}())
	self.params = hx_arr_19
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
		hx_obj_21 := map[string]any{}
		hx_obj_21["fileName"] = hxrt.StringFromLiteral("../../../../std/haxe/http/HttpBase.cross.hx")
		hx_obj_21["lineNumber"] = 105
		hx_obj_21["className"] = hxrt.StringFromLiteral("haxe.http.HttpBase")
		hx_obj_21["methodName"] = hxrt.StringFromLiteral("request")
		return hx_obj_21
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
