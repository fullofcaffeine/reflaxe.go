package main

import "snapshot/hxrt"

func main() {
	loop := New_sys__net__Host(hxrt.StringFromLiteral("127.0.0.1"))
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("loop.to_nonempty="), hxrt.StdString(!hxrt.StringEqualStringPtr(loop.toString(), hxrt.StringFromLiteral("")))))
	hxrt.Println(v)
	named := New_sys__net__Host(hxrt.StringFromLiteral("localhost"))
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("named.to_nonempty="), hxrt.StdString(!hxrt.StringEqualStringPtr(named.toString(), hxrt.StringFromLiteral("")))))
	hxrt.Println(v_1)
	invalidThrows := false
	hxrt.TryCatch(func() {
		New_sys__net__Host(hxrt.StringFromLiteral("256.256.256.256"))
	}, func(hx_caught_1 any) {
		hx_tmp := hx_caught_1
		_ = hx_tmp
		invalidThrows = true
	})
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("invalid_throws="), hxrt.StdString(invalidThrows)))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("localhost_nonempty="), hxrt.StdString(!hxrt.StringEqualStringPtr(sys__net__Host_localhost(), hxrt.StringFromLiteral("")))))
	hxrt.Println(v_3)
	hxrt.TryCatch(func() {
		loop.reverse()
	}, func(hx_caught_3 any) {
		hx_tmp_1 := hx_caught_3
		_ = hx_tmp_1
	})
	hxrt.Println(any(hxrt.StringFromLiteral("reverse_called=true")))
}
