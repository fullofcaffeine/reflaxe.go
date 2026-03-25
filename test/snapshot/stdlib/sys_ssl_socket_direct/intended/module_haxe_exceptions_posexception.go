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
		hx_obj_23 := map[string]any{}
		hx_obj_23["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_23["lineNumber"] = 0
		hx_obj_23["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_23["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_23
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_24 map[string]any) *string {
		hx_field_25 := hx_obj_24["className"]
		if hx_field_25 == nil {
			var hx_zero_26 *string
			return hx_zero_26
		}
		return hx_field_25.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_27 map[string]any) *string {
		hx_field_28 := hx_obj_27["methodName"]
		if hx_field_28 == nil {
			var hx_zero_29 *string
			return hx_zero_29
		}
		return hx_field_28.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_30 map[string]any) *string {
		hx_field_31 := hx_obj_30["fileName"]
		if hx_field_31 == nil {
			var hx_zero_32 *string
			return hx_zero_32
		}
		return hx_field_31.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_33 map[string]any) int {
		hx_field_34 := hx_obj_33["lineNumber"]
		if hx_field_34 == nil {
			var hx_zero_35 int
			return hx_zero_35
		}
		return hx_field_34.(int)
	}(self.posInfos))
}
