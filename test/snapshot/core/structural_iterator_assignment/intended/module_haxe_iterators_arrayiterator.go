package main

import "snapshot/hxrt"

type I_haxe__iterators__ArrayIterator interface {
	hasNext() bool
	next() any
}

type haxe__iterators__ArrayIterator struct {
	__hx_this I_haxe__iterators__ArrayIterator
	array     *hxrt.Array
	current   int
}

func New_haxe__iterators__ArrayIterator(array *hxrt.Array) *haxe__iterators__ArrayIterator {
	self := &haxe__iterators__ArrayIterator{}
	self.__hx_this = self
	self.current = 0
	self.array = array
	return self
}

func (self *haxe__iterators__ArrayIterator) hasNext() bool {
	return (self.current < self.array.Len())
}

func (self *haxe__iterators__ArrayIterator) next() any {
	hx_post_49 := self.current
	self.current = int(int32((self.current + 1)))
	return self.array.Get(hx_post_49)
}
