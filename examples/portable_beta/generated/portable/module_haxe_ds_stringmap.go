package main

import "examples_portable_beta/hxrt"

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
	hx_obj_1 := map[string]any{}
	hx_obj_1["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_1["next"] = func() *string {
		hx_post_2 := index
		index = int(int32((index + 1)))
		return keys[hx_post_2]
	}
	return hx_obj_1
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_3 := map[string]any{}
	hx_obj_3["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_3["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_4 := index
			index = int(int32((index + 1)))
			return hx_post_4
		}()])
	}
	return hx_obj_3
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
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
		key := func(hx_obj_11 map[string]any) func() *string {
			hx_field_12 := hx_obj_11["next"]
			if hx_field_12 == nil {
				var hx_zero_13 func() *string
				return hx_zero_13
			}
			return hx_field_12.(func() *string)
		}(keys)()
		hx_obj_14 := map[string]any{}
		hx_obj_14["key"] = key
		hx_obj_14["value"] = _gthis.__hx_this.get(key)
		return hx_obj_14
	}
	return hx_obj_7
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_15 any) *string {
		if hx_value_15 == nil {
			var hx_zero_16 *string
			return hx_zero_16
		}
		return hx_value_15.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_17 any) *string {
		if hx_value_17 == nil {
			var hx_zero_18 *string
			return hx_zero_18
		}
		return hx_value_17.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_21 any) bool {
		if hx_value_21 == nil {
			var hx_zero_22 bool
			return hx_zero_22
		}
		return hx_value_21.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_19 any) *string {
		if hx_value_19 == nil {
			var hx_zero_20 *string
			return hx_zero_20
		}
		return hx_value_19.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_25 any) bool {
		if hx_value_25 == nil {
			var hx_zero_26 bool
			return hx_zero_26
		}
		return hx_value_25.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_23 any) *string {
		if hx_value_23 == nil {
			var hx_zero_24 *string
			return hx_zero_24
		}
		return hx_value_23.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_27 any) *haxe__ds__StringMap {
		if hx_value_27 == nil {
			var hx_zero_28 *haxe__ds__StringMap
			return hx_zero_28
		}
		return hx_value_27.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_29 any) map[string]any {
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
	}(key)() {
		key_1 := func(hx_obj_34 map[string]any) func() *string {
			hx_field_35 := hx_obj_34["next"]
			if hx_field_35 == nil {
				var hx_zero_36 func() *string
				return hx_zero_36
			}
			return hx_field_35.(func() *string)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_37 any) map[string]any {
		if hx_value_37 == nil {
			var hx_zero_38 map[string]any
			return hx_zero_38
		}
		return hx_value_37.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_39 map[string]any) func() bool {
		hx_field_40 := hx_obj_39["hasNext"]
		if hx_field_40 == nil {
			var hx_zero_41 func() bool
			return hx_zero_41
		}
		return hx_field_40.(func() bool)
	}(iterator)() {
		key := func(hx_obj_42 map[string]any) func() *string {
			hx_field_43 := hx_obj_42["next"]
			if hx_field_43 == nil {
				var hx_zero_44 func() *string
				return hx_zero_44
			}
			return hx_field_43.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_45 map[string]any) func() bool {
			hx_field_46 := hx_obj_45["hasNext"]
			if hx_field_46 == nil {
				var hx_zero_47 func() bool
				return hx_zero_47
			}
			return hx_field_46.(func() bool)
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
