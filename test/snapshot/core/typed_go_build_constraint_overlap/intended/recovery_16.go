//go:build (!gamma && (alpha || beta)) || alpha || beta

package main

import (
	"image"
	"snapshot/hxrt"
)

func reflaxe__go___internal__CompilerReflect_callMethod__probe_1(value any, arguments []any) (any, bool) {
	switch source := value.(type) {
	case *hxrt.Function[func(*image.Point) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_1 any) *image.Point {
			if hx_value_1 == nil {
				var hx_zero_2 *image.Point
				return hx_zero_2
			}
			return hx_value_1.(*image.Point)
		}(argument_0)), true
	default:
		var zero any
		return zero, false
	}
}

func hx_recover_callable_50c151b6__probe_1(value any) (*hxrt.Function[func(*hxrt.Function[func(any) *string]) *string], bool) {
	switch source := value.(type) {
	case *hxrt.Function[func(*image.Point) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *hxrt.Function[func(any) *string]) *string {
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
		var zero *hxrt.Function[func(*hxrt.Function[func(any) *string]) *string]
		return zero, false
	}
}

func hx_recover_callable_ae7f85da__probe_1(value any) (*hxrt.Function[func(any) int], bool) {
	switch source := value.(type) {
	case *hxrt.Function[func(*image.Point) int]:
		return hxrt.ProjectFunction(source, func(arg_0 any) int {
			return hxrt.FunctionCall(source)(func(hx_value_1 any) *image.Point {
				if hx_value_1 == nil {
					var hx_zero_2 *image.Point
					return hx_zero_2
				}
				return hx_value_1.(*image.Point)
			}(arg_0))
		}), true
	default:
		var zero *hxrt.Function[func(any) int]
		return zero, false
	}
}
