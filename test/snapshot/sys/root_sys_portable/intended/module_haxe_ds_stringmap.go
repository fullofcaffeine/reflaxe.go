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
	hx_obj_28 := map[string]any{}
	hx_obj_28["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_28["next"] = func() *string {
		hx_post_29 := index
		index = int(int32((index + 1)))
		return keys[hx_post_29]
	}
	return hx_obj_28
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_30 := map[string]any{}
	hx_obj_30["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_30["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_31 := index
			index = int(int32((index + 1)))
			return hx_post_31
		}()])
	}
	return hx_obj_30
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_32 any) map[string]any {
		if hx_value_32 == nil {
			var hx_zero_33 map[string]any
			return hx_zero_33
		}
		return hx_value_32.(map[string]any)
	}(self.keys())
	hx_obj_34 := map[string]any{}
	hx_obj_34["hasNext"] = func() bool {
		return func(hx_obj_35 map[string]any) func() bool {
			hx_field_36 := hx_obj_35["hasNext"]
			if hx_field_36 == nil {
				var hx_zero_37 func() bool
				return hx_zero_37
			}
			return hx_field_36.(func() bool)
		}(keys)()
	}
	hx_obj_34["next"] = func() map[string]any {
		key := func(hx_obj_38 map[string]any) func() *string {
			hx_field_39 := hx_obj_38["next"]
			if hx_field_39 == nil {
				var hx_zero_40 func() *string
				return hx_zero_40
			}
			return hx_field_39.(func() *string)
		}(keys)()
		hx_obj_41 := map[string]any{}
		hx_obj_41["key"] = key
		hx_obj_41["value"] = _gthis.get(key)
		return hx_obj_41
	}
	return hx_obj_34
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_42 any) *string {
		if hx_value_42 == nil {
			var hx_zero_43 *string
			return hx_zero_43
		}
		return hx_value_42.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_44 any) *string {
		if hx_value_44 == nil {
			var hx_zero_45 *string
			return hx_zero_45
		}
		return hx_value_44.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_48 any) bool {
		if hx_value_48 == nil {
			var hx_zero_49 bool
			return hx_zero_49
		}
		return hx_value_48.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_46 any) *string {
		if hx_value_46 == nil {
			var hx_zero_47 *string
			return hx_zero_47
		}
		return hx_value_46.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_52 any) bool {
		if hx_value_52 == nil {
			var hx_zero_53 bool
			return hx_zero_53
		}
		return hx_value_52.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_50 any) *string {
		if hx_value_50 == nil {
			var hx_zero_51 *string
			return hx_zero_51
		}
		return hx_value_50.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_54 any) *haxe__ds__StringMap {
		if hx_value_54 == nil {
			var hx_zero_55 *haxe__ds__StringMap
			return hx_zero_55
		}
		return hx_value_54.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_56 any) map[string]any {
		if hx_value_56 == nil {
			var hx_zero_57 map[string]any
			return hx_zero_57
		}
		return hx_value_56.(map[string]any)
	}(self.keys())
	for func(hx_obj_58 map[string]any) func() bool {
		hx_field_59 := hx_obj_58["hasNext"]
		if hx_field_59 == nil {
			var hx_zero_60 func() bool
			return hx_zero_60
		}
		return hx_field_59.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_61 map[string]any) func() *string {
			hx_field_62 := hx_obj_61["next"]
			if hx_field_62 == nil {
				var hx_zero_63 func() *string
				return hx_zero_63
			}
			return hx_field_62.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_64 any) map[string]any {
		if hx_value_64 == nil {
			var hx_zero_65 map[string]any
			return hx_zero_65
		}
		return hx_value_64.(map[string]any)
	}(self.keys())
	for func(hx_obj_66 map[string]any) func() bool {
		hx_field_67 := hx_obj_66["hasNext"]
		if hx_field_67 == nil {
			var hx_zero_68 func() bool
			return hx_zero_68
		}
		return hx_field_67.(func() bool)
	}(iterator)() {
		key := func(hx_obj_69 map[string]any) func() *string {
			hx_field_70 := hx_obj_69["next"]
			if hx_field_70 == nil {
				var hx_zero_71 func() *string
				return hx_zero_71
			}
			return hx_field_70.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_72 map[string]any) func() bool {
			hx_field_73 := hx_obj_72["hasNext"]
			if hx_field_73 == nil {
				var hx_zero_74 func() bool
				return hx_zero_74
			}
			return hx_field_73.(func() bool)
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
