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
		hx_obj_304 := map[string]any{}
		hx_obj_304["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_304["lineNumber"] = 0
		hx_obj_304["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_304["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_304
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_305 map[string]any) *string {
		hx_field_306 := hx_obj_305["className"]
		if hx_field_306 == nil {
			var hx_zero_307 *string
			return hx_zero_307
		}
		return hx_field_306.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_308 map[string]any) *string {
		hx_field_309 := hx_obj_308["methodName"]
		if hx_field_309 == nil {
			var hx_zero_310 *string
			return hx_zero_310
		}
		return hx_field_309.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_311 map[string]any) *string {
		hx_field_312 := hx_obj_311["fileName"]
		if hx_field_312 == nil {
			var hx_zero_313 *string
			return hx_zero_313
		}
		return hx_field_312.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_314 map[string]any) int {
		hx_field_315 := hx_obj_314["lineNumber"]
		if hx_field_315 == nil {
			var hx_zero_316 int
			return hx_zero_316
		}
		return hx_field_315.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
