package main

import "snapshot/hxrt"

func main() {
	var parsed any = hxrt.JsonParse(hxrt.StringFromLiteral("[1,true,\"x\"]"))
	var v any = any(hxrt.StdString(hxrt.JsonStringify(parsed)))
	hxrt.Println(v)
}
