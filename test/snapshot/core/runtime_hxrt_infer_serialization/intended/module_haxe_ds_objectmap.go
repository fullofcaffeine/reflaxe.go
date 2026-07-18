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
	hx_obj_52 := map[string]any{}
	hx_obj_52["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_52["next"] = func() any {
		return keys[func() int {
			hx_post_53 := index
			index = int(int32((index + 1)))
			return hx_post_53
		}()]
	}
	return hx_obj_52
}

func (self *haxe__ds__ObjectMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.ObjectMapKeys(self.h)
	index := 0
	hx_obj_54 := map[string]any{}
	hx_obj_54["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_54["next"] = func() any {
		return hxrt.ObjectMapGet(_gthis.h, keys[func() int {
			hx_post_55 := index
			index = int(int32((index + 1)))
			return hx_post_55
		}()])
	}
	return hx_obj_54
}

func (self *haxe__ds__ObjectMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_56 any) map[string]any {
		if hx_value_56 == nil {
			var hx_zero_57 map[string]any
			return hx_zero_57
		}
		return hx_value_56.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_58 := map[string]any{}
	hx_obj_58["hasNext"] = func() bool {
		return func(hx_obj_59 map[string]any) func() bool {
			hx_field_60 := hx_obj_59["hasNext"]
			if hx_field_60 == nil {
				var hx_zero_61 func() bool
				return hx_zero_61
			}
			return hx_field_60.(func() bool)
		}(keys)()
	}
	hx_obj_58["next"] = func() map[string]any {
		var key any = func(hx_obj_62 map[string]any) func() any {
			hx_field_63 := hx_obj_62["next"]
			if hx_field_63 == nil {
				var hx_zero_64 func() any
				return hx_zero_64
			}
			return hx_field_63.(func() any)
		}(keys)()
		hx_obj_65 := map[string]any{}
		hx_obj_65["key"] = key
		hx_obj_65["value"] = _gthis.__hx_this.get(key)
		return hx_obj_65
	}
	return hx_obj_58
}

func (self *haxe__ds__ObjectMap) getIMap(key any) any {
	return self.__hx_this.get(key)
}

func (self *haxe__ds__ObjectMap) setIMap(key any, value any) {
	self.__hx_this.set(key, value)
}

func (self *haxe__ds__ObjectMap) existsIMap(key any) bool {
	return func(hx_value_66 any) bool {
		if hx_value_66 == nil {
			var hx_zero_67 bool
			return hx_zero_67
		}
		return hx_value_66.(bool)
	}(self.__hx_this.exists(key))
}

func (self *haxe__ds__ObjectMap) removeIMap(key any) bool {
	return func(hx_value_68 any) bool {
		if hx_value_68 == nil {
			var hx_zero_69 bool
			return hx_zero_69
		}
		return hx_value_68.(bool)
	}(self.__hx_this.remove(key))
}

func (self *haxe__ds__ObjectMap) copyIMap() haxe__IMap {
	return func(hx_value_70 any) *haxe__ds__ObjectMap {
		if hx_value_70 == nil {
			var hx_zero_71 *haxe__ds__ObjectMap
			return hx_zero_71
		}
		return hx_value_70.(*haxe__ds__ObjectMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__ObjectMap) copy() *haxe__ds__ObjectMap {
	copied := New_haxe__ds__ObjectMap()
	key := func(hx_value_72 any) map[string]any {
		if hx_value_72 == nil {
			var hx_zero_73 map[string]any
			return hx_zero_73
		}
		return hx_value_72.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_74 map[string]any) func() bool {
		hx_field_75 := hx_obj_74["hasNext"]
		if hx_field_75 == nil {
			var hx_zero_76 func() bool
			return hx_zero_76
		}
		return hx_field_75.(func() bool)
	}(key)() {
		var key_1 any = func(hx_obj_77 map[string]any) func() any {
			hx_field_78 := hx_obj_77["next"]
			if hx_field_78 == nil {
				var hx_zero_79 func() any
				return hx_zero_79
			}
			return hx_field_78.(func() any)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__ObjectMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_80 any) map[string]any {
		if hx_value_80 == nil {
			var hx_zero_81 map[string]any
			return hx_zero_81
		}
		return hx_value_80.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_82 map[string]any) func() bool {
		hx_field_83 := hx_obj_82["hasNext"]
		if hx_field_83 == nil {
			var hx_zero_84 func() bool
			return hx_zero_84
		}
		return hx_field_83.(func() bool)
	}(iterator)() {
		var key any = func(hx_obj_85 map[string]any) func() any {
			hx_field_86 := hx_obj_85["next"]
			if hx_field_86 == nil {
				var hx_zero_87 func() any
				return hx_zero_87
			}
			return hx_field_86.(func() any)
		}(iterator)()
		x := hxrt.StdString(key)
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x_1 := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x_1))
		if func(hx_obj_88 map[string]any) func() bool {
			hx_field_89 := hx_obj_88["hasNext"]
			if hx_field_89 == nil {
				var hx_zero_90 func() bool
				return hx_zero_90
			}
			return hx_field_89.(func() bool)
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
