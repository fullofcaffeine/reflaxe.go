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
	hx_obj_43 := map[string]any{}
	hx_obj_43["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_43["next"] = func() *string {
		hx_post_44 := index
		index = int(int32((index + 1)))
		return keys[hx_post_44]
	}
	return hx_obj_43
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_45 := map[string]any{}
	hx_obj_45["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_45["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_46 := index
			index = int(int32((index + 1)))
			return hx_post_46
		}()])
	}
	return hx_obj_45
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_47 any) map[string]any {
		if hx_value_47 == nil {
			var hx_zero_48 map[string]any
			return hx_zero_48
		}
		return hx_value_47.(map[string]any)
	}(self.keys())
	hx_obj_49 := map[string]any{}
	hx_obj_49["hasNext"] = func() bool {
		return func(hx_obj_50 map[string]any) func() bool {
			hx_field_51 := hx_obj_50["hasNext"]
			if hx_field_51 == nil {
				var hx_zero_52 func() bool
				return hx_zero_52
			}
			return hx_field_51.(func() bool)
		}(keys)()
	}
	hx_obj_49["next"] = func() map[string]any {
		key := func(hx_obj_53 map[string]any) func() *string {
			hx_field_54 := hx_obj_53["next"]
			if hx_field_54 == nil {
				var hx_zero_55 func() *string
				return hx_zero_55
			}
			return hx_field_54.(func() *string)
		}(keys)()
		hx_obj_56 := map[string]any{}
		hx_obj_56["key"] = key
		hx_obj_56["value"] = _gthis.get(key)
		return hx_obj_56
	}
	return hx_obj_49
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_57 any) *string {
		if hx_value_57 == nil {
			var hx_zero_58 *string
			return hx_zero_58
		}
		return hx_value_57.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_59 any) *string {
		if hx_value_59 == nil {
			var hx_zero_60 *string
			return hx_zero_60
		}
		return hx_value_59.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_63 any) bool {
		if hx_value_63 == nil {
			var hx_zero_64 bool
			return hx_zero_64
		}
		return hx_value_63.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_61 any) *string {
		if hx_value_61 == nil {
			var hx_zero_62 *string
			return hx_zero_62
		}
		return hx_value_61.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_67 any) bool {
		if hx_value_67 == nil {
			var hx_zero_68 bool
			return hx_zero_68
		}
		return hx_value_67.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_65 any) *string {
		if hx_value_65 == nil {
			var hx_zero_66 *string
			return hx_zero_66
		}
		return hx_value_65.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_69 any) *haxe__ds__StringMap {
		if hx_value_69 == nil {
			var hx_zero_70 *haxe__ds__StringMap
			return hx_zero_70
		}
		return hx_value_69.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_71 any) map[string]any {
		if hx_value_71 == nil {
			var hx_zero_72 map[string]any
			return hx_zero_72
		}
		return hx_value_71.(map[string]any)
	}(self.keys())
	for func(hx_obj_73 map[string]any) func() bool {
		hx_field_74 := hx_obj_73["hasNext"]
		if hx_field_74 == nil {
			var hx_zero_75 func() bool
			return hx_zero_75
		}
		return hx_field_74.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_76 map[string]any) func() *string {
			hx_field_77 := hx_obj_76["next"]
			if hx_field_77 == nil {
				var hx_zero_78 func() *string
				return hx_zero_78
			}
			return hx_field_77.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_79 any) map[string]any {
		if hx_value_79 == nil {
			var hx_zero_80 map[string]any
			return hx_zero_80
		}
		return hx_value_79.(map[string]any)
	}(self.keys())
	for func(hx_obj_81 map[string]any) func() bool {
		hx_field_82 := hx_obj_81["hasNext"]
		if hx_field_82 == nil {
			var hx_zero_83 func() bool
			return hx_zero_83
		}
		return hx_field_82.(func() bool)
	}(iterator)() {
		key := func(hx_obj_84 map[string]any) func() *string {
			hx_field_85 := hx_obj_84["next"]
			if hx_field_85 == nil {
				var hx_zero_86 func() *string
				return hx_zero_86
			}
			return hx_field_85.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_87 map[string]any) func() bool {
			hx_field_88 := hx_obj_87["hasNext"]
			if hx_field_88 == nil {
				var hx_zero_89 func() bool
				return hx_zero_89
			}
			return hx_field_88.(func() bool)
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
