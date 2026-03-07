package main

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
	self.keysByHash.set(hash, k)
	self.valuesByHash.set(hash, v)
}

func (self *haxe__ds__HashMap) get(k interface{ hashCode() int }) any {
	return self.valuesByHash.get(k.hashCode())
}

func (self *haxe__ds__HashMap) exists(k interface{ hashCode() int }) bool {
	return func(hx_value_5 any) bool {
		if hx_value_5 == nil {
			var hx_zero_6 bool
			return hx_zero_6
		}
		return hx_value_5.(bool)
	}(self.valuesByHash.exists(k.hashCode()))
}

func (self *haxe__ds__HashMap) remove(k interface{ hashCode() int }) bool {
	hash := k.hashCode()
	func(hx_value_7 any) bool {
		if hx_value_7 == nil {
			var hx_zero_8 bool
			return hx_zero_8
		}
		return hx_value_7.(bool)
	}(self.valuesByHash.remove(hash))
	return func(hx_value_9 any) bool {
		if hx_value_9 == nil {
			var hx_zero_10 bool
			return hx_zero_10
		}
		return hx_value_9.(bool)
	}(self.keysByHash.remove(hash))
}

func (self *haxe__ds__HashMap) keys() map[string]any {
	_gthis := self
	hashes := func(hx_value_11 any) map[string]any {
		if hx_value_11 == nil {
			var hx_zero_12 map[string]any
			return hx_zero_12
		}
		return hx_value_11.(map[string]any)
	}(self.keysByHash.keys())
	hx_obj_13 := map[string]any{}
	hx_obj_13["hasNext"] = func() bool {
		return func(hx_obj_14 map[string]any) func() bool {
			hx_field_15 := hx_obj_14["hasNext"]
			if hx_field_15 == nil {
				var hx_zero_16 func() bool
				return hx_zero_16
			}
			return hx_field_15.(func() bool)
		}(hashes)()
	}
	hx_obj_13["next"] = func() interface{ hashCode() int } {
		return func(hx_value_20 any) interface{ hashCode() int } {
			if hx_value_20 == nil {
				var hx_zero_21 interface{ hashCode() int }
				return hx_zero_21
			}
			return hx_value_20.(interface{ hashCode() int })
		}(_gthis.keysByHash.get(func(hx_obj_17 map[string]any) func() int {
			hx_field_18 := hx_obj_17["next"]
			if hx_field_18 == nil {
				var hx_zero_19 func() int
				return hx_zero_19
			}
			return hx_field_18.(func() int)
		}(hashes)()))
	}
	return hx_obj_13
}

func (self *haxe__ds__HashMap) iterator() map[string]any {
	_gthis := self
	hashes := func(hx_value_22 any) map[string]any {
		if hx_value_22 == nil {
			var hx_zero_23 map[string]any
			return hx_zero_23
		}
		return hx_value_22.(map[string]any)
	}(self.valuesByHash.keys())
	hx_obj_24 := map[string]any{}
	hx_obj_24["hasNext"] = func() bool {
		return func(hx_obj_25 map[string]any) func() bool {
			hx_field_26 := hx_obj_25["hasNext"]
			if hx_field_26 == nil {
				var hx_zero_27 func() bool
				return hx_zero_27
			}
			return hx_field_26.(func() bool)
		}(hashes)()
	}
	hx_obj_24["next"] = func() any {
		return _gthis.valuesByHash.get(func(hx_obj_28 map[string]any) func() int {
			hx_field_29 := hx_obj_28["next"]
			if hx_field_29 == nil {
				var hx_zero_30 func() int
				return hx_zero_30
			}
			return hx_field_29.(func() int)
		}(hashes)())
	}
	return hx_obj_24
}

func (self *haxe__ds__HashMap) keyValueIterator() *haxe__iterators__HashMapKeyValueIterator {
	return New_haxe__iterators__HashMapKeyValueIterator(self)
}

func (self *haxe__ds__HashMap) copy() *haxe__ds__HashMap {
	copied := New_haxe__ds__HashMap()
	hash := func(hx_value_31 any) map[string]any {
		if hx_value_31 == nil {
			var hx_zero_32 map[string]any
			return hx_zero_32
		}
		return hx_value_31.(map[string]any)
	}(self.keysByHash.keys())
	for func(hx_obj_33 map[string]any) func() bool {
		hx_field_34 := hx_obj_33["hasNext"]
		if hx_field_34 == nil {
			var hx_zero_35 func() bool
			return hx_zero_35
		}
		return hx_field_34.(func() bool)
	}(hash)() {
		hash_1 := func(hx_obj_36 map[string]any) func() int {
			hx_field_37 := hx_obj_36["next"]
			if hx_field_37 == nil {
				var hx_zero_38 func() int
				return hx_zero_38
			}
			return hx_field_37.(func() int)
		}(hash)()
		copied.keysByHash.set(hash_1, func(hx_value_39 any) interface{ hashCode() int } {
			if hx_value_39 == nil {
				var hx_zero_40 interface{ hashCode() int }
				return hx_zero_40
			}
			return hx_value_39.(interface{ hashCode() int })
		}(self.keysByHash.get(hash_1)))
		copied.valuesByHash.set(hash_1, self.valuesByHash.get(hash_1))
	}
	return copied
}

func (self *haxe__ds__HashMap) clear() {
	self.keysByHash.clear()
	self.valuesByHash.clear()
}

func haxe__ds__HashMap_hashOf(key map[string]any) int {
	return func(hx_obj_41 map[string]any) func() int {
		hx_field_42 := hx_obj_41["hashCode"]
		if hx_field_42 == nil {
			var hx_zero_43 func() int
			return hx_zero_43
		}
		return hx_field_42.(func() int)
	}(key)()
}
