package main

import "snapshot/hxrt"

type I_haxe__iterators__MapKeyValueIterator interface {
	hasNext() bool
	next() hxrt.Anonymous
}

type haxe__iterators__MapKeyValueIterator struct {
	__hx_this I_haxe__iterators__MapKeyValueIterator
	map_      haxe__IMap
	keys      hxrt.Anonymous
}

func (self *haxe__iterators__MapKeyValueIterator) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe__iterators__MapKeyValueIterator(map_ haxe__IMap) *haxe__iterators__MapKeyValueIterator {
	return hx_init_haxe__iterators__MapKeyValueIterator(hx_alloc_haxe__iterators__MapKeyValueIterator(nil), map_)
}

func hx_alloc_haxe__iterators__MapKeyValueIterator(owner any) *haxe__iterators__MapKeyValueIterator {
	self := &haxe__iterators__MapKeyValueIterator{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_haxe__iterators__MapKeyValueIterator)
	}
	return self
}

func hx_init_haxe__iterators__MapKeyValueIterator(self *haxe__iterators__MapKeyValueIterator, map_ haxe__IMap) *haxe__iterators__MapKeyValueIterator {
	self.map_ = map_
	self.keys = map_.keys()
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func() bool {
		var hx_callable_1 *hxrt.Function[func() bool] = hx_recover_callable_2(hxrt.AnonymousGet(self.keys, "hasNext"))
		return hxrt.FunctionCall(hx_callable_1)()
	}()
}

func (self *haxe__iterators__MapKeyValueIterator) next() hxrt.Anonymous {
	var key any = func() any {
		var hx_callable_2 *hxrt.Function[func() any] = hx_recover_callable_3(hxrt.AnonymousGet(self.keys, "next"))
		return hxrt.FunctionCall(hx_callable_2)()
	}()
	hx_obj_3 := map[string]any{}
	hx_obj_3["key"] = key
	hx_obj_3["value"] = self.map_.getIMap(key)
	return hxrt.AnonymousFromMap(hx_obj_3)
}
