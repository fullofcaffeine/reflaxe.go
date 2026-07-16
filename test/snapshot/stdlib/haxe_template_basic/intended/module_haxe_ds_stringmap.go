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
	hx_obj_249 := map[string]any{}
	hx_obj_249["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_249["next"] = func() *string {
		hx_post_250 := index
		index = int(int32((index + 1)))
		return keys[hx_post_250]
	}
	return hx_obj_249
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_251 := map[string]any{}
	hx_obj_251["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_251["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_252 := index
			index = int(int32((index + 1)))
			return hx_post_252
		}()])
	}
	return hx_obj_251
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_253 any) map[string]any {
		if hx_value_253 == nil {
			var hx_zero_254 map[string]any
			return hx_zero_254
		}
		return hx_value_253.(map[string]any)
	}(self.keys())
	hx_obj_255 := map[string]any{}
	hx_obj_255["hasNext"] = func() bool {
		return func(hx_obj_256 map[string]any) func() bool {
			hx_field_257 := hx_obj_256["hasNext"]
			if hx_field_257 == nil {
				var hx_zero_258 func() bool
				return hx_zero_258
			}
			return hx_field_257.(func() bool)
		}(keys)()
	}
	hx_obj_255["next"] = func() map[string]any {
		key := func(hx_obj_259 map[string]any) func() *string {
			hx_field_260 := hx_obj_259["next"]
			if hx_field_260 == nil {
				var hx_zero_261 func() *string
				return hx_zero_261
			}
			return hx_field_260.(func() *string)
		}(keys)()
		hx_obj_262 := map[string]any{}
		hx_obj_262["key"] = key
		hx_obj_262["value"] = _gthis.get(key)
		return hx_obj_262
	}
	return hx_obj_255
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_263 any) *string {
		if hx_value_263 == nil {
			var hx_zero_264 *string
			return hx_zero_264
		}
		return hx_value_263.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_265 any) *string {
		if hx_value_265 == nil {
			var hx_zero_266 *string
			return hx_zero_266
		}
		return hx_value_265.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_269 any) bool {
		if hx_value_269 == nil {
			var hx_zero_270 bool
			return hx_zero_270
		}
		return hx_value_269.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_267 any) *string {
		if hx_value_267 == nil {
			var hx_zero_268 *string
			return hx_zero_268
		}
		return hx_value_267.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_273 any) bool {
		if hx_value_273 == nil {
			var hx_zero_274 bool
			return hx_zero_274
		}
		return hx_value_273.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_271 any) *string {
		if hx_value_271 == nil {
			var hx_zero_272 *string
			return hx_zero_272
		}
		return hx_value_271.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_275 any) *haxe__ds__StringMap {
		if hx_value_275 == nil {
			var hx_zero_276 *haxe__ds__StringMap
			return hx_zero_276
		}
		return hx_value_275.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_277 any) map[string]any {
		if hx_value_277 == nil {
			var hx_zero_278 map[string]any
			return hx_zero_278
		}
		return hx_value_277.(map[string]any)
	}(self.keys())
	for func(hx_obj_279 map[string]any) func() bool {
		hx_field_280 := hx_obj_279["hasNext"]
		if hx_field_280 == nil {
			var hx_zero_281 func() bool
			return hx_zero_281
		}
		return hx_field_280.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_282 map[string]any) func() *string {
			hx_field_283 := hx_obj_282["next"]
			if hx_field_283 == nil {
				var hx_zero_284 func() *string
				return hx_zero_284
			}
			return hx_field_283.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_285 any) map[string]any {
		if hx_value_285 == nil {
			var hx_zero_286 map[string]any
			return hx_zero_286
		}
		return hx_value_285.(map[string]any)
	}(self.keys())
	for func(hx_obj_287 map[string]any) func() bool {
		hx_field_288 := hx_obj_287["hasNext"]
		if hx_field_288 == nil {
			var hx_zero_289 func() bool
			return hx_zero_289
		}
		return hx_field_288.(func() bool)
	}(iterator)() {
		key := func(hx_obj_290 map[string]any) func() *string {
			hx_field_291 := hx_obj_290["next"]
			if hx_field_291 == nil {
				var hx_zero_292 func() *string
				return hx_zero_292
			}
			return hx_field_291.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_293 map[string]any) func() bool {
			hx_field_294 := hx_obj_293["hasNext"]
			if hx_field_294 == nil {
				var hx_zero_295 func() bool
				return hx_zero_295
			}
			return hx_field_294.(func() bool)
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
