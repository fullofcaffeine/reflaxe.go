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
	hx_obj_389 := map[string]any{}
	hx_obj_389["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_389["next"] = func() int {
		hx_post_390 := index
		index = int(int32((index + 1)))
		return keys[hx_post_390]
	}
	return hx_obj_389
}

func (self *haxe__ds__IntMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.IntMapKeys(self.h)
	index := 0
	hx_obj_391 := map[string]any{}
	hx_obj_391["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_391["next"] = func() any {
		return hxrt.IntMapGet(_gthis.h, keys[func() int {
			hx_post_392 := index
			index = int(int32((index + 1)))
			return hx_post_392
		}()])
	}
	return hx_obj_391
}

func (self *haxe__ds__IntMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_393 any) map[string]any {
		if hx_value_393 == nil {
			var hx_zero_394 map[string]any
			return hx_zero_394
		}
		return hx_value_393.(map[string]any)
	}(self.keys())
	hx_obj_395 := map[string]any{}
	hx_obj_395["hasNext"] = func() bool {
		return func(hx_obj_396 map[string]any) func() bool {
			hx_field_397 := hx_obj_396["hasNext"]
			if hx_field_397 == nil {
				var hx_zero_398 func() bool
				return hx_zero_398
			}
			return hx_field_397.(func() bool)
		}(keys)()
	}
	hx_obj_395["next"] = func() map[string]any {
		key := func(hx_obj_399 map[string]any) func() int {
			hx_field_400 := hx_obj_399["next"]
			if hx_field_400 == nil {
				var hx_zero_401 func() int
				return hx_zero_401
			}
			return hx_field_400.(func() int)
		}(keys)()
		hx_obj_402 := map[string]any{}
		hx_obj_402["key"] = key
		hx_obj_402["value"] = _gthis.get(key)
		return hx_obj_402
	}
	return hx_obj_395
}

func (self *haxe__ds__IntMap) getIMap(key any) any {
	return self.get(hxrt.IntFromNullableAny(func(hx_value_403 any) int {
		if hx_value_403 == nil {
			var hx_zero_404 int
			return hx_zero_404
		}
		return hx_value_403.(int)
	}(key)))
}

func (self *haxe__ds__IntMap) setIMap(key any, value any) {
	self.set(hxrt.IntFromNullableAny(func(hx_value_405 any) int {
		if hx_value_405 == nil {
			var hx_zero_406 int
			return hx_zero_406
		}
		return hx_value_405.(int)
	}(key)), value)
}

func (self *haxe__ds__IntMap) existsIMap(key any) bool {
	return func(hx_value_409 any) bool {
		if hx_value_409 == nil {
			var hx_zero_410 bool
			return hx_zero_410
		}
		return hx_value_409.(bool)
	}(self.exists(hxrt.IntFromNullableAny(func(hx_value_407 any) int {
		if hx_value_407 == nil {
			var hx_zero_408 int
			return hx_zero_408
		}
		return hx_value_407.(int)
	}(key))))
}

func (self *haxe__ds__IntMap) removeIMap(key any) bool {
	return func(hx_value_413 any) bool {
		if hx_value_413 == nil {
			var hx_zero_414 bool
			return hx_zero_414
		}
		return hx_value_413.(bool)
	}(self.remove(hxrt.IntFromNullableAny(func(hx_value_411 any) int {
		if hx_value_411 == nil {
			var hx_zero_412 int
			return hx_zero_412
		}
		return hx_value_411.(int)
	}(key))))
}

func (self *haxe__ds__IntMap) copyIMap() haxe__IMap {
	return func(hx_value_415 any) *haxe__ds__IntMap {
		if hx_value_415 == nil {
			var hx_zero_416 *haxe__ds__IntMap
			return hx_zero_416
		}
		return hx_value_415.(*haxe__ds__IntMap)
	}(self.copy())
}

func (self *haxe__ds__IntMap) copy() *haxe__ds__IntMap {
	copied := New_haxe__ds__IntMap()
	key := func(hx_value_417 any) map[string]any {
		if hx_value_417 == nil {
			var hx_zero_418 map[string]any
			return hx_zero_418
		}
		return hx_value_417.(map[string]any)
	}(self.keys())
	for func(hx_obj_419 map[string]any) func() bool {
		hx_field_420 := hx_obj_419["hasNext"]
		if hx_field_420 == nil {
			var hx_zero_421 func() bool
			return hx_zero_421
		}
		return hx_field_420.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_422 map[string]any) func() int {
			hx_field_423 := hx_obj_422["next"]
			if hx_field_423 == nil {
				var hx_zero_424 func() int
				return hx_zero_424
			}
			return hx_field_423.(func() int)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__IntMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_425 any) map[string]any {
		if hx_value_425 == nil {
			var hx_zero_426 map[string]any
			return hx_zero_426
		}
		return hx_value_425.(map[string]any)
	}(self.keys())
	for func(hx_obj_427 map[string]any) func() bool {
		hx_field_428 := hx_obj_427["hasNext"]
		if hx_field_428 == nil {
			var hx_zero_429 func() bool
			return hx_zero_429
		}
		return hx_field_428.(func() bool)
	}(iterator)() {
		key := func(hx_obj_430 map[string]any) func() int {
			hx_field_431 := hx_obj_430["next"]
			if hx_field_431 == nil {
				var hx_zero_432 func() int
				return hx_zero_432
			}
			return hx_field_431.(func() int)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_433 map[string]any) func() bool {
			hx_field_434 := hx_obj_433["hasNext"]
			if hx_field_434 == nil {
				var hx_zero_435 func() bool
				return hx_zero_435
			}
			return hx_field_434.(func() bool)
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
