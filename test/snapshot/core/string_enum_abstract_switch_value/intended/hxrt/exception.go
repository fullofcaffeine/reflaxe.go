package hxrt

import (
	"fmt"
	"os"
)

type ExceptionValue struct {
	Value   any
	Message *string
}

type ExceptionCarrier interface {
	HxExceptionValue() *ExceptionValue
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

func BindException(value any, message *string, previous *ExceptionValue, native any) *ExceptionValue {
	_ = previous
	_ = native
	return &ExceptionValue{
		Value:   value,
		Message: message,
	}
}

func unwrapHaxeException(recovered any) (any, bool) {
	switch v := recovered.(type) {
	case HaxeException:
		return v.Value, true
	case *HaxeException:
		return v.Value, true
	default:
		return nil, false
	}
}

// UnwrapException converts only the carrier created by Throw. A foreign Go
// panic is a backend/native failure, not a Haxe value, so it resumes unwinding
// unchanged instead of entering a Haxe catch clause.
func UnwrapException(recovered any) any {
	value, ok := unwrapHaxeException(recovered)
	if !ok {
		panic(recovered)
	}
	return value
}

func TryCatch(tryFn func(), catchFn func(any)) {
	defer func() {
		if recovered := recover(); recovered != nil {
			catchFn(UnwrapException(recovered))
		}
	}()
	tryFn()
}

// ReportUncaughtException is the stable portable-thread failure report. Native
// Go panics never reach this function; their ordinary Go panic output remains
// intact for explicit Go-native boundaries.
func ReportUncaughtException(value any) {
	message := ExceptionMessage(value)
	fmt.Fprintf(os.Stderr, "Uncaught exception %s\n", *StdString(message))
}

func ExceptionCaught(value any) *ExceptionValue {
	switch v := value.(type) {
	case *ExceptionValue:
		return v
	case ExceptionValue:
		copy := v
		return &copy
	case ExceptionCarrier:
		return v.HxExceptionValue()
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
