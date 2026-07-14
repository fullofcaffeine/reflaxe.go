package main

import (
	"log"
	"snapshot/hxrt"
)

func main() {
	hxrt.Println(hxrt.StringFromLiteral("native-start"))
	hxrt.TryCatch(func() {
		log.Panic(hxrt.StringFromLiteral("native-failure"))
	}, func(hx_caught_1 any) {
		error := hxrt.ExceptionCaught(hx_caught_1)
		hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("incorrectly-caught="), hxrt.StdString(error)))
	})
}
