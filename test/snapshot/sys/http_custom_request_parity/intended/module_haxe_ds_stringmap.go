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
	hx_obj_236 := map[string]any{}
	hx_obj_236["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_236["next"] = func() *string {
		hx_post_237 := index
		index = int(int32((index + 1)))
		return keys[hx_post_237]
	}
	return hx_obj_236
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_238 := map[string]any{}
	hx_obj_238["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_238["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_239 := index
			index = int(int32((index + 1)))
			return hx_post_239
		}()])
	}
	return hx_obj_238
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_240 any) map[string]any {
		if hx_value_240 == nil {
			var hx_zero_241 map[string]any
			return hx_zero_241
		}
		return hx_value_240.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_242 := map[string]any{}
	hx_obj_242["hasNext"] = func() bool {
		return func(hx_obj_243 map[string]any) func() bool {
			hx_field_244 := hx_obj_243["hasNext"]
			if hx_field_244 == nil {
				var hx_zero_245 func() bool
				return hx_zero_245
			}
			return hx_field_244.(func() bool)
		}(keys)()
	}
	hx_obj_242["next"] = func() map[string]any {
		key := func(hx_obj_246 map[string]any) func() *string {
			hx_field_247 := hx_obj_246["next"]
			if hx_field_247 == nil {
				var hx_zero_248 func() *string
				return hx_zero_248
			}
			return hx_field_247.(func() *string)
		}(keys)()
		hx_obj_249 := map[string]any{}
		hx_obj_249["key"] = key
		hx_obj_249["value"] = _gthis.__hx_this.get(key)
		return hx_obj_249
	}
	return hx_obj_242
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_250 any) *string {
		if hx_value_250 == nil {
			var hx_zero_251 *string
			return hx_zero_251
		}
		return hx_value_250.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_252 any) *string {
		if hx_value_252 == nil {
			var hx_zero_253 *string
			return hx_zero_253
		}
		return hx_value_252.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_256 any) bool {
		if hx_value_256 == nil {
			var hx_zero_257 bool
			return hx_zero_257
		}
		return hx_value_256.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_254 any) *string {
		if hx_value_254 == nil {
			var hx_zero_255 *string
			return hx_zero_255
		}
		return hx_value_254.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_260 any) bool {
		if hx_value_260 == nil {
			var hx_zero_261 bool
			return hx_zero_261
		}
		return hx_value_260.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_258 any) *string {
		if hx_value_258 == nil {
			var hx_zero_259 *string
			return hx_zero_259
		}
		return hx_value_258.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_262 any) *haxe__ds__StringMap {
		if hx_value_262 == nil {
			var hx_zero_263 *haxe__ds__StringMap
			return hx_zero_263
		}
		return hx_value_262.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_264 any) map[string]any {
		if hx_value_264 == nil {
			var hx_zero_265 map[string]any
			return hx_zero_265
		}
		return hx_value_264.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_266 map[string]any) func() bool {
		hx_field_267 := hx_obj_266["hasNext"]
		if hx_field_267 == nil {
			var hx_zero_268 func() bool
			return hx_zero_268
		}
		return hx_field_267.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_269 map[string]any) func() *string {
			hx_field_270 := hx_obj_269["next"]
			if hx_field_270 == nil {
				var hx_zero_271 func() *string
				return hx_zero_271
			}
			return hx_field_270.(func() *string)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_272 any) map[string]any {
		if hx_value_272 == nil {
			var hx_zero_273 map[string]any
			return hx_zero_273
		}
		return hx_value_272.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_274 map[string]any) func() bool {
		hx_field_275 := hx_obj_274["hasNext"]
		if hx_field_275 == nil {
			var hx_zero_276 func() bool
			return hx_zero_276
		}
		return hx_field_275.(func() bool)
	}(iterator)() {
		key := func(hx_obj_277 map[string]any) func() *string {
			hx_field_278 := hx_obj_277["next"]
			if hx_field_278 == nil {
				var hx_zero_279 func() *string
				return hx_zero_279
			}
			return hx_field_278.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_280 map[string]any) func() bool {
			hx_field_281 := hx_obj_280["hasNext"]
			if hx_field_281 == nil {
				var hx_zero_282 func() bool
				return hx_zero_282
			}
			return hx_field_281.(func() bool)
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
