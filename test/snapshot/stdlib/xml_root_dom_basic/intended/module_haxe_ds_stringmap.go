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
	hx_obj_104 := map[string]any{}
	hx_obj_104["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_104["next"] = func() *string {
		hx_post_105 := index
		index = int(int32((index + 1)))
		return keys[hx_post_105]
	}
	return hx_obj_104
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_106 := map[string]any{}
	hx_obj_106["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_106["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_107 := index
			index = int(int32((index + 1)))
			return hx_post_107
		}()])
	}
	return hx_obj_106
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_108 any) map[string]any {
		if hx_value_108 == nil {
			var hx_zero_109 map[string]any
			return hx_zero_109
		}
		return hx_value_108.(map[string]any)
	}(self.keys())
	hx_obj_110 := map[string]any{}
	hx_obj_110["hasNext"] = func() bool {
		return func(hx_obj_111 map[string]any) func() bool {
			hx_field_112 := hx_obj_111["hasNext"]
			if hx_field_112 == nil {
				var hx_zero_113 func() bool
				return hx_zero_113
			}
			return hx_field_112.(func() bool)
		}(keys)()
	}
	hx_obj_110["next"] = func() map[string]any {
		key := func(hx_obj_114 map[string]any) func() *string {
			hx_field_115 := hx_obj_114["next"]
			if hx_field_115 == nil {
				var hx_zero_116 func() *string
				return hx_zero_116
			}
			return hx_field_115.(func() *string)
		}(keys)()
		hx_obj_117 := map[string]any{}
		hx_obj_117["key"] = key
		hx_obj_117["value"] = _gthis.get(key)
		return hx_obj_117
	}
	return hx_obj_110
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_118 any) *string {
		if hx_value_118 == nil {
			var hx_zero_119 *string
			return hx_zero_119
		}
		return hx_value_118.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_120 any) *string {
		if hx_value_120 == nil {
			var hx_zero_121 *string
			return hx_zero_121
		}
		return hx_value_120.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_124 any) bool {
		if hx_value_124 == nil {
			var hx_zero_125 bool
			return hx_zero_125
		}
		return hx_value_124.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_122 any) *string {
		if hx_value_122 == nil {
			var hx_zero_123 *string
			return hx_zero_123
		}
		return hx_value_122.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_128 any) bool {
		if hx_value_128 == nil {
			var hx_zero_129 bool
			return hx_zero_129
		}
		return hx_value_128.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_126 any) *string {
		if hx_value_126 == nil {
			var hx_zero_127 *string
			return hx_zero_127
		}
		return hx_value_126.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_130 any) *haxe__ds__StringMap {
		if hx_value_130 == nil {
			var hx_zero_131 *haxe__ds__StringMap
			return hx_zero_131
		}
		return hx_value_130.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_132 any) map[string]any {
		if hx_value_132 == nil {
			var hx_zero_133 map[string]any
			return hx_zero_133
		}
		return hx_value_132.(map[string]any)
	}(self.keys())
	for func(hx_obj_134 map[string]any) func() bool {
		hx_field_135 := hx_obj_134["hasNext"]
		if hx_field_135 == nil {
			var hx_zero_136 func() bool
			return hx_zero_136
		}
		return hx_field_135.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_137 map[string]any) func() *string {
			hx_field_138 := hx_obj_137["next"]
			if hx_field_138 == nil {
				var hx_zero_139 func() *string
				return hx_zero_139
			}
			return hx_field_138.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_140 any) map[string]any {
		if hx_value_140 == nil {
			var hx_zero_141 map[string]any
			return hx_zero_141
		}
		return hx_value_140.(map[string]any)
	}(self.keys())
	for func(hx_obj_142 map[string]any) func() bool {
		hx_field_143 := hx_obj_142["hasNext"]
		if hx_field_143 == nil {
			var hx_zero_144 func() bool
			return hx_zero_144
		}
		return hx_field_143.(func() bool)
	}(iterator)() {
		key := func(hx_obj_145 map[string]any) func() *string {
			hx_field_146 := hx_obj_145["next"]
			if hx_field_146 == nil {
				var hx_zero_147 func() *string
				return hx_zero_147
			}
			return hx_field_146.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_148 map[string]any) func() bool {
			hx_field_149 := hx_obj_148["hasNext"]
			if hx_field_149 == nil {
				var hx_zero_150 func() bool
				return hx_zero_150
			}
			return hx_field_149.(func() bool)
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
