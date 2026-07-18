package main

import "snapshot/hxrt"

func main() {
	bytes := haxe__io__Bytes_ofString(hxrt.StringFromLiteral("abc"), nil)
	bytes.b[1] = 122
	bytes.__hx_rawValid = false
	var v any = any(bytes.__hx_this.toString())
	hxrt.Println(v)
	var v_1 any = any(bytes.length)
	hxrt.Println(v_1)
	buffer := New_haxe__io__BytesBuffer()
	var encoding *haxe__io__Encoding = nil
	source := haxe__io__Bytes_ofString(hxrt.StringFromLiteral("Hi"), encoding)
	buffer.b = hxrt.BytesBufferAdd(buffer.b, source.__hx_this.getData())
	buffer.b = hxrt.BytesBufferAddByte(buffer.b, 33)
	out := buffer.__hx_this.getBytes()
	var v_2 any = any(out.__hx_this.toString())
	hxrt.Println(v_2)
}
