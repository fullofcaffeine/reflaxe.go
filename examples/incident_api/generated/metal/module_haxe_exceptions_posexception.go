package main

import "examples_incident_api_metal/hxrt"

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
		hx_obj_217 := map[string]any{}
		hx_obj_217["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_217["lineNumber"] = 0
		hx_obj_217["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_217["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_217
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_218 map[string]any) *string {
		hx_field_219 := hx_obj_218["className"]
		if hx_field_219 == nil {
			var hx_zero_220 *string
			return hx_zero_220
		}
		return hx_field_219.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_221 map[string]any) *string {
		hx_field_222 := hx_obj_221["methodName"]
		if hx_field_222 == nil {
			var hx_zero_223 *string
			return hx_zero_223
		}
		return hx_field_222.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_224 map[string]any) *string {
		hx_field_225 := hx_obj_224["fileName"]
		if hx_field_225 == nil {
			var hx_zero_226 *string
			return hx_zero_226
		}
		return hx_field_225.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_227 map[string]any) int {
		hx_field_228 := hx_obj_227["lineNumber"]
		if hx_field_228 == nil {
			var hx_zero_229 int
			return hx_zero_229
		}
		return hx_field_228.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
