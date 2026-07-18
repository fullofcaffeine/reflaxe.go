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
		hx_obj_200 := map[string]any{}
		hx_obj_200["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_200["lineNumber"] = 0
		hx_obj_200["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_200["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_200
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_201 map[string]any) *string {
		hx_field_202 := hx_obj_201["className"]
		if hx_field_202 == nil {
			var hx_zero_203 *string
			return hx_zero_203
		}
		return hx_field_202.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_204 map[string]any) *string {
		hx_field_205 := hx_obj_204["methodName"]
		if hx_field_205 == nil {
			var hx_zero_206 *string
			return hx_zero_206
		}
		return hx_field_205.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_207 map[string]any) *string {
		hx_field_208 := hx_obj_207["fileName"]
		if hx_field_208 == nil {
			var hx_zero_209 *string
			return hx_zero_209
		}
		return hx_field_208.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_210 map[string]any) int {
		hx_field_211 := hx_obj_210["lineNumber"]
		if hx_field_211 == nil {
			var hx_zero_212 int
			return hx_zero_212
		}
		return hx_field_211.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
