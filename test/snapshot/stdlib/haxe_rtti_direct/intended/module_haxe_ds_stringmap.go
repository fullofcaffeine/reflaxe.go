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
	hx_obj_773 := map[string]any{}
	hx_obj_773["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_773["next"] = func() *string {
		hx_post_774 := index
		index = int(int32((index + 1)))
		return keys[hx_post_774]
	}
	return hx_obj_773
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_775 := map[string]any{}
	hx_obj_775["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_775["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_776 := index
			index = int(int32((index + 1)))
			return hx_post_776
		}()])
	}
	return hx_obj_775
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_777 any) map[string]any {
		if hx_value_777 == nil {
			var hx_zero_778 map[string]any
			return hx_zero_778
		}
		return hx_value_777.(map[string]any)
	}(self.keys())
	hx_obj_779 := map[string]any{}
	hx_obj_779["hasNext"] = func() bool {
		return func(hx_obj_780 map[string]any) func() bool {
			hx_field_781 := hx_obj_780["hasNext"]
			if hx_field_781 == nil {
				var hx_zero_782 func() bool
				return hx_zero_782
			}
			return hx_field_781.(func() bool)
		}(keys)()
	}
	hx_obj_779["next"] = func() map[string]any {
		key := func(hx_obj_783 map[string]any) func() *string {
			hx_field_784 := hx_obj_783["next"]
			if hx_field_784 == nil {
				var hx_zero_785 func() *string
				return hx_zero_785
			}
			return hx_field_784.(func() *string)
		}(keys)()
		hx_obj_786 := map[string]any{}
		hx_obj_786["key"] = key
		hx_obj_786["value"] = _gthis.get(key)
		return hx_obj_786
	}
	return hx_obj_779
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_787 any) *string {
		if hx_value_787 == nil {
			var hx_zero_788 *string
			return hx_zero_788
		}
		return hx_value_787.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_789 any) *string {
		if hx_value_789 == nil {
			var hx_zero_790 *string
			return hx_zero_790
		}
		return hx_value_789.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_793 any) bool {
		if hx_value_793 == nil {
			var hx_zero_794 bool
			return hx_zero_794
		}
		return hx_value_793.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_791 any) *string {
		if hx_value_791 == nil {
			var hx_zero_792 *string
			return hx_zero_792
		}
		return hx_value_791.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_797 any) bool {
		if hx_value_797 == nil {
			var hx_zero_798 bool
			return hx_zero_798
		}
		return hx_value_797.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_795 any) *string {
		if hx_value_795 == nil {
			var hx_zero_796 *string
			return hx_zero_796
		}
		return hx_value_795.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_799 any) *haxe__ds__StringMap {
		if hx_value_799 == nil {
			var hx_zero_800 *haxe__ds__StringMap
			return hx_zero_800
		}
		return hx_value_799.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_801 any) map[string]any {
		if hx_value_801 == nil {
			var hx_zero_802 map[string]any
			return hx_zero_802
		}
		return hx_value_801.(map[string]any)
	}(self.keys())
	for func(hx_obj_803 map[string]any) func() bool {
		hx_field_804 := hx_obj_803["hasNext"]
		if hx_field_804 == nil {
			var hx_zero_805 func() bool
			return hx_zero_805
		}
		return hx_field_804.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_806 map[string]any) func() *string {
			hx_field_807 := hx_obj_806["next"]
			if hx_field_807 == nil {
				var hx_zero_808 func() *string
				return hx_zero_808
			}
			return hx_field_807.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_809 any) map[string]any {
		if hx_value_809 == nil {
			var hx_zero_810 map[string]any
			return hx_zero_810
		}
		return hx_value_809.(map[string]any)
	}(self.keys())
	for func(hx_obj_811 map[string]any) func() bool {
		hx_field_812 := hx_obj_811["hasNext"]
		if hx_field_812 == nil {
			var hx_zero_813 func() bool
			return hx_zero_813
		}
		return hx_field_812.(func() bool)
	}(iterator)() {
		key := func(hx_obj_814 map[string]any) func() *string {
			hx_field_815 := hx_obj_814["next"]
			if hx_field_815 == nil {
				var hx_zero_816 func() *string
				return hx_zero_816
			}
			return hx_field_815.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_817 map[string]any) func() bool {
			hx_field_818 := hx_obj_817["hasNext"]
			if hx_field_818 == nil {
				var hx_zero_819 func() bool
				return hx_zero_819
			}
			return hx_field_818.(func() bool)
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
