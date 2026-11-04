package main

import "snapshot/hxrt"

func main() {
	hxrt.TryCatch(func() {
		hxrt.TryCatch(func() {
			hxrt.Throw(haxe__Exception_thrown(true))
		}, func(hx_caught_4 any) {
			hx_exception_5 := haxe__Exception_caught(hx_caught_4)
			hx_payload_6 := haxe__Exception_catchValue(hx_caught_4)
			_ = hx_exception_5
			if func(hx_value any) bool {
				switch hx_value.(type) {
				case int:
					return true
				default:
					return false
				}
			}(hx_payload_6) {
				var i int = func(hx_value_7 any) int {
					if hx_value_7 == nil {
						var hx_zero_8 int
						return hx_zero_8
					}
					return hx_value_7.(int)
				}(hx_payload_6)
				hxrt.Println(any(i))
			} else {
				hxrt.Throw(hx_caught_4)
			}
		})
	}, func(hx_caught_1 any) {
		hx_exception_2 := haxe__Exception_caught(hx_caught_1)
		hx_payload_3 := haxe__Exception_catchValue(hx_caught_1)
		_ = hx_exception_2
		var e any = hx_payload_3
		_ = e
		hxrt.Println(any(hxrt.StringFromLiteral("outer")))
	})
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
