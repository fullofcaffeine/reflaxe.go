package main

import "snapshot/hxrt"

func Stable_run(value any, callback any) int {
	apply := hx_recover_callable_408c8b7a(callback)
	if !hx_is_nominal_Circle(value) {
		return 0
	}
	circle := hx_recover_class_Circle(value)
	return func() int {
		var hx_callable_1 *hxrt.Function[func(int) int] = apply
		var hx_argument_2 int = circle.size
		return hxrt.FunctionCall(hx_callable_1)(hx_argument_2)
	}()
}
