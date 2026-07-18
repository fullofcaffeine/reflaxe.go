package main

import "snapshot/hxrt"

func emit(label *string, value *haxe___Int64_____Int64) {
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(label, hxrt.StringFromLiteral("=")), haxe___Int64__Int64_Impl__toString(value)))
	hxrt.Println(v)
}

func main() {
	max := haxe__Int64Helper_parseString(hxrt.StringFromLiteral("9223372036854775807"))
	min := haxe__Int64Helper_parseString(hxrt.StringFromLiteral("-9223372036854775808"))
	var b_low int
	var b_high int
	b_high = 0
	b_low = 1
	high := int(int32((hxrt.Int32Wrap(max.high) + hxrt.Int32Wrap(b_high))))
	low := int(int32((hxrt.Int32Wrap(max.low) + hxrt.Int32Wrap(b_low))))
	if haxe___Int32__Int32_Impl__ucompare(low, max.low) < 0 {
		hx_post_1 := high
		high = int(int32((high + 1)))
		ret := hx_post_1
		_ = ret
		high = high
	}
	x := New_haxe___Int64_____Int64(high, low)
	var this1 *haxe___Int64_____Int64
	this1 = x
	value := this1
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("wrap_add="), haxe___Int64__Int64_Impl__toString(value)))
	hxrt.Println(v)
	var b_low_1 int
	var b_high_1 int
	b_high_1 = 0
	b_low_1 = 1
	high_1 := int(int32((hxrt.Int32Wrap(min.high) - hxrt.Int32Wrap(b_high_1))))
	low_1 := int(int32((hxrt.Int32Wrap(min.low) - hxrt.Int32Wrap(b_low_1))))
	if haxe___Int32__Int32_Impl__ucompare(min.low, b_low_1) < 0 {
		hx_post_2 := high_1
		high_1 = int(int32((high_1 - 1)))
		ret_1 := hx_post_2
		_ = ret_1
		high_1 = high_1
	}
	x_1 := New_haxe___Int64_____Int64(high_1, low_1)
	var this1_1 *haxe___Int64_____Int64
	this1_1 = x_1
	value_1 := this1_1
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("wrap_sub="), haxe___Int64__Int64_Impl__toString(value_1)))
	hxrt.Println(v_1)
	a := haxe__Int64Helper_parseString(hxrt.StringFromLiteral("1234567890123"))
	b := haxe__Int64Helper_parseString(hxrt.StringFromLiteral("-987654321"))
	high_2 := int(int32((hxrt.Int32Wrap(a.high) + hxrt.Int32Wrap(b.high))))
	low_2 := int(int32((hxrt.Int32Wrap(a.low) + hxrt.Int32Wrap(b.low))))
	if haxe___Int32__Int32_Impl__ucompare(low_2, a.low) < 0 {
		hx_post_3 := high_2
		high_2 = int(int32((high_2 + 1)))
		ret_2 := hx_post_3
		_ = ret_2
		high_2 = high_2
	}
	x_2 := New_haxe___Int64_____Int64(high_2, low_2)
	var this1_2 *haxe___Int64_____Int64
	this1_2 = x_2
	value_2 := this1_2
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("sum="), haxe___Int64__Int64_Impl__toString(value_2)))
	hxrt.Println(v_2)
	high_3 := int(int32((hxrt.Int32Wrap(a.high) - hxrt.Int32Wrap(b.high))))
	low_3 := int(int32((hxrt.Int32Wrap(a.low) - hxrt.Int32Wrap(b.low))))
	if haxe___Int32__Int32_Impl__ucompare(a.low, b.low) < 0 {
		hx_post_4 := high_3
		high_3 = int(int32((high_3 - 1)))
		ret_3 := hx_post_4
		_ = ret_3
		high_3 = high_3
	}
	x_3 := New_haxe___Int64_____Int64(high_3, low_3)
	var this1_3 *haxe___Int64_____Int64
	this1_3 = x_3
	value_3 := this1_3
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("diff="), haxe___Int64__Int64_Impl__toString(value_3)))
	hxrt.Println(v_3)
	var a_low int
	var a_high int
	a_high = 0
	a_low = 30000
	var b_low_2 int
	var b_high_2 int
	b_high_2 = 0
	b_low_2 = 30000
	mask := 65535
	al := int(int32((hxrt.Int32Wrap(a_low) & hxrt.Int32Wrap(mask))))
	ah := int(int32(int32((uint32(hxrt.Int32Wrap(a_low)) >> uint(16)))))
	bl := int(int32((hxrt.Int32Wrap(b_low_2) & hxrt.Int32Wrap(mask))))
	bh := int(int32(int32((uint32(hxrt.Int32Wrap(b_low_2)) >> uint(16)))))
	p00 := int(int32((hxrt.Int32Wrap(al) * hxrt.Int32Wrap(bl))))
	p10 := int(int32((hxrt.Int32Wrap(ah) * hxrt.Int32Wrap(bl))))
	p01 := int(int32((hxrt.Int32Wrap(al) * hxrt.Int32Wrap(bh))))
	p11 := int(int32((hxrt.Int32Wrap(ah) * hxrt.Int32Wrap(bh))))
	low_4 := p00
	high_4 := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(p11) + hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(p01)) >> uint(16)))))))))) + hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(p10)) >> uint(16)))))))))
	p01 = int(int32((hxrt.Int32Wrap(p01) << uint(16))))
	low_4 = int(int32((hxrt.Int32Wrap(low_4) + hxrt.Int32Wrap(p01))))
	if haxe___Int32__Int32_Impl__ucompare(low_4, p01) < 0 {
		hx_post_5 := high_4
		high_4 = int(int32((high_4 + 1)))
		ret_4 := hx_post_5
		_ = ret_4
		high_4 = high_4
	}
	p10 = int(int32((hxrt.Int32Wrap(p10) << uint(16))))
	low_4 = int(int32((hxrt.Int32Wrap(low_4) + hxrt.Int32Wrap(p10))))
	if haxe___Int32__Int32_Impl__ucompare(low_4, p10) < 0 {
		hx_post_6 := high_4
		high_4 = int(int32((high_4 + 1)))
		ret_5 := hx_post_6
		_ = ret_5
		high_4 = high_4
	}
	high_4 = int(int32((hxrt.Int32Wrap(high_4) + hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(a_low) * hxrt.Int32Wrap(b_high_2))))) + hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(a_high) * hxrt.Int32Wrap(b_low_2))))))))))))
	x_4 := New_haxe___Int64_____Int64(high_4, low_4)
	var this1_4 *haxe___Int64_____Int64
	this1_4 = x_4
	value_4 := this1_4
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("mul="), haxe___Int64__Int64_Impl__toString(value_4)))
	hxrt.Println(v_4)
	positive := haxe___Int64__Int64_Impl__divMod(haxe__Int64Helper_parseString(hxrt.StringFromLiteral("123456789")), func() *haxe___Int64_____Int64 {
		x_5 := New_haxe___Int64_____Int64(0, 97)
		var this1_5 *haxe___Int64_____Int64
		this1_5 = x_5
		return this1_5
	}())
	value_5 := func(hx_obj_7 map[string]any) *haxe___Int64_____Int64 {
		hx_field_8 := hx_obj_7["quotient"]
		if hx_field_8 == nil {
			var hx_zero_9 *haxe___Int64_____Int64
			return hx_zero_9
		}
		return hx_field_8.(*haxe___Int64_____Int64)
	}(positive)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("div_q="), haxe___Int64__Int64_Impl__toString(value_5)))
	hxrt.Println(v_5)
	value_6 := func(hx_obj_10 map[string]any) *haxe___Int64_____Int64 {
		hx_field_11 := hx_obj_10["modulus"]
		if hx_field_11 == nil {
			var hx_zero_12 *haxe___Int64_____Int64
			return hx_zero_12
		}
		return hx_field_11.(*haxe___Int64_____Int64)
	}(positive)
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("div_r="), haxe___Int64__Int64_Impl__toString(value_6)))
	hxrt.Println(v_6)
	negative := haxe___Int64__Int64_Impl__divMod(haxe__Int64Helper_parseString(hxrt.StringFromLiteral("-123456789")), func() *haxe___Int64_____Int64 {
		x_6 := New_haxe___Int64_____Int64(0, 97)
		var this1_6 *haxe___Int64_____Int64
		this1_6 = x_6
		return this1_6
	}())
	value_7 := func(hx_obj_13 map[string]any) *haxe___Int64_____Int64 {
		hx_field_14 := hx_obj_13["quotient"]
		if hx_field_14 == nil {
			var hx_zero_15 *haxe___Int64_____Int64
			return hx_zero_15
		}
		return hx_field_14.(*haxe___Int64_____Int64)
	}(negative)
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("div_neg_q="), haxe___Int64__Int64_Impl__toString(value_7)))
	hxrt.Println(v_7)
	value_8 := func(hx_obj_16 map[string]any) *haxe___Int64_____Int64 {
		hx_field_17 := hx_obj_16["modulus"]
		if hx_field_17 == nil {
			var hx_zero_18 *haxe___Int64_____Int64
			return hx_zero_18
		}
		return hx_field_17.(*haxe___Int64_____Int64)
	}(negative)
	var v_8 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("div_neg_r="), haxe___Int64__Int64_Impl__toString(value_8)))
	hxrt.Println(v_8)
	var v_9 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("cmp="), func() int {
		var a_low_1 int
		var a_high_1 int
		a_high_1 = -1
		a_low_1 = -1
		var b_low_3 int
		var b_high_3 int
		b_high_3 = 0
		b_low_3 = 1
		v_10 := int(int32((hxrt.Int32Wrap(a_high_1) - hxrt.Int32Wrap(b_high_3))))
		var hx_if_19 int
		if v_10 != 0 {
			hx_if_19 = v_10
		} else {
			hx_if_19 = haxe___Int32__Int32_Impl__ucompare(a_low_1, b_low_3)
		}
		v_10 = hx_if_19
		var hx_if_22 int
		if a_high_1 < 0 {
			var hx_if_20 int
			if b_high_3 < 0 {
				hx_if_20 = v_10
			} else {
				hx_if_20 = -1
			}
			hx_if_22 = hx_if_20
		} else {
			var hx_if_21 int
			if b_high_3 >= 0 {
				hx_if_21 = v_10
			} else {
				hx_if_21 = 1
			}
			hx_if_22 = hx_if_21
		}
		return hx_if_22
	}()))
	hxrt.Println(v_9)
	var v_11 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("ucmp="), func() int {
		var a_low_2 int
		var a_high_2 int
		a_high_2 = -1
		a_low_2 = -1
		var b_low_4 int
		var b_high_4 int
		b_high_4 = 0
		b_low_4 = 1
		v_12 := haxe___Int32__Int32_Impl__ucompare(a_high_2, b_high_4)
		var hx_if_23 int
		if v_12 != 0 {
			hx_if_23 = v_12
		} else {
			hx_if_23 = haxe___Int32__Int32_Impl__ucompare(a_low_2, b_low_4)
		}
		return hx_if_23
	}()))
	hxrt.Println(v_11)
	var a_low_3 int
	var a_high_3 int
	a_high_3 = 0
	a_low_3 = 1
	b_1 := 40
	b_1 = int(int32((hxrt.Int32Wrap(b_1) & hxrt.Int32Wrap(63))))
	var hx_if_25 *haxe___Int64_____Int64
	if b_1 == 0 {
		high_5 := a_high_3
		low_5 := a_low_3
		x_7 := New_haxe___Int64_____Int64(high_5, low_5)
		var this1_7 *haxe___Int64_____Int64
		this1_7 = x_7
		hx_if_25 = this1_7
	} else {
		var hx_if_24 *haxe___Int64_____Int64
		if b_1 < 32 {
			high_6 := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(a_high_3) << uint(b_1))))) | hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(a_low_3)) >> uint(int(int32((hxrt.Int32Wrap(32) - hxrt.Int32Wrap(b_1)))))))))))))
			low_6 := int(int32((hxrt.Int32Wrap(a_low_3) << uint(b_1))))
			x_8 := New_haxe___Int64_____Int64(high_6, low_6)
			var this1_8 *haxe___Int64_____Int64
			this1_8 = x_8
			hx_if_24 = this1_8
		} else {
			high_7 := int(int32((hxrt.Int32Wrap(a_low_3) << uint(int(int32((hxrt.Int32Wrap(b_1) - hxrt.Int32Wrap(32))))))))
			x_9 := New_haxe___Int64_____Int64(high_7, 0)
			var this1_9 *haxe___Int64_____Int64
			this1_9 = x_9
			hx_if_24 = this1_9
		}
		hx_if_25 = hx_if_24
	}
	value_9 := hx_if_25
	var v_13 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("shl="), haxe___Int64__Int64_Impl__toString(value_9)))
	hxrt.Println(v_13)
	a_1 := haxe__Int64Helper_parseString(hxrt.StringFromLiteral("-8"))
	b_2 := 1
	b_2 = int(int32((hxrt.Int32Wrap(b_2) & hxrt.Int32Wrap(63))))
	var hx_if_27 *haxe___Int64_____Int64
	if b_2 == 0 {
		high_8 := a_1.high
		low_7 := a_1.low
		x_10 := New_haxe___Int64_____Int64(high_8, low_7)
		var this1_10 *haxe___Int64_____Int64
		this1_10 = x_10
		hx_if_27 = this1_10
	} else {
		var hx_if_26 *haxe___Int64_____Int64
		if b_2 < 32 {
			high_9 := int(int32((hxrt.Int32Wrap(a_1.high) >> uint(b_2))))
			low_8 := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(a_1.high) << uint(int(int32((hxrt.Int32Wrap(32) - hxrt.Int32Wrap(b_2))))))))) | hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(a_1.low)) >> uint(b_2)))))))))
			x_11 := New_haxe___Int64_____Int64(high_9, low_8)
			var this1_11 *haxe___Int64_____Int64
			this1_11 = x_11
			hx_if_26 = this1_11
		} else {
			high_10 := int(int32((hxrt.Int32Wrap(a_1.high) >> uint(31))))
			low_9 := int(int32((hxrt.Int32Wrap(a_1.high) >> uint(int(int32((hxrt.Int32Wrap(b_2) - hxrt.Int32Wrap(32))))))))
			x_12 := New_haxe___Int64_____Int64(high_10, low_9)
			var this1_12 *haxe___Int64_____Int64
			this1_12 = x_12
			hx_if_26 = this1_12
		}
		hx_if_27 = hx_if_26
	}
	value_10 := hx_if_27
	var v_14 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("shr="), haxe___Int64__Int64_Impl__toString(value_10)))
	hxrt.Println(v_14)
	a_2 := haxe__Int64Helper_parseString(hxrt.StringFromLiteral("-1"))
	b_3 := 1
	b_3 = int(int32((hxrt.Int32Wrap(b_3) & hxrt.Int32Wrap(63))))
	var hx_if_29 *haxe___Int64_____Int64
	if b_3 == 0 {
		high_11 := a_2.high
		low_10 := a_2.low
		x_13 := New_haxe___Int64_____Int64(high_11, low_10)
		var this1_13 *haxe___Int64_____Int64
		this1_13 = x_13
		hx_if_29 = this1_13
	} else {
		var hx_if_28 *haxe___Int64_____Int64
		if b_3 < 32 {
			high_12 := int(int32(int32((uint32(hxrt.Int32Wrap(a_2.high)) >> uint(b_3)))))
			low_11 := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(a_2.high) << uint(int(int32((hxrt.Int32Wrap(32) - hxrt.Int32Wrap(b_3))))))))) | hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(a_2.low)) >> uint(b_3)))))))))
			x_14 := New_haxe___Int64_____Int64(high_12, low_11)
			var this1_14 *haxe___Int64_____Int64
			this1_14 = x_14
			hx_if_28 = this1_14
		} else {
			low_12 := int(int32(int32((uint32(hxrt.Int32Wrap(a_2.high)) >> uint(int(int32((hxrt.Int32Wrap(b_3) - hxrt.Int32Wrap(32)))))))))
			x_15 := New_haxe___Int64_____Int64(0, low_12)
			var this1_15 *haxe___Int64_____Int64
			this1_15 = x_15
			hx_if_28 = this1_15
		}
		hx_if_29 = hx_if_28
	}
	value_11 := hx_if_29
	var v_15 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("ushr="), haxe___Int64__Int64_Impl__toString(value_11)))
	hxrt.Println(v_15)
	value_12 := haxe__Int64Helper_fromFloat(9007199254740991.0)
	var v_16 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("from_float="), haxe___Int64__Int64_Impl__toString(value_12)))
	hxrt.Println(v_16)
	var v_17 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("to_int_ok="), func() int {
		var x_low int
		var x_high int
		x_high = 0
		x_low = 2147483647
		if x_high != int(int32((hxrt.Int32Wrap(x_low) >> uint(31)))) {
			hxrt.Throw(hxrt.StringFromLiteral("Overflow"))
		}
		return x_low
	}()))
	hxrt.Println(v_17)
	hxrt.TryCatch(func() {
		x_16 := haxe__Int64Helper_parseString(hxrt.StringFromLiteral("2147483648"))
		if x_16.high != int(int32((hxrt.Int32Wrap(x_16.low) >> uint(31)))) {
			hxrt.Throw(hxrt.StringFromLiteral("Overflow"))
		}
		_ = x_16.low
		hxrt.Println(any(hxrt.StringFromLiteral("to_int_overflow=missing")))
	}, func(hx_caught_30 any) {
		e := hx_caught_30
		var v_18 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("to_int_overflow="), hxrt.StdString(e)))
		hxrt.Println(v_18)
	})
	var round_low int
	var round_high int
	round_high = 2147483647
	round_low = -12345
	var v_19 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("round_high="), round_high))
	hxrt.Println(v_19)
	var v_20 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("round_low="), round_low))
	hxrt.Println(v_20)
}
