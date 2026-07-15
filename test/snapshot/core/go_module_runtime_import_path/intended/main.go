package main

import "example.com/acme/reflaxe-go-smoke/hxrt"

func main() {
	hxrt.Println(any(hxrt.StringFromLiteral("go-module-runtime-import-ok")))
}
