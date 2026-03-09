package main

import "snapshot/hxrt"

func main() {
	error := hxrt.NewValueException(hxrt.StringFromLiteral("boom"), nil, nil)
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("msg="), hxrt.ExceptionMessage(error)))
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("value="), func() *string {
		var hx_if_1 *string
		if hxrt.AnyEqualsNull(error.Value) {
			hx_if_1 = hxrt.StringFromLiteral("null")
		} else {
			var this1 any = error.Value
			hx_if_1 = hxrt.StdString(this1)
		}
		return hx_if_1
	}()))
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("str="), hxrt.ExceptionMessage(error)))
}
