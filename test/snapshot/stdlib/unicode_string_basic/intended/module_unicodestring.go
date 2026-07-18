package main

import "snapshot/hxrt"

func _UnicodeString__UnicodeString_Impl___new(value *string) any {
	var this1 *string
	this1 = value
	return any(this1)
}

func _UnicodeString__UnicodeString_Impl__addString(a any, b *string) any {
	return hxrt.StringConcatStringPtr(func(hx_value_57 any) *string {
		if hx_value_57 == nil {
			var hx_zero_58 *string
			return hx_zero_58
		}
		return hx_value_57.(*string)
	}(a), b)
}

func _UnicodeString__UnicodeString_Impl__addUnicode(a any, b any) any {
	return hxrt.StringConcatStringPtr(func(hx_value_59 any) *string {
		if hx_value_59 == nil {
			var hx_zero_60 *string
			return hx_zero_60
		}
		return hx_value_59.(*string)
	}(a), func(hx_value_61 any) *string {
		if hx_value_61 == nil {
			var hx_zero_62 *string
			return hx_zero_62
		}
		return hx_value_61.(*string)
	}(b))
}

func _UnicodeString__UnicodeString_Impl__charAt(this1 *string, index int) *string {
	if (index < 0) || (index >= hxrt.StringLengthStringPtr(func(hx_value_63 any) *string {
		if hx_value_63 == nil {
			var hx_zero_64 *string
			return hx_zero_64
		}
		return hx_value_63.(*string)
	}(this1))) {
		return hxrt.StringFromLiteral("")
	}
	return hxrt.StdString(hxrt.StringSliceCodePointsStringPtr(this1, index, int(int32((hxrt.Int32Wrap(index) + hxrt.Int32Wrap(1))))))
}

