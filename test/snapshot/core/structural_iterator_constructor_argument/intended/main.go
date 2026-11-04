package main

import "snapshot/hxrt"

func checkedIterator(values *hxrt.Array) hxrt.Anonymous {
	hx_arr_1 := events
	hx_arr_1.Push(hxrt.StringFromLiteral("iterator"))
	return func(hx_structural_source_2 *haxe__iterators__ArrayIterator) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_2, func(name string) hxrt.AnonymousField {
			switch name {
			case "array":
				return hxrt.AnonymousField{Value: hx_structural_source_2.array, Found: true}
			case "current":
				return hxrt.AnonymousField{Value: hx_structural_source_2.current, Found: true}
			case "hasNext":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
					hx_receiver_5 := hx_structural_source_2
					return hxrt.BoundFunction(func() bool {
						return hx_receiver_5.__hx_this.hasNext()
					}, func() any {
						return hx_receiver_5.__hx_this
					}, "hasNext")
				}(), Found: true}
			case "next":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
					hx_receiver_6 := hx_structural_source_2
					return hxrt.BoundFunction(func() any {
						return hx_receiver_6.__hx_this.next()
					}, func() any {
						return hx_receiver_6.__hx_this
					}, "next")
				}(), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			case "current":
				hx_structural_source_2.current = func(hx_value_3 any) int {
					if hx_value_3 == nil {
						var hx_zero_4 int
						return hx_zero_4
					}
					return hx_value_3.(int)
				}(incoming)
				return true
			}
			return false
		})
	}(New_haxe__iterators__ArrayIterator(values))
}

var events *hxrt.Array = hxrt.NewArray()

func main() {
	arrayValues := hxrt.NewArray(1, 2)
	arrayConsumer := New_SnapshotIntConsumer(mark(hxrt.StringFromLiteral("before")), func() hxrt.Anonymous {
		hx_arr_12 := events
		hx_arr_12.Push(hxrt.StringFromLiteral("iterator"))
		return func(hx_structural_source_7 *haxe__iterators__ArrayIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_7, func(name string) hxrt.AnonymousField {
				switch name {
				case "array":
					return hxrt.AnonymousField{Value: hx_structural_source_7.array, Found: true}
				case "current":
					return hxrt.AnonymousField{Value: hx_structural_source_7.current, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_10 := hx_structural_source_7
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_10.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_10.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_11 := hx_structural_source_7
						return hxrt.BoundFunction(func() any {
							return hx_receiver_11.__hx_this.next()
						}, func() any {
							return hx_receiver_11.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "current":
					hx_structural_source_7.current = func(hx_value_8 any) int {
						if hx_value_8 == nil {
							var hx_zero_9 int
							return hx_zero_9
						}
						return hx_value_8.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(New_haxe__iterators__ArrayIterator(arrayValues))
	}(), mark(hxrt.StringFromLiteral("after")))
	hx_array_target_13 := arrayValues
	hx_array_index_14 := 0
	hx_array_target_13.Set(hx_array_index_14, 9)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("order="), hxrt.StringJoinAny(events.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("array="), arrayConsumer.collect()))
	hxrt.Println(v_1)
	genericConsumer := New_SnapshotGenericConsumer(func(hx_structural_source_15 *SnapshotGenericIterator) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_15, func(name string) hxrt.AnonymousField {
			switch name {
			case "first":
				return hxrt.AnonymousField{Value: hx_structural_source_15.first, Found: true}
			case "second":
				return hxrt.AnonymousField{Value: hx_structural_source_15.second, Found: true}
			case "index":
				return hxrt.AnonymousField{Value: hx_structural_source_15.index, Found: true}
			case "hasNext":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
					hx_receiver_18 := hx_structural_source_15
					return hxrt.BoundFunction(func() bool {
						return hx_receiver_18.__hx_this.hasNext()
					}, func() any {
						return hx_receiver_18.__hx_this
					}, "hasNext")
				}(), Found: true}
			case "next":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
					hx_receiver_19 := hx_structural_source_15
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
			case "index":
				hx_structural_source_15.index = func(hx_value_16 any) int {
					if hx_value_16 == nil {
						var hx_zero_17 int
						return hx_zero_17
					}
					return hx_value_16.(int)
				}(incoming)
				return true
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
			case "values":
				return hxrt.AnonymousField{Value: hx_structural_source_22.values, Found: true}
			case "index":
				return hxrt.AnonymousField{Value: hx_structural_source_22.index, Found: true}
			case "hasNext":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
					hx_receiver_25 := hx_structural_source_22
					return hxrt.BoundFunction(func() bool {
						return hx_receiver_25.__hx_this.hasNext()
					}, func() any {
						return hx_receiver_25.__hx_this
					}, "hasNext")
				}(), Found: true}
			case "next":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
					hx_receiver_26 := hx_structural_source_22
					return hxrt.BoundFunction(func() *string {
						return hx_receiver_26.__hx_this.next()
					}, func() any {
						return hx_receiver_26.__hx_this
					}, "next")
				}(), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			case "index":
				hx_structural_source_22.index = func(hx_value_23 any) int {
					if hx_value_23 == nil {
						var hx_zero_24 int
						return hx_zero_24
					}
					return hx_value_23.(int)
				}(incoming)
				return true
			}
			return false
		})
	}(baseIterator))
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("virtual="), func(hx_value_27 any) *string {
		if hx_value_27 == nil {
			var hx_zero_28 *string
			return hx_zero_28
		}
		return hx_value_27.(*string)
	}(virtualConsumer.collect())))
	hxrt.Println(v_3)
}

func mark(label *string) *string {
	hx_arr_29 := events
	hx_arr_29.Push(label)
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
	return hx_init_SnapshotBaseIterator(hx_alloc_SnapshotBaseIterator(nil), values)
}

