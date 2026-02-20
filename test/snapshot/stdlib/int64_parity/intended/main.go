package main

import (
	"bytes"
	"compress/zlib"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"io"
	"math"
	"path/filepath"
	"reflect"
	"snapshot/hxrt"
	"strings"
	"time"
)

func emit(label *string, value *haxe___Int64_____Int64) {
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringConcatAny(label, hxrt.StringFromLiteral("=")), haxe___Int64__Int64_Impl__toString(value)))
}

func main() {
	max := haxe__Int64Helper_parseString(hxrt.StringFromLiteral("9223372036854775807"))
	_ = max
	min := haxe__Int64Helper_parseString(hxrt.StringFromLiteral("-9223372036854775808"))
	_ = min
	var b_low int
	_ = b_low
	var b_high int
	_ = b_high
	b_high = 0
	b_low = 1
	high := int(int32((hxrt.Int32Wrap(max.high) + hxrt.Int32Wrap(b_high))))
	_ = high
	low := int(int32((hxrt.Int32Wrap(max.low) + hxrt.Int32Wrap(b_low))))
	_ = low
	if haxe___Int32__Int32_Impl__ucompare(low, max.low) < 0 {
		hx_post_1 := high
		high = int(int32((high + 1)))
		ret := hx_post_1
		_ = ret
		high = high
		_ = ret
	}
	x := New_haxe___Int64_____Int64(high, low)
	_ = x
	var this1 *haxe___Int64_____Int64
	_ = this1
	this1 = x
	value := this1
	_ = value
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringFromLiteral("wrap_add"), hxrt.StringFromLiteral("=")), haxe___Int64__Int64_Impl__toString(value)))
	var b_low_1 int
	_ = b_low_1
	var b_high_1 int
	_ = b_high_1
	b_high_1 = 0
	b_low_1 = 1
	high_1 := int(int32((hxrt.Int32Wrap(min.high) - hxrt.Int32Wrap(b_high_1))))
	_ = high_1
	low_1 := int(int32((hxrt.Int32Wrap(min.low) - hxrt.Int32Wrap(b_low_1))))
	_ = low_1
	if haxe___Int32__Int32_Impl__ucompare(min.low, b_low_1) < 0 {
		hx_post_2 := high_1
		high_1 = int(int32((high_1 - 1)))
		ret_1 := hx_post_2
		_ = ret_1
		high_1 = high_1
		_ = ret_1
	}
	x_1 := New_haxe___Int64_____Int64(high_1, low_1)
	_ = x_1
	var this1_1 *haxe___Int64_____Int64
	_ = this1_1
	this1_1 = x_1
	value_1 := this1_1
	_ = value_1
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringFromLiteral("wrap_sub"), hxrt.StringFromLiteral("=")), haxe___Int64__Int64_Impl__toString(value_1)))
	a := haxe__Int64Helper_parseString(hxrt.StringFromLiteral("1234567890123"))
	_ = a
	b := haxe__Int64Helper_parseString(hxrt.StringFromLiteral("-987654321"))
	_ = b
	high_2 := int(int32((hxrt.Int32Wrap(a.high) + hxrt.Int32Wrap(b.high))))
	_ = high_2
	low_2 := int(int32((hxrt.Int32Wrap(a.low) + hxrt.Int32Wrap(b.low))))
	_ = low_2
	if haxe___Int32__Int32_Impl__ucompare(low_2, a.low) < 0 {
		hx_post_3 := high_2
		high_2 = int(int32((high_2 + 1)))
		ret_2 := hx_post_3
		_ = ret_2
		high_2 = high_2
		_ = ret_2
	}
	x_2 := New_haxe___Int64_____Int64(high_2, low_2)
	_ = x_2
	var this1_2 *haxe___Int64_____Int64
	_ = this1_2
	this1_2 = x_2
	value_2 := this1_2
	_ = value_2
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringFromLiteral("sum"), hxrt.StringFromLiteral("=")), haxe___Int64__Int64_Impl__toString(value_2)))
	high_3 := int(int32((hxrt.Int32Wrap(a.high) - hxrt.Int32Wrap(b.high))))
	_ = high_3
	low_3 := int(int32((hxrt.Int32Wrap(a.low) - hxrt.Int32Wrap(b.low))))
	_ = low_3
	if haxe___Int32__Int32_Impl__ucompare(a.low, b.low) < 0 {
		hx_post_4 := high_3
		high_3 = int(int32((high_3 - 1)))
		ret_3 := hx_post_4
		_ = ret_3
		high_3 = high_3
		_ = ret_3
	}
	x_3 := New_haxe___Int64_____Int64(high_3, low_3)
	_ = x_3
	var this1_3 *haxe___Int64_____Int64
	_ = this1_3
	this1_3 = x_3
	value_3 := this1_3
	_ = value_3
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringFromLiteral("diff"), hxrt.StringFromLiteral("=")), haxe___Int64__Int64_Impl__toString(value_3)))
	var a_low int
	_ = a_low
	var a_high int
	_ = a_high
	a_high = 0
	a_low = 30000
	var b_low_2 int
	_ = b_low_2
	var b_high_2 int
	_ = b_high_2
	b_high_2 = 0
	b_low_2 = 30000
	mask := 65535
	_ = mask
	al := int(int32((hxrt.Int32Wrap(a_low) & hxrt.Int32Wrap(mask))))
	_ = al
	ah := int(int32(int32((uint32(hxrt.Int32Wrap(a_low)) >> uint(16)))))
	_ = ah
	bl := int(int32((hxrt.Int32Wrap(b_low_2) & hxrt.Int32Wrap(mask))))
	_ = bl
	bh := int(int32(int32((uint32(hxrt.Int32Wrap(b_low_2)) >> uint(16)))))
	_ = bh
	p00 := int(int32((hxrt.Int32Wrap(al) * hxrt.Int32Wrap(bl))))
	_ = p00
	p10 := int(int32((hxrt.Int32Wrap(ah) * hxrt.Int32Wrap(bl))))
	_ = p10
	p01 := int(int32((hxrt.Int32Wrap(al) * hxrt.Int32Wrap(bh))))
	_ = p01
	p11 := int(int32((hxrt.Int32Wrap(ah) * hxrt.Int32Wrap(bh))))
	_ = p11
	low_4 := p00
	_ = low_4
	high_4 := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(p11) + hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(p01)) >> uint(16)))))))))) + hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(p10)) >> uint(16)))))))))
	_ = high_4
	p01 = int(int32((hxrt.Int32Wrap(p01) << uint(16))))
	low_4 = int(int32((hxrt.Int32Wrap(low_4) + hxrt.Int32Wrap(p01))))
	if haxe___Int32__Int32_Impl__ucompare(low_4, p01) < 0 {
		hx_post_5 := high_4
		high_4 = int(int32((high_4 + 1)))
		ret_4 := hx_post_5
		_ = ret_4
		high_4 = high_4
		_ = ret_4
	}
	p10 = int(int32((hxrt.Int32Wrap(p10) << uint(16))))
	low_4 = int(int32((hxrt.Int32Wrap(low_4) + hxrt.Int32Wrap(p10))))
	if haxe___Int32__Int32_Impl__ucompare(low_4, p10) < 0 {
		hx_post_6 := high_4
		high_4 = int(int32((high_4 + 1)))
		ret_5 := hx_post_6
		_ = ret_5
		high_4 = high_4
		_ = ret_5
	}
	high_4 = int(int32((hxrt.Int32Wrap(high_4) + hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(a_low) * hxrt.Int32Wrap(b_high_2))))) + hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(a_high) * hxrt.Int32Wrap(b_low_2))))))))))))
	x_4 := New_haxe___Int64_____Int64(high_4, low_4)
	_ = x_4
	var this1_4 *haxe___Int64_____Int64
	_ = this1_4
	this1_4 = x_4
	value_4 := this1_4
	_ = value_4
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringFromLiteral("mul"), hxrt.StringFromLiteral("=")), haxe___Int64__Int64_Impl__toString(value_4)))
	positive := haxe___Int64__Int64_Impl__divMod(haxe__Int64Helper_parseString(hxrt.StringFromLiteral("123456789")), func() *haxe___Int64_____Int64 {
		x_5 := New_haxe___Int64_____Int64(0, 97)
		_ = x_5
		var this1_5 *haxe___Int64_____Int64
		_ = this1_5
		this1_5 = x_5
		return this1_5
	}())
	_ = positive
	value_5 := func(hx_obj_7 map[string]any) *haxe___Int64_____Int64 {
		hx_field_8 := hx_obj_7["quotient"]
		if hx_field_8 == nil {
			var hx_zero_9 *haxe___Int64_____Int64
			return hx_zero_9
		}
		return hx_field_8.(*haxe___Int64_____Int64)
	}(positive)
	_ = value_5
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringFromLiteral("div_q"), hxrt.StringFromLiteral("=")), haxe___Int64__Int64_Impl__toString(value_5)))
	value_6 := func(hx_obj_10 map[string]any) *haxe___Int64_____Int64 {
		hx_field_11 := hx_obj_10["modulus"]
		if hx_field_11 == nil {
			var hx_zero_12 *haxe___Int64_____Int64
			return hx_zero_12
		}
		return hx_field_11.(*haxe___Int64_____Int64)
	}(positive)
	_ = value_6
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringFromLiteral("div_r"), hxrt.StringFromLiteral("=")), haxe___Int64__Int64_Impl__toString(value_6)))
	negative := haxe___Int64__Int64_Impl__divMod(haxe__Int64Helper_parseString(hxrt.StringFromLiteral("-123456789")), func() *haxe___Int64_____Int64 {
		x_6 := New_haxe___Int64_____Int64(0, 97)
		_ = x_6
		var this1_6 *haxe___Int64_____Int64
		_ = this1_6
		this1_6 = x_6
		return this1_6
	}())
	_ = negative
	value_7 := func(hx_obj_13 map[string]any) *haxe___Int64_____Int64 {
		hx_field_14 := hx_obj_13["quotient"]
		if hx_field_14 == nil {
			var hx_zero_15 *haxe___Int64_____Int64
			return hx_zero_15
		}
		return hx_field_14.(*haxe___Int64_____Int64)
	}(negative)
	_ = value_7
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringFromLiteral("div_neg_q"), hxrt.StringFromLiteral("=")), haxe___Int64__Int64_Impl__toString(value_7)))
	value_8 := func(hx_obj_16 map[string]any) *haxe___Int64_____Int64 {
		hx_field_17 := hx_obj_16["modulus"]
		if hx_field_17 == nil {
			var hx_zero_18 *haxe___Int64_____Int64
			return hx_zero_18
		}
		return hx_field_17.(*haxe___Int64_____Int64)
	}(negative)
	_ = value_8
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringFromLiteral("div_neg_r"), hxrt.StringFromLiteral("=")), haxe___Int64__Int64_Impl__toString(value_8)))
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("cmp="), func() int {
		var a_low_1 int
		_ = a_low_1
		var a_high_1 int
		_ = a_high_1
		a_high_1 = -1
		a_low_1 = -1
		var b_low_3 int
		_ = b_low_3
		var b_high_3 int
		_ = b_high_3
		b_high_3 = 0
		b_low_3 = 1
		v := int(int32((hxrt.Int32Wrap(a_high_1) - hxrt.Int32Wrap(b_high_3))))
		_ = v
		var hx_if_19 int
		if v != 0 {
			hx_if_19 = v
		} else {
			hx_if_19 = haxe___Int32__Int32_Impl__ucompare(a_low_1, b_low_3)
		}
		v = hx_if_19
		var hx_if_22 int
		if a_high_1 < 0 {
			var hx_if_20 int
			if b_high_3 < 0 {
				hx_if_20 = v
			} else {
				hx_if_20 = -1
			}
			hx_if_22 = hx_if_20
		} else {
			var hx_if_21 int
			if b_high_3 >= 0 {
				hx_if_21 = v
			} else {
				hx_if_21 = 1
			}
			hx_if_22 = hx_if_21
		}
		return hx_if_22
	}()))
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("ucmp="), func() int {
		var a_low_2 int
		_ = a_low_2
		var a_high_2 int
		_ = a_high_2
		a_high_2 = -1
		a_low_2 = -1
		var b_low_4 int
		_ = b_low_4
		var b_high_4 int
		_ = b_high_4
		b_high_4 = 0
		b_low_4 = 1
		v_1 := haxe___Int32__Int32_Impl__ucompare(a_high_2, b_high_4)
		_ = v_1
		var hx_if_23 int
		if v_1 != 0 {
			hx_if_23 = v_1
		} else {
			hx_if_23 = haxe___Int32__Int32_Impl__ucompare(a_low_2, b_low_4)
		}
		return hx_if_23
	}()))
	var a_low_3 int
	_ = a_low_3
	var a_high_3 int
	_ = a_high_3
	a_high_3 = 0
	a_low_3 = 1
	b_1 := 40
	_ = b_1
	b_1 = int(int32((hxrt.Int32Wrap(b_1) & hxrt.Int32Wrap(63))))
	var hx_if_25 *haxe___Int64_____Int64
	if b_1 == 0 {
		high_5 := a_high_3
		_ = high_5
		low_5 := a_low_3
		_ = low_5
		x_7 := New_haxe___Int64_____Int64(high_5, low_5)
		_ = x_7
		var this1_7 *haxe___Int64_____Int64
		_ = this1_7
		this1_7 = x_7
		hx_if_25 = this1_7
	} else {
		var hx_if_24 *haxe___Int64_____Int64
		if b_1 < 32 {
			high_6 := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(a_high_3) << uint(b_1))))) | hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(a_low_3)) >> uint(int(int32((hxrt.Int32Wrap(32) - hxrt.Int32Wrap(b_1)))))))))))))
			_ = high_6
			low_6 := int(int32((hxrt.Int32Wrap(a_low_3) << uint(b_1))))
			_ = low_6
			x_8 := New_haxe___Int64_____Int64(high_6, low_6)
			_ = x_8
			var this1_8 *haxe___Int64_____Int64
			_ = this1_8
			this1_8 = x_8
			hx_if_24 = this1_8
		} else {
			high_7 := int(int32((hxrt.Int32Wrap(a_low_3) << uint(int(int32((hxrt.Int32Wrap(b_1) - hxrt.Int32Wrap(32))))))))
			_ = high_7
			x_9 := New_haxe___Int64_____Int64(high_7, 0)
			_ = x_9
			var this1_9 *haxe___Int64_____Int64
			_ = this1_9
			this1_9 = x_9
			hx_if_24 = this1_9
		}
		hx_if_25 = hx_if_24
	}
	value_9 := hx_if_25
	_ = value_9
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringFromLiteral("shl"), hxrt.StringFromLiteral("=")), haxe___Int64__Int64_Impl__toString(value_9)))
	a_1 := haxe__Int64Helper_parseString(hxrt.StringFromLiteral("-8"))
	_ = a_1
	b_2 := 1
	_ = b_2
	b_2 = int(int32((hxrt.Int32Wrap(b_2) & hxrt.Int32Wrap(63))))
	var hx_if_27 *haxe___Int64_____Int64
	if b_2 == 0 {
		high_8 := a_1.high
		_ = high_8
		low_7 := a_1.low
		_ = low_7
		x_10 := New_haxe___Int64_____Int64(high_8, low_7)
		_ = x_10
		var this1_10 *haxe___Int64_____Int64
		_ = this1_10
		this1_10 = x_10
		hx_if_27 = this1_10
	} else {
		var hx_if_26 *haxe___Int64_____Int64
		if b_2 < 32 {
			high_9 := int(int32((hxrt.Int32Wrap(a_1.high) >> uint(b_2))))
			_ = high_9
			low_8 := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(a_1.high) << uint(int(int32((hxrt.Int32Wrap(32) - hxrt.Int32Wrap(b_2))))))))) | hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(a_1.low)) >> uint(b_2)))))))))
			_ = low_8
			x_11 := New_haxe___Int64_____Int64(high_9, low_8)
			_ = x_11
			var this1_11 *haxe___Int64_____Int64
			_ = this1_11
			this1_11 = x_11
			hx_if_26 = this1_11
		} else {
			high_10 := int(int32((hxrt.Int32Wrap(a_1.high) >> uint(31))))
			_ = high_10
			low_9 := int(int32((hxrt.Int32Wrap(a_1.high) >> uint(int(int32((hxrt.Int32Wrap(b_2) - hxrt.Int32Wrap(32))))))))
			_ = low_9
			x_12 := New_haxe___Int64_____Int64(high_10, low_9)
			_ = x_12
			var this1_12 *haxe___Int64_____Int64
			_ = this1_12
			this1_12 = x_12
			hx_if_26 = this1_12
		}
		hx_if_27 = hx_if_26
	}
	value_10 := hx_if_27
	_ = value_10
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringFromLiteral("shr"), hxrt.StringFromLiteral("=")), haxe___Int64__Int64_Impl__toString(value_10)))
	a_2 := haxe__Int64Helper_parseString(hxrt.StringFromLiteral("-1"))
	_ = a_2
	b_3 := 1
	_ = b_3
	b_3 = int(int32((hxrt.Int32Wrap(b_3) & hxrt.Int32Wrap(63))))
	var hx_if_29 *haxe___Int64_____Int64
	if b_3 == 0 {
		high_11 := a_2.high
		_ = high_11
		low_10 := a_2.low
		_ = low_10
		x_13 := New_haxe___Int64_____Int64(high_11, low_10)
		_ = x_13
		var this1_13 *haxe___Int64_____Int64
		_ = this1_13
		this1_13 = x_13
		hx_if_29 = this1_13
	} else {
		var hx_if_28 *haxe___Int64_____Int64
		if b_3 < 32 {
			high_12 := int(int32(int32((uint32(hxrt.Int32Wrap(a_2.high)) >> uint(b_3)))))
			_ = high_12
			low_11 := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(a_2.high) << uint(int(int32((hxrt.Int32Wrap(32) - hxrt.Int32Wrap(b_3))))))))) | hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(a_2.low)) >> uint(b_3)))))))))
			_ = low_11
			x_14 := New_haxe___Int64_____Int64(high_12, low_11)
			_ = x_14
			var this1_14 *haxe___Int64_____Int64
			_ = this1_14
			this1_14 = x_14
			hx_if_28 = this1_14
		} else {
			low_12 := int(int32(int32((uint32(hxrt.Int32Wrap(a_2.high)) >> uint(int(int32((hxrt.Int32Wrap(b_3) - hxrt.Int32Wrap(32)))))))))
			_ = low_12
			x_15 := New_haxe___Int64_____Int64(0, low_12)
			_ = x_15
			var this1_15 *haxe___Int64_____Int64
			_ = this1_15
			this1_15 = x_15
			hx_if_28 = this1_15
		}
		hx_if_29 = hx_if_28
	}
	value_11 := hx_if_29
	_ = value_11
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringFromLiteral("ushr"), hxrt.StringFromLiteral("=")), haxe___Int64__Int64_Impl__toString(value_11)))
	value_12 := haxe__Int64Helper_fromFloat(9007199254740991.0)
	_ = value_12
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringFromLiteral("from_float"), hxrt.StringFromLiteral("=")), haxe___Int64__Int64_Impl__toString(value_12)))
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("to_int_ok="), func() int {
		var x_low int
		_ = x_low
		var x_high int
		_ = x_high
		x_high = 0
		x_low = 2147483647
		if x_high != int(int32((hxrt.Int32Wrap(x_low) >> uint(31)))) {
			hxrt.Throw(hxrt.StringFromLiteral("Overflow"))
		}
		return x_low
	}()))
	hxrt.TryCatch(func() {
		x_16 := haxe__Int64Helper_parseString(hxrt.StringFromLiteral("2147483648"))
		_ = x_16
		if x_16.high != int(int32((hxrt.Int32Wrap(x_16.low) >> uint(31)))) {
			hxrt.Throw(hxrt.StringFromLiteral("Overflow"))
		}
		_ = x_16.low
		hxrt.Println(hxrt.StringFromLiteral("to_int_overflow=missing"))
	}, func(hx_caught_30 any) {
		e := hx_caught_30
		_ = e
		hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("to_int_overflow="), hxrt.StdString(e)))
	})
	var round_low int
	_ = round_low
	var round_high int
	_ = round_high
	round_high = 2147483647
	round_low = -12345
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("round_high="), round_high))
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("round_low="), round_low))
}

