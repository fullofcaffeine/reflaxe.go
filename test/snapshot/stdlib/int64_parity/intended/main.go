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

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

func emit(label *string, value *haxe___Int64_____Int64) {
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(label, hxrt.StringFromLiteral("=")), haxe___Int64__Int64_Impl__toString(value)))
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
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("wrap_add="), haxe___Int64__Int64_Impl__toString(value)))
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
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("wrap_sub="), haxe___Int64__Int64_Impl__toString(value_1)))
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
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("sum="), haxe___Int64__Int64_Impl__toString(value_2)))
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
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("diff="), haxe___Int64__Int64_Impl__toString(value_3)))
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
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("mul="), haxe___Int64__Int64_Impl__toString(value_4)))
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
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("div_q="), haxe___Int64__Int64_Impl__toString(value_5)))
	value_6 := func(hx_obj_10 map[string]any) *haxe___Int64_____Int64 {
		hx_field_11 := hx_obj_10["modulus"]
		if hx_field_11 == nil {
			var hx_zero_12 *haxe___Int64_____Int64
			return hx_zero_12
		}
		return hx_field_11.(*haxe___Int64_____Int64)
	}(positive)
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("div_r="), haxe___Int64__Int64_Impl__toString(value_6)))
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
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("div_neg_q="), haxe___Int64__Int64_Impl__toString(value_7)))
	value_8 := func(hx_obj_16 map[string]any) *haxe___Int64_____Int64 {
		hx_field_17 := hx_obj_16["modulus"]
		if hx_field_17 == nil {
			var hx_zero_18 *haxe___Int64_____Int64
			return hx_zero_18
		}
		return hx_field_17.(*haxe___Int64_____Int64)
	}(negative)
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("div_neg_r="), haxe___Int64__Int64_Impl__toString(value_8)))
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("cmp="), func() int {
		var a_low_1 int
		var a_high_1 int
		a_high_1 = -1
		a_low_1 = -1
		var b_low_3 int
		var b_high_3 int
		b_high_3 = 0
		b_low_3 = 1
		v := int(int32((hxrt.Int32Wrap(a_high_1) - hxrt.Int32Wrap(b_high_3))))
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
		var a_high_2 int
		a_high_2 = -1
		a_low_2 = -1
		var b_low_4 int
		var b_high_4 int
		b_high_4 = 0
		b_low_4 = 1
		v_1 := haxe___Int32__Int32_Impl__ucompare(a_high_2, b_high_4)
		var hx_if_23 int
		if v_1 != 0 {
			hx_if_23 = v_1
		} else {
			hx_if_23 = haxe___Int32__Int32_Impl__ucompare(a_low_2, b_low_4)
		}
		return hx_if_23
	}()))
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
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("shl="), haxe___Int64__Int64_Impl__toString(value_9)))
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
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("shr="), haxe___Int64__Int64_Impl__toString(value_10)))
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
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("ushr="), haxe___Int64__Int64_Impl__toString(value_11)))
	value_12 := haxe__Int64Helper_fromFloat(9007199254740991.0)
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("from_float="), haxe___Int64__Int64_Impl__toString(value_12)))
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("to_int_ok="), func() int {
		var x_low int
		var x_high int
		x_high = 0
		x_low = 2147483647
		if x_high != int(int32((hxrt.Int32Wrap(x_low) >> uint(31)))) {
			hxrt.Throw(hxrt.StringFromLiteral("Overflow"))
		}
		return x_low
	}()))
	hxrt.TryCatch(func() {
		x_16 := haxe__Int64Helper_parseString(hxrt.StringFromLiteral("2147483648"))
		if x_16.high != int(int32((hxrt.Int32Wrap(x_16.low) >> uint(31)))) {
			hxrt.Throw(hxrt.StringFromLiteral("Overflow"))
		}
		_ = x_16.low
		hxrt.Println(hxrt.StringFromLiteral("to_int_overflow=missing"))
	}, func(hx_caught_30 any) {
		e := hx_caught_30
		hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("to_int_overflow="), hxrt.StdString(e)))
	})
	var round_low int
	var round_high int
	round_high = 2147483647
	round_low = -12345
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("round_high="), round_high))
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("round_low="), round_low))
}

type haxe__io__Encoding struct {
	tag int
}

type haxe__io__Input interface {
	get_bigEndian() bool
	set_bigEndian(e bool) bool
	readByte() int
	readBytes(buf *haxe__io__Bytes, pos int, len int) int
	close()
}

type haxe__io__Output interface {
	get_bigEndian() bool
	set_bigEndian(e bool) bool
	writeByte(c int)
	writeBytes(s *haxe__io__Bytes, pos int, len int) int
	flush()
	close()
}

type haxe__io__Eof struct {
}

