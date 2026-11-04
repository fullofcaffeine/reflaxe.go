package main

import "snapshot/hxrt"

func collectInts(iterator hxrt.Anonymous) *string {
	values := hxrt.NewArray()
	for func() bool {
		var hx_callable_1 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(iterator, "hasNext"))
		return hxrt.FunctionCall(hx_callable_1)()
	}() {
		values.Push(func() int {
			var hx_callable_3 *hxrt.Function[func() int] = hx_recover_callable_1(hxrt.AnonymousGet(iterator, "next"))
			return hxrt.FunctionCall(hx_callable_3)()
		}())
	}
	return hxrt.StringJoinAny(values.Values(), hxrt.StringFromLiteral(","))
}

func collectStrings(iterator hxrt.Anonymous) *string {
	values := hxrt.NewArray()
	for func() bool {
		var hx_callable_4 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(iterator, "hasNext"))
		return hxrt.FunctionCall(hx_callable_4)()
	}() {
		values.Push(func() *string {
			var hx_callable_6 *hxrt.Function[func() *string] = hx_recover_callable_2(hxrt.AnonymousGet(iterator, "next"))
			return hxrt.FunctionCall(hx_callable_6)()
		}())
	}
	return hxrt.StringJoinAny(values.Values(), hxrt.StringFromLiteral(","))
}

func main() {
	arrayValues := hxrt.NewArray(4, 5)
	arrayIterator := func(hx_structural_source_7 *haxe__iterators__ArrayIterator) hxrt.Anonymous {
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
	hx_array_target_12 := arrayValues
	hx_array_index_13 := 0
	hx_array_target_12.Set(hx_array_index_13, 8)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("array="), collectInts(arrayIterator)))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("arrayArgument="), collectInts(func(hx_structural_source_14 *haxe__iterators__ArrayIterator) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_14, func(name string) hxrt.AnonymousField {
			switch name {
			case "array":
				return hxrt.AnonymousField{Value: hx_structural_source_14.array, Found: true}
			case "current":
				return hxrt.AnonymousField{Value: hx_structural_source_14.current, Found: true}
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
			case "current":
				hx_structural_source_14.current = func(hx_value_15 any) int {
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
	}(New_haxe__iterators__ArrayIterator(hxrt.NewArray(6, 7))))))
	hxrt.Println(v_1)
	genericIterator := makeGenericIterator()
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("generic="), collectStrings(genericIterator)))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("genericArgument="), collectStrings(func(hx_structural_source_19 *SnapshotGenericIterator) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_19, func(name string) hxrt.AnonymousField {
			switch name {
			case "first":
				return hxrt.AnonymousField{Value: hx_structural_source_19.first, Found: true}
			case "second":
				return hxrt.AnonymousField{Value: hx_structural_source_19.second, Found: true}
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
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
					hx_receiver_23 := hx_structural_source_19
					return hxrt.BoundFunction(func() any {
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
	}(New_SnapshotGenericIterator(hxrt.StringFromLiteral("c"), hxrt.StringFromLiteral("d"))))))
	hxrt.Println(v_3)
	baseIterator := func(hx_value *SnapshotSpecializedIterator) *SnapshotBaseIterator {
		if hx_value == nil {
			return nil
		}
		return hx_value.SnapshotBaseIterator
	}(New_SnapshotSpecializedIterator(hxrt.NewArray(hxrt.StringFromLiteral("z"))))
	virtualIterator := func(hx_structural_source_24 *SnapshotBaseIterator) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_24, func(name string) hxrt.AnonymousField {
			switch name {
			case "values":
				return hxrt.AnonymousField{Value: hx_structural_source_24.values, Found: true}
			case "index":
				return hxrt.AnonymousField{Value: hx_structural_source_24.index, Found: true}
			case "hasNext":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
					hx_receiver_27 := hx_structural_source_24
					return hxrt.BoundFunction(func() bool {
						return hx_receiver_27.__hx_this.hasNext()
					}, func() any {
						return hx_receiver_27.__hx_this
					}, "hasNext")
				}(), Found: true}
			case "next":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
					hx_receiver_28 := hx_structural_source_24
					return hxrt.BoundFunction(func() *string {
						return hx_receiver_28.__hx_this.next()
					}, func() any {
						return hx_receiver_28.__hx_this
					}, "next")
				}(), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			case "index":
				hx_structural_source_24.index = func(hx_value_25 any) int {
					if hx_value_25 == nil {
						var hx_zero_26 int
						return hx_zero_26
					}
					return hx_value_25.(int)
				}(incoming)
				return true
			}
			return false
		})
	}(baseIterator)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("virtual="), collectStrings(virtualIterator)))
	hxrt.Println(v_4)
}

func makeGenericIterator() hxrt.Anonymous {
	return func(hx_structural_source_29 *SnapshotGenericIterator) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_29, func(name string) hxrt.AnonymousField {
			switch name {
			case "first":
				return hxrt.AnonymousField{Value: hx_structural_source_29.first, Found: true}
			case "second":
				return hxrt.AnonymousField{Value: hx_structural_source_29.second, Found: true}
			case "index":
				return hxrt.AnonymousField{Value: hx_structural_source_29.index, Found: true}
			case "hasNext":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
					hx_receiver_32 := hx_structural_source_29
					return hxrt.BoundFunction(func() bool {
						return hx_receiver_32.__hx_this.hasNext()
					}, func() any {
						return hx_receiver_32.__hx_this
					}, "hasNext")
				}(), Found: true}
			case "next":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
					hx_receiver_33 := hx_structural_source_29
					return hxrt.BoundFunction(func() any {
						return hx_receiver_33.__hx_this.next()
					}, func() any {
						return hx_receiver_33.__hx_this
					}, "next")
				}(), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			case "index":
				hx_structural_source_29.index = func(hx_value_30 any) int {
					if hx_value_30 == nil {
						var hx_zero_31 int
						return hx_zero_31
					}
					return hx_value_30.(int)
				}(incoming)
				return true
			}
			return false
		})
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
	self.values = values
	self.index = 0
	return self
}

func (self *SnapshotBaseIterator) hasNext() bool {
	return (self.index < self.values.Len())
}

func (self *SnapshotBaseIterator) next() *string {
	return hxrt.StringConcatAny(hxrt.StringFromLiteral("base:"), self.values.Get(func() int {
		hx_post_37 := self.index
		self.index = int(int32((self.index + 1)))
		return hx_post_37
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
	self.first = first
	self.second = second
	self.index = 0
	return self
}

func (self *SnapshotGenericIterator) hasNext() bool {
	return (self.index < 2)
}

func (self *SnapshotGenericIterator) next() any {
	var hx_if_39 any
	if func() int {
		hx_post_38 := self.index
		self.index = int(int32((self.index + 1)))
		return hx_post_38
	}() == 0 {
		hx_if_39 = self.first
	} else {
		hx_if_39 = self.second
	}
	return hx_if_39
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
		hx_post_43 := self.index
		self.index = int(int32((self.index + 1)))
		return hx_post_43
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

func hx_recover_callable_1(value any) *hxrt.Function[func() int] {
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

func hx_recover_callable_2(value any) *hxrt.Function[func() *string] {
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
			return func(hx_value_11 any) *string {
				if hx_value_11 == nil {
					var hx_zero_12 *string
					return hx_zero_12
				}
				return hx_value_11.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_13 any) *string {
				if hx_value_13 == nil {
					var hx_zero_14 *string
					return hx_zero_14
				}
				return hx_value_13.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *string]:
		return source
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}
