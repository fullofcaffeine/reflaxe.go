package main

import "examples_portable_beta/hxrt"

type I_haxe__iterators__StringKeyValueIterator interface {
	hasNext() bool
	next() hxrt.Anonymous
}

type haxe__iterators__StringKeyValueIterator struct {
	__hx_this I_haxe__iterators__StringKeyValueIterator
	offset    int
	s         *string
}

func (self *haxe__iterators__StringKeyValueIterator) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe__iterators__StringKeyValueIterator(s *string) *haxe__iterators__StringKeyValueIterator {
	return hx_init_haxe__iterators__StringKeyValueIterator(hx_alloc_haxe__iterators__StringKeyValueIterator(nil), s)
}

func hx_alloc_haxe__iterators__StringKeyValueIterator(owner any) *haxe__iterators__StringKeyValueIterator {
	self := &haxe__iterators__StringKeyValueIterator{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_haxe__iterators__StringKeyValueIterator)
	}
	return self
}

func hx_init_haxe__iterators__StringKeyValueIterator(self *haxe__iterators__StringKeyValueIterator, s *string) *haxe__iterators__StringKeyValueIterator {
	self.offset = 0
	self.s = s
	return self
}

func (self *haxe__iterators__StringKeyValueIterator) hasNext() bool {
	return (self.offset < hxrt.StringLengthStringPtr(self.s))
}

func (self *haxe__iterators__StringKeyValueIterator) next() hxrt.Anonymous {
	current := self.offset
	self.offset = int(int32((self.offset + 1)))
	code := hxrt.StringCharCodeAtStringPtr(self.s, current)
	hx_obj_1 := map[string]any{}
	hx_obj_1["key"] = current
	hx_obj_1["value"] = code
	return hxrt.AnonymousFromMap(hx_obj_1)
}
