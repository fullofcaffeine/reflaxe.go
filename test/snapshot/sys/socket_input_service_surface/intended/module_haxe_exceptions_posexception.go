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
		hx_obj_80 := map[string]any{}
		hx_obj_80["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_80["lineNumber"] = 0
		hx_obj_80["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_80["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_80
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_81 map[string]any) *string {
		hx_field_82 := hx_obj_81["className"]
		if hx_field_82 == nil {
			var hx_zero_83 *string
			return hx_zero_83
		}
		return hx_field_82.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_84 map[string]any) *string {
		hx_field_85 := hx_obj_84["methodName"]
		if hx_field_85 == nil {
			var hx_zero_86 *string
			return hx_zero_86
		}
		return hx_field_85.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_87 map[string]any) *string {
		hx_field_88 := hx_obj_87["fileName"]
		if hx_field_88 == nil {
			var hx_zero_89 *string
			return hx_zero_89
		}
		return hx_field_88.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_90 map[string]any) int {
		hx_field_91 := hx_obj_90["lineNumber"]
		if hx_field_91 == nil {
			var hx_zero_92 int
			return hx_zero_92
		}
		return hx_field_91.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
