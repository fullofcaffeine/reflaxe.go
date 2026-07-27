package main

import "snapshot/hxrt"

type I_haxe__ds__ObjectMap interface {
	set(key any, value any)
	get(key any) any
	exists(key any) bool
	remove(key any) bool
	keys() map[string]any
	iterator() map[string]any
	keyValueIterator() map[string]any
	getIMap(key any) any
	setIMap(key any, value any)
	existsIMap(key any) bool
	removeIMap(key any) bool
	copyIMap() haxe__IMap
	copy() *haxe__ds__ObjectMap
	toString() *string
	clear()
}

type haxe__ds__ObjectMap struct {
	__hx_this I_haxe__ds__ObjectMap
	h         *hxrt.ObjectMapCell
}

func New_haxe__ds__ObjectMap() *haxe__ds__ObjectMap {
	self := &haxe__ds__ObjectMap{}
	self.__hx_this = self
	self.h = hxrt.ObjectMapNew()
	return self
}

func (self *haxe__ds__ObjectMap) set(key any, value any) {
	hxrt.ObjectMapSet(self.h, key, value)
}

func (self *haxe__ds__ObjectMap) get(key any) any {
	return hxrt.ObjectMapGet(self.h, key)
}

func (self *haxe__ds__ObjectMap) exists(key any) bool {
	return hxrt.ObjectMapExists(self.h, key)
}

func (self *haxe__ds__ObjectMap) remove(key any) bool {
	return hxrt.ObjectMapRemove(self.h, key)
}

func (self *haxe__ds__ObjectMap) keys() map[string]any {
	keys := hxrt.ObjectMapKeys(self.h)
	index := 0
	hx_obj_53 := map[string]any{}
	hx_obj_53["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_53["next"] = func() any {
		return keys[func() int {
			hx_post_54 := index
			index = int(int32((index + 1)))
			return hx_post_54
		}()]
	}
	return hx_obj_53
}

func (self *haxe__ds__ObjectMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.ObjectMapKeys(self.h)
	index := 0
	hx_obj_55 := map[string]any{}
	hx_obj_55["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_55["next"] = func() any {
		return hxrt.ObjectMapGet(_gthis.h, keys[func() int {
			hx_post_56 := index
			index = int(int32((index + 1)))
			return hx_post_56
		}()])
	}
	return hx_obj_55
}

func (self *haxe__ds__ObjectMap) keyValueIterator() map[string]any {
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
		var key any = func(hx_obj_63 map[string]any) func() any {
			hx_field_64 := hx_obj_63["next"]
			if hx_field_64 == nil {
				var hx_zero_65 func() any
				return hx_zero_65
			}
			return hx_field_64.(func() any)
		}(keys)()
		hx_obj_66 := map[string]any{}
		hx_obj_66["key"] = key
		hx_obj_66["value"] = _gthis.__hx_this.get(key)
		return hx_obj_66
	}
	return hx_obj_59
}

func (self *haxe__ds__ObjectMap) getIMap(key any) any {
	return self.__hx_this.get(key)
}

func (self *haxe__ds__ObjectMap) setIMap(key any, value any) {
	self.__hx_this.set(key, value)
}

func (self *haxe__ds__ObjectMap) existsIMap(key any) bool {
	return func(hx_value_67 any) bool {
		if hx_value_67 == nil {
			var hx_zero_68 bool
			return hx_zero_68
		}
		return hx_value_67.(bool)
	}(self.__hx_this.exists(key))
}

func (self *haxe__ds__ObjectMap) removeIMap(key any) bool {
	return func(hx_value_69 any) bool {
		if hx_value_69 == nil {
			var hx_zero_70 bool
			return hx_zero_70
		}
		return hx_value_69.(bool)
	}(self.__hx_this.remove(key))
}

func (self *haxe__ds__ObjectMap) copyIMap() haxe__IMap {
	return func(hx_value_71 any) *haxe__ds__ObjectMap {
		if hx_value_71 == nil {
			var hx_zero_72 *haxe__ds__ObjectMap
			return hx_zero_72
		}
		return hx_value_71.(*haxe__ds__ObjectMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__ObjectMap) copy() *haxe__ds__ObjectMap {
	copied := New_haxe__ds__ObjectMap()
	key := func(hx_value_73 any) map[string]any {
		if hx_value_73 == nil {
			var hx_zero_74 map[string]any
			return hx_zero_74
		}
		return hx_value_73.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_75 map[string]any) func() bool {
		hx_field_76 := hx_obj_75["hasNext"]
		if hx_field_76 == nil {
			var hx_zero_77 func() bool
			return hx_zero_77
		}
		return hx_field_76.(func() bool)
	}(key)() {
		var key_1 any = func(hx_obj_78 map[string]any) func() any {
			hx_field_79 := hx_obj_78["next"]
			if hx_field_79 == nil {
				var hx_zero_80 func() any
				return hx_zero_80
			}
			return hx_field_79.(func() any)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__ObjectMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_81 any) map[string]any {
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
	}(iterator)() {
		var key any = func(hx_obj_86 map[string]any) func() any {
			hx_field_87 := hx_obj_86["next"]
			if hx_field_87 == nil {
				var hx_zero_88 func() any
				return hx_zero_88
			}
			return hx_field_87.(func() any)
		}(iterator)()
		x := hxrt.StdString(key)
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x_1 := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x_1))
		if func(hx_obj_89 map[string]any) func() bool {
			hx_field_90 := hx_obj_89["hasNext"]
			if hx_field_90 == nil {
				var hx_zero_91 func() bool
				return hx_zero_91
			}
			return hx_field_90.(func() bool)
		}(iterator)() {
			out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(", "))
		}
	}
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("]"))
	return out_b
}

func (self *haxe__ds__ObjectMap) clear() {
	hxrt.ObjectMapClear(self.h)
}

func (self *haxe__ds__ObjectMap) String() string {
	return *self.__hx_this.toString()
}
