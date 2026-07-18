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
		hx_obj_59 := map[string]any{}
		hx_obj_59["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_59["lineNumber"] = 0
		hx_obj_59["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_59["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_59
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_60 map[string]any) *string {
		hx_field_61 := hx_obj_60["className"]
		if hx_field_61 == nil {
			var hx_zero_62 *string
			return hx_zero_62
		}
		return hx_field_61.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_63 map[string]any) *string {
		hx_field_64 := hx_obj_63["methodName"]
		if hx_field_64 == nil {
			var hx_zero_65 *string
			return hx_zero_65
		}
		return hx_field_64.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_66 map[string]any) *string {
		hx_field_67 := hx_obj_66["fileName"]
		if hx_field_67 == nil {
			var hx_zero_68 *string
			return hx_zero_68
		}
		return hx_field_67.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_69 map[string]any) int {
		hx_field_70 := hx_obj_69["lineNumber"]
		if hx_field_70 == nil {
			var hx_zero_71 int
			return hx_zero_71
		}
		return hx_field_70.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
