package main

import "snapshot/hxrt"

func main() {
	var v any = any(Registry_count)
	hxrt.Println(v)
}

func reflaxe__go___internal__CompilerReflect_callMethod(value any, arguments []any) any {
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Cannot call null"))
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(any) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(argument_0)
	}
	hx_recovered, hx_matched := reflaxe__go___internal__CompilerReflect_callMethod__probe_0(origin, arguments)
	if hx_matched {
		return hx_recovered
	}
	hx_recovered, hx_matched = reflaxe__go___internal__CompilerReflect_callMethod__probe_1(origin, arguments)
	if hx_matched {
		return hx_recovered
	}
	hx_recovered, hx_matched = reflaxe__go___internal__CompilerReflect_callMethod__probe_2(origin, arguments)
	if hx_matched {
		return hx_recovered
	}
	hx_recovered, hx_matched = reflaxe__go___internal__CompilerReflect_callMethod__probe_3(origin, arguments)
	if hx_matched {
		return hx_recovered
	}
	return hxrt.ReflectCallMethod(value, arguments)
}

func hx_recover_callable_0(value any) *hxrt.Function[func(any) int] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(any) int]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(any) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 any) int {
			return func(hx_value_35 any) int {
				if hx_value_35 == nil {
					var hx_zero_36 int
					return hx_zero_36
				}
				return hx_value_35.(int)
			}(hxrt.FunctionCall(source)(arg_0))
		})
	}
	hx_recovered, hx_matched := hx_recover_callable_0__probe_0(origin)
	if hx_matched {
		return hx_recovered
	}
	hx_recovered, hx_matched = hx_recover_callable_0__probe_1(origin)
	if hx_matched {
		return hx_recovered
	}
	hx_recovered, hx_matched = hx_recover_callable_0__probe_2(origin)
	if hx_matched {
		return hx_recovered
	}
	hx_recovered, hx_matched = hx_recover_callable_0__probe_3(origin)
	if hx_matched {
		return hx_recovered
	}
	hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
	return nil
}

func hx_recover_callable_1(value any) *hxrt.Function[func(*hxrt.Function[func(any) *string]) *string] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(*hxrt.Function[func(any) *string]) *string]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(any) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 *hxrt.Function[func(any) *string]) *string {
			return hxrt.FunctionCall(source)(arg_0)
		})
	}
	hx_recovered, hx_matched := hx_recover_callable_1__probe_0(origin)
	if hx_matched {
		return hx_recovered
	}
	hx_recovered, hx_matched = hx_recover_callable_1__probe_1(origin)
	if hx_matched {
		return hx_recovered
	}
	hx_recovered, hx_matched = hx_recover_callable_1__probe_2(origin)
	if hx_matched {
		return hx_recovered
	}
	hx_recovered, hx_matched = hx_recover_callable_1__probe_3(origin)
	if hx_matched {
		return hx_recovered
	}
	hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
	return nil
}

func hx_recover_class_0(value any) *haxe__ValueException {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ValueException:
		return source
	default:
		return value.(*haxe__ValueException)
	}
}

func hx_recover_class_1(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__Exception:
		return source
	case *haxe__ValueException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}

func hx_is_nominal_0(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_1(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__Exception:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}
