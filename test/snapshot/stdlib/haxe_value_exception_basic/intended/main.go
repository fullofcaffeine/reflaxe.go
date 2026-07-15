package main

import "snapshot/hxrt"

func main() {
	error := hxrt.NewValueException(hxrt.StringFromLiteral("boom"), nil, nil)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("msg="), hxrt.ExceptionMessage(error)))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("value="), func() *string {
		var hx_if_1 *string
		if hxrt.AnyEqualsNull(error.Value) {
			hx_if_1 = hxrt.StringFromLiteral("null")
		} else {
			var this1 any = error.Value
			hx_if_1 = hxrt.StdString(this1)
		}
		return hx_if_1
	}()))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("str="), hxrt.ExceptionMessage(error)))
	hxrt.Println(v_2)
}
