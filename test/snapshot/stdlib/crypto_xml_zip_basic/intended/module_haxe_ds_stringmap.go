package main

import "snapshot/hxrt"

type I_haxe__ds__StringMap interface {
	set(key *string, value any)
	get(key *string) any
	exists(key *string) bool
	remove(key *string) bool
	keys() hxrt.Anonymous
	iterator() hxrt.Anonymous
	keyValueIterator() hxrt.Anonymous
	getIMap(key any) any
	setIMap(key any, value any)
	existsIMap(key any) bool
	removeIMap(key any) bool
	copyIMap() haxe__IMap
	copy() *haxe__ds__StringMap
	toString() *string
	clear()
}

type haxe__ds__StringMap struct {
	__hx_this I_haxe__ds__StringMap
	h         *hxrt.StringMapCell
}

func (self *haxe__ds__StringMap) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe__ds__StringMap() *haxe__ds__StringMap {
	return hx_init_haxe__ds__StringMap(hx_alloc_haxe__ds__StringMap(nil))
}

func hx_alloc_haxe__ds__StringMap(owner any) *haxe__ds__StringMap {
	self := &haxe__ds__StringMap{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_haxe__ds__StringMap)
	}
	return self
}

func hx_init_haxe__ds__StringMap(self *haxe__ds__StringMap) *haxe__ds__StringMap {
	self.h = hxrt.StringMapNew()
	return self
}

func (self *haxe__ds__StringMap) set(key *string, value any) {
	hxrt.StringMapSet(self.h, key, value)
}

func (self *haxe__ds__StringMap) get(key *string) any {
	return hxrt.StringMapGet(self.h, key)
}

func (self *haxe__ds__StringMap) exists(key *string) bool {
	return hxrt.StringMapExists(self.h, key)
}

func (self *haxe__ds__StringMap) remove(key *string) bool {
	return hxrt.StringMapRemove(self.h, key)
}

func (self *haxe__ds__StringMap) keys() hxrt.Anonymous {
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_1 := map[string]any{}
	hx_obj_1["hasNext"] = hxrt.NewFunction(func() bool {
		return (index < len(keys))
	})
	hx_obj_1["next"] = hxrt.NewFunction(func() *string {
		hx_post_2 := index
		index = int(int32((index + 1)))
		return keys[hx_post_2]
	})
	return hxrt.AnonymousFromMap(hx_obj_1)
}

func (self *haxe__ds__StringMap) iterator() hxrt.Anonymous {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_3 := map[string]any{}
	hx_obj_3["hasNext"] = hxrt.NewFunction(func() bool {
		return (index < len(keys))
	})
	hx_obj_3["next"] = hxrt.NewFunction(func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_4 := index
			index = int(int32((index + 1)))
			return hx_post_4
		}()])
	})
	return hxrt.AnonymousFromMap(hx_obj_3)
}

func (self *haxe__ds__StringMap) keyValueIterator() hxrt.Anonymous {
	_gthis := self
	keys := self.__hx_this.keys()
	hx_obj_5 := map[string]any{}
	hx_obj_5["hasNext"] = hxrt.NewFunction(func() bool {
		return func() bool {
			var hx_callable_6 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(keys, "hasNext"))
			return hxrt.FunctionCall(hx_callable_6)()
		}()
	})
	hx_obj_5["next"] = hxrt.NewFunction(func() hxrt.Anonymous {
		key := func() *string {
			var hx_callable_7 *hxrt.Function[func() *string] = hx_recover_callable_1(hxrt.AnonymousGet(keys, "next"))
			return hxrt.FunctionCall(hx_callable_7)()
		}()
		hx_obj_8 := map[string]any{}
		hx_obj_8["key"] = key
		hx_obj_8["value"] = _gthis.__hx_this.get(key)
		return hxrt.AnonymousFromMap(hx_obj_8)
	})
	return hxrt.AnonymousFromMap(hx_obj_5)
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_9 any) *string {
		if hx_value_9 == nil {
			var hx_zero_10 *string
			return hx_zero_10
		}
		return hx_value_9.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_11 any) *string {
		if hx_value_11 == nil {
			var hx_zero_12 *string
			return hx_zero_12
		}
		return hx_value_11.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_15 any) bool {
		if hx_value_15 == nil {
			var hx_zero_16 bool
			return hx_zero_16
		}
		return hx_value_15.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_13 any) *string {
		if hx_value_13 == nil {
			var hx_zero_14 *string
			return hx_zero_14
		}
		return hx_value_13.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_19 any) bool {
		if hx_value_19 == nil {
			var hx_zero_20 bool
			return hx_zero_20
		}
		return hx_value_19.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_17 any) *string {
		if hx_value_17 == nil {
			var hx_zero_18 *string
			return hx_zero_18
		}
		return hx_value_17.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return hx_recover_class_3(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := self.__hx_this.keys()
	for func() bool {
		var hx_callable_21 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(key, "hasNext"))
		return hxrt.FunctionCall(hx_callable_21)()
	}() {
		key_1 := func() *string {
			var hx_callable_22 *hxrt.Function[func() *string] = hx_recover_callable_1(hxrt.AnonymousGet(key, "next"))
			return hxrt.FunctionCall(hx_callable_22)()
		}()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := self.__hx_this.keys()
	for func() bool {
		var hx_callable_23 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(iterator, "hasNext"))
		return hxrt.FunctionCall(hx_callable_23)()
	}() {
		key := func() *string {
			var hx_callable_24 *hxrt.Function[func() *string] = hx_recover_callable_1(hxrt.AnonymousGet(iterator, "next"))
			return hxrt.FunctionCall(hx_callable_24)()
		}()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func() bool {
			var hx_callable_25 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(iterator, "hasNext"))
			return hxrt.FunctionCall(hx_callable_25)()
		}() {
			out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(", "))
		}
	}
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("]"))
	return out_b
}

func (self *haxe__ds__StringMap) clear() {
	hxrt.StringMapClear(self.h)
}

func (self *haxe__ds__StringMap) String() string {
	return *self.__hx_this.toString()
}
