package main

import "snapshot/hxrt"

func main() {
	s := hxrt.StringFromLiteral("héllo")
	out := hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringCharAtStringPtr(s, 1), hxrt.StdString(hxrt.StringCharCodeAtAnyStringPtr(s, 1))), hxrt.StringSubstringStringPtr(s, 0, 3)), hxrt.StringSubstrStringPtr(s, -2, 0, false))
	hxrt.Println(int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(s)) + hxrt.Int32Wrap(hxrt.StringLengthStringPtr(out))))))
}
