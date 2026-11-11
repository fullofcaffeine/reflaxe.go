package main

import "snapshot/hxrt"

func main() {
	report := hxrt.NewFunction(func(expected *hxrt.Array, value any) {
		var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StdString(expected.Get(0)), hxrt.StringFromLiteral(":")), hxrt.StdString(value)))
		hxrt.Println(v)
	})
	func() {
		var hx_callable_1 *hxrt.Function[func(*hxrt.Array, int)] = hx_recover_callable_0(report)
		var hx_argument_2 *hxrt.Array = hxrt.NewArray(2)
		var hx_argument_3 int = 2
		hxrt.FunctionCall(hx_callable_1)(hx_argument_2, hx_argument_3)
	}()
	func() {
		var hx_callable_4 *hxrt.Function[func(*hxrt.Array, *string)] = hx_recover_callable_1(report)
		var hx_argument_5 *hxrt.Array = hxrt.NewArray(hxrt.StringFromLiteral("word"))
		var hx_argument_6 *string = hxrt.StringFromLiteral("word")
		hxrt.FunctionCall(hx_callable_4)(hx_argument_5, hx_argument_6)
	}()
	outer := hxrt.NewFunction(func(expected *hxrt.Array, value any) {
		inner := hxrt.NewFunction(func(items *hxrt.Array, item any) {
			var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StdString(items.Get(0)), hxrt.StringFromLiteral(":")), hxrt.StdString(item)))
			hxrt.Println(v)
		})
		func() {
			var hx_callable_7 *hxrt.Function[func(*hxrt.Array, any)] = inner
			var hx_argument_8 *hxrt.Array = expected
			var hx_argument_9 any = value
			hxrt.FunctionCall(hx_callable_7)(hx_argument_8, hx_argument_9)
		}()
	})
	func() {
		var hx_callable_10 *hxrt.Function[func(*hxrt.Array, int)] = hx_recover_callable_0(outer)
		var hx_argument_11 *hxrt.Array = hxrt.NewArray(42)
		var hx_argument_12 int = 42
		hxrt.FunctionCall(hx_callable_10)(hx_argument_11, hx_argument_12)
	}()
}

func hx_recover_callable_0(value any) *hxrt.Function[func(*hxrt.Array, int)] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(*hxrt.Array, int)]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(*hxrt.Array, any)]:
		return hxrt.ProjectFunction(source, func(arg_0 *hxrt.Array, arg_1 int) {
			hxrt.FunctionCall(source)(arg_0, arg_1)
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_1(value any) *hxrt.Function[func(*hxrt.Array, *string)] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(*hxrt.Array, *string)]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(*hxrt.Array, any)]:
		return hxrt.ProjectFunction(source, func(arg_0 *hxrt.Array, arg_1 *string) {
			hxrt.FunctionCall(source)(arg_0, arg_1)
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}
