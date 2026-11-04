package main

import "snapshot/hxrt"

func expect(actual int, expected int) {
	if actual != expected {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("expected "), expected), hxrt.StringFromLiteral(", got ")), actual)))
	}
}

func main() {
	value := hxrt.StringFromLiteral("a😀bé😀")
	expect(hxrt.StringIndexOfStringPtr(value, hxrt.StringFromLiteral("😀"), 0, false), 1)
	expect(hxrt.StringIndexOfStringPtr(value, hxrt.StringFromLiteral("😀"), 2, true), 4)
	expect(hxrt.StringIndexOfStringPtr(value, hxrt.StringFromLiteral(""), 3, true), 3)
	expect(hxrt.StringIndexOfStringPtr(value, hxrt.StringFromLiteral(""), 99, true), 5)
	expect(hxrt.StringIndexOfStringPtr(value, hxrt.StringFromLiteral("😀"), -1, true), 4)
	expect(hxrt.StringIndexOfStringPtr(value, hxrt.StringFromLiteral(""), -1, true), 0)
	expect(hxrt.StringIndexOfStringPtr(value, hxrt.StringFromLiteral("😀"), -99, true), 1)
	expect(hxrt.StringIndexOfStringPtr(value, hxrt.StringFromLiteral("missing"), 0, false), -1)
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
