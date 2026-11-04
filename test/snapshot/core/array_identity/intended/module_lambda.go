package main

import "snapshot/hxrt"

func Lambda_has(it hxrt.Anonymous, expected any) bool {
	found := false
	value := func() hxrt.Anonymous {
		var hx_callable_1 *hxrt.Function[func() hxrt.Anonymous] = hx_recover_callable_0(hxrt.AnonymousGet(it, "iterator"))
		return hxrt.FunctionCall(hx_callable_1)()
	}()
	for func() bool {
		var hx_callable_2 *hxrt.Function[func() bool] = hx_recover_callable_1(hxrt.AnonymousGet(value, "hasNext"))
		return hxrt.FunctionCall(hx_callable_2)()
	}() {
		var value_1 any = func() any {
			var hx_callable_3 *hxrt.Function[func() any] = hx_recover_callable_2(hxrt.AnonymousGet(value, "next"))
			return hxrt.FunctionCall(hx_callable_3)()
		}()
		if hxrt.HaxeEqual(value_1, expected) {
			found = true
			break
		}
	}
	return found
}
