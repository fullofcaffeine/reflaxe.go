package hxrt

import (
	"fmt"
	"math"
	"math/bits"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"
)

func isAnyNil(value any) bool {
	if value == nil {
		return true
	}

	switch value.(type) {
	case bool, string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64:
		return false
	}

	inner := reflect.ValueOf(value)
	if !inner.IsValid() {
		return true
	}

	switch inner.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return inner.IsNil()
	default:
		return false
	}
}

// ObjectSource resolves structural views without importing their optional runtime
// implementation. It then resolves the live logical receiver of generated base
// carriers. Nil checks precede protocol calls, including typed nil pointers.
// Function carriers do not implement either protocol and keep their own identity.
func ObjectSource(value any) any {
	for !isAnyNil(value) {
		source, ok := value.(interface{ sourceObject() any })
		if !ok {
			break
		}
		value = source.sourceObject()
	}
	if isAnyNil(value) {
		return nil
	}
	if receiver, ok := value.(interface{ HxFunctionReceiver() any }); ok {
		return receiver.HxFunctionReceiver()
	}
	return value
}

// StdString implements Haxe's `Std.string` for an erased value.
//
// What: Strings, numbers, booleans, and null keep their direct conversion.
// Arrays, enum values, anonymous objects, and class instances print the way
// the Haxe interpreter prints them: `[a,b]`, `Some(1)`, `{a: 1, b: x}`, the
// result of `toString()`, or the class path (`pack.Point`).
// Why: Go's own formatting (`&{1 [1]}`, `map[a:1]`) leaks the runtime carriers.
// How: Scalars return early. Other values go through one depth-limited writer;
// generated enums and classes supply their Haxe names through small exported
// methods (see haxeEnumValue and haxeClassName).
func StdString(value any) *string {
	value = ObjectSource(value)
	if isAnyNil(value) {
		return StringFromLiteral("null")
	}
	if text, ok := value.(*string); ok {
		return text
	}
	if text, ok := scalarHaxeString(value); ok {
		return StringFromLiteral(text)
	}
	var text strings.Builder
	writeResolvedHaxeString(&text, value, 0)
	return StringFromLiteral(text.String())
}

func scalarHaxeString(value any) (string, bool) {
	switch v := value.(type) {
	case string:
		return v, true
	case bool:
		return strconv.FormatBool(v), true
	case float32, float64:
		return fmt.Sprint(v), true
	case int:
		return strconv.FormatInt(int64(v), 10), true
	case int8:
		return strconv.FormatInt(int64(v), 10), true
	case int16:
		return strconv.FormatInt(int64(v), 10), true
	case int32:
		return strconv.FormatInt(int64(v), 10), true
	case int64:
		return strconv.FormatInt(v, 10), true
	case uint:
		return strconv.FormatUint(uint64(v), 10), true
	case uint8:
		return strconv.FormatUint(uint64(v), 10), true
	case uint16:
		return strconv.FormatUint(uint64(v), 10), true
	case uint32:
		return strconv.FormatUint(uint64(v), 10), true
	case uint64:
		return strconv.FormatUint(v, 10), true
	case uintptr:
		return strconv.FormatUint(uint64(v), 10), true
	default:
		return "", false
	}
}

// haxeStringMaxDepth is how deep Std.string looks into nested values.
//
// What: A value nested more than five levels inside Arrays, enum arguments, or
// anonymous objects prints as `<...>`; so does a `toString()` result at depth 5.
// Why: This is the Haxe interpreter's default print depth (`eval-print-depth`),
// and the interpreter is this repository's semantic reference. The limit also
// stops an Array or object that contains itself from overflowing the Go stack.
// How: Every container passes depth+1 to its children, as the interpreter does.
const haxeStringMaxDepth = 5

// haxeEnumValue is implemented by every generated Haxe enum carrier. It returns
// the constructor name and its arguments, which hold Haxe values.
type haxeEnumValue interface {
	HxEnumValue() (string, []any)
}

// haxeClassName is implemented by generated classes that have no `toString()`.
// It returns the dotted Haxe class path. Classes with `toString()` implement
// fmt.Stringer instead.
type haxeClassName interface {
	HxClassName() string
}

