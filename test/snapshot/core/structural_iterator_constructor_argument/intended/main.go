package main

import "snapshot/hxrt"

func checkedIterator(values []int) map[string]any {
	hx_arr_1 := events
	hx_arr_1 = append(hx_arr_1, hxrt.StringFromLiteral("iterator"))
	events = hx_arr_1
	return func() map[string]any {
		hx_structural_array_2 := values
		hx_structural_array_index_3 := 0
		hx_structural_iterator_map_4 := map[string]any{}
		hx_structural_iterator_map_4["hasNext"] = func() bool {
			return (hx_structural_array_index_3 < len(hx_structural_array_2))
		}
		hx_structural_iterator_map_4["next"] = func() int {
			hx_structural_array_value_5 := hx_structural_array_2[hx_structural_array_index_3]
			hx_structural_array_index_3 = (hx_structural_array_index_3 + 1)
			return hx_structural_array_value_5
		}
		return hx_structural_iterator_map_4
	}()
}

var events []*string = []*string{}

func main() {
	arrayValues := []int{1, 2}
	arrayConsumer := New_SnapshotIntConsumer(mark(hxrt.StringFromLiteral("before")), func() map[string]any {
		hx_arr_6 := events
		hx_arr_6 = append(hx_arr_6, hxrt.StringFromLiteral("iterator"))
		events = hx_arr_6
		return func() map[string]any {
			hx_structural_array_7 := arrayValues
			hx_structural_array_index_8 := 0
			hx_structural_iterator_map_9 := map[string]any{}
			hx_structural_iterator_map_9["hasNext"] = func() bool {
				return (hx_structural_array_index_8 < len(hx_structural_array_7))
			}
			hx_structural_iterator_map_9["next"] = func() int {
				hx_structural_array_value_10 := hx_structural_array_7[hx_structural_array_index_8]
				hx_structural_array_index_8 = (hx_structural_array_index_8 + 1)
				return hx_structural_array_value_10
			}
			return hx_structural_iterator_map_9
		}()
	}(), mark(hxrt.StringFromLiteral("after")))
	arrayValues[0] = 9
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("order="), hxrt.StringJoinAny(func(hx_sort_src_11 []*string) []any {
		hx_sort_out_13 := make([]any, 0, len(hx_sort_src_11))
		for _, hx_sort_item_12 := range hx_sort_src_11 {
			hx_sort_out_13 = append(hx_sort_out_13, hx_sort_item_12)
		}
		return hx_sort_out_13
	}(events), hxrt.StringFromLiteral(","))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("array="), arrayConsumer.collect()))
	hxrt.Println(v_1)
	genericConsumer := New_SnapshotGenericConsumer(func(hx_structural_iterator_14 *SnapshotGenericIterator) map[string]any {
		hx_structural_iterator_map_15 := map[string]any{}
		hx_structural_iterator_map_15["hasNext"] = func() bool {
			return hx_structural_iterator_14.__hx_this.hasNext()
		}
		hx_structural_iterator_map_15["next"] = func() any {
			return any(hx_structural_iterator_14.__hx_this.next())
		}
		return hx_structural_iterator_map_15
	}(New_SnapshotGenericIterator(hxrt.StringFromLiteral("g1"), hxrt.StringFromLiteral("g2"))))
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("generic="), func(hx_value_16 any) *string {
		if hx_value_16 == nil {
			var hx_zero_17 *string
			return hx_zero_17
		}
		return hx_value_16.(*string)
	}(genericConsumer.collect())))
	hxrt.Println(v_2)
	baseIterator := New_SnapshotSpecializedIterator([]*string{hxrt.StringFromLiteral("v")}).SnapshotBaseIterator
	virtualConsumer := New_SnapshotGenericConsumer(func(hx_structural_iterator_18 *SnapshotBaseIterator) map[string]any {
		hx_structural_iterator_map_19 := map[string]any{}
		hx_structural_iterator_map_19["hasNext"] = func() bool {
			return hx_structural_iterator_18.__hx_this.hasNext()
		}
		hx_structural_iterator_map_19["next"] = func() any {
			return any(hx_structural_iterator_18.__hx_this.next())
		}
		return hx_structural_iterator_map_19
	}(baseIterator))
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("virtual="), func(hx_value_20 any) *string {
		if hx_value_20 == nil {
			var hx_zero_21 *string
			return hx_zero_21
		}
		return hx_value_20.(*string)
	}(virtualConsumer.collect())))
	hxrt.Println(v_3)
}

func mark(label *string) *string {
	hx_arr_22 := events
	hx_arr_22 = append(hx_arr_22, label)
	events = hx_arr_22
	return label
}

