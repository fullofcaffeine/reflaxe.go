package main

import "snapshot/hxrt"

func main() {
	first := New__Main__AtomicBox(1)
	second := New__Main__AtomicBox(2)
	var this1 *hxrt.AtomicObjectCell
	this1 = hxrt.AtomicObjectNew(first)
	value := this1
	hxrt.AtomicObjectCompareExchange(value, first, second)
}

type I__Main__AtomicBox interface {
}

type _Main__AtomicBox struct {
	__hx_this I__Main__AtomicBox
	value     int
}

func New__Main__AtomicBox(value int) *_Main__AtomicBox {
	self := &_Main__AtomicBox{}
	self.__hx_this = self
	self.value = value
	return self
}
