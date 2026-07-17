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
	hx_obj_91 := map[string]any{}
	hx_obj_91["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_91["next"] = func() *string {
		hx_post_92 := index
		index = int(int32((index + 1)))
		return keys[hx_post_92]
	}
	return hx_obj_91
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_93 := map[string]any{}
	hx_obj_93["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_93["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_94 := index
			index = int(int32((index + 1)))
			return hx_post_94
		}()])
	}
	return hx_obj_93
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_95 any) map[string]any {
		if hx_value_95 == nil {
			var hx_zero_96 map[string]any
			return hx_zero_96
		}
		return hx_value_95.(map[string]any)
	}(self.keys())
	hx_obj_97 := map[string]any{}
	hx_obj_97["hasNext"] = func() bool {
		return func(hx_obj_98 map[string]any) func() bool {
			hx_field_99 := hx_obj_98["hasNext"]
			if hx_field_99 == nil {
				var hx_zero_100 func() bool
				return hx_zero_100
			}
			return hx_field_99.(func() bool)
		}(keys)()
	}
	hx_obj_97["next"] = func() map[string]any {
		key := func(hx_obj_101 map[string]any) func() *string {
			hx_field_102 := hx_obj_101["next"]
			if hx_field_102 == nil {
				var hx_zero_103 func() *string
				return hx_zero_103
			}
			return hx_field_102.(func() *string)
		}(keys)()
		hx_obj_104 := map[string]any{}
		hx_obj_104["key"] = key
		hx_obj_104["value"] = _gthis.get(key)
		return hx_obj_104
	}
	return hx_obj_97
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_105 any) *string {
		if hx_value_105 == nil {
			var hx_zero_106 *string
			return hx_zero_106
		}
		return hx_value_105.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_107 any) *string {
		if hx_value_107 == nil {
			var hx_zero_108 *string
			return hx_zero_108
		}
		return hx_value_107.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_111 any) bool {
		if hx_value_111 == nil {
			var hx_zero_112 bool
			return hx_zero_112
		}
		return hx_value_111.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_109 any) *string {
		if hx_value_109 == nil {
			var hx_zero_110 *string
			return hx_zero_110
		}
		return hx_value_109.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_115 any) bool {
		if hx_value_115 == nil {
			var hx_zero_116 bool
			return hx_zero_116
		}
		return hx_value_115.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_113 any) *string {
		if hx_value_113 == nil {
			var hx_zero_114 *string
			return hx_zero_114
		}
		return hx_value_113.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_117 any) *haxe__ds__StringMap {
		if hx_value_117 == nil {
			var hx_zero_118 *haxe__ds__StringMap
			return hx_zero_118
		}
		return hx_value_117.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_119 any) map[string]any {
		if hx_value_119 == nil {
			var hx_zero_120 map[string]any
			return hx_zero_120
		}
		return hx_value_119.(map[string]any)
	}(self.keys())
	for func(hx_obj_121 map[string]any) func() bool {
		hx_field_122 := hx_obj_121["hasNext"]
		if hx_field_122 == nil {
			var hx_zero_123 func() bool
			return hx_zero_123
		}
		return hx_field_122.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_124 map[string]any) func() *string {
			hx_field_125 := hx_obj_124["next"]
			if hx_field_125 == nil {
				var hx_zero_126 func() *string
				return hx_zero_126
			}
			return hx_field_125.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_127 any) map[string]any {
		if hx_value_127 == nil {
			var hx_zero_128 map[string]any
			return hx_zero_128
		}
		return hx_value_127.(map[string]any)
	}(self.keys())
	for func(hx_obj_129 map[string]any) func() bool {
		hx_field_130 := hx_obj_129["hasNext"]
		if hx_field_130 == nil {
			var hx_zero_131 func() bool
			return hx_zero_131
		}
		return hx_field_130.(func() bool)
	}(iterator)() {
		key := func(hx_obj_132 map[string]any) func() *string {
			hx_field_133 := hx_obj_132["next"]
			if hx_field_133 == nil {
				var hx_zero_134 func() *string
				return hx_zero_134
			}
			return hx_field_133.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_135 map[string]any) func() bool {
			hx_field_136 := hx_obj_135["hasNext"]
			if hx_field_136 == nil {
				var hx_zero_137 func() bool
				return hx_zero_137
			}
			return hx_field_136.(func() bool)
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
