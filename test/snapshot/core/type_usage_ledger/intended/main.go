package main

import (
	"fmt"
	"snapshot/hxrt"
)

func consumeDetails(details map[string]any) int {
	return func(hx_obj_1 map[string]any) int {
		hx_field_2 := hx_obj_1["count"]
		if hx_field_2 == nil {
			var hx_zero_3 int
			return hx_zero_3
		}
		return hx_field_2.(int)
	}(details)
}

func main() {
	box := New_UsedBox(39)
	nested := New_UsedBox(box)
	sibling := New_UsedSibling(1)
	callback := func(value int) *string {
		return hxrt.StdString(value)
	}
	hx_obj_4 := map[string]any{}
	hx_obj_4["label"] = callback(39)
	hx_obj_4["count"] = sibling.delta
	details := hx_obj_4
	fmt.Println(box.value)
	fmt.Println(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(38) + hxrt.Int32Wrap(consumeDetails(details)))))) + hxrt.Int32Wrap(nestedDepth(nested)))))) + hxrt.Int32Wrap(hxrt.StringLengthStringPtr(hxrt.StringFromLiteral("x")))))))
}

func nestedDepth(value *UsedBox) int {
	var hx_if_5 int
	if value == nil {
		hx_if_5 = 0
	} else {
		hx_if_5 = 1
	}
	return hx_if_5
}
