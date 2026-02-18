package main

import "snapshot/hxrt"

type I_Base interface {
	ping() int
}

type Base struct {
	__hx_this I_Base
}

func New_Base() *Base {
	self := &Base{}
	self.__hx_this = self
	return self
}

func (self *Base) ping() int {
	return 1
}

type I_Child interface {
	ping() int
}

type Child struct {
	*Base
	__hx_this I_Child
}

func New_Child() *Child {
	self := &Child{}
	self.Base = &Base{}
	self.Base.__hx_this = self
	self.__hx_this = self
	return self
}

func (self *Child) ping() int {
	return 2
}

func main() {
	child := New_Child()
	hxrt.Println(child.ping())
}
