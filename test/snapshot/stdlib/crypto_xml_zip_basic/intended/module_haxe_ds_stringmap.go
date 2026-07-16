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
	hx_obj_142 := map[string]any{}
	hx_obj_142["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_142["next"] = func() *string {
		hx_post_143 := index
		index = int(int32((index + 1)))
		return keys[hx_post_143]
	}
	return hx_obj_142
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_144 := map[string]any{}
	hx_obj_144["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_144["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_145 := index
			index = int(int32((index + 1)))
			return hx_post_145
		}()])
	}
	return hx_obj_144
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_146 any) map[string]any {
		if hx_value_146 == nil {
			var hx_zero_147 map[string]any
			return hx_zero_147
		}
		return hx_value_146.(map[string]any)
	}(self.keys())
	hx_obj_148 := map[string]any{}
	hx_obj_148["hasNext"] = func() bool {
		return func(hx_obj_149 map[string]any) func() bool {
			hx_field_150 := hx_obj_149["hasNext"]
			if hx_field_150 == nil {
				var hx_zero_151 func() bool
				return hx_zero_151
			}
			return hx_field_150.(func() bool)
		}(keys)()
	}
	hx_obj_148["next"] = func() map[string]any {
		key := func(hx_obj_152 map[string]any) func() *string {
			hx_field_153 := hx_obj_152["next"]
			if hx_field_153 == nil {
				var hx_zero_154 func() *string
				return hx_zero_154
			}
			return hx_field_153.(func() *string)
		}(keys)()
		hx_obj_155 := map[string]any{}
		hx_obj_155["key"] = key
		hx_obj_155["value"] = _gthis.get(key)
		return hx_obj_155
	}
	return hx_obj_148
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_156 any) *string {
		if hx_value_156 == nil {
			var hx_zero_157 *string
			return hx_zero_157
		}
		return hx_value_156.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_158 any) *string {
		if hx_value_158 == nil {
			var hx_zero_159 *string
			return hx_zero_159
		}
		return hx_value_158.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_162 any) bool {
		if hx_value_162 == nil {
			var hx_zero_163 bool
			return hx_zero_163
		}
		return hx_value_162.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_160 any) *string {
		if hx_value_160 == nil {
			var hx_zero_161 *string
			return hx_zero_161
		}
		return hx_value_160.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_166 any) bool {
		if hx_value_166 == nil {
			var hx_zero_167 bool
			return hx_zero_167
		}
		return hx_value_166.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_164 any) *string {
		if hx_value_164 == nil {
			var hx_zero_165 *string
			return hx_zero_165
		}
		return hx_value_164.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_168 any) *haxe__ds__StringMap {
		if hx_value_168 == nil {
			var hx_zero_169 *haxe__ds__StringMap
			return hx_zero_169
		}
		return hx_value_168.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_170 any) map[string]any {
		if hx_value_170 == nil {
			var hx_zero_171 map[string]any
			return hx_zero_171
		}
		return hx_value_170.(map[string]any)
	}(self.keys())
	for func(hx_obj_172 map[string]any) func() bool {
		hx_field_173 := hx_obj_172["hasNext"]
		if hx_field_173 == nil {
			var hx_zero_174 func() bool
			return hx_zero_174
		}
		return hx_field_173.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_175 map[string]any) func() *string {
			hx_field_176 := hx_obj_175["next"]
			if hx_field_176 == nil {
				var hx_zero_177 func() *string
				return hx_zero_177
			}
			return hx_field_176.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_178 any) map[string]any {
		if hx_value_178 == nil {
			var hx_zero_179 map[string]any
			return hx_zero_179
		}
		return hx_value_178.(map[string]any)
	}(self.keys())
	for func(hx_obj_180 map[string]any) func() bool {
		hx_field_181 := hx_obj_180["hasNext"]
		if hx_field_181 == nil {
			var hx_zero_182 func() bool
			return hx_zero_182
		}
		return hx_field_181.(func() bool)
	}(iterator)() {
		key := func(hx_obj_183 map[string]any) func() *string {
			hx_field_184 := hx_obj_183["next"]
			if hx_field_184 == nil {
				var hx_zero_185 func() *string
				return hx_zero_185
			}
			return hx_field_184.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_186 map[string]any) func() bool {
			hx_field_187 := hx_obj_186["hasNext"]
			if hx_field_187 == nil {
				var hx_zero_188 func() bool
				return hx_zero_188
			}
			return hx_field_187.(func() bool)
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
