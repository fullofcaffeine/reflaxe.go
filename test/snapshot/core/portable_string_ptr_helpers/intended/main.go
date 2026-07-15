package main

import "snapshot/hxrt"

func main() {
	left := hxrt.StringFromLiteral("porta")
	right := hxrt.StringFromLiteral("ble")
	joined := hxrt.StringConcatStringPtr(left, right)
	eq := hxrt.StringEqualStringPtr(left, right)
	neq := !hxrt.StringEqualStringPtr(left, right)
	hxrt.Println(any(joined))
	hxrt.Println(any(eq))
	hxrt.Println(any(neq))
}
