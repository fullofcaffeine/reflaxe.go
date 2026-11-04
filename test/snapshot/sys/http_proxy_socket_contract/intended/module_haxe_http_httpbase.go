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
	emptyOnData      *hxrt.Function[func(*string)]
	onData           *hxrt.Function[func(*string)]
	onBytes          *hxrt.Function[func(*haxe__io__Bytes)]
	onError          *hxrt.Function[func(*string)]
	onStatus         *hxrt.Function[func(int)]
}

func (self *haxe__http__HttpBase) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe__http__HttpBase(url *string) *haxe__http__HttpBase {
	return hx_init_haxe__http__HttpBase(hx_alloc_haxe__http__HttpBase(nil), url)
}

func hx_alloc_haxe__http__HttpBase(owner any) *haxe__http__HttpBase {
	self := &haxe__http__HttpBase{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_haxe__http__HttpBase)
	}
	return self
}

func hx_init_haxe__http__HttpBase(self *haxe__http__HttpBase, url *string) *haxe__http__HttpBase {
	self.onData = hxrt.DefaultDynamicFunction(hxrt.NewFunction(func(data *string) {
	}), func() any {
		return self.__hx_this
	}, "onData")
	self.onBytes = hxrt.DefaultDynamicFunction(hxrt.NewFunction(func(data *haxe__io__Bytes) {
	}), func() any {
		return self.__hx_this
	}, "onBytes")
	self.onError = hxrt.DefaultDynamicFunction(hxrt.NewFunction(func(msg *string) {
	}), func() any {
		return self.__hx_this
	}, "onError")
	self.onStatus = hxrt.DefaultDynamicFunction(hxrt.NewFunction(func(status int) {
	}), func() any {
		return self.__hx_this
	}, "onStatus")
	self.url = url
	self.headers = hxrt.NewArray()
	self.params = hxrt.NewArray()
	self.emptyOnData = hxrt.ReadDynamicFunction(self.onData)
	return self
}

func (self *haxe__http__HttpBase) setHeader(name *string, value *string) {
	_g := 0
	_g1 := self.headers.Len()
	for _g < _g1 {
		hx_post_1 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_1
		if hxrt.StringEqualStringPtr(func(hx_value_5 any) *string {
			if hx_value_5 == nil {
				var hx_zero_6 *string
				return hx_zero_6
			}
			return hx_value_5.(*string)
		}(hxrt.AnonymousGet(hx_structural_recover(self.headers.Get(i)), "name")), name) {
			hx_array_target_3 := self.headers
			hx_array_index_4 := i
			hx_obj_2 := map[string]any{}
			hx_obj_2["name"] = name
			hx_obj_2["value"] = value
			hx_array_target_3.Set(hx_array_index_4, hxrt.AnonymousFromMap(hx_obj_2))
			return
		}
	}
	hx_arr_7 := self.headers
	hx_obj_8 := map[string]any{}
	hx_obj_8["name"] = name
	hx_obj_8["value"] = value
	hx_arr_7.Push(hxrt.AnonymousFromMap(hx_obj_8))
}

func (self *haxe__http__HttpBase) addHeader(header *string, value *string) {
	hx_arr_9 := self.headers
	hx_obj_10 := map[string]any{}
	hx_obj_10["name"] = header
	hx_obj_10["value"] = value
	hx_arr_9.Push(hxrt.AnonymousFromMap(hx_obj_10))
}

func (self *haxe__http__HttpBase) setParameter(name *string, value *string) {
	_g := 0
	_g1 := self.params.Len()
	for _g < _g1 {
		hx_post_11 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_11
		if hxrt.StringEqualStringPtr(func(hx_value_15 any) *string {
			if hx_value_15 == nil {
				var hx_zero_16 *string
				return hx_zero_16
			}
			return hx_value_15.(*string)
		}(hxrt.AnonymousGet(hx_structural_recover(self.params.Get(i)), "name")), name) {
			hx_array_target_13 := self.params
			hx_array_index_14 := i
			hx_obj_12 := map[string]any{}
			hx_obj_12["name"] = name
			hx_obj_12["value"] = value
			hx_array_target_13.Set(hx_array_index_14, hxrt.AnonymousFromMap(hx_obj_12))
			return
		}
	}
	hx_arr_17 := self.params
	hx_obj_18 := map[string]any{}
	hx_obj_18["name"] = name
	hx_obj_18["value"] = value
	hx_arr_17.Push(hxrt.AnonymousFromMap(hx_obj_18))
}

func (self *haxe__http__HttpBase) addParameter(name *string, value *string) {
	hx_arr_19 := self.params
	hx_obj_20 := map[string]any{}
	hx_obj_20["name"] = name
	hx_obj_20["value"] = value
	hx_arr_19.Push(hxrt.AnonymousFromMap(hx_obj_20))
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
	hxrt.Throw(haxe__Exception_thrown(New_haxe__exceptions__NotImplementedException(nil, nil, func() hxrt.Anonymous {
		hx_obj_21 := map[string]any{}
		hx_obj_21["fileName"] = hxrt.StringFromLiteral("haxe/http/HttpBase.hx")
		hx_obj_21["lineNumber"] = 106
		hx_obj_21["className"] = hxrt.StringFromLiteral("haxe.http.HttpBase")
		hx_obj_21["methodName"] = hxrt.StringFromLiteral("request")
		return hxrt.AnonymousFromMap(hx_obj_21)
	}())))
}

func (self *haxe__http__HttpBase) hasOnData() bool {
	return !func() bool {
		var f1 any = any(hxrt.ReadDynamicFunction(self.onData))
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
			func() {
				var hx_callable_22 *hxrt.Function[func(*string)] = self.onData
				var hx_argument_23 *string = s
				hxrt.FunctionCall(hx_callable_22)(hx_argument_23)
			}()
		}
	}
	func() {
		var hx_callable_24 *hxrt.Function[func(*haxe__io__Bytes)] = self.onBytes
		var hx_argument_25 *haxe__io__Bytes = data
		hxrt.FunctionCall(hx_callable_24)(hx_argument_25)
	}()
}

func (self *haxe__http__HttpBase) get_responseData() *string {
	if hxrt.StringEqualStringPtr(self.responseAsString, nil) && (self.responseBytes != nil) {
		self.responseAsString = self.responseBytes.__hx_this.getString(0, self.responseBytes.length, haxe__io__Encoding_UTF8)
	}
	return self.responseAsString
}
