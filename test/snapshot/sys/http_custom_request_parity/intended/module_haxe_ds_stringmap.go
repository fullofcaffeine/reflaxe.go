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
	hx_obj_240 := map[string]any{}
	hx_obj_240["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_240["next"] = func() *string {
		hx_post_241 := index
		index = int(int32((index + 1)))
		return keys[hx_post_241]
	}
	return hx_obj_240
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_242 := map[string]any{}
	hx_obj_242["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_242["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_243 := index
			index = int(int32((index + 1)))
			return hx_post_243
		}()])
	}
	return hx_obj_242
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_244 any) map[string]any {
		if hx_value_244 == nil {
			var hx_zero_245 map[string]any
			return hx_zero_245
		}
		return hx_value_244.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_246 := map[string]any{}
	hx_obj_246["hasNext"] = func() bool {
		return func(hx_obj_247 map[string]any) func() bool {
			hx_field_248 := hx_obj_247["hasNext"]
			if hx_field_248 == nil {
				var hx_zero_249 func() bool
				return hx_zero_249
			}
			return hx_field_248.(func() bool)
		}(keys)()
	}
	hx_obj_246["next"] = func() map[string]any {
		key := func(hx_obj_250 map[string]any) func() *string {
			hx_field_251 := hx_obj_250["next"]
			if hx_field_251 == nil {
				var hx_zero_252 func() *string
				return hx_zero_252
			}
			return hx_field_251.(func() *string)
		}(keys)()
		hx_obj_253 := map[string]any{}
		hx_obj_253["key"] = key
		hx_obj_253["value"] = _gthis.__hx_this.get(key)
		return hx_obj_253
	}
	return hx_obj_246
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_254 any) *string {
		if hx_value_254 == nil {
			var hx_zero_255 *string
			return hx_zero_255
		}
		return hx_value_254.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_256 any) *string {
		if hx_value_256 == nil {
			var hx_zero_257 *string
			return hx_zero_257
		}
		return hx_value_256.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_260 any) bool {
		if hx_value_260 == nil {
			var hx_zero_261 bool
			return hx_zero_261
		}
		return hx_value_260.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_258 any) *string {
		if hx_value_258 == nil {
			var hx_zero_259 *string
			return hx_zero_259
		}
		return hx_value_258.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_264 any) bool {
		if hx_value_264 == nil {
			var hx_zero_265 bool
			return hx_zero_265
		}
		return hx_value_264.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_262 any) *string {
		if hx_value_262 == nil {
			var hx_zero_263 *string
			return hx_zero_263
		}
		return hx_value_262.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_266 any) *haxe__ds__StringMap {
		if hx_value_266 == nil {
			var hx_zero_267 *haxe__ds__StringMap
			return hx_zero_267
		}
		return hx_value_266.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_268 any) map[string]any {
		if hx_value_268 == nil {
			var hx_zero_269 map[string]any
			return hx_zero_269
		}
		return hx_value_268.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_270 map[string]any) func() bool {
		hx_field_271 := hx_obj_270["hasNext"]
		if hx_field_271 == nil {
			var hx_zero_272 func() bool
			return hx_zero_272
		}
		return hx_field_271.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_273 map[string]any) func() *string {
			hx_field_274 := hx_obj_273["next"]
			if hx_field_274 == nil {
				var hx_zero_275 func() *string
				return hx_zero_275
			}
			return hx_field_274.(func() *string)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_276 any) map[string]any {
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
	}(iterator)() {
		key := func(hx_obj_281 map[string]any) func() *string {
			hx_field_282 := hx_obj_281["next"]
			if hx_field_282 == nil {
				var hx_zero_283 func() *string
				return hx_zero_283
			}
			return hx_field_282.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_284 map[string]any) func() bool {
			hx_field_285 := hx_obj_284["hasNext"]
			if hx_field_285 == nil {
				var hx_zero_286 func() bool
				return hx_zero_286
			}
			return hx_field_285.(func() bool)
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
