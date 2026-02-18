package main

import "snapshot/hxrt"

type I_Base interface {
	tag() int
}

type Base struct {
	__hx_this I_Base
}

func New_Base() *Base {
	self := &Base{}
	self.__hx_this = self
	return self
}

func (self *Base) tag() int {
	return 1
}

type I_Child interface {
	tag() int
}

type Child struct {
	*Base
	__hx_this I_Child
}

func New_Child() *Child {
	self := &Child{}
	self.Base = New_Base()
	self.Base.__hx_this = self
	self.__hx_this = self
	return self
}

func (self *Child) tag() int {
	return 2
}

func Factory_makeBase(flag bool) *Base {
	if flag {
		return New_Child().Base
	}
	return New_Base()
}

func main() {
	asBase := Factory_makeBase(true)
	_ = asBase
	hxrt.Println(asBase.__hx_this.tag())
}
