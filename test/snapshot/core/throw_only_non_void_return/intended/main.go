package main

import "snapshot/hxrt"

func failInt() int {
	hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("boom-int")))
	var hx_throw_zero_1 int
	return hx_throw_zero_1
}

func failString() *string {
	hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("boom-string")))
	var hx_throw_zero_2 *string
	return hx_throw_zero_2
}

func main() {
	var hx_try_3 int
	hxrt.TryCatch(func() {
		hx_try_3 = failInt()
	}, func(hx_caught_4 any) {
		hx_exception_5 := haxe__Exception_caught(hx_caught_4)
		hx_payload_6 := haxe__Exception_catchValue(hx_caught_4)
		_ = hx_exception_5
		var hx_tmp any = hx_payload_6
		_ = hx_tmp
		hx_try_3 = 7
	})
	a := hx_try_3
	var hx_try_7 *string
	hxrt.TryCatch(func() {
		hx_try_7 = failString()
	}, func(hx_caught_8 any) {
		hx_exception_9 := haxe__Exception_caught(hx_caught_8)
		hx_payload_10 := haxe__Exception_catchValue(hx_caught_8)
		_ = hx_exception_9
		var hx_tmp_1 any = hx_payload_10
		_ = hx_tmp_1
		hx_try_7 = hxrt.StringFromLiteral("ok")
	})
	b := hx_try_7
	var v any = any(hxrt.StdString(a))
	hxrt.Println(v)
	hxrt.Println(any(b))
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
