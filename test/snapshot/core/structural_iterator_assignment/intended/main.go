package main

import "snapshot/hxrt"

func collectInts(iterator map[string]any) *string {
	values := []int{}
	for func(hx_obj_1 map[string]any) func() bool {
		hx_field_2 := hx_obj_1["hasNext"]
		if hx_field_2 == nil {
			var hx_zero_3 func() bool
			return hx_zero_3
		}
		return hx_field_2.(func() bool)
	}(iterator)() {
		values = append(values, func(hx_obj_5 map[string]any) func() int {
			hx_field_6 := hx_obj_5["next"]
			if hx_field_6 == nil {
				var hx_zero_7 func() int
				return hx_zero_7
			}
			return hx_field_6.(func() int)
		}(iterator)())
	}
	return hxrt.StringJoinAny(func(hx_sort_src_8 []int) []any {
		hx_sort_out_10 := make([]any, 0, len(hx_sort_src_8))
		for _, hx_sort_item_9 := range hx_sort_src_8 {
			hx_sort_out_10 = append(hx_sort_out_10, hx_sort_item_9)
		}
		return hx_sort_out_10
	}(values), hxrt.StringFromLiteral(","))
}

func collectStrings(iterator map[string]any) *string {
	values := []*string{}
	for func(hx_obj_11 map[string]any) func() bool {
		hx_field_12 := hx_obj_11["hasNext"]
		if hx_field_12 == nil {
			var hx_zero_13 func() bool
			return hx_zero_13
		}
		return hx_field_12.(func() bool)
	}(iterator)() {
		values = append(values, func(hx_obj_15 map[string]any) func() *string {
			hx_field_16 := hx_obj_15["next"]
			if hx_field_16 == nil {
				var hx_zero_17 func() *string
				return hx_zero_17
			}
			return hx_field_16.(func() *string)
		}(iterator)())
	}
	return hxrt.StringJoinAny(func(hx_sort_src_18 []*string) []any {
		hx_sort_out_20 := make([]any, 0, len(hx_sort_src_18))
		for _, hx_sort_item_19 := range hx_sort_src_18 {
			hx_sort_out_20 = append(hx_sort_out_20, hx_sort_item_19)
		}
		return hx_sort_out_20
	}(values), hxrt.StringFromLiteral(","))
}

func main() {
	arrayValues := []int{4, 5}
	arrayIterator := func() map[string]any {
		hx_structural_array_21 := arrayValues
		hx_structural_array_index_22 := 0
		hx_structural_iterator_map_23 := map[string]any{}
		hx_structural_iterator_map_23["hasNext"] = func() bool {
			return (hx_structural_array_index_22 < len(hx_structural_array_21))
		}
		hx_structural_iterator_map_23["next"] = func() int {
			hx_structural_array_value_24 := hx_structural_array_21[hx_structural_array_index_22]
			hx_structural_array_index_22 = (hx_structural_array_index_22 + 1)
			return hx_structural_array_value_24
		}
		return hx_structural_iterator_map_23
	}()
	arrayValues[0] = 8
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("array="), collectInts(arrayIterator)))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("arrayArgument="), collectInts(func() map[string]any {
		hx_structural_array_25 := []int{6, 7}
		hx_structural_array_index_26 := 0
		hx_structural_iterator_map_27 := map[string]any{}
		hx_structural_iterator_map_27["hasNext"] = func() bool {
			return (hx_structural_array_index_26 < len(hx_structural_array_25))
		}
		hx_structural_iterator_map_27["next"] = func() int {
			hx_structural_array_value_28 := hx_structural_array_25[hx_structural_array_index_26]
			hx_structural_array_index_26 = (hx_structural_array_index_26 + 1)
			return hx_structural_array_value_28
		}
		return hx_structural_iterator_map_27
	}())))
	hxrt.Println(v_1)
	genericIterator := makeGenericIterator()
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("generic="), collectStrings(genericIterator)))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("genericArgument="), collectStrings(func(hx_structural_iterator_29 *SnapshotGenericIterator) map[string]any {
		hx_structural_iterator_map_30 := map[string]any{}
		hx_structural_iterator_map_30["hasNext"] = func() bool {
			return hx_structural_iterator_29.__hx_this.hasNext()
		}
		hx_structural_iterator_map_30["next"] = func() *string {
			return func(hx_value_31 any) *string {
				if hx_value_31 == nil {
					var hx_zero_32 *string
					return hx_zero_32
				}
				return hx_value_31.(*string)
			}(any(hx_structural_iterator_29.__hx_this.next()))
		}
		return hx_structural_iterator_map_30
	}(New_SnapshotGenericIterator(hxrt.StringFromLiteral("c"), hxrt.StringFromLiteral("d"))))))
	hxrt.Println(v_3)
	baseIterator := New_SnapshotSpecializedIterator([]*string{hxrt.StringFromLiteral("z")}).SnapshotBaseIterator
	virtualIterator := func(hx_structural_iterator_33 *SnapshotBaseIterator) map[string]any {
		hx_structural_iterator_map_34 := map[string]any{}
		hx_structural_iterator_map_34["hasNext"] = func() bool {
			return hx_structural_iterator_33.__hx_this.hasNext()
		}
		hx_structural_iterator_map_34["next"] = func() *string {
			return hx_structural_iterator_33.__hx_this.next()
		}
		return hx_structural_iterator_map_34
	}(baseIterator)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("virtual="), collectStrings(virtualIterator)))
	hxrt.Println(v_4)
}

func makeGenericIterator() map[string]any {
	return func(hx_structural_iterator_35 *SnapshotGenericIterator) map[string]any {
		hx_structural_iterator_map_36 := map[string]any{}
		hx_structural_iterator_map_36["hasNext"] = func() bool {
			return hx_structural_iterator_35.__hx_this.hasNext()
		}
		hx_structural_iterator_map_36["next"] = func() *string {
			return func(hx_value_37 any) *string {
				if hx_value_37 == nil {
					var hx_zero_38 *string
					return hx_zero_38
				}
				return hx_value_37.(*string)
			}(any(hx_structural_iterator_35.__hx_this.next()))
		}
		return hx_structural_iterator_map_36
	}(New_SnapshotGenericIterator(hxrt.StringFromLiteral("u"), hxrt.StringFromLiteral("v")))
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
	self.values = values
	self.index = 0
	return self
}

func (self *SnapshotBaseIterator) hasNext() bool {
	return (self.index < len(self.values))
}

func (self *SnapshotBaseIterator) next() *string {
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("base:"), self.values[func() int {
		hx_post_39 := self.index
		self.index = int(int32((self.index + 1)))
		return hx_post_39
	}()])
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
	self.first = first
	self.second = second
	self.index = 0
	return self
}

func (self *SnapshotGenericIterator) hasNext() bool {
	return (self.index < 2)
}

func (self *SnapshotGenericIterator) next() any {
	var hx_if_41 any
	if func() int {
		hx_post_40 := self.index
		self.index = int(int32((self.index + 1)))
		return hx_post_40
	}() == 0 {
		hx_if_41 = self.first
	} else {
		hx_if_41 = self.second
	}
	return hx_if_41
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
		hx_post_42 := self.index
		self.index = int(int32((self.index + 1)))
		return hx_post_42
	}()])
}
