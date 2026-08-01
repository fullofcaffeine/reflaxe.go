package hxrt

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"
)

func isAnyNil(value any) bool {
	if value == nil {
		return true
	}

	switch value.(type) {
	case bool, string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64:
		return false
	}

	inner := reflect.ValueOf(value)
	if !inner.IsValid() {
		return true
	}

	switch inner.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return inner.IsNil()
	default:
		return false
	}
}

func StdString(value any) *string {
	if isAnyNil(value) {
		return StringFromLiteral("null")
	}

	switch v := value.(type) {
	case *string:
		return v
	case string:
		return StringFromLiteral(v)
	case bool:
		return StringFromLiteral(strconv.FormatBool(v))
	case int:
		return StringFromLiteral(strconv.FormatInt(int64(v), 10))
	case int8:
		return StringFromLiteral(strconv.FormatInt(int64(v), 10))
	case int16:
		return StringFromLiteral(strconv.FormatInt(int64(v), 10))
	case int32:
		return StringFromLiteral(strconv.FormatInt(int64(v), 10))
	case int64:
		return StringFromLiteral(strconv.FormatInt(v, 10))
	case uint:
		return StringFromLiteral(strconv.FormatUint(uint64(v), 10))
	case uint8:
		return StringFromLiteral(strconv.FormatUint(uint64(v), 10))
	case uint16:
		return StringFromLiteral(strconv.FormatUint(uint64(v), 10))
	case uint32:
		return StringFromLiteral(strconv.FormatUint(uint64(v), 10))
	case uint64:
		return StringFromLiteral(strconv.FormatUint(v, 10))
	case uintptr:
		return StringFromLiteral(strconv.FormatUint(uint64(v), 10))
	default:
		return StringFromLiteral(fmt.Sprint(v))
	}
}

// StringParseFloatExact converts one source-validated numeric token.
//
// What: Returns the Go float64 value for an exact decimal/exponent token.
// Why: Native IEEE-754 conversion is representation work, while staged Std owns
// whitespace, prefix scanning, malformed exponents, and trailing input policy.
// How: Parse the complete supplied token and map every native error to Haxe NaN.
func StringParseFloatExact(value *string) float64 {
	if value == nil {
		return math.NaN()
	}
	parsed, err := strconv.ParseFloat(*value, 64)
	if err != nil {
		return math.NaN()
	}
	return parsed
}

func StringSlice(values []*string) []string {
	out := make([]string, len(values))
	for i := 0; i < len(values); i++ {
		out[i] = *StdString(values[i])
	}
	return out
}

func StringConcatAny(left any, right any) *string {
	l := StdString(left)
	r := StdString(right)
	return StringFromLiteral(*l + *r)
}

func StringEqualAny(left any, right any) bool {
	l := StdString(left)
	r := StdString(right)
	return *l == *r
}

func AnyEqualsNull(value any) bool {
	return isAnyNil(value)
}

func stringValueOrNullToken(value *string) string {
	if value == nil {
		return "null"
	}
	return *value
}

func StringConcatStringPtr(left *string, right *string) *string {
	return StringFromLiteral(stringValueOrNullToken(left) + stringValueOrNullToken(right))
}

// StringEqualStringPtr preserves Haxe equality for pointer-backed strings.
//
// What: Compare nullable string carriers by value.
// Why: Haxe null is distinct from the literal string "null" even though string
// conversion and concatenation intentionally render both with the same text.
// How: Resolve nil identity before dereferencing two present string values.
func StringEqualStringPtr(left *string, right *string) bool {
	if left == nil || right == nil {
		return left == nil && right == nil
	}
	return *left == *right
}

func StringLengthStringPtr(value *string) int {
	return utf8.RuneCountInString(stringValueOrNullToken(value))
}

func StringCharAtStringPtr(value *string, index int) *string {
	if index < 0 {
		return StringFromLiteral("")
	}
	raw := stringValueOrNullToken(value)
	runeIndex := 0
	for _, runeValue := range raw {
		if runeIndex == index {
			return StringFromLiteral(string(runeValue))
		}
		runeIndex++
	}
	return StringFromLiteral("")
}

func StringCharCodeAtStringPtr(value *string, index int) int {
	if index < 0 {
		return -1
	}
	raw := stringValueOrNullToken(value)
	runeIndex := 0
	for _, runeValue := range raw {
		if runeIndex == index {
			return int(runeValue)
		}
		runeIndex++
	}
	return -1
}

// StringSliceCodePointsStringPtr exposes Go's rune-aware string representation
// to typed staged std code. Haxe owns bounds and range policy; this helper only
// converts an already-normalized half-open code-point range into a string.
func StringSliceCodePointsStringPtr(value *string, start int, end int) *string {
	return StringFromLiteral(sliceStringByRuneRange(stringValueOrNullToken(value), start, end))
}

func StringCharCodeAtAnyStringPtr(value *string, index int) any {
	code := StringCharCodeAtStringPtr(value, index)
	if code < 0 {
		return nil
	}
	return code
}

func StringSubstringStringPtr(value *string, start int, end int) *string {
	raw := stringValueOrNullToken(value)
	total := utf8.RuneCountInString(raw)
	clampedStart, clampedEnd := clampSubstringRange(start, end, total)
	return StringFromLiteral(sliceStringByRuneRange(raw, clampedStart, clampedEnd))
}

