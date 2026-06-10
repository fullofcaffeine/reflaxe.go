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
		hx_obj_15 := map[string]any{}
		hx_obj_15["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_15["lineNumber"] = 0
		hx_obj_15["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_15["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_15
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_16 map[string]any) *string {
		hx_field_17 := hx_obj_16["className"]
		if hx_field_17 == nil {
			var hx_zero_18 *string
			return hx_zero_18
		}
		return hx_field_17.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_19 map[string]any) *string {
		hx_field_20 := hx_obj_19["methodName"]
		if hx_field_20 == nil {
			var hx_zero_21 *string
			return hx_zero_21
		}
		return hx_field_20.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_22 map[string]any) *string {
		hx_field_23 := hx_obj_22["fileName"]
		if hx_field_23 == nil {
			var hx_zero_24 *string
			return hx_zero_24
		}
		return hx_field_23.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_25 map[string]any) int {
		hx_field_26 := hx_obj_25["lineNumber"]
		if hx_field_26 == nil {
			var hx_zero_27 int
			return hx_zero_27
		}
		return hx_field_26.(int)
	}(self.posInfos))
}
