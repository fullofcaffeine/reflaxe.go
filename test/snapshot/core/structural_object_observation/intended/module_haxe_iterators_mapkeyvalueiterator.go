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
	self := &haxe__iterators__MapKeyValueIterator{}
	self.__hx_this = self
	self.map_ = map_
	self.keys = map_.keys()
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func() bool {
		var hx_callable_4 *hxrt.Function[func() bool] = func(hx_obj_1 hxrt.Anonymous) *hxrt.Function[func() bool] {
			hx_field_2 := hxrt.AnonymousGet(hx_obj_1, "hasNext")
			if hx_field_2 == nil {
				var hx_zero_3 *hxrt.Function[func() bool]
				return hx_zero_3
			}
			return hx_field_2.(*hxrt.Function[func() bool])
		}(self.keys)
		return hxrt.FunctionCall(hx_callable_4)()
	}()
}

func (self *haxe__iterators__MapKeyValueIterator) next() hxrt.Anonymous {
	var key any = func() any {
		var hx_callable_8 *hxrt.Function[func() any] = func(hx_obj_5 hxrt.Anonymous) *hxrt.Function[func() any] {
			hx_field_6 := hxrt.AnonymousGet(hx_obj_5, "next")
			if hx_field_6 == nil {
				var hx_zero_7 *hxrt.Function[func() any]
				return hx_zero_7
			}
			return hx_field_6.(*hxrt.Function[func() any])
		}(self.keys)
		return hxrt.FunctionCall(hx_callable_8)()
	}()
	hx_obj_9 := map[string]any{}
	hx_obj_9["key"] = key
	hx_obj_9["value"] = self.map_.getIMap(key)
	return hxrt.AnonymousFromMap(hx_obj_9)
}
