package main

import "snapshot/hxrt"

func checkedGeneric(first any, second any) hxrt.Anonymous {
	note(hxrt.StringFromLiteral("generic:effect"))
	return func(hx_structural_source_1 *SnapshotInlineGenericIterator) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
			switch name {
			case "first":
				return hxrt.AnonymousField{Value: hx_structural_source_1.first, Found: true}
			case "second":
				return hxrt.AnonymousField{Value: hx_structural_source_1.second, Found: true}
			case "index":
				return hxrt.AnonymousField{Value: hx_structural_source_1.index, Found: true}
			case "hasNext":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
					hx_receiver_4 := hx_structural_source_1
					return hxrt.BoundFunction(func() bool {
						return hx_receiver_4.__hx_this.hasNext()
					}, func() any {
						return hx_receiver_4.__hx_this
					}, "hasNext")
				}(), Found: true}
			case "next":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
					hx_receiver_5 := hx_structural_source_1
					return hxrt.BoundFunction(func() any {
						return hx_receiver_5.__hx_this.next()
					}, func() any {
						return hx_receiver_5.__hx_this
					}, "next")
				}(), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			case "index":
				hx_structural_source_1.index = func(hx_value_2 any) int {
					if hx_value_2 == nil {
						var hx_zero_3 int
						return hx_zero_3
					}
					return hx_value_2.(int)
				}(incoming)
				return true
			}
			return false
		})
	}(New_SnapshotInlineGenericIterator(first, second))
}

func checkedVirtual(values *hxrt.Array) hxrt.Anonymous {
	note(hxrt.StringFromLiteral("virtual:effect"))
	iterator := func(hx_value *SnapshotInlineSpecializedStringIterator) *SnapshotInlineBaseStringIterator {
		if hx_value == nil {
			return nil
		}
		return hx_value.SnapshotInlineBaseStringIterator
	}(New_SnapshotInlineSpecializedStringIterator(values))
	return func(hx_structural_source_6 *SnapshotInlineBaseStringIterator) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_6, func(name string) hxrt.AnonymousField {
			switch name {
			case "values":
				return hxrt.AnonymousField{Value: hx_structural_source_6.values, Found: true}
			case "index":
				return hxrt.AnonymousField{Value: hx_structural_source_6.index, Found: true}
			case "hasNext":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
					hx_receiver_9 := hx_structural_source_6
					return hxrt.BoundFunction(func() bool {
						return hx_receiver_9.__hx_this.hasNext()
					}, func() any {
						return hx_receiver_9.__hx_this
					}, "hasNext")
				}(), Found: true}
			case "next":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
					hx_receiver_10 := hx_structural_source_6
					return hxrt.BoundFunction(func() *string {
						return hx_receiver_10.__hx_this.next()
					}, func() any {
						return hx_receiver_10.__hx_this
					}, "next")
				}(), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			case "index":
				hx_structural_source_6.index = func(hx_value_7 any) int {
					if hx_value_7 == nil {
						var hx_zero_8 int
						return hx_zero_8
					}
					return hx_value_7.(int)
				}(incoming)
				return true
			}
			return false
		})
	}(iterator)
}

func collect(iterator hxrt.Anonymous) *string {
	values := hxrt.NewArray()
	for func() bool {
		var hx_callable_11 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(iterator, "hasNext"))
		return hxrt.FunctionCall(hx_callable_11)()
	}() {
		values.Push(func() *string {
			var hx_callable_13 *hxrt.Function[func() *string] = hx_recover_callable_1(hxrt.AnonymousGet(iterator, "next"))
			return hxrt.FunctionCall(hx_callable_13)()
		}())
	}
	return hxrt.StringJoinAny(values.Values(), hxrt.StringFromLiteral(","))
}

var events *hxrt.Array = hxrt.NewArray()

