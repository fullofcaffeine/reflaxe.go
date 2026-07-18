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
		hx_obj_70 := map[string]any{}
		hx_obj_70["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_70["lineNumber"] = 0
		hx_obj_70["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_70["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_70
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_71 map[string]any) *string {
		hx_field_72 := hx_obj_71["className"]
		if hx_field_72 == nil {
			var hx_zero_73 *string
			return hx_zero_73
		}
		return hx_field_72.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_74 map[string]any) *string {
		hx_field_75 := hx_obj_74["methodName"]
		if hx_field_75 == nil {
			var hx_zero_76 *string
			return hx_zero_76
		}
		return hx_field_75.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_77 map[string]any) *string {
		hx_field_78 := hx_obj_77["fileName"]
		if hx_field_78 == nil {
			var hx_zero_79 *string
			return hx_zero_79
		}
		return hx_field_78.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_80 map[string]any) int {
		hx_field_81 := hx_obj_80["lineNumber"]
		if hx_field_81 == nil {
			var hx_zero_82 int
			return hx_zero_82
		}
		return hx_field_81.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
