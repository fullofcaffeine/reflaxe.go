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
		hx_obj_338 := map[string]any{}
		hx_obj_338["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_338["lineNumber"] = 0
		hx_obj_338["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_338["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_338
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_339 map[string]any) *string {
		hx_field_340 := hx_obj_339["className"]
		if hx_field_340 == nil {
			var hx_zero_341 *string
			return hx_zero_341
		}
		return hx_field_340.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_342 map[string]any) *string {
		hx_field_343 := hx_obj_342["methodName"]
		if hx_field_343 == nil {
			var hx_zero_344 *string
			return hx_zero_344
		}
		return hx_field_343.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_345 map[string]any) *string {
		hx_field_346 := hx_obj_345["fileName"]
		if hx_field_346 == nil {
			var hx_zero_347 *string
			return hx_zero_347
		}
		return hx_field_346.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_348 map[string]any) int {
		hx_field_349 := hx_obj_348["lineNumber"]
		if hx_field_349 == nil {
			var hx_zero_350 int
			return hx_zero_350
		}
		return hx_field_349.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
