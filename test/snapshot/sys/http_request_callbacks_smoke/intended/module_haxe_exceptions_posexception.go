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
		hx_obj_308 := map[string]any{}
		hx_obj_308["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_308["lineNumber"] = 0
		hx_obj_308["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_308["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_308
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_309 map[string]any) *string {
		hx_field_310 := hx_obj_309["className"]
		if hx_field_310 == nil {
			var hx_zero_311 *string
			return hx_zero_311
		}
		return hx_field_310.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_312 map[string]any) *string {
		hx_field_313 := hx_obj_312["methodName"]
		if hx_field_313 == nil {
			var hx_zero_314 *string
			return hx_zero_314
		}
		return hx_field_313.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_315 map[string]any) *string {
		hx_field_316 := hx_obj_315["fileName"]
		if hx_field_316 == nil {
			var hx_zero_317 *string
			return hx_zero_317
		}
		return hx_field_316.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_318 map[string]any) int {
		hx_field_319 := hx_obj_318["lineNumber"]
		if hx_field_319 == nil {
			var hx_zero_320 int
			return hx_zero_320
		}
		return hx_field_319.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
