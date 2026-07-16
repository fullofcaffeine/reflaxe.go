package main

import "snapshot/hxrt"

type I_haxe__ds__ObjectMap interface {
	set(key any, value any)
	get(key any) any
	exists(key any) bool
	remove(key any) bool
	keys() map[string]any
	iterator() map[string]any
	keyValueIterator() map[string]any
	getIMap(key any) any
	setIMap(key any, value any)
	existsIMap(key any) bool
	removeIMap(key any) bool
	copyIMap() haxe__IMap
	copy() *haxe__ds__ObjectMap
	toString() *string
	clear()
}

type haxe__ds__ObjectMap struct {
	__hx_this I_haxe__ds__ObjectMap
	h         *hxrt.ObjectMapCell
}

func New_haxe__ds__ObjectMap() *haxe__ds__ObjectMap {
	self := &haxe__ds__ObjectMap{}
	self.__hx_this = self
	self.h = hxrt.ObjectMapNew()
	return self
}

func (self *haxe__ds__ObjectMap) set(key any, value any) {
	hxrt.ObjectMapSet(self.h, key, value)
}

func (self *haxe__ds__ObjectMap) get(key any) any {
	return hxrt.ObjectMapGet(self.h, key)
}

func (self *haxe__ds__ObjectMap) exists(key any) bool {
	return hxrt.ObjectMapExists(self.h, key)
}

func (self *haxe__ds__ObjectMap) remove(key any) bool {
	return hxrt.ObjectMapRemove(self.h, key)
}

func (self *haxe__ds__ObjectMap) keys() map[string]any {
	keys := hxrt.ObjectMapKeys(self.h)
	index := 0
	hx_obj_296 := map[string]any{}
	hx_obj_296["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_296["next"] = func() any {
		return keys[func() int {
			hx_post_297 := index
			index = int(int32((index + 1)))
			return hx_post_297
		}()]
	}
	return hx_obj_296
}

func (self *haxe__ds__ObjectMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.ObjectMapKeys(self.h)
	index := 0
	hx_obj_298 := map[string]any{}
	hx_obj_298["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_298["next"] = func() any {
		return hxrt.ObjectMapGet(_gthis.h, keys[func() int {
			hx_post_299 := index
			index = int(int32((index + 1)))
			return hx_post_299
		}()])
	}
	return hx_obj_298
}

func (self *haxe__ds__ObjectMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_300 any) map[string]any {
		if hx_value_300 == nil {
			var hx_zero_301 map[string]any
			return hx_zero_301
		}
		return hx_value_300.(map[string]any)
	}(self.keys())
	hx_obj_302 := map[string]any{}
	hx_obj_302["hasNext"] = func() bool {
		return func(hx_obj_303 map[string]any) func() bool {
			hx_field_304 := hx_obj_303["hasNext"]
			if hx_field_304 == nil {
				var hx_zero_305 func() bool
				return hx_zero_305
			}
			return hx_field_304.(func() bool)
		}(keys)()
	}
	hx_obj_302["next"] = func() map[string]any {
		var key any = func(hx_obj_306 map[string]any) func() any {
			hx_field_307 := hx_obj_306["next"]
			if hx_field_307 == nil {
				var hx_zero_308 func() any
				return hx_zero_308
			}
			return hx_field_307.(func() any)
		}(keys)()
		hx_obj_309 := map[string]any{}
		hx_obj_309["key"] = key
		hx_obj_309["value"] = _gthis.get(key)
		return hx_obj_309
	}
	return hx_obj_302
}

func (self *haxe__ds__ObjectMap) getIMap(key any) any {
	return self.get(key)
}

func (self *haxe__ds__ObjectMap) setIMap(key any, value any) {
	self.set(key, value)
}

func (self *haxe__ds__ObjectMap) existsIMap(key any) bool {
	return func(hx_value_310 any) bool {
		if hx_value_310 == nil {
			var hx_zero_311 bool
			return hx_zero_311
		}
		return hx_value_310.(bool)
	}(self.exists(key))
}

func (self *haxe__ds__ObjectMap) removeIMap(key any) bool {
	return func(hx_value_312 any) bool {
		if hx_value_312 == nil {
			var hx_zero_313 bool
			return hx_zero_313
		}
		return hx_value_312.(bool)
	}(self.remove(key))
}

func (self *haxe__ds__ObjectMap) copyIMap() haxe__IMap {
	return func(hx_value_314 any) *haxe__ds__ObjectMap {
		if hx_value_314 == nil {
			var hx_zero_315 *haxe__ds__ObjectMap
			return hx_zero_315
		}
		return hx_value_314.(*haxe__ds__ObjectMap)
	}(self.copy())
}

func (self *haxe__ds__ObjectMap) copy() *haxe__ds__ObjectMap {
	copied := New_haxe__ds__ObjectMap()
	key := func(hx_value_316 any) map[string]any {
		if hx_value_316 == nil {
			var hx_zero_317 map[string]any
			return hx_zero_317
		}
		return hx_value_316.(map[string]any)
	}(self.keys())
	for func(hx_obj_318 map[string]any) func() bool {
		hx_field_319 := hx_obj_318["hasNext"]
		if hx_field_319 == nil {
			var hx_zero_320 func() bool
			return hx_zero_320
		}
		return hx_field_319.(func() bool)
	}(key)() {
		var key_1 any = func(hx_obj_321 map[string]any) func() any {
			hx_field_322 := hx_obj_321["next"]
			if hx_field_322 == nil {
				var hx_zero_323 func() any
				return hx_zero_323
			}
			return hx_field_322.(func() any)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__ObjectMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_324 any) map[string]any {
		if hx_value_324 == nil {
			var hx_zero_325 map[string]any
			return hx_zero_325
		}
		return hx_value_324.(map[string]any)
	}(self.keys())
	for func(hx_obj_326 map[string]any) func() bool {
		hx_field_327 := hx_obj_326["hasNext"]
		if hx_field_327 == nil {
			var hx_zero_328 func() bool
			return hx_zero_328
		}
		return hx_field_327.(func() bool)
	}(iterator)() {
		var key any = func(hx_obj_329 map[string]any) func() any {
			hx_field_330 := hx_obj_329["next"]
			if hx_field_330 == nil {
				var hx_zero_331 func() any
				return hx_zero_331
			}
			return hx_field_330.(func() any)
		}(iterator)()
		x := hxrt.StdString(key)
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x_1 := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x_1))
		if func(hx_obj_332 map[string]any) func() bool {
			hx_field_333 := hx_obj_332["hasNext"]
			if hx_field_333 == nil {
				var hx_zero_334 func() bool
				return hx_zero_334
			}
			return hx_field_333.(func() bool)
		}(iterator)() {
			out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(", "))
		}
	}
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("]"))
	return out_b
}

func (self *haxe__ds__ObjectMap) clear() {
	hxrt.ObjectMapClear(self.h)
}
