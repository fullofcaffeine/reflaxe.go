package main

import "snapshot/hxrt"

var haxe__crypto__Base64_BYTES *haxe__io__Bytes = haxe__io__Bytes_ofString(haxe__crypto__Base64_CHARS, nil)

var haxe__crypto__Base64_CHARS *string = hxrt.StringFromLiteral("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")

var haxe__crypto__Base64_URL_BYTES *haxe__io__Bytes = haxe__io__Bytes_ofString(haxe__crypto__Base64_URL_CHARS, nil)

var haxe__crypto__Base64_URL_CHARS *string = hxrt.StringFromLiteral("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_")

func haxe__crypto__Base64_addPadding(value *string, byteLength int, complement bool) *string {
	if !complement {
		return value
	}
	_g := int((hxrt.Int32Wrap(byteLength) % hxrt.Int32Wrap(3)))
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
	return haxe__io__Bytes___hx_fromNativeView(hxrt.CryptoBase64Decode(haxe__crypto__Base64_removePadding(value, complement), false))
}

func haxe__crypto__Base64_encode(bytes *haxe__io__Bytes, complement bool) *string {
	return haxe__crypto__Base64_addPadding(hxrt.StdString(hxrt.CryptoBase64Encode(bytes.__hx_this.__hx_nativeView(), false)), bytes.length, complement)
}

func haxe__crypto__Base64_removePadding(value *string, complement bool) *string {
	if complement {
		for (hxrt.StringLengthStringPtr(value) > 0) && (hxrt.StringCharCodeAtAnyStringPtr(value, int((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(value))-hxrt.Int32Wrap(1)))) == 61) {
			value = hxrt.StringSubstrStringPtr(value, 0, -1, true)
		}
	}
	return value
}

func haxe__crypto__Base64_urlDecode(value *string, complement bool) *haxe__io__Bytes {
	return haxe__io__Bytes___hx_fromNativeView(hxrt.CryptoBase64Decode(haxe__crypto__Base64_removePadding(value, complement), true))
}

func haxe__crypto__Base64_urlEncode(bytes *haxe__io__Bytes, complement bool) *string {
	return haxe__crypto__Base64_addPadding(hxrt.StdString(hxrt.CryptoBase64Encode(bytes.__hx_this.__hx_nativeView(), true)), bytes.length, complement)
}
