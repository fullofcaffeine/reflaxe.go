package main

import "snapshot/hxrt"

func main() {
	var hx_try_1 int
	hxrt.TryCatch(func() {
		hx_try_1 = risky(0)
	}, func(hx_caught_2 any) {
		hx_exception_3 := haxe__Exception_caught(hx_caught_2)
		hx_payload_4 := haxe__Exception_catchValue(hx_caught_2)
		_ = hx_exception_3
		var e any = hx_payload_4
		_ = e
		hx_try_1 = 11
	})
	a := hx_try_1
	var hx_try_5 int
	hxrt.TryCatch(func() {
		hx_try_5 = risky(4)
	}, func(hx_caught_6 any) {
		hx_exception_7 := haxe__Exception_caught(hx_caught_6)
		hx_payload_8 := haxe__Exception_catchValue(hx_caught_6)
		_ = hx_exception_7
		var e_1 any = hx_payload_8
		_ = e_1
		hx_try_5 = 11
	})
	b := hx_try_5
	var v any = any(hxrt.StdString(a))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StdString(b))
	hxrt.Println(v_1)
}

func risky(v int) int {
	if v == 0 {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("bad")))
	}
	return int((hxrt.Int32Wrap(v) + hxrt.Int32Wrap(1)))
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

func hx_is_nominal_0(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_1(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__Exception:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}