type haxe__io__Error struct {
	tag    int
	params []any
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

type haxe__io__BytesInput struct {
	bigEndian bool
	b         []int
	pos       int
	len       int
	totlen    int
}

type haxe__io__BytesOutput struct {
	bigEndian bool
	b         *haxe__io__BytesBuffer
}

func New_haxe__io__Input() haxe__io__Input {
	return New_haxe__io__BytesInput(&haxe__io__Bytes{b: []int{}, length: 0})
}

func New_haxe__io__Output() haxe__io__Output {
	return New_haxe__io__BytesOutput()
}

func New_haxe__io__Eof() *haxe__io__Eof {
	return &haxe__io__Eof{}
}

var haxe__io__Encoding_UTF8 *haxe__io__Encoding = &haxe__io__Encoding{tag: 0}

var haxe__io__Encoding_RawNative *haxe__io__Encoding = &haxe__io__Encoding{tag: 1}

func (self *haxe__io__Encoding) String() string {
	if self == nil {
		return "null"
	}
	switch self.tag {
	case 0:
		return "UTF8"
	case 1:
		return "RawNative"
	default:
		return "Encoding"
	}
}

func (self *haxe__io__Encoding) toString() *string {
	return hxrt.StringFromLiteral(self.String())
}

func haxe__io__resolveEncodingTag(encoding ...*haxe__io__Encoding) int {
	if len(encoding) == 0 || encoding[0] == nil {
		return 0
	}
	return encoding[0].tag
}

func haxe__io__bytes_fromStringRawNativeUTF16LE(value *string) *haxe__io__Bytes {
	raw := []byte(*hxrt.StdString(value))
	converted := make([]int, len(raw))
	for i := 0; i < len(raw); i++ {
		converted[i] = int(raw[i])
	}
	return &haxe__io__Bytes{b: converted, length: len(converted), __hx_raw: raw, __hx_rawValid: true}
}

func haxe__io__bytes_toStringRawNativeUTF16LE(value []int) *string {
	return hxrt.BytesToString(value)
}

func (self *haxe__io__Eof) toString() *string {
	return hxrt.StringFromLiteral("Eof")
}

var haxe__io__Error_Blocked *haxe__io__Error = &haxe__io__Error{tag: 0}

var haxe__io__Error_Overflow *haxe__io__Error = &haxe__io__Error{tag: 1}

var haxe__io__Error_OutsideBounds *haxe__io__Error = &haxe__io__Error{tag: 2}

func haxe__io__Error_Custom(e any) *haxe__io__Error {
	return &haxe__io__Error{tag: 3, params: []any{e}}
}

func (self *haxe__io__Error) String() string {
	if self == nil {
		return "null"
	}
	switch self.tag {
	case 0:
		return "Blocked"
	case 1:
		return "Overflow"
	case 2:
		return "OutsideBounds"
	case 3:
		if len(self.params) == 0 {
			return "Custom(null)"
		}
		return "Custom(" + *hxrt.StdString(self.params[0]) + ")"
	default:
		return "Error"
	}
}

func (self *haxe__io__Error) toString() *string {
	return hxrt.StringFromLiteral(self.String())
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
	if haxe__io__resolveEncodingTag(encoding...) == 1 {
		return haxe__io__bytes_fromStringRawNativeUTF16LE(value)
	}
	raw := []byte(*hxrt.StdString(value))
	converted := make([]int, len(raw))
	for i := 0; i < len(raw); i++ {
		converted[i] = int(raw[i])
	}
	return &haxe__io__Bytes{b: converted, length: len(converted), __hx_raw: raw, __hx_rawValid: true}
}

func haxe__io__Bytes_ofData(b []int) *haxe__io__Bytes {
	if b == nil {
		return &haxe__io__Bytes{b: []int{}, length: 0}
	}
	return &haxe__io__Bytes{b: b, length: len(b)}
}

func haxe__io__Bytes_ofHex(s *string) *haxe__io__Bytes {
	raw := *hxrt.StdString(s)
	lenValue := len(raw)
	if (lenValue & 1) != 0 {
		hxrt.Throw(hxrt.StringFromLiteral("Not a hex string (odd number of digits)"))
		return &haxe__io__Bytes{b: []int{}, length: 0}
	}
	ret := haxe__io__Bytes_alloc(lenValue >> 1)
	for i := 0; i < ret.length; i++ {
		high := int(raw[i*2])
		low := int(raw[i*2+1])
		high = (high & 0xF) + ((high&0x40)>>6)*9
		low = (low & 0xF) + ((low&0x40)>>6)*9
		ret.set(i, ((high<<4)|low)&0xFF)
	}
	return ret
}

func (self *haxe__io__Bytes) toString() *string {
	if self == nil {
		return hxrt.StringFromLiteral("")
	}
	return hxrt.BytesToString(self.b)
}

func (self *haxe__io__Bytes) toHex() *string {
	if self == nil || self.length == 0 {
		return hxrt.StringFromLiteral("")
	}
	hexChars := "0123456789abcdef"
	out := make([]byte, self.length*2)
	for i := 0; i < self.length; i++ {
		c := self.b[i] & 0xFF
		out[i*2] = hexChars[c>>4]
		out[i*2+1] = hexChars[c&15]
	}
	return hxrt.StringFromLiteral(string(out))
}

func (self *haxe__io__Bytes) getData() []int {
	if self == nil {
		return []int{}
	}
	return self.b
}

func (self *haxe__io__Bytes) getString(pos int, len int, encoding ...*haxe__io__Encoding) *string {
	if self == nil || pos < 0 || len < 0 || pos+len > self.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return hxrt.StringFromLiteral("")
	}
	slice := self.b[pos : pos+len]
	if haxe__io__resolveEncodingTag(encoding...) == 1 {
		return haxe__io__bytes_toStringRawNativeUTF16LE(slice)
	}
	return hxrt.BytesToString(slice)
}

func (self *haxe__io__Bytes) readString(pos int, len int) *string {
	return self.getString(pos, len)
}

func (self *haxe__io__Bytes) get(pos int) int {
	return self.b[pos]
}

func (self *haxe__io__Bytes) set(pos int, value int) {
	self.b[pos] = value
	self.__hx_rawValid = false
}

func (self *haxe__io__Bytes) blit(pos int, src *haxe__io__Bytes, srcpos int, len int) {
	if self == nil || src == nil || pos < 0 || srcpos < 0 || len < 0 || pos+len > self.length || srcpos+len > src.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return
	}
	if len == 0 {
		return
	}
	if self == src && pos > srcpos {
		for i := len - 1; i >= 0; i-- {
			self.b[pos+i] = src.b[srcpos+i]
		}
	} else {
		for i := 0; i < len; i++ {
			self.b[pos+i] = src.b[srcpos+i]
		}
	}
	self.__hx_rawValid = false
}

func (self *haxe__io__Bytes) fill(pos int, len int, value int) {
	if self == nil || pos < 0 || len < 0 || pos+len > self.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return
	}
	masked := value & 255
	for i := 0; i < len; i++ {
		self.b[pos+i] = masked
	}
	self.__hx_rawValid = false
}

