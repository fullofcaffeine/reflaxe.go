package hxrt

import (
	"math"
	"unicode/utf16"
)

// ByteView is an opaque immutable []byte view cached by staged haxe.io.Bytes.
// Public bounds, encoding selection, and cache invalidation remain in Haxe.
type ByteView struct {
	raw []byte
}

// Float32FromBits and the related helpers expose Go's IEEE-754 bit
// reinterpretation as a narrow typed runtime capability. Portable haxe.io
// owns byte ordering and stream behavior; the runtime only performs the
// target-native reinterpretation that Haxe source cannot express directly.
func Float32FromBits(value int) float64 {
	return float64(math.Float32frombits(uint32(value)))
}

func Float32Bits(value float64) int {
	return int(int32(math.Float32bits(float32(value))))
}

func Float64FromWords(low int, high int) float64 {
	bits := uint64(uint32(low)) | uint64(uint32(high))<<32
	return math.Float64frombits(bits)
}

func Float64LowWord(value float64) int {
	return int(int32(uint32(math.Float64bits(value))))
}

func Float64HighWord(value float64) int {
	return int(int32(uint32(math.Float64bits(value) >> 32)))
}

func BytesAllocValues(length int) []int {
	if length < 0 {
		Throw(StringFromLiteral("Negative Bytes length"))
		return []int{}
	}
	return make([]int, length)
}

func BytesViewFromValues(values []int) *ByteView {
	raw := make([]byte, len(values))
	for index, value := range values {
		raw[index] = byte(value)
	}
	return &ByteView{raw: raw}
}

func BytesViewFromString(value *string, utf16LE bool) *ByteView {
	text := *StdString(value)
	if !utf16LE {
		return &ByteView{raw: []byte(text)}
	}
	units := utf16.Encode([]rune(text))
	raw := make([]byte, len(units)*2)
	for index, unit := range units {
		raw[index*2] = byte(unit)
		raw[index*2+1] = byte(unit >> 8)
	}
	return &ByteView{raw: raw}
}

func BytesValuesFromView(view *ByteView) []int {
	if view == nil || len(view.raw) == 0 {
		return []int{}
	}
	values := make([]int, len(view.raw))
	for index, value := range view.raw {
		values[index] = int(value)
	}
	return values
}

func BytesViewLength(view *ByteView) int {
	if view == nil {
		return 0
	}
	return len(view.raw)
}

func BytesViewMatchesValues(view *ByteView, values []int) bool {
	if view == nil || len(view.raw) != len(values) {
		return false
	}
	for index, value := range values {
		if view.raw[index] != byte(value) {
			return false
		}
	}
	return true
}

func BytesStringFromView(view *ByteView, pos int, length int, utf16LE bool) *string {
	if view == nil {
		return StringFromLiteral("")
	}
	raw := view.raw[pos : pos+length]
	if !utf16LE {
		return StringFromLiteral(string(raw))
	}
	limit := len(raw)
	if limit&1 != 0 {
		limit--
	}
	units := make([]uint16, limit/2)
	for index := range units {
		units[index] = uint16(raw[index*2]) | uint16(raw[index*2+1])<<8
	}
	return StringFromLiteral(string(utf16.Decode(units)))
}

func BytesClone(values []int) []int {
	if len(values) == 0 {
		return []int{}
	}
	out := make([]int, len(values))
	copy(out, values)
	return out
}

// BytesBlitValues applies Go's overlap-safe copy after staged Bytes has checked
// the public Haxe bounds contract.
func BytesBlitValues(target []int, pos int, source []int, sourcePos int, length int) {
	copy(target[pos:pos+length], source[sourcePos:sourcePos+length])
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
