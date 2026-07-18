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
		hx_obj_51 := map[string]any{}
		hx_obj_51["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_51["lineNumber"] = 0
		hx_obj_51["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_51["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_51
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_52 map[string]any) *string {
		hx_field_53 := hx_obj_52["className"]
		if hx_field_53 == nil {
			var hx_zero_54 *string
			return hx_zero_54
		}
		return hx_field_53.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_55 map[string]any) *string {
		hx_field_56 := hx_obj_55["methodName"]
		if hx_field_56 == nil {
			var hx_zero_57 *string
			return hx_zero_57
		}
		return hx_field_56.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_58 map[string]any) *string {
		hx_field_59 := hx_obj_58["fileName"]
		if hx_field_59 == nil {
			var hx_zero_60 *string
			return hx_zero_60
		}
		return hx_field_59.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_61 map[string]any) int {
		hx_field_62 := hx_obj_61["lineNumber"]
		if hx_field_62 == nil {
			var hx_zero_63 int
			return hx_zero_63
		}
		return hx_field_62.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
