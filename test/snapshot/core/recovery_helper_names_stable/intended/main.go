package main

import "snapshot/hxrt"

func main() {
	var v any = any(Edited_run(New_Square(2), New_Circle(3), hxrt.NewFunction(func(text *string) int {
		return hxrt.StringLengthStringPtr(text)
	})))
	hxrt.Println(v)
	var v_1 any = any(Stable_run(New_Circle(4), hxrt.NewFunction(func(size int) int {
		return int((hxrt.Int32Wrap(size) * hxrt.Int32Wrap(10)))
	})))
	hxrt.Println(v_1)
}

func hx_recover_callable_408c8b7a(value any) *hxrt.Function[func(int) int] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(int) int]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(*string) int]:
		return hxrt.ProjectFunction(source, func(arg_0 int) int {
			return hxrt.FunctionCall(source)(func(hx_value_1 any) *string {
				if hx_value_1 == nil {
					var hx_zero_2 *string
					return hx_zero_2
				}
				return hx_value_1.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(int) int]:
		return source
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_class_Circle(value any) *Circle {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *Circle:
		return source
	default:
		return value.(*Circle)
	}
}

func hx_recover_class_Square(value any) *Square {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *Square:
		return source
	default:
		return value.(*Square)
	}
}

func hx_is_nominal_Circle(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *Circle:
		return (source != nil)
	default:
		return false
	}
}
