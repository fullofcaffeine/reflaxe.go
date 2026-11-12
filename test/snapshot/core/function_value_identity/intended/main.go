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
		var f2 any = any(hx_recover_callable_0(reuse(stored_1)))
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
	restored := hx_recover_callable_0(hx_recover_callable_0(boxed))
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("record="), func() int {
		var hx_callable_7 *hxrt.Function[func(int) int] = record_callback
		var hx_argument_8 int = 0
		return hxrt.FunctionCall(hx_callable_7)(hx_argument_8)
	}()), hxrt.StringFromLiteral(":")), hxrt.StdString(hxrt.ReflectCompareMethods(any(first), any(restored)))))
	hxrt.Println(v_3)
	var missing *hxrt.Function[func(int) int] = nil
	evaluated := 0
	caught := false
	hxrt.TryCatch(func() {
		func() int {
			var hx_callable_12 *hxrt.Function[func(int) int] = missing
			var hx_argument_13 int = func() int {
				evaluated = int(int32((evaluated + 1)))
				return evaluated
			}()
			return hxrt.FunctionCall(hx_callable_12)(hx_argument_13)
		}()
	}, func(hx_caught_9 any) {
		hx_exception_10 := haxe__Exception_caught(hx_caught_9)
		hx_payload_11 := haxe__Exception_catchValue(hx_caught_9)
		_ = hx_exception_10
		var hx_tmp any = hx_payload_11
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

func hxrt__generated_method_field(obj any, key string) any {
	var receiver any
	switch value := hxrt.ObjectSource(obj).(type) {
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *haxe__Exception:
		return hxrt__generated_method_field__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_method_field__haxe__ValueException(value, key)
	default:
		return nil
	}
}

func hxrt__generated_method_field__haxe__Exception(value *haxe__Exception, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "details":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_13 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_13.__hx_this.details()
			}, func() any {
				return hx_receiver_13.__hx_this
			}, "details")
		}()
	case "get_message":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_14 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_14.__hx_this.get_message()
			}, func() any {
				return hx_receiver_14.__hx_this
			}, "get_message")
		}()
	case "get_native":
		return func() *hxrt.Function[func() any] {
			hx_receiver_15 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_15.__hx_this.get_native()
			}, func() any {
				return hx_receiver_15.__hx_this
			}, "get_native")
		}()
	case "get_previous":
		return func() *hxrt.Function[func() *haxe__Exception] {
			hx_receiver_16 := value
			return hxrt.BoundFunction(func() *haxe__Exception {
				return hx_receiver_16.__hx_this.get_previous()
			}, func() any {
				return hx_receiver_16.__hx_this
			}, "get_previous")
		}()
	case "get_stack":
		return func() *hxrt.Function[func() any] {
			hx_receiver_17 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_17.__hx_this.get_stack()
			}, func() any {
				return hx_receiver_17.__hx_this
			}, "get_stack")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_18 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_18.__hx_this.toString()
			}, func() any {
				return hx_receiver_18.__hx_this
			}, "toString")
		}()
	case "unwrap":
		return func() *hxrt.Function[func() any] {
			hx_receiver_19 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_19.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_19.__hx_this
			}, "unwrap")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__ValueException(value *haxe__ValueException, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "unwrap":
		return func() *hxrt.Function[func() any] {
			hx_receiver_20 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_20.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_20.__hx_this
			}, "unwrap")
		}()
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_method_field__haxe__Exception(value.haxe__Exception, key)
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
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *haxe__Exception:
		return hxrt__generated_field_lookup__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_lookup__haxe__ValueException(value, key)
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
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *haxe__Exception:
		return hxrt__generated_field_has__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_has__haxe__ValueException(value, key)
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
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *haxe__Exception:
		return hxrt__generated_field_set__haxe__Exception(value, key, incoming)
	case *haxe__ValueException:
		return hxrt__generated_field_set__haxe__ValueException(value, key, incoming)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_set__haxe___Int64_____Int64(value, key, incoming)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedFields(object any) *hxrt.Array {
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch receiver.(type) {
	case *haxe__Exception:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"))
	case *haxe__ValueException:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("value"))
	case *haxe___Int64_____Int64:
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	default:
		return nil
	}
}

func hxrt__generated_field_lookup__haxe__Exception(value *haxe__Exception, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "savedMessage":
		return hxrt.ReadStoredField(value.savedMessage)
	case "savedNative":
		return hxrt.ReadStoredField(value.savedNative)
	case "savedPrevious":
		return hxrt.ReadStoredField(value.savedPrevious)
	case "savedStack":
		return hxrt.ReadStoredField(value.savedStack)
	}
	return nil
}

func hxrt__generated_field_has__haxe__Exception(value *haxe__Exception, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "savedMessage":
		return true
	case "savedNative":
		return true
	case "savedPrevious":
		return true
	case "savedStack":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__Exception(value *haxe__Exception, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "savedMessage":
		if incoming == nil {
			var zero *string
			value.savedMessage = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.savedMessage = typed
			return true
		default:
			return false
		}
	case "savedNative":
		if incoming == nil {
			var zero any
			value.savedNative = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.savedNative = typed
			return true
		default:
			return false
		}
	case "savedPrevious":
		if incoming == nil {
			var zero *haxe__Exception
			value.savedPrevious = zero
			return true
		}
		switch typed := incoming.(type) {
		case *haxe__Exception:
			value.savedPrevious = typed
			return true
		default:
			return false
		}
	case "savedStack":
		if incoming == nil {
			var zero any
			value.savedStack = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.savedStack = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__ValueException(value *haxe__ValueException, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "value":
		return hxrt.ReadStoredField(value.value)
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_field_lookup__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_has__haxe__ValueException(value *haxe__ValueException, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "value":
		return true
	}
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_has__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_set__haxe__ValueException(value *haxe__ValueException, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "value":
		if incoming == nil {
			var zero any
			value.value = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.value = typed
			return true
		default:
			return false
		}
	}
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_set__haxe__Exception(value.haxe__Exception, key, incoming)
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
	key := *hxrt.StdString(field)
	return hxrt__generated_method_field(object, key)
}

func reflaxe__go___internal__CompilerReflect_isEnumValue(value any) bool {
	switch enumValue := value.(type) {
	case *haxe__StackItem:
		return (enumValue != nil)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_callMethod(value any, arguments []any) any {
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Cannot call null"))
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(int) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_21 any) int {
			if hx_value_21 == nil {
				var hx_zero_22 int
				return hx_zero_22
			}
			return hx_value_21.(int)
		}(argument_0))
	case *hxrt.Function[func(...any) any]:
		var argument_0 []any = nil
		for index, argument := range arguments {
			if index >= 0 {
				argument_0 = append(argument_0, argument)
			}
		}
		return hxrt.FunctionCall(source)(argument_0...)
	case *hxrt.Function[func() *string]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() any]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.FunctionCall(source)()
	default:
		return hxrt.ReflectCallMethod(value, arguments)
	}
}

func hx_recover_callable_0(value any) *hxrt.Function[func(int) int] {
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
	case *hxrt.Function[func(int) int]:
		return source
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 int) int {
			return func(hx_value_23 any) int {
				if hx_value_23 == nil {
					var hx_zero_24 int
					return hx_zero_24
				}
				return hx_value_23.(int)
			}(hxrt.FunctionCall(source)(arg_0))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_class_0(value any) *haxe__ValueException {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ValueException:
		return source
	default:
		return value.(*haxe__ValueException)
	}
}

func hx_recover_class_1(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__Exception:
		return source
	case *haxe__ValueException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}

func hx_is_nominal_0(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_1(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__Exception:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}
