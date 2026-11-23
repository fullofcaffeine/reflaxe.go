package hxrt

// String methods whose optional Int argument is a Haxe `Null<Int>` value held
// in Go `any` storage: `substr(pos, ?len)`, `substring(startIndex, ?endIndex)`,
// `indexOf(str, ?startIndex)` and `lastIndexOf(str, ?startIndex)`.
//
// What: Each helper calls the matching String helper in string.go and reads a
// null optional argument as omitted.
// Why: Haxe (`--interp`, neko and python) reads `s.substr(1, n)` with a null
// `n` as `s.substr(1)`. The string.go helpers take that argument as a Go
// `int`, and a Go type assertion on a nil `any` panics.
// How: The generated code passes the receiver and the arguments as arguments,
// so each runs once and in source order. A nil argument takes the omitted
// path. Any other value is read with `IntFromNullableAny`. Arguments that
// cannot be null (Int values and constants) do not use these helpers; they
// call the string.go helpers directly. The `any`-receiver variants serve the
// generic string lowering, as `StringSubstr` does in string.go.

// StringSubstrOptionalStringPtr is `value.substr(pos, length)` for a
// `Null<Int>` length.
func StringSubstrOptionalStringPtr(value *string, pos int, length any) *string {
	if length == nil {
		return StringSubstrStringPtr(value, pos, 0, false)
	}
	return StringSubstrStringPtr(value, pos, IntFromNullableAny(length), true)
}

// StringSubstringOptionalStringPtr is `value.substring(start, end)` for a
// `Null<Int>` end. A null end is the length of `value`.
func StringSubstringOptionalStringPtr(value *string, start int, end any) *string {
	if end == nil {
		return StringSubstringStringPtr(value, start, StringLengthStringPtr(value))
	}
	return StringSubstringStringPtr(value, start, IntFromNullableAny(end))
}

// StringIndexOfOptionalStringPtr is `value.indexOf(search, startIndex)` for a
// `Null<Int>` start index.
func StringIndexOfOptionalStringPtr(value *string, search *string, startIndex any) int {
	if startIndex == nil {
		return StringIndexOfStringPtr(value, search, 0, false)
	}
	return StringIndexOfStringPtr(value, search, IntFromNullableAny(startIndex), true)
}

// StringLastIndexOfOptionalStringPtr is `value.lastIndexOf(search,
// startIndex)` for a `Null<Int>` start index.
func StringLastIndexOfOptionalStringPtr(value *string, search *string, startIndex any) int {
	if startIndex == nil {
		return StringLastIndexOfStringPtr(value, search, 0, false)
	}
	return StringLastIndexOfStringPtr(value, search, IntFromNullableAny(startIndex), true)
}

// StringSubstrOptional is StringSubstrOptionalStringPtr for an `any` receiver.
func StringSubstrOptional(value any, pos int, length any) *string {
	return StringSubstrOptionalStringPtr(StdString(value), pos, length)
}

// StringSubstringOptional is StringSubstringOptionalStringPtr for an `any` receiver.
func StringSubstringOptional(value any, start int, end any) *string {
	return StringSubstringOptionalStringPtr(StdString(value), start, end)
}

// StringIndexOfOptional is StringIndexOfOptionalStringPtr for `any` operands.
func StringIndexOfOptional(value any, search any, startIndex any) int {
	return StringIndexOfOptionalStringPtr(StdString(value), StdString(search), startIndex)
}

// StringLastIndexOfOptional is StringLastIndexOfOptionalStringPtr for `any` operands.
func StringLastIndexOfOptional(value any, search any, startIndex any) int {
	return StringLastIndexOfOptionalStringPtr(StdString(value), StdString(search), startIndex)
}