func haxe__Int64Helper_fromFloat(f float64) *haxe___Int64_____Int64 {
	if Math_isNaN(f) || !Math_isFinite(f) {
		hxrt.Throw(hxrt.StringFromLiteral("Number is NaN or Infinite"))
	}
	noFractions := (f - hxrt.FloatMod(f, float64(1)))
	_ = noFractions
	if noFractions > 9007199254740991 {
		hxrt.Throw(hxrt.StringFromLiteral("Conversion overflow"))
	}
	if noFractions < -9007199254740991 {
		hxrt.Throw(hxrt.StringFromLiteral("Conversion underflow"))
	}
	x := New_haxe___Int64_____Int64(0, 0)
	_ = x
	var this1 *haxe___Int64_____Int64
	_ = this1
	this1 = x
	result := this1
	_ = result
	neg := (noFractions < 0)
	_ = neg
	var hx_if_32 float64
	if neg {
		hx_if_32 = -noFractions
	} else {
		hx_if_32 = noFractions
	}
	rest := hx_if_32
	_ = rest
	i := 0
	_ = i
	for rest >= 1 {
		curr := hxrt.FloatMod(rest, float64(2))
		_ = curr
		rest = (rest / float64(2))
		if curr >= 1 {
			var a_low int
			_ = a_low
			var a_high int
			_ = a_high
			a_high = 0
			a_low = 1
			b_1 := i
			_ = b_1
			b_1 = int(int32((hxrt.Int32Wrap(b_1) & hxrt.Int32Wrap(63))))
			var hx_if_34 *haxe___Int64_____Int64
			if b_1 == 0 {
				high := a_high
				_ = high
				low := a_low
				_ = low
				x_1 := New_haxe___Int64_____Int64(high, low)
				_ = x_1
				var this1_1 *haxe___Int64_____Int64
				_ = this1_1
				this1_1 = x_1
				hx_if_34 = this1_1
			} else {
				var hx_if_33 *haxe___Int64_____Int64
				if b_1 < 32 {
					high_1 := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(a_high) << uint(b_1))))) | hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(a_low)) >> uint(int(int32((hxrt.Int32Wrap(32) - hxrt.Int32Wrap(b_1)))))))))))))
					_ = high_1
					low_1 := int(int32((hxrt.Int32Wrap(a_low) << uint(b_1))))
					_ = low_1
					x_2 := New_haxe___Int64_____Int64(high_1, low_1)
					_ = x_2
					var this1_2 *haxe___Int64_____Int64
					_ = this1_2
					this1_2 = x_2
					hx_if_33 = this1_2
				} else {
					high_2 := int(int32((hxrt.Int32Wrap(a_low) << uint(int(int32((hxrt.Int32Wrap(b_1) - hxrt.Int32Wrap(32))))))))
					_ = high_2
					x_3 := New_haxe___Int64_____Int64(high_2, 0)
					_ = x_3
					var this1_3 *haxe___Int64_____Int64
					_ = this1_3
					this1_3 = x_3
					hx_if_33 = this1_3
				}
				hx_if_34 = hx_if_33
			}
			b := hx_if_34
			_ = b
			high_3 := int(int32((hxrt.Int32Wrap(result.high) + hxrt.Int32Wrap(b.high))))
			_ = high_3
			low_2 := int(int32((hxrt.Int32Wrap(result.low) + hxrt.Int32Wrap(b.low))))
			_ = low_2
			if haxe___Int32__Int32_Impl__ucompare(low_2, result.low) < 0 {
				hx_post_35 := high_3
				high_3 = int(int32((high_3 + 1)))
				ret := hx_post_35
				_ = ret
				high_3 = high_3
				_ = ret
			}
			x_4 := New_haxe___Int64_____Int64(high_3, low_2)
			_ = x_4
			var this1_4 *haxe___Int64_____Int64
			_ = this1_4
			this1_4 = x_4
			result = this1_4
		}
		i = int(int32((i + 1)))
	}
	if neg {
		high_4 := int(int32(^int32(result.high)))
		_ = high_4
		low_3 := int(int32((hxrt.Int32Wrap(int(int32(^int32(result.low)))) + hxrt.Int32Wrap(1))))
		_ = low_3
		if low_3 == 0 {
			hx_post_36 := high_4
			high_4 = int(int32((high_4 + 1)))
			ret_1 := hx_post_36
			_ = ret_1
			high_4 = high_4
			_ = ret_1
		}
		x_5 := New_haxe___Int64_____Int64(high_4, low_3)
		_ = x_5
		var this1_5 *haxe___Int64_____Int64
		_ = this1_5
		this1_5 = x_5
		result = this1_5
	}
	return result
}