// haxeStringWriter is implemented by the optional Array carrier (array.go). It
// writes its elements inside the same depth-limited pass, and string.go stays
// independent of that file.
type haxeStringWriter interface {
	writeHaxeString(text *strings.Builder, depth int)
}

// haxeFunction matches the optional Function carrier (function.go) through a
// method it already has. A new method would be copied into every Function[F]
// instantiation and grow each program's binary; this needs none.
type haxeFunction interface {
	functionPayload() any
}

func writeHaxeString(text *strings.Builder, value any, depth int) {
	if depth > haxeStringMaxDepth {
		text.WriteString("<...>")
		return
	}
	value = ObjectSource(value)
	if isAnyNil(value) {
		text.WriteString("null")
		return
	}
	writeResolvedHaxeString(text, value, depth)
}

// writeResolvedHaxeString writes a value that ObjectSource already resolved and
// that is not null. Type tests run in this order: runtime carriers, enums,
// anonymous objects, toString() adapters, class names, then scalars and Go's
// own formatting for native values.
func writeResolvedHaxeString(text *strings.Builder, value any, depth int) {
	switch v := value.(type) {
	case *string:
		text.WriteString(*v)
	case haxeStringWriter:
		v.writeHaxeString(text, depth)
	case haxeFunction:
		// Haxe prints a function as `#fun` at its own depth, unlike a toString()
		// result, which is printed one level deeper.
		text.WriteString("#fun")
	case haxeEnumValue:
		name, params := v.HxEnumValue()
		text.WriteString(name)
		if len(params) > 0 {
			text.WriteByte('(')
			for index, param := range params {
				if index > 0 {
					text.WriteByte(',')
				}
				writeHaxeString(text, param, depth+1)
			}
			text.WriteByte(')')
		}
	case map[string]any:
		writeAnonymousHaxeString(text, v, depth+1)
	case fmt.Stringer:
		// A generated toString() adapter. Haxe calls toString first and then
		// prints its result one level deeper.
		result := v.String()
		writeHaxeString(text, &result, depth+1)
	case haxeClassName:
		text.WriteString(v.HxClassName())
	default:
		if scalar, ok := scalarHaxeString(value); ok {
			text.WriteString(scalar)
		} else {
			text.WriteString(fmt.Sprint(value))
		}
	}
}

// writeAnonymousHaxeString prints a Haxe anonymous object.
//
// What: An object with a callable `toString` field prints that call's result.
// Otherwise it prints `{name: value, ...}`, or `{}` when it has no fields.
// Why: Haxe does not define the field order. Go map order is random, so the
// writer needs a fixed order, and it uses the interpreter's order so generated
// programs print the same text as `haxe --interp`.
// How: The interpreter keeps fields in a little-endian Patricia tree keyed by
// OCaml's `Hashtbl.hash` of the name; reading that tree visits keys by their
// bit-reversed value. haxeInterpFieldOrderKey computes the same key.
func writeAnonymousHaxeString(text *strings.Builder, fields map[string]any, depth int) {
	if method, found := fields["toString"].(haxeFunction); found {
		// Only the `() -> String` and `() -> Dynamic` shapes are called; typed
		// assertions keep reflection out of this path.
		switch call := method.functionPayload().(type) {
		case func() *string:
			writeHaxeString(text, call(), depth)
			return
		case func() any:
			writeHaxeString(text, call(), depth)
			return
		}
	}
	ordered := make([]anonymousFieldOrder, 0, len(fields))
	for name := range fields {
		ordered = append(ordered, anonymousFieldOrder{key: haxeInterpFieldOrderKey(name), name: name})
	}
	sortAnonymousFieldOrder(ordered)
	text.WriteByte('{')
	for index, field := range ordered {
		if index > 0 {
			text.WriteString(", ")
		}
		text.WriteString(field.name)
		text.WriteString(": ")
		writeHaxeString(text, fields[field.name], depth)
	}
	text.WriteByte('}')
}

type anonymousFieldOrder struct {
	key  uint32
	name string
}

