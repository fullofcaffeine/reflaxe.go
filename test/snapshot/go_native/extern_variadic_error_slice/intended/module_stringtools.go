package main

import "snapshot/hxrt"

func StringTools_replace(s *string, sub *string, by *string) *string {
	if hxrt.StringLengthStringPtr(sub) == 0 {
		var every_b *string
		every_b = hxrt.StringFromLiteral("")
		every_b = hxrt.StringConcatStringPtr(every_b, hxrt.StdString(by))
		_g := 0
		_g1 := hxrt.StringLengthStringPtr(s)
		for _g < _g1 {
			hx_post_1 := _g
			_g = int(int32((_g + 1)))
			index := hx_post_1
			x := hxrt.StringSubstrStringPtr(s, index, 1, true)
			every_b = hxrt.StringConcatStringPtr(every_b, hxrt.StdString(x))
			every_b = hxrt.StringConcatStringPtr(every_b, hxrt.StdString(by))
		}
		return every_b
	}
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	index_1 := 0
	for index_1 < hxrt.StringLengthStringPtr(s) {
		if (int((hxrt.Int32Wrap(index_1) + hxrt.Int32Wrap(hxrt.StringLengthStringPtr(sub)))) <= hxrt.StringLengthStringPtr(s)) && hxrt.StringEqualStringPtr(hxrt.StringSubstrStringPtr(s, index_1, hxrt.StringLengthStringPtr(sub), true), sub) {
			out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(by))
			index_1 = int((hxrt.Int32Wrap(index_1) + hxrt.Int32Wrap(hxrt.StringLengthStringPtr(sub))))
		} else {
			x_1 := hxrt.StringSubstrStringPtr(s, index_1, 1, true)
			out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x_1))
			index_1 = int(int32((index_1 + 1)))
		}
	}
	return out_b
}
