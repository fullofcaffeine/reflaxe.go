package main

import "snapshot/hxrt"

func invoke(callback *hxrt.Function[func(any) int]) int {
	return func(hx_fn *hxrt.Function[func(any) int], hx_arg_0 any) int {
		return hxrt.FunctionCall(hx_fn)(hx_arg_0)
	}(callback, nil)
}

func main() {
	func(hx_fn *hxrt.Function[func(*hxrt.Function[func()], any)], hx_arg_0 *hxrt.Function[func()], hx_arg_1 any) {
		hxrt.FunctionCall(hx_fn)(hx_arg_0, hx_arg_1)
	}(schedule, hxrt.NewFunction(func() {
		func(hx_fn *hxrt.Function[func(any, map[string]any)], hx_arg_0 any, hx_arg_1 map[string]any) {
			hxrt.FunctionCall(hx_fn)(hx_arg_0, hx_arg_1)
		}(haxe__Log_trace, hxrt.StringFromLiteral("scheduled"), func() map[string]any {
			hx_obj_1 := map[string]any{}
			hx_obj_1["fileName"] = hxrt.StringFromLiteral("Main.hx")
			hx_obj_1["lineNumber"] = 26
			hx_obj_1["className"] = hxrt.StringFromLiteral("Main")
			hx_obj_1["methodName"] = hxrt.StringFromLiteral("main")
			return hx_obj_1
		}())
	}), nil)
	func(hx_fn *hxrt.Function[func(*hxrt.Function[func()], any)], hx_arg_0 *hxrt.Function[func()], hx_arg_1 any) {
		hxrt.FunctionCall(hx_fn)(hx_arg_0, hx_arg_1)
	}(schedule, hxrt.NewFunction(func() {
	}), 10)
	func(hx_fn *hxrt.Function[func(any, map[string]any)], hx_arg_0 any, hx_arg_1 map[string]any) {
		hxrt.FunctionCall(hx_fn)(hx_arg_0, hx_arg_1)
	}(haxe__Log_trace, func(hx_fn *hxrt.Function[func(any, any, any) float64], hx_arg_0 any, hx_arg_1 any, hx_arg_2 any) float64 {
		return hxrt.FunctionCall(hx_fn)(hx_arg_0, hx_arg_1, hx_arg_2)
	}(summarize, nil, nil, nil), func() map[string]any {
		hx_obj_2 := map[string]any{}
		hx_obj_2["fileName"] = hxrt.StringFromLiteral("Main.hx")
		hx_obj_2["lineNumber"] = 28
		hx_obj_2["className"] = hxrt.StringFromLiteral("Main")
		hx_obj_2["methodName"] = hxrt.StringFromLiteral("main")
		return hx_obj_2
	}())
	func(hx_fn *hxrt.Function[func(any, map[string]any)], hx_arg_0 any, hx_arg_1 map[string]any) {
		hxrt.FunctionCall(hx_fn)(hx_arg_0, hx_arg_1)
	}(haxe__Log_trace, func(hx_fn *hxrt.Function[func(any, any, any) float64], hx_arg_0 any, hx_arg_1 any, hx_arg_2 any) float64 {
		return hxrt.FunctionCall(hx_fn)(hx_arg_0, hx_arg_1, hx_arg_2)
	}(summarize, false, 9, 2.5), func() map[string]any {
		hx_obj_3 := map[string]any{}
		hx_obj_3["fileName"] = hxrt.StringFromLiteral("Main.hx")
		hx_obj_3["lineNumber"] = 29
		hx_obj_3["className"] = hxrt.StringFromLiteral("Main")
		hx_obj_3["methodName"] = hxrt.StringFromLiteral("main")
		return hx_obj_3
	}())
	summarize = hxrt.NewFunction(func(enabled any, count any, ratio any) float64 {
		if enabled == nil {
			enabled = false
		}
		if count == nil {
			count = 4
		}
		if ratio == nil {
			ratio = 2.0
		}
		var hx_if_4 float64
		if enabled.(bool) {
			hx_if_4 = (float64(count.(int)) + ratio.(float64))
		} else {
			hx_if_4 = ratio.(float64)
		}
		return hx_if_4
	})
	func(hx_fn *hxrt.Function[func(any, map[string]any)], hx_arg_0 any, hx_arg_1 map[string]any) {
		hxrt.FunctionCall(hx_fn)(hx_arg_0, hx_arg_1)
	}(haxe__Log_trace, func(hx_fn *hxrt.Function[func(any, any, any) float64], hx_arg_0 any, hx_arg_1 any, hx_arg_2 any) float64 {
		return hxrt.FunctionCall(hx_fn)(hx_arg_0, hx_arg_1, hx_arg_2)
	}(summarize, false, 8, 3.5), func() map[string]any {
		hx_obj_5 := map[string]any{}
		hx_obj_5["fileName"] = hxrt.StringFromLiteral("Main.hx")
		hx_obj_5["lineNumber"] = 34
		hx_obj_5["className"] = hxrt.StringFromLiteral("Main")
		hx_obj_5["methodName"] = hxrt.StringFromLiteral("main")
		return hx_obj_5
	}())
	var v any = any(func(hx_fn *hxrt.Function[func(any, any, any) float64], hx_arg_0 any, hx_arg_1 any, hx_arg_2 any) float64 {
		return hxrt.FunctionCall(hx_fn)(hx_arg_0, hx_arg_1, hx_arg_2)
	}(summarize, nil, nil, nil))
	hxrt.Println(v)
	var local *hxrt.Function[func(any, any) int] = nil
	local = hxrt.NewFunction(func(enabled any, count any) int {
		if enabled == nil {
			enabled = true
		}
		if count == nil {
			count = 6
		}
		var hx_if_6 int
		if enabled.(bool) {
			hx_if_6 = count.(int)
		} else {
			hx_if_6 = int(-int32(count.(int)))
		}
		return hx_if_6
	})
	func(hx_fn *hxrt.Function[func(any, map[string]any)], hx_arg_0 any, hx_arg_1 map[string]any) {
		hxrt.FunctionCall(hx_fn)(hx_arg_0, hx_arg_1)
	}(haxe__Log_trace, func(hx_fn *hxrt.Function[func(any, any) int], hx_arg_0 any, hx_arg_1 any) int {
		return hxrt.FunctionCall(hx_fn)(hx_arg_0, hx_arg_1)
	}(local, false, 11), func() map[string]any {
		hx_obj_7 := map[string]any{}
		hx_obj_7["fileName"] = hxrt.StringFromLiteral("Main.hx")
		hx_obj_7["lineNumber"] = 41
		hx_obj_7["className"] = hxrt.StringFromLiteral("Main")
		hx_obj_7["methodName"] = hxrt.StringFromLiteral("main")
		return hx_obj_7
	}())
	var initialized *hxrt.Function[func(any, any) int] = hxrt.NewFunction(func(enabled any, count any) int {
		if enabled == nil {
			enabled = true
		}
		if count == nil {
			count = 6
		}
		var hx_if_8 int
		if enabled.(bool) {
			hx_if_8 = count.(int)
		} else {
			hx_if_8 = int(-int32(count.(int)))
		}
		return hx_if_8
	})
	var v_1 any = any(func(hx_fn *hxrt.Function[func(any, any) int], hx_arg_0 any, hx_arg_1 any) int {
		return hxrt.FunctionCall(hx_fn)(hx_arg_0, hx_arg_1)
	}(initialized, nil, nil))
	hxrt.Println(v_1)
	initialized = hxrt.NewFunction(func(enabled any, count any) int {
		if enabled == nil {
			enabled = false
		}
		if count == nil {
			count = 4
		}
		var hx_if_9 int
		if enabled.(bool) {
			hx_if_9 = count.(int)
		} else {
			hx_if_9 = int(-int32(count.(int)))
		}
		return hx_if_9
	})
	var v_2 any = any(func(hx_fn *hxrt.Function[func(any, any) int], hx_arg_0 any, hx_arg_1 any) int {
		return hxrt.FunctionCall(hx_fn)(hx_arg_0, hx_arg_1)
	}(initialized, nil, nil))
	hxrt.Println(v_2)
	var casted *hxrt.Function[func(any) int] = hxrt.NewFunction(func(count any) int {
		if count == nil {
			count = 2
		}
		return count.(int)
	})
	var v_3 any = any(func(hx_fn *hxrt.Function[func(any) int], hx_arg_0 any) int {
		return hxrt.FunctionCall(hx_fn)(hx_arg_0)
	}(casted, 2))
	hxrt.Println(v_3)
	hx_callable_10 := hxrt.NamedFunction(namedDefault, "Main.namedDefault")
	var namedCarrier *hxrt.Function[func(any) int] = hxrt.AdaptFunction(hx_callable_10, func(count any) int {
		if count == nil {
			count = 8
		}
		return hxrt.FunctionCall(hx_callable_10)(hxrt.IntFromNullableAny(count))
	})
	var v_4 any = any(func(hx_fn *hxrt.Function[func(any) int], hx_arg_0 any) int {
		return hxrt.FunctionCall(hx_fn)(hx_arg_0)
	}(namedCarrier, nil))
	hxrt.Println(v_4)
	var localSource *hxrt.Function[func(any) int] = hxrt.NewFunction(func(count any) int {
		if count == nil {
			count = 5
		}
		return count.(int)
	})
	hx_callable_11 := localSource
	var localAlias *hxrt.Function[func(any) int] = hxrt.AdaptFunction(hx_callable_11, func(count any) int {
		if count == nil {
			count = 5
		}
		return hxrt.FunctionCall(hx_callable_11)(hxrt.IntFromNullableAny(count))
	})
	var v_5 any = any(func(hx_fn *hxrt.Function[func(any) int], hx_arg_0 any) int {
		return hxrt.FunctionCall(hx_fn)(hx_arg_0)
	}(localAlias, nil))
	hxrt.Println(v_5)
	var v_6 any = any(invoke(func() *hxrt.Function[func(any) int] {
		hx_callable_12 := localSource
		return hxrt.AdaptFunction(hx_callable_12, func(count any) int {
			if count == nil {
				count = 5
			}
			return hxrt.FunctionCall(hx_callable_12)(hxrt.IntFromNullableAny(count))
		})
	}()))
	hxrt.Println(v_6)
	var v_7 any = any(invoke(hxrt.NewFunction(func(count any) int {
		if count == nil {
			count = 12
		}
		return count.(int)
	})))
	hxrt.Println(v_7)
	var v_8 any = any(func(hx_fn *hxrt.Function[func(any) int], hx_arg_0 any) int {
		return hxrt.FunctionCall(hx_fn)(hx_arg_0)
	}(makeLocalCarrier(), nil))
	hxrt.Println(v_8)
	var localOther *hxrt.Function[func(any) int] = hxrt.NewFunction(func(count any) int {
		if count == nil {
			count = 6
		}
		return count.(int)
	})
	var hx_if_15 *hxrt.Function[func(any) int]
	if func(hx_fn *hxrt.Function[func(any) int], hx_arg_0 any) int {
		return hxrt.FunctionCall(hx_fn)(hx_arg_0)
	}(localAlias, nil) == 5 {
		hx_callable_13 := localSource
		hx_if_15 = hxrt.AdaptFunction(hx_callable_13, func(count any) int {
			if count == nil {
				count = 5
			}
			return hxrt.FunctionCall(hx_callable_13)(hxrt.IntFromNullableAny(count))
		})
	} else {
		hx_callable_14 := localOther
		hx_if_15 = hxrt.AdaptFunction(hx_callable_14, func(count any) int {
			if count == nil {
				count = 6
			}
			return hxrt.FunctionCall(hx_callable_14)(hxrt.IntFromNullableAny(count))
		})
	}
	var conditional *hxrt.Function[func(any) int] = hx_if_15
	var v_9 any = any(func(hx_fn *hxrt.Function[func(any) int], hx_arg_0 any) int {
		return hxrt.FunctionCall(hx_fn)(hx_arg_0)
	}(conditional, nil))
	hxrt.Println(v_9)
	_g := func(hx_fn *hxrt.Function[func(any) int], hx_arg_0 any) int {
		return hxrt.FunctionCall(hx_fn)(hx_arg_0)
	}(localAlias, nil)
	var hx_if_18 *hxrt.Function[func(any) int]
	if _g == 5 {
		hx_callable_16 := localOther
		hx_if_18 = hxrt.AdaptFunction(hx_callable_16, func(count any) int {
			if count == nil {
				count = 6
			}
			return hxrt.FunctionCall(hx_callable_16)(hxrt.IntFromNullableAny(count))
		})
	} else {
		hx_callable_17 := localSource
		hx_if_18 = hxrt.AdaptFunction(hx_callable_17, func(count any) int {
			if count == nil {
				count = 5
			}
			return hxrt.FunctionCall(hx_callable_17)(hxrt.IntFromNullableAny(count))
		})
	}
	var switched *hxrt.Function[func(any) int] = hx_if_18
	var v_10 any = any(func(hx_fn *hxrt.Function[func(any) int], hx_arg_0 any) int {
		return hxrt.FunctionCall(hx_fn)(hx_arg_0)
	}(switched, nil))
	hxrt.Println(v_10)
	observed := func(hx_fn *hxrt.Function[func(any) int], hx_arg_0 any) int {
		return hxrt.FunctionCall(hx_fn)(hx_arg_0)
	}(localAlias, nil)
	if observed != 5 {
		hxrt.Throw(hxrt.StringFromLiteral("unexpected"))
	}
	hx_callable_19 := localSource
	var blocked *hxrt.Function[func(any) int] = hxrt.AdaptFunction(hx_callable_19, func(count any) int {
		if count == nil {
			count = 5
		}
		return hxrt.FunctionCall(hx_callable_19)(hxrt.IntFromNullableAny(count))
	})
	var v_11 any = any(func(hx_fn *hxrt.Function[func(any) int], hx_arg_0 any) int {
		return hxrt.FunctionCall(hx_fn)(hx_arg_0)
	}(blocked, nil))
	hxrt.Println(v_11)
	var hx_try_20 *hxrt.Function[func(any) int]
	hxrt.TryCatch(func() {
		hx_callable_21 := localOther
		hx_try_20 = hxrt.AdaptFunction(hx_callable_21, func(count any) int {
			if count == nil {
				count = 6
			}
			return hxrt.FunctionCall(hx_callable_21)(hxrt.IntFromNullableAny(count))
		})
	}, func(hx_caught_22 any) {
		hx_tmp := hx_caught_22
		_ = hx_tmp
		hx_callable_24 := localSource
		hx_try_20 = hxrt.AdaptFunction(hx_callable_24, func(count any) int {
			if count == nil {
				count = 5
			}
			return hxrt.FunctionCall(hx_callable_24)(hxrt.IntFromNullableAny(count))
		})
	})
	var caught *hxrt.Function[func(any) int] = hx_try_20
	var v_12 any = any(func(hx_fn *hxrt.Function[func(any) int], hx_arg_0 any) int {
		return hxrt.FunctionCall(hx_fn)(hx_arg_0)
	}(caught, nil))
	hxrt.Println(v_12)
	var callbacks_0 *hxrt.Function[func(any) int]
	hx_callable_25 := localSource
	callbacks_0 = hxrt.AdaptFunction(hx_callable_25, func(count any) int {
		if count == nil {
			count = 5
		}
		return hxrt.FunctionCall(hx_callable_25)(hxrt.IntFromNullableAny(count))
	})
	var v_13 any = any(func(hx_fn *hxrt.Function[func(any) int], hx_arg_0 any) int {
		return hxrt.FunctionCall(hx_fn)(hx_arg_0)
	}(callbacks_0, nil))
	hxrt.Println(v_13)
	hx_callable_26 := localOther
	callbacks_0 = hxrt.AdaptFunction(hx_callable_26, func(count any) int {
		if count == nil {
			count = 6
		}
		return hxrt.FunctionCall(hx_callable_26)(hxrt.IntFromNullableAny(count))
	})
	var v_14 any = any(func(hx_fn *hxrt.Function[func(any) int], hx_arg_0 any) int {
		return hxrt.FunctionCall(hx_fn)(hx_arg_0)
	}(callbacks_0, nil))
	hxrt.Println(v_14)
	var holder_callback *hxrt.Function[func(any) int]
	hx_callable_27 := localSource
	holder_callback = hxrt.AdaptFunction(hx_callable_27, func(count any) int {
		if count == nil {
			count = 5
		}
		return hxrt.FunctionCall(hx_callable_27)(hxrt.IntFromNullableAny(count))
	})
	var v_15 any = any(func(hx_fn *hxrt.Function[func(any) int], hx_arg_0 any) int {
		return hxrt.FunctionCall(hx_fn)(hx_arg_0)
	}(holder_callback, nil))
	hxrt.Println(v_15)
}

