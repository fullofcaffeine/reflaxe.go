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
	hx_obj_738 := map[string]any{}
	hx_obj_738["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_738["next"] = func() *string {
		hx_post_739 := index
		index = int(int32((index + 1)))
		return keys[hx_post_739]
	}
	return hx_obj_738
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_740 := map[string]any{}
	hx_obj_740["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_740["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_741 := index
			index = int(int32((index + 1)))
			return hx_post_741
		}()])
	}
	return hx_obj_740
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_742 any) map[string]any {
		if hx_value_742 == nil {
			var hx_zero_743 map[string]any
			return hx_zero_743
		}
		return hx_value_742.(map[string]any)
	}(self.keys())
	hx_obj_744 := map[string]any{}
	hx_obj_744["hasNext"] = func() bool {
		return func(hx_obj_745 map[string]any) func() bool {
			hx_field_746 := hx_obj_745["hasNext"]
			if hx_field_746 == nil {
				var hx_zero_747 func() bool
				return hx_zero_747
			}
			return hx_field_746.(func() bool)
		}(keys)()
	}
	hx_obj_744["next"] = func() map[string]any {
		key := func(hx_obj_748 map[string]any) func() *string {
			hx_field_749 := hx_obj_748["next"]
			if hx_field_749 == nil {
				var hx_zero_750 func() *string
				return hx_zero_750
			}
			return hx_field_749.(func() *string)
		}(keys)()
		hx_obj_751 := map[string]any{}
		hx_obj_751["key"] = key
		hx_obj_751["value"] = _gthis.get(key)
		return hx_obj_751
	}
	return hx_obj_744
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_752 any) *string {
		if hx_value_752 == nil {
			var hx_zero_753 *string
			return hx_zero_753
		}
		return hx_value_752.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_754 any) *string {
		if hx_value_754 == nil {
			var hx_zero_755 *string
			return hx_zero_755
		}
		return hx_value_754.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_758 any) bool {
		if hx_value_758 == nil {
			var hx_zero_759 bool
			return hx_zero_759
		}
		return hx_value_758.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_756 any) *string {
		if hx_value_756 == nil {
			var hx_zero_757 *string
			return hx_zero_757
		}
		return hx_value_756.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_762 any) bool {
		if hx_value_762 == nil {
			var hx_zero_763 bool
			return hx_zero_763
		}
		return hx_value_762.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_760 any) *string {
		if hx_value_760 == nil {
			var hx_zero_761 *string
			return hx_zero_761
		}
		return hx_value_760.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_764 any) *haxe__ds__StringMap {
		if hx_value_764 == nil {
			var hx_zero_765 *haxe__ds__StringMap
			return hx_zero_765
		}
		return hx_value_764.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_766 any) map[string]any {
		if hx_value_766 == nil {
			var hx_zero_767 map[string]any
			return hx_zero_767
		}
		return hx_value_766.(map[string]any)
	}(self.keys())
	for func(hx_obj_768 map[string]any) func() bool {
		hx_field_769 := hx_obj_768["hasNext"]
		if hx_field_769 == nil {
			var hx_zero_770 func() bool
			return hx_zero_770
		}
		return hx_field_769.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_771 map[string]any) func() *string {
			hx_field_772 := hx_obj_771["next"]
			if hx_field_772 == nil {
				var hx_zero_773 func() *string
				return hx_zero_773
			}
			return hx_field_772.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_774 any) map[string]any {
		if hx_value_774 == nil {
			var hx_zero_775 map[string]any
			return hx_zero_775
		}
		return hx_value_774.(map[string]any)
	}(self.keys())
	for func(hx_obj_776 map[string]any) func() bool {
		hx_field_777 := hx_obj_776["hasNext"]
		if hx_field_777 == nil {
			var hx_zero_778 func() bool
			return hx_zero_778
		}
		return hx_field_777.(func() bool)
	}(iterator)() {
		key := func(hx_obj_779 map[string]any) func() *string {
			hx_field_780 := hx_obj_779["next"]
			if hx_field_780 == nil {
				var hx_zero_781 func() *string
				return hx_zero_781
			}
			return hx_field_780.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_782 map[string]any) func() bool {
			hx_field_783 := hx_obj_782["hasNext"]
			if hx_field_783 == nil {
				var hx_zero_784 func() bool
				return hx_zero_784
			}
			return hx_field_783.(func() bool)
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