func haxe__Int64Helper_parseString(sParam *string) *haxe___Int64_____Int64 {
	var base_low int
	_ = base_low
	var base_high int
	_ = base_high
	base_high = 0
	base_low = 10
	x := New_haxe___Int64_____Int64(0, 0)
	_ = x
	var this1 *haxe___Int64_____Int64
	_ = this1
	this1 = x
	current := this1
	_ = current
	x_1 := New_haxe___Int64_____Int64(0, 1)
	_ = x_1
	var this1_1 *haxe___Int64_____Int64
	_ = this1_1
	this1_1 = x_1
	multiplier := this1_1
	_ = multiplier
	sIsNegative := false
	_ = sIsNegative
	s := StringTools_trim(sParam)
	_ = s
	if hxrt.StringEqualAny(hxrt.StringCharAt(s, 0), hxrt.StringFromLiteral("-")) {
		sIsNegative = true
		s = hxrt.StringSubstring(s, 1, hxrt.StringLength(s))
	}
	len := hxrt.StringLength(s)
	_ = len
	_g := 0
	_ = _g
	_g1 := len
	_ = _g1
	for _g < _g1 {
		hx_post_37 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_37
		_ = i
		digitInt := int(int32((hxrt.Int32Wrap(hxrt.StringCharCodeAt(s, int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(len)-hxrt.Int32Wrap(1)))))-hxrt.Int32Wrap(i)))))) - hxrt.Int32Wrap(48))))
		_ = digitInt
		if (digitInt < 0) || (digitInt > 9) {
			hxrt.Throw(hxrt.StringFromLiteral("NumberFormatError"))
		}
		if digitInt != 0 {
			var digit_low int
			_ = digit_low
			var digit_high int
			_ = digit_high
			digit_high = int(int32((hxrt.Int32Wrap(digitInt) >> uint(31))))
			digit_low = digitInt
			if sIsNegative {
				var b_low int
				_ = b_low
				var b_high int
				_ = b_high
				mask := 65535
				_ = mask
				al := int(int32((hxrt.Int32Wrap(multiplier.low) & hxrt.Int32Wrap(mask))))
				_ = al
				ah := int(int32(int32((uint32(hxrt.Int32Wrap(multiplier.low)) >> uint(16)))))
				_ = ah
				bl := int(int32((hxrt.Int32Wrap(digit_low) & hxrt.Int32Wrap(mask))))
				_ = bl
				bh := int(int32(int32((uint32(hxrt.Int32Wrap(digit_low)) >> uint(16)))))
				_ = bh
				p00 := int(int32((hxrt.Int32Wrap(al) * hxrt.Int32Wrap(bl))))
				_ = p00
				p10 := int(int32((hxrt.Int32Wrap(ah) * hxrt.Int32Wrap(bl))))
				_ = p10
				p01 := int(int32((hxrt.Int32Wrap(al) * hxrt.Int32Wrap(bh))))
				_ = p01
				p11 := int(int32((hxrt.Int32Wrap(ah) * hxrt.Int32Wrap(bh))))
				_ = p11
				low := p00
				_ = low
				high := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(p11) + hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(p01)) >> uint(16)))))))))) + hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(p10)) >> uint(16)))))))))
				_ = high
				p01 = int(int32((hxrt.Int32Wrap(p01) << uint(16))))
				low = int(int32((hxrt.Int32Wrap(low) + hxrt.Int32Wrap(p01))))
				if haxe___Int32__Int32_Impl__ucompare(low, p01) < 0 {
					hx_post_38 := high
					high = int(int32((high + 1)))
					ret := hx_post_38
					_ = ret
					high = high
					_ = ret
				}
				p10 = int(int32((hxrt.Int32Wrap(p10) << uint(16))))
				low = int(int32((hxrt.Int32Wrap(low) + hxrt.Int32Wrap(p10))))
				if haxe___Int32__Int32_Impl__ucompare(low, p10) < 0 {
					hx_post_39 := high
					high = int(int32((high + 1)))
					ret_1 := hx_post_39
					_ = ret_1
					high = high
					_ = ret_1
				}
				high = int(int32((hxrt.Int32Wrap(high) + hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(multiplier.low) * hxrt.Int32Wrap(digit_high))))) + hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(multiplier.high) * hxrt.Int32Wrap(digit_low))))))))))))
				b_high = high
				b_low = low
				high_1 := int(int32((hxrt.Int32Wrap(current.high) - hxrt.Int32Wrap(b_high))))
				_ = high_1
				low_1 := int(int32((hxrt.Int32Wrap(current.low) - hxrt.Int32Wrap(b_low))))
				_ = low_1
				if haxe___Int32__Int32_Impl__ucompare(current.low, b_low) < 0 {
					hx_post_40 := high_1
					high_1 = int(int32((high_1 - 1)))
					ret_2 := hx_post_40
					_ = ret_2
					high_1 = high_1
					_ = ret_2
				}
				x_2 := New_haxe___Int64_____Int64(high_1, low_1)
				_ = x_2
				var this1_2 *haxe___Int64_____Int64
				_ = this1_2
				this1_2 = x_2
				current = this1_2
				if !(current.high < 0) {
					hxrt.Throw(hxrt.StringFromLiteral("NumberFormatError: Underflow"))
				}
			} else {
				var b_low_1 int
				_ = b_low_1
				var b_high_1 int
				_ = b_high_1
				mask_1 := 65535
				_ = mask_1
				al_1 := int(int32((hxrt.Int32Wrap(multiplier.low) & hxrt.Int32Wrap(mask_1))))
				_ = al_1
				ah_1 := int(int32(int32((uint32(hxrt.Int32Wrap(multiplier.low)) >> uint(16)))))
				_ = ah_1
				bl_1 := int(int32((hxrt.Int32Wrap(digit_low) & hxrt.Int32Wrap(mask_1))))
				_ = bl_1
				bh_1 := int(int32(int32((uint32(hxrt.Int32Wrap(digit_low)) >> uint(16)))))
				_ = bh_1
				p00_1 := int(int32((hxrt.Int32Wrap(al_1) * hxrt.Int32Wrap(bl_1))))
				_ = p00_1
				p10_1 := int(int32((hxrt.Int32Wrap(ah_1) * hxrt.Int32Wrap(bl_1))))
				_ = p10_1
				p01_1 := int(int32((hxrt.Int32Wrap(al_1) * hxrt.Int32Wrap(bh_1))))
				_ = p01_1
				p11_1 := int(int32((hxrt.Int32Wrap(ah_1) * hxrt.Int32Wrap(bh_1))))
				_ = p11_1
				low_2 := p00_1
				_ = low_2
				high_2 := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(p11_1) + hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(p01_1)) >> uint(16)))))))))) + hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(p10_1)) >> uint(16)))))))))
				_ = high_2
				p01_1 = int(int32((hxrt.Int32Wrap(p01_1) << uint(16))))
				low_2 = int(int32((hxrt.Int32Wrap(low_2) + hxrt.Int32Wrap(p01_1))))
				if haxe___Int32__Int32_Impl__ucompare(low_2, p01_1) < 0 {
					hx_post_41 := high_2
					high_2 = int(int32((high_2 + 1)))
					ret_3 := hx_post_41
					_ = ret_3
					high_2 = high_2
					_ = ret_3
				}
				p10_1 = int(int32((hxrt.Int32Wrap(p10_1) << uint(16))))
				low_2 = int(int32((hxrt.Int32Wrap(low_2) + hxrt.Int32Wrap(p10_1))))
				if haxe___Int32__Int32_Impl__ucompare(low_2, p10_1) < 0 {
					hx_post_42 := high_2
					high_2 = int(int32((high_2 + 1)))
					ret_4 := hx_post_42
					_ = ret_4
					high_2 = high_2
					_ = ret_4
				}
				high_2 = int(int32((hxrt.Int32Wrap(high_2) + hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(multiplier.low) * hxrt.Int32Wrap(digit_high))))) + hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(multiplier.high) * hxrt.Int32Wrap(digit_low))))))))))))
				b_high_1 = high_2
				b_low_1 = low_2
				high_3 := int(int32((hxrt.Int32Wrap(current.high) + hxrt.Int32Wrap(b_high_1))))
				_ = high_3
				low_3 := int(int32((hxrt.Int32Wrap(current.low) + hxrt.Int32Wrap(b_low_1))))
				_ = low_3
				if haxe___Int32__Int32_Impl__ucompare(low_3, current.low) < 0 {
					hx_post_43 := high_3
					high_3 = int(int32((high_3 + 1)))
					ret_5 := hx_post_43
					_ = ret_5
					high_3 = high_3
					_ = ret_5
				}
				x_3 := New_haxe___Int64_____Int64(high_3, low_3)
				_ = x_3
				var this1_3 *haxe___Int64_____Int64
				_ = this1_3
				this1_3 = x_3
				current = this1_3
				if current.high < 0 {
					hxrt.Throw(hxrt.StringFromLiteral("NumberFormatError: Overflow"))
				}
			}
		}
		mask_2 := 65535
		_ = mask_2
		al_2 := int(int32((hxrt.Int32Wrap(multiplier.low) & hxrt.Int32Wrap(mask_2))))
		_ = al_2
		ah_2 := int(int32(int32((uint32(hxrt.Int32Wrap(multiplier.low)) >> uint(16)))))
		_ = ah_2
		bl_2 := int(int32((hxrt.Int32Wrap(base_low) & hxrt.Int32Wrap(mask_2))))
		_ = bl_2
		bh_2 := int(int32(int32((uint32(hxrt.Int32Wrap(base_low)) >> uint(16)))))
		_ = bh_2
		p00_2 := int(int32((hxrt.Int32Wrap(al_2) * hxrt.Int32Wrap(bl_2))))
		_ = p00_2
		p10_2 := int(int32((hxrt.Int32Wrap(ah_2) * hxrt.Int32Wrap(bl_2))))
		_ = p10_2
		p01_2 := int(int32((hxrt.Int32Wrap(al_2) * hxrt.Int32Wrap(bh_2))))
		_ = p01_2
		p11_2 := int(int32((hxrt.Int32Wrap(ah_2) * hxrt.Int32Wrap(bh_2))))
		_ = p11_2
		low_4 := p00_2
		_ = low_4
		high_4 := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(p11_2) + hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(p01_2)) >> uint(16)))))))))) + hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(p10_2)) >> uint(16)))))))))
		_ = high_4
		p01_2 = int(int32((hxrt.Int32Wrap(p01_2) << uint(16))))
		low_4 = int(int32((hxrt.Int32Wrap(low_4) + hxrt.Int32Wrap(p01_2))))
		if haxe___Int32__Int32_Impl__ucompare(low_4, p01_2) < 0 {
			hx_post_44 := high_4
			high_4 = int(int32((high_4 + 1)))
			ret_6 := hx_post_44
			_ = ret_6
			high_4 = high_4
			_ = ret_6
		}
		p10_2 = int(int32((hxrt.Int32Wrap(p10_2) << uint(16))))
		low_4 = int(int32((hxrt.Int32Wrap(low_4) + hxrt.Int32Wrap(p10_2))))
		if haxe___Int32__Int32_Impl__ucompare(low_4, p10_2) < 0 {
			hx_post_45 := high_4
			high_4 = int(int32((high_4 + 1)))
			ret_7 := hx_post_45
			_ = ret_7
			high_4 = high_4
			_ = ret_7
		}
		high_4 = int(int32((hxrt.Int32Wrap(high_4) + hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(multiplier.low) * hxrt.Int32Wrap(base_high))))) + hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(multiplier.high) * hxrt.Int32Wrap(base_low))))))))))))
		x_4 := New_haxe___Int64_____Int64(high_4, low_4)
		_ = x_4
		var this1_4 *haxe___Int64_____Int64
		_ = this1_4
		this1_4 = x_4
		multiplier = this1_4
	}
	return current
}

