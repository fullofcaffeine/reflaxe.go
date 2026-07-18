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
		hx_obj_89 := map[string]any{}
		hx_obj_89["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_89["lineNumber"] = 0
		hx_obj_89["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_89["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_89
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_90 map[string]any) *string {
		hx_field_91 := hx_obj_90["className"]
		if hx_field_91 == nil {
			var hx_zero_92 *string
			return hx_zero_92
		}
		return hx_field_91.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_93 map[string]any) *string {
		hx_field_94 := hx_obj_93["methodName"]
		if hx_field_94 == nil {
			var hx_zero_95 *string
			return hx_zero_95
		}
		return hx_field_94.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_96 map[string]any) *string {
		hx_field_97 := hx_obj_96["fileName"]
		if hx_field_97 == nil {
			var hx_zero_98 *string
			return hx_zero_98
		}
		return hx_field_97.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_99 map[string]any) int {
		hx_field_100 := hx_obj_99["lineNumber"]
		if hx_field_100 == nil {
			var hx_zero_101 int
			return hx_zero_101
		}
		return hx_field_100.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
