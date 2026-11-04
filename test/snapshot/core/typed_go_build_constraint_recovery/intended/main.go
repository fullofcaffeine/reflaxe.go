package main

import "snapshot/hxrt"

func main() {
	if !hxrt.StringEqualStringPtr(Registry_selected, hxrt.StringFromLiteral("ok")) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("selected recovery did not run")))
	}
	var v any = any(Registry_selected)
	hxrt.Println(v)
}

func hx_recover_class_2(value any) *haxe__ValueException {
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

func hx_recover_class_3(value any) *haxe__Exception {
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
