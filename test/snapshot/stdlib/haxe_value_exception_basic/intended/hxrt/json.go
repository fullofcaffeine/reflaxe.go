package hxrt

import "encoding/json"

func JsonParse(source *string) any {
	if source == nil {
		return nil
	}

	var decoded any
	if err := json.Unmarshal([]byte(*source), &decoded); err != nil {
		return nil
	}
	return decoded
}

func JsonStringify(value any) *string {
	encoded, err := json.Marshal(value)
	if err != nil {
		return StringFromLiteral("null")
	}
	return StringFromLiteral(string(encoded))
}
