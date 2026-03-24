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
		hx_obj_29 := map[string]any{}
		hx_obj_29["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_29["lineNumber"] = 0
		hx_obj_29["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_29["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_29
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_30 map[string]any) *string {
		hx_field_31 := hx_obj_30["className"]
		if hx_field_31 == nil {
			var hx_zero_32 *string
			return hx_zero_32
		}
		return hx_field_31.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_33 map[string]any) *string {
		hx_field_34 := hx_obj_33["methodName"]
		if hx_field_34 == nil {
			var hx_zero_35 *string
			return hx_zero_35
		}
		return hx_field_34.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_36 map[string]any) *string {
		hx_field_37 := hx_obj_36["fileName"]
		if hx_field_37 == nil {
			var hx_zero_38 *string
			return hx_zero_38
		}
		return hx_field_37.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_39 map[string]any) int {
		hx_field_40 := hx_obj_39["lineNumber"]
		if hx_field_40 == nil {
			var hx_zero_41 int
			return hx_zero_41
		}
		return hx_field_40.(int)
	}(self.posInfos))
}
