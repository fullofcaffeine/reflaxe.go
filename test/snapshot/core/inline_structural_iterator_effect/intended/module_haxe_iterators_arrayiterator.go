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

func (self *haxe__iterators__ArrayIterator) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe__iterators__ArrayIterator(array *hxrt.Array) *haxe__iterators__ArrayIterator {
	return hx_init_haxe__iterators__ArrayIterator(hx_alloc_haxe__iterators__ArrayIterator(nil), array)
}

func hx_alloc_haxe__iterators__ArrayIterator(owner any) *haxe__iterators__ArrayIterator {
	self := &haxe__iterators__ArrayIterator{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_haxe__iterators__ArrayIterator)
	}
	return self
}

func hx_init_haxe__iterators__ArrayIterator(self *haxe__iterators__ArrayIterator, array *hxrt.Array) *haxe__iterators__ArrayIterator {
	self.current = 0
	self.array = array
	return self
}

func (self *haxe__iterators__ArrayIterator) hasNext() bool {
	return (self.current < self.array.Len())
}

func (self *haxe__iterators__ArrayIterator) next() any {
	hx_post_1 := self.current
	self.current = int(int32((self.current + 1)))
	return self.array.Get(hx_post_1)
}