func (self *haxe__io__Bytes) sub(pos int, len int) *haxe__io__Bytes {
	if self == nil || pos < 0 || len < 0 || pos+len > self.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return &haxe__io__Bytes{b: []int{}, length: 0}
	}
	if len == 0 {
		return &haxe__io__Bytes{b: []int{}, length: 0}
	}
	copied := make([]int, len)
	copy(copied, self.b[pos:pos+len])
	return &haxe__io__Bytes{b: copied, length: len}
}

func (self *haxe__io__Bytes) compare(other *haxe__io__Bytes) int {
	if self == nil && other == nil {
		return 0
	}
	if self == nil {
		return -1
	}
	if other == nil {
		return 1
	}
	limit := self.length
	if other.length < limit {
		limit = other.length
	}
	for i := 0; i < limit; i++ {
		if self.b[i] < other.b[i] {
			return -1
		}
		if self.b[i] > other.b[i] {
			return 1
		}
	}
	if self.length < other.length {
		return -1
	}
	if self.length > other.length {
		return 1
	}
	return 0
}

func New_haxe__io__BytesBuffer() *haxe__io__BytesBuffer {
	return &haxe__io__BytesBuffer{b: []int{}}
}

func (self *haxe__io__BytesBuffer) addByte(value int) {
	self.b = append(self.b, (value & 255))
}

func (self *haxe__io__BytesBuffer) add(src *haxe__io__Bytes) {
	if src == nil {
		return
	}
	self.b = append(self.b, src.b...)
}

func (self *haxe__io__BytesBuffer) addBytes(src *haxe__io__Bytes, pos int, len int) {
	if src == nil || pos < 0 || len < 0 || pos+len > src.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return
	}
	if len == 0 {
		return
	}
	self.b = append(self.b, src.b[pos:pos+len]...)
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

func New_haxe__io__BytesInput(b *haxe__io__Bytes, opts ...int) *haxe__io__BytesInput {
	if b == nil {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return &haxe__io__BytesInput{}
	}
	start := 0
	if len(opts) > 0 {
		start = opts[0]
	}
	sliceLen := (b.length - start)
	if len(opts) > 1 {
		sliceLen = opts[1]
	}
	if start < 0 || sliceLen < 0 || start+sliceLen > b.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return &haxe__io__BytesInput{}
	}
	return &haxe__io__BytesInput{b: b.b, pos: start, len: sliceLen, totlen: sliceLen}
}

func (self *haxe__io__BytesInput) get_position() int {
	return self.pos
}

func (self *haxe__io__BytesInput) set_position(p int) int {
	if p < 0 {
		p = 0
	} else {
		if p > self.totlen {
			p = self.totlen
		}
	}
	self.len = (self.totlen - p)
	self.pos = p
	return p
}

func (self *haxe__io__BytesInput) get_length() int {
	return self.totlen
}

func (self *haxe__io__BytesInput) readByte() int {
	if self == nil || self.len == 0 {
		hxrt.Throw(&haxe__io__Eof{})
		return 0
	}
	self.len = (self.len - 1)
	value := self.b[self.pos]
	self.pos = (self.pos + 1)
	return value
}

func (self *haxe__io__BytesInput) readBytes(buf *haxe__io__Bytes, pos int, len int) int {
	if buf == nil || pos < 0 || len < 0 || pos+len > buf.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return 0
	}
	if len > 0 && (self == nil || self.len == 0) {
		hxrt.Throw(&haxe__io__Eof{})
		return 0
	}
	if self == nil {
		return 0
	}
	if self.len < len {
		len = self.len
	}
	for i := 0; i < len; i++ {
		buf.b[pos+i] = self.b[self.pos+i]
	}
	self.pos += len
	self.len -= len
	return len
}

func (self *haxe__io__BytesInput) get_bigEndian() bool {
	if self == nil {
		return false
	}
	return self.bigEndian
}

func (self *haxe__io__BytesInput) set_bigEndian(e bool) bool {
	if self != nil {
		self.bigEndian = e
	}
	return e
}

func (self *haxe__io__BytesInput) close() {
	_ = self
}

func New_haxe__io__BytesOutput() *haxe__io__BytesOutput {
	return &haxe__io__BytesOutput{b: &haxe__io__BytesBuffer{b: []int{}}}
}

func (self *haxe__io__BytesOutput) get_length() int {
	if self == nil || self.b == nil {
		return 0
	}
	return self.b.get_length()
}

func (self *haxe__io__BytesOutput) writeByte(c int) {
	if self == nil || self.b == nil {
		return
	}
	self.b.addByte(c)
}

func (self *haxe__io__BytesOutput) writeBytes(buf *haxe__io__Bytes, pos int, len int) int {
	if buf == nil || pos < 0 || len < 0 || pos+len > buf.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return 0
	}
	if self == nil || self.b == nil {
		return 0
	}
	self.b.addBytes(buf, pos, len)
	return len
}

func (self *haxe__io__BytesOutput) get_bigEndian() bool {
	if self == nil {
		return false
	}
	return self.bigEndian
}

func (self *haxe__io__BytesOutput) set_bigEndian(e bool) bool {
	if self != nil {
		self.bigEndian = e
	}
	return e
}

func (self *haxe__io__BytesOutput) flush() {
	_ = self
}

func (self *haxe__io__BytesOutput) close() {
	_ = self
}

