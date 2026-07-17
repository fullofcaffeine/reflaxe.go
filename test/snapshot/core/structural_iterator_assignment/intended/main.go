package main

import "snapshot/hxrt"

func collectInts(iterator map[string]any) *string {
	values := hxrt.NewArray()
	for func(hx_obj_1 map[string]any) func() bool {
		hx_field_2 := hx_obj_1["hasNext"]
		if hx_field_2 == nil {
			var hx_zero_3 func() bool
			return hx_zero_3
		}
		return hx_field_2.(func() bool)
	}(iterator)() {
		values.Push(func(hx_obj_5 map[string]any) func() int {
			hx_field_6 := hx_obj_5["next"]
			if hx_field_6 == nil {
				var hx_zero_7 func() int
				return hx_zero_7
			}
			return hx_field_6.(func() int)
		}(iterator)())
	}
	return hxrt.StringJoinAny(values.Values(), hxrt.StringFromLiteral(","))
}

func collectStrings(iterator map[string]any) *string {
	values := hxrt.NewArray()
	for func(hx_obj_8 map[string]any) func() bool {
		hx_field_9 := hx_obj_8["hasNext"]
		if hx_field_9 == nil {
			var hx_zero_10 func() bool
			return hx_zero_10
		}
		return hx_field_9.(func() bool)
	}(iterator)() {
		values.Push(func(hx_obj_12 map[string]any) func() *string {
			hx_field_13 := hx_obj_12["next"]
			if hx_field_13 == nil {
				var hx_zero_14 func() *string
				return hx_zero_14
			}
			return hx_field_13.(func() *string)
		}(iterator)())
	}
	return hxrt.StringJoinAny(values.Values(), hxrt.StringFromLiteral(","))
}

func main() {
	arrayValues := hxrt.NewArray(4, 5)
	arrayIterator := func() map[string]any {
		hx_structural_array_15 := arrayValues
		hx_structural_array_index_16 := 0
		hx_structural_iterator_map_17 := map[string]any{}
		hx_structural_iterator_map_17["hasNext"] = func() bool {
			return (hx_structural_array_index_16 < hx_structural_array_15.Len())
		}
		hx_structural_iterator_map_17["next"] = func() int {
			hx_structural_array_value_18 := hx_structural_array_15.Get(hx_structural_array_index_16)
			hx_structural_array_index_16 = (hx_structural_array_index_16 + 1)
			return func(hx_value_19 any) int {
				if hx_value_19 == nil {
					var hx_zero_20 int
					return hx_zero_20
				}
				return hx_value_19.(int)
			}(any(hx_structural_array_value_18))
		}
		return hx_structural_iterator_map_17
	}()
	hx_array_target_21 := arrayValues
	hx_array_index_22 := 0
	hx_array_target_21.Set(hx_array_index_22, 8)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("array="), collectInts(arrayIterator)))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("arrayArgument="), collectInts(func() map[string]any {
		hx_structural_array_23 := hxrt.NewArray(6, 7)
		hx_structural_array_index_24 := 0
		hx_structural_iterator_map_25 := map[string]any{}
		hx_structural_iterator_map_25["hasNext"] = func() bool {
			return (hx_structural_array_index_24 < hx_structural_array_23.Len())
		}
		hx_structural_iterator_map_25["next"] = func() int {
			hx_structural_array_value_26 := hx_structural_array_23.Get(hx_structural_array_index_24)
			hx_structural_array_index_24 = (hx_structural_array_index_24 + 1)
			return func(hx_value_27 any) int {
				if hx_value_27 == nil {
					var hx_zero_28 int
					return hx_zero_28
				}
				return hx_value_27.(int)
			}(any(hx_structural_array_value_26))
		}
		return hx_structural_iterator_map_25
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
	baseIterator := New_SnapshotSpecializedIterator(hxrt.NewArray(hxrt.StringFromLiteral("z"))).SnapshotBaseIterator
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
	values    *hxrt.Array
	index     int
}

func New_SnapshotBaseIterator(values *hxrt.Array) *SnapshotBaseIterator {
	self := &SnapshotBaseIterator{}
	self.__hx_this = self
	self.values = values
	self.index = 0
	return self
}

func (self *SnapshotBaseIterator) hasNext() bool {
	return (self.index < self.values.Len())
}

func (self *SnapshotBaseIterator) next() *string {
	return hxrt.StringConcatAny(hxrt.StringFromLiteral("base:"), self.values.Get(func() int {
		hx_post_42 := self.index
		self.index = int(int32((self.index + 1)))
		return hx_post_42
	}()))
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
	var hx_if_44 any
	if func() int {
		hx_post_43 := self.index
		self.index = int(int32((self.index + 1)))
		return hx_post_43
	}() == 0 {
		hx_if_44 = self.first
	} else {
		hx_if_44 = self.second
	}
	return hx_if_44
}

type I_SnapshotSpecializedIterator interface {
	hasNext() bool
	next() *string
}

type SnapshotSpecializedIterator struct {
	*SnapshotBaseIterator
	__hx_this I_SnapshotSpecializedIterator
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
		hx_post_48 := self.index
		self.index = int(int32((self.index + 1)))
		return hx_post_48
	}()))
}
