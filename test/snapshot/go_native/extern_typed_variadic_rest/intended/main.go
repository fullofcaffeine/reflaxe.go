package main

import (
	"fmt"
	"os"
	"snapshot/hxrt"
	"text/template"
	"unicode"
)

func main() {
	letter := unicode.Letter
	if !unicode.In(65, []*unicode.RangeTable{letter, unicode.Digit}...) {
		os.Exit(1)
	}
	if !unicode.In(65, []*unicode.RangeTable{letter}...) {
		os.Exit(4)
	}
	if unicode.In(65, []*unicode.RangeTable{unicode.Digit}...) {
		os.Exit(2)
	}
	template_local := template.New(*hxrt.StdString(hxrt.StringFromLiteral("t")))
	option := hxrt.StringFromLiteral("missingkey=zero")
	template_local.Option([]string{*hxrt.StdString(option), *hxrt.StdString(hxrt.StringFromLiteral("missingkey=default"))}...)
	if !hxrt.StringEqualStringPtr(hxrt.StdString(fmt.Sprint([]any{*hxrt.StdString(hxrt.StringFromLiteral("left")), 7, *hxrt.StdString(hxrt.StringFromLiteral("right"))}...)), hxrt.StringFromLiteral("left7right")) {
		os.Exit(3)
	}
	if (sum([]int{1, 2, 3}) != 6) || (sum([]int{}) != 0) {
		os.Exit(5)
	}
}

func sum(values []int) int {
	total := 0
	var _g_current int
	var _g_args []int
	_g_current = 0
	_g_args = func(hx_value_1 any) []int {
		if hx_value_1 == nil {
			var hx_zero_2 []int
			return hx_zero_2
		}
		return hx_value_1.([]int)
	}(any(values))
	for _g_current < len(_g_args) {
		this1 := hxrt.ArrayFromValues(func(hx_sort_src_3 []int) []any {
			hx_sort_out_5 := make([]any, 0, len(hx_sort_src_3))
			for _, hx_sort_item_4 := range hx_sort_src_3 {
				hx_sort_out_5 = append(hx_sort_out_5, hx_sort_item_4)
			}
			return hx_sort_out_5
		}(_g_args))
		hx_post_6 := _g_current
		_g_current = int(int32((_g_current + 1)))
		index := hx_post_6
		value := hxrt.IntFromNullableAny(hxrt.IntFromNullableAny(this1.Get(index)))
		total = int((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(value)))
	}
	return total
}
