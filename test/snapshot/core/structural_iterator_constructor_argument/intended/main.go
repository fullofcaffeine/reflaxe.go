package main

import "snapshot/hxrt"

func checkedIterator(values *hxrt.Array) hxrt.Anonymous {
	hx_arr_1 := events
	hx_arr_1.Push(hxrt.StringFromLiteral("iterator"))
	return func() hxrt.Anonymous {
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
		return hxrt.AnonymousFromMap(hx_structural_iterator_map_4)
	}()
}

var events *hxrt.Array = hxrt.NewArray()

func main() {
	arrayValues := hxrt.NewArray(1, 2)
	arrayConsumer := New_SnapshotIntConsumer(mark(hxrt.StringFromLiteral("before")), func() hxrt.Anonymous {
		hx_arr_8 := events
		hx_arr_8.Push(hxrt.StringFromLiteral("iterator"))
		return func() hxrt.Anonymous {
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
			return hxrt.AnonymousFromMap(hx_structural_iterator_map_11)
		}()
	}(), mark(hxrt.StringFromLiteral("after")))
	hx_array_target_15 := arrayValues
	hx_array_index_16 := 0
	hx_array_target_15.Set(hx_array_index_16, 9)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("order="), hxrt.StringJoinAny(events.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("array="), arrayConsumer.collect()))
	hxrt.Println(v_1)
	genericConsumer := New_SnapshotGenericConsumer(func(hx_structural_source_17 *SnapshotGenericIterator) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_17, func(name string) hxrt.AnonymousField {
			switch name {
			case "hasNext":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
					hx_receiver_18 := hx_structural_source_17
					return hxrt.BoundFunction(func() bool {
						return hx_receiver_18.__hx_this.hasNext()
					}, func() any {
						return hx_receiver_18.__hx_this
					}, "hasNext")
				}(), Found: true}
			case "next":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
					hx_receiver_19 := hx_structural_source_17
					return hxrt.BoundFunction(func() any {
						return hx_receiver_19.__hx_this.next()
					}, func() any {
						return hx_receiver_19.__hx_this
					}, "next")
				}(), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			}
			return false
		})
	}(New_SnapshotGenericIterator(hxrt.StringFromLiteral("g1"), hxrt.StringFromLiteral("g2"))))
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("generic="), func(hx_value_20 any) *string {
		if hx_value_20 == nil {
			var hx_zero_21 *string
			return hx_zero_21
		}
		return hx_value_20.(*string)
	}(genericConsumer.collect())))
	hxrt.Println(v_2)
	baseIterator := New_SnapshotSpecializedIterator(hxrt.NewArray(hxrt.StringFromLiteral("v"))).SnapshotBaseIterator
	virtualConsumer := New_SnapshotGenericConsumer(func(hx_structural_source_22 *SnapshotBaseIterator) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_22, func(name string) hxrt.AnonymousField {
			switch name {
			case "hasNext":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
					hx_receiver_23 := hx_structural_source_22
					return hxrt.BoundFunction(func() bool {
						return hx_receiver_23.__hx_this.hasNext()
					}, func() any {
						return hx_receiver_23.__hx_this
					}, "hasNext")
				}(), Found: true}
			case "next":
				return hxrt.AnonymousField{Value: func(hx_callback_source_25 *hxrt.Function[func() *string]) *hxrt.Function[func() any] {
					return hxrt.AdaptFunction(hx_callback_source_25, func() any {
						return hxrt.FunctionCall(hx_callback_source_25)()
					})
				}(func() *hxrt.Function[func() *string] {
					hx_receiver_24 := hx_structural_source_22
					return hxrt.BoundFunction(func() *string {
						return hx_receiver_24.__hx_this.next()
					}, func() any {
						return hx_receiver_24.__hx_this
					}, "next")
				}()), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			}
			return false
		})
	}(baseIterator))
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("virtual="), func(hx_value_26 any) *string {
		if hx_value_26 == nil {
			var hx_zero_27 *string
			return hx_zero_27
		}
		return hx_value_26.(*string)
	}(virtualConsumer.collect())))
	hxrt.Println(v_3)
}

