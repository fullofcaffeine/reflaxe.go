//go:build alpha && beta

package main

import "snapshot/hxrt"

func hx_recover_callable_2__probe_2(value any) (*hxrt.Function[func(*AlphaFeature) *string], bool) {
	switch source := value.(type) {
	case *hxrt.Function[func(*hxrt.Function[func(*BetaFeature) *string]) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 *AlphaFeature) *string {
			return hxrt.FunctionCall(source)(hx_recover_callable_3(arg_0))
		}), true
	default:
		var zero *hxrt.Function[func(*AlphaFeature) *string]
		return zero, false
	}
}

func hx_recover_callable_3__probe_2(value any) (*hxrt.Function[func(*BetaFeature) *string], bool) {
	switch source := value.(type) {
	case *hxrt.Function[func(*hxrt.Function[func(*AlphaFeature) *string]) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 *BetaFeature) *string {
			return hxrt.FunctionCall(source)(hx_recover_callable_2(arg_0))
		}), true
	default:
		var zero *hxrt.Function[func(*BetaFeature) *string]
		return zero, false
	}
}
