package hxrt

import "fmt"

func StringFromLiteral(value string) *string {
	copy := value
	return &copy
}

func StdString(value any) *string {
	switch v := value.(type) {
	case nil:
		return StringFromLiteral("null")
	case *string:
		if v == nil {
			return StringFromLiteral("null")
		}
		return v
	case string:
		return StringFromLiteral(v)
	default:
		return StringFromLiteral(fmt.Sprint(v))
	}
}

func StringConcatAny(left any, right any) *string {
	l := StdString(left)
	r := StdString(right)
	return StringFromLiteral(*l + *r)
}

func StringEqualAny(left any, right any) bool {
	l := StdString(left)
	r := StdString(right)
	return *l == *r
}

func Println(value any) {
	fmt.Println(*StdString(value))
}

type HaxeException struct {
	Value any
}

type ExceptionValue struct {
	Value   any
	Message *string
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