func makeLocalCarrier() *hxrt.Function[func(any) int] {
	var source *hxrt.Function[func(any) int] = hxrt.NewFunction(func(count any) int {
		if count == nil {
			count = 10
		}
		return count.(int)
	})
	hx_callable_28 := source
	return hxrt.AdaptFunction(hx_callable_28, func(count any) int {
		if count == nil {
			count = 10
		}
		return hxrt.FunctionCall(hx_callable_28)(hxrt.IntFromNullableAny(count))
	})
}

func namedDefault(count int) int {
	return count
}

var schedule *hxrt.Function[func(*hxrt.Function[func()], any)] = hxrt.NewFunction(func(callback *hxrt.Function[func()], timeout any) {
	func(hx_fn *hxrt.Function[func()]) {
		hxrt.FunctionCall(hx_fn)()
	}(callback)
})

var summarize *hxrt.Function[func(any, any, any) float64] = hxrt.NewFunction(func(enabled any, count any, ratio any) float64 {
	if enabled == nil {
		enabled = true
	}
	if count == nil {
		count = 3
	}
	if ratio == nil {
		ratio = 1.5
	}
	var hx_if_29 float64
	if enabled.(bool) {
		hx_if_29 = (float64(count.(int)) + ratio.(float64))
	} else {
		hx_if_29 = ratio.(float64)
	}
	return hx_if_29
})
