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
	hx_obj_129 := map[string]any{}
	hx_obj_129["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_129["next"] = func() *string {
		hx_post_130 := index
		index = int(int32((index + 1)))
		return keys[hx_post_130]
	}
	return hx_obj_129
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_131 := map[string]any{}
	hx_obj_131["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_131["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_132 := index
			index = int(int32((index + 1)))
			return hx_post_132
		}()])
	}
	return hx_obj_131
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_133 any) map[string]any {
		if hx_value_133 == nil {
			var hx_zero_134 map[string]any
			return hx_zero_134
		}
		return hx_value_133.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_135 := map[string]any{}
	hx_obj_135["hasNext"] = func() bool {
		return func(hx_obj_136 map[string]any) func() bool {
			hx_field_137 := hx_obj_136["hasNext"]
			if hx_field_137 == nil {
				var hx_zero_138 func() bool
				return hx_zero_138
			}
			return hx_field_137.(func() bool)
		}(keys)()
	}
	hx_obj_135["next"] = func() map[string]any {
		key := func(hx_obj_139 map[string]any) func() *string {
			hx_field_140 := hx_obj_139["next"]
			if hx_field_140 == nil {
				var hx_zero_141 func() *string
				return hx_zero_141
			}
			return hx_field_140.(func() *string)
		}(keys)()
		hx_obj_142 := map[string]any{}
		hx_obj_142["key"] = key
		hx_obj_142["value"] = _gthis.__hx_this.get(key)
		return hx_obj_142
	}
	return hx_obj_135
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_143 any) *string {
		if hx_value_143 == nil {
			var hx_zero_144 *string
			return hx_zero_144
		}
		return hx_value_143.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_145 any) *string {
		if hx_value_145 == nil {
			var hx_zero_146 *string
			return hx_zero_146
		}
		return hx_value_145.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_149 any) bool {
		if hx_value_149 == nil {
			var hx_zero_150 bool
			return hx_zero_150
		}
		return hx_value_149.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_147 any) *string {
		if hx_value_147 == nil {
			var hx_zero_148 *string
			return hx_zero_148
		}
		return hx_value_147.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_153 any) bool {
		if hx_value_153 == nil {
			var hx_zero_154 bool
			return hx_zero_154
		}
		return hx_value_153.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_151 any) *string {
		if hx_value_151 == nil {
			var hx_zero_152 *string
			return hx_zero_152
		}
		return hx_value_151.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_155 any) *haxe__ds__StringMap {
		if hx_value_155 == nil {
			var hx_zero_156 *haxe__ds__StringMap
			return hx_zero_156
		}
		return hx_value_155.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_157 any) map[string]any {
		if hx_value_157 == nil {
			var hx_zero_158 map[string]any
			return hx_zero_158
		}
		return hx_value_157.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_159 map[string]any) func() bool {
		hx_field_160 := hx_obj_159["hasNext"]
		if hx_field_160 == nil {
			var hx_zero_161 func() bool
			return hx_zero_161
		}
		return hx_field_160.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_162 map[string]any) func() *string {
			hx_field_163 := hx_obj_162["next"]
			if hx_field_163 == nil {
				var hx_zero_164 func() *string
				return hx_zero_164
			}
			return hx_field_163.(func() *string)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_165 any) map[string]any {
		if hx_value_165 == nil {
			var hx_zero_166 map[string]any
			return hx_zero_166
		}
		return hx_value_165.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_167 map[string]any) func() bool {
		hx_field_168 := hx_obj_167["hasNext"]
		if hx_field_168 == nil {
			var hx_zero_169 func() bool
			return hx_zero_169
		}
		return hx_field_168.(func() bool)
	}(iterator)() {
		key := func(hx_obj_170 map[string]any) func() *string {
			hx_field_171 := hx_obj_170["next"]
			if hx_field_171 == nil {
				var hx_zero_172 func() *string
				return hx_zero_172
			}
			return hx_field_171.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_173 map[string]any) func() bool {
			hx_field_174 := hx_obj_173["hasNext"]
			if hx_field_174 == nil {
				var hx_zero_175 func() bool
				return hx_zero_175
			}
			return hx_field_174.(func() bool)
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
