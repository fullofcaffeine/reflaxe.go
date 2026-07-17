package main

import "snapshot/hxrt"

func main() {
	args := hxrt.NewArray(hxrt.StringFromLiteral("-c"), hxrt.StringFromLiteral("printf 'wrapper-out\\n'; exit 7"))
	code := hxrt.SysCommand(hxrt.StringFromLiteral("sh"), func() []*string {
		var hx_if_6 []*string
		if args == nil {
			hx_if_6 = nil
		} else {
			hx_if_6 = func(hx_lambda_raw_1 []any) []*string {
				hx_lambda_out_2 := make([]*string, 0, len(hx_lambda_raw_1))
				for _, hx_lambda_item_3 := range hx_lambda_raw_1 {
					hx_lambda_out_2 = append(hx_lambda_out_2, func(hx_value_4 any) *string {
						if hx_value_4 == nil {
							var hx_zero_5 *string
							return hx_zero_5
						}
						return hx_value_4.(*string)
					}(hx_lambda_item_3))
				}
				return hx_lambda_out_2
			}(args.Values())
		}
		return hx_if_6
	}())
	hxrt.SysExit(code)
}
