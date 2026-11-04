package main

import "snapshot/hxrt"

func absent() any {
	return nil
}

func main() {
	callback := hx_recover_callable_0(absent())
	hxrt.Println(any((callback == nil)))
	var erased any = absent()
	hxrt.TryCatch(func() {
		reflaxe__go___internal__CompilerReflect_callMethod(erased, []any{})
		hxrt.Println(any(hxrt.StringFromLiteral("missed null call")))
	}, func(hx_caught_1 any) {
		error := hx_caught_1
		hxrt.Println(error)
	})
}

func reflaxe__go___internal__CompilerReflect_callMethod(value any, arguments []any) any {
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Cannot call null"))
		return nil
	}
	return hxrt.ReflectCallMethod(value, arguments)
}

func hx_recover_callable_0(value any) *hxrt.Function[func() int] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() int]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
	return nil
}
