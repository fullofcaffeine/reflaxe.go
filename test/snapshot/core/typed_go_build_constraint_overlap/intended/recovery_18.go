//go:build ((!gamma && (alpha || beta)) || alpha || beta) && alpha

package main

import (
	"image"
	"snapshot/hxrt"
)

func hx_recover_callable_dbaff400__probe_1(value any) (*hxrt.Function[func(*AlphaFeature) *string], bool) {
	switch source := value.(type) {
	case *hxrt.Function[func(*image.Point) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *AlphaFeature) *string {
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
		var zero *hxrt.Function[func(*AlphaFeature) *string]
		return zero, false
	}
}
