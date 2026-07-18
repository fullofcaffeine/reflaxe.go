package main

import "snapshot/hxrt"

func main() {
	parsed := haxe__io__Bytes_ofHex(hxrt.StringFromLiteral("0fDA"))
	var v any = any(parsed.__hx_this.toHex())
	hxrt.Println(v)
	buffer := New_haxe__io__BytesBuffer()
	buffer.b = hxrt.BytesBufferAddByte(buffer.b, 260)
	if 2 > parsed.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
	}
	buffer.b = hxrt.BytesBufferAddSlice(buffer.b, parsed.__hx_this.getData(), 1, 1)
	var encoding *haxe__io__Encoding = nil
	source := haxe__io__Bytes_ofString(hxrt.StringFromLiteral("Z"), encoding)
	buffer.b = hxrt.BytesBufferAdd(buffer.b, source.__hx_this.getData())
	out := buffer.__hx_this.getBytes()
	var v_1 any = any(out.__hx_this.toHex())
	hxrt.Println(v_1)
}
