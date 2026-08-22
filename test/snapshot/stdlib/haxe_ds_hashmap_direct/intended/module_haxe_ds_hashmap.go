package main

import "snapshot/hxrt"

type I_haxe__ds__HashMap interface {
	set(k interface{ hashCode() int }, v any)
	get(k interface{ hashCode() int }) any
	exists(k interface{ hashCode() int }) bool
	remove(k interface{ hashCode() int }) bool
	keys() map[string]any
	iterator() map[string]any
	keyValueIterator() *haxe__iterators__HashMapKeyValueIterator
	copy() *haxe__ds__HashMap
	clear()
}

type haxe__ds__HashMap struct {
	__hx_this    I_haxe__ds__HashMap
	keysByHash   *haxe__ds__IntMap
	valuesByHash *haxe__ds__IntMap
}

func New_haxe__ds__HashMap() *haxe__ds__HashMap {
	self := &haxe__ds__HashMap{}
	self.__hx_this = self
	self.keysByHash = New_haxe__ds__IntMap()
	self.valuesByHash = New_haxe__ds__IntMap()
	return self
}

func (self *haxe__ds__HashMap) set(k interface{ hashCode() int }, v any) {
	hash := k.hashCode()
	self.keysByHash.__hx_this.set(hash, k)
	self.valuesByHash.__hx_this.set(hash, v)
}

func (self *haxe__ds__HashMap) get(k interface{ hashCode() int }) any {
	return self.valuesByHash.__hx_this.get(k.hashCode())
}

func (self *haxe__ds__HashMap) exists(k interface{ hashCode() int }) bool {
	return func(hx_value_1 any) bool {
		if hx_value_1 == nil {
			var hx_zero_2 bool
			return hx_zero_2
		}
		return hx_value_1.(bool)
	}(self.valuesByHash.__hx_this.exists(k.hashCode()))
}

func (self *haxe__ds__HashMap) remove(k interface{ hashCode() int }) bool {
	hash := k.hashCode()
	func(hx_value_3 any) bool {
		if hx_value_3 == nil {
			var hx_zero_4 bool
			return hx_zero_4
		}
		return hx_value_3.(bool)
	}(self.valuesByHash.__hx_this.remove(hash))
	return func(hx_value_5 any) bool {
		if hx_value_5 == nil {
			var hx_zero_6 bool
			return hx_zero_6
		}
		return hx_value_5.(bool)
	}(self.keysByHash.__hx_this.remove(hash))
}

func (self *haxe__ds__HashMap) keys() map[string]any {
	_gthis := self
	hashes := func(hx_value_7 any) map[string]any {
		if hx_value_7 == nil {
			var hx_zero_8 map[string]any
			return hx_zero_8
		}
		return hx_value_7.(map[string]any)
	}(self.keysByHash.__hx_this.keys())
	hx_obj_9 := map[string]any{}
	hx_obj_9["hasNext"] = func() bool {
		return func(hx_obj_10 map[string]any) func() bool {
			hx_field_11 := hx_obj_10["hasNext"]
			if hx_field_11 == nil {
				var hx_zero_12 func() bool
				return hx_zero_12
			}
			return hx_field_11.(func() bool)
		}(hashes)()
	}
	hx_obj_9["next"] = func() interface{ hashCode() int } {
		return func(hx_value_16 any) interface{ hashCode() int } {
			if hx_value_16 == nil {
				var hx_zero_17 interface{ hashCode() int }
				return hx_zero_17
			}
			return hx_value_16.(interface{ hashCode() int })
		}(_gthis.keysByHash.__hx_this.get(func(hx_obj_13 map[string]any) func() int {
			hx_field_14 := hx_obj_13["next"]
			if hx_field_14 == nil {
				var hx_zero_15 func() int
				return hx_zero_15
			}
			return hx_field_14.(func() int)
		}(hashes)()))
	}
	return hx_obj_9
}

func (self *haxe__ds__HashMap) iterator() map[string]any {
	_gthis := self
	hashes := func(hx_value_18 any) map[string]any {
		if hx_value_18 == nil {
			var hx_zero_19 map[string]any
			return hx_zero_19
		}
		return hx_value_18.(map[string]any)
	}(self.valuesByHash.__hx_this.keys())
	hx_obj_20 := map[string]any{}
	hx_obj_20["hasNext"] = func() bool {
		return func(hx_obj_21 map[string]any) func() bool {
			hx_field_22 := hx_obj_21["hasNext"]
			if hx_field_22 == nil {
				var hx_zero_23 func() bool
				return hx_zero_23
			}
			return hx_field_22.(func() bool)
		}(hashes)()
	}
	hx_obj_20["next"] = func() any {
		return _gthis.valuesByHash.__hx_this.get(func(hx_obj_24 map[string]any) func() int {
			hx_field_25 := hx_obj_24["next"]
			if hx_field_25 == nil {
				var hx_zero_26 func() int
				return hx_zero_26
			}
			return hx_field_25.(func() int)
		}(hashes)())
	}
	return hx_obj_20
}

func (self *haxe__ds__HashMap) keyValueIterator() *haxe__iterators__HashMapKeyValueIterator {
	return New_haxe__iterators__HashMapKeyValueIterator(self)
}

func (self *haxe__ds__HashMap) copy() *haxe__ds__HashMap {
	copied := New_haxe__ds__HashMap()
	hash := func(hx_value_27 any) map[string]any {
		if hx_value_27 == nil {
			var hx_zero_28 map[string]any
			return hx_zero_28
		}
		return hx_value_27.(map[string]any)
	}(self.keysByHash.__hx_this.keys())
	for func(hx_obj_29 map[string]any) func() bool {
		hx_field_30 := hx_obj_29["hasNext"]
		if hx_field_30 == nil {
			var hx_zero_31 func() bool
			return hx_zero_31
		}
		return hx_field_30.(func() bool)
	}(hash)() {
		hash_1 := func(hx_obj_32 map[string]any) func() int {
			hx_field_33 := hx_obj_32["next"]
			if hx_field_33 == nil {
				var hx_zero_34 func() int
				return hx_zero_34
			}
			return hx_field_33.(func() int)
		}(hash)()
		copied.keysByHash.__hx_this.set(hash_1, func(hx_value_35 any) interface{ hashCode() int } {
			if hx_value_35 == nil {
				var hx_zero_36 interface{ hashCode() int }
				return hx_zero_36
			}
			return hx_value_35.(interface{ hashCode() int })
		}(self.keysByHash.__hx_this.get(hash_1)))
		copied.valuesByHash.__hx_this.set(hash_1, self.valuesByHash.__hx_this.get(hash_1))
	}
	return copied
}

func (self *haxe__ds__HashMap) clear() {
	_this := self.keysByHash
	hxrt.IntMapClear(_this.h)
	_this_1 := self.valuesByHash
	hxrt.IntMapClear(_this_1.h)
}

func haxe__ds__HashMap_hashOf(key map[string]any) int {
	return func(hx_obj_37 map[string]any) func() int {
		hx_field_38 := hx_obj_37["hashCode"]
		if hx_field_38 == nil {
			var hx_zero_39 func() int
			return hx_zero_39
		}
		return hx_field_38.(func() int)
	}(key)()
}
