package main

import "snapshot/hxrt"

func main() {
	safe(hxrt.StringFromLiteral("fromCharCode"), hxrt.NewFunction(func() *string {
		str := hxrt.StringFromLiteral("A")
		var this1_1 *string
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("Ucs2 String not supported on this platform")))
		this1_1 = str
		this1 := hxrt.StdString(any(this1_1))
		return this1
	}))
}

func safe(label *string, fn *hxrt.Function[func() *string]) {
	hxrt.TryCatch(func() {
		var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(label, hxrt.StringFromLiteral("=")), func() *string {
			var hx_callable_4 *hxrt.Function[func() *string] = fn
			return hxrt.FunctionCall(hx_callable_4)()
		}()))
		hxrt.Println(v_1)
	}, func(hx_caught_1 any) {
		hx_exception_2 := haxe__Exception_caught(hx_caught_1)
		hx_payload_3 := haxe__Exception_catchValue(hx_caught_1)
		_ = hx_exception_2
		var error any = hx_payload_3
		var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(label, hxrt.StringFromLiteral("=!")), hxrt.StdString(error)))
		hxrt.Println(v)
	})
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
