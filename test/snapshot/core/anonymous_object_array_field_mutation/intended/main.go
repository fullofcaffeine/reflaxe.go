package main

import "snapshot/hxrt"

func main() {
	var box_items []int
	box_items = []int{1, 2}
	hx_arr_1 := box_items
	hx_arr_1 = append(hx_arr_1, 3)
	box_items = hx_arr_1
	hxrt.Println(len(box_items))
	hxrt.Println(box_items[2])
	hx_arr_2 := box_items
	if len(hx_arr_2) > 0 {
		hx_arr_2 = hx_arr_2[:(len(hx_arr_2) - 1)]
	}
	box_items = hx_arr_2
	hxrt.Println(len(box_items))
	var nested_inner_items []*string
	nested_inner_items = []*string{hxrt.StringFromLiteral("a")}
	hx_arr_3 := nested_inner_items
	hx_arr_3 = append(hx_arr_3, hxrt.StringFromLiteral("b"))
	nested_inner_items = hx_arr_3
	hxrt.Println(len(nested_inner_items))
	hxrt.Println(nested_inner_items[1])
	hx_arr_4 := nested_inner_items
	if len(hx_arr_4) > 0 {
		hx_arr_4 = hx_arr_4[:(len(hx_arr_4) - 1)]
	}
	nested_inner_items = hx_arr_4
	hxrt.Println(len(nested_inner_items))
}
