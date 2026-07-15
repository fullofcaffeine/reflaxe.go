package main

import "snapshot/hxrt"

func main() {
	greeter := New_helper__Greeter()
	var v any = any(greeter.hello(hxrt.StringFromLiteral("go")))
	hxrt.Println(v)
}
