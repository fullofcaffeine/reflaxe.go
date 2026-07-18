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
	hx_obj_112 := map[string]any{}
	hx_obj_112["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_112["next"] = func() int {
		hx_post_113 := index
		index = int(int32((index + 1)))
		return keys[hx_post_113]
	}
	return hx_obj_112
}

func (self *haxe__ds__IntMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.IntMapKeys(self.h)
	index := 0
	hx_obj_114 := map[string]any{}
	hx_obj_114["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_114["next"] = func() any {
		return hxrt.IntMapGet(_gthis.h, keys[func() int {
			hx_post_115 := index
			index = int(int32((index + 1)))
			return hx_post_115
		}()])
	}
	return hx_obj_114
}

func (self *haxe__ds__IntMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_116 any) map[string]any {
		if hx_value_116 == nil {
			var hx_zero_117 map[string]any
			return hx_zero_117
		}
		return hx_value_116.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_118 := map[string]any{}
	hx_obj_118["hasNext"] = func() bool {
		return func(hx_obj_119 map[string]any) func() bool {
			hx_field_120 := hx_obj_119["hasNext"]
			if hx_field_120 == nil {
				var hx_zero_121 func() bool
				return hx_zero_121
			}
			return hx_field_120.(func() bool)
		}(keys)()
	}
	hx_obj_118["next"] = func() map[string]any {
		key := func(hx_obj_122 map[string]any) func() int {
			hx_field_123 := hx_obj_122["next"]
			if hx_field_123 == nil {
				var hx_zero_124 func() int
				return hx_zero_124
			}
			return hx_field_123.(func() int)
		}(keys)()
		hx_obj_125 := map[string]any{}
		hx_obj_125["key"] = key
		hx_obj_125["value"] = _gthis.__hx_this.get(key)
		return hx_obj_125
	}
	return hx_obj_118
}

func (self *haxe__ds__IntMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.IntFromNullableAny(func(hx_value_126 any) int {
		if hx_value_126 == nil {
			var hx_zero_127 int
			return hx_zero_127
		}
		return hx_value_126.(int)
	}(key)))
}

func (self *haxe__ds__IntMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.IntFromNullableAny(func(hx_value_128 any) int {
		if hx_value_128 == nil {
			var hx_zero_129 int
			return hx_zero_129
		}
		return hx_value_128.(int)
	}(key)), value)
}

func (self *haxe__ds__IntMap) existsIMap(key any) bool {
	return func(hx_value_132 any) bool {
		if hx_value_132 == nil {
			var hx_zero_133 bool
			return hx_zero_133
		}
		return hx_value_132.(bool)
	}(self.__hx_this.exists(hxrt.IntFromNullableAny(func(hx_value_130 any) int {
		if hx_value_130 == nil {
			var hx_zero_131 int
			return hx_zero_131
		}
		return hx_value_130.(int)
	}(key))))
}

func (self *haxe__ds__IntMap) removeIMap(key any) bool {
	return func(hx_value_136 any) bool {
		if hx_value_136 == nil {
			var hx_zero_137 bool
			return hx_zero_137
		}
		return hx_value_136.(bool)
	}(self.__hx_this.remove(hxrt.IntFromNullableAny(func(hx_value_134 any) int {
		if hx_value_134 == nil {
			var hx_zero_135 int
			return hx_zero_135
		}
		return hx_value_134.(int)
	}(key))))
}

func (self *haxe__ds__IntMap) copyIMap() haxe__IMap {
	return func(hx_value_138 any) *haxe__ds__IntMap {
		if hx_value_138 == nil {
			var hx_zero_139 *haxe__ds__IntMap
			return hx_zero_139
		}
		return hx_value_138.(*haxe__ds__IntMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__IntMap) copy() *haxe__ds__IntMap {
	copied := New_haxe__ds__IntMap()
	key := func(hx_value_140 any) map[string]any {
		if hx_value_140 == nil {
			var hx_zero_141 map[string]any
			return hx_zero_141
		}
		return hx_value_140.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_142 map[string]any) func() bool {
		hx_field_143 := hx_obj_142["hasNext"]
		if hx_field_143 == nil {
			var hx_zero_144 func() bool
			return hx_zero_144
		}
		return hx_field_143.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_145 map[string]any) func() int {
			hx_field_146 := hx_obj_145["next"]
			if hx_field_146 == nil {
				var hx_zero_147 func() int
				return hx_zero_147
			}
			return hx_field_146.(func() int)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__IntMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_148 any) map[string]any {
		if hx_value_148 == nil {
			var hx_zero_149 map[string]any
			return hx_zero_149
		}
		return hx_value_148.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_150 map[string]any) func() bool {
		hx_field_151 := hx_obj_150["hasNext"]
		if hx_field_151 == nil {
			var hx_zero_152 func() bool
			return hx_zero_152
		}
		return hx_field_151.(func() bool)
	}(iterator)() {
		key := func(hx_obj_153 map[string]any) func() int {
			hx_field_154 := hx_obj_153["next"]
			if hx_field_154 == nil {
				var hx_zero_155 func() int
				return hx_zero_155
			}
			return hx_field_154.(func() int)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_156 map[string]any) func() bool {
			hx_field_157 := hx_obj_156["hasNext"]
			if hx_field_157 == nil {
				var hx_zero_158 func() bool
				return hx_zero_158
			}
			return hx_field_157.(func() bool)
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
