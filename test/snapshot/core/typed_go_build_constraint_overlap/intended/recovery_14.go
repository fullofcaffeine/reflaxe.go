//go:build !gamma && (alpha || beta)

package main

import "snapshot/hxrt"

func reflaxe__go___internal__CompilerReflect_callMethod__probe_0(value any, arguments []any) (any, bool) {
	switch source := value.(type) {
	case *hxrt.Function[func(*hxrt.Function[func(*CompoundFeature) *string]) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_callable_b6d1e8b6(argument_0)), true
	default:
		var zero any
		return zero, false
	}
}

func hx_recover_callable_50c151b6__probe_0(value any) (*hxrt.Function[func(*hxrt.Function[func(any) *string]) *string], bool) {
	switch source := value.(type) {
	case *hxrt.Function[func(*hxrt.Function[func(*CompoundFeature) *string]) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 *hxrt.Function[func(any) *string]) *string {
			return hxrt.FunctionCall(source)(hx_recover_callable_b6d1e8b6(arg_0))
		}), true
	default:
		var zero *hxrt.Function[func(*hxrt.Function[func(any) *string]) *string]
		return zero, false
	}
}

func hx_recover_callable_ae7f85da__probe_0(value any) (*hxrt.Function[func(any) int], bool) {
	switch source := value.(type) {
	case *hxrt.Function[func(*hxrt.Function[func(*CompoundFeature) *string]) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 any) int {
			return func(hx_value_9 any) int {
				if hx_value_9 == nil {
					var hx_zero_10 int
					return hx_zero_10
				}
				return hx_value_9.(int)
			}(hxrt.FunctionCall(source)(hx_recover_callable_b6d1e8b6(arg_0)))
		}), true
	default:
		var zero *hxrt.Function[func(any) int]
		return zero, false
	}
}

func hx_recover_callable_b6d1e8b6(value any) *hxrt.Function[func(*CompoundFeature) *string] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(*CompoundFeature) *string]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(any) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 *CompoundFeature) *string {
			return hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*hxrt.Function[func(*CompoundFeature) *string]) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 *CompoundFeature) *string {
			return hxrt.FunctionCall(source)(hx_recover_callable_b6d1e8b6(arg_0))
		})
	}
	hx_recovered, hx_matched := hx_recover_callable_b6d1e8b6__probe_0(origin)
	if hx_matched {
		return hx_recovered
	}
	hx_recovered, hx_matched = hx_recover_callable_b6d1e8b6__probe_1(origin)
	if hx_matched {
		return hx_recovered
	}
	hx_recovered, hx_matched = hx_recover_callable_b6d1e8b6__probe_2(origin)
	if hx_matched {
		return hx_recovered
	}
	hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
	return nil
}
