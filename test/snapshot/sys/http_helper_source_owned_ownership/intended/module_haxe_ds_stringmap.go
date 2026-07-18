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
	hx_obj_261 := map[string]any{}
	hx_obj_261["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_261["next"] = func() *string {
		hx_post_262 := index
		index = int(int32((index + 1)))
		return keys[hx_post_262]
	}
	return hx_obj_261
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_263 := map[string]any{}
	hx_obj_263["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_263["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_264 := index
			index = int(int32((index + 1)))
			return hx_post_264
		}()])
	}
	return hx_obj_263
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_265 any) map[string]any {
		if hx_value_265 == nil {
			var hx_zero_266 map[string]any
			return hx_zero_266
		}
		return hx_value_265.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_267 := map[string]any{}
	hx_obj_267["hasNext"] = func() bool {
		return func(hx_obj_268 map[string]any) func() bool {
			hx_field_269 := hx_obj_268["hasNext"]
			if hx_field_269 == nil {
				var hx_zero_270 func() bool
				return hx_zero_270
			}
			return hx_field_269.(func() bool)
		}(keys)()
	}
	hx_obj_267["next"] = func() map[string]any {
		key := func(hx_obj_271 map[string]any) func() *string {
			hx_field_272 := hx_obj_271["next"]
			if hx_field_272 == nil {
				var hx_zero_273 func() *string
				return hx_zero_273
			}
			return hx_field_272.(func() *string)
		}(keys)()
		hx_obj_274 := map[string]any{}
		hx_obj_274["key"] = key
		hx_obj_274["value"] = _gthis.__hx_this.get(key)
		return hx_obj_274
	}
	return hx_obj_267
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_275 any) *string {
		if hx_value_275 == nil {
			var hx_zero_276 *string
			return hx_zero_276
		}
		return hx_value_275.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_277 any) *string {
		if hx_value_277 == nil {
			var hx_zero_278 *string
			return hx_zero_278
		}
		return hx_value_277.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_281 any) bool {
		if hx_value_281 == nil {
			var hx_zero_282 bool
			return hx_zero_282
		}
		return hx_value_281.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_279 any) *string {
		if hx_value_279 == nil {
			var hx_zero_280 *string
			return hx_zero_280
		}
		return hx_value_279.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_285 any) bool {
		if hx_value_285 == nil {
			var hx_zero_286 bool
			return hx_zero_286
		}
		return hx_value_285.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_283 any) *string {
		if hx_value_283 == nil {
			var hx_zero_284 *string
			return hx_zero_284
		}
		return hx_value_283.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_287 any) *haxe__ds__StringMap {
		if hx_value_287 == nil {
			var hx_zero_288 *haxe__ds__StringMap
			return hx_zero_288
		}
		return hx_value_287.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_289 any) map[string]any {
		if hx_value_289 == nil {
			var hx_zero_290 map[string]any
			return hx_zero_290
		}
		return hx_value_289.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_291 map[string]any) func() bool {
		hx_field_292 := hx_obj_291["hasNext"]
		if hx_field_292 == nil {
			var hx_zero_293 func() bool
			return hx_zero_293
		}
		return hx_field_292.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_294 map[string]any) func() *string {
			hx_field_295 := hx_obj_294["next"]
			if hx_field_295 == nil {
				var hx_zero_296 func() *string
				return hx_zero_296
			}
			return hx_field_295.(func() *string)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_297 any) map[string]any {
		if hx_value_297 == nil {
			var hx_zero_298 map[string]any
			return hx_zero_298
		}
		return hx_value_297.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_299 map[string]any) func() bool {
		hx_field_300 := hx_obj_299["hasNext"]
		if hx_field_300 == nil {
			var hx_zero_301 func() bool
			return hx_zero_301
		}
		return hx_field_300.(func() bool)
	}(iterator)() {
		key := func(hx_obj_302 map[string]any) func() *string {
			hx_field_303 := hx_obj_302["next"]
			if hx_field_303 == nil {
				var hx_zero_304 func() *string
				return hx_zero_304
			}
			return hx_field_303.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_305 map[string]any) func() bool {
			hx_field_306 := hx_obj_305["hasNext"]
			if hx_field_306 == nil {
				var hx_zero_307 func() bool
				return hx_zero_307
			}
			return hx_field_306.(func() bool)
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
