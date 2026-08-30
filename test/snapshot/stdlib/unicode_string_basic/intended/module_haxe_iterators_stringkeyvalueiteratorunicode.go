package main

import "snapshot/hxrt"

type I_haxe__iterators__StringKeyValueIteratorUnicode interface {
	hasNext() bool
	next() map[string]any
}

type haxe__iterators__StringKeyValueIteratorUnicode struct {
	__hx_this  I_haxe__iterators__StringKeyValueIteratorUnicode
	byteOffset int
	charOffset int
	s          *string
}

func New_haxe__iterators__StringKeyValueIteratorUnicode(s *string) *haxe__iterators__StringKeyValueIteratorUnicode {
	self := &haxe__iterators__StringKeyValueIteratorUnicode{}
	self.__hx_this = self
	self.charOffset = 0
	self.byteOffset = 0
	self.s = s
	return self
}

func (self *haxe__iterators__StringKeyValueIteratorUnicode) hasNext() bool {
	return (self.byteOffset < hxrt.StringLengthStringPtr(self.s))
}

func (self *haxe__iterators__StringKeyValueIteratorUnicode) next() map[string]any {
	value := self.s
	hx_post_1 := self.byteOffset
	self.byteOffset = int(int32((self.byteOffset + 1)))
	index := hx_post_1
	c := hxrt.StringCharCodeAtStringPtr(value, index)
	if ((c >= 55296) && (c <= 56319)) && (self.byteOffset < hxrt.StringLengthStringPtr(self.s)) {
		c = int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(c) - hxrt.Int32Wrap(55232)))) << uint(10)))) | hxrt.Int32Wrap(int((hxrt.Int32Wrap(func() int {
			value_1 := self.s
			hx_post_2 := self.byteOffset
			self.byteOffset = int(int32((self.byteOffset + 1)))
			index_1 := hx_post_2
			return hxrt.StringCharCodeAtStringPtr(value_1, index_1)
		}()) & hxrt.Int32Wrap(1023))))))
	}
	hx_obj_3 := map[string]any{}
	hx_post_4 := self.charOffset
	self.charOffset = int(int32((self.charOffset + 1)))
	hx_obj_3["key"] = hx_post_4
	hx_obj_3["value"] = c
	return hx_obj_3
}

func haxe__iterators__StringKeyValueIteratorUnicode_codeAt(value *string, index int) int {
	return hxrt.StringCharCodeAtStringPtr(value, index)
}

func haxe__iterators__StringKeyValueIteratorUnicode_unicodeKeyValueIterator(s *string) *haxe__iterators__StringKeyValueIteratorUnicode {
	return New_haxe__iterators__StringKeyValueIteratorUnicode(s)
}
