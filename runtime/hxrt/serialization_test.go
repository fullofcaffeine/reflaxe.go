package hxrt

import "testing"

func TestSerializationParseFloatUsesBoundedHostConversion(t *testing.T) {
	for raw, expected := range map[string]float64{
		"0":       0,
		"-12.5":   -12.5,
		"6.25e+2": 625,
	} {
		value := raw
		if actual := SerializationParseFloat(&value); actual != expected {
			t.Fatalf("SerializationParseFloat(%q) = %v, want %v", raw, actual, expected)
		}
	}
}
