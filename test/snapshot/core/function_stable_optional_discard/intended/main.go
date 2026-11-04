package main

import "snapshot/hxrt"

func f(x int) int {
	hxrt.Println(any(x))
	return x
}

func main() {
	hx_callable_1 := hxrt.NamedFunction(f, "Main.f")
	var callback *hxrt.Function[func(any)] = hxrt.AdaptFunction(hx_callable_1, func(x any) {
		if x == nil {
			x = 9
		}
		hxrt.FunctionCall(hx_callable_1)(hxrt.IntFromNullableAny(x))
	})
	func() {
		var hx_callable_2 *hxrt.Function[func(any)] = callback
		var hx_argument_3 any = nil
		hxrt.FunctionCall(hx_callable_2)(hx_argument_3)
	}()
}
