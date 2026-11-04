package main

import "snapshot/hxrt"

func handle(v int) {
	hxrt.TryCatch(func() {
		raise(v)
	}, func(hx_caught_1 any) {
		hx_exception_2 := haxe__Exception_caught(hx_caught_1)
		hx_payload_3 := haxe__Exception_catchValue(hx_caught_1)
		_ = hx_exception_2
		if func(hx_value any) bool {
			switch hx_value.(type) {
			case *string:
				return true
			default:
				return false
			}
		}(hx_payload_3) {
			var s *string = func(hx_value_6 any) *string {
				if hx_value_6 == nil {
					var hx_zero_7 *string
					return hx_zero_7
				}
				return hx_value_6.(*string)
			}(hx_payload_3)
			hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("S:"), s)))
		} else {
			if func(hx_value any) bool {
				switch hx_value.(type) {
				case int:
					return true
				default:
					return false
				}
			}(hx_payload_3) {
				var i int = func(hx_value_4 any) int {
					if hx_value_4 == nil {
						var hx_zero_5 int
						return hx_zero_5
					}
					return hx_value_4.(int)
				}(hx_payload_3)
				hxrt.Println(any(i))
			} else {
				var e any = hx_payload_3
				_ = e
				hxrt.Println(any(hxrt.StringFromLiteral("D")))
			}
		}
	})
}

func main() {
	handle(0)
	handle(1)
	handle(2)
}

func raise(v int) {
	if v == 0 {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("text")))
	}
	if v == 1 {
		hxrt.Throw(haxe__Exception_thrown(11))
	}
	hxrt.Throw(haxe__Exception_thrown(true))
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
