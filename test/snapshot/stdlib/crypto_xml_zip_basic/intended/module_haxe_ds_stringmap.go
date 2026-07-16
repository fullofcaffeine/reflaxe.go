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
	hx_obj_167 := map[string]any{}
	hx_obj_167["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_167["next"] = func() *string {
		hx_post_168 := index
		index = int(int32((index + 1)))
		return keys[hx_post_168]
	}
	return hx_obj_167
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_169 := map[string]any{}
	hx_obj_169["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_169["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_170 := index
			index = int(int32((index + 1)))
			return hx_post_170
		}()])
	}
	return hx_obj_169
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_171 any) map[string]any {
		if hx_value_171 == nil {
			var hx_zero_172 map[string]any
			return hx_zero_172
		}
		return hx_value_171.(map[string]any)
	}(self.keys())
	hx_obj_173 := map[string]any{}
	hx_obj_173["hasNext"] = func() bool {
		return func(hx_obj_174 map[string]any) func() bool {
			hx_field_175 := hx_obj_174["hasNext"]
			if hx_field_175 == nil {
				var hx_zero_176 func() bool
				return hx_zero_176
			}
			return hx_field_175.(func() bool)
		}(keys)()
	}
	hx_obj_173["next"] = func() map[string]any {
		key := func(hx_obj_177 map[string]any) func() *string {
			hx_field_178 := hx_obj_177["next"]
			if hx_field_178 == nil {
				var hx_zero_179 func() *string
				return hx_zero_179
			}
			return hx_field_178.(func() *string)
		}(keys)()
		hx_obj_180 := map[string]any{}
		hx_obj_180["key"] = key
		hx_obj_180["value"] = _gthis.get(key)
		return hx_obj_180
	}
	return hx_obj_173
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_181 any) *string {
		if hx_value_181 == nil {
			var hx_zero_182 *string
			return hx_zero_182
		}
		return hx_value_181.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_183 any) *string {
		if hx_value_183 == nil {
			var hx_zero_184 *string
			return hx_zero_184
		}
		return hx_value_183.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_187 any) bool {
		if hx_value_187 == nil {
			var hx_zero_188 bool
			return hx_zero_188
		}
		return hx_value_187.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_185 any) *string {
		if hx_value_185 == nil {
			var hx_zero_186 *string
			return hx_zero_186
		}
		return hx_value_185.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_191 any) bool {
		if hx_value_191 == nil {
			var hx_zero_192 bool
			return hx_zero_192
		}
		return hx_value_191.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_189 any) *string {
		if hx_value_189 == nil {
			var hx_zero_190 *string
			return hx_zero_190
		}
		return hx_value_189.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_193 any) *haxe__ds__StringMap {
		if hx_value_193 == nil {
			var hx_zero_194 *haxe__ds__StringMap
			return hx_zero_194
		}
		return hx_value_193.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_195 any) map[string]any {
		if hx_value_195 == nil {
			var hx_zero_196 map[string]any
			return hx_zero_196
		}
		return hx_value_195.(map[string]any)
	}(self.keys())
	for func(hx_obj_197 map[string]any) func() bool {
		hx_field_198 := hx_obj_197["hasNext"]
		if hx_field_198 == nil {
			var hx_zero_199 func() bool
			return hx_zero_199
		}
		return hx_field_198.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_200 map[string]any) func() *string {
			hx_field_201 := hx_obj_200["next"]
			if hx_field_201 == nil {
				var hx_zero_202 func() *string
				return hx_zero_202
			}
			return hx_field_201.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_203 any) map[string]any {
		if hx_value_203 == nil {
			var hx_zero_204 map[string]any
			return hx_zero_204
		}
		return hx_value_203.(map[string]any)
	}(self.keys())
	for func(hx_obj_205 map[string]any) func() bool {
		hx_field_206 := hx_obj_205["hasNext"]
		if hx_field_206 == nil {
			var hx_zero_207 func() bool
			return hx_zero_207
		}
		return hx_field_206.(func() bool)
	}(iterator)() {
		key := func(hx_obj_208 map[string]any) func() *string {
			hx_field_209 := hx_obj_208["next"]
			if hx_field_209 == nil {
				var hx_zero_210 func() *string
				return hx_zero_210
			}
			return hx_field_209.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_211 map[string]any) func() bool {
			hx_field_212 := hx_obj_211["hasNext"]
			if hx_field_212 == nil {
				var hx_zero_213 func() bool
				return hx_zero_213
			}
			return hx_field_212.(func() bool)
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
