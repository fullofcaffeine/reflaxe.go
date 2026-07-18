package main

import "snapshot/hxrt"

func haxe__crypto__Sha1_encode(value *string) *string {
	return hxrt.StdString(hxrt.CryptoSha1String(value))
}

func haxe__crypto__Sha1_make(value *haxe__io__Bytes) *haxe__io__Bytes {
	return haxe__io__Bytes___hx_fromNativeView(hxrt.CryptoSha1Values(value.__hx_this.__hx_nativeView()))
}
