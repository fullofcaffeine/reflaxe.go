package main

import "snapshot/hxrt"

func checkedIterator(values *hxrt.Array) map[string]any {
	hx_arr_1 := events
	hx_arr_1.Push(hxrt.StringFromLiteral("iterator"))
	return func() map[string]any {
		hx_structural_array_2 := values
		hx_structural_array_index_3 := 0
		hx_structural_iterator_map_4 := map[string]any{}
		hx_structural_iterator_map_4["hasNext"] = hxrt.NewFunction(func() bool {
			return (hx_structural_array_index_3 < hx_structural_array_2.Len())
		})
		hx_structural_iterator_map_4["next"] = hxrt.NewFunction(func() int {
			hx_structural_array_value_5 := hx_structural_array_2.Get(hx_structural_array_index_3)
			hx_structural_array_index_3 = (hx_structural_array_index_3 + 1)
			return func(hx_value_6 any) int {
				if hx_value_6 == nil {
					var hx_zero_7 int
					return hx_zero_7
				}
				return hx_value_6.(int)
			}(any(hx_structural_array_value_5))
		})
		return hx_structural_iterator_map_4
	}()
}

var events *hxrt.Array = hxrt.NewArray()

func main() {
	arrayValues := hxrt.NewArray(1, 2)
	arrayConsumer := New_SnapshotIntConsumer(mark(hxrt.StringFromLiteral("before")), func() map[string]any {
		hx_arr_8 := events
		hx_arr_8.Push(hxrt.StringFromLiteral("iterator"))
		return func() map[string]any {
			hx_structural_array_9 := arrayValues
			hx_structural_array_index_10 := 0
			hx_structural_iterator_map_11 := map[string]any{}
			hx_structural_iterator_map_11["hasNext"] = hxrt.NewFunction(func() bool {
				return (hx_structural_array_index_10 < hx_structural_array_9.Len())
			})
			hx_structural_iterator_map_11["next"] = hxrt.NewFunction(func() int {
				hx_structural_array_value_12 := hx_structural_array_9.Get(hx_structural_array_index_10)
				hx_structural_array_index_10 = (hx_structural_array_index_10 + 1)
				return func(hx_value_13 any) int {
					if hx_value_13 == nil {
						var hx_zero_14 int
						return hx_zero_14
					}
					return hx_value_13.(int)
				}(any(hx_structural_array_value_12))
			})
			return hx_structural_iterator_map_11
		}()
	}(), mark(hxrt.StringFromLiteral("after")))
	hx_array_target_15 := arrayValues
	hx_array_index_16 := 0
	hx_array_target_15.Set(hx_array_index_16, 9)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("order="), hxrt.StringJoinAny(events.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("array="), arrayConsumer.collect()))
	hxrt.Println(v_1)
	genericConsumer := New_SnapshotGenericConsumer(func(hx_structural_iterator_17 *SnapshotGenericIterator) map[string]any {
		hx_structural_iterator_map_18 := map[string]any{}
		hx_structural_iterator_map_18["hasNext"] = hxrt.BoundFunction(func() bool {
			return hx_structural_iterator_17.__hx_this.hasNext()
		}, func() any {
			return hx_structural_iterator_17.__hx_this
		}, "hasNext")
		hx_structural_iterator_map_18["next"] = hxrt.BoundFunction(func() any {
			return any(hx_structural_iterator_17.__hx_this.next())
		}, func() any {
			return hx_structural_iterator_17.__hx_this
		}, "next")
		return hx_structural_iterator_map_18
	}(New_SnapshotGenericIterator(hxrt.StringFromLiteral("g1"), hxrt.StringFromLiteral("g2"))))
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("generic="), func(hx_value_19 any) *string {
		if hx_value_19 == nil {
			var hx_zero_20 *string
			return hx_zero_20
		}
		return hx_value_19.(*string)
	}(genericConsumer.collect())))
	hxrt.Println(v_2)
	baseIterator := New_SnapshotSpecializedIterator(hxrt.NewArray(hxrt.StringFromLiteral("v"))).SnapshotBaseIterator
	virtualConsumer := New_SnapshotGenericConsumer(func(hx_structural_iterator_21 *SnapshotBaseIterator) map[string]any {
		hx_structural_iterator_map_22 := map[string]any{}
		hx_structural_iterator_map_22["hasNext"] = hxrt.BoundFunction(func() bool {
			return hx_structural_iterator_21.__hx_this.hasNext()
		}, func() any {
			return hx_structural_iterator_21.__hx_this
		}, "hasNext")
		hx_structural_iterator_map_22["next"] = hxrt.BoundFunction(func() any {
			return any(hx_structural_iterator_21.__hx_this.next())
		}, func() any {
			return hx_structural_iterator_21.__hx_this
		}, "next")
		return hx_structural_iterator_map_22
	}(baseIterator))
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("virtual="), func(hx_value_23 any) *string {
		if hx_value_23 == nil {
			var hx_zero_24 *string
			return hx_zero_24
		}
		return hx_value_23.(*string)
	}(virtualConsumer.collect())))
	hxrt.Println(v_3)
}

func mark(label *string) *string {
	hx_arr_25 := events
	hx_arr_25.Push(label)
	return label
}

type I_SnapshotBaseIterator interface {
	hasNext() bool
	next() *string
}

type SnapshotBaseIterator struct {
	__hx_this I_SnapshotBaseIterator
	values    *hxrt.Array
	index     int
}

func (self *SnapshotBaseIterator) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_SnapshotBaseIterator(values *hxrt.Array) *SnapshotBaseIterator {
	self := &SnapshotBaseIterator{}
	self.__hx_this = self
	self.index = 0
	self.values = values
	return self
}

