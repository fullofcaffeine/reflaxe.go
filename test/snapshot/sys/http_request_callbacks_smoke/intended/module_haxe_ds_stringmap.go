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
	hx_obj_65 := map[string]any{}
	hx_obj_65["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_65["next"] = func() *string {
		hx_post_66 := index
		index = int(int32((index + 1)))
		return keys[hx_post_66]
	}
	return hx_obj_65
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_67 := map[string]any{}
	hx_obj_67["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_67["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_68 := index
			index = int(int32((index + 1)))
			return hx_post_68
		}()])
	}
	return hx_obj_67
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_69 any) map[string]any {
		if hx_value_69 == nil {
			var hx_zero_70 map[string]any
			return hx_zero_70
		}
		return hx_value_69.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_71 := map[string]any{}
	hx_obj_71["hasNext"] = func() bool {
		return func(hx_obj_72 map[string]any) func() bool {
			hx_field_73 := hx_obj_72["hasNext"]
			if hx_field_73 == nil {
				var hx_zero_74 func() bool
				return hx_zero_74
			}
			return hx_field_73.(func() bool)
		}(keys)()
	}
	hx_obj_71["next"] = func() map[string]any {
		key := func(hx_obj_75 map[string]any) func() *string {
			hx_field_76 := hx_obj_75["next"]
			if hx_field_76 == nil {
				var hx_zero_77 func() *string
				return hx_zero_77
			}
			return hx_field_76.(func() *string)
		}(keys)()
		hx_obj_78 := map[string]any{}
		hx_obj_78["key"] = key
		hx_obj_78["value"] = _gthis.__hx_this.get(key)
		return hx_obj_78
	}
	return hx_obj_71
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_79 any) *string {
		if hx_value_79 == nil {
			var hx_zero_80 *string
			return hx_zero_80
		}
		return hx_value_79.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_81 any) *string {
		if hx_value_81 == nil {
			var hx_zero_82 *string
			return hx_zero_82
		}
		return hx_value_81.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_85 any) bool {
		if hx_value_85 == nil {
			var hx_zero_86 bool
			return hx_zero_86
		}
		return hx_value_85.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_83 any) *string {
		if hx_value_83 == nil {
			var hx_zero_84 *string
			return hx_zero_84
		}
		return hx_value_83.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_89 any) bool {
		if hx_value_89 == nil {
			var hx_zero_90 bool
			return hx_zero_90
		}
		return hx_value_89.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_87 any) *string {
		if hx_value_87 == nil {
			var hx_zero_88 *string
			return hx_zero_88
		}
		return hx_value_87.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_91 any) *haxe__ds__StringMap {
		if hx_value_91 == nil {
			var hx_zero_92 *haxe__ds__StringMap
			return hx_zero_92
		}
		return hx_value_91.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_93 any) map[string]any {
		if hx_value_93 == nil {
			var hx_zero_94 map[string]any
			return hx_zero_94
		}
		return hx_value_93.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_95 map[string]any) func() bool {
		hx_field_96 := hx_obj_95["hasNext"]
		if hx_field_96 == nil {
			var hx_zero_97 func() bool
			return hx_zero_97
		}
		return hx_field_96.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_98 map[string]any) func() *string {
			hx_field_99 := hx_obj_98["next"]
			if hx_field_99 == nil {
				var hx_zero_100 func() *string
				return hx_zero_100
			}
			return hx_field_99.(func() *string)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_101 any) map[string]any {
		if hx_value_101 == nil {
			var hx_zero_102 map[string]any
			return hx_zero_102
		}
		return hx_value_101.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_103 map[string]any) func() bool {
		hx_field_104 := hx_obj_103["hasNext"]
		if hx_field_104 == nil {
			var hx_zero_105 func() bool
			return hx_zero_105
		}
		return hx_field_104.(func() bool)
	}(iterator)() {
		key := func(hx_obj_106 map[string]any) func() *string {
			hx_field_107 := hx_obj_106["next"]
			if hx_field_107 == nil {
				var hx_zero_108 func() *string
				return hx_zero_108
			}
			return hx_field_107.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_109 map[string]any) func() bool {
			hx_field_110 := hx_obj_109["hasNext"]
			if hx_field_110 == nil {
				var hx_zero_111 func() bool
				return hx_zero_111
			}
			return hx_field_110.(func() bool)
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
