package main

import "snapshot/hxrt"

func haxe___Int64__Int64_Impl__divMod(dividend *haxe___Int64_____Int64, divisor *haxe___Int64_____Int64) map[string]any {
	if divisor.high == 0 {
		_g := divisor.low
		switch _g {
		case 0:
			hxrt.Throw(hxrt.StringFromLiteral("divide by zero"))
		case 1:
			hx_obj_1 := map[string]any{}
			high := dividend.high
			low := dividend.low
			x := New_haxe___Int64_____Int64(high, low)
			var this1 *haxe___Int64_____Int64
			this1 = x
			hx_obj_1["quotient"] = this1
			x_1 := New_haxe___Int64_____Int64(0, 0)
			var this1_1 *haxe___Int64_____Int64
			this1_1 = x_1
			hx_obj_1["modulus"] = this1_1
			return hx_obj_1
		}
	}
	divSign := ((dividend.high < 0) != (divisor.high < 0))
	var hx_if_3 *haxe___Int64_____Int64
	if dividend.high < 0 {
		high_1 := int(int32(^int32(dividend.high)))
		low_1 := int(int32((hxrt.Int32Wrap(int(int32(^int32(dividend.low)))) + hxrt.Int32Wrap(1))))
		if low_1 == 0 {
			hx_post_2 := high_1
			high_1 = int(int32((high_1 + 1)))
			ret := hx_post_2
			_ = ret
			high_1 = high_1
		}
		x_2 := New_haxe___Int64_____Int64(high_1, low_1)
		var this1_2 *haxe___Int64_____Int64
		this1_2 = x_2
		hx_if_3 = this1_2
	} else {
		high_2 := dividend.high
		low_2 := dividend.low
		x_3 := New_haxe___Int64_____Int64(high_2, low_2)
		var this1_3 *haxe___Int64_____Int64
		this1_3 = x_3
		hx_if_3 = this1_3
	}
	modulus := hx_if_3
	var hx_if_5 *haxe___Int64_____Int64
	if divisor.high < 0 {
		high_3 := int(int32(^int32(divisor.high)))
		low_3 := int(int32((hxrt.Int32Wrap(int(int32(^int32(divisor.low)))) + hxrt.Int32Wrap(1))))
		if low_3 == 0 {
			hx_post_4 := high_3
			high_3 = int(int32((high_3 + 1)))
			ret_1 := hx_post_4
			_ = ret_1
			high_3 = high_3
		}
		x_4 := New_haxe___Int64_____Int64(high_3, low_3)
		var this1_4 *haxe___Int64_____Int64
		this1_4 = x_4
		hx_if_5 = this1_4
	} else {
		hx_if_5 = divisor
	}
	divisor = hx_if_5
	x_5 := New_haxe___Int64_____Int64(0, 0)
	var this1_5 *haxe___Int64_____Int64
	this1_5 = x_5
	quotient := this1_5
	x_6 := New_haxe___Int64_____Int64(0, 1)
	var this1_6 *haxe___Int64_____Int64
	this1_6 = x_6
	mask := this1_6
	for !(divisor.high < 0) {
		v := haxe___Int32__Int32_Impl__ucompare(divisor.high, modulus.high)
		var hx_if_6 int
		if v != 0 {
			hx_if_6 = v
		} else {
			hx_if_6 = haxe___Int32__Int32_Impl__ucompare(divisor.low, modulus.low)
		}
		cmp := hx_if_6
		b := 1
		b = int(int32((hxrt.Int32Wrap(b) & hxrt.Int32Wrap(63))))
		var hx_if_8 *haxe___Int64_____Int64
		if b == 0 {
			high_4 := divisor.high
			low_4 := divisor.low
			x_7 := New_haxe___Int64_____Int64(high_4, low_4)
			var this1_7 *haxe___Int64_____Int64
			this1_7 = x_7
			hx_if_8 = this1_7
		} else {
			var hx_if_7 *haxe___Int64_____Int64
			if b < 32 {
				high_5 := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(divisor.high) << uint(b))))) | hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(divisor.low)) >> uint(int(int32((hxrt.Int32Wrap(32) - hxrt.Int32Wrap(b)))))))))))))
				low_5 := int(int32((hxrt.Int32Wrap(divisor.low) << uint(b))))
				x_8 := New_haxe___Int64_____Int64(high_5, low_5)
				var this1_8 *haxe___Int64_____Int64
				this1_8 = x_8
				hx_if_7 = this1_8
			} else {
				high_6 := int(int32((hxrt.Int32Wrap(divisor.low) << uint(int(int32((hxrt.Int32Wrap(b) - hxrt.Int32Wrap(32))))))))
				x_9 := New_haxe___Int64_____Int64(high_6, 0)
				var this1_9 *haxe___Int64_____Int64
				this1_9 = x_9
				hx_if_7 = this1_9
			}
			hx_if_8 = hx_if_7
		}
		divisor = hx_if_8
		b_1 := 1
		b_1 = int(int32((hxrt.Int32Wrap(b_1) & hxrt.Int32Wrap(63))))
		var hx_if_10 *haxe___Int64_____Int64
		if b_1 == 0 {
			high_7 := mask.high
			low_6 := mask.low
			x_10 := New_haxe___Int64_____Int64(high_7, low_6)
			var this1_10 *haxe___Int64_____Int64
			this1_10 = x_10
			hx_if_10 = this1_10
		} else {
			var hx_if_9 *haxe___Int64_____Int64
			if b_1 < 32 {
				high_8 := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(mask.high) << uint(b_1))))) | hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(mask.low)) >> uint(int(int32((hxrt.Int32Wrap(32) - hxrt.Int32Wrap(b_1)))))))))))))
				low_7 := int(int32((hxrt.Int32Wrap(mask.low) << uint(b_1))))
				x_11 := New_haxe___Int64_____Int64(high_8, low_7)
				var this1_11 *haxe___Int64_____Int64
				this1_11 = x_11
				hx_if_9 = this1_11
			} else {
				high_9 := int(int32((hxrt.Int32Wrap(mask.low) << uint(int(int32((hxrt.Int32Wrap(b_1) - hxrt.Int32Wrap(32))))))))
				x_12 := New_haxe___Int64_____Int64(high_9, 0)
				var this1_12 *haxe___Int64_____Int64
				this1_12 = x_12
				hx_if_9 = this1_12
			}
			hx_if_10 = hx_if_9
		}
		mask = hx_if_10
		if cmp >= 0 {
			break
		}
	}
	for func() bool {
		var b_low int
		var b_high int
		b_high = 0
		b_low = 0
		return ((mask.high != b_high) || (mask.low != b_low))
	}() {
		if func() int {
			v_1 := haxe___Int32__Int32_Impl__ucompare(modulus.high, divisor.high)
			var hx_if_12 int
			if v_1 != 0 {
				hx_if_12 = v_1
			} else {
				hx_if_12 = haxe___Int32__Int32_Impl__ucompare(modulus.low, divisor.low)
			}
			return hx_if_12
		}() >= 0 {
			high_10 := int(int32((hxrt.Int32Wrap(quotient.high) | hxrt.Int32Wrap(mask.high))))
			low_8 := int(int32((hxrt.Int32Wrap(quotient.low) | hxrt.Int32Wrap(mask.low))))
			x_13 := New_haxe___Int64_____Int64(high_10, low_8)
			var this1_13 *haxe___Int64_____Int64
			this1_13 = x_13
			quotient = this1_13
			high_11 := int(int32((hxrt.Int32Wrap(modulus.high) - hxrt.Int32Wrap(divisor.high))))
			low_9 := int(int32((hxrt.Int32Wrap(modulus.low) - hxrt.Int32Wrap(divisor.low))))
			if haxe___Int32__Int32_Impl__ucompare(modulus.low, divisor.low) < 0 {
				hx_post_11 := high_11
				high_11 = int(int32((high_11 - 1)))
				ret_2 := hx_post_11
				_ = ret_2
				high_11 = high_11
			}
			x_14 := New_haxe___Int64_____Int64(high_11, low_9)
			var this1_14 *haxe___Int64_____Int64
			this1_14 = x_14
			modulus = this1_14
		}
		b_2 := 1
		b_2 = int(int32((hxrt.Int32Wrap(b_2) & hxrt.Int32Wrap(63))))
		var hx_if_14 *haxe___Int64_____Int64
		if b_2 == 0 {
			high_12 := mask.high
			low_10 := mask.low
			x_15 := New_haxe___Int64_____Int64(high_12, low_10)
			var this1_15 *haxe___Int64_____Int64
			this1_15 = x_15
			hx_if_14 = this1_15
		} else {
			var hx_if_13 *haxe___Int64_____Int64
			if b_2 < 32 {
				high_13 := int(int32(int32((uint32(hxrt.Int32Wrap(mask.high)) >> uint(b_2)))))
				low_11 := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(mask.high) << uint(int(int32((hxrt.Int32Wrap(32) - hxrt.Int32Wrap(b_2))))))))) | hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(mask.low)) >> uint(b_2)))))))))
				x_16 := New_haxe___Int64_____Int64(high_13, low_11)
				var this1_16 *haxe___Int64_____Int64
				this1_16 = x_16
				hx_if_13 = this1_16
			} else {
				low_12 := int(int32(int32((uint32(hxrt.Int32Wrap(mask.high)) >> uint(int(int32((hxrt.Int32Wrap(b_2) - hxrt.Int32Wrap(32)))))))))
				x_17 := New_haxe___Int64_____Int64(0, low_12)
				var this1_17 *haxe___Int64_____Int64
				this1_17 = x_17
				hx_if_13 = this1_17
			}
			hx_if_14 = hx_if_13
		}
		mask = hx_if_14
		b_3 := 1
		b_3 = int(int32((hxrt.Int32Wrap(b_3) & hxrt.Int32Wrap(63))))
		var hx_if_16 *haxe___Int64_____Int64
		if b_3 == 0 {
			high_14 := divisor.high
			low_13 := divisor.low
			x_18 := New_haxe___Int64_____Int64(high_14, low_13)
			var this1_18 *haxe___Int64_____Int64
			this1_18 = x_18
			hx_if_16 = this1_18
		} else {
			var hx_if_15 *haxe___Int64_____Int64
			if b_3 < 32 {
				high_15 := int(int32(int32((uint32(hxrt.Int32Wrap(divisor.high)) >> uint(b_3)))))
				low_14 := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(divisor.high) << uint(int(int32((hxrt.Int32Wrap(32) - hxrt.Int32Wrap(b_3))))))))) | hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(divisor.low)) >> uint(b_3)))))))))
				x_19 := New_haxe___Int64_____Int64(high_15, low_14)
				var this1_19 *haxe___Int64_____Int64
				this1_19 = x_19
				hx_if_15 = this1_19
			} else {
				low_15 := int(int32(int32((uint32(hxrt.Int32Wrap(divisor.high)) >> uint(int(int32((hxrt.Int32Wrap(b_3) - hxrt.Int32Wrap(32)))))))))
				x_20 := New_haxe___Int64_____Int64(0, low_15)
				var this1_20 *haxe___Int64_____Int64
				this1_20 = x_20
				hx_if_15 = this1_20
			}
			hx_if_16 = hx_if_15
		}
		divisor = hx_if_16
	}
	if divSign {
		high_16 := int(int32(^int32(quotient.high)))
		low_16 := int(int32((hxrt.Int32Wrap(int(int32(^int32(quotient.low)))) + hxrt.Int32Wrap(1))))
		if low_16 == 0 {
			hx_post_17 := high_16
			high_16 = int(int32((high_16 + 1)))
			ret_3 := hx_post_17
			_ = ret_3
			high_16 = high_16
		}
		x_21 := New_haxe___Int64_____Int64(high_16, low_16)
		var this1_21 *haxe___Int64_____Int64
		this1_21 = x_21
		quotient = this1_21
	}
	if dividend.high < 0 {
		high_17 := int(int32(^int32(modulus.high)))
		low_17 := int(int32((hxrt.Int32Wrap(int(int32(^int32(modulus.low)))) + hxrt.Int32Wrap(1))))
		if low_17 == 0 {
			hx_post_18 := high_17
			high_17 = int(int32((high_17 + 1)))
			ret_4 := hx_post_18
			_ = ret_4
			high_17 = high_17
		}
		x_22 := New_haxe___Int64_____Int64(high_17, low_17)
		var this1_22 *haxe___Int64_____Int64
		this1_22 = x_22
		modulus = this1_22
	}
	hx_obj_19 := map[string]any{}
	hx_obj_19["quotient"] = quotient
	hx_obj_19["modulus"] = modulus
	return hx_obj_19
}

