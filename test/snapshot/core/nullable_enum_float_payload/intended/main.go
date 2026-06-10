package main

import "snapshot/hxrt"

type MaybeFloat struct {
	tag    int
	params []any
}

func MaybeFloat_Some(value any) *MaybeFloat {
	enumValue := &MaybeFloat{tag: 0}
	enumValue.params = []any{value}
	return enumValue
}

func describe(value *MaybeFloat) *string {
	var _g any = value.params[0]
	var payload any = _g
	var hx_if_1 *string
	if payload == nil {
		hx_if_1 = hxrt.StringFromLiteral("missing")
	} else {
		hx_if_1 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("value="), hxrt.StdString(payload.(float64)))
	}
	return hx_if_1
}

func main() {
	hxrt.Println(describe(MaybeFloat_Some(nil)))
	hxrt.Println(describe(MaybeFloat_Some(1.5)))
}
