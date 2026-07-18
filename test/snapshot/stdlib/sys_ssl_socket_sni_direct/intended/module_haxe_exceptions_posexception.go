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
		hx_obj_132 := map[string]any{}
		hx_obj_132["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_132["lineNumber"] = 0
		hx_obj_132["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_132["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_132
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_133 map[string]any) *string {
		hx_field_134 := hx_obj_133["className"]
		if hx_field_134 == nil {
			var hx_zero_135 *string
			return hx_zero_135
		}
		return hx_field_134.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_136 map[string]any) *string {
		hx_field_137 := hx_obj_136["methodName"]
		if hx_field_137 == nil {
			var hx_zero_138 *string
			return hx_zero_138
		}
		return hx_field_137.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_139 map[string]any) *string {
		hx_field_140 := hx_obj_139["fileName"]
		if hx_field_140 == nil {
			var hx_zero_141 *string
			return hx_zero_141
		}
		return hx_field_140.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_142 map[string]any) int {
		hx_field_143 := hx_obj_142["lineNumber"]
		if hx_field_143 == nil {
			var hx_zero_144 int
			return hx_zero_144
		}
		return hx_field_143.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
