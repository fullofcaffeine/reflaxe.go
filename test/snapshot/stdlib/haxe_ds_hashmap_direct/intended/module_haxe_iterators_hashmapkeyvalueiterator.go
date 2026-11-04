package main

import "snapshot/hxrt"

type I_haxe__iterators__HashMapKeyValueIterator interface {
	hasNext() bool
	next() hxrt.Anonymous
}

type haxe__iterators__HashMapKeyValueIterator struct {
	__hx_this I_haxe__iterators__HashMapKeyValueIterator
	map_      *haxe__ds__HashMap
	keys      hxrt.Anonymous
}

func (self *haxe__iterators__HashMapKeyValueIterator) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe__iterators__HashMapKeyValueIterator(map_ *haxe__ds__HashMap) *haxe__iterators__HashMapKeyValueIterator {
	return hx_init_haxe__iterators__HashMapKeyValueIterator(hx_alloc_haxe__iterators__HashMapKeyValueIterator(nil), map_)
}

func hx_alloc_haxe__iterators__HashMapKeyValueIterator(owner any) *haxe__iterators__HashMapKeyValueIterator {
	self := &haxe__iterators__HashMapKeyValueIterator{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_haxe__iterators__HashMapKeyValueIterator)
	}
	return self
}

func hx_init_haxe__iterators__HashMapKeyValueIterator(self *haxe__iterators__HashMapKeyValueIterator, map_ *haxe__ds__HashMap) *haxe__iterators__HashMapKeyValueIterator {
	self.map_ = map_
	_gthis := map_
	hashes := map_.keysByHash.__hx_this.keys()
	hx_obj_1 := map[string]any{}
	hx_obj_1["hasNext"] = hxrt.NewFunction(func() bool {
		return func() bool {
			var hx_callable_2 *hxrt.Function[func() bool] = hx_recover_callable_1(hxrt.AnonymousGet(hashes, "hasNext"))
			return hxrt.FunctionCall(hx_callable_2)()
		}()
	})
	hx_obj_1["next"] = hxrt.NewFunction(func() hxrt.Anonymous {
		return hx_structural_recover(_gthis.keysByHash.__hx_this.get(func() int {
			var hx_callable_3 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(hashes, "next"))
			return hxrt.FunctionCall(hx_callable_3)()
		}()))
	})
	self.keys = hxrt.AnonymousFromMap(hx_obj_1)
	return self
}

func (self *haxe__iterators__HashMapKeyValueIterator) hasNext() bool {
	return func() bool {
		var hx_callable_4 *hxrt.Function[func() bool] = hx_recover_callable_1(hxrt.AnonymousGet(self.keys, "hasNext"))
		return hxrt.FunctionCall(hx_callable_4)()
	}()
}

func (self *haxe__iterators__HashMapKeyValueIterator) next() hxrt.Anonymous {
	key := func() hxrt.Anonymous {
		var hx_callable_5 *hxrt.Function[func() hxrt.Anonymous] = hx_recover_callable_2(hxrt.AnonymousGet(self.keys, "next"))
		return hxrt.FunctionCall(hx_callable_5)()
	}()
	_this := self.map_
	var value any = _this.valuesByHash.__hx_this.get(func() int {
		var hx_callable_6 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(key, "hashCode"))
		return hxrt.FunctionCall(hx_callable_6)()
	}())
	hx_obj_7 := map[string]any{}
	hx_obj_7["key"] = key
	hx_obj_7["value"] = value
	return hxrt.AnonymousFromMap(hx_obj_7)
}
