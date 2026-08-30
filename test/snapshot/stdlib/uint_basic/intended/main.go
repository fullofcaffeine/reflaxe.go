package main

import "snapshot/hxrt"

func main() {
	base := 3
	scaled := int((hxrt.Int32Wrap(base) * hxrt.Int32Wrap(7)))
	shifted := int(int32((uint32(hxrt.Int32Wrap(scaled)) >> uint(1))))
	var v any = any(hxrt.StdString(func() float64 {
		int := base
		var hx_if_1 float64
		if int < 0 {
			hx_if_1 = (4294967296.0 + float64(int))
		} else {
			hx_if_1 = (float64(int) + 0.0)
		}
		return hx_if_1
	}()))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StdString(func() float64 {
		int_1 := scaled
		var hx_if_2 float64
		if int_1 < 0 {
			hx_if_2 = (4294967296.0 + float64(int_1))
		} else {
			hx_if_2 = (float64(int_1) + 0.0)
		}
		return hx_if_2
	}()))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StdString(func() float64 {
		int_2 := shifted
		var hx_if_3 float64
		if int_2 < 0 {
			hx_if_3 = (4294967296.0 + float64(int_2))
		} else {
			hx_if_3 = (float64(int_2) + 0.0)
		}
		return hx_if_3
	}()))
	hxrt.Println(v_2)
}
