package main

import "snapshot/hxrt"

func main() {
	var parsed any = hxrt.JsonParse(hxrt.StringFromLiteral("[1,true,\"x\"]"))
	hxrt.Println(hxrt.StdString(hxrt.JsonStringify(parsed)))
}
