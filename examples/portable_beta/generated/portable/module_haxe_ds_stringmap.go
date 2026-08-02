package main

import "examples_portable_beta/hxrt"

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
	hx_obj_108 := map[string]any{}
	hx_obj_108["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_108["next"] = func() *string {
		hx_post_109 := index
		index = int(int32((index + 1)))
		return keys[hx_post_109]
	}
	return hx_obj_108
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_110 := map[string]any{}
	hx_obj_110["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_110["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_111 := index
			index = int(int32((index + 1)))
			return hx_post_111
		}()])
	}
	return hx_obj_110
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_112 any) map[string]any {
		if hx_value_112 == nil {
			var hx_zero_113 map[string]any
			return hx_zero_113
		}
		return hx_value_112.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_114 := map[string]any{}
	hx_obj_114["hasNext"] = func() bool {
		return func(hx_obj_115 map[string]any) func() bool {
			hx_field_116 := hx_obj_115["hasNext"]
			if hx_field_116 == nil {
				var hx_zero_117 func() bool
				return hx_zero_117
			}
			return hx_field_116.(func() bool)
		}(keys)()
	}
	hx_obj_114["next"] = func() map[string]any {
		key := func(hx_obj_118 map[string]any) func() *string {
			hx_field_119 := hx_obj_118["next"]
			if hx_field_119 == nil {
				var hx_zero_120 func() *string
				return hx_zero_120
			}
			return hx_field_119.(func() *string)
		}(keys)()
		hx_obj_121 := map[string]any{}
		hx_obj_121["key"] = key
		hx_obj_121["value"] = _gthis.__hx_this.get(key)
		return hx_obj_121
	}
	return hx_obj_114
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_122 any) *string {
		if hx_value_122 == nil {
			var hx_zero_123 *string
			return hx_zero_123
		}
		return hx_value_122.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_124 any) *string {
		if hx_value_124 == nil {
			var hx_zero_125 *string
			return hx_zero_125
		}
		return hx_value_124.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_128 any) bool {
		if hx_value_128 == nil {
			var hx_zero_129 bool
			return hx_zero_129
		}
		return hx_value_128.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_126 any) *string {
		if hx_value_126 == nil {
			var hx_zero_127 *string
			return hx_zero_127
		}
		return hx_value_126.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_132 any) bool {
		if hx_value_132 == nil {
			var hx_zero_133 bool
			return hx_zero_133
		}
		return hx_value_132.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_130 any) *string {
		if hx_value_130 == nil {
			var hx_zero_131 *string
			return hx_zero_131
		}
		return hx_value_130.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_134 any) *haxe__ds__StringMap {
		if hx_value_134 == nil {
			var hx_zero_135 *haxe__ds__StringMap
			return hx_zero_135
		}
		return hx_value_134.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_136 any) map[string]any {
		if hx_value_136 == nil {
			var hx_zero_137 map[string]any
			return hx_zero_137
		}
		return hx_value_136.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_138 map[string]any) func() bool {
		hx_field_139 := hx_obj_138["hasNext"]
		if hx_field_139 == nil {
			var hx_zero_140 func() bool
			return hx_zero_140
		}
		return hx_field_139.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_141 map[string]any) func() *string {
			hx_field_142 := hx_obj_141["next"]
			if hx_field_142 == nil {
				var hx_zero_143 func() *string
				return hx_zero_143
			}
			return hx_field_142.(func() *string)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_144 any) map[string]any {
		if hx_value_144 == nil {
			var hx_zero_145 map[string]any
			return hx_zero_145
		}
		return hx_value_144.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_146 map[string]any) func() bool {
		hx_field_147 := hx_obj_146["hasNext"]
		if hx_field_147 == nil {
			var hx_zero_148 func() bool
			return hx_zero_148
		}
		return hx_field_147.(func() bool)
	}(iterator)() {
		key := func(hx_obj_149 map[string]any) func() *string {
			hx_field_150 := hx_obj_149["next"]
			if hx_field_150 == nil {
				var hx_zero_151 func() *string
				return hx_zero_151
			}
			return hx_field_150.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_152 map[string]any) func() bool {
			hx_field_153 := hx_obj_152["hasNext"]
			if hx_field_153 == nil {
				var hx_zero_154 func() bool
				return hx_zero_154
			}
			return hx_field_153.(func() bool)
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
