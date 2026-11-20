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
		var hx_callable_1 *hxrt.Function[func() *string] = hx_recover_callable_b9d17ba0(hxrt.AnonymousGet(view, "get_message"))
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
		var hx_callable_4 *hxrt.Function[func() *string] = hx_recover_callable_b9d17ba0(hxrt.AnonymousGet(view, "toString"))
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
	restored := hx_recover_class_LocalException(returned)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("same="), hxrt.StdString((restored == source))))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("direct-previous="), hxrt.StdString((source.get_previous() == cause))))
	hxrt.Println(v_5)
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("previous="), hxrt.StdString((func() *haxe__Exception {
		var hx_callable_7 *hxrt.Function[func() *haxe__Exception] = hx_recover_callable_5932719d(hxrt.AnonymousGet(view, "get_previous"))
		return hxrt.FunctionCall(hx_callable_7)()
	}() == cause))))
	hxrt.Println(v_6)
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("native="), hxrt.StdString((func() *LocalException {
		var this1 any = source.__hx_this.get_native()
		return hx_recover_class_LocalException(this1)
	}() == source))))
	hxrt.Println(v_7)
	directDetails := source.details()
	details := func() *string {
		var hx_callable_8 *hxrt.Function[func() *string] = hx_recover_callable_b9d17ba0(hxrt.AnonymousGet(view, "details"))
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
		return func(hx_structural_source_1 *LocalException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
				switch name {
				case "code":
					return hxrt.AnonymousField{Value: hx_structural_source_1.code, Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_4 := hx_structural_source_1
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_4.__hx_this.toString()
						}, func() any {
							return hx_receiver_4.__hx_this
						}, "toString")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_1.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_1.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_1.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_1.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_7 := hx_structural_source_1
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_7.__hx_this.get_message()
						}, func() any {
							return hx_receiver_7.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_8 := hx_structural_source_1
						return hxrt.BoundFunction(func() any {
							return hx_receiver_8.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_8.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_9 := hx_structural_source_1
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_9.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_9.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_10 := hx_structural_source_1
						return hxrt.BoundFunction(func() any {
							return hx_receiver_10.__hx_this.get_native()
						}, func() any {
							return hx_receiver_10.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_11 := hx_structural_source_1
						return hxrt.BoundFunction(func() any {
							return hx_receiver_11.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_11.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_12 := hx_structural_source_1
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_12.__hx_this.details()
						}, func() any {
							return hx_receiver_12.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "code":
					hx_structural_source_1.code = func(hx_value_2 any) int {
						if hx_value_2 == nil {
							var hx_zero_3 int
							return hx_zero_3
						}
						return hx_value_2.(int)
					}(incoming)
					return true
				case "savedMessage":
					hx_structural_source_1.savedMessage = func(hx_value_5 any) *string {
						if hx_value_5 == nil {
							var hx_zero_6 *string
							return hx_zero_6
						}
						return hx_value_5.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_1.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_1.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_1.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_13 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_13, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_13.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_13.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_13.high = func(hx_value_14 any) int {
						if hx_value_14 == nil {
							var hx_zero_15 int
							return hx_zero_15
						}
						return hx_value_14.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_13.low = func(hx_value_16 any) int {
						if hx_value_16 == nil {
							var hx_zero_17 int
							return hx_zero_17
						}
						return hx_value_16.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__ValueException:
		return func(hx_structural_source_18 *haxe__ValueException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_18, func(name string) hxrt.AnonymousField {
				switch name {
				case "value":
					return hxrt.AnonymousField{Value: hx_structural_source_18.value, Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_19 := hx_structural_source_18
						return hxrt.BoundFunction(func() any {
							return hx_receiver_19.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_19.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_18.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_18.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_18.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_18.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_22 := hx_structural_source_18
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_22.__hx_this.get_message()
						}, func() any {
							return hx_receiver_22.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_23 := hx_structural_source_18
						return hxrt.BoundFunction(func() any {
							return hx_receiver_23.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_23.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_24 := hx_structural_source_18
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_24.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_24.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_25 := hx_structural_source_18
						return hxrt.BoundFunction(func() any {
							return hx_receiver_25.__hx_this.get_native()
						}, func() any {
							return hx_receiver_25.__hx_this
						}, "get_native")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_26 := hx_structural_source_18
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_26.__hx_this.toString()
						}, func() any {
							return hx_receiver_26.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_27 := hx_structural_source_18
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_27.__hx_this.details()
						}, func() any {
							return hx_receiver_27.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_18.savedMessage = func(hx_value_20 any) *string {
						if hx_value_20 == nil {
							var hx_zero_21 *string
							return hx_zero_21
						}
						return hx_value_20.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_18.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_18.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_18.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__Exception:
		return func(hx_structural_source_28 *haxe__Exception) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_28, func(name string) hxrt.AnonymousField {
				switch name {
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_28.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_28.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_28.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_28.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_31 := hx_structural_source_28
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_31.__hx_this.get_message()
						}, func() any {
							return hx_receiver_31.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_32 := hx_structural_source_28
						return hxrt.BoundFunction(func() any {
							return hx_receiver_32.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_32.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_33 := hx_structural_source_28
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_33.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_33.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_34 := hx_structural_source_28
						return hxrt.BoundFunction(func() any {
							return hx_receiver_34.__hx_this.get_native()
						}, func() any {
							return hx_receiver_34.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_35 := hx_structural_source_28
						return hxrt.BoundFunction(func() any {
							return hx_receiver_35.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_35.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_36 := hx_structural_source_28
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_36.__hx_this.toString()
						}, func() any {
							return hx_receiver_36.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_37 := hx_structural_source_28
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_37.__hx_this.details()
						}, func() any {
							return hx_receiver_37.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_28.savedMessage = func(hx_value_29 any) *string {
						if hx_value_29 == nil {
							var hx_zero_30 *string
							return hx_zero_30
						}
						return hx_value_29.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_28.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_28.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_28.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	default:
		return hxrt.AnonymousFromValue(value)
	}
}

func hx_recover_callable_5932719d(value any) *hxrt.Function[func() *haxe__Exception] {
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
			return hx_recover_class_haxe__Exception(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() *haxe__Exception {
			return hx_recover_class_haxe__Exception(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return source
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_b9d17ba0(value any) *hxrt.Function[func() *string] {
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
			return func(hx_value_1 any) *string {
				if hx_value_1 == nil {
					var hx_zero_2 *string
					return hx_zero_2
				}
				return hx_value_1.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_3 any) *string {
				if hx_value_3 == nil {
					var hx_zero_4 *string
					return hx_zero_4
				}
				return hx_value_3.(*string)
			}(hxrt.FunctionCall(source)())
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_class_LocalException(value any) *LocalException {
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

func hx_recover_class_haxe__Exception(value any) *haxe__Exception {
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

func hx_recover_class_haxe__ValueException(value any) *haxe__ValueException {
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

func hx_is_nominal_haxe__Exception(value any) bool {
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

func hx_is_nominal_haxe__ValueException(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}
