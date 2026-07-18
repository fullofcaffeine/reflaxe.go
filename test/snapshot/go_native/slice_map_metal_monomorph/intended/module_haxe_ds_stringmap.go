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
	hx_obj_27 := map[string]any{}
	hx_obj_27["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_27["next"] = func() *string {
		hx_post_28 := index
		index = int(int32((index + 1)))
		return keys[hx_post_28]
	}
	return hx_obj_27
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_29 := map[string]any{}
	hx_obj_29["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_29["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_30 := index
			index = int(int32((index + 1)))
			return hx_post_30
		}()])
	}
	return hx_obj_29
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_31 any) map[string]any {
		if hx_value_31 == nil {
			var hx_zero_32 map[string]any
			return hx_zero_32
		}
		return hx_value_31.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_33 := map[string]any{}
	hx_obj_33["hasNext"] = func() bool {
		return func(hx_obj_34 map[string]any) func() bool {
			hx_field_35 := hx_obj_34["hasNext"]
			if hx_field_35 == nil {
				var hx_zero_36 func() bool
				return hx_zero_36
			}
			return hx_field_35.(func() bool)
		}(keys)()
	}
	hx_obj_33["next"] = func() map[string]any {
		key := func(hx_obj_37 map[string]any) func() *string {
			hx_field_38 := hx_obj_37["next"]
			if hx_field_38 == nil {
				var hx_zero_39 func() *string
				return hx_zero_39
			}
			return hx_field_38.(func() *string)
		}(keys)()
		hx_obj_40 := map[string]any{}
		hx_obj_40["key"] = key
		hx_obj_40["value"] = _gthis.__hx_this.get(key)
		return hx_obj_40
	}
	return hx_obj_33
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_41 any) *string {
		if hx_value_41 == nil {
			var hx_zero_42 *string
			return hx_zero_42
		}
		return hx_value_41.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_43 any) *string {
		if hx_value_43 == nil {
			var hx_zero_44 *string
			return hx_zero_44
		}
		return hx_value_43.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_47 any) bool {
		if hx_value_47 == nil {
			var hx_zero_48 bool
			return hx_zero_48
		}
		return hx_value_47.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_45 any) *string {
		if hx_value_45 == nil {
			var hx_zero_46 *string
			return hx_zero_46
		}
		return hx_value_45.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_51 any) bool {
		if hx_value_51 == nil {
			var hx_zero_52 bool
			return hx_zero_52
		}
		return hx_value_51.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_49 any) *string {
		if hx_value_49 == nil {
			var hx_zero_50 *string
			return hx_zero_50
		}
		return hx_value_49.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_53 any) *haxe__ds__StringMap {
		if hx_value_53 == nil {
			var hx_zero_54 *haxe__ds__StringMap
			return hx_zero_54
		}
		return hx_value_53.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_55 any) map[string]any {
		if hx_value_55 == nil {
			var hx_zero_56 map[string]any
			return hx_zero_56
		}
		return hx_value_55.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_57 map[string]any) func() bool {
		hx_field_58 := hx_obj_57["hasNext"]
		if hx_field_58 == nil {
			var hx_zero_59 func() bool
			return hx_zero_59
		}
		return hx_field_58.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_60 map[string]any) func() *string {
			hx_field_61 := hx_obj_60["next"]
			if hx_field_61 == nil {
				var hx_zero_62 func() *string
				return hx_zero_62
			}
			return hx_field_61.(func() *string)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_63 any) map[string]any {
		if hx_value_63 == nil {
			var hx_zero_64 map[string]any
			return hx_zero_64
		}
		return hx_value_63.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_65 map[string]any) func() bool {
		hx_field_66 := hx_obj_65["hasNext"]
		if hx_field_66 == nil {
			var hx_zero_67 func() bool
			return hx_zero_67
		}
		return hx_field_66.(func() bool)
	}(iterator)() {
		key := func(hx_obj_68 map[string]any) func() *string {
			hx_field_69 := hx_obj_68["next"]
			if hx_field_69 == nil {
				var hx_zero_70 func() *string
				return hx_zero_70
			}
			return hx_field_69.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_71 map[string]any) func() bool {
			hx_field_72 := hx_obj_71["hasNext"]
			if hx_field_72 == nil {
				var hx_zero_73 func() bool
				return hx_zero_73
			}
			return hx_field_72.(func() bool)
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
