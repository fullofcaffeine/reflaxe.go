package main

import "snapshot/hxrt"

func add(a int, b int) int {
	return int((hxrt.Int32Wrap(a) + hxrt.Int32Wrap(b)))
}

func main() {
	var local *hxrt.Function[func(any) int] = hxrt.NewFunction(func(v any) int {
		if v == nil {
			v = 10
		}
		return int((hxrt.Int32Wrap(v.(int)) + hxrt.Int32Wrap(1)))
	})
	var v any = any(add(1, 2))
	hxrt.Println(v)
	var v_1 any = any(add(5, 2))
	hxrt.Println(v_1)
	var v_2 any = any(add(5, 6))
	hxrt.Println(v_2)
	var v_3 any = any(func() int {
		var hx_callable_1 *hxrt.Function[func(any) int] = local
		var hx_argument_2 any = 10
		return hxrt.FunctionCall(hx_callable_1)(hx_argument_2)
	}())
	hxrt.Println(v_3)
	var v_4 any = any(func() int {
		var hx_callable_3 *hxrt.Function[func(any) int] = local
		var hx_argument_4 any = 20
		return hxrt.FunctionCall(hx_callable_3)(hx_argument_4)
	}())
	hxrt.Println(v_4)
}
