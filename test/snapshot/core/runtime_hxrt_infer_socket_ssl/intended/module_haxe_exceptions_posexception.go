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
		hx_obj_96 := map[string]any{}
		hx_obj_96["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_96["lineNumber"] = 0
		hx_obj_96["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_96["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_96
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_97 map[string]any) *string {
		hx_field_98 := hx_obj_97["className"]
		if hx_field_98 == nil {
			var hx_zero_99 *string
			return hx_zero_99
		}
		return hx_field_98.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_100 map[string]any) *string {
		hx_field_101 := hx_obj_100["methodName"]
		if hx_field_101 == nil {
			var hx_zero_102 *string
			return hx_zero_102
		}
		return hx_field_101.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_103 map[string]any) *string {
		hx_field_104 := hx_obj_103["fileName"]
		if hx_field_104 == nil {
			var hx_zero_105 *string
			return hx_zero_105
		}
		return hx_field_104.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_106 map[string]any) int {
		hx_field_107 := hx_obj_106["lineNumber"]
		if hx_field_107 == nil {
			var hx_zero_108 int
			return hx_zero_108
		}
		return hx_field_107.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