func mark(label *string) *string {
	hx_arr_28 := events
	hx_arr_28.Push(label)
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
		hx_post_32 := self.index
		self.index = int(int32((self.index + 1)))
		return hx_post_32
	}()))
}

type I_SnapshotGenericConsumer interface {
	collect() *string
}

type SnapshotGenericConsumer struct {
	__hx_this I_SnapshotGenericConsumer
	iterator  hxrt.Anonymous
}

func (self *SnapshotGenericConsumer) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_SnapshotGenericConsumer(iterator hxrt.Anonymous) *SnapshotGenericConsumer {
	self := &SnapshotGenericConsumer{}
	self.__hx_this = self
	self.iterator = iterator
	return self
}

func (self *SnapshotGenericConsumer) collect() *string {
	values := hxrt.NewArray()
	for func() bool {
		var hx_callable_36 *hxrt.Function[func() bool] = func(hx_obj_33 hxrt.Anonymous) *hxrt.Function[func() bool] {
			hx_field_34 := hxrt.AnonymousGet(hx_obj_33, "hasNext")
			if hx_field_34 == nil {
				var hx_zero_35 *hxrt.Function[func() bool]
				return hx_zero_35
			}
			return hx_field_34.(*hxrt.Function[func() bool])
		}(self.iterator)
		return hxrt.FunctionCall(hx_callable_36)()
	}() {
		values.Push(func() any {
			var hx_callable_41 *hxrt.Function[func() any] = func(hx_obj_38 hxrt.Anonymous) *hxrt.Function[func() any] {
				hx_field_39 := hxrt.AnonymousGet(hx_obj_38, "next")
				if hx_field_39 == nil {
					var hx_zero_40 *hxrt.Function[func() any]
					return hx_zero_40
				}
				return hx_field_39.(*hxrt.Function[func() any])
			}(self.iterator)
			return hxrt.FunctionCall(hx_callable_41)()
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
	var hx_if_43 any
	if func() int {
		hx_post_42 := self.index
		self.index = int(int32((self.index + 1)))
		return hx_post_42
	}() == 0 {
		hx_if_43 = self.first
	} else {
		hx_if_43 = self.second
	}
	return hx_if_43
}

type I_SnapshotIntConsumer interface {
	collect() *string
}

type SnapshotIntConsumer struct {
	__hx_this I_SnapshotIntConsumer
	iterator  hxrt.Anonymous
}

func (self *SnapshotIntConsumer) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_SnapshotIntConsumer(before *string, iterator hxrt.Anonymous, after *string) *SnapshotIntConsumer {
	self := &SnapshotIntConsumer{}
	self.__hx_this = self
	self.iterator = iterator
	return self
}

func (self *SnapshotIntConsumer) collect() *string {
	values := hxrt.NewArray()
	for func() bool {
		var hx_callable_47 *hxrt.Function[func() bool] = func(hx_obj_44 hxrt.Anonymous) *hxrt.Function[func() bool] {
			hx_field_45 := hxrt.AnonymousGet(hx_obj_44, "hasNext")
			if hx_field_45 == nil {
				var hx_zero_46 *hxrt.Function[func() bool]
				return hx_zero_46
			}
			return hx_field_45.(*hxrt.Function[func() bool])
		}(self.iterator)
		return hxrt.FunctionCall(hx_callable_47)()
	}() {
		values.Push(func() int {
			var hx_callable_52 *hxrt.Function[func() int] = func(hx_obj_49 hxrt.Anonymous) *hxrt.Function[func() int] {
				hx_field_50 := hxrt.AnonymousGet(hx_obj_49, "next")
				if hx_field_50 == nil {
					var hx_zero_51 *hxrt.Function[func() int]
					return hx_zero_51
				}
				return hx_field_50.(*hxrt.Function[func() int])
			}(self.iterator)
			return hxrt.FunctionCall(hx_callable_52)()
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
		hx_post_56 := self.index
		self.index = int(int32((self.index + 1)))
		return hx_post_56
	}()))
}
