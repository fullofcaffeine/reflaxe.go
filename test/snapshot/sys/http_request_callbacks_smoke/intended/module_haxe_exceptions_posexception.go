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
		hx_obj_41 := map[string]any{}
		hx_obj_41["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_41["lineNumber"] = 0
		hx_obj_41["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_41["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_41
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_42 map[string]any) *string {
		hx_field_43 := hx_obj_42["className"]
		if hx_field_43 == nil {
			var hx_zero_44 *string
			return hx_zero_44
		}
		return hx_field_43.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_45 map[string]any) *string {
		hx_field_46 := hx_obj_45["methodName"]
		if hx_field_46 == nil {
			var hx_zero_47 *string
			return hx_zero_47
		}
		return hx_field_46.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_48 map[string]any) *string {
		hx_field_49 := hx_obj_48["fileName"]
		if hx_field_49 == nil {
			var hx_zero_50 *string
			return hx_zero_50
		}
		return hx_field_49.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_51 map[string]any) int {
		hx_field_52 := hx_obj_51["lineNumber"]
		if hx_field_52 == nil {
			var hx_zero_53 int
			return hx_zero_53
		}
		return hx_field_52.(int)
	}(self.posInfos))
}
