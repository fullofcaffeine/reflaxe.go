package main

import "snapshot/hxrt"

func main() {
	if (!hxrt.StringEqualStringPtr(_GuardedTry__GuardedTry_Fields__readValue(true, false), hxrt.StringFromLiteral("")) || !hxrt.StringEqualStringPtr(_GuardedTry__GuardedTry_Fields__readValue(false, false), hxrt.StringFromLiteral("value"))) || !hxrt.StringEqualStringPtr(_GuardedTry__GuardedTry_Fields__readValue(false, true), hxrt.StringFromLiteral("fallback")) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("unexpected result")))
	}
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
