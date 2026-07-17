package main

import "snapshot/hxrt"

func checkedGeneric(first any, second any) map[string]any {
	note(hxrt.StringFromLiteral("generic:effect"))
	return func(hx_structural_iterator_1 *SnapshotInlineGenericIterator) map[string]any {
		hx_structural_iterator_map_2 := map[string]any{}
		hx_structural_iterator_map_2["hasNext"] = func() bool {
			return hx_structural_iterator_1.__hx_this.hasNext()
		}
		hx_structural_iterator_map_2["next"] = func() any {
			return any(hx_structural_iterator_1.__hx_this.next())
		}
		return hx_structural_iterator_map_2
	}(New_SnapshotInlineGenericIterator(first, second))
}

func checkedVirtual(values *hxrt.Array) map[string]any {
	note(hxrt.StringFromLiteral("virtual:effect"))
	iterator := New_SnapshotInlineSpecializedStringIterator(values).SnapshotInlineBaseStringIterator
	return func(hx_structural_iterator_3 *SnapshotInlineBaseStringIterator) map[string]any {
		hx_structural_iterator_map_4 := map[string]any{}
		hx_structural_iterator_map_4["hasNext"] = func() bool {
			return hx_structural_iterator_3.__hx_this.hasNext()
		}
		hx_structural_iterator_map_4["next"] = func() *string {
			return hx_structural_iterator_3.__hx_this.next()
		}
		return hx_structural_iterator_map_4
	}(iterator)
}

func collect(iterator map[string]any) *string {
	values := hxrt.NewArray()
	for func(hx_obj_5 map[string]any) func() bool {
		hx_field_6 := hx_obj_5["hasNext"]
		if hx_field_6 == nil {
			var hx_zero_7 func() bool
			return hx_zero_7
		}
		return hx_field_6.(func() bool)
	}(iterator)() {
		values.Push(func(hx_obj_9 map[string]any) func() *string {
			hx_field_10 := hx_obj_9["next"]
			if hx_field_10 == nil {
				var hx_zero_11 func() *string
				return hx_zero_11
			}
			return hx_field_10.(func() *string)
		}(iterator)())
	}
	return hxrt.StringJoinAny(values.Values(), hxrt.StringFromLiteral(","))
}

var events *hxrt.Array = hxrt.NewArray()

func main() {
	note(hxrt.StringFromLiteral("generic:effect"))
	generic := func(hx_structural_iterator_12 *SnapshotInlineGenericIterator) map[string]any {
		hx_structural_iterator_map_13 := map[string]any{}
		hx_structural_iterator_map_13["hasNext"] = func() bool {
			return hx_structural_iterator_12.__hx_this.hasNext()
		}
		hx_structural_iterator_map_13["next"] = func() *string {
			return func(hx_value_14 any) *string {
				if hx_value_14 == nil {
					var hx_zero_15 *string
					return hx_zero_15
				}
				return hx_value_14.(*string)
			}(any(hx_structural_iterator_12.__hx_this.next()))
		}
		return hx_structural_iterator_map_13
	}(New_SnapshotInlineGenericIterator(hxrt.StringFromLiteral("g1"), hxrt.StringFromLiteral("g2")))
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("events-generic="), hxrt.StringJoinAny(events.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("generic="), collect(generic)))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("virtual="), collect(func() map[string]any {
		note(hxrt.StringFromLiteral("virtual:effect"))
		iterator := New_SnapshotInlineSpecializedStringIterator(hxrt.NewArray(hxrt.StringFromLiteral("v"))).SnapshotInlineBaseStringIterator
		return func(hx_structural_iterator_16 *SnapshotInlineBaseStringIterator) map[string]any {
			hx_structural_iterator_map_17 := map[string]any{}
			hx_structural_iterator_map_17["hasNext"] = func() bool {
				return hx_structural_iterator_16.__hx_this.hasNext()
			}
			hx_structural_iterator_map_17["next"] = func() *string {
				return hx_structural_iterator_16.__hx_this.next()
			}
			return hx_structural_iterator_map_17
		}(iterator)
	}())))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("events-final="), hxrt.StringJoinAny(events.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_3)
}

func note(event *string) {
	hx_arr_18 := events
	hx_arr_18.Push(event)
}

type I_SnapshotInlineBaseStringIterator interface {
	hasNext() bool
	next() *string
}

type SnapshotInlineBaseStringIterator struct {
	__hx_this I_SnapshotInlineBaseStringIterator
	values    *hxrt.Array
	index     int
}

func New_SnapshotInlineBaseStringIterator(values *hxrt.Array) *SnapshotInlineBaseStringIterator {
	self := &SnapshotInlineBaseStringIterator{}
	self.__hx_this = self
	self.index = 0
	note(hxrt.StringFromLiteral("virtual:new"))
	self.values = values
	return self
}

func (self *SnapshotInlineBaseStringIterator) hasNext() bool {
	return (self.index < self.values.Len())
}

func (self *SnapshotInlineBaseStringIterator) next() *string {
	return hxrt.StringConcatAny(hxrt.StringFromLiteral("base:"), self.values.Get(func() int {
		hx_post_22 := self.index
		self.index = int(int32((self.index + 1)))
		return hx_post_22
	}()))
}

type I_SnapshotInlineGenericIterator interface {
	hasNext() bool
	next() any
}

type SnapshotInlineGenericIterator struct {
	__hx_this I_SnapshotInlineGenericIterator
	first     any
	second    any
	index     int
}

func New_SnapshotInlineGenericIterator(first any, second any) *SnapshotInlineGenericIterator {
	self := &SnapshotInlineGenericIterator{}
	self.__hx_this = self
	self.index = 0
	note(hxrt.StringFromLiteral("generic:new"))
	self.first = first
	self.second = second
	return self
}

func (self *SnapshotInlineGenericIterator) hasNext() bool {
	return (self.index < 2)
}

func (self *SnapshotInlineGenericIterator) next() any {
	var hx_if_24 any
	if func() int {
		hx_post_23 := self.index
		self.index = int(int32((self.index + 1)))
		return hx_post_23
	}() == 0 {
		hx_if_24 = self.first
	} else {
		hx_if_24 = self.second
	}
	return hx_if_24
}

type I_SnapshotInlineSpecializedStringIterator interface {
	hasNext() bool
	next() *string
}

type SnapshotInlineSpecializedStringIterator struct {
	*SnapshotInlineBaseStringIterator
	__hx_this I_SnapshotInlineSpecializedStringIterator
}

func New_SnapshotInlineSpecializedStringIterator(values *hxrt.Array) *SnapshotInlineSpecializedStringIterator {
	self := &SnapshotInlineSpecializedStringIterator{}
	self.SnapshotInlineBaseStringIterator = New_SnapshotInlineBaseStringIterator(values)
	self.SnapshotInlineBaseStringIterator.__hx_this = self
	self.__hx_this = self
	return self
}

func (self *SnapshotInlineSpecializedStringIterator) next() *string {
	return hxrt.StringConcatAny(hxrt.StringFromLiteral("special:"), self.values.Get(func() int {
		hx_post_28 := self.index
		self.index = int(int32((self.index + 1)))
		return hx_post_28
	}()))
}
