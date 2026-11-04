package main

import (
	"regexp"
	"snapshot/hxrt"
)

func main() {
	calls := 0
	callback := hxrt.NewFunction(func(value *string) *string {
		calls = int(int32((calls + 1)))
		return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("["), value), hxrt.StringFromLiteral("]"))
	})
	expression := regexp.MustCompile(*hxrt.StdString(hxrt.StringFromLiteral("x")))
	var v any = any(hxrt.StdString(expression.ReplaceAllStringFunc(*hxrt.StdString(hxrt.StringFromLiteral("x x")), func(hx_native_callback_1 *hxrt.Function[func(*string) *string]) func(string) string {
		if hx_native_callback_1 == nil {
			return nil
		}
		return func(hx_native_arg_2 string) string {
			return *hxrt.StdString(hxrt.FunctionCall(hx_native_callback_1)(hxrt.StdString(hx_native_arg_2)))
		}
	}(callback))))
	hxrt.Println(v)
	replace := func() *hxrt.Function[func(*string, *hxrt.Function[func(*string) *string]) *string] {
		hx_extern_method_3 := expression.ReplaceAllStringFunc
		if hx_extern_method_3 == nil {
			return nil
		}
		return hxrt.NewFunction(func(hx_extern_arg_4 *string, hx_extern_arg_5 *hxrt.Function[func(*string) *string]) *string {
			return hxrt.StdString(hx_extern_method_3(*hxrt.StdString(hx_extern_arg_4), func(hx_native_callback_6 *hxrt.Function[func(*string) *string]) func(string) string {
				if hx_native_callback_6 == nil {
					return nil
				}
				return func(hx_native_arg_7 string) string {
					return *hxrt.StdString(hxrt.FunctionCall(hx_native_callback_6)(hxrt.StdString(hx_native_arg_7)))
				}
			}(hx_extern_arg_5)))
		})
	}()
	var v_1 any = any(func() *string {
		var hx_callable_8 *hxrt.Function[func(*string, *hxrt.Function[func(*string) *string]) *string] = replace
		var hx_argument_9 *string = hxrt.StringFromLiteral("x")
		var hx_argument_10 *hxrt.Function[func(*string) *string] = callback
		return hxrt.FunctionCall(hx_callable_8)(hx_argument_9, hx_argument_10)
	}())
	hxrt.Println(v_1)
	hxrt.Println(any(calls))
	var missing *hxrt.Function[func(*string) *string] = nil
	var v_2 any = any(hxrt.StdString(expression.ReplaceAllStringFunc(*hxrt.StdString(hxrt.StringFromLiteral("z")), func(hx_native_callback_11 *hxrt.Function[func(*string) *string]) func(string) string {
		if hx_native_callback_11 == nil {
			return nil
		}
		return func(hx_native_arg_12 string) string {
			return *hxrt.StdString(hxrt.FunctionCall(hx_native_callback_11)(hxrt.StdString(hx_native_arg_12)))
		}
	}(missing))))
	hxrt.Println(v_2)
	var v_3 any = any(func() *string {
		var hx_callable_13 *hxrt.Function[func(*string, *hxrt.Function[func(*string) *string]) *string] = replace
		var hx_argument_14 *string = hxrt.StringFromLiteral("z")
		var hx_argument_15 *hxrt.Function[func(*string) *string] = missing
		return hxrt.FunctionCall(hx_callable_13)(hx_argument_14, hx_argument_15)
	}())
	hxrt.Println(v_3)
}
