package main

import "snapshot/hxrt"

func haxe__crypto__Md5_encode(value *string) *string {
	return hxrt.StdString(hxrt.CryptoMd5String(value))
}

func haxe__crypto__Md5_make(value *haxe__io__Bytes) *haxe__io__Bytes {
	return haxe__io__Bytes___hx_fromNativeView(hxrt.CryptoMd5Values(value.__hx_this.__hx_nativeView()))
}
