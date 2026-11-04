package main

import (
	"fmt"
	"snapshot/hxrt"
)

func consumeDetails(details hxrt.Anonymous) int {
	return func(hx_value_1 any) int {
		if hx_value_1 == nil {
			var hx_zero_2 int
			return hx_zero_2
		}
		return hx_value_1.(int)
	}(hxrt.AnonymousGet(details, "count"))
}

func main() {
	box := New_UsedBox(39)
	nested := New_UsedBox(box)
	sibling := New_UsedSibling(1)
	callback := hxrt.NewFunction(func(value int) *string {
		return hxrt.StdString(value)
	})
	hx_obj_3 := map[string]any{}
	hx_obj_3["label"] = func() *string {
		var hx_callable_4 *hxrt.Function[func(int) *string] = callback
		var hx_argument_5 int = 39
		return hxrt.FunctionCall(hx_callable_4)(hx_argument_5)
	}()
	hx_obj_3["count"] = sibling.delta
	details := hxrt.AnonymousFromMap(hx_obj_3)
	fmt.Println(box.value)
	fmt.Println(int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(38) + hxrt.Int32Wrap(consumeDetails(details))))) + hxrt.Int32Wrap(nestedDepth(nested))))) + hxrt.Int32Wrap(hxrt.StringLengthStringPtr(hxrt.StringFromLiteral("x"))))))
}

func nestedDepth(value *UsedBox) int {
	var hx_if_6 int
	if value == nil {
		hx_if_6 = 0
	} else {
		hx_if_6 = 1
	}
	return hx_if_6
}
