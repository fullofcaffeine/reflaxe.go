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
	hx_obj_748 := map[string]any{}
	hx_obj_748["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_748["next"] = func() *string {
		hx_post_749 := index
		index = int(int32((index + 1)))
		return keys[hx_post_749]
	}
	return hx_obj_748
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_750 := map[string]any{}
	hx_obj_750["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_750["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_751 := index
			index = int(int32((index + 1)))
			return hx_post_751
		}()])
	}
	return hx_obj_750
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_752 any) map[string]any {
		if hx_value_752 == nil {
			var hx_zero_753 map[string]any
			return hx_zero_753
		}
		return hx_value_752.(map[string]any)
	}(self.keys())
	hx_obj_754 := map[string]any{}
	hx_obj_754["hasNext"] = func() bool {
		return func(hx_obj_755 map[string]any) func() bool {
			hx_field_756 := hx_obj_755["hasNext"]
			if hx_field_756 == nil {
				var hx_zero_757 func() bool
				return hx_zero_757
			}
			return hx_field_756.(func() bool)
		}(keys)()
	}
	hx_obj_754["next"] = func() map[string]any {
		key := func(hx_obj_758 map[string]any) func() *string {
			hx_field_759 := hx_obj_758["next"]
			if hx_field_759 == nil {
				var hx_zero_760 func() *string
				return hx_zero_760
			}
			return hx_field_759.(func() *string)
		}(keys)()
		hx_obj_761 := map[string]any{}
		hx_obj_761["key"] = key
		hx_obj_761["value"] = _gthis.get(key)
		return hx_obj_761
	}
	return hx_obj_754
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_762 any) *string {
		if hx_value_762 == nil {
			var hx_zero_763 *string
			return hx_zero_763
		}
		return hx_value_762.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_764 any) *string {
		if hx_value_764 == nil {
			var hx_zero_765 *string
			return hx_zero_765
		}
		return hx_value_764.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_768 any) bool {
		if hx_value_768 == nil {
			var hx_zero_769 bool
			return hx_zero_769
		}
		return hx_value_768.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_766 any) *string {
		if hx_value_766 == nil {
			var hx_zero_767 *string
			return hx_zero_767
		}
		return hx_value_766.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_772 any) bool {
		if hx_value_772 == nil {
			var hx_zero_773 bool
			return hx_zero_773
		}
		return hx_value_772.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_770 any) *string {
		if hx_value_770 == nil {
			var hx_zero_771 *string
			return hx_zero_771
		}
		return hx_value_770.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_774 any) *haxe__ds__StringMap {
		if hx_value_774 == nil {
			var hx_zero_775 *haxe__ds__StringMap
			return hx_zero_775
		}
		return hx_value_774.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_776 any) map[string]any {
		if hx_value_776 == nil {
			var hx_zero_777 map[string]any
			return hx_zero_777
		}
		return hx_value_776.(map[string]any)
	}(self.keys())
	for func(hx_obj_778 map[string]any) func() bool {
		hx_field_779 := hx_obj_778["hasNext"]
		if hx_field_779 == nil {
			var hx_zero_780 func() bool
			return hx_zero_780
		}
		return hx_field_779.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_781 map[string]any) func() *string {
			hx_field_782 := hx_obj_781["next"]
			if hx_field_782 == nil {
				var hx_zero_783 func() *string
				return hx_zero_783
			}
			return hx_field_782.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_784 any) map[string]any {
		if hx_value_784 == nil {
			var hx_zero_785 map[string]any
			return hx_zero_785
		}
		return hx_value_784.(map[string]any)
	}(self.keys())
	for func(hx_obj_786 map[string]any) func() bool {
		hx_field_787 := hx_obj_786["hasNext"]
		if hx_field_787 == nil {
			var hx_zero_788 func() bool
			return hx_zero_788
		}
		return hx_field_787.(func() bool)
	}(iterator)() {
		key := func(hx_obj_789 map[string]any) func() *string {
			hx_field_790 := hx_obj_789["next"]
			if hx_field_790 == nil {
				var hx_zero_791 func() *string
				return hx_zero_791
			}
			return hx_field_790.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_792 map[string]any) func() bool {
			hx_field_793 := hx_obj_792["hasNext"]
			if hx_field_793 == nil {
				var hx_zero_794 func() bool
				return hx_zero_794
			}
			return hx_field_793.(func() bool)
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
