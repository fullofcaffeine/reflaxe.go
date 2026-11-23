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
	object = ObjectSource(object)
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
	object = ObjectSource(object)
	// Carrier methods implement runtime protocols, not Haxe function members.
	// The Haxe members of an Array come from ReflectLookupBuiltin instead.
	switch object.(type) {
	case functionValue, *Array:
		return &ReflectFieldLookup{}
	}
	key, valid := reflectFieldName(field)
	switch key {
	case "HxFunctionReceiver":
		if _, generated := object.(interface{ HxFunctionReceiver() any }); generated {
			return &ReflectFieldLookup{}
		}
	case "HxClassName":
		if _, generated := object.(haxeClassName); generated {
			return &ReflectFieldLookup{}
		}
	case "HxEnumValue":
		if _, generated := object.(haxeEnumValue); generated {
			return &ReflectFieldLookup{}
		}
	}
	if object == nil || !valid {
		return &ReflectFieldLookup{}
	}
	method := reflect.ValueOf(object).MethodByName(key)
	if !method.IsValid() || !method.CanInterface() {
		return &ReflectFieldLookup{}
	}
	return &ReflectFieldLookup{Found: true, Value: method.Interface()}
}

// ReflectLookupBuiltin finds a member of the portable Array or String carrier:
// `length`, or a Haxe method as a function bound to the receiver.
//
// What: Array and String are core Haxe classes. Neither the compiler's adapters
// for generated classes nor native struct inspection can see their members.
// Why: --interp returns these members from Reflect.field, which is also how a
// field of a Dynamic value is read, so `d.length` and `d.charAt(0)` work there.
// Staged Reflect.hasField does not call this lookup: --interp reports none of
// these members as a field.
// How: Each method calls the carrier operation or String helper that typed code
// calls, or composes them as the typed lowering does. A missing optional
// argument, or a null one, reads as omitted. `iterator` and `keyValueIterator`
// return generated Haxe classes, which hxrt cannot construct, so they stay absent.
func ReflectLookupBuiltin(object any, field *string) *ReflectFieldLookup {
	key, valid := reflectFieldName(field)
	if !valid {
		return &ReflectFieldLookup{}
	}
	source := ObjectSource(object)
	var method func(...any) any
	var slot string
	switch receiver := source.(type) {
	case *Array:
		if key == "length" {
			return &ReflectFieldLookup{Found: true, Value: receiver.Len()}
		}
		method, slot = reflectArrayMethod(receiver, key), "Array."+key
	case *string:
		if key == "length" {
			return &ReflectFieldLookup{Found: true, Value: StringLengthStringPtr(receiver)}
		}
		method, slot = reflectStringMethod(receiver, key), "String."+key
	}
	if method == nil {
		return &ReflectFieldLookup{}
	}
	// A bound method keeps its receiver and member name, so Reflect.compareMethods
	// recognizes two reads of the same member, while == sees two fresh values.
	return &ReflectFieldLookup{Found: true, Value: BoundFunction(method, func() any { return source }, slot)}
}

func reflectArgument(arguments []any, index int) any {
	if index < len(arguments) {
		return arguments[index]
	}
	return nil
}

func reflectIntArgument(arguments []any, index int) int {
	return IntFromNullableAny(reflectArgument(arguments, index))
}

// reflectOptionalIntArgument reads a missing or null optional Int as omitted.
func reflectOptionalIntArgument(arguments []any, index int) (int, bool) {
	value := reflectArgument(arguments, index)
	if value == nil {
		return 0, false
	}
	return IntFromNullableAny(value), true
}

func reflectStringArgument(arguments []any, index int) *string {
	text, _ := reflectArgument(arguments, index).(*string)
	return text
}

func reflectArrayArgument(arguments []any, index int) *Array {
	array, _ := ObjectSource(reflectArgument(arguments, index)).(*Array)
	if array == nil {
		Throw(StringFromLiteral("Invalid argument: expected Array"))
	}
	return array
}

// reflectCallback calls a Haxe function passed to a reflected Array method.
func reflectCallback(callback any, arguments ...any) any {
	if FunctionOrigin(callback) == nil {
		Throw(StringFromLiteral("Cannot call null"))
	}
	return ReflectCallMethod(callback, arguments)
}

