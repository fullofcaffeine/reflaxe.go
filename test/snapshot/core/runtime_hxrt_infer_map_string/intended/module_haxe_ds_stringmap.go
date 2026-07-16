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
	hx_obj_12 := map[string]any{}
	hx_obj_12["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_12["next"] = func() *string {
		hx_post_13 := index
		index = int(int32((index + 1)))
		return keys[hx_post_13]
	}
	return hx_obj_12
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_14 := map[string]any{}
	hx_obj_14["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_14["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_15 := index
			index = int(int32((index + 1)))
			return hx_post_15
		}()])
	}
	return hx_obj_14
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_16 any) map[string]any {
		if hx_value_16 == nil {
			var hx_zero_17 map[string]any
			return hx_zero_17
		}
		return hx_value_16.(map[string]any)
	}(self.keys())
	hx_obj_18 := map[string]any{}
	hx_obj_18["hasNext"] = func() bool {
		return func(hx_obj_19 map[string]any) func() bool {
			hx_field_20 := hx_obj_19["hasNext"]
			if hx_field_20 == nil {
				var hx_zero_21 func() bool
				return hx_zero_21
			}
			return hx_field_20.(func() bool)
		}(keys)()
	}
	hx_obj_18["next"] = func() map[string]any {
		key := func(hx_obj_22 map[string]any) func() *string {
			hx_field_23 := hx_obj_22["next"]
			if hx_field_23 == nil {
				var hx_zero_24 func() *string
				return hx_zero_24
			}
			return hx_field_23.(func() *string)
		}(keys)()
		hx_obj_25 := map[string]any{}
		hx_obj_25["key"] = key
		hx_obj_25["value"] = _gthis.get(key)
		return hx_obj_25
	}
	return hx_obj_18
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_26 any) *string {
		if hx_value_26 == nil {
			var hx_zero_27 *string
			return hx_zero_27
		}
		return hx_value_26.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_28 any) *string {
		if hx_value_28 == nil {
			var hx_zero_29 *string
			return hx_zero_29
		}
		return hx_value_28.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_32 any) bool {
		if hx_value_32 == nil {
			var hx_zero_33 bool
			return hx_zero_33
		}
		return hx_value_32.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_30 any) *string {
		if hx_value_30 == nil {
			var hx_zero_31 *string
			return hx_zero_31
		}
		return hx_value_30.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_36 any) bool {
		if hx_value_36 == nil {
			var hx_zero_37 bool
			return hx_zero_37
		}
		return hx_value_36.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_34 any) *string {
		if hx_value_34 == nil {
			var hx_zero_35 *string
			return hx_zero_35
		}
		return hx_value_34.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_38 any) *haxe__ds__StringMap {
		if hx_value_38 == nil {
			var hx_zero_39 *haxe__ds__StringMap
			return hx_zero_39
		}
		return hx_value_38.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_40 any) map[string]any {
		if hx_value_40 == nil {
			var hx_zero_41 map[string]any
			return hx_zero_41
		}
		return hx_value_40.(map[string]any)
	}(self.keys())
	for func(hx_obj_42 map[string]any) func() bool {
		hx_field_43 := hx_obj_42["hasNext"]
		if hx_field_43 == nil {
			var hx_zero_44 func() bool
			return hx_zero_44
		}
		return hx_field_43.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_45 map[string]any) func() *string {
			hx_field_46 := hx_obj_45["next"]
			if hx_field_46 == nil {
				var hx_zero_47 func() *string
				return hx_zero_47
			}
			return hx_field_46.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_48 any) map[string]any {
		if hx_value_48 == nil {
			var hx_zero_49 map[string]any
			return hx_zero_49
		}
		return hx_value_48.(map[string]any)
	}(self.keys())
	for func(hx_obj_50 map[string]any) func() bool {
		hx_field_51 := hx_obj_50["hasNext"]
		if hx_field_51 == nil {
			var hx_zero_52 func() bool
			return hx_zero_52
		}
		return hx_field_51.(func() bool)
	}(iterator)() {
		key := func(hx_obj_53 map[string]any) func() *string {
			hx_field_54 := hx_obj_53["next"]
			if hx_field_54 == nil {
				var hx_zero_55 func() *string
				return hx_zero_55
			}
			return hx_field_54.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_56 map[string]any) func() bool {
			hx_field_57 := hx_obj_56["hasNext"]
			if hx_field_57 == nil {
				var hx_zero_58 func() bool
				return hx_zero_58
			}
			return hx_field_57.(func() bool)
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
