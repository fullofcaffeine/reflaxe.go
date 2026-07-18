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
		hx_obj_50 := map[string]any{}
		hx_obj_50["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_50["lineNumber"] = 0
		hx_obj_50["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_50["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_50
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_51 map[string]any) *string {
		hx_field_52 := hx_obj_51["className"]
		if hx_field_52 == nil {
			var hx_zero_53 *string
			return hx_zero_53
		}
		return hx_field_52.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_54 map[string]any) *string {
		hx_field_55 := hx_obj_54["methodName"]
		if hx_field_55 == nil {
			var hx_zero_56 *string
			return hx_zero_56
		}
		return hx_field_55.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_57 map[string]any) *string {
		hx_field_58 := hx_obj_57["fileName"]
		if hx_field_58 == nil {
			var hx_zero_59 *string
			return hx_zero_59
		}
		return hx_field_58.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_60 map[string]any) int {
		hx_field_61 := hx_obj_60["lineNumber"]
		if hx_field_61 == nil {
			var hx_zero_62 int
			return hx_zero_62
		}
		return hx_field_61.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