func StringSubstrStringPtr(value *string, pos int, length int, hasLength bool) *string {
	raw := stringValueOrNullToken(value)
	total := utf8.RuneCountInString(raw)

	start := pos
	if start < 0 {
		start = total + start
	}
	if start < 0 {
		start = 0
	}
	if start > total {
		start = total
	}

	end := total
	if hasLength {
		if length < 0 {
			end = total + length
		} else {
			end = start + length
		}
		if end > total {
			end = total
		}
	}
	if end < start {
		end = start
	}
	return StringFromLiteral(sliceStringByRuneRange(raw, start, end))
}

func StringLastIndexOfStringPtr(value *string, search *string, startIndex int, hasStart bool) int {
	raw := stringValueOrNullToken(value)
	needle := stringValueOrNullToken(search)
	total := utf8.RuneCountInString(raw)

	start := total
	if hasStart {
		start = startIndex
		if start < 0 {
			return -1
		}
		if start > total {
			start = total
		}
	}

	if needle == "" {
		return start
	}

	needleLen := utf8.RuneCountInString(needle)
	if needleLen > total || start < 0 {
		return -1
	}
	maxStart := start
	if maxStart > total-needleLen {
		maxStart = total - needleLen
	}
	for index := maxStart; index >= 0; index-- {
		if sliceStringByRuneRange(raw, index, index+needleLen) == needle {
			return index
		}
	}
	return -1
}

func StringSplitStringPtr(value *string, delimiter *string) []*string {
	raw := stringValueOrNullToken(value)
	sep := stringValueOrNullToken(delimiter)
	if raw == "" {
		return []*string{}
	}
	if sep == "" {
		out := make([]*string, 0, utf8.RuneCountInString(raw))
		for _, runeValue := range raw {
			out = append(out, StringFromLiteral(string(runeValue)))
		}
		return out
	}
	parts := strings.Split(raw, sep)
	out := make([]*string, 0, len(parts))
	for _, part := range parts {
		out = append(out, StringFromLiteral(part))
	}
	return out
}

func StringToLowerCaseStringPtr(value *string) *string {
	return StringFromLiteral(strings.ToLower(stringValueOrNullToken(value)))
}

func StringToUpperCaseStringPtr(value *string) *string {
	return StringFromLiteral(strings.ToUpper(stringValueOrNullToken(value)))
}

func StringJoinAny(values []any, delimiter *string) *string {
	sep := stringValueOrNullToken(delimiter)
	parts := make([]string, 0, len(values))
	for _, value := range values {
		parts = append(parts, *StdString(value))
	}
	return StringFromLiteral(strings.Join(parts, sep))
}

func StringFromCharCode(code int) *string {
	return StringFromLiteral(string(rune(code)))
}

func clampSubstringRange(start int, end int, total int) (int, int) {
	if start < 0 {
		start = 0
	}
	if end < 0 {
		end = 0
	}
	if start > total {
		start = total
	}
	if end > total {
		end = total
	}
	if end < start {
		end = start
	}
	return start, end
}

func sliceStringByRuneRange(raw string, start int, end int) string {
	if start >= end {
		return ""
	}

	runeIndex := 0
	byteIndex := 0
	startByte := -1
	endByte := -1
	for byteIndex < len(raw) {
		if runeIndex == start && startByte == -1 {
			startByte = byteIndex
		}
		if runeIndex == end {
			endByte = byteIndex
			break
		}
		_, size := utf8.DecodeRuneInString(raw[byteIndex:])
		byteIndex += size
		runeIndex++
	}

	if startByte == -1 {
		startByte = len(raw)
	}
	if endByte == -1 {
		if runeIndex == end {
			endByte = byteIndex
		} else {
			endByte = len(raw)
		}
	}
	if endByte < startByte {
		endByte = startByte
	}
	return raw[startByte:endByte]
}

func StringLength(value any) int {
	return StringLengthStringPtr(StdString(value))
}

func StringCharAt(value any, index int) *string {
	return StringCharAtStringPtr(StdString(value), index)
}

func StringCharCodeAt(value any, index int) int {
	return StringCharCodeAtStringPtr(StdString(value), index)
}

func StringCharCodeAtAny(value any, index int) any {
	return StringCharCodeAtAnyStringPtr(StdString(value), index)
}

func StringSubstring(value any, start int, end int) *string {
	return StringSubstringStringPtr(StdString(value), start, end)
}

func StringSubstr(value any, pos int, length int, hasLength bool) *string {
	return StringSubstrStringPtr(StdString(value), pos, length, hasLength)
}

func StringLastIndexOf(value any, search any, startIndex int, hasStart bool) int {
	return StringLastIndexOfStringPtr(StdString(value), StdString(search), startIndex, hasStart)
}

func StringSplit(value any, delimiter any) []*string {
	return StringSplitStringPtr(StdString(value), StdString(delimiter))
}

func StringToLowerCase(value any) *string {
	return StringToLowerCaseStringPtr(StdString(value))
}

func StringToUpperCase(value any) *string {
	return StringToUpperCaseStringPtr(StdString(value))
}
