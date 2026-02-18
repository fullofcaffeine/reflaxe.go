package main

import "snapshot/hxrt"

func main() {
	left := hxrt.StringFromLiteral("go")
	_ = left
	right := hxrt.StringFromLiteral("pher")
	joined := hxrt.StringConcatStringPtr(left, right)
	_ = joined
	eq := hxrt.StringEqualStringPtr(left, right)
	_ = eq
	neq := !hxrt.StringEqualStringPtr(left, right)
	_ = neq
	hxrt.Println(joined)
	hxrt.Println(eq)
	hxrt.Println(neq)
}
