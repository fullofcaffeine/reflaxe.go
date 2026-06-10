package main

import "snapshot/hxrt"

func main() {
	var box_items []int
	box_items = []int{1, 2}
	box_items = append(box_items, 3)
	hxrt.Println(len(box_items))
	hxrt.Println(box_items[2])
	if len(box_items) > 0 {
		box_items = box_items[:(len(box_items) - 1)]
	}
	hxrt.Println(len(box_items))
	var nested_inner_items []*string
	nested_inner_items = []*string{hxrt.StringFromLiteral("a")}
	nested_inner_items = append(nested_inner_items, hxrt.StringFromLiteral("b"))
	hxrt.Println(len(nested_inner_items))
	hxrt.Println(nested_inner_items[1])
	if len(nested_inner_items) > 0 {
		nested_inner_items = nested_inner_items[:(len(nested_inner_items) - 1)]
	}
	hxrt.Println(len(nested_inner_items))
}
