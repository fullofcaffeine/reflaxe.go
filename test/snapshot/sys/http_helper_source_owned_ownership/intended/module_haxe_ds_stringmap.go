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
	hx_obj_63 := map[string]any{}
	hx_obj_63["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_63["next"] = func() *string {
		hx_post_64 := index
		index = int(int32((index + 1)))
		return keys[hx_post_64]
	}
	return hx_obj_63
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_65 := map[string]any{}
	hx_obj_65["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_65["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_66 := index
			index = int(int32((index + 1)))
			return hx_post_66
		}()])
	}
	return hx_obj_65
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_67 any) map[string]any {
		if hx_value_67 == nil {
			var hx_zero_68 map[string]any
			return hx_zero_68
		}
		return hx_value_67.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_69 := map[string]any{}
	hx_obj_69["hasNext"] = func() bool {
		return func(hx_obj_70 map[string]any) func() bool {
			hx_field_71 := hx_obj_70["hasNext"]
			if hx_field_71 == nil {
				var hx_zero_72 func() bool
				return hx_zero_72
			}
			return hx_field_71.(func() bool)
		}(keys)()
	}
	hx_obj_69["next"] = func() map[string]any {
		key := func(hx_obj_73 map[string]any) func() *string {
			hx_field_74 := hx_obj_73["next"]
			if hx_field_74 == nil {
				var hx_zero_75 func() *string
				return hx_zero_75
			}
			return hx_field_74.(func() *string)
		}(keys)()
		hx_obj_76 := map[string]any{}
		hx_obj_76["key"] = key
		hx_obj_76["value"] = _gthis.__hx_this.get(key)
		return hx_obj_76
	}
	return hx_obj_69
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_77 any) *string {
		if hx_value_77 == nil {
			var hx_zero_78 *string
			return hx_zero_78
		}
		return hx_value_77.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_79 any) *string {
		if hx_value_79 == nil {
			var hx_zero_80 *string
			return hx_zero_80
		}
		return hx_value_79.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_83 any) bool {
		if hx_value_83 == nil {
			var hx_zero_84 bool
			return hx_zero_84
		}
		return hx_value_83.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_81 any) *string {
		if hx_value_81 == nil {
			var hx_zero_82 *string
			return hx_zero_82
		}
		return hx_value_81.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_87 any) bool {
		if hx_value_87 == nil {
			var hx_zero_88 bool
			return hx_zero_88
		}
		return hx_value_87.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_85 any) *string {
		if hx_value_85 == nil {
			var hx_zero_86 *string
			return hx_zero_86
		}
		return hx_value_85.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_89 any) *haxe__ds__StringMap {
		if hx_value_89 == nil {
			var hx_zero_90 *haxe__ds__StringMap
			return hx_zero_90
		}
		return hx_value_89.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_91 any) map[string]any {
		if hx_value_91 == nil {
			var hx_zero_92 map[string]any
			return hx_zero_92
		}
		return hx_value_91.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_93 map[string]any) func() bool {
		hx_field_94 := hx_obj_93["hasNext"]
		if hx_field_94 == nil {
			var hx_zero_95 func() bool
			return hx_zero_95
		}
		return hx_field_94.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_96 map[string]any) func() *string {
			hx_field_97 := hx_obj_96["next"]
			if hx_field_97 == nil {
				var hx_zero_98 func() *string
				return hx_zero_98
			}
			return hx_field_97.(func() *string)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_99 any) map[string]any {
		if hx_value_99 == nil {
			var hx_zero_100 map[string]any
			return hx_zero_100
		}
		return hx_value_99.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_101 map[string]any) func() bool {
		hx_field_102 := hx_obj_101["hasNext"]
		if hx_field_102 == nil {
			var hx_zero_103 func() bool
			return hx_zero_103
		}
		return hx_field_102.(func() bool)
	}(iterator)() {
		key := func(hx_obj_104 map[string]any) func() *string {
			hx_field_105 := hx_obj_104["next"]
			if hx_field_105 == nil {
				var hx_zero_106 func() *string
				return hx_zero_106
			}
			return hx_field_105.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_107 map[string]any) func() bool {
			hx_field_108 := hx_obj_107["hasNext"]
			if hx_field_108 == nil {
				var hx_zero_109 func() bool
				return hx_zero_109
			}
			return hx_field_108.(func() bool)
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
