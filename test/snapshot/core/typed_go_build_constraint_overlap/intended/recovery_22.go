//go:build beta

package main

import "snapshot/hxrt"

func reflaxe__go___internal__CompilerReflect_callMethod__probe_3(value any, arguments []any) (any, bool) {
	switch source := value.(type) {
	case *hxrt.Function[func(*hxrt.Function[func(*BetaFeature) *string]) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_callable_3(argument_0)), true
	default:
		var zero any
		return zero, false
	}
}

func hx_recover_callable_0__probe_3(value any) (*hxrt.Function[func(any) int], bool) {
	switch source := value.(type) {
	case *hxrt.Function[func(*hxrt.Function[func(*BetaFeature) *string]) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 any) int {
			return func(hx_value_37 any) int {
				if hx_value_37 == nil {
					var hx_zero_38 int
					return hx_zero_38
				}
				return hx_value_37.(int)
			}(any(hxrt.FunctionCall(source)(hx_recover_callable_3(arg_0))))
		}), true
	default:
		var zero *hxrt.Function[func(any) int]
		return zero, false
	}
}

func hx_recover_callable_1__probe_3(value any) (*hxrt.Function[func(*hxrt.Function[func(any) *string]) *string], bool) {
	switch source := value.(type) {
	case *hxrt.Function[func(*hxrt.Function[func(*BetaFeature) *string]) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 *hxrt.Function[func(any) *string]) *string {
			return hxrt.FunctionCall(source)(hx_recover_callable_3(arg_0))
		}), true
	default:
		var zero *hxrt.Function[func(*hxrt.Function[func(any) *string]) *string]
		return zero, false
	}
}

func hx_recover_callable_3(value any) *hxrt.Function[func(*BetaFeature) *string] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(*BetaFeature) *string]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(any) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 *BetaFeature) *string {
			return hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*hxrt.Function[func(*BetaFeature) *string]) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 *BetaFeature) *string {
			return hxrt.FunctionCall(source)(hx_recover_callable_3(arg_0))
		})
	}
	hx_recovered, hx_matched := hx_recover_callable_3__probe_0(origin)
	if hx_matched {
		return hx_recovered
	}
	hx_recovered, hx_matched = hx_recover_callable_3__probe_1(origin)
	if hx_matched {
		return hx_recovered
	}
	hx_recovered, hx_matched = hx_recover_callable_3__probe_2(origin)
	if hx_matched {
		return hx_recovered
	}
	hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
	return nil
}
