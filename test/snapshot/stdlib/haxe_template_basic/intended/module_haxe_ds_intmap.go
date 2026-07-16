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
	hx_obj_362 := map[string]any{}
	hx_obj_362["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_362["next"] = func() int {
		hx_post_363 := index
		index = int(int32((index + 1)))
		return keys[hx_post_363]
	}
	return hx_obj_362
}

func (self *haxe__ds__IntMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.IntMapKeys(self.h)
	index := 0
	hx_obj_364 := map[string]any{}
	hx_obj_364["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_364["next"] = func() any {
		return hxrt.IntMapGet(_gthis.h, keys[func() int {
			hx_post_365 := index
			index = int(int32((index + 1)))
			return hx_post_365
		}()])
	}
	return hx_obj_364
}

func (self *haxe__ds__IntMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_366 any) map[string]any {
		if hx_value_366 == nil {
			var hx_zero_367 map[string]any
			return hx_zero_367
		}
		return hx_value_366.(map[string]any)
	}(self.keys())
	hx_obj_368 := map[string]any{}
	hx_obj_368["hasNext"] = func() bool {
		return func(hx_obj_369 map[string]any) func() bool {
			hx_field_370 := hx_obj_369["hasNext"]
			if hx_field_370 == nil {
				var hx_zero_371 func() bool
				return hx_zero_371
			}
			return hx_field_370.(func() bool)
		}(keys)()
	}
	hx_obj_368["next"] = func() map[string]any {
		key := func(hx_obj_372 map[string]any) func() int {
			hx_field_373 := hx_obj_372["next"]
			if hx_field_373 == nil {
				var hx_zero_374 func() int
				return hx_zero_374
			}
			return hx_field_373.(func() int)
		}(keys)()
		hx_obj_375 := map[string]any{}
		hx_obj_375["key"] = key
		hx_obj_375["value"] = _gthis.get(key)
		return hx_obj_375
	}
	return hx_obj_368
}

func (self *haxe__ds__IntMap) getIMap(key any) any {
	return self.get(hxrt.IntFromNullableAny(func(hx_value_376 any) int {
		if hx_value_376 == nil {
			var hx_zero_377 int
			return hx_zero_377
		}
		return hx_value_376.(int)
	}(key)))
}

func (self *haxe__ds__IntMap) setIMap(key any, value any) {
	self.set(hxrt.IntFromNullableAny(func(hx_value_378 any) int {
		if hx_value_378 == nil {
			var hx_zero_379 int
			return hx_zero_379
		}
		return hx_value_378.(int)
	}(key)), value)
}

func (self *haxe__ds__IntMap) existsIMap(key any) bool {
	return func(hx_value_382 any) bool {
		if hx_value_382 == nil {
			var hx_zero_383 bool
			return hx_zero_383
		}
		return hx_value_382.(bool)
	}(self.exists(hxrt.IntFromNullableAny(func(hx_value_380 any) int {
		if hx_value_380 == nil {
			var hx_zero_381 int
			return hx_zero_381
		}
		return hx_value_380.(int)
	}(key))))
}

func (self *haxe__ds__IntMap) removeIMap(key any) bool {
	return func(hx_value_386 any) bool {
		if hx_value_386 == nil {
			var hx_zero_387 bool
			return hx_zero_387
		}
		return hx_value_386.(bool)
	}(self.remove(hxrt.IntFromNullableAny(func(hx_value_384 any) int {
		if hx_value_384 == nil {
			var hx_zero_385 int
			return hx_zero_385
		}
		return hx_value_384.(int)
	}(key))))
}

func (self *haxe__ds__IntMap) copyIMap() haxe__IMap {
	return func(hx_value_388 any) *haxe__ds__IntMap {
		if hx_value_388 == nil {
			var hx_zero_389 *haxe__ds__IntMap
			return hx_zero_389
		}
		return hx_value_388.(*haxe__ds__IntMap)
	}(self.copy())
}

func (self *haxe__ds__IntMap) copy() *haxe__ds__IntMap {
	copied := New_haxe__ds__IntMap()
	key := func(hx_value_390 any) map[string]any {
		if hx_value_390 == nil {
			var hx_zero_391 map[string]any
			return hx_zero_391
		}
		return hx_value_390.(map[string]any)
	}(self.keys())
	for func(hx_obj_392 map[string]any) func() bool {
		hx_field_393 := hx_obj_392["hasNext"]
		if hx_field_393 == nil {
			var hx_zero_394 func() bool
			return hx_zero_394
		}
		return hx_field_393.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_395 map[string]any) func() int {
			hx_field_396 := hx_obj_395["next"]
			if hx_field_396 == nil {
				var hx_zero_397 func() int
				return hx_zero_397
			}
			return hx_field_396.(func() int)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__IntMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_398 any) map[string]any {
		if hx_value_398 == nil {
			var hx_zero_399 map[string]any
			return hx_zero_399
		}
		return hx_value_398.(map[string]any)
	}(self.keys())
	for func(hx_obj_400 map[string]any) func() bool {
		hx_field_401 := hx_obj_400["hasNext"]
		if hx_field_401 == nil {
			var hx_zero_402 func() bool
			return hx_zero_402
		}
		return hx_field_401.(func() bool)
	}(iterator)() {
		key := func(hx_obj_403 map[string]any) func() int {
			hx_field_404 := hx_obj_403["next"]
			if hx_field_404 == nil {
				var hx_zero_405 func() int
				return hx_zero_405
			}
			return hx_field_404.(func() int)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_406 map[string]any) func() bool {
			hx_field_407 := hx_obj_406["hasNext"]
			if hx_field_407 == nil {
				var hx_zero_408 func() bool
				return hx_zero_408
			}
			return hx_field_407.(func() bool)
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
