package main

import "snapshot/hxrt"

func sys__ssl__Digest_make(data *haxe__io__Bytes, algorithm any) *haxe__io__Bytes {
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
	}(input.Values()), hxrt.StdString(func(hx_value_8 any) *string {
		if hx_value_8 == nil {
			var hx_zero_9 *string
			return hx_zero_9
		}
		return hx_value_8.(*string)
	}(algorithm)))
	result := haxe__io__Bytes_alloc(len(nativeResult))
	_g_1 := 0
	_g1_1 := len(nativeResult)
	for _g_1 < _g1_1 {
		hx_post_10 := _g_1
		_g_1 = int(int32((_g_1 + 1)))
		index_1 := hx_post_10
		result.b[index_1] = int(int32((hxrt.Int32Wrap(nativeResult[index_1]) & hxrt.Int32Wrap(255))))
		result.__hx_rawValid = false
	}
	return result
}

func sys__ssl__Digest_sign(data *haxe__io__Bytes, privateKey *sys__ssl__Key, algorithm any) *haxe__io__Bytes {
	input := hxrt.NewArray()
	_g := 0
	_g1 := data.length
	for _g < _g1 {
		hx_post_11 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_11
		input.Push(data.b[index])
	}
	nativeResult := hxrt.SslDigestSignValues(func(hx_lambda_raw_13 []any) []int {
		hx_lambda_out_14 := make([]int, 0, len(hx_lambda_raw_13))
		for _, hx_lambda_item_15 := range hx_lambda_raw_13 {
			hx_lambda_out_14 = append(hx_lambda_out_14, func(hx_value_16 any) int {
				if hx_value_16 == nil {
					var hx_zero_17 int
					return hx_zero_17
				}
				return hx_value_16.(int)
			}(hx_lambda_item_15))
		}
		return hx_lambda_out_14
	}(input.Values()), privateKey.handle, hxrt.StdString(func(hx_value_18 any) *string {
		if hx_value_18 == nil {
			var hx_zero_19 *string
			return hx_zero_19
		}
		return hx_value_18.(*string)
	}(algorithm)))
	result := haxe__io__Bytes_alloc(len(nativeResult))
	_g_1 := 0
	_g1_1 := len(nativeResult)
	for _g_1 < _g1_1 {
		hx_post_20 := _g_1
		_g_1 = int(int32((_g_1 + 1)))
		index_1 := hx_post_20
		result.b[index_1] = int(int32((hxrt.Int32Wrap(nativeResult[index_1]) & hxrt.Int32Wrap(255))))
		result.__hx_rawValid = false
	}
	return result
}

func sys__ssl__Digest_verify(data *haxe__io__Bytes, signature *haxe__io__Bytes, publicKey *sys__ssl__Key, algorithm any) bool {
	input := hxrt.NewArray()
	_g := 0
	_g1 := data.length
	for _g < _g1 {
		hx_post_21 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_21
		input.Push(data.b[index])
	}
	signatureValues := hxrt.NewArray()
	_g_1 := 0
	_g1_1 := signature.length
	for _g_1 < _g1_1 {
		hx_post_23 := _g_1
		_g_1 = int(int32((_g_1 + 1)))
		index_1 := hx_post_23
		signatureValues.Push(signature.b[index_1])
	}
	return hxrt.SslDigestVerifyValues(func(hx_lambda_raw_25 []any) []int {
		hx_lambda_out_26 := make([]int, 0, len(hx_lambda_raw_25))
		for _, hx_lambda_item_27 := range hx_lambda_raw_25 {
			hx_lambda_out_26 = append(hx_lambda_out_26, func(hx_value_28 any) int {
				if hx_value_28 == nil {
					var hx_zero_29 int
					return hx_zero_29
				}
				return hx_value_28.(int)
			}(hx_lambda_item_27))
		}
		return hx_lambda_out_26
	}(input.Values()), func(hx_lambda_raw_30 []any) []int {
		hx_lambda_out_31 := make([]int, 0, len(hx_lambda_raw_30))
		for _, hx_lambda_item_32 := range hx_lambda_raw_30 {
			hx_lambda_out_31 = append(hx_lambda_out_31, func(hx_value_33 any) int {
				if hx_value_33 == nil {
					var hx_zero_34 int
					return hx_zero_34
				}
				return hx_value_33.(int)
			}(hx_lambda_item_32))
		}
		return hx_lambda_out_31
	}(signatureValues.Values()), publicKey.handle, hxrt.StdString(func(hx_value_35 any) *string {
		if hx_value_35 == nil {
			var hx_zero_36 *string
			return hx_zero_36
		}
		return hx_value_35.(*string)
	}(algorithm)))
}
