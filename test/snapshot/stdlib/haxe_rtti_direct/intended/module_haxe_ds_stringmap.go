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
	hx_obj_719 := map[string]any{}
	hx_obj_719["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_719["next"] = func() *string {
		hx_post_720 := index
		index = int(int32((index + 1)))
		return keys[hx_post_720]
	}
	return hx_obj_719
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_721 := map[string]any{}
	hx_obj_721["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_721["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_722 := index
			index = int(int32((index + 1)))
			return hx_post_722
		}()])
	}
	return hx_obj_721
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_723 any) map[string]any {
		if hx_value_723 == nil {
			var hx_zero_724 map[string]any
			return hx_zero_724
		}
		return hx_value_723.(map[string]any)
	}(self.keys())
	hx_obj_725 := map[string]any{}
	hx_obj_725["hasNext"] = func() bool {
		return func(hx_obj_726 map[string]any) func() bool {
			hx_field_727 := hx_obj_726["hasNext"]
			if hx_field_727 == nil {
				var hx_zero_728 func() bool
				return hx_zero_728
			}
			return hx_field_727.(func() bool)
		}(keys)()
	}
	hx_obj_725["next"] = func() map[string]any {
		key := func(hx_obj_729 map[string]any) func() *string {
			hx_field_730 := hx_obj_729["next"]
			if hx_field_730 == nil {
				var hx_zero_731 func() *string
				return hx_zero_731
			}
			return hx_field_730.(func() *string)
		}(keys)()
		hx_obj_732 := map[string]any{}
		hx_obj_732["key"] = key
		hx_obj_732["value"] = _gthis.get(key)
		return hx_obj_732
	}
	return hx_obj_725
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_733 any) *string {
		if hx_value_733 == nil {
			var hx_zero_734 *string
			return hx_zero_734
		}
		return hx_value_733.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_735 any) *string {
		if hx_value_735 == nil {
			var hx_zero_736 *string
			return hx_zero_736
		}
		return hx_value_735.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_739 any) bool {
		if hx_value_739 == nil {
			var hx_zero_740 bool
			return hx_zero_740
		}
		return hx_value_739.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_737 any) *string {
		if hx_value_737 == nil {
			var hx_zero_738 *string
			return hx_zero_738
		}
		return hx_value_737.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_743 any) bool {
		if hx_value_743 == nil {
			var hx_zero_744 bool
			return hx_zero_744
		}
		return hx_value_743.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_741 any) *string {
		if hx_value_741 == nil {
			var hx_zero_742 *string
			return hx_zero_742
		}
		return hx_value_741.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_745 any) *haxe__ds__StringMap {
		if hx_value_745 == nil {
			var hx_zero_746 *haxe__ds__StringMap
			return hx_zero_746
		}
		return hx_value_745.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_747 any) map[string]any {
		if hx_value_747 == nil {
			var hx_zero_748 map[string]any
			return hx_zero_748
		}
		return hx_value_747.(map[string]any)
	}(self.keys())
	for func(hx_obj_749 map[string]any) func() bool {
		hx_field_750 := hx_obj_749["hasNext"]
		if hx_field_750 == nil {
			var hx_zero_751 func() bool
			return hx_zero_751
		}
		return hx_field_750.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_752 map[string]any) func() *string {
			hx_field_753 := hx_obj_752["next"]
			if hx_field_753 == nil {
				var hx_zero_754 func() *string
				return hx_zero_754
			}
			return hx_field_753.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_755 any) map[string]any {
		if hx_value_755 == nil {
			var hx_zero_756 map[string]any
			return hx_zero_756
		}
		return hx_value_755.(map[string]any)
	}(self.keys())
	for func(hx_obj_757 map[string]any) func() bool {
		hx_field_758 := hx_obj_757["hasNext"]
		if hx_field_758 == nil {
			var hx_zero_759 func() bool
			return hx_zero_759
		}
		return hx_field_758.(func() bool)
	}(iterator)() {
		key := func(hx_obj_760 map[string]any) func() *string {
			hx_field_761 := hx_obj_760["next"]
			if hx_field_761 == nil {
				var hx_zero_762 func() *string
				return hx_zero_762
			}
			return hx_field_761.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_763 map[string]any) func() bool {
			hx_field_764 := hx_obj_763["hasNext"]
			if hx_field_764 == nil {
				var hx_zero_765 func() bool
				return hx_zero_765
			}
			return hx_field_764.(func() bool)
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
