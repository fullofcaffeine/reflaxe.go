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
		hx_obj_116 := map[string]any{}
		hx_obj_116["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_116["lineNumber"] = 0
		hx_obj_116["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_116["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_116
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_117 map[string]any) *string {
		hx_field_118 := hx_obj_117["className"]
		if hx_field_118 == nil {
			var hx_zero_119 *string
			return hx_zero_119
		}
		return hx_field_118.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_120 map[string]any) *string {
		hx_field_121 := hx_obj_120["methodName"]
		if hx_field_121 == nil {
			var hx_zero_122 *string
			return hx_zero_122
		}
		return hx_field_121.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_123 map[string]any) *string {
		hx_field_124 := hx_obj_123["fileName"]
		if hx_field_124 == nil {
			var hx_zero_125 *string
			return hx_zero_125
		}
		return hx_field_124.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_126 map[string]any) int {
		hx_field_127 := hx_obj_126["lineNumber"]
		if hx_field_127 == nil {
			var hx_zero_128 int
			return hx_zero_128
		}
		return hx_field_127.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
