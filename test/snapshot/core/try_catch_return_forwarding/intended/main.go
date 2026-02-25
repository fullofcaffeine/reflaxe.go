package main

import "snapshot/hxrt"

func main() {
	hxrt.Println(withTry(true))
	hxrt.Println(withTry(false))
}

func withTry(flag bool) *string {
	state := hxrt.StringFromLiteral("start")
	_ = state
	hx_try_return_1 := false
	var hx_try_value_2 *string
	hxrt.TryCatch(func() {
		state = hxrt.StringFromLiteral("try")
		if flag {
			hx_try_value_2 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("try:"), state)
			hx_try_return_1 = true
			return
		}
		hxrt.Throw(hxrt.StringFromLiteral("boom"))
	}, func(hx_caught_3 any) {
		e := hx_caught_3
		_ = e
		state = hxrt.StringFromLiteral("catch")
		hx_try_value_2 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("catch:"), state), hxrt.StringFromLiteral(":")), hxrt.StdString(e))
		hx_try_return_1 = true
		return
	})
	if hx_try_return_1 {
		return hx_try_value_2
	}
	state = hxrt.StringFromLiteral("tail")
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("tail:"), state)
}
