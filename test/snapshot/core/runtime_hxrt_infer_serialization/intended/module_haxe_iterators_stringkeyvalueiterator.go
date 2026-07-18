package main

import "snapshot/hxrt"

type I_haxe__iterators__StringKeyValueIterator interface {
	hasNext() bool
	next() map[string]any
}

type haxe__iterators__StringKeyValueIterator struct {
	__hx_this I_haxe__iterators__StringKeyValueIterator
	offset    int
	s         *string
}

func New_haxe__iterators__StringKeyValueIterator(s *string) *haxe__iterators__StringKeyValueIterator {
	self := &haxe__iterators__StringKeyValueIterator{}
	self.__hx_this = self
	self.offset = 0
	self.s = s
	return self
}

func (self *haxe__iterators__StringKeyValueIterator) hasNext() bool {
	return (self.offset < hxrt.StringLengthStringPtr(self.s))
}

func (self *haxe__iterators__StringKeyValueIterator) next() map[string]any {
	current := self.offset
	self.offset = int(int32((self.offset + 1)))
	code := hxrt.StringCharCodeAtStringPtr(self.s, current)
	hx_obj_235 := map[string]any{}
	hx_obj_235["key"] = current
	hx_obj_235["value"] = code
	return hx_obj_235
}
