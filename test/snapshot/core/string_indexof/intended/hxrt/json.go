package hxrt

import (
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
	switch current := value.(type) {
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

func JsonStringify(value any) *string {
	normalized, ok := jsonFromHaxeValue(value, make(map[jsonVisit]bool))
	if !ok {
		return StringFromLiteral("null")
	}
	encoded, err := json.Marshal(normalized)
	if err != nil {
		return StringFromLiteral("null")
	}
	return StringFromLiteral(string(encoded))
}
