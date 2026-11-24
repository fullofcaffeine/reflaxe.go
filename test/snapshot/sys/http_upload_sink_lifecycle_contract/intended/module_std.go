package main

import (
	"math/rand"
	"snapshot/hxrt"
)

func Std_digitValue(code int, hexadecimal bool) int {
	if (code >= 48) && (code <= 57) {
		return int((hxrt.Int32Wrap(code) - hxrt.Int32Wrap(48)))
	}
	if (hexadecimal && (code >= 97)) && (code <= 102) {
		return int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(code) - hxrt.Int32Wrap(97)))) + hxrt.Int32Wrap(10)))
	}
	if (hexadecimal && (code >= 65)) && (code <= 70) {
		return int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(code) - hxrt.Int32Wrap(65)))) + hxrt.Int32Wrap(10)))
	}
	return -1
}

func Std_downcast(value any, c any) any {
	var hx_if_1 any
	if hx_is_of_type_token(value, c) {
		hx_if_1 = value
	} else {
		hx_if_1 = nil
	}
	return hx_if_1
}

func Std_instance(value any, c any) any {
	return func() any {
		var hx_if_2 any
		if hx_is_of_type_token(value, c) {
			hx_if_2 = value
		} else {
			hx_if_2 = nil
		}
		return hx_if_2
	}()
}

func Std_int(x float64) int {
	return hxrt.MathTruncInt(x)
}

func Std_invalidFloat() float64 {
	return hxrt.StringParseFloatExact(hxrt.StringFromLiteral(""))
}

func Std_is(v any, t any) bool {
	return hx_is_of_type_token(v, t)
}

func Std_isDecimalDigit(code int) bool {
	return ((code >= 48) && (code <= 57))
}

func Std_isSpaceCode(code int) bool {
	return (((code >= 9) && (code <= 13)) || (code == 32))
}

func Std_parseFloat(x *string) float64 {
	if hxrt.StringEqualStringPtr(x, nil) {
		return hxrt.StringParseFloatExact(hxrt.StringFromLiteral(""))
	}
	index := 0
	length := hxrt.StringLengthStringPtr(x)
	for (index < length) && func() bool {
		code := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(x, index))
		return (((code >= 9) && (code <= 13)) || (code == 32))
	}() {
		index = int(int32((index + 1)))
	}
	start := index
	if index < length {
		var sign any = hxrt.StringCharCodeAtAnyStringPtr(x, index)
		if (sign == 45) || (sign == 43) {
			index = int(int32((index + 1)))
		}
	}
	digits := 0
	for (index < length) && func() bool {
		code_1 := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(x, index))
		return ((code_1 >= 48) && (code_1 <= 57))
	}() {
		digits = int(int32((digits + 1)))
		index = int(int32((index + 1)))
	}
	if (index < length) && (hxrt.StringCharCodeAtAnyStringPtr(x, index) == 46) {
		index = int(int32((index + 1)))
		for (index < length) && func() bool {
			code_2 := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(x, index))
			return ((code_2 >= 48) && (code_2 <= 57))
		}() {
			digits = int(int32((digits + 1)))
			index = int(int32((index + 1)))
		}
	}
	if digits == 0 {
		return hxrt.StringParseFloatExact(hxrt.StringFromLiteral(""))
	}
	if (index < length) && ((hxrt.StringCharCodeAtAnyStringPtr(x, index) == 101) || (hxrt.StringCharCodeAtAnyStringPtr(x, index) == 69)) {
		index = int(int32((index + 1)))
		if (index < length) && ((hxrt.StringCharCodeAtAnyStringPtr(x, index) == 45) || (hxrt.StringCharCodeAtAnyStringPtr(x, index) == 43)) {
			index = int(int32((index + 1)))
		}
		exponentStart := index
		for (index < length) && func() bool {
			code_3 := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(x, index))
			return ((code_3 >= 48) && (code_3 <= 57))
		}() {
			index = int(int32((index + 1)))
		}
		if index == exponentStart {
			return hxrt.StringParseFloatExact(hxrt.StringFromLiteral(""))
		}
	}
	return hxrt.StringParseFloatExact(hxrt.StringSubstringStringPtr(x, start, index))
}

