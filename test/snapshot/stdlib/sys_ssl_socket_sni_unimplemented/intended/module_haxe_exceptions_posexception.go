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
		hx_obj_13 := map[string]any{}
		hx_obj_13["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_13["lineNumber"] = 0
		hx_obj_13["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_13["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_13
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_14 map[string]any) *string {
		hx_field_15 := hx_obj_14["className"]
		if hx_field_15 == nil {
			var hx_zero_16 *string
			return hx_zero_16
		}
		return hx_field_15.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_17 map[string]any) *string {
		hx_field_18 := hx_obj_17["methodName"]
		if hx_field_18 == nil {
			var hx_zero_19 *string
			return hx_zero_19
		}
		return hx_field_18.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_20 map[string]any) *string {
		hx_field_21 := hx_obj_20["fileName"]
		if hx_field_21 == nil {
			var hx_zero_22 *string
			return hx_zero_22
		}
		return hx_field_21.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_23 map[string]any) int {
		hx_field_24 := hx_obj_23["lineNumber"]
		if hx_field_24 == nil {
			var hx_zero_25 int
			return hx_zero_25
		}
		return hx_field_24.(int)
	}(self.posInfos))
}
