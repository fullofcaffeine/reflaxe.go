package main

import (
	"encoding/json"
	"snapshot/hxrt"
)

func keepForeign(value any) any {
	return value
}

func keepNumber(value *json.Number) *json.Number {
	return value
}

func main() {
	if keepNumber(nil) != nil {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("native carrier null changed")))
	}
	if !hxrt.AnyEqualsNull(keepForeign(nil)) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("extern abstract null changed")))
	}
	hxrt.Println(any(hxrt.StringFromLiteral("string abstract native boundary: PASS")))
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
