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

func (self *_Main__AtomicBox) HxFunctionReceiver() any {
	return self.__hx_this
}

func New__Main__AtomicBox(value int) *_Main__AtomicBox {
	return hx_init__Main__AtomicBox(hx_alloc__Main__AtomicBox(nil), value)
}

func hx_alloc__Main__AtomicBox(owner any) *_Main__AtomicBox {
	self := &_Main__AtomicBox{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I__Main__AtomicBox)
	}
	return self
}

func hx_init__Main__AtomicBox(self *_Main__AtomicBox, value int) *_Main__AtomicBox {
	self.value = value
	return self
}
