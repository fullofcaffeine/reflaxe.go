package main

import "snapshot/hxrt"

func main() {
	var box_items *hxrt.Array
	box_items = hxrt.NewArray(1, 2)
	box_items.Push(3)
	var v any = any(box_items.Len())
	hxrt.Println(v)
	var v_1 any = any(box_items.Get(2))
	hxrt.Println(v_1)
	box_items.Pop()
	var v_2 any = any(box_items.Len())
	hxrt.Println(v_2)
	var nested_inner_items *hxrt.Array
	nested_inner_items = hxrt.NewArray(hxrt.StringFromLiteral("a"))
	nested_inner_items.Push(hxrt.StringFromLiteral("b"))
	var v_3 any = any(nested_inner_items.Len())
	hxrt.Println(v_3)
	var v_4 any = any(nested_inner_items.Get(1))
	hxrt.Println(v_4)
	nested_inner_items.Pop()
	var v_5 any = any(nested_inner_items.Len())
	hxrt.Println(v_5)
}
