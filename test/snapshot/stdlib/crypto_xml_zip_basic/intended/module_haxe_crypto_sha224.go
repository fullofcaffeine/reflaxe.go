package main

import "snapshot/hxrt"

func haxe__crypto__Sha224_encode(value *string) *string {
	return hxrt.StdString(hxrt.CryptoSha224String(value))
}

func haxe__crypto__Sha224_fromValues(values []int) *haxe__io__Bytes {
	bytes := haxe__io__Bytes_alloc(len(values))
	_g := 0
	_g1 := len(values)
	for _g < _g1 {
		hx_post_78 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_78
		bytes.b[index] = int(int32((hxrt.Int32Wrap(values[index]) & hxrt.Int32Wrap(255))))
	}
	return bytes
}

func haxe__crypto__Sha224_make(value *haxe__io__Bytes) *haxe__io__Bytes {
	return haxe__crypto__Sha224_fromValues(hxrt.CryptoSha224Values(haxe__crypto__Sha224_toValues(value)))
}

func haxe__crypto__Sha224_toValues(bytes *haxe__io__Bytes) []int {
	values := hxrt.NewArray()
	_g := 0
	_g1 := bytes.length
	for _g < _g1 {
		hx_post_79 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_79
		values.Push(bytes.b[index])
	}
	return func(hx_lambda_raw_81 []any) []int {
		hx_lambda_out_82 := make([]int, 0, len(hx_lambda_raw_81))
		for _, hx_lambda_item_83 := range hx_lambda_raw_81 {
			hx_lambda_out_82 = append(hx_lambda_out_82, func(hx_value_84 any) int {
				if hx_value_84 == nil {
					var hx_zero_85 int
					return hx_zero_85
				}
				return hx_value_84.(int)
			}(hx_lambda_item_83))
		}
		return hx_lambda_out_82
	}(values.Values())
}
