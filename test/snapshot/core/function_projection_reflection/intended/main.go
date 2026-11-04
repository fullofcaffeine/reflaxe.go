package main

import "snapshot/hxrt"

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

var calls int = 0

func higher(callback *hxrt.Function[func(int) int]) int {
	return func() int {
		var hx_callable_1 *hxrt.Function[func(int) int] = callback
		var hx_argument_2 int = 3
		return hxrt.FunctionCall(hx_callable_1)(hx_argument_2)
	}()
}

func main() {
	hx_callable_3 := hxrt.NamedFunction(source, "Main.source")
	hx_normalized_callable_4 := hxrt.AdaptFunction(hx_callable_3, func(x any) int {
		if x == nil {
			x = 7
		}
		return hxrt.FunctionCall(hx_callable_3)(hxrt.IntFromNullableAny(x))
	})
	var projected *hxrt.Function[func(any)] = func(hx_callback_source_6 *hxrt.Function[func(any) int]) *hxrt.Function[func(any)] {
		return hxrt.ProjectFunction(hx_callback_source_6, func(hx_erased_callback_arg_5 any) {
			hxrt.FunctionCall(hx_callback_source_6)(hx_erased_callback_arg_5)
		})
	}(hx_normalized_callable_4)
	func() {
		var hx_callable_7 *hxrt.Function[func(any)] = projected
		var hx_argument_8 any = nil
		hxrt.FunctionCall(hx_callable_7)(hx_argument_8)
	}()
	var v any = Reflect_callMethod(nil, projected, hxrt.NewArray(nil))
	hxrt.Println(v)
	var v_1 any = Reflect_callMethod(nil, projected, hxrt.NewArray(3))
	hxrt.Println(v_1)
	var v_2 any = any(calls)
	hxrt.Println(v_2)
	var literal *hxrt.Function[func(any)] = hx_recover_callable_0(hxrt.NewFunction(func(x any) int {
		if x == nil {
			x = 7
		}
		return int((hxrt.Int32Wrap(x.(int)) + hxrt.Int32Wrap(2)))
	}))
	var v_3 any = Reflect_callMethod(nil, literal, hxrt.NewArray())
	hxrt.Println(v_3)
	higherView := func(hx_callback_source_14 *hxrt.Function[func(*hxrt.Function[func(int) int]) int]) *hxrt.Function[func(*hxrt.Function[func(any) any])] {
		return hxrt.ProjectFunction(hx_callback_source_14, func(hx_erased_callback_arg_9 *hxrt.Function[func(any) any]) {
			hxrt.FunctionCall(hx_callback_source_14)(func(hx_callback_source_11 *hxrt.Function[func(any) any]) *hxrt.Function[func(int) int] {
				return hxrt.ProjectFunction(hx_callback_source_11, func(hx_erased_callback_arg_10 int) int {
					return func(hx_value_12 any) int {
						if hx_value_12 == nil {
							var hx_zero_13 int
							return hx_zero_13
						}
						return hx_value_12.(int)
					}(hxrt.FunctionCall(hx_callback_source_11)(hx_erased_callback_arg_10))
				})
			}(hx_erased_callback_arg_9))
		})
	}(hxrt.NamedFunction(higher, "Main.higher"))
	var v_4 any = Reflect_callMethod(nil, higherView, hxrt.NewArray(hxrt.NewFunction(func(x int) int {
		return int((hxrt.Int32Wrap(x) + hxrt.Int32Wrap(2)))
	})))
	hxrt.Println(v_4)
	hx_callable_15 := hxrt.NamedFunction(source, "Main.source")
	hx_normalized_callable_16 := hxrt.AdaptFunction(hx_callable_15, func(x any) int {
		if x == nil {
			x = 7
		}
		return hxrt.FunctionCall(hx_callable_15)(hxrt.IntFromNullableAny(x))
	})
	var dynamicSource any = hx_normalized_callable_16
	var v_5 any = reflaxe__go___internal__CompilerReflect_callMethod(dynamicSource, []any{})
	hxrt.Println(v_5)
}

func source(x int) int {
	calls = int(int32((calls + 1)))
	return int((hxrt.Int32Wrap(x) + hxrt.Int32Wrap(2)))
}

