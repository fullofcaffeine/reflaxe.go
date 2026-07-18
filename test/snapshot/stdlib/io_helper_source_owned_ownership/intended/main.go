package main

import "snapshot/hxrt"

func main() {
	lines := New_haxe__io__BytesInput(haxe__io__Bytes_ofString(hxrt.StringFromLiteral("first\r\nsecond"), nil), nil, nil)
	var v any = any(lines.__hx_this.readLine())
	hxrt.Println(v)
	var v_1 any = any(lines.__hx_this.readLine())
	hxrt.Println(v_1)
	replay := New_haxe__io__BytesInput(haxe__io__Bytes_ofString(hxrt.StringFromLiteral("012345"), nil), nil, nil)
	buf := haxe__io__Bytes_alloc(4)
	replay.__hx_this.readFullBytes(buf, 0, 4)
	var v_2 any = any(buf.__hx_this.toString())
	hxrt.Println(v_2)
	copyOut := New_haxe__io__BytesOutput()
	copyOut.__hx_this.writeInput(New_haxe__io__BytesInput(haxe__io__Bytes_ofString(hxrt.StringFromLiteral("xy"), nil), nil, nil).haxe__io__Input, nil)
	var v_3 any = any(copyOut.__hx_this.getBytes().__hx_this.toString())
	hxrt.Println(v_3)
	all := New_haxe__io__BytesInput(haxe__io__Bytes_ofString(hxrt.StringFromLiteral("zz"), nil), nil, nil).__hx_this.readAll(nil)
	var v_4 any = any(all.__hx_this.toString())
	hxrt.Println(v_4)
}
