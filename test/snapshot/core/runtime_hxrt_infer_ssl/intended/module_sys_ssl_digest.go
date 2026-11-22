package main

import "snapshot/hxrt"

func sys__ssl__Digest_make(data *haxe__io__Bytes, algorithm *string) *haxe__io__Bytes {
	input := hxrt.NewArray()
	_g := 0
	_g1 := data.length
	for _g < _g1 {
		hx_post_1 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_1
		input.Push(data.b[index])
	}
	nativeResult := hxrt.SslDigestMakeValues(func(hx_lambda_raw_3 []any) []int {
		hx_lambda_out_4 := make([]int, 0, len(hx_lambda_raw_3))
		for _, hx_lambda_item_5 := range hx_lambda_raw_3 {
			hx_lambda_out_4 = append(hx_lambda_out_4, func(hx_value_6 any) int {
				if hx_value_6 == nil {
					var hx_zero_7 int
					return hx_zero_7
				}
				return hx_value_6.(int)
			}(hx_lambda_item_5))
		}
		return hx_lambda_out_4
	}(input.Values()), algorithm)
	result := haxe__io__Bytes_alloc(len(nativeResult))
	_g_1 := 0
	_g1_1 := len(nativeResult)
	for _g_1 < _g1_1 {
		hx_post_8 := _g_1
		_g_1 = int(int32((_g_1 + 1)))
		index_1 := hx_post_8
		result.b[index_1] = int((hxrt.Int32Wrap(nativeResult[index_1]) & hxrt.Int32Wrap(255)))
		result.__hx_rawValid = false
	}
	return result
}

func sys__ssl__Digest_sign(data *haxe__io__Bytes, privateKey *sys__ssl__Key, algorithm *string) *haxe__io__Bytes {
	input := hxrt.NewArray()
	_g := 0
	_g1 := data.length
	for _g < _g1 {
		hx_post_9 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_9
		input.Push(data.b[index])
	}
	nativeResult := hxrt.SslDigestSignValues(func(hx_lambda_raw_11 []any) []int {
		hx_lambda_out_12 := make([]int, 0, len(hx_lambda_raw_11))
		for _, hx_lambda_item_13 := range hx_lambda_raw_11 {
			hx_lambda_out_12 = append(hx_lambda_out_12, func(hx_value_14 any) int {
				if hx_value_14 == nil {
					var hx_zero_15 int
					return hx_zero_15
				}
				return hx_value_14.(int)
			}(hx_lambda_item_13))
		}
		return hx_lambda_out_12
	}(input.Values()), privateKey.handle, algorithm)
	result := haxe__io__Bytes_alloc(len(nativeResult))
	_g_1 := 0
	_g1_1 := len(nativeResult)
	for _g_1 < _g1_1 {
		hx_post_16 := _g_1
		_g_1 = int(int32((_g_1 + 1)))
		index_1 := hx_post_16
		result.b[index_1] = int((hxrt.Int32Wrap(nativeResult[index_1]) & hxrt.Int32Wrap(255)))
		result.__hx_rawValid = false
	}
	return result
}

func sys__ssl__Digest_verify(data *haxe__io__Bytes, signature *haxe__io__Bytes, publicKey *sys__ssl__Key, algorithm *string) bool {
	input := hxrt.NewArray()
	_g := 0
	_g1 := data.length
	for _g < _g1 {
		hx_post_17 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_17
		input.Push(data.b[index])
	}
	signatureValues := hxrt.NewArray()
	_g_1 := 0
	_g1_1 := signature.length
	for _g_1 < _g1_1 {
		hx_post_19 := _g_1
		_g_1 = int(int32((_g_1 + 1)))
		index_1 := hx_post_19
		signatureValues.Push(signature.b[index_1])
	}
	return hxrt.SslDigestVerifyValues(func(hx_lambda_raw_21 []any) []int {
		hx_lambda_out_22 := make([]int, 0, len(hx_lambda_raw_21))
		for _, hx_lambda_item_23 := range hx_lambda_raw_21 {
			hx_lambda_out_22 = append(hx_lambda_out_22, func(hx_value_24 any) int {
				if hx_value_24 == nil {
					var hx_zero_25 int
					return hx_zero_25
				}
				return hx_value_24.(int)
			}(hx_lambda_item_23))
		}
		return hx_lambda_out_22
	}(input.Values()), func(hx_lambda_raw_26 []any) []int {
		hx_lambda_out_27 := make([]int, 0, len(hx_lambda_raw_26))
		for _, hx_lambda_item_28 := range hx_lambda_raw_26 {
			hx_lambda_out_27 = append(hx_lambda_out_27, func(hx_value_29 any) int {
				if hx_value_29 == nil {
					var hx_zero_30 int
					return hx_zero_30
				}
				return hx_value_29.(int)
			}(hx_lambda_item_28))
		}
		return hx_lambda_out_27
	}(signatureValues.Values()), publicKey.handle, algorithm)
}
