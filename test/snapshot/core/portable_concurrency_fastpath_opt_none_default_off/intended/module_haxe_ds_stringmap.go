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
	hx_obj_18 := map[string]any{}
	hx_obj_18["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_18["next"] = func() *string {
		hx_post_19 := index
		index = int(int32((index + 1)))
		return keys[hx_post_19]
	}
	return hx_obj_18
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_20 := map[string]any{}
	hx_obj_20["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_20["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_21 := index
			index = int(int32((index + 1)))
			return hx_post_21
		}()])
	}
	return hx_obj_20
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_22 any) map[string]any {
		if hx_value_22 == nil {
			var hx_zero_23 map[string]any
			return hx_zero_23
		}
		return hx_value_22.(map[string]any)
	}(self.keys())
	hx_obj_24 := map[string]any{}
	hx_obj_24["hasNext"] = func() bool {
		return func(hx_obj_25 map[string]any) func() bool {
			hx_field_26 := hx_obj_25["hasNext"]
			if hx_field_26 == nil {
				var hx_zero_27 func() bool
				return hx_zero_27
			}
			return hx_field_26.(func() bool)
		}(keys)()
	}
	hx_obj_24["next"] = func() map[string]any {
		key := func(hx_obj_28 map[string]any) func() *string {
			hx_field_29 := hx_obj_28["next"]
			if hx_field_29 == nil {
				var hx_zero_30 func() *string
				return hx_zero_30
			}
			return hx_field_29.(func() *string)
		}(keys)()
		hx_obj_31 := map[string]any{}
		hx_obj_31["key"] = key
		hx_obj_31["value"] = _gthis.get(key)
		return hx_obj_31
	}
	return hx_obj_24
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_32 any) *string {
		if hx_value_32 == nil {
			var hx_zero_33 *string
			return hx_zero_33
		}
		return hx_value_32.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_34 any) *string {
		if hx_value_34 == nil {
			var hx_zero_35 *string
			return hx_zero_35
		}
		return hx_value_34.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_38 any) bool {
		if hx_value_38 == nil {
			var hx_zero_39 bool
			return hx_zero_39
		}
		return hx_value_38.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_36 any) *string {
		if hx_value_36 == nil {
			var hx_zero_37 *string
			return hx_zero_37
		}
		return hx_value_36.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_42 any) bool {
		if hx_value_42 == nil {
			var hx_zero_43 bool
			return hx_zero_43
		}
		return hx_value_42.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_40 any) *string {
		if hx_value_40 == nil {
			var hx_zero_41 *string
			return hx_zero_41
		}
		return hx_value_40.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_44 any) *haxe__ds__StringMap {
		if hx_value_44 == nil {
			var hx_zero_45 *haxe__ds__StringMap
			return hx_zero_45
		}
		return hx_value_44.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_46 any) map[string]any {
		if hx_value_46 == nil {
			var hx_zero_47 map[string]any
			return hx_zero_47
		}
		return hx_value_46.(map[string]any)
	}(self.keys())
	for func(hx_obj_48 map[string]any) func() bool {
		hx_field_49 := hx_obj_48["hasNext"]
		if hx_field_49 == nil {
			var hx_zero_50 func() bool
			return hx_zero_50
		}
		return hx_field_49.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_51 map[string]any) func() *string {
			hx_field_52 := hx_obj_51["next"]
			if hx_field_52 == nil {
				var hx_zero_53 func() *string
				return hx_zero_53
			}
			return hx_field_52.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_54 any) map[string]any {
		if hx_value_54 == nil {
			var hx_zero_55 map[string]any
			return hx_zero_55
		}
		return hx_value_54.(map[string]any)
	}(self.keys())
	for func(hx_obj_56 map[string]any) func() bool {
		hx_field_57 := hx_obj_56["hasNext"]
		if hx_field_57 == nil {
			var hx_zero_58 func() bool
			return hx_zero_58
		}
		return hx_field_57.(func() bool)
	}(iterator)() {
		key := func(hx_obj_59 map[string]any) func() *string {
			hx_field_60 := hx_obj_59["next"]
			if hx_field_60 == nil {
				var hx_zero_61 func() *string
				return hx_zero_61
			}
			return hx_field_60.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_62 map[string]any) func() bool {
			hx_field_63 := hx_obj_62["hasNext"]
			if hx_field_63 == nil {
				var hx_zero_64 func() bool
				return hx_zero_64
			}
			return hx_field_63.(func() bool)
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
