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
	hx_obj_79 := map[string]any{}
	hx_obj_79["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_79["next"] = func() *string {
		hx_post_80 := index
		index = int(int32((index + 1)))
		return keys[hx_post_80]
	}
	return hx_obj_79
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_81 := map[string]any{}
	hx_obj_81["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_81["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_82 := index
			index = int(int32((index + 1)))
			return hx_post_82
		}()])
	}
	return hx_obj_81
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_83 any) map[string]any {
		if hx_value_83 == nil {
			var hx_zero_84 map[string]any
			return hx_zero_84
		}
		return hx_value_83.(map[string]any)
	}(self.keys())
	hx_obj_85 := map[string]any{}
	hx_obj_85["hasNext"] = func() bool {
		return func(hx_obj_86 map[string]any) func() bool {
			hx_field_87 := hx_obj_86["hasNext"]
			if hx_field_87 == nil {
				var hx_zero_88 func() bool
				return hx_zero_88
			}
			return hx_field_87.(func() bool)
		}(keys)()
	}
	hx_obj_85["next"] = func() map[string]any {
		key := func(hx_obj_89 map[string]any) func() *string {
			hx_field_90 := hx_obj_89["next"]
			if hx_field_90 == nil {
				var hx_zero_91 func() *string
				return hx_zero_91
			}
			return hx_field_90.(func() *string)
		}(keys)()
		hx_obj_92 := map[string]any{}
		hx_obj_92["key"] = key
		hx_obj_92["value"] = _gthis.get(key)
		return hx_obj_92
	}
	return hx_obj_85
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_93 any) *string {
		if hx_value_93 == nil {
			var hx_zero_94 *string
			return hx_zero_94
		}
		return hx_value_93.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_95 any) *string {
		if hx_value_95 == nil {
			var hx_zero_96 *string
			return hx_zero_96
		}
		return hx_value_95.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_99 any) bool {
		if hx_value_99 == nil {
			var hx_zero_100 bool
			return hx_zero_100
		}
		return hx_value_99.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_97 any) *string {
		if hx_value_97 == nil {
			var hx_zero_98 *string
			return hx_zero_98
		}
		return hx_value_97.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_103 any) bool {
		if hx_value_103 == nil {
			var hx_zero_104 bool
			return hx_zero_104
		}
		return hx_value_103.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_101 any) *string {
		if hx_value_101 == nil {
			var hx_zero_102 *string
			return hx_zero_102
		}
		return hx_value_101.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_105 any) *haxe__ds__StringMap {
		if hx_value_105 == nil {
			var hx_zero_106 *haxe__ds__StringMap
			return hx_zero_106
		}
		return hx_value_105.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_107 any) map[string]any {
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
	}(key)() {
		key_1 := func(hx_obj_112 map[string]any) func() *string {
			hx_field_113 := hx_obj_112["next"]
			if hx_field_113 == nil {
				var hx_zero_114 func() *string
				return hx_zero_114
			}
			return hx_field_113.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_115 any) map[string]any {
		if hx_value_115 == nil {
			var hx_zero_116 map[string]any
			return hx_zero_116
		}
		return hx_value_115.(map[string]any)
	}(self.keys())
	for func(hx_obj_117 map[string]any) func() bool {
		hx_field_118 := hx_obj_117["hasNext"]
		if hx_field_118 == nil {
			var hx_zero_119 func() bool
			return hx_zero_119
		}
		return hx_field_118.(func() bool)
	}(iterator)() {
		key := func(hx_obj_120 map[string]any) func() *string {
			hx_field_121 := hx_obj_120["next"]
			if hx_field_121 == nil {
				var hx_zero_122 func() *string
				return hx_zero_122
			}
			return hx_field_121.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_123 map[string]any) func() bool {
			hx_field_124 := hx_obj_123["hasNext"]
			if hx_field_124 == nil {
				var hx_zero_125 func() bool
				return hx_zero_125
			}
			return hx_field_124.(func() bool)
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
