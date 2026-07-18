package main

import "snapshot/hxrt"

func main() {
	utf8 := New_haxe__Utf8(0)
	utf8.__b = hxrt.StringConcatStringPtr(utf8.__b, haxe__Utf8_codePointToString(97))
	utf8.__b = hxrt.StringConcatStringPtr(utf8.__b, haxe__Utf8_codePointToString(128512))
	utf8.__b = hxrt.StringConcatStringPtr(utf8.__b, haxe__Utf8_codePointToString(233))
	value := utf8.__b
	sized := New_haxe__Utf8(4)
	sized.__b = hxrt.StringConcatStringPtr(sized.__b, haxe__Utf8_codePointToString(122))
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("string="), value)))
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("sized="), sized.__b))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("length="), haxe__io__Bytes_ofString(value, nil).length))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("sub="), func() *string {
		var unicode any = value
		return _UnicodeString__UnicodeString_Impl__substr(hxrt.StdString(unicode), 1, 2)
	}()))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("encode="), haxe__Utf8_encode(hxrt.StringFromLiteral("é"))))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("decode="), haxe__Utf8_decode(haxe__Utf8_encode(hxrt.StringFromLiteral("é")))))
	hxrt.Println(v_4)
}
