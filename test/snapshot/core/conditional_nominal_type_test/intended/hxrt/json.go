package hxrt

import (
	"bytes"
	"encoding/json"
	"reflect"
)

// jsonVisit identifies one reference-shaped value along the active encoding path.
type jsonVisit struct {
	kind byte
	ptr  uintptr
}

// jsonToHaxeValue converts Go JSON containers into Haxe's runtime containers.
//
// What: It replaces decoded strings and []any values with the Haxe runtime carriers.
// Why: encoding/json has no knowledge of Haxe's nullable string pointer or Array
// identity. Raw Go strings panic when a typed Haxe Dynamic cast expects *string,
// while raw slices bypass Array APIs and later stringify as {}.
// How: Objects remain map[string]any, strings use the canonical string constructor,
// and every nested JSON array uses the portable Array constructor from array.go.
func jsonToHaxeValue(value any) any {
	switch current := value.(type) {
	case string:
		return StringFromLiteral(current)
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
func jsonFromHaxeValue(value any, visiting map[jsonVisit]bool) (any, bool) {
	value = ObjectSource(value)
	switch current := value.(type) {
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

func JsonParse(source *string) any {
	if source == nil {
		return nil
	}

	var decoded any
	if err := json.Unmarshal([]byte(*source), &decoded); err != nil {
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
