package main

import "snapshot/hxrt"

func counter() *hxrt.Function[func(int) int] {
	total := 0
	return hxrt.NewFunction(func(delta int) int {
		return func() int {
			total = int((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(delta)))
			return total
		}()
	})
}

func main() {
	first := counter()
	alias := first
	other := counter()
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("identity="), hxrt.StdString(hxrt.ReflectCompareMethods(any(first), any(alias)))), hxrt.StringFromLiteral(":")), hxrt.StdString(hxrt.ReflectCompareMethods(any(first), any(other)))))
	hxrt.Println(v)
	var stored_2 *hxrt.Function[func(int) int]
	var stored_1 *hxrt.Function[func(int) int]
	var stored_0 *hxrt.Function[func(int) int]
	stored_0 = first
	stored_1 = alias
	stored_2 = other
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("transport="), hxrt.StdString(func() bool {
		var f2 any = any(reuse(stored_1))
		return hxrt.ReflectCompareMethods(any(stored_0), f2)
	}())), hxrt.StringFromLiteral(":")), hxrt.StdString(hxrt.ReflectCompareMethods(any(stored_0), any(stored_2)))))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("capture="), func() int {
		var hx_callable_1 *hxrt.Function[func(int) int] = first
		var hx_argument_2 int = 2
		return hxrt.FunctionCall(hx_callable_1)(hx_argument_2)
	}()), hxrt.StringFromLiteral(":")), func() int {
		var hx_callable_3 *hxrt.Function[func(int) int] = alias
		var hx_argument_4 int = 3
		return hxrt.FunctionCall(hx_callable_3)(hx_argument_4)
	}()), hxrt.StringFromLiteral(":")), func() int {
		var hx_callable_5 *hxrt.Function[func(int) int] = other
		var hx_argument_6 int = 1
		return hxrt.FunctionCall(hx_callable_5)(hx_argument_6)
	}()))
	hxrt.Println(v_2)
	var record_callback *hxrt.Function[func(int) int]
	record_callback = first
	extracted := record_callback
	var boxed any = extracted
	restored := func(hx_value_7 any) *hxrt.Function[func(int) int] {
		if hx_value_7 == nil {
			var hx_zero_8 *hxrt.Function[func(int) int]
			return hx_zero_8
		}
		return hx_value_7.(*hxrt.Function[func(int) int])
	}(boxed)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("record="), func() int {
		var hx_callable_9 *hxrt.Function[func(int) int] = record_callback
		var hx_argument_10 int = 0
		return hxrt.FunctionCall(hx_callable_9)(hx_argument_10)
	}()), hxrt.StringFromLiteral(":")), hxrt.StdString(hxrt.ReflectCompareMethods(any(first), any(restored)))))
	hxrt.Println(v_3)
	var missing *hxrt.Function[func(int) int] = nil
	evaluated := 0
	caught := false
	hxrt.TryCatch(func() {
		func() int {
			var hx_callable_13 *hxrt.Function[func(int) int] = missing
			var hx_argument_14 int = func() int {
				evaluated = int(int32((evaluated + 1)))
				return evaluated
			}()
			return hxrt.FunctionCall(hx_callable_13)(hx_argument_14)
		}()
	}, func(hx_caught_11 any) {
		hx_tmp := hx_caught_11
		_ = hx_tmp
		caught = true
	})
	var v_4 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("null="), hxrt.StdString(caught)), hxrt.StringFromLiteral(":")), evaluated))
	hxrt.Println(v_4)
}

func reuse(value any) any {
	return value
}
