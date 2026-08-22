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
	hx_obj_1 := map[string]any{}
	hx_obj_1["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_1["next"] = func() any {
		return keys[func() int {
			hx_post_2 := index
			index = int(int32((index + 1)))
			return hx_post_2
		}()]
	}
	return hx_obj_1
}

func (self *haxe__ds__ObjectMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.ObjectMapKeys(self.h)
	index := 0
	hx_obj_3 := map[string]any{}
	hx_obj_3["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_3["next"] = func() any {
		return hxrt.ObjectMapGet(_gthis.h, keys[func() int {
			hx_post_4 := index
			index = int(int32((index + 1)))
			return hx_post_4
		}()])
	}
	return hx_obj_3
}

func (self *haxe__ds__ObjectMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_5 any) map[string]any {
		if hx_value_5 == nil {
			var hx_zero_6 map[string]any
			return hx_zero_6
		}
		return hx_value_5.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_7 := map[string]any{}
	hx_obj_7["hasNext"] = func() bool {
		return func(hx_obj_8 map[string]any) func() bool {
			hx_field_9 := hx_obj_8["hasNext"]
			if hx_field_9 == nil {
				var hx_zero_10 func() bool
				return hx_zero_10
			}
			return hx_field_9.(func() bool)
		}(keys)()
	}
	hx_obj_7["next"] = func() map[string]any {
		var key any = func(hx_obj_11 map[string]any) func() any {
			hx_field_12 := hx_obj_11["next"]
			if hx_field_12 == nil {
				var hx_zero_13 func() any
				return hx_zero_13
			}
			return hx_field_12.(func() any)
		}(keys)()
		hx_obj_14 := map[string]any{}
		hx_obj_14["key"] = key
		hx_obj_14["value"] = _gthis.__hx_this.get(key)
		return hx_obj_14
	}
	return hx_obj_7
}

func (self *haxe__ds__ObjectMap) getIMap(key any) any {
	return self.__hx_this.get(key)
}

func (self *haxe__ds__ObjectMap) setIMap(key any, value any) {
	self.__hx_this.set(key, value)
}

func (self *haxe__ds__ObjectMap) existsIMap(key any) bool {
	return func(hx_value_15 any) bool {
		if hx_value_15 == nil {
			var hx_zero_16 bool
			return hx_zero_16
		}
		return hx_value_15.(bool)
	}(self.__hx_this.exists(key))
}

func (self *haxe__ds__ObjectMap) removeIMap(key any) bool {
	return func(hx_value_17 any) bool {
		if hx_value_17 == nil {
			var hx_zero_18 bool
			return hx_zero_18
		}
		return hx_value_17.(bool)
	}(self.__hx_this.remove(key))
}

func (self *haxe__ds__ObjectMap) copyIMap() haxe__IMap {
	return func(hx_value_19 any) *haxe__ds__ObjectMap {
		if hx_value_19 == nil {
			var hx_zero_20 *haxe__ds__ObjectMap
			return hx_zero_20
		}
		return hx_value_19.(*haxe__ds__ObjectMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__ObjectMap) copy() *haxe__ds__ObjectMap {
	copied := New_haxe__ds__ObjectMap()
	key := func(hx_value_21 any) map[string]any {
		if hx_value_21 == nil {
			var hx_zero_22 map[string]any
			return hx_zero_22
		}
		return hx_value_21.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_23 map[string]any) func() bool {
		hx_field_24 := hx_obj_23["hasNext"]
		if hx_field_24 == nil {
			var hx_zero_25 func() bool
			return hx_zero_25
		}
		return hx_field_24.(func() bool)
	}(key)() {
		var key_1 any = func(hx_obj_26 map[string]any) func() any {
			hx_field_27 := hx_obj_26["next"]
			if hx_field_27 == nil {
				var hx_zero_28 func() any
				return hx_zero_28
			}
			return hx_field_27.(func() any)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__ObjectMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_29 any) map[string]any {
		if hx_value_29 == nil {
			var hx_zero_30 map[string]any
			return hx_zero_30
		}
		return hx_value_29.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_31 map[string]any) func() bool {
		hx_field_32 := hx_obj_31["hasNext"]
		if hx_field_32 == nil {
			var hx_zero_33 func() bool
			return hx_zero_33
		}
		return hx_field_32.(func() bool)
	}(iterator)() {
		var key any = func(hx_obj_34 map[string]any) func() any {
			hx_field_35 := hx_obj_34["next"]
			if hx_field_35 == nil {
				var hx_zero_36 func() any
				return hx_zero_36
			}
			return hx_field_35.(func() any)
		}(iterator)()
		x := hxrt.StdString(key)
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x_1 := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x_1))
		if func(hx_obj_37 map[string]any) func() bool {
			hx_field_38 := hx_obj_37["hasNext"]
			if hx_field_38 == nil {
				var hx_zero_39 func() bool
				return hx_zero_39
			}
			return hx_field_38.(func() bool)
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

func (self *haxe__ds__ObjectMap) String() string {
	return *self.__hx_this.toString()
}
