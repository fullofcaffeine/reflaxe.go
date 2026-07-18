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
		hx_obj_42 := map[string]any{}
		hx_obj_42["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_42["lineNumber"] = 0
		hx_obj_42["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_42["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_42
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_43 map[string]any) *string {
		hx_field_44 := hx_obj_43["className"]
		if hx_field_44 == nil {
			var hx_zero_45 *string
			return hx_zero_45
		}
		return hx_field_44.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_46 map[string]any) *string {
		hx_field_47 := hx_obj_46["methodName"]
		if hx_field_47 == nil {
			var hx_zero_48 *string
			return hx_zero_48
		}
		return hx_field_47.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_49 map[string]any) *string {
		hx_field_50 := hx_obj_49["fileName"]
		if hx_field_50 == nil {
			var hx_zero_51 *string
			return hx_zero_51
		}
		return hx_field_50.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_52 map[string]any) int {
		hx_field_53 := hx_obj_52["lineNumber"]
		if hx_field_53 == nil {
			var hx_zero_54 int
			return hx_zero_54
		}
		return hx_field_53.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
