package main

import "snapshot/hxrt"

func main() {
	var value any = hxrt.JsonParse(hxrt.StringFromLiteral("{\"name\":\"reflaxe.go\"}"))
	var v any = any(hxrt.StdString(!hxrt.AnyEqualsNull(value)))
	hxrt.Println(v)
}
