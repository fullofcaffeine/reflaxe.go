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
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("capture="), func(hx_fn *hxrt.Function[func(int) int], hx_arg_0 int) int {
		return hxrt.FunctionCall(hx_fn)(hx_arg_0)
	}(first, 2)), hxrt.StringFromLiteral(":")), func(hx_fn *hxrt.Function[func(int) int], hx_arg_0 int) int {
		return hxrt.FunctionCall(hx_fn)(hx_arg_0)
	}(alias, 3)), hxrt.StringFromLiteral(":")), func(hx_fn *hxrt.Function[func(int) int], hx_arg_0 int) int {
		return hxrt.FunctionCall(hx_fn)(hx_arg_0)
	}(other, 1)))
	hxrt.Println(v_2)
	var record_callback *hxrt.Function[func(int) int]
	record_callback = first
	extracted := record_callback
	var boxed any = extracted
	restored := func(hx_value_1 any) *hxrt.Function[func(int) int] {
		if hx_value_1 == nil {
			var hx_zero_2 *hxrt.Function[func(int) int]
			return hx_zero_2
		}
		return hx_value_1.(*hxrt.Function[func(int) int])
	}(boxed)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("record="), func(hx_fn *hxrt.Function[func(int) int], hx_arg_0 int) int {
		return hxrt.FunctionCall(hx_fn)(hx_arg_0)
	}(record_callback, 0)), hxrt.StringFromLiteral(":")), hxrt.StdString(hxrt.ReflectCompareMethods(any(first), any(restored)))))
	hxrt.Println(v_3)
	var missing *hxrt.Function[func(int) int] = nil
	evaluated := 0
	caught := false
	hxrt.TryCatch(func() {
		func(hx_fn *hxrt.Function[func(int) int], hx_arg_0 int) int {
			return hxrt.FunctionCall(hx_fn)(hx_arg_0)
		}(missing, func() int {
			evaluated = int(int32((evaluated + 1)))
			return evaluated
		}())
	}, func(hx_caught_3 any) {
		hx_tmp := hx_caught_3
		_ = hx_tmp
		caught = true
	})
	var v_4 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("null="), hxrt.StdString(caught)), hxrt.StringFromLiteral(":")), evaluated))
	hxrt.Println(v_4)
}

func reuse(value any) any {
	return value
}
