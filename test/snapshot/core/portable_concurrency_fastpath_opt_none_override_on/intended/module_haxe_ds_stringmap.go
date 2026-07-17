package main

import "snapshot/hxrt"

type I_haxe__ds__StringMap interface {
	set(key *string, value any)
	get(key *string) any
	exists(key *string) bool
	remove(key *string) bool
	keys() map[string]any
	iterator() map[string]any
	keyValueIterator() map[string]any
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

func New_haxe__ds__StringMap() *haxe__ds__StringMap {
	self := &haxe__ds__StringMap{}
	self.__hx_this = self
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

func (self *haxe__ds__StringMap) keys() map[string]any {
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_15 := map[string]any{}
	hx_obj_15["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_15["next"] = func() *string {
		hx_post_16 := index
		index = int(int32((index + 1)))
		return keys[hx_post_16]
	}
	return hx_obj_15
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_17 := map[string]any{}
	hx_obj_17["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_17["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_18 := index
			index = int(int32((index + 1)))
			return hx_post_18
		}()])
	}
	return hx_obj_17
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_19 any) map[string]any {
		if hx_value_19 == nil {
			var hx_zero_20 map[string]any
			return hx_zero_20
		}
		return hx_value_19.(map[string]any)
	}(self.keys())
	hx_obj_21 := map[string]any{}
	hx_obj_21["hasNext"] = func() bool {
		return func(hx_obj_22 map[string]any) func() bool {
			hx_field_23 := hx_obj_22["hasNext"]
			if hx_field_23 == nil {
				var hx_zero_24 func() bool
				return hx_zero_24
			}
			return hx_field_23.(func() bool)
		}(keys)()
	}
	hx_obj_21["next"] = func() map[string]any {
		key := func(hx_obj_25 map[string]any) func() *string {
			hx_field_26 := hx_obj_25["next"]
			if hx_field_26 == nil {
				var hx_zero_27 func() *string
				return hx_zero_27
			}
			return hx_field_26.(func() *string)
		}(keys)()
		hx_obj_28 := map[string]any{}
		hx_obj_28["key"] = key
		hx_obj_28["value"] = _gthis.get(key)
		return hx_obj_28
	}
	return hx_obj_21
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_29 any) *string {
		if hx_value_29 == nil {
			var hx_zero_30 *string
			return hx_zero_30
		}
		return hx_value_29.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_31 any) *string {
		if hx_value_31 == nil {
			var hx_zero_32 *string
			return hx_zero_32
		}
		return hx_value_31.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_35 any) bool {
		if hx_value_35 == nil {
			var hx_zero_36 bool
			return hx_zero_36
		}
		return hx_value_35.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_33 any) *string {
		if hx_value_33 == nil {
			var hx_zero_34 *string
			return hx_zero_34
		}
		return hx_value_33.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_39 any) bool {
		if hx_value_39 == nil {
			var hx_zero_40 bool
			return hx_zero_40
		}
		return hx_value_39.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_37 any) *string {
		if hx_value_37 == nil {
			var hx_zero_38 *string
			return hx_zero_38
		}
		return hx_value_37.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_41 any) *haxe__ds__StringMap {
		if hx_value_41 == nil {
			var hx_zero_42 *haxe__ds__StringMap
			return hx_zero_42
		}
		return hx_value_41.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_43 any) map[string]any {
		if hx_value_43 == nil {
			var hx_zero_44 map[string]any
			return hx_zero_44
		}
		return hx_value_43.(map[string]any)
	}(self.keys())
	for func(hx_obj_45 map[string]any) func() bool {
		hx_field_46 := hx_obj_45["hasNext"]
		if hx_field_46 == nil {
			var hx_zero_47 func() bool
			return hx_zero_47
		}
		return hx_field_46.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_48 map[string]any) func() *string {
			hx_field_49 := hx_obj_48["next"]
			if hx_field_49 == nil {
				var hx_zero_50 func() *string
				return hx_zero_50
			}
			return hx_field_49.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_51 any) map[string]any {
		if hx_value_51 == nil {
			var hx_zero_52 map[string]any
			return hx_zero_52
		}
		return hx_value_51.(map[string]any)
	}(self.keys())
	for func(hx_obj_53 map[string]any) func() bool {
		hx_field_54 := hx_obj_53["hasNext"]
		if hx_field_54 == nil {
			var hx_zero_55 func() bool
			return hx_zero_55
		}
		return hx_field_54.(func() bool)
	}(iterator)() {
		key := func(hx_obj_56 map[string]any) func() *string {
			hx_field_57 := hx_obj_56["next"]
			if hx_field_57 == nil {
				var hx_zero_58 func() *string
				return hx_zero_58
			}
			return hx_field_57.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_59 map[string]any) func() bool {
			hx_field_60 := hx_obj_59["hasNext"]
			if hx_field_60 == nil {
				var hx_zero_61 func() bool
				return hx_zero_61
			}
			return hx_field_60.(func() bool)
		}(iterator)() {
			out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(", "))
		}
	}
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("]"))
	return out_b
}

func (self *haxe__ds__StringMap) clear() {
	hxrt.StringMapClear(self.h)
}
