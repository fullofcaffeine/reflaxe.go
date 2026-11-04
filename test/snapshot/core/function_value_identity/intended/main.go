package main

import "snapshot/hxrt"

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

func counter() *hxrt.Function[func(int) int] {
	total := 0
	return hxrt.NewFunction(func(delta int) int {
		return func() int {
			total = int((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(delta)))
			return total
		}()
	})
}

func main() {
	first := counter()
	alias := first
	other := counter()
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("identity="), hxrt.StdString(hxrt.ReflectCompareMethods(any(first), any(alias)))), hxrt.StringFromLiteral(":")), hxrt.StdString(hxrt.ReflectCompareMethods(any(first), any(other)))))
	hxrt.Println(v)
	var stored_2 *hxrt.Function[func(int) int]
	var stored_1 *hxrt.Function[func(int) int]
	var stored_0 *hxrt.Function[func(int) int]
	stored_0 = first
	stored_1 = alias
	stored_2 = other
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("transport="), hxrt.StdString(func() bool {
		var f2 any = any(reuse(stored_1))
		return hxrt.ReflectCompareMethods(any(stored_0), f2)
	}())), hxrt.StringFromLiteral(":")), hxrt.StdString(hxrt.ReflectCompareMethods(any(stored_0), any(stored_2)))))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("capture="), func() int {
		var hx_callable_1 *hxrt.Function[func(int) int] = first
		var hx_argument_2 int = 2
		return hxrt.FunctionCall(hx_callable_1)(hx_argument_2)
	}()), hxrt.StringFromLiteral(":")), func() int {
		var hx_callable_3 *hxrt.Function[func(int) int] = alias
		var hx_argument_4 int = 3
		return hxrt.FunctionCall(hx_callable_3)(hx_argument_4)
	}()), hxrt.StringFromLiteral(":")), func() int {
		var hx_callable_5 *hxrt.Function[func(int) int] = other
		var hx_argument_6 int = 1
		return hxrt.FunctionCall(hx_callable_5)(hx_argument_6)
	}()))
	hxrt.Println(v_2)
	var record_callback *hxrt.Function[func(int) int]
	record_callback = first
	extracted := record_callback
	var boxed any = extracted
	restored := func(hx_value_7 any) *hxrt.Function[func(int) int] {
		if hx_value_7 == nil {
			var hx_zero_8 *hxrt.Function[func(int) int]
			return hx_zero_8
		}
		return hx_value_7.(*hxrt.Function[func(int) int])
	}(boxed)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("record="), func() int {
		var hx_callable_9 *hxrt.Function[func(int) int] = record_callback
		var hx_argument_10 int = 0
		return hxrt.FunctionCall(hx_callable_9)(hx_argument_10)
	}()), hxrt.StringFromLiteral(":")), hxrt.StdString(hxrt.ReflectCompareMethods(any(first), any(restored)))))
	hxrt.Println(v_3)
	var missing *hxrt.Function[func(int) int] = nil
	evaluated := 0
	caught := false
	hxrt.TryCatch(func() {
		func() int {
			var hx_callable_13 *hxrt.Function[func(int) int] = missing
			var hx_argument_14 int = func() int {
				evaluated = int(int32((evaluated + 1)))
				return evaluated
			}()
			return hxrt.FunctionCall(hx_callable_13)(hx_argument_14)
		}()
	}, func(hx_caught_11 any) {
		hx_tmp := hx_caught_11
		_ = hx_tmp
		caught = true
	})
	var v_4 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("null="), hxrt.StdString(caught)), hxrt.StringFromLiteral(":")), evaluated))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("opaque="), hxrt.StdString(hxrt.AnyEqualsNull(Reflect_field(first, hxrt.StringFromLiteral("String"))))), hxrt.StringFromLiteral(":")), hxrt.StdString(Reflect_hasField(first, hxrt.StringFromLiteral("String")))))
	hxrt.Println(v_5)
}

func reuse(value any) any {
	return value
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
	switch value := object.(type) {
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
	switch value := object.(type) {
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
	switch value := object.(type) {
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
	switch value := object.(type) {
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
		return value.high
	case "low":
		return value.low
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
