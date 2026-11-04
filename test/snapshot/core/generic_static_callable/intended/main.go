package main

import "snapshot/hxrt"

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

func callback(value *hxrt.Function[func(any) any]) *hxrt.Function[func(any) any] {
	return value
}

func identity(value any) any {
	return value
}

func main() {
	typed := func(hx_callback_source_2 *hxrt.Function[func(any) any]) *hxrt.Function[func(int) int] {
		return hxrt.AdaptFunction(hx_callback_source_2, func(hx_erased_callback_arg_1 int) int {
			return func(hx_value_3 any) int {
				if hx_value_3 == nil {
					var hx_zero_4 int
					return hx_zero_4
				}
				return hx_value_3.(int)
			}(hxrt.FunctionCall(hx_callback_source_2)(hx_erased_callback_arg_1))
		})
	}(hxrt.NamedFunction(identity, "Main.identity"))
	var v any = any(func() int {
		var hx_callable_5 *hxrt.Function[func(int) int] = typed
		var hx_argument_6 int = 7
		return hxrt.FunctionCall(hx_callable_5)(hx_argument_6)
	}())
	hxrt.Println(v)
	values := hxrt.NewArray(hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("a"))
	hxrt.ArraySort(values, func(hx_comparator_12 *hxrt.Function[func(*string, *string) int]) *hxrt.Function[func(any, any) int] {
		return hxrt.AdaptFunction(hx_comparator_12, func(hx_cmp_left_10 any, hx_cmp_right_11 any) int {
			return hxrt.FunctionCall(hx_comparator_12)(func(hx_value_13 any) *string {
				if hx_value_13 == nil {
					var hx_zero_14 *string
					return hx_zero_14
				}
				return hx_value_13.(*string)
			}(hx_cmp_left_10), func(hx_value_15 any) *string {
				if hx_value_15 == nil {
					var hx_zero_16 *string
					return hx_zero_16
				}
				return hx_value_15.(*string)
			}(hx_cmp_right_11))
		})
	}(func(hx_callback_source_9 *hxrt.Function[func(any, any) int]) *hxrt.Function[func(*string, *string) int] {
		return hxrt.AdaptFunction(hx_callback_source_9, func(hx_erased_callback_arg_7 *string, hx_erased_callback_arg_8 *string) int {
			return hxrt.FunctionCall(hx_callback_source_9)(hx_erased_callback_arg_7, hx_erased_callback_arg_8)
		})
	}(hxrt.NamedFunction(Reflect_compare, "Reflect.compare"))))
	var v_1 any = any(hxrt.StringJoinAny(values.Values(), hxrt.StringFromLiteral(",")))
	hxrt.Println(v_1)
	hx_obj_17 := map[string]any{}
	hx_obj_17["value"] = 9
	source := hxrt.AnonymousFromMap(hx_obj_17)
	result := hx_structural_recover_0(identity(source))
	var v_2 any = any(func(hx_obj_19 hxrt.Anonymous) int {
		hx_field_20 := hxrt.AnonymousGet(hx_obj_19, "value")
		if hx_field_20 == nil {
			var hx_zero_21 int
			return hx_zero_21
		}
		return hx_field_20.(int)
	}(result))
	hxrt.Println(v_2)
	hxrt.Println(any(hxrt.HaxeEqual(result, source)))
	cb := func(hx_callback_source_27 *hxrt.Function[func(any) any]) *hxrt.Function[func(int) int] {
		return hxrt.AdaptFunction(hx_callback_source_27, func(hx_erased_callback_arg_26 int) int {
			return func(hx_value_28 any) int {
				if hx_value_28 == nil {
					var hx_zero_29 int
					return hx_zero_29
				}
				return hx_value_28.(int)
			}(hxrt.FunctionCall(hx_callback_source_27)(hx_erased_callback_arg_26))
		})
	}(callback(func(hx_callback_source_25 *hxrt.Function[func(int) int]) *hxrt.Function[func(any) any] {
		return hxrt.AdaptFunction(hx_callback_source_25, func(hx_erased_callback_arg_22 any) any {
			return hxrt.FunctionCall(hx_callback_source_25)(func(hx_value_23 any) int {
				if hx_value_23 == nil {
					var hx_zero_24 int
					return hx_zero_24
				}
				return hx_value_23.(int)
			}(hx_erased_callback_arg_22))
		})
	}(hxrt.NewFunction(func(value int) int {
		return int((hxrt.Int32Wrap(value) + hxrt.Int32Wrap(1)))
	}))))
	var v_3 any = any(func() int {
		var hx_callable_30 *hxrt.Function[func(int) int] = cb
		var hx_argument_31 int = 4
		return hxrt.FunctionCall(hx_callable_30)(hx_argument_31)
	}())
	hxrt.Println(v_3)
}

func hx_structural_recover_0(value any) hxrt.Anonymous {
	return hxrt.AnonymousFromValue(value)
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
