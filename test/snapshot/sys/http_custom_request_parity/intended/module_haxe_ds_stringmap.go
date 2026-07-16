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
	hx_obj_56 := map[string]any{}
	hx_obj_56["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_56["next"] = func() *string {
		hx_post_57 := index
		index = int(int32((index + 1)))
		return keys[hx_post_57]
	}
	return hx_obj_56
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_58 := map[string]any{}
	hx_obj_58["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_58["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_59 := index
			index = int(int32((index + 1)))
			return hx_post_59
		}()])
	}
	return hx_obj_58
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_60 any) map[string]any {
		if hx_value_60 == nil {
			var hx_zero_61 map[string]any
			return hx_zero_61
		}
		return hx_value_60.(map[string]any)
	}(self.keys())
	hx_obj_62 := map[string]any{}
	hx_obj_62["hasNext"] = func() bool {
		return func(hx_obj_63 map[string]any) func() bool {
			hx_field_64 := hx_obj_63["hasNext"]
			if hx_field_64 == nil {
				var hx_zero_65 func() bool
				return hx_zero_65
			}
			return hx_field_64.(func() bool)
		}(keys)()
	}
	hx_obj_62["next"] = func() map[string]any {
		key := func(hx_obj_66 map[string]any) func() *string {
			hx_field_67 := hx_obj_66["next"]
			if hx_field_67 == nil {
				var hx_zero_68 func() *string
				return hx_zero_68
			}
			return hx_field_67.(func() *string)
		}(keys)()
		hx_obj_69 := map[string]any{}
		hx_obj_69["key"] = key
		hx_obj_69["value"] = _gthis.get(key)
		return hx_obj_69
	}
	return hx_obj_62
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_70 any) *string {
		if hx_value_70 == nil {
			var hx_zero_71 *string
			return hx_zero_71
		}
		return hx_value_70.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_72 any) *string {
		if hx_value_72 == nil {
			var hx_zero_73 *string
			return hx_zero_73
		}
		return hx_value_72.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_76 any) bool {
		if hx_value_76 == nil {
			var hx_zero_77 bool
			return hx_zero_77
		}
		return hx_value_76.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_74 any) *string {
		if hx_value_74 == nil {
			var hx_zero_75 *string
			return hx_zero_75
		}
		return hx_value_74.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_80 any) bool {
		if hx_value_80 == nil {
			var hx_zero_81 bool
			return hx_zero_81
		}
		return hx_value_80.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_78 any) *string {
		if hx_value_78 == nil {
			var hx_zero_79 *string
			return hx_zero_79
		}
		return hx_value_78.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_82 any) *haxe__ds__StringMap {
		if hx_value_82 == nil {
			var hx_zero_83 *haxe__ds__StringMap
			return hx_zero_83
		}
		return hx_value_82.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_84 any) map[string]any {
		if hx_value_84 == nil {
			var hx_zero_85 map[string]any
			return hx_zero_85
		}
		return hx_value_84.(map[string]any)
	}(self.keys())
	for func(hx_obj_86 map[string]any) func() bool {
		hx_field_87 := hx_obj_86["hasNext"]
		if hx_field_87 == nil {
			var hx_zero_88 func() bool
			return hx_zero_88
		}
		return hx_field_87.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_89 map[string]any) func() *string {
			hx_field_90 := hx_obj_89["next"]
			if hx_field_90 == nil {
				var hx_zero_91 func() *string
				return hx_zero_91
			}
			return hx_field_90.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_92 any) map[string]any {
		if hx_value_92 == nil {
			var hx_zero_93 map[string]any
			return hx_zero_93
		}
		return hx_value_92.(map[string]any)
	}(self.keys())
	for func(hx_obj_94 map[string]any) func() bool {
		hx_field_95 := hx_obj_94["hasNext"]
		if hx_field_95 == nil {
			var hx_zero_96 func() bool
			return hx_zero_96
		}
		return hx_field_95.(func() bool)
	}(iterator)() {
		key := func(hx_obj_97 map[string]any) func() *string {
			hx_field_98 := hx_obj_97["next"]
			if hx_field_98 == nil {
				var hx_zero_99 func() *string
				return hx_zero_99
			}
			return hx_field_98.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_100 map[string]any) func() bool {
			hx_field_101 := hx_obj_100["hasNext"]
			if hx_field_101 == nil {
				var hx_zero_102 func() bool
				return hx_zero_102
			}
			return hx_field_101.(func() bool)
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
