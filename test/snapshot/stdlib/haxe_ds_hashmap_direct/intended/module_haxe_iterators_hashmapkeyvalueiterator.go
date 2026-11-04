package main

import "snapshot/hxrt"

type I_haxe__iterators__HashMapKeyValueIterator interface {
	hasNext() bool
	next() hxrt.Anonymous
}

type haxe__iterators__HashMapKeyValueIterator struct {
	__hx_this I_haxe__iterators__HashMapKeyValueIterator
	map_      *haxe__ds__HashMap
	keys      hxrt.Anonymous
}

func (self *haxe__iterators__HashMapKeyValueIterator) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe__iterators__HashMapKeyValueIterator(map_ *haxe__ds__HashMap) *haxe__iterators__HashMapKeyValueIterator {
	self := &haxe__iterators__HashMapKeyValueIterator{}
	self.__hx_this = self
	self.map_ = map_
	_gthis := map_
	hashes := map_.keysByHash.__hx_this.keys()
	hx_obj_1 := map[string]any{}
	hx_obj_1["hasNext"] = hxrt.NewFunction(func() bool {
		return func() bool {
			var hx_callable_5 *hxrt.Function[func() bool] = func(hx_obj_2 hxrt.Anonymous) *hxrt.Function[func() bool] {
				hx_field_3 := hxrt.AnonymousGet(hx_obj_2, "hasNext")
				if hx_field_3 == nil {
					var hx_zero_4 *hxrt.Function[func() bool]
					return hx_zero_4
				}
				return hx_field_3.(*hxrt.Function[func() bool])
			}(hashes)
			return hxrt.FunctionCall(hx_callable_5)()
		}()
	})
	hx_obj_1["next"] = hxrt.NewFunction(func() hxrt.Anonymous {
		return hx_structural_recover_1(_gthis.keysByHash.__hx_this.get(func() int {
			var hx_callable_9 *hxrt.Function[func() int] = func(hx_obj_6 hxrt.Anonymous) *hxrt.Function[func() int] {
				hx_field_7 := hxrt.AnonymousGet(hx_obj_6, "next")
				if hx_field_7 == nil {
					var hx_zero_8 *hxrt.Function[func() int]
					return hx_zero_8
				}
				return hx_field_7.(*hxrt.Function[func() int])
			}(hashes)
			return hxrt.FunctionCall(hx_callable_9)()
		}()))
	})
	self.keys = hxrt.AnonymousFromMap(hx_obj_1)
	return self
}

func (self *haxe__iterators__HashMapKeyValueIterator) hasNext() bool {
	return func() bool {
		var hx_callable_14 *hxrt.Function[func() bool] = func(hx_obj_11 hxrt.Anonymous) *hxrt.Function[func() bool] {
			hx_field_12 := hxrt.AnonymousGet(hx_obj_11, "hasNext")
			if hx_field_12 == nil {
				var hx_zero_13 *hxrt.Function[func() bool]
				return hx_zero_13
			}
			return hx_field_12.(*hxrt.Function[func() bool])
		}(self.keys)
		return hxrt.FunctionCall(hx_callable_14)()
	}()
}

func (self *haxe__iterators__HashMapKeyValueIterator) next() hxrt.Anonymous {
	key := func() hxrt.Anonymous {
		var hx_callable_18 *hxrt.Function[func() hxrt.Anonymous] = func(hx_obj_15 hxrt.Anonymous) *hxrt.Function[func() hxrt.Anonymous] {
			hx_field_16 := hxrt.AnonymousGet(hx_obj_15, "next")
			if hx_field_16 == nil {
				var hx_zero_17 *hxrt.Function[func() hxrt.Anonymous]
				return hx_zero_17
			}
			return hx_field_16.(*hxrt.Function[func() hxrt.Anonymous])
		}(self.keys)
		return hxrt.FunctionCall(hx_callable_18)()
	}()
	_this := self.map_
	var value any = _this.valuesByHash.__hx_this.get(func() int {
		var hx_callable_22 *hxrt.Function[func() int] = func(hx_obj_19 hxrt.Anonymous) *hxrt.Function[func() int] {
			hx_field_20 := hxrt.AnonymousGet(hx_obj_19, "hashCode")
			if hx_field_20 == nil {
				var hx_zero_21 *hxrt.Function[func() int]
				return hx_zero_21
			}
			return hx_field_20.(*hxrt.Function[func() int])
		}(key)
		return hxrt.FunctionCall(hx_callable_22)()
	}())
	hx_obj_23 := map[string]any{}
	hx_obj_23["key"] = key
	hx_obj_23["value"] = value
	return hxrt.AnonymousFromMap(hx_obj_23)
}
