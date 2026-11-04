package main

import "snapshot/hxrt"

func inspect(value *hxrt.Function[func(any) bool]) bool {
	return func() bool {
		var hx_callable_1 *hxrt.Function[func(any) bool] = value
		var hx_argument_2 any = nil
		return hxrt.FunctionCall(hx_callable_1)(hx_argument_2)
	}()
}

func main() {
	integer := hxrt.NewFunction(func() any {
		return nil
	})
	decimal := hxrt.NewFunction(func() any {
		return nil
	})
	flag := hxrt.NewFunction(func() any {
		return nil
	})
	var v any = any((func() any {
		var hx_callable_3 *hxrt.Function[func() any] = integer
		return hxrt.FunctionCall(hx_callable_3)()
	}() == nil))
	hxrt.Println(v)
	var v_1 any = any((func() any {
		var hx_callable_4 *hxrt.Function[func() any] = decimal
		return hxrt.FunctionCall(hx_callable_4)()
	}() == nil))
	hxrt.Println(v_1)
	var v_2 any = any((func() any {
		var hx_callable_5 *hxrt.Function[func() any] = flag
		return hxrt.FunctionCall(hx_callable_5)()
	}() == nil))
	hxrt.Println(v_2)
	var v_3 any = any(inspect(hxrt.NewFunction(func(value any) bool {
		return (value == nil)
	})))
	hxrt.Println(v_3)
}