// sortAnonymousFieldOrder sorts by key, then by name, with an in-place heap sort.
// A small hand-written sort keeps package sort and its reflection helper out of
// every generated program: sort.Slice added about 35 KB to an example binary.
func sortAnonymousFieldOrder(fields []anonymousFieldOrder) {
	less := func(left, right anonymousFieldOrder) bool {
		return left.key < right.key || left.key == right.key && left.name < right.name
	}
	siftDown := func(root, end int) {
		for {
			child := 2*root + 1
			if child >= end {
				return
			}
			if child+1 < end && less(fields[child], fields[child+1]) {
				child++
			}
			if !less(fields[root], fields[child]) {
				return
			}
			fields[root], fields[child] = fields[child], fields[root]
			root = child
		}
	}
	for root := len(fields)/2 - 1; root >= 0; root-- {
		siftDown(root, len(fields))
	}
	for end := len(fields) - 1; end > 0; end-- {
		fields[0], fields[end] = fields[end], fields[0]
		siftDown(0, end)
	}
}

// haxeInterpFieldOrderKey returns the sort key of one anonymous-object field in
// the Haxe interpreter: OCaml's `Hashtbl.hash` of the name (MurmurHash3 mixing
// over little-endian 32-bit blocks, folded to 30 bits), bit-reversed.
func haxeInterpFieldOrderKey(name string) uint32 {
	mix := func(hash, data uint32) uint32 {
		data *= 0xcc9e2d51
		data = bits.RotateLeft32(data, 15)
		data *= 0x1b873593
		hash ^= data
		hash = bits.RotateLeft32(hash, 13)
		return hash*5 + 0xe6546b64
	}
	var hash uint32
	index := 0
	for ; index+4 <= len(name); index += 4 {
		hash = mix(hash, uint32(name[index])|uint32(name[index+1])<<8|uint32(name[index+2])<<16|uint32(name[index+3])<<24)
	}
	if rest := len(name) - index; rest > 0 {
		var data uint32
		for offset := rest - 1; offset >= 0; offset-- {
			data = data<<8 | uint32(name[index+offset])
		}
		hash = mix(hash, data)
	}
	hash ^= uint32(len(name))
	hash ^= hash >> 16
	hash *= 0x85ebca6b
	hash ^= hash >> 13
	hash *= 0xc2b2ae35
	hash ^= hash >> 16
	return bits.Reverse32(hash & 0x3fffffff)
}

// StringParseFloatExact converts one source-validated numeric token.
//
// What: Returns the Go float64 value for an exact decimal/exponent token.
// Why: Native IEEE-754 conversion is representation work, while staged Std owns
// whitespace, prefix scanning, malformed exponents, and trailing input policy.
// How: Parse the complete supplied token and map every native error to Haxe NaN.
func StringParseFloatExact(value *string) float64 {
	if value == nil {
		return math.NaN()
	}
	parsed, err := strconv.ParseFloat(*value, 64)
	if err != nil {
		return math.NaN()
	}
	return parsed
}

func StringSlice(values []*string) []string {
	out := make([]string, len(values))
	for i := 0; i < len(values); i++ {
		out[i] = *StdString(values[i])
	}
	return out
}

func StringConcatAny(left any, right any) *string {
	l := StdString(left)
	r := StdString(right)
	return StringFromLiteral(*l + *r)
}

func StringEqualAny(left any, right any) bool {
	l := StdString(left)
	r := StdString(right)
	return *l == *r
}

func AnyEqualsNull(value any) bool {
	return isAnyNil(value)
}

func stringValueOrNullToken(value *string) string {
	if value == nil {
		return "null"
	}
	return *value
}

func StringConcatStringPtr(left *string, right *string) *string {
	return StringFromLiteral(stringValueOrNullToken(left) + stringValueOrNullToken(right))
}

// StringEqualStringPtr preserves Haxe equality for pointer-backed strings.
//
// What: Compare nullable string carriers by value.
// Why: Haxe null is distinct from the literal string "null" even though string
// conversion and concatenation intentionally render both with the same text.
// How: Resolve nil identity before dereferencing two present string values.
func StringEqualStringPtr(left *string, right *string) bool {
	if left == nil || right == nil {
		return left == nil && right == nil
	}
	return *left == *right
}

