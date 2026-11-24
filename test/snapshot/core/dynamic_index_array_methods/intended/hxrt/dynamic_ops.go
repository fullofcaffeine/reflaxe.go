package hxrt

import (
	"math"
	"strings"
)

// Operators on Haxe `Dynamic` values.
//
// What: Arithmetic, bitwise, shift, unary, truth and ordered-comparison
// helpers for operands that Haxe types as `Dynamic`. Each helper takes the
// operands as Go `any` and applies the rule of Haxe `--interp` to the value
// the operand holds at run time.
// Why: A `Dynamic` value is Go `any`, and Go defines no operators on `any`.
// Haxe decides at run time what `d + 1` means: Int arithmetic with 32-bit
// wrap-around when both values are Int, Float arithmetic when one value is a
// Float, string concatenation when one value is a String, and an "Invalid
// operation" error for any other value (null and Bool included).
// How: The generated code passes each operand once, in source order, as an
// argument. Int results are Go `int` values in the 32-bit range, as other
// generated code stores Haxe Int values; Float results are `float64`. The
// compiler converts a result to the static Haxe type of the expression. Typed
// operands without a `Dynamic` side never reach these helpers; they keep the
// Go operators. The rules match `--interp` and neko:
//   - `+`, `-` and `*` of two Int values wrap to 32 bits; with a Float value
//     they are Float operations. `/` is always a Float division.
//   - `%` of two Int values is the truncated remainder, and NaN for a zero
//     divisor; with a Float value it is `math.Mod`.
//   - `&`, `|`, `^`, `<<`, `>>` and `>>>` need two Int values (a Float value
//     raises the error even when it is whole); a shift uses the low five bits
//     of its count.
//   - Unary `-` negates an Int (with wrap-around) or a Float; `~x` is `-1 - x`,
//     so it also accepts a Float.
//   - A condition is true only for the Bool `true`; `!x` is true only for
//     `false` and null.
//   - An ordered comparison orders numbers by value, Strings by their text and
//     Bools as `false < true`. Two values that Haxe `==` calls equal (null and
//     null, or one object) compare equal; any other pair is unordered, so
//     every ordered comparison of it is false.
//   - `target[index]` reads and writes an Array through the Array carrier;
//     a null target raises "Null Access" and any other target raises
//     "Unexpected value <value>, expected Array". So this slice also copies
//     the Array carrier (array.go).
//   - A value stored in a Float, Bool or String slot (`var f:Float = d`)
//     is converted to the Go type of the slot; see DynamicToFloatSlot.
//   - A field of an element read from `target[index]` (`d[0].length`)
//     raises "Null Access" when the element is null; see
//     DynamicFieldReceiver.

// dynamicInt reads an Int operand. Go integer types other than `int` can reach
// a `Dynamic` value through native boundaries; they are read as their low 32
// bits, the Haxe Int range.
func dynamicInt(value any) (int32, bool) {
	switch number := value.(type) {
	case int:
		return int32(number), true
	case int8:
		return int32(number), true
	case int16:
		return int32(number), true
	case int32:
		return number, true
	case int64:
		return int32(number), true
	case uint:
		return int32(number), true
	case uint8:
		return int32(number), true
	case uint16:
		return int32(number), true
	case uint32:
		return int32(number), true
	case uint64:
		return int32(number), true
	default:
		return 0, false
	}
}

// dynamicFloat reads a Float operand.
func dynamicFloat(value any) (float64, bool) {
	switch number := value.(type) {
	case float64:
		return number, true
	case float32:
		return float64(number), true
	default:
		return 0, false
	}
}

// dynamicNumber reads an Int or Float operand as a Float.
func dynamicNumber(value any) (float64, bool) {
	if number, ok := dynamicInt(value); ok {
		return float64(number), true
	}
	return dynamicFloat(value)
}

// dynamicString reads a String operand.
func dynamicString(value any) (string, bool) {
	switch text := value.(type) {
	case *string:
		if text == nil {
			return "", false
		}
		return *text, true
	case string:
		return text, true
	default:
		return "", false
	}
}

