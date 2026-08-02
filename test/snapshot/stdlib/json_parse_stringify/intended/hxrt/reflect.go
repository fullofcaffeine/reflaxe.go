package hxrt

import (
	"reflect"
	"sort"
)

// ReflectFieldLookup distinguishes an absent dynamic member from a present
// member whose Haxe value is null. Staged Reflect consumes it immediately.
type ReflectFieldLookup struct {
	Found bool
	Value any
}

func reflectFieldName(field *string) (string, bool) {
	if field == nil {
		return "", false
	}
	return *StdString(field), true
}

// ReflectLookupField inspects only native map and exported struct fields.
// Compiler-known RTTI and generated lowercase members remain same-package
// metadata adapters so hxrt never becomes a registry of generated program types.
func ReflectLookupField(object any, field *string) *ReflectFieldLookup {
	key, valid := reflectFieldName(field)
	if object == nil || !valid {
		return &ReflectFieldLookup{}
	}
	switch value := object.(type) {
	case map[string]any:
		fieldValue, found := value[key]
		return &ReflectFieldLookup{Found: found, Value: fieldValue}
	case map[any]any:
		fieldValue, found := value[key]
		return &ReflectFieldLookup{Found: found, Value: fieldValue}
	case *map[string]any:
		if value == nil {
			return &ReflectFieldLookup{}
		}
		fieldValue, found := (*value)[key]
		return &ReflectFieldLookup{Found: found, Value: fieldValue}
	case *map[any]any:
		if value == nil {
			return &ReflectFieldLookup{}
		}
		fieldValue, found := (*value)[key]
		return &ReflectFieldLookup{Found: found, Value: fieldValue}
	}

	current := reflect.ValueOf(object)
	for current.IsValid() && (current.Kind() == reflect.Interface || current.Kind() == reflect.Pointer) {
		if current.IsNil() {
			return &ReflectFieldLookup{}
		}
		current = current.Elem()
	}
	if !current.IsValid() || current.Kind() != reflect.Struct {
		return &ReflectFieldLookup{}
	}
	fieldValue := current.FieldByName(key)
	if !fieldValue.IsValid() || !fieldValue.CanInterface() {
		return &ReflectFieldLookup{}
	}
	return &ReflectFieldLookup{Found: true, Value: fieldValue.Interface()}
}

// ReflectLookupMethod exposes only exported native Go methods. Generated Haxe
// methods are lowercase and are resolved by the compiler metadata adapter first.
func ReflectLookupMethod(object any, field *string) *ReflectFieldLookup {
	key, valid := reflectFieldName(field)
	if object == nil || !valid {
		return &ReflectFieldLookup{}
	}
	method := reflect.ValueOf(object).MethodByName(key)
	if !method.IsValid() || !method.CanInterface() {
		return &ReflectFieldLookup{}
	}
	return &ReflectFieldLookup{Found: true, Value: method.Interface()}
}

func reflectAssignField(field reflect.Value, value any) bool {
	if !field.IsValid() || !field.CanSet() {
		return false
	}
	if value == nil {
		field.Set(reflect.Zero(field.Type()))
		return true
	}
	incoming := reflect.ValueOf(value)
	if incoming.Type().AssignableTo(field.Type()) {
		field.Set(incoming)
		return true
	}
	if incoming.Type().ConvertibleTo(field.Type()) {
		field.Set(incoming.Convert(field.Type()))
		return true
	}
	if field.Kind() == reflect.Interface && incoming.Type().Implements(field.Type()) {
		field.Set(incoming)
		return true
	}
	return false
}

// ReflectSetField mutates anonymous maps and exported native struct fields. A
// false result lets staged Reflect try the typed generated-field adapter.
func ReflectSetField(object any, field *string, value any) bool {
	key, valid := reflectFieldName(field)
	if object == nil || !valid {
		return false
	}
	switch target := object.(type) {
	case map[string]any:
		target[key] = value
		return true
	case map[any]any:
		target[key] = value
		return true
	case *map[string]any:
		if target == nil {
			return false
		}
		(*target)[key] = value
		return true
	case *map[any]any:
		if target == nil {
			return false
		}
		(*target)[key] = value
		return true
	}
	current := reflect.ValueOf(object)
	if !current.IsValid() || current.Kind() != reflect.Pointer || current.IsNil() {
		return false
	}
	current = current.Elem()
	if !current.IsValid() || current.Kind() != reflect.Struct {
		return false
	}
	return reflectAssignField(current.FieldByName(key), value)
}

func reflectCallArgument(value any, target reflect.Type) reflect.Value {
	if value == nil {
		return reflect.Zero(target)
	}
	incoming := reflect.ValueOf(value)
	if incoming.Type().AssignableTo(target) {
		return incoming
	}
	if incoming.Type().ConvertibleTo(target) {
		return incoming.Convert(target)
	}
	if target.Kind() == reflect.Interface && incoming.Type().Implements(target) {
		return incoming
	}
	return incoming
}

// ReflectCallMethod invokes an already-resolved function. Object context is not
// needed on Go because generated and reflected method values are already bound.
func ReflectCallMethod(function any, arguments []any) any {
	if function == nil {
		return nil
	}
	callable := reflect.ValueOf(function)
	if !callable.IsValid() || callable.Kind() != reflect.Func {
		return nil
	}
	functionType := callable.Type()
	callArguments := make([]reflect.Value, 0, len(arguments))
	for index, argument := range arguments {
		parameterIndex := index
		if functionType.IsVariadic() && parameterIndex >= functionType.NumIn()-1 {
			callArguments = append(callArguments, reflectCallArgument(argument, functionType.In(functionType.NumIn()-1).Elem()))
			continue
		}
		callArguments = append(callArguments, reflectCallArgument(argument, functionType.In(parameterIndex)))
	}
	results := callable.Call(callArguments)
	if len(results) == 0 {
		return nil
	}
	return results[0].Interface()
}

