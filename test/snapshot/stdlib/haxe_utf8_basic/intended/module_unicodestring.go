package main

import "snapshot/hxrt"

func _UnicodeString__UnicodeString_Impl___new(value *string) *string {
	var this1 *string
	this1 = value
	return this1
}

func _UnicodeString__UnicodeString_Impl__addString(a *string, b *string) *string {
	return hxrt.StringConcatStringPtr(a, b)
}

func _UnicodeString__UnicodeString_Impl__addUnicode(a *string, b *string) *string {
	return hxrt.StringConcatStringPtr(a, b)
}

func _UnicodeString__UnicodeString_Impl__charAt(this1 *string, index int) *string {
	if (index < 0) || (index >= hxrt.StringLengthStringPtr(this1)) {
		return hxrt.StringFromLiteral("")
	}
	return hxrt.StdString(hxrt.StringSliceCodePointsStringPtr(this1, index, int((hxrt.Int32Wrap(index) + hxrt.Int32Wrap(1)))))
}

func _UnicodeString__UnicodeString_Impl__charCodeAt(this1 *string, index int) any {
	if (index < 0) || (index >= hxrt.StringLengthStringPtr(this1)) {
		return nil
	}
	return hxrt.StringCharCodeAtStringPtr(this1, index)
}

func _UnicodeString__UnicodeString_Impl__compare(a *string, b *string) int {
	left := a
	right := b
	leftLength := hxrt.StringLengthStringPtr(left)
	rightLength := hxrt.StringLengthStringPtr(right)
	var hx_if_1 int
	if leftLength < rightLength {
		hx_if_1 = leftLength
	} else {
		hx_if_1 = rightLength
	}
	limit := hx_if_1
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
	var hx_if_3 int
	if leftLength < rightLength {
		hx_if_3 = -1
	} else {
		var hx_if_2 int
		if leftLength > rightLength {
			hx_if_2 = 1
		} else {
			hx_if_2 = 0
		}
		hx_if_3 = hx_if_2
	}
	return hx_if_3
}

func _UnicodeString__UnicodeString_Impl__equal(a *string, b *string) bool {
	return hxrt.StringEqualStringPtr(a, b)
}

func _UnicodeString__UnicodeString_Impl__get_length(this1 *string) int {
	return hxrt.StringLengthStringPtr(this1)
}

func _UnicodeString__UnicodeString_Impl__greaterThan(a *string, b *string) bool {
	return (_UnicodeString__UnicodeString_Impl__compare(a, b) > 0)
}

func _UnicodeString__UnicodeString_Impl__greaterThanOrEqual(a *string, b *string) bool {
	return (_UnicodeString__UnicodeString_Impl__compare(a, b) >= 0)
}

func _UnicodeString__UnicodeString_Impl__indexOf(this1 *string, str *string, startIndex any) int {
	total := hxrt.StringLengthStringPtr(this1)
	var hx_if_4 int
	if startIndex == nil {
		hx_if_4 = 0
	} else {
		hx_if_4 = startIndex.(int)
	}
	start := hx_if_4
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
		if _UnicodeString__UnicodeString_Impl__matchesAt(this1, str, needleLength, candidate) {
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
	total := hxrt.StringLengthStringPtr(this1)
	var hx_if_5 int
	if startIndex == nil {
		hx_if_5 = total
	} else {
		hx_if_5 = startIndex.(int)
	}
	start := hx_if_5
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
		if _UnicodeString__UnicodeString_Impl__matchesAt(this1, str, needleLength, candidate) {
			return candidate
		}
		candidate = int(int32((candidate - 1)))
	}
	return -1
}

var _UnicodeString__UnicodeString_Impl__length int

func _UnicodeString__UnicodeString_Impl__lessThan(a *string, b *string) bool {
	return (_UnicodeString__UnicodeString_Impl__compare(a, b) < 0)
}

func _UnicodeString__UnicodeString_Impl__lessThanOrEqual(a *string, b *string) bool {
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

func _UnicodeString__UnicodeString_Impl__notEqual(a *string, b *string) bool {
	return !hxrt.StringEqualStringPtr(a, b)
}

func _UnicodeString__UnicodeString_Impl__substr(this1 *string, pos int, len any) *string {
	total := hxrt.StringLengthStringPtr(this1)
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
		var hx_if_6 int
		if len.(int) < 0 {
			hx_if_6 = int((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(len.(int))))
		} else {
			hx_if_6 = int((hxrt.Int32Wrap(start) + hxrt.Int32Wrap(len.(int))))
		}
		end = hx_if_6
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
	total := hxrt.StringLengthStringPtr(this1)
	var hx_if_7 int
	if startIndex < 0 {
		hx_if_7 = 0
	} else {
		hx_if_7 = startIndex
	}
	start := hx_if_7
	end := total
	if endIndex != nil {
		var hx_if_8 int
		if endIndex.(int) < 0 {
			hx_if_8 = 0
		} else {
			hx_if_8 = endIndex.(int)
		}
		end = hx_if_8
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
			hx_post_9 := pos
			pos = int(int32((pos + 1)))
			pos_1 := hx_post_9
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
						hx_post_10 := pos
						pos = int(int32((pos + 1)))
						pos_2 := hx_post_10
						c2 := data[pos_2]
						if (c2 < 128) || (c2 > 191) {
							return false
						}
					} else {
						if c < 240 {
							if int((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(2))) > max {
								return false
							}
							hx_post_11 := pos
							pos = int(int32((pos + 1)))
							pos_3 := hx_post_11
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
							hx_post_12 := pos
							pos = int(int32((pos + 1)))
							pos_4 := hx_post_12
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
								hx_post_13 := pos
								pos = int(int32((pos + 1)))
								pos_5 := hx_post_13
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
								hx_post_14 := pos
								pos = int(int32((pos + 1)))
								pos_6 := hx_post_14
								c3_1 := data[pos_6]
								if (c3_1 < 128) || (c3_1 > 191) {
									return false
								}
								hx_post_15 := pos
								pos = int(int32((pos + 1)))
								pos_7 := hx_post_15
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
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("UnicodeString.validate: RawNative encoding is not supported")))
	}
	return false
}
