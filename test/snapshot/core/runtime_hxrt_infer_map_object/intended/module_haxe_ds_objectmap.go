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
	hx_obj_12 := map[string]any{}
	hx_obj_12["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_12["next"] = func() any {
		return keys[func() int {
			hx_post_13 := index
			index = int(int32((index + 1)))
			return hx_post_13
		}()]
	}
	return hx_obj_12
}

func (self *haxe__ds__ObjectMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.ObjectMapKeys(self.h)
	index := 0
	hx_obj_14 := map[string]any{}
	hx_obj_14["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_14["next"] = func() any {
		return hxrt.ObjectMapGet(_gthis.h, keys[func() int {
			hx_post_15 := index
			index = int(int32((index + 1)))
			return hx_post_15
		}()])
	}
	return hx_obj_14
}

func (self *haxe__ds__ObjectMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_16 any) map[string]any {
		if hx_value_16 == nil {
			var hx_zero_17 map[string]any
			return hx_zero_17
		}
		return hx_value_16.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_18 := map[string]any{}
	hx_obj_18["hasNext"] = func() bool {
		return func(hx_obj_19 map[string]any) func() bool {
			hx_field_20 := hx_obj_19["hasNext"]
			if hx_field_20 == nil {
				var hx_zero_21 func() bool
				return hx_zero_21
			}
			return hx_field_20.(func() bool)
		}(keys)()
	}
	hx_obj_18["next"] = func() map[string]any {
		var key any = func(hx_obj_22 map[string]any) func() any {
			hx_field_23 := hx_obj_22["next"]
			if hx_field_23 == nil {
				var hx_zero_24 func() any
				return hx_zero_24
			}
			return hx_field_23.(func() any)
		}(keys)()
		hx_obj_25 := map[string]any{}
		hx_obj_25["key"] = key
		hx_obj_25["value"] = _gthis.__hx_this.get(key)
		return hx_obj_25
	}
	return hx_obj_18
}

func (self *haxe__ds__ObjectMap) getIMap(key any) any {
	return self.__hx_this.get(key)
}

func (self *haxe__ds__ObjectMap) setIMap(key any, value any) {
	self.__hx_this.set(key, value)
}

func (self *haxe__ds__ObjectMap) existsIMap(key any) bool {
	return func(hx_value_26 any) bool {
		if hx_value_26 == nil {
			var hx_zero_27 bool
			return hx_zero_27
		}
		return hx_value_26.(bool)
	}(self.__hx_this.exists(key))
}

func (self *haxe__ds__ObjectMap) removeIMap(key any) bool {
	return func(hx_value_28 any) bool {
		if hx_value_28 == nil {
			var hx_zero_29 bool
			return hx_zero_29
		}
		return hx_value_28.(bool)
	}(self.__hx_this.remove(key))
}

func (self *haxe__ds__ObjectMap) copyIMap() haxe__IMap {
	return func(hx_value_30 any) *haxe__ds__ObjectMap {
		if hx_value_30 == nil {
			var hx_zero_31 *haxe__ds__ObjectMap
			return hx_zero_31
		}
		return hx_value_30.(*haxe__ds__ObjectMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__ObjectMap) copy() *haxe__ds__ObjectMap {
	copied := New_haxe__ds__ObjectMap()
	key := func(hx_value_32 any) map[string]any {
		if hx_value_32 == nil {
			var hx_zero_33 map[string]any
			return hx_zero_33
		}
		return hx_value_32.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_34 map[string]any) func() bool {
		hx_field_35 := hx_obj_34["hasNext"]
		if hx_field_35 == nil {
			var hx_zero_36 func() bool
			return hx_zero_36
		}
		return hx_field_35.(func() bool)
	}(key)() {
		var key_1 any = func(hx_obj_37 map[string]any) func() any {
			hx_field_38 := hx_obj_37["next"]
			if hx_field_38 == nil {
				var hx_zero_39 func() any
				return hx_zero_39
			}
			return hx_field_38.(func() any)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__ObjectMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_40 any) map[string]any {
		if hx_value_40 == nil {
			var hx_zero_41 map[string]any
			return hx_zero_41
		}
		return hx_value_40.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_42 map[string]any) func() bool {
		hx_field_43 := hx_obj_42["hasNext"]
		if hx_field_43 == nil {
			var hx_zero_44 func() bool
			return hx_zero_44
		}
		return hx_field_43.(func() bool)
	}(iterator)() {
		var key any = func(hx_obj_45 map[string]any) func() any {
			hx_field_46 := hx_obj_45["next"]
			if hx_field_46 == nil {
				var hx_zero_47 func() any
				return hx_zero_47
			}
			return hx_field_46.(func() any)
		}(iterator)()
		x := hxrt.StdString(key)
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x_1 := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x_1))
		if func(hx_obj_48 map[string]any) func() bool {
			hx_field_49 := hx_obj_48["hasNext"]
			if hx_field_49 == nil {
				var hx_zero_50 func() bool
				return hx_zero_50
			}
			return hx_field_49.(func() bool)
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
