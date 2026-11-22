package main

import "snapshot/hxrt"

var getters *hxrt.Array = func() *hxrt.Array {
	_g := hxrt.NewArray()
	name := hxrt.StringFromLiteral("first")
	_g.Push(hxrt.NewFunction(func() *string {
		return name
	}))
	name_1 := hxrt.StringFromLiteral("second")
	_g.Push(hxrt.NewFunction(func() *string {
		return name_1
	}))
	return _g
}()

func main() {
	_g := 0
	_g1 := _Main__Main_Fields__moduleFields
	for _g < _g1.Len() {
		field := hx_structural_recover(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		var v any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("module="), func(hx_value_3 any) *string {
			if hx_value_3 == nil {
				var hx_zero_4 *string
				return hx_zero_4
			}
			return hx_value_3.(*string)
		}(hxrt.AnonymousGet(field, "tag"))), hxrt.StringFromLiteral(":")), func(hx_value_5 any) *string {
			if hx_value_5 == nil {
				var hx_zero_6 *string
				return hx_zero_6
			}
			return hx_value_5.(*string)
		}(hxrt.AnonymousGet(field, "kind"))))
		hxrt.Println(v)
	}
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("nested="), hxrt.StringJoinAny(nestedPairs.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("getters="), hxrt.StringJoinAny(func() *hxrt.Array {
		_g_1 := hxrt.NewArray()
		_g1_1 := 0
		_g2 := getters
		for _g1_1 < _g2.Len() {
			getter := hx_recover_callable_b9d17ba0(_g2.Get(_g1_1))
			_g1_1 = int(int32((_g1_1 + 1)))
			_g_1.Push(func() *string {
				var hx_callable_8 *hxrt.Function[func() *string] = getter
				return hxrt.FunctionCall(hx_callable_8)()
			}())
		}
		return _g_1
	}().Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("mutated="), hxrt.StringJoinAny(func() *hxrt.Array {
		_g_2 := hxrt.NewArray()
		_g1_2 := 0
		_g2_1 := mutatedGetters
		for _g1_2 < _g2_1.Len() {
			getter_1 := hx_recover_callable_e273723c(_g2_1.Get(_g1_2))
			_g1_2 = int(int32((_g1_2 + 1)))
			_g_2.Push(func() int {
				var hx_callable_10 *hxrt.Function[func() int] = getter_1
				return hxrt.FunctionCall(hx_callable_10)()
			}())
		}
		return _g_2
	}().Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_3)
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("shadowed="), shadowedBlocks)))
}

var mutatedGetters *hxrt.Array = func() *hxrt.Array {
	out := hxrt.NewArray()
	x := 1
	x = int((hxrt.Int32Wrap(x) + hxrt.Int32Wrap(10)))
	out.Push(hxrt.NewFunction(func() int {
		return x
	}))
	x_1 := 2
	x_1 = int((hxrt.Int32Wrap(x_1) + hxrt.Int32Wrap(10)))
	out.Push(hxrt.NewFunction(func() int {
		return x_1
	}))
	return out
}()

var nestedPairs *hxrt.Array = func() *hxrt.Array {
	_g := hxrt.NewArray()
	x := 1
	y := hxrt.StringFromLiteral("p")
	_g.Push(hxrt.StringConcatAny(x, y))
	y_1 := hxrt.StringFromLiteral("q")
	_g.Push(hxrt.StringConcatAny(x, y_1))
	x_1 := 2
	y_2 := hxrt.StringFromLiteral("p")
	_g.Push(hxrt.StringConcatAny(x_1, y_2))
	y_3 := hxrt.StringFromLiteral("q")
	_g.Push(hxrt.StringConcatAny(x_1, y_3))
	return _g
}()

var shadowedBlocks *string = func() *string {
	out := hxrt.NewArray()
	x := hxrt.StringFromLiteral("left")
	out.Push(x)
	x_1 := hxrt.StringFromLiteral("right")
	out.Push(x_1)
	return hxrt.StringJoinAny(out.Values(), hxrt.StringFromLiteral("+"))
}()

var _Main__Main_Fields__moduleFields *hxrt.Array = func() *hxrt.Array {
	_g := hxrt.NewArray()
	tag := hxrt.StringFromLiteral("a")
	hx_obj_20 := map[string]any{}
	hx_obj_20["tag"] = tag
	hx_obj_20["kind"] = hxrt.StringFromLiteral("text")
	_g.Push(hxrt.AnonymousFromMap(hx_obj_20))
	tag_1 := hxrt.StringFromLiteral("b")
	hx_obj_22 := map[string]any{}
	hx_obj_22["tag"] = tag_1
	hx_obj_22["kind"] = hxrt.StringFromLiteral("text")
	_g.Push(hxrt.AnonymousFromMap(hx_obj_22))
	return _g
}()

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

func hx_recover_callable_b9d17ba0(value any) *hxrt.Function[func() *string] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() *string]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() *string]:
		return source
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_1 any) *string {
				if hx_value_1 == nil {
					var hx_zero_2 *string
					return hx_zero_2
				}
				return hx_value_1.(*string)
			}(hxrt.FunctionCall(source)())
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_e273723c(value any) *hxrt.Function[func() int] {
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
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_1 any) int {
				if hx_value_1 == nil {
					var hx_zero_2 int
					return hx_zero_2
				}
				return hx_value_1.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() int]:
		return source
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}
