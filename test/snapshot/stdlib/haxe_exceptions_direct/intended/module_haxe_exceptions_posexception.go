package main

import "snapshot/hxrt"

type I_haxe__exceptions__PosException interface {
	toString() *string
}

type haxe__exceptions__PosException struct {
	__hx_this      I_haxe__exceptions__PosException
	posInfos       map[string]any
	__hx_exception *hxrt.ExceptionValue
}

func New_haxe__exceptions__PosException(message *string, previous *hxrt.ExceptionValue, pos map[string]any) *haxe__exceptions__PosException {
	self := &haxe__exceptions__PosException{}
	self.__hx_exception = hxrt.BindException(self, message, previous, nil)
	self.__hx_this = self
	if pos == nil {
		hx_obj_2 := map[string]any{}
		hx_obj_2["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_2["lineNumber"] = 0
		hx_obj_2["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_2["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_2
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_3 map[string]any) *string {
		hx_field_4 := hx_obj_3["className"]
		if hx_field_4 == nil {
			var hx_zero_5 *string
			return hx_zero_5
		}
		return hx_field_4.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_6 map[string]any) *string {
		hx_field_7 := hx_obj_6["methodName"]
		if hx_field_7 == nil {
			var hx_zero_8 *string
			return hx_zero_8
		}
		return hx_field_7.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_9 map[string]any) *string {
		hx_field_10 := hx_obj_9["fileName"]
		if hx_field_10 == nil {
			var hx_zero_11 *string
			return hx_zero_11
		}
		return hx_field_10.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_12 map[string]any) int {
		hx_field_13 := hx_obj_12["lineNumber"]
		if hx_field_13 == nil {
			var hx_zero_14 int
			return hx_zero_14
		}
		return hx_field_13.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
