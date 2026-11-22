package hxrt

import "reflect"

// IsEnumValue recognizes the generated Go carrier shared by Haxe enum values.
// This is intentionally a representation predicate: recursive comparison and
// all EnumValueMap algorithms remain in staged Haxe source.
func IsEnumValue(value any) bool {
	ref := reflect.ValueOf(value)
	for ref.IsValid() && (ref.Kind() == reflect.Interface || ref.Kind() == reflect.Pointer) {
		if ref.IsNil() {
			return false
		}
		ref = ref.Elem()
	}
	if !ref.IsValid() || ref.Kind() != reflect.Struct {
		return false
	}
	tag := ref.FieldByName("tag")
	params := ref.FieldByName("params")
	return tag.IsValid() && tag.Kind() == reflect.Int && params.IsValid() && params.Kind() == reflect.Slice
}
