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
		hx_obj_328 := map[string]any{}
		hx_obj_328["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_328["lineNumber"] = 0
		hx_obj_328["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_328["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_328
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_329 map[string]any) *string {
		hx_field_330 := hx_obj_329["className"]
		if hx_field_330 == nil {
			var hx_zero_331 *string
			return hx_zero_331
		}
		return hx_field_330.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_332 map[string]any) *string {
		hx_field_333 := hx_obj_332["methodName"]
		if hx_field_333 == nil {
			var hx_zero_334 *string
			return hx_zero_334
		}
		return hx_field_333.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_335 map[string]any) *string {
		hx_field_336 := hx_obj_335["fileName"]
		if hx_field_336 == nil {
			var hx_zero_337 *string
			return hx_zero_337
		}
		return hx_field_336.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_338 map[string]any) int {
		hx_field_339 := hx_obj_338["lineNumber"]
		if hx_field_339 == nil {
			var hx_zero_340 int
			return hx_zero_340
		}
		return hx_field_339.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
