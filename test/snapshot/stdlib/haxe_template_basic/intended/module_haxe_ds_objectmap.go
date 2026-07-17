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
	hx_obj_287 := map[string]any{}
	hx_obj_287["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_287["next"] = func() any {
		return keys[func() int {
			hx_post_288 := index
			index = int(int32((index + 1)))
			return hx_post_288
		}()]
	}
	return hx_obj_287
}

func (self *haxe__ds__ObjectMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.ObjectMapKeys(self.h)
	index := 0
	hx_obj_289 := map[string]any{}
	hx_obj_289["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_289["next"] = func() any {
		return hxrt.ObjectMapGet(_gthis.h, keys[func() int {
			hx_post_290 := index
			index = int(int32((index + 1)))
			return hx_post_290
		}()])
	}
	return hx_obj_289
}

func (self *haxe__ds__ObjectMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_291 any) map[string]any {
		if hx_value_291 == nil {
			var hx_zero_292 map[string]any
			return hx_zero_292
		}
		return hx_value_291.(map[string]any)
	}(self.keys())
	hx_obj_293 := map[string]any{}
	hx_obj_293["hasNext"] = func() bool {
		return func(hx_obj_294 map[string]any) func() bool {
			hx_field_295 := hx_obj_294["hasNext"]
			if hx_field_295 == nil {
				var hx_zero_296 func() bool
				return hx_zero_296
			}
			return hx_field_295.(func() bool)
		}(keys)()
	}
	hx_obj_293["next"] = func() map[string]any {
		var key any = func(hx_obj_297 map[string]any) func() any {
			hx_field_298 := hx_obj_297["next"]
			if hx_field_298 == nil {
				var hx_zero_299 func() any
				return hx_zero_299
			}
			return hx_field_298.(func() any)
		}(keys)()
		hx_obj_300 := map[string]any{}
		hx_obj_300["key"] = key
		hx_obj_300["value"] = _gthis.get(key)
		return hx_obj_300
	}
	return hx_obj_293
}

func (self *haxe__ds__ObjectMap) getIMap(key any) any {
	return self.get(key)
}

func (self *haxe__ds__ObjectMap) setIMap(key any, value any) {
	self.set(key, value)
}

func (self *haxe__ds__ObjectMap) existsIMap(key any) bool {
	return func(hx_value_301 any) bool {
		if hx_value_301 == nil {
			var hx_zero_302 bool
			return hx_zero_302
		}
		return hx_value_301.(bool)
	}(self.exists(key))
}

func (self *haxe__ds__ObjectMap) removeIMap(key any) bool {
	return func(hx_value_303 any) bool {
		if hx_value_303 == nil {
			var hx_zero_304 bool
			return hx_zero_304
		}
		return hx_value_303.(bool)
	}(self.remove(key))
}

func (self *haxe__ds__ObjectMap) copyIMap() haxe__IMap {
	return func(hx_value_305 any) *haxe__ds__ObjectMap {
		if hx_value_305 == nil {
			var hx_zero_306 *haxe__ds__ObjectMap
			return hx_zero_306
		}
		return hx_value_305.(*haxe__ds__ObjectMap)
	}(self.copy())
}

func (self *haxe__ds__ObjectMap) copy() *haxe__ds__ObjectMap {
	copied := New_haxe__ds__ObjectMap()
	key := func(hx_value_307 any) map[string]any {
		if hx_value_307 == nil {
			var hx_zero_308 map[string]any
			return hx_zero_308
		}
		return hx_value_307.(map[string]any)
	}(self.keys())
	for func(hx_obj_309 map[string]any) func() bool {
		hx_field_310 := hx_obj_309["hasNext"]
		if hx_field_310 == nil {
			var hx_zero_311 func() bool
			return hx_zero_311
		}
		return hx_field_310.(func() bool)
	}(key)() {
		var key_1 any = func(hx_obj_312 map[string]any) func() any {
			hx_field_313 := hx_obj_312["next"]
			if hx_field_313 == nil {
				var hx_zero_314 func() any
				return hx_zero_314
			}
			return hx_field_313.(func() any)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__ObjectMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_315 any) map[string]any {
		if hx_value_315 == nil {
			var hx_zero_316 map[string]any
			return hx_zero_316
		}
		return hx_value_315.(map[string]any)
	}(self.keys())
	for func(hx_obj_317 map[string]any) func() bool {
		hx_field_318 := hx_obj_317["hasNext"]
		if hx_field_318 == nil {
			var hx_zero_319 func() bool
			return hx_zero_319
		}
		return hx_field_318.(func() bool)
	}(iterator)() {
		var key any = func(hx_obj_320 map[string]any) func() any {
			hx_field_321 := hx_obj_320["next"]
			if hx_field_321 == nil {
				var hx_zero_322 func() any
				return hx_zero_322
			}
			return hx_field_321.(func() any)
		}(iterator)()
		x := hxrt.StdString(key)
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x_1 := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x_1))
		if func(hx_obj_323 map[string]any) func() bool {
			hx_field_324 := hx_obj_323["hasNext"]
			if hx_field_324 == nil {
				var hx_zero_325 func() bool
				return hx_zero_325
			}
			return hx_field_324.(func() bool)
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
