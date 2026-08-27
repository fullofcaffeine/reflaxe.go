package main

import (
	"fmt"
	"snapshot/hxrt"
)

func main() {
	stringSource := hxrt.NewArray(hxrt.StringFromLiteral("alpha"), hxrt.StringFromLiteral("beta"))
	nativeStrings := func(hx_lambda_raw_1 []any) []string {
		hx_lambda_out_2 := make([]string, 0, len(hx_lambda_raw_1))
		for _, hx_lambda_item_3 := range hx_lambda_raw_1 {
			hx_lambda_out_2 = append(hx_lambda_out_2, *hxrt.StdString(hx_lambda_item_3))
		}
		return hx_lambda_out_2
	}(stringSource.Values())
	inlineStrings := []string{*hxrt.StdString(hxrt.StringFromLiteral("gamma"))}
	nativeStrings[0] = *hxrt.StdString(hxrt.StringFromLiteral("delta"))
	fmt.Println(*hxrt.StdString(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("strings="), renderStrings(nativeStrings))))
	fmt.Println(*hxrt.StdString(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("inline="), renderStrings(inlineStrings))))
}

func renderStrings(values []string) *string {
	rendered := hxrt.StringFromLiteral("")
	_g := 0
	_g1 := len(values)
	for _g < _g1 {
		hx_post_4 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_4
		if index > 0 {
			rendered = hxrt.StringConcatStringPtr(rendered, hxrt.StringFromLiteral(","))
		}
		rendered = hxrt.StringConcatStringPtr(rendered, hxrt.StdString(values[index]))
	}
	return rendered
}
