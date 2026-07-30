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
	hx_obj_235 := map[string]any{}
	hx_obj_235["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_235["next"] = func() *string {
		hx_post_236 := index
		index = int(int32((index + 1)))
		return keys[hx_post_236]
	}
	return hx_obj_235
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_237 := map[string]any{}
	hx_obj_237["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_237["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_238 := index
			index = int(int32((index + 1)))
			return hx_post_238
		}()])
	}
	return hx_obj_237
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_239 any) map[string]any {
		if hx_value_239 == nil {
			var hx_zero_240 map[string]any
			return hx_zero_240
		}
		return hx_value_239.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_241 := map[string]any{}
	hx_obj_241["hasNext"] = func() bool {
		return func(hx_obj_242 map[string]any) func() bool {
			hx_field_243 := hx_obj_242["hasNext"]
			if hx_field_243 == nil {
				var hx_zero_244 func() bool
				return hx_zero_244
			}
			return hx_field_243.(func() bool)
		}(keys)()
	}
	hx_obj_241["next"] = func() map[string]any {
		key := func(hx_obj_245 map[string]any) func() *string {
			hx_field_246 := hx_obj_245["next"]
			if hx_field_246 == nil {
				var hx_zero_247 func() *string
				return hx_zero_247
			}
			return hx_field_246.(func() *string)
		}(keys)()
		hx_obj_248 := map[string]any{}
		hx_obj_248["key"] = key
		hx_obj_248["value"] = _gthis.__hx_this.get(key)
		return hx_obj_248
	}
	return hx_obj_241
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_249 any) *string {
		if hx_value_249 == nil {
			var hx_zero_250 *string
			return hx_zero_250
		}
		return hx_value_249.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_251 any) *string {
		if hx_value_251 == nil {
			var hx_zero_252 *string
			return hx_zero_252
		}
		return hx_value_251.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_255 any) bool {
		if hx_value_255 == nil {
			var hx_zero_256 bool
			return hx_zero_256
		}
		return hx_value_255.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_253 any) *string {
		if hx_value_253 == nil {
			var hx_zero_254 *string
			return hx_zero_254
		}
		return hx_value_253.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_259 any) bool {
		if hx_value_259 == nil {
			var hx_zero_260 bool
			return hx_zero_260
		}
		return hx_value_259.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_257 any) *string {
		if hx_value_257 == nil {
			var hx_zero_258 *string
			return hx_zero_258
		}
		return hx_value_257.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_261 any) *haxe__ds__StringMap {
		if hx_value_261 == nil {
			var hx_zero_262 *haxe__ds__StringMap
			return hx_zero_262
		}
		return hx_value_261.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_263 any) map[string]any {
		if hx_value_263 == nil {
			var hx_zero_264 map[string]any
			return hx_zero_264
		}
		return hx_value_263.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_265 map[string]any) func() bool {
		hx_field_266 := hx_obj_265["hasNext"]
		if hx_field_266 == nil {
			var hx_zero_267 func() bool
			return hx_zero_267
		}
		return hx_field_266.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_268 map[string]any) func() *string {
			hx_field_269 := hx_obj_268["next"]
			if hx_field_269 == nil {
				var hx_zero_270 func() *string
				return hx_zero_270
			}
			return hx_field_269.(func() *string)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_271 any) map[string]any {
		if hx_value_271 == nil {
			var hx_zero_272 map[string]any
			return hx_zero_272
		}
		return hx_value_271.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_273 map[string]any) func() bool {
		hx_field_274 := hx_obj_273["hasNext"]
		if hx_field_274 == nil {
			var hx_zero_275 func() bool
			return hx_zero_275
		}
		return hx_field_274.(func() bool)
	}(iterator)() {
		key := func(hx_obj_276 map[string]any) func() *string {
			hx_field_277 := hx_obj_276["next"]
			if hx_field_277 == nil {
				var hx_zero_278 func() *string
				return hx_zero_278
			}
			return hx_field_277.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_279 map[string]any) func() bool {
			hx_field_280 := hx_obj_279["hasNext"]
			if hx_field_280 == nil {
				var hx_zero_281 func() bool
				return hx_zero_281
			}
			return hx_field_280.(func() bool)
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
