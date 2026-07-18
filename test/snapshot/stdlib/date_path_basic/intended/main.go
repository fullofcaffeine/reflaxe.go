package main

import "snapshot/hxrt"

func main() {
	d := Date_fromString(hxrt.StringFromLiteral("2024-02-03 04:05:06"))
	var v any = any(d.__hx_this.getFullYear())
	hxrt.Println(v)
	var v_1 any = any(d.__hx_this.getMonth())
	hxrt.Println(v_1)
	var v_2 any = any(d.__hx_this.getDate())
	hxrt.Println(v_2)
	var v_3 any = any(d.__hx_this.getHours())
	hxrt.Println(v_3)
	var v_4 any = any(haxe__io__Path_join(hxrt.NewArray(hxrt.StringFromLiteral("a"), hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("c.txt"))))
	hxrt.Println(v_4)
	p := New_haxe__io__Path(hxrt.StringFromLiteral("/tmp/demo.txt"))
	var v_5 any = any(p.dir)
	hxrt.Println(v_5)
	var v_6 any = any(p.file)
	hxrt.Println(v_6)
	var v_7 any = any(p.ext)
	hxrt.Println(v_7)
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
