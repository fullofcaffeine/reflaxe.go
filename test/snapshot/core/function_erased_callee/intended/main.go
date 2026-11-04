package main

import "snapshot/hxrt"

func invoke(values *hxrt.Array, index int) int {
	return func() int {
		var hx_callable_3 *hxrt.Function[func(int) int] = func(hx_value_1 any) *hxrt.Function[func(int) int] {
			if hx_value_1 == nil {
				var hx_zero_2 *hxrt.Function[func(int) int]
				return hx_zero_2
			}
			return hx_value_1.(*hxrt.Function[func(int) int])
		}(values.Get(index))
		var hx_argument_4 int = 20
		return hxrt.FunctionCall(hx_callable_3)(hx_argument_4)
	}()
}

func main() {
	var v any = any(invoke(hxrt.NewArray(hxrt.NewFunction(func(value int) int {
		return int((hxrt.Int32Wrap(value) + hxrt.Int32Wrap(3)))
	})), 0))
	hxrt.Println(v)
	var boxed any = hxrt.NewFunction(func(value int) int {
		return int((hxrt.Int32Wrap(value) * hxrt.Int32Wrap(2)))
	})
	var v_1 any = any(func() int {
		var hx_callable_7 *hxrt.Function[func(int) int] = func(hx_value_5 any) *hxrt.Function[func(int) int] {
			if hx_value_5 == nil {
				var hx_zero_6 *hxrt.Function[func(int) int]
				return hx_zero_6
			}
			return hx_value_5.(*hxrt.Function[func(int) int])
		}(boxed)
		var hx_argument_8 int = 11
		return hxrt.FunctionCall(hx_callable_7)(hx_argument_8)
	}())
	hxrt.Println(v_1)
}
