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
	func(hx_fn *hxrt.Function[func()]) {
		hxrt.FunctionCall(hx_fn)()
	}(callback)
	func(hx_fn *hxrt.Function[func()]) {
		hxrt.FunctionCall(hx_fn)()
	}(alias)
	factory := func() *hxrt.Function[func(*hxrt.Function[func()]) *hxrt.Function[func()]] {
		hx_extern_method_2 := sync.OnceFunc
		if hx_extern_method_2 == nil {
			return nil
		}
		return hxrt.NewFunction(func(hx_extern_arg_3 *hxrt.Function[func()]) *hxrt.Function[func()] {
			return func() *hxrt.Function[func()] {
				hx_extern_method_4 := hx_extern_method_2(hxrt.FunctionNative(hx_extern_arg_3))
				if hx_extern_method_4 == nil {
					return nil
				}
				return hxrt.NewFunction(func() {
					hx_extern_method_4()
				})
			}()
		})
	}()
	second := func(hx_fn *hxrt.Function[func(*hxrt.Function[func()]) *hxrt.Function[func()]], hx_arg_0 *hxrt.Function[func()]) *hxrt.Function[func()] {
		return hxrt.FunctionCall(hx_fn)(hx_arg_0)
	}(factory, hxrt.NewFunction(func() {
		calls = int(int32((calls + 1)))
	}))
	func(hx_fn *hxrt.Function[func()]) {
		hxrt.FunctionCall(hx_fn)()
	}(second)
	func(hx_fn *hxrt.Function[func()]) {
		hxrt.FunctionCall(hx_fn)()
	}(second)
	hxrt.Println(any(calls))
	var v any = any(hxrt.ReflectIsFunction(any(callback)))
	hxrt.Println(v)
	var v_1 any = any(hxrt.ReflectCompareMethods(any(callback), any(alias)))
	hxrt.Println(v_1)
}
