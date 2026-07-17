package main

import "snapshot/hxrt"

type I_haxe__iterators__StringIterator interface {
	hasNext() bool
	next() int
}

type haxe__iterators__StringIterator struct {
	__hx_this I_haxe__iterators__StringIterator
	offset    int
	s         *string
}

func New_haxe__iterators__StringIterator(s *string) *haxe__iterators__StringIterator {
	self := &haxe__iterators__StringIterator{}
	self.__hx_this = self
	self.offset = 0
	self.s = s
	return self
}

func (self *haxe__iterators__StringIterator) hasNext() bool {
	return (self.offset < hxrt.StringLengthStringPtr(self.s))
}

func (self *haxe__iterators__StringIterator) next() int {
	return hxrt.StringCharCodeAtStringPtr(self.s, func() int {
		hx_post_127 := self.offset
		self.offset = int(int32((self.offset + 1)))
		return hx_post_127
	}())
}
