package main

import "snapshot/hxrt"

func main() {
	var parsed any = hxrt.JsonParse(hxrt.StringFromLiteral("[1,true,\"x\"]"))
	_ = parsed
	hxrt.Println(hxrt.JsonStringify(parsed))
}
