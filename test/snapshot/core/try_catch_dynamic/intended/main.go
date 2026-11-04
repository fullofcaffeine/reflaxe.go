package main

import "snapshot/hxrt"

func main() {
	hxrt.TryCatch(func() {
		var v any = any(mayFail(false))
		hxrt.Println(v)
		var v_1 any = any(mayFail(true))
		hxrt.Println(v_1)
	}, func(hx_caught_1 any) {
		hx_exception_2 := haxe__Exception_caught(hx_caught_1)
		hx_payload_3 := haxe__Exception_catchValue(hx_caught_1)
		_ = hx_exception_2
		var e any = hx_payload_3
		hxrt.Println(e)
	})
	hxrt.Println(any(9))
}

func mayFail(flag bool) int {
	if flag {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("boom")))
	}
	return 7
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