func haxe___Int32__Int32_Impl__ucompare(a int, b int) int {
	if a < 0 {
		var hx_if_46 int
		if b < 0 {
			hx_if_46 = int(int32((hxrt.Int32Wrap(int(int32(^int32(b)))) - hxrt.Int32Wrap(int(int32(^int32(a)))))))
		} else {
			hx_if_46 = 1
		}
		return hx_if_46
	}
	var hx_if_47 int
	if b < 0 {
		hx_if_47 = -1
	} else {
		hx_if_47 = int(int32((hxrt.Int32Wrap(a) - hxrt.Int32Wrap(b))))
	}
	return hx_if_47
}

func haxe___Int64__Int64_Impl__divMod(dividend *haxe___Int64_____Int64, divisor *haxe___Int64_____Int64) map[string]any {
	if divisor.high == 0 {
		_g := divisor.low
		_ = _g
		switch _g {
		case 0:
			hxrt.Throw(hxrt.StringFromLiteral("divide by zero"))
		case 1:
			hx_obj_48 := map[string]any{}
			high := dividend.high
			_ = high
			low := dividend.low
			_ = low
			x := New_haxe___Int64_____Int64(high, low)
			_ = x
			var this1 *haxe___Int64_____Int64
			_ = this1
			this1 = x
			hx_obj_48["quotient"] = this1
			x_1 := New_haxe___Int64_____Int64(0, 0)
			_ = x_1
			var this1_1 *haxe___Int64_____Int64
			_ = this1_1
			this1_1 = x_1
			hx_obj_48["modulus"] = this1_1
			return hx_obj_48
		}
	}
	divSign := ((dividend.high < 0) != (divisor.high < 0))
	_ = divSign
	var hx_if_50 *haxe___Int64_____Int64
	if dividend.high < 0 {
		high_1 := int(int32(^int32(dividend.high)))
		_ = high_1
		low_1 := int(int32((hxrt.Int32Wrap(int(int32(^int32(dividend.low)))) + hxrt.Int32Wrap(1))))
		_ = low_1
		if low_1 == 0 {
			hx_post_49 := high_1
			high_1 = int(int32((high_1 + 1)))
			ret := hx_post_49
			_ = ret
			high_1 = high_1
			_ = ret
		}
		x_2 := New_haxe___Int64_____Int64(high_1, low_1)
		_ = x_2
		var this1_2 *haxe___Int64_____Int64
		_ = this1_2
		this1_2 = x_2
		hx_if_50 = this1_2
	} else {
		high_2 := dividend.high
		_ = high_2
		low_2 := dividend.low
		_ = low_2
		x_3 := New_haxe___Int64_____Int64(high_2, low_2)
		_ = x_3
		var this1_3 *haxe___Int64_____Int64
		_ = this1_3
		this1_3 = x_3
		hx_if_50 = this1_3
	}
	modulus := hx_if_50
	_ = modulus
	var hx_if_52 *haxe___Int64_____Int64
	if divisor.high < 0 {
		high_3 := int(int32(^int32(divisor.high)))
		_ = high_3
		low_3 := int(int32((hxrt.Int32Wrap(int(int32(^int32(divisor.low)))) + hxrt.Int32Wrap(1))))
		_ = low_3
		if low_3 == 0 {
			hx_post_51 := high_3
			high_3 = int(int32((high_3 + 1)))
			ret_1 := hx_post_51
			_ = ret_1
			high_3 = high_3
			_ = ret_1
		}
		x_4 := New_haxe___Int64_____Int64(high_3, low_3)
		_ = x_4
		var this1_4 *haxe___Int64_____Int64
		_ = this1_4
		this1_4 = x_4
		hx_if_52 = this1_4
	} else {
		hx_if_52 = divisor
	}
	divisor = hx_if_52
	x_5 := New_haxe___Int64_____Int64(0, 0)
	_ = x_5
	var this1_5 *haxe___Int64_____Int64
	_ = this1_5
	this1_5 = x_5
	quotient := this1_5
	_ = quotient
	x_6 := New_haxe___Int64_____Int64(0, 1)
	_ = x_6
	var this1_6 *haxe___Int64_____Int64
	_ = this1_6
	this1_6 = x_6
	mask := this1_6
	_ = mask
	for !(divisor.high < 0) {
		v := haxe___Int32__Int32_Impl__ucompare(divisor.high, modulus.high)
		_ = v
		var hx_if_53 int
		if v != 0 {
			hx_if_53 = v
		} else {
			hx_if_53 = haxe___Int32__Int32_Impl__ucompare(divisor.low, modulus.low)
		}
		cmp := hx_if_53
		_ = cmp
		b := 1
		_ = b
		b = int(int32((hxrt.Int32Wrap(b) & hxrt.Int32Wrap(63))))
		var hx_if_55 *haxe___Int64_____Int64
		if b == 0 {
			high_4 := divisor.high
			_ = high_4
			low_4 := divisor.low
			_ = low_4
			x_7 := New_haxe___Int64_____Int64(high_4, low_4)
			_ = x_7
			var this1_7 *haxe___Int64_____Int64
			_ = this1_7
			this1_7 = x_7
			hx_if_55 = this1_7
		} else {
			var hx_if_54 *haxe___Int64_____Int64
			if b < 32 {
				high_5 := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(divisor.high) << uint(b))))) | hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(divisor.low)) >> uint(int(int32((hxrt.Int32Wrap(32) - hxrt.Int32Wrap(b)))))))))))))
				_ = high_5
				low_5 := int(int32((hxrt.Int32Wrap(divisor.low) << uint(b))))
				_ = low_5
				x_8 := New_haxe___Int64_____Int64(high_5, low_5)
				_ = x_8
				var this1_8 *haxe___Int64_____Int64
				_ = this1_8
				this1_8 = x_8
				hx_if_54 = this1_8
			} else {
				high_6 := int(int32((hxrt.Int32Wrap(divisor.low) << uint(int(int32((hxrt.Int32Wrap(b) - hxrt.Int32Wrap(32))))))))
				_ = high_6
				x_9 := New_haxe___Int64_____Int64(high_6, 0)
				_ = x_9
				var this1_9 *haxe___Int64_____Int64
				_ = this1_9
				this1_9 = x_9
				hx_if_54 = this1_9
			}
			hx_if_55 = hx_if_54
		}
		divisor = hx_if_55
		b_1 := 1
		_ = b_1
		b_1 = int(int32((hxrt.Int32Wrap(b_1) & hxrt.Int32Wrap(63))))
		var hx_if_57 *haxe___Int64_____Int64
		if b_1 == 0 {
			high_7 := mask.high
			_ = high_7
			low_6 := mask.low
			_ = low_6
			x_10 := New_haxe___Int64_____Int64(high_7, low_6)
			_ = x_10
			var this1_10 *haxe___Int64_____Int64
			_ = this1_10
			this1_10 = x_10
			hx_if_57 = this1_10
		} else {
			var hx_if_56 *haxe___Int64_____Int64
			if b_1 < 32 {
				high_8 := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(mask.high) << uint(b_1))))) | hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(mask.low)) >> uint(int(int32((hxrt.Int32Wrap(32) - hxrt.Int32Wrap(b_1)))))))))))))
				_ = high_8
				low_7 := int(int32((hxrt.Int32Wrap(mask.low) << uint(b_1))))
				_ = low_7
				x_11 := New_haxe___Int64_____Int64(high_8, low_7)
				_ = x_11
				var this1_11 *haxe___Int64_____Int64
				_ = this1_11
				this1_11 = x_11
				hx_if_56 = this1_11
			} else {
				high_9 := int(int32((hxrt.Int32Wrap(mask.low) << uint(int(int32((hxrt.Int32Wrap(b_1) - hxrt.Int32Wrap(32))))))))
				_ = high_9
				x_12 := New_haxe___Int64_____Int64(high_9, 0)
				_ = x_12
				var this1_12 *haxe___Int64_____Int64
				_ = this1_12
				this1_12 = x_12
				hx_if_56 = this1_12
			}
			hx_if_57 = hx_if_56
		}
		mask = hx_if_57
		if cmp >= 0 {
			break
		}
	}
	for func() bool {
		var b_low int
		_ = b_low
		var b_high int
		_ = b_high
		b_high = 0
		b_low = 0
		return ((mask.high != b_high) || (mask.low != b_low))
	}() {
		if func() int {
			v_1 := haxe___Int32__Int32_Impl__ucompare(modulus.high, divisor.high)
			_ = v_1
			var hx_if_58 int
			if v_1 != 0 {
				hx_if_58 = v_1
			} else {
				hx_if_58 = haxe___Int32__Int32_Impl__ucompare(modulus.low, divisor.low)
			}
			return hx_if_58
		}() >= 0 {
			high_10 := int(int32((hxrt.Int32Wrap(quotient.high) | hxrt.Int32Wrap(mask.high))))
			_ = high_10
			low_8 := int(int32((hxrt.Int32Wrap(quotient.low) | hxrt.Int32Wrap(mask.low))))
			_ = low_8
			x_13 := New_haxe___Int64_____Int64(high_10, low_8)
			_ = x_13
			var this1_13 *haxe___Int64_____Int64
			_ = this1_13
			this1_13 = x_13
			quotient = this1_13
			high_11 := int(int32((hxrt.Int32Wrap(modulus.high) - hxrt.Int32Wrap(divisor.high))))
			_ = high_11
			low_9 := int(int32((hxrt.Int32Wrap(modulus.low) - hxrt.Int32Wrap(divisor.low))))
			_ = low_9
			if haxe___Int32__Int32_Impl__ucompare(modulus.low, divisor.low) < 0 {
				hx_post_59 := high_11
				high_11 = int(int32((high_11 - 1)))
				ret_2 := hx_post_59
				_ = ret_2
				high_11 = high_11
				_ = ret_2
			}
			x_14 := New_haxe___Int64_____Int64(high_11, low_9)
			_ = x_14
			var this1_14 *haxe___Int64_____Int64
			_ = this1_14
			this1_14 = x_14
			modulus = this1_14
		}
		b_2 := 1
		_ = b_2
		b_2 = int(int32((hxrt.Int32Wrap(b_2) & hxrt.Int32Wrap(63))))
		var hx_if_61 *haxe___Int64_____Int64
		if b_2 == 0 {
			high_12 := mask.high
			_ = high_12
			low_10 := mask.low
			_ = low_10
			x_15 := New_haxe___Int64_____Int64(high_12, low_10)
			_ = x_15
			var this1_15 *haxe___Int64_____Int64
			_ = this1_15
			this1_15 = x_15
			hx_if_61 = this1_15
		} else {
			var hx_if_60 *haxe___Int64_____Int64
			if b_2 < 32 {
				high_13 := int(int32(int32((uint32(hxrt.Int32Wrap(mask.high)) >> uint(b_2)))))
				_ = high_13
				low_11 := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(mask.high) << uint(int(int32((hxrt.Int32Wrap(32) - hxrt.Int32Wrap(b_2))))))))) | hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(mask.low)) >> uint(b_2)))))))))
				_ = low_11
				x_16 := New_haxe___Int64_____Int64(high_13, low_11)
				_ = x_16
				var this1_16 *haxe___Int64_____Int64
				_ = this1_16
				this1_16 = x_16
				hx_if_60 = this1_16
			} else {
				low_12 := int(int32(int32((uint32(hxrt.Int32Wrap(mask.high)) >> uint(int(int32((hxrt.Int32Wrap(b_2) - hxrt.Int32Wrap(32)))))))))
				_ = low_12
				x_17 := New_haxe___Int64_____Int64(0, low_12)
				_ = x_17
				var this1_17 *haxe___Int64_____Int64
				_ = this1_17
				this1_17 = x_17
				hx_if_60 = this1_17
			}
			hx_if_61 = hx_if_60
		}
		mask = hx_if_61
		b_3 := 1
		_ = b_3
		b_3 = int(int32((hxrt.Int32Wrap(b_3) & hxrt.Int32Wrap(63))))
		var hx_if_63 *haxe___Int64_____Int64
		if b_3 == 0 {
			high_14 := divisor.high
			_ = high_14
			low_13 := divisor.low
			_ = low_13
			x_18 := New_haxe___Int64_____Int64(high_14, low_13)
			_ = x_18
			var this1_18 *haxe___Int64_____Int64
			_ = this1_18
			this1_18 = x_18
			hx_if_63 = this1_18
		} else {
			var hx_if_62 *haxe___Int64_____Int64
			if b_3 < 32 {
				high_15 := int(int32(int32((uint32(hxrt.Int32Wrap(divisor.high)) >> uint(b_3)))))
				_ = high_15
				low_14 := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(divisor.high) << uint(int(int32((hxrt.Int32Wrap(32) - hxrt.Int32Wrap(b_3))))))))) | hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(divisor.low)) >> uint(b_3)))))))))
				_ = low_14
				x_19 := New_haxe___Int64_____Int64(high_15, low_14)
				_ = x_19
				var this1_19 *haxe___Int64_____Int64
				_ = this1_19
				this1_19 = x_19
				hx_if_62 = this1_19
			} else {
				low_15 := int(int32(int32((uint32(hxrt.Int32Wrap(divisor.high)) >> uint(int(int32((hxrt.Int32Wrap(b_3) - hxrt.Int32Wrap(32)))))))))
				_ = low_15
				x_20 := New_haxe___Int64_____Int64(0, low_15)
				_ = x_20
				var this1_20 *haxe___Int64_____Int64
				_ = this1_20
				this1_20 = x_20
				hx_if_62 = this1_20
			}
			hx_if_63 = hx_if_62
		}
		divisor = hx_if_63
	}
	if divSign {
		high_16 := int(int32(^int32(quotient.high)))
		_ = high_16
		low_16 := int(int32((hxrt.Int32Wrap(int(int32(^int32(quotient.low)))) + hxrt.Int32Wrap(1))))
		_ = low_16
		if low_16 == 0 {
			hx_post_64 := high_16
			high_16 = int(int32((high_16 + 1)))
			ret_3 := hx_post_64
			_ = ret_3
			high_16 = high_16
			_ = ret_3
		}
		x_21 := New_haxe___Int64_____Int64(high_16, low_16)
		_ = x_21
		var this1_21 *haxe___Int64_____Int64
		_ = this1_21
		this1_21 = x_21
		quotient = this1_21
	}
	if dividend.high < 0 {
		high_17 := int(int32(^int32(modulus.high)))
		_ = high_17
		low_17 := int(int32((hxrt.Int32Wrap(int(int32(^int32(modulus.low)))) + hxrt.Int32Wrap(1))))
		_ = low_17
		if low_17 == 0 {
			hx_post_65 := high_17
			high_17 = int(int32((high_17 + 1)))
			ret_4 := hx_post_65
			_ = ret_4
			high_17 = high_17
			_ = ret_4
		}
		x_22 := New_haxe___Int64_____Int64(high_17, low_17)
		_ = x_22
		var this1_22 *haxe___Int64_____Int64
		_ = this1_22
		this1_22 = x_22
		modulus = this1_22
	}
	hx_obj_66 := map[string]any{}
	hx_obj_66["quotient"] = quotient
	hx_obj_66["modulus"] = modulus
	return hx_obj_66
}

