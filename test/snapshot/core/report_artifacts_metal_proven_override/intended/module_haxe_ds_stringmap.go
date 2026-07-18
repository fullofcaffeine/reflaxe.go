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
	hx_obj_16 := map[string]any{}
	hx_obj_16["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_16["next"] = func() *string {
		hx_post_17 := index
		index = int(int32((index + 1)))
		return keys[hx_post_17]
	}
	return hx_obj_16
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_18 := map[string]any{}
	hx_obj_18["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_18["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_19 := index
			index = int(int32((index + 1)))
			return hx_post_19
		}()])
	}
	return hx_obj_18
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_20 any) map[string]any {
		if hx_value_20 == nil {
			var hx_zero_21 map[string]any
			return hx_zero_21
		}
		return hx_value_20.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_22 := map[string]any{}
	hx_obj_22["hasNext"] = func() bool {
		return func(hx_obj_23 map[string]any) func() bool {
			hx_field_24 := hx_obj_23["hasNext"]
			if hx_field_24 == nil {
				var hx_zero_25 func() bool
				return hx_zero_25
			}
			return hx_field_24.(func() bool)
		}(keys)()
	}
	hx_obj_22["next"] = func() map[string]any {
		key := func(hx_obj_26 map[string]any) func() *string {
			hx_field_27 := hx_obj_26["next"]
			if hx_field_27 == nil {
				var hx_zero_28 func() *string
				return hx_zero_28
			}
			return hx_field_27.(func() *string)
		}(keys)()
		hx_obj_29 := map[string]any{}
		hx_obj_29["key"] = key
		hx_obj_29["value"] = _gthis.__hx_this.get(key)
		return hx_obj_29
	}
	return hx_obj_22
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_30 any) *string {
		if hx_value_30 == nil {
			var hx_zero_31 *string
			return hx_zero_31
		}
		return hx_value_30.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_32 any) *string {
		if hx_value_32 == nil {
			var hx_zero_33 *string
			return hx_zero_33
		}
		return hx_value_32.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_36 any) bool {
		if hx_value_36 == nil {
			var hx_zero_37 bool
			return hx_zero_37
		}
		return hx_value_36.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_34 any) *string {
		if hx_value_34 == nil {
			var hx_zero_35 *string
			return hx_zero_35
		}
		return hx_value_34.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_40 any) bool {
		if hx_value_40 == nil {
			var hx_zero_41 bool
			return hx_zero_41
		}
		return hx_value_40.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_38 any) *string {
		if hx_value_38 == nil {
			var hx_zero_39 *string
			return hx_zero_39
		}
		return hx_value_38.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_42 any) *haxe__ds__StringMap {
		if hx_value_42 == nil {
			var hx_zero_43 *haxe__ds__StringMap
			return hx_zero_43
		}
		return hx_value_42.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_44 any) map[string]any {
		if hx_value_44 == nil {
			var hx_zero_45 map[string]any
			return hx_zero_45
		}
		return hx_value_44.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_46 map[string]any) func() bool {
		hx_field_47 := hx_obj_46["hasNext"]
		if hx_field_47 == nil {
			var hx_zero_48 func() bool
			return hx_zero_48
		}
		return hx_field_47.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_49 map[string]any) func() *string {
			hx_field_50 := hx_obj_49["next"]
			if hx_field_50 == nil {
				var hx_zero_51 func() *string
				return hx_zero_51
			}
			return hx_field_50.(func() *string)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_52 any) map[string]any {
		if hx_value_52 == nil {
			var hx_zero_53 map[string]any
			return hx_zero_53
		}
		return hx_value_52.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_54 map[string]any) func() bool {
		hx_field_55 := hx_obj_54["hasNext"]
		if hx_field_55 == nil {
			var hx_zero_56 func() bool
			return hx_zero_56
		}
		return hx_field_55.(func() bool)
	}(iterator)() {
		key := func(hx_obj_57 map[string]any) func() *string {
			hx_field_58 := hx_obj_57["next"]
			if hx_field_58 == nil {
				var hx_zero_59 func() *string
				return hx_zero_59
			}
			return hx_field_58.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_60 map[string]any) func() bool {
			hx_field_61 := hx_obj_60["hasNext"]
			if hx_field_61 == nil {
				var hx_zero_62 func() bool
				return hx_zero_62
			}
			return hx_field_61.(func() bool)
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
