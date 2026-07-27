package hxrt

import (
	"strconv"
)

// SerializationParseFloat performs the host numeric conversion after staged
// Unserializer has selected and bounded one floating-point token.
func SerializationParseFloat(value *string) float64 {
	parsed, err := strconv.ParseFloat(*StdString(value), 64)
	if err != nil {
		Throw(err)
		return 0
	}
	return parsed
}
