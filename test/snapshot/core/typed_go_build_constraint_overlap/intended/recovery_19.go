//go:build ((!gamma && (alpha || beta)) || alpha || beta) && beta

package main

import (
	"image"
	"snapshot/hxrt"
)

func hx_recover_callable_3__probe_1(value any) (*hxrt.Function[func(*BetaFeature) *string], bool) {
	switch source := value.(type) {
	case *hxrt.Function[func(*image.Point) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *BetaFeature) *string {
			return func(hx_value_51 any) *string {
				if hx_value_51 == nil {
					var hx_zero_52 *string
					return hx_zero_52
				}
				return hx_value_51.(*string)
			}(hxrt.FunctionCall(source)(func(hx_value_49 any) *image.Point {
				if hx_value_49 == nil {
					var hx_zero_50 *image.Point
					return hx_zero_50
				}
				return hx_value_49.(*image.Point)
			}(arg_0)))
		}), true
	default:
		var zero *hxrt.Function[func(*BetaFeature) *string]
		return zero, false
	}
}
