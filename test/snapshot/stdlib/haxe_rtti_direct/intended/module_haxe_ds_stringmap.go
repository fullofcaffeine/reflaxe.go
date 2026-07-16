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
	hx_obj_735 := map[string]any{}
	hx_obj_735["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_735["next"] = func() *string {
		hx_post_736 := index
		index = int(int32((index + 1)))
		return keys[hx_post_736]
	}
	return hx_obj_735
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_737 := map[string]any{}
	hx_obj_737["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_737["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_738 := index
			index = int(int32((index + 1)))
			return hx_post_738
		}()])
	}
	return hx_obj_737
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_739 any) map[string]any {
		if hx_value_739 == nil {
			var hx_zero_740 map[string]any
			return hx_zero_740
		}
		return hx_value_739.(map[string]any)
	}(self.keys())
	hx_obj_741 := map[string]any{}
	hx_obj_741["hasNext"] = func() bool {
		return func(hx_obj_742 map[string]any) func() bool {
			hx_field_743 := hx_obj_742["hasNext"]
			if hx_field_743 == nil {
				var hx_zero_744 func() bool
				return hx_zero_744
			}
			return hx_field_743.(func() bool)
		}(keys)()
	}
	hx_obj_741["next"] = func() map[string]any {
		key := func(hx_obj_745 map[string]any) func() *string {
			hx_field_746 := hx_obj_745["next"]
			if hx_field_746 == nil {
				var hx_zero_747 func() *string
				return hx_zero_747
			}
			return hx_field_746.(func() *string)
		}(keys)()
		hx_obj_748 := map[string]any{}
		hx_obj_748["key"] = key
		hx_obj_748["value"] = _gthis.get(key)
		return hx_obj_748
	}
	return hx_obj_741
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_749 any) *string {
		if hx_value_749 == nil {
			var hx_zero_750 *string
			return hx_zero_750
		}
		return hx_value_749.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_751 any) *string {
		if hx_value_751 == nil {
			var hx_zero_752 *string
			return hx_zero_752
		}
		return hx_value_751.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_755 any) bool {
		if hx_value_755 == nil {
			var hx_zero_756 bool
			return hx_zero_756
		}
		return hx_value_755.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_753 any) *string {
		if hx_value_753 == nil {
			var hx_zero_754 *string
			return hx_zero_754
		}
		return hx_value_753.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_759 any) bool {
		if hx_value_759 == nil {
			var hx_zero_760 bool
			return hx_zero_760
		}
		return hx_value_759.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_757 any) *string {
		if hx_value_757 == nil {
			var hx_zero_758 *string
			return hx_zero_758
		}
		return hx_value_757.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_761 any) *haxe__ds__StringMap {
		if hx_value_761 == nil {
			var hx_zero_762 *haxe__ds__StringMap
			return hx_zero_762
		}
		return hx_value_761.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_763 any) map[string]any {
		if hx_value_763 == nil {
			var hx_zero_764 map[string]any
			return hx_zero_764
		}
		return hx_value_763.(map[string]any)
	}(self.keys())
	for func(hx_obj_765 map[string]any) func() bool {
		hx_field_766 := hx_obj_765["hasNext"]
		if hx_field_766 == nil {
			var hx_zero_767 func() bool
			return hx_zero_767
		}
		return hx_field_766.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_768 map[string]any) func() *string {
			hx_field_769 := hx_obj_768["next"]
			if hx_field_769 == nil {
				var hx_zero_770 func() *string
				return hx_zero_770
			}
			return hx_field_769.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_771 any) map[string]any {
		if hx_value_771 == nil {
			var hx_zero_772 map[string]any
			return hx_zero_772
		}
		return hx_value_771.(map[string]any)
	}(self.keys())
	for func(hx_obj_773 map[string]any) func() bool {
		hx_field_774 := hx_obj_773["hasNext"]
		if hx_field_774 == nil {
			var hx_zero_775 func() bool
			return hx_zero_775
		}
		return hx_field_774.(func() bool)
	}(iterator)() {
		key := func(hx_obj_776 map[string]any) func() *string {
			hx_field_777 := hx_obj_776["next"]
			if hx_field_777 == nil {
				var hx_zero_778 func() *string
				return hx_zero_778
			}
			return hx_field_777.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_779 map[string]any) func() bool {
			hx_field_780 := hx_obj_779["hasNext"]
			if hx_field_780 == nil {
				var hx_zero_781 func() bool
				return hx_zero_781
			}
			return hx_field_780.(func() bool)
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
