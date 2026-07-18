package main

import "snapshot/hxrt"

type I_haxe__exceptions__ArgumentException interface {
	toString() *string
}

type haxe__exceptions__ArgumentException struct {
	*haxe__exceptions__PosException
	__hx_this I_haxe__exceptions__ArgumentException
	argument  *string
}

func New_haxe__exceptions__ArgumentException(argument *string, message *string, previous *hxrt.ExceptionValue, pos map[string]any) *haxe__exceptions__ArgumentException {
	self := &haxe__exceptions__ArgumentException{}
	self.haxe__exceptions__PosException = New_haxe__exceptions__PosException(func() *string {
		var hx_if_16 *string
		if hxrt.StringEqualStringPtr(message, nil) {
			hx_if_16 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Invalid argument \""), argument), hxrt.StringFromLiteral("\""))
		} else {
			hx_if_16 = message
		}
		return hx_if_16
	}(), previous, pos)
	self.haxe__exceptions__PosException.__hx_this = self
	self.__hx_this = self
	self.argument = argument
	return self
}

func (self *haxe__exceptions__ArgumentException) String() string {
	return *self.__hx_this.toString()
}
