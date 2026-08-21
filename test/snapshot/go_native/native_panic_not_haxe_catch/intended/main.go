package main

import (
	"log"
	"snapshot/hxrt"
)

func main() {
	hxrt.Println(any(hxrt.StringFromLiteral("native-start")))
	hxrt.TryCatch(func() {
		log.Panic(*hxrt.StdString(hxrt.StringFromLiteral("native-failure")))
	}, func(hx_caught_1 any) {
		error := hxrt.ExceptionCaught(hx_caught_1)
		var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("incorrectly-caught="), hxrt.StdString(error)))
		hxrt.Println(v)
	})
}
