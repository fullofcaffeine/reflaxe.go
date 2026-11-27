package hxrt

import (
	"bytes"
	"encoding/json"
	"math"
	"reflect"
	"strconv"
	"strings"
)

// jsonVisit identifies one reference-shaped value along the active encoding path.
type jsonVisit struct {
	kind byte
	ptr  uintptr
}

// jsonToHaxeValue converts Go JSON containers into Haxe's runtime containers.
//
// What: It replaces decoded strings, numbers and []any values with the Haxe
// runtime carriers.
// Why: encoding/json has no knowledge of Haxe's nullable string pointer, Array
// identity or Int. Raw Go strings panic when a typed Haxe Dynamic cast expects
// *string, while raw slices bypass Array APIs and later stringify as {}.
// How: Objects remain map[string]any, strings use the canonical string constructor,
// numbers become an Int or a Float through jsonNumberToHaxe, and every nested JSON
// array uses the portable Array constructor from array.go.
func jsonToHaxeValue(value any) any {
	switch current := value.(type) {
	case string:
		return StringFromLiteral(current)
	case json.Number:
		return jsonNumberToHaxe(string(current))
	case []any:
		values := make([]any, len(current))
		for index, item := range current {
			values[index] = jsonToHaxeValue(item)
		}
		return portableArrayFromValues(values)
	case map[string]any:
		out := make(map[string]any, len(current))
		for key, item := range current {
			out[key] = jsonToHaxeValue(item)
		}
		return out
	default:
		return value
	}
}

// jsonFromHaxeValue exposes Haxe containers in the shapes encoding/json accepts.
//
// What: It recursively converts portable Array carriers to ordinary JSON slices.
// Why: Array deliberately hides its erased storage; marshaling the carrier struct
// directly would emit {} instead of a JSON array.
// How: Copy only the container graph needed by the encoder and track active maps,
// slices, and carriers that expose the internal Values protocol, so cyclic Dynamic
// values retain the previous "null on error" behavior instead of recursing forever.
// A Float becomes the number text Haxe's JsonPrinter writes: Std.string's text
// when it is finite, else null (encoding/json would reject the whole document).
func jsonFromHaxeValue(value any, visiting map[jsonVisit]bool) (any, bool) {
	value = ObjectSource(value)
	switch current := value.(type) {
	case float64:
		if math.IsNaN(current) || math.IsInf(current, 0) {
			return nil, true
		}
		return json.Number(formatHaxeFloat(current, 64)), true
	case functionValue:
		if current.functionIdentity() == nil {
			return nil, true
		}
		return "<fun>", true
	case interface{ Values() []any }:
		ref := reflect.ValueOf(current)
		if !ref.IsValid() || (ref.Kind() == reflect.Pointer && ref.IsNil()) {
			return nil, true
		}
		if ref.Kind() != reflect.Pointer {
			return value, true
		}
		visit := jsonVisit{kind: 'a', ptr: ref.Pointer()}
		if visiting[visit] {
			return nil, false
		}
		visiting[visit] = true
		defer delete(visiting, visit)

		values := current.Values()
		out := make([]any, len(values))
		for index, item := range values {
			converted, ok := jsonFromHaxeValue(item, visiting)
			if !ok {
				return nil, false
			}
			out[index] = converted
		}
		return out, true
	case []any:
		if current == nil {
			return nil, true
		}
		visit := jsonVisit{kind: 's', ptr: reflect.ValueOf(current).Pointer()}
		if visiting[visit] {
			return nil, false
		}
		visiting[visit] = true
		defer delete(visiting, visit)

		out := make([]any, len(current))
		for index, item := range current {
			converted, ok := jsonFromHaxeValue(item, visiting)
			if !ok {
				return nil, false
			}
			out[index] = converted
		}
		return out, true
	case map[string]any:
		if current == nil {
			return nil, true
		}
		visit := jsonVisit{kind: 'm', ptr: reflect.ValueOf(current).Pointer()}
		if visiting[visit] {
			return nil, false
		}
		visiting[visit] = true
		defer delete(visiting, visit)

		out := make(map[string]any, len(current))
		for key, item := range current {
			if fn, ok := item.(functionValue); ok && fn.functionIdentity() != nil {
				continue
			}
			converted, ok := jsonFromHaxeValue(item, visiting)
			if !ok {
				return nil, false
			}
			out[key] = converted
		}
		return out, true
	default:
		return value, true
	}
}

// jsonNumberToHaxe returns the Haxe value of one JSON number.
//
// What: An Int for a number written without a `.` whose value is integral and
// fits in Haxe's 32-bit Int (`4`, `-0`, `1e3`), and a Float for every other
// number (`4.0`, `2.5e0`, `2147483648`, `1e400`).
// Why: Haxe's JsonParser, which `--interp`, neko, Java, C++ and HashLink use,
// makes this choice: it reads the number as a Float and returns `Std.int(f)`
// when the text has no `.` and that Int equals `f`. A typed `{count:Int}` read
// needs a Go int, and Std.isOfType, Type.typeof, Std.string (`-0` is the Int 0)
// and the Serializer (`i4`) report the Int. An integral value beyond 32 bits
// stays a Float, as Haxe's Int is 32-bit on every target.
// How: The decoder keeps each number's text (json.Decoder.UseNumber), and
// encoding/json has already checked its grammar, so ParseFloat can only report
// a value out of range. It then returns the infinity that `Std.parseFloat`
// gives in `--interp`; json.Unmarshal used to reject the whole document.
func jsonNumberToHaxe(text string) any {
	value, _ := strconv.ParseFloat(text, 64)
	if !strings.Contains(text, ".") && value == math.Trunc(value) && value >= math.MinInt32 && value <= math.MaxInt32 {
		return int(value)
	}
	return value
}

// JsonParse decodes one JSON document into Haxe runtime values, or returns nil
// when the text is not exactly one JSON value.
//
// A json.Decoder keeps each number's text for jsonNumberToHaxe. Unlike
// json.Unmarshal, it stops after the first value, so anything but JSON white
// space after that value rejects the document, as json.Unmarshal did.
func JsonParse(source *string) any {
	if source == nil {
		return nil
	}

	decoder := json.NewDecoder(strings.NewReader(*source))
	decoder.UseNumber()
	var decoded any
	if err := decoder.Decode(&decoded); err != nil {
		return nil
	}
	if strings.TrimLeft((*source)[decoder.InputOffset():], " \t\r\n") != "" {
		return nil
	}
	return jsonToHaxeValue(decoded)
}

func JsonStringify(value any, space *string) *string {
	normalized, ok := jsonFromHaxeValue(value, make(map[jsonVisit]bool))
	if !ok {
		return StringFromLiteral("null")
	}
	// Haxe JSON does not use HTML escaping. Preserve a present empty indent as
	// pretty-printing too; Encoder.SetIndent alone cannot represent that state.
	var encoded bytes.Buffer
	encoder := json.NewEncoder(&encoded)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(normalized); err != nil {
		return StringFromLiteral("null")
	}
	data := bytes.TrimSuffix(encoded.Bytes(), []byte{'\n'})
	if space != nil {
		var indented bytes.Buffer
		if err := json.Indent(&indented, data, "", *space); err != nil {
			return StringFromLiteral("null")
		}
		data = indented.Bytes()
	}
	return StringFromLiteral(string(data))
}
