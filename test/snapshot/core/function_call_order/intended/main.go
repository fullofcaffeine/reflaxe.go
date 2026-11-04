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
			var hx_callable_11 *hxrt.Function[func(int) int] = absent
			var hx_argument_12 int = func() int {
				var hx_callable_10 *hxrt.Function[func() int] = install
				return hxrt.FunctionCall(hx_callable_10)()
			}()
			return hxrt.FunctionCall(hx_callable_11)(hx_argument_12)
		}())
		hxrt.Println(v_3)
	}, func(hx_caught_8 any) {
		error := hxrt.ExceptionCaught(hx_caught_8)
		var v_2 any = any(hxrt.ExceptionMessage(error))
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
			var hx_callable_16 *hxrt.Function[func(int) int] = present
			var hx_argument_17 int = func() int {
				var hx_callable_15 *hxrt.Function[func() int] = clear
				return hxrt.FunctionCall(hx_callable_15)()
			}()
			return hxrt.FunctionCall(hx_callable_16)(hx_argument_17)
		}())
		hxrt.Println(v_5)
	}, func(hx_caught_13 any) {
		error_1 := hxrt.ExceptionCaught(hx_caught_13)
		var v_4 any = any(hxrt.ExceptionMessage(error_1))
		hxrt.Println(v_4)
	})
}
