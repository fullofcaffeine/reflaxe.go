package hxrt

// Anonymous is the storage boundary for Haxe structural objects. A map owns
// fields; a view retains a source object and uses compiler-authored member access.
// Erasure is confined to this heterogeneous field boundary, never invocation.
type Anonymous interface {
	sourceObject() any
	readField(string) AnonymousField
	writeField(string, any) bool
}

// AnonymousField distinguishes a missing field from a present null value.
type AnonymousField struct {
	Value any
	Found bool
}

type anonymousMap map[string]any

func (value anonymousMap) sourceObject() any { return map[string]any(value) }
func (value anonymousMap) readField(name string) AnonymousField {
	field, found := value[name]
	return AnonymousField{Value: field, Found: found}
}
func (value anonymousMap) writeField(name string, field any) bool {
	value[name] = field
	return true
}

// AnonymousFromMap admits existing backing storage without copying its fields.
// A nil map is a null source, not an allocated anonymous object.
func AnonymousFromMap(value map[string]any) Anonymous {
	if value == nil {
		return nil
	}
	return anonymousMap(value)
}

type anonymousView struct {
	source any
	read   func(string) AnonymousField
	write  func(string, any) bool
}

func (value *anonymousView) sourceObject() any { return value.source }
func (value *anonymousView) readField(name string) AnonymousField {
	return value.read(name)
}
func (value *anonymousView) writeField(name string, field any) bool {
	return value.write != nil && value.write(name, field)
}

// NewAnonymousView retains the physical source strongly. Member reads remain
// live; this constructor neither extracts methods nor snapshots mutable slots.
// Logical receiver rebinding is resolved by ObjectSource at observation time.
func NewAnonymousView(source any, read func(string) AnonymousField, write func(string, any) bool) Anonymous {
	if isAnyNil(source) {
		return nil
	}
	return &anonymousView{source: source, read: read, write: write}
}

// AnonymousLookup is the presence-preserving generated field access boundary.
func AnonymousLookup(value Anonymous, name string) AnonymousField {
	if value == nil {
		Throw(StringFromLiteral("Null Access"))
	}
	return value.readField(name)
}

// AnonymousGet reads a value; Haxe typed access treats a missing field as null.
func AnonymousGet(value Anonymous, name string) any {
	return AnonymousLookup(value, name).Value
}

// AnonymousSet writes through the original storage. Typed source admission must
// provide a writer for writable fields; a missing writer is not silent success.
func AnonymousSet(value Anonymous, name string, field any) {
	if value == nil {
		Throw(StringFromLiteral("Null Access"))
	}
	if !value.writeField(name, field) {
		Throw(StringFromLiteral("Invalid field write: " + name))
	}
}

// AnonymousFromValue recovers anonymous storage after erasure. Generated class
// admission belongs to compiler member plans; this boundary admits only values
// whose structural representation is already known.
func AnonymousFromValue(value any) Anonymous {
	if isAnyNil(value) {
		return nil
	}
	switch object := value.(type) {
	case Anonymous:
		return object
	case map[string]any:
		return AnonymousFromMap(object)
	default:
		Throw(StringFromLiteral("Invalid structural object"))
		return nil
	}
}
