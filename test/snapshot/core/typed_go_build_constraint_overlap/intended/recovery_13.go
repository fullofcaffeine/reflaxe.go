//go:build !gamma && ((!gamma && (alpha || beta)) || alpha || beta) && (alpha || beta)

package main

import (
	"image"
	"snapshot/hxrt"
)

func hx_recover_callable_b6d1e8b6__probe_0(value any) (*hxrt.Function[func(*CompoundFeature) *string], bool) {
	switch source := value.(type) {
	case *hxrt.Function[func(*image.Point) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *CompoundFeature) *string {
			return func(hx_value_3 any) *string {
				if hx_value_3 == nil {
					var hx_zero_4 *string
					return hx_zero_4
				}
				return hx_value_3.(*string)
			}(hxrt.FunctionCall(source)(func(hx_value_1 any) *image.Point {
				if hx_value_1 == nil {
					var hx_zero_2 *image.Point
					return hx_zero_2
				}
				return hx_value_1.(*image.Point)
			}(arg_0)))
		}), true
	default:
		var zero *hxrt.Function[func(*CompoundFeature) *string]
		return zero, false
	}
}