func haxe___Int64__Int64_Impl__toString(this1 *haxe___Int64_____Int64) *string {
	i := this1
	_ = i
	if func() bool {
		var b_low int
		_ = b_low
		var b_high int
		_ = b_high
		b_high = 0
		b_low = 0
		return ((i.high == b_high) && (i.low == b_low))
	}() {
		return hxrt.StringFromLiteral("0")
	}
	str := hxrt.StringFromLiteral("")
	_ = str
	neg := false
	_ = neg
	if i.high < 0 {
		neg = true
	}
	x := New_haxe___Int64_____Int64(0, 10)
	_ = x
	var this1_1 *haxe___Int64_____Int64
	_ = this1_1
	this1_1 = x
	ten := this1_1
	_ = ten
	for func() bool {
		var b_low_1 int
		_ = b_low_1
		var b_high_1 int
		_ = b_high_1
		b_high_1 = 0
		b_low_1 = 0
		return ((i.high != b_high_1) || (i.low != b_low_1))
	}() {
		r := haxe___Int64__Int64_Impl__divMod(i, ten)
		_ = r
		if func(hx_obj_67 map[string]any) *haxe___Int64_____Int64 {
			hx_field_68 := hx_obj_67["modulus"]
			if hx_field_68 == nil {
				var hx_zero_69 *haxe___Int64_____Int64
				return hx_zero_69
			}
			return hx_field_68.(*haxe___Int64_____Int64)
		}(r).high < 0 {
			str = hxrt.StringConcatAny(func() int {
				var this_low int
				_ = this_low
				var this_high int
				_ = this_high
				x_1 := func(hx_obj_70 map[string]any) *haxe___Int64_____Int64 {
					hx_field_71 := hx_obj_70["modulus"]
					if hx_field_71 == nil {
						var hx_zero_72 *haxe___Int64_____Int64
						return hx_zero_72
					}
					return hx_field_71.(*haxe___Int64_____Int64)
				}(r)
				_ = x_1
				high := int(int32(^int32(x_1.high)))
				_ = high
				low := int(int32((hxrt.Int32Wrap(int(int32(^int32(x_1.low)))) + hxrt.Int32Wrap(1))))
				_ = low
				if low == 0 {
					hx_post_73 := high
					high = int(int32((high + 1)))
					ret := hx_post_73
					_ = ret
					high = high
					_ = ret
				}
				this_high = high
				this_low = low
				return this_low
			}(), str)
			x_2 := func(hx_obj_74 map[string]any) *haxe___Int64_____Int64 {
				hx_field_75 := hx_obj_74["quotient"]
				if hx_field_75 == nil {
					var hx_zero_76 *haxe___Int64_____Int64
					return hx_zero_76
				}
				return hx_field_75.(*haxe___Int64_____Int64)
			}(r)
			_ = x_2
			high_1 := int(int32(^int32(x_2.high)))
			_ = high_1
			low_1 := int(int32((hxrt.Int32Wrap(int(int32(^int32(x_2.low)))) + hxrt.Int32Wrap(1))))
			_ = low_1
			if low_1 == 0 {
				hx_post_77 := high_1
				high_1 = int(int32((high_1 + 1)))
				ret_1 := hx_post_77
				_ = ret_1
				high_1 = high_1
				_ = ret_1
			}
			x_3 := New_haxe___Int64_____Int64(high_1, low_1)
			_ = x_3
			var this1_2 *haxe___Int64_____Int64
			_ = this1_2
			this1_2 = x_3
			i = this1_2
		} else {
			str = hxrt.StringConcatAny(func(hx_obj_78 map[string]any) *haxe___Int64_____Int64 {
				hx_field_79 := hx_obj_78["modulus"]
				if hx_field_79 == nil {
					var hx_zero_80 *haxe___Int64_____Int64
					return hx_zero_80
				}
				return hx_field_79.(*haxe___Int64_____Int64)
			}(r).low, str)
			i = func(hx_obj_81 map[string]any) *haxe___Int64_____Int64 {
				hx_field_82 := hx_obj_81["quotient"]
				if hx_field_82 == nil {
					var hx_zero_83 *haxe___Int64_____Int64
					return hx_zero_83
				}
				return hx_field_82.(*haxe___Int64_____Int64)
			}(r)
		}
	}
	if neg {
		str = hxrt.StringConcatAny(hxrt.StringFromLiteral("-"), str)
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

type haxe__io__Encoding struct {
}

type haxe__io__Input struct {
}

type haxe__io__Output struct {
}

type haxe__io__Bytes struct {
	b             []int
	length        int
	__hx_raw      []byte
	__hx_rawValid bool
}

type haxe__io__BytesBuffer struct {
	b []int
}

func New_haxe__io__Input() *haxe__io__Input {
	return &haxe__io__Input{}
}

func New_haxe__io__Output() *haxe__io__Output {
	return &haxe__io__Output{}
}

func New_haxe__io__Bytes(length int, b []int) *haxe__io__Bytes {
	if b == nil {
		b = make([]int, length)
	}
	return &haxe__io__Bytes{b: b, length: len(b)}
}

func haxe__io__Bytes_alloc(length int) *haxe__io__Bytes {
	return &haxe__io__Bytes{b: make([]int, length), length: length}
}

func haxe__io__Bytes_ofString(value *string, encoding ...*haxe__io__Encoding) *haxe__io__Bytes {
	raw := []byte(*hxrt.StdString(value))
	converted := make([]int, len(raw))
	for i := 0; i < len(raw); i++ {
		converted[i] = int(raw[i])
	}
	return &haxe__io__Bytes{b: converted, length: len(converted), __hx_raw: raw, __hx_rawValid: true}
}

func (self *haxe__io__Bytes) toString() *string {
	if self == nil {
		return hxrt.StringFromLiteral("")
	}
	return hxrt.BytesToString(self.b)
}

func (self *haxe__io__Bytes) get(pos int) int {
	return self.b[pos]
}

func (self *haxe__io__Bytes) set(pos int, value int) {
	self.b[pos] = value
	self.__hx_rawValid = false
}

func New_haxe__io__BytesBuffer() *haxe__io__BytesBuffer {
	return &haxe__io__BytesBuffer{b: []int{}}
}

func (self *haxe__io__BytesBuffer) addByte(value int) {
	self.b = append(self.b, value)
}

func (self *haxe__io__BytesBuffer) add(src *haxe__io__Bytes) {
	if src == nil {
		return
	}
	self.b = append(self.b, src.b...)
}

func (self *haxe__io__BytesBuffer) addString(value *string, encoding ...*haxe__io__Encoding) {
	self.add(haxe__io__Bytes_ofString(value))
}

func (self *haxe__io__BytesBuffer) getBytes() *haxe__io__Bytes {
	copied := hxrt.BytesClone(self.b)
	return &haxe__io__Bytes{b: copied, length: len(copied)}
}

func (self *haxe__io__BytesBuffer) get_length() int {
	return len(self.b)
}

type Std struct {
}

type StringTools struct {
}

func StringTools_trim(value *string) *string {
	return hxrt.StringFromLiteral(strings.TrimSpace(*hxrt.StdString(value)))
}

func StringTools_startsWith(value *string, prefix *string) bool {
	return strings.HasPrefix(*hxrt.StdString(value), *hxrt.StdString(prefix))
}

func StringTools_replace(value *string, sub *string, by *string) *string {
	return hxrt.StringFromLiteral(strings.ReplaceAll(*hxrt.StdString(value), *hxrt.StdString(sub), *hxrt.StdString(by)))
}

type Date struct {
	value time.Time
}

func Date_fromString(source *string) *Date {
	raw := *hxrt.StdString(source)
	parsed, err := time.ParseInLocation("2006-01-02 15:04:05", raw, time.Local)
	if err != nil {
		parsedDateOnly, errDateOnly := time.ParseInLocation("2006-01-02", raw, time.Local)
		if errDateOnly == nil {
			parsed = parsedDateOnly
		} else {
			parsed = time.Unix(0, 0)
		}
	}
	return &Date{value: parsed}
}

func Date_now() *Date {
	return &Date{value: time.Now()}
}

func (self *Date) getFullYear() int {
	return self.value.Year()
}

func (self *Date) getMonth() int {
	return int(self.value.Month()) - 1
}

func (self *Date) getDate() int {
	return self.value.Day()
}

func (self *Date) getHours() int {
	return self.value.Hour()
}

type Math struct {
}

func Math_floor(value float64) int {
	return int(math.Floor(value))
}

func Math_ceil(value float64) int {
	return int(math.Ceil(value))
}

func Math_round(value float64) int {
	return int(math.Floor(value + 0.5))
}

func Math_abs(value float64) float64 {
	return math.Abs(value)
}

func Math_isNaN(value float64) bool {
	return math.IsNaN(value)
}

func Math_isFinite(value float64) bool {
	return !math.IsInf(value, 0)
}

func Math_min(a float64, b float64) float64 {
	return math.Min(a, b)
}

func Math_max(a float64, b float64) float64 {
	return math.Max(a, b)
}

type Type struct {
}

type Reflect struct {
}

func Reflect_compare(a any, b any) int {
	toFloat := func(value any) (float64, bool) {
		switch v := value.(type) {
		case int:
			return float64(v), true
		case int8:
			return float64(v), true
		case int16:
			return float64(v), true
		case int32:
			return float64(v), true
		case int64:
			return float64(v), true
		case uint:
			return float64(v), true
		case uint8:
			return float64(v), true
		case uint16:
			return float64(v), true
		case uint32:
			return float64(v), true
		case uint64:
			return float64(v), true
		case float32:
			return float64(v), true
		case float64:
			return v, true
		default:
			return 0, false
		}
	}
	if af, ok := toFloat(a); ok {
		if bf, okB := toFloat(b); okB {
			if af < bf {
				return -1
			}
			if af > bf {
				return 1
			}
			return 0
		}
	}
	aStr := *hxrt.StdString(a)
	bStr := *hxrt.StdString(b)
	if aStr < bStr {
		return -1
	}
	if aStr > bStr {
		return 1
	}
	return 0
}

func Reflect_field(obj any, field *string) any {
	if obj == nil {
		return nil
	}
	key := *hxrt.StdString(field)
	switch value := obj.(type) {
	case map[string]any:
		return value[key]
	case map[any]any:
		return value[key]
	case *map[string]any:
		if value == nil {
			return nil
		}
		return (*value)[key]
	case *map[any]any:
		if value == nil {
			return nil
		}
		return (*value)[key]
	}
	rv := reflect.ValueOf(obj)
	if !rv.IsValid() {
		return nil
	}
	if rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return nil
		}
		rv = rv.Elem()
	}
	if rv.Kind() == reflect.Struct {
		if fieldValue := rv.FieldByName(key); fieldValue.IsValid() && fieldValue.CanInterface() {
			return fieldValue.Interface()
		}
	}
	method := reflect.ValueOf(obj).MethodByName(key)
	if method.IsValid() {
		return method.Interface()
	}
	return nil
}

