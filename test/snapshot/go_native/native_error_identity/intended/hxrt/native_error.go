package hxrt

import "errors"

// RawError keeps native signatures identical to Go error even when a source
// local shadows that predeclared identifier.
type RawError = error

// NativeErrorValue preserves the original interface before Haxe generic storage
// erases its static type. A typed nil error therefore remains a non-null carrier.
// The payload is private and the carrier deliberately has no Error/Unwrap method.
type NativeErrorValue struct {
	original RawError
}

// ImportNativeError preserves nil interfaces and retains every other original
// value without invoking formatting, matching, or unwrapping methods.
func ImportNativeError(value RawError) *NativeErrorValue {
	if value == nil {
		return nil
	}
	return &NativeErrorValue{original: value}
}

// ExportNativeError returns the exact saved root. Repeated crossings do not add
// wrappers, copy the concrete error, or require it to be comparable.
func ExportNativeError(value *NativeErrorValue) RawError {
	if value == nil {
		return nil
	}
	return value.original
}

// ExportLegacyError gives Go a native error for one go.Error value. A go.Error
// is a message snapshot, so the result is a new error with the same text (from
// String, the generated adapter for the Haxe toString method). It is not the
// same value, so errors.Is cannot match two crossings; use go.NativeError when
// identity matters. A null go.Error becomes a nil interface, never a typed nil
// pointer that Go would treat as a present error. The type parameter lets the
// generated class stay in the program's package while the nil check stays exact.
func ExportLegacyError[E interface {
	comparable
	String() string
}](value E) RawError {
	var none E
	if value == none {
		return nil
	}
	return errors.New(value.String())
}

// ExportErrors converts a Haxe slice of error carriers, such as a spread Array,
// into the []error that a Go ...error parameter accepts. Each element uses the
// same export as a single argument, so null elements stay nil interfaces.
func ExportErrors[E any](values []E, export func(E) RawError) []RawError {
	natives := make([]RawError, len(values))
	for index, value := range values {
		natives[index] = export(value)
	}
	return natives
}

// Text explicitly observes native formatting, including its side effects and
// panic behavior. Transport never calls it. The result uses Haxe string storage.
func (value *NativeErrorValue) Text() *string {
	return StringFromLiteral(value.original.Error())
}
