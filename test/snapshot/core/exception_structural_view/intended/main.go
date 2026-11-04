package main

import "snapshot/hxrt"

type I_LocalException interface {
	get_message() *string
	get_stack() any
	get_previous() *haxe__Exception
	get_native() any
	unwrap() any
	toString() *string
	details() *string
}

type LocalException struct {
	*haxe__Exception
	__hx_this I_LocalException
	code      int
}

func (self *LocalException) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_LocalException(message *string, code int, previous *haxe__Exception) *LocalException {
	return hx_init_LocalException(hx_alloc_LocalException(nil), message, code, previous)
}

func hx_alloc_LocalException(owner any) *LocalException {
	self := &LocalException{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_LocalException)
	}
	self.haxe__Exception = hx_alloc_haxe__Exception(self.__hx_this)
	return self
}

func hx_init_LocalException(self *LocalException, message *string, code int, previous *haxe__Exception) *LocalException {
	hx_init_haxe__Exception(self.haxe__Exception, message, previous, nil)
	self.code = code
	return self
}

func (self *LocalException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(self.get_message(), hxrt.StringFromLiteral(":")), self.code)
}

func (self *LocalException) String() string {
	return *self.toString()
}

func main() {
	cause := New_haxe__Exception(hxrt.StringFromLiteral("root cause"), nil, nil)
	source := New_LocalException(hxrt.StringFromLiteral("boom"), 7, cause)
	var erased any = source
	view := hx_structural_recover(hx_structural_recover(erased))
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("message="), func() *string {
		var hx_callable_1 *hxrt.Function[func() *string] = hx_recover_callable_0(hxrt.AnonymousGet(view, "get_message"))
		return hxrt.FunctionCall(hx_callable_1)()
	}()))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("code="), func(hx_value_2 any) int {
		if hx_value_2 == nil {
			var hx_zero_3 int
			return hx_zero_3
		}
		return hx_value_2.(int)
	}(hxrt.AnonymousGet(view, "code"))))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("string="), func() *string {
		var hx_callable_4 *hxrt.Function[func() *string] = hx_recover_callable_0(hxrt.AnonymousGet(view, "toString"))
		return hxrt.FunctionCall(hx_callable_4)()
	}()))
	hxrt.Println(v_2)
	source.code = 8
	var v_3 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("live="), func(hx_value_5 any) int {
		if hx_value_5 == nil {
			var hx_zero_6 int
			return hx_zero_6
		}
		return hx_value_5.(int)
	}(hxrt.AnonymousGet(view, "code"))))
	hxrt.Println(v_3)
	var returned any = view
	restored := hx_recover_class_0(returned)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("same="), hxrt.StdString((restored == source))))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("direct-previous="), hxrt.StdString((source.get_previous() == cause))))
	hxrt.Println(v_5)
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("previous="), hxrt.StdString((func() *haxe__Exception {
		var hx_callable_7 *hxrt.Function[func() *haxe__Exception] = hx_recover_callable_1(hxrt.AnonymousGet(view, "get_previous"))
		return hxrt.FunctionCall(hx_callable_7)()
	}() == cause))))
	hxrt.Println(v_6)
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("native="), hxrt.StdString((func() *LocalException {
		var this1 any = source.__hx_this.get_native()
		return hx_recover_class_0(this1)
	}() == source))))
	hxrt.Println(v_7)
	directDetails := source.details()
	details := func() *string {
		var hx_callable_8 *hxrt.Function[func() *string] = hx_recover_callable_0(hxrt.AnonymousGet(view, "details"))
		return hxrt.FunctionCall(hx_callable_8)()
	}()
	var v_8 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("details-same="), hxrt.StdString(hxrt.StringEqualStringPtr(details, directDetails))))
	hxrt.Println(v_8)
	var v_9 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("details-message="), hxrt.StdString((hxrt.StringIndexOfStringPtr(details, hxrt.StringFromLiteral("boom"), 0, false) >= 0))))
	hxrt.Println(v_9)
	var v_10 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("details-cause="), hxrt.StdString((hxrt.StringIndexOfStringPtr(details, hxrt.StringFromLiteral("root cause"), 0, false) >= 0))))
	hxrt.Println(v_10)
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *LocalException:
		return func(hx_structural_source_42 *LocalException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_42, func(name string) hxrt.AnonymousField {
				switch name {
				case "code":
					return hxrt.AnonymousField{Value: hx_structural_source_42.code, Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_45 := hx_structural_source_42
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_45.__hx_this.toString()
						}, func() any {
							return hx_receiver_45.__hx_this
						}, "toString")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_42.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_42.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_42.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_42.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_48 := hx_structural_source_42
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_48.__hx_this.get_message()
						}, func() any {
							return hx_receiver_48.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_49 := hx_structural_source_42
						return hxrt.BoundFunction(func() any {
							return hx_receiver_49.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_49.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_50 := hx_structural_source_42
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_50.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_50.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_51 := hx_structural_source_42
						return hxrt.BoundFunction(func() any {
							return hx_receiver_51.__hx_this.get_native()
						}, func() any {
							return hx_receiver_51.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_52 := hx_structural_source_42
						return hxrt.BoundFunction(func() any {
							return hx_receiver_52.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_52.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_53 := hx_structural_source_42
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_53.__hx_this.details()
						}, func() any {
							return hx_receiver_53.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "code":
					hx_structural_source_42.code = func(hx_value_43 any) int {
						if hx_value_43 == nil {
							var hx_zero_44 int
							return hx_zero_44
						}
						return hx_value_43.(int)
					}(incoming)
					return true
				case "savedMessage":
					hx_structural_source_42.savedMessage = func(hx_value_46 any) *string {
						if hx_value_46 == nil {
							var hx_zero_47 *string
							return hx_zero_47
						}
						return hx_value_46.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_42.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_42.savedPrevious = hx_recover_class_2(incoming)
					return true
				case "savedNative":
					hx_structural_source_42.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_54 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_54, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_54.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_54.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_54.high = func(hx_value_55 any) int {
						if hx_value_55 == nil {
							var hx_zero_56 int
							return hx_zero_56
						}
						return hx_value_55.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_54.low = func(hx_value_57 any) int {
						if hx_value_57 == nil {
							var hx_zero_58 int
							return hx_zero_58
						}
						return hx_value_57.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__ValueException:
		return func(hx_structural_source_59 *haxe__ValueException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_59, func(name string) hxrt.AnonymousField {
				switch name {
				case "value":
					return hxrt.AnonymousField{Value: hx_structural_source_59.value, Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_60 := hx_structural_source_59
						return hxrt.BoundFunction(func() any {
							return hx_receiver_60.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_60.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_59.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_59.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_59.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_59.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_63 := hx_structural_source_59
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_63.__hx_this.get_message()
						}, func() any {
							return hx_receiver_63.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_64 := hx_structural_source_59
						return hxrt.BoundFunction(func() any {
							return hx_receiver_64.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_64.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_65 := hx_structural_source_59
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_65.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_65.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_66 := hx_structural_source_59
						return hxrt.BoundFunction(func() any {
							return hx_receiver_66.__hx_this.get_native()
						}, func() any {
							return hx_receiver_66.__hx_this
						}, "get_native")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_67 := hx_structural_source_59
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_67.__hx_this.toString()
						}, func() any {
							return hx_receiver_67.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_68 := hx_structural_source_59
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_68.__hx_this.details()
						}, func() any {
							return hx_receiver_68.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_59.savedMessage = func(hx_value_61 any) *string {
						if hx_value_61 == nil {
							var hx_zero_62 *string
							return hx_zero_62
						}
						return hx_value_61.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_59.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_59.savedPrevious = hx_recover_class_2(incoming)
					return true
				case "savedNative":
					hx_structural_source_59.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__Exception:
		return func(hx_structural_source_69 *haxe__Exception) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_69, func(name string) hxrt.AnonymousField {
				switch name {
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_69.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_69.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_69.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_69.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_72 := hx_structural_source_69
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_72.__hx_this.get_message()
						}, func() any {
							return hx_receiver_72.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_73 := hx_structural_source_69
						return hxrt.BoundFunction(func() any {
							return hx_receiver_73.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_73.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_74 := hx_structural_source_69
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_74.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_74.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_75 := hx_structural_source_69
						return hxrt.BoundFunction(func() any {
							return hx_receiver_75.__hx_this.get_native()
						}, func() any {
							return hx_receiver_75.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_76 := hx_structural_source_69
						return hxrt.BoundFunction(func() any {
							return hx_receiver_76.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_76.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_77 := hx_structural_source_69
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_77.__hx_this.toString()
						}, func() any {
							return hx_receiver_77.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_78 := hx_structural_source_69
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_78.__hx_this.details()
						}, func() any {
							return hx_receiver_78.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_69.savedMessage = func(hx_value_70 any) *string {
						if hx_value_70 == nil {
							var hx_zero_71 *string
							return hx_zero_71
						}
						return hx_value_70.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_69.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_69.savedPrevious = hx_recover_class_2(incoming)
					return true
				case "savedNative":
					hx_structural_source_69.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	default:
		return hxrt.AnonymousFromValue(value)
	}
}

func hx_recover_callable_0(value any) *hxrt.Function[func() *string] {
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
	case *hxrt.Function[func() *string]:
		return source
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_79 any) *string {
				if hx_value_79 == nil {
					var hx_zero_80 *string
					return hx_zero_80
				}
				return hx_value_79.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_81 any) *string {
				if hx_value_81 == nil {
					var hx_zero_82 *string
					return hx_zero_82
				}
				return hx_value_81.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_1(value any) *hxrt.Function[func() *haxe__Exception] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() *haxe__Exception]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() *haxe__Exception {
			return hx_recover_class_2(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() *haxe__Exception {
			return hx_recover_class_2(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return source
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_class_0(value any) *LocalException {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *LocalException:
		return source
	default:
		return value.(*LocalException)
	}
}

func hx_recover_class_1(value any) *haxe__ValueException {
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

func hx_recover_class_2(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *LocalException:
		return source.haxe__Exception
	case *haxe__ValueException:
		return source.haxe__Exception
	case *haxe__Exception:
		return source
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
	case *LocalException:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	case *haxe__Exception:
		return (source != nil)
	default:
		return false
	}
}
