package main

import "snapshot/hxrt"

type I_haxe__ds__HashMap interface {
	set(k hxrt.Anonymous, v any)
	get(k hxrt.Anonymous) any
	exists(k hxrt.Anonymous) bool
	remove(k hxrt.Anonymous) bool
	keys() hxrt.Anonymous
	iterator() hxrt.Anonymous
	keyValueIterator() *haxe__iterators__HashMapKeyValueIterator
	copy() *haxe__ds__HashMap
	clear()
}

type haxe__ds__HashMap struct {
	__hx_this    I_haxe__ds__HashMap
	keysByHash   *haxe__ds__IntMap
	valuesByHash *haxe__ds__IntMap
}

func (self *haxe__ds__HashMap) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe__ds__HashMap() *haxe__ds__HashMap {
	return hx_init_haxe__ds__HashMap(hx_alloc_haxe__ds__HashMap(nil))
}

func hx_alloc_haxe__ds__HashMap(owner any) *haxe__ds__HashMap {
	self := &haxe__ds__HashMap{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_haxe__ds__HashMap)
	}
	return self
}

func hx_init_haxe__ds__HashMap(self *haxe__ds__HashMap) *haxe__ds__HashMap {
	self.keysByHash = New_haxe__ds__IntMap()
	self.valuesByHash = New_haxe__ds__IntMap()
	return self
}

func (self *haxe__ds__HashMap) set(k hxrt.Anonymous, v any) {
	hash := func() int {
		var hx_callable_1 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(k, "hashCode"))
		return hxrt.FunctionCall(hx_callable_1)()
	}()
	self.keysByHash.__hx_this.set(hash, k)
	self.valuesByHash.__hx_this.set(hash, v)
}

func (self *haxe__ds__HashMap) get(k hxrt.Anonymous) any {
	return self.valuesByHash.__hx_this.get(func() int {
		var hx_callable_2 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(k, "hashCode"))
		return hxrt.FunctionCall(hx_callable_2)()
	}())
}

func (self *haxe__ds__HashMap) exists(k hxrt.Anonymous) bool {
	return func(hx_value_4 any) bool {
		if hx_value_4 == nil {
			var hx_zero_5 bool
			return hx_zero_5
		}
		return hx_value_4.(bool)
	}(self.valuesByHash.__hx_this.exists(func() int {
		var hx_callable_3 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(k, "hashCode"))
		return hxrt.FunctionCall(hx_callable_3)()
	}()))
}

func (self *haxe__ds__HashMap) remove(k hxrt.Anonymous) bool {
	hash := func() int {
		var hx_callable_6 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(k, "hashCode"))
		return hxrt.FunctionCall(hx_callable_6)()
	}()
	func(hx_value_7 any) bool {
		if hx_value_7 == nil {
			var hx_zero_8 bool
			return hx_zero_8
		}
		return hx_value_7.(bool)
	}(self.valuesByHash.__hx_this.remove(hash))
	return func(hx_value_9 any) bool {
		if hx_value_9 == nil {
			var hx_zero_10 bool
			return hx_zero_10
		}
		return hx_value_9.(bool)
	}(self.keysByHash.__hx_this.remove(hash))
}

func (self *haxe__ds__HashMap) keys() hxrt.Anonymous {
	_gthis := self
	hashes := self.keysByHash.__hx_this.keys()
	hx_obj_11 := map[string]any{}
	hx_obj_11["hasNext"] = hxrt.NewFunction(func() bool {
		return func() bool {
			var hx_callable_12 *hxrt.Function[func() bool] = hx_recover_callable_1(hxrt.AnonymousGet(hashes, "hasNext"))
			return hxrt.FunctionCall(hx_callable_12)()
		}()
	})
	hx_obj_11["next"] = hxrt.NewFunction(func() hxrt.Anonymous {
		return hx_structural_recover(_gthis.keysByHash.__hx_this.get(func() int {
			var hx_callable_13 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(hashes, "next"))
			return hxrt.FunctionCall(hx_callable_13)()
		}()))
	})
	return hxrt.AnonymousFromMap(hx_obj_11)
}

func (self *haxe__ds__HashMap) iterator() hxrt.Anonymous {
	_gthis := self
	hashes := self.valuesByHash.__hx_this.keys()
	hx_obj_14 := map[string]any{}
	hx_obj_14["hasNext"] = hxrt.NewFunction(func() bool {
		return func() bool {
			var hx_callable_15 *hxrt.Function[func() bool] = hx_recover_callable_1(hxrt.AnonymousGet(hashes, "hasNext"))
			return hxrt.FunctionCall(hx_callable_15)()
		}()
	})
	hx_obj_14["next"] = hxrt.NewFunction(func() any {
		return _gthis.valuesByHash.__hx_this.get(func() int {
			var hx_callable_16 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(hashes, "next"))
			return hxrt.FunctionCall(hx_callable_16)()
		}())
	})
	return hxrt.AnonymousFromMap(hx_obj_14)
}

func (self *haxe__ds__HashMap) keyValueIterator() *haxe__iterators__HashMapKeyValueIterator {
	return New_haxe__iterators__HashMapKeyValueIterator(self)
}

func (self *haxe__ds__HashMap) copy() *haxe__ds__HashMap {
	copied := New_haxe__ds__HashMap()
	hash := self.keysByHash.__hx_this.keys()
	for func() bool {
		var hx_callable_17 *hxrt.Function[func() bool] = hx_recover_callable_1(hxrt.AnonymousGet(hash, "hasNext"))
		return hxrt.FunctionCall(hx_callable_17)()
	}() {
		hash_1 := func() int {
			var hx_callable_18 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(hash, "next"))
			return hxrt.FunctionCall(hx_callable_18)()
		}()
		copied.keysByHash.__hx_this.set(hash_1, hx_structural_recover(self.keysByHash.__hx_this.get(hash_1)))
		copied.valuesByHash.__hx_this.set(hash_1, self.valuesByHash.__hx_this.get(hash_1))
	}
	return copied
}

func (self *haxe__ds__HashMap) clear() {
	_this := self.keysByHash
	hxrt.IntMapClear(_this.h)
	_this_1 := self.valuesByHash
	hxrt.IntMapClear(_this_1.h)
}

func haxe__ds__HashMap_hashOf(key hxrt.Anonymous) int {
	return func() int {
		var hx_callable_19 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(key, "hashCode"))
		return hxrt.FunctionCall(hx_callable_19)()
	}()
}