func StringLengthStringPtr(value *string) int {
	return utf8.RuneCountInString(stringValueOrNullToken(value))
}

func StringCharAtStringPtr(value *string, index int) *string {
	if index < 0 {
		return StringFromLiteral("")
	}
	raw := stringValueOrNullToken(value)
	runeIndex := 0
	for _, runeValue := range raw {
		if runeIndex == index {
			return StringFromLiteral(string(runeValue))
		}
		runeIndex++
	}
	return StringFromLiteral("")
}

func StringCharCodeAtStringPtr(value *string, index int) int {
	if index < 0 {
		return -1
	}
	raw := stringValueOrNullToken(value)
	runeIndex := 0
	for _, runeValue := range raw {
		if runeIndex == index {
			return int(runeValue)
		}
		runeIndex++
	}
	return -1
}

// StringSliceCodePointsStringPtr exposes Go's rune-aware string representation
// to typed staged std code. Haxe owns bounds and range policy; this helper only
// converts an already-normalized half-open code-point range into a string.
func StringSliceCodePointsStringPtr(value *string, start int, end int) *string {
	return StringFromLiteral(sliceStringByRuneRange(stringValueOrNullToken(value), start, end))
}

func StringCharCodeAtAnyStringPtr(value *string, index int) any {
	code := StringCharCodeAtStringPtr(value, index)
	if code < 0 {
		return nil
	}
	return code
}

func StringSubstringStringPtr(value *string, start int, end int) *string {
	raw := stringValueOrNullToken(value)
	total := utf8.RuneCountInString(raw)
	clampedStart, clampedEnd := clampSubstringRange(start, end, total)
	return StringFromLiteral(sliceStringByRuneRange(raw, clampedStart, clampedEnd))
}

func StringSubstrStringPtr(value *string, pos int, length int, hasLength bool) *string {
	raw := stringValueOrNullToken(value)
	total := utf8.RuneCountInString(raw)

	start := pos
	if start < 0 {
		start = total + start
	}
	if start < 0 {
		start = 0
	}
	if start > total {
		start = total
	}

	end := total
	if hasLength {
		if length < 0 {
			end = total + length
		} else {
			end = start + length
		}
		if end > total {
			end = total
		}
	}
	if end < start {
		end = start
	}
	return StringFromLiteral(sliceStringByRuneRange(raw, start, end))
}

// StringIndexOfStringPtr implements Haxe String.indexOf in code-point space.
// Why: Go strings expose byte indexes, but portable Haxe string operations use
// character indexes. What: the helper returns the first matching code-point
// offset at or after the optional start. How: it clamps the start like Haxe and
// compares rune-range slices so multi-byte UTF-8 text keeps portable indexes.
func StringIndexOfStringPtr(value *string, search *string, startIndex int, hasStart bool) int {
	raw := stringValueOrNullToken(value)
	needle := stringValueOrNullToken(search)
	total := utf8.RuneCountInString(raw)

	start := 0
	if hasStart {
		start = startIndex
		if start < 0 {
			if needle == "" {
				return 0
			}
			start = total + start
			if start < 0 {
				start = 0
			}
		}
		if start > total {
			start = total
		}
	}

	if needle == "" {
		return start
	}

	needleLen := utf8.RuneCountInString(needle)
	if needleLen > total-start {
		return -1
	}
	for index := start; index <= total-needleLen; index++ {
		if sliceStringByRuneRange(raw, index, index+needleLen) == needle {
			return index
		}
	}
	return -1
}

func StringLastIndexOfStringPtr(value *string, search *string, startIndex int, hasStart bool) int {
	raw := stringValueOrNullToken(value)
	needle := stringValueOrNullToken(search)
	total := utf8.RuneCountInString(raw)

	start := total
	if hasStart {
		start = startIndex
		if start < 0 {
			return -1
		}
		if start > total {
			start = total
		}
	}

	if needle == "" {
		return start
	}

	needleLen := utf8.RuneCountInString(needle)
	if needleLen > total || start < 0 {
		return -1
	}
	maxStart := start
	if maxStart > total-needleLen {
		maxStart = total - needleLen
	}
	for index := maxStart; index >= 0; index-- {
		if sliceStringByRuneRange(raw, index, index+needleLen) == needle {
			return index
		}
	}
	return -1
}

