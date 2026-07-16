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
	hx_obj_152 := map[string]any{}
	hx_obj_152["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_152["next"] = func() *string {
		hx_post_153 := index
		index = int(int32((index + 1)))
		return keys[hx_post_153]
	}
	return hx_obj_152
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_154 := map[string]any{}
	hx_obj_154["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_154["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_155 := index
			index = int(int32((index + 1)))
			return hx_post_155
		}()])
	}
	return hx_obj_154
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_156 any) map[string]any {
		if hx_value_156 == nil {
			var hx_zero_157 map[string]any
			return hx_zero_157
		}
		return hx_value_156.(map[string]any)
	}(self.keys())
	hx_obj_158 := map[string]any{}
	hx_obj_158["hasNext"] = func() bool {
		return func(hx_obj_159 map[string]any) func() bool {
			hx_field_160 := hx_obj_159["hasNext"]
			if hx_field_160 == nil {
				var hx_zero_161 func() bool
				return hx_zero_161
			}
			return hx_field_160.(func() bool)
		}(keys)()
	}
	hx_obj_158["next"] = func() map[string]any {
		key := func(hx_obj_162 map[string]any) func() *string {
			hx_field_163 := hx_obj_162["next"]
			if hx_field_163 == nil {
				var hx_zero_164 func() *string
				return hx_zero_164
			}
			return hx_field_163.(func() *string)
		}(keys)()
		hx_obj_165 := map[string]any{}
		hx_obj_165["key"] = key
		hx_obj_165["value"] = _gthis.get(key)
		return hx_obj_165
	}
	return hx_obj_158
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_166 any) *string {
		if hx_value_166 == nil {
			var hx_zero_167 *string
			return hx_zero_167
		}
		return hx_value_166.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_168 any) *string {
		if hx_value_168 == nil {
			var hx_zero_169 *string
			return hx_zero_169
		}
		return hx_value_168.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_172 any) bool {
		if hx_value_172 == nil {
			var hx_zero_173 bool
			return hx_zero_173
		}
		return hx_value_172.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_170 any) *string {
		if hx_value_170 == nil {
			var hx_zero_171 *string
			return hx_zero_171
		}
		return hx_value_170.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_176 any) bool {
		if hx_value_176 == nil {
			var hx_zero_177 bool
			return hx_zero_177
		}
		return hx_value_176.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_174 any) *string {
		if hx_value_174 == nil {
			var hx_zero_175 *string
			return hx_zero_175
		}
		return hx_value_174.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_178 any) *haxe__ds__StringMap {
		if hx_value_178 == nil {
			var hx_zero_179 *haxe__ds__StringMap
			return hx_zero_179
		}
		return hx_value_178.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_180 any) map[string]any {
		if hx_value_180 == nil {
			var hx_zero_181 map[string]any
			return hx_zero_181
		}
		return hx_value_180.(map[string]any)
	}(self.keys())
	for func(hx_obj_182 map[string]any) func() bool {
		hx_field_183 := hx_obj_182["hasNext"]
		if hx_field_183 == nil {
			var hx_zero_184 func() bool
			return hx_zero_184
		}
		return hx_field_183.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_185 map[string]any) func() *string {
			hx_field_186 := hx_obj_185["next"]
			if hx_field_186 == nil {
				var hx_zero_187 func() *string
				return hx_zero_187
			}
			return hx_field_186.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_188 any) map[string]any {
		if hx_value_188 == nil {
			var hx_zero_189 map[string]any
			return hx_zero_189
		}
		return hx_value_188.(map[string]any)
	}(self.keys())
	for func(hx_obj_190 map[string]any) func() bool {
		hx_field_191 := hx_obj_190["hasNext"]
		if hx_field_191 == nil {
			var hx_zero_192 func() bool
			return hx_zero_192
		}
		return hx_field_191.(func() bool)
	}(iterator)() {
		key := func(hx_obj_193 map[string]any) func() *string {
			hx_field_194 := hx_obj_193["next"]
			if hx_field_194 == nil {
				var hx_zero_195 func() *string
				return hx_zero_195
			}
			return hx_field_194.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_196 map[string]any) func() bool {
			hx_field_197 := hx_obj_196["hasNext"]
			if hx_field_197 == nil {
				var hx_zero_198 func() bool
				return hx_zero_198
			}
			return hx_field_197.(func() bool)
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
