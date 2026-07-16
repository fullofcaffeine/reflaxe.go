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
	hx_obj_157 := map[string]any{}
	hx_obj_157["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_157["next"] = func() *string {
		hx_post_158 := index
		index = int(int32((index + 1)))
		return keys[hx_post_158]
	}
	return hx_obj_157
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_159 := map[string]any{}
	hx_obj_159["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_159["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_160 := index
			index = int(int32((index + 1)))
			return hx_post_160
		}()])
	}
	return hx_obj_159
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_161 any) map[string]any {
		if hx_value_161 == nil {
			var hx_zero_162 map[string]any
			return hx_zero_162
		}
		return hx_value_161.(map[string]any)
	}(self.keys())
	hx_obj_163 := map[string]any{}
	hx_obj_163["hasNext"] = func() bool {
		return func(hx_obj_164 map[string]any) func() bool {
			hx_field_165 := hx_obj_164["hasNext"]
			if hx_field_165 == nil {
				var hx_zero_166 func() bool
				return hx_zero_166
			}
			return hx_field_165.(func() bool)
		}(keys)()
	}
	hx_obj_163["next"] = func() map[string]any {
		key := func(hx_obj_167 map[string]any) func() *string {
			hx_field_168 := hx_obj_167["next"]
			if hx_field_168 == nil {
				var hx_zero_169 func() *string
				return hx_zero_169
			}
			return hx_field_168.(func() *string)
		}(keys)()
		hx_obj_170 := map[string]any{}
		hx_obj_170["key"] = key
		hx_obj_170["value"] = _gthis.get(key)
		return hx_obj_170
	}
	return hx_obj_163
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_171 any) *string {
		if hx_value_171 == nil {
			var hx_zero_172 *string
			return hx_zero_172
		}
		return hx_value_171.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_173 any) *string {
		if hx_value_173 == nil {
			var hx_zero_174 *string
			return hx_zero_174
		}
		return hx_value_173.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_177 any) bool {
		if hx_value_177 == nil {
			var hx_zero_178 bool
			return hx_zero_178
		}
		return hx_value_177.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_175 any) *string {
		if hx_value_175 == nil {
			var hx_zero_176 *string
			return hx_zero_176
		}
		return hx_value_175.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_181 any) bool {
		if hx_value_181 == nil {
			var hx_zero_182 bool
			return hx_zero_182
		}
		return hx_value_181.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_179 any) *string {
		if hx_value_179 == nil {
			var hx_zero_180 *string
			return hx_zero_180
		}
		return hx_value_179.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_183 any) *haxe__ds__StringMap {
		if hx_value_183 == nil {
			var hx_zero_184 *haxe__ds__StringMap
			return hx_zero_184
		}
		return hx_value_183.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_185 any) map[string]any {
		if hx_value_185 == nil {
			var hx_zero_186 map[string]any
			return hx_zero_186
		}
		return hx_value_185.(map[string]any)
	}(self.keys())
	for func(hx_obj_187 map[string]any) func() bool {
		hx_field_188 := hx_obj_187["hasNext"]
		if hx_field_188 == nil {
			var hx_zero_189 func() bool
			return hx_zero_189
		}
		return hx_field_188.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_190 map[string]any) func() *string {
			hx_field_191 := hx_obj_190["next"]
			if hx_field_191 == nil {
				var hx_zero_192 func() *string
				return hx_zero_192
			}
			return hx_field_191.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_193 any) map[string]any {
		if hx_value_193 == nil {
			var hx_zero_194 map[string]any
			return hx_zero_194
		}
		return hx_value_193.(map[string]any)
	}(self.keys())
	for func(hx_obj_195 map[string]any) func() bool {
		hx_field_196 := hx_obj_195["hasNext"]
		if hx_field_196 == nil {
			var hx_zero_197 func() bool
			return hx_zero_197
		}
		return hx_field_196.(func() bool)
	}(iterator)() {
		key := func(hx_obj_198 map[string]any) func() *string {
			hx_field_199 := hx_obj_198["next"]
			if hx_field_199 == nil {
				var hx_zero_200 func() *string
				return hx_zero_200
			}
			return hx_field_199.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_201 map[string]any) func() bool {
			hx_field_202 := hx_obj_201["hasNext"]
			if hx_field_202 == nil {
				var hx_zero_203 func() bool
				return hx_zero_203
			}
			return hx_field_202.(func() bool)
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