func hxrt_typeClassMetadataField(value any, key string) (any, bool) {
	classValue, ok := value.(*hxrt__TypeClassValue)
	if !ok || classValue == nil {
		return nil, false
	}
	className := *hxrt.StdString(classValue.name)
	switch className {
	default:
		return nil, false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedField(object any, field *string) any {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_lookup__haxe___Int64_____Int64(value, key)
	default:
		return nil
	}
}

func reflaxe__go___internal__CompilerReflect_hasGeneratedField(object any, field *string) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_has__haxe___Int64_____Int64(value, key)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_setGeneratedField(object any, field *string, incoming any) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_set__haxe___Int64_____Int64(value, key, incoming)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedFields(object any) *hxrt.Array {
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch receiver.(type) {
	case *haxe___Int64_____Int64:
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	default:
		return nil
	}
}

func hxrt__generated_field_lookup__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "high":
		return hxrt.ReadStoredField(value.high)
	case "low":
		return hxrt.ReadStoredField(value.low)
	}
	return nil
}

func hxrt__generated_field_has__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "high":
		return true
	case "low":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "high":
		if incoming == nil {
			var zero int
			value.high = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.high = typed
			return true
		default:
			return false
		}
	case "low":
		if incoming == nil {
			var zero int
			value.low = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.low = typed
			return true
		default:
			return false
		}
	}
	return false
}

func reflaxe__go___internal__CompilerReflect_typeField(object any, field *string) any {
	key := *hxrt.StdString(field)
	value, found := hxrt_typeClassMetadataField(object, key)
	if !found {
		return nil
	}
	return value
}

func reflaxe__go___internal__CompilerReflect_hasTypeField(object any, field *string) bool {
	key := *hxrt.StdString(field)
	_, found := hxrt_typeClassMetadataField(object, key)
	return found
}

func reflaxe__go___internal__CompilerReflect_generatedMethod(object any, field *string) any {
	return nil
}

func reflaxe__go___internal__CompilerReflect_isEnumValue(value any) bool {
	return false
}

func reflaxe__go___internal__CompilerReflect_callMethod(value any, arguments []any) any {
	switch source := hxrt.FunctionOrigin(value).(type) {
	case *hxrt.Function[func(any) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(argument_0)
	case *hxrt.Function[func(int) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_7 any) int {
			if hx_value_7 == nil {
				var hx_zero_8 int
				return hx_zero_8
			}
			return hx_value_7.(int)
		}(any(argument_0)))
	case *hxrt.Function[func(*hxrt.Function[func(int) int]) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_callable_1(argument_0))
	case *hxrt.Function[func(...any) any]:
		var argument_0 []any = nil
		for index, argument := range arguments {
			if index >= 0 {
				argument_0 = append(argument_0, argument)
			}
		}
		return hxrt.FunctionCall(source)(argument_0...)
	default:
		return hxrt.ReflectCallMethod(value, arguments)
	}
}

func hx_recover_callable_0(value any) *hxrt.Function[func(any)] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(any)]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(any) int]:
		return hxrt.ProjectFunction(source, func(arg_0 any) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 any) {
			hxrt.FunctionCall(source)(func(hx_value_9 any) int {
				if hx_value_9 == nil {
					var hx_zero_10 int
					return hx_zero_10
				}
				return hx_value_9.(int)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*hxrt.Function[func(int) int]) int]:
		return hxrt.ProjectFunction(source, func(arg_0 any) {
			hxrt.FunctionCall(source)(hx_recover_callable_1(arg_0))
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 any) {
			hxrt.FunctionCall(source)(arg_0)
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_1(value any) *hxrt.Function[func(int) int] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(int) int]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(any) int]:
		return hxrt.ProjectFunction(source, func(arg_0 int) int {
			return hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) int]:
		return source
	case *hxrt.Function[func(*hxrt.Function[func(int) int]) int]:
		return hxrt.ProjectFunction(source, func(arg_0 int) int {
			return hxrt.FunctionCall(source)(hx_recover_callable_1(arg_0))
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 int) int {
			return func(hx_value_11 any) int {
				if hx_value_11 == nil {
					var hx_zero_12 int
					return hx_zero_12
				}
				return hx_value_11.(int)
			}(any(hxrt.FunctionCall(source)(arg_0)))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}
