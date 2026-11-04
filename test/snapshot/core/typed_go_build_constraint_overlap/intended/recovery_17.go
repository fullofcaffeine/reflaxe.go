//go:build !gamma && alpha && (alpha || beta)

package main

import "snapshot/hxrt"

func hx_recover_callable_2__probe_0(value any) (*hxrt.Function[func(*AlphaFeature) *string], bool) {
	switch source := value.(type) {
	case *hxrt.Function[func(*hxrt.Function[func(*CompoundFeature) *string]) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 *AlphaFeature) *string {
			return hxrt.FunctionCall(source)(hx_recover_callable_4(arg_0))
		}), true
	default:
		var zero *hxrt.Function[func(*AlphaFeature) *string]
		return zero, false
	}
}

func hx_recover_callable_4__probe_1(value any) (*hxrt.Function[func(*CompoundFeature) *string], bool) {
	switch source := value.(type) {
	case *hxrt.Function[func(*hxrt.Function[func(*AlphaFeature) *string]) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 *CompoundFeature) *string {
			return hxrt.FunctionCall(source)(hx_recover_callable_2(arg_0))
		}), true
	default:
		var zero *hxrt.Function[func(*CompoundFeature) *string]
		return zero, false
	}
}
