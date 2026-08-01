package main

import "snapshot/hxrt"

func invoke(callback func(any) int) int {
	return callback(nil)
}

func main() {
	func(hx_fn func(func(), any), hx_arg_0 func(), hx_arg_1 any) {
		if hx_fn == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
			return
		}
		hx_fn(hx_arg_0, hx_arg_1)
	}(schedule, func() {
		func(hx_fn func(any, map[string]any), hx_arg_0 any, hx_arg_1 map[string]any) {
			if hx_fn == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
				return
			}
			hx_fn(hx_arg_0, hx_arg_1)
		}(haxe__Log_trace, hxrt.StringFromLiteral("scheduled"), func() map[string]any {
			hx_obj_1 := map[string]any{}
			hx_obj_1["fileName"] = hxrt.StringFromLiteral("Main.hx")
			hx_obj_1["lineNumber"] = 26
			hx_obj_1["className"] = hxrt.StringFromLiteral("Main")
			hx_obj_1["methodName"] = hxrt.StringFromLiteral("main")
			return hx_obj_1
		}())
	}, nil)
	func(hx_fn func(func(), any), hx_arg_0 func(), hx_arg_1 any) {
		if hx_fn == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
			return
		}
		hx_fn(hx_arg_0, hx_arg_1)
	}(schedule, func() {
	}, 10)
	func(hx_fn func(any, map[string]any), hx_arg_0 any, hx_arg_1 map[string]any) {
		if hx_fn == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
			return
		}
		hx_fn(hx_arg_0, hx_arg_1)
	}(haxe__Log_trace, func(hx_fn func(any, any, any) float64, hx_arg_0 any, hx_arg_1 any, hx_arg_2 any) float64 {
		if hx_fn == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
			var hx_null_call_zero_2 float64
			return hx_null_call_zero_2
		}
		return hx_fn(hx_arg_0, hx_arg_1, hx_arg_2)
	}(summarize, nil, nil, nil), func() map[string]any {
		hx_obj_3 := map[string]any{}
		hx_obj_3["fileName"] = hxrt.StringFromLiteral("Main.hx")
		hx_obj_3["lineNumber"] = 28
		hx_obj_3["className"] = hxrt.StringFromLiteral("Main")
		hx_obj_3["methodName"] = hxrt.StringFromLiteral("main")
		return hx_obj_3
	}())
	func(hx_fn func(any, map[string]any), hx_arg_0 any, hx_arg_1 map[string]any) {
		if hx_fn == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
			return
		}
		hx_fn(hx_arg_0, hx_arg_1)
	}(haxe__Log_trace, func(hx_fn func(any, any, any) float64, hx_arg_0 any, hx_arg_1 any, hx_arg_2 any) float64 {
		if hx_fn == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
			var hx_null_call_zero_4 float64
			return hx_null_call_zero_4
		}
		return hx_fn(hx_arg_0, hx_arg_1, hx_arg_2)
	}(summarize, false, 9, 2.5), func() map[string]any {
		hx_obj_5 := map[string]any{}
		hx_obj_5["fileName"] = hxrt.StringFromLiteral("Main.hx")
		hx_obj_5["lineNumber"] = 29
		hx_obj_5["className"] = hxrt.StringFromLiteral("Main")
		hx_obj_5["methodName"] = hxrt.StringFromLiteral("main")
		return hx_obj_5
	}())
	summarize = func(enabled any, count any, ratio any) float64 {
		if enabled == nil {
			enabled = false
		}
		if count == nil {
			count = 4
		}
		if ratio == nil {
			ratio = 2.0
		}
		var hx_if_6 float64
		if enabled.(bool) {
			hx_if_6 = (float64(count.(int)) + ratio.(float64))
		} else {
			hx_if_6 = ratio.(float64)
		}
		return hx_if_6
	}
	func(hx_fn func(any, map[string]any), hx_arg_0 any, hx_arg_1 map[string]any) {
		if hx_fn == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
			return
		}
		hx_fn(hx_arg_0, hx_arg_1)
	}(haxe__Log_trace, func(hx_fn func(any, any, any) float64, hx_arg_0 any, hx_arg_1 any, hx_arg_2 any) float64 {
		if hx_fn == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
			var hx_null_call_zero_7 float64
			return hx_null_call_zero_7
		}
		return hx_fn(hx_arg_0, hx_arg_1, hx_arg_2)
	}(summarize, false, 8, 3.5), func() map[string]any {
		hx_obj_8 := map[string]any{}
		hx_obj_8["fileName"] = hxrt.StringFromLiteral("Main.hx")
		hx_obj_8["lineNumber"] = 34
		hx_obj_8["className"] = hxrt.StringFromLiteral("Main")
		hx_obj_8["methodName"] = hxrt.StringFromLiteral("main")
		return hx_obj_8
	}())
	var v any = any(func(hx_fn func(any, any, any) float64, hx_arg_0 any, hx_arg_1 any, hx_arg_2 any) float64 {
		if hx_fn == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
			var hx_null_call_zero_9 float64
			return hx_null_call_zero_9
		}
		return hx_fn(hx_arg_0, hx_arg_1, hx_arg_2)
	}(summarize, nil, nil, nil))
	hxrt.Println(v)
	var local func(any, any) int = nil
	local = func(enabled any, count any) int {
		if enabled == nil {
			enabled = true
		}
		if count == nil {
			count = 6
		}
		var hx_if_10 int
		if enabled.(bool) {
			hx_if_10 = count.(int)
		} else {
			hx_if_10 = int(int32(-int32(count.(int))))
		}
		return hx_if_10
	}
	func(hx_fn func(any, map[string]any), hx_arg_0 any, hx_arg_1 map[string]any) {
		if hx_fn == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
			return
		}
		hx_fn(hx_arg_0, hx_arg_1)
	}(haxe__Log_trace, local(false, 11), func() map[string]any {
		hx_obj_11 := map[string]any{}
		hx_obj_11["fileName"] = hxrt.StringFromLiteral("Main.hx")
		hx_obj_11["lineNumber"] = 41
		hx_obj_11["className"] = hxrt.StringFromLiteral("Main")
		hx_obj_11["methodName"] = hxrt.StringFromLiteral("main")
		return hx_obj_11
	}())
	var initialized func(any, any) int = func(enabled any, count any) int {
		if enabled == nil {
			enabled = true
		}
		if count == nil {
			count = 6
		}
		var hx_if_12 int
		if enabled.(bool) {
			hx_if_12 = count.(int)
		} else {
			hx_if_12 = int(int32(-int32(count.(int))))
		}
		return hx_if_12
	}
	var v_1 any = any(initialized(nil, nil))
	hxrt.Println(v_1)
	initialized = func(enabled any, count any) int {
		if enabled == nil {
			enabled = false
		}
		if count == nil {
			count = 4
		}
		var hx_if_13 int
		if enabled.(bool) {
			hx_if_13 = count.(int)
		} else {
			hx_if_13 = int(int32(-int32(count.(int))))
		}
		return hx_if_13
	}
	var v_2 any = any(initialized(nil, nil))
	hxrt.Println(v_2)
	casted := func(count int) int {
		return count
	}
	var v_3 any = any(casted(2))
	hxrt.Println(v_3)
	hx_callable_14 := namedDefault
	var namedCarrier func(any) int = func(count any) int {
		if count == nil {
			count = 8
		}
		return hx_callable_14(hxrt.IntFromNullableAny(count))
	}
	var v_4 any = any(namedCarrier(nil))
	hxrt.Println(v_4)
	localSource := func(count int) int {
		return count
	}
	hx_callable_15 := localSource
	var localAlias func(any) int = func(count any) int {
		if count == nil {
			count = 5
		}
		return hx_callable_15(hxrt.IntFromNullableAny(count))
	}
	var v_5 any = any(localAlias(nil))
	hxrt.Println(v_5)
	var v_6 any = any(invoke(func() func(any) int {
		hx_callable_16 := localSource
		return func(count any) int {
			if count == nil {
				count = 5
			}
			return hx_callable_16(hxrt.IntFromNullableAny(count))
		}
	}()))
	hxrt.Println(v_6)
	var v_7 any = any(invoke(func(count any) int {
		if count == nil {
			count = 12
		}
		return count.(int)
	}))
	hxrt.Println(v_7)
	var v_8 any = any(makeLocalCarrier()(nil))
	hxrt.Println(v_8)
	localOther := func(count int) int {
		return count
	}
	var hx_if_19 func(any) int
	if localAlias(nil) == 5 {
		hx_callable_17 := localSource
		hx_if_19 = func(count any) int {
			if count == nil {
				count = 5
			}
			return hx_callable_17(hxrt.IntFromNullableAny(count))
		}
	} else {
		hx_callable_18 := localOther
		hx_if_19 = func(count any) int {
			if count == nil {
				count = 6
			}
			return hx_callable_18(hxrt.IntFromNullableAny(count))
		}
	}
	conditional := hx_if_19
	var v_9 any = any(conditional(nil))
	hxrt.Println(v_9)
	_g := localAlias(nil)
	var hx_if_22 func(any) int
	if _g == 5 {
		hx_callable_20 := localOther
		hx_if_22 = func(count any) int {
			if count == nil {
				count = 6
			}
			return hx_callable_20(hxrt.IntFromNullableAny(count))
		}
	} else {
		hx_callable_21 := localSource
		hx_if_22 = func(count any) int {
			if count == nil {
				count = 5
			}
			return hx_callable_21(hxrt.IntFromNullableAny(count))
		}
	}
	switched := hx_if_22
	var v_10 any = any(switched(nil))
	hxrt.Println(v_10)
	observed := localAlias(nil)
	if observed != 5 {
		hxrt.Throw(hxrt.StringFromLiteral("unexpected"))
	}
	hx_callable_23 := localSource
	blocked := func(count any) int {
		if count == nil {
			count = 5
		}
		return hx_callable_23(hxrt.IntFromNullableAny(count))
	}
	var v_11 any = any(blocked(nil))
	hxrt.Println(v_11)
	var hx_try_24 func(any) int
	hxrt.TryCatch(func() {
		hx_callable_25 := localOther
		hx_try_24 = func(count any) int {
			if count == nil {
				count = 6
			}
			return hx_callable_25(hxrt.IntFromNullableAny(count))
		}
	}, func(hx_caught_26 any) {
		hx_tmp := hx_caught_26
		_ = hx_tmp
		hx_callable_28 := localSource
		hx_try_24 = func(count any) int {
			if count == nil {
				count = 5
			}
			return hx_callable_28(hxrt.IntFromNullableAny(count))
		}
	})
	caught := hx_try_24
	var v_12 any = any(caught(nil))
	hxrt.Println(v_12)
	var callbacks_0 func(any) int
	hx_callable_29 := localSource
	callbacks_0 = func(count any) int {
		if count == nil {
			count = 5
		}
		return hx_callable_29(hxrt.IntFromNullableAny(count))
	}
	var v_13 any = any(callbacks_0(nil))
	hxrt.Println(v_13)
	hx_callable_30 := localOther
	callbacks_0 = func(count any) int {
		if count == nil {
			count = 6
		}
		return hx_callable_30(hxrt.IntFromNullableAny(count))
	}
	var v_14 any = any(callbacks_0(nil))
	hxrt.Println(v_14)
	var holder_callback func(any) int
	hx_callable_31 := localSource
	holder_callback = func(count any) int {
		if count == nil {
			count = 5
		}
		return hx_callable_31(hxrt.IntFromNullableAny(count))
	}
	var v_15 any = any(holder_callback(nil))
	hxrt.Println(v_15)
}

func makeLocalCarrier() func(any) int {
	source := func(count int) int {
		return count
	}
	hx_callable_32 := source
	return func(count any) int {
		if count == nil {
			count = 10
		}
		return hx_callable_32(hxrt.IntFromNullableAny(count))
	}
}

func namedDefault(count int) int {
	return count
}

var schedule func(func(), any) = func(callback func(), timeout any) {
	callback()
}

var summarize func(any, any, any) float64 = func(enabled any, count any, ratio any) float64 {
	if enabled == nil {
		enabled = true
	}
	if count == nil {
		count = 3
	}
	if ratio == nil {
		ratio = 1.5
	}
	var hx_if_33 float64
	if enabled.(bool) {
		hx_if_33 = (float64(count.(int)) + ratio.(float64))
	} else {
		hx_if_33 = ratio.(float64)
	}
	return hx_if_33
}
