package main

import "snapshot/hxrt"

func _UnicodeString__UnicodeString_Impl___new(value *string) any {
	var this1 *string
	this1 = value
	return any(this1)
}

func _UnicodeString__UnicodeString_Impl__addString(a any, b *string) any {
	return hxrt.StringConcatStringPtr(func(hx_value_1 any) *string {
		if hx_value_1 == nil {
			var hx_zero_2 *string
			return hx_zero_2
		}
		return hx_value_1.(*string)
	}(a), b)
}

func _UnicodeString__UnicodeString_Impl__addUnicode(a any, b any) any {
	return hxrt.StringConcatStringPtr(func(hx_value_3 any) *string {
		if hx_value_3 == nil {
			var hx_zero_4 *string
			return hx_zero_4
		}
		return hx_value_3.(*string)
	}(a), func(hx_value_5 any) *string {
		if hx_value_5 == nil {
			var hx_zero_6 *string
			return hx_zero_6
		}
		return hx_value_5.(*string)
	}(b))
}

func _UnicodeString__UnicodeString_Impl__charAt(this1 *string, index int) *string {
	if (index < 0) || (index >= hxrt.StringLengthStringPtr(func(hx_value_7 any) *string {
		if hx_value_7 == nil {
			var hx_zero_8 *string
			return hx_zero_8
		}
		return hx_value_7.(*string)
	}(this1))) {
		return hxrt.StringFromLiteral("")
	}
	return hxrt.StdString(hxrt.StringSliceCodePointsStringPtr(this1, index, int((hxrt.Int32Wrap(index) + hxrt.Int32Wrap(1)))))
}

func _UnicodeString__UnicodeString_Impl__charCodeAt(this1 *string, index int) any {
	if (index < 0) || (index >= hxrt.StringLengthStringPtr(func(hx_value_9 any) *string {
		if hx_value_9 == nil {
			var hx_zero_10 *string
			return hx_zero_10
		}
		return hx_value_9.(*string)
	}(this1))) {
		return nil
	}
	return hxrt.StringCharCodeAtStringPtr(this1, index)
}

func _UnicodeString__UnicodeString_Impl__compare(a any, b any) int {
	left := hxrt.StdString(a)
	right := hxrt.StdString(b)
	leftLength := hxrt.StringLengthStringPtr(left)
	rightLength := hxrt.StringLengthStringPtr(right)
	var hx_if_11 int
	if leftLength < rightLength {
		hx_if_11 = leftLength
	} else {
		hx_if_11 = rightLength
	}
	limit := hx_if_11
	index := 0
	for index < limit {
		leftCode := hxrt.StringCharCodeAtStringPtr(left, index)
		rightCode := hxrt.StringCharCodeAtStringPtr(right, index)
		if leftCode < rightCode {
			return -1
		}
		if leftCode > rightCode {
			return 1
		}
		index = int(int32((index + 1)))
	}
	var hx_if_13 int
	if leftLength < rightLength {
		hx_if_13 = -1
	} else {
		var hx_if_12 int
		if leftLength > rightLength {
			hx_if_12 = 1
		} else {
			hx_if_12 = 0
		}
		hx_if_13 = hx_if_12
	}
	return hx_if_13
}

func _UnicodeString__UnicodeString_Impl__equal(a any, b any) bool {
	return hxrt.StringEqualStringPtr(func(hx_value_14 any) *string {
		if hx_value_14 == nil {
			var hx_zero_15 *string
			return hx_zero_15
		}
		return hx_value_14.(*string)
	}(a), func(hx_value_16 any) *string {
		if hx_value_16 == nil {
			var hx_zero_17 *string
			return hx_zero_17
		}
		return hx_value_16.(*string)
	}(b))
}

func _UnicodeString__UnicodeString_Impl__get_length(this1 *string) int {
	return hxrt.StringLengthStringPtr(this1)
}

func _UnicodeString__UnicodeString_Impl__greaterThan(a any, b any) bool {
	return (_UnicodeString__UnicodeString_Impl__compare(a, b) > 0)
}

func _UnicodeString__UnicodeString_Impl__greaterThanOrEqual(a any, b any) bool {
	return (_UnicodeString__UnicodeString_Impl__compare(a, b) >= 0)
}

