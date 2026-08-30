package main

import "snapshot/hxrt"

func main() {
	var this1 []int
	this1 = []int{}
	hx_len_1 := 4
	if hx_len_1 < 0 {
		hx_len_1 = 0
	}
	if hx_len_1 <= len(this1) {
		this1 = this1[:hx_len_1]
	} else {
		var hx_zero_2 int
		for len(this1) < hx_len_1 {
			this1 = append(this1, hx_zero_2)
		}
	}
	v := this1
	v[0] = 3
	v[1] = 1
	v[2] = 4
	v[3] = 1
	var v1 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("len:"), len(v)))
	hxrt.Println(v1)
	render(hxrt.StringFromLiteral("base"), v)
	v[1] = 9
	var v1_1 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("len_after_set:"), len(v)))
	hxrt.Println(v1_1)
	render(hxrt.StringFromLiteral("mut"), v)
	var this1_1 []int
	this1_1 = []int{}
	hx_len_3 := 6
	if hx_len_3 < 0 {
		hx_len_3 = 0
	}
	if hx_len_3 <= len(this1_1) {
		this1_1 = this1_1[:hx_len_3]
	} else {
		var hx_zero_4 int
		for len(this1_1) < hx_len_3 {
			this1_1 = append(this1_1, hx_zero_4)
		}
	}
	w := this1_1
	_g := 0
	_g1 := len(w)
	for _g < _g1 {
		hx_post_5 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_5
		w[i] = int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(1)))) * hxrt.Int32Wrap(2)))
	}
	var v_1 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("w_len:"), len(w)))
	hxrt.Println(v_1)
	render(hxrt.StringFromLiteral("w"), w)
}

func render(label *string, v []int) {
	out := hxrt.StringFromLiteral("")
	sum := 0
	_g := 0
	_g1 := len(v)
	for _g < _g1 {
		hx_post_6 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_6
		if i > 0 {
			out = hxrt.StringConcatStringPtr(out, hxrt.StringFromLiteral(","))
		}
		out = hxrt.StringConcatStringPtr(out, hxrt.StdString(v[i]))
		sum = int((hxrt.Int32Wrap(sum) + hxrt.Int32Wrap(v[i])))
	}
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(label, hxrt.StringFromLiteral(":")), out)))
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(label, hxrt.StringFromLiteral("_sum:")), sum)))
}
