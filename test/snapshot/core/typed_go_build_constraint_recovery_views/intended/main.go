package main

import "snapshot/hxrt"

func main() {
	if !hxrt.StringEqualStringPtr(Registry_selected, hxrt.StringFromLiteral("ok")) {
		hxrt.Throw(hxrt.StringFromLiteral("selected recovery did not run"))
	}
	var v any = any(Registry_selected)
	hxrt.Println(v)
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *FeatureBase:
		return func(hx_structural_source_18 *FeatureBase) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_18, func(name string) hxrt.AnonymousField {
				switch name {
				case "label":
					return hxrt.AnonymousField{Value: hx_structural_source_18.label, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "label":
					hx_structural_source_18.label = func(hx_value_19 any) *string {
						if hx_value_19 == nil {
							var hx_zero_20 *string
							return hx_zero_20
						}
						return hx_value_19.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_24 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_24, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_24.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_24.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_24.high = func(hx_value_25 any) int {
						if hx_value_25 == nil {
							var hx_zero_26 int
							return hx_zero_26
						}
						return hx_value_25.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_24.low = func(hx_value_27 any) int {
						if hx_value_27 == nil {
							var hx_zero_28 int
							return hx_zero_28
						}
						return hx_value_27.(int)
					}(incoming)
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
