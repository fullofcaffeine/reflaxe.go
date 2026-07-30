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
		hx_obj_133 := map[string]any{}
		hx_obj_133["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_133["lineNumber"] = 0
		hx_obj_133["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_133["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_133
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_134 map[string]any) *string {
		hx_field_135 := hx_obj_134["className"]
		if hx_field_135 == nil {
			var hx_zero_136 *string
			return hx_zero_136
		}
		return hx_field_135.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_137 map[string]any) *string {
		hx_field_138 := hx_obj_137["methodName"]
		if hx_field_138 == nil {
			var hx_zero_139 *string
			return hx_zero_139
		}
		return hx_field_138.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_140 map[string]any) *string {
		hx_field_141 := hx_obj_140["fileName"]
		if hx_field_141 == nil {
			var hx_zero_142 *string
			return hx_zero_142
		}
		return hx_field_141.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_143 map[string]any) int {
		hx_field_144 := hx_obj_143["lineNumber"]
		if hx_field_144 == nil {
			var hx_zero_145 int
			return hx_zero_145
		}
		return hx_field_144.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
