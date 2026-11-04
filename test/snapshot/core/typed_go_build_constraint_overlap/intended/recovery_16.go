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
		return hxrt.FunctionCall(source)(func(hx_value_29 any) *image.Point {
			if hx_value_29 == nil {
				var hx_zero_30 *image.Point
				return hx_zero_30
			}
			return hx_value_29.(*image.Point)
		}(any(argument_0))), true
	default:
		var zero any
		return zero, false
	}
}

func hx_recover_callable_0__probe_1(value any) (*hxrt.Function[func(any) int], bool) {
	switch source := value.(type) {
	case *hxrt.Function[func(*image.Point) int]:
		return hxrt.ProjectFunction(source, func(arg_0 any) int {
			return hxrt.FunctionCall(source)(func(hx_value_31 any) *image.Point {
				if hx_value_31 == nil {
					var hx_zero_32 *image.Point
					return hx_zero_32
				}
				return hx_value_31.(*image.Point)
			}(any(arg_0)))
		}), true
	default:
		var zero *hxrt.Function[func(any) int]
		return zero, false
	}
}

func hx_recover_callable_1__probe_1(value any) (*hxrt.Function[func(*hxrt.Function[func(any) *string]) *string], bool) {
	switch source := value.(type) {
	case *hxrt.Function[func(*image.Point) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *hxrt.Function[func(any) *string]) *string {
			return func(hx_value_43 any) *string {
				if hx_value_43 == nil {
					var hx_zero_44 *string
					return hx_zero_44
				}
				return hx_value_43.(*string)
			}(any(hxrt.FunctionCall(source)(func(hx_value_41 any) *image.Point {
				if hx_value_41 == nil {
					var hx_zero_42 *image.Point
					return hx_zero_42
				}
				return hx_value_41.(*image.Point)
			}(any(arg_0)))))
		}), true
	default:
		var zero *hxrt.Function[func(*hxrt.Function[func(any) *string]) *string]
		return zero, false
	}
}
