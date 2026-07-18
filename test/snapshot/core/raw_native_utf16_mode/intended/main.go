package main

import "snapshot/hxrt"

func bytesHex(value *haxe__io__Bytes) *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	_g := 0
	_g1 := value.length
	for _g < _g1 {
		hx_post_1 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_1
		if i > 0 {
			out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(","))
		}
		x := value.b[i]
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
	}
	return out_b
}

func main() {
	sample := hxrt.StringFromLiteral("hé")
	utf8 := haxe__io__Bytes_ofString(sample, haxe__io__Encoding_UTF8)
	rawNative := haxe__io__Bytes_ofString(sample, haxe__io__Encoding_RawNative)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("utf8.len="), utf8.length), hxrt.StringFromLiteral(" hex=")), bytesHex(utf8)))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("raw.len="), rawNative.length), hxrt.StringFromLiteral(" hex=")), bytesHex(rawNative)))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("raw.get="), rawNative.__hx_this.getString(0, rawNative.length, haxe__io__Encoding_RawNative)))
	hxrt.Println(v_2)
	output := New_haxe__io__BytesOutput()
	output.__hx_this.writeString(sample, haxe__io__Encoding_RawNative)
	written := output.__hx_this.getBytes()
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("out.raw.hex="), bytesHex(written)))
	hxrt.Println(v_3)
	input := New_haxe__io__BytesInput(written, nil, nil)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("in.raw="), input.__hx_this.readString(written.length, haxe__io__Encoding_RawNative)))
	hxrt.Println(v_4)
}
