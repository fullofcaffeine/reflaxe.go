package main

import "snapshot/hxrt"

func haxe__crypto__Sha256_encode(value *string) *string {
	return hxrt.StdString(hxrt.CryptoSha256String(value))
}

func haxe__crypto__Sha256_fromValues(values []int) *haxe__io__Bytes {
	bytes := haxe__io__Bytes_alloc(len(values))
	_g := 0
	_g1 := len(values)
	for _g < _g1 {
		hx_post_70 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_70
		bytes.b[index] = int(int32((hxrt.Int32Wrap(values[index]) & hxrt.Int32Wrap(255))))
	}
	return bytes
}

func haxe__crypto__Sha256_make(value *haxe__io__Bytes) *haxe__io__Bytes {
	return haxe__crypto__Sha256_fromValues(hxrt.CryptoSha256Values(haxe__crypto__Sha256_toValues(value)))
}

func haxe__crypto__Sha256_toValues(bytes *haxe__io__Bytes) []int {
	values := hxrt.NewArray()
	_g := 0
	_g1 := bytes.length
	for _g < _g1 {
		hx_post_71 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_71
		values.Push(bytes.b[index])
	}
	return func(hx_lambda_raw_73 []any) []int {
		hx_lambda_out_74 := make([]int, 0, len(hx_lambda_raw_73))
		for _, hx_lambda_item_75 := range hx_lambda_raw_73 {
			hx_lambda_out_74 = append(hx_lambda_out_74, func(hx_value_76 any) int {
				if hx_value_76 == nil {
					var hx_zero_77 int
					return hx_zero_77
				}
				return hx_value_76.(int)
			}(hx_lambda_item_75))
		}
		return hx_lambda_out_74
	}(values.Values())
}
