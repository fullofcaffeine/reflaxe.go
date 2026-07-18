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
	hx_obj_33 := map[string]any{}
	hx_obj_33["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_33["next"] = func() *string {
		hx_post_34 := index
		index = int(int32((index + 1)))
		return keys[hx_post_34]
	}
	return hx_obj_33
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_35 := map[string]any{}
	hx_obj_35["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_35["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_36 := index
			index = int(int32((index + 1)))
			return hx_post_36
		}()])
	}
	return hx_obj_35
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_37 any) map[string]any {
		if hx_value_37 == nil {
			var hx_zero_38 map[string]any
			return hx_zero_38
		}
		return hx_value_37.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_39 := map[string]any{}
	hx_obj_39["hasNext"] = func() bool {
		return func(hx_obj_40 map[string]any) func() bool {
			hx_field_41 := hx_obj_40["hasNext"]
			if hx_field_41 == nil {
				var hx_zero_42 func() bool
				return hx_zero_42
			}
			return hx_field_41.(func() bool)
		}(keys)()
	}
	hx_obj_39["next"] = func() map[string]any {
		key := func(hx_obj_43 map[string]any) func() *string {
			hx_field_44 := hx_obj_43["next"]
			if hx_field_44 == nil {
				var hx_zero_45 func() *string
				return hx_zero_45
			}
			return hx_field_44.(func() *string)
		}(keys)()
		hx_obj_46 := map[string]any{}
		hx_obj_46["key"] = key
		hx_obj_46["value"] = _gthis.__hx_this.get(key)
		return hx_obj_46
	}
	return hx_obj_39
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_47 any) *string {
		if hx_value_47 == nil {
			var hx_zero_48 *string
			return hx_zero_48
		}
		return hx_value_47.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_49 any) *string {
		if hx_value_49 == nil {
			var hx_zero_50 *string
			return hx_zero_50
		}
		return hx_value_49.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_53 any) bool {
		if hx_value_53 == nil {
			var hx_zero_54 bool
			return hx_zero_54
		}
		return hx_value_53.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_51 any) *string {
		if hx_value_51 == nil {
			var hx_zero_52 *string
			return hx_zero_52
		}
		return hx_value_51.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_57 any) bool {
		if hx_value_57 == nil {
			var hx_zero_58 bool
			return hx_zero_58
		}
		return hx_value_57.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_55 any) *string {
		if hx_value_55 == nil {
			var hx_zero_56 *string
			return hx_zero_56
		}
		return hx_value_55.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_59 any) *haxe__ds__StringMap {
		if hx_value_59 == nil {
			var hx_zero_60 *haxe__ds__StringMap
			return hx_zero_60
		}
		return hx_value_59.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_61 any) map[string]any {
		if hx_value_61 == nil {
			var hx_zero_62 map[string]any
			return hx_zero_62
		}
		return hx_value_61.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_63 map[string]any) func() bool {
		hx_field_64 := hx_obj_63["hasNext"]
		if hx_field_64 == nil {
			var hx_zero_65 func() bool
			return hx_zero_65
		}
		return hx_field_64.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_66 map[string]any) func() *string {
			hx_field_67 := hx_obj_66["next"]
			if hx_field_67 == nil {
				var hx_zero_68 func() *string
				return hx_zero_68
			}
			return hx_field_67.(func() *string)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_69 any) map[string]any {
		if hx_value_69 == nil {
			var hx_zero_70 map[string]any
			return hx_zero_70
		}
		return hx_value_69.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_71 map[string]any) func() bool {
		hx_field_72 := hx_obj_71["hasNext"]
		if hx_field_72 == nil {
			var hx_zero_73 func() bool
			return hx_zero_73
		}
		return hx_field_72.(func() bool)
	}(iterator)() {
		key := func(hx_obj_74 map[string]any) func() *string {
			hx_field_75 := hx_obj_74["next"]
			if hx_field_75 == nil {
				var hx_zero_76 func() *string
				return hx_zero_76
			}
			return hx_field_75.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_77 map[string]any) func() bool {
			hx_field_78 := hx_obj_77["hasNext"]
			if hx_field_78 == nil {
				var hx_zero_79 func() bool
				return hx_zero_79
			}
			return hx_field_78.(func() bool)
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
