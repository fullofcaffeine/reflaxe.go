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
		hx_obj_94 := map[string]any{}
		hx_obj_94["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_94["lineNumber"] = 0
		hx_obj_94["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_94["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_94
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_95 map[string]any) *string {
		hx_field_96 := hx_obj_95["className"]
		if hx_field_96 == nil {
			var hx_zero_97 *string
			return hx_zero_97
		}
		return hx_field_96.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_98 map[string]any) *string {
		hx_field_99 := hx_obj_98["methodName"]
		if hx_field_99 == nil {
			var hx_zero_100 *string
			return hx_zero_100
		}
		return hx_field_99.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_101 map[string]any) *string {
		hx_field_102 := hx_obj_101["fileName"]
		if hx_field_102 == nil {
			var hx_zero_103 *string
			return hx_zero_103
		}
		return hx_field_102.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_104 map[string]any) int {
		hx_field_105 := hx_obj_104["lineNumber"]
		if hx_field_105 == nil {
			var hx_zero_106 int
			return hx_zero_106
		}
		return hx_field_105.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
