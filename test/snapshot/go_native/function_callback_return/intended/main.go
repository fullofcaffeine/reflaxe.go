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
	builds := 0
	total := 0
	nested := func() *hxrt.Function[func() *hxrt.Function[func(int) int]] {
		hx_extern_method_12 := sync.OnceValue(func(hx_native_callback_11 *hxrt.Function[func() *hxrt.Function[func(int) int]]) func() func(int) int {
			if hx_native_callback_11 == nil {
				return nil
			}
			return func() func(int) int {
				return hxrt.FunctionNative(hxrt.FunctionCall(hx_native_callback_11)())
			}
		}(hxrt.NewFunction(func() *hxrt.Function[func(int) int] {
			builds = int(int32((builds + 1)))
			return hxrt.NewFunction(func(delta int) int {
				total = int((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(delta)))
				return total
			})
		})))
		if hx_extern_method_12 == nil {
			return nil
		}
		return hxrt.NewFunction(func() *hxrt.Function[func(int) int] {
			return func() *hxrt.Function[func(int) int] {
				hx_extern_method_13 := hx_extern_method_12()
				if hx_extern_method_13 == nil {
					return nil
				}
				return hxrt.NewFunction(func(hx_extern_arg_14 int) int {
					return int(hx_extern_method_13(hx_extern_arg_14))
				})
			}()
		})
	}()
	first := func() *hxrt.Function[func(int) int] {
		var hx_callable_15 *hxrt.Function[func() *hxrt.Function[func(int) int]] = nested
		return hxrt.FunctionCall(hx_callable_15)()
	}()
	second_1 := func() *hxrt.Function[func(int) int] {
		var hx_callable_16 *hxrt.Function[func() *hxrt.Function[func(int) int]] = nested
		return hxrt.FunctionCall(hx_callable_16)()
	}()
	var v_2 any = any(func() int {
		var hx_callable_17 *hxrt.Function[func(int) int] = first
		var hx_argument_18 int = 2
		return hxrt.FunctionCall(hx_callable_17)(hx_argument_18)
	}())
	hxrt.Println(v_2)
	var v_3 any = any(func() int {
		var hx_callable_19 *hxrt.Function[func(int) int] = second_1
		var hx_argument_20 int = 3
		return hxrt.FunctionCall(hx_callable_19)(hx_argument_20)
	}())
	hxrt.Println(v_3)
	hxrt.Println(any(builds))
	absent := func() *hxrt.Function[func() *hxrt.Function[func(int) int]] {
		hx_extern_method_22 := sync.OnceValue(func(hx_native_callback_21 *hxrt.Function[func() *hxrt.Function[func(int) int]]) func() func(int) int {
			if hx_native_callback_21 == nil {
				return nil
			}
			return func() func(int) int {
				return hxrt.FunctionNative(hxrt.FunctionCall(hx_native_callback_21)())
			}
		}(hxrt.NewFunction(func() *hxrt.Function[func(int) int] {
			return nil
		})))
		if hx_extern_method_22 == nil {
			return nil
		}
		return hxrt.NewFunction(func() *hxrt.Function[func(int) int] {
			return func() *hxrt.Function[func(int) int] {
				hx_extern_method_23 := hx_extern_method_22()
				if hx_extern_method_23 == nil {
					return nil
				}
				return hxrt.NewFunction(func(hx_extern_arg_24 int) int {
					return int(hx_extern_method_23(hx_extern_arg_24))
				})
			}()
		})
	}()
	var v_4 any = any((func() *hxrt.Function[func(int) int] {
		var hx_callable_25 *hxrt.Function[func() *hxrt.Function[func(int) int]] = absent
		return hxrt.FunctionCall(hx_callable_25)()
	}() == nil))
	hxrt.Println(v_4)
	nestedFactory := hxrt.NewFunction(func(hx_extern_arg_27 *hxrt.Function[func() *hxrt.Function[func(int) int]]) *hxrt.Function[func() *hxrt.Function[func(int) int]] {
		return func() *hxrt.Function[func() *hxrt.Function[func(int) int]] {
			hx_extern_method_29 := sync.OnceValue(func(hx_native_callback_28 *hxrt.Function[func() *hxrt.Function[func(int) int]]) func() func(int) int {
				if hx_native_callback_28 == nil {
					return nil
				}
				return func() func(int) int {
					return hxrt.FunctionNative(hxrt.FunctionCall(hx_native_callback_28)())
				}
			}(hx_extern_arg_27))
			if hx_extern_method_29 == nil {
				return nil
			}
			return hxrt.NewFunction(func() *hxrt.Function[func(int) int] {
				return func() *hxrt.Function[func(int) int] {
					hx_extern_method_30 := hx_extern_method_29()
					if hx_extern_method_30 == nil {
						return nil
					}
					return hxrt.NewFunction(func(hx_extern_arg_31 int) int {
						return int(hx_extern_method_30(hx_extern_arg_31))
					})
				}()
			})
		}()
	})
	fromFactory := func() *hxrt.Function[func() *hxrt.Function[func(int) int]] {
		var hx_callable_32 *hxrt.Function[func(*hxrt.Function[func() *hxrt.Function[func(int) int]]) *hxrt.Function[func() *hxrt.Function[func(int) int]]] = nestedFactory
		var hx_argument_33 *hxrt.Function[func() *hxrt.Function[func(int) int]] = hxrt.NewFunction(func() *hxrt.Function[func(int) int] {
			return hxrt.NewFunction(func(value int) int {
				return int((hxrt.Int32Wrap(value) + hxrt.Int32Wrap(3)))
			})
		})
		return hxrt.FunctionCall(hx_callable_32)(hx_argument_33)
	}()
	var v_5 any = any(func() int {
		var hx_callable_35 *hxrt.Function[func(int) int] = func() *hxrt.Function[func(int) int] {
			var hx_callable_34 *hxrt.Function[func() *hxrt.Function[func(int) int]] = fromFactory
			return hxrt.FunctionCall(hx_callable_34)()
		}()
		var hx_argument_36 int = 7
		return hxrt.FunctionCall(hx_callable_35)(hx_argument_36)
	}())
	hxrt.Println(v_5)
}
