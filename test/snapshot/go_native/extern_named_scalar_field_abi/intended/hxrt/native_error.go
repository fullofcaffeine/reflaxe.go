package hxrt

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

// Text explicitly observes native formatting, including its side effects and
// panic behavior. Transport never calls it. The result uses Haxe string storage.
func (value *NativeErrorValue) Text() *string {
	return StringFromLiteral(value.original.Error())
}
