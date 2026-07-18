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
	hx_obj_72 := map[string]any{}
	hx_obj_72["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_72["next"] = func() *string {
		hx_post_73 := index
		index = int(int32((index + 1)))
		return keys[hx_post_73]
	}
	return hx_obj_72
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_74 := map[string]any{}
	hx_obj_74["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_74["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_75 := index
			index = int(int32((index + 1)))
			return hx_post_75
		}()])
	}
	return hx_obj_74
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_76 any) map[string]any {
		if hx_value_76 == nil {
			var hx_zero_77 map[string]any
			return hx_zero_77
		}
		return hx_value_76.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_78 := map[string]any{}
	hx_obj_78["hasNext"] = func() bool {
		return func(hx_obj_79 map[string]any) func() bool {
			hx_field_80 := hx_obj_79["hasNext"]
			if hx_field_80 == nil {
				var hx_zero_81 func() bool
				return hx_zero_81
			}
			return hx_field_80.(func() bool)
		}(keys)()
	}
	hx_obj_78["next"] = func() map[string]any {
		key := func(hx_obj_82 map[string]any) func() *string {
			hx_field_83 := hx_obj_82["next"]
			if hx_field_83 == nil {
				var hx_zero_84 func() *string
				return hx_zero_84
			}
			return hx_field_83.(func() *string)
		}(keys)()
		hx_obj_85 := map[string]any{}
		hx_obj_85["key"] = key
		hx_obj_85["value"] = _gthis.__hx_this.get(key)
		return hx_obj_85
	}
	return hx_obj_78
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_86 any) *string {
		if hx_value_86 == nil {
			var hx_zero_87 *string
			return hx_zero_87
		}
		return hx_value_86.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_88 any) *string {
		if hx_value_88 == nil {
			var hx_zero_89 *string
			return hx_zero_89
		}
		return hx_value_88.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_92 any) bool {
		if hx_value_92 == nil {
			var hx_zero_93 bool
			return hx_zero_93
		}
		return hx_value_92.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_90 any) *string {
		if hx_value_90 == nil {
			var hx_zero_91 *string
			return hx_zero_91
		}
		return hx_value_90.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_96 any) bool {
		if hx_value_96 == nil {
			var hx_zero_97 bool
			return hx_zero_97
		}
		return hx_value_96.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_94 any) *string {
		if hx_value_94 == nil {
			var hx_zero_95 *string
			return hx_zero_95
		}
		return hx_value_94.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_98 any) *haxe__ds__StringMap {
		if hx_value_98 == nil {
			var hx_zero_99 *haxe__ds__StringMap
			return hx_zero_99
		}
		return hx_value_98.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_100 any) map[string]any {
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
	}(key)() {
		key_1 := func(hx_obj_105 map[string]any) func() *string {
			hx_field_106 := hx_obj_105["next"]
			if hx_field_106 == nil {
				var hx_zero_107 func() *string
				return hx_zero_107
			}
			return hx_field_106.(func() *string)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_108 any) map[string]any {
		if hx_value_108 == nil {
			var hx_zero_109 map[string]any
			return hx_zero_109
		}
		return hx_value_108.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_110 map[string]any) func() bool {
		hx_field_111 := hx_obj_110["hasNext"]
		if hx_field_111 == nil {
			var hx_zero_112 func() bool
			return hx_zero_112
		}
		return hx_field_111.(func() bool)
	}(iterator)() {
		key := func(hx_obj_113 map[string]any) func() *string {
			hx_field_114 := hx_obj_113["next"]
			if hx_field_114 == nil {
				var hx_zero_115 func() *string
				return hx_zero_115
			}
			return hx_field_114.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_116 map[string]any) func() bool {
			hx_field_117 := hx_obj_116["hasNext"]
			if hx_field_117 == nil {
				var hx_zero_118 func() bool
				return hx_zero_118
			}
			return hx_field_117.(func() bool)
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
