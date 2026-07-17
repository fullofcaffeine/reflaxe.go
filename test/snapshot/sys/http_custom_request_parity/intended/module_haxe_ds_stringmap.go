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
	hx_obj_70 := map[string]any{}
	hx_obj_70["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_70["next"] = func() *string {
		hx_post_71 := index
		index = int(int32((index + 1)))
		return keys[hx_post_71]
	}
	return hx_obj_70
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_72 := map[string]any{}
	hx_obj_72["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_72["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_73 := index
			index = int(int32((index + 1)))
			return hx_post_73
		}()])
	}
	return hx_obj_72
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_74 any) map[string]any {
		if hx_value_74 == nil {
			var hx_zero_75 map[string]any
			return hx_zero_75
		}
		return hx_value_74.(map[string]any)
	}(self.keys())
	hx_obj_76 := map[string]any{}
	hx_obj_76["hasNext"] = func() bool {
		return func(hx_obj_77 map[string]any) func() bool {
			hx_field_78 := hx_obj_77["hasNext"]
			if hx_field_78 == nil {
				var hx_zero_79 func() bool
				return hx_zero_79
			}
			return hx_field_78.(func() bool)
		}(keys)()
	}
	hx_obj_76["next"] = func() map[string]any {
		key := func(hx_obj_80 map[string]any) func() *string {
			hx_field_81 := hx_obj_80["next"]
			if hx_field_81 == nil {
				var hx_zero_82 func() *string
				return hx_zero_82
			}
			return hx_field_81.(func() *string)
		}(keys)()
		hx_obj_83 := map[string]any{}
		hx_obj_83["key"] = key
		hx_obj_83["value"] = _gthis.get(key)
		return hx_obj_83
	}
	return hx_obj_76
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_84 any) *string {
		if hx_value_84 == nil {
			var hx_zero_85 *string
			return hx_zero_85
		}
		return hx_value_84.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_86 any) *string {
		if hx_value_86 == nil {
			var hx_zero_87 *string
			return hx_zero_87
		}
		return hx_value_86.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_90 any) bool {
		if hx_value_90 == nil {
			var hx_zero_91 bool
			return hx_zero_91
		}
		return hx_value_90.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_88 any) *string {
		if hx_value_88 == nil {
			var hx_zero_89 *string
			return hx_zero_89
		}
		return hx_value_88.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_94 any) bool {
		if hx_value_94 == nil {
			var hx_zero_95 bool
			return hx_zero_95
		}
		return hx_value_94.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_92 any) *string {
		if hx_value_92 == nil {
			var hx_zero_93 *string
			return hx_zero_93
		}
		return hx_value_92.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_96 any) *haxe__ds__StringMap {
		if hx_value_96 == nil {
			var hx_zero_97 *haxe__ds__StringMap
			return hx_zero_97
		}
		return hx_value_96.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_98 any) map[string]any {
		if hx_value_98 == nil {
			var hx_zero_99 map[string]any
			return hx_zero_99
		}
		return hx_value_98.(map[string]any)
	}(self.keys())
	for func(hx_obj_100 map[string]any) func() bool {
		hx_field_101 := hx_obj_100["hasNext"]
		if hx_field_101 == nil {
			var hx_zero_102 func() bool
			return hx_zero_102
		}
		return hx_field_101.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_103 map[string]any) func() *string {
			hx_field_104 := hx_obj_103["next"]
			if hx_field_104 == nil {
				var hx_zero_105 func() *string
				return hx_zero_105
			}
			return hx_field_104.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_106 any) map[string]any {
		if hx_value_106 == nil {
			var hx_zero_107 map[string]any
			return hx_zero_107
		}
		return hx_value_106.(map[string]any)
	}(self.keys())
	for func(hx_obj_108 map[string]any) func() bool {
		hx_field_109 := hx_obj_108["hasNext"]
		if hx_field_109 == nil {
			var hx_zero_110 func() bool
			return hx_zero_110
		}
		return hx_field_109.(func() bool)
	}(iterator)() {
		key := func(hx_obj_111 map[string]any) func() *string {
			hx_field_112 := hx_obj_111["next"]
			if hx_field_112 == nil {
				var hx_zero_113 func() *string
				return hx_zero_113
			}
			return hx_field_112.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_114 map[string]any) func() bool {
			hx_field_115 := hx_obj_114["hasNext"]
			if hx_field_115 == nil {
				var hx_zero_116 func() bool
				return hx_zero_116
			}
			return hx_field_115.(func() bool)
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
