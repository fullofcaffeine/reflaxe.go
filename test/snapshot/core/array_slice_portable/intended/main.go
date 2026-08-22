package main

import "snapshot/hxrt"

func main() {
	values := hxrt.NewArray(0, 1, 2, 3)
	printValues(values.SliceFrom(1))
	printValues(values.SliceOptional(1, 3))
	printValues(values.SliceFrom(-2))
	printValues(values.SliceOptional(0, -1))
	var v any = any(values.SliceFrom(7).Len())
	hxrt.Println(v)
	printValues(values.SliceOptional(1, nil))
	var nullableEnd any = nil
	printValues(values.SliceOptional(1, nullableEnd))
	var concreteEnd any = 3
	printValues(values.SliceOptional(1, concreteEnd))
	copy := values.SliceOptional(0, 2)
	hx_array_target_1 := copy
	hx_array_index_2 := 0
	hx_array_target_1.Set(hx_array_index_2, 9)
	hxrt.Println(any(values.Get(0)))
	printValues(NativeSlicer_middle(values))
}

func printValues(values *hxrt.Array) {
	var v any = any(hxrt.StringJoinAny(values.Values(), hxrt.StringFromLiteral(",")))
	hxrt.Println(v)
}
