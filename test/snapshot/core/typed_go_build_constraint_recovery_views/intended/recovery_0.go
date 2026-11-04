//go:build !cgo

package main

import "snapshot/hxrt"

func hx_structural_recover__probe_0(value any) (hxrt.Anonymous, bool) {
	switch source := value.(type) {
	case *PureGoFeature:
		return func(hx_structural_source_21 *PureGoFeature) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_21, func(name string) hxrt.AnonymousField {
				switch name {
				case "label":
					return hxrt.AnonymousField{Value: hx_structural_source_21.label, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "label":
					hx_structural_source_21.label = func(hx_value_22 any) *string {
						if hx_value_22 == nil {
							var hx_zero_23 *string
							return hx_zero_23
						}
						return hx_value_22.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source), true
	default:
		var zero hxrt.Anonymous
		return zero, false
	}
}

func hx_structural_recover__probe_1(value any) (hxrt.Anonymous, bool) {
	var zero hxrt.Anonymous
	return zero, false
}

func reflaxe__go___internal__CompilerReflect_callMethod__probe_0(value any, arguments []any) (any, bool) {
	switch source := value.(type) {
	case *hxrt.Function[func(*PureGoFeature) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_2(argument_0)), true
	default:
		var zero any
		return zero, false
	}
}

func reflaxe__go___internal__CompilerReflect_callMethod__probe_1(value any, arguments []any) (any, bool) {
	var zero any
	return zero, false
}

func hx_recover_callable_0__probe_0(value any) (*hxrt.Function[func(any) *string], bool) {
	switch source := value.(type) {
	case *hxrt.Function[func(*PureGoFeature) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 any) *string {
			return hxrt.FunctionCall(source)(hx_recover_class_2(arg_0))
		}), true
	default:
		var zero *hxrt.Function[func(any) *string]
		return zero, false
	}
}

func hx_recover_callable_0__probe_1(value any) (*hxrt.Function[func(any) *string], bool) {
	var zero *hxrt.Function[func(any) *string]
	return zero, false
}

func hx_recover_class_0__probe_0(value any) (*FeatureBase, bool) {
	switch source := value.(type) {
	case *PureGoFeature:
		return source.FeatureBase, true
	default:
		var zero *FeatureBase
		return zero, false
	}
}

func hx_recover_class_0__probe_1(value any) (*FeatureBase, bool) {
	var zero *FeatureBase
	return zero, false
}

func hx_recover_class_2(value any) *PureGoFeature {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *PureGoFeature:
		return source
	default:
		return value.(*PureGoFeature)
	}
}
