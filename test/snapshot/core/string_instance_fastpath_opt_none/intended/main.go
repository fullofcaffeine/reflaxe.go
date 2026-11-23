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
	var none any = noInt()
	var v_6 any = any(hxrt.StringSubstrOptional(s, 1, none))
	hxrt.Println(v_6)
	var v_7 any = any(hxrt.StringSubstringOptional(s, 1, none))
	hxrt.Println(v_7)
	var v_8 any = any(hxrt.StringIndexOfOptional(s, hxrt.StringFromLiteral("l"), none))
	hxrt.Println(v_8)
	var v_9 any = any(hxrt.StringLastIndexOfOptional(s, hxrt.StringFromLiteral("l"), none))
	hxrt.Println(v_9)
}

func noInt() any {
	return nil
}
