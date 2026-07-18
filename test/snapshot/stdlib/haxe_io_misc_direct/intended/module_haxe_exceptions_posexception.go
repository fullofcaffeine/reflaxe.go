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
		hx_obj_76 := map[string]any{}
		hx_obj_76["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_76["lineNumber"] = 0
		hx_obj_76["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_76["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_76
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_77 map[string]any) *string {
		hx_field_78 := hx_obj_77["className"]
		if hx_field_78 == nil {
			var hx_zero_79 *string
			return hx_zero_79
		}
		return hx_field_78.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_80 map[string]any) *string {
		hx_field_81 := hx_obj_80["methodName"]
		if hx_field_81 == nil {
			var hx_zero_82 *string
			return hx_zero_82
		}
		return hx_field_81.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_83 map[string]any) *string {
		hx_field_84 := hx_obj_83["fileName"]
		if hx_field_84 == nil {
			var hx_zero_85 *string
			return hx_zero_85
		}
		return hx_field_84.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_86 map[string]any) int {
		hx_field_87 := hx_obj_86["lineNumber"]
		if hx_field_87 == nil {
			var hx_zero_88 int
			return hx_zero_88
		}
		return hx_field_87.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
