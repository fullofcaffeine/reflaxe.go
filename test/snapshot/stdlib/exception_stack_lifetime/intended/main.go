package main

import "snapshot/hxrt"

func firstOrigin() *haxe__Exception {
	return New_haxe__Exception(hxrt.StringFromLiteral("first"), nil, nil)
}

func main() {
	first := firstOrigin()
	second := secondOrigin()
	initial := haxe___CallStack__CallStack_Impl__toString(first.__hx_this.get_stack())
	later := readLater(first)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("captured="), hxrt.StdString((hxrt.StringIndexOfStringPtr(initial, hxrt.StringFromLiteral("firstOrigin"), 0, false) >= 0))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("stable="), hxrt.StdString(hxrt.StringEqualStringPtr(initial, later))))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("separate="), hxrt.StdString((hxrt.StringIndexOfStringPtr(readLater(second), hxrt.StringFromLiteral("secondOrigin"), 0, false) >= 0))))
	hxrt.Println(v_2)
	hxrt.TryCatch(func() {
		hxrt.Throw(haxe__Exception_thrown(first))
	}, func(hx_caught_1 any) {
		hx_exception_2 := haxe__Exception_caught(hx_caught_1)
		hx_payload_3 := haxe__Exception_catchValue(hx_caught_1)
		_ = hx_payload_3
		var caught *haxe__Exception = hx_recover_class_0(hx_exception_2)
		var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("rethrow-same="), hxrt.StdString(((caught == first) && hxrt.StringEqualStringPtr(readLater(caught), initial)))))
		hxrt.Println(v_3)
	})
}

func readLater(error *haxe__Exception) *string {
	return haxe___CallStack__CallStack_Impl__toString(error.__hx_this.get_stack())
}

func secondOrigin() *haxe__Exception {
	return New_haxe__Exception(hxrt.StringFromLiteral("second"), nil, nil)
}

func hx_recover_class_0(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ValueException:
		return source.haxe__Exception
	case *haxe__Exception:
		return source
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
	case *haxe__ValueException:
		return (source != nil)
	case *haxe__Exception:
		return (source != nil)
	default:
		return false
	}
}
