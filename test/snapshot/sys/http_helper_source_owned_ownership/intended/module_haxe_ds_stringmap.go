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
	hx_obj_13 := map[string]any{}
	hx_obj_13["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_13["next"] = func() *string {
		hx_post_14 := index
		index = int(int32((index + 1)))
		return keys[hx_post_14]
	}
	return hx_obj_13
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_15 := map[string]any{}
	hx_obj_15["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_15["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_16 := index
			index = int(int32((index + 1)))
			return hx_post_16
		}()])
	}
	return hx_obj_15
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_17 any) map[string]any {
		if hx_value_17 == nil {
			var hx_zero_18 map[string]any
			return hx_zero_18
		}
		return hx_value_17.(map[string]any)
	}(self.keys())
	hx_obj_19 := map[string]any{}
	hx_obj_19["hasNext"] = func() bool {
		return func(hx_obj_20 map[string]any) func() bool {
			hx_field_21 := hx_obj_20["hasNext"]
			if hx_field_21 == nil {
				var hx_zero_22 func() bool
				return hx_zero_22
			}
			return hx_field_21.(func() bool)
		}(keys)()
	}
	hx_obj_19["next"] = func() map[string]any {
		key := func(hx_obj_23 map[string]any) func() *string {
			hx_field_24 := hx_obj_23["next"]
			if hx_field_24 == nil {
				var hx_zero_25 func() *string
				return hx_zero_25
			}
			return hx_field_24.(func() *string)
		}(keys)()
		hx_obj_26 := map[string]any{}
		hx_obj_26["key"] = key
		hx_obj_26["value"] = _gthis.get(key)
		return hx_obj_26
	}
	return hx_obj_19
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_27 any) *string {
		if hx_value_27 == nil {
			var hx_zero_28 *string
			return hx_zero_28
		}
		return hx_value_27.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_29 any) *string {
		if hx_value_29 == nil {
			var hx_zero_30 *string
			return hx_zero_30
		}
		return hx_value_29.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_33 any) bool {
		if hx_value_33 == nil {
			var hx_zero_34 bool
			return hx_zero_34
		}
		return hx_value_33.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_31 any) *string {
		if hx_value_31 == nil {
			var hx_zero_32 *string
			return hx_zero_32
		}
		return hx_value_31.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_37 any) bool {
		if hx_value_37 == nil {
			var hx_zero_38 bool
			return hx_zero_38
		}
		return hx_value_37.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_35 any) *string {
		if hx_value_35 == nil {
			var hx_zero_36 *string
			return hx_zero_36
		}
		return hx_value_35.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_39 any) *haxe__ds__StringMap {
		if hx_value_39 == nil {
			var hx_zero_40 *haxe__ds__StringMap
			return hx_zero_40
		}
		return hx_value_39.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_41 any) map[string]any {
		if hx_value_41 == nil {
			var hx_zero_42 map[string]any
			return hx_zero_42
		}
		return hx_value_41.(map[string]any)
	}(self.keys())
	for func(hx_obj_43 map[string]any) func() bool {
		hx_field_44 := hx_obj_43["hasNext"]
		if hx_field_44 == nil {
			var hx_zero_45 func() bool
			return hx_zero_45
		}
		return hx_field_44.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_46 map[string]any) func() *string {
			hx_field_47 := hx_obj_46["next"]
			if hx_field_47 == nil {
				var hx_zero_48 func() *string
				return hx_zero_48
			}
			return hx_field_47.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_49 any) map[string]any {
		if hx_value_49 == nil {
			var hx_zero_50 map[string]any
			return hx_zero_50
		}
		return hx_value_49.(map[string]any)
	}(self.keys())
	for func(hx_obj_51 map[string]any) func() bool {
		hx_field_52 := hx_obj_51["hasNext"]
		if hx_field_52 == nil {
			var hx_zero_53 func() bool
			return hx_zero_53
		}
		return hx_field_52.(func() bool)
	}(iterator)() {
		key := func(hx_obj_54 map[string]any) func() *string {
			hx_field_55 := hx_obj_54["next"]
			if hx_field_55 == nil {
				var hx_zero_56 func() *string
				return hx_zero_56
			}
			return hx_field_55.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_57 map[string]any) func() bool {
			hx_field_58 := hx_obj_57["hasNext"]
			if hx_field_58 == nil {
				var hx_zero_59 func() bool
				return hx_zero_59
			}
			return hx_field_58.(func() bool)
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
