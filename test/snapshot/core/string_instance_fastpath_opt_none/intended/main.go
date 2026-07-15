package main

import "snapshot/hxrt"

func main() {
	s := hxrt.StringFromLiteral("héllo")
	var v any = any(hxrt.StringLength(s))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringCharAt(s, 1))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringCharCodeAtAny(s, 1))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringSubstring(s, 1, 4))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringSubstr(s, 2, 2, true))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringSubstr(s, -2, 0, false))
	hxrt.Println(v_5)
}
