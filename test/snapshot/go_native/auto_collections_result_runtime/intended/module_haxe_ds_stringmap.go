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
	hx_obj_46 := map[string]any{}
	hx_obj_46["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_46["next"] = func() *string {
		hx_post_47 := index
		index = int(int32((index + 1)))
		return keys[hx_post_47]
	}
	return hx_obj_46
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_48 := map[string]any{}
	hx_obj_48["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_48["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_49 := index
			index = int(int32((index + 1)))
			return hx_post_49
		}()])
	}
	return hx_obj_48
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_50 any) map[string]any {
		if hx_value_50 == nil {
			var hx_zero_51 map[string]any
			return hx_zero_51
		}
		return hx_value_50.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_52 := map[string]any{}
	hx_obj_52["hasNext"] = func() bool {
		return func(hx_obj_53 map[string]any) func() bool {
			hx_field_54 := hx_obj_53["hasNext"]
			if hx_field_54 == nil {
				var hx_zero_55 func() bool
				return hx_zero_55
			}
			return hx_field_54.(func() bool)
		}(keys)()
	}
	hx_obj_52["next"] = func() map[string]any {
		key := func(hx_obj_56 map[string]any) func() *string {
			hx_field_57 := hx_obj_56["next"]
			if hx_field_57 == nil {
				var hx_zero_58 func() *string
				return hx_zero_58
			}
			return hx_field_57.(func() *string)
		}(keys)()
		hx_obj_59 := map[string]any{}
		hx_obj_59["key"] = key
		hx_obj_59["value"] = _gthis.__hx_this.get(key)
		return hx_obj_59
	}
	return hx_obj_52
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_60 any) *string {
		if hx_value_60 == nil {
			var hx_zero_61 *string
			return hx_zero_61
		}
		return hx_value_60.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_62 any) *string {
		if hx_value_62 == nil {
			var hx_zero_63 *string
			return hx_zero_63
		}
		return hx_value_62.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_66 any) bool {
		if hx_value_66 == nil {
			var hx_zero_67 bool
			return hx_zero_67
		}
		return hx_value_66.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_64 any) *string {
		if hx_value_64 == nil {
			var hx_zero_65 *string
			return hx_zero_65
		}
		return hx_value_64.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_70 any) bool {
		if hx_value_70 == nil {
			var hx_zero_71 bool
			return hx_zero_71
		}
		return hx_value_70.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_68 any) *string {
		if hx_value_68 == nil {
			var hx_zero_69 *string
			return hx_zero_69
		}
		return hx_value_68.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_72 any) *haxe__ds__StringMap {
		if hx_value_72 == nil {
			var hx_zero_73 *haxe__ds__StringMap
			return hx_zero_73
		}
		return hx_value_72.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_74 any) map[string]any {
		if hx_value_74 == nil {
			var hx_zero_75 map[string]any
			return hx_zero_75
		}
		return hx_value_74.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_76 map[string]any) func() bool {
		hx_field_77 := hx_obj_76["hasNext"]
		if hx_field_77 == nil {
			var hx_zero_78 func() bool
			return hx_zero_78
		}
		return hx_field_77.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_79 map[string]any) func() *string {
			hx_field_80 := hx_obj_79["next"]
			if hx_field_80 == nil {
				var hx_zero_81 func() *string
				return hx_zero_81
			}
			return hx_field_80.(func() *string)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_82 any) map[string]any {
		if hx_value_82 == nil {
			var hx_zero_83 map[string]any
			return hx_zero_83
		}
		return hx_value_82.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_84 map[string]any) func() bool {
		hx_field_85 := hx_obj_84["hasNext"]
		if hx_field_85 == nil {
			var hx_zero_86 func() bool
			return hx_zero_86
		}
		return hx_field_85.(func() bool)
	}(iterator)() {
		key := func(hx_obj_87 map[string]any) func() *string {
			hx_field_88 := hx_obj_87["next"]
			if hx_field_88 == nil {
				var hx_zero_89 func() *string
				return hx_zero_89
			}
			return hx_field_88.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_90 map[string]any) func() bool {
			hx_field_91 := hx_obj_90["hasNext"]
			if hx_field_91 == nil {
				var hx_zero_92 func() bool
				return hx_zero_92
			}
			return hx_field_91.(func() bool)
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
