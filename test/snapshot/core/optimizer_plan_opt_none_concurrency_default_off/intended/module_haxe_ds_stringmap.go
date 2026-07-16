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
	hx_obj_17 := map[string]any{}
	hx_obj_17["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_17["next"] = func() *string {
		hx_post_18 := index
		index = int(int32((index + 1)))
		return keys[hx_post_18]
	}
	return hx_obj_17
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_19 := map[string]any{}
	hx_obj_19["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_19["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_20 := index
			index = int(int32((index + 1)))
			return hx_post_20
		}()])
	}
	return hx_obj_19
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_21 any) map[string]any {
		if hx_value_21 == nil {
			var hx_zero_22 map[string]any
			return hx_zero_22
		}
		return hx_value_21.(map[string]any)
	}(self.keys())
	hx_obj_23 := map[string]any{}
	hx_obj_23["hasNext"] = func() bool {
		return func(hx_obj_24 map[string]any) func() bool {
			hx_field_25 := hx_obj_24["hasNext"]
			if hx_field_25 == nil {
				var hx_zero_26 func() bool
				return hx_zero_26
			}
			return hx_field_25.(func() bool)
		}(keys)()
	}
	hx_obj_23["next"] = func() map[string]any {
		key := func(hx_obj_27 map[string]any) func() *string {
			hx_field_28 := hx_obj_27["next"]
			if hx_field_28 == nil {
				var hx_zero_29 func() *string
				return hx_zero_29
			}
			return hx_field_28.(func() *string)
		}(keys)()
		hx_obj_30 := map[string]any{}
		hx_obj_30["key"] = key
		hx_obj_30["value"] = _gthis.get(key)
		return hx_obj_30
	}
	return hx_obj_23
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_31 any) *string {
		if hx_value_31 == nil {
			var hx_zero_32 *string
			return hx_zero_32
		}
		return hx_value_31.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_33 any) *string {
		if hx_value_33 == nil {
			var hx_zero_34 *string
			return hx_zero_34
		}
		return hx_value_33.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_37 any) bool {
		if hx_value_37 == nil {
			var hx_zero_38 bool
			return hx_zero_38
		}
		return hx_value_37.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_35 any) *string {
		if hx_value_35 == nil {
			var hx_zero_36 *string
			return hx_zero_36
		}
		return hx_value_35.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_41 any) bool {
		if hx_value_41 == nil {
			var hx_zero_42 bool
			return hx_zero_42
		}
		return hx_value_41.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_39 any) *string {
		if hx_value_39 == nil {
			var hx_zero_40 *string
			return hx_zero_40
		}
		return hx_value_39.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_43 any) *haxe__ds__StringMap {
		if hx_value_43 == nil {
			var hx_zero_44 *haxe__ds__StringMap
			return hx_zero_44
		}
		return hx_value_43.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_45 any) map[string]any {
		if hx_value_45 == nil {
			var hx_zero_46 map[string]any
			return hx_zero_46
		}
		return hx_value_45.(map[string]any)
	}(self.keys())
	for func(hx_obj_47 map[string]any) func() bool {
		hx_field_48 := hx_obj_47["hasNext"]
		if hx_field_48 == nil {
			var hx_zero_49 func() bool
			return hx_zero_49
		}
		return hx_field_48.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_50 map[string]any) func() *string {
			hx_field_51 := hx_obj_50["next"]
			if hx_field_51 == nil {
				var hx_zero_52 func() *string
				return hx_zero_52
			}
			return hx_field_51.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_53 any) map[string]any {
		if hx_value_53 == nil {
			var hx_zero_54 map[string]any
			return hx_zero_54
		}
		return hx_value_53.(map[string]any)
	}(self.keys())
	for func(hx_obj_55 map[string]any) func() bool {
		hx_field_56 := hx_obj_55["hasNext"]
		if hx_field_56 == nil {
			var hx_zero_57 func() bool
			return hx_zero_57
		}
		return hx_field_56.(func() bool)
	}(iterator)() {
		key := func(hx_obj_58 map[string]any) func() *string {
			hx_field_59 := hx_obj_58["next"]
			if hx_field_59 == nil {
				var hx_zero_60 func() *string
				return hx_zero_60
			}
			return hx_field_59.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_61 map[string]any) func() bool {
			hx_field_62 := hx_obj_61["hasNext"]
			if hx_field_62 == nil {
				var hx_zero_63 func() bool
				return hx_zero_63
			}
			return hx_field_62.(func() bool)
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
