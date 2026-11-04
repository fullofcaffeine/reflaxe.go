package main

import "snapshot/hxrt"

type I_haxe__ds__IntMap interface {
	set(key int, value any)
	get(key int) any
	exists(key int) bool
	remove(key int) bool
	keys() hxrt.Anonymous
	iterator() hxrt.Anonymous
	keyValueIterator() hxrt.Anonymous
	getIMap(key any) any
	setIMap(key any, value any)
	existsIMap(key any) bool
	removeIMap(key any) bool
	copyIMap() haxe__IMap
	copy() *haxe__ds__IntMap
	toString() *string
	clear()
}

type haxe__ds__IntMap struct {
	__hx_this I_haxe__ds__IntMap
	h         *hxrt.IntMapCell
}

func (self *haxe__ds__IntMap) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe__ds__IntMap() *haxe__ds__IntMap {
	return hx_init_haxe__ds__IntMap(hx_alloc_haxe__ds__IntMap(nil))
}

func hx_alloc_haxe__ds__IntMap(owner any) *haxe__ds__IntMap {
	self := &haxe__ds__IntMap{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_haxe__ds__IntMap)
	}
	return self
}

func hx_init_haxe__ds__IntMap(self *haxe__ds__IntMap) *haxe__ds__IntMap {
	self.h = hxrt.IntMapNew()
	return self
}

func (self *haxe__ds__IntMap) set(key int, value any) {
	hxrt.IntMapSet(self.h, key, value)
}

func (self *haxe__ds__IntMap) get(key int) any {
	return hxrt.IntMapGet(self.h, key)
}

func (self *haxe__ds__IntMap) exists(key int) bool {
	return hxrt.IntMapExists(self.h, key)
}

func (self *haxe__ds__IntMap) remove(key int) bool {
	return hxrt.IntMapRemove(self.h, key)
}

func (self *haxe__ds__IntMap) keys() hxrt.Anonymous {
	keys := hxrt.IntMapKeys(self.h)
	index := 0
	hx_obj_1 := map[string]any{}
	hx_obj_1["hasNext"] = hxrt.NewFunction(func() bool {
		return (index < len(keys))
	})
	hx_obj_1["next"] = hxrt.NewFunction(func() int {
		hx_post_2 := index
		index = int(int32((index + 1)))
		return keys[hx_post_2]
	})
	return hxrt.AnonymousFromMap(hx_obj_1)
}

func (self *haxe__ds__IntMap) iterator() hxrt.Anonymous {
	_gthis := self
	keys := hxrt.IntMapKeys(self.h)
	index := 0
	hx_obj_3 := map[string]any{}
	hx_obj_3["hasNext"] = hxrt.NewFunction(func() bool {
		return (index < len(keys))
	})
	hx_obj_3["next"] = hxrt.NewFunction(func() any {
		return hxrt.IntMapGet(_gthis.h, keys[func() int {
			hx_post_4 := index
			index = int(int32((index + 1)))
			return hx_post_4
		}()])
	})
	return hxrt.AnonymousFromMap(hx_obj_3)
}

func (self *haxe__ds__IntMap) keyValueIterator() hxrt.Anonymous {
	_gthis := self
	keys := self.__hx_this.keys()
	hx_obj_5 := map[string]any{}
	hx_obj_5["hasNext"] = hxrt.NewFunction(func() bool {
		return func() bool {
			var hx_callable_6 *hxrt.Function[func() bool] = hx_recover_callable_1(hxrt.AnonymousGet(keys, "hasNext"))
			return hxrt.FunctionCall(hx_callable_6)()
		}()
	})
	hx_obj_5["next"] = hxrt.NewFunction(func() hxrt.Anonymous {
		key := func() int {
			var hx_callable_7 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(keys, "next"))
			return hxrt.FunctionCall(hx_callable_7)()
		}()
		hx_obj_8 := map[string]any{}
		hx_obj_8["key"] = key
		hx_obj_8["value"] = _gthis.__hx_this.get(key)
		return hxrt.AnonymousFromMap(hx_obj_8)
	})
	return hxrt.AnonymousFromMap(hx_obj_5)
}

func (self *haxe__ds__IntMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.IntFromNullableAny(func(hx_value_9 any) int {
		if hx_value_9 == nil {
			var hx_zero_10 int
			return hx_zero_10
		}
		return hx_value_9.(int)
	}(key)))
}

func (self *haxe__ds__IntMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.IntFromNullableAny(func(hx_value_11 any) int {
		if hx_value_11 == nil {
			var hx_zero_12 int
			return hx_zero_12
		}
		return hx_value_11.(int)
	}(key)), value)
}

func (self *haxe__ds__IntMap) existsIMap(key any) bool {
	return func(hx_value_15 any) bool {
		if hx_value_15 == nil {
			var hx_zero_16 bool
			return hx_zero_16
		}
		return hx_value_15.(bool)
	}(self.__hx_this.exists(hxrt.IntFromNullableAny(func(hx_value_13 any) int {
		if hx_value_13 == nil {
			var hx_zero_14 int
			return hx_zero_14
		}
		return hx_value_13.(int)
	}(key))))
}

func (self *haxe__ds__IntMap) removeIMap(key any) bool {
	return func(hx_value_19 any) bool {
		if hx_value_19 == nil {
			var hx_zero_20 bool
			return hx_zero_20
		}
		return hx_value_19.(bool)
	}(self.__hx_this.remove(hxrt.IntFromNullableAny(func(hx_value_17 any) int {
		if hx_value_17 == nil {
			var hx_zero_18 int
			return hx_zero_18
		}
		return hx_value_17.(int)
	}(key))))
}

func (self *haxe__ds__IntMap) copyIMap() haxe__IMap {
	return hx_recover_class_0(self.__hx_this.copy())
}

func (self *haxe__ds__IntMap) copy() *haxe__ds__IntMap {
	copied := New_haxe__ds__IntMap()
	key := self.__hx_this.keys()
	for func() bool {
		var hx_callable_21 *hxrt.Function[func() bool] = hx_recover_callable_1(hxrt.AnonymousGet(key, "hasNext"))
		return hxrt.FunctionCall(hx_callable_21)()
	}() {
		key_1 := func() int {
			var hx_callable_22 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(key, "next"))
			return hxrt.FunctionCall(hx_callable_22)()
		}()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__IntMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := self.__hx_this.keys()
	for func() bool {
		var hx_callable_23 *hxrt.Function[func() bool] = hx_recover_callable_1(hxrt.AnonymousGet(iterator, "hasNext"))
		return hxrt.FunctionCall(hx_callable_23)()
	}() {
		key := func() int {
			var hx_callable_24 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(iterator, "next"))
			return hxrt.FunctionCall(hx_callable_24)()
		}()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func() bool {
			var hx_callable_25 *hxrt.Function[func() bool] = hx_recover_callable_1(hxrt.AnonymousGet(iterator, "hasNext"))
			return hxrt.FunctionCall(hx_callable_25)()
		}() {
			out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(", "))
		}
	}
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("]"))
	return out_b
}

func (self *haxe__ds__IntMap) clear() {
	hxrt.IntMapClear(self.h)
}

func (self *haxe__ds__IntMap) String() string {
	return *self.__hx_this.toString()
}