// ReflectFields returns deterministic native object field names. The public API
// does not promise ordering, but sorting stabilizes generated tests and callers.
func ReflectFields(object any) []*string {
	if object == nil {
		return nil
	}
	keys := make([]string, 0)
	switch value := object.(type) {
	case map[string]any:
		for key := range value {
			keys = append(keys, key)
		}
	case map[any]any:
		for key := range value {
			keys = append(keys, *StdString(key))
		}
	case *map[string]any:
		if value != nil {
			for key := range *value {
				keys = append(keys, key)
			}
		}
	case *map[any]any:
		if value != nil {
			for key := range *value {
				keys = append(keys, *StdString(key))
			}
		}
	default:
		current := reflect.ValueOf(object)
		for current.IsValid() && (current.Kind() == reflect.Interface || current.Kind() == reflect.Pointer) {
			if current.IsNil() {
				return nil
			}
			current = current.Elem()
		}
		if current.IsValid() && current.Kind() == reflect.Struct {
			currentType := current.Type()
			for index := 0; index < currentType.NumField(); index++ {
				field := currentType.Field(index)
				if field.PkgPath == "" {
					keys = append(keys, field.Name)
				}
			}
		}
	}
	sort.Strings(keys)
	result := make([]*string, len(keys))
	for index, key := range keys {
		result[index] = StringFromLiteral(key)
	}
	return result
}

func ReflectIsFunction(value any) bool {
	if value == nil {
		return false
	}
	current := reflect.ValueOf(value)
	return current.IsValid() && current.Kind() == reflect.Func
}

func ReflectCompare(left any, right any) int {
	toFloat := func(value any) (float64, bool) {
		switch number := value.(type) {
		case int:
			return float64(number), true
		case int8:
			return float64(number), true
		case int16:
			return float64(number), true
		case int32:
			return float64(number), true
		case int64:
			return float64(number), true
		case uint:
			return float64(number), true
		case uint8:
			return float64(number), true
		case uint16:
			return float64(number), true
		case uint32:
			return float64(number), true
		case uint64:
			return float64(number), true
		case float32:
			return float64(number), true
		case float64:
			return number, true
		default:
			return 0, false
		}
	}
	if leftNumber, ok := toFloat(left); ok {
		if rightNumber, rightOK := toFloat(right); rightOK {
			if leftNumber < rightNumber {
				return -1
			}
			if leftNumber > rightNumber {
				return 1
			}
			return 0
		}
	}
	leftString := *StdString(left)
	rightString := *StdString(right)
	if leftString < rightString {
		return -1
	}
	if leftString > rightString {
		return 1
	}
	return 0
}

func ReflectCompareMethods(left any, right any) bool {
	if left == nil || right == nil {
		return left == nil && right == nil
	}
	leftValue := reflect.ValueOf(left)
	rightValue := reflect.ValueOf(right)
	if !leftValue.IsValid() || !rightValue.IsValid() {
		return !leftValue.IsValid() && !rightValue.IsValid()
	}
	if leftValue.Kind() == reflect.Func && rightValue.Kind() == reflect.Func {
		if leftValue.IsNil() || rightValue.IsNil() {
			return leftValue.IsNil() && rightValue.IsNil()
		}
		return leftValue.Pointer() == rightValue.Pointer()
	}
	return reflect.DeepEqual(left, right)
}

func ReflectIsObject(value any) bool {
	if value == nil {
		return false
	}
	if stringValue, ok := value.(*string); ok {
		return stringValue != nil
	}
	current := reflect.ValueOf(value)
	for current.IsValid() && (current.Kind() == reflect.Interface || current.Kind() == reflect.Pointer) {
		if current.IsNil() {
			return false
		}
		current = current.Elem()
	}
	return current.IsValid() && (current.Kind() == reflect.Struct || current.Kind() == reflect.Map)
}

func ReflectDeleteField(object any, field *string) bool {
	key, valid := reflectFieldName(field)
	if object == nil || !valid {
		return false
	}
	switch value := object.(type) {
	case map[string]any:
		_, found := value[key]
		delete(value, key)
		return found
	case map[any]any:
		_, found := value[key]
		delete(value, key)
		return found
	case *map[string]any:
		if value == nil {
			return false
		}
		_, found := (*value)[key]
		delete(*value, key)
		return found
	case *map[any]any:
		if value == nil {
			return false
		}
		_, found := (*value)[key]
		delete(*value, key)
		return found
	default:
		return false
	}
}

func ReflectCopy(object any) any {
	if object == nil {
		return nil
	}
	switch value := object.(type) {
	case map[string]any:
		copy := make(map[string]any, len(value))
		for key, fieldValue := range value {
			copy[key] = fieldValue
		}
		return copy
	case map[any]any:
		copy := make(map[any]any, len(value))
		for key, fieldValue := range value {
			copy[key] = fieldValue
		}
		return copy
	case *map[string]any:
		if value == nil {
			return nil
		}
		copy := ReflectCopy(*value).(map[string]any)
		return &copy
	case *map[any]any:
		if value == nil {
			return nil
		}
		copy := ReflectCopy(*value).(map[any]any)
		return &copy
	default:
		return nil
	}
}

// ReflectMakeVarArgs adapts Haxe's Array-taking callback to one Go variadic
// function. The callback receives a fresh portable Array on every invocation.
func ReflectMakeVarArgs(callback any) any {
	if callback == nil {
		return nil
	}
	return func(arguments ...any) any {
		return ReflectCallMethod(callback, []any{ArrayFromValues(arguments)})
	}
}
