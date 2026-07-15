package main

import "snapshot/hxrt"

func main() {
	var box_items []int
	box_items = []int{1, 2}
	box_items = append(box_items, 3)
	var v any = any(len(box_items))
	hxrt.Println(v)
	var v_1 any = any(box_items[2])
	hxrt.Println(v_1)
	if len(box_items) > 0 {
		box_items = box_items[:(len(box_items) - 1)]
	}
	var v_2 any = any(len(box_items))
	hxrt.Println(v_2)
	var nested_inner_items []*string
	nested_inner_items = []*string{hxrt.StringFromLiteral("a")}
	nested_inner_items = append(nested_inner_items, hxrt.StringFromLiteral("b"))
	var v_3 any = any(len(nested_inner_items))
	hxrt.Println(v_3)
	var v_4 any = any(nested_inner_items[1])
	hxrt.Println(v_4)
	if len(nested_inner_items) > 0 {
		nested_inner_items = nested_inner_items[:(len(nested_inner_items) - 1)]
	}
	var v_5 any = any(len(nested_inner_items))
	hxrt.Println(v_5)
}
