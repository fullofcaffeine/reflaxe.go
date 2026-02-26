package main

import "snapshot/hxrt"

func main() {
	s := hxrt.StringFromLiteral("héllo")
	out := hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringCharAt(s, 1), hxrt.StdString(hxrt.StringCharCodeAtAny(s, 1))), hxrt.StringSubstring(s, 0, 3)), hxrt.StringSubstr(s, -2, 0, false))
	hxrt.Println(int(int32((hxrt.Int32Wrap(hxrt.StringLength(s)) + hxrt.Int32Wrap(hxrt.StringLength(out))))))
}
