package hxrt

import "reflect"

type objectMapIdentity struct {
	typeOf  reflect.Type
	pointer uintptr
	value   any
	byValue bool
}

type objectMapEntry struct {
	key   any
	value any
}

// ObjectMapEntry is the temporary typed serializer view of one identity-map
// entry. The public ObjectMap API never exposes this runtime carrier.
type ObjectMapEntry struct {
	Key   any
	Value any
}

// ObjectMapCell retains original keys strongly and indexes them by reference
// identity. order preserves the first insertion of each live key.
type ObjectMapCell struct {
	entries map[objectMapIdentity]objectMapEntry
	order   []objectMapIdentity
}

func ObjectMapNew() *ObjectMapCell {
	return &ObjectMapCell{entries: make(map[objectMapIdentity]objectMapEntry)}
}

func objectMapIdentityOf(key any) (objectMapIdentity, bool) {
	value := reflect.ValueOf(key)
	for value.IsValid() && value.Kind() == reflect.Interface {
		if value.IsNil() {
			return objectMapIdentity{}, true
		}
		value = value.Elem()
	}
	if !value.IsValid() {
		return objectMapIdentity{}, true
	}

	identity := objectMapIdentity{typeOf: value.Type()}
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.Slice, reflect.UnsafePointer:
		if value.IsNil() {
			return identity, true
		}
		identity.pointer = value.Pointer()
		return identity, true
	default:
		if value.Type().Comparable() {
			identity.byValue = true
			identity.value = value.Interface()
			return identity, true
		}
		return objectMapIdentity{}, false
	}
}

func objectMapRequireIdentity(key any) objectMapIdentity {
	identity, ok := objectMapIdentityOf(key)
	if !ok {
		Throw(StringFromLiteral("ObjectMap key does not have stable reference identity"))
		return objectMapIdentity{}
	}
	return identity
}

func ObjectMapSet(cell *ObjectMapCell, key any, value any) {
	identity := objectMapRequireIdentity(key)
	if _, exists := cell.entries[identity]; !exists {
		cell.order = append(cell.order, identity)
	}
	cell.entries[identity] = objectMapEntry{key: key, value: value}
}

func ObjectMapGet(cell *ObjectMapCell, key any) any {
	entry, exists := cell.entries[objectMapRequireIdentity(key)]
	if !exists {
		return nil
	}
	return entry.value
}

func ObjectMapExists(cell *ObjectMapCell, key any) bool {
	_, exists := cell.entries[objectMapRequireIdentity(key)]
	return exists
}

func ObjectMapRemove(cell *ObjectMapCell, key any) bool {
	identity := objectMapRequireIdentity(key)
	if _, exists := cell.entries[identity]; !exists {
		return false
	}
	delete(cell.entries, identity)
	for index, orderedIdentity := range cell.order {
		if orderedIdentity == identity {
			cell.order = append(cell.order[:index], cell.order[index+1:]...)
			break
		}
	}
	return true
}

func ObjectMapKeys(cell *ObjectMapCell) []any {
	out := make([]any, 0, len(cell.order))
	for _, identity := range cell.order {
		out = append(out, cell.entries[identity].key)
	}
	return out
}

func ObjectMapClear(cell *ObjectMapCell) {
	cell.entries = make(map[objectMapIdentity]objectMapEntry)
	cell.order = nil
}

// ObjectMapSnapshot is a narrow serializer bridge retained until Serializer
// itself moves to staged source. Its insertion-ordered slice also removes Go
// map iteration nondeterminism from object-map wire output.
func ObjectMapSnapshot(cell *ObjectMapCell) []ObjectMapEntry {
	out := make([]ObjectMapEntry, 0, len(cell.order))
	for _, identity := range cell.order {
		entry := cell.entries[identity]
		out = append(out, ObjectMapEntry{Key: entry.key, Value: entry.value})
	}
	return out
}
