package hxrt

import "math"

// portableArrayFromValues is the narrow registration point used by optional
// runtime features that must construct a Haxe Array without making their Go file
// uncompilable when array.go is intentionally omitted. array.go replaces this
// fallback during package initialization; feature dependency planning guarantees
// that any program which calls such a feature also receives the Array carrier.
// The erased values remain localized to the runtime representation boundary.
var portableArrayFromValues = func(values []any) any {
	return values
}

// HaxeException is the foundational panic carrier shared by core validation
// failures and the optional exception helpers. Keeping the carrier in core
// lets selectively copied runtimes preserve Haxe failures without importing
// the broader catch/message surface.
type HaxeException struct {
	Value any
}

func Throw(value any) {
	panic(HaxeException{Value: value})
}

func StringFromLiteral(value string) *string {
	copy := value
	return &copy
}

func FloatMod(a float64, b float64) float64 {
	return math.Mod(a, b)
}

func Int32Wrap(value int) int32 {
	return int32(value)
}

// IsInt reports whether Std.isOfType(value, Int) holds.
//
// What: true for a Go integer and for a Float with an Int value (FloatIsInt),
// false for every other value, including null.
// Why: Haxe has no run-time Int for a Float such as 4.0, so `--interp`, neko,
// JS, Java, Python and PHP answer Std.isOfType(4.0, Int) with true
// (haxe_go-jxxd). Type.typeof still reports TFloat for it.
// How: A Go int, the common case, returns before the type switch over the
// other integer and float types.
func IsInt(value any) bool {
	if _, ok := value.(int); ok {
		return true
	}
	switch current := value.(type) {
	case float64:
		return FloatIsInt(current)
	case int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr:
		return true
	case float32:
		return FloatIsInt(float64(current))
	default:
		return false
	}
}

// FloatIsInt reports whether a Float has an Int value: integral and within
// Haxe's 32-bit Int range, as on `--interp`, JS and Python. -0.0 counts;
// 2147483648.0, NaN and the infinities do not.
func FloatIsInt(value float64) bool {
	return value >= math.MinInt32 && value <= math.MaxInt32 && value == math.Trunc(value)
}

// IntFromNullableAny is a portable dynamic-to-Int boundary. Conversion
// failures use Throw so generated Haxe try/catch can handle them; raw Go panics
// are reserved for failures originating in explicit native Go authority.
func IntFromNullableAny(value any) int {
	if value == nil {
		Throw(StringFromLiteral("Invalid operation: null"))
		return 0
	}
	switch v := value.(type) {
	case int:
		return v
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case uint:
		return int(v)
	case uint8:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	default:
		Throw(StringFromLiteral("Invalid operation: expected Int-compatible value"))
		return 0
	}
}
