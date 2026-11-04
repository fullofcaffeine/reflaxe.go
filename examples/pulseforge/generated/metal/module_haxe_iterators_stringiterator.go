package main

import "examples_pulseforge_metal/hxrt"

type I_haxe__iterators__StringIterator interface {
	hasNext() bool
	next() int
}

type haxe__iterators__StringIterator struct {
	__hx_this I_haxe__iterators__StringIterator
	offset    int
	s         *string
}

func (self *haxe__iterators__StringIterator) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe__iterators__StringIterator(s *string) *haxe__iterators__StringIterator {
	return hx_init_haxe__iterators__StringIterator(hx_alloc_haxe__iterators__StringIterator(nil), s)
}

func hx_alloc_haxe__iterators__StringIterator(owner any) *haxe__iterators__StringIterator {
	self := &haxe__iterators__StringIterator{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_haxe__iterators__StringIterator)
	}
	return self
}

func hx_init_haxe__iterators__StringIterator(self *haxe__iterators__StringIterator, s *string) *haxe__iterators__StringIterator {
	self.offset = 0
	self.s = s
	return self
}

func (self *haxe__iterators__StringIterator) hasNext() bool {
	return (self.offset < hxrt.StringLengthStringPtr(self.s))
}

func (self *haxe__iterators__StringIterator) next() int {
	return hxrt.StringCharCodeAtStringPtr(self.s, func() int {
		hx_post_1 := self.offset
		self.offset = int(int32((self.offset + 1)))
		return hx_post_1
	}())
}
