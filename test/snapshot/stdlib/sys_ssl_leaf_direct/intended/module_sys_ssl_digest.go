package main

import "snapshot/hxrt"

func sys__ssl__Digest_make(data *haxe__io__Bytes, algorithm any) *haxe__io__Bytes {
	input := hxrt.NewArray()
	_g := 0
	_g1 := data.length
	for _g < _g1 {
		hx_post_8 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_8
		input.Push(data.b[index])
	}
	nativeResult := hxrt.SslDigestMakeValues(func(hx_lambda_raw_10 []any) []int {
		hx_lambda_out_11 := make([]int, 0, len(hx_lambda_raw_10))
		for _, hx_lambda_item_12 := range hx_lambda_raw_10 {
			hx_lambda_out_11 = append(hx_lambda_out_11, func(hx_value_13 any) int {
				if hx_value_13 == nil {
					var hx_zero_14 int
					return hx_zero_14
				}
				return hx_value_13.(int)
			}(hx_lambda_item_12))
		}
		return hx_lambda_out_11
	}(input.Values()), hxrt.StdString(func(hx_value_15 any) *string {
		if hx_value_15 == nil {
			var hx_zero_16 *string
			return hx_zero_16
		}
		return hx_value_15.(*string)
	}(algorithm)))
	result := haxe__io__Bytes_alloc(len(nativeResult))
	_g_1 := 0
	_g1_1 := len(nativeResult)
	for _g_1 < _g1_1 {
		hx_post_17 := _g_1
		_g_1 = int(int32((_g_1 + 1)))
		index_1 := hx_post_17
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
		hx_post_18 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_18
		input.Push(data.b[index])
	}
	nativeResult := hxrt.SslDigestSignValues(func(hx_lambda_raw_20 []any) []int {
		hx_lambda_out_21 := make([]int, 0, len(hx_lambda_raw_20))
		for _, hx_lambda_item_22 := range hx_lambda_raw_20 {
			hx_lambda_out_21 = append(hx_lambda_out_21, func(hx_value_23 any) int {
				if hx_value_23 == nil {
					var hx_zero_24 int
					return hx_zero_24
				}
				return hx_value_23.(int)
			}(hx_lambda_item_22))
		}
		return hx_lambda_out_21
	}(input.Values()), privateKey.handle, hxrt.StdString(func(hx_value_25 any) *string {
		if hx_value_25 == nil {
			var hx_zero_26 *string
			return hx_zero_26
		}
		return hx_value_25.(*string)
	}(algorithm)))
	result := haxe__io__Bytes_alloc(len(nativeResult))
	_g_1 := 0
	_g1_1 := len(nativeResult)
	for _g_1 < _g1_1 {
		hx_post_27 := _g_1
		_g_1 = int(int32((_g_1 + 1)))
		index_1 := hx_post_27
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
		hx_post_28 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_28
		input.Push(data.b[index])
	}
	signatureValues := hxrt.NewArray()
	_g_1 := 0
	_g1_1 := signature.length
	for _g_1 < _g1_1 {
		hx_post_30 := _g_1
		_g_1 = int(int32((_g_1 + 1)))
		index_1 := hx_post_30
		signatureValues.Push(signature.b[index_1])
	}
	return hxrt.SslDigestVerifyValues(func(hx_lambda_raw_32 []any) []int {
		hx_lambda_out_33 := make([]int, 0, len(hx_lambda_raw_32))
		for _, hx_lambda_item_34 := range hx_lambda_raw_32 {
			hx_lambda_out_33 = append(hx_lambda_out_33, func(hx_value_35 any) int {
				if hx_value_35 == nil {
					var hx_zero_36 int
					return hx_zero_36
				}
				return hx_value_35.(int)
			}(hx_lambda_item_34))
		}
		return hx_lambda_out_33
	}(input.Values()), func(hx_lambda_raw_37 []any) []int {
		hx_lambda_out_38 := make([]int, 0, len(hx_lambda_raw_37))
		for _, hx_lambda_item_39 := range hx_lambda_raw_37 {
			hx_lambda_out_38 = append(hx_lambda_out_38, func(hx_value_40 any) int {
				if hx_value_40 == nil {
					var hx_zero_41 int
					return hx_zero_41
				}
				return hx_value_40.(int)
			}(hx_lambda_item_39))
		}
		return hx_lambda_out_38
	}(signatureValues.Values()), publicKey.handle, hxrt.StdString(func(hx_value_42 any) *string {
		if hx_value_42 == nil {
			var hx_zero_43 *string
			return hx_zero_43
		}
		return hx_value_42.(*string)
	}(algorithm)))
}