func Std_parseInt(x *string) any {
	if hxrt.StringEqualStringPtr(x, nil) {
		return nil
	}
	index := 0
	length := hxrt.StringLengthStringPtr(x)
	for (index < length) && func() bool {
		code := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(x, index))
		return (((code >= 9) && (code <= 13)) || (code == 32))
	}() {
		index = int(int32((index + 1)))
	}
	negative := false
	if index < length {
		var sign any = hxrt.StringCharCodeAtAnyStringPtr(x, index)
		if (sign == 45) || (sign == 43) {
			negative = (sign == 45)
			index = int(int32((index + 1)))
		}
	}
	hexadecimal := (((int((hxrt.Int32Wrap(index) + hxrt.Int32Wrap(1))) < length) && (hxrt.StringCharCodeAtAnyStringPtr(x, index) == 48)) && ((hxrt.StringCharCodeAtAnyStringPtr(x, int((hxrt.Int32Wrap(index)+hxrt.Int32Wrap(1)))) == 120) || (hxrt.StringCharCodeAtAnyStringPtr(x, int((hxrt.Int32Wrap(index)+hxrt.Int32Wrap(1)))) == 88)))
	if hexadecimal {
		index = int((hxrt.Int32Wrap(index) + hxrt.Int32Wrap(2)))
	}
	digitStart := index
	if hexadecimal {
		value := 0
		significantDigits := 0
		sawNonZero := false
		for index < length {
			code_1 := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(x, index))
			var hx_if_5 int
			if (code_1 >= 48) && (code_1 <= 57) {
				hx_if_5 = int((hxrt.Int32Wrap(code_1) - hxrt.Int32Wrap(48)))
			} else {
				var hx_if_4 int
				if (code_1 >= 97) && (code_1 <= 102) {
					hx_if_4 = int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(code_1) - hxrt.Int32Wrap(97)))) + hxrt.Int32Wrap(10)))
				} else {
					var hx_if_3 int
					if (code_1 >= 65) && (code_1 <= 70) {
						hx_if_3 = int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(code_1) - hxrt.Int32Wrap(65)))) + hxrt.Int32Wrap(10)))
					} else {
						hx_if_3 = -1
					}
					hx_if_4 = hx_if_3
				}
				hx_if_5 = hx_if_4
			}
			digit := hx_if_5
			if digit < 0 {
				break
			}
			if (digit != 0) || sawNonZero {
				sawNonZero = true
				significantDigits = int(int32((significantDigits + 1)))
				if significantDigits > 8 {
					return nil
				}
			}
			value = int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(value) << uint(4)))) | hxrt.Int32Wrap(digit)))
			index = int(int32((index + 1)))
		}
		if index == digitStart {
			return nil
		}
		var hx_if_6 any
		if negative {
			hx_if_6 = int(-int32(value))
		} else {
			hx_if_6 = value
		}
		return hx_if_6
	}
	result := 0
	for index < length {
		code_2 := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(x, index))
		var hx_if_7 int
		if (code_2 >= 48) && (code_2 <= 57) {
			hx_if_7 = int((hxrt.Int32Wrap(code_2) - hxrt.Int32Wrap(48)))
		} else {
			hx_if_7 = -1
		}
		digit_1 := hx_if_7
		if digit_1 < 0 {
			break
		}
		var hx_if_8 int
		if negative {
			hx_if_8 = 8
		} else {
			hx_if_8 = 7
		}
		lastAllowedDigit := hx_if_8
		if (result < -214748364) || ((result == -214748364) && (digit_1 > lastAllowedDigit)) {
			return nil
		}
		result = int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(result) * hxrt.Int32Wrap(10)))) - hxrt.Int32Wrap(digit_1)))
		index = int(int32((index + 1)))
	}
	if index == digitStart {
		return nil
	}
	var hx_if_9 any
	if negative {
		hx_if_9 = result
	} else {
		hx_if_9 = int(-int32(result))
	}
	return hx_if_9
}

func Std_random(x int) int {
	var hx_if_10 int
	if x <= 1 {
		hx_if_10 = 0
	} else {
		hx_if_10 = rand.Intn(x)
	}
	return hx_if_10
}