func (self *haxe__io__BytesOutput) getBytes() *haxe__io__Bytes {
	if self == nil || self.b == nil {
		return &haxe__io__Bytes{b: []int{}, length: 0}
	}
	return self.b.getBytes()
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

func _UnicodeString__UnicodeString_Impl__get_length(value any) int {
	return len([]rune(*hxrt.StdString(value)))
}

func _UnicodeString__UnicodeString_Impl__charAt(value any, index int) *string {
	if index < 0 {
		return hxrt.StringFromLiteral("")
	}
	runes := []rune(*hxrt.StdString(value))
	if index >= len(runes) {
		return hxrt.StringFromLiteral("")
	}
	return hxrt.StringFromLiteral(string(runes[index]))
}

func _UnicodeString__UnicodeString_Impl__charCodeAt(value any, index int) any {
	if index < 0 {
		return nil
	}
	runes := []rune(*hxrt.StdString(value))
	if index >= len(runes) {
		return nil
	}
	return int(runes[index])
}

func _UnicodeString__UnicodeString_Impl__substring(value any, startIndex int, endIndex ...int) *string {
	runes := []rune(*hxrt.StdString(value))
	end := len(runes)
	if len(endIndex) > 0 {
		end = endIndex[0]
	}
	if startIndex < 0 {
		startIndex = 0
	}
	if end < 0 {
		end = 0
	}
	if startIndex == end {
		return hxrt.StringFromLiteral("")
	}
	if startIndex > end {
		startIndex, end = end, startIndex
	}
	if startIndex > len(runes) {
		return hxrt.StringFromLiteral("")
	}
	if end > len(runes) {
		end = len(runes)
	}
	return hxrt.StringFromLiteral(string(runes[startIndex:end]))
}

func _UnicodeString__UnicodeString_Impl__substr(value any, pos int, lengthArgs ...int) *string {
	runes := []rune(*hxrt.StdString(value))
	unicodeLength := len(runes)
	if pos < 0 {
		pos = unicodeLength + pos
		if pos < 0 {
			pos = 0
		}
	}
	if pos > unicodeLength {
		return hxrt.StringFromLiteral("")
	}
	if len(lengthArgs) == 0 {
		return hxrt.StringFromLiteral(string(runes[pos:]))
	}
	lengthValue := lengthArgs[0]
	end := unicodeLength
	if lengthValue < 0 {
		end = unicodeLength + lengthValue
	} else {
		end = pos + lengthValue
	}
	if end < pos {
		return hxrt.StringFromLiteral("")
	}
	if end > unicodeLength {
		end = unicodeLength
	}
	return hxrt.StringFromLiteral(string(runes[pos:end]))
}

func _UnicodeString__UnicodeString_Impl__indexOf(value any, str *string, startIndex ...int) int {
	runes := []rune(*hxrt.StdString(value))
	needle := []rune(*hxrt.StdString(str))
	start := 0
	if len(startIndex) > 0 {
		start = startIndex[0]
		if start < 0 {
			start = len(runes) + start
		}
	}
	if start < 0 {
		start = 0
	}
	if len(needle) == 0 {
		if start > len(runes) {
			return len(runes)
		}
		return start
	}
	if start > len(runes) || len(needle) > len(runes) {
		return -1
	}
	for i := start; i+len(needle) <= len(runes); i++ {
		matched := true
		for j := 0; j < len(needle); j++ {
			if runes[i+j] != needle[j] {
				matched = false
				break
			}
		}
		if matched {
			return i
		}
	}
	return -1
}

func _UnicodeString__UnicodeString_Impl__lastIndexOf(value any, str *string, startIndex ...int) int {
	runes := []rune(*hxrt.StdString(value))
	needle := []rune(*hxrt.StdString(str))
	if len(needle) == 0 {
		if len(startIndex) == 0 {
			return len(runes)
		}
		start := startIndex[0]
		if start < 0 {
			start = 0
		}
		if start > len(runes) {
			return len(runes)
		}
		return start
	}
	start := len(runes)
	if len(startIndex) > 0 {
		start = startIndex[0]
		if start < 0 {
			start = 0
		}
	}
	limit := start + len(needle)
	if limit > len(runes) {
		limit = len(runes)
	}
	for i := limit - len(needle); i >= 0; i-- {
		matched := true
		for j := 0; j < len(needle); j++ {
			if runes[i+j] != needle[j] {
				matched = false
				break
			}
		}
		if matched {
			return i
		}
	}
	return -1
}

func _UnicodeString__UnicodeString_Impl__validate(value *haxe__io__Bytes, encoding *haxe__io__Encoding) bool {
	if haxe__io__resolveEncodingTag(encoding) == 1 {
		hxrt.Throw(hxrt.StringFromLiteral("UnicodeString.validate: RawNative encoding is not supported"))
		return false
	}
	raw := hxrt_haxeBytesToRaw(value)
	pos := 0
	max := len(raw)
	for pos < max {
		c := int(raw[pos])
		pos++
		if c < 0x80 {
			continue
		} else if c < 0xC2 {
			return false
		} else if c < 0xE0 {
			if pos+1 > max {
				return false
			}
			c2 := int(raw[pos])
			pos++
			if c2 < 0x80 || c2 > 0xBF {
				return false
			}
		} else if c < 0xF0 {
			if pos+2 > max {
				return false
			}
			c2 := int(raw[pos])
			pos++
			if c == 0xE0 {
				if c2 < 0xA0 || c2 > 0xBF {
					return false
				}
			} else if c2 < 0x80 || c2 > 0xBF {
				return false
			}
			c3 := int(raw[pos])
			pos++
			if c3 < 0x80 || c3 > 0xBF {
				return false
			}
			c = (c << 16) | (c2 << 8) | c3
			if 0xEDA080 <= c && c <= 0xEDBFBF {
				return false
			}
		} else if c > 0xF4 {
			return false
		} else {
			if pos+3 > max {
				return false
			}
			c2 := int(raw[pos])
			pos++
			if c == 0xF0 {
				if c2 < 0x90 || c2 > 0xBF {
					return false
				}
			} else if c == 0xF4 {
				if c2 < 0x80 || c2 > 0x8F {
					return false
				}
			} else if c2 < 0x80 || c2 > 0xBF {
				return false
			}
			c3 := int(raw[pos])
			pos++
			if c3 < 0x80 || c3 > 0xBF {
				return false
			}
			c4 := int(raw[pos])
			pos++
			if c4 < 0x80 || c4 > 0xBF {
				return false
			}
		}
	}
	return true
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

func Date_fromTime(ms float64) *Date {
	nanos := int64(ms * 1e6)
	return &Date{value: time.Unix(0, nanos).In(time.Local)}
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

func (self *Date) getTime() float64 {
	return float64(self.value.UnixNano()) / 1e6
}

type DateTools struct {
}

func DateTools_format(date *Date, format *string) *string {
	layout := *hxrt.StdString(format)
	layout = strings.ReplaceAll(layout, "%%", "__HX_PERCENT__")
	layout = strings.ReplaceAll(layout, "%Y", "2006")
	layout = strings.ReplaceAll(layout, "%m", "01")
	layout = strings.ReplaceAll(layout, "%d", "02")
	layout = strings.ReplaceAll(layout, "%H", "15")
	layout = strings.ReplaceAll(layout, "%M", "04")
	layout = strings.ReplaceAll(layout, "%S", "05")
	layout = strings.ReplaceAll(layout, "__HX_PERCENT__", "%")
	return hxrt.StringFromLiteral(date.value.Format(layout))
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

var Xml_Element int = 0

var Xml_PCData int = 1

var Xml_CData int = 2

var Xml_Comment int = 3

var Xml_DocType int = 4

var Xml_ProcessingInstruction int = 5

var Xml_Document int = 6

type Xml struct {
	nodeType       int
	nodeName       *string
	nodeValue      *string
	parent         *Xml
	children       []*Xml
	attributeMap   map[string]string
	attributeOrder []string
}

func _Xml__XmlType_Impl__toString(value int) *string {
	switch value {
	case Xml_Element:
		return hxrt.StringFromLiteral("Element")
	case Xml_PCData:
		return hxrt.StringFromLiteral("PCData")
	case Xml_CData:
		return hxrt.StringFromLiteral("CData")
	case Xml_Comment:
		return hxrt.StringFromLiteral("Comment")
	case Xml_DocType:
		return hxrt.StringFromLiteral("DocType")
	case Xml_ProcessingInstruction:
		return hxrt.StringFromLiteral("ProcessingInstruction")
	case Xml_Document:
		return hxrt.StringFromLiteral("Document")
	default:
		return hxrt.StringFromLiteral("XmlType")
	}
}

func New_Xml(nodeType int) *Xml {
	return &Xml{nodeType: nodeType, children: []*Xml{}, attributeMap: map[string]string{}, attributeOrder: []string{}}
}

func Xml_createElement(name *string) *Xml {
	xml := New_Xml(Xml_Element)
	xml.nodeName = name
	return xml
}

func Xml_createPCData(data *string) *Xml {
	xml := New_Xml(Xml_PCData)
	xml.nodeValue = data
	return xml
}

func Xml_createCData(data *string) *Xml {
	xml := New_Xml(Xml_CData)
	xml.nodeValue = data
	return xml
}

func Xml_createComment(data *string) *Xml {
	xml := New_Xml(Xml_Comment)
	xml.nodeValue = data
	return xml
}

func Xml_createDocType(data *string) *Xml {
	xml := New_Xml(Xml_DocType)
	xml.nodeValue = data
	return xml
}

func Xml_createProcessingInstruction(data *string) *Xml {
	xml := New_Xml(Xml_ProcessingInstruction)
	xml.nodeValue = data
	return xml
}

func Xml_createDocument() *Xml {
	return New_Xml(Xml_Document)
}

func Xml_ensureElementType(self *Xml) {
	if self.nodeType != Xml_Document && self.nodeType != Xml_Element {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element or Document but found "), _Xml__XmlType_Impl__toString(self.nodeType)))
	}
}

func Xml_parse(source *string) *Xml {
	return haxe__xml__Parser_parse(source)
}

func (self *Xml) get(att *string) *string {
	if self.nodeType != Xml_Element {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(self.nodeType)))
		return nil
	}
	key := *hxrt.StdString(att)
	value, ok := self.attributeMap[key]
	if !ok {
		return nil
	}
	return hxrt.StringFromLiteral(value)
}

func (self *Xml) set(att *string, value *string) {
	if self.nodeType != Xml_Element {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(self.nodeType)))
		return
	}
	key := *hxrt.StdString(att)
	if _, ok := self.attributeMap[key]; !ok {
		self.attributeOrder = append(self.attributeOrder, key)
	}
	self.attributeMap[key] = *hxrt.StdString(value)
}