func main() {
	note(hxrt.StringFromLiteral("generic:effect"))
	generic := func(hx_structural_source_14 *SnapshotInlineGenericIterator) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_14, func(name string) hxrt.AnonymousField {
			switch name {
			case "first":
				return hxrt.AnonymousField{Value: hx_structural_source_14.first, Found: true}
			case "second":
				return hxrt.AnonymousField{Value: hx_structural_source_14.second, Found: true}
			case "index":
				return hxrt.AnonymousField{Value: hx_structural_source_14.index, Found: true}
			case "hasNext":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
					hx_receiver_17 := hx_structural_source_14
					return hxrt.BoundFunction(func() bool {
						return hx_receiver_17.__hx_this.hasNext()
					}, func() any {
						return hx_receiver_17.__hx_this
					}, "hasNext")
				}(), Found: true}
			case "next":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
					hx_receiver_18 := hx_structural_source_14
					return hxrt.BoundFunction(func() any {
						return hx_receiver_18.__hx_this.next()
					}, func() any {
						return hx_receiver_18.__hx_this
					}, "next")
				}(), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			case "index":
				hx_structural_source_14.index = func(hx_value_15 any) int {
					if hx_value_15 == nil {
						var hx_zero_16 int
						return hx_zero_16
					}
					return hx_value_15.(int)
				}(incoming)
				return true
			}
			return false
		})
	}(New_SnapshotInlineGenericIterator(hxrt.StringFromLiteral("g1"), hxrt.StringFromLiteral("g2")))
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("events-generic="), hxrt.StringJoinAny(events.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("generic="), collect(generic)))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("virtual="), collect(func() hxrt.Anonymous {
		note(hxrt.StringFromLiteral("virtual:effect"))
		iterator := func(hx_value *SnapshotInlineSpecializedStringIterator) *SnapshotInlineBaseStringIterator {
			if hx_value == nil {
				return nil
			}
			return hx_value.SnapshotInlineBaseStringIterator
		}(New_SnapshotInlineSpecializedStringIterator(hxrt.NewArray(hxrt.StringFromLiteral("v"))))
		return func(hx_structural_source_19 *SnapshotInlineBaseStringIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_19, func(name string) hxrt.AnonymousField {
				switch name {
				case "values":
					return hxrt.AnonymousField{Value: hx_structural_source_19.values, Found: true}
				case "index":
					return hxrt.AnonymousField{Value: hx_structural_source_19.index, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_22 := hx_structural_source_19
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_22.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_22.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_23 := hx_structural_source_19
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_23.__hx_this.next()
						}, func() any {
							return hx_receiver_23.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "index":
					hx_structural_source_19.index = func(hx_value_20 any) int {
						if hx_value_20 == nil {
							var hx_zero_21 int
							return hx_zero_21
						}
						return hx_value_20.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(iterator)
	}())))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("events-final="), hxrt.StringJoinAny(events.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_3)
}

func note(event *string) {
	hx_arr_24 := events
	hx_arr_24.Push(event)
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

func (self *SnapshotInlineBaseStringIterator) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_SnapshotInlineBaseStringIterator(values *hxrt.Array) *SnapshotInlineBaseStringIterator {
	return hx_init_SnapshotInlineBaseStringIterator(hx_alloc_SnapshotInlineBaseStringIterator(nil), values)
}

func hx_alloc_SnapshotInlineBaseStringIterator(owner any) *SnapshotInlineBaseStringIterator {
	self := &SnapshotInlineBaseStringIterator{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_SnapshotInlineBaseStringIterator)
	}
	return self
}

func hx_init_SnapshotInlineBaseStringIterator(self *SnapshotInlineBaseStringIterator, values *hxrt.Array) *SnapshotInlineBaseStringIterator {
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
		hx_post_28 := self.index
		self.index = int(int32((self.index + 1)))
		return hx_post_28
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

func (self *SnapshotInlineGenericIterator) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_SnapshotInlineGenericIterator(first any, second any) *SnapshotInlineGenericIterator {
	return hx_init_SnapshotInlineGenericIterator(hx_alloc_SnapshotInlineGenericIterator(nil), first, second)
}

func hx_alloc_SnapshotInlineGenericIterator(owner any) *SnapshotInlineGenericIterator {
	self := &SnapshotInlineGenericIterator{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_SnapshotInlineGenericIterator)
	}
	return self
}

func hx_init_SnapshotInlineGenericIterator(self *SnapshotInlineGenericIterator, first any, second any) *SnapshotInlineGenericIterator {
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
	var hx_if_30 any
	if func() int {
		hx_post_29 := self.index
		self.index = int(int32((self.index + 1)))
		return hx_post_29
	}() == 0 {
		hx_if_30 = self.first
	} else {
		hx_if_30 = self.second
	}
	return hx_if_30
}

type I_SnapshotInlineSpecializedStringIterator interface {
	hasNext() bool
	next() *string
}

type SnapshotInlineSpecializedStringIterator struct {
	*SnapshotInlineBaseStringIterator
	__hx_this I_SnapshotInlineSpecializedStringIterator
}

func (self *SnapshotInlineSpecializedStringIterator) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_SnapshotInlineSpecializedStringIterator(values *hxrt.Array) *SnapshotInlineSpecializedStringIterator {
	return hx_init_SnapshotInlineSpecializedStringIterator(hx_alloc_SnapshotInlineSpecializedStringIterator(nil), values)
}

func hx_alloc_SnapshotInlineSpecializedStringIterator(owner any) *SnapshotInlineSpecializedStringIterator {
	self := &SnapshotInlineSpecializedStringIterator{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_SnapshotInlineSpecializedStringIterator)
	}
	self.SnapshotInlineBaseStringIterator = hx_alloc_SnapshotInlineBaseStringIterator(self.__hx_this)
	return self
}

func hx_init_SnapshotInlineSpecializedStringIterator(self *SnapshotInlineSpecializedStringIterator, values *hxrt.Array) *SnapshotInlineSpecializedStringIterator {
	hx_init_SnapshotInlineBaseStringIterator(self.SnapshotInlineBaseStringIterator, values)
	return self
}

func (self *SnapshotInlineSpecializedStringIterator) next() *string {
	return hxrt.StringConcatAny(hxrt.StringFromLiteral("special:"), self.values.Get(func() int {
		hx_post_34 := self.index
		self.index = int(int32((self.index + 1)))
		return hx_post_34
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

func hx_recover_callable_1(value any) *hxrt.Function[func() *string] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() *string]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_5 any) *string {
				if hx_value_5 == nil {
					var hx_zero_6 *string
					return hx_zero_6
				}
				return hx_value_5.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_7 any) *string {
				if hx_value_7 == nil {
					var hx_zero_8 *string
					return hx_zero_8
				}
				return hx_value_7.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *string]:
		return source
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}
