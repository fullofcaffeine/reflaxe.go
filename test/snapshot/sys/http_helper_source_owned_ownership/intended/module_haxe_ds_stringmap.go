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
	hx_obj_258 := map[string]any{}
	hx_obj_258["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_258["next"] = func() *string {
		hx_post_259 := index
		index = int(int32((index + 1)))
		return keys[hx_post_259]
	}
	return hx_obj_258
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_260 := map[string]any{}
	hx_obj_260["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_260["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_261 := index
			index = int(int32((index + 1)))
			return hx_post_261
		}()])
	}
	return hx_obj_260
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_262 any) map[string]any {
		if hx_value_262 == nil {
			var hx_zero_263 map[string]any
			return hx_zero_263
		}
		return hx_value_262.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_264 := map[string]any{}
	hx_obj_264["hasNext"] = func() bool {
		return func(hx_obj_265 map[string]any) func() bool {
			hx_field_266 := hx_obj_265["hasNext"]
			if hx_field_266 == nil {
				var hx_zero_267 func() bool
				return hx_zero_267
			}
			return hx_field_266.(func() bool)
		}(keys)()
	}
	hx_obj_264["next"] = func() map[string]any {
		key := func(hx_obj_268 map[string]any) func() *string {
			hx_field_269 := hx_obj_268["next"]
			if hx_field_269 == nil {
				var hx_zero_270 func() *string
				return hx_zero_270
			}
			return hx_field_269.(func() *string)
		}(keys)()
		hx_obj_271 := map[string]any{}
		hx_obj_271["key"] = key
		hx_obj_271["value"] = _gthis.__hx_this.get(key)
		return hx_obj_271
	}
	return hx_obj_264
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_272 any) *string {
		if hx_value_272 == nil {
			var hx_zero_273 *string
			return hx_zero_273
		}
		return hx_value_272.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_274 any) *string {
		if hx_value_274 == nil {
			var hx_zero_275 *string
			return hx_zero_275
		}
		return hx_value_274.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_278 any) bool {
		if hx_value_278 == nil {
			var hx_zero_279 bool
			return hx_zero_279
		}
		return hx_value_278.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_276 any) *string {
		if hx_value_276 == nil {
			var hx_zero_277 *string
			return hx_zero_277
		}
		return hx_value_276.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_282 any) bool {
		if hx_value_282 == nil {
			var hx_zero_283 bool
			return hx_zero_283
		}
		return hx_value_282.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_280 any) *string {
		if hx_value_280 == nil {
			var hx_zero_281 *string
			return hx_zero_281
		}
		return hx_value_280.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_284 any) *haxe__ds__StringMap {
		if hx_value_284 == nil {
			var hx_zero_285 *haxe__ds__StringMap
			return hx_zero_285
		}
		return hx_value_284.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_286 any) map[string]any {
		if hx_value_286 == nil {
			var hx_zero_287 map[string]any
			return hx_zero_287
		}
		return hx_value_286.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_288 map[string]any) func() bool {
		hx_field_289 := hx_obj_288["hasNext"]
		if hx_field_289 == nil {
			var hx_zero_290 func() bool
			return hx_zero_290
		}
		return hx_field_289.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_291 map[string]any) func() *string {
			hx_field_292 := hx_obj_291["next"]
			if hx_field_292 == nil {
				var hx_zero_293 func() *string
				return hx_zero_293
			}
			return hx_field_292.(func() *string)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_294 any) map[string]any {
		if hx_value_294 == nil {
			var hx_zero_295 map[string]any
			return hx_zero_295
		}
		return hx_value_294.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_296 map[string]any) func() bool {
		hx_field_297 := hx_obj_296["hasNext"]
		if hx_field_297 == nil {
			var hx_zero_298 func() bool
			return hx_zero_298
		}
		return hx_field_297.(func() bool)
	}(iterator)() {
		key := func(hx_obj_299 map[string]any) func() *string {
			hx_field_300 := hx_obj_299["next"]
			if hx_field_300 == nil {
				var hx_zero_301 func() *string
				return hx_zero_301
			}
			return hx_field_300.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_302 map[string]any) func() bool {
			hx_field_303 := hx_obj_302["hasNext"]
			if hx_field_303 == nil {
				var hx_zero_304 func() bool
				return hx_zero_304
			}
			return hx_field_303.(func() bool)
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
