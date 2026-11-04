package main

import "snapshot/hxrt"

func main() {
	factor := 3
	mul := hxrt.NewFunction(func(v int) int {
		return int((hxrt.Int32Wrap(v) * hxrt.Int32Wrap(factor)))
	})
	factor = 4
	var v any = any(func() int {
		var hx_callable_1 *hxrt.Function[func(int) int] = mul
		var hx_argument_2 int = 2
		return hxrt.FunctionCall(hx_callable_1)(hx_argument_2)
	}())
	hxrt.Println(v)
}