func (self *Xml) remove(att *string) {
	if self.nodeType != Xml_Element {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(self.nodeType)))
		return
	}
	key := *hxrt.StdString(att)
	delete(self.attributeMap, key)
	filtered := make([]string, 0, len(self.attributeOrder))
	for _, existing := range self.attributeOrder {
		if existing != key {
			filtered = append(filtered, existing)
		}
	}
	self.attributeOrder = filtered
}

func (self *Xml) exists(att *string) bool {
	if self.nodeType != Xml_Element {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(self.nodeType)))
		return false
	}
	_, ok := self.attributeMap[*hxrt.StdString(att)]
	return ok
}

func (self *Xml) attributes() map[string]any {
	if self.nodeType != Xml_Element {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(self.nodeType)))
		return map[string]any{}
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(self.attributeOrder) }
	iter["next"] = func() *string { key := self.attributeOrder[index]; index++; return hxrt.StringFromLiteral(key) }
	return iter
}

func (self *Xml) iterator() map[string]any {
	Xml_ensureElementType(self)
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(self.children) }
	iter["next"] = func() *Xml { child := self.children[index]; index++; return child }
	return iter
}

func (self *Xml) elements() map[string]any {
	Xml_ensureElementType(self)
	matches := make([]*Xml, 0, len(self.children))
	for _, child := range self.children {
		if child.nodeType == Xml_Element {
			matches = append(matches, child)
		}
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(matches) }
	iter["next"] = func() *Xml { child := matches[index]; index++; return child }
	return iter
}

func (self *Xml) elementsNamed(name *string) map[string]any {
	Xml_ensureElementType(self)
	wanted := *hxrt.StdString(name)
	matches := make([]*Xml, 0, len(self.children))
	for _, child := range self.children {
		if child.nodeType == Xml_Element && *hxrt.StdString(child.nodeName) == wanted {
			matches = append(matches, child)
		}
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(matches) }
	iter["next"] = func() *Xml { child := matches[index]; index++; return child }
	return iter
}

func (self *Xml) firstChild() *Xml {
	Xml_ensureElementType(self)
	if len(self.children) == 0 {
		return nil
	}
	return self.children[0]
}

