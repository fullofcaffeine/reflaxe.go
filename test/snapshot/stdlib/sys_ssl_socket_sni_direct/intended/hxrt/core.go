package hxrt

import "math"

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

func IntFromNullableAny(value any) int {
	if value == nil {
		panic("Invalid operation: null")
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
		panic("Invalid operation: expected Int-compatible value")
	}
}
