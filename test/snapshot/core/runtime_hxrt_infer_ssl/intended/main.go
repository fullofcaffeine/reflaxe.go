package main

import "snapshot/hxrt"

func main() {
	out := sys__ssl__Digest_make(haxe__io__Bytes_ofString(hxrt.StringFromLiteral("ssl"), nil), any(hxrt.StringFromLiteral("SHA256")))
	var v any = any(out.__hx_this.toHex())
	hxrt.Println(v)
}
