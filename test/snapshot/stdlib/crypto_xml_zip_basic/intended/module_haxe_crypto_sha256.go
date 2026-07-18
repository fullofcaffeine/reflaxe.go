package main

import "snapshot/hxrt"

func haxe__crypto__Sha256_encode(value *string) *string {
	return hxrt.StdString(hxrt.CryptoSha256String(value))
}

func haxe__crypto__Sha256_make(value *haxe__io__Bytes) *haxe__io__Bytes {
	return haxe__io__Bytes___hx_fromNativeView(hxrt.CryptoSha256Values(value.__hx_this.__hx_nativeView()))
}