func (self *Xml) firstElement() *Xml {
	Xml_ensureElementType(self)
	for _, child := range self.children {
		if child.nodeType == Xml_Element {
			return child
		}
	}
	return nil
}

func (self *Xml) addChild(x *Xml) {
	Xml_ensureElementType(self)
	if x == nil {
		return
	}
	if x.parent != nil {
		x.parent.removeChild(x)
	}
	self.children = append(self.children, x)
	x.parent = self
}

func (self *Xml) removeChild(x *Xml) bool {
	Xml_ensureElementType(self)
	for i, child := range self.children {
		if child == x {
			self.children = append(self.children[:i], self.children[i+1:]...)
			x.parent = nil
			return true
		}
	}
	return false
}

func (self *Xml) insertChild(x *Xml, pos int) {
	Xml_ensureElementType(self)
	if x == nil {
		return
	}
	if x.parent != nil {
		x.parent.removeChild(x)
	}
	if pos < 0 {
		pos = 0
	}
	if pos > len(self.children) {
		pos = len(self.children)
	}
	self.children = append(self.children, nil)
	copy(self.children[pos+1:], self.children[pos:])
	self.children[pos] = x
	x.parent = self
}

func (self *Xml) toString() *string {
	if self == nil {
		return hxrt.StringFromLiteral("")
	}
	return haxe__xml__Printer_print(self)
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
	doc := Xml_createDocument()
	stack := []*Xml{doc}
	decoder := xml.NewDecoder(strings.NewReader(raw))
	for {
		tokenStart := decoder.InputOffset()
		token, err := decoder.Token()
		tokenEnd := decoder.InputOffset()
		if err == io.EOF {
			break
		}
		if err != nil {
			hxrt.Throw(err)
			return Xml_createDocument()
		}
		current := stack[len(stack)-1]
		switch value := token.(type) {
		case xml.StartElement:
			node := Xml_createElement(hxrt.StringFromLiteral(value.Name.Local))
			for _, attr := range value.Attr {
				node.set(hxrt.StringFromLiteral(attr.Name.Local), hxrt.StringFromLiteral(attr.Value))
			}
			current.addChild(node)
			stack = append(stack, node)
		case xml.EndElement:
			if len(stack) > 1 {
				stack = stack[:len(stack)-1]
			}
		case xml.CharData:
			text := string([]byte(value))
			if len(text) != 0 {
				tokenSource := raw[tokenStart:tokenEnd]
				if strings.HasPrefix(tokenSource, "<![CDATA[") && strings.HasSuffix(tokenSource, "]]>") {
					current.addChild(Xml_createCData(hxrt.StringFromLiteral(text)))
				} else {
					current.addChild(Xml_createPCData(hxrt.StringFromLiteral(text)))
				}
			}
		case xml.Comment:
			current.addChild(Xml_createComment(hxrt.StringFromLiteral(string([]byte(value)))))
		case xml.Directive:
			directive := strings.TrimSpace(string([]byte(value)))
			upper := strings.ToUpper(directive)
			if strings.HasPrefix(upper, "DOCTYPE") {
				directive = strings.TrimSpace(directive[len("DOCTYPE"):])
			}
			current.addChild(Xml_createDocType(hxrt.StringFromLiteral(directive)))
		case xml.ProcInst:
			payload := value.Target
			if len(value.Inst) != 0 {
				payload += " " + string(value.Inst)
			}
			current.addChild(Xml_createProcessingInstruction(hxrt.StringFromLiteral(strings.TrimSpace(payload))))
		}
	}
	return doc
}

func haxe__xml__Printer_escapeText(value string) string {
	return strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;").Replace(value)
}

func haxe__xml__Printer_escapeAttr(value string) string {
	return strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;", "\"", "&quot;").Replace(value)
}

func haxe__xml__Printer_hasChildren(value *Xml) bool {
	for _, child := range value.children {
		switch child.nodeType {
		case Xml_Element, Xml_PCData:
			return true
		case Xml_CData, Xml_Comment:
			if len(strings.TrimLeft(*hxrt.StdString(child.nodeValue), " \n\r\t")) != 0 {
				return true
			}
		}
	}
	return false
}

func haxe__xml__Printer_writeNode(output *strings.Builder, value *Xml, tabs string, pretty bool) {
	newline := func() {
		if pretty {
			output.WriteString("\n")
		}
	}
	switch value.nodeType {
	case Xml_CData:
		output.WriteString(tabs + "<![CDATA[")
		output.WriteString(*hxrt.StdString(value.nodeValue))
		output.WriteString("]]>")
		newline()
	case Xml_Comment:
		commentContent := strings.NewReplacer("\n", "", "\r", "", "\t", "").Replace(*hxrt.StdString(value.nodeValue))
		output.WriteString(tabs)
		output.WriteString(strings.TrimSpace("<!--" + commentContent + "-->"))
		newline()
	case Xml_Document:
		for _, child := range value.children {
			haxe__xml__Printer_writeNode(output, child, tabs, pretty)
		}
	case Xml_Element:
		output.WriteString(tabs + "<")
		output.WriteString(*hxrt.StdString(value.nodeName))
		for _, attribute := range value.attributeOrder {
			output.WriteString(" " + attribute + "=\"")
			output.WriteString(haxe__xml__Printer_escapeAttr(value.attributeMap[attribute]))
			output.WriteString("\"")
		}
		if haxe__xml__Printer_hasChildren(value) {
			output.WriteString(">")
			newline()
			childTabs := tabs
			if pretty {
				childTabs = tabs + "\t"
			}
			for _, child := range value.children {
				haxe__xml__Printer_writeNode(output, child, childTabs, pretty)
			}
			output.WriteString(tabs + "</")
			output.WriteString(*hxrt.StdString(value.nodeName))
			output.WriteString(">")
			newline()
		} else {
			output.WriteString("/>")
			newline()
		}
	case Xml_PCData:
		nodeValue := *hxrt.StdString(value.nodeValue)
		if len(nodeValue) != 0 {
			output.WriteString(tabs + haxe__xml__Printer_escapeText(nodeValue))
			newline()
		}
	case Xml_ProcessingInstruction:
		output.WriteString("<?" + *hxrt.StdString(value.nodeValue) + "?>")
		newline()
	case Xml_DocType:
		output.WriteString("<!DOCTYPE " + *hxrt.StdString(value.nodeValue) + ">")
		newline()
	}
}

