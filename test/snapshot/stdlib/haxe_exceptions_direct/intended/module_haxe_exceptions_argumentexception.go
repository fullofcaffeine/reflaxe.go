package main

import "snapshot/hxrt"

type I_haxe__exceptions__ArgumentException interface {
	get_message() *string
	get_stack() any
	get_previous() *haxe__Exception
	get_native() any
	unwrap() any
	toString() *string
	details() *string
}

type haxe__exceptions__ArgumentException struct {
	*haxe__exceptions__PosException
	__hx_this I_haxe__exceptions__ArgumentException
	argument  *string
}

func (self *haxe__exceptions__ArgumentException) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe__exceptions__ArgumentException(argument *string, message *string, previous *haxe__Exception, pos hxrt.Anonymous) *haxe__exceptions__ArgumentException {
	return hx_init_haxe__exceptions__ArgumentException(hx_alloc_haxe__exceptions__ArgumentException(nil), argument, message, previous, pos)
}

func hx_alloc_haxe__exceptions__ArgumentException(owner any) *haxe__exceptions__ArgumentException {
	self := &haxe__exceptions__ArgumentException{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_haxe__exceptions__ArgumentException)
	}
	self.haxe__exceptions__PosException = hx_alloc_haxe__exceptions__PosException(self.__hx_this)
	return self
}

func hx_init_haxe__exceptions__ArgumentException(self *haxe__exceptions__ArgumentException, argument *string, message *string, previous *haxe__Exception, pos hxrt.Anonymous) *haxe__exceptions__ArgumentException {
	hx_init_haxe__exceptions__PosException(self.haxe__exceptions__PosException, func() *string {
		var hx_if_1 *string
		if hxrt.StringEqualStringPtr(message, nil) {
			hx_if_1 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Invalid argument \""), argument), hxrt.StringFromLiteral("\""))
		} else {
			hx_if_1 = message
		}
		return hx_if_1
	}(), previous, pos)
	self.argument = argument
	return self
}

func (self *haxe__exceptions__ArgumentException) String() string {
	return *self.__hx_this.toString()
}
