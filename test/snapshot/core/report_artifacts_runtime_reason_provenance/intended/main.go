package main

import "snapshot/hxrt"

func main() {
	var value any = hxrt.JsonParse(hxrt.StringFromLiteral("{\"name\":\"reflaxe.go\"}"))
	if hxrt.AnyEqualsNull(value) {
		hxrt.Throw(hxrt.StringFromLiteral("unexpected"))
	}
}
