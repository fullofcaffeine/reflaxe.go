package main

import "snapshot/hxrt"

func main() {
	var v any = any(withTry(true))
	hxrt.Println(v)
	var v_1 any = any(withTry(false))
	hxrt.Println(v_1)
}

func withTry(flag bool) *string {
	state := hxrt.StringFromLiteral("start")
	hx_try_return_1 := false
	var hx_try_value_2 *string
	hxrt.TryCatch(func() {
		state = hxrt.StringFromLiteral("try")
		if flag {
			hx_try_value_2 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("try:"), state)
			hx_try_return_1 = true
			return
		}
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("boom")))
	}, func(hx_caught_3 any) {
		hx_exception_4 := haxe__Exception_caught(hx_caught_3)
		hx_payload_5 := haxe__Exception_catchValue(hx_caught_3)
		_ = hx_exception_4
		var e any = hx_payload_5
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