func hx_alloc_SnapshotBaseIterator(owner any) *SnapshotBaseIterator {
	self := &SnapshotBaseIterator{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_SnapshotBaseIterator)
	}
	return self
}

func hx_init_SnapshotBaseIterator(self *SnapshotBaseIterator, values *hxrt.Array) *SnapshotBaseIterator {
	self.index = 0
	self.values = values
	return self
}

func (self *SnapshotBaseIterator) hasNext() bool {
	return (self.index < self.values.Len())
}

func (self *SnapshotBaseIterator) next() *string {
	return hxrt.StringConcatAny(hxrt.StringFromLiteral("base:"), self.values.Get(func() int {
		hx_post_33 := self.index
		self.index = int(int32((self.index + 1)))
		return hx_post_33
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
	return hx_init_SnapshotGenericConsumer(hx_alloc_SnapshotGenericConsumer(nil), iterator)
}

func hx_alloc_SnapshotGenericConsumer(owner any) *SnapshotGenericConsumer {
	self := &SnapshotGenericConsumer{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_SnapshotGenericConsumer)
	}
	return self
}

func hx_init_SnapshotGenericConsumer(self *SnapshotGenericConsumer, iterator hxrt.Anonymous) *SnapshotGenericConsumer {
	self.iterator = iterator
	return self
}

func (self *SnapshotGenericConsumer) collect() *string {
	values := hxrt.NewArray()
	for func() bool {
		var hx_callable_34 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(self.iterator, "hasNext"))
		return hxrt.FunctionCall(hx_callable_34)()
	}() {
		values.Push(func() any {
			var hx_callable_36 *hxrt.Function[func() any] = hx_recover_callable_1(hxrt.AnonymousGet(self.iterator, "next"))
			return hxrt.FunctionCall(hx_callable_36)()
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
	return hx_init_SnapshotGenericIterator(hx_alloc_SnapshotGenericIterator(nil), first, second)
}

func hx_alloc_SnapshotGenericIterator(owner any) *SnapshotGenericIterator {
	self := &SnapshotGenericIterator{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_SnapshotGenericIterator)
	}
	return self
}

func hx_init_SnapshotGenericIterator(self *SnapshotGenericIterator, first any, second any) *SnapshotGenericIterator {
	self.index = 0
	self.first = first
	self.second = second
	return self
}

func (self *SnapshotGenericIterator) hasNext() bool {
	return (self.index < 2)
}

func (self *SnapshotGenericIterator) next() any {
	var hx_if_38 any
	if func() int {
		hx_post_37 := self.index
		self.index = int(int32((self.index + 1)))
		return hx_post_37
	}() == 0 {
		hx_if_38 = self.first
	} else {
		hx_if_38 = self.second
	}
	return hx_if_38
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
	return hx_init_SnapshotIntConsumer(hx_alloc_SnapshotIntConsumer(nil), before, iterator, after)
}

func hx_alloc_SnapshotIntConsumer(owner any) *SnapshotIntConsumer {
	self := &SnapshotIntConsumer{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_SnapshotIntConsumer)
	}
	return self
}

func hx_init_SnapshotIntConsumer(self *SnapshotIntConsumer, before *string, iterator hxrt.Anonymous, after *string) *SnapshotIntConsumer {
	self.iterator = iterator
	return self
}

func (self *SnapshotIntConsumer) collect() *string {
	values := hxrt.NewArray()
	for func() bool {
		var hx_callable_39 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(self.iterator, "hasNext"))
		return hxrt.FunctionCall(hx_callable_39)()
	}() {
		values.Push(func() int {
			var hx_callable_41 *hxrt.Function[func() int] = hx_recover_callable_2(hxrt.AnonymousGet(self.iterator, "next"))
			return hxrt.FunctionCall(hx_callable_41)()
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
	return hx_init_SnapshotSpecializedIterator(hx_alloc_SnapshotSpecializedIterator(nil), values)
}

func hx_alloc_SnapshotSpecializedIterator(owner any) *SnapshotSpecializedIterator {
	self := &SnapshotSpecializedIterator{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_SnapshotSpecializedIterator)
	}
	self.SnapshotBaseIterator = hx_alloc_SnapshotBaseIterator(self.__hx_this)
	return self
}

func hx_init_SnapshotSpecializedIterator(self *SnapshotSpecializedIterator, values *hxrt.Array) *SnapshotSpecializedIterator {
	hx_init_SnapshotBaseIterator(self.SnapshotBaseIterator, values)
	return self
}

func (self *SnapshotSpecializedIterator) next() *string {
	return hxrt.StringConcatAny(hxrt.StringFromLiteral("special:"), self.values.Get(func() int {
		hx_post_45 := self.index
		self.index = int(int32((self.index + 1)))
		return hx_post_45
	}()))
}

func hx_recover_callable_0(value any) *hxrt.Function[func() bool] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() bool]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return source
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_1 any) bool {
				if hx_value_1 == nil {
					var hx_zero_2 bool
					return hx_zero_2
				}
				return hx_value_1.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_3 any) bool {
				if hx_value_3 == nil {
					var hx_zero_4 bool
					return hx_zero_4
				}
				return hx_value_3.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_1(value any) *hxrt.Function[func() any] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() any]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() any]:
		return source
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_2(value any) *hxrt.Function[func() int] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() int]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_5 any) int {
				if hx_value_5 == nil {
					var hx_zero_6 int
					return hx_zero_6
				}
				return hx_value_5.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_7 any) int {
				if hx_value_7 == nil {
					var hx_zero_8 int
					return hx_zero_8
				}
				return hx_value_7.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_9 any) int {
				if hx_value_9 == nil {
					var hx_zero_10 int
					return hx_zero_10
				}
				return hx_value_9.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}
