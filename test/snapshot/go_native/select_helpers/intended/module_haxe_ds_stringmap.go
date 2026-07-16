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
	hx_obj_34 := map[string]any{}
	hx_obj_34["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_34["next"] = func() *string {
		hx_post_35 := index
		index = int(int32((index + 1)))
		return keys[hx_post_35]
	}
	return hx_obj_34
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_36 := map[string]any{}
	hx_obj_36["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_36["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_37 := index
			index = int(int32((index + 1)))
			return hx_post_37
		}()])
	}
	return hx_obj_36
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_38 any) map[string]any {
		if hx_value_38 == nil {
			var hx_zero_39 map[string]any
			return hx_zero_39
		}
		return hx_value_38.(map[string]any)
	}(self.keys())
	hx_obj_40 := map[string]any{}
	hx_obj_40["hasNext"] = func() bool {
		return func(hx_obj_41 map[string]any) func() bool {
			hx_field_42 := hx_obj_41["hasNext"]
			if hx_field_42 == nil {
				var hx_zero_43 func() bool
				return hx_zero_43
			}
			return hx_field_42.(func() bool)
		}(keys)()
	}
	hx_obj_40["next"] = func() map[string]any {
		key := func(hx_obj_44 map[string]any) func() *string {
			hx_field_45 := hx_obj_44["next"]
			if hx_field_45 == nil {
				var hx_zero_46 func() *string
				return hx_zero_46
			}
			return hx_field_45.(func() *string)
		}(keys)()
		hx_obj_47 := map[string]any{}
		hx_obj_47["key"] = key
		hx_obj_47["value"] = _gthis.get(key)
		return hx_obj_47
	}
	return hx_obj_40
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_48 any) *string {
		if hx_value_48 == nil {
			var hx_zero_49 *string
			return hx_zero_49
		}
		return hx_value_48.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_50 any) *string {
		if hx_value_50 == nil {
			var hx_zero_51 *string
			return hx_zero_51
		}
		return hx_value_50.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_54 any) bool {
		if hx_value_54 == nil {
			var hx_zero_55 bool
			return hx_zero_55
		}
		return hx_value_54.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_52 any) *string {
		if hx_value_52 == nil {
			var hx_zero_53 *string
			return hx_zero_53
		}
		return hx_value_52.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_58 any) bool {
		if hx_value_58 == nil {
			var hx_zero_59 bool
			return hx_zero_59
		}
		return hx_value_58.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_56 any) *string {
		if hx_value_56 == nil {
			var hx_zero_57 *string
			return hx_zero_57
		}
		return hx_value_56.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_60 any) *haxe__ds__StringMap {
		if hx_value_60 == nil {
			var hx_zero_61 *haxe__ds__StringMap
			return hx_zero_61
		}
		return hx_value_60.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_62 any) map[string]any {
		if hx_value_62 == nil {
			var hx_zero_63 map[string]any
			return hx_zero_63
		}
		return hx_value_62.(map[string]any)
	}(self.keys())
	for func(hx_obj_64 map[string]any) func() bool {
		hx_field_65 := hx_obj_64["hasNext"]
		if hx_field_65 == nil {
			var hx_zero_66 func() bool
			return hx_zero_66
		}
		return hx_field_65.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_67 map[string]any) func() *string {
			hx_field_68 := hx_obj_67["next"]
			if hx_field_68 == nil {
				var hx_zero_69 func() *string
				return hx_zero_69
			}
			return hx_field_68.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_70 any) map[string]any {
		if hx_value_70 == nil {
			var hx_zero_71 map[string]any
			return hx_zero_71
		}
		return hx_value_70.(map[string]any)
	}(self.keys())
	for func(hx_obj_72 map[string]any) func() bool {
		hx_field_73 := hx_obj_72["hasNext"]
		if hx_field_73 == nil {
			var hx_zero_74 func() bool
			return hx_zero_74
		}
		return hx_field_73.(func() bool)
	}(iterator)() {
		key := func(hx_obj_75 map[string]any) func() *string {
			hx_field_76 := hx_obj_75["next"]
			if hx_field_76 == nil {
				var hx_zero_77 func() *string
				return hx_zero_77
			}
			return hx_field_76.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_78 map[string]any) func() bool {
			hx_field_79 := hx_obj_78["hasNext"]
			if hx_field_79 == nil {
				var hx_zero_80 func() bool
				return hx_zero_80
			}
			return hx_field_79.(func() bool)
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