func _UnicodeString__UnicodeString_Impl__charCodeAt(this1 *string, index int) any {
	if (index < 0) || (index >= hxrt.StringLengthStringPtr(func(hx_value_65 any) *string {
		if hx_value_65 == nil {
			var hx_zero_66 *string
			return hx_zero_66
		}
		return hx_value_65.(*string)
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
	var hx_if_67 int
	if leftLength < rightLength {
		hx_if_67 = leftLength
	} else {
		hx_if_67 = rightLength
	}
	limit := hx_if_67
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
	var hx_if_69 int
	if leftLength < rightLength {
		hx_if_69 = -1
	} else {
		var hx_if_68 int
		if leftLength > rightLength {
			hx_if_68 = 1
		} else {
			hx_if_68 = 0
		}
		hx_if_69 = hx_if_68
	}
	return hx_if_69
}

func _UnicodeString__UnicodeString_Impl__equal(a any, b any) bool {
	return hxrt.StringEqualStringPtr(func(hx_value_70 any) *string {
		if hx_value_70 == nil {
			var hx_zero_71 *string
			return hx_zero_71
		}
		return hx_value_70.(*string)
	}(a), func(hx_value_72 any) *string {
		if hx_value_72 == nil {
			var hx_zero_73 *string
			return hx_zero_73
		}
		return hx_value_72.(*string)
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
	total := hxrt.StringLengthStringPtr(func(hx_value_74 any) *string {
		if hx_value_74 == nil {
			var hx_zero_75 *string
			return hx_zero_75
		}
		return hx_value_74.(*string)
	}(this1))
	var hx_if_76 int
	if startIndex == nil {
		hx_if_76 = 0
	} else {
		hx_if_76 = startIndex.(int)
	}
	start := hx_if_76
	if start < 0 {
		start = int(int32((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(start))))
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
	if needleLength > int(int32((hxrt.Int32Wrap(total) - hxrt.Int32Wrap(start)))) {
		return -1
	}
	candidate := start
	lastCandidate := int(int32((hxrt.Int32Wrap(total) - hxrt.Int32Wrap(needleLength))))
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
	total := hxrt.StringLengthStringPtr(func(hx_value_77 any) *string {
		if hx_value_77 == nil {
			var hx_zero_78 *string
			return hx_zero_78
		}
		return hx_value_77.(*string)
	}(this1))
	var hx_if_79 int
	if startIndex == nil {
		hx_if_79 = total
	} else {
		hx_if_79 = startIndex.(int)
	}
	start := hx_if_79
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
	lastCandidate := int(int32((hxrt.Int32Wrap(total) - hxrt.Int32Wrap(needleLength))))
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
		if hxrt.StringCharCodeAtStringPtr(this1, int(int32((hxrt.Int32Wrap(candidate)+hxrt.Int32Wrap(offset))))) != hxrt.StringCharCodeAtStringPtr(str, offset) {
			return false
		}
		offset = int(int32((offset + 1)))
	}
	return true
}

func _UnicodeString__UnicodeString_Impl__notEqual(a any, b any) bool {
	return !hxrt.StringEqualStringPtr(func(hx_value_80 any) *string {
		if hx_value_80 == nil {
			var hx_zero_81 *string
			return hx_zero_81
		}
		return hx_value_80.(*string)
	}(a), func(hx_value_82 any) *string {
		if hx_value_82 == nil {
			var hx_zero_83 *string
			return hx_zero_83
		}
		return hx_value_82.(*string)
	}(b))
}

func _UnicodeString__UnicodeString_Impl__substr(this1 *string, pos int, len any) *string {
	total := hxrt.StringLengthStringPtr(func(hx_value_84 any) *string {
		if hx_value_84 == nil {
			var hx_zero_85 *string
			return hx_zero_85
		}
		return hx_value_84.(*string)
	}(this1))
	start := pos
	if start < 0 {
		start = int(int32((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(start))))
		if start < 0 {
			start = 0
		}
	}
	if start > total {
		return hxrt.StringFromLiteral("")
	}
	end := total
	if len != nil {
		var hx_if_86 int
		if len.(int) < 0 {
			hx_if_86 = int(int32((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(len.(int)))))
		} else {
			hx_if_86 = int(int32((hxrt.Int32Wrap(start) + hxrt.Int32Wrap(len.(int)))))
		}
		end = hx_if_86
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
	total := hxrt.StringLengthStringPtr(func(hx_value_87 any) *string {
		if hx_value_87 == nil {
			var hx_zero_88 *string
			return hx_zero_88
		}
		return hx_value_87.(*string)
	}(this1))
	var hx_if_89 int
	if startIndex < 0 {
		hx_if_89 = 0
	} else {
		hx_if_89 = startIndex
	}
	start := hx_if_89
	end := total
	if endIndex != nil {
		var hx_if_90 int
		if endIndex.(int) < 0 {
			hx_if_90 = 0
		} else {
			hx_if_90 = endIndex.(int)
		}
		end = hx_if_90
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
			hx_post_91 := pos
			pos = int(int32((pos + 1)))
			pos_1 := hx_post_91
			c := data[pos_1]
			if c < 128 {
			} else {
				if c < 194 {
					return false
				} else {
					if c < 224 {
						if int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(1)))) > max {
							return false
						}
						hx_post_92 := pos
						pos = int(int32((pos + 1)))
						pos_2 := hx_post_92
						c2 := data[pos_2]
						if (c2 < 128) || (c2 > 191) {
							return false
						}
					} else {
						if c < 240 {
							if int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(2)))) > max {
								return false
							}
							hx_post_93 := pos
							pos = int(int32((pos + 1)))
							pos_3 := hx_post_93
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
							hx_post_94 := pos
							pos = int(int32((pos + 1)))
							pos_4 := hx_post_94
							c3 := data[pos_4]
							if (c3 < 128) || (c3 > 191) {
								return false
							}
							c = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(c) << uint(16))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(c2_1) << uint(8))))))))) | hxrt.Int32Wrap(c3))))
							if (15573120 <= c) && (c <= 15581119) {
								return false
							}
						} else {
							if c > 244 {
								return false
							} else {
								if int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(3)))) > max {
									return false
								}
								hx_post_95 := pos
								pos = int(int32((pos + 1)))
								pos_5 := hx_post_95
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
								hx_post_96 := pos
								pos = int(int32((pos + 1)))
								pos_6 := hx_post_96
								c3_1 := data[pos_6]
								if (c3_1 < 128) || (c3_1 > 191) {
									return false
								}
								hx_post_97 := pos
								pos = int(int32((pos + 1)))
								pos_7 := hx_post_97
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
