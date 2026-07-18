package main

import (
	"math"
	"snapshot/hxrt"
)

func invalidDateThrows() bool {
	hx_try_return_1 := false
	var hx_try_value_2 bool
	hxrt.TryCatch(func() {
		Date_fromString(hxrt.StringFromLiteral("not-a-date"))
	}, func(hx_caught_3 any) {
		hx_tmp := hx_caught_3
		_ = hx_tmp
		hx_try_value_2 = true
		hx_try_return_1 = true
		return
	})
	if hx_try_return_1 {
		return hx_try_value_2
	}
	return false
}

func main() {
	value := New_Date(2024, 1, 29, 12, 34, 56)
	var v any = any(value.__hx_this.toString())
	hxrt.Println(v)
	var v_1 any = any((func() bool {
		f := math.Sqrt(4.0)
		return (!math.IsInf(f, 0) && !math.IsNaN(f))
	}() && (Math_min(Math_PI, 4.0) == Math_PI)))
	hxrt.Println(v_1)
	var v_2 any = any(invalidDateThrows())
	hxrt.Println(v_2)
}
