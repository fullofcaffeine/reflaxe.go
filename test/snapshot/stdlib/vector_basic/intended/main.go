package main

import "snapshot/hxrt"

func main() {
	var this1 []int
	_ = this1
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
	_ = v
	v[0] = 3
	v[1] = 1
	v[2] = 4
	v[3] = 1
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("len:"), len(v)))
	render(hxrt.StringFromLiteral("base"), v)
	v[1] = 9
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("len_after_set:"), len(v)))
	render(hxrt.StringFromLiteral("mut"), v)
	var this1_1 []int
	_ = this1_1
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
	_ = w
	_g := 0
	_ = _g
	_g1 := len(w)
	_ = _g1
	for _g < _g1 {
		hx_post_5 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_5
		_ = i
		w[i] = int(int32((int32(int(int32((int32(i) + int32(1))))) * int32(2))))
	}
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("w_len:"), len(w)))
	render(hxrt.StringFromLiteral("w"), w)
}

func render(label *string, v []int) {
	out := hxrt.StringFromLiteral("")
	_ = out
	sum := 0
	_ = sum
	_g := 0
	_ = _g
	_g1 := len(v)
	_ = _g1
	for _g < _g1 {
		hx_post_6 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_6
		_ = i
		if i > 0 {
			out = hxrt.StringConcatAny(out, hxrt.StringFromLiteral(","))
		}
		out = hxrt.StringConcatAny(out, hxrt.StdString(v[i]))
		sum = int(int32((int32(sum) + int32(v[i]))))
	}
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringConcatAny(label, hxrt.StringFromLiteral(":")), out))
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringConcatAny(label, hxrt.StringFromLiteral("_sum:")), sum))
}
