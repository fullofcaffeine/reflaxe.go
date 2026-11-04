package main

import "snapshot/hxrt"

func main() {
	add := hxrt.NewFunction(func(a int, b int) int {
		return int((hxrt.Int32Wrap(a) + hxrt.Int32Wrap(b)))
	})
	mul := hxrt.NewFunction(func(v int) int {
		return int((hxrt.Int32Wrap(v) * hxrt.Int32Wrap(3)))
	})
	var v any = any(twice(5))
	hxrt.Println(v)
	var v_1 any = any(func() int {
		var hx_callable_1 *hxrt.Function[func(int, int) int] = add
		var hx_argument_2 int = 2
		var hx_argument_3 int = 7
		return hxrt.FunctionCall(hx_callable_1)(hx_argument_2, hx_argument_3)
	}())
	hxrt.Println(v_1)
	var v_2 any = any(func() int {
		var hx_callable_4 *hxrt.Function[func(int) int] = mul
		var hx_argument_5 int = 4
		return hxrt.FunctionCall(hx_callable_4)(hx_argument_5)
	}())
	hxrt.Println(v_2)
}

func twice(value int) int {
	return int((hxrt.Int32Wrap(value) * hxrt.Int32Wrap(2)))
}
