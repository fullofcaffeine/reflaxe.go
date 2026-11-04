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
