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
	hx_obj_113 := map[string]any{}
	hx_obj_113["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_113["next"] = func() int {
		hx_post_114 := index
		index = int(int32((index + 1)))
		return keys[hx_post_114]
	}
	return hx_obj_113
}

func (self *haxe__ds__IntMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.IntMapKeys(self.h)
	index := 0
	hx_obj_115 := map[string]any{}
	hx_obj_115["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_115["next"] = func() any {
		return hxrt.IntMapGet(_gthis.h, keys[func() int {
			hx_post_116 := index
			index = int(int32((index + 1)))
			return hx_post_116
		}()])
	}
	return hx_obj_115
}

func (self *haxe__ds__IntMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_117 any) map[string]any {
		if hx_value_117 == nil {
			var hx_zero_118 map[string]any
			return hx_zero_118
		}
		return hx_value_117.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_119 := map[string]any{}
	hx_obj_119["hasNext"] = func() bool {
		return func(hx_obj_120 map[string]any) func() bool {
			hx_field_121 := hx_obj_120["hasNext"]
			if hx_field_121 == nil {
				var hx_zero_122 func() bool
				return hx_zero_122
			}
			return hx_field_121.(func() bool)
		}(keys)()
	}
	hx_obj_119["next"] = func() map[string]any {
		key := func(hx_obj_123 map[string]any) func() int {
			hx_field_124 := hx_obj_123["next"]
			if hx_field_124 == nil {
				var hx_zero_125 func() int
				return hx_zero_125
			}
			return hx_field_124.(func() int)
		}(keys)()
		hx_obj_126 := map[string]any{}
		hx_obj_126["key"] = key
		hx_obj_126["value"] = _gthis.__hx_this.get(key)
		return hx_obj_126
	}
	return hx_obj_119
}

func (self *haxe__ds__IntMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.IntFromNullableAny(func(hx_value_127 any) int {
		if hx_value_127 == nil {
			var hx_zero_128 int
			return hx_zero_128
		}
		return hx_value_127.(int)
	}(key)))
}

func (self *haxe__ds__IntMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.IntFromNullableAny(func(hx_value_129 any) int {
		if hx_value_129 == nil {
			var hx_zero_130 int
			return hx_zero_130
		}
		return hx_value_129.(int)
	}(key)), value)
}

func (self *haxe__ds__IntMap) existsIMap(key any) bool {
	return func(hx_value_133 any) bool {
		if hx_value_133 == nil {
			var hx_zero_134 bool
			return hx_zero_134
		}
		return hx_value_133.(bool)
	}(self.__hx_this.exists(hxrt.IntFromNullableAny(func(hx_value_131 any) int {
		if hx_value_131 == nil {
			var hx_zero_132 int
			return hx_zero_132
		}
		return hx_value_131.(int)
	}(key))))
}

func (self *haxe__ds__IntMap) removeIMap(key any) bool {
	return func(hx_value_137 any) bool {
		if hx_value_137 == nil {
			var hx_zero_138 bool
			return hx_zero_138
		}
		return hx_value_137.(bool)
	}(self.__hx_this.remove(hxrt.IntFromNullableAny(func(hx_value_135 any) int {
		if hx_value_135 == nil {
			var hx_zero_136 int
			return hx_zero_136
		}
		return hx_value_135.(int)
	}(key))))
}

func (self *haxe__ds__IntMap) copyIMap() haxe__IMap {
	return func(hx_value_139 any) *haxe__ds__IntMap {
		if hx_value_139 == nil {
			var hx_zero_140 *haxe__ds__IntMap
			return hx_zero_140
		}
		return hx_value_139.(*haxe__ds__IntMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__IntMap) copy() *haxe__ds__IntMap {
	copied := New_haxe__ds__IntMap()
	key := func(hx_value_141 any) map[string]any {
		if hx_value_141 == nil {
			var hx_zero_142 map[string]any
			return hx_zero_142
		}
		return hx_value_141.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_143 map[string]any) func() bool {
		hx_field_144 := hx_obj_143["hasNext"]
		if hx_field_144 == nil {
			var hx_zero_145 func() bool
			return hx_zero_145
		}
		return hx_field_144.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_146 map[string]any) func() int {
			hx_field_147 := hx_obj_146["next"]
			if hx_field_147 == nil {
				var hx_zero_148 func() int
				return hx_zero_148
			}
			return hx_field_147.(func() int)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__IntMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_149 any) map[string]any {
		if hx_value_149 == nil {
			var hx_zero_150 map[string]any
			return hx_zero_150
		}
		return hx_value_149.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_151 map[string]any) func() bool {
		hx_field_152 := hx_obj_151["hasNext"]
		if hx_field_152 == nil {
			var hx_zero_153 func() bool
			return hx_zero_153
		}
		return hx_field_152.(func() bool)
	}(iterator)() {
		key := func(hx_obj_154 map[string]any) func() int {
			hx_field_155 := hx_obj_154["next"]
			if hx_field_155 == nil {
				var hx_zero_156 func() int
				return hx_zero_156
			}
			return hx_field_155.(func() int)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_157 map[string]any) func() bool {
			hx_field_158 := hx_obj_157["hasNext"]
			if hx_field_158 == nil {
				var hx_zero_159 func() bool
				return hx_zero_159
			}
			return hx_field_158.(func() bool)
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

func (self *haxe__ds__IntMap) String() string {
	return *self.__hx_this.toString()
}
