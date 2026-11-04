package main

import "snapshot/hxrt"

func invoke(values *hxrt.Array, index int) int {
	return func() int {
		var hx_callable_1 *hxrt.Function[func(int) int] = hx_recover_callable_0(values.Get(index))
		var hx_argument_2 int = 20
		return hxrt.FunctionCall(hx_callable_1)(hx_argument_2)
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
		var hx_callable_3 *hxrt.Function[func(int) int] = hx_recover_callable_0(boxed)
		var hx_argument_4 int = 11
		return hxrt.FunctionCall(hx_callable_3)(hx_argument_4)
	}())
	hxrt.Println(v_1)
}

func hx_recover_callable_0(value any) *hxrt.Function[func(int) int] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(int) int]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(int) int]:
		return source
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}
