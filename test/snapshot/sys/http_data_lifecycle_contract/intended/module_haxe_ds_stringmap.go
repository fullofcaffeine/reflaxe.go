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
	hx_obj_239 := map[string]any{}
	hx_obj_239["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_239["next"] = func() *string {
		hx_post_240 := index
		index = int(int32((index + 1)))
		return keys[hx_post_240]
	}
	return hx_obj_239
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_241 := map[string]any{}
	hx_obj_241["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_241["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_242 := index
			index = int(int32((index + 1)))
			return hx_post_242
		}()])
	}
	return hx_obj_241
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_243 any) map[string]any {
		if hx_value_243 == nil {
			var hx_zero_244 map[string]any
			return hx_zero_244
		}
		return hx_value_243.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_245 := map[string]any{}
	hx_obj_245["hasNext"] = func() bool {
		return func(hx_obj_246 map[string]any) func() bool {
			hx_field_247 := hx_obj_246["hasNext"]
			if hx_field_247 == nil {
				var hx_zero_248 func() bool
				return hx_zero_248
			}
			return hx_field_247.(func() bool)
		}(keys)()
	}
	hx_obj_245["next"] = func() map[string]any {
		key := func(hx_obj_249 map[string]any) func() *string {
			hx_field_250 := hx_obj_249["next"]
			if hx_field_250 == nil {
				var hx_zero_251 func() *string
				return hx_zero_251
			}
			return hx_field_250.(func() *string)
		}(keys)()
		hx_obj_252 := map[string]any{}
		hx_obj_252["key"] = key
		hx_obj_252["value"] = _gthis.__hx_this.get(key)
		return hx_obj_252
	}
	return hx_obj_245
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_253 any) *string {
		if hx_value_253 == nil {
			var hx_zero_254 *string
			return hx_zero_254
		}
		return hx_value_253.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_255 any) *string {
		if hx_value_255 == nil {
			var hx_zero_256 *string
			return hx_zero_256
		}
		return hx_value_255.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_259 any) bool {
		if hx_value_259 == nil {
			var hx_zero_260 bool
			return hx_zero_260
		}
		return hx_value_259.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_257 any) *string {
		if hx_value_257 == nil {
			var hx_zero_258 *string
			return hx_zero_258
		}
		return hx_value_257.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_263 any) bool {
		if hx_value_263 == nil {
			var hx_zero_264 bool
			return hx_zero_264
		}
		return hx_value_263.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_261 any) *string {
		if hx_value_261 == nil {
			var hx_zero_262 *string
			return hx_zero_262
		}
		return hx_value_261.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_265 any) *haxe__ds__StringMap {
		if hx_value_265 == nil {
			var hx_zero_266 *haxe__ds__StringMap
			return hx_zero_266
		}
		return hx_value_265.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_267 any) map[string]any {
		if hx_value_267 == nil {
			var hx_zero_268 map[string]any
			return hx_zero_268
		}
		return hx_value_267.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_269 map[string]any) func() bool {
		hx_field_270 := hx_obj_269["hasNext"]
		if hx_field_270 == nil {
			var hx_zero_271 func() bool
			return hx_zero_271
		}
		return hx_field_270.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_272 map[string]any) func() *string {
			hx_field_273 := hx_obj_272["next"]
			if hx_field_273 == nil {
				var hx_zero_274 func() *string
				return hx_zero_274
			}
			return hx_field_273.(func() *string)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_275 any) map[string]any {
		if hx_value_275 == nil {
			var hx_zero_276 map[string]any
			return hx_zero_276
		}
		return hx_value_275.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_277 map[string]any) func() bool {
		hx_field_278 := hx_obj_277["hasNext"]
		if hx_field_278 == nil {
			var hx_zero_279 func() bool
			return hx_zero_279
		}
		return hx_field_278.(func() bool)
	}(iterator)() {
		key := func(hx_obj_280 map[string]any) func() *string {
			hx_field_281 := hx_obj_280["next"]
			if hx_field_281 == nil {
				var hx_zero_282 func() *string
				return hx_zero_282
			}
			return hx_field_281.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_283 map[string]any) func() bool {
			hx_field_284 := hx_obj_283["hasNext"]
			if hx_field_284 == nil {
				var hx_zero_285 func() bool
				return hx_zero_285
			}
			return hx_field_284.(func() bool)
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
