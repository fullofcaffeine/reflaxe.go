package main

import "snapshot/hxrt"

func haxe__crypto__Md5_encode(value *string) *string {
	return hxrt.StdString(hxrt.CryptoMd5String(value))
}

func haxe__crypto__Md5_fromValues(values []int) *haxe__io__Bytes {
	bytes := haxe__io__Bytes_alloc(len(values))
	_g := 0
	_g1 := len(values)
	for _g < _g1 {
		hx_post_94 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_94
		bytes.b[index] = int(int32((hxrt.Int32Wrap(values[index]) & hxrt.Int32Wrap(255))))
	}
	return bytes
}

func haxe__crypto__Md5_make(value *haxe__io__Bytes) *haxe__io__Bytes {
	return haxe__crypto__Md5_fromValues(hxrt.CryptoMd5Values(haxe__crypto__Md5_toValues(value)))
}

func haxe__crypto__Md5_toValues(bytes *haxe__io__Bytes) []int {
	values := hxrt.NewArray()
	_g := 0
	_g1 := bytes.length
	for _g < _g1 {
		hx_post_95 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_95
		values.Push(bytes.b[index])
	}
	return func(hx_lambda_raw_97 []any) []int {
		hx_lambda_out_98 := make([]int, 0, len(hx_lambda_raw_97))
		for _, hx_lambda_item_99 := range hx_lambda_raw_97 {
			hx_lambda_out_98 = append(hx_lambda_out_98, func(hx_value_100 any) int {
				if hx_value_100 == nil {
					var hx_zero_101 int
					return hx_zero_101
				}
				return hx_value_100.(int)
			}(hx_lambda_item_99))
		}
		return hx_lambda_out_98
	}(values.Values())
}
