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
	var callback *hxrt.Function[func(any)] = hx_recover_callable_0(hx_normalized_callable_2)
	func() {
		var hx_callable_3 *hxrt.Function[func(any)] = callback
		var hx_argument_4 any = nil
		hxrt.FunctionCall(hx_callable_3)(hx_argument_4)
	}()
}

func hx_recover_callable_0(value any) *hxrt.Function[func(any)] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(any)]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(any) int]:
		return hxrt.ProjectFunction(source, func(arg_0 any) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 any) {
			hxrt.FunctionCall(source)(func(hx_value_1 any) int {
				if hx_value_1 == nil {
					var hx_zero_2 int
					return hx_zero_2
				}
				return hx_value_1.(int)
			}(arg_0))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}
