package hxrt

// Ordered comparisons (`<`, `<=`, `>`, `>=`) whose operands include a Haxe
// `Null<Int>` or `Null<Float>` value held in Go `any` storage.
//
// What: Each helper evaluates one Haxe comparison operator for one operand
// shape: a nullable Int with an Int (in either order), two nullable Ints, or
// two nullable numbers compared as Floats.
// Why: Haxe's `--interp` and neko give `false` when exactly one operand is null
// and raise no error. When both operands are null, `<=` and `>=` give `true`
// and `<` and `>` give `false`, as for two equal numbers. Go cannot order `any`
// values, and `IntFromNullableAny` raises an error on null.
// How: The generated code passes the operands as arguments, so each runs once
// and in source order. A null operand gives `false`; two null operands read as
// two equal zeros. Other values are read with `IntFromNullableAny` (or as
// `float64`) and compared with the Go operator. The exported helpers are small
// enough for the Go compiler to inline; only the shared pair readers are
// called out of line. Comparisons whose operands cannot be null do not use
// these helpers.

// NullableIntLess reports `left < right` for a Haxe `Null<Int>` and an Int.
func NullableIntLess(left any, right int) bool {
	if left == nil {
		return false
	}
	return IntFromNullableAny(left) < right
}

// NullableIntLessEqual reports `left <= right` for a Haxe `Null<Int>` and an Int.
func NullableIntLessEqual(left any, right int) bool {
	if left == nil {
		return false
	}
	return IntFromNullableAny(left) <= right
}

// NullableIntGreater reports `left > right` for a Haxe `Null<Int>` and an Int.
func NullableIntGreater(left any, right int) bool {
	if left == nil {
		return false
	}
	return IntFromNullableAny(left) > right
}

// NullableIntGreaterEqual reports `left >= right` for a Haxe `Null<Int>` and an Int.
func NullableIntGreaterEqual(left any, right int) bool {
	if left == nil {
		return false
	}
	return IntFromNullableAny(left) >= right
}

// IntLessNullable reports `left < right` for an Int and a Haxe `Null<Int>`.
func IntLessNullable(left int, right any) bool {
	if right == nil {
		return false
	}
	return left < IntFromNullableAny(right)
}

// IntLessEqualNullable reports `left <= right` for an Int and a Haxe `Null<Int>`.
func IntLessEqualNullable(left int, right any) bool {
	if right == nil {
		return false
	}
	return left <= IntFromNullableAny(right)
}

// IntGreaterNullable reports `left > right` for an Int and a Haxe `Null<Int>`.
func IntGreaterNullable(left int, right any) bool {
	if right == nil {
		return false
	}
	return left > IntFromNullableAny(right)
}

// IntGreaterEqualNullable reports `left >= right` for an Int and a Haxe `Null<Int>`.
func IntGreaterEqualNullable(left int, right any) bool {
	if right == nil {
		return false
	}
	return left >= IntFromNullableAny(right)
}

// nullableIntPair reads two `Null<Int>` operands. ok is false when exactly one
// operand is null; two null operands read as (0, 0, true), so `<=` and `>=`
// hold and `<` and `>` do not.
func nullableIntPair(left any, right any) (int, int, bool) {
	if left == nil || right == nil {
		return 0, 0, left == nil && right == nil
	}
	return IntFromNullableAny(left), IntFromNullableAny(right), true
}

// NullableIntsLess reports `left < right` for two Haxe `Null<Int>` values.
func NullableIntsLess(left any, right any) bool {
	l, r, ok := nullableIntPair(left, right)
	return ok && l < r
}

// NullableIntsLessEqual reports `left <= right` for two Haxe `Null<Int>` values.
func NullableIntsLessEqual(left any, right any) bool {
	l, r, ok := nullableIntPair(left, right)
	return ok && l <= r
}

// NullableIntsGreater reports `left > right` for two Haxe `Null<Int>` values.
func NullableIntsGreater(left any, right any) bool {
	l, r, ok := nullableIntPair(left, right)
	return ok && l > r
}

// NullableIntsGreaterEqual reports `left >= right` for two Haxe `Null<Int>` values.
func NullableIntsGreaterEqual(left any, right any) bool {
	l, r, ok := nullableIntPair(left, right)
	return ok && l >= r
}

// nullableNumber reads a non-null `Null<Float>` (stored as float64) or
// `Null<Int>` value as a Float.
func nullableNumber(value any) float64 {
	if number, ok := value.(float64); ok {
		return number
	}
	return float64(IntFromNullableAny(value))
}

// nullableFloatPair reads two `Null<Float>` or `Null<Int>` operands as Floats,
// with the null rules of nullableIntPair. A NaN operand stays NaN, so every
// comparison with it is false.
func nullableFloatPair(left any, right any) (float64, float64, bool) {
	if left == nil || right == nil {
		return 0, 0, left == nil && right == nil
	}
	return nullableNumber(left), nullableNumber(right), true
}

// NullableFloatsLess reports `left < right` for two nullable numbers compared
// as Floats (`Null<Float>`, or `Null<Int>` next to `Null<Float>`).
func NullableFloatsLess(left any, right any) bool {
	l, r, ok := nullableFloatPair(left, right)
	return ok && l < r
}

// NullableFloatsLessEqual reports `left <= right` for two nullable numbers
// compared as Floats.
func NullableFloatsLessEqual(left any, right any) bool {
	l, r, ok := nullableFloatPair(left, right)
	return ok && l <= r
}

// NullableFloatsGreater reports `left > right` for two nullable numbers
// compared as Floats.
func NullableFloatsGreater(left any, right any) bool {
	l, r, ok := nullableFloatPair(left, right)
	return ok && l > r
}

// NullableFloatsGreaterEqual reports `left >= right` for two nullable numbers
// compared as Floats.
func NullableFloatsGreaterEqual(left any, right any) bool {
	l, r, ok := nullableFloatPair(left, right)
	return ok && l >= r
}
