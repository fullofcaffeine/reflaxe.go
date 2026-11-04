package main

import "snapshot/hxrt"

type I_helper__Greeter interface {
	hello(name *string) *string
}

type helper__Greeter struct {
	__hx_this I_helper__Greeter
}

func (self *helper__Greeter) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_helper__Greeter() *helper__Greeter {
	return hx_init_helper__Greeter(hx_alloc_helper__Greeter(nil))
}

func hx_alloc_helper__Greeter(owner any) *helper__Greeter {
	self := &helper__Greeter{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_helper__Greeter)
	}
	return self
}

func hx_init_helper__Greeter(self *helper__Greeter) *helper__Greeter {
	return self
}

func (self *helper__Greeter) hello(name *string) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("hello,"), name)
}
