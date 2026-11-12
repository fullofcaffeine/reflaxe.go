//go:build ((!gamma && (alpha || beta)) || alpha || beta) && alpha

package main

import (
	"image"
	"snapshot/hxrt"
)

func hx_recover_callable_2__probe_1(value any) (*hxrt.Function[func(*AlphaFeature) *string], bool) {
	switch source := value.(type) {
	case *hxrt.Function[func(*image.Point) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *AlphaFeature) *string {
			return func(hx_value_47 any) *string {
				if hx_value_47 == nil {
					var hx_zero_48 *string
					return hx_zero_48
				}
				return hx_value_47.(*string)
			}(hxrt.FunctionCall(source)(func(hx_value_45 any) *image.Point {
				if hx_value_45 == nil {
					var hx_zero_46 *image.Point
					return hx_zero_46
				}
				return hx_value_45.(*image.Point)
			}(arg_0)))
		}), true
	default:
		var zero *hxrt.Function[func(*AlphaFeature) *string]
		return zero, false
	}
}
