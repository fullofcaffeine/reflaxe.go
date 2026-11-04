package main

import "snapshot/hxrt"

func invoke(callback *hxrt.Function[func(any)]) {
	func() {
		var hx_callable_1 *hxrt.Function[func(any)] = callback
		var hx_argument_2 any = nil
		hxrt.FunctionCall(hx_callable_1)(hx_argument_2)
	}()
	func() {
		var hx_callable_3 *hxrt.Function[func(any)] = callback
		var hx_argument_4 any = 0
		hxrt.FunctionCall(hx_callable_3)(hx_argument_4)
	}()
	func() {
		var hx_callable_5 *hxrt.Function[func(any)] = callback
		var hx_argument_6 any = 4
		hxrt.FunctionCall(hx_callable_5)(hx_argument_6)
	}()
}

func main() {
	invoke(func(hx_callback_source_8 *hxrt.Function[func(any) int]) *hxrt.Function[func(any)] {
		return hxrt.AdaptFunction(hx_callback_source_8, func(hx_erased_callback_arg_7 any) {
			hxrt.FunctionCall(hx_callback_source_8)(hx_erased_callback_arg_7)
		})
	}(make()))
}

func make() *hxrt.Function[func(any) int] {
	return hxrt.NewFunction(func(value any) int {
		if value == nil {
			value = 9
		}
		hxrt.Println(any(value.(int)))
		return value.(int)
	})
}
