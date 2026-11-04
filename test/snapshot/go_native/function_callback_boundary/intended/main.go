package main

import (
	"snapshot/hxrt"
	"sort"
)

func main() {
	threshold := 3
	predicate := hxrt.NewFunction(func(index int) bool {
		return (index >= threshold)
	})
	alias := predicate
	var v any = any(sort.Search(10, hxrt.FunctionNative(predicate)))
	hxrt.Println(v)
	threshold = 4
	var v_1 any = any(sort.Search(10, hxrt.FunctionNative(predicate)))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.ReflectCompareMethods(any(predicate), any(alias)))
	hxrt.Println(v_2)
	search := func() *hxrt.Function[func(int, *hxrt.Function[func(int) bool]) int] {
		hx_extern_method_1 := sort.Search
		if hx_extern_method_1 == nil {
			return nil
		}
		return hxrt.NewFunction(func(hx_extern_arg_2 int, hx_extern_arg_3 *hxrt.Function[func(int) bool]) int {
			return int(hx_extern_method_1(hx_extern_arg_2, hxrt.FunctionNative(hx_extern_arg_3)))
		})
	}()
	var v_3 any = any(func() int {
		var hx_callable_4 *hxrt.Function[func(int, *hxrt.Function[func(int) bool]) int] = search
		var hx_argument_5 int = 10
		var hx_argument_6 *hxrt.Function[func(int) bool] = predicate
		return hxrt.FunctionCall(hx_callable_4)(hx_argument_5, hx_argument_6)
	}())
	hxrt.Println(v_3)
}