func (self *SnapshotBaseIterator) hasNext() bool {
	return (self.index < self.values.Len())
}

func (self *SnapshotBaseIterator) next() *string {
	return hxrt.StringConcatAny(hxrt.StringFromLiteral("base:"), self.values.Get(func() int {
		hx_post_29 := self.index
		self.index = int(int32((self.index + 1)))
		return hx_post_29
	}()))
}

type I_SnapshotGenericConsumer interface {
	collect() *string
}

type SnapshotGenericConsumer struct {
	__hx_this I_SnapshotGenericConsumer
	iterator  map[string]any
}

func (self *SnapshotGenericConsumer) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_SnapshotGenericConsumer(iterator map[string]any) *SnapshotGenericConsumer {
	self := &SnapshotGenericConsumer{}
	self.__hx_this = self
	self.iterator = iterator
	return self
}

func (self *SnapshotGenericConsumer) collect() *string {
	values := hxrt.NewArray()
	for func() bool {
		var hx_callable_33 *hxrt.Function[func() bool] = func(hx_obj_30 map[string]any) *hxrt.Function[func() bool] {
			hx_field_31 := hx_obj_30["hasNext"]
			if hx_field_31 == nil {
				var hx_zero_32 *hxrt.Function[func() bool]
				return hx_zero_32
			}
			return hx_field_31.(*hxrt.Function[func() bool])
		}(self.iterator)
		return hxrt.FunctionCall(hx_callable_33)()
	}() {
		values.Push(func() any {
			var hx_callable_38 *hxrt.Function[func() any] = func(hx_obj_35 map[string]any) *hxrt.Function[func() any] {
				hx_field_36 := hx_obj_35["next"]
				if hx_field_36 == nil {
					var hx_zero_37 *hxrt.Function[func() any]
					return hx_zero_37
				}
				return hx_field_36.(*hxrt.Function[func() any])
			}(self.iterator)
			return hxrt.FunctionCall(hx_callable_38)()
		}())
	}
	return hxrt.StringJoinAny(values.Values(), hxrt.StringFromLiteral(","))
}

type I_SnapshotGenericIterator interface {
	hasNext() bool
	next() any
}

type SnapshotGenericIterator struct {
	__hx_this I_SnapshotGenericIterator
	first     any
	second    any
	index     int
}

func (self *SnapshotGenericIterator) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_SnapshotGenericIterator(first any, second any) *SnapshotGenericIterator {
	self := &SnapshotGenericIterator{}
	self.__hx_this = self
	self.index = 0
	self.first = first
	self.second = second
	return self
}

func (self *SnapshotGenericIterator) hasNext() bool {
	return (self.index < 2)
}

func (self *SnapshotGenericIterator) next() any {
	var hx_if_40 any
	if func() int {
		hx_post_39 := self.index
		self.index = int(int32((self.index + 1)))
		return hx_post_39
	}() == 0 {
		hx_if_40 = self.first
	} else {
		hx_if_40 = self.second
	}
	return hx_if_40
}

type I_SnapshotIntConsumer interface {
	collect() *string
}

type SnapshotIntConsumer struct {
	__hx_this I_SnapshotIntConsumer
	iterator  map[string]any
}

func (self *SnapshotIntConsumer) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_SnapshotIntConsumer(before *string, iterator map[string]any, after *string) *SnapshotIntConsumer {
	self := &SnapshotIntConsumer{}
	self.__hx_this = self
	self.iterator = iterator
	return self
}

func (self *SnapshotIntConsumer) collect() *string {
	values := hxrt.NewArray()
	for func() bool {
		var hx_callable_44 *hxrt.Function[func() bool] = func(hx_obj_41 map[string]any) *hxrt.Function[func() bool] {
			hx_field_42 := hx_obj_41["hasNext"]
			if hx_field_42 == nil {
				var hx_zero_43 *hxrt.Function[func() bool]
				return hx_zero_43
			}
			return hx_field_42.(*hxrt.Function[func() bool])
		}(self.iterator)
		return hxrt.FunctionCall(hx_callable_44)()
	}() {
		values.Push(func() int {
			var hx_callable_49 *hxrt.Function[func() int] = func(hx_obj_46 map[string]any) *hxrt.Function[func() int] {
				hx_field_47 := hx_obj_46["next"]
				if hx_field_47 == nil {
					var hx_zero_48 *hxrt.Function[func() int]
					return hx_zero_48
				}
				return hx_field_47.(*hxrt.Function[func() int])
			}(self.iterator)
			return hxrt.FunctionCall(hx_callable_49)()
		}())
	}
	return hxrt.StringJoinAny(values.Values(), hxrt.StringFromLiteral(","))
}

type I_SnapshotSpecializedIterator interface {
	hasNext() bool
	next() *string
}

type SnapshotSpecializedIterator struct {
	*SnapshotBaseIterator
	__hx_this I_SnapshotSpecializedIterator
}

func (self *SnapshotSpecializedIterator) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_SnapshotSpecializedIterator(values *hxrt.Array) *SnapshotSpecializedIterator {
	self := &SnapshotSpecializedIterator{}
	self.SnapshotBaseIterator = New_SnapshotBaseIterator(values)
	self.SnapshotBaseIterator.__hx_this = self
	self.__hx_this = self
	return self
}

func (self *SnapshotSpecializedIterator) next() *string {
	return hxrt.StringConcatAny(hxrt.StringFromLiteral("special:"), self.values.Get(func() int {
		hx_post_53 := self.index
		self.index = int(int32((self.index + 1)))
		return hx_post_53
	}()))
}
