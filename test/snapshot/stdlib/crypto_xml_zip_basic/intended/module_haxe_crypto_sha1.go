package main

import "snapshot/hxrt"

func haxe__crypto__Sha1_encode(value *string) *string {
	return hxrt.StdString(hxrt.CryptoSha1String(value))
}

func haxe__crypto__Sha1_fromValues(values []int) *haxe__io__Bytes {
	bytes := haxe__io__Bytes_alloc(len(values))
	_g := 0
	_g1 := len(values)
	for _g < _g1 {
		hx_post_86 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_86
		bytes.b[index] = int(int32((hxrt.Int32Wrap(values[index]) & hxrt.Int32Wrap(255))))
	}
	return bytes
}

func haxe__crypto__Sha1_make(value *haxe__io__Bytes) *haxe__io__Bytes {
	return haxe__crypto__Sha1_fromValues(hxrt.CryptoSha1Values(haxe__crypto__Sha1_toValues(value)))
}

func haxe__crypto__Sha1_toValues(bytes *haxe__io__Bytes) []int {
	values := hxrt.NewArray()
	_g := 0
	_g1 := bytes.length
	for _g < _g1 {
		hx_post_87 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_87
		values.Push(bytes.b[index])
	}
	return func(hx_lambda_raw_89 []any) []int {
		hx_lambda_out_90 := make([]int, 0, len(hx_lambda_raw_89))
		for _, hx_lambda_item_91 := range hx_lambda_raw_89 {
			hx_lambda_out_90 = append(hx_lambda_out_90, func(hx_value_92 any) int {
				if hx_value_92 == nil {
					var hx_zero_93 int
					return hx_zero_93
				}
				return hx_value_92.(int)
			}(hx_lambda_item_91))
		}
		return hx_lambda_out_90
	}(values.Values())
}
