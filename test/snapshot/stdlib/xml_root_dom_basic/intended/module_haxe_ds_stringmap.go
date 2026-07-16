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
	hx_obj_119 := map[string]any{}
	hx_obj_119["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_119["next"] = func() *string {
		hx_post_120 := index
		index = int(int32((index + 1)))
		return keys[hx_post_120]
	}
	return hx_obj_119
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_121 := map[string]any{}
	hx_obj_121["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_121["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_122 := index
			index = int(int32((index + 1)))
			return hx_post_122
		}()])
	}
	return hx_obj_121
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_123 any) map[string]any {
		if hx_value_123 == nil {
			var hx_zero_124 map[string]any
			return hx_zero_124
		}
		return hx_value_123.(map[string]any)
	}(self.keys())
	hx_obj_125 := map[string]any{}
	hx_obj_125["hasNext"] = func() bool {
		return func(hx_obj_126 map[string]any) func() bool {
			hx_field_127 := hx_obj_126["hasNext"]
			if hx_field_127 == nil {
				var hx_zero_128 func() bool
				return hx_zero_128
			}
			return hx_field_127.(func() bool)
		}(keys)()
	}
	hx_obj_125["next"] = func() map[string]any {
		key := func(hx_obj_129 map[string]any) func() *string {
			hx_field_130 := hx_obj_129["next"]
			if hx_field_130 == nil {
				var hx_zero_131 func() *string
				return hx_zero_131
			}
			return hx_field_130.(func() *string)
		}(keys)()
		hx_obj_132 := map[string]any{}
		hx_obj_132["key"] = key
		hx_obj_132["value"] = _gthis.get(key)
		return hx_obj_132
	}
	return hx_obj_125
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_133 any) *string {
		if hx_value_133 == nil {
			var hx_zero_134 *string
			return hx_zero_134
		}
		return hx_value_133.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_135 any) *string {
		if hx_value_135 == nil {
			var hx_zero_136 *string
			return hx_zero_136
		}
		return hx_value_135.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_139 any) bool {
		if hx_value_139 == nil {
			var hx_zero_140 bool
			return hx_zero_140
		}
		return hx_value_139.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_137 any) *string {
		if hx_value_137 == nil {
			var hx_zero_138 *string
			return hx_zero_138
		}
		return hx_value_137.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_143 any) bool {
		if hx_value_143 == nil {
			var hx_zero_144 bool
			return hx_zero_144
		}
		return hx_value_143.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_141 any) *string {
		if hx_value_141 == nil {
			var hx_zero_142 *string
			return hx_zero_142
		}
		return hx_value_141.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_145 any) *haxe__ds__StringMap {
		if hx_value_145 == nil {
			var hx_zero_146 *haxe__ds__StringMap
			return hx_zero_146
		}
		return hx_value_145.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_147 any) map[string]any {
		if hx_value_147 == nil {
			var hx_zero_148 map[string]any
			return hx_zero_148
		}
		return hx_value_147.(map[string]any)
	}(self.keys())
	for func(hx_obj_149 map[string]any) func() bool {
		hx_field_150 := hx_obj_149["hasNext"]
		if hx_field_150 == nil {
			var hx_zero_151 func() bool
			return hx_zero_151
		}
		return hx_field_150.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_152 map[string]any) func() *string {
			hx_field_153 := hx_obj_152["next"]
			if hx_field_153 == nil {
				var hx_zero_154 func() *string
				return hx_zero_154
			}
			return hx_field_153.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_155 any) map[string]any {
		if hx_value_155 == nil {
			var hx_zero_156 map[string]any
			return hx_zero_156
		}
		return hx_value_155.(map[string]any)
	}(self.keys())
	for func(hx_obj_157 map[string]any) func() bool {
		hx_field_158 := hx_obj_157["hasNext"]
		if hx_field_158 == nil {
			var hx_zero_159 func() bool
			return hx_zero_159
		}
		return hx_field_158.(func() bool)
	}(iterator)() {
		key := func(hx_obj_160 map[string]any) func() *string {
			hx_field_161 := hx_obj_160["next"]
			if hx_field_161 == nil {
				var hx_zero_162 func() *string
				return hx_zero_162
			}
			return hx_field_161.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_163 map[string]any) func() bool {
			hx_field_164 := hx_obj_163["hasNext"]
			if hx_field_164 == nil {
				var hx_zero_165 func() bool
				return hx_zero_165
			}
			return hx_field_164.(func() bool)
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
