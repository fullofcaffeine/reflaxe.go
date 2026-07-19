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
	hx_obj_53 := map[string]any{}
	hx_obj_53["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_53["next"] = func() *string {
		hx_post_54 := index
		index = int(int32((index + 1)))
		return keys[hx_post_54]
	}
	return hx_obj_53
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_55 := map[string]any{}
	hx_obj_55["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_55["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_56 := index
			index = int(int32((index + 1)))
			return hx_post_56
		}()])
	}
	return hx_obj_55
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_57 any) map[string]any {
		if hx_value_57 == nil {
			var hx_zero_58 map[string]any
			return hx_zero_58
		}
		return hx_value_57.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_59 := map[string]any{}
	hx_obj_59["hasNext"] = func() bool {
		return func(hx_obj_60 map[string]any) func() bool {
			hx_field_61 := hx_obj_60["hasNext"]
			if hx_field_61 == nil {
				var hx_zero_62 func() bool
				return hx_zero_62
			}
			return hx_field_61.(func() bool)
		}(keys)()
	}
	hx_obj_59["next"] = func() map[string]any {
		key := func(hx_obj_63 map[string]any) func() *string {
			hx_field_64 := hx_obj_63["next"]
			if hx_field_64 == nil {
				var hx_zero_65 func() *string
				return hx_zero_65
			}
			return hx_field_64.(func() *string)
		}(keys)()
		hx_obj_66 := map[string]any{}
		hx_obj_66["key"] = key
		hx_obj_66["value"] = _gthis.__hx_this.get(key)
		return hx_obj_66
	}
	return hx_obj_59
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_67 any) *string {
		if hx_value_67 == nil {
			var hx_zero_68 *string
			return hx_zero_68
		}
		return hx_value_67.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_69 any) *string {
		if hx_value_69 == nil {
			var hx_zero_70 *string
			return hx_zero_70
		}
		return hx_value_69.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_73 any) bool {
		if hx_value_73 == nil {
			var hx_zero_74 bool
			return hx_zero_74
		}
		return hx_value_73.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_71 any) *string {
		if hx_value_71 == nil {
			var hx_zero_72 *string
			return hx_zero_72
		}
		return hx_value_71.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_77 any) bool {
		if hx_value_77 == nil {
			var hx_zero_78 bool
			return hx_zero_78
		}
		return hx_value_77.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_75 any) *string {
		if hx_value_75 == nil {
			var hx_zero_76 *string
			return hx_zero_76
		}
		return hx_value_75.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_79 any) *haxe__ds__StringMap {
		if hx_value_79 == nil {
			var hx_zero_80 *haxe__ds__StringMap
			return hx_zero_80
		}
		return hx_value_79.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_81 any) map[string]any {
		if hx_value_81 == nil {
			var hx_zero_82 map[string]any
			return hx_zero_82
		}
		return hx_value_81.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_83 map[string]any) func() bool {
		hx_field_84 := hx_obj_83["hasNext"]
		if hx_field_84 == nil {
			var hx_zero_85 func() bool
			return hx_zero_85
		}
		return hx_field_84.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_86 map[string]any) func() *string {
			hx_field_87 := hx_obj_86["next"]
			if hx_field_87 == nil {
				var hx_zero_88 func() *string
				return hx_zero_88
			}
			return hx_field_87.(func() *string)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_89 any) map[string]any {
		if hx_value_89 == nil {
			var hx_zero_90 map[string]any
			return hx_zero_90
		}
		return hx_value_89.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_91 map[string]any) func() bool {
		hx_field_92 := hx_obj_91["hasNext"]
		if hx_field_92 == nil {
			var hx_zero_93 func() bool
			return hx_zero_93
		}
		return hx_field_92.(func() bool)
	}(iterator)() {
		key := func(hx_obj_94 map[string]any) func() *string {
			hx_field_95 := hx_obj_94["next"]
			if hx_field_95 == nil {
				var hx_zero_96 func() *string
				return hx_zero_96
			}
			return hx_field_95.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_97 map[string]any) func() bool {
			hx_field_98 := hx_obj_97["hasNext"]
			if hx_field_98 == nil {
				var hx_zero_99 func() bool
				return hx_zero_99
			}
			return hx_field_98.(func() bool)
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
