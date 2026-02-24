package main

import "snapshot/hxrt"

func main() {
	ok := go___Result_ok(7)
	_ = ok
	hxrt.Println(ok.__hx_this.isOk())
	hxrt.Println(ok.__hx_this.unwrap())
	err := go___Result_failure(hxrt.StringFromLiteral("boom"))
	_ = err
	hxrt.Println(err.__hx_this.isErr())
	hxrt.Println(err.__hx_this.error())
}
