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

// Why:
// - These byte-oriented helpers used to live directly in GoCompiler as large raw Go blocks.
// - They only operate on plain []int buffers and do not need compiler context.
//
// What:
//   - Move the pure hex and buffer helper algorithms into hxrt so compiler-emitted haxe.io.Bytes
//     wrappers can stay thin.
//
// How:
//   - Keep behavior identical to the old compiler-owned helpers, including odd-length errors and
//     the permissive nibble mapping used by Bytes.ofHex today.
func BytesOfHex(value *string) []int {
	raw := *StdString(value)
	lenValue := len(raw)
	if (lenValue & 1) != 0 {
		Throw(StringFromLiteral("Not a hex string (odd number of digits)"))
		return []int{}
	}

	out := make([]int, lenValue>>1)
	for i := 0; i < len(out); i++ {
		high := int(raw[i*2])
		low := int(raw[i*2+1])
		high = (high & 0xF) + (((high & 0x40) >> 6) * 9)
		low = (low & 0xF) + (((low & 0x40) >> 6) * 9)
		out[i] = ((high << 4) | low) & 0xFF
	}
	return out
}

func BytesToHex(values []int, length int) *string {
	if length <= 0 {
		return StringFromLiteral("")
	}
	if length > len(values) {
		length = len(values)
	}

	hexChars := "0123456789abcdef"
	out := make([]byte, length*2)
	for i := 0; i < length; i++ {
		c := values[i] & 0xFF
		out[i*2] = hexChars[c>>4]
		out[i*2+1] = hexChars[c&15]
	}
	return StringFromLiteral(string(out))
}

func BytesBufferAddByte(buffer []int, value int) []int {
	return append(buffer, value&0xFF)
}

func BytesBufferAdd(buffer []int, src []int) []int {
	if len(src) == 0 {
		return buffer
	}
	return append(buffer, src...)
}

func BytesBufferAddSlice(buffer []int, src []int, pos int, length int) []int {
	if length == 0 {
		return buffer
	}
	return append(buffer, src[pos:pos+length]...)
}

func BytesBufferLength(buffer []int) int {
	return len(buffer)
}
