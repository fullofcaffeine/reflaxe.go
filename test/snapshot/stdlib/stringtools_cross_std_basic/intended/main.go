package main

import "snapshot/hxrt"

func main() {
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("trim="), StringTools_trim(hxrt.StringFromLiteral("  hi  "))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("starts="), hxrt.StdString(StringTools_startsWith(hxrt.StringFromLiteral("hello"), hxrt.StringFromLiteral("he")))))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("replace="), StringTools_replace(hxrt.StringFromLiteral("a-b-c"), hxrt.StringFromLiteral("-"), hxrt.StringFromLiteral(":"))))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("contains="), hxrt.StdString(StringTools_contains(hxrt.StringFromLiteral("banana"), hxrt.StringFromLiteral("nan")))))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("ends="), hxrt.StdString(StringTools_endsWith(hxrt.StringFromLiteral("banana"), hxrt.StringFromLiteral("na")))))
	hxrt.Println(v_4)
}
