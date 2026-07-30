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
	hx_obj_233 := map[string]any{}
	hx_obj_233["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_233["next"] = func() *string {
		hx_post_234 := index
		index = int(int32((index + 1)))
		return keys[hx_post_234]
	}
	return hx_obj_233
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_235 := map[string]any{}
	hx_obj_235["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_235["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_236 := index
			index = int(int32((index + 1)))
			return hx_post_236
		}()])
	}
	return hx_obj_235
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_237 any) map[string]any {
		if hx_value_237 == nil {
			var hx_zero_238 map[string]any
			return hx_zero_238
		}
		return hx_value_237.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_239 := map[string]any{}
	hx_obj_239["hasNext"] = func() bool {
		return func(hx_obj_240 map[string]any) func() bool {
			hx_field_241 := hx_obj_240["hasNext"]
			if hx_field_241 == nil {
				var hx_zero_242 func() bool
				return hx_zero_242
			}
			return hx_field_241.(func() bool)
		}(keys)()
	}
	hx_obj_239["next"] = func() map[string]any {
		key := func(hx_obj_243 map[string]any) func() *string {
			hx_field_244 := hx_obj_243["next"]
			if hx_field_244 == nil {
				var hx_zero_245 func() *string
				return hx_zero_245
			}
			return hx_field_244.(func() *string)
		}(keys)()
		hx_obj_246 := map[string]any{}
		hx_obj_246["key"] = key
		hx_obj_246["value"] = _gthis.__hx_this.get(key)
		return hx_obj_246
	}
	return hx_obj_239
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_247 any) *string {
		if hx_value_247 == nil {
			var hx_zero_248 *string
			return hx_zero_248
		}
		return hx_value_247.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_249 any) *string {
		if hx_value_249 == nil {
			var hx_zero_250 *string
			return hx_zero_250
		}
		return hx_value_249.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_253 any) bool {
		if hx_value_253 == nil {
			var hx_zero_254 bool
			return hx_zero_254
		}
		return hx_value_253.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_251 any) *string {
		if hx_value_251 == nil {
			var hx_zero_252 *string
			return hx_zero_252
		}
		return hx_value_251.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_257 any) bool {
		if hx_value_257 == nil {
			var hx_zero_258 bool
			return hx_zero_258
		}
		return hx_value_257.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_255 any) *string {
		if hx_value_255 == nil {
			var hx_zero_256 *string
			return hx_zero_256
		}
		return hx_value_255.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_259 any) *haxe__ds__StringMap {
		if hx_value_259 == nil {
			var hx_zero_260 *haxe__ds__StringMap
			return hx_zero_260
		}
		return hx_value_259.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_261 any) map[string]any {
		if hx_value_261 == nil {
			var hx_zero_262 map[string]any
			return hx_zero_262
		}
		return hx_value_261.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_263 map[string]any) func() bool {
		hx_field_264 := hx_obj_263["hasNext"]
		if hx_field_264 == nil {
			var hx_zero_265 func() bool
			return hx_zero_265
		}
		return hx_field_264.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_266 map[string]any) func() *string {
			hx_field_267 := hx_obj_266["next"]
			if hx_field_267 == nil {
				var hx_zero_268 func() *string
				return hx_zero_268
			}
			return hx_field_267.(func() *string)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_269 any) map[string]any {
		if hx_value_269 == nil {
			var hx_zero_270 map[string]any
			return hx_zero_270
		}
		return hx_value_269.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_271 map[string]any) func() bool {
		hx_field_272 := hx_obj_271["hasNext"]
		if hx_field_272 == nil {
			var hx_zero_273 func() bool
			return hx_zero_273
		}
		return hx_field_272.(func() bool)
	}(iterator)() {
		key := func(hx_obj_274 map[string]any) func() *string {
			hx_field_275 := hx_obj_274["next"]
			if hx_field_275 == nil {
				var hx_zero_276 func() *string
				return hx_zero_276
			}
			return hx_field_275.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_277 map[string]any) func() bool {
			hx_field_278 := hx_obj_277["hasNext"]
			if hx_field_278 == nil {
				var hx_zero_279 func() bool
				return hx_zero_279
			}
			return hx_field_278.(func() bool)
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
