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
	hx_obj_248 := map[string]any{}
	hx_obj_248["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_248["next"] = func() *string {
		hx_post_249 := index
		index = int(int32((index + 1)))
		return keys[hx_post_249]
	}
	return hx_obj_248
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_250 := map[string]any{}
	hx_obj_250["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_250["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_251 := index
			index = int(int32((index + 1)))
			return hx_post_251
		}()])
	}
	return hx_obj_250
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_252 any) map[string]any {
		if hx_value_252 == nil {
			var hx_zero_253 map[string]any
			return hx_zero_253
		}
		return hx_value_252.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_254 := map[string]any{}
	hx_obj_254["hasNext"] = func() bool {
		return func(hx_obj_255 map[string]any) func() bool {
			hx_field_256 := hx_obj_255["hasNext"]
			if hx_field_256 == nil {
				var hx_zero_257 func() bool
				return hx_zero_257
			}
			return hx_field_256.(func() bool)
		}(keys)()
	}
	hx_obj_254["next"] = func() map[string]any {
		key := func(hx_obj_258 map[string]any) func() *string {
			hx_field_259 := hx_obj_258["next"]
			if hx_field_259 == nil {
				var hx_zero_260 func() *string
				return hx_zero_260
			}
			return hx_field_259.(func() *string)
		}(keys)()
		hx_obj_261 := map[string]any{}
		hx_obj_261["key"] = key
		hx_obj_261["value"] = _gthis.__hx_this.get(key)
		return hx_obj_261
	}
	return hx_obj_254
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_262 any) *string {
		if hx_value_262 == nil {
			var hx_zero_263 *string
			return hx_zero_263
		}
		return hx_value_262.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_264 any) *string {
		if hx_value_264 == nil {
			var hx_zero_265 *string
			return hx_zero_265
		}
		return hx_value_264.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_268 any) bool {
		if hx_value_268 == nil {
			var hx_zero_269 bool
			return hx_zero_269
		}
		return hx_value_268.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_266 any) *string {
		if hx_value_266 == nil {
			var hx_zero_267 *string
			return hx_zero_267
		}
		return hx_value_266.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_272 any) bool {
		if hx_value_272 == nil {
			var hx_zero_273 bool
			return hx_zero_273
		}
		return hx_value_272.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_270 any) *string {
		if hx_value_270 == nil {
			var hx_zero_271 *string
			return hx_zero_271
		}
		return hx_value_270.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_274 any) *haxe__ds__StringMap {
		if hx_value_274 == nil {
			var hx_zero_275 *haxe__ds__StringMap
			return hx_zero_275
		}
		return hx_value_274.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_276 any) map[string]any {
		if hx_value_276 == nil {
			var hx_zero_277 map[string]any
			return hx_zero_277
		}
		return hx_value_276.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_278 map[string]any) func() bool {
		hx_field_279 := hx_obj_278["hasNext"]
		if hx_field_279 == nil {
			var hx_zero_280 func() bool
			return hx_zero_280
		}
		return hx_field_279.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_281 map[string]any) func() *string {
			hx_field_282 := hx_obj_281["next"]
			if hx_field_282 == nil {
				var hx_zero_283 func() *string
				return hx_zero_283
			}
			return hx_field_282.(func() *string)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_284 any) map[string]any {
		if hx_value_284 == nil {
			var hx_zero_285 map[string]any
			return hx_zero_285
		}
		return hx_value_284.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_286 map[string]any) func() bool {
		hx_field_287 := hx_obj_286["hasNext"]
		if hx_field_287 == nil {
			var hx_zero_288 func() bool
			return hx_zero_288
		}
		return hx_field_287.(func() bool)
	}(iterator)() {
		key := func(hx_obj_289 map[string]any) func() *string {
			hx_field_290 := hx_obj_289["next"]
			if hx_field_290 == nil {
				var hx_zero_291 func() *string
				return hx_zero_291
			}
			return hx_field_290.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_292 map[string]any) func() bool {
			hx_field_293 := hx_obj_292["hasNext"]
			if hx_field_293 == nil {
				var hx_zero_294 func() bool
				return hx_zero_294
			}
			return hx_field_293.(func() bool)
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