func Reflect_hasField(obj any, field *string) bool {
	if obj == nil {
		return false
	}
	key := *hxrt.StdString(field)
	switch value := obj.(type) {
	case map[string]any:
		_, ok := value[key]
		return ok
	case map[any]any:
		_, ok := value[key]
		return ok
	case *map[string]any:
		if value == nil {
			return false
		}
		_, ok := (*value)[key]
		return ok
	case *map[any]any:
		if value == nil {
			return false
		}
		_, ok := (*value)[key]
		return ok
	}
	rv := reflect.ValueOf(obj)
	if !rv.IsValid() {
		return false
	}
	if rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return false
		}
		rv = rv.Elem()
	}
	if rv.Kind() == reflect.Struct {
		if rv.FieldByName(key).IsValid() {
			return true
		}
	}
	return reflect.ValueOf(obj).MethodByName(key).IsValid()
}

func Reflect_setField(obj any, field *string, value any) {
	if obj == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		return
	}
	key := *hxrt.StdString(field)
	switch target := obj.(type) {
	case map[string]any:
		target[key] = value
		return
	case map[any]any:
		target[key] = value
		return
	case *map[string]any:
		if target == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
			return
		}
		(*target)[key] = value
		return
	case *map[any]any:
		if target == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
			return
		}
		(*target)[key] = value
		return
	}
	rv := reflect.ValueOf(obj)
	if !rv.IsValid() || rv.Kind() != reflect.Pointer {
		return
	}
	if rv.IsNil() {
		hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		return
	}
	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return
	}
	fieldValue := rv.FieldByName(key)
	if !fieldValue.IsValid() || !fieldValue.CanSet() {
		return
	}
	if value == nil {
		fieldValue.Set(reflect.Zero(fieldValue.Type()))
		return
	}
	incoming := reflect.ValueOf(value)
	if incoming.Type().AssignableTo(fieldValue.Type()) {
		fieldValue.Set(incoming)
		return
	}
	if incoming.Type().ConvertibleTo(fieldValue.Type()) {
		fieldValue.Set(incoming.Convert(fieldValue.Type()))
		return
	}
	if fieldValue.Kind() == reflect.Interface {
		fieldValue.Set(incoming)
	}
}

