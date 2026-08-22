package main

import "snapshot/hxrt"

type I_haxe__ds__IntMap interface {
	set(key int, value any)
	get(key int) any
	exists(key int) bool
	remove(key int) bool
	keys() map[string]any
	iterator() map[string]any
	keyValueIterator() map[string]any
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

func (self *haxe__ds__IntMap) keys() map[string]any {
	keys := hxrt.IntMapKeys(self.h)
	index := 0
	hx_obj_1 := map[string]any{}
	hx_obj_1["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_1["next"] = func() int {
		hx_post_2 := index
		index = int(int32((index + 1)))
		return keys[hx_post_2]
	}
	return hx_obj_1
}

func (self *haxe__ds__IntMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.IntMapKeys(self.h)
	index := 0
	hx_obj_3 := map[string]any{}
	hx_obj_3["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_3["next"] = func() any {
		return hxrt.IntMapGet(_gthis.h, keys[func() int {
			hx_post_4 := index
			index = int(int32((index + 1)))
			return hx_post_4
		}()])
	}
	return hx_obj_3
}

func (self *haxe__ds__IntMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_5 any) map[string]any {
		if hx_value_5 == nil {
			var hx_zero_6 map[string]any
			return hx_zero_6
		}
		return hx_value_5.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_7 := map[string]any{}
	hx_obj_7["hasNext"] = func() bool {
		return func(hx_obj_8 map[string]any) func() bool {
			hx_field_9 := hx_obj_8["hasNext"]
			if hx_field_9 == nil {
				var hx_zero_10 func() bool
				return hx_zero_10
			}
			return hx_field_9.(func() bool)
		}(keys)()
	}
	hx_obj_7["next"] = func() map[string]any {
		key := func(hx_obj_11 map[string]any) func() int {
			hx_field_12 := hx_obj_11["next"]
			if hx_field_12 == nil {
				var hx_zero_13 func() int
				return hx_zero_13
			}
			return hx_field_12.(func() int)
		}(keys)()
		hx_obj_14 := map[string]any{}
		hx_obj_14["key"] = key
		hx_obj_14["value"] = _gthis.__hx_this.get(key)
		return hx_obj_14
	}
	return hx_obj_7
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
	key := func(hx_value_29 any) map[string]any {
		if hx_value_29 == nil {
			var hx_zero_30 map[string]any
			return hx_zero_30
		}
		return hx_value_29.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_31 map[string]any) func() bool {
		hx_field_32 := hx_obj_31["hasNext"]
		if hx_field_32 == nil {
			var hx_zero_33 func() bool
			return hx_zero_33
		}
		return hx_field_32.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_34 map[string]any) func() int {
			hx_field_35 := hx_obj_34["next"]
			if hx_field_35 == nil {
				var hx_zero_36 func() int
				return hx_zero_36
			}
			return hx_field_35.(func() int)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__IntMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_37 any) map[string]any {
		if hx_value_37 == nil {
			var hx_zero_38 map[string]any
			return hx_zero_38
		}
		return hx_value_37.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_39 map[string]any) func() bool {
		hx_field_40 := hx_obj_39["hasNext"]
		if hx_field_40 == nil {
			var hx_zero_41 func() bool
			return hx_zero_41
		}
		return hx_field_40.(func() bool)
	}(iterator)() {
		key := func(hx_obj_42 map[string]any) func() int {
			hx_field_43 := hx_obj_42["next"]
			if hx_field_43 == nil {
				var hx_zero_44 func() int
				return hx_zero_44
			}
			return hx_field_43.(func() int)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_45 map[string]any) func() bool {
			hx_field_46 := hx_obj_45["hasNext"]
			if hx_field_46 == nil {
				var hx_zero_47 func() bool
				return hx_zero_47
			}
			return hx_field_46.(func() bool)
		}(iterator)() {
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
