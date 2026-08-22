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
		hx_obj_1 := map[string]any{}
		hx_obj_1["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_1["lineNumber"] = 0
		hx_obj_1["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_1["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_1
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_2 map[string]any) *string {
		hx_field_3 := hx_obj_2["className"]
		if hx_field_3 == nil {
			var hx_zero_4 *string
			return hx_zero_4
		}
		return hx_field_3.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_5 map[string]any) *string {
		hx_field_6 := hx_obj_5["methodName"]
		if hx_field_6 == nil {
			var hx_zero_7 *string
			return hx_zero_7
		}
		return hx_field_6.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_8 map[string]any) *string {
		hx_field_9 := hx_obj_8["fileName"]
		if hx_field_9 == nil {
			var hx_zero_10 *string
			return hx_zero_10
		}
		return hx_field_9.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_11 map[string]any) int {
		hx_field_12 := hx_obj_11["lineNumber"]
		if hx_field_12 == nil {
			var hx_zero_13 int
			return hx_zero_13
		}
		return hx_field_12.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
