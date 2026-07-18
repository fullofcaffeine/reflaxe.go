package main

import "snapshot/hxrt"

func main() {
	socket := New_sys__ssl__Socket()
	socket.verifyCert = false
	socket.__hx_this.close()
	hxrt.Println(any(hxrt.StringFromLiteral("tls-socket-ready")))
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
