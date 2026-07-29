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
	hx_obj_259 := map[string]any{}
	hx_obj_259["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_259["next"] = func() *string {
		hx_post_260 := index
		index = int(int32((index + 1)))
		return keys[hx_post_260]
	}
	return hx_obj_259
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_261 := map[string]any{}
	hx_obj_261["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_261["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_262 := index
			index = int(int32((index + 1)))
			return hx_post_262
		}()])
	}
	return hx_obj_261
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_263 any) map[string]any {
		if hx_value_263 == nil {
			var hx_zero_264 map[string]any
			return hx_zero_264
		}
		return hx_value_263.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_265 := map[string]any{}
	hx_obj_265["hasNext"] = func() bool {
		return func(hx_obj_266 map[string]any) func() bool {
			hx_field_267 := hx_obj_266["hasNext"]
			if hx_field_267 == nil {
				var hx_zero_268 func() bool
				return hx_zero_268
			}
			return hx_field_267.(func() bool)
		}(keys)()
	}
	hx_obj_265["next"] = func() map[string]any {
		key := func(hx_obj_269 map[string]any) func() *string {
			hx_field_270 := hx_obj_269["next"]
			if hx_field_270 == nil {
				var hx_zero_271 func() *string
				return hx_zero_271
			}
			return hx_field_270.(func() *string)
		}(keys)()
		hx_obj_272 := map[string]any{}
		hx_obj_272["key"] = key
		hx_obj_272["value"] = _gthis.__hx_this.get(key)
		return hx_obj_272
	}
	return hx_obj_265
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_273 any) *string {
		if hx_value_273 == nil {
			var hx_zero_274 *string
			return hx_zero_274
		}
		return hx_value_273.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_275 any) *string {
		if hx_value_275 == nil {
			var hx_zero_276 *string
			return hx_zero_276
		}
		return hx_value_275.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_279 any) bool {
		if hx_value_279 == nil {
			var hx_zero_280 bool
			return hx_zero_280
		}
		return hx_value_279.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_277 any) *string {
		if hx_value_277 == nil {
			var hx_zero_278 *string
			return hx_zero_278
		}
		return hx_value_277.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_283 any) bool {
		if hx_value_283 == nil {
			var hx_zero_284 bool
			return hx_zero_284
		}
		return hx_value_283.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_281 any) *string {
		if hx_value_281 == nil {
			var hx_zero_282 *string
			return hx_zero_282
		}
		return hx_value_281.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_285 any) *haxe__ds__StringMap {
		if hx_value_285 == nil {
			var hx_zero_286 *haxe__ds__StringMap
			return hx_zero_286
		}
		return hx_value_285.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_287 any) map[string]any {
		if hx_value_287 == nil {
			var hx_zero_288 map[string]any
			return hx_zero_288
		}
		return hx_value_287.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_289 map[string]any) func() bool {
		hx_field_290 := hx_obj_289["hasNext"]
		if hx_field_290 == nil {
			var hx_zero_291 func() bool
			return hx_zero_291
		}
		return hx_field_290.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_292 map[string]any) func() *string {
			hx_field_293 := hx_obj_292["next"]
			if hx_field_293 == nil {
				var hx_zero_294 func() *string
				return hx_zero_294
			}
			return hx_field_293.(func() *string)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_295 any) map[string]any {
		if hx_value_295 == nil {
			var hx_zero_296 map[string]any
			return hx_zero_296
		}
		return hx_value_295.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_297 map[string]any) func() bool {
		hx_field_298 := hx_obj_297["hasNext"]
		if hx_field_298 == nil {
			var hx_zero_299 func() bool
			return hx_zero_299
		}
		return hx_field_298.(func() bool)
	}(iterator)() {
		key := func(hx_obj_300 map[string]any) func() *string {
			hx_field_301 := hx_obj_300["next"]
			if hx_field_301 == nil {
				var hx_zero_302 func() *string
				return hx_zero_302
			}
			return hx_field_301.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_303 map[string]any) func() bool {
			hx_field_304 := hx_obj_303["hasNext"]
			if hx_field_304 == nil {
				var hx_zero_305 func() bool
				return hx_zero_305
			}
			return hx_field_304.(func() bool)
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
