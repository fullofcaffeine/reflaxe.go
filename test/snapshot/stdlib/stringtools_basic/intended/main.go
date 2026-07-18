package main

import "snapshot/hxrt"

func main() {
	s := hxrt.StringFromLiteral("  hi  ")
	var v any = any(StringTools_trim(s))
	hxrt.Println(v)
	var v_1 any = any(StringTools_startsWith(hxrt.StringFromLiteral("hello"), hxrt.StringFromLiteral("he")))
	hxrt.Println(v_1)
	var v_2 any = any(StringTools_replace(hxrt.StringFromLiteral("a-b-c"), hxrt.StringFromLiteral("-"), hxrt.StringFromLiteral(":")))
	hxrt.Println(v_2)
}
