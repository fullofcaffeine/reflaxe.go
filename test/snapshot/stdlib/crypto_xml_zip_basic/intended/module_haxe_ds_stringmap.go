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
	hx_obj_166 := map[string]any{}
	hx_obj_166["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_166["next"] = func() *string {
		hx_post_167 := index
		index = int(int32((index + 1)))
		return keys[hx_post_167]
	}
	return hx_obj_166
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_168 := map[string]any{}
	hx_obj_168["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_168["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_169 := index
			index = int(int32((index + 1)))
			return hx_post_169
		}()])
	}
	return hx_obj_168
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_170 any) map[string]any {
		if hx_value_170 == nil {
			var hx_zero_171 map[string]any
			return hx_zero_171
		}
		return hx_value_170.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_172 := map[string]any{}
	hx_obj_172["hasNext"] = func() bool {
		return func(hx_obj_173 map[string]any) func() bool {
			hx_field_174 := hx_obj_173["hasNext"]
			if hx_field_174 == nil {
				var hx_zero_175 func() bool
				return hx_zero_175
			}
			return hx_field_174.(func() bool)
		}(keys)()
	}
	hx_obj_172["next"] = func() map[string]any {
		key := func(hx_obj_176 map[string]any) func() *string {
			hx_field_177 := hx_obj_176["next"]
			if hx_field_177 == nil {
				var hx_zero_178 func() *string
				return hx_zero_178
			}
			return hx_field_177.(func() *string)
		}(keys)()
		hx_obj_179 := map[string]any{}
		hx_obj_179["key"] = key
		hx_obj_179["value"] = _gthis.__hx_this.get(key)
		return hx_obj_179
	}
	return hx_obj_172
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_180 any) *string {
		if hx_value_180 == nil {
			var hx_zero_181 *string
			return hx_zero_181
		}
		return hx_value_180.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_182 any) *string {
		if hx_value_182 == nil {
			var hx_zero_183 *string
			return hx_zero_183
		}
		return hx_value_182.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_186 any) bool {
		if hx_value_186 == nil {
			var hx_zero_187 bool
			return hx_zero_187
		}
		return hx_value_186.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_184 any) *string {
		if hx_value_184 == nil {
			var hx_zero_185 *string
			return hx_zero_185
		}
		return hx_value_184.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_190 any) bool {
		if hx_value_190 == nil {
			var hx_zero_191 bool
			return hx_zero_191
		}
		return hx_value_190.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_188 any) *string {
		if hx_value_188 == nil {
			var hx_zero_189 *string
			return hx_zero_189
		}
		return hx_value_188.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_192 any) *haxe__ds__StringMap {
		if hx_value_192 == nil {
			var hx_zero_193 *haxe__ds__StringMap
			return hx_zero_193
		}
		return hx_value_192.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_194 any) map[string]any {
		if hx_value_194 == nil {
			var hx_zero_195 map[string]any
			return hx_zero_195
		}
		return hx_value_194.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_196 map[string]any) func() bool {
		hx_field_197 := hx_obj_196["hasNext"]
		if hx_field_197 == nil {
			var hx_zero_198 func() bool
			return hx_zero_198
		}
		return hx_field_197.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_199 map[string]any) func() *string {
			hx_field_200 := hx_obj_199["next"]
			if hx_field_200 == nil {
				var hx_zero_201 func() *string
				return hx_zero_201
			}
			return hx_field_200.(func() *string)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_202 any) map[string]any {
		if hx_value_202 == nil {
			var hx_zero_203 map[string]any
			return hx_zero_203
		}
		return hx_value_202.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_204 map[string]any) func() bool {
		hx_field_205 := hx_obj_204["hasNext"]
		if hx_field_205 == nil {
			var hx_zero_206 func() bool
			return hx_zero_206
		}
		return hx_field_205.(func() bool)
	}(iterator)() {
		key := func(hx_obj_207 map[string]any) func() *string {
			hx_field_208 := hx_obj_207["next"]
			if hx_field_208 == nil {
				var hx_zero_209 func() *string
				return hx_zero_209
			}
			return hx_field_208.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_210 map[string]any) func() bool {
			hx_field_211 := hx_obj_210["hasNext"]
			if hx_field_211 == nil {
				var hx_zero_212 func() bool
				return hx_zero_212
			}
			return hx_field_211.(func() bool)
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

func (self *haxe__ds__StringMap) String() string {
	return *self.__hx_this.toString()
}
