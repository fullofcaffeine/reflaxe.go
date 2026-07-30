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
		hx_obj_309 := map[string]any{}
		hx_obj_309["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_309["lineNumber"] = 0
		hx_obj_309["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_309["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_309
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_310 map[string]any) *string {
		hx_field_311 := hx_obj_310["className"]
		if hx_field_311 == nil {
			var hx_zero_312 *string
			return hx_zero_312
		}
		return hx_field_311.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_313 map[string]any) *string {
		hx_field_314 := hx_obj_313["methodName"]
		if hx_field_314 == nil {
			var hx_zero_315 *string
			return hx_zero_315
		}
		return hx_field_314.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_316 map[string]any) *string {
		hx_field_317 := hx_obj_316["fileName"]
		if hx_field_317 == nil {
			var hx_zero_318 *string
			return hx_zero_318
		}
		return hx_field_317.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_319 map[string]any) int {
		hx_field_320 := hx_obj_319["lineNumber"]
		if hx_field_320 == nil {
			var hx_zero_321 int
			return hx_zero_321
		}
		return hx_field_320.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
