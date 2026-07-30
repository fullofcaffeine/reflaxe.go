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
		hx_obj_314 := map[string]any{}
		hx_obj_314["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_314["lineNumber"] = 0
		hx_obj_314["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_314["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_314
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_315 map[string]any) *string {
		hx_field_316 := hx_obj_315["className"]
		if hx_field_316 == nil {
			var hx_zero_317 *string
			return hx_zero_317
		}
		return hx_field_316.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_318 map[string]any) *string {
		hx_field_319 := hx_obj_318["methodName"]
		if hx_field_319 == nil {
			var hx_zero_320 *string
			return hx_zero_320
		}
		return hx_field_319.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_321 map[string]any) *string {
		hx_field_322 := hx_obj_321["fileName"]
		if hx_field_322 == nil {
			var hx_zero_323 *string
			return hx_zero_323
		}
		return hx_field_322.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_324 map[string]any) int {
		hx_field_325 := hx_obj_324["lineNumber"]
		if hx_field_325 == nil {
			var hx_zero_326 int
			return hx_zero_326
		}
		return hx_field_325.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
