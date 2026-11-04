package main

import "snapshot/hxrt"

func main() {
	s := hxrt.StringFromLiteral("  hi  ")
	var v any = any(StringTools_trim(s))
	hxrt.Println(v)
	var v_1 any = any(StringTools_startsWith(hxrt.StringFromLiteral("hello"), hxrt.StringFromLiteral("he")))
	hxrt.Println(v_1)
	var v_2 any = any(StringTools_replace(hxrt.StringFromLiteral("a-b-c"), hxrt.StringFromLiteral("-"), hxrt.StringFromLiteral(":")))
	hxrt.Println(v_2)
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