type Xml struct {
	raw *string
}

func Xml_parse(source *string) *Xml {
	return haxe__xml__Parser_parse(source)
}

func (self *Xml) toString() *string {
	if self == nil || self.raw == nil {
		return hxrt.StringFromLiteral("")
	}
	return hxrt.StringFromLiteral(*self.raw)
}

type haxe__crypto__Base64 struct {
}

type haxe__crypto__Md5 struct {
}

type haxe__crypto__Sha1 struct {
}

type haxe__crypto__Sha224 struct {
}

type haxe__crypto__Sha256 struct {
}

func hxrt_haxeBytesToRaw(value *haxe__io__Bytes) []byte {
	if value == nil {
		return []byte{}
	}
	if value.__hx_rawValid && len(value.__hx_raw) == len(value.b) {
		return value.__hx_raw
	}
	raw := make([]byte, len(value.b))
	for i := 0; i < len(value.b); i++ {
		raw[i] = byte(value.b[i])
	}
	value.__hx_raw = raw
	value.__hx_rawValid = true
	return raw
}

func hxrt_rawToHaxeBytes(value []byte) *haxe__io__Bytes {
	converted := make([]int, len(value))
	for i := 0; i < len(value); i++ {
		converted[i] = int(value[i])
	}
	return &haxe__io__Bytes{b: converted, length: len(converted), __hx_raw: value, __hx_rawValid: true}
}

