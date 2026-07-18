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
		hx_obj_34 := map[string]any{}
		hx_obj_34["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_34["lineNumber"] = 0
		hx_obj_34["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_34["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_34
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_35 map[string]any) *string {
		hx_field_36 := hx_obj_35["className"]
		if hx_field_36 == nil {
			var hx_zero_37 *string
			return hx_zero_37
		}
		return hx_field_36.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_38 map[string]any) *string {
		hx_field_39 := hx_obj_38["methodName"]
		if hx_field_39 == nil {
			var hx_zero_40 *string
			return hx_zero_40
		}
		return hx_field_39.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_41 map[string]any) *string {
		hx_field_42 := hx_obj_41["fileName"]
		if hx_field_42 == nil {
			var hx_zero_43 *string
			return hx_zero_43
		}
		return hx_field_42.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_44 map[string]any) int {
		hx_field_45 := hx_obj_44["lineNumber"]
		if hx_field_45 == nil {
			var hx_zero_46 int
			return hx_zero_46
		}
		return hx_field_45.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
