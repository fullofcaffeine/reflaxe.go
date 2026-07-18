package main

import "examples_tui_todo_portable/hxrt"

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
		hx_obj_143 := map[string]any{}
		hx_obj_143["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_143["lineNumber"] = 0
		hx_obj_143["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_143["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_143
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_144 map[string]any) *string {
		hx_field_145 := hx_obj_144["className"]
		if hx_field_145 == nil {
			var hx_zero_146 *string
			return hx_zero_146
		}
		return hx_field_145.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_147 map[string]any) *string {
		hx_field_148 := hx_obj_147["methodName"]
		if hx_field_148 == nil {
			var hx_zero_149 *string
			return hx_zero_149
		}
		return hx_field_148.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_150 map[string]any) *string {
		hx_field_151 := hx_obj_150["fileName"]
		if hx_field_151 == nil {
			var hx_zero_152 *string
			return hx_zero_152
		}
		return hx_field_151.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_153 map[string]any) int {
		hx_field_154 := hx_obj_153["lineNumber"]
		if hx_field_154 == nil {
			var hx_zero_155 int
			return hx_zero_155
		}
		return hx_field_154.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
