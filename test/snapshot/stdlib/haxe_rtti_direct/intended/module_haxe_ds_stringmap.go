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
	hx_obj_778 := map[string]any{}
	hx_obj_778["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_778["next"] = func() *string {
		hx_post_779 := index
		index = int(int32((index + 1)))
		return keys[hx_post_779]
	}
	return hx_obj_778
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_780 := map[string]any{}
	hx_obj_780["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_780["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_781 := index
			index = int(int32((index + 1)))
			return hx_post_781
		}()])
	}
	return hx_obj_780
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_782 any) map[string]any {
		if hx_value_782 == nil {
			var hx_zero_783 map[string]any
			return hx_zero_783
		}
		return hx_value_782.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_784 := map[string]any{}
	hx_obj_784["hasNext"] = func() bool {
		return func(hx_obj_785 map[string]any) func() bool {
			hx_field_786 := hx_obj_785["hasNext"]
			if hx_field_786 == nil {
				var hx_zero_787 func() bool
				return hx_zero_787
			}
			return hx_field_786.(func() bool)
		}(keys)()
	}
	hx_obj_784["next"] = func() map[string]any {
		key := func(hx_obj_788 map[string]any) func() *string {
			hx_field_789 := hx_obj_788["next"]
			if hx_field_789 == nil {
				var hx_zero_790 func() *string
				return hx_zero_790
			}
			return hx_field_789.(func() *string)
		}(keys)()
		hx_obj_791 := map[string]any{}
		hx_obj_791["key"] = key
		hx_obj_791["value"] = _gthis.__hx_this.get(key)
		return hx_obj_791
	}
	return hx_obj_784
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_792 any) *string {
		if hx_value_792 == nil {
			var hx_zero_793 *string
			return hx_zero_793
		}
		return hx_value_792.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_794 any) *string {
		if hx_value_794 == nil {
			var hx_zero_795 *string
			return hx_zero_795
		}
		return hx_value_794.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_798 any) bool {
		if hx_value_798 == nil {
			var hx_zero_799 bool
			return hx_zero_799
		}
		return hx_value_798.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_796 any) *string {
		if hx_value_796 == nil {
			var hx_zero_797 *string
			return hx_zero_797
		}
		return hx_value_796.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_802 any) bool {
		if hx_value_802 == nil {
			var hx_zero_803 bool
			return hx_zero_803
		}
		return hx_value_802.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_800 any) *string {
		if hx_value_800 == nil {
			var hx_zero_801 *string
			return hx_zero_801
		}
		return hx_value_800.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_804 any) *haxe__ds__StringMap {
		if hx_value_804 == nil {
			var hx_zero_805 *haxe__ds__StringMap
			return hx_zero_805
		}
		return hx_value_804.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_806 any) map[string]any {
		if hx_value_806 == nil {
			var hx_zero_807 map[string]any
			return hx_zero_807
		}
		return hx_value_806.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_808 map[string]any) func() bool {
		hx_field_809 := hx_obj_808["hasNext"]
		if hx_field_809 == nil {
			var hx_zero_810 func() bool
			return hx_zero_810
		}
		return hx_field_809.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_811 map[string]any) func() *string {
			hx_field_812 := hx_obj_811["next"]
			if hx_field_812 == nil {
				var hx_zero_813 func() *string
				return hx_zero_813
			}
			return hx_field_812.(func() *string)
		}(key)()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_814 any) map[string]any {
		if hx_value_814 == nil {
			var hx_zero_815 map[string]any
			return hx_zero_815
		}
		return hx_value_814.(map[string]any)
	}(self.__hx_this.keys())
	for func(hx_obj_816 map[string]any) func() bool {
		hx_field_817 := hx_obj_816["hasNext"]
		if hx_field_817 == nil {
			var hx_zero_818 func() bool
			return hx_zero_818
		}
		return hx_field_817.(func() bool)
	}(iterator)() {
		key := func(hx_obj_819 map[string]any) func() *string {
			hx_field_820 := hx_obj_819["next"]
			if hx_field_820 == nil {
				var hx_zero_821 func() *string
				return hx_zero_821
			}
			return hx_field_820.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_822 map[string]any) func() bool {
			hx_field_823 := hx_obj_822["hasNext"]
			if hx_field_823 == nil {
				var hx_zero_824 func() bool
				return hx_zero_824
			}
			return hx_field_823.(func() bool)
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
