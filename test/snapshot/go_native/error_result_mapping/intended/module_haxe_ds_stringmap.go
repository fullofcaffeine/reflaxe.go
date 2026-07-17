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
	hx_obj_23 := map[string]any{}
	hx_obj_23["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_23["next"] = func() *string {
		hx_post_24 := index
		index = int(int32((index + 1)))
		return keys[hx_post_24]
	}
	return hx_obj_23
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_25 := map[string]any{}
	hx_obj_25["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_25["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_26 := index
			index = int(int32((index + 1)))
			return hx_post_26
		}()])
	}
	return hx_obj_25
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_27 any) map[string]any {
		if hx_value_27 == nil {
			var hx_zero_28 map[string]any
			return hx_zero_28
		}
		return hx_value_27.(map[string]any)
	}(self.keys())
	hx_obj_29 := map[string]any{}
	hx_obj_29["hasNext"] = func() bool {
		return func(hx_obj_30 map[string]any) func() bool {
			hx_field_31 := hx_obj_30["hasNext"]
			if hx_field_31 == nil {
				var hx_zero_32 func() bool
				return hx_zero_32
			}
			return hx_field_31.(func() bool)
		}(keys)()
	}
	hx_obj_29["next"] = func() map[string]any {
		key := func(hx_obj_33 map[string]any) func() *string {
			hx_field_34 := hx_obj_33["next"]
			if hx_field_34 == nil {
				var hx_zero_35 func() *string
				return hx_zero_35
			}
			return hx_field_34.(func() *string)
		}(keys)()
		hx_obj_36 := map[string]any{}
		hx_obj_36["key"] = key
		hx_obj_36["value"] = _gthis.get(key)
		return hx_obj_36
	}
	return hx_obj_29
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_37 any) *string {
		if hx_value_37 == nil {
			var hx_zero_38 *string
			return hx_zero_38
		}
		return hx_value_37.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_39 any) *string {
		if hx_value_39 == nil {
			var hx_zero_40 *string
			return hx_zero_40
		}
		return hx_value_39.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_43 any) bool {
		if hx_value_43 == nil {
			var hx_zero_44 bool
			return hx_zero_44
		}
		return hx_value_43.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_41 any) *string {
		if hx_value_41 == nil {
			var hx_zero_42 *string
			return hx_zero_42
		}
		return hx_value_41.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_47 any) bool {
		if hx_value_47 == nil {
			var hx_zero_48 bool
			return hx_zero_48
		}
		return hx_value_47.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_45 any) *string {
		if hx_value_45 == nil {
			var hx_zero_46 *string
			return hx_zero_46
		}
		return hx_value_45.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_49 any) *haxe__ds__StringMap {
		if hx_value_49 == nil {
			var hx_zero_50 *haxe__ds__StringMap
			return hx_zero_50
		}
		return hx_value_49.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_51 any) map[string]any {
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
	}(key)() {
		key_1 := func(hx_obj_56 map[string]any) func() *string {
			hx_field_57 := hx_obj_56["next"]
			if hx_field_57 == nil {
				var hx_zero_58 func() *string
				return hx_zero_58
			}
			return hx_field_57.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_59 any) map[string]any {
		if hx_value_59 == nil {
			var hx_zero_60 map[string]any
			return hx_zero_60
		}
		return hx_value_59.(map[string]any)
	}(self.keys())
	for func(hx_obj_61 map[string]any) func() bool {
		hx_field_62 := hx_obj_61["hasNext"]
		if hx_field_62 == nil {
			var hx_zero_63 func() bool
			return hx_zero_63
		}
		return hx_field_62.(func() bool)
	}(iterator)() {
		key := func(hx_obj_64 map[string]any) func() *string {
			hx_field_65 := hx_obj_64["next"]
			if hx_field_65 == nil {
				var hx_zero_66 func() *string
				return hx_zero_66
			}
			return hx_field_65.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_67 map[string]any) func() bool {
			hx_field_68 := hx_obj_67["hasNext"]
			if hx_field_68 == nil {
				var hx_zero_69 func() bool
				return hx_zero_69
			}
			return hx_field_68.(func() bool)
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
