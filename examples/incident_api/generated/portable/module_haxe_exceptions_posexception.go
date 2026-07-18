package main

import "examples_incident_api_portable/hxrt"

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
		hx_obj_205 := map[string]any{}
		hx_obj_205["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_205["lineNumber"] = 0
		hx_obj_205["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_205["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_205
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_206 map[string]any) *string {
		hx_field_207 := hx_obj_206["className"]
		if hx_field_207 == nil {
			var hx_zero_208 *string
			return hx_zero_208
		}
		return hx_field_207.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_209 map[string]any) *string {
		hx_field_210 := hx_obj_209["methodName"]
		if hx_field_210 == nil {
			var hx_zero_211 *string
			return hx_zero_211
		}
		return hx_field_210.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_212 map[string]any) *string {
		hx_field_213 := hx_obj_212["fileName"]
		if hx_field_213 == nil {
			var hx_zero_214 *string
			return hx_zero_214
		}
		return hx_field_213.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_215 map[string]any) int {
		hx_field_216 := hx_obj_215["lineNumber"]
		if hx_field_216 == nil {
			var hx_zero_217 int
			return hx_zero_217
		}
		return hx_field_216.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