func _UnicodeString__UnicodeString_Impl__indexOf(this1 *string, str *string, startIndex any) int {
	total := hxrt.StringLengthStringPtr(func(hx_value_18 any) *string {
		if hx_value_18 == nil {
			var hx_zero_19 *string
			return hx_zero_19
		}
		return hx_value_18.(*string)
	}(this1))
	var hx_if_20 int
	if startIndex == nil {
		hx_if_20 = 0
	} else {
		hx_if_20 = startIndex.(int)
	}
	start := hx_if_20
	if start < 0 {
		start = int((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(start)))
		if start < 0 {
			start = 0
		}
	}
	if start > total {
		start = total
	}
	needleLength := hxrt.StringLengthStringPtr(str)
	if needleLength == 0 {
		return start
	}
	if needleLength > int((hxrt.Int32Wrap(total) - hxrt.Int32Wrap(start))) {
		return -1
	}
	candidate := start
	lastCandidate := int((hxrt.Int32Wrap(total) - hxrt.Int32Wrap(needleLength)))
	for candidate <= lastCandidate {
		if _UnicodeString__UnicodeString_Impl__matchesAt(hxrt.StdString(this1), str, needleLength, candidate) {
			return candidate
		}
		candidate = int(int32((candidate + 1)))
	}
	return -1
}

func _UnicodeString__UnicodeString_Impl__iterator(this1 *string) *haxe__iterators__StringIteratorUnicode {
	return New_haxe__iterators__StringIteratorUnicode(this1)
}

func _UnicodeString__UnicodeString_Impl__keyValueIterator(this1 *string) *haxe__iterators__StringKeyValueIteratorUnicode {
	return New_haxe__iterators__StringKeyValueIteratorUnicode(this1)
}

func _UnicodeString__UnicodeString_Impl__lastIndexOf(this1 *string, str *string, startIndex any) int {
	total := hxrt.StringLengthStringPtr(func(hx_value_21 any) *string {
		if hx_value_21 == nil {
			var hx_zero_22 *string
			return hx_zero_22
		}
		return hx_value_21.(*string)
	}(this1))
	var hx_if_23 int
	if startIndex == nil {
		hx_if_23 = total
	} else {
		hx_if_23 = startIndex.(int)
	}
	start := hx_if_23
	if start < 0 {
		return -1
	}
	if start > total {
		start = total
	}
	needleLength := hxrt.StringLengthStringPtr(str)
	if needleLength == 0 {
		return start
	}
	if needleLength > total {
		return -1
	}
	candidate := start
	lastCandidate := int((hxrt.Int32Wrap(total) - hxrt.Int32Wrap(needleLength)))
	if candidate > lastCandidate {
		candidate = lastCandidate
	}
	for candidate >= 0 {
		if _UnicodeString__UnicodeString_Impl__matchesAt(hxrt.StdString(this1), str, needleLength, candidate) {
			return candidate
		}
		candidate = int(int32((candidate - 1)))
	}
	return -1
}

var _UnicodeString__UnicodeString_Impl__length int

func _UnicodeString__UnicodeString_Impl__lessThan(a any, b any) bool {
	return (_UnicodeString__UnicodeString_Impl__compare(a, b) < 0)
}

func _UnicodeString__UnicodeString_Impl__lessThanOrEqual(a any, b any) bool {
	return (_UnicodeString__UnicodeString_Impl__compare(a, b) <= 0)
}

func _UnicodeString__UnicodeString_Impl__matchesAt(this1 *string, str *string, needleLength int, candidate int) bool {
	offset := 0
	for offset < needleLength {
		if hxrt.StringCharCodeAtStringPtr(this1, int((hxrt.Int32Wrap(candidate)+hxrt.Int32Wrap(offset)))) != hxrt.StringCharCodeAtStringPtr(str, offset) {
			return false
		}
		offset = int(int32((offset + 1)))
	}
	return true
}

func _UnicodeString__UnicodeString_Impl__notEqual(a any, b any) bool {
	return !hxrt.StringEqualStringPtr(func(hx_value_24 any) *string {
		if hx_value_24 == nil {
			var hx_zero_25 *string
			return hx_zero_25
		}
		return hx_value_24.(*string)
	}(a), func(hx_value_26 any) *string {
		if hx_value_26 == nil {
			var hx_zero_27 *string
			return hx_zero_27
		}
		return hx_value_26.(*string)
	}(b))
}

