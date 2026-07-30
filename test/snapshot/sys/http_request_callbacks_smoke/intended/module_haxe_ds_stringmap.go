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
	hx_obj_229 := map[string]any{}
	hx_obj_229["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_229["next"] = func() *string {
		hx_post_230 := index
		index = int(int32((index + 1)))
		return keys[hx_post_230]
	}
	return hx_obj_229
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_231 := map[string]any{}
	hx_obj_231["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_231["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_232 := index
			index = int(int32((index + 1)))
			return hx_post_232
		}()])
	}
	return hx_obj_231
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_233 any) map[string]any {
		if hx_value_233 == nil {
			var hx_zero_234 map[string]any
			return hx_zero_234
		}
		return hx_value_233.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_235 := map[string]any{}
	hx_obj_235["hasNext"] = func() bool {
		return func(hx_obj_236 map[string]any) func() bool {
			hx_field_237 := hx_obj_236["hasNext"]
			if hx_field_237 == nil {
				var hx_zero_238 func() bool
				return hx_zero_238
			}
			return hx_field_237.(func() bool)
		}(keys)()
	}
	hx_obj_235["next"] = func() map[string]any {
		key := func(hx_obj_239 map[string]any) func() *string {
			hx_field_240 := hx_obj_239["next"]
			if hx_field_240 == nil {
				var hx_zero_241 func() *string
				return hx_zero_241
			}
			return hx_field_240.(func() *string)
		}(keys)()
		hx_obj_242 := map[string]any{}
		hx_obj_242["key"] = key
		hx_obj_242["value"] = _gthis.__hx_this.get(key)
		return hx_obj_242
	}
	return hx_obj_235
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_243 any) *string {
		if hx_value_243 == nil {
			var hx_zero_244 *string
			return hx_zero_244
		}
		return hx_value_243.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_245 any) *string {
		if hx_value_245 == nil {
			var hx_zero_246 *string
			return hx_zero_246
		}
		return hx_value_245.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_249 any) bool {
		if hx_value_249 == nil {
			var hx_zero_250 bool
			return hx_zero_250
		}
		return hx_value_249.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_247 any) *string {
		if hx_value_247 == nil {
			var hx_zero_248 *string
			return hx_zero_248
		}
		return hx_value_247.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_253 any) bool {
		if hx_value_253 == nil {
			var hx_zero_254 bool
			return hx_zero_254
		}
		return hx_value_253.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_251 any) *string {
		if hx_value_251 == nil {
			var hx_zero_252 *string
			return hx_zero_252
		}
		return hx_value_251.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_255 any) *haxe__ds__StringMap {
		if hx_value_255 == nil {
			var hx_zero_256 *haxe__ds__StringMap
			return hx_zero_256
		}
		return hx_value_255.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_257 any) map[string]any {
		if hx_value_257 == nil {
			var hx_zero_258 map[string]any
			return hx_zero_258
		}
		return hx_value_257.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_259 map[string]any) func() bool {
		hx_field_260 := hx_obj_259["hasNext"]
		if hx_field_260 == nil {
			var hx_zero_261 func() bool
			return hx_zero_261
		}
		return hx_field_260.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_262 map[string]any) func() *string {
			hx_field_263 := hx_obj_262["next"]
			if hx_field_263 == nil {
				var hx_zero_264 func() *string
				return hx_zero_264
			}
			return hx_field_263.(func() *string)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_265 any) map[string]any {
		if hx_value_265 == nil {
			var hx_zero_266 map[string]any
			return hx_zero_266
		}
		return hx_value_265.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_267 map[string]any) func() bool {
		hx_field_268 := hx_obj_267["hasNext"]
		if hx_field_268 == nil {
			var hx_zero_269 func() bool
			return hx_zero_269
		}
		return hx_field_268.(func() bool)
	}(iterator)() {
		key := func(hx_obj_270 map[string]any) func() *string {
			hx_field_271 := hx_obj_270["next"]
			if hx_field_271 == nil {
				var hx_zero_272 func() *string
				return hx_zero_272
			}
			return hx_field_271.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_273 map[string]any) func() bool {
			hx_field_274 := hx_obj_273["hasNext"]
			if hx_field_274 == nil {
				var hx_zero_275 func() bool
				return hx_zero_275
			}
			return hx_field_274.(func() bool)
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
