package main

import "snapshot/hxrt"

func main() {
	expression := New_EReg(hxrt.StringFromLiteral("a"), hxrt.StringFromLiteral("g"))
	var v any = any(expression.__hx_this.replace(hxrt.StringFromLiteral("a-a"), hxrt.StringFromLiteral("b")))
	hxrt.Println(v)
}