func _UnicodeString__UnicodeString_Impl__substr(this1 *string, pos int, len any) *string {
	total := hxrt.StringLengthStringPtr(func(hx_value_28 any) *string {
		if hx_value_28 == nil {
			var hx_zero_29 *string
			return hx_zero_29
		}
		return hx_value_28.(*string)
	}(this1))
	start := pos
	if start < 0 {
		start = int((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(start)))
		if start < 0 {
			start = 0
		}
	}
	if start > total {
		return hxrt.StringFromLiteral("")
	}
	end := total
	if len != nil {
		var hx_if_30 int
		if len.(int) < 0 {
			hx_if_30 = int((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(len.(int))))
		} else {
			hx_if_30 = int((hxrt.Int32Wrap(start) + hxrt.Int32Wrap(len.(int))))
		}
		end = hx_if_30
		if end > total {
			end = total
		}
		if end <= start {
			return hxrt.StringFromLiteral("")
		}
	}
	return hxrt.StdString(hxrt.StringSliceCodePointsStringPtr(this1, start, end))
}

func _UnicodeString__UnicodeString_Impl__substring(this1 *string, startIndex int, endIndex any) *string {
	total := hxrt.StringLengthStringPtr(func(hx_value_31 any) *string {
		if hx_value_31 == nil {
			var hx_zero_32 *string
			return hx_zero_32
		}
		return hx_value_31.(*string)
	}(this1))
	var hx_if_33 int
	if startIndex < 0 {
		hx_if_33 = 0
	} else {
		hx_if_33 = startIndex
	}
	start := hx_if_33
	end := total
	if endIndex != nil {
		var hx_if_34 int
		if endIndex.(int) < 0 {
			hx_if_34 = 0
		} else {
			hx_if_34 = endIndex.(int)
		}
		end = hx_if_34
		if start == end {
			return hxrt.StringFromLiteral("")
		}
		if start > end {
			swap := start
			start = end
			end = swap
		}
	}
	if start > total {
		return hxrt.StringFromLiteral("")
	}
	if end > total {
		end = total
	}
	return hxrt.StdString(hxrt.StringSliceCodePointsStringPtr(this1, start, end))
}

func _UnicodeString__UnicodeString_Impl__validate(bytes *haxe__io__Bytes, encoding *haxe__io__Encoding) bool {
	switch encoding.tag {
	case 0:
		data := bytes.__hx_this.getData()
		pos := 0
		max := bytes.length
		for pos < max {
			hx_post_35 := pos
			pos = int(int32((pos + 1)))
			pos_1 := hx_post_35
			c := data[pos_1]
			if c < 128 {
			} else {
				if c < 194 {
					return false
				} else {
					if c < 224 {
						if int((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(1))) > max {
							return false
						}
						hx_post_36 := pos
						pos = int(int32((pos + 1)))
						pos_2 := hx_post_36
						c2 := data[pos_2]
						if (c2 < 128) || (c2 > 191) {
							return false
						}
					} else {
						if c < 240 {
							if int((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(2))) > max {
								return false
							}
							hx_post_37 := pos
							pos = int(int32((pos + 1)))
							pos_3 := hx_post_37
							c2_1 := data[pos_3]
							if c == 224 {
								if (c2_1 < 160) || (c2_1 > 191) {
									return false
								}
							} else {
								if (c2_1 < 128) || (c2_1 > 191) {
									return false
								}
							}
							hx_post_38 := pos
							pos = int(int32((pos + 1)))
							pos_4 := hx_post_38
							c3 := data[pos_4]
							if (c3 < 128) || (c3 > 191) {
								return false
							}
							c = int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(c) << uint(16)))) | hxrt.Int32Wrap(int((hxrt.Int32Wrap(c2_1) << uint(8))))))) | hxrt.Int32Wrap(c3)))
							if (15573120 <= c) && (c <= 15581119) {
								return false
							}
						} else {
							if c > 244 {
								return false
							} else {
								if int((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(3))) > max {
									return false
								}
								hx_post_39 := pos
								pos = int(int32((pos + 1)))
								pos_5 := hx_post_39
								c2_2 := data[pos_5]
								if c == 240 {
									if (c2_2 < 144) || (c2_2 > 191) {
										return false
									}
								} else {
									if c == 244 {
										if (c2_2 < 128) || (c2_2 > 143) {
											return false
										}
									} else {
										if (c2_2 < 128) || (c2_2 > 191) {
											return false
										}
									}
								}
								hx_post_40 := pos
								pos = int(int32((pos + 1)))
								pos_6 := hx_post_40
								c3_1 := data[pos_6]
								if (c3_1 < 128) || (c3_1 > 191) {
									return false
								}
								hx_post_41 := pos
								pos = int(int32((pos + 1)))
								pos_7 := hx_post_41
								c4 := data[pos_7]
								if (c4 < 128) || (c4 > 191) {
									return false
								}
							}
						}
					}
				}
			}
		}
		return true
	case 1:
		hxrt.Throw(hxrt.StringFromLiteral("UnicodeString.validate: RawNative encoding is not supported"))
	}
	return false
}