type I_SnapshotBaseIterator interface {
	hasNext() bool
	next() *string
}

type SnapshotBaseIterator struct {
	__hx_this I_SnapshotBaseIterator
	values    []*string
	index     int
}

func New_SnapshotBaseIterator(values []*string) *SnapshotBaseIterator {
	self := &SnapshotBaseIterator{}
	self.__hx_this = self
	self.index = 0
	self.values = values
	return self
}

func (self *SnapshotBaseIterator) hasNext() bool {
	return (self.index < len(self.values))
}

func (self *SnapshotBaseIterator) next() *string {
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("base:"), self.values[func() int {
		hx_post_23 := self.index
		self.index = int(int32((self.index + 1)))
		return hx_post_23
	}()])
}

type I_SnapshotGenericConsumer interface {
	collect() *string
}

type SnapshotGenericConsumer struct {
	__hx_this I_SnapshotGenericConsumer
	iterator  map[string]any
}

func New_SnapshotGenericConsumer(iterator map[string]any) *SnapshotGenericConsumer {
	self := &SnapshotGenericConsumer{}
	self.__hx_this = self
	self.iterator = iterator
	return self
}

func (self *SnapshotGenericConsumer) collect() *string {
	values := []any{}
	for func(hx_obj_24 map[string]any) func() bool {
		hx_field_25 := hx_obj_24["hasNext"]
		if hx_field_25 == nil {
			var hx_zero_26 func() bool
			return hx_zero_26
		}
		return hx_field_25.(func() bool)
	}(self.iterator)() {
		values = append(values, func(hx_obj_28 map[string]any) func() any {
			hx_field_29 := hx_obj_28["next"]
			if hx_field_29 == nil {
				var hx_zero_30 func() any
				return hx_zero_30
			}
			return hx_field_29.(func() any)
		}(self.iterator)())
	}
	return hxrt.StringJoinAny(values, hxrt.StringFromLiteral(","))
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
	var hx_if_32 any
	if func() int {
		hx_post_31 := self.index
		self.index = int(int32((self.index + 1)))
		return hx_post_31
	}() == 0 {
		hx_if_32 = self.first
	} else {
		hx_if_32 = self.second
	}
	return hx_if_32
}

type I_SnapshotIntConsumer interface {
	collect() *string
}

type SnapshotIntConsumer struct {
	__hx_this I_SnapshotIntConsumer
	iterator  map[string]any
}

func New_SnapshotIntConsumer(before *string, iterator map[string]any, after *string) *SnapshotIntConsumer {
	self := &SnapshotIntConsumer{}
	self.__hx_this = self
	self.iterator = iterator
	return self
}

func (self *SnapshotIntConsumer) collect() *string {
	values := []int{}
	for func(hx_obj_33 map[string]any) func() bool {
		hx_field_34 := hx_obj_33["hasNext"]
		if hx_field_34 == nil {
			var hx_zero_35 func() bool
			return hx_zero_35
		}
		return hx_field_34.(func() bool)
	}(self.iterator)() {
		values = append(values, func(hx_obj_37 map[string]any) func() int {
			hx_field_38 := hx_obj_37["next"]
			if hx_field_38 == nil {
				var hx_zero_39 func() int
				return hx_zero_39
			}
			return hx_field_38.(func() int)
		}(self.iterator)())
	}
	return hxrt.StringJoinAny(func(hx_sort_src_40 []int) []any {
		hx_sort_out_42 := make([]any, 0, len(hx_sort_src_40))
		for _, hx_sort_item_41 := range hx_sort_src_40 {
			hx_sort_out_42 = append(hx_sort_out_42, hx_sort_item_41)
		}
		return hx_sort_out_42
	}(values), hxrt.StringFromLiteral(","))
}

type I_SnapshotSpecializedIterator interface {
	hasNext() bool
	next() *string
}

type SnapshotSpecializedIterator struct {
	*SnapshotBaseIterator
	__hx_this I_SnapshotSpecializedIterator
}

func New_SnapshotSpecializedIterator(values []*string) *SnapshotSpecializedIterator {
	self := &SnapshotSpecializedIterator{}
	self.SnapshotBaseIterator = New_SnapshotBaseIterator(values)
	self.SnapshotBaseIterator.__hx_this = self
	self.__hx_this = self
	return self
}

func (self *SnapshotSpecializedIterator) next() *string {
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("special:"), self.values[func() int {
		hx_post_43 := self.index
		self.index = int(int32((self.index + 1)))
		return hx_post_43
	}()])
}
