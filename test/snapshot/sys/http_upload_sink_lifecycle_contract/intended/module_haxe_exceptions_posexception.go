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
		hx_obj_457 := map[string]any{}
		hx_obj_457["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_457["lineNumber"] = 0
		hx_obj_457["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_457["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_457
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_458 map[string]any) *string {
		hx_field_459 := hx_obj_458["className"]
		if hx_field_459 == nil {
			var hx_zero_460 *string
			return hx_zero_460
		}
		return hx_field_459.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_461 map[string]any) *string {
		hx_field_462 := hx_obj_461["methodName"]
		if hx_field_462 == nil {
			var hx_zero_463 *string
			return hx_zero_463
		}
		return hx_field_462.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_464 map[string]any) *string {
		hx_field_465 := hx_obj_464["fileName"]
		if hx_field_465 == nil {
			var hx_zero_466 *string
			return hx_zero_466
		}
		return hx_field_465.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_467 map[string]any) int {
		hx_field_468 := hx_obj_467["lineNumber"]
		if hx_field_468 == nil {
			var hx_zero_469 int
			return hx_zero_469
		}
		return hx_field_468.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
