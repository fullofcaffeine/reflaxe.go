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
		hx_obj_337 := map[string]any{}
		hx_obj_337["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_337["lineNumber"] = 0
		hx_obj_337["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_337["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_337
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_338 map[string]any) *string {
		hx_field_339 := hx_obj_338["className"]
		if hx_field_339 == nil {
			var hx_zero_340 *string
			return hx_zero_340
		}
		return hx_field_339.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_341 map[string]any) *string {
		hx_field_342 := hx_obj_341["methodName"]
		if hx_field_342 == nil {
			var hx_zero_343 *string
			return hx_zero_343
		}
		return hx_field_342.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_344 map[string]any) *string {
		hx_field_345 := hx_obj_344["fileName"]
		if hx_field_345 == nil {
			var hx_zero_346 *string
			return hx_zero_346
		}
		return hx_field_345.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_347 map[string]any) int {
		hx_field_348 := hx_obj_347["lineNumber"]
		if hx_field_348 == nil {
			var hx_zero_349 int
			return hx_zero_349
		}
		return hx_field_348.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
