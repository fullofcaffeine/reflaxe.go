package main

import "snapshot/hxrt"

func main() {
	expression := New_EReg(hxrt.StringFromLiteral("a"), hxrt.StringFromLiteral("g"))
	var v any = any(expression.__hx_this.replace(hxrt.StringFromLiteral("a-a"), hxrt.StringFromLiteral("b")))
	hxrt.Println(v)
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
