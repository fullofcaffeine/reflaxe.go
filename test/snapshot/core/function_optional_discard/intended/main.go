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
	invoke(hx_recover_callable_0(make()))
}

func make() *hxrt.Function[func(any) int] {
	return hxrt.NewFunction(func(value any) int {
		if value == nil {
			value = 9
		}
		var hx_if_9 int
		if value == nil {
			hx_if_9 = 9
		} else {
			hx_if_9 = value.(int)
		}
		result := hx_if_9
		hxrt.Println(any(result))
		return result
	})
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
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}
