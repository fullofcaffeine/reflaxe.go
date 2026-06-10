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
		hx_obj_25 := map[string]any{}
		hx_obj_25["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_25["lineNumber"] = 0
		hx_obj_25["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_25["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_25
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_26 map[string]any) *string {
		hx_field_27 := hx_obj_26["className"]
		if hx_field_27 == nil {
			var hx_zero_28 *string
			return hx_zero_28
		}
		return hx_field_27.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_29 map[string]any) *string {
		hx_field_30 := hx_obj_29["methodName"]
		if hx_field_30 == nil {
			var hx_zero_31 *string
			return hx_zero_31
		}
		return hx_field_30.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_32 map[string]any) *string {
		hx_field_33 := hx_obj_32["fileName"]
		if hx_field_33 == nil {
			var hx_zero_34 *string
			return hx_zero_34
		}
		return hx_field_33.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_35 map[string]any) int {
		hx_field_36 := hx_obj_35["lineNumber"]
		if hx_field_36 == nil {
			var hx_zero_37 int
			return hx_zero_37
		}
		return hx_field_36.(int)
	}(self.posInfos))
}
