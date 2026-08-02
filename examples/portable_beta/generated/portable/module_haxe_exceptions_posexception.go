package main

import "examples_portable_beta/hxrt"

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
		hx_obj_177 := map[string]any{}
		hx_obj_177["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_177["lineNumber"] = 0
		hx_obj_177["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_177["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_177
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_178 map[string]any) *string {
		hx_field_179 := hx_obj_178["className"]
		if hx_field_179 == nil {
			var hx_zero_180 *string
			return hx_zero_180
		}
		return hx_field_179.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_181 map[string]any) *string {
		hx_field_182 := hx_obj_181["methodName"]
		if hx_field_182 == nil {
			var hx_zero_183 *string
			return hx_zero_183
		}
		return hx_field_182.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_184 map[string]any) *string {
		hx_field_185 := hx_obj_184["fileName"]
		if hx_field_185 == nil {
			var hx_zero_186 *string
			return hx_zero_186
		}
		return hx_field_185.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_187 map[string]any) int {
		hx_field_188 := hx_obj_187["lineNumber"]
		if hx_field_188 == nil {
			var hx_zero_189 int
			return hx_zero_189
		}
		return hx_field_188.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
