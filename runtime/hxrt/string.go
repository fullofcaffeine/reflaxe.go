package hxrt

import (
	"fmt"
	"reflect"
)

func isAnyNil(value any) bool {
	if value == nil {
		return true
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
	default:
		return StringFromLiteral(fmt.Sprint(v))
	}
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

func StringEqualStringPtr(left *string, right *string) bool {
	return stringValueOrNullToken(left) == stringValueOrNullToken(right)
}

func StringLength(value any) int {
	runes := []rune(*StdString(value))
	return len(runes)
}

func StringCharAt(value any, index int) *string {
	runes := []rune(*StdString(value))
	if index < 0 || index >= len(runes) {
		return StringFromLiteral("")
	}
	return StringFromLiteral(string(runes[index]))
}

func StringCharCodeAt(value any, index int) int {
	runes := []rune(*StdString(value))
	if index < 0 || index >= len(runes) {
		return -1
	}
	return int(runes[index])
}

func StringCharCodeAtAny(value any, index int) any {
	code := StringCharCodeAt(value, index)
	if code < 0 {
		return nil
	}
	return code
}

func StringSubstring(value any, start int, end int) *string {
	runes := []rune(*StdString(value))
	if start < 0 {
		start = 0
	}
	if end < 0 {
		end = 0
	}
	if start > len(runes) {
		start = len(runes)
	}
	if end > len(runes) {
		end = len(runes)
	}
	if end < start {
		end = start
	}
	return StringFromLiteral(string(runes[start:end]))
}
