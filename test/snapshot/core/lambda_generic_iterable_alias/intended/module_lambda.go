package main

import "snapshot/hxrt"

func Lambda_count(it hxrt.Anonymous, predicate *hxrt.Function[func(any) bool]) int {
	result := 0
	if predicate == nil {
		hx_tmp := func() hxrt.Anonymous {
			var hx_callable_1 *hxrt.Function[func() hxrt.Anonymous] = hx_recover_callable_0(hxrt.AnonymousGet(it, "iterator"))
			return hxrt.FunctionCall(hx_callable_1)()
		}()
		for func() bool {
			var hx_callable_2 *hxrt.Function[func() bool] = hx_recover_callable_1(hxrt.AnonymousGet(hx_tmp, "hasNext"))
			return hxrt.FunctionCall(hx_callable_2)()
		}() {
			var hx_tmp_1 any = func() any {
				var hx_callable_3 *hxrt.Function[func() any] = hx_recover_callable_2(hxrt.AnonymousGet(hx_tmp, "next"))
				return hxrt.FunctionCall(hx_callable_3)()
			}()
			_ = hx_tmp_1
			result = int(int32((result + 1)))
		}
	} else {
		value := func() hxrt.Anonymous {
			var hx_callable_4 *hxrt.Function[func() hxrt.Anonymous] = hx_recover_callable_0(hxrt.AnonymousGet(it, "iterator"))
			return hxrt.FunctionCall(hx_callable_4)()
		}()
		for func() bool {
			var hx_callable_5 *hxrt.Function[func() bool] = hx_recover_callable_1(hxrt.AnonymousGet(value, "hasNext"))
			return hxrt.FunctionCall(hx_callable_5)()
		}() {
			var value_1 any = func() any {
				var hx_callable_6 *hxrt.Function[func() any] = hx_recover_callable_2(hxrt.AnonymousGet(value, "next"))
				return hxrt.FunctionCall(hx_callable_6)()
			}()
			if func() bool {
				var hx_callable_7 *hxrt.Function[func(any) bool] = predicate
				var hx_argument_8 any = value_1
				return hxrt.FunctionCall(hx_callable_7)(hx_argument_8)
			}() {
				result = int(int32((result + 1)))
			}
		}
	}
	return result
}

func Lambda_empty(it hxrt.Anonymous) bool {
	result := true
	hx_tmp := func() hxrt.Anonymous {
		var hx_callable_9 *hxrt.Function[func() hxrt.Anonymous] = hx_recover_callable_0(hxrt.AnonymousGet(it, "iterator"))
		return hxrt.FunctionCall(hx_callable_9)()
	}()
	for func() bool {
		var hx_callable_10 *hxrt.Function[func() bool] = hx_recover_callable_1(hxrt.AnonymousGet(hx_tmp, "hasNext"))
		return hxrt.FunctionCall(hx_callable_10)()
	}() {
		var hx_tmp_1 any = func() any {
			var hx_callable_11 *hxrt.Function[func() any] = hx_recover_callable_2(hxrt.AnonymousGet(hx_tmp, "next"))
			return hxrt.FunctionCall(hx_callable_11)()
		}()
		_ = hx_tmp_1
		result = false
		break
	}
	return result
}

func Lambda_exists(it hxrt.Anonymous, predicate *hxrt.Function[func(any) bool]) bool {
	found := false
	value := func() hxrt.Anonymous {
		var hx_callable_12 *hxrt.Function[func() hxrt.Anonymous] = hx_recover_callable_0(hxrt.AnonymousGet(it, "iterator"))
		return hxrt.FunctionCall(hx_callable_12)()
	}()
	for func() bool {
		var hx_callable_13 *hxrt.Function[func() bool] = hx_recover_callable_1(hxrt.AnonymousGet(value, "hasNext"))
		return hxrt.FunctionCall(hx_callable_13)()
	}() {
		var value_1 any = func() any {
			var hx_callable_14 *hxrt.Function[func() any] = hx_recover_callable_2(hxrt.AnonymousGet(value, "next"))
			return hxrt.FunctionCall(hx_callable_14)()
		}()
		if func() bool {
			var hx_callable_15 *hxrt.Function[func(any) bool] = predicate
			var hx_argument_16 any = value_1
			return hxrt.FunctionCall(hx_callable_15)(hx_argument_16)
		}() {
			found = true
			break
		}
	}
	return found
}

