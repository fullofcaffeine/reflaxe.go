package main

import (
	"fmt"
	"snapshot/hxrt"
)

func main() {
	rendered := hxrt.StdString(fmt.Sprint([]any{*hxrt.StdString(hxrt.StringFromLiteral("left")), 7, *hxrt.StdString(hxrt.StringFromLiteral("right"))}...))
	if !hxrt.StringEqualStringPtr(rendered, hxrt.StringFromLiteral("left7right")) {
		hxrt.Throw(hxrt.StringFromLiteral("variadic extern arguments were not expanded as native values"))
	}
}
