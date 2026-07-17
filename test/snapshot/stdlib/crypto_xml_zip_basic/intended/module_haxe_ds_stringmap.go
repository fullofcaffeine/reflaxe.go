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
	hx_obj_126 := map[string]any{}
	hx_obj_126["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_126["next"] = func() *string {
		hx_post_127 := index
		index = int(int32((index + 1)))
		return keys[hx_post_127]
	}
	return hx_obj_126
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_128 := map[string]any{}
	hx_obj_128["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_128["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_129 := index
			index = int(int32((index + 1)))
			return hx_post_129
		}()])
	}
	return hx_obj_128
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_130 any) map[string]any {
		if hx_value_130 == nil {
			var hx_zero_131 map[string]any
			return hx_zero_131
		}
		return hx_value_130.(map[string]any)
	}(self.keys())
	hx_obj_132 := map[string]any{}
	hx_obj_132["hasNext"] = func() bool {
		return func(hx_obj_133 map[string]any) func() bool {
			hx_field_134 := hx_obj_133["hasNext"]
			if hx_field_134 == nil {
				var hx_zero_135 func() bool
				return hx_zero_135
			}
			return hx_field_134.(func() bool)
		}(keys)()
	}
	hx_obj_132["next"] = func() map[string]any {
		key := func(hx_obj_136 map[string]any) func() *string {
			hx_field_137 := hx_obj_136["next"]
			if hx_field_137 == nil {
				var hx_zero_138 func() *string
				return hx_zero_138
			}
			return hx_field_137.(func() *string)
		}(keys)()
		hx_obj_139 := map[string]any{}
		hx_obj_139["key"] = key
		hx_obj_139["value"] = _gthis.get(key)
		return hx_obj_139
	}
	return hx_obj_132
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_140 any) *string {
		if hx_value_140 == nil {
			var hx_zero_141 *string
			return hx_zero_141
		}
		return hx_value_140.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_142 any) *string {
		if hx_value_142 == nil {
			var hx_zero_143 *string
			return hx_zero_143
		}
		return hx_value_142.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_146 any) bool {
		if hx_value_146 == nil {
			var hx_zero_147 bool
			return hx_zero_147
		}
		return hx_value_146.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_144 any) *string {
		if hx_value_144 == nil {
			var hx_zero_145 *string
			return hx_zero_145
		}
		return hx_value_144.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_150 any) bool {
		if hx_value_150 == nil {
			var hx_zero_151 bool
			return hx_zero_151
		}
		return hx_value_150.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_148 any) *string {
		if hx_value_148 == nil {
			var hx_zero_149 *string
			return hx_zero_149
		}
		return hx_value_148.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_152 any) *haxe__ds__StringMap {
		if hx_value_152 == nil {
			var hx_zero_153 *haxe__ds__StringMap
			return hx_zero_153
		}
		return hx_value_152.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_154 any) map[string]any {
		if hx_value_154 == nil {
			var hx_zero_155 map[string]any
			return hx_zero_155
		}
		return hx_value_154.(map[string]any)
	}(self.keys())
	for func(hx_obj_156 map[string]any) func() bool {
		hx_field_157 := hx_obj_156["hasNext"]
		if hx_field_157 == nil {
			var hx_zero_158 func() bool
			return hx_zero_158
		}
		return hx_field_157.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_159 map[string]any) func() *string {
			hx_field_160 := hx_obj_159["next"]
			if hx_field_160 == nil {
				var hx_zero_161 func() *string
				return hx_zero_161
			}
			return hx_field_160.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_162 any) map[string]any {
		if hx_value_162 == nil {
			var hx_zero_163 map[string]any
			return hx_zero_163
		}
		return hx_value_162.(map[string]any)
	}(self.keys())
	for func(hx_obj_164 map[string]any) func() bool {
		hx_field_165 := hx_obj_164["hasNext"]
		if hx_field_165 == nil {
			var hx_zero_166 func() bool
			return hx_zero_166
		}
		return hx_field_165.(func() bool)
	}(iterator)() {
		key := func(hx_obj_167 map[string]any) func() *string {
			hx_field_168 := hx_obj_167["next"]
			if hx_field_168 == nil {
				var hx_zero_169 func() *string
				return hx_zero_169
			}
			return hx_field_168.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_170 map[string]any) func() bool {
			hx_field_171 := hx_obj_170["hasNext"]
			if hx_field_171 == nil {
				var hx_zero_172 func() bool
				return hx_zero_172
			}
			return hx_field_171.(func() bool)
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
