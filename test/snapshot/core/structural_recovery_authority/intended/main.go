package main

import "snapshot/hxrt"

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

type I_C interface {
}

type C struct {
	__hx_this I_C
	n         int
	m         int
}

func (self *C) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_C() *C {
	return hx_init_C(hx_alloc_C(nil))
}

func hx_alloc_C(owner any) *C {
	self := &C{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_C)
	}
	return self
}

func hx_init_C(self *C) *C {
	self.m = 2
	self.n = 1
	return self
}

func main() {
	narrow := func(hx_structural_source_1 *C) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
			switch name {
			case "n":
				return hxrt.AnonymousField{Value: hx_structural_source_1.n, Found: true}
			case "m":
				return hxrt.AnonymousField{Value: hx_structural_source_1.m, Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			case "n":
				hx_structural_source_1.n = func(hx_value_2 any) int {
					if hx_value_2 == nil {
						var hx_zero_3 int
						return hx_zero_3
					}
					return hx_value_2.(int)
				}(incoming)
				return true
			case "m":
				hx_structural_source_1.m = func(hx_value_4 any) int {
					if hx_value_4 == nil {
						var hx_zero_5 int
						return hx_zero_5
					}
					return hx_value_4.(int)
				}(incoming)
				return true
			}
			return false
		})
	}(New_C())
	var erased any = narrow
	wide := hx_structural_recover(hx_structural_recover(erased))
	var v any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("wide="), func(hx_value_6 any) int {
		if hx_value_6 == nil {
			var hx_zero_7 int
			return hx_zero_7
		}
		return hx_value_6.(int)
	}(hxrt.AnonymousGet(wide, "m"))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("array="), func(hx_value_9 any) int {
		if hx_value_9 == nil {
			var hx_zero_10 int
			return hx_zero_10
		}
		return hx_value_9.(int)
	}(read(func() hxrt.Anonymous {
		hx_obj_8 := map[string]any{}
		hx_obj_8["callbacks"] = hxrt.NewArray(hxrt.NewFunction(func() int {
			return 3
		}))
		return hxrt.AnonymousFromMap(hx_obj_8)
	}()))))
	hxrt.Println(v_1)
	hx_callable_11 := hxrt.NamedFunction(source, "Main.source")
	hx_normalized_callable_12 := hxrt.AdaptFunction(hx_callable_11, func(x any) int {
		if x == nil {
			x = 7
		}
		return hxrt.FunctionCall(hx_callable_11)(hxrt.IntFromNullableAny(x))
	})
	var projected *hxrt.Function[func(any)] = hx_recover_callable_0(hx_normalized_callable_12)
	func() {
		var hx_callable_13 *hxrt.Function[func(any)] = projected
		var hx_argument_14 any = nil
		hxrt.FunctionCall(hx_callable_13)(hx_argument_14)
	}()
	var callable any = projected
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("dynamic="), hxrt.StdString(reflaxe__go___internal__CompilerReflect_callMethod(callable, []any{}))))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("reflect="), hxrt.StdString(Reflect_callMethod(nil, projected, hxrt.NewArray()))))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("explicitNull="), hxrt.StdString(Reflect_callMethod(nil, projected, hxrt.NewArray(nil)))))
	hxrt.Println(v_4)
	var variadic any = Reflect_makeVarArgs(hx_recover_callable_1(hxrt.NewFunction(func(args *hxrt.Array) int {
		return args.Len()
	})))
	noArgs := hx_recover_callable_2(variadic)
	var v_5 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("variadic="), func() int {
		var hx_callable_15 *hxrt.Function[func() int] = noArgs
		return hxrt.FunctionCall(hx_callable_15)()
	}()))
	hxrt.Println(v_5)
}

func read(value hxrt.Anonymous) any {
	return func() any {
		var hx_callable_18 *hxrt.Function[func() any] = hx_recover_callable_3(func(hx_value_16 any) *hxrt.Array {
			if hx_value_16 == nil {
				var hx_zero_17 *hxrt.Array
				return hx_zero_17
			}
			return hx_value_16.(*hxrt.Array)
		}(hxrt.AnonymousGet(value, "callbacks")).Get(0))
		return hxrt.FunctionCall(hx_callable_18)()
	}()
}