// reflectArraySearch applies the typed Array search lowering with Haxe equality:
// a negative start counts from the end, then clamps into the Array.
func reflectArraySearch(array *Array, value any, start int, hasStart bool, backward bool) int {
	length := array.Len()
	index := 0
	if backward {
		index = length - 1
	}
	if hasStart {
		index = start
		if index < 0 {
			index += length
		}
		if backward && index >= length {
			index = length - 1
		}
		if !backward && index < 0 {
			index = 0
		}
	}
	step := 1
	if backward {
		step = -1
	}
	for ; index >= 0 && index < length; index += step {
		if HaxeEqual(array.Get(index), value) {
			return index
		}
	}
	return -1
}

// reflectArrayMethod returns the Haxe Array method `key` on one Array, or nil.
func reflectArrayMethod(array *Array, key string) func(...any) any {
	switch key {
	case "concat":
		return func(arguments ...any) any {
			result := array.Copy()
			result.Push(reflectArrayArgument(arguments, 0).Values()...)
			return result
		}
	case "contains":
		return func(arguments ...any) any {
			return reflectArraySearch(array, reflectArgument(arguments, 0), 0, false, false) >= 0
		}
	case "copy":
		return func(...any) any { return array.Copy() }
	case "filter":
		return func(arguments ...any) any {
			result := NewArray()
			for index := 0; index < array.Len(); index++ {
				value := array.Get(index)
				if keep, _ := reflectCallback(reflectArgument(arguments, 0), value).(bool); keep {
					result.Push(value)
				}
			}
			return result
		}
	case "indexOf", "lastIndexOf":
		backward := key == "lastIndexOf"
		return func(arguments ...any) any {
			start, hasStart := reflectOptionalIntArgument(arguments, 1)
			return reflectArraySearch(array, reflectArgument(arguments, 0), start, hasStart, backward)
		}
	case "insert":
		return func(arguments ...any) any {
			array.Insert(reflectIntArgument(arguments, 0), reflectArgument(arguments, 1))
			return nil
		}
	case "join":
		return func(arguments ...any) any { return StringJoinAny(array.Values(), reflectStringArgument(arguments, 0)) }
	case "map":
		return func(arguments ...any) any {
			result := NewArray()
			for index := 0; index < array.Len(); index++ {
				result.Push(reflectCallback(reflectArgument(arguments, 0), array.Get(index)))
			}
			return result
		}
	case "pop":
		return func(...any) any { return array.Pop() }
	case "push":
		return func(arguments ...any) any { return array.Push(reflectArgument(arguments, 0)) }
	case "remove":
		return func(arguments ...any) any {
			index := reflectArraySearch(array, reflectArgument(arguments, 0), 0, false, false)
			if index < 0 {
				return false
			}
			array.RemoveAt(index)
			return true
		}
	case "resize":
		return func(arguments ...any) any {
			array.SetLength(reflectIntArgument(arguments, 0))
			return nil
		}
	case "reverse":
		return func(...any) any {
			array.Reverse()
			return nil
		}
	case "shift":
		return func(...any) any {
			value := array.Get(0)
			array.RemoveAt(0)
			return value
		}
	case "slice":
		return func(arguments ...any) any {
			start := reflectIntArgument(arguments, 0)
			if end, hasEnd := reflectOptionalIntArgument(arguments, 1); hasEnd {
				return array.Slice(start, end)
			}
			return array.SliceFrom(start)
		}
	case "sort":
		return func(arguments ...any) any {
			compare := reflectArgument(arguments, 0)
			ArraySort(array, NewFunction(func(left any, right any) int {
				return IntFromNullableAny(reflectCallback(compare, left, right))
			}))
			return nil
		}
	case "splice":
		return func(arguments ...any) any {
			return array.Splice(reflectIntArgument(arguments, 0), reflectIntArgument(arguments, 1))
		}
	case "toString":
		return func(...any) any { return StdString(array) }
	case "unshift":
		return func(arguments ...any) any {
			array.Insert(0, reflectArgument(arguments, 0))
			return nil
		}
	}
	return nil
}

