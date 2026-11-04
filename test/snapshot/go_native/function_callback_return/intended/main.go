package main

import (
	"snapshot/hxrt"
	"sync"
)

func main() {
	calls := 0
	callback := func() *hxrt.Function[func()] {
		hx_extern_method_1 := sync.OnceFunc(hxrt.FunctionNative(hxrt.NewFunction(func() {
			calls = int(int32((calls + 1)))
		})))
		if hx_extern_method_1 == nil {
			return nil
		}
		return hxrt.NewFunction(func() {
			hx_extern_method_1()
		})
	}()
	alias := callback
	func() {
		var hx_callable_2 *hxrt.Function[func()] = callback
		hxrt.FunctionCall(hx_callable_2)()
	}()
	func() {
		var hx_callable_3 *hxrt.Function[func()] = alias
		hxrt.FunctionCall(hx_callable_3)()
	}()
	factory := func() *hxrt.Function[func(*hxrt.Function[func()]) *hxrt.Function[func()]] {
		hx_extern_method_4 := sync.OnceFunc
		if hx_extern_method_4 == nil {
			return nil
		}
		return hxrt.NewFunction(func(hx_extern_arg_5 *hxrt.Function[func()]) *hxrt.Function[func()] {
			return func() *hxrt.Function[func()] {
				hx_extern_method_6 := hx_extern_method_4(hxrt.FunctionNative(hx_extern_arg_5))
				if hx_extern_method_6 == nil {
					return nil
				}
				return hxrt.NewFunction(func() {
					hx_extern_method_6()
				})
			}()
		})
	}()
	second := func() *hxrt.Function[func()] {
		var hx_callable_7 *hxrt.Function[func(*hxrt.Function[func()]) *hxrt.Function[func()]] = factory
		var hx_argument_8 *hxrt.Function[func()] = hxrt.NewFunction(func() {
			calls = int(int32((calls + 1)))
		})
		return hxrt.FunctionCall(hx_callable_7)(hx_argument_8)
	}()
	func() {
		var hx_callable_9 *hxrt.Function[func()] = second
		hxrt.FunctionCall(hx_callable_9)()
	}()
	func() {
		var hx_callable_10 *hxrt.Function[func()] = second
		hxrt.FunctionCall(hx_callable_10)()
	}()
	hxrt.Println(any(calls))
	var v any = any(hxrt.ReflectIsFunction(any(callback)))
	hxrt.Println(v)
	var v_1 any = any(hxrt.ReflectCompareMethods(any(callback), any(alias)))
	hxrt.Println(v_1)
}
