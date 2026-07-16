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
	hx_obj_71 := map[string]any{}
	hx_obj_71["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_71["next"] = func() int {
		hx_post_72 := index
		index = int(int32((index + 1)))
		return keys[hx_post_72]
	}
	return hx_obj_71
}

func (self *haxe__ds__IntMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.IntMapKeys(self.h)
	index := 0
	hx_obj_73 := map[string]any{}
	hx_obj_73["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_73["next"] = func() any {
		return hxrt.IntMapGet(_gthis.h, keys[func() int {
			hx_post_74 := index
			index = int(int32((index + 1)))
			return hx_post_74
		}()])
	}
	return hx_obj_73
}

func (self *haxe__ds__IntMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_75 any) map[string]any {
		if hx_value_75 == nil {
			var hx_zero_76 map[string]any
			return hx_zero_76
		}
		return hx_value_75.(map[string]any)
	}(self.keys())
	hx_obj_77 := map[string]any{}
	hx_obj_77["hasNext"] = func() bool {
		return func(hx_obj_78 map[string]any) func() bool {
			hx_field_79 := hx_obj_78["hasNext"]
			if hx_field_79 == nil {
				var hx_zero_80 func() bool
				return hx_zero_80
			}
			return hx_field_79.(func() bool)
		}(keys)()
	}
	hx_obj_77["next"] = func() map[string]any {
		key := func(hx_obj_81 map[string]any) func() int {
			hx_field_82 := hx_obj_81["next"]
			if hx_field_82 == nil {
				var hx_zero_83 func() int
				return hx_zero_83
			}
			return hx_field_82.(func() int)
		}(keys)()
		hx_obj_84 := map[string]any{}
		hx_obj_84["key"] = key
		hx_obj_84["value"] = _gthis.get(key)
		return hx_obj_84
	}
	return hx_obj_77
}

func (self *haxe__ds__IntMap) getIMap(key any) any {
	return self.get(hxrt.IntFromNullableAny(func(hx_value_85 any) int {
		if hx_value_85 == nil {
			var hx_zero_86 int
			return hx_zero_86
		}
		return hx_value_85.(int)
	}(key)))
}

func (self *haxe__ds__IntMap) setIMap(key any, value any) {
	self.set(hxrt.IntFromNullableAny(func(hx_value_87 any) int {
		if hx_value_87 == nil {
			var hx_zero_88 int
			return hx_zero_88
		}
		return hx_value_87.(int)
	}(key)), value)
}

func (self *haxe__ds__IntMap) existsIMap(key any) bool {
	return func(hx_value_91 any) bool {
		if hx_value_91 == nil {
			var hx_zero_92 bool
			return hx_zero_92
		}
		return hx_value_91.(bool)
	}(self.exists(hxrt.IntFromNullableAny(func(hx_value_89 any) int {
		if hx_value_89 == nil {
			var hx_zero_90 int
			return hx_zero_90
		}
		return hx_value_89.(int)
	}(key))))
}

func (self *haxe__ds__IntMap) removeIMap(key any) bool {
	return func(hx_value_95 any) bool {
		if hx_value_95 == nil {
			var hx_zero_96 bool
			return hx_zero_96
		}
		return hx_value_95.(bool)
	}(self.remove(hxrt.IntFromNullableAny(func(hx_value_93 any) int {
		if hx_value_93 == nil {
			var hx_zero_94 int
			return hx_zero_94
		}
		return hx_value_93.(int)
	}(key))))
}

func (self *haxe__ds__IntMap) copyIMap() haxe__IMap {
	return func(hx_value_97 any) *haxe__ds__IntMap {
		if hx_value_97 == nil {
			var hx_zero_98 *haxe__ds__IntMap
			return hx_zero_98
		}
		return hx_value_97.(*haxe__ds__IntMap)
	}(self.copy())
}

func (self *haxe__ds__IntMap) copy() *haxe__ds__IntMap {
	copied := New_haxe__ds__IntMap()
	key := func(hx_value_99 any) map[string]any {
		if hx_value_99 == nil {
			var hx_zero_100 map[string]any
			return hx_zero_100
		}
		return hx_value_99.(map[string]any)
	}(self.keys())
	for func(hx_obj_101 map[string]any) func() bool {
		hx_field_102 := hx_obj_101["hasNext"]
		if hx_field_102 == nil {
			var hx_zero_103 func() bool
			return hx_zero_103
		}
		return hx_field_102.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_104 map[string]any) func() int {
			hx_field_105 := hx_obj_104["next"]
			if hx_field_105 == nil {
				var hx_zero_106 func() int
				return hx_zero_106
			}
			return hx_field_105.(func() int)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__IntMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_107 any) map[string]any {
		if hx_value_107 == nil {
			var hx_zero_108 map[string]any
			return hx_zero_108
		}
		return hx_value_107.(map[string]any)
	}(self.keys())
	for func(hx_obj_109 map[string]any) func() bool {
		hx_field_110 := hx_obj_109["hasNext"]
		if hx_field_110 == nil {
			var hx_zero_111 func() bool
			return hx_zero_111
		}
		return hx_field_110.(func() bool)
	}(iterator)() {
		key := func(hx_obj_112 map[string]any) func() int {
			hx_field_113 := hx_obj_112["next"]
			if hx_field_113 == nil {
				var hx_zero_114 func() int
				return hx_zero_114
			}
			return hx_field_113.(func() int)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_115 map[string]any) func() bool {
			hx_field_116 := hx_obj_115["hasNext"]
			if hx_field_116 == nil {
				var hx_zero_117 func() bool
				return hx_zero_117
			}
			return hx_field_116.(func() bool)
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
