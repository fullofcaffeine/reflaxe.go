package main

import "snapshot/hxrt"

func main() {
	var s *string = nil
	s = hxrt.StringConcatStringPtr(s, hxrt.StringFromLiteral("a"))
	s = hxrt.StringConcatAny(s, 2)
	hxrt.Println(s)
}
