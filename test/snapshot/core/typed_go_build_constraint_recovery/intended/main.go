package main

import "snapshot/hxrt"

func main() {
	if !hxrt.StringEqualStringPtr(Registry_selected, hxrt.StringFromLiteral("ok")) {
		hxrt.Throw(hxrt.StringFromLiteral("selected recovery did not run"))
	}
	var v any = any(Registry_selected)
	hxrt.Println(v)
}
