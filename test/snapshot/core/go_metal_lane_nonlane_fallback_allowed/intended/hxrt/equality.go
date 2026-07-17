package hxrt

import "reflect"

// haxeNumericValue normalizes one erased Haxe number for equality.
//
// What: Recognize the Go integer and floating-point carriers that can reach an
// erased Haxe value.
// Why: Haxe `==` compares numeric values even when generic erasure leaves the two
// operands in different Go number types.
// How: Convert admitted carriers to the shared float64 range used by portable
// Haxe Int and Float values and report whether the input was numeric.
func haxeNumericValue(value any) (float64, bool) {
	switch number := value.(type) {
	case int:
		return float64(number), true
	case int8:
		return float64(number), true
	case int16:
		return float64(number), true
	case int32:
		return float64(number), true
	case int64:
		return float64(number), true
	case uint:
		return float64(number), true
	case uint8:
		return float64(number), true
	case uint16:
		return float64(number), true
	case uint32:
		return float64(number), true
	case uint64:
		return float64(number), true
	case uintptr:
		return float64(number), true
	case float32:
		return float64(number), true
	case float64:
		return number, true
	default:
		return 0, false
	}
}

// referenceEqual compares erased reference-shaped carriers by identity.
//
// What: Compare pointers, maps, slices, functions, and channels without inspecting
// their contents.
// Why: Haxe reference equality must not become deep equality, while applying Go
// interface `==` directly to a non-comparable carrier would panic.
// How: Keep reflection in this selectively copied helper, reject unlike dynamic
// types, preserve nil behavior, and compare only each reference carrier's pointer.
func referenceEqual(left any, right any) bool {
	leftValue := reflect.ValueOf(left)
	rightValue := reflect.ValueOf(right)
	if !leftValue.IsValid() || !rightValue.IsValid() {
		return !leftValue.IsValid() && !rightValue.IsValid()
	}
	if leftValue.Type() != rightValue.Type() {
		return false
	}

	switch leftValue.Kind() {
	case reflect.Interface:
		if leftValue.IsNil() || rightValue.IsNil() {
			return leftValue.IsNil() && rightValue.IsNil()
		}
		return referenceEqual(leftValue.Elem().Interface(), rightValue.Elem().Interface())
	case reflect.Ptr, reflect.UnsafePointer, reflect.Map, reflect.Slice, reflect.Func, reflect.Chan:
		if leftValue.IsNil() || rightValue.IsNil() {
			return leftValue.IsNil() && rightValue.IsNil()
		}
		return leftValue.Pointer() == rightValue.Pointer()
	default:
		if leftValue.Type().Comparable() {
			return left == right
		}
		return false
	}
}

// HaxeEqual preserves portable Haxe `==` behavior at interface-backed boundaries.
//
// What: Compare values stored as Go any because of generic erasure, nullable
// primitive storage, or another non-comparable carrier.
// Why: Direct Go interface equality panics for non-comparable carriers and
// compares pointer-backed Haxe strings by identity instead of contents.
// How: Normalize nil and numeric values, compare strings by contents, and delegate
// reference-shaped carriers to the identity-only reflection island above.
func HaxeEqual(left any, right any) bool {
	leftNil := isAnyNil(left)
	rightNil := isAnyNil(right)
	if leftNil || rightNil {
		return leftNil && rightNil
	}

	if leftNumber, ok := haxeNumericValue(left); ok {
		rightNumber, rightOK := haxeNumericValue(right)
		return rightOK && leftNumber == rightNumber
	}

	switch leftValue := left.(type) {
	case *string:
		rightValue, ok := right.(*string)
		return ok && *leftValue == *rightValue
	case string:
		rightValue, ok := right.(string)
		return ok && leftValue == rightValue
	case bool:
		rightValue, ok := right.(bool)
		return ok && leftValue == rightValue
	default:
		return referenceEqual(left, right)
	}
}
