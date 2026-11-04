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
	self := &haxe__ds__IntMap{}
	self.__hx_this = self
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
			var hx_callable_9 *hxrt.Function[func() bool] = func(hx_obj_6 hxrt.Anonymous) *hxrt.Function[func() bool] {
				hx_field_7 := hxrt.AnonymousGet(hx_obj_6, "hasNext")
				if hx_field_7 == nil {
					var hx_zero_8 *hxrt.Function[func() bool]
					return hx_zero_8
				}
				return hx_field_7.(*hxrt.Function[func() bool])
			}(keys)
			return hxrt.FunctionCall(hx_callable_9)()
		}()
	})
	hx_obj_5["next"] = hxrt.NewFunction(func() hxrt.Anonymous {
		key := func() int {
			var hx_callable_13 *hxrt.Function[func() int] = func(hx_obj_10 hxrt.Anonymous) *hxrt.Function[func() int] {
				hx_field_11 := hxrt.AnonymousGet(hx_obj_10, "next")
				if hx_field_11 == nil {
					var hx_zero_12 *hxrt.Function[func() int]
					return hx_zero_12
				}
				return hx_field_11.(*hxrt.Function[func() int])
			}(keys)
			return hxrt.FunctionCall(hx_callable_13)()
		}()
		hx_obj_14 := map[string]any{}
		hx_obj_14["key"] = key
		hx_obj_14["value"] = _gthis.__hx_this.get(key)
		return hxrt.AnonymousFromMap(hx_obj_14)
	})
	return hxrt.AnonymousFromMap(hx_obj_5)
}

func (self *haxe__ds__IntMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.IntFromNullableAny(func(hx_value_15 any) int {
		if hx_value_15 == nil {
			var hx_zero_16 int
			return hx_zero_16
		}
		return hx_value_15.(int)
	}(key)))
}

func (self *haxe__ds__IntMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.IntFromNullableAny(func(hx_value_17 any) int {
		if hx_value_17 == nil {
			var hx_zero_18 int
			return hx_zero_18
		}
		return hx_value_17.(int)
	}(key)), value)
}

func (self *haxe__ds__IntMap) existsIMap(key any) bool {
	return func(hx_value_21 any) bool {
		if hx_value_21 == nil {
			var hx_zero_22 bool
			return hx_zero_22
		}
		return hx_value_21.(bool)
	}(self.__hx_this.exists(hxrt.IntFromNullableAny(func(hx_value_19 any) int {
		if hx_value_19 == nil {
			var hx_zero_20 int
			return hx_zero_20
		}
		return hx_value_19.(int)
	}(key))))
}

func (self *haxe__ds__IntMap) removeIMap(key any) bool {
	return func(hx_value_25 any) bool {
		if hx_value_25 == nil {
			var hx_zero_26 bool
			return hx_zero_26
		}
		return hx_value_25.(bool)
	}(self.__hx_this.remove(hxrt.IntFromNullableAny(func(hx_value_23 any) int {
		if hx_value_23 == nil {
			var hx_zero_24 int
			return hx_zero_24
		}
		return hx_value_23.(int)
	}(key))))
}

func (self *haxe__ds__IntMap) copyIMap() haxe__IMap {
	return func(hx_value_27 any) *haxe__ds__IntMap {
		if hx_value_27 == nil {
			var hx_zero_28 *haxe__ds__IntMap
			return hx_zero_28
		}
		return hx_value_27.(*haxe__ds__IntMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__IntMap) copy() *haxe__ds__IntMap {
	copied := New_haxe__ds__IntMap()
	key := self.__hx_this.keys()
	for func() bool {
		var hx_callable_32 *hxrt.Function[func() bool] = func(hx_obj_29 hxrt.Anonymous) *hxrt.Function[func() bool] {
			hx_field_30 := hxrt.AnonymousGet(hx_obj_29, "hasNext")
			if hx_field_30 == nil {
				var hx_zero_31 *hxrt.Function[func() bool]
				return hx_zero_31
			}
			return hx_field_30.(*hxrt.Function[func() bool])
		}(key)
		return hxrt.FunctionCall(hx_callable_32)()
	}() {
		key_1 := func() int {
			var hx_callable_36 *hxrt.Function[func() int] = func(hx_obj_33 hxrt.Anonymous) *hxrt.Function[func() int] {
				hx_field_34 := hxrt.AnonymousGet(hx_obj_33, "next")
				if hx_field_34 == nil {
					var hx_zero_35 *hxrt.Function[func() int]
					return hx_zero_35
				}
				return hx_field_34.(*hxrt.Function[func() int])
			}(key)
			return hxrt.FunctionCall(hx_callable_36)()
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
		var hx_callable_40 *hxrt.Function[func() bool] = func(hx_obj_37 hxrt.Anonymous) *hxrt.Function[func() bool] {
			hx_field_38 := hxrt.AnonymousGet(hx_obj_37, "hasNext")
			if hx_field_38 == nil {
				var hx_zero_39 *hxrt.Function[func() bool]
				return hx_zero_39
			}
			return hx_field_38.(*hxrt.Function[func() bool])
		}(iterator)
		return hxrt.FunctionCall(hx_callable_40)()
	}() {
		key := func() int {
			var hx_callable_44 *hxrt.Function[func() int] = func(hx_obj_41 hxrt.Anonymous) *hxrt.Function[func() int] {
				hx_field_42 := hxrt.AnonymousGet(hx_obj_41, "next")
				if hx_field_42 == nil {
					var hx_zero_43 *hxrt.Function[func() int]
					return hx_zero_43
				}
				return hx_field_42.(*hxrt.Function[func() int])
			}(iterator)
			return hxrt.FunctionCall(hx_callable_44)()
		}()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func() bool {
			var hx_callable_48 *hxrt.Function[func() bool] = func(hx_obj_45 hxrt.Anonymous) *hxrt.Function[func() bool] {
				hx_field_46 := hxrt.AnonymousGet(hx_obj_45, "hasNext")
				if hx_field_46 == nil {
					var hx_zero_47 *hxrt.Function[func() bool]
					return hx_zero_47
				}
				return hx_field_46.(*hxrt.Function[func() bool])
			}(iterator)
			return hxrt.FunctionCall(hx_callable_48)()
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
