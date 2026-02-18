package main

import "snapshot/hxrt"

type I_Base interface {
	read() int
}

type Base struct {
	__hx_this I_Base
	value     int
}

func New_Base(value int) *Base {
	self := &Base{}
	self.__hx_this = self
	self.value = value
	return self
}

func (self *Base) read() int {
	return self.value
}

type I_Child interface {
	read() int
}

type Child struct {
	*Base
	__hx_this I_Child
}

func New_Child(value int) *Child {
	self := &Child{}
	self.Base = New_Base((value + 1))
	self.Base.__hx_this = self
	self.__hx_this = self
	return self
}

func main() {
	child := New_Child(4)
	base := child.Base
	show(child.Base)
	hxrt.Println(base.__hx_this.read())
}

func show(base *Base) {
	hxrt.Println(base.__hx_this.read())
}