// dynamicOperationError raises the error of `--interp` for an operator that
// does not accept its operand values, for example `Invalid operation: true + 1`.
func dynamicOperationError(left any, operator string, right any) {
	Throw(StringFromLiteral("Invalid operation: " + *StdString(left) + " " + operator + " " + *StdString(right)))
}

// dynamicArithmetic applies a `+`, `-` or `*` whose operands are numbers:
// `ints` for two Int values (Go int32 arithmetic wraps as Haxe Int does) and
// `floats` otherwise. The result is false when an operand is not a number.
func dynamicArithmetic(left any, right any, ints func(int32, int32) int32, floats func(float64, float64) float64) (any, bool) {
	if l, ok := dynamicInt(left); ok {
		if r, ok := dynamicInt(right); ok {
			return int(ints(l, r)), true
		}
	}
	l, leftOK := dynamicNumber(left)
	r, rightOK := dynamicNumber(right)
	if !leftOK || !rightOK {
		return nil, false
	}
	return floats(l, r), true
}

// DynamicAdd is `left + right`: a number sum, or string concatenation when
// either value is a String.
func DynamicAdd(left any, right any) any {
	if sum, ok := dynamicArithmetic(left, right, func(l int32, r int32) int32 { return l + r }, func(l float64, r float64) float64 { return l + r }); ok {
		return sum
	}
	_, leftString := dynamicString(left)
	_, rightString := dynamicString(right)
	if leftString || rightString {
		return StringConcatAny(left, right)
	}
	dynamicOperationError(left, "+", right)
	return nil
}

// DynamicSub is `left - right`.
func DynamicSub(left any, right any) any {
	if difference, ok := dynamicArithmetic(left, right, func(l int32, r int32) int32 { return l - r }, func(l float64, r float64) float64 { return l - r }); ok {
		return difference
	}
	dynamicOperationError(left, "-", right)
	return nil
}

// DynamicMul is `left * right`.
func DynamicMul(left any, right any) any {
	if product, ok := dynamicArithmetic(left, right, func(l int32, r int32) int32 { return l * r }, func(l float64, r float64) float64 { return l * r }); ok {
		return product
	}
	dynamicOperationError(left, "*", right)
	return nil
}

// DynamicDiv is `left / right`, a Float division for Int values too.
func DynamicDiv(left any, right any) float64 {
	l, leftOK := dynamicNumber(left)
	r, rightOK := dynamicNumber(right)
	if !leftOK || !rightOK {
		dynamicOperationError(left, "/", right)
	}
	return l / r
}

// DynamicMod is `left % right`. Two Int values give the remainder with the
// sign of `left`, or NaN for a zero `right`; a Float value gives `math.Mod`.
func DynamicMod(left any, right any) any {
	if l, ok := dynamicInt(left); ok {
		if r, ok := dynamicInt(right); ok {
			if r == 0 {
				return math.NaN()
			}
			return int(l % r)
		}
	}
	l, leftOK := dynamicNumber(left)
	r, rightOK := dynamicNumber(right)
	if !leftOK || !rightOK {
		dynamicOperationError(left, "%", right)
		return nil
	}
	return math.Mod(l, r)
}

// dynamicInts reads the two Int operands of a bitwise or shift operator.
func dynamicInts(left any, operator string, right any) (int32, int32) {
	l, leftOK := dynamicInt(left)
	r, rightOK := dynamicInt(right)
	if !leftOK || !rightOK {
		dynamicOperationError(left, operator, right)
	}
	return l, r
}

// DynamicAnd is `left & right`.
func DynamicAnd(left any, right any) int {
	l, r := dynamicInts(left, "&", right)
	return int(l & r)
}

// DynamicOr is `left | right`.
func DynamicOr(left any, right any) int {
	l, r := dynamicInts(left, "|", right)
	return int(l | r)
}

// DynamicXor is `left ^ right`.
func DynamicXor(left any, right any) int {
	l, r := dynamicInts(left, "^", right)
	return int(l ^ r)
}

// DynamicShl is `left << right`.
func DynamicShl(left any, right any) int {
	l, r := dynamicInts(left, "<<", right)
	return int(l << (uint32(r) & 31))
}

