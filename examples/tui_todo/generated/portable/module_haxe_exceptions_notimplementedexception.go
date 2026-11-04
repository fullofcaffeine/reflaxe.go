package main

import "examples_tui_todo_portable/hxrt"

type I_haxe__exceptions__NotImplementedException interface {
	get_message() *string
	get_stack() any
	get_previous() *haxe__Exception
	get_native() any
	unwrap() any
	toString() *string
	details() *string
}

type haxe__exceptions__NotImplementedException struct {
	*haxe__exceptions__PosException
	__hx_this I_haxe__exceptions__NotImplementedException
}

func (self *haxe__exceptions__NotImplementedException) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe__exceptions__NotImplementedException(message *string, previous *haxe__Exception, pos hxrt.Anonymous) *haxe__exceptions__NotImplementedException {
	return hx_init_haxe__exceptions__NotImplementedException(hx_alloc_haxe__exceptions__NotImplementedException(nil), message, previous, pos)
}

func hx_alloc_haxe__exceptions__NotImplementedException(owner any) *haxe__exceptions__NotImplementedException {
	self := &haxe__exceptions__NotImplementedException{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_haxe__exceptions__NotImplementedException)
	}
	self.haxe__exceptions__PosException = hx_alloc_haxe__exceptions__PosException(self.__hx_this)
	return self
}

func hx_init_haxe__exceptions__NotImplementedException(self *haxe__exceptions__NotImplementedException, message *string, previous *haxe__Exception, pos hxrt.Anonymous) *haxe__exceptions__NotImplementedException {
	hx_init_haxe__exceptions__PosException(self.haxe__exceptions__PosException, func() *string {
		var hx_if_1 *string
		if hxrt.StringEqualStringPtr(message, nil) {
			hx_if_1 = hxrt.StringFromLiteral("Not implemented")
		} else {
			hx_if_1 = message
		}
		return hx_if_1
	}(), previous, pos)
	return self
}

func (self *haxe__exceptions__NotImplementedException) String() string {
	return *self.__hx_this.toString()
}
