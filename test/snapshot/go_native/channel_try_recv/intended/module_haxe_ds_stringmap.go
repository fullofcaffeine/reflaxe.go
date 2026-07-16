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
	hx_obj_32 := map[string]any{}
	hx_obj_32["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_32["next"] = func() *string {
		hx_post_33 := index
		index = int(int32((index + 1)))
		return keys[hx_post_33]
	}
	return hx_obj_32
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_34 := map[string]any{}
	hx_obj_34["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_34["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_35 := index
			index = int(int32((index + 1)))
			return hx_post_35
		}()])
	}
	return hx_obj_34
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_36 any) map[string]any {
		if hx_value_36 == nil {
			var hx_zero_37 map[string]any
			return hx_zero_37
		}
		return hx_value_36.(map[string]any)
	}(self.keys())
	hx_obj_38 := map[string]any{}
	hx_obj_38["hasNext"] = func() bool {
		return func(hx_obj_39 map[string]any) func() bool {
			hx_field_40 := hx_obj_39["hasNext"]
			if hx_field_40 == nil {
				var hx_zero_41 func() bool
				return hx_zero_41
			}
			return hx_field_40.(func() bool)
		}(keys)()
	}
	hx_obj_38["next"] = func() map[string]any {
		key := func(hx_obj_42 map[string]any) func() *string {
			hx_field_43 := hx_obj_42["next"]
			if hx_field_43 == nil {
				var hx_zero_44 func() *string
				return hx_zero_44
			}
			return hx_field_43.(func() *string)
		}(keys)()
		hx_obj_45 := map[string]any{}
		hx_obj_45["key"] = key
		hx_obj_45["value"] = _gthis.get(key)
		return hx_obj_45
	}
	return hx_obj_38
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_46 any) *string {
		if hx_value_46 == nil {
			var hx_zero_47 *string
			return hx_zero_47
		}
		return hx_value_46.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_48 any) *string {
		if hx_value_48 == nil {
			var hx_zero_49 *string
			return hx_zero_49
		}
		return hx_value_48.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_52 any) bool {
		if hx_value_52 == nil {
			var hx_zero_53 bool
			return hx_zero_53
		}
		return hx_value_52.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_50 any) *string {
		if hx_value_50 == nil {
			var hx_zero_51 *string
			return hx_zero_51
		}
		return hx_value_50.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_56 any) bool {
		if hx_value_56 == nil {
			var hx_zero_57 bool
			return hx_zero_57
		}
		return hx_value_56.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_54 any) *string {
		if hx_value_54 == nil {
			var hx_zero_55 *string
			return hx_zero_55
		}
		return hx_value_54.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_58 any) *haxe__ds__StringMap {
		if hx_value_58 == nil {
			var hx_zero_59 *haxe__ds__StringMap
			return hx_zero_59
		}
		return hx_value_58.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_60 any) map[string]any {
		if hx_value_60 == nil {
			var hx_zero_61 map[string]any
			return hx_zero_61
		}
		return hx_value_60.(map[string]any)
	}(self.keys())
	for func(hx_obj_62 map[string]any) func() bool {
		hx_field_63 := hx_obj_62["hasNext"]
		if hx_field_63 == nil {
			var hx_zero_64 func() bool
			return hx_zero_64
		}
		return hx_field_63.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_65 map[string]any) func() *string {
			hx_field_66 := hx_obj_65["next"]
			if hx_field_66 == nil {
				var hx_zero_67 func() *string
				return hx_zero_67
			}
			return hx_field_66.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_68 any) map[string]any {
		if hx_value_68 == nil {
			var hx_zero_69 map[string]any
			return hx_zero_69
		}
		return hx_value_68.(map[string]any)
	}(self.keys())
	for func(hx_obj_70 map[string]any) func() bool {
		hx_field_71 := hx_obj_70["hasNext"]
		if hx_field_71 == nil {
			var hx_zero_72 func() bool
			return hx_zero_72
		}
		return hx_field_71.(func() bool)
	}(iterator)() {
		key := func(hx_obj_73 map[string]any) func() *string {
			hx_field_74 := hx_obj_73["next"]
			if hx_field_74 == nil {
				var hx_zero_75 func() *string
				return hx_zero_75
			}
			return hx_field_74.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_76 map[string]any) func() bool {
			hx_field_77 := hx_obj_76["hasNext"]
			if hx_field_77 == nil {
				var hx_zero_78 func() bool
				return hx_zero_78
			}
			return hx_field_77.(func() bool)
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
