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
		hx_obj_330 := map[string]any{}
		hx_obj_330["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_330["lineNumber"] = 0
		hx_obj_330["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_330["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_330
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_331 map[string]any) *string {
		hx_field_332 := hx_obj_331["className"]
		if hx_field_332 == nil {
			var hx_zero_333 *string
			return hx_zero_333
		}
		return hx_field_332.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_334 map[string]any) *string {
		hx_field_335 := hx_obj_334["methodName"]
		if hx_field_335 == nil {
			var hx_zero_336 *string
			return hx_zero_336
		}
		return hx_field_335.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_337 map[string]any) *string {
		hx_field_338 := hx_obj_337["fileName"]
		if hx_field_338 == nil {
			var hx_zero_339 *string
			return hx_zero_339
		}
		return hx_field_338.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_340 map[string]any) int {
		hx_field_341 := hx_obj_340["lineNumber"]
		if hx_field_341 == nil {
			var hx_zero_342 int
			return hx_zero_342
		}
		return hx_field_341.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
