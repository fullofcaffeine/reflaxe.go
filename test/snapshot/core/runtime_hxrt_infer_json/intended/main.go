package main

import "snapshot/hxrt"

func main() {
	var value any = hxrt.JsonParse(hxrt.StringFromLiteral("{\"name\":\"reflaxe.go\"}"))
	hxrt.Println(hxrt.StdString(!hxrt.AnyEqualsNull(value)))
}
