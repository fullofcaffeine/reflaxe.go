package main

import "snapshot/hxrt"

func main() {
	callback := hxrt.NewFunction(func(hx_tmp int) int {
		return 1
	})
	replace := hxrt.NewFunction(func() int {
		callback = hxrt.NewFunction(func(hx_tmp int) int {
			return 2
		})
		return 0
	})
	var v any = any(func() int {
		var hx_callable_2 *hxrt.Function[func(int) int] = callback
		var hx_argument_3 int = func() int {
			var hx_callable_1 *hxrt.Function[func() int] = replace
			return hxrt.FunctionCall(hx_callable_1)()
		}()
		return hxrt.FunctionCall(hx_callable_2)(hx_argument_3)
	}())
	hxrt.Println(v)
	earlier := 3
	mutate := hxrt.NewFunction(func() int {
		earlier = 9
		return 4
	})
	pair := hxrt.NewFunction(func(left int, right int) int {
		return int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(left) * hxrt.Int32Wrap(10)))) + hxrt.Int32Wrap(right)))
	})
	var v_1 any = any(func() int {
		var hx_callable_5 *hxrt.Function[func(int, int) int] = pair
		var hx_argument_6 int = earlier
		var hx_argument_7 int = func() int {
			var hx_callable_4 *hxrt.Function[func() int] = mutate
			return hxrt.FunctionCall(hx_callable_4)()
		}()
		return hxrt.FunctionCall(hx_callable_5)(hx_argument_6, hx_argument_7)
	}())
	hxrt.Println(v_1)
	var absent *hxrt.Function[func(int) int] = nil
	install := hxrt.NewFunction(func() int {
		absent = hxrt.NewFunction(func(hx_tmp int) int {
			return 8
		})
		return 0
	})
	hxrt.TryCatch(func() {
		var v_3 any = any(func() int {
			var hx_callable_12 *hxrt.Function[func(int) int] = absent
			var hx_argument_13 int = func() int {
				var hx_callable_11 *hxrt.Function[func() int] = install
				return hxrt.FunctionCall(hx_callable_11)()
			}()
			return hxrt.FunctionCall(hx_callable_12)(hx_argument_13)
		}())
		hxrt.Println(v_3)
	}, func(hx_caught_8 any) {
		hx_exception_9 := haxe__Exception_caught(hx_caught_8)
		hx_payload_10 := haxe__Exception_catchValue(hx_caught_8)
		_ = hx_payload_10
		var error *haxe__Exception = hx_recover_class_0(hx_exception_9)
		var v_2 any = any(error.__hx_this.get_message())
		hxrt.Println(v_2)
	})
	present := hxrt.NewFunction(func(hx_tmp int) int {
		return 7
	})
	clear := hxrt.NewFunction(func() int {
		present = nil
		return 0
	})
	hxrt.TryCatch(func() {
		var v_5 any = any(func() int {
			var hx_callable_18 *hxrt.Function[func(int) int] = present
			var hx_argument_19 int = func() int {
				var hx_callable_17 *hxrt.Function[func() int] = clear
				return hxrt.FunctionCall(hx_callable_17)()
			}()
			return hxrt.FunctionCall(hx_callable_18)(hx_argument_19)
		}())
		hxrt.Println(v_5)
	}, func(hx_caught_14 any) {
		hx_exception_15 := haxe__Exception_caught(hx_caught_14)
		hx_payload_16 := haxe__Exception_catchValue(hx_caught_14)
		_ = hx_payload_16
		var error_1 *haxe__Exception = hx_recover_class_0(hx_exception_15)
		var v_4 any = any(error_1.__hx_this.get_message())
		hxrt.Println(v_4)
	})
}

func hx_recover_class_0(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ValueException:
		return source.haxe__Exception
	case *haxe__Exception:
		return source
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
	case *haxe__ValueException:
		return (source != nil)
	case *haxe__Exception:
		return (source != nil)
	default:
		return false
	}
}
