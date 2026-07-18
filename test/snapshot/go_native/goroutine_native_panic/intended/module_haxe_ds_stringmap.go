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
	hx_obj_19 := map[string]any{}
	hx_obj_19["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_19["next"] = func() *string {
		hx_post_20 := index
		index = int(int32((index + 1)))
		return keys[hx_post_20]
	}
	return hx_obj_19
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_21 := map[string]any{}
	hx_obj_21["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_21["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_22 := index
			index = int(int32((index + 1)))
			return hx_post_22
		}()])
	}
	return hx_obj_21
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_23 any) map[string]any {
		if hx_value_23 == nil {
			var hx_zero_24 map[string]any
			return hx_zero_24
		}
		return hx_value_23.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_25 := map[string]any{}
	hx_obj_25["hasNext"] = func() bool {
		return func(hx_obj_26 map[string]any) func() bool {
			hx_field_27 := hx_obj_26["hasNext"]
			if hx_field_27 == nil {
				var hx_zero_28 func() bool
				return hx_zero_28
			}
			return hx_field_27.(func() bool)
		}(keys)()
	}
	hx_obj_25["next"] = func() map[string]any {
		key := func(hx_obj_29 map[string]any) func() *string {
			hx_field_30 := hx_obj_29["next"]
			if hx_field_30 == nil {
				var hx_zero_31 func() *string
				return hx_zero_31
			}
			return hx_field_30.(func() *string)
		}(keys)()
		hx_obj_32 := map[string]any{}
		hx_obj_32["key"] = key
		hx_obj_32["value"] = _gthis.__hx_this.get(key)
		return hx_obj_32
	}
	return hx_obj_25
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_33 any) *string {
		if hx_value_33 == nil {
			var hx_zero_34 *string
			return hx_zero_34
		}
		return hx_value_33.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_35 any) *string {
		if hx_value_35 == nil {
			var hx_zero_36 *string
			return hx_zero_36
		}
		return hx_value_35.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_39 any) bool {
		if hx_value_39 == nil {
			var hx_zero_40 bool
			return hx_zero_40
		}
		return hx_value_39.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_37 any) *string {
		if hx_value_37 == nil {
			var hx_zero_38 *string
			return hx_zero_38
		}
		return hx_value_37.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_43 any) bool {
		if hx_value_43 == nil {
			var hx_zero_44 bool
			return hx_zero_44
		}
		return hx_value_43.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_41 any) *string {
		if hx_value_41 == nil {
			var hx_zero_42 *string
			return hx_zero_42
		}
		return hx_value_41.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_45 any) *haxe__ds__StringMap {
		if hx_value_45 == nil {
			var hx_zero_46 *haxe__ds__StringMap
			return hx_zero_46
		}
		return hx_value_45.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_47 any) map[string]any {
		if hx_value_47 == nil {
			var hx_zero_48 map[string]any
			return hx_zero_48
		}
		return hx_value_47.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_49 map[string]any) func() bool {
		hx_field_50 := hx_obj_49["hasNext"]
		if hx_field_50 == nil {
			var hx_zero_51 func() bool
			return hx_zero_51
		}
		return hx_field_50.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_52 map[string]any) func() *string {
			hx_field_53 := hx_obj_52["next"]
			if hx_field_53 == nil {
				var hx_zero_54 func() *string
				return hx_zero_54
			}
			return hx_field_53.(func() *string)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_55 any) map[string]any {
		if hx_value_55 == nil {
			var hx_zero_56 map[string]any
			return hx_zero_56
		}
		return hx_value_55.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_57 map[string]any) func() bool {
		hx_field_58 := hx_obj_57["hasNext"]
		if hx_field_58 == nil {
			var hx_zero_59 func() bool
			return hx_zero_59
		}
		return hx_field_58.(func() bool)
	}(iterator)() {
		key := func(hx_obj_60 map[string]any) func() *string {
			hx_field_61 := hx_obj_60["next"]
			if hx_field_61 == nil {
				var hx_zero_62 func() *string
				return hx_zero_62
			}
			return hx_field_61.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_63 map[string]any) func() bool {
			hx_field_64 := hx_obj_63["hasNext"]
			if hx_field_64 == nil {
				var hx_zero_65 func() bool
				return hx_zero_65
			}
			return hx_field_64.(func() bool)
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
