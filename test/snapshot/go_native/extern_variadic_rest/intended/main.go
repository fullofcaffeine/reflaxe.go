package main

import (
	"fmt"
	"snapshot/hxrt"
)

func main() {
	rendered := hxrt.StdString(fmt.Sprint([]any{*hxrt.StdString(hxrt.StringFromLiteral("left")), 7, *hxrt.StdString(hxrt.StringFromLiteral("right"))}...))
	if !hxrt.StringEqualStringPtr(rendered, hxrt.StringFromLiteral("left7right")) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("variadic extern arguments were not expanded as native values")))
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
