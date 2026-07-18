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
		hx_obj_35 := map[string]any{}
		hx_obj_35["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_35["lineNumber"] = 0
		hx_obj_35["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_35["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_35
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_36 map[string]any) *string {
		hx_field_37 := hx_obj_36["className"]
		if hx_field_37 == nil {
			var hx_zero_38 *string
			return hx_zero_38
		}
		return hx_field_37.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_39 map[string]any) *string {
		hx_field_40 := hx_obj_39["methodName"]
		if hx_field_40 == nil {
			var hx_zero_41 *string
			return hx_zero_41
		}
		return hx_field_40.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_42 map[string]any) *string {
		hx_field_43 := hx_obj_42["fileName"]
		if hx_field_43 == nil {
			var hx_zero_44 *string
			return hx_zero_44
		}
		return hx_field_43.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_45 map[string]any) int {
		hx_field_46 := hx_obj_45["lineNumber"]
		if hx_field_46 == nil {
			var hx_zero_47 int
			return hx_zero_47
		}
		return hx_field_46.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