func StringSplitStringPtr(value *string, delimiter *string) []*string {
	raw := stringValueOrNullToken(value)
	sep := stringValueOrNullToken(delimiter)
	if raw == "" {
		return []*string{}
	}
	if sep == "" {
		out := make([]*string, 0, utf8.RuneCountInString(raw))
		for _, runeValue := range raw {
			out = append(out, StringFromLiteral(string(runeValue)))
		}
		return out
	}
	parts := strings.Split(raw, sep)
	out := make([]*string, 0, len(parts))
	for _, part := range parts {
		out = append(out, StringFromLiteral(part))
	}
	return out
}

func StringToLowerCaseStringPtr(value *string) *string {
	return StringFromLiteral(strings.ToLower(stringValueOrNullToken(value)))
}

func StringToUpperCaseStringPtr(value *string) *string {
	return StringFromLiteral(strings.ToUpper(stringValueOrNullToken(value)))
}

func StringJoinAny(values []any, delimiter *string) *string {
	sep := stringValueOrNullToken(delimiter)
	parts := make([]string, 0, len(values))
	for _, value := range values {
		parts = append(parts, *StdString(value))
	}
	return StringFromLiteral(strings.Join(parts, sep))
}

func StringFromCharCode(code int) *string {
	return StringFromLiteral(string(rune(code)))
}

func clampSubstringRange(start int, end int, total int) (int, int) {
	if start < 0 {
		start = 0
	}
	if end < 0 {
		end = 0
	}
	if start > total {
		start = total
	}
	if end > total {
		end = total
	}
	if end < start {
		end = start
	}
	return start, end
}

func sliceStringByRuneRange(raw string, start int, end int) string {
	if start >= end {
		return ""
	}

	runeIndex := 0
	byteIndex := 0
	startByte := -1
	endByte := -1
	for byteIndex < len(raw) {
		if runeIndex == start && startByte == -1 {
			startByte = byteIndex
		}
		if runeIndex == end {
			endByte = byteIndex
			break
		}
		_, size := utf8.DecodeRuneInString(raw[byteIndex:])
		byteIndex += size
		runeIndex++
	}

	if startByte == -1 {
		startByte = len(raw)
	}
	if endByte == -1 {
		if runeIndex == end {
			endByte = byteIndex
		} else {
			endByte = len(raw)
		}
	}
	if endByte < startByte {
		endByte = startByte
	}
	return raw[startByte:endByte]
}

func StringLength(value any) int {
	return StringLengthStringPtr(StdString(value))
}

func StringCharAt(value any, index int) *string {
	return StringCharAtStringPtr(StdString(value), index)
}

func StringCharCodeAt(value any, index int) int {
	return StringCharCodeAtStringPtr(StdString(value), index)
}

func StringCharCodeAtAny(value any, index int) any {
	return StringCharCodeAtAnyStringPtr(StdString(value), index)
}

func StringSubstring(value any, start int, end int) *string {
	return StringSubstringStringPtr(StdString(value), start, end)
}

func StringSubstr(value any, pos int, length int, hasLength bool) *string {
	return StringSubstrStringPtr(StdString(value), pos, length, hasLength)
}

func StringIndexOf(value any, search any, startIndex int, hasStart bool) int {
	return StringIndexOfStringPtr(StdString(value), StdString(search), startIndex, hasStart)
}

func StringLastIndexOf(value any, search any, startIndex int, hasStart bool) int {
	return StringLastIndexOfStringPtr(StdString(value), StdString(search), startIndex, hasStart)
}

func StringSplit(value any, delimiter any) []*string {
	return StringSplitStringPtr(StdString(value), StdString(delimiter))
}

func StringToLowerCase(value any) *string {
	return StringToLowerCaseStringPtr(StdString(value))
}

func StringToUpperCase(value any) *string {
	return StringToUpperCaseStringPtr(StdString(value))
}
