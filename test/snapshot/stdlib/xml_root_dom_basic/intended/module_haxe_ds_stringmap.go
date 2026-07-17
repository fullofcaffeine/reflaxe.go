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
	hx_obj_105 := map[string]any{}
	hx_obj_105["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_105["next"] = func() *string {
		hx_post_106 := index
		index = int(int32((index + 1)))
		return keys[hx_post_106]
	}
	return hx_obj_105
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_107 := map[string]any{}
	hx_obj_107["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_107["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_108 := index
			index = int(int32((index + 1)))
			return hx_post_108
		}()])
	}
	return hx_obj_107
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_109 any) map[string]any {
		if hx_value_109 == nil {
			var hx_zero_110 map[string]any
			return hx_zero_110
		}
		return hx_value_109.(map[string]any)
	}(self.keys())
	hx_obj_111 := map[string]any{}
	hx_obj_111["hasNext"] = func() bool {
		return func(hx_obj_112 map[string]any) func() bool {
			hx_field_113 := hx_obj_112["hasNext"]
			if hx_field_113 == nil {
				var hx_zero_114 func() bool
				return hx_zero_114
			}
			return hx_field_113.(func() bool)
		}(keys)()
	}
	hx_obj_111["next"] = func() map[string]any {
		key := func(hx_obj_115 map[string]any) func() *string {
			hx_field_116 := hx_obj_115["next"]
			if hx_field_116 == nil {
				var hx_zero_117 func() *string
				return hx_zero_117
			}
			return hx_field_116.(func() *string)
		}(keys)()
		hx_obj_118 := map[string]any{}
		hx_obj_118["key"] = key
		hx_obj_118["value"] = _gthis.get(key)
		return hx_obj_118
	}
	return hx_obj_111
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_119 any) *string {
		if hx_value_119 == nil {
			var hx_zero_120 *string
			return hx_zero_120
		}
		return hx_value_119.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_121 any) *string {
		if hx_value_121 == nil {
			var hx_zero_122 *string
			return hx_zero_122
		}
		return hx_value_121.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_125 any) bool {
		if hx_value_125 == nil {
			var hx_zero_126 bool
			return hx_zero_126
		}
		return hx_value_125.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_123 any) *string {
		if hx_value_123 == nil {
			var hx_zero_124 *string
			return hx_zero_124
		}
		return hx_value_123.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_129 any) bool {
		if hx_value_129 == nil {
			var hx_zero_130 bool
			return hx_zero_130
		}
		return hx_value_129.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_127 any) *string {
		if hx_value_127 == nil {
			var hx_zero_128 *string
			return hx_zero_128
		}
		return hx_value_127.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_131 any) *haxe__ds__StringMap {
		if hx_value_131 == nil {
			var hx_zero_132 *haxe__ds__StringMap
			return hx_zero_132
		}
		return hx_value_131.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_133 any) map[string]any {
		if hx_value_133 == nil {
			var hx_zero_134 map[string]any
			return hx_zero_134
		}
		return hx_value_133.(map[string]any)
	}(self.keys())
	for func(hx_obj_135 map[string]any) func() bool {
		hx_field_136 := hx_obj_135["hasNext"]
		if hx_field_136 == nil {
			var hx_zero_137 func() bool
			return hx_zero_137
		}
		return hx_field_136.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_138 map[string]any) func() *string {
			hx_field_139 := hx_obj_138["next"]
			if hx_field_139 == nil {
				var hx_zero_140 func() *string
				return hx_zero_140
			}
			return hx_field_139.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_141 any) map[string]any {
		if hx_value_141 == nil {
			var hx_zero_142 map[string]any
			return hx_zero_142
		}
		return hx_value_141.(map[string]any)
	}(self.keys())
	for func(hx_obj_143 map[string]any) func() bool {
		hx_field_144 := hx_obj_143["hasNext"]
		if hx_field_144 == nil {
			var hx_zero_145 func() bool
			return hx_zero_145
		}
		return hx_field_144.(func() bool)
	}(iterator)() {
		key := func(hx_obj_146 map[string]any) func() *string {
			hx_field_147 := hx_obj_146["next"]
			if hx_field_147 == nil {
				var hx_zero_148 func() *string
				return hx_zero_148
			}
			return hx_field_147.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_149 map[string]any) func() bool {
			hx_field_150 := hx_obj_149["hasNext"]
			if hx_field_150 == nil {
				var hx_zero_151 func() bool
				return hx_zero_151
			}
			return hx_field_150.(func() bool)
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
