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
	hx_obj_21 := map[string]any{}
	hx_obj_21["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_21["next"] = func() *string {
		hx_post_22 := index
		index = int(int32((index + 1)))
		return keys[hx_post_22]
	}
	return hx_obj_21
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_23 := map[string]any{}
	hx_obj_23["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_23["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_24 := index
			index = int(int32((index + 1)))
			return hx_post_24
		}()])
	}
	return hx_obj_23
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_25 any) map[string]any {
		if hx_value_25 == nil {
			var hx_zero_26 map[string]any
			return hx_zero_26
		}
		return hx_value_25.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_27 := map[string]any{}
	hx_obj_27["hasNext"] = func() bool {
		return func(hx_obj_28 map[string]any) func() bool {
			hx_field_29 := hx_obj_28["hasNext"]
			if hx_field_29 == nil {
				var hx_zero_30 func() bool
				return hx_zero_30
			}
			return hx_field_29.(func() bool)
		}(keys)()
	}
	hx_obj_27["next"] = func() map[string]any {
		key := func(hx_obj_31 map[string]any) func() *string {
			hx_field_32 := hx_obj_31["next"]
			if hx_field_32 == nil {
				var hx_zero_33 func() *string
				return hx_zero_33
			}
			return hx_field_32.(func() *string)
		}(keys)()
		hx_obj_34 := map[string]any{}
		hx_obj_34["key"] = key
		hx_obj_34["value"] = _gthis.__hx_this.get(key)
		return hx_obj_34
	}
	return hx_obj_27
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_35 any) *string {
		if hx_value_35 == nil {
			var hx_zero_36 *string
			return hx_zero_36
		}
		return hx_value_35.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_37 any) *string {
		if hx_value_37 == nil {
			var hx_zero_38 *string
			return hx_zero_38
		}
		return hx_value_37.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_41 any) bool {
		if hx_value_41 == nil {
			var hx_zero_42 bool
			return hx_zero_42
		}
		return hx_value_41.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_39 any) *string {
		if hx_value_39 == nil {
			var hx_zero_40 *string
			return hx_zero_40
		}
		return hx_value_39.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_45 any) bool {
		if hx_value_45 == nil {
			var hx_zero_46 bool
			return hx_zero_46
		}
		return hx_value_45.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_43 any) *string {
		if hx_value_43 == nil {
			var hx_zero_44 *string
			return hx_zero_44
		}
		return hx_value_43.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_47 any) *haxe__ds__StringMap {
		if hx_value_47 == nil {
			var hx_zero_48 *haxe__ds__StringMap
			return hx_zero_48
		}
		return hx_value_47.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_49 any) map[string]any {
		if hx_value_49 == nil {
			var hx_zero_50 map[string]any
			return hx_zero_50
		}
		return hx_value_49.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_51 map[string]any) func() bool {
		hx_field_52 := hx_obj_51["hasNext"]
		if hx_field_52 == nil {
			var hx_zero_53 func() bool
			return hx_zero_53
		}
		return hx_field_52.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_54 map[string]any) func() *string {
			hx_field_55 := hx_obj_54["next"]
			if hx_field_55 == nil {
				var hx_zero_56 func() *string
				return hx_zero_56
			}
			return hx_field_55.(func() *string)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_57 any) map[string]any {
		if hx_value_57 == nil {
			var hx_zero_58 map[string]any
			return hx_zero_58
		}
		return hx_value_57.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_59 map[string]any) func() bool {
		hx_field_60 := hx_obj_59["hasNext"]
		if hx_field_60 == nil {
			var hx_zero_61 func() bool
			return hx_zero_61
		}
		return hx_field_60.(func() bool)
	}(iterator)() {
		key := func(hx_obj_62 map[string]any) func() *string {
			hx_field_63 := hx_obj_62["next"]
			if hx_field_63 == nil {
				var hx_zero_64 func() *string
				return hx_zero_64
			}
			return hx_field_63.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_65 map[string]any) func() bool {
			hx_field_66 := hx_obj_65["hasNext"]
			if hx_field_66 == nil {
				var hx_zero_67 func() bool
				return hx_zero_67
			}
			return hx_field_66.(func() bool)
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
