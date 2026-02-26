package main

import "snapshot/hxrt"

func main() {
	s := hxrt.StringFromLiteral("héllo")
	hxrt.Println(hxrt.StringLengthStringPtr(s))
	hxrt.Println(hxrt.StringCharAtStringPtr(s, 1))
	hxrt.Println(hxrt.StringCharCodeAtAnyStringPtr(s, 1))
	hxrt.Println(hxrt.StringSubstringStringPtr(s, 1, 4))
	hxrt.Println(hxrt.StringSubstrStringPtr(s, 2, 2, true))
	hxrt.Println(hxrt.StringSubstrStringPtr(s, -2, 0, false))
}
