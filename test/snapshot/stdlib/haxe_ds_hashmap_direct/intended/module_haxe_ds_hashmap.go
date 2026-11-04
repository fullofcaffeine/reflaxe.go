package main

import "snapshot/hxrt"

type I_haxe__ds__HashMap interface {
	set(k hxrt.Anonymous, v any)
	get(k hxrt.Anonymous) any
	exists(k hxrt.Anonymous) bool
	remove(k hxrt.Anonymous) bool
	keys() hxrt.Anonymous
	iterator() hxrt.Anonymous
	keyValueIterator() *haxe__iterators__HashMapKeyValueIterator
	copy() *haxe__ds__HashMap
	clear()
}

type haxe__ds__HashMap struct {
	__hx_this    I_haxe__ds__HashMap
	keysByHash   *haxe__ds__IntMap
	valuesByHash *haxe__ds__IntMap
}

func (self *haxe__ds__HashMap) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe__ds__HashMap() *haxe__ds__HashMap {
	self := &haxe__ds__HashMap{}
	self.__hx_this = self
	self.keysByHash = New_haxe__ds__IntMap()
	self.valuesByHash = New_haxe__ds__IntMap()
	return self
}

func (self *haxe__ds__HashMap) set(k hxrt.Anonymous, v any) {
	hash := func() int {
		var hx_callable_4 *hxrt.Function[func() int] = func(hx_obj_1 hxrt.Anonymous) *hxrt.Function[func() int] {
			hx_field_2 := hxrt.AnonymousGet(hx_obj_1, "hashCode")
			if hx_field_2 == nil {
				var hx_zero_3 *hxrt.Function[func() int]
				return hx_zero_3
			}
			return hx_field_2.(*hxrt.Function[func() int])
		}(k)
		return hxrt.FunctionCall(hx_callable_4)()
	}()
	self.keysByHash.__hx_this.set(hash, k)
	self.valuesByHash.__hx_this.set(hash, v)
}

func (self *haxe__ds__HashMap) get(k hxrt.Anonymous) any {
	return self.valuesByHash.__hx_this.get(func() int {
		var hx_callable_8 *hxrt.Function[func() int] = func(hx_obj_5 hxrt.Anonymous) *hxrt.Function[func() int] {
			hx_field_6 := hxrt.AnonymousGet(hx_obj_5, "hashCode")
			if hx_field_6 == nil {
				var hx_zero_7 *hxrt.Function[func() int]
				return hx_zero_7
			}
			return hx_field_6.(*hxrt.Function[func() int])
		}(k)
		return hxrt.FunctionCall(hx_callable_8)()
	}())
}

func (self *haxe__ds__HashMap) exists(k hxrt.Anonymous) bool {
	return func(hx_value_13 any) bool {
		if hx_value_13 == nil {
			var hx_zero_14 bool
			return hx_zero_14
		}
		return hx_value_13.(bool)
	}(self.valuesByHash.__hx_this.exists(func() int {
		var hx_callable_12 *hxrt.Function[func() int] = func(hx_obj_9 hxrt.Anonymous) *hxrt.Function[func() int] {
			hx_field_10 := hxrt.AnonymousGet(hx_obj_9, "hashCode")
			if hx_field_10 == nil {
				var hx_zero_11 *hxrt.Function[func() int]
				return hx_zero_11
			}
			return hx_field_10.(*hxrt.Function[func() int])
		}(k)
		return hxrt.FunctionCall(hx_callable_12)()
	}()))
}

func (self *haxe__ds__HashMap) remove(k hxrt.Anonymous) bool {
	hash := func() int {
		var hx_callable_18 *hxrt.Function[func() int] = func(hx_obj_15 hxrt.Anonymous) *hxrt.Function[func() int] {
			hx_field_16 := hxrt.AnonymousGet(hx_obj_15, "hashCode")
			if hx_field_16 == nil {
				var hx_zero_17 *hxrt.Function[func() int]
				return hx_zero_17
			}
			return hx_field_16.(*hxrt.Function[func() int])
		}(k)
		return hxrt.FunctionCall(hx_callable_18)()
	}()
	func(hx_value_19 any) bool {
		if hx_value_19 == nil {
			var hx_zero_20 bool
			return hx_zero_20
		}
		return hx_value_19.(bool)
	}(self.valuesByHash.__hx_this.remove(hash))
	return func(hx_value_21 any) bool {
		if hx_value_21 == nil {
			var hx_zero_22 bool
			return hx_zero_22
		}
		return hx_value_21.(bool)
	}(self.keysByHash.__hx_this.remove(hash))
}

