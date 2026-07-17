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
	hx_obj_282 := map[string]any{}
	hx_obj_282["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_282["next"] = func() *string {
		hx_post_283 := index
		index = int(int32((index + 1)))
		return keys[hx_post_283]
	}
	return hx_obj_282
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_284 := map[string]any{}
	hx_obj_284["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_284["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_285 := index
			index = int(int32((index + 1)))
			return hx_post_285
		}()])
	}
	return hx_obj_284
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_286 any) map[string]any {
		if hx_value_286 == nil {
			var hx_zero_287 map[string]any
			return hx_zero_287
		}
		return hx_value_286.(map[string]any)
	}(self.keys())
	hx_obj_288 := map[string]any{}
	hx_obj_288["hasNext"] = func() bool {
		return func(hx_obj_289 map[string]any) func() bool {
			hx_field_290 := hx_obj_289["hasNext"]
			if hx_field_290 == nil {
				var hx_zero_291 func() bool
				return hx_zero_291
			}
			return hx_field_290.(func() bool)
		}(keys)()
	}
	hx_obj_288["next"] = func() map[string]any {
		key := func(hx_obj_292 map[string]any) func() *string {
			hx_field_293 := hx_obj_292["next"]
			if hx_field_293 == nil {
				var hx_zero_294 func() *string
				return hx_zero_294
			}
			return hx_field_293.(func() *string)
		}(keys)()
		hx_obj_295 := map[string]any{}
		hx_obj_295["key"] = key
		hx_obj_295["value"] = _gthis.get(key)
		return hx_obj_295
	}
	return hx_obj_288
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_296 any) *string {
		if hx_value_296 == nil {
			var hx_zero_297 *string
			return hx_zero_297
		}
		return hx_value_296.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_298 any) *string {
		if hx_value_298 == nil {
			var hx_zero_299 *string
			return hx_zero_299
		}
		return hx_value_298.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_302 any) bool {
		if hx_value_302 == nil {
			var hx_zero_303 bool
			return hx_zero_303
		}
		return hx_value_302.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_300 any) *string {
		if hx_value_300 == nil {
			var hx_zero_301 *string
			return hx_zero_301
		}
		return hx_value_300.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_306 any) bool {
		if hx_value_306 == nil {
			var hx_zero_307 bool
			return hx_zero_307
		}
		return hx_value_306.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_304 any) *string {
		if hx_value_304 == nil {
			var hx_zero_305 *string
			return hx_zero_305
		}
		return hx_value_304.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_308 any) *haxe__ds__StringMap {
		if hx_value_308 == nil {
			var hx_zero_309 *haxe__ds__StringMap
			return hx_zero_309
		}
		return hx_value_308.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_310 any) map[string]any {
		if hx_value_310 == nil {
			var hx_zero_311 map[string]any
			return hx_zero_311
		}
		return hx_value_310.(map[string]any)
	}(self.keys())
	for func(hx_obj_312 map[string]any) func() bool {
		hx_field_313 := hx_obj_312["hasNext"]
		if hx_field_313 == nil {
			var hx_zero_314 func() bool
			return hx_zero_314
		}
		return hx_field_313.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_315 map[string]any) func() *string {
			hx_field_316 := hx_obj_315["next"]
			if hx_field_316 == nil {
				var hx_zero_317 func() *string
				return hx_zero_317
			}
			return hx_field_316.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_318 any) map[string]any {
		if hx_value_318 == nil {
			var hx_zero_319 map[string]any
			return hx_zero_319
		}
		return hx_value_318.(map[string]any)
	}(self.keys())
	for func(hx_obj_320 map[string]any) func() bool {
		hx_field_321 := hx_obj_320["hasNext"]
		if hx_field_321 == nil {
			var hx_zero_322 func() bool
			return hx_zero_322
		}
		return hx_field_321.(func() bool)
	}(iterator)() {
		key := func(hx_obj_323 map[string]any) func() *string {
			hx_field_324 := hx_obj_323["next"]
			if hx_field_324 == nil {
				var hx_zero_325 func() *string
				return hx_zero_325
			}
			return hx_field_324.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_326 map[string]any) func() bool {
			hx_field_327 := hx_obj_326["hasNext"]
			if hx_field_327 == nil {
				var hx_zero_328 func() bool
				return hx_zero_328
			}
			return hx_field_327.(func() bool)
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
