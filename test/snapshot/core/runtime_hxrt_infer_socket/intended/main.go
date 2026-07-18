package main

import "snapshot/hxrt"

func main() {
	var v any = any(New_sys__net__Host(hxrt.StringFromLiteral("127.0.0.1")).toString())
	hxrt.Println(v)
}