func (self *haxe__ds__HashMap) keys() hxrt.Anonymous {
	_gthis := self
	hashes := self.keysByHash.__hx_this.keys()
	hx_obj_23 := map[string]any{}
	hx_obj_23["hasNext"] = hxrt.NewFunction(func() bool {
		return func() bool {
			var hx_callable_27 *hxrt.Function[func() bool] = func(hx_obj_24 hxrt.Anonymous) *hxrt.Function[func() bool] {
				hx_field_25 := hxrt.AnonymousGet(hx_obj_24, "hasNext")
				if hx_field_25 == nil {
					var hx_zero_26 *hxrt.Function[func() bool]
					return hx_zero_26
				}
				return hx_field_25.(*hxrt.Function[func() bool])
			}(hashes)
			return hxrt.FunctionCall(hx_callable_27)()
		}()
	})
	hx_obj_23["next"] = hxrt.NewFunction(func() hxrt.Anonymous {
		return hx_structural_recover_0(_gthis.keysByHash.__hx_this.get(func() int {
			var hx_callable_31 *hxrt.Function[func() int] = func(hx_obj_28 hxrt.Anonymous) *hxrt.Function[func() int] {
				hx_field_29 := hxrt.AnonymousGet(hx_obj_28, "next")
				if hx_field_29 == nil {
					var hx_zero_30 *hxrt.Function[func() int]
					return hx_zero_30
				}
				return hx_field_29.(*hxrt.Function[func() int])
			}(hashes)
			return hxrt.FunctionCall(hx_callable_31)()
		}()))
	})
	return hxrt.AnonymousFromMap(hx_obj_23)
}

func (self *haxe__ds__HashMap) iterator() hxrt.Anonymous {
	_gthis := self
	hashes := self.valuesByHash.__hx_this.keys()
	hx_obj_33 := map[string]any{}
	hx_obj_33["hasNext"] = hxrt.NewFunction(func() bool {
		return func() bool {
			var hx_callable_37 *hxrt.Function[func() bool] = func(hx_obj_34 hxrt.Anonymous) *hxrt.Function[func() bool] {
				hx_field_35 := hxrt.AnonymousGet(hx_obj_34, "hasNext")
				if hx_field_35 == nil {
					var hx_zero_36 *hxrt.Function[func() bool]
					return hx_zero_36
				}
				return hx_field_35.(*hxrt.Function[func() bool])
			}(hashes)
			return hxrt.FunctionCall(hx_callable_37)()
		}()
	})
	hx_obj_33["next"] = hxrt.NewFunction(func() any {
		return _gthis.valuesByHash.__hx_this.get(func() int {
			var hx_callable_41 *hxrt.Function[func() int] = func(hx_obj_38 hxrt.Anonymous) *hxrt.Function[func() int] {
				hx_field_39 := hxrt.AnonymousGet(hx_obj_38, "next")
				if hx_field_39 == nil {
					var hx_zero_40 *hxrt.Function[func() int]
					return hx_zero_40
				}
				return hx_field_39.(*hxrt.Function[func() int])
			}(hashes)
			return hxrt.FunctionCall(hx_callable_41)()
		}())
	})
	return hxrt.AnonymousFromMap(hx_obj_33)
}

func (self *haxe__ds__HashMap) keyValueIterator() *haxe__iterators__HashMapKeyValueIterator {
	return New_haxe__iterators__HashMapKeyValueIterator(self)
}

func (self *haxe__ds__HashMap) copy() *haxe__ds__HashMap {
	copied := New_haxe__ds__HashMap()
	hash := self.keysByHash.__hx_this.keys()
	for func() bool {
		var hx_callable_45 *hxrt.Function[func() bool] = func(hx_obj_42 hxrt.Anonymous) *hxrt.Function[func() bool] {
			hx_field_43 := hxrt.AnonymousGet(hx_obj_42, "hasNext")
			if hx_field_43 == nil {
				var hx_zero_44 *hxrt.Function[func() bool]
				return hx_zero_44
			}
			return hx_field_43.(*hxrt.Function[func() bool])
		}(hash)
		return hxrt.FunctionCall(hx_callable_45)()
	}() {
		hash_1 := func() int {
			var hx_callable_49 *hxrt.Function[func() int] = func(hx_obj_46 hxrt.Anonymous) *hxrt.Function[func() int] {
				hx_field_47 := hxrt.AnonymousGet(hx_obj_46, "next")
				if hx_field_47 == nil {
					var hx_zero_48 *hxrt.Function[func() int]
					return hx_zero_48
				}
				return hx_field_47.(*hxrt.Function[func() int])
			}(hash)
			return hxrt.FunctionCall(hx_callable_49)()
		}()
		copied.keysByHash.__hx_this.set(hash_1, hx_structural_recover_0(self.keysByHash.__hx_this.get(hash_1)))
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

func haxe__ds__HashMap_hashOf(key hxrt.Anonymous) int {
	return func() int {
		var hx_callable_54 *hxrt.Function[func() int] = func(hx_obj_51 hxrt.Anonymous) *hxrt.Function[func() int] {
			hx_field_52 := hxrt.AnonymousGet(hx_obj_51, "hashCode")
			if hx_field_52 == nil {
				var hx_zero_53 *hxrt.Function[func() int]
				return hx_zero_53
			}
			return hx_field_52.(*hxrt.Function[func() int])
		}(key)
		return hxrt.FunctionCall(hx_callable_54)()
	}()
}
