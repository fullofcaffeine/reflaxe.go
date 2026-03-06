package hxrt

func BytesFromString(value *string) []int {
	if value == nil {
		return []int{}
	}

	raw := []byte(*value)
	out := make([]int, len(raw))
	for i := 0; i < len(raw); i++ {
		out[i] = int(raw[i])
	}
	return out
}

func BytesToString(values []int) *string {
	raw := make([]byte, len(values))
	for i := 0; i < len(values); i++ {
		raw[i] = byte(values[i])
	}
	return StringFromLiteral(string(raw))
}

func BytesClone(values []int) []int {
	if len(values) == 0 {
		return []int{}
	}
	out := make([]int, len(values))
	copy(out, values)
	return out
}
