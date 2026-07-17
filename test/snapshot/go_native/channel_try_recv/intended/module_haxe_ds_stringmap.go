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
	hx_obj_31 := map[string]any{}
	hx_obj_31["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_31["next"] = func() *string {
		hx_post_32 := index
		index = int(int32((index + 1)))
		return keys[hx_post_32]
	}
	return hx_obj_31
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_33 := map[string]any{}
	hx_obj_33["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_33["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_34 := index
			index = int(int32((index + 1)))
			return hx_post_34
		}()])
	}
	return hx_obj_33
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_35 any) map[string]any {
		if hx_value_35 == nil {
			var hx_zero_36 map[string]any
			return hx_zero_36
		}
		return hx_value_35.(map[string]any)
	}(self.keys())
	hx_obj_37 := map[string]any{}
	hx_obj_37["hasNext"] = func() bool {
		return func(hx_obj_38 map[string]any) func() bool {
			hx_field_39 := hx_obj_38["hasNext"]
			if hx_field_39 == nil {
				var hx_zero_40 func() bool
				return hx_zero_40
			}
			return hx_field_39.(func() bool)
		}(keys)()
	}
	hx_obj_37["next"] = func() map[string]any {
		key := func(hx_obj_41 map[string]any) func() *string {
			hx_field_42 := hx_obj_41["next"]
			if hx_field_42 == nil {
				var hx_zero_43 func() *string
				return hx_zero_43
			}
			return hx_field_42.(func() *string)
		}(keys)()
		hx_obj_44 := map[string]any{}
		hx_obj_44["key"] = key
		hx_obj_44["value"] = _gthis.get(key)
		return hx_obj_44
	}
	return hx_obj_37
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_45 any) *string {
		if hx_value_45 == nil {
			var hx_zero_46 *string
			return hx_zero_46
		}
		return hx_value_45.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_47 any) *string {
		if hx_value_47 == nil {
			var hx_zero_48 *string
			return hx_zero_48
		}
		return hx_value_47.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_51 any) bool {
		if hx_value_51 == nil {
			var hx_zero_52 bool
			return hx_zero_52
		}
		return hx_value_51.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_49 any) *string {
		if hx_value_49 == nil {
			var hx_zero_50 *string
			return hx_zero_50
		}
		return hx_value_49.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_55 any) bool {
		if hx_value_55 == nil {
			var hx_zero_56 bool
			return hx_zero_56
		}
		return hx_value_55.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_53 any) *string {
		if hx_value_53 == nil {
			var hx_zero_54 *string
			return hx_zero_54
		}
		return hx_value_53.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_57 any) *haxe__ds__StringMap {
		if hx_value_57 == nil {
			var hx_zero_58 *haxe__ds__StringMap
			return hx_zero_58
		}
		return hx_value_57.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_59 any) map[string]any {
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
	}(key)() {
		key_1 := func(hx_obj_64 map[string]any) func() *string {
			hx_field_65 := hx_obj_64["next"]
			if hx_field_65 == nil {
				var hx_zero_66 func() *string
				return hx_zero_66
			}
			return hx_field_65.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_67 any) map[string]any {
		if hx_value_67 == nil {
			var hx_zero_68 map[string]any
			return hx_zero_68
		}
		return hx_value_67.(map[string]any)
	}(self.keys())
	for func(hx_obj_69 map[string]any) func() bool {
		hx_field_70 := hx_obj_69["hasNext"]
		if hx_field_70 == nil {
			var hx_zero_71 func() bool
			return hx_zero_71
		}
		return hx_field_70.(func() bool)
	}(iterator)() {
		key := func(hx_obj_72 map[string]any) func() *string {
			hx_field_73 := hx_obj_72["next"]
			if hx_field_73 == nil {
				var hx_zero_74 func() *string
				return hx_zero_74
			}
			return hx_field_73.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_75 map[string]any) func() bool {
			hx_field_76 := hx_obj_75["hasNext"]
			if hx_field_76 == nil {
				var hx_zero_77 func() bool
				return hx_zero_77
			}
			return hx_field_76.(func() bool)
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
