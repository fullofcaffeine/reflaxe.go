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
	hx_obj_254 := map[string]any{}
	hx_obj_254["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_254["next"] = func() *string {
		hx_post_255 := index
		index = int(int32((index + 1)))
		return keys[hx_post_255]
	}
	return hx_obj_254
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_256 := map[string]any{}
	hx_obj_256["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_256["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_257 := index
			index = int(int32((index + 1)))
			return hx_post_257
		}()])
	}
	return hx_obj_256
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_258 any) map[string]any {
		if hx_value_258 == nil {
			var hx_zero_259 map[string]any
			return hx_zero_259
		}
		return hx_value_258.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_260 := map[string]any{}
	hx_obj_260["hasNext"] = func() bool {
		return func(hx_obj_261 map[string]any) func() bool {
			hx_field_262 := hx_obj_261["hasNext"]
			if hx_field_262 == nil {
				var hx_zero_263 func() bool
				return hx_zero_263
			}
			return hx_field_262.(func() bool)
		}(keys)()
	}
	hx_obj_260["next"] = func() map[string]any {
		key := func(hx_obj_264 map[string]any) func() *string {
			hx_field_265 := hx_obj_264["next"]
			if hx_field_265 == nil {
				var hx_zero_266 func() *string
				return hx_zero_266
			}
			return hx_field_265.(func() *string)
		}(keys)()
		hx_obj_267 := map[string]any{}
		hx_obj_267["key"] = key
		hx_obj_267["value"] = _gthis.__hx_this.get(key)
		return hx_obj_267
	}
	return hx_obj_260
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_268 any) *string {
		if hx_value_268 == nil {
			var hx_zero_269 *string
			return hx_zero_269
		}
		return hx_value_268.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_270 any) *string {
		if hx_value_270 == nil {
			var hx_zero_271 *string
			return hx_zero_271
		}
		return hx_value_270.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_274 any) bool {
		if hx_value_274 == nil {
			var hx_zero_275 bool
			return hx_zero_275
		}
		return hx_value_274.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_272 any) *string {
		if hx_value_272 == nil {
			var hx_zero_273 *string
			return hx_zero_273
		}
		return hx_value_272.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_278 any) bool {
		if hx_value_278 == nil {
			var hx_zero_279 bool
			return hx_zero_279
		}
		return hx_value_278.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_276 any) *string {
		if hx_value_276 == nil {
			var hx_zero_277 *string
			return hx_zero_277
		}
		return hx_value_276.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_280 any) *haxe__ds__StringMap {
		if hx_value_280 == nil {
			var hx_zero_281 *haxe__ds__StringMap
			return hx_zero_281
		}
		return hx_value_280.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_282 any) map[string]any {
		if hx_value_282 == nil {
			var hx_zero_283 map[string]any
			return hx_zero_283
		}
		return hx_value_282.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_284 map[string]any) func() bool {
		hx_field_285 := hx_obj_284["hasNext"]
		if hx_field_285 == nil {
			var hx_zero_286 func() bool
			return hx_zero_286
		}
		return hx_field_285.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_287 map[string]any) func() *string {
			hx_field_288 := hx_obj_287["next"]
			if hx_field_288 == nil {
				var hx_zero_289 func() *string
				return hx_zero_289
			}
			return hx_field_288.(func() *string)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_290 any) map[string]any {
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
	}(iterator)() {
		key := func(hx_obj_295 map[string]any) func() *string {
			hx_field_296 := hx_obj_295["next"]
			if hx_field_296 == nil {
				var hx_zero_297 func() *string
				return hx_zero_297
			}
			return hx_field_296.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_298 map[string]any) func() bool {
			hx_field_299 := hx_obj_298["hasNext"]
			if hx_field_299 == nil {
				var hx_zero_300 func() bool
				return hx_zero_300
			}
			return hx_field_299.(func() bool)
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
