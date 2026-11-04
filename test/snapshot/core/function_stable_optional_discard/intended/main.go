package main

import "snapshot/hxrt"

func f(x int) int {
	hxrt.Println(any(x))
	return x
}

func main() {
	hx_callable_1 := hxrt.NamedFunction(f, "Main.f")
	hx_normalized_callable_2 := hxrt.AdaptFunction(hx_callable_1, func(x any) int {
		if x == nil {
			x = 9
		}
		return hxrt.FunctionCall(hx_callable_1)(hxrt.IntFromNullableAny(x))
	})
	var callback *hxrt.Function[func(any)] = func(hx_callback_source_4 *hxrt.Function[func(any) int]) *hxrt.Function[func(any)] {
		return hxrt.ProjectFunction(hx_callback_source_4, func(hx_erased_callback_arg_3 any) {
			hxrt.FunctionCall(hx_callback_source_4)(hx_erased_callback_arg_3)
		})
	}(hx_normalized_callable_2)
	func() {
		var hx_callable_5 *hxrt.Function[func(any)] = callback
		var hx_argument_6 any = nil
		hxrt.FunctionCall(hx_callable_5)(hx_argument_6)
	}()
}
