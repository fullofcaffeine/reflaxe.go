package main

import "snapshot/hxrt"

func main() {
	if !hxrt.StringEqualStringPtr(Registry_selected, hxrt.StringFromLiteral("ok")) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("selected recovery did not run")))
	}
	var v any = any(Registry_selected)
	hxrt.Println(v)
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *FeatureBase:
		return func(hx_structural_source_38 *FeatureBase) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_38, func(name string) hxrt.AnonymousField {
				switch name {
				case "label":
					return hxrt.AnonymousField{Value: hx_structural_source_38.label, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "label":
					hx_structural_source_38.label = func(hx_value_39 any) *string {
						if hx_value_39 == nil {
							var hx_zero_40 *string
							return hx_zero_40
						}
						return hx_value_39.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_44 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_44, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_44.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_44.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_44.high = func(hx_value_45 any) int {
						if hx_value_45 == nil {
							var hx_zero_46 int
							return hx_zero_46
						}
						return hx_value_45.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_44.low = func(hx_value_47 any) int {
						if hx_value_47 == nil {
							var hx_zero_48 int
							return hx_zero_48
						}
						return hx_value_47.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__Exception:
		return func(hx_structural_source_49 *haxe__Exception) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_49, func(name string) hxrt.AnonymousField {
				switch name {
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_49.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_49.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_49.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_49.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_52 := hx_structural_source_49
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_52.__hx_this.get_message()
						}, func() any {
							return hx_receiver_52.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_53 := hx_structural_source_49
						return hxrt.BoundFunction(func() any {
							return hx_receiver_53.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_53.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_54 := hx_structural_source_49
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_54.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_54.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_55 := hx_structural_source_49
						return hxrt.BoundFunction(func() any {
							return hx_receiver_55.__hx_this.get_native()
						}, func() any {
							return hx_receiver_55.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_56 := hx_structural_source_49
						return hxrt.BoundFunction(func() any {
							return hx_receiver_56.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_56.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_57 := hx_structural_source_49
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_57.__hx_this.toString()
						}, func() any {
							return hx_receiver_57.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_58 := hx_structural_source_49
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_58.__hx_this.details()
						}, func() any {
							return hx_receiver_58.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_49.savedMessage = func(hx_value_50 any) *string {
						if hx_value_50 == nil {
							var hx_zero_51 *string
							return hx_zero_51
						}
						return hx_value_50.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_49.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_49.savedPrevious = hx_recover_class_2(incoming)
					return true
				case "savedNative":
					hx_structural_source_49.savedNative = incoming
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
	}
	hx_recovered, hx_matched := hx_structural_recover__probe_0(value)
	if hx_matched {
		return hx_recovered
	}
	hx_recovered, hx_matched = hx_structural_recover__probe_1(value)
	if hx_matched {
		return hx_recovered
	}
	return hxrt.AnonymousFromValue(value)
}

func reflaxe__go___internal__CompilerReflect_callMethod(value any, arguments []any) any {
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Cannot call null"))
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() *string]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() any]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.FunctionCall(source)()
	}
	hx_recovered, hx_matched := reflaxe__go___internal__CompilerReflect_callMethod__probe_0(origin, arguments)
	if hx_matched {
		return hx_recovered
	}
	hx_recovered, hx_matched = reflaxe__go___internal__CompilerReflect_callMethod__probe_1(origin, arguments)
	if hx_matched {
		return hx_recovered
	}
	return hxrt.ReflectCallMethod(value, arguments)
}

func hx_recover_callable_0(value any) *hxrt.Function[func(any) *string] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(any) *string]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	hx_recovered, hx_matched := hx_recover_callable_0__probe_0(origin)
	if hx_matched {
		return hx_recovered
	}
	hx_recovered, hx_matched = hx_recover_callable_0__probe_1(origin)
	if hx_matched {
		return hx_recovered
	}
	hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
	return nil
}

func hx_recover_class_0(value any) *FeatureBase {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *FeatureBase:
		return source
	}
	hx_recovered, hx_matched := hx_recover_class_0__probe_0(value)
	if hx_matched {
		return hx_recovered
	}
	hx_recovered, hx_matched = hx_recover_class_0__probe_1(value)
	if hx_matched {
		return hx_recovered
	}
	return value.(*FeatureBase)
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
