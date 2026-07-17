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
	hx_obj_187 := map[string]any{}
	hx_obj_187["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_187["next"] = func() *string {
		hx_post_188 := index
		index = int(int32((index + 1)))
		return keys[hx_post_188]
	}
	return hx_obj_187
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_189 := map[string]any{}
	hx_obj_189["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_189["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_190 := index
			index = int(int32((index + 1)))
			return hx_post_190
		}()])
	}
	return hx_obj_189
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_191 any) map[string]any {
		if hx_value_191 == nil {
			var hx_zero_192 map[string]any
			return hx_zero_192
		}
		return hx_value_191.(map[string]any)
	}(self.keys())
	hx_obj_193 := map[string]any{}
	hx_obj_193["hasNext"] = func() bool {
		return func(hx_obj_194 map[string]any) func() bool {
			hx_field_195 := hx_obj_194["hasNext"]
			if hx_field_195 == nil {
				var hx_zero_196 func() bool
				return hx_zero_196
			}
			return hx_field_195.(func() bool)
		}(keys)()
	}
	hx_obj_193["next"] = func() map[string]any {
		key := func(hx_obj_197 map[string]any) func() *string {
			hx_field_198 := hx_obj_197["next"]
			if hx_field_198 == nil {
				var hx_zero_199 func() *string
				return hx_zero_199
			}
			return hx_field_198.(func() *string)
		}(keys)()
		hx_obj_200 := map[string]any{}
		hx_obj_200["key"] = key
		hx_obj_200["value"] = _gthis.get(key)
		return hx_obj_200
	}
	return hx_obj_193
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_201 any) *string {
		if hx_value_201 == nil {
			var hx_zero_202 *string
			return hx_zero_202
		}
		return hx_value_201.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_203 any) *string {
		if hx_value_203 == nil {
			var hx_zero_204 *string
			return hx_zero_204
		}
		return hx_value_203.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_207 any) bool {
		if hx_value_207 == nil {
			var hx_zero_208 bool
			return hx_zero_208
		}
		return hx_value_207.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_205 any) *string {
		if hx_value_205 == nil {
			var hx_zero_206 *string
			return hx_zero_206
		}
		return hx_value_205.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_211 any) bool {
		if hx_value_211 == nil {
			var hx_zero_212 bool
			return hx_zero_212
		}
		return hx_value_211.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_209 any) *string {
		if hx_value_209 == nil {
			var hx_zero_210 *string
			return hx_zero_210
		}
		return hx_value_209.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_213 any) *haxe__ds__StringMap {
		if hx_value_213 == nil {
			var hx_zero_214 *haxe__ds__StringMap
			return hx_zero_214
		}
		return hx_value_213.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_215 any) map[string]any {
		if hx_value_215 == nil {
			var hx_zero_216 map[string]any
			return hx_zero_216
		}
		return hx_value_215.(map[string]any)
	}(self.keys())
	for func(hx_obj_217 map[string]any) func() bool {
		hx_field_218 := hx_obj_217["hasNext"]
		if hx_field_218 == nil {
			var hx_zero_219 func() bool
			return hx_zero_219
		}
		return hx_field_218.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_220 map[string]any) func() *string {
			hx_field_221 := hx_obj_220["next"]
			if hx_field_221 == nil {
				var hx_zero_222 func() *string
				return hx_zero_222
			}
			return hx_field_221.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_223 any) map[string]any {
		if hx_value_223 == nil {
			var hx_zero_224 map[string]any
			return hx_zero_224
		}
		return hx_value_223.(map[string]any)
	}(self.keys())
	for func(hx_obj_225 map[string]any) func() bool {
		hx_field_226 := hx_obj_225["hasNext"]
		if hx_field_226 == nil {
			var hx_zero_227 func() bool
			return hx_zero_227
		}
		return hx_field_226.(func() bool)
	}(iterator)() {
		key := func(hx_obj_228 map[string]any) func() *string {
			hx_field_229 := hx_obj_228["next"]
			if hx_field_229 == nil {
				var hx_zero_230 func() *string
				return hx_zero_230
			}
			return hx_field_229.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_231 map[string]any) func() bool {
			hx_field_232 := hx_obj_231["hasNext"]
			if hx_field_232 == nil {
				var hx_zero_233 func() bool
				return hx_zero_233
			}
			return hx_field_232.(func() bool)
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
