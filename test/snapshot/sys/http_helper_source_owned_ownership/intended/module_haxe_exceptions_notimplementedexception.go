package main

import "snapshot/hxrt"

type I_haxe__exceptions__NotImplementedException interface {
	toString() *string
}

type haxe__exceptions__NotImplementedException struct {
	*haxe__exceptions__PosException
	__hx_this I_haxe__exceptions__NotImplementedException
}

func New_haxe__exceptions__NotImplementedException(message *string, previous *hxrt.ExceptionValue, pos map[string]any) *haxe__exceptions__NotImplementedException {
	self := &haxe__exceptions__NotImplementedException{}
	self.haxe__exceptions__PosException = New_haxe__exceptions__PosException(func() *string {
		var hx_if_32 *string
		if hxrt.StringEqualStringPtr(message, nil) {
			hx_if_32 = hxrt.StringFromLiteral("Not implemented")
		} else {
			hx_if_32 = message
		}
		return hx_if_32
	}(), previous, pos)
	self.haxe__exceptions__PosException.__hx_this = self
	self.__hx_this = self
	return self
}

func (self *haxe__exceptions__NotImplementedException) String() string {
	return *self.__hx_this.toString()
}
