package main

import "examples_portable_beta/hxrt"

type I_haxe__ValueException interface {
	get_message() *string
	get_stack() any
	get_previous() *haxe__Exception
	get_native() any
	unwrap() any
	toString() *string
	details() *string
}

type haxe__ValueException struct {
	*haxe__Exception
	__hx_this I_haxe__ValueException
	value     any
}

func (self *haxe__ValueException) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe__ValueException(value any, previous *haxe__Exception, native any) *haxe__ValueException {
	return hx_init_haxe__ValueException(hx_alloc_haxe__ValueException(nil), value, previous, native)
}

func hx_alloc_haxe__ValueException(owner any) *haxe__ValueException {
	self := &haxe__ValueException{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_haxe__ValueException)
	}
	self.haxe__Exception = hx_alloc_haxe__Exception(self.__hx_this)
	return self
}

func hx_init_haxe__ValueException(self *haxe__ValueException, value any, previous *haxe__Exception, native any) *haxe__ValueException {
	hx_init_haxe__Exception(self.haxe__Exception, func() *string {
		var hx_if_1 *string
		if hxrt.AnyEqualsNull(value) {
			hx_if_1 = hxrt.StringFromLiteral("null")
		} else {
			hx_if_1 = hxrt.StdString(value)
		}
		return hx_if_1
	}(), previous, native)
	self.value = value
	return self
}

func (self *haxe__ValueException) unwrap() any {
	return self.value
}

func (self *haxe__ValueException) String() string {
	return *self.__hx_this.toString()
}
