package hxrt

type HaxeException struct {
	Value any
}

type ExceptionValue struct {
	Value   any
	Message *string
}

// NewValueException keeps direct haxe.ValueException construction on the same
// runtime carrier used by catch/wildcard exception paths. The current portable
// contract only needs value + message parity here; broader previous/native
// object-model parity is tracked separately in the direct exception-subclass
// follow-up work.
func NewValueException(value any, previous *ExceptionValue, native any) *ExceptionValue {
	_ = previous
	_ = native
	return &ExceptionValue{
		Value:   value,
		Message: StdString(value),
	}
}

func Throw(value any) {
	panic(HaxeException{Value: value})
}

func UnwrapException(recovered any) any {
	switch v := recovered.(type) {
	case HaxeException:
		return v.Value
	case *HaxeException:
		return v.Value
	default:
		return v
	}
}

func TryCatch(tryFn func(), catchFn func(any)) {
	defer func() {
		if recovered := recover(); recovered != nil {
			catchFn(UnwrapException(recovered))
		}
	}()
	tryFn()
}

func ExceptionCaught(value any) *ExceptionValue {
	switch v := value.(type) {
	case *ExceptionValue:
		return v
	case ExceptionValue:
		copy := v
		return &copy
	default:
		return &ExceptionValue{
			Value:   value,
			Message: StdString(value),
		}
	}
}

func ExceptionThrown(value any) any {
	return ExceptionCaught(value).Value
}

func ExceptionMessage(value any) *string {
	return ExceptionCaught(value).Message
}
