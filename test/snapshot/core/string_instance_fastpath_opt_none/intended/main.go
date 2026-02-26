package main

import "snapshot/hxrt"

func main() {
	s := hxrt.StringFromLiteral("héllo")
	hxrt.Println(hxrt.StringLength(s))
	hxrt.Println(hxrt.StringCharAt(s, 1))
	hxrt.Println(hxrt.StringCharCodeAtAny(s, 1))
	hxrt.Println(hxrt.StringSubstring(s, 1, 4))
	hxrt.Println(hxrt.StringSubstr(s, 2, 2, true))
	hxrt.Println(hxrt.StringSubstr(s, -2, 0, false))
}
