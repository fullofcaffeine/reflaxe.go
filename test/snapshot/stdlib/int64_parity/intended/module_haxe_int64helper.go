package main

import "snapshot/hxrt"

func haxe__Int64Helper_fromFloat(f float64) *haxe___Int64_____Int64 {
	if Math_isNaN(f) || !Math_isFinite(f) {
		hxrt.Throw(hxrt.StringFromLiteral("Number is NaN or Infinite"))
		var hx_throw_zero_32 *haxe___Int64_____Int64
		return hx_throw_zero_32
	}
	noFractions := (f - hxrt.FloatMod(f, float64(1)))
	if noFractions > 9007199254740991 {
		hxrt.Throw(hxrt.StringFromLiteral("Conversion overflow"))
		var hx_throw_zero_33 *haxe___Int64_____Int64
		return hx_throw_zero_33
	}
	if noFractions < -9007199254740991 {
		hxrt.Throw(hxrt.StringFromLiteral("Conversion underflow"))
		var hx_throw_zero_34 *haxe___Int64_____Int64
		return hx_throw_zero_34
	}
	x := New_haxe___Int64_____Int64(0, 0)
	var this1 *haxe___Int64_____Int64
	this1 = x
	result := this1
	neg := (noFractions < 0)
	var hx_if_35 float64
	if neg {
		hx_if_35 = -noFractions
	} else {
		hx_if_35 = noFractions
	}
	rest := hx_if_35
	i := 0
	for rest >= 1 {
		curr := hxrt.FloatMod(rest, float64(2))
		rest = (rest / float64(2))
		if curr >= 1 {
			var a_low int
			var a_high int
			a_high = 0
			a_low = 1
			b_1 := i
			b_1 = int(int32((hxrt.Int32Wrap(b_1) & hxrt.Int32Wrap(63))))
			var hx_if_37 *haxe___Int64_____Int64
			if b_1 == 0 {
				high := a_high
				low := a_low
				x_1 := New_haxe___Int64_____Int64(high, low)
				var this1_1 *haxe___Int64_____Int64
				this1_1 = x_1
				hx_if_37 = this1_1
			} else {
				var hx_if_36 *haxe___Int64_____Int64
				if b_1 < 32 {
					high_1 := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(a_high) << uint(b_1))))) | hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(a_low)) >> uint(int(int32((hxrt.Int32Wrap(32) - hxrt.Int32Wrap(b_1)))))))))))))
					low_1 := int(int32((hxrt.Int32Wrap(a_low) << uint(b_1))))
					x_2 := New_haxe___Int64_____Int64(high_1, low_1)
					var this1_2 *haxe___Int64_____Int64
					this1_2 = x_2
					hx_if_36 = this1_2
				} else {
					high_2 := int(int32((hxrt.Int32Wrap(a_low) << uint(int(int32((hxrt.Int32Wrap(b_1) - hxrt.Int32Wrap(32))))))))
					x_3 := New_haxe___Int64_____Int64(high_2, 0)
					var this1_3 *haxe___Int64_____Int64
					this1_3 = x_3
					hx_if_36 = this1_3
				}
				hx_if_37 = hx_if_36
			}
			b := hx_if_37
			high_3 := int(int32((hxrt.Int32Wrap(result.high) + hxrt.Int32Wrap(b.high))))
			low_2 := int(int32((hxrt.Int32Wrap(result.low) + hxrt.Int32Wrap(b.low))))
			if haxe___Int32__Int32_Impl__ucompare(low_2, result.low) < 0 {
				hx_post_38 := high_3
				high_3 = int(int32((high_3 + 1)))
				ret := hx_post_38
				_ = ret
				high_3 = high_3
			}
			x_4 := New_haxe___Int64_____Int64(high_3, low_2)
			var this1_4 *haxe___Int64_____Int64
			this1_4 = x_4
			result = this1_4
		}
		i = int(int32((i + 1)))
	}
	if neg {
		high_4 := int(int32(^int32(result.high)))
		low_3 := int(int32((hxrt.Int32Wrap(int(int32(^int32(result.low)))) + hxrt.Int32Wrap(1))))
		if low_3 == 0 {
			hx_post_39 := high_4
			high_4 = int(int32((high_4 + 1)))
			ret_1 := hx_post_39
			_ = ret_1
			high_4 = high_4
		}
		x_5 := New_haxe___Int64_____Int64(high_4, low_3)
		var this1_5 *haxe___Int64_____Int64
		this1_5 = x_5
		result = this1_5
	}
	return result
}

