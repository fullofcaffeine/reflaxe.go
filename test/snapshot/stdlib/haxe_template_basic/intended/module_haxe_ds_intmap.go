package main

import "snapshot/hxrt"

type I_haxe__ds__IntMap interface {
	set(key int, value any)
	get(key int) any
	exists(key int) bool
	remove(key int) bool
	keys() map[string]any
	iterator() map[string]any
	keyValueIterator() map[string]any
	getIMap(key any) any
	setIMap(key any, value any)
	existsIMap(key any) bool
	removeIMap(key any) bool
	copyIMap() haxe__IMap
	copy() *haxe__ds__IntMap
	toString() *string
	clear()
}

type haxe__ds__IntMap struct {
	__hx_this I_haxe__ds__IntMap
	h         *hxrt.IntMapCell
}

func New_haxe__ds__IntMap() *haxe__ds__IntMap {
	self := &haxe__ds__IntMap{}
	self.__hx_this = self
	self.h = hxrt.IntMapNew()
	return self
}

func (self *haxe__ds__IntMap) set(key int, value any) {
	hxrt.IntMapSet(self.h, key, value)
}

func (self *haxe__ds__IntMap) get(key int) any {
	return hxrt.IntMapGet(self.h, key)
}

func (self *haxe__ds__IntMap) exists(key int) bool {
	return hxrt.IntMapExists(self.h, key)
}

func (self *haxe__ds__IntMap) remove(key int) bool {
	return hxrt.IntMapRemove(self.h, key)
}

func (self *haxe__ds__IntMap) keys() map[string]any {
	keys := hxrt.IntMapKeys(self.h)
	index := 0
	hx_obj_353 := map[string]any{}
	hx_obj_353["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_353["next"] = func() int {
		hx_post_354 := index
		index = int(int32((index + 1)))
		return keys[hx_post_354]
	}
	return hx_obj_353
}

func (self *haxe__ds__IntMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.IntMapKeys(self.h)
	index := 0
	hx_obj_355 := map[string]any{}
	hx_obj_355["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_355["next"] = func() any {
		return hxrt.IntMapGet(_gthis.h, keys[func() int {
			hx_post_356 := index
			index = int(int32((index + 1)))
			return hx_post_356
		}()])
	}
	return hx_obj_355
}

func (self *haxe__ds__IntMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_357 any) map[string]any {
		if hx_value_357 == nil {
			var hx_zero_358 map[string]any
			return hx_zero_358
		}
		return hx_value_357.(map[string]any)
	}(self.keys())
	hx_obj_359 := map[string]any{}
	hx_obj_359["hasNext"] = func() bool {
		return func(hx_obj_360 map[string]any) func() bool {
			hx_field_361 := hx_obj_360["hasNext"]
			if hx_field_361 == nil {
				var hx_zero_362 func() bool
				return hx_zero_362
			}
			return hx_field_361.(func() bool)
		}(keys)()
	}
	hx_obj_359["next"] = func() map[string]any {
		key := func(hx_obj_363 map[string]any) func() int {
			hx_field_364 := hx_obj_363["next"]
			if hx_field_364 == nil {
				var hx_zero_365 func() int
				return hx_zero_365
			}
			return hx_field_364.(func() int)
		}(keys)()
		hx_obj_366 := map[string]any{}
		hx_obj_366["key"] = key
		hx_obj_366["value"] = _gthis.get(key)
		return hx_obj_366
	}
	return hx_obj_359
}

func (self *haxe__ds__IntMap) getIMap(key any) any {
	return self.get(hxrt.IntFromNullableAny(func(hx_value_367 any) int {
		if hx_value_367 == nil {
			var hx_zero_368 int
			return hx_zero_368
		}
		return hx_value_367.(int)
	}(key)))
}

func (self *haxe__ds__IntMap) setIMap(key any, value any) {
	self.set(hxrt.IntFromNullableAny(func(hx_value_369 any) int {
		if hx_value_369 == nil {
			var hx_zero_370 int
			return hx_zero_370
		}
		return hx_value_369.(int)
	}(key)), value)
}

func (self *haxe__ds__IntMap) existsIMap(key any) bool {
	return func(hx_value_373 any) bool {
		if hx_value_373 == nil {
			var hx_zero_374 bool
			return hx_zero_374
		}
		return hx_value_373.(bool)
	}(self.exists(hxrt.IntFromNullableAny(func(hx_value_371 any) int {
		if hx_value_371 == nil {
			var hx_zero_372 int
			return hx_zero_372
		}
		return hx_value_371.(int)
	}(key))))
}

func (self *haxe__ds__IntMap) removeIMap(key any) bool {
	return func(hx_value_377 any) bool {
		if hx_value_377 == nil {
			var hx_zero_378 bool
			return hx_zero_378
		}
		return hx_value_377.(bool)
	}(self.remove(hxrt.IntFromNullableAny(func(hx_value_375 any) int {
		if hx_value_375 == nil {
			var hx_zero_376 int
			return hx_zero_376
		}
		return hx_value_375.(int)
	}(key))))
}

func (self *haxe__ds__IntMap) copyIMap() haxe__IMap {
	return func(hx_value_379 any) *haxe__ds__IntMap {
		if hx_value_379 == nil {
			var hx_zero_380 *haxe__ds__IntMap
			return hx_zero_380
		}
		return hx_value_379.(*haxe__ds__IntMap)
	}(self.copy())
}

func (self *haxe__ds__IntMap) copy() *haxe__ds__IntMap {
	copied := New_haxe__ds__IntMap()
	key := func(hx_value_381 any) map[string]any {
		if hx_value_381 == nil {
			var hx_zero_382 map[string]any
			return hx_zero_382
		}
		return hx_value_381.(map[string]any)
	}(self.keys())
	for func(hx_obj_383 map[string]any) func() bool {
		hx_field_384 := hx_obj_383["hasNext"]
		if hx_field_384 == nil {
			var hx_zero_385 func() bool
			return hx_zero_385
		}
		return hx_field_384.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_386 map[string]any) func() int {
			hx_field_387 := hx_obj_386["next"]
			if hx_field_387 == nil {
				var hx_zero_388 func() int
				return hx_zero_388
			}
			return hx_field_387.(func() int)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__IntMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_389 any) map[string]any {
		if hx_value_389 == nil {
			var hx_zero_390 map[string]any
			return hx_zero_390
		}
		return hx_value_389.(map[string]any)
	}(self.keys())
	for func(hx_obj_391 map[string]any) func() bool {
		hx_field_392 := hx_obj_391["hasNext"]
		if hx_field_392 == nil {
			var hx_zero_393 func() bool
			return hx_zero_393
		}
		return hx_field_392.(func() bool)
	}(iterator)() {
		key := func(hx_obj_394 map[string]any) func() int {
			hx_field_395 := hx_obj_394["next"]
			if hx_field_395 == nil {
				var hx_zero_396 func() int
				return hx_zero_396
			}
			return hx_field_395.(func() int)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_397 map[string]any) func() bool {
			hx_field_398 := hx_obj_397["hasNext"]
			if hx_field_398 == nil {
				var hx_zero_399 func() bool
				return hx_zero_399
			}
			return hx_field_398.(func() bool)
		}(iterator)() {
			out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(", "))
		}
	}
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("]"))
	return out_b
}

func (self *haxe__ds__IntMap) clear() {
	hxrt.IntMapClear(self.h)
}