func haxe__xml__Printer_print(value *Xml, pretty ...bool) *string {
	if value == nil {
		return hxrt.StringFromLiteral("")
	}
	usePretty := false
	if len(pretty) > 0 {
		usePretty = pretty[0]
	}
	var output strings.Builder
	haxe__xml__Printer_writeNode(&output, value, "", usePretty)
	return hxrt.StringFromLiteral(output.String())
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

type ValueType struct {
	tag    int
	params []any
}

var ValueType_TNull *ValueType = &ValueType{tag: 0, params: []any{}}

var ValueType_TInt *ValueType = &ValueType{tag: 1, params: []any{}}

var ValueType_TFloat *ValueType = &ValueType{tag: 2, params: []any{}}

var ValueType_TBool *ValueType = &ValueType{tag: 3, params: []any{}}

var ValueType_TObject *ValueType = &ValueType{tag: 4, params: []any{}}

var ValueType_TFunction *ValueType = &ValueType{tag: 5, params: []any{}}

var ValueType_TUnknown *ValueType = &ValueType{tag: 8, params: []any{}}

func ValueType_TClass(c any) *ValueType {
	return &ValueType{tag: 6, params: []any{c}}
}

func ValueType_TEnum(e any) *ValueType {
	return &ValueType{tag: 7, params: []any{e}}
}

func hxrt_typeCallAny(callable any, args []any) (any, bool) {
	result := any(nil)
	ok := false
	defer func() {
		if recover() != nil {
			result = nil
			ok = false
		}
	}()
	if callable == nil {
		return nil, false
	}
	fn := reflect.ValueOf(callable)
	if !fn.IsValid() || fn.Kind() != reflect.Func {
		return nil, false
	}
	fnType := fn.Type()
	if fnType.NumIn() != len(args) {
		return nil, false
	}
	in := make([]reflect.Value, len(args))
	for i := 0; i < len(args); i++ {
		paramType := fnType.In(i)
		arg := args[i]
		if arg == nil {
			in[i] = reflect.Zero(paramType)
			continue
		}
		v := reflect.ValueOf(arg)
		if v.IsValid() && v.Type().AssignableTo(paramType) {
			in[i] = v
			continue
		}
		if v.IsValid() && v.Type().ConvertibleTo(paramType) {
			in[i] = v.Convert(paramType)
			continue
		}
		if paramType.Kind() == reflect.Interface && v.IsValid() {
			in[i] = v
			continue
		}
		return nil, false
	}
	out := fn.Call(in)
	if len(out) == 0 {
		return nil, true
	}
	first := out[0]
	if !first.IsValid() {
		return nil, true
	}
	result = first.Interface()
	ok = true
	return result, ok
}

func hxrt_typeResolvedClassName(value any) (string, bool) {
	switch current := value.(type) {
	case *hxrt__TypeClassValue:
		if current == nil || current.name == nil {
			return "", false
		}
		return *current.name, true
	case hxrt__TypeClassValue:
		if current.name == nil {
			return "", false
		}
		return *current.name, true
	case string:
		return current, true
	case *string:
		if current == nil {
			return "", false
		}
		return *current, true
	default:
		return "", false
	}
}

func hxrt_typeResolvedEnumName(value any) (string, bool) {
	switch current := value.(type) {
	case *hxrt__TypeEnumValue:
		if current == nil || current.name == nil {
			return "", false
		}
		return *current.name, true
	case hxrt__TypeEnumValue:
		if current.name == nil {
			return "", false
		}
		return *current.name, true
	case string:
		return current, true
	case *string:
		if current == nil {
			return "", false
		}
		return *current, true
	default:
		return "", false
	}
}

func hxrt_typeCreateClassInstance(className string, args []any) (any, bool) {
	switch className {
	case "haxe._Int64.___Int64":
		return hxrt_typeCallAny(New_haxe___Int64_____Int64, args)
	default:
		return nil, false
	}
}

func hxrt_typeCreateClassEmptyInstance(className string) (any, bool) {
	switch className {
	case "haxe._Int64.___Int64":
		return &haxe___Int64_____Int64{}, true
	default:
		return nil, false
	}
}

func hxrt_typeCreateEnumInstance(enumName string, constructorName string, constructorIndex int, useIndex bool, args []any) (any, bool) {
	switch enumName {
	case "ValueType":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TNull, true
			case 1:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TInt, true
			case 2:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TFloat, true
			case 3:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TBool, true
			case 4:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TObject, true
			case 5:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TFunction, true
			case 6:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(ValueType_TClass, args)
			case 7:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(ValueType_TEnum, args)
			case 8:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TUnknown, true
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "TNull":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TNull, true
		case "TInt":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TInt, true
		case "TFloat":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TFloat, true
		case "TBool":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TBool, true
		case "TObject":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TObject, true
		case "TFunction":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TFunction, true
		case "TClass":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(ValueType_TClass, args)
		case "TEnum":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(ValueType_TEnum, args)
		case "TUnknown":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TUnknown, true
		default:
			return nil, false
		}
	default:
		return nil, false
	}
}

func Type_getClass(o any) any {
	if hxrt.AnyEqualsNull(o) {
		return nil
	}
	switch value := o.(type) {
	case *hxrt__TypeClassValue:
		if value == nil {
			return nil
		}
		return value
	case hxrt__TypeClassValue:
		copyValue := value
		return &copyValue
	case *haxe___Int64_____Int64:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe._Int64.___Int64")}
	default:
		return nil
	}
}

