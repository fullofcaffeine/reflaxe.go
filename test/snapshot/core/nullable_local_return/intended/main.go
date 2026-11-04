package main

import "snapshot/hxrt"

func decimal(value any) float64 {
	var hx_if_1 any
	if value == nil {
		hx_if_1 = 1.5
	} else {
		hx_if_1 = value.(float64)
	}
	var result any = hx_if_1
	return func(hx_value_2 any) float64 {
		if hx_value_2 == nil {
			var hx_zero_3 float64
			return hx_zero_3
		}
		return hx_value_2.(float64)
	}(result)
}

func flag(value any) bool {
	var hx_if_4 any
	if value == nil {
		hx_if_4 = true
	} else {
		hx_if_4 = value.(bool)
	}
	var result any = hx_if_4
	return func(hx_value_5 any) bool {
		if hx_value_5 == nil {
			var hx_zero_6 bool
			return hx_zero_6
		}
		return hx_value_5.(bool)
	}(result)
}

func integer(value any) int {
	var hx_if_7 any
	if value == nil {
		hx_if_7 = 9
	} else {
		hx_if_7 = value.(int)
	}
	var result any = hx_if_7
	return hxrt.IntFromNullableAny(result)
}

func main() {
	var v any = any(integer(nil))
	hxrt.Println(v)
	var v_1 any = any(integer(0))
	hxrt.Println(v_1)
	var v_2 any = any(integer(4))
	hxrt.Println(v_2)
	var v_3 any = any(decimal(nil))
	hxrt.Println(v_3)
	var v_4 any = any(decimal(0.0))
	hxrt.Println(v_4)
	var v_5 any = any(decimal(2.5))
	hxrt.Println(v_5)
	var v_6 any = any(flag(nil))
	hxrt.Println(v_6)
	var v_7 any = any(flag(false))
	hxrt.Println(v_7)
	var v_8 any = any(flag(true))
	hxrt.Println(v_8)
	var v_9 any = any(redirected(nil))
	hxrt.Println(v_9)
	var v_10 any = any(redirected(0))
	hxrt.Println(v_10)
	var v_11 any = any((nullable(nil) == nil))
	hxrt.Println(v_11)
	var v_12 any = any(nullable(3))
	hxrt.Println(v_12)
}

func nullable(value any) any {
	var result any = value
	return result
}

func redirected(value any) int {
	hx_try_return_8 := false
	var hx_try_value_9 int
	hxrt.TryCatch(func() {
		var hx_if_13 any
		if value == nil {
			hx_if_13 = 7
		} else {
			hx_if_13 = value.(int)
		}
		var result any = hx_if_13
		hx_try_value_9 = hxrt.IntFromNullableAny(result)
		hx_try_return_8 = true
		return
	}, func(hx_caught_10 any) {
		hx_exception_11 := haxe__Exception_caught(hx_caught_10)
		hx_payload_12 := haxe__Exception_catchValue(hx_caught_10)
		_ = hx_payload_12
		var hx_tmp *haxe__Exception = hx_recover_class_0(hx_exception_11)
		_ = hx_tmp
		hx_try_value_9 = -1
		hx_try_return_8 = true
		return
	})
	if hx_try_return_8 {
		return hx_try_value_9
	}
	var hx_try_zero_14 int
	return hx_try_zero_14
}

func hx_recover_class_0(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__Exception:
		return source
	case *haxe__ValueException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}

func hx_recover_class_1(value any) *haxe__ValueException {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ValueException:
		return source
	default:
		return value.(*haxe__ValueException)
	}
}
