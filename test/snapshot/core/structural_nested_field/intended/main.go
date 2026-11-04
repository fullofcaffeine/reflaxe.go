package main

import "snapshot/hxrt"

func main() {
	hx_obj_1 := map[string]any{}
	hx_obj_2 := map[string]any{}
	hx_obj_2["next"] = hxrt.NewFunction(func() int {
		return 3
	})
	hx_obj_1["inner"] = hxrt.AnonymousFromMap(hx_obj_2)
	source := hxrt.AnonymousFromMap(hx_obj_1)
	func() hxrt.Anonymous {
		hx_structural_target_3 := source
		hx_structural_value_5 := source
		hxrt.AnonymousSet(hx_structural_target_3, "link", hx_structural_value_5)
		return hx_structural_value_5
	}()
	view := through(source)
	var v any = any(func() int {
		var hx_callable_6 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(hx_structural_recover(hxrt.AnonymousGet(view, "inner")), "next"))
		return hxrt.FunctionCall(hx_callable_6)()
	}())
	hxrt.Println(v)
	var v_1 any = any(func(hx_value_7 any) int {
		if hx_value_7 == nil {
			var hx_zero_8 int
			return hx_zero_8
		}
		return hx_value_7.(int)
	}(read(source)))
	hxrt.Println(v_1)
	func() *hxrt.Function[func() int] {
		hx_structural_target_9 := hx_structural_recover(hxrt.AnonymousGet(source, "inner"))
		hx_structural_value_11 := hxrt.NewFunction(func() int {
			return 9
		})
		hxrt.AnonymousSet(hx_structural_target_9, "next", hx_structural_value_11)
		return hx_structural_value_11
	}()
	var v_2 any = any(func() int {
		var hx_callable_12 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(hx_structural_recover(hxrt.AnonymousGet(view, "inner")), "next"))
		return hxrt.FunctionCall(hx_callable_12)()
	}())
	hxrt.Println(v_2)
	var v_3 any = any(func(hx_value_13 any) int {
		if hx_value_13 == nil {
			var hx_zero_14 int
			return hx_zero_14
		}
		return hx_value_13.(int)
	}(read(source)))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.HaxeEqual(hx_structural_recover(hxrt.AnonymousGet(view, "link")), source))
	hxrt.Println(v_4)
	var v_5 any = any(func(hx_value_15 any) int {
		if hx_value_15 == nil {
			var hx_zero_16 int
			return hx_zero_16
		}
		return hx_value_15.(int)
	}(read(hx_structural_recover(hxrt.AnonymousGet(view, "link")))))
	hxrt.Println(v_5)
}

func read(value hxrt.Anonymous) any {
	return func() any {
		var hx_callable_17 *hxrt.Function[func() any] = hx_recover_callable_1(hxrt.AnonymousGet(hx_structural_recover(hxrt.AnonymousGet(value, "inner")), "next"))
		return hxrt.FunctionCall(hx_callable_17)()
	}()
}

func through(value hxrt.Anonymous) hxrt.Anonymous {
	return value
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_1 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_1.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_1.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_1.high = func(hx_value_2 any) int {
						if hx_value_2 == nil {
							var hx_zero_3 int
							return hx_zero_3
						}
						return hx_value_2.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_1.low = func(hx_value_4 any) int {
						if hx_value_4 == nil {
							var hx_zero_5 int
							return hx_zero_5
						}
						return hx_value_4.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	default:
		return hxrt.AnonymousFromValue(value)
	}
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
	switch source := origin.(type) {
	case *hxrt.Function[func() int]:
		return source
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_1(value any) *hxrt.Function[func() any] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() any]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}