func Lambda_filter(it hxrt.Anonymous, predicate *hxrt.Function[func(any) bool]) *hxrt.Array {
	out := hxrt.NewArray()
	value := func() hxrt.Anonymous {
		var hx_callable_17 *hxrt.Function[func() hxrt.Anonymous] = hx_recover_callable_0(hxrt.AnonymousGet(it, "iterator"))
		return hxrt.FunctionCall(hx_callable_17)()
	}()
	for func() bool {
		var hx_callable_18 *hxrt.Function[func() bool] = hx_recover_callable_1(hxrt.AnonymousGet(value, "hasNext"))
		return hxrt.FunctionCall(hx_callable_18)()
	}() {
		var value_1 any = func() any {
			var hx_callable_19 *hxrt.Function[func() any] = hx_recover_callable_2(hxrt.AnonymousGet(value, "next"))
			return hxrt.FunctionCall(hx_callable_19)()
		}()
		if func() bool {
			var hx_callable_21 *hxrt.Function[func(any) bool] = predicate
			var hx_argument_22 any = value_1
			return hxrt.FunctionCall(hx_callable_21)(hx_argument_22)
		}() {
			out.Push(value_1)
		}
	}
	return out
}

func Lambda_fold(it hxrt.Anonymous, combine *hxrt.Function[func(any, any) any], first any) any {
	value := func() hxrt.Anonymous {
		var hx_callable_23 *hxrt.Function[func() hxrt.Anonymous] = hx_recover_callable_0(hxrt.AnonymousGet(it, "iterator"))
		return hxrt.FunctionCall(hx_callable_23)()
	}()
	for func() bool {
		var hx_callable_24 *hxrt.Function[func() bool] = hx_recover_callable_1(hxrt.AnonymousGet(value, "hasNext"))
		return hxrt.FunctionCall(hx_callable_24)()
	}() {
		var value_1 any = func() any {
			var hx_callable_25 *hxrt.Function[func() any] = hx_recover_callable_2(hxrt.AnonymousGet(value, "next"))
			return hxrt.FunctionCall(hx_callable_25)()
		}()
		first = func() any {
			var hx_callable_26 *hxrt.Function[func(any, any) any] = combine
			var hx_argument_27 any = value_1
			var hx_argument_28 any = first
			return hxrt.FunctionCall(hx_callable_26)(hx_argument_27, hx_argument_28)
		}()
	}
	return first
}

func Lambda_has(it hxrt.Anonymous, expected any) bool {
	found := false
	value := func() hxrt.Anonymous {
		var hx_callable_29 *hxrt.Function[func() hxrt.Anonymous] = hx_recover_callable_0(hxrt.AnonymousGet(it, "iterator"))
		return hxrt.FunctionCall(hx_callable_29)()
	}()
	for func() bool {
		var hx_callable_30 *hxrt.Function[func() bool] = hx_recover_callable_1(hxrt.AnonymousGet(value, "hasNext"))
		return hxrt.FunctionCall(hx_callable_30)()
	}() {
		var value_1 any = func() any {
			var hx_callable_31 *hxrt.Function[func() any] = hx_recover_callable_2(hxrt.AnonymousGet(value, "next"))
			return hxrt.FunctionCall(hx_callable_31)()
		}()
		if hxrt.HaxeEqual(value_1, expected) {
			found = true
			break
		}
	}
	return found
}

func Lambda_map_(it hxrt.Anonymous, transform *hxrt.Function[func(any) any]) *hxrt.Array {
	out := hxrt.NewArray()
	value := func() hxrt.Anonymous {
		var hx_callable_32 *hxrt.Function[func() hxrt.Anonymous] = hx_recover_callable_0(hxrt.AnonymousGet(it, "iterator"))
		return hxrt.FunctionCall(hx_callable_32)()
	}()
	for func() bool {
		var hx_callable_33 *hxrt.Function[func() bool] = hx_recover_callable_1(hxrt.AnonymousGet(value, "hasNext"))
		return hxrt.FunctionCall(hx_callable_33)()
	}() {
		var value_1 any = func() any {
			var hx_callable_34 *hxrt.Function[func() any] = hx_recover_callable_2(hxrt.AnonymousGet(value, "next"))
			return hxrt.FunctionCall(hx_callable_34)()
		}()
		out.Push(func() any {
			var hx_callable_36 *hxrt.Function[func(any) any] = transform
			var hx_argument_37 any = value_1
			return hxrt.FunctionCall(hx_callable_36)(hx_argument_37)
		}())
	}
	return out
}
