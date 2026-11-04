package main

import "snapshot/hxrt"

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

type I_Counter interface {
	get_value() int
	set_value(value int) int
}

type Counter struct {
	__hx_this I_Counter
	value     int
	reads     int
	writes    int
	stored    int
}

func (self *Counter) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Counter() *Counter {
	return hx_init_Counter(hx_alloc_Counter(nil))
}

func hx_alloc_Counter(owner any) *Counter {
	self := &Counter{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Counter)
	}
	return self
}

func hx_init_Counter(self *Counter) *Counter {
	self.stored = 4
	self.writes = 0
	self.reads = 0
	return self
}

func (self *Counter) get_value() int {
	self.reads = int(int32((self.reads + 1)))
	return self.stored
}

func (self *Counter) set_value(value int) int {
	self.writes = int(int32((self.writes + 1)))
	self.stored = value
	return value
}

func main() {
	source := New_Counter()
	var erased any = source
	view := hx_structural_recover(hx_structural_recover(erased))
	var v any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("before="), func() int {
		var hx_callable_1 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(view, "get_value"))
		return hxrt.FunctionCall(hx_callable_1)()
	}()))
	hxrt.Println(v)
	assigned := func() int {
		var hx_callable_2 *hxrt.Function[func(int) int] = hx_recover_callable_1(hxrt.AnonymousGet(view, "set_value"))
		var hx_argument_3 int = 9
		return hxrt.FunctionCall(hx_callable_2)(hx_argument_3)
	}()
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("assigned="), assigned)))
	func() int {
		var hx_callable_5 *hxrt.Function[func(int) int] = hx_recover_callable_1(hxrt.AnonymousGet(view, "set_value"))
		var hx_argument_6 int = int((hxrt.Int32Wrap(func() int {
			var hx_callable_4 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(view, "get_value"))
			return hxrt.FunctionCall(hx_callable_4)()
		}()) + hxrt.Int32Wrap(3)))
		return hxrt.FunctionCall(hx_callable_5)(hx_argument_6)
	}()
	var v_1 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("after="), func() int {
		var hx_callable_7 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(view, "get_value"))
		return hxrt.FunctionCall(hx_callable_7)()
	}()))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("reads="), source.reads))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("writes="), source.writes))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("raw-absent="), hxrt.StdString(!Reflect_hasField(view, hxrt.StringFromLiteral("value")))))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("raw-null="), hxrt.StdString(hxrt.AnyEqualsNull(Reflect_field(view, hxrt.StringFromLiteral("value"))))))
	hxrt.Println(v_5)
	raw := hx_structural_recover(hx_structural_recover(erased))
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("plain-null="), hxrt.StdString((func(hx_value_8 any) any {
		if hx_value_8 == nil {
			return nil
		}
		return hx_value_8.(int)
	}(hxrt.AnonymousGet(raw, "value")) == nil))))
	hxrt.Println(v_6)
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *Counter:
		return func(hx_structural_source_57 *Counter) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_57, func(name string) hxrt.AnonymousField {
				switch name {
				case "reads":
					return hxrt.AnonymousField{Value: hx_structural_source_57.reads, Found: true}
				case "writes":
					return hxrt.AnonymousField{Value: hx_structural_source_57.writes, Found: true}
				case "stored":
					return hxrt.AnonymousField{Value: hx_structural_source_57.stored, Found: true}
				case "get_value":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_64 := hx_structural_source_57
						return hxrt.BoundFunction(func() int {
							return hx_receiver_64.__hx_this.get_value()
						}, func() any {
							return hx_receiver_64.__hx_this
						}, "get_value")
					}(), Found: true}
				case "set_value":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_65 := hx_structural_source_57
						return hxrt.BoundFunction(func(value int) int {
							return hx_receiver_65.__hx_this.set_value(value)
						}, func() any {
							return hx_receiver_65.__hx_this
						}, "set_value")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "reads":
					hx_structural_source_57.reads = func(hx_value_58 any) int {
						if hx_value_58 == nil {
							var hx_zero_59 int
							return hx_zero_59
						}
						return hx_value_58.(int)
					}(incoming)
					return true
				case "writes":
					hx_structural_source_57.writes = func(hx_value_60 any) int {
						if hx_value_60 == nil {
							var hx_zero_61 int
							return hx_zero_61
						}
						return hx_value_60.(int)
					}(incoming)
					return true
				case "stored":
					hx_structural_source_57.stored = func(hx_value_62 any) int {
						if hx_value_62 == nil {
							var hx_zero_63 int
							return hx_zero_63
						}
						return hx_value_62.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_66 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_66, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_66.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_66.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_66.high = func(hx_value_67 any) int {
						if hx_value_67 == nil {
							var hx_zero_68 int
							return hx_zero_68
						}
						return hx_value_67.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_66.low = func(hx_value_69 any) int {
						if hx_value_69 == nil {
							var hx_zero_70 int
							return hx_zero_70
						}
						return hx_value_69.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__Exception:
		return func(hx_structural_source_71 *haxe__Exception) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_71, func(name string) hxrt.AnonymousField {
				switch name {
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_71.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_71.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_71.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_71.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_74 := hx_structural_source_71
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_74.__hx_this.get_message()
						}, func() any {
							return hx_receiver_74.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_75 := hx_structural_source_71
						return hxrt.BoundFunction(func() any {
							return hx_receiver_75.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_75.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_76 := hx_structural_source_71
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_76.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_76.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_77 := hx_structural_source_71
						return hxrt.BoundFunction(func() any {
							return hx_receiver_77.__hx_this.get_native()
						}, func() any {
							return hx_receiver_77.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_78 := hx_structural_source_71
						return hxrt.BoundFunction(func() any {
							return hx_receiver_78.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_78.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_79 := hx_structural_source_71
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_79.__hx_this.toString()
						}, func() any {
							return hx_receiver_79.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_80 := hx_structural_source_71
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_80.__hx_this.details()
						}, func() any {
							return hx_receiver_80.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_71.savedMessage = func(hx_value_72 any) *string {
						if hx_value_72 == nil {
							var hx_zero_73 *string
							return hx_zero_73
						}
						return hx_value_72.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_71.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_71.savedPrevious = hx_recover_class_1(incoming)
					return true
				case "savedNative":
					hx_structural_source_71.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__ValueException:
		return func(hx_structural_source_81 *haxe__ValueException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_81, func(name string) hxrt.AnonymousField {
				switch name {
				case "value":
					return hxrt.AnonymousField{Value: hx_structural_source_81.value, Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_82 := hx_structural_source_81
						return hxrt.BoundFunction(func() any {
							return hx_receiver_82.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_82.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_81.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_81.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_81.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_81.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_85 := hx_structural_source_81
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_85.__hx_this.get_message()
						}, func() any {
							return hx_receiver_85.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_86 := hx_structural_source_81
						return hxrt.BoundFunction(func() any {
							return hx_receiver_86.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_86.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_87 := hx_structural_source_81
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_87.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_87.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_88 := hx_structural_source_81
						return hxrt.BoundFunction(func() any {
							return hx_receiver_88.__hx_this.get_native()
						}, func() any {
							return hx_receiver_88.__hx_this
						}, "get_native")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_89 := hx_structural_source_81
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_89.__hx_this.toString()
						}, func() any {
							return hx_receiver_89.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_90 := hx_structural_source_81
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_90.__hx_this.details()
						}, func() any {
							return hx_receiver_90.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_81.savedMessage = func(hx_value_83 any) *string {
						if hx_value_83 == nil {
							var hx_zero_84 *string
							return hx_zero_84
						}
						return hx_value_83.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_81.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_81.savedPrevious = hx_recover_class_1(incoming)
					return true
				case "savedNative":
					hx_structural_source_81.savedNative = incoming
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
	case *Counter:
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
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *Counter:
		return hxrt__generated_method_field__Counter(value, key)
	case *haxe__Exception:
		return hxrt__generated_method_field__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_method_field__haxe__ValueException(value, key)
	default:
		return nil
	}
}

func hxrt__generated_method_field__Counter(value *Counter, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "get_value":
		return func() *hxrt.Function[func() int] {
			hx_receiver_91 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_91.__hx_this.get_value()
			}, func() any {
				return hx_receiver_91.__hx_this
			}, "get_value")
		}()
	case "set_value":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_92 := value
			return hxrt.BoundFunction(func(value int) int {
				return hx_receiver_92.__hx_this.set_value(value)
			}, func() any {
				return hx_receiver_92.__hx_this
			}, "set_value")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__Exception(value *haxe__Exception, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "details":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_93 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_93.__hx_this.details()
			}, func() any {
				return hx_receiver_93.__hx_this
			}, "details")
		}()
	case "get_message":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_94 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_94.__hx_this.get_message()
			}, func() any {
				return hx_receiver_94.__hx_this
			}, "get_message")
		}()
	case "get_native":
		return func() *hxrt.Function[func() any] {
			hx_receiver_95 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_95.__hx_this.get_native()
			}, func() any {
				return hx_receiver_95.__hx_this
			}, "get_native")
		}()
	case "get_previous":
		return func() *hxrt.Function[func() *haxe__Exception] {
			hx_receiver_96 := value
			return hxrt.BoundFunction(func() *haxe__Exception {
				return hx_receiver_96.__hx_this.get_previous()
			}, func() any {
				return hx_receiver_96.__hx_this
			}, "get_previous")
		}()
	case "get_stack":
		return func() *hxrt.Function[func() any] {
			hx_receiver_97 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_97.__hx_this.get_stack()
			}, func() any {
				return hx_receiver_97.__hx_this
			}, "get_stack")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_98 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_98.__hx_this.toString()
			}, func() any {
				return hx_receiver_98.__hx_this
			}, "toString")
		}()
	case "unwrap":
		return func() *hxrt.Function[func() any] {
			hx_receiver_99 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_99.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_99.__hx_this
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
			hx_receiver_100 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_100.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_100.__hx_this
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
	case *Counter:
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
	case *Counter:
		return hxrt__generated_field_lookup__Counter(value, key)
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
	case *Counter:
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
	case *Counter:
		return hxrt__generated_field_has__Counter(value, key)
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
	case *Counter:
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
	case *Counter:
		return hxrt__generated_field_set__Counter(value, key, incoming)
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
	case *Counter:
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
	case *Counter:
		return hxrt.NewArray(hxrt.StringFromLiteral("reads"), hxrt.StringFromLiteral("writes"), hxrt.StringFromLiteral("stored"))
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

func hxrt__generated_field_lookup__Counter(value *Counter, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "reads":
		return hxrt.ReadStoredField(value.reads)
	case "stored":
		return hxrt.ReadStoredField(value.stored)
	case "writes":
		return hxrt.ReadStoredField(value.writes)
	}
	return nil
}

func hxrt__generated_field_has__Counter(value *Counter, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "reads":
		return true
	case "stored":
		return true
	case "writes":
		return true
	}
	return false
}

func hxrt__generated_field_set__Counter(value *Counter, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "reads":
		if incoming == nil {
			var zero int
			value.reads = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.reads = typed
			return true
		default:
			return false
		}
	case "stored":
		if incoming == nil {
			var zero int
			value.stored = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.stored = typed
			return true
		default:
			return false
		}
	case "writes":
		if incoming == nil {
			var zero int
			value.writes = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.writes = typed
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
	case *hxrt.Function[func(...any) any]:
		var argument_0 []any = nil
		for index, argument := range arguments {
			if index >= 0 {
				argument_0 = append(argument_0, argument)
			}
		}
		return hxrt.FunctionCall(source)(argument_0...)
	case *hxrt.Function[func() int]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(int) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_101 any) int {
			if hx_value_101 == nil {
				var hx_zero_102 int
				return hx_zero_102
			}
			return hx_value_101.(int)
		}(any(argument_0)))
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

func hx_recover_callable_0(value any) *hxrt.Function[func() int] {
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
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_103 any) int {
				if hx_value_103 == nil {
					var hx_zero_104 int
					return hx_zero_104
				}
				return hx_value_103.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() int]:
		return source
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_105 any) int {
				if hx_value_105 == nil {
					var hx_zero_106 int
					return hx_zero_106
				}
				return hx_value_105.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_107 any) int {
				if hx_value_107 == nil {
					var hx_zero_108 int
					return hx_zero_108
				}
				return hx_value_107.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_109 any) int {
				if hx_value_109 == nil {
					var hx_zero_110 int
					return hx_zero_110
				}
				return hx_value_109.(int)
			}(any(hxrt.FunctionCall(source)()))
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
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 int) int {
			return func(hx_value_111 any) int {
				if hx_value_111 == nil {
					var hx_zero_112 int
					return hx_zero_112
				}
				return hx_value_111.(int)
			}(any(hxrt.FunctionCall(source)(arg_0)))
		})
	case *hxrt.Function[func(int) int]:
		return source
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
