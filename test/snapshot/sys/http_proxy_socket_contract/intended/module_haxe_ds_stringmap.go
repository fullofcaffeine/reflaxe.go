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
	hx_obj_255 := map[string]any{}
	hx_obj_255["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_255["next"] = func() *string {
		hx_post_256 := index
		index = int(int32((index + 1)))
		return keys[hx_post_256]
	}
	return hx_obj_255
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_257 := map[string]any{}
	hx_obj_257["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_257["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_258 := index
			index = int(int32((index + 1)))
			return hx_post_258
		}()])
	}
	return hx_obj_257
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_259 any) map[string]any {
		if hx_value_259 == nil {
			var hx_zero_260 map[string]any
			return hx_zero_260
		}
		return hx_value_259.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_261 := map[string]any{}
	hx_obj_261["hasNext"] = func() bool {
		return func(hx_obj_262 map[string]any) func() bool {
			hx_field_263 := hx_obj_262["hasNext"]
			if hx_field_263 == nil {
				var hx_zero_264 func() bool
				return hx_zero_264
			}
			return hx_field_263.(func() bool)
		}(keys)()
	}
	hx_obj_261["next"] = func() map[string]any {
		key := func(hx_obj_265 map[string]any) func() *string {
			hx_field_266 := hx_obj_265["next"]
			if hx_field_266 == nil {
				var hx_zero_267 func() *string
				return hx_zero_267
			}
			return hx_field_266.(func() *string)
		}(keys)()
		hx_obj_268 := map[string]any{}
		hx_obj_268["key"] = key
		hx_obj_268["value"] = _gthis.__hx_this.get(key)
		return hx_obj_268
	}
	return hx_obj_261
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_269 any) *string {
		if hx_value_269 == nil {
			var hx_zero_270 *string
			return hx_zero_270
		}
		return hx_value_269.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_271 any) *string {
		if hx_value_271 == nil {
			var hx_zero_272 *string
			return hx_zero_272
		}
		return hx_value_271.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_275 any) bool {
		if hx_value_275 == nil {
			var hx_zero_276 bool
			return hx_zero_276
		}
		return hx_value_275.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_273 any) *string {
		if hx_value_273 == nil {
			var hx_zero_274 *string
			return hx_zero_274
		}
		return hx_value_273.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_279 any) bool {
		if hx_value_279 == nil {
			var hx_zero_280 bool
			return hx_zero_280
		}
		return hx_value_279.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_277 any) *string {
		if hx_value_277 == nil {
			var hx_zero_278 *string
			return hx_zero_278
		}
		return hx_value_277.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_281 any) *haxe__ds__StringMap {
		if hx_value_281 == nil {
			var hx_zero_282 *haxe__ds__StringMap
			return hx_zero_282
		}
		return hx_value_281.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_283 any) map[string]any {
		if hx_value_283 == nil {
			var hx_zero_284 map[string]any
			return hx_zero_284
		}
		return hx_value_283.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_285 map[string]any) func() bool {
		hx_field_286 := hx_obj_285["hasNext"]
		if hx_field_286 == nil {
			var hx_zero_287 func() bool
			return hx_zero_287
		}
		return hx_field_286.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_288 map[string]any) func() *string {
			hx_field_289 := hx_obj_288["next"]
			if hx_field_289 == nil {
				var hx_zero_290 func() *string
				return hx_zero_290
			}
			return hx_field_289.(func() *string)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_291 any) map[string]any {
		if hx_value_291 == nil {
			var hx_zero_292 map[string]any
			return hx_zero_292
		}
		return hx_value_291.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_293 map[string]any) func() bool {
		hx_field_294 := hx_obj_293["hasNext"]
		if hx_field_294 == nil {
			var hx_zero_295 func() bool
			return hx_zero_295
		}
		return hx_field_294.(func() bool)
	}(iterator)() {
		key := func(hx_obj_296 map[string]any) func() *string {
			hx_field_297 := hx_obj_296["next"]
			if hx_field_297 == nil {
				var hx_zero_298 func() *string
				return hx_zero_298
			}
			return hx_field_297.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_299 map[string]any) func() bool {
			hx_field_300 := hx_obj_299["hasNext"]
			if hx_field_300 == nil {
				var hx_zero_301 func() bool
				return hx_zero_301
			}
			return hx_field_300.(func() bool)
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
