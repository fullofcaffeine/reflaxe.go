package main

import "snapshot/hxrt"

type I_helper__Greeter interface {
	hello(name *string) *string
}

type helper__Greeter struct {
	__hx_this I_helper__Greeter
}

func New_helper__Greeter() *helper__Greeter {
	self := &helper__Greeter{}
	self.__hx_this = self
	return self
}

func (self *helper__Greeter) hello(name *string) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("hello,"), name)
}
