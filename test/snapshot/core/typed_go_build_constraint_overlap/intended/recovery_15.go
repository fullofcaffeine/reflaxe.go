//go:build !gamma && (alpha || beta) && beta

package main

import "snapshot/hxrt"

func hx_recover_callable_3__probe_0(value any) (*hxrt.Function[func(*BetaFeature) *string], bool) {
	switch source := value.(type) {
	case *hxrt.Function[func(*hxrt.Function[func(*CompoundFeature) *string]) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 *BetaFeature) *string {
			return hxrt.FunctionCall(source)(hx_recover_callable_4(arg_0))
		}), true
	default:
		var zero *hxrt.Function[func(*BetaFeature) *string]
		return zero, false
	}
}

func hx_recover_callable_4__probe_2(value any) (*hxrt.Function[func(*CompoundFeature) *string], bool) {
	switch source := value.(type) {
	case *hxrt.Function[func(*hxrt.Function[func(*BetaFeature) *string]) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 *CompoundFeature) *string {
			return hxrt.FunctionCall(source)(hx_recover_callable_3(arg_0))
		}), true
	default:
		var zero *hxrt.Function[func(*CompoundFeature) *string]
		return zero, false
	}
}
