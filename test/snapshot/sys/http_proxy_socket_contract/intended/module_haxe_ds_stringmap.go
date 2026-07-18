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
	hx_obj_103 := map[string]any{}
	hx_obj_103["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_103["next"] = func() *string {
		hx_post_104 := index
		index = int(int32((index + 1)))
		return keys[hx_post_104]
	}
	return hx_obj_103
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_105 := map[string]any{}
	hx_obj_105["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_105["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_106 := index
			index = int(int32((index + 1)))
			return hx_post_106
		}()])
	}
	return hx_obj_105
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_107 any) map[string]any {
		if hx_value_107 == nil {
			var hx_zero_108 map[string]any
			return hx_zero_108
		}
		return hx_value_107.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_109 := map[string]any{}
	hx_obj_109["hasNext"] = func() bool {
		return func(hx_obj_110 map[string]any) func() bool {
			hx_field_111 := hx_obj_110["hasNext"]
			if hx_field_111 == nil {
				var hx_zero_112 func() bool
				return hx_zero_112
			}
			return hx_field_111.(func() bool)
		}(keys)()
	}
	hx_obj_109["next"] = func() map[string]any {
		key := func(hx_obj_113 map[string]any) func() *string {
			hx_field_114 := hx_obj_113["next"]
			if hx_field_114 == nil {
				var hx_zero_115 func() *string
				return hx_zero_115
			}
			return hx_field_114.(func() *string)
		}(keys)()
		hx_obj_116 := map[string]any{}
		hx_obj_116["key"] = key
		hx_obj_116["value"] = _gthis.__hx_this.get(key)
		return hx_obj_116
	}
	return hx_obj_109
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_117 any) *string {
		if hx_value_117 == nil {
			var hx_zero_118 *string
			return hx_zero_118
		}
		return hx_value_117.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_119 any) *string {
		if hx_value_119 == nil {
			var hx_zero_120 *string
			return hx_zero_120
		}
		return hx_value_119.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_123 any) bool {
		if hx_value_123 == nil {
			var hx_zero_124 bool
			return hx_zero_124
		}
		return hx_value_123.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_121 any) *string {
		if hx_value_121 == nil {
			var hx_zero_122 *string
			return hx_zero_122
		}
		return hx_value_121.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_127 any) bool {
		if hx_value_127 == nil {
			var hx_zero_128 bool
			return hx_zero_128
		}
		return hx_value_127.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_125 any) *string {
		if hx_value_125 == nil {
			var hx_zero_126 *string
			return hx_zero_126
		}
		return hx_value_125.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_129 any) *haxe__ds__StringMap {
		if hx_value_129 == nil {
			var hx_zero_130 *haxe__ds__StringMap
			return hx_zero_130
		}
		return hx_value_129.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_131 any) map[string]any {
		if hx_value_131 == nil {
			var hx_zero_132 map[string]any
			return hx_zero_132
		}
		return hx_value_131.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_133 map[string]any) func() bool {
		hx_field_134 := hx_obj_133["hasNext"]
		if hx_field_134 == nil {
			var hx_zero_135 func() bool
			return hx_zero_135
		}
		return hx_field_134.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_136 map[string]any) func() *string {
			hx_field_137 := hx_obj_136["next"]
			if hx_field_137 == nil {
				var hx_zero_138 func() *string
				return hx_zero_138
			}
			return hx_field_137.(func() *string)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_139 any) map[string]any {
		if hx_value_139 == nil {
			var hx_zero_140 map[string]any
			return hx_zero_140
		}
		return hx_value_139.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_141 map[string]any) func() bool {
		hx_field_142 := hx_obj_141["hasNext"]
		if hx_field_142 == nil {
			var hx_zero_143 func() bool
			return hx_zero_143
		}
		return hx_field_142.(func() bool)
	}(iterator)() {
		key := func(hx_obj_144 map[string]any) func() *string {
			hx_field_145 := hx_obj_144["next"]
			if hx_field_145 == nil {
				var hx_zero_146 func() *string
				return hx_zero_146
			}
			return hx_field_145.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_147 map[string]any) func() bool {
			hx_field_148 := hx_obj_147["hasNext"]
			if hx_field_148 == nil {
				var hx_zero_149 func() bool
				return hx_zero_149
			}
			return hx_field_148.(func() bool)
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
