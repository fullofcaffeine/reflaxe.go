package main

import "snapshot/hxrt"

func main() {
	left := hxrt.StringFromLiteral("porta")
	_ = left
	right := hxrt.StringFromLiteral("ble")
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