func source(x int) int {
	return int((hxrt.Int32Wrap(x) + hxrt.Int32Wrap(2)))
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *C:
		return func(hx_structural_source_23 *C) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_23, func(name string) hxrt.AnonymousField {
				switch name {
				case "n":
					return hxrt.AnonymousField{Value: hx_structural_source_23.n, Found: true}
				case "m":
					return hxrt.AnonymousField{Value: hx_structural_source_23.m, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "n":
					hx_structural_source_23.n = func(hx_value_24 any) int {
						if hx_value_24 == nil {
							var hx_zero_25 int
							return hx_zero_25
						}
						return hx_value_24.(int)
					}(incoming)
					return true
				case "m":
					hx_structural_source_23.m = func(hx_value_26 any) int {
						if hx_value_26 == nil {
							var hx_zero_27 int
							return hx_zero_27
						}
						return hx_value_26.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_28 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_28, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_28.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_28.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_28.high = func(hx_value_29 any) int {
						if hx_value_29 == nil {
							var hx_zero_30 int
							return hx_zero_30
						}
						return hx_value_29.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_28.low = func(hx_value_31 any) int {
						if hx_value_31 == nil {
							var hx_zero_32 int
							return hx_zero_32
						}
						return hx_value_31.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	default:
		return hxrt.AnonymousFromValue(value)
	}
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
	case *C:
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
	case *C:
		return hxrt__generated_field_lookup__C(value, key)
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
	case *C:
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
	case *C:
		return hxrt__generated_field_has__C(value, key)
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
	case *C:
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
	case *C:
		return hxrt__generated_field_set__C(value, key, incoming)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_set__haxe___Int64_____Int64(value, key, incoming)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedFields(object any) *hxrt.Array {
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *C:
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
	case *C:
		return hxrt.NewArray(hxrt.StringFromLiteral("n"), hxrt.StringFromLiteral("m"))
	case *haxe___Int64_____Int64:
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	default:
		return nil
	}
}

func hxrt__generated_field_lookup__C(value *C, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "m":
		return hxrt.ReadStoredField(value.m)
	case "n":
		return hxrt.ReadStoredField(value.n)
	}
	return nil
}

func hxrt__generated_field_has__C(value *C, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "m":
		return true
	case "n":
		return true
	}
	return false
}

func hxrt__generated_field_set__C(value *C, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "m":
		if incoming == nil {
			var zero int
			value.m = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.m = typed
			return true
		default:
			return false
		}
	case "n":
		if incoming == nil {
			var zero int
			value.n = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.n = typed
			return true
		default:
			return false
		}
	}
	return false
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
	case *hxrt.Function[func() int]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(any) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(argument_0)
	case *hxrt.Function[func(*hxrt.Array) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_33 any) *hxrt.Array {
			if hx_value_33 == nil {
				var hx_zero_34 *hxrt.Array
				return hx_zero_34
			}
			return hx_value_33.(*hxrt.Array)
		}(any(argument_0)))
	case *hxrt.Function[func(int) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_35 any) int {
			if hx_value_35 == nil {
				var hx_zero_36 int
				return hx_zero_36
			}
			return hx_value_35.(int)
		}(any(argument_0)))
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
	case *hxrt.Function[func(*hxrt.Array) int]:
		return hxrt.ProjectFunction(source, func(arg_0 any) {
			hxrt.FunctionCall(source)(func(hx_value_37 any) *hxrt.Array {
				if hx_value_37 == nil {
					var hx_zero_38 *hxrt.Array
					return hx_zero_38
				}
				return hx_value_37.(*hxrt.Array)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 any) {
			hxrt.FunctionCall(source)(func(hx_value_39 any) int {
				if hx_value_39 == nil {
					var hx_zero_40 int
					return hx_zero_40
				}
				return hx_value_39.(int)
			}(any(arg_0)))
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

func hx_recover_callable_1(value any) *hxrt.Function[func(*hxrt.Array) any] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(*hxrt.Array) any]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(any) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *hxrt.Array) any {
			return hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*hxrt.Array) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *hxrt.Array) any {
			return hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *hxrt.Array) any {
			return hxrt.FunctionCall(source)(func(hx_value_41 any) int {
				if hx_value_41 == nil {
					var hx_zero_42 int
					return hx_zero_42
				}
				return hx_value_41.(int)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 *hxrt.Array) any {
			return hxrt.FunctionCall(source)(arg_0)
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
	case *hxrt.Function[func() int]:
		return source
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_43 any) int {
				if hx_value_43 == nil {
					var hx_zero_44 int
					return hx_zero_44
				}
				return hx_value_43.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_3(value any) *hxrt.Function[func() any] {
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
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}
