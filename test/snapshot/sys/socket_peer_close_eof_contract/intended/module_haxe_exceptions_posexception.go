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
		hx_obj_109 := map[string]any{}
		hx_obj_109["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_109["lineNumber"] = 0
		hx_obj_109["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_109["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_109
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_110 map[string]any) *string {
		hx_field_111 := hx_obj_110["className"]
		if hx_field_111 == nil {
			var hx_zero_112 *string
			return hx_zero_112
		}
		return hx_field_111.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_113 map[string]any) *string {
		hx_field_114 := hx_obj_113["methodName"]
		if hx_field_114 == nil {
			var hx_zero_115 *string
			return hx_zero_115
		}
		return hx_field_114.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_116 map[string]any) *string {
		hx_field_117 := hx_obj_116["fileName"]
		if hx_field_117 == nil {
			var hx_zero_118 *string
			return hx_zero_118
		}
		return hx_field_117.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_119 map[string]any) int {
		hx_field_120 := hx_obj_119["lineNumber"]
		if hx_field_120 == nil {
			var hx_zero_121 int
			return hx_zero_121
		}
		return hx_field_120.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
