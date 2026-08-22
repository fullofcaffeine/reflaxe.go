package hxrt

import "strings"

// StringCompareStringPtr provides ordered Haxe comparison for string carriers.
//
// What: Return a negative, zero, or positive result for two nullable strings.
// Why: Generated strings are pointers, which Go cannot order directly, and null
// must remain distinct from the literal text "null".
// How: Order null before present values, then delegate present UTF-8 values to
// Go's lexical string comparison.
func StringCompareStringPtr(left *string, right *string) int {
	if left == nil {
		if right == nil {
			return 0
		}
		return -1
	}
	if right == nil {
		return 1
	}
	return strings.Compare(*left, *right)
}
