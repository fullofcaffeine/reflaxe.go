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
		hx_obj_310 := map[string]any{}
		hx_obj_310["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_310["lineNumber"] = 0
		hx_obj_310["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_310["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_310
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_311 map[string]any) *string {
		hx_field_312 := hx_obj_311["className"]
		if hx_field_312 == nil {
			var hx_zero_313 *string
			return hx_zero_313
		}
		return hx_field_312.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_314 map[string]any) *string {
		hx_field_315 := hx_obj_314["methodName"]
		if hx_field_315 == nil {
			var hx_zero_316 *string
			return hx_zero_316
		}
		return hx_field_315.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_317 map[string]any) *string {
		hx_field_318 := hx_obj_317["fileName"]
		if hx_field_318 == nil {
			var hx_zero_319 *string
			return hx_zero_319
		}
		return hx_field_318.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_320 map[string]any) int {
		hx_field_321 := hx_obj_320["lineNumber"]
		if hx_field_321 == nil {
			var hx_zero_322 int
			return hx_zero_322
		}
		return hx_field_321.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
