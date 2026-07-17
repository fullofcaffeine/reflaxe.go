package main

import "snapshot/hxrt"

var haxe__crypto__Base64_BYTES *haxe__io__Bytes = haxe__io__Bytes_ofString(haxe__crypto__Base64_CHARS)

var haxe__crypto__Base64_CHARS *string = hxrt.StringFromLiteral("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")

var haxe__crypto__Base64_URL_BYTES *haxe__io__Bytes = haxe__io__Bytes_ofString(haxe__crypto__Base64_URL_CHARS)

var haxe__crypto__Base64_URL_CHARS *string = hxrt.StringFromLiteral("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_")

func haxe__crypto__Base64_addPadding(value *string, byteLength int, complement bool) *string {
	if !complement {
		return value
	}
	_g := int(int32((hxrt.Int32Wrap(byteLength) % hxrt.Int32Wrap(3))))
	switch _g {
	case 1:
		return hxrt.StringConcatStringPtr(value, hxrt.StringFromLiteral("=="))
	case 2:
		return hxrt.StringConcatStringPtr(value, hxrt.StringFromLiteral("="))
	default:
		return value
	}
}

func haxe__crypto__Base64_decode(value *string, complement bool) *haxe__io__Bytes {
	return haxe__crypto__Base64_fromValues(hxrt.CryptoBase64Decode(haxe__crypto__Base64_removePadding(value, complement), false))
}

func haxe__crypto__Base64_encode(bytes *haxe__io__Bytes, complement bool) *string {
	return haxe__crypto__Base64_addPadding(hxrt.StdString(hxrt.CryptoBase64Encode(haxe__crypto__Base64_toValues(bytes), false)), bytes.length, complement)
}

func haxe__crypto__Base64_fromValues(values []int) *haxe__io__Bytes {
	bytes := haxe__io__Bytes_alloc(len(values))
	_g := 0
	_g1 := len(values)
	for _g < _g1 {
		hx_post_55 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_55
		bytes.b[index] = int(int32((hxrt.Int32Wrap(values[index]) & hxrt.Int32Wrap(255))))
	}
	return bytes
}

func haxe__crypto__Base64_removePadding(value *string, complement bool) *string {
	if complement {
		for (hxrt.StringLengthStringPtr(value) > 0) && (hxrt.StringCharCodeAtAnyStringPtr(value, int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(value))-hxrt.Int32Wrap(1))))) == 61) {
			value = hxrt.StringSubstrStringPtr(value, 0, -1, true)
		}
	}
	return value
}

func haxe__crypto__Base64_toValues(bytes *haxe__io__Bytes) []int {
	values := hxrt.NewArray()
	_g := 0
	_g1 := bytes.length
	for _g < _g1 {
		hx_post_56 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_56
		values.Push(bytes.b[index])
	}
	return func(hx_lambda_raw_58 []any) []int {
		hx_lambda_out_59 := make([]int, 0, len(hx_lambda_raw_58))
		for _, hx_lambda_item_60 := range hx_lambda_raw_58 {
			hx_lambda_out_59 = append(hx_lambda_out_59, func(hx_value_61 any) int {
				if hx_value_61 == nil {
					var hx_zero_62 int
					return hx_zero_62
				}
				return hx_value_61.(int)
			}(hx_lambda_item_60))
		}
		return hx_lambda_out_59
	}(values.Values())
}

func haxe__crypto__Base64_urlDecode(value *string, complement bool) *haxe__io__Bytes {
	return haxe__crypto__Base64_fromValues(hxrt.CryptoBase64Decode(haxe__crypto__Base64_removePadding(value, complement), true))
}

func haxe__crypto__Base64_urlEncode(bytes *haxe__io__Bytes, complement bool) *string {
	return haxe__crypto__Base64_addPadding(hxrt.StdString(hxrt.CryptoBase64Encode(haxe__crypto__Base64_toValues(bytes), true)), bytes.length, complement)
}
