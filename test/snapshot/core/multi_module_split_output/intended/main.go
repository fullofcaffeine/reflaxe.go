package main

import "snapshot/hxrt"

func main() {
	greeter := New_helper__Greeter()
	hxrt.Println(greeter.hello(hxrt.StringFromLiteral("go")))
}
