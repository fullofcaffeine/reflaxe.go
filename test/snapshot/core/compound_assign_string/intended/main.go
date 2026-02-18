package main

import "snapshot/hxrt"

func main() {
	var s *string = nil
	_ = s
	s = hxrt.StringConcatAny(s, hxrt.StringFromLiteral("a"))
	s = hxrt.StringConcatAny(s, 2)
	hxrt.Println(s)
}
