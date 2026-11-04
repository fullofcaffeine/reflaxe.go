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
	var v_1 any = any(func(hx_fn *hxrt.Function[func(int, int) int], hx_arg_0 int, hx_arg_1 int) int {
		return hxrt.FunctionCall(hx_fn)(hx_arg_0, hx_arg_1)
	}(add, 2, 7))
	hxrt.Println(v_1)
	var v_2 any = any(func(hx_fn *hxrt.Function[func(int) int], hx_arg_0 int) int {
		return hxrt.FunctionCall(hx_fn)(hx_arg_0)
	}(mul, 4))
	hxrt.Println(v_2)
}

func twice(value int) int {
	return int((hxrt.Int32Wrap(value) * hxrt.Int32Wrap(2)))
}
