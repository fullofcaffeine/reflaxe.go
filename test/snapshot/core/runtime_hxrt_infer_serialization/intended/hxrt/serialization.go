package hxrt

import (
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unsafe"
)

// SerializationField is one deterministic erased field snapshot.
//
// What: Carries a generated Haxe field name and its current value.
// Why: Generated fields are intentionally package-private Go fields, so ordinary
// cross-package reflection cannot expose them to staged Serializer source.
// How: SerializationFields performs the narrow unsafe lift once, then source owns
// field ordering in the wire stream and recursive traversal.
type SerializationField struct {
	Name  *string
	Value any
}

// SerializationParseFloat performs the host numeric conversion after staged
// Unserializer has selected and bounded one floating-point token.
func SerializationParseFloat(value *string) float64 {
	parsed, err := strconv.ParseFloat(*StdString(value), 64)
	if err != nil {
		Throw(err)
		return 0
	}
	return parsed
}

func serializationStruct(value any) (reflect.Value, bool) {
	current := reflect.ValueOf(value)
	for current.IsValid() && (current.Kind() == reflect.Interface || current.Kind() == reflect.Pointer) {
		if current.IsNil() {
			return reflect.Value{}, false
		}
		current = current.Elem()
	}
	return current, current.IsValid() && current.Kind() == reflect.Struct
}

func serializationFieldValue(value reflect.Value) (any, bool) {
	if !value.IsValid() {
		return nil, false
	}
	if value.CanInterface() {
		return value.Interface(), true
	}
	if !value.CanAddr() {
		return nil, false
	}
	lifted := serializationAccessibleField(value)
	if !lifted.IsValid() || !lifted.CanInterface() {
		return nil, false
	}
	return lifted.Interface(), true
}

func serializationAccessibleField(value reflect.Value) reflect.Value {
	return reflect.NewAt(value.Type(), unsafe.Pointer(value.UnsafeAddr())).Elem()
}

// SerializationFields returns the portable instance-field view of one object.
func SerializationFields(value any) []*SerializationField {
	switch object := value.(type) {
	case map[string]any:
		return serializationStringMapFields(object)
	case *map[string]any:
		if object == nil {
			return nil
		}
		return serializationStringMapFields(*object)
	case map[any]any:
		converted := make(map[string]any, len(object))
		for key, fieldValue := range object {
			converted[*StdString(key)] = fieldValue
		}
		return serializationStringMapFields(converted)
	case *map[any]any:
		if object == nil {
			return nil
		}
		converted := make(map[string]any, len(*object))
		for key, fieldValue := range *object {
			converted[*StdString(key)] = fieldValue
		}
		return serializationStringMapFields(converted)
	}
	current, ok := serializationStruct(value)
	if !ok {
		return nil
	}
	typeInfo := current.Type()
	out := make([]*SerializationField, 0, current.NumField())
	for index := 0; index < current.NumField(); index++ {
		fieldInfo := typeInfo.Field(index)
		if strings.HasPrefix(fieldInfo.Name, "__hx_") {
			continue
		}
		fieldValue, available := serializationFieldValue(current.Field(index))
		if !available {
			continue
		}
		out = append(out, &SerializationField{
			Name:  StringFromLiteral(fieldInfo.Name),
			Value: fieldValue,
		})
	}
	return out
}

func serializationStringMapFields(object map[string]any) []*SerializationField {
	keys := make([]string, 0, len(object))
	for key := range object {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	out := make([]*SerializationField, 0, len(keys))
	for _, key := range keys {
		out = append(out, &SerializationField{Name: StringFromLiteral(key), Value: object[key]})
	}
	return out
}

// SerializationSetField assigns one decoded field to an object or anonymous map.
//
// Dynamic is deliberate at this runtime boundary: the serialization format erases
// field types, while the reflected destination restores its exact generated type.
func SerializationSetField(target any, name *string, value any) {
	if target == nil || name == nil {
		return
	}
	fieldName := *name
	switch object := target.(type) {
	case map[string]any:
		object[fieldName] = value
		return
	case map[any]any:
		object[fieldName] = value
		return
	case *map[string]any:
		if object != nil {
			(*object)[fieldName] = value
		}
		return
	case *map[any]any:
		if object != nil {
			(*object)[fieldName] = value
		}
		return
	}
	current := reflect.ValueOf(target)
	if !current.IsValid() || current.Kind() != reflect.Pointer || current.IsNil() {
		return
	}
	current = current.Elem()
	if !current.IsValid() || current.Kind() != reflect.Struct {
		return
	}
	field := current.FieldByName(fieldName)
	if !field.IsValid() {
		return
	}
	if !field.CanSet() {
		if !field.CanAddr() {
			return
		}
		field = serializationAccessibleField(field)
	}
	if value == nil {
		field.Set(reflect.Zero(field.Type()))
		return
	}
	incoming := reflect.ValueOf(value)
	if incoming.Type().AssignableTo(field.Type()) {
		field.Set(incoming)
		return
	}
	if incoming.Type().ConvertibleTo(field.Type()) {
		field.Set(incoming.Convert(field.Type()))
		return
	}
	if field.Kind() == reflect.Interface {
		field.Set(incoming)
	}
}

// SerializationBindSelf repairs the hidden virtual-dispatch self pointer after
// Type.createEmptyInstance bypasses a generated constructor.
func SerializationBindSelf(instance any) {
	current := reflect.ValueOf(instance)
	if !current.IsValid() || current.Kind() != reflect.Pointer || current.IsNil() {
		return
	}
	element := current.Elem()
	if !element.IsValid() || element.Kind() != reflect.Struct {
		return
	}
	field := element.FieldByName("__hx_this")
	if !field.IsValid() || !current.Type().AssignableTo(field.Type()) {
		return
	}
	if !field.CanSet() {
		if !field.CanAddr() {
			return
		}
		field = serializationAccessibleField(field)
	}
	field.Set(current)
}