// reflectStringMethod returns the Haxe String method `key` on one String, or nil.
func reflectStringMethod(text *string, key string) func(...any) any {
	switch key {
	case "charAt":
		return func(arguments ...any) any { return StringCharAtStringPtr(text, reflectIntArgument(arguments, 0)) }
	case "charCodeAt":
		return func(arguments ...any) any {
			return StringCharCodeAtAnyStringPtr(text, reflectIntArgument(arguments, 0))
		}
	case "indexOf":
		return func(arguments ...any) any {
			start, hasStart := reflectOptionalIntArgument(arguments, 1)
			return StringIndexOfStringPtr(text, reflectStringArgument(arguments, 0), start, hasStart)
		}
	case "lastIndexOf":
		return func(arguments ...any) any {
			start, hasStart := reflectOptionalIntArgument(arguments, 1)
			return StringLastIndexOfStringPtr(text, reflectStringArgument(arguments, 0), start, hasStart)
		}
	case "split":
		return func(arguments ...any) any {
			parts := StringSplitStringPtr(text, reflectStringArgument(arguments, 0))
			values := make([]any, len(parts))
			for index, part := range parts {
				values[index] = part
			}
			return &Array{values: values}
		}
	case "substr":
		return func(arguments ...any) any {
			length, hasLength := reflectOptionalIntArgument(arguments, 1)
			return StringSubstrStringPtr(text, reflectIntArgument(arguments, 0), length, hasLength)
		}
	case "substring":
		return func(arguments ...any) any {
			end, hasEnd := reflectOptionalIntArgument(arguments, 1)
			if !hasEnd {
				end = StringLengthStringPtr(text)
			}
			return StringSubstringStringPtr(text, reflectIntArgument(arguments, 0), end)
		}
	case "toLowerCase":
		return func(...any) any { return StringToLowerCaseStringPtr(text) }
	case "toUpperCase":
		return func(...any) any { return StringToUpperCaseStringPtr(text) }
	case "toString":
		return func(...any) any { return text }
	}
	return nil
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

// ReflectSetField mutates anonymous maps, exported native struct fields and the
// length of an Array, which truncates it or extends it with nulls as in
// --interp. A false result lets staged Reflect try the typed generated-field
// adapter.
func ReflectSetField(object any, field *string, value any) bool {
	object = ObjectSource(object)
	key, valid := reflectFieldName(field)
	if object == nil || !valid {
		return false
	}
	switch target := object.(type) {
	case *Array:
		if key != "length" {
			return false
		}
		target.SetLength(IntFromNullableAny(value))
		return true
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
	function = FunctionOrigin(function)
	if value, ok := function.(functionValue); ok {
		function = value.functionPayload()
	}
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
// An Array or a String lists only `length`, as in --interp.
func ReflectFields(object any) []*string {
	object = ObjectSource(object)
	if object == nil {
		return nil
	}
	keys := make([]string, 0)
	switch value := object.(type) {
	case *Array, *string:
		keys = append(keys, "length")
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
	if function, ok := value.(functionValue); ok {
		return function.functionIdentity() != nil
	}
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
	leftNil, rightNil := isAnyNil(left), isAnyNil(right)
	if leftNil || rightNil {
		return leftNil && rightNil
	}
	if function, ok := left.(functionValue); ok {
		other, otherOK := right.(functionValue)
		return otherOK && sameFunctionIdentity(function, other)
	}
	if _, ok := right.(functionValue); ok {
		return false
	}
	if left == nil || right == nil {
		return left == nil && right == nil
	}
	leftValue := reflect.ValueOf(left)
	rightValue := reflect.ValueOf(right)
	if !leftValue.IsValid() || !rightValue.IsValid() {
		return !leftValue.IsValid() && !rightValue.IsValid()
	}
	if leftValue.Kind() == reflect.Func && rightValue.Kind() == reflect.Func {
		// Raw native functions have no source identity. A Haxe import must create
		// a Function value before comparison; executable addresses are not keys.
		return false
	}
	return reflect.DeepEqual(left, right)
}

func ReflectIsObject(value any) bool {
	value = ObjectSource(value)
	if _, ok := value.(functionValue); ok {
		return false
	}
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
	object = ObjectSource(object)
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
	object = ObjectSource(object)
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
	return NewFunction(func(arguments ...any) any {
		return ReflectCallMethod(callback, []any{ArrayFromValues(arguments)})
	})
}
