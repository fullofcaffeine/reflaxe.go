package main

import "snapshot/hxrt"

func main() {
	arguments := Sys_args()
	echo := ((arguments.Len() > 0) && hxrt.StringEqualAny(arguments.Get(0), hxrt.StringFromLiteral("echo")))
	hxrt.Print(any(hxrt.StringFromLiteral("ready|")))
	hxrt.TryCatch(func() {
		value := Sys_getChar(echo)
		hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("|"), value), hxrt.StringFromLiteral("|"))))
	}, func(hx_caught_1 any) {
		hx_exception_2 := haxe__Exception_caught(hx_caught_1)
		hx_payload_3 := haxe__Exception_catchValue(hx_caught_1)
		_ = hx_exception_2
		if hx_is_nominal_haxe__io__Eof(hx_payload_3) {
			var hx_tmp *haxe__io__Eof = hx_recover_class_haxe__io__Eof(hx_payload_3)
			_ = hx_tmp
			hxrt.Println(any(hxrt.StringFromLiteral("eof|")))
		} else {
			hxrt.Throw(hx_caught_1)
		}
	})
}

func hx_recover_class_haxe__Exception(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__Exception:
		return source
	case *haxe__ValueException:
		return source.haxe__Exception
	case *haxe__exceptions__NotImplementedException:
		return source.haxe__exceptions__PosException.haxe__Exception
	case *haxe__exceptions__PosException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}

func hx_recover_class_haxe__ValueException(value any) *haxe__ValueException {
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

func hx_recover_class_haxe__io__Eof(value any) *haxe__io__Eof {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__io__Eof:
		return source
	default:
		return value.(*haxe__io__Eof)
	}
}

func hx_is_nominal_haxe__Exception(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__Exception:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	case *haxe__exceptions__NotImplementedException:
		return (source != nil)
	case *haxe__exceptions__PosException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__ValueException(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__io__Eof(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__io__Eof:
		return (source != nil)
	default:
		return false
	}
}
