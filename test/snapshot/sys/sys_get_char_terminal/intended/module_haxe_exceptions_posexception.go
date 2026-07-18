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
		hx_obj_111 := map[string]any{}
		hx_obj_111["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_111["lineNumber"] = 0
		hx_obj_111["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_111["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_111
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_112 map[string]any) *string {
		hx_field_113 := hx_obj_112["className"]
		if hx_field_113 == nil {
			var hx_zero_114 *string
			return hx_zero_114
		}
		return hx_field_113.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_115 map[string]any) *string {
		hx_field_116 := hx_obj_115["methodName"]
		if hx_field_116 == nil {
			var hx_zero_117 *string
			return hx_zero_117
		}
		return hx_field_116.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_118 map[string]any) *string {
		hx_field_119 := hx_obj_118["fileName"]
		if hx_field_119 == nil {
			var hx_zero_120 *string
			return hx_zero_120
		}
		return hx_field_119.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_121 map[string]any) int {
		hx_field_122 := hx_obj_121["lineNumber"]
		if hx_field_122 == nil {
			var hx_zero_123 int
			return hx_zero_123
		}
		return hx_field_122.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