func haxe__Int64Helper_parseString(sParam *string) *haxe___Int64_____Int64 {
	var base_low int
	var base_high int
	base_high = 0
	base_low = 10
	x := New_haxe___Int64_____Int64(0, 0)
	var this1 *haxe___Int64_____Int64
	this1 = x
	current := this1
	x_1 := New_haxe___Int64_____Int64(0, 1)
	var this1_1 *haxe___Int64_____Int64
	this1_1 = x_1
	multiplier := this1_1
	sIsNegative := false
	s := StringTools_trim(sParam)
	if hxrt.StringEqualStringPtr(hxrt.StringCharAtStringPtr(s, 0), hxrt.StringFromLiteral("-")) {
		sIsNegative = true
		s = hxrt.StringSubstringStringPtr(s, 1, hxrt.StringLengthStringPtr(s))
	}
	len := hxrt.StringLengthStringPtr(s)
	_g := 0
	_g1 := len
	for _g < _g1 {
		hx_post_40 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_40
		digitInt := int(int32((hxrt.Int32Wrap(hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(s, int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(len)-hxrt.Int32Wrap(1)))))-hxrt.Int32Wrap(i))))))) - hxrt.Int32Wrap(48))))
		if (digitInt < 0) || (digitInt > 9) {
			hxrt.Throw(hxrt.StringFromLiteral("NumberFormatError"))
			var hx_throw_zero_41 *haxe___Int64_____Int64
			return hx_throw_zero_41
		}
		if digitInt != 0 {
			var digit_low int
			var digit_high int
			digit_high = int(int32((hxrt.Int32Wrap(digitInt) >> uint(31))))
			digit_low = digitInt
			if sIsNegative {
				var b_low int
				var b_high int
				mask := 65535
				al := int(int32((hxrt.Int32Wrap(multiplier.low) & hxrt.Int32Wrap(mask))))
				ah := int(int32(int32((uint32(hxrt.Int32Wrap(multiplier.low)) >> uint(16)))))
				bl := int(int32((hxrt.Int32Wrap(digit_low) & hxrt.Int32Wrap(mask))))
				bh := int(int32(int32((uint32(hxrt.Int32Wrap(digit_low)) >> uint(16)))))
				p00 := int(int32((hxrt.Int32Wrap(al) * hxrt.Int32Wrap(bl))))
				p10 := int(int32((hxrt.Int32Wrap(ah) * hxrt.Int32Wrap(bl))))
				p01 := int(int32((hxrt.Int32Wrap(al) * hxrt.Int32Wrap(bh))))
				p11 := int(int32((hxrt.Int32Wrap(ah) * hxrt.Int32Wrap(bh))))
				low := p00
				high := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(p11) + hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(p01)) >> uint(16)))))))))) + hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(p10)) >> uint(16)))))))))
				p01 = int(int32((hxrt.Int32Wrap(p01) << uint(16))))
				low = int(int32((hxrt.Int32Wrap(low) + hxrt.Int32Wrap(p01))))
				if haxe___Int32__Int32_Impl__ucompare(low, p01) < 0 {
					hx_post_42 := high
					high = int(int32((high + 1)))
					ret := hx_post_42
					_ = ret
					high = high
				}
				p10 = int(int32((hxrt.Int32Wrap(p10) << uint(16))))
				low = int(int32((hxrt.Int32Wrap(low) + hxrt.Int32Wrap(p10))))
				if haxe___Int32__Int32_Impl__ucompare(low, p10) < 0 {
					hx_post_43 := high
					high = int(int32((high + 1)))
					ret_1 := hx_post_43
					_ = ret_1
					high = high
				}
				high = int(int32((hxrt.Int32Wrap(high) + hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(multiplier.low) * hxrt.Int32Wrap(digit_high))))) + hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(multiplier.high) * hxrt.Int32Wrap(digit_low))))))))))))
				b_high = high
				b_low = low
				high_1 := int(int32((hxrt.Int32Wrap(current.high) - hxrt.Int32Wrap(b_high))))
				low_1 := int(int32((hxrt.Int32Wrap(current.low) - hxrt.Int32Wrap(b_low))))
				if haxe___Int32__Int32_Impl__ucompare(current.low, b_low) < 0 {
					hx_post_44 := high_1
					high_1 = int(int32((high_1 - 1)))
					ret_2 := hx_post_44
					_ = ret_2
					high_1 = high_1
				}
				x_2 := New_haxe___Int64_____Int64(high_1, low_1)
				var this1_2 *haxe___Int64_____Int64
				this1_2 = x_2
				current = this1_2
				if !(current.high < 0) {
					hxrt.Throw(hxrt.StringFromLiteral("NumberFormatError: Underflow"))
					var hx_throw_zero_45 *haxe___Int64_____Int64
					return hx_throw_zero_45
				}
			} else {
				var b_low_1 int
				var b_high_1 int
				mask_1 := 65535
				al_1 := int(int32((hxrt.Int32Wrap(multiplier.low) & hxrt.Int32Wrap(mask_1))))
				ah_1 := int(int32(int32((uint32(hxrt.Int32Wrap(multiplier.low)) >> uint(16)))))
				bl_1 := int(int32((hxrt.Int32Wrap(digit_low) & hxrt.Int32Wrap(mask_1))))
				bh_1 := int(int32(int32((uint32(hxrt.Int32Wrap(digit_low)) >> uint(16)))))
				p00_1 := int(int32((hxrt.Int32Wrap(al_1) * hxrt.Int32Wrap(bl_1))))
				p10_1 := int(int32((hxrt.Int32Wrap(ah_1) * hxrt.Int32Wrap(bl_1))))
				p01_1 := int(int32((hxrt.Int32Wrap(al_1) * hxrt.Int32Wrap(bh_1))))
				p11_1 := int(int32((hxrt.Int32Wrap(ah_1) * hxrt.Int32Wrap(bh_1))))
				low_2 := p00_1
				high_2 := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(p11_1) + hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(p01_1)) >> uint(16)))))))))) + hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(p10_1)) >> uint(16)))))))))
				p01_1 = int(int32((hxrt.Int32Wrap(p01_1) << uint(16))))
				low_2 = int(int32((hxrt.Int32Wrap(low_2) + hxrt.Int32Wrap(p01_1))))
				if haxe___Int32__Int32_Impl__ucompare(low_2, p01_1) < 0 {
					hx_post_46 := high_2
					high_2 = int(int32((high_2 + 1)))
					ret_3 := hx_post_46
					_ = ret_3
					high_2 = high_2
				}
				p10_1 = int(int32((hxrt.Int32Wrap(p10_1) << uint(16))))
				low_2 = int(int32((hxrt.Int32Wrap(low_2) + hxrt.Int32Wrap(p10_1))))
				if haxe___Int32__Int32_Impl__ucompare(low_2, p10_1) < 0 {
					hx_post_47 := high_2
					high_2 = int(int32((high_2 + 1)))
					ret_4 := hx_post_47
					_ = ret_4
					high_2 = high_2
				}
				high_2 = int(int32((hxrt.Int32Wrap(high_2) + hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(multiplier.low) * hxrt.Int32Wrap(digit_high))))) + hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(multiplier.high) * hxrt.Int32Wrap(digit_low))))))))))))
				b_high_1 = high_2
				b_low_1 = low_2
				high_3 := int(int32((hxrt.Int32Wrap(current.high) + hxrt.Int32Wrap(b_high_1))))
				low_3 := int(int32((hxrt.Int32Wrap(current.low) + hxrt.Int32Wrap(b_low_1))))
				if haxe___Int32__Int32_Impl__ucompare(low_3, current.low) < 0 {
					hx_post_48 := high_3
					high_3 = int(int32((high_3 + 1)))
					ret_5 := hx_post_48
					_ = ret_5
					high_3 = high_3
				}
				x_3 := New_haxe___Int64_____Int64(high_3, low_3)
				var this1_3 *haxe___Int64_____Int64
				this1_3 = x_3
				current = this1_3
				if current.high < 0 {
					hxrt.Throw(hxrt.StringFromLiteral("NumberFormatError: Overflow"))
					var hx_throw_zero_49 *haxe___Int64_____Int64
					return hx_throw_zero_49
				}
			}
		}
		mask_2 := 65535
		al_2 := int(int32((hxrt.Int32Wrap(multiplier.low) & hxrt.Int32Wrap(mask_2))))
		ah_2 := int(int32(int32((uint32(hxrt.Int32Wrap(multiplier.low)) >> uint(16)))))
		bl_2 := int(int32((hxrt.Int32Wrap(base_low) & hxrt.Int32Wrap(mask_2))))
		bh_2 := int(int32(int32((uint32(hxrt.Int32Wrap(base_low)) >> uint(16)))))
		p00_2 := int(int32((hxrt.Int32Wrap(al_2) * hxrt.Int32Wrap(bl_2))))
		p10_2 := int(int32((hxrt.Int32Wrap(ah_2) * hxrt.Int32Wrap(bl_2))))
		p01_2 := int(int32((hxrt.Int32Wrap(al_2) * hxrt.Int32Wrap(bh_2))))
		p11_2 := int(int32((hxrt.Int32Wrap(ah_2) * hxrt.Int32Wrap(bh_2))))
		low_4 := p00_2
		high_4 := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(p11_2) + hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(p01_2)) >> uint(16)))))))))) + hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(p10_2)) >> uint(16)))))))))
		p01_2 = int(int32((hxrt.Int32Wrap(p01_2) << uint(16))))
		low_4 = int(int32((hxrt.Int32Wrap(low_4) + hxrt.Int32Wrap(p01_2))))
		if haxe___Int32__Int32_Impl__ucompare(low_4, p01_2) < 0 {
			hx_post_50 := high_4
			high_4 = int(int32((high_4 + 1)))
			ret_6 := hx_post_50
			_ = ret_6
			high_4 = high_4
		}
		p10_2 = int(int32((hxrt.Int32Wrap(p10_2) << uint(16))))
		low_4 = int(int32((hxrt.Int32Wrap(low_4) + hxrt.Int32Wrap(p10_2))))
		if haxe___Int32__Int32_Impl__ucompare(low_4, p10_2) < 0 {
			hx_post_51 := high_4
			high_4 = int(int32((high_4 + 1)))
			ret_7 := hx_post_51
			_ = ret_7
			high_4 = high_4
		}
		high_4 = int(int32((hxrt.Int32Wrap(high_4) + hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(multiplier.low) * hxrt.Int32Wrap(base_high))))) + hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(multiplier.high) * hxrt.Int32Wrap(base_low))))))))))))
		x_4 := New_haxe___Int64_____Int64(high_4, low_4)
		var this1_4 *haxe___Int64_____Int64
		this1_4 = x_4
		multiplier = this1_4
	}
	return current
}