func Type_getEnum(o any) any {
	if hxrt.AnyEqualsNull(o) {
		return nil
	}
	switch value := o.(type) {
	case *hxrt__TypeEnumValue:
		if value == nil {
			return nil
		}
		return value
	case hxrt__TypeEnumValue:
		copyValue := value
		return &copyValue
	case *ValueType:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("ValueType")}
	default:
		return nil
	}
}

func Type_getSuperClass(c any) any {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return nil
	}
	switch className {
	case "haxe._Int64.___Int64":
		return nil
	default:
		return nil
	}
}

func Type_getClassName(c any) *string {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return nil
	}
	return hxrt.StringFromLiteral(className)
}

func Type_getClassFields(c any) []*string {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return []*string{}
	}
	switch className {
	case "haxe._Int64.___Int64":
		return []*string{}
	default:
		return []*string{}
	}
}

func Type_getInstanceFields(c any) []*string {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return []*string{}
	}
	switch className {
	case "haxe._Int64.___Int64":
		return []*string{hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low")}
	default:
		return []*string{}
	}
}

func Type_getEnumName(e any) *string {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	return hxrt.StringFromLiteral(enumName)
}

func Type_resolveClass(name *string) any {
	if name == nil {
		return nil
	}
	rawName := *hxrt.StdString(name)
	switch rawName {
	case "haxe._Int64.___Int64":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	default:
		return nil
	}
}

func Type_resolveEnum(name *string) any {
	if name == nil {
		return nil
	}
	rawName := *hxrt.StdString(name)
	switch rawName {
	case "ValueType":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	default:
		return nil
	}
}

func Type_createInstance(cl any, args []any) any {
	className, ok := hxrt_typeResolvedClassName(cl)
	if !ok {
		return nil
	}
	instance, ok := hxrt_typeCreateClassInstance(className, args)
	if !ok {
		return nil
	}
	return instance
}

func Type_createEmptyInstance(cl any) any {
	className, ok := hxrt_typeResolvedClassName(cl)
	if !ok {
		return nil
	}
	instance, ok := hxrt_typeCreateClassEmptyInstance(className)
	if !ok {
		return nil
	}
	return instance
}

func Type_createEnum(e any, constr *string, params []any) any {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	constructorName := ""
	if constr != nil {
		constructorName = *hxrt.StdString(constr)
	}
	enumValue, ok := hxrt_typeCreateEnumInstance(enumName, constructorName, 0, false, params)
	if !ok {
		return nil
	}
	return enumValue
}

func Type_createEnumIndex(e any, index int, params []any) any {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	enumValue, ok := hxrt_typeCreateEnumInstance(enumName, "", index, true, params)
	if !ok {
		return nil
	}
	return enumValue
}

func Type_enumConstructor(e any) *string {
	if hxrt.AnyEqualsNull(e) {
		return nil
	}
	switch value := e.(type) {
	case *ValueType:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("TNull")
		case 1:
			return hxrt.StringFromLiteral("TInt")
		case 2:
			return hxrt.StringFromLiteral("TFloat")
		case 3:
			return hxrt.StringFromLiteral("TBool")
		case 4:
			return hxrt.StringFromLiteral("TObject")
		case 5:
			return hxrt.StringFromLiteral("TFunction")
		case 6:
			return hxrt.StringFromLiteral("TClass")
		case 7:
			return hxrt.StringFromLiteral("TEnum")
		case 8:
			return hxrt.StringFromLiteral("TUnknown")
		default:
			return nil
		}
	default:
		return nil
	}
}

func Type_enumIndex(e any) int {
	if hxrt.AnyEqualsNull(e) {
		return -1
	}
	switch value := e.(type) {
	case *ValueType:
		if value == nil {
			return -1
		}
		return value.tag
	default:
		return -1
	}
}

func Type_getEnumConstructs(e any) []*string {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return []*string{}
	}
	switch enumName {
	case "ValueType":
		return []*string{hxrt.StringFromLiteral("TNull"), hxrt.StringFromLiteral("TInt"), hxrt.StringFromLiteral("TFloat"), hxrt.StringFromLiteral("TBool"), hxrt.StringFromLiteral("TObject"), hxrt.StringFromLiteral("TFunction"), hxrt.StringFromLiteral("TClass"), hxrt.StringFromLiteral("TEnum"), hxrt.StringFromLiteral("TUnknown")}
	default:
		return []*string{}
	}
}

func Type_enumParameters(e any) []any {
	if hxrt.AnyEqualsNull(e) {
		return []any{}
	}
	switch value := e.(type) {
	case *ValueType:
		if value == nil || value.params == nil {
			return []any{}
		}
		out := make([]any, len(value.params))
		copy(out, value.params)
		return out
	default:
		return []any{}
	}
}

func Type_allEnums(e any) []any {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return []any{}
	}
	switch enumName {
	case "ValueType":
		return []any{ValueType_TNull, ValueType_TInt, ValueType_TFloat, ValueType_TBool, ValueType_TObject, ValueType_TFunction, ValueType_TUnknown}
	default:
		return []any{}
	}
}

func Type_typeof(v any) any {
	if hxrt.AnyEqualsNull(v) {
		return ValueType_TNull
	}
	if enumValue := Type_getEnum(v); enumValue != nil {
		return ValueType_TEnum(enumValue)
	}
	if classValue := Type_getClass(v); classValue != nil {
		return ValueType_TClass(classValue)
	}
	switch v.(type) {
	case bool:
		return ValueType_TBool
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr:
		return ValueType_TInt
	case float32, float64:
		return ValueType_TFloat
	case string, *string:
		return ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("String")})
	}
	ref := reflect.ValueOf(v)
	if !ref.IsValid() {
		return ValueType_TNull
	}
	switch ref.Kind() {
	case reflect.Func:
		return ValueType_TFunction
	case reflect.Slice, reflect.Array:
		return ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Array")})
	case reflect.Map, reflect.Struct, reflect.Interface, reflect.Pointer:
		return ValueType_TObject
	default:
		return ValueType_TUnknown
	}
}

func Type_enumEq(a any, b any) bool {
	return reflect.DeepEqual(a, b)
}