// DynamicShr is `left >> right`, a shift that keeps the sign.
func DynamicShr(left any, right any) int {
	l, r := dynamicInts(left, ">>", right)
	return int(l >> (uint32(r) & 31))
}

// DynamicUShr is `left >>> right`, a shift that fills with zero bits.
func DynamicUShr(left any, right any) int {
	l, r := dynamicInts(left, ">>>", right)
	return int(int32(uint32(l) >> (uint32(r) & 31)))
}

// DynamicNeg is `-value`.
func DynamicNeg(value any) any {
	if number, ok := dynamicInt(value); ok {
		return int(-number)
	}
	if number, ok := dynamicFloat(value); ok {
		return -number
	}
	Throw(StringFromLiteral("Invalid operation"))
	return nil
}

// DynamicComplement is `~value`, which `--interp` computes as `-1 - value`.
func DynamicComplement(value any) any {
	return DynamicSub(-1, value)
}

// DynamicIsTrue reads `value` as a condition: only the Bool `true` is true.
func DynamicIsTrue(value any) bool {
	truth, ok := value.(bool)
	return ok && truth
}

// DynamicNot is `!value`: true for `false` and null, and false for any other
// value, including Int 0 and the empty String.
func DynamicNot(value any) bool {
	if isAnyNil(value) {
		return true
	}
	truth, ok := value.(bool)
	return ok && !truth
}

// DynamicToFloat reads the result of a `Dynamic` operator that Haxe types as
// Float (`d - 1`, `-d`, `d++`). The value is an Int or a Float.
func DynamicToFloat(value any) float64 {
	if number, ok := dynamicNumber(value); ok {
		return number
	}
	if isAnyNil(value) {
		Throw(StringFromLiteral("Invalid operation: null"))
	}
	Throw(StringFromLiteral("Invalid operation: expected Float-compatible value"))
	return 0
}

// Slots of a static type that receive a `Dynamic` value.
//
// What: DynamicToFloatSlot, DynamicToNullableFloat, DynamicToBool and
// DynamicToNullableBool convert a `Dynamic` value for a Float,
// `Null<Float>`, Bool or `Null<Bool>` slot: a local, an argument, a return
// value, a field, an Array element or a Map value. DynamicToString
// (string.go) does the same for a String slot.
// Why: Haxe puts no check or conversion on `var f:Float = d`, so `--interp`
// keeps the value unchanged: an Int stays an Int, and a String in a Float
// slot stays a String until an operator rejects it. A Go `float64`, `bool`
// or `*string` slot cannot hold another kind of value, and a Go type
// assertion stops the program with a panic that Haxe `try`/`catch` cannot
// handle. So the helpers convert what the slot can represent and raise a
// Haxe error for the rest:
//   - An Int becomes a Float, as Haxe converts an Int typed value into a
//     Float slot.
//   - Null stays null in `Null<Float>`, `Null<Bool>` and String slots. In a
//     Float or Bool slot, which Go cannot set to null, it reads as 0 or
//     false, the default value of the type on the static Haxe targets.
//   - A String slot stores another value as its `Std.string` text, which is
//     what `--interp` prints for the kept value.
//   - Any other value raises "Invalid operation: expected Float-compatible
//     value" or "Invalid operation: expected Bool value".
// How: Null is checked with isAnyNil, so a null pointer stored in the `any`
// (a null String) also counts as null.

// DynamicToFloatSlot converts a `Dynamic` value for a Float slot.
func DynamicToFloatSlot(value any) float64 {
	if isAnyNil(value) {
		return 0
	}
	return DynamicToFloat(value)
}

// DynamicToNullableFloat converts a `Dynamic` value for a `Null<Float>`
// slot: null stays null, and a number is stored as a Go float64.
func DynamicToNullableFloat(value any) any {
	if isAnyNil(value) {
		return nil
	}
	return DynamicToFloat(value)
}

