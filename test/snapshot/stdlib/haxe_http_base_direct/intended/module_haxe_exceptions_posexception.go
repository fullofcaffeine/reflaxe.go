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
		hx_obj_30 := map[string]any{}
		hx_obj_30["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_30["lineNumber"] = 0
		hx_obj_30["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_30["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_30
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_31 map[string]any) *string {
		hx_field_32 := hx_obj_31["className"]
		if hx_field_32 == nil {
			var hx_zero_33 *string
			return hx_zero_33
		}
		return hx_field_32.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_34 map[string]any) *string {
		hx_field_35 := hx_obj_34["methodName"]
		if hx_field_35 == nil {
			var hx_zero_36 *string
			return hx_zero_36
		}
		return hx_field_35.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_37 map[string]any) *string {
		hx_field_38 := hx_obj_37["fileName"]
		if hx_field_38 == nil {
			var hx_zero_39 *string
			return hx_zero_39
		}
		return hx_field_38.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_40 map[string]any) int {
		hx_field_41 := hx_obj_40["lineNumber"]
		if hx_field_41 == nil {
			var hx_zero_42 int
			return hx_zero_42
		}
		return hx_field_41.(int)
	}(self.posInfos))
}
