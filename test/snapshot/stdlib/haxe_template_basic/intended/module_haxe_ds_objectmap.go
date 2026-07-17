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
	hx_obj_329 := map[string]any{}
	hx_obj_329["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_329["next"] = func() any {
		return keys[func() int {
			hx_post_330 := index
			index = int(int32((index + 1)))
			return hx_post_330
		}()]
	}
	return hx_obj_329
}

func (self *haxe__ds__ObjectMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.ObjectMapKeys(self.h)
	index := 0
	hx_obj_331 := map[string]any{}
	hx_obj_331["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_331["next"] = func() any {
		return hxrt.ObjectMapGet(_gthis.h, keys[func() int {
			hx_post_332 := index
			index = int(int32((index + 1)))
			return hx_post_332
		}()])
	}
	return hx_obj_331
}

func (self *haxe__ds__ObjectMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_333 any) map[string]any {
		if hx_value_333 == nil {
			var hx_zero_334 map[string]any
			return hx_zero_334
		}
		return hx_value_333.(map[string]any)
	}(self.keys())
	hx_obj_335 := map[string]any{}
	hx_obj_335["hasNext"] = func() bool {
		return func(hx_obj_336 map[string]any) func() bool {
			hx_field_337 := hx_obj_336["hasNext"]
			if hx_field_337 == nil {
				var hx_zero_338 func() bool
				return hx_zero_338
			}
			return hx_field_337.(func() bool)
		}(keys)()
	}
	hx_obj_335["next"] = func() map[string]any {
		var key any = func(hx_obj_339 map[string]any) func() any {
			hx_field_340 := hx_obj_339["next"]
			if hx_field_340 == nil {
				var hx_zero_341 func() any
				return hx_zero_341
			}
			return hx_field_340.(func() any)
		}(keys)()
		hx_obj_342 := map[string]any{}
		hx_obj_342["key"] = key
		hx_obj_342["value"] = _gthis.get(key)
		return hx_obj_342
	}
	return hx_obj_335
}

func (self *haxe__ds__ObjectMap) getIMap(key any) any {
	return self.get(key)
}

func (self *haxe__ds__ObjectMap) setIMap(key any, value any) {
	self.set(key, value)
}

func (self *haxe__ds__ObjectMap) existsIMap(key any) bool {
	return func(hx_value_343 any) bool {
		if hx_value_343 == nil {
			var hx_zero_344 bool
			return hx_zero_344
		}
		return hx_value_343.(bool)
	}(self.exists(key))
}

func (self *haxe__ds__ObjectMap) removeIMap(key any) bool {
	return func(hx_value_345 any) bool {
		if hx_value_345 == nil {
			var hx_zero_346 bool
			return hx_zero_346
		}
		return hx_value_345.(bool)
	}(self.remove(key))
}

func (self *haxe__ds__ObjectMap) copyIMap() haxe__IMap {
	return func(hx_value_347 any) *haxe__ds__ObjectMap {
		if hx_value_347 == nil {
			var hx_zero_348 *haxe__ds__ObjectMap
			return hx_zero_348
		}
		return hx_value_347.(*haxe__ds__ObjectMap)
	}(self.copy())
}

func (self *haxe__ds__ObjectMap) copy() *haxe__ds__ObjectMap {
	copied := New_haxe__ds__ObjectMap()
	key := func(hx_value_349 any) map[string]any {
		if hx_value_349 == nil {
			var hx_zero_350 map[string]any
			return hx_zero_350
		}
		return hx_value_349.(map[string]any)
	}(self.keys())
	for func(hx_obj_351 map[string]any) func() bool {
		hx_field_352 := hx_obj_351["hasNext"]
		if hx_field_352 == nil {
			var hx_zero_353 func() bool
			return hx_zero_353
		}
		return hx_field_352.(func() bool)
	}(key)() {
		var key_1 any = func(hx_obj_354 map[string]any) func() any {
			hx_field_355 := hx_obj_354["next"]
			if hx_field_355 == nil {
				var hx_zero_356 func() any
				return hx_zero_356
			}
			return hx_field_355.(func() any)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__ObjectMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_357 any) map[string]any {
		if hx_value_357 == nil {
			var hx_zero_358 map[string]any
			return hx_zero_358
		}
		return hx_value_357.(map[string]any)
	}(self.keys())
	for func(hx_obj_359 map[string]any) func() bool {
		hx_field_360 := hx_obj_359["hasNext"]
		if hx_field_360 == nil {
			var hx_zero_361 func() bool
			return hx_zero_361
		}
		return hx_field_360.(func() bool)
	}(iterator)() {
		var key any = func(hx_obj_362 map[string]any) func() any {
			hx_field_363 := hx_obj_362["next"]
			if hx_field_363 == nil {
				var hx_zero_364 func() any
				return hx_zero_364
			}
			return hx_field_363.(func() any)
		}(iterator)()
		x := hxrt.StdString(key)
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x_1 := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x_1))
		if func(hx_obj_365 map[string]any) func() bool {
			hx_field_366 := hx_obj_365["hasNext"]
			if hx_field_366 == nil {
				var hx_zero_367 func() bool
				return hx_zero_367
			}
			return hx_field_366.(func() bool)
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