func haxe__crypto__Base64_encode(bytes *haxe__io__Bytes, complement ...bool) *string {
	useComplement := true
	if len(complement) > 0 {
		useComplement = complement[0]
	}
	encoded := base64.StdEncoding.EncodeToString(hxrt_haxeBytesToRaw(bytes))
	if !useComplement {
		encoded = strings.TrimRight(encoded, "=")
	}
	return hxrt.StringFromLiteral(encoded)
}

func haxe__crypto__Base64_decode(value *string, complement ...bool) *haxe__io__Bytes {
	useComplement := true
	if len(complement) > 0 {
		useComplement = complement[0]
	}
	rawValue := *hxrt.StdString(value)
	if useComplement {
		rawValue = strings.TrimRight(rawValue, "=")
	}
	decoded, err := base64.RawStdEncoding.DecodeString(rawValue)
	if err != nil {
		decoded, err = base64.StdEncoding.DecodeString(*hxrt.StdString(value))
		if err != nil {
			hxrt.Throw(err)
			return &haxe__io__Bytes{b: []int{}, length: 0}
		}
	}
	return hxrt_rawToHaxeBytes(decoded)
}

func haxe__crypto__Base64_urlEncode(bytes *haxe__io__Bytes, complement ...bool) *string {
	useComplement := false
	if len(complement) > 0 {
		useComplement = complement[0]
	}
	encoded := base64.RawURLEncoding.EncodeToString(hxrt_haxeBytesToRaw(bytes))
	if useComplement {
		missing := len(encoded) % 4
		if missing != 0 {
			encoded = (encoded + strings.Repeat("=", (4-missing)))
		}
	}
	return hxrt.StringFromLiteral(encoded)
}

func haxe__crypto__Base64_urlDecode(value *string, complement ...bool) *haxe__io__Bytes {
	rawValue := *hxrt.StdString(value)
	decoded, err := base64.RawURLEncoding.DecodeString(strings.TrimRight(rawValue, "="))
	if err != nil {
		hxrt.Throw(err)
		return &haxe__io__Bytes{b: []int{}, length: 0}
	}
	return hxrt_rawToHaxeBytes(decoded)
}

func haxe__crypto__Md5_encode(value *string) *string {
	sum := md5.Sum([]byte(*hxrt.StdString(value)))
	return hxrt.StringFromLiteral(hex.EncodeToString(sum[:]))
}

func haxe__crypto__Md5_make(value *haxe__io__Bytes) *haxe__io__Bytes {
	sum := md5.Sum(hxrt_haxeBytesToRaw(value))
	return hxrt_rawToHaxeBytes(sum[:])
}

func haxe__crypto__Sha1_encode(value *string) *string {
	sum := sha1.Sum([]byte(*hxrt.StdString(value)))
	return hxrt.StringFromLiteral(hex.EncodeToString(sum[:]))
}

func haxe__crypto__Sha1_make(value *haxe__io__Bytes) *haxe__io__Bytes {
	sum := sha1.Sum(hxrt_haxeBytesToRaw(value))
	return hxrt_rawToHaxeBytes(sum[:])
}

func haxe__crypto__Sha224_encode(value *string) *string {
	sum := sha256.Sum224([]byte(*hxrt.StdString(value)))
	return hxrt.StringFromLiteral(hex.EncodeToString(sum[:]))
}

func haxe__crypto__Sha224_make(value *haxe__io__Bytes) *haxe__io__Bytes {
	sum := sha256.Sum224(hxrt_haxeBytesToRaw(value))
	return hxrt_rawToHaxeBytes(sum[:])
}

func haxe__crypto__Sha256_encode(value *string) *string {
	sum := sha256.Sum256([]byte(*hxrt.StdString(value)))
	return hxrt.StringFromLiteral(hex.EncodeToString(sum[:]))
}

func haxe__crypto__Sha256_make(value *haxe__io__Bytes) *haxe__io__Bytes {
	sum := sha256.Sum256(hxrt_haxeBytesToRaw(value))
	return hxrt_rawToHaxeBytes(sum[:])
}

type haxe__ds__BalancedTree struct {
}

type haxe__ds__Option struct {
	tag    int
	params []any
}

var haxe__ds__Option_None *haxe__ds__Option = &haxe__ds__Option{tag: 1, params: []any{}}

func haxe__ds__Option_Some(value any) *haxe__ds__Option {
	return &haxe__ds__Option{tag: 0, params: []any{value}}
}

type haxe__io__BytesInput struct {
}

type haxe__io__BytesOutput struct {
}

type haxe__io__Eof struct {
}

type haxe__io__Error struct {
}

type haxe__io__Path struct {
	dir       *string
	file      *string
	ext       *string
	backslash bool
}

func New_haxe__io__Path(path *string) *haxe__io__Path {
	raw := *hxrt.StdString(path)
	dir := filepath.Dir(raw)
	if dir == "." {
		dir = ""
	}
	base := filepath.Base(raw)
	dotExt := filepath.Ext(base)
	file := base
	if dotExt != "" {
		file = strings.TrimSuffix(base, dotExt)
	}
	ext := strings.TrimPrefix(dotExt, ".")
	return &haxe__io__Path{dir: hxrt.StringFromLiteral(dir), file: hxrt.StringFromLiteral(file), ext: hxrt.StringFromLiteral(ext), backslash: strings.Contains(raw, "\\")}
}

func haxe__io__Path_join(parts []*string) *string {
	if len(parts) == 0 {
		return hxrt.StringFromLiteral("")
	}
	joined := filepath.ToSlash(filepath.Join(hxrt.StringSlice(parts)...))
	return hxrt.StringFromLiteral(joined)
}

type haxe__io__StringInput struct {
}

type haxe__xml__Parser struct {
}

type haxe__xml__Printer struct {
}

func haxe__xml__Parser_parse(source *string, strict ...bool) *Xml {
	raw := *hxrt.StdString(source)
	decoder := xml.NewDecoder(strings.NewReader(raw))
	for {
		_, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			hxrt.Throw(err)
			return &Xml{raw: hxrt.StringFromLiteral("")}
		}
	}
	return &Xml{raw: hxrt.StringFromLiteral(raw)}
}

func haxe__xml__Printer_print(value *Xml, pretty ...bool) *string {
	if value == nil || value.raw == nil {
		return hxrt.StringFromLiteral("")
	}
	return hxrt.StringFromLiteral(*value.raw)
}

type haxe__zip__Compress struct {
}

type haxe__zip__Uncompress struct {
}

func haxe__zip__Compress_run(src *haxe__io__Bytes, level int) *haxe__io__Bytes {
	raw := hxrt_haxeBytesToRaw(src)
	var buffer bytes.Buffer
	writer, err := zlib.NewWriterLevel(&buffer, level)
	if err != nil {
		hxrt.Throw(err)
		return nil
	}
	if _, err := writer.Write(raw); err != nil {
		_ = writer.Close()
		hxrt.Throw(err)
		return nil
	}
	if err := writer.Close(); err != nil {
		hxrt.Throw(err)
		return nil
	}
	return hxrt_rawToHaxeBytes(buffer.Bytes())
}

func haxe__zip__Uncompress_run(src *haxe__io__Bytes, bufsize ...int) *haxe__io__Bytes {
	raw := hxrt_haxeBytesToRaw(src)
	reader, err := zlib.NewReader(bytes.NewReader(raw))
	if err != nil {
		hxrt.Throw(err)
		return nil
	}
	defer reader.Close()
	decoded, err := io.ReadAll(reader)
	if err != nil {
		hxrt.Throw(err)
		return nil
	}
	return hxrt_rawToHaxeBytes(decoded)
}

type sys__FileSystem struct {
}