func haxe___Int64__Int64_Impl__toString(this1 *haxe___Int64_____Int64) *string {
	i := this1
	if func() bool {
		var b_low int
		var b_high int
		b_high = 0
		b_low = 0
		return ((i.high == b_high) && (i.low == b_low))
	}() {
		return hxrt.StringFromLiteral("0")
	}
	str := hxrt.StringFromLiteral("")
	neg := false
	if i.high < 0 {
		neg = true
	}
	x := New_haxe___Int64_____Int64(0, 10)
	var this1_1 *haxe___Int64_____Int64
	this1_1 = x
	ten := this1_1
	for func() bool {
		var b_low_1 int
		var b_high_1 int
		b_high_1 = 0
		b_low_1 = 0
		return ((i.high != b_high_1) || (i.low != b_low_1))
	}() {
		r := haxe___Int64__Int64_Impl__divMod(i, ten)
		if func(hx_obj_34 map[string]any) *haxe___Int64_____Int64 {
			hx_field_35 := hx_obj_34["modulus"]
			if hx_field_35 == nil {
				var hx_zero_36 *haxe___Int64_____Int64
				return hx_zero_36
			}
			return hx_field_35.(*haxe___Int64_____Int64)
		}(r).high < 0 {
			str = hxrt.StringConcatAny(func() int {
				var this_low int
				var this_high int
				_ = this_high
				x_1 := func(hx_obj_20 map[string]any) *haxe___Int64_____Int64 {
					hx_field_21 := hx_obj_20["modulus"]
					if hx_field_21 == nil {
						var hx_zero_22 *haxe___Int64_____Int64
						return hx_zero_22
					}
					return hx_field_21.(*haxe___Int64_____Int64)
				}(r)
				high := int(int32(^int32(x_1.high)))
				low := int(int32((hxrt.Int32Wrap(int(int32(^int32(x_1.low)))) + hxrt.Int32Wrap(1))))
				if low == 0 {
					hx_post_23 := high
					high = int(int32((high + 1)))
					ret := hx_post_23
					_ = ret
					high = high
				}
				this_high = high
				this_low = low
				return this_low
			}(), str)
			x_2 := func(hx_obj_24 map[string]any) *haxe___Int64_____Int64 {
				hx_field_25 := hx_obj_24["quotient"]
				if hx_field_25 == nil {
					var hx_zero_26 *haxe___Int64_____Int64
					return hx_zero_26
				}
				return hx_field_25.(*haxe___Int64_____Int64)
			}(r)
			high_1 := int(int32(^int32(x_2.high)))
			low_1 := int(int32((hxrt.Int32Wrap(int(int32(^int32(x_2.low)))) + hxrt.Int32Wrap(1))))
			if low_1 == 0 {
				hx_post_27 := high_1
				high_1 = int(int32((high_1 + 1)))
				ret_1 := hx_post_27
				_ = ret_1
				high_1 = high_1
			}
			x_3 := New_haxe___Int64_____Int64(high_1, low_1)
			var this1_2 *haxe___Int64_____Int64
			this1_2 = x_3
			i = this1_2
		} else {
			str = hxrt.StringConcatAny(func(hx_obj_28 map[string]any) *haxe___Int64_____Int64 {
				hx_field_29 := hx_obj_28["modulus"]
				if hx_field_29 == nil {
					var hx_zero_30 *haxe___Int64_____Int64
					return hx_zero_30
				}
				return hx_field_29.(*haxe___Int64_____Int64)
			}(r).low, str)
			i = func(hx_obj_31 map[string]any) *haxe___Int64_____Int64 {
				hx_field_32 := hx_obj_31["quotient"]
				if hx_field_32 == nil {
					var hx_zero_33 *haxe___Int64_____Int64
					return hx_zero_33
				}
				return hx_field_32.(*haxe___Int64_____Int64)
			}(r)
		}
	}
	if neg {
		str = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("-"), str)
	}
	return str
}

type I_haxe___Int64_____Int64 interface {
}

type haxe___Int64_____Int64 struct {
	__hx_this I_haxe___Int64_____Int64
	high      int
	low       int
}

func New_haxe___Int64_____Int64(high int, low int) *haxe___Int64_____Int64 {
	self := &haxe___Int64_____Int64{}
	self.__hx_this = self
	self.high = high
	self.low = low
	return self
}
