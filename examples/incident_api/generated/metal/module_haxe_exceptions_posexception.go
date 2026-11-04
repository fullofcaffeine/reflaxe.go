package main

import "examples_incident_api_metal/hxrt"

type I_haxe__exceptions__PosException interface {
	get_message() *string
	get_stack() any
	get_previous() *haxe__Exception
	get_native() any
	unwrap() any
	toString() *string
	details() *string
}

type haxe__exceptions__PosException struct {
	*haxe__Exception
	__hx_this I_haxe__exceptions__PosException
	posInfos  hxrt.Anonymous
}

func (self *haxe__exceptions__PosException) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe__exceptions__PosException(message *string, previous *haxe__Exception, pos hxrt.Anonymous) *haxe__exceptions__PosException {
	return hx_init_haxe__exceptions__PosException(hx_alloc_haxe__exceptions__PosException(nil), message, previous, pos)
}

func hx_alloc_haxe__exceptions__PosException(owner any) *haxe__exceptions__PosException {
	self := &haxe__exceptions__PosException{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_haxe__exceptions__PosException)
	}
	self.haxe__Exception = hx_alloc_haxe__Exception(self.__hx_this)
	return self
}

func hx_init_haxe__exceptions__PosException(self *haxe__exceptions__PosException, message *string, previous *haxe__Exception, pos hxrt.Anonymous) *haxe__exceptions__PosException {
	hx_init_haxe__Exception(self.haxe__Exception, message, previous, nil)
	if pos == nil {
		hx_obj_1 := map[string]any{}
		hx_obj_1["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_1["lineNumber"] = 0
		hx_obj_1["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_1["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hxrt.AnonymousFromMap(hx_obj_1)
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(self.__hx_this.get_message(), hxrt.StringFromLiteral(" in ")), func(hx_value_2 any) *string {
		if hx_value_2 == nil {
			var hx_zero_3 *string
			return hx_zero_3
		}
		return hx_value_2.(*string)
	}(hxrt.AnonymousGet(self.posInfos, "className"))), hxrt.StringFromLiteral(".")), func(hx_value_4 any) *string {
		if hx_value_4 == nil {
			var hx_zero_5 *string
			return hx_zero_5
		}
		return hx_value_4.(*string)
	}(hxrt.AnonymousGet(self.posInfos, "methodName"))), hxrt.StringFromLiteral(" at ")), func(hx_value_6 any) *string {
		if hx_value_6 == nil {
			var hx_zero_7 *string
			return hx_zero_7
		}
		return hx_value_6.(*string)
	}(hxrt.AnonymousGet(self.posInfos, "fileName"))), hxrt.StringFromLiteral(":")), func(hx_value_8 any) int {
		if hx_value_8 == nil {
			var hx_zero_9 int
			return hx_zero_9
		}
		return hx_value_8.(int)
	}(hxrt.AnonymousGet(self.posInfos, "lineNumber")))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
