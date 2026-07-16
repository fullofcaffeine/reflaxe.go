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
	hx_obj_703 := map[string]any{}
	hx_obj_703["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_703["next"] = func() *string {
		hx_post_704 := index
		index = int(int32((index + 1)))
		return keys[hx_post_704]
	}
	return hx_obj_703
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_705 := map[string]any{}
	hx_obj_705["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_705["next"] = func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_706 := index
			index = int(int32((index + 1)))
			return hx_post_706
		}()])
	}
	return hx_obj_705
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_707 any) map[string]any {
		if hx_value_707 == nil {
			var hx_zero_708 map[string]any
			return hx_zero_708
		}
		return hx_value_707.(map[string]any)
	}(self.keys())
	hx_obj_709 := map[string]any{}
	hx_obj_709["hasNext"] = func() bool {
		return func(hx_obj_710 map[string]any) func() bool {
			hx_field_711 := hx_obj_710["hasNext"]
			if hx_field_711 == nil {
				var hx_zero_712 func() bool
				return hx_zero_712
			}
			return hx_field_711.(func() bool)
		}(keys)()
	}
	hx_obj_709["next"] = func() map[string]any {
		key := func(hx_obj_713 map[string]any) func() *string {
			hx_field_714 := hx_obj_713["next"]
			if hx_field_714 == nil {
				var hx_zero_715 func() *string
				return hx_zero_715
			}
			return hx_field_714.(func() *string)
		}(keys)()
		hx_obj_716 := map[string]any{}
		hx_obj_716["key"] = key
		hx_obj_716["value"] = _gthis.get(key)
		return hx_obj_716
	}
	return hx_obj_709
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.get(hxrt.StdString(func(hx_value_717 any) *string {
		if hx_value_717 == nil {
			var hx_zero_718 *string
			return hx_zero_718
		}
		return hx_value_717.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.set(hxrt.StdString(func(hx_value_719 any) *string {
		if hx_value_719 == nil {
			var hx_zero_720 *string
			return hx_zero_720
		}
		return hx_value_719.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_723 any) bool {
		if hx_value_723 == nil {
			var hx_zero_724 bool
			return hx_zero_724
		}
		return hx_value_723.(bool)
	}(self.exists(hxrt.StdString(func(hx_value_721 any) *string {
		if hx_value_721 == nil {
			var hx_zero_722 *string
			return hx_zero_722
		}
		return hx_value_721.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_727 any) bool {
		if hx_value_727 == nil {
			var hx_zero_728 bool
			return hx_zero_728
		}
		return hx_value_727.(bool)
	}(self.remove(hxrt.StdString(func(hx_value_725 any) *string {
		if hx_value_725 == nil {
			var hx_zero_726 *string
			return hx_zero_726
		}
		return hx_value_725.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_729 any) *haxe__ds__StringMap {
		if hx_value_729 == nil {
			var hx_zero_730 *haxe__ds__StringMap
			return hx_zero_730
		}
		return hx_value_729.(*haxe__ds__StringMap)
	}(self.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_731 any) map[string]any {
		if hx_value_731 == nil {
			var hx_zero_732 map[string]any
			return hx_zero_732
		}
		return hx_value_731.(map[string]any)
	}(self.keys())
	for func(hx_obj_733 map[string]any) func() bool {
		hx_field_734 := hx_obj_733["hasNext"]
		if hx_field_734 == nil {
			var hx_zero_735 func() bool
			return hx_zero_735
		}
		return hx_field_734.(func() bool)
	}(key)() {
		key_1 := func(hx_obj_736 map[string]any) func() *string {
			hx_field_737 := hx_obj_736["next"]
			if hx_field_737 == nil {
				var hx_zero_738 func() *string
				return hx_zero_738
			}
			return hx_field_737.(func() *string)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_739 any) map[string]any {
		if hx_value_739 == nil {
			var hx_zero_740 map[string]any
			return hx_zero_740
		}
		return hx_value_739.(map[string]any)
	}(self.keys())
	for func(hx_obj_741 map[string]any) func() bool {
		hx_field_742 := hx_obj_741["hasNext"]
		if hx_field_742 == nil {
			var hx_zero_743 func() bool
			return hx_zero_743
		}
		return hx_field_742.(func() bool)
	}(iterator)() {
		key := func(hx_obj_744 map[string]any) func() *string {
			hx_field_745 := hx_obj_744["next"]
			if hx_field_745 == nil {
				var hx_zero_746 func() *string
				return hx_zero_746
			}
			return hx_field_745.(func() *string)
		}(iterator)()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func(hx_obj_747 map[string]any) func() bool {
			hx_field_748 := hx_obj_747["hasNext"]
			if hx_field_748 == nil {
				var hx_zero_749 func() bool
				return hx_zero_749
			}
			return hx_field_748.(func() bool)
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
