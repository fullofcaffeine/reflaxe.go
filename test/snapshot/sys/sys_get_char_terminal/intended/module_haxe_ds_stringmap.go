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
	hx_obj_64 := map[string]any{}
	hx_obj_64["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_64["next"] = func() *string {
		hx_post_65 := index
		index = int(int32((index + 1)))
		return keys[hx_post_65]
	}
	return hx_obj_64
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_66 := map[string]any{}
	hx_obj_66["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_66["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_67 := index
			index = int(int32((index + 1)))
			return hx_post_67
		}()])
	}
	return hx_obj_66
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_68 any) map[string]any {
		if hx_value_68 == nil {
			var hx_zero_69 map[string]any
			return hx_zero_69
		}
		return hx_value_68.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_70 := map[string]any{}
	hx_obj_70["hasNext"] = func() bool {
		return func(hx_obj_71 map[string]any) func() bool {
			hx_field_72 := hx_obj_71["hasNext"]
			if hx_field_72 == nil {
				var hx_zero_73 func() bool
				return hx_zero_73
			}
			return hx_field_72.(func() bool)
		}(keys)()
	}
	hx_obj_70["next"] = func() map[string]any {
		key := func(hx_obj_74 map[string]any) func() *string {
			hx_field_75 := hx_obj_74["next"]
			if hx_field_75 == nil {
				var hx_zero_76 func() *string
				return hx_zero_76
			}
			return hx_field_75.(func() *string)
		}(keys)()
		hx_obj_77 := map[string]any{}
		hx_obj_77["key"] = key
		hx_obj_77["value"] = _gthis.__hx_this.get(key)
		return hx_obj_77
	}
	return hx_obj_70
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_78 any) *string {
		if hx_value_78 == nil {
			var hx_zero_79 *string
			return hx_zero_79
		}
		return hx_value_78.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_80 any) *string {
		if hx_value_80 == nil {
			var hx_zero_81 *string
			return hx_zero_81
		}
		return hx_value_80.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_84 any) bool {
		if hx_value_84 == nil {
			var hx_zero_85 bool
			return hx_zero_85
		}
		return hx_value_84.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_82 any) *string {
		if hx_value_82 == nil {
			var hx_zero_83 *string
			return hx_zero_83
		}
		return hx_value_82.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_88 any) bool {
		if hx_value_88 == nil {
			var hx_zero_89 bool
			return hx_zero_89
		}
		return hx_value_88.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_86 any) *string {
		if hx_value_86 == nil {
			var hx_zero_87 *string
			return hx_zero_87
		}
		return hx_value_86.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_90 any) *haxe__ds__StringMap {
		if hx_value_90 == nil {
			var hx_zero_91 *haxe__ds__StringMap
			return hx_zero_91
		}
		return hx_value_90.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_92 any) map[string]any {
		if hx_value_92 == nil {
			var hx_zero_93 map[string]any
			return hx_zero_93
		}
		return hx_value_92.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_94 map[string]any) func() bool {
		hx_field_95 := hx_obj_94["hasNext"]
		if hx_field_95 == nil {
			var hx_zero_96 func() bool
			return hx_zero_96
		}
		return hx_field_95.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_97 map[string]any) func() *string {
			hx_field_98 := hx_obj_97["next"]
			if hx_field_98 == nil {
				var hx_zero_99 func() *string
				return hx_zero_99
			}
			return hx_field_98.(func() *string)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_100 any) map[string]any {
		if hx_value_100 == nil {
			var hx_zero_101 map[string]any
			return hx_zero_101
		}
		return hx_value_100.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_102 map[string]any) func() bool {
		hx_field_103 := hx_obj_102["hasNext"]
		if hx_field_103 == nil {
			var hx_zero_104 func() bool
			return hx_zero_104
		}
		return hx_field_103.(func() bool)
	}(iterator)() {
		key := func(hx_obj_105 map[string]any) func() *string {
			hx_field_106 := hx_obj_105["next"]
			if hx_field_106 == nil {
				var hx_zero_107 func() *string
				return hx_zero_107
			}
			return hx_field_106.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_108 map[string]any) func() bool {
			hx_field_109 := hx_obj_108["hasNext"]
			if hx_field_109 == nil {
				var hx_zero_110 func() bool
				return hx_zero_110
			}
			return hx_field_109.(func() bool)
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
