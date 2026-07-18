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
		hx_post_88 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_88
		if hxrt.StringEqualStringPtr(func(hx_obj_94 map[string]any) *string {
			hx_field_95 := hx_obj_94["name"]
			if hx_field_95 == nil {
				var hx_zero_96 *string
				return hx_zero_96
			}
			return hx_field_95.(*string)
		}(func(hx_value_92 any) map[string]any {
			if hx_value_92 == nil {
				var hx_zero_93 map[string]any
				return hx_zero_93
			}
			return hx_value_92.(map[string]any)
		}(self.headers.Get(i))), name) {
			hx_array_target_90 := self.headers
			hx_array_index_91 := i
			hx_obj_89 := map[string]any{}
			hx_obj_89["name"] = name
			hx_obj_89["value"] = value
			hx_array_target_90.Set(hx_array_index_91, hx_obj_89)
			return
		}
	}
	hx_arr_97 := self.headers
	hx_obj_98 := map[string]any{}
	hx_obj_98["name"] = name
	hx_obj_98["value"] = value
	hx_arr_97.Push(hx_obj_98)
}

func (self *haxe__http__HttpBase) addHeader(header *string, value *string) {
	hx_arr_99 := self.headers
	hx_obj_100 := map[string]any{}
	hx_obj_100["name"] = header
	hx_obj_100["value"] = value
	hx_arr_99.Push(hx_obj_100)
}

func (self *haxe__http__HttpBase) setParameter(name *string, value *string) {
	_g := 0
	_g1 := self.params.Len()
	for _g < _g1 {
		hx_post_101 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_101
		if hxrt.StringEqualStringPtr(func(hx_obj_107 map[string]any) *string {
			hx_field_108 := hx_obj_107["name"]
			if hx_field_108 == nil {
				var hx_zero_109 *string
				return hx_zero_109
			}
			return hx_field_108.(*string)
		}(func(hx_value_105 any) map[string]any {
			if hx_value_105 == nil {
				var hx_zero_106 map[string]any
				return hx_zero_106
			}
			return hx_value_105.(map[string]any)
		}(self.params.Get(i))), name) {
			hx_array_target_103 := self.params
			hx_array_index_104 := i
			hx_obj_102 := map[string]any{}
			hx_obj_102["name"] = name
			hx_obj_102["value"] = value
			hx_array_target_103.Set(hx_array_index_104, hx_obj_102)
			return
		}
	}
	hx_arr_110 := self.params
	hx_obj_111 := map[string]any{}
	hx_obj_111["name"] = name
	hx_obj_111["value"] = value
	hx_arr_110.Push(hx_obj_111)
}

func (self *haxe__http__HttpBase) addParameter(name *string, value *string) {
	hx_arr_112 := self.params
	hx_obj_113 := map[string]any{}
	hx_obj_113["name"] = name
	hx_obj_113["value"] = value
	hx_arr_112.Push(hx_obj_113)
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
		hx_obj_114 := map[string]any{}
		hx_obj_114["fileName"] = hxrt.StringFromLiteral("haxe/http/HttpBase.hx")
		hx_obj_114["lineNumber"] = 106
		hx_obj_114["className"] = hxrt.StringFromLiteral("haxe.http.HttpBase")
		hx_obj_114["methodName"] = hxrt.StringFromLiteral("request")
		return hx_obj_114
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
