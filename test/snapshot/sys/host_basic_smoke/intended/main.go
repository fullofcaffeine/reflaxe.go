package main

import "snapshot/hxrt"

func main() {
	loop := New_sys__net__Host(hxrt.StringFromLiteral("127.0.0.1"))
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("loop.to_nonempty="), hxrt.StdString(!hxrt.StringEqualStringPtr(loop.__hx_this.toString(), hxrt.StringFromLiteral("")))))
	hxrt.Println(v)
	named := New_sys__net__Host(hxrt.StringFromLiteral("localhost"))
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("named.to_nonempty="), hxrt.StdString(!hxrt.StringEqualStringPtr(named.__hx_this.toString(), hxrt.StringFromLiteral("")))))
	hxrt.Println(v_1)
	invalidThrows := false
	hxrt.TryCatch(func() {
		New_sys__net__Host(hxrt.StringFromLiteral("256.256.256.256"))
	}, func(hx_caught_1 any) {
		hx_exception_2 := haxe__Exception_caught(hx_caught_1)
		hx_payload_3 := haxe__Exception_catchValue(hx_caught_1)
		_ = hx_exception_2
		var hx_tmp any = hx_payload_3
		_ = hx_tmp
		invalidThrows = true
	})
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("invalid_throws="), hxrt.StdString(invalidThrows)))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("localhost_nonempty="), hxrt.StdString(!hxrt.StringEqualStringPtr(sys__net__Host_localhost(), hxrt.StringFromLiteral("")))))
	hxrt.Println(v_3)
	hxrt.TryCatch(func() {
		loop.__hx_this.reverse()
	}, func(hx_caught_4 any) {
		hx_exception_5 := haxe__Exception_caught(hx_caught_4)
		hx_payload_6 := haxe__Exception_catchValue(hx_caught_4)
		_ = hx_exception_5
		var hx_tmp_1 any = hx_payload_6
		_ = hx_tmp_1
	})
	hxrt.Println(any(hxrt.StringFromLiteral("reverse_called=true")))
}

func hx_recover_class_0(value any) *haxe__ValueException {
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

func hx_recover_class_1(value any) *haxe__Exception {
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
