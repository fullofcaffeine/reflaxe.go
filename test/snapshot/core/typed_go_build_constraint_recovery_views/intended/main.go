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
		return func(hx_structural_source_4 *FeatureBase) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_4, func(name string) hxrt.AnonymousField {
				switch name {
				case "label":
					return hxrt.AnonymousField{Value: hx_structural_source_4.label, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "label":
					hx_structural_source_4.label = func(hx_value_5 any) *string {
						if hx_value_5 == nil {
							var hx_zero_6 *string
							return hx_zero_6
						}
						return hx_value_5.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_10 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_10, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_10.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_10.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_10.high = func(hx_value_11 any) int {
						if hx_value_11 == nil {
							var hx_zero_12 int
							return hx_zero_12
						}
						return hx_value_11.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_10.low = func(hx_value_13 any) int {
						if hx_value_13 == nil {
							var hx_zero_14 int
							return hx_zero_14
						}
						return hx_value_13.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__Exception:
		return func(hx_structural_source_15 *haxe__Exception) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_15, func(name string) hxrt.AnonymousField {
				switch name {
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_15.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_15.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_15.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_15.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_18 := hx_structural_source_15
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_18.__hx_this.get_message()
						}, func() any {
							return hx_receiver_18.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_19 := hx_structural_source_15
						return hxrt.BoundFunction(func() any {
							return hx_receiver_19.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_19.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_20 := hx_structural_source_15
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_20.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_20.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_21 := hx_structural_source_15
						return hxrt.BoundFunction(func() any {
							return hx_receiver_21.__hx_this.get_native()
						}, func() any {
							return hx_receiver_21.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_22 := hx_structural_source_15
						return hxrt.BoundFunction(func() any {
							return hx_receiver_22.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_22.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_23 := hx_structural_source_15
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_23.__hx_this.toString()
						}, func() any {
							return hx_receiver_23.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_24 := hx_structural_source_15
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_24.__hx_this.details()
						}, func() any {
							return hx_receiver_24.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_15.savedMessage = func(hx_value_16 any) *string {
						if hx_value_16 == nil {
							var hx_zero_17 *string
							return hx_zero_17
						}
						return hx_value_16.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_15.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_15.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_15.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__ValueException:
		return func(hx_structural_source_25 *haxe__ValueException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_25, func(name string) hxrt.AnonymousField {
				switch name {
				case "value":
					return hxrt.AnonymousField{Value: hx_structural_source_25.value, Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_26 := hx_structural_source_25
						return hxrt.BoundFunction(func() any {
							return hx_receiver_26.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_26.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_25.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_25.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_25.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_25.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_29 := hx_structural_source_25
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_29.__hx_this.get_message()
						}, func() any {
							return hx_receiver_29.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_30 := hx_structural_source_25
						return hxrt.BoundFunction(func() any {
							return hx_receiver_30.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_30.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_31 := hx_structural_source_25
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_31.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_31.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_32 := hx_structural_source_25
						return hxrt.BoundFunction(func() any {
							return hx_receiver_32.__hx_this.get_native()
						}, func() any {
							return hx_receiver_32.__hx_this
						}, "get_native")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_33 := hx_structural_source_25
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_33.__hx_this.toString()
						}, func() any {
							return hx_receiver_33.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_34 := hx_structural_source_25
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_34.__hx_this.details()
						}, func() any {
							return hx_receiver_34.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_25.savedMessage = func(hx_value_27 any) *string {
						if hx_value_27 == nil {
							var hx_zero_28 *string
							return hx_zero_28
						}
						return hx_value_27.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_25.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_25.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_25.savedNative = incoming
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

func hx_recover_callable_6be299fc(value any) *hxrt.Function[func(any) *string] {
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
	hx_recovered, hx_matched := hx_recover_callable_6be299fc__probe_0(origin)
	if hx_matched {
		return hx_recovered
	}
	hx_recovered, hx_matched = hx_recover_callable_6be299fc__probe_1(origin)
	if hx_matched {
		return hx_recovered
	}
	hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
	return nil
}

func hx_recover_class_FeatureBase(value any) *FeatureBase {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *FeatureBase:
		return source
	}
	hx_recovered, hx_matched := hx_recover_class_FeatureBase__probe_0(value)
	if hx_matched {
		return hx_recovered
	}
	hx_recovered, hx_matched = hx_recover_class_FeatureBase__probe_1(value)
	if hx_matched {
		return hx_recovered
	}
	return value.(*FeatureBase)
}

func hx_recover_class_haxe__Exception(value any) *haxe__Exception {
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
	case *haxe__Exception:
		return (source != nil)
	case *haxe__ValueException:
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
