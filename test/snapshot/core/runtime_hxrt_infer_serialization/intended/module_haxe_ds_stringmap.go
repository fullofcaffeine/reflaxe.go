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
	hx_obj_6 := map[string]any{}
	hx_obj_6["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_6["next"] = func() *string {
		hx_post_7 := index
		index = int(int32((index + 1)))
		return keys[hx_post_7]
	}
	return hx_obj_6
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_8 := map[string]any{}
	hx_obj_8["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_8["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_9 := index
			index = int(int32((index + 1)))
			return hx_post_9
		}()])
	}
	return hx_obj_8
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_10 any) map[string]any {
		if hx_value_10 == nil {
			var hx_zero_11 map[string]any
			return hx_zero_11
		}
		return hx_value_10.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_12 := map[string]any{}
	hx_obj_12["hasNext"] = func() bool {
		return func(hx_obj_13 map[string]any) func() bool {
			hx_field_14 := hx_obj_13["hasNext"]
			if hx_field_14 == nil {
				var hx_zero_15 func() bool
				return hx_zero_15
			}
			return hx_field_14.(func() bool)
		}(keys)()
	}
	hx_obj_12["next"] = func() map[string]any {
		key := func(hx_obj_16 map[string]any) func() *string {
			hx_field_17 := hx_obj_16["next"]
			if hx_field_17 == nil {
				var hx_zero_18 func() *string
				return hx_zero_18
			}
			return hx_field_17.(func() *string)
		}(keys)()
		hx_obj_19 := map[string]any{}
		hx_obj_19["key"] = key
		hx_obj_19["value"] = _gthis.__hx_this.get(key)
		return hx_obj_19
	}
	return hx_obj_12
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_20 any) *string {
		if hx_value_20 == nil {
			var hx_zero_21 *string
			return hx_zero_21
		}
		return hx_value_20.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_22 any) *string {
		if hx_value_22 == nil {
			var hx_zero_23 *string
			return hx_zero_23
		}
		return hx_value_22.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_26 any) bool {
		if hx_value_26 == nil {
			var hx_zero_27 bool
			return hx_zero_27
		}
		return hx_value_26.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_24 any) *string {
		if hx_value_24 == nil {
			var hx_zero_25 *string
			return hx_zero_25
		}
		return hx_value_24.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_30 any) bool {
		if hx_value_30 == nil {
			var hx_zero_31 bool
			return hx_zero_31
		}
		return hx_value_30.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_28 any) *string {
		if hx_value_28 == nil {
			var hx_zero_29 *string
			return hx_zero_29
		}
		return hx_value_28.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_32 any) *haxe__ds__StringMap {
		if hx_value_32 == nil {
			var hx_zero_33 *haxe__ds__StringMap
			return hx_zero_33
		}
		return hx_value_32.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_34 any) map[string]any {
		if hx_value_34 == nil {
			var hx_zero_35 map[string]any
			return hx_zero_35
		}
		return hx_value_34.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_36 map[string]any) func() bool {
		hx_field_37 := hx_obj_36["hasNext"]
		if hx_field_37 == nil {
			var hx_zero_38 func() bool
			return hx_zero_38
		}
		return hx_field_37.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_39 map[string]any) func() *string {
			hx_field_40 := hx_obj_39["next"]
			if hx_field_40 == nil {
				var hx_zero_41 func() *string
				return hx_zero_41
			}
			return hx_field_40.(func() *string)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_42 any) map[string]any {
		if hx_value_42 == nil {
			var hx_zero_43 map[string]any
			return hx_zero_43
		}
		return hx_value_42.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_44 map[string]any) func() bool {
		hx_field_45 := hx_obj_44["hasNext"]
		if hx_field_45 == nil {
			var hx_zero_46 func() bool
			return hx_zero_46
		}
		return hx_field_45.(func() bool)
	}(iterator)() {
		key := func(hx_obj_47 map[string]any) func() *string {
			hx_field_48 := hx_obj_47["next"]
			if hx_field_48 == nil {
				var hx_zero_49 func() *string
				return hx_zero_49
			}
			return hx_field_48.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_50 map[string]any) func() bool {
			hx_field_51 := hx_obj_50["hasNext"]
			if hx_field_51 == nil {
				var hx_zero_52 func() bool
				return hx_zero_52
			}
			return hx_field_51.(func() bool)
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
