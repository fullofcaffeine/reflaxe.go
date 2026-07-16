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
	hx_obj_14 := map[string]any{}
	hx_obj_14["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_14["next"] = func() int {
		hx_post_15 := index
		index = int(int32((index + 1)))
		return keys[hx_post_15]
	}
	return hx_obj_14
}

func (self *haxe__ds__IntMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.IntMapKeys(self.h)
	index := 0
	hx_obj_16 := map[string]any{}
	hx_obj_16["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_16["next"] = func() any {
		return hxrt.IntMapGet(_gthis.h, keys[func() int {
			hx_post_17 := index
			index = int(int32((index + 1)))
			return hx_post_17
		}()])
	}
	return hx_obj_16
}

func (self *haxe__ds__IntMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_18 any) map[string]any {
		if hx_value_18 == nil {
			var hx_zero_19 map[string]any
			return hx_zero_19
		}
		return hx_value_18.(map[string]any)
	}(self.keys())
	hx_obj_20 := map[string]any{}
	hx_obj_20["hasNext"] = func() bool {
		return func(hx_obj_21 map[string]any) func() bool {
			hx_field_22 := hx_obj_21["hasNext"]
			if hx_field_22 == nil {
				var hx_zero_23 func() bool
				return hx_zero_23
			}
			return hx_field_22.(func() bool)
		}(keys)()
	}
	hx_obj_20["next"] = func() map[string]any {
		key := func(hx_obj_24 map[string]any) func() int {
			hx_field_25 := hx_obj_24["next"]
			if hx_field_25 == nil {
				var hx_zero_26 func() int
				return hx_zero_26
			}
			return hx_field_25.(func() int)
		}(keys)()
		hx_obj_27 := map[string]any{}
		hx_obj_27["key"] = key
		hx_obj_27["value"] = _gthis.get(key)
		return hx_obj_27
	}
	return hx_obj_20
}

func (self *haxe__ds__IntMap) getIMap(key any) any {
	return self.get(hxrt.IntFromNullableAny(func(hx_value_28 any) int {
		if hx_value_28 == nil {
			var hx_zero_29 int
			return hx_zero_29
		}
		return hx_value_28.(int)
	}(key)))
}

func (self *haxe__ds__IntMap) setIMap(key any, value any) {
	self.set(hxrt.IntFromNullableAny(func(hx_value_30 any) int {
		if hx_value_30 == nil {
			var hx_zero_31 int
			return hx_zero_31
		}
		return hx_value_30.(int)
	}(key)), value)
}

func (self *haxe__ds__IntMap) existsIMap(key any) bool {
	return func(hx_value_34 any) bool {
		if hx_value_34 == nil {
			var hx_zero_35 bool
			return hx_zero_35
		}
		return hx_value_34.(bool)
	}(self.exists(hxrt.IntFromNullableAny(func(hx_value_32 any) int {
		if hx_value_32 == nil {
			var hx_zero_33 int
			return hx_zero_33
		}
		return hx_value_32.(int)
	}(key))))
}

func (self *haxe__ds__IntMap) removeIMap(key any) bool {
	return func(hx_value_38 any) bool {
		if hx_value_38 == nil {
			var hx_zero_39 bool
			return hx_zero_39
		}
		return hx_value_38.(bool)
	}(self.remove(hxrt.IntFromNullableAny(func(hx_value_36 any) int {
		if hx_value_36 == nil {
			var hx_zero_37 int
			return hx_zero_37
		}
		return hx_value_36.(int)
	}(key))))
}

func (self *haxe__ds__IntMap) copyIMap() haxe__IMap {
	return func(hx_value_40 any) *haxe__ds__IntMap {
		if hx_value_40 == nil {
			var hx_zero_41 *haxe__ds__IntMap
			return hx_zero_41
		}
		return hx_value_40.(*haxe__ds__IntMap)
	}(self.copy())
}

func (self *haxe__ds__IntMap) copy() *haxe__ds__IntMap {
	copied := New_haxe__ds__IntMap()
	key := func(hx_value_42 any) map[string]any {
		if hx_value_42 == nil {
			var hx_zero_43 map[string]any
			return hx_zero_43
		}
		return hx_value_42.(map[string]any)
	}(self.keys())
	for func(hx_obj_44 map[string]any) func() bool {
		hx_field_45 := hx_obj_44["hasNext"]
		if hx_field_45 == nil {
			var hx_zero_46 func() bool
			return hx_zero_46
		}
		return hx_field_45.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_47 map[string]any) func() int {
			hx_field_48 := hx_obj_47["next"]
			if hx_field_48 == nil {
				var hx_zero_49 func() int
				return hx_zero_49
			}
			return hx_field_48.(func() int)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__IntMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_50 any) map[string]any {
		if hx_value_50 == nil {
			var hx_zero_51 map[string]any
			return hx_zero_51
		}
		return hx_value_50.(map[string]any)
	}(self.keys())
	for func(hx_obj_52 map[string]any) func() bool {
		hx_field_53 := hx_obj_52["hasNext"]
		if hx_field_53 == nil {
			var hx_zero_54 func() bool
			return hx_zero_54
		}
		return hx_field_53.(func() bool)
	}(iterator)() {
		key := func(hx_obj_55 map[string]any) func() int {
			hx_field_56 := hx_obj_55["next"]
			if hx_field_56 == nil {
				var hx_zero_57 func() int
				return hx_zero_57
			}
			return hx_field_56.(func() int)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_58 map[string]any) func() bool {
			hx_field_59 := hx_obj_58["hasNext"]
			if hx_field_59 == nil {
				var hx_zero_60 func() bool
				return hx_zero_60
			}
			return hx_field_59.(func() bool)
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
