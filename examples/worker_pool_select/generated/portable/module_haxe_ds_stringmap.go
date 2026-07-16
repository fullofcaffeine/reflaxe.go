package main

import "examples_worker_pool_select_portable/hxrt"

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
	hx_obj_36 := map[string]any{}
	hx_obj_36["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_36["next"] = func() *string {
		hx_post_37 := index
		index = int(int32((index + 1)))
		return keys[hx_post_37]
	}
	return hx_obj_36
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_38 := map[string]any{}
	hx_obj_38["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_38["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_39 := index
			index = int(int32((index + 1)))
			return hx_post_39
		}()])
	}
	return hx_obj_38
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_40 any) map[string]any {
		if hx_value_40 == nil {
			var hx_zero_41 map[string]any
			return hx_zero_41
		}
		return hx_value_40.(map[string]any)
	}(self.keys())
	hx_obj_42 := map[string]any{}
	hx_obj_42["hasNext"] = func() bool {
		return func(hx_obj_43 map[string]any) func() bool {
			hx_field_44 := hx_obj_43["hasNext"]
			if hx_field_44 == nil {
				var hx_zero_45 func() bool
				return hx_zero_45
			}
			return hx_field_44.(func() bool)
		}(keys)()
	}
	hx_obj_42["next"] = func() map[string]any {
		key := func(hx_obj_46 map[string]any) func() *string {
			hx_field_47 := hx_obj_46["next"]
			if hx_field_47 == nil {
				var hx_zero_48 func() *string
				return hx_zero_48
			}
			return hx_field_47.(func() *string)
		}(keys)()
		hx_obj_49 := map[string]any{}
		hx_obj_49["key"] = key
		hx_obj_49["value"] = _gthis.get(key)
		return hx_obj_49
	}
	return hx_obj_42
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_50 any) *string {
		if hx_value_50 == nil {
			var hx_zero_51 *string
			return hx_zero_51
		}
		return hx_value_50.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_52 any) *string {
		if hx_value_52 == nil {
			var hx_zero_53 *string
			return hx_zero_53
		}
		return hx_value_52.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_56 any) bool {
		if hx_value_56 == nil {
			var hx_zero_57 bool
			return hx_zero_57
		}
		return hx_value_56.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_54 any) *string {
		if hx_value_54 == nil {
			var hx_zero_55 *string
			return hx_zero_55
		}
		return hx_value_54.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_60 any) bool {
		if hx_value_60 == nil {
			var hx_zero_61 bool
			return hx_zero_61
		}
		return hx_value_60.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_58 any) *string {
		if hx_value_58 == nil {
			var hx_zero_59 *string
			return hx_zero_59
		}
		return hx_value_58.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_62 any) *haxe__ds__StringMap {
		if hx_value_62 == nil {
			var hx_zero_63 *haxe__ds__StringMap
			return hx_zero_63
		}
		return hx_value_62.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_64 any) map[string]any {
		if hx_value_64 == nil {
			var hx_zero_65 map[string]any
			return hx_zero_65
		}
		return hx_value_64.(map[string]any)
	}(self.keys())
	for func(hx_obj_66 map[string]any) func() bool {
		hx_field_67 := hx_obj_66["hasNext"]
		if hx_field_67 == nil {
			var hx_zero_68 func() bool
			return hx_zero_68
		}
		return hx_field_67.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_69 map[string]any) func() *string {
			hx_field_70 := hx_obj_69["next"]
			if hx_field_70 == nil {
				var hx_zero_71 func() *string
				return hx_zero_71
			}
			return hx_field_70.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_72 any) map[string]any {
		if hx_value_72 == nil {
			var hx_zero_73 map[string]any
			return hx_zero_73
		}
		return hx_value_72.(map[string]any)
	}(self.keys())
	for func(hx_obj_74 map[string]any) func() bool {
		hx_field_75 := hx_obj_74["hasNext"]
		if hx_field_75 == nil {
			var hx_zero_76 func() bool
			return hx_zero_76
		}
		return hx_field_75.(func() bool)
	}(iterator)() {
		key := func(hx_obj_77 map[string]any) func() *string {
			hx_field_78 := hx_obj_77["next"]
			if hx_field_78 == nil {
				var hx_zero_79 func() *string
				return hx_zero_79
			}
			return hx_field_78.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_80 map[string]any) func() bool {
			hx_field_81 := hx_obj_80["hasNext"]
			if hx_field_81 == nil {
				var hx_zero_82 func() bool
				return hx_zero_82
			}
			return hx_field_81.(func() bool)
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
