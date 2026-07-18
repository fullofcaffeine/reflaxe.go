package main

import "snapshot/hxrt"

func main() {
	out := sys__ssl__Digest_make(haxe__io__Bytes_ofString(hxrt.StringFromLiteral("ssl"), nil), any(hxrt.StringFromLiteral("SHA256")))
	var v any = any(out.__hx_this.toHex())
	hxrt.Println(v)
}

type Std struct {
}

type haxe__ds__Option struct {
	tag    int
	params []any
}

var haxe__ds__Option_None *haxe__ds__Option = &haxe__ds__Option{tag: 1, params: []any{}}

func haxe__ds__Option_Some(value any) *haxe__ds__Option {
	return &haxe__ds__Option{tag: 0, params: []any{value}}
}
