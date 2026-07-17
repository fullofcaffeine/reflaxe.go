package main

import "snapshot/hxrt"

type I_go___Slice interface {
	get_length() int
	push(value any)
	get(index int) any
	set(index int, value any)
	toArray() *hxrt.Array
}

type go___Slice struct {
	__hx_this I_go___Slice
	data      []any
	length    int
}

func New_go___Slice() *go___Slice {
	self := &go___Slice{}
	self.__hx_this = self
	self.data = []any{}
	return self
}

func (self *go___Slice) get_length() int {
	return len(self.data)
}

func (self *go___Slice) push(value any) {
	self.data = func() []any {
		hx_native_slice_17 := self.data
		return append(hx_native_slice_17, value)
	}()
}

func (self *go___Slice) get(index int) any {
	return self.data[index]
}

func (self *go___Slice) set(index int, value any) {
	self.data[index] = value
}

func (self *go___Slice) toArray() *hxrt.Array {
	return hxrt.ArrayFromValues(self.data)
}
