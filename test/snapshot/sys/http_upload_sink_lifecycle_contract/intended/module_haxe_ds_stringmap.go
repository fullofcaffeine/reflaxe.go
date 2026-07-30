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
	hx_obj_373 := map[string]any{}
	hx_obj_373["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_373["next"] = func() *string {
		hx_post_374 := index
		index = int(int32((index + 1)))
		return keys[hx_post_374]
	}
	return hx_obj_373
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_375 := map[string]any{}
	hx_obj_375["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_375["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_376 := index
			index = int(int32((index + 1)))
			return hx_post_376
		}()])
	}
	return hx_obj_375
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_377 any) map[string]any {
		if hx_value_377 == nil {
			var hx_zero_378 map[string]any
			return hx_zero_378
		}
		return hx_value_377.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_379 := map[string]any{}
	hx_obj_379["hasNext"] = func() bool {
		return func(hx_obj_380 map[string]any) func() bool {
			hx_field_381 := hx_obj_380["hasNext"]
			if hx_field_381 == nil {
				var hx_zero_382 func() bool
				return hx_zero_382
			}
			return hx_field_381.(func() bool)
		}(keys)()
	}
	hx_obj_379["next"] = func() map[string]any {
		key := func(hx_obj_383 map[string]any) func() *string {
			hx_field_384 := hx_obj_383["next"]
			if hx_field_384 == nil {
				var hx_zero_385 func() *string
				return hx_zero_385
			}
			return hx_field_384.(func() *string)
		}(keys)()
		hx_obj_386 := map[string]any{}
		hx_obj_386["key"] = key
		hx_obj_386["value"] = _gthis.__hx_this.get(key)
		return hx_obj_386
	}
	return hx_obj_379
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_387 any) *string {
		if hx_value_387 == nil {
			var hx_zero_388 *string
			return hx_zero_388
		}
		return hx_value_387.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_389 any) *string {
		if hx_value_389 == nil {
			var hx_zero_390 *string
			return hx_zero_390
		}
		return hx_value_389.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_393 any) bool {
		if hx_value_393 == nil {
			var hx_zero_394 bool
			return hx_zero_394
		}
		return hx_value_393.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_391 any) *string {
		if hx_value_391 == nil {
			var hx_zero_392 *string
			return hx_zero_392
		}
		return hx_value_391.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_397 any) bool {
		if hx_value_397 == nil {
			var hx_zero_398 bool
			return hx_zero_398
		}
		return hx_value_397.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_395 any) *string {
		if hx_value_395 == nil {
			var hx_zero_396 *string
			return hx_zero_396
		}
		return hx_value_395.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_399 any) *haxe__ds__StringMap {
		if hx_value_399 == nil {
			var hx_zero_400 *haxe__ds__StringMap
			return hx_zero_400
		}
		return hx_value_399.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_401 any) map[string]any {
		if hx_value_401 == nil {
			var hx_zero_402 map[string]any
			return hx_zero_402
		}
		return hx_value_401.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_403 map[string]any) func() bool {
		hx_field_404 := hx_obj_403["hasNext"]
		if hx_field_404 == nil {
			var hx_zero_405 func() bool
			return hx_zero_405
		}
		return hx_field_404.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_406 map[string]any) func() *string {
			hx_field_407 := hx_obj_406["next"]
			if hx_field_407 == nil {
				var hx_zero_408 func() *string
				return hx_zero_408
			}
			return hx_field_407.(func() *string)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_409 any) map[string]any {
		if hx_value_409 == nil {
			var hx_zero_410 map[string]any
			return hx_zero_410
		}
		return hx_value_409.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_411 map[string]any) func() bool {
		hx_field_412 := hx_obj_411["hasNext"]
		if hx_field_412 == nil {
			var hx_zero_413 func() bool
			return hx_zero_413
		}
		return hx_field_412.(func() bool)
	}(iterator)() {
		key := func(hx_obj_414 map[string]any) func() *string {
			hx_field_415 := hx_obj_414["next"]
			if hx_field_415 == nil {
				var hx_zero_416 func() *string
				return hx_zero_416
			}
			return hx_field_415.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_417 map[string]any) func() bool {
			hx_field_418 := hx_obj_417["hasNext"]
			if hx_field_418 == nil {
				var hx_zero_419 func() bool
				return hx_zero_419
			}
			return hx_field_418.(func() bool)
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
