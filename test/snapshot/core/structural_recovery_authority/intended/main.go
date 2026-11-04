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
		return func(hx_structural_source_57 *C) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_57, func(name string) hxrt.AnonymousField {
				switch name {
				case "n":
					return hxrt.AnonymousField{Value: hx_structural_source_57.n, Found: true}
				case "m":
					return hxrt.AnonymousField{Value: hx_structural_source_57.m, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "n":
					hx_structural_source_57.n = func(hx_value_58 any) int {
						if hx_value_58 == nil {
							var hx_zero_59 int
							return hx_zero_59
						}
						return hx_value_58.(int)
					}(incoming)
					return true
				case "m":
					hx_structural_source_57.m = func(hx_value_60 any) int {
						if hx_value_60 == nil {
							var hx_zero_61 int
							return hx_zero_61
						}
						return hx_value_60.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_62 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_62, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_62.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_62.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_62.high = func(hx_value_63 any) int {
						if hx_value_63 == nil {
							var hx_zero_64 int
							return hx_zero_64
						}
						return hx_value_63.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_62.low = func(hx_value_65 any) int {
						if hx_value_65 == nil {
							var hx_zero_66 int
							return hx_zero_66
						}
						return hx_value_65.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__Exception:
		return func(hx_structural_source_67 *haxe__Exception) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_67, func(name string) hxrt.AnonymousField {
				switch name {
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_67.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_67.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_67.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_67.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_70 := hx_structural_source_67
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_70.__hx_this.get_message()
						}, func() any {
							return hx_receiver_70.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_71 := hx_structural_source_67
						return hxrt.BoundFunction(func() any {
							return hx_receiver_71.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_71.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_72 := hx_structural_source_67
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_72.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_72.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_73 := hx_structural_source_67
						return hxrt.BoundFunction(func() any {
							return hx_receiver_73.__hx_this.get_native()
						}, func() any {
							return hx_receiver_73.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_74 := hx_structural_source_67
						return hxrt.BoundFunction(func() any {
							return hx_receiver_74.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_74.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_75 := hx_structural_source_67
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_75.__hx_this.toString()
						}, func() any {
							return hx_receiver_75.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_76 := hx_structural_source_67
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_76.__hx_this.details()
						}, func() any {
							return hx_receiver_76.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_67.savedMessage = func(hx_value_68 any) *string {
						if hx_value_68 == nil {
							var hx_zero_69 *string
							return hx_zero_69
						}
						return hx_value_68.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_67.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_67.savedPrevious = hx_recover_class_1(incoming)
					return true
				case "savedNative":
					hx_structural_source_67.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__ValueException:
		return func(hx_structural_source_77 *haxe__ValueException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_77, func(name string) hxrt.AnonymousField {
				switch name {
				case "value":
					return hxrt.AnonymousField{Value: hx_structural_source_77.value, Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_78 := hx_structural_source_77
						return hxrt.BoundFunction(func() any {
							return hx_receiver_78.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_78.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_77.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_77.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_77.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_77.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_81 := hx_structural_source_77
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_81.__hx_this.get_message()
						}, func() any {
							return hx_receiver_81.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_82 := hx_structural_source_77
						return hxrt.BoundFunction(func() any {
							return hx_receiver_82.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_82.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_83 := hx_structural_source_77
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_83.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_83.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_84 := hx_structural_source_77
						return hxrt.BoundFunction(func() any {
							return hx_receiver_84.__hx_this.get_native()
						}, func() any {
							return hx_receiver_84.__hx_this
						}, "get_native")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_85 := hx_structural_source_77
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_85.__hx_this.toString()
						}, func() any {
							return hx_receiver_85.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_86 := hx_structural_source_77
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_86.__hx_this.details()
						}, func() any {
							return hx_receiver_86.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_77.savedMessage = func(hx_value_79 any) *string {
						if hx_value_79 == nil {
							var hx_zero_80 *string
							return hx_zero_80
						}
						return hx_value_79.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_77.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_77.savedPrevious = hx_recover_class_1(incoming)
					return true
				case "savedNative":
					hx_structural_source_77.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	default:
		return hxrt.AnonymousFromValue(value)
	}
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
			hx_receiver_87 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_87.__hx_this.details()
			}, func() any {
				return hx_receiver_87.__hx_this
			}, "details")
		}()
	case "get_message":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_88 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_88.__hx_this.get_message()
			}, func() any {
				return hx_receiver_88.__hx_this
			}, "get_message")
		}()
	case "get_native":
		return func() *hxrt.Function[func() any] {
			hx_receiver_89 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_89.__hx_this.get_native()
			}, func() any {
				return hx_receiver_89.__hx_this
			}, "get_native")
		}()
	case "get_previous":
		return func() *hxrt.Function[func() *haxe__Exception] {
			hx_receiver_90 := value
			return hxrt.BoundFunction(func() *haxe__Exception {
				return hx_receiver_90.__hx_this.get_previous()
			}, func() any {
				return hx_receiver_90.__hx_this
			}, "get_previous")
		}()
	case "get_stack":
		return func() *hxrt.Function[func() any] {
			hx_receiver_91 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_91.__hx_this.get_stack()
			}, func() any {
				return hx_receiver_91.__hx_this
			}, "get_stack")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_92 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_92.__hx_this.toString()
			}, func() any {
				return hx_receiver_92.__hx_this
			}, "toString")
		}()
	case "unwrap":
		return func() *hxrt.Function[func() any] {
			hx_receiver_93 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_93.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_93.__hx_this
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
			hx_receiver_94 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_94.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_94.__hx_this
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
	case *C:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
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
	case *C:
		return hxrt__generated_field_lookup__C(value, key)
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
	case *C:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
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
	case *C:
		return hxrt__generated_field_has__C(value, key)
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
	case *C:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
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
	case *C:
		return hxrt__generated_field_set__C(value, key, incoming)
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
	case *C:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
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
	case *C:
		return hxrt.NewArray(hxrt.StringFromLiteral("n"), hxrt.StringFromLiteral("m"))
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
		return hxrt.FunctionCall(source)(func(hx_value_95 any) *hxrt.Array {
			if hx_value_95 == nil {
				var hx_zero_96 *hxrt.Array
				return hx_zero_96
			}
			return hx_value_95.(*hxrt.Array)
		}(any(argument_0)))
	case *hxrt.Function[func(int) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_97 any) int {
			if hx_value_97 == nil {
				var hx_zero_98 int
				return hx_zero_98
			}
			return hx_value_97.(int)
		}(any(argument_0)))
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
			hxrt.FunctionCall(source)(func(hx_value_99 any) *hxrt.Array {
				if hx_value_99 == nil {
					var hx_zero_100 *hxrt.Array
					return hx_zero_100
				}
				return hx_value_99.(*hxrt.Array)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 any) {
			hxrt.FunctionCall(source)(func(hx_value_101 any) int {
				if hx_value_101 == nil {
					var hx_zero_102 int
					return hx_zero_102
				}
				return hx_value_101.(int)
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
			return hxrt.FunctionCall(source)(func(hx_value_103 any) int {
				if hx_value_103 == nil {
					var hx_zero_104 int
					return hx_zero_104
				}
				return hx_value_103.(int)
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
			return func(hx_value_105 any) int {
				if hx_value_105 == nil {
					var hx_zero_106 int
					return hx_zero_106
				}
				return hx_value_105.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_107 any) int {
				if hx_value_107 == nil {
					var hx_zero_108 int
					return hx_zero_108
				}
				return hx_value_107.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_109 any) int {
				if hx_value_109 == nil {
					var hx_zero_110 int
					return hx_zero_110
				}
				return hx_value_109.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_111 any) int {
				if hx_value_111 == nil {
					var hx_zero_112 int
					return hx_zero_112
				}
				return hx_value_111.(int)
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
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() any]:
		return source
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
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
