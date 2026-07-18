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
		hx_obj_68 := map[string]any{}
		hx_obj_68["fileName"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_68["lineNumber"] = 0
		hx_obj_68["className"] = hxrt.StringFromLiteral("(unknown)")
		hx_obj_68["methodName"] = hxrt.StringFromLiteral("(unknown)")
		self.posInfos = hx_obj_68
	} else {
		self.posInfos = pos
	}
	return self
}

func (self *haxe__exceptions__PosException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func (self *haxe__exceptions__PosException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.ExceptionMessage(self), hxrt.StringFromLiteral(" in ")), func(hx_obj_69 map[string]any) *string {
		hx_field_70 := hx_obj_69["className"]
		if hx_field_70 == nil {
			var hx_zero_71 *string
			return hx_zero_71
		}
		return hx_field_70.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(".")), func(hx_obj_72 map[string]any) *string {
		hx_field_73 := hx_obj_72["methodName"]
		if hx_field_73 == nil {
			var hx_zero_74 *string
			return hx_zero_74
		}
		return hx_field_73.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(" at ")), func(hx_obj_75 map[string]any) *string {
		hx_field_76 := hx_obj_75["fileName"]
		if hx_field_76 == nil {
			var hx_zero_77 *string
			return hx_zero_77
		}
		return hx_field_76.(*string)
	}(self.posInfos)), hxrt.StringFromLiteral(":")), func(hx_obj_78 map[string]any) int {
		hx_field_79 := hx_obj_78["lineNumber"]
		if hx_field_79 == nil {
			var hx_zero_80 int
			return hx_zero_80
		}
		return hx_field_79.(int)
	}(self.posInfos))
}

func (self *haxe__exceptions__PosException) String() string {
	return *self.__hx_this.toString()
}
