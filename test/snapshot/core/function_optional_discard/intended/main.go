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
		var hx_argument_4 any = nil
		hxrt.FunctionCall(hx_callable_3)(hx_argument_4)
	}()
	func() {
		var hx_callable_5 *hxrt.Function[func(any)] = callback
		var hx_argument_6 any = 0
		hxrt.FunctionCall(hx_callable_5)(hx_argument_6)
	}()
	func() {
		var hx_callable_7 *hxrt.Function[func(any)] = callback
		var hx_argument_8 any = 4
		hxrt.FunctionCall(hx_callable_7)(hx_argument_8)
	}()
}

func main() {
	invoke(func(hx_callback_source_10 *hxrt.Function[func(any) int]) *hxrt.Function[func(any)] {
		return hxrt.AdaptFunction(hx_callback_source_10, func(hx_erased_callback_arg_9 any) {
			hxrt.FunctionCall(hx_callback_source_10)(hx_erased_callback_arg_9)
		})
	}(make()))
}

func make() *hxrt.Function[func(any) int] {
	return hxrt.NewFunction(func(value any) int {
		if value == nil {
			value = 9
		}
		var hx_if_11 int
		if value == nil {
			hx_if_11 = 9
		} else {
			hx_if_11 = value.(int)
		}
		result := hx_if_11
		hxrt.Println(any(result))
		return result
	})
}
