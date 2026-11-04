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
		return hxrt.NewFunction(func(hx_extern_arg_2 int, hx_extern_arg_3 *hxrt.Function[func(int) bool]) int {
			return int(hx_extern_method_1(hx_extern_arg_2, hxrt.FunctionNative(hx_extern_arg_3)))
		})
	}()
	var v_3 any = any(func(hx_fn *hxrt.Function[func(int, *hxrt.Function[func(int) bool]) int], hx_arg_0 int, hx_arg_1 *hxrt.Function[func(int) bool]) int {
		return hxrt.FunctionCall(hx_fn)(hx_arg_0, hx_arg_1)
	}(search, 10, predicate))
	hxrt.Println(v_3)
}
