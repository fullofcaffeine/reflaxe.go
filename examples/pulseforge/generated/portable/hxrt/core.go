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
