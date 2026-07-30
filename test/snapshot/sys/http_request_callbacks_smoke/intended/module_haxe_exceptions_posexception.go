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
		hx_obj_323 := map[string]any{}
		hx_obj_323["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_323["lineNumber"] = 0
		hx_obj_323["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_323["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_323
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_324 map[string]any) *string {
		hx_field_325 := hx_obj_324["className"]
		if hx_field_325 == nil {
			var hx_zero_326 *string
			return hx_zero_326
		}
		return hx_field_325.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_327 map[string]any) *string {
		hx_field_328 := hx_obj_327["methodName"]
		if hx_field_328 == nil {
			var hx_zero_329 *string
			return hx_zero_329
		}
		return hx_field_328.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_330 map[string]any) *string {
		hx_field_331 := hx_obj_330["fileName"]
		if hx_field_331 == nil {
			var hx_zero_332 *string
			return hx_zero_332
		}
		return hx_field_331.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_333 map[string]any) int {
		hx_field_334 := hx_obj_333["lineNumber"]
		if hx_field_334 == nil {
			var hx_zero_335 int
			return hx_zero_335
		}
		return hx_field_334.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
