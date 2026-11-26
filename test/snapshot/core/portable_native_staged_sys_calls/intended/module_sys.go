package main

import "snapshot/hxrt"

func Sys_args() *hxrt.Array {
	return hxrt.ArrayFromValues(func(hx_sort_src_1 []*string) []any {
		hx_sort_out_3 := make([]any, 0, len(hx_sort_src_1))
		for _, hx_sort_item_2 := range hx_sort_src_1 {
			hx_sort_out_3 = append(hx_sort_out_3, hx_sort_item_2)
		}
		return hx_sort_out_3
	}(hxrt.SysArgs()))
}

func Sys_command(cmd *string, args *hxrt.Array) int {
	return hxrt.SysCommand(cmd, func() []*string {
		var hx_if_9 []*string
		if args == nil {
			hx_if_9 = nil
		} else {
			hx_if_9 = func(hx_lambda_raw_4 []any) []*string {
				hx_lambda_out_5 := make([]*string, 0, len(hx_lambda_raw_4))
				for _, hx_lambda_item_6 := range hx_lambda_raw_4 {
					hx_lambda_out_5 = append(hx_lambda_out_5, func(hx_value_7 any) *string {
						if hx_value_7 == nil {
							var hx_zero_8 *string
							return hx_zero_8
						}
						return hx_value_7.(*string)
					}(hx_lambda_item_6))
				}
				return hx_lambda_out_5
			}(args.Values())
		}
		return hx_if_9
	}())
}
