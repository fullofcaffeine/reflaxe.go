package main

import (
	"log"
	"snapshot/hxrt"
)

func main() {
	hxrt.Println(any(hxrt.StringFromLiteral("native-start")))
	hxrt.TryCatch(func() {
		log.Panic(*hxrt.StdString(hxrt.StringFromLiteral("native-failure")))
	}, func(hx_caught_1 any) {
		hx_exception_2 := haxe__Exception_caught(hx_caught_1)
		hx_payload_3 := haxe__Exception_catchValue(hx_caught_1)
		_ = hx_payload_3
		var error *haxe__Exception = hx_recover_class_0(hx_exception_2)
		var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("incorrectly-caught="), hxrt.StdString(error)))
		hxrt.Println(v)
	})
}

func hx_recover_class_0(value any) *haxe__Exception {
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
