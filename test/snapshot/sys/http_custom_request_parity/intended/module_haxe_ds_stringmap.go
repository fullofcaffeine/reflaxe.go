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
	hx_obj_61 := map[string]any{}
	hx_obj_61["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_61["next"] = func() *string {
		hx_post_62 := index
		index = int(int32((index + 1)))
		return keys[hx_post_62]
	}
	return hx_obj_61
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_63 := map[string]any{}
	hx_obj_63["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_63["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_64 := index
			index = int(int32((index + 1)))
			return hx_post_64
		}()])
	}
	return hx_obj_63
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_65 any) map[string]any {
		if hx_value_65 == nil {
			var hx_zero_66 map[string]any
			return hx_zero_66
		}
		return hx_value_65.(map[string]any)
	}(self.keys())
	hx_obj_67 := map[string]any{}
	hx_obj_67["hasNext"] = func() bool {
		return func(hx_obj_68 map[string]any) func() bool {
			hx_field_69 := hx_obj_68["hasNext"]
			if hx_field_69 == nil {
				var hx_zero_70 func() bool
				return hx_zero_70
			}
			return hx_field_69.(func() bool)
		}(keys)()
	}
	hx_obj_67["next"] = func() map[string]any {
		key := func(hx_obj_71 map[string]any) func() *string {
			hx_field_72 := hx_obj_71["next"]
			if hx_field_72 == nil {
				var hx_zero_73 func() *string
				return hx_zero_73
			}
			return hx_field_72.(func() *string)
		}(keys)()
		hx_obj_74 := map[string]any{}
		hx_obj_74["key"] = key
		hx_obj_74["value"] = _gthis.get(key)
		return hx_obj_74
	}
	return hx_obj_67
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_75 any) *string {
		if hx_value_75 == nil {
			var hx_zero_76 *string
			return hx_zero_76
		}
		return hx_value_75.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_77 any) *string {
		if hx_value_77 == nil {
			var hx_zero_78 *string
			return hx_zero_78
		}
		return hx_value_77.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_81 any) bool {
		if hx_value_81 == nil {
			var hx_zero_82 bool
			return hx_zero_82
		}
		return hx_value_81.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_79 any) *string {
		if hx_value_79 == nil {
			var hx_zero_80 *string
			return hx_zero_80
		}
		return hx_value_79.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_85 any) bool {
		if hx_value_85 == nil {
			var hx_zero_86 bool
			return hx_zero_86
		}
		return hx_value_85.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_83 any) *string {
		if hx_value_83 == nil {
			var hx_zero_84 *string
			return hx_zero_84
		}
		return hx_value_83.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_87 any) *haxe__ds__StringMap {
		if hx_value_87 == nil {
			var hx_zero_88 *haxe__ds__StringMap
			return hx_zero_88
		}
		return hx_value_87.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_89 any) map[string]any {
		if hx_value_89 == nil {
			var hx_zero_90 map[string]any
			return hx_zero_90
		}
		return hx_value_89.(map[string]any)
	}(self.keys())
	for func(hx_obj_91 map[string]any) func() bool {
		hx_field_92 := hx_obj_91["hasNext"]
		if hx_field_92 == nil {
			var hx_zero_93 func() bool
			return hx_zero_93
		}
		return hx_field_92.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_94 map[string]any) func() *string {
			hx_field_95 := hx_obj_94["next"]
			if hx_field_95 == nil {
				var hx_zero_96 func() *string
				return hx_zero_96
			}
			return hx_field_95.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_97 any) map[string]any {
		if hx_value_97 == nil {
			var hx_zero_98 map[string]any
			return hx_zero_98
		}
		return hx_value_97.(map[string]any)
	}(self.keys())
	for func(hx_obj_99 map[string]any) func() bool {
		hx_field_100 := hx_obj_99["hasNext"]
		if hx_field_100 == nil {
			var hx_zero_101 func() bool
			return hx_zero_101
		}
		return hx_field_100.(func() bool)
	}(iterator)() {
		key := func(hx_obj_102 map[string]any) func() *string {
			hx_field_103 := hx_obj_102["next"]
			if hx_field_103 == nil {
				var hx_zero_104 func() *string
				return hx_zero_104
			}
			return hx_field_103.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_105 map[string]any) func() bool {
			hx_field_106 := hx_obj_105["hasNext"]
			if hx_field_106 == nil {
				var hx_zero_107 func() bool
				return hx_zero_107
			}
			return hx_field_106.(func() bool)
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
