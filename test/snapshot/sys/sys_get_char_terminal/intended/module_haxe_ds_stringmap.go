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
	hx_obj_42 := map[string]any{}
	hx_obj_42["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_42["next"] = func() *string {
		hx_post_43 := index
		index = int(int32((index + 1)))
		return keys[hx_post_43]
	}
	return hx_obj_42
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_44 := map[string]any{}
	hx_obj_44["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_44["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_45 := index
			index = int(int32((index + 1)))
			return hx_post_45
		}()])
	}
	return hx_obj_44
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_46 any) map[string]any {
		if hx_value_46 == nil {
			var hx_zero_47 map[string]any
			return hx_zero_47
		}
		return hx_value_46.(map[string]any)
	}(self.keys())
	hx_obj_48 := map[string]any{}
	hx_obj_48["hasNext"] = func() bool {
		return func(hx_obj_49 map[string]any) func() bool {
			hx_field_50 := hx_obj_49["hasNext"]
			if hx_field_50 == nil {
				var hx_zero_51 func() bool
				return hx_zero_51
			}
			return hx_field_50.(func() bool)
		}(keys)()
	}
	hx_obj_48["next"] = func() map[string]any {
		key := func(hx_obj_52 map[string]any) func() *string {
			hx_field_53 := hx_obj_52["next"]
			if hx_field_53 == nil {
				var hx_zero_54 func() *string
				return hx_zero_54
			}
			return hx_field_53.(func() *string)
		}(keys)()
		hx_obj_55 := map[string]any{}
		hx_obj_55["key"] = key
		hx_obj_55["value"] = _gthis.get(key)
		return hx_obj_55
	}
	return hx_obj_48
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_56 any) *string {
		if hx_value_56 == nil {
			var hx_zero_57 *string
			return hx_zero_57
		}
		return hx_value_56.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_58 any) *string {
		if hx_value_58 == nil {
			var hx_zero_59 *string
			return hx_zero_59
		}
		return hx_value_58.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_62 any) bool {
		if hx_value_62 == nil {
			var hx_zero_63 bool
			return hx_zero_63
		}
		return hx_value_62.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_60 any) *string {
		if hx_value_60 == nil {
			var hx_zero_61 *string
			return hx_zero_61
		}
		return hx_value_60.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_66 any) bool {
		if hx_value_66 == nil {
			var hx_zero_67 bool
			return hx_zero_67
		}
		return hx_value_66.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_64 any) *string {
		if hx_value_64 == nil {
			var hx_zero_65 *string
			return hx_zero_65
		}
		return hx_value_64.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_68 any) *haxe__ds__StringMap {
		if hx_value_68 == nil {
			var hx_zero_69 *haxe__ds__StringMap
			return hx_zero_69
		}
		return hx_value_68.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_70 any) map[string]any {
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
	}(key)() {
		key_1 := func(hx_obj_75 map[string]any) func() *string {
			hx_field_76 := hx_obj_75["next"]
			if hx_field_76 == nil {
				var hx_zero_77 func() *string
				return hx_zero_77
			}
			return hx_field_76.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_78 any) map[string]any {
		if hx_value_78 == nil {
			var hx_zero_79 map[string]any
			return hx_zero_79
		}
		return hx_value_78.(map[string]any)
	}(self.keys())
	for func(hx_obj_80 map[string]any) func() bool {
		hx_field_81 := hx_obj_80["hasNext"]
		if hx_field_81 == nil {
			var hx_zero_82 func() bool
			return hx_zero_82
		}
		return hx_field_81.(func() bool)
	}(iterator)() {
		key := func(hx_obj_83 map[string]any) func() *string {
			hx_field_84 := hx_obj_83["next"]
			if hx_field_84 == nil {
				var hx_zero_85 func() *string
				return hx_zero_85
			}
			return hx_field_84.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_86 map[string]any) func() bool {
			hx_field_87 := hx_obj_86["hasNext"]
			if hx_field_87 == nil {
				var hx_zero_88 func() bool
				return hx_zero_88
			}
			return hx_field_87.(func() bool)
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
