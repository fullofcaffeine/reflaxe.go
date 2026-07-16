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
	hx_obj_109 := map[string]any{}
	hx_obj_109["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_109["next"] = func() *string {
		hx_post_110 := index
		index = int(int32((index + 1)))
		return keys[hx_post_110]
	}
	return hx_obj_109
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_111 := map[string]any{}
	hx_obj_111["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_111["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_112 := index
			index = int(int32((index + 1)))
			return hx_post_112
		}()])
	}
	return hx_obj_111
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_113 any) map[string]any {
		if hx_value_113 == nil {
			var hx_zero_114 map[string]any
			return hx_zero_114
		}
		return hx_value_113.(map[string]any)
	}(self.keys())
	hx_obj_115 := map[string]any{}
	hx_obj_115["hasNext"] = func() bool {
		return func(hx_obj_116 map[string]any) func() bool {
			hx_field_117 := hx_obj_116["hasNext"]
			if hx_field_117 == nil {
				var hx_zero_118 func() bool
				return hx_zero_118
			}
			return hx_field_117.(func() bool)
		}(keys)()
	}
	hx_obj_115["next"] = func() map[string]any {
		key := func(hx_obj_119 map[string]any) func() *string {
			hx_field_120 := hx_obj_119["next"]
			if hx_field_120 == nil {
				var hx_zero_121 func() *string
				return hx_zero_121
			}
			return hx_field_120.(func() *string)
		}(keys)()
		hx_obj_122 := map[string]any{}
		hx_obj_122["key"] = key
		hx_obj_122["value"] = _gthis.get(key)
		return hx_obj_122
	}
	return hx_obj_115
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_123 any) *string {
		if hx_value_123 == nil {
			var hx_zero_124 *string
			return hx_zero_124
		}
		return hx_value_123.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_125 any) *string {
		if hx_value_125 == nil {
			var hx_zero_126 *string
			return hx_zero_126
		}
		return hx_value_125.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_129 any) bool {
		if hx_value_129 == nil {
			var hx_zero_130 bool
			return hx_zero_130
		}
		return hx_value_129.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_127 any) *string {
		if hx_value_127 == nil {
			var hx_zero_128 *string
			return hx_zero_128
		}
		return hx_value_127.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_133 any) bool {
		if hx_value_133 == nil {
			var hx_zero_134 bool
			return hx_zero_134
		}
		return hx_value_133.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_131 any) *string {
		if hx_value_131 == nil {
			var hx_zero_132 *string
			return hx_zero_132
		}
		return hx_value_131.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_135 any) *haxe__ds__StringMap {
		if hx_value_135 == nil {
			var hx_zero_136 *haxe__ds__StringMap
			return hx_zero_136
		}
		return hx_value_135.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_137 any) map[string]any {
		if hx_value_137 == nil {
			var hx_zero_138 map[string]any
			return hx_zero_138
		}
		return hx_value_137.(map[string]any)
	}(self.keys())
	for func(hx_obj_139 map[string]any) func() bool {
		hx_field_140 := hx_obj_139["hasNext"]
		if hx_field_140 == nil {
			var hx_zero_141 func() bool
			return hx_zero_141
		}
		return hx_field_140.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_142 map[string]any) func() *string {
			hx_field_143 := hx_obj_142["next"]
			if hx_field_143 == nil {
				var hx_zero_144 func() *string
				return hx_zero_144
			}
			return hx_field_143.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_145 any) map[string]any {
		if hx_value_145 == nil {
			var hx_zero_146 map[string]any
			return hx_zero_146
		}
		return hx_value_145.(map[string]any)
	}(self.keys())
	for func(hx_obj_147 map[string]any) func() bool {
		hx_field_148 := hx_obj_147["hasNext"]
		if hx_field_148 == nil {
			var hx_zero_149 func() bool
			return hx_zero_149
		}
		return hx_field_148.(func() bool)
	}(iterator)() {
		key := func(hx_obj_150 map[string]any) func() *string {
			hx_field_151 := hx_obj_150["next"]
			if hx_field_151 == nil {
				var hx_zero_152 func() *string
				return hx_zero_152
			}
			return hx_field_151.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_153 map[string]any) func() bool {
			hx_field_154 := hx_obj_153["hasNext"]
			if hx_field_154 == nil {
				var hx_zero_155 func() bool
				return hx_zero_155
			}
			return hx_field_154.(func() bool)
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
