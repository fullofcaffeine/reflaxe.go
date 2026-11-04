package main

import "snapshot/hxrt"

func describe(mode any) *string {
	var hx_switch_1 *string
	switch *hxrt.StdString(mode) {
	case *hxrt.StdString(hxrt.StringFromLiteral("human")):
		hx_switch_1 = hxrt.StringFromLiteral("human")
	case *hxrt.StdString(hxrt.StringFromLiteral("json")):
		hx_switch_1 = hxrt.StringFromLiteral("json")
	}
	return hx_switch_1
}

func main() {
	if !hxrt.StringEqualStringPtr(describe(any(hxrt.StringFromLiteral("human"))), hxrt.StringFromLiteral("human")) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("human enum-abstract switch mismatch")))
	}
	if !hxrt.StringEqualStringPtr(describe(any(hxrt.StringFromLiteral("json"))), hxrt.StringFromLiteral("json")) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("json enum-abstract switch mismatch")))
	}
	statementMatched := false
	var _g any = any(hxrt.StringFromLiteral("json"))
	switch *hxrt.StdString(_g) {
	case *hxrt.StdString(hxrt.StringFromLiteral("human")):
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("statement switch matched the wrong string value")))
	case *hxrt.StdString(hxrt.StringFromLiteral("json")):
		statementMatched = true
	}
	if !statementMatched {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("statement enum-abstract switch missed")))
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
