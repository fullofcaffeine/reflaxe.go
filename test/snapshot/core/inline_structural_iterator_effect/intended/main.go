package main

import "snapshot/hxrt"

var argumentEffectCount int = 0

func checkedArgumentIterator(values *hxrt.Array) hxrt.Anonymous {
	argumentEffectCount = int(int32((argumentEffectCount + 1)))
	return func(hx_structural_source_1 *haxe__iterators__ArrayIterator) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
			switch name {
			case "array":
				return hxrt.AnonymousField{Value: hx_structural_source_1.array, Found: true}
			case "current":
				return hxrt.AnonymousField{Value: hx_structural_source_1.current, Found: true}
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
			case "current":
				hx_structural_source_1.current = func(hx_value_2 any) int {
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
	}(New_haxe__iterators__ArrayIterator(values))
}

func checkedIterator(values *hxrt.Array) hxrt.Anonymous {
	effectCount = int(int32((effectCount + 1)))
	return func(hx_structural_source_6 *haxe__iterators__ArrayIterator) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_6, func(name string) hxrt.AnonymousField {
			switch name {
			case "array":
				return hxrt.AnonymousField{Value: hx_structural_source_6.array, Found: true}
			case "current":
				return hxrt.AnonymousField{Value: hx_structural_source_6.current, Found: true}
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
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
					hx_receiver_10 := hx_structural_source_6
					return hxrt.BoundFunction(func() any {
						return hx_receiver_10.__hx_this.next()
					}, func() any {
						return hx_receiver_10.__hx_this
					}, "next")
				}(), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			case "current":
				hx_structural_source_6.current = func(hx_value_7 any) int {
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
	}(New_haxe__iterators__ArrayIterator(values))
}

func collect(iterator hxrt.Anonymous) *string {
	values := hxrt.NewArray()
	for func() bool {
		var hx_callable_11 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(iterator, "hasNext"))
		return hxrt.FunctionCall(hx_callable_11)()
	}() {
		values.Push(func() int {
			var hx_callable_13 *hxrt.Function[func() int] = hx_recover_callable_1(hxrt.AnonymousGet(iterator, "next"))
			return hxrt.FunctionCall(hx_callable_13)()
		}())
	}
	return hxrt.StringJoinAny(values.Values(), hxrt.StringFromLiteral(","))
}

var effectCount int = 0

func main() {
	values := hxrt.NewArray(1, 2)
	effectCount = int(int32((effectCount + 1)))
	iterator := func(hx_structural_source_14 *haxe__iterators__ArrayIterator) hxrt.Anonymous {
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
	}(New_haxe__iterators__ArrayIterator(values))
	hx_array_target_19 := values
	hx_array_index_20 := 0
	hx_array_target_19.Set(hx_array_index_20, 9)
	var v any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("effect-before="), effectCount))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("values="), collect(iterator)))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("effect-after="), effectCount))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("argument-values="), collect(func() hxrt.Anonymous {
		argumentEffectCount = int(int32((argumentEffectCount + 1)))
		return func(hx_structural_source_21 *haxe__iterators__ArrayIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_21, func(name string) hxrt.AnonymousField {
				switch name {
				case "array":
					return hxrt.AnonymousField{Value: hx_structural_source_21.array, Found: true}
				case "current":
					return hxrt.AnonymousField{Value: hx_structural_source_21.current, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_24 := hx_structural_source_21
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_24.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_24.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_25 := hx_structural_source_21
						return hxrt.BoundFunction(func() any {
							return hx_receiver_25.__hx_this.next()
						}, func() any {
							return hx_receiver_25.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "current":
					hx_structural_source_21.current = func(hx_value_22 any) int {
						if hx_value_22 == nil {
							var hx_zero_23 int
							return hx_zero_23
						}
						return hx_value_22.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(New_haxe__iterators__ArrayIterator(hxrt.NewArray(3, 4)))
	}())))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("argument-effects="), argumentEffectCount))
	hxrt.Println(v_4)
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
			return func(hx_value_3 any) int {
				if hx_value_3 == nil {
					var hx_zero_4 int
					return hx_zero_4
				}
				return hx_value_3.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_5 any) int {
				if hx_value_5 == nil {
					var hx_zero_6 int
					return hx_zero_6
				}
				return hx_value_5.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}
