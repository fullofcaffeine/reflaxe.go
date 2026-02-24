package main

import "snapshot/hxrt"

func main() {
	greeter := New_helper__Greeter()
	_ = greeter
	hxrt.Println(greeter.__hx_this.hello(hxrt.StringFromLiteral("go")))
}
