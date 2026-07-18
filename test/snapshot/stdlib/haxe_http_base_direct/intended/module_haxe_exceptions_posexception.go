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
		hx_obj_40 := map[string]any{}
		hx_obj_40["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_40["lineNumber"] = 0
		hx_obj_40["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_40["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_40
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_41 map[string]any) *string {
		hx_field_42 := hx_obj_41["className"]
		if hx_field_42 == nil {
			var hx_zero_43 *string
			return hx_zero_43
		}
		return hx_field_42.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_44 map[string]any) *string {
		hx_field_45 := hx_obj_44["methodName"]
		if hx_field_45 == nil {
			var hx_zero_46 *string
			return hx_zero_46
		}
		return hx_field_45.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_47 map[string]any) *string {
		hx_field_48 := hx_obj_47["fileName"]
		if hx_field_48 == nil {
			var hx_zero_49 *string
			return hx_zero_49
		}
		return hx_field_48.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_50 map[string]any) int {
		hx_field_51 := hx_obj_50["lineNumber"]
		if hx_field_51 == nil {
			var hx_zero_52 int
			return hx_zero_52
		}
		return hx_field_51.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
