package main

import "snapshot/hxrt"

type I_haxe__ds__ObjectMap interface {
	set(key any, value any)
	get(key any) any
	exists(key any) bool
	remove(key any) bool
	keys() hxrt.Anonymous
	iterator() hxrt.Anonymous
	keyValueIterator() hxrt.Anonymous
	getIMap(key any) any
	setIMap(key any, value any)
	existsIMap(key any) bool
	removeIMap(key any) bool
	copyIMap() haxe__IMap
	copy() *haxe__ds__ObjectMap
	toString() *string
	clear()
}

type haxe__ds__ObjectMap struct {
	__hx_this I_haxe__ds__ObjectMap
	h         *hxrt.ObjectMapCell
}

func (self *haxe__ds__ObjectMap) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe__ds__ObjectMap() *haxe__ds__ObjectMap {
	return hx_init_haxe__ds__ObjectMap(hx_alloc_haxe__ds__ObjectMap(nil))
}

func hx_alloc_haxe__ds__ObjectMap(owner any) *haxe__ds__ObjectMap {
	self := &haxe__ds__ObjectMap{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_haxe__ds__ObjectMap)
	}
	return self
}

func hx_init_haxe__ds__ObjectMap(self *haxe__ds__ObjectMap) *haxe__ds__ObjectMap {
	self.h = hxrt.ObjectMapNew()
	return self
}

func (self *haxe__ds__ObjectMap) set(key any, value any) {
	hxrt.ObjectMapSet(self.h, key, value)
}

func (self *haxe__ds__ObjectMap) get(key any) any {
	return hxrt.ObjectMapGet(self.h, key)
}

func (self *haxe__ds__ObjectMap) exists(key any) bool {
	return hxrt.ObjectMapExists(self.h, key)
}

func (self *haxe__ds__ObjectMap) remove(key any) bool {
	return hxrt.ObjectMapRemove(self.h, key)
}

func (self *haxe__ds__ObjectMap) keys() hxrt.Anonymous {
	keys := hxrt.ObjectMapKeys(self.h)
	index := 0
	hx_obj_1 := map[string]any{}
	hx_obj_1["hasNext"] = hxrt.NewFunction(func() bool {
		return (index < len(keys))
	})
	hx_obj_1["next"] = hxrt.NewFunction(func() any {
		return keys[func() int {
			hx_post_2 := index
			index = int(int32((index + 1)))
			return hx_post_2
		}()]
	})
	return hxrt.AnonymousFromMap(hx_obj_1)
}

func (self *haxe__ds__ObjectMap) iterator() hxrt.Anonymous {
	_gthis := self
	keys := hxrt.ObjectMapKeys(self.h)
	index := 0
	hx_obj_3 := map[string]any{}
	hx_obj_3["hasNext"] = hxrt.NewFunction(func() bool {
		return (index < len(keys))
	})
	hx_obj_3["next"] = hxrt.NewFunction(func() any {
		return hxrt.ObjectMapGet(_gthis.h, keys[func() int {
			hx_post_4 := index
			index = int(int32((index + 1)))
			return hx_post_4
		}()])
	})
	return hxrt.AnonymousFromMap(hx_obj_3)
}

func (self *haxe__ds__ObjectMap) keyValueIterator() hxrt.Anonymous {
	_gthis := self
	keys := self.__hx_this.keys()
	hx_obj_5 := map[string]any{}
	hx_obj_5["hasNext"] = hxrt.NewFunction(func() bool {
		return func() bool {
			var hx_callable_6 *hxrt.Function[func() bool] = hx_recover_callable_2(hxrt.AnonymousGet(keys, "hasNext"))
			return hxrt.FunctionCall(hx_callable_6)()
		}()
	})
	hx_obj_5["next"] = hxrt.NewFunction(func() hxrt.Anonymous {
		var key any = func() any {
			var hx_callable_7 *hxrt.Function[func() any] = hx_recover_callable_3(hxrt.AnonymousGet(keys, "next"))
			return hxrt.FunctionCall(hx_callable_7)()
		}()
		hx_obj_8 := map[string]any{}
		hx_obj_8["key"] = key
		hx_obj_8["value"] = _gthis.__hx_this.get(key)
		return hxrt.AnonymousFromMap(hx_obj_8)
	})
	return hxrt.AnonymousFromMap(hx_obj_5)
}

func (self *haxe__ds__ObjectMap) getIMap(key any) any {
	return self.__hx_this.get(key)
}

func (self *haxe__ds__ObjectMap) setIMap(key any, value any) {
	self.__hx_this.set(key, value)
}

func (self *haxe__ds__ObjectMap) existsIMap(key any) bool {
	return func(hx_value_9 any) bool {
		if hx_value_9 == nil {
			var hx_zero_10 bool
			return hx_zero_10
		}
		return hx_value_9.(bool)
	}(self.__hx_this.exists(key))
}

func (self *haxe__ds__ObjectMap) removeIMap(key any) bool {
	return func(hx_value_11 any) bool {
		if hx_value_11 == nil {
			var hx_zero_12 bool
			return hx_zero_12
		}
		return hx_value_11.(bool)
	}(self.__hx_this.remove(key))
}

func (self *haxe__ds__ObjectMap) copyIMap() haxe__IMap {
	return hx_recover_class_0(self.__hx_this.copy())
}

func (self *haxe__ds__ObjectMap) copy() *haxe__ds__ObjectMap {
	copied := New_haxe__ds__ObjectMap()
	key := self.__hx_this.keys()
	for func() bool {
		var hx_callable_13 *hxrt.Function[func() bool] = hx_recover_callable_2(hxrt.AnonymousGet(key, "hasNext"))
		return hxrt.FunctionCall(hx_callable_13)()
	}() {
		var key_1 any = func() any {
			var hx_callable_14 *hxrt.Function[func() any] = hx_recover_callable_3(hxrt.AnonymousGet(key, "next"))
			return hxrt.FunctionCall(hx_callable_14)()
		}()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__ObjectMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := self.__hx_this.keys()
	for func() bool {
		var hx_callable_15 *hxrt.Function[func() bool] = hx_recover_callable_2(hxrt.AnonymousGet(iterator, "hasNext"))
		return hxrt.FunctionCall(hx_callable_15)()
	}() {
		var key any = func() any {
			var hx_callable_16 *hxrt.Function[func() any] = hx_recover_callable_3(hxrt.AnonymousGet(iterator, "next"))
			return hxrt.FunctionCall(hx_callable_16)()
		}()
		x := hxrt.StdString(key)
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x_1 := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x_1))
		if func() bool {
			var hx_callable_17 *hxrt.Function[func() bool] = hx_recover_callable_2(hxrt.AnonymousGet(iterator, "hasNext"))
			return hxrt.FunctionCall(hx_callable_17)()
		}() {
			out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(", "))
		}
	}
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("]"))
	return out_b
}

func (self *haxe__ds__ObjectMap) clear() {
	hxrt.ObjectMapClear(self.h)
}

func (self *haxe__ds__ObjectMap) String() string {
	return *self.__hx_this.toString()
}
