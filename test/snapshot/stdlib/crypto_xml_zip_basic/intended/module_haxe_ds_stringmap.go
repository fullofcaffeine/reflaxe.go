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
	hx_obj_138 := map[string]any{}
	hx_obj_138["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_138["next"] = func() *string {
		hx_post_139 := index
		index = int(int32((index + 1)))
		return keys[hx_post_139]
	}
	return hx_obj_138
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_140 := map[string]any{}
	hx_obj_140["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_140["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_141 := index
			index = int(int32((index + 1)))
			return hx_post_141
		}()])
	}
	return hx_obj_140
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_142 any) map[string]any {
		if hx_value_142 == nil {
			var hx_zero_143 map[string]any
			return hx_zero_143
		}
		return hx_value_142.(map[string]any)
	}(self.keys())
	hx_obj_144 := map[string]any{}
	hx_obj_144["hasNext"] = func() bool {
		return func(hx_obj_145 map[string]any) func() bool {
			hx_field_146 := hx_obj_145["hasNext"]
			if hx_field_146 == nil {
				var hx_zero_147 func() bool
				return hx_zero_147
			}
			return hx_field_146.(func() bool)
		}(keys)()
	}
	hx_obj_144["next"] = func() map[string]any {
		key := func(hx_obj_148 map[string]any) func() *string {
			hx_field_149 := hx_obj_148["next"]
			if hx_field_149 == nil {
				var hx_zero_150 func() *string
				return hx_zero_150
			}
			return hx_field_149.(func() *string)
		}(keys)()
		hx_obj_151 := map[string]any{}
		hx_obj_151["key"] = key
		hx_obj_151["value"] = _gthis.get(key)
		return hx_obj_151
	}
	return hx_obj_144
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_152 any) *string {
		if hx_value_152 == nil {
			var hx_zero_153 *string
			return hx_zero_153
		}
		return hx_value_152.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_154 any) *string {
		if hx_value_154 == nil {
			var hx_zero_155 *string
			return hx_zero_155
		}
		return hx_value_154.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_158 any) bool {
		if hx_value_158 == nil {
			var hx_zero_159 bool
			return hx_zero_159
		}
		return hx_value_158.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_156 any) *string {
		if hx_value_156 == nil {
			var hx_zero_157 *string
			return hx_zero_157
		}
		return hx_value_156.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_162 any) bool {
		if hx_value_162 == nil {
			var hx_zero_163 bool
			return hx_zero_163
		}
		return hx_value_162.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_160 any) *string {
		if hx_value_160 == nil {
			var hx_zero_161 *string
			return hx_zero_161
		}
		return hx_value_160.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_164 any) *haxe__ds__StringMap {
		if hx_value_164 == nil {
			var hx_zero_165 *haxe__ds__StringMap
			return hx_zero_165
		}
		return hx_value_164.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_166 any) map[string]any {
		if hx_value_166 == nil {
			var hx_zero_167 map[string]any
			return hx_zero_167
		}
		return hx_value_166.(map[string]any)
	}(self.keys())
	for func(hx_obj_168 map[string]any) func() bool {
		hx_field_169 := hx_obj_168["hasNext"]
		if hx_field_169 == nil {
			var hx_zero_170 func() bool
			return hx_zero_170
		}
		return hx_field_169.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_171 map[string]any) func() *string {
			hx_field_172 := hx_obj_171["next"]
			if hx_field_172 == nil {
				var hx_zero_173 func() *string
				return hx_zero_173
			}
			return hx_field_172.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_174 any) map[string]any {
		if hx_value_174 == nil {
			var hx_zero_175 map[string]any
			return hx_zero_175
		}
		return hx_value_174.(map[string]any)
	}(self.keys())
	for func(hx_obj_176 map[string]any) func() bool {
		hx_field_177 := hx_obj_176["hasNext"]
		if hx_field_177 == nil {
			var hx_zero_178 func() bool
			return hx_zero_178
		}
		return hx_field_177.(func() bool)
	}(iterator)() {
		key := func(hx_obj_179 map[string]any) func() *string {
			hx_field_180 := hx_obj_179["next"]
			if hx_field_180 == nil {
				var hx_zero_181 func() *string
				return hx_zero_181
			}
			return hx_field_180.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_182 map[string]any) func() bool {
			hx_field_183 := hx_obj_182["hasNext"]
			if hx_field_183 == nil {
				var hx_zero_184 func() bool
				return hx_zero_184
			}
			return hx_field_183.(func() bool)
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
