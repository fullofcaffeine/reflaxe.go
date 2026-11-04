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
	self.n = 1
	return self
}

func main() {
	var erased any = New_C()
	view := hx_structural_recover(hx_structural_recover(erased))
	var v any = any(func(hx_value_1 any) int {
		if hx_value_1 == nil {
			var hx_zero_2 int
			return hx_zero_2
		}
		return hx_value_1.(int)
	}(hxrt.AnonymousGet(view, "n")))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringEqualStringPtr(func(hx_value_3 any) *string {
		if hx_value_3 == nil {
			var hx_zero_4 *string
			return hx_zero_4
		}
		return hx_value_3.(*string)
	}(hxrt.AnonymousGet(view, "label")), nil))
	hxrt.Println(v_1)
	var v_2 any = any(Reflect_hasField(view, hxrt.StringFromLiteral("label")))
	hxrt.Println(v_2)
	present := func(hx_structural_source_5 *Present) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_5, func(name string) hxrt.AnonymousField {
			switch name {
			case "label":
				return hxrt.AnonymousField{Value: hx_structural_source_5.label, Found: true}
			case "n":
				return hxrt.AnonymousField{Value: hx_structural_source_5.n, Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			case "label":
				hx_structural_source_5.label = func(hx_value_6 any) *string {
					if hx_value_6 == nil {
						var hx_zero_7 *string
						return hx_zero_7
					}
					return hx_value_6.(*string)
				}(incoming)
				return true
			case "n":
				hx_structural_source_5.n = func(hx_value_8 any) int {
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
	}(New_Present())
	var v_3 any = any(hxrt.StringEqualStringPtr(func(hx_value_10 any) *string {
		if hx_value_10 == nil {
			var hx_zero_11 *string
			return hx_zero_11
		}
		return hx_value_10.(*string)
	}(hxrt.AnonymousGet(present, "label")), nil))
	hxrt.Println(v_3)
	var v_4 any = any(Reflect_hasField(present, hxrt.StringFromLiteral("label")))
	hxrt.Println(v_4)
}

type I_Present interface {
}

type Present struct {
	*C
	__hx_this I_Present
	label     *string
}

func (self *Present) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Present() *Present {
	return hx_init_Present(hx_alloc_Present(nil))
}

func hx_alloc_Present(owner any) *Present {
	self := &Present{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Present)
	}
	self.C = hx_alloc_C(self.__hx_this)
	return self
}

func hx_init_Present(self *Present) *Present {
	hx_init_C(self.C)
	self.label = nil
	return self
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *C:
		return func(hx_structural_source_42 *C) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_42, func(name string) hxrt.AnonymousField {
				switch name {
				case "n":
					return hxrt.AnonymousField{Value: hx_structural_source_42.n, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "n":
					hx_structural_source_42.n = func(hx_value_43 any) int {
						if hx_value_43 == nil {
							var hx_zero_44 int
							return hx_zero_44
						}
						return hx_value_43.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *Present:
		return func(hx_structural_source_45 *Present) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_45, func(name string) hxrt.AnonymousField {
				switch name {
				case "label":
					return hxrt.AnonymousField{Value: hx_structural_source_45.label, Found: true}
				case "n":
					return hxrt.AnonymousField{Value: hx_structural_source_45.n, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "label":
					hx_structural_source_45.label = func(hx_value_46 any) *string {
						if hx_value_46 == nil {
							var hx_zero_47 *string
							return hx_zero_47
						}
						return hx_value_46.(*string)
					}(incoming)
					return true
				case "n":
					hx_structural_source_45.n = func(hx_value_48 any) int {
						if hx_value_48 == nil {
							var hx_zero_49 int
							return hx_zero_49
						}
						return hx_value_48.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_50 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_50, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_50.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_50.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_50.high = func(hx_value_51 any) int {
						if hx_value_51 == nil {
							var hx_zero_52 int
							return hx_zero_52
						}
						return hx_value_51.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_50.low = func(hx_value_53 any) int {
						if hx_value_53 == nil {
							var hx_zero_54 int
							return hx_zero_54
						}
						return hx_value_53.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__Exception:
		return func(hx_structural_source_55 *haxe__Exception) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_55, func(name string) hxrt.AnonymousField {
				switch name {
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_55.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_55.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_55.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_55.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_58 := hx_structural_source_55
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_58.__hx_this.get_message()
						}, func() any {
							return hx_receiver_58.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_59 := hx_structural_source_55
						return hxrt.BoundFunction(func() any {
							return hx_receiver_59.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_59.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_60 := hx_structural_source_55
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_60.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_60.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_61 := hx_structural_source_55
						return hxrt.BoundFunction(func() any {
							return hx_receiver_61.__hx_this.get_native()
						}, func() any {
							return hx_receiver_61.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_62 := hx_structural_source_55
						return hxrt.BoundFunction(func() any {
							return hx_receiver_62.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_62.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_63 := hx_structural_source_55
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_63.__hx_this.toString()
						}, func() any {
							return hx_receiver_63.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_64 := hx_structural_source_55
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_64.__hx_this.details()
						}, func() any {
							return hx_receiver_64.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_55.savedMessage = func(hx_value_56 any) *string {
						if hx_value_56 == nil {
							var hx_zero_57 *string
							return hx_zero_57
						}
						return hx_value_56.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_55.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_55.savedPrevious = hx_recover_class_1(incoming)
					return true
				case "savedNative":
					hx_structural_source_55.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__ValueException:
		return func(hx_structural_source_65 *haxe__ValueException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_65, func(name string) hxrt.AnonymousField {
				switch name {
				case "value":
					return hxrt.AnonymousField{Value: hx_structural_source_65.value, Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_66 := hx_structural_source_65
						return hxrt.BoundFunction(func() any {
							return hx_receiver_66.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_66.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_65.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_65.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_65.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_65.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_69 := hx_structural_source_65
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_69.__hx_this.get_message()
						}, func() any {
							return hx_receiver_69.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_70 := hx_structural_source_65
						return hxrt.BoundFunction(func() any {
							return hx_receiver_70.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_70.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_71 := hx_structural_source_65
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_71.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_71.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_72 := hx_structural_source_65
						return hxrt.BoundFunction(func() any {
							return hx_receiver_72.__hx_this.get_native()
						}, func() any {
							return hx_receiver_72.__hx_this
						}, "get_native")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_73 := hx_structural_source_65
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_73.__hx_this.toString()
						}, func() any {
							return hx_receiver_73.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_74 := hx_structural_source_65
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_74.__hx_this.details()
						}, func() any {
							return hx_receiver_74.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_65.savedMessage = func(hx_value_67 any) *string {
						if hx_value_67 == nil {
							var hx_zero_68 *string
							return hx_zero_68
						}
						return hx_value_67.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_65.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_65.savedPrevious = hx_recover_class_1(incoming)
					return true
				case "savedNative":
					hx_structural_source_65.savedNative = incoming
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
			hx_receiver_75 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_75.__hx_this.details()
			}, func() any {
				return hx_receiver_75.__hx_this
			}, "details")
		}()
	case "get_message":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_76 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_76.__hx_this.get_message()
			}, func() any {
				return hx_receiver_76.__hx_this
			}, "get_message")
		}()
	case "get_native":
		return func() *hxrt.Function[func() any] {
			hx_receiver_77 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_77.__hx_this.get_native()
			}, func() any {
				return hx_receiver_77.__hx_this
			}, "get_native")
		}()
	case "get_previous":
		return func() *hxrt.Function[func() *haxe__Exception] {
			hx_receiver_78 := value
			return hxrt.BoundFunction(func() *haxe__Exception {
				return hx_receiver_78.__hx_this.get_previous()
			}, func() any {
				return hx_receiver_78.__hx_this
			}, "get_previous")
		}()
	case "get_stack":
		return func() *hxrt.Function[func() any] {
			hx_receiver_79 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_79.__hx_this.get_stack()
			}, func() any {
				return hx_receiver_79.__hx_this
			}, "get_stack")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_80 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_80.__hx_this.toString()
			}, func() any {
				return hx_receiver_80.__hx_this
			}, "toString")
		}()
	case "unwrap":
		return func() *hxrt.Function[func() any] {
			hx_receiver_81 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_81.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_81.__hx_this
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
			hx_receiver_82 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_82.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_82.__hx_this
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
	case *Present:
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
	case *Present:
		return hxrt__generated_field_lookup__Present(value, key)
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
	case *Present:
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
	case *Present:
		return hxrt__generated_field_has__Present(value, key)
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
	case *Present:
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
	case *Present:
		return hxrt__generated_field_set__Present(value, key, incoming)
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
	case *Present:
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
		return hxrt.NewArray(hxrt.StringFromLiteral("n"))
	case *Present:
		return hxrt.NewArray(hxrt.StringFromLiteral("n"), hxrt.StringFromLiteral("label"))
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

func hxrt__generated_field_lookup__Present(value *Present, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "label":
		return hxrt.ReadStoredField(value.label)
	}
	if value.C == nil {
		return nil
	}
	return hxrt__generated_field_lookup__C(value.C, key)
}

func hxrt__generated_field_has__Present(value *Present, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "label":
		return true
	}
	if value.C == nil {
		return false
	}
	return hxrt__generated_field_has__C(value.C, key)
}

func hxrt__generated_field_set__Present(value *Present, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "label":
		if incoming == nil {
			var zero *string
			value.label = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.label = typed
			return true
		default:
			return false
		}
	}
	if value.C == nil {
		return false
	}
	return hxrt__generated_field_set__C(value.C, key, incoming)
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
