package hxrt

import "reflect"

// TemplateArrayValues exposes a Go slice or array as the erased element slice
// consumed by staged haxe.Template foreach loops. The runtime owns only this
// representation inspection; iteration and rendering remain in Haxe source.
func TemplateArrayValues(value any) []any {
	if value == nil {
		return nil
	}
	ref := reflect.ValueOf(value)
	if !ref.IsValid() {
		return nil
	}
	if ref.Kind() == reflect.Pointer {
		if ref.IsNil() {
			return nil
		}
		ref = ref.Elem()
	}
	if ref.Kind() != reflect.Slice && ref.Kind() != reflect.Array {
		return nil
	}
	out := make([]any, ref.Len())
	for index := 0; index < ref.Len(); index++ {
		item := ref.Index(index)
		if item.CanInterface() {
			out[index] = item.Interface()
		}
	}
	return out
}

// TemplateIsObject classifies the dynamic record carriers that staged
// haxe.Template may search for fields. This is a runtime representation fact,
// not Template lookup or fallback policy.
func TemplateIsObject(value any) bool {
	if value == nil {
		return false
	}
	ref := reflect.ValueOf(value)
	if !ref.IsValid() {
		return false
	}
	for ref.Kind() == reflect.Pointer || ref.Kind() == reflect.Interface {
		if ref.IsNil() {
			return false
		}
		ref = ref.Elem()
	}
	return ref.Kind() == reflect.Struct || ref.Kind() == reflect.Map
}

// TemplateCall invokes an already-resolved Haxe function with the dynamic
// argument list assembled by staged haxe.Template. Function discovery and all
// macro/iterator policy remain in Haxe source and generated field metadata.
func TemplateCall(funcValue any, args []any) any {
	if funcValue == nil {
		return nil
	}
	function := reflect.ValueOf(funcValue)
	if !function.IsValid() || function.Kind() != reflect.Func {
		return nil
	}
	callArgs := make([]reflect.Value, 0, len(args))
	for _, arg := range args {
		callArgs = append(callArgs, reflect.ValueOf(arg))
	}
	results := function.Call(callArgs)
	if len(results) == 0 {
		return nil
	}
	return results[0].Interface()
}
