package main

import "snapshot/hxrt"

type I_haxe__iterators__MapKeyValueIterator interface {
	hasNext() bool
	next() map[string]any
}

type haxe__iterators__MapKeyValueIterator struct {
	__hx_this I_haxe__iterators__MapKeyValueIterator
	map_      haxe__IMap
	keys      map[string]any
}

func (self *haxe__iterators__MapKeyValueIterator) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe__iterators__MapKeyValueIterator(map_ haxe__IMap) *haxe__iterators__MapKeyValueIterator {
	self := &haxe__iterators__MapKeyValueIterator{}
	self.__hx_this = self
	self.map_ = map_
	self.keys = func(hx_value_1 any) map[string]any {
		if hx_value_1 == nil {
			var hx_zero_2 map[string]any
			return hx_zero_2
		}
		return hx_value_1.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func() bool {
		var hx_callable_6 *hxrt.Function[func() bool] = func(hx_obj_3 map[string]any) *hxrt.Function[func() bool] {
			hx_field_4 := hx_obj_3["hasNext"]
			if hx_field_4 == nil {
				var hx_zero_5 *hxrt.Function[func() bool]
				return hx_zero_5
			}
			return hx_field_4.(*hxrt.Function[func() bool])
		}(self.keys)
		return hxrt.FunctionCall(hx_callable_6)()
	}()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func() any {
		var hx_callable_10 *hxrt.Function[func() any] = func(hx_obj_7 map[string]any) *hxrt.Function[func() any] {
			hx_field_8 := hx_obj_7["next"]
			if hx_field_8 == nil {
				var hx_zero_9 *hxrt.Function[func() any]
				return hx_zero_9
			}
			return hx_field_8.(*hxrt.Function[func() any])
		}(self.keys)
		return hxrt.FunctionCall(hx_callable_10)()
	}()
	hx_obj_11 := map[string]any{}
	hx_obj_11["key"] = key
	hx_obj_11["value"] = self.map_.getIMap(key)
	return hx_obj_11
}
