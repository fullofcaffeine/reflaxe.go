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
	hx_obj_262 := map[string]any{}
	hx_obj_262["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_262["next"] = func() *string {
		hx_post_263 := index
		index = int(int32((index + 1)))
		return keys[hx_post_263]
	}
	return hx_obj_262
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_264 := map[string]any{}
	hx_obj_264["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_264["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_265 := index
			index = int(int32((index + 1)))
			return hx_post_265
		}()])
	}
	return hx_obj_264
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_266 any) map[string]any {
		if hx_value_266 == nil {
			var hx_zero_267 map[string]any
			return hx_zero_267
		}
		return hx_value_266.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_268 := map[string]any{}
	hx_obj_268["hasNext"] = func() bool {
		return func(hx_obj_269 map[string]any) func() bool {
			hx_field_270 := hx_obj_269["hasNext"]
			if hx_field_270 == nil {
				var hx_zero_271 func() bool
				return hx_zero_271
			}
			return hx_field_270.(func() bool)
		}(keys)()
	}
	hx_obj_268["next"] = func() map[string]any {
		key := func(hx_obj_272 map[string]any) func() *string {
			hx_field_273 := hx_obj_272["next"]
			if hx_field_273 == nil {
				var hx_zero_274 func() *string
				return hx_zero_274
			}
			return hx_field_273.(func() *string)
		}(keys)()
		hx_obj_275 := map[string]any{}
		hx_obj_275["key"] = key
		hx_obj_275["value"] = _gthis.__hx_this.get(key)
		return hx_obj_275
	}
	return hx_obj_268
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_276 any) *string {
		if hx_value_276 == nil {
			var hx_zero_277 *string
			return hx_zero_277
		}
		return hx_value_276.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_278 any) *string {
		if hx_value_278 == nil {
			var hx_zero_279 *string
			return hx_zero_279
		}
		return hx_value_278.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_282 any) bool {
		if hx_value_282 == nil {
			var hx_zero_283 bool
			return hx_zero_283
		}
		return hx_value_282.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_280 any) *string {
		if hx_value_280 == nil {
			var hx_zero_281 *string
			return hx_zero_281
		}
		return hx_value_280.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_286 any) bool {
		if hx_value_286 == nil {
			var hx_zero_287 bool
			return hx_zero_287
		}
		return hx_value_286.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_284 any) *string {
		if hx_value_284 == nil {
			var hx_zero_285 *string
			return hx_zero_285
		}
		return hx_value_284.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_288 any) *haxe__ds__StringMap {
		if hx_value_288 == nil {
			var hx_zero_289 *haxe__ds__StringMap
			return hx_zero_289
		}
		return hx_value_288.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_290 any) map[string]any {
		if hx_value_290 == nil {
			var hx_zero_291 map[string]any
			return hx_zero_291
		}
		return hx_value_290.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_292 map[string]any) func() bool {
		hx_field_293 := hx_obj_292["hasNext"]
		if hx_field_293 == nil {
			var hx_zero_294 func() bool
			return hx_zero_294
		}
		return hx_field_293.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_295 map[string]any) func() *string {
			hx_field_296 := hx_obj_295["next"]
			if hx_field_296 == nil {
				var hx_zero_297 func() *string
				return hx_zero_297
			}
			return hx_field_296.(func() *string)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_298 any) map[string]any {
		if hx_value_298 == nil {
			var hx_zero_299 map[string]any
			return hx_zero_299
		}
		return hx_value_298.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_300 map[string]any) func() bool {
		hx_field_301 := hx_obj_300["hasNext"]
		if hx_field_301 == nil {
			var hx_zero_302 func() bool
			return hx_zero_302
		}
		return hx_field_301.(func() bool)
	}(iterator)() {
		key := func(hx_obj_303 map[string]any) func() *string {
			hx_field_304 := hx_obj_303["next"]
			if hx_field_304 == nil {
				var hx_zero_305 func() *string
				return hx_zero_305
			}
			return hx_field_304.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_306 map[string]any) func() bool {
			hx_field_307 := hx_obj_306["hasNext"]
			if hx_field_307 == nil {
				var hx_zero_308 func() bool
				return hx_zero_308
			}
			return hx_field_307.(func() bool)
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
