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
	hx_obj_22 := map[string]any{}
	hx_obj_22["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_22["next"] = func() *string {
		hx_post_23 := index
		index = int(int32((index + 1)))
		return keys[hx_post_23]
	}
	return hx_obj_22
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_24 := map[string]any{}
	hx_obj_24["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_24["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_25 := index
			index = int(int32((index + 1)))
			return hx_post_25
		}()])
	}
	return hx_obj_24
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_26 any) map[string]any {
		if hx_value_26 == nil {
			var hx_zero_27 map[string]any
			return hx_zero_27
		}
		return hx_value_26.(map[string]any)
	}(self.keys())
	hx_obj_28 := map[string]any{}
	hx_obj_28["hasNext"] = func() bool {
		return func(hx_obj_29 map[string]any) func() bool {
			hx_field_30 := hx_obj_29["hasNext"]
			if hx_field_30 == nil {
				var hx_zero_31 func() bool
				return hx_zero_31
			}
			return hx_field_30.(func() bool)
		}(keys)()
	}
	hx_obj_28["next"] = func() map[string]any {
		key := func(hx_obj_32 map[string]any) func() *string {
			hx_field_33 := hx_obj_32["next"]
			if hx_field_33 == nil {
				var hx_zero_34 func() *string
				return hx_zero_34
			}
			return hx_field_33.(func() *string)
		}(keys)()
		hx_obj_35 := map[string]any{}
		hx_obj_35["key"] = key
		hx_obj_35["value"] = _gthis.get(key)
		return hx_obj_35
	}
	return hx_obj_28
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_36 any) *string {
		if hx_value_36 == nil {
			var hx_zero_37 *string
			return hx_zero_37
		}
		return hx_value_36.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_38 any) *string {
		if hx_value_38 == nil {
			var hx_zero_39 *string
			return hx_zero_39
		}
		return hx_value_38.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_42 any) bool {
		if hx_value_42 == nil {
			var hx_zero_43 bool
			return hx_zero_43
		}
		return hx_value_42.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_40 any) *string {
		if hx_value_40 == nil {
			var hx_zero_41 *string
			return hx_zero_41
		}
		return hx_value_40.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_46 any) bool {
		if hx_value_46 == nil {
			var hx_zero_47 bool
			return hx_zero_47
		}
		return hx_value_46.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_44 any) *string {
		if hx_value_44 == nil {
			var hx_zero_45 *string
			return hx_zero_45
		}
		return hx_value_44.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_48 any) *haxe__ds__StringMap {
		if hx_value_48 == nil {
			var hx_zero_49 *haxe__ds__StringMap
			return hx_zero_49
		}
		return hx_value_48.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_50 any) map[string]any {
		if hx_value_50 == nil {
			var hx_zero_51 map[string]any
			return hx_zero_51
		}
		return hx_value_50.(map[string]any)
	}(self.keys())
	for func(hx_obj_52 map[string]any) func() bool {
		hx_field_53 := hx_obj_52["hasNext"]
		if hx_field_53 == nil {
			var hx_zero_54 func() bool
			return hx_zero_54
		}
		return hx_field_53.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_55 map[string]any) func() *string {
			hx_field_56 := hx_obj_55["next"]
			if hx_field_56 == nil {
				var hx_zero_57 func() *string
				return hx_zero_57
			}
			return hx_field_56.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_58 any) map[string]any {
		if hx_value_58 == nil {
			var hx_zero_59 map[string]any
			return hx_zero_59
		}
		return hx_value_58.(map[string]any)
	}(self.keys())
	for func(hx_obj_60 map[string]any) func() bool {
		hx_field_61 := hx_obj_60["hasNext"]
		if hx_field_61 == nil {
			var hx_zero_62 func() bool
			return hx_zero_62
		}
		return hx_field_61.(func() bool)
	}(iterator)() {
		key := func(hx_obj_63 map[string]any) func() *string {
			hx_field_64 := hx_obj_63["next"]
			if hx_field_64 == nil {
				var hx_zero_65 func() *string
				return hx_zero_65
			}
			return hx_field_64.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_66 map[string]any) func() bool {
			hx_field_67 := hx_obj_66["hasNext"]
			if hx_field_67 == nil {
				var hx_zero_68 func() bool
				return hx_zero_68
			}
			return hx_field_67.(func() bool)
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
