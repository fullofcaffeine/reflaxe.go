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
	if !unicode.In(65, func() []*unicode.RangeTable {
		array := hxrt.NewArray(letter, unicode.Digit)
		var this1 *hxrt.Array
		this1 = array
		return func(hx_lambda_raw_1 []any) []*unicode.RangeTable {
			hx_lambda_out_2 := make([]*unicode.RangeTable, 0, len(hx_lambda_raw_1))
			for _, hx_lambda_item_3 := range hx_lambda_raw_1 {
				hx_lambda_out_2 = append(hx_lambda_out_2, func(hx_value_4 any) *unicode.RangeTable {
					if hx_value_4 == nil {
						var hx_zero_5 *unicode.RangeTable
						return hx_zero_5
					}
					return hx_value_4.(*unicode.RangeTable)
				}(hx_lambda_item_3))
			}
			return hx_lambda_out_2
		}(this1.Values())
	}()...) {
		os.Exit(1)
	}
	if !unicode.In(65, []*unicode.RangeTable{letter}...) {
		os.Exit(4)
	}
	if unicode.In(65, func() []*unicode.RangeTable {
		array_1 := hxrt.NewArray(unicode.Digit)
		var this1_1 *hxrt.Array
		this1_1 = array_1
		return func(hx_lambda_raw_6 []any) []*unicode.RangeTable {
			hx_lambda_out_7 := make([]*unicode.RangeTable, 0, len(hx_lambda_raw_6))
			for _, hx_lambda_item_8 := range hx_lambda_raw_6 {
				hx_lambda_out_7 = append(hx_lambda_out_7, func(hx_value_9 any) *unicode.RangeTable {
					if hx_value_9 == nil {
						var hx_zero_10 *unicode.RangeTable
						return hx_zero_10
					}
					return hx_value_9.(*unicode.RangeTable)
				}(hx_lambda_item_8))
			}
			return hx_lambda_out_7
		}(this1_1.Values())
	}()...) {
		os.Exit(2)
	}
	template := template.New(*hxrt.StdString(hxrt.StringFromLiteral("t")))
	option := hxrt.StringFromLiteral("missingkey=zero")
	template.Option([]string{*hxrt.StdString(option), *hxrt.StdString(hxrt.StringFromLiteral("missingkey=default"))}...)
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
	_g_args = func(hx_value_11 any) []int {
		if hx_value_11 == nil {
			var hx_zero_12 []int
			return hx_zero_12
		}
		return hx_value_11.([]int)
	}(any(values))
	for _g_current < len(_g_args) {
		this1 := hxrt.ArrayFromValues(func(hx_sort_src_13 []int) []any {
			hx_sort_out_15 := make([]any, 0, len(hx_sort_src_13))
			for _, hx_sort_item_14 := range hx_sort_src_13 {
				hx_sort_out_15 = append(hx_sort_out_15, hx_sort_item_14)
			}
			return hx_sort_out_15
		}(_g_args))
		hx_post_16 := _g_current
		_g_current = int(int32((_g_current + 1)))
		index := hx_post_16
		value := hxrt.IntFromNullableAny(hxrt.IntFromNullableAny(this1.Get(index)))
		total = int((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(value)))
	}
	return total
}
