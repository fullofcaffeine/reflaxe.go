package main

import "snapshot/hxrt"

type I_haxe__iterators__StringIteratorUnicode interface {
	hasNext() bool
	next() int
}

type haxe__iterators__StringIteratorUnicode struct {
	__hx_this I_haxe__iterators__StringIteratorUnicode
	offset    int
	s         *string
}

func New_haxe__iterators__StringIteratorUnicode(s *string) *haxe__iterators__StringIteratorUnicode {
	self := &haxe__iterators__StringIteratorUnicode{}
	self.__hx_this = self
	self.offset = 0
	self.s = s
	return self
}

func (self *haxe__iterators__StringIteratorUnicode) hasNext() bool {
	return (self.offset < hxrt.StringLengthStringPtr(self.s))
}

func (self *haxe__iterators__StringIteratorUnicode) next() int {
	value := self.s
	hx_post_45 := self.offset
	self.offset = int(int32((self.offset + 1)))
	index := hx_post_45
	c := hxrt.StringCharCodeAtStringPtr(value, index)
	if ((c >= 55296) && (c <= 56319)) && (self.offset < hxrt.StringLengthStringPtr(self.s)) {
		c = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(c) - hxrt.Int32Wrap(55232))))) << uint(10))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(func() int {
			value_1 := self.s
			hx_post_46 := self.offset
			self.offset = int(int32((self.offset + 1)))
			index_1 := hx_post_46
			return hxrt.StringCharCodeAtStringPtr(value_1, index_1)
		}()) & hxrt.Int32Wrap(1023))))))))
	}
	return c
}

func haxe__iterators__StringIteratorUnicode_codeAt(value *string, index int) int {
	return hxrt.StringCharCodeAtStringPtr(value, index)
}

func haxe__iterators__StringIteratorUnicode_unicodeIterator(s *string) *haxe__iterators__StringIteratorUnicode {
	return New_haxe__iterators__StringIteratorUnicode(s)
}
