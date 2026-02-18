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
