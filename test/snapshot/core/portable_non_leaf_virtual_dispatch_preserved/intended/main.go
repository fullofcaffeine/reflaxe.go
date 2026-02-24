package main

import "snapshot/hxrt"

type I_Base interface {
	who() int
	callWho() int
}

type Base struct {
	__hx_this I_Base
}

func New_Base() *Base {
	self := &Base{}
	self.__hx_this = self
	return self
}

func (self *Base) who() int {
	return 1
}

func (self *Base) callWho() int {
	return self.__hx_this.who()
}

type I_Child interface {
	who() int
	callWho() int
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

func (self *Child) who() int {
	return 2
}

func main() {
	child := New_Child()
	base := child.Base
	hxrt.Println(base.__hx_this.callWho())
}
