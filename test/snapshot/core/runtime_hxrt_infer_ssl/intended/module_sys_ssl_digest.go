package main

import "snapshot/hxrt"

func sys__ssl__Digest_make(data *haxe__io__Bytes, alg any) *haxe__io__Bytes {
	algName := hxrt.StdString(func(hx_value_1 any) *string {
		if hx_value_1 == nil {
			var hx_zero_2 *string
			return hx_zero_2
		}
		return hx_value_1.(*string)
	}(alg))
	_ = algName
	return hxrt_rawToHaxeBytes(hxrt.SslDigestMake(hxrt_haxeBytesToRaw(data), algName))
}

func sys__ssl__Digest_sign(data *haxe__io__Bytes, privKey *sys__ssl__Key, alg any) *haxe__io__Bytes {
	algName := hxrt.StdString(func(hx_value_3 any) *string {
		if hx_value_3 == nil {
			var hx_zero_4 *string
			return hx_zero_4
		}
		return hx_value_3.(*string)
	}(alg))
	_ = algName
	return hxrt_rawToHaxeBytes(hxrt.SslDigestSign(hxrt_haxeBytesToRaw(data), privKey.handle, algName))
}

func sys__ssl__Digest_verify(data *haxe__io__Bytes, signature *haxe__io__Bytes, pubKey *sys__ssl__Key, alg any) bool {
	algName := hxrt.StdString(func(hx_value_5 any) *string {
		if hx_value_5 == nil {
			var hx_zero_6 *string
			return hx_zero_6
		}
		return hx_value_5.(*string)
	}(alg))
	_ = algName
	return hxrt.SslDigestVerify(hxrt_haxeBytesToRaw(data), hxrt_haxeBytesToRaw(signature), pubKey.handle, algName)
}
