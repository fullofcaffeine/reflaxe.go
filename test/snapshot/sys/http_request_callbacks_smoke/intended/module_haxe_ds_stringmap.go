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
	hx_obj_55 := map[string]any{}
	hx_obj_55["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_55["next"] = func() *string {
		hx_post_56 := index
		index = int(int32((index + 1)))
		return keys[hx_post_56]
	}
	return hx_obj_55
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_57 := map[string]any{}
	hx_obj_57["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_57["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_58 := index
			index = int(int32((index + 1)))
			return hx_post_58
		}()])
	}
	return hx_obj_57
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_59 any) map[string]any {
		if hx_value_59 == nil {
			var hx_zero_60 map[string]any
			return hx_zero_60
		}
		return hx_value_59.(map[string]any)
	}(self.keys())
	hx_obj_61 := map[string]any{}
	hx_obj_61["hasNext"] = func() bool {
		return func(hx_obj_62 map[string]any) func() bool {
			hx_field_63 := hx_obj_62["hasNext"]
			if hx_field_63 == nil {
				var hx_zero_64 func() bool
				return hx_zero_64
			}
			return hx_field_63.(func() bool)
		}(keys)()
	}
	hx_obj_61["next"] = func() map[string]any {
		key := func(hx_obj_65 map[string]any) func() *string {
			hx_field_66 := hx_obj_65["next"]
			if hx_field_66 == nil {
				var hx_zero_67 func() *string
				return hx_zero_67
			}
			return hx_field_66.(func() *string)
		}(keys)()
		hx_obj_68 := map[string]any{}
		hx_obj_68["key"] = key
		hx_obj_68["value"] = _gthis.get(key)
		return hx_obj_68
	}
	return hx_obj_61
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_69 any) *string {
		if hx_value_69 == nil {
			var hx_zero_70 *string
			return hx_zero_70
		}
		return hx_value_69.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_71 any) *string {
		if hx_value_71 == nil {
			var hx_zero_72 *string
			return hx_zero_72
		}
		return hx_value_71.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_75 any) bool {
		if hx_value_75 == nil {
			var hx_zero_76 bool
			return hx_zero_76
		}
		return hx_value_75.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_73 any) *string {
		if hx_value_73 == nil {
			var hx_zero_74 *string
			return hx_zero_74
		}
		return hx_value_73.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_79 any) bool {
		if hx_value_79 == nil {
			var hx_zero_80 bool
			return hx_zero_80
		}
		return hx_value_79.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_77 any) *string {
		if hx_value_77 == nil {
			var hx_zero_78 *string
			return hx_zero_78
		}
		return hx_value_77.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_81 any) *haxe__ds__StringMap {
		if hx_value_81 == nil {
			var hx_zero_82 *haxe__ds__StringMap
			return hx_zero_82
		}
		return hx_value_81.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_83 any) map[string]any {
		if hx_value_83 == nil {
			var hx_zero_84 map[string]any
			return hx_zero_84
		}
		return hx_value_83.(map[string]any)
	}(self.keys())
	for func(hx_obj_85 map[string]any) func() bool {
		hx_field_86 := hx_obj_85["hasNext"]
		if hx_field_86 == nil {
			var hx_zero_87 func() bool
			return hx_zero_87
		}
		return hx_field_86.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_88 map[string]any) func() *string {
			hx_field_89 := hx_obj_88["next"]
			if hx_field_89 == nil {
				var hx_zero_90 func() *string
				return hx_zero_90
			}
			return hx_field_89.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_91 any) map[string]any {
		if hx_value_91 == nil {
			var hx_zero_92 map[string]any
			return hx_zero_92
		}
		return hx_value_91.(map[string]any)
	}(self.keys())
	for func(hx_obj_93 map[string]any) func() bool {
		hx_field_94 := hx_obj_93["hasNext"]
		if hx_field_94 == nil {
			var hx_zero_95 func() bool
			return hx_zero_95
		}
		return hx_field_94.(func() bool)
	}(iterator)() {
		key := func(hx_obj_96 map[string]any) func() *string {
			hx_field_97 := hx_obj_96["next"]
			if hx_field_97 == nil {
				var hx_zero_98 func() *string
				return hx_zero_98
			}
			return hx_field_97.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_99 map[string]any) func() bool {
			hx_field_100 := hx_obj_99["hasNext"]
			if hx_field_100 == nil {
				var hx_zero_101 func() bool
				return hx_zero_101
			}
			return hx_field_100.(func() bool)
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