// DynamicToBool converts a `Dynamic` value for a Bool slot.
func DynamicToBool(value any) bool {
	if isAnyNil(value) {
		return false
	}
	truth, ok := value.(bool)
	if !ok {
		Throw(StringFromLiteral("Invalid operation: expected Bool value"))
	}
	return truth
}

// DynamicToNullableBool converts a `Dynamic` value for a `Null<Bool>` slot:
// null stays null.
func DynamicToNullableBool(value any) any {
	if isAnyNil(value) {
		return nil
	}
	return DynamicToBool(value)
}

// dynamicCompare orders two values. The second result is false when the
// pair is unordered.
func dynamicCompare(left any, right any) (int, bool) {
	if l, ok := dynamicNumber(left); ok {
		r, ok := dynamicNumber(right)
		switch {
		case !ok:
			return 0, false
		case l < r:
			return -1, true
		case l > r:
			return 1, true
		case l == r:
			return 0, true
		default:
			return 0, false
		}
	}
	if l, ok := dynamicString(left); ok {
		if r, ok := dynamicString(right); ok {
			return strings.Compare(l, r), true
		}
		return 0, false
	}
	if l, ok := left.(bool); ok {
		if r, ok := right.(bool); ok {
			switch {
			case l == r:
				return 0, true
			case r:
				return -1, true
			default:
				return 1, true
			}
		}
		return 0, false
	}
	if HaxeEqual(left, right) {
		return 0, true
	}
	return 0, false
}

// DynamicLess is `left < right`.
func DynamicLess(left any, right any) bool {
	order, ok := dynamicCompare(left, right)
	return ok && order < 0
}

// DynamicLessEqual is `left <= right`.
func DynamicLessEqual(left any, right any) bool {
	order, ok := dynamicCompare(left, right)
	return ok && order <= 0
}

// DynamicGreater is `left > right`.
func DynamicGreater(left any, right any) bool {
	order, ok := dynamicCompare(left, right)
	return ok && order > 0
}

// DynamicGreaterEqual is `left >= right`.
func DynamicGreaterEqual(left any, right any) bool {
	order, ok := dynamicCompare(left, right)
	return ok && order >= 0
}

// dynamicIndexArray reads the Array that `target[index]` indexes when Haxe
// types `target` as `Dynamic`. `--interp` indexes only an Array this way: a
// null target raises "Null Access", and any other value (a String, a Map, an
// anonymous object, a number) raises "Unexpected value <value>, expected
// Array".
func dynamicIndexArray(target any) *Array {
	if isAnyNil(target) {
		Throw(StringFromLiteral("Null Access"))
	}
	array, ok := target.(*Array)
	if !ok {
		Throw(StringFromLiteral("Unexpected value " + *StdString(target) + ", expected Array"))
	}
	return array
}

// DynamicIndexGet is `target[index]` on a `Dynamic` target. The compiler
// converts the index to an Int as it does for a typed Array, and the read is
// the Array carrier's: an index outside the Array reads null.
func DynamicIndexGet(target any, index int) any {
	return dynamicIndexArray(target).Get(index)
}

// DynamicIndexSet is `target[index] = value` on a `Dynamic` target and returns
// `value`. The write is the Array carrier's: an index past the end fills the
// gap with nulls, and a negative index changes nothing, as for a typed Array.
func DynamicIndexSet(target any, index int, value any) any {
	return dynamicIndexArray(target).Set(index, value)
}

// DynamicFieldReceiver is the receiver of a field read, write or call on an
// element read from a `Dynamic` index (`d[0].length`, `d[0].push(x)`), on a
// field of one, or on the result of a call of one.
//
// What: It returns `target` unchanged, and raises "Null Access" when
// `target` is null (also a null pointer stored in the `any`).
// Why: Haxe types such an element as a structure, so `--interp` raises an
// error for a field of a null element, as for any null structure. The
// compiler reads the field through `Reflect.field`, which reads null from a
// null receiver, so it checks the receiver here first.
// How: The compiler passes the receiver once, before the field read or the
// arguments of the call.
func DynamicFieldReceiver(target any) any {
	if isAnyNil(target) {
		Throw(StringFromLiteral("Null Access"))
	}
	return target
}
