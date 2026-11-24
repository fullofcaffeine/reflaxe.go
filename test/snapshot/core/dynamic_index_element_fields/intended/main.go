package main

import "snapshot/hxrt"

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

func attempt(name *string, run *hxrt.Function[func() any]) {
	var hx_try_1 *string
	hxrt.TryCatch(func() {
		hx_try_1 = hxrt.StdString(func() any {
			var hx_callable_2 *hxrt.Function[func() any] = run
			return hxrt.FunctionCall(hx_callable_2)()
		}())
	}, func(hx_caught_3 any) {
		hx_exception_4 := haxe__Exception_caught(hx_caught_3)
		hx_payload_5 := haxe__Exception_catchValue(hx_caught_3)
		_ = hx_exception_4
		var e any = hx_payload_5
		hx_try_1 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("error "), hxrt.StdString(e))
	})
	text := hx_try_1
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(name, hxrt.StringFromLiteral(" = ")), text)))
}

func dyn(value any) any {
	return value
}

func evaluationOrder() {
	var rows any = dyn(hxrt.NewArray(hxrt.NewArray(1), hxrt.NewArray(2)))
	reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(step(hxrt.StringFromLiteral("r"), rows), hxrt.IntFromNullableAny(step(hxrt.StringFromLiteral("i"), 0)))), hxrt.StringFromLiteral("push")), []any{step(hxrt.StringFromLiteral("a"), 5)})
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("call order "), order), hxrt.StringFromLiteral(" ")), hxrt.StdString(rows)))
	hxrt.Println(v)
	order = hxrt.StringFromLiteral("")
	read := hxrt.StdString(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(step(hxrt.StringFromLiteral("r"), rows), hxrt.IntFromNullableAny(step(hxrt.StringFromLiteral("i"), 1)))), hxrt.StringFromLiteral("length")))
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("read order "), order), hxrt.StringFromLiteral(" ")), read))
	hxrt.Println(v_1)
	order = hxrt.StringFromLiteral("")
	var objs any = dyn(hxrt.NewArray(func() hxrt.Anonymous {
		hx_obj_6 := map[string]any{}
		hx_obj_6["n"] = 1
		return hxrt.AnonymousFromMap(hx_obj_6)
	}()))
	fh := hx_structural_recover(hxrt.DynamicIndexGet(step(hxrt.StringFromLiteral("r"), objs), hxrt.IntFromNullableAny(step(hxrt.StringFromLiteral("i"), 0))))
	func() any {
		hx_structural_target_7 := fh
		hx_structural_old_8 := hxrt.AnonymousGet(hx_structural_target_7, "n")
		hx_structural_value_9 := hxrt.DynamicAdd(hx_structural_old_8, step(hxrt.StringFromLiteral("v"), 2))
		hxrt.AnonymousSet(hx_structural_target_7, "n", hx_structural_value_9)
		return hx_structural_value_9
	}()
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("write order "), order), hxrt.StringFromLiteral(" ")), Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(objs, 0)), hxrt.StringFromLiteral("n"))))
	hxrt.Println(v_2)
}

func fails(name *string, run *hxrt.Function[func() any]) {
	var hx_try_10 *string
	hxrt.TryCatch(func() {
		func() any {
			var hx_callable_11 *hxrt.Function[func() any] = run
			return hxrt.FunctionCall(hx_callable_11)()
		}()
		hx_try_10 = hxrt.StringFromLiteral("no error")
	}, func(hx_caught_12 any) {
		hx_exception_13 := haxe__Exception_caught(hx_caught_12)
		hx_payload_14 := haxe__Exception_catchValue(hx_caught_12)
		_ = hx_exception_13
		var e any = hx_payload_14
		_ = e
		hx_try_10 = hxrt.StringFromLiteral("error")
	})
	text := hx_try_10
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(name, hxrt.StringFromLiteral(" = ")), text)))
}

func lengths() {
	var words any = dyn(hxrt.NewArray(hxrt.StringFromLiteral("alpha"), hxrt.StringFromLiteral("be"), hxrt.StringFromLiteral("")))
	var rows any = dyn(hxrt.NewArray(hxrt.NewArray(1, 2), hxrt.NewArray(3), hxrt.NewArray()))
	var v any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("length "), Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(words, 0)), hxrt.StringFromLiteral("length"))), hxrt.StringFromLiteral(" ")), Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(words, 2)), hxrt.StringFromLiteral("length"))), hxrt.StringFromLiteral(" ")), Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(rows, 0)), hxrt.StringFromLiteral("length"))), hxrt.StringFromLiteral(" ")), Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(rows, 2)), hxrt.StringFromLiteral("length"))))
	hxrt.Println(v)
	var v_1 any = Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(words, 1)), hxrt.StringFromLiteral("length"))
	hxrt.Println(v_1)
	n := hxrt.IntFromNullableAny(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(words, 0)), hxrt.StringFromLiteral("length")))
	f := hxrt.DynamicToFloat(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(rows, 0)), hxrt.StringFromLiteral("length")))
	text := hxrt.StdString(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(words, 0)), hxrt.StringFromLiteral("length")))
	var maybe any = Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(rows, 1)), hxrt.StringFromLiteral("length"))
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("typed "), int((hxrt.Int32Wrap(n)*hxrt.Int32Wrap(2)))), hxrt.StringFromLiteral(" ")), (f/float64(4))), hxrt.StringFromLiteral(" ")), text), hxrt.StringFromLiteral(" ")), maybe)))
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("operators "), hxrt.DynamicToFloat(hxrt.DynamicAdd(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(words, 0)), hxrt.StringFromLiteral("length")), 1))), hxrt.StringFromLiteral(" ")), hxrt.DynamicToFloat(hxrt.DynamicSub(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(words, 0)), hxrt.StringFromLiteral("length")), Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(rows, 0)), hxrt.StringFromLiteral("length"))))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.DynamicGreater(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(words, 0)), hxrt.StringFromLiteral("length")), 3))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.HaxeEqual(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(rows, 1)), hxrt.StringFromLiteral("length")), 1))), hxrt.StringFromLiteral(" ")), hxrt.DynamicDiv(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(words, 0)), hxrt.StringFromLiteral("length")), 2)), hxrt.StringFromLiteral(" ")), hxrt.DynamicToFloat(hxrt.DynamicNeg(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(rows, 0)), hxrt.StringFromLiteral("length"))))))
	hxrt.Println(v_2)
	if hxrt.HaxeEqual(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(rows, 1)), hxrt.StringFromLiteral("length")), 1) {
		hxrt.Println(any(hxrt.StringFromLiteral("condition one")))
	}
	var total any = 0
	_g := 0
	_g1 := hxrt.IntFromNullableAny(Reflect_field(rows, hxrt.StringFromLiteral("length")))
	for _g < _g1 {
		hx_post_15 := _g
		_g = int(int32((_g + 1)))
		k := hx_post_15
		total = hxrt.DynamicAdd(total, Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(rows, k)), hxrt.StringFromLiteral("length")))
	}
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("loop "), hxrt.StdString(total)), hxrt.StringConcatAny(hxrt.StringFromLiteral(" "), Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(words, 1)), hxrt.StringFromLiteral("length")))))
	hxrt.Println(v_3)
	func() any {
		hx_dynamic_target_16 := hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(rows, 0))
		hx_dynamic_value_18 := 1
		Reflect_setField(hx_dynamic_target_16, hxrt.StringFromLiteral("length"), hx_dynamic_value_18)
		return hx_dynamic_value_18
	}()
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("length write "), hxrt.DynamicIndexGet(rows, 0))))
}

func main() {
	lengths()
	methods()
	structures()
	mutations()
	parsed()
	nulls()
	evaluationOrder()
}

func methods() {
	var words any = dyn(hxrt.NewArray(hxrt.StringFromLiteral("alpha"), hxrt.StringFromLiteral("be")))
	var rows any = dyn(hxrt.NewArray(hxrt.NewArray(1, 2), hxrt.NewArray(3)))
	reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(rows, 0)), hxrt.StringFromLiteral("push")), []any{9})
	reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(rows, 1)), hxrt.StringFromLiteral("unshift")), []any{0})
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("statements "), hxrt.StdString(rows)))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("push "), reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(rows, 0)), hxrt.StringFromLiteral("push")), []any{10})), hxrt.StringFromLiteral(" pop ")), reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(rows, 0)), hxrt.StringFromLiteral("pop")), []any{})), hxrt.StringFromLiteral(" ")), hxrt.DynamicIndexGet(rows, 0)))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("string "), reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(words, 1)), hxrt.StringFromLiteral("charAt")), []any{0})), hxrt.StringFromLiteral(" ")), reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(words, 0)), hxrt.StringFromLiteral("indexOf")), []any{hxrt.StringFromLiteral("p")})), hxrt.StringFromLiteral(" ")), reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(words, 0)), hxrt.StringFromLiteral("split")), []any{hxrt.StringFromLiteral("l")})), hxrt.StringFromLiteral(" ")), reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(words, 0)), hxrt.StringFromLiteral("substr")), []any{1, 2})))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("join "), reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(rows, 1)), hxrt.StringFromLiteral("join")), []any{hxrt.StringFromLiteral("-")})), hxrt.StringFromLiteral(" ")), reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(rows, 0)), hxrt.StringFromLiteral("indexOf")), []any{9})), hxrt.StringFromLiteral(" ")), reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(rows, 0)), hxrt.StringFromLiteral("contains")), []any{2})))
	hxrt.Println(v_3)
	c := hxrt.StdString(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(words, 0)), hxrt.StringFromLiteral("charAt")), []any{1}))
	size := hxrt.IntFromNullableAny(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(rows, 0)), hxrt.StringFromLiteral("push")), []any{11}))
	var v_4 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("typed "), hxrt.StringToUpperCaseStringPtr(c)), hxrt.StringFromLiteral(" ")), int((hxrt.Int32Wrap(size) + hxrt.Int32Wrap(1)))))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("chain "), Reflect_field(hxrt.DynamicFieldReceiver(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(words, 0)), hxrt.StringFromLiteral("toUpperCase")), []any{})), hxrt.StringFromLiteral("length"))), hxrt.StringFromLiteral(" ")), reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(hxrt.DynamicFieldReceiver(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(words, 0)), hxrt.StringFromLiteral("toUpperCase")), []any{})), hxrt.StringFromLiteral("charAt")), []any{0})), hxrt.StringFromLiteral(" ")), Reflect_field(hxrt.DynamicFieldReceiver(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(rows, 0)), hxrt.StringFromLiteral("concat")), []any{hxrt.NewArray(7)})), hxrt.StringFromLiteral("length"))))
	hxrt.Println(v_5)
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("call result "), hxrt.IntFromNullableAny(hxrt.DynamicAdd(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(rows, 1)), hxrt.StringFromLiteral("pop")), []any{}), 1))), hxrt.StringFromLiteral(" ")), hxrt.StringConcatAny(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(words, 1)), hxrt.StringFromLiteral("charAt")), []any{1}), hxrt.StringFromLiteral("!"))))
	hxrt.Println(v_6)
	attempt(hxrt.StringFromLiteral("missing method"), hxrt.NewFunction(func() any {
		return reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(words, 0)), hxrt.StringFromLiteral("nope")), []any{})
	}))
}

func mutations() {
	var objs any = dyn(hxrt.NewArray(func() hxrt.Anonymous {
		hx_obj_19 := map[string]any{}
		hx_obj_19["name"] = hxrt.StringFromLiteral("ann")
		hx_obj_19["n"] = 1
		return hxrt.AnonymousFromMap(hx_obj_19)
	}(), New_Point(4)))
	func() any {
		hx_dynamic_target_20 := hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(objs, 0))
		hx_dynamic_value_22 := hxrt.StringFromLiteral("bob")
		Reflect_setField(hx_dynamic_target_20, hxrt.StringFromLiteral("name"), hx_dynamic_value_22)
		return hx_dynamic_value_22
	}()
	fh := hx_structural_recover(hxrt.DynamicIndexGet(objs, 0))
	func() int {
		hx_structural_target_23 := fh
		hx_structural_old_24 := func(hx_value_26 any) int {
			if hx_value_26 == nil {
				var hx_zero_27 int
				return hx_zero_27
			}
			return hx_value_26.(int)
		}(hxrt.AnonymousGet(hx_structural_target_23, "n"))
		hx_structural_value_25 := int(int32((hx_structural_old_24 + 1)))
		hxrt.AnonymousSet(hx_structural_target_23, "n", hx_structural_value_25)
		return hx_structural_old_24
	}()
	fh_1 := hx_structural_recover(hxrt.DynamicIndexGet(objs, 0))
	func() int {
		hx_structural_target_28 := fh_1
		hx_structural_old_29 := func(hx_value_31 any) int {
			if hx_value_31 == nil {
				var hx_zero_32 int
				return hx_zero_32
			}
			return hx_value_31.(int)
		}(hxrt.AnonymousGet(hx_structural_target_28, "n"))
		hx_structural_value_30 := int((hxrt.Int32Wrap(hx_structural_old_29) + hxrt.Int32Wrap(5)))
		hxrt.AnonymousSet(hx_structural_target_28, "n", hx_structural_value_30)
		return hx_structural_value_30
	}()
	fh_2 := hx_structural_recover(hxrt.DynamicIndexGet(objs, 0))
	func() *string {
		hx_structural_target_33 := fh_2
		hx_structural_old_34 := func(hx_value_36 any) *string {
			if hx_value_36 == nil {
				var hx_zero_37 *string
				return hx_zero_37
			}
			return hx_value_36.(*string)
		}(hxrt.AnonymousGet(hx_structural_target_33, "name"))
		hx_structural_value_35 := hxrt.StringConcatStringPtr(hx_structural_old_34, hxrt.StringFromLiteral("!"))
		hxrt.AnonymousSet(hx_structural_target_33, "name", hx_structural_value_35)
		return hx_structural_value_35
	}()
	func() any {
		hx_dynamic_target_38 := hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(objs, 1))
		hx_dynamic_value_40 := 9
		Reflect_setField(hx_dynamic_target_38, hxrt.StringFromLiteral("x"), hx_dynamic_value_40)
		return hx_dynamic_value_40
	}()
	fh_3 := hx_structural_recover(hxrt.DynamicIndexGet(objs, 1))
	func() int {
		hx_structural_target_41 := fh_3
		hx_structural_old_42 := func(hx_value_44 any) int {
			if hx_value_44 == nil {
				var hx_zero_45 int
				return hx_zero_45
			}
			return hx_value_44.(int)
		}(hxrt.AnonymousGet(hx_structural_target_41, "x"))
		hx_structural_value_43 := int((hxrt.Int32Wrap(hx_structural_old_42) * hxrt.Int32Wrap(2)))
		hxrt.AnonymousSet(hx_structural_target_41, "x", hx_structural_value_43)
		return hx_structural_value_43
	}()
	var v any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("writes "), Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(objs, 0)), hxrt.StringFromLiteral("name"))), hxrt.StringFromLiteral(" ")), Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(objs, 0)), hxrt.StringFromLiteral("n"))), hxrt.StringFromLiteral(" ")), Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(objs, 1)), hxrt.StringFromLiteral("x"))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("values "), func() int {
		fh_4 := hx_structural_recover(hxrt.DynamicIndexGet(objs, 0))
		return func() int {
			hx_structural_target_46 := fh_4
			hx_structural_old_47 := func(hx_value_49 any) int {
				if hx_value_49 == nil {
					var hx_zero_50 int
					return hx_zero_50
				}
				return hx_value_49.(int)
			}(hxrt.AnonymousGet(hx_structural_target_46, "n"))
			hx_structural_value_48 := int(int32((hx_structural_old_47 + 1)))
			hxrt.AnonymousSet(hx_structural_target_46, "n", hx_structural_value_48)
			return hx_structural_old_47
		}()
	}()), hxrt.StringFromLiteral(" ")), func() int {
		fh_5 := hx_structural_recover(hxrt.DynamicIndexGet(objs, 0))
		return func() int {
			hx_structural_target_51 := fh_5
			hx_structural_old_52 := func(hx_value_54 any) int {
				if hx_value_54 == nil {
					var hx_zero_55 int
					return hx_zero_55
				}
				return hx_value_54.(int)
			}(hxrt.AnonymousGet(hx_structural_target_51, "n"))
			hx_structural_value_53 := int(int32((hx_structural_old_52 + 1)))
			hxrt.AnonymousSet(hx_structural_target_51, "n", hx_structural_value_53)
			return hx_structural_value_53
		}()
	}()), hxrt.StringFromLiteral(" ")), func() int {
		fh_6 := hx_structural_recover(hxrt.DynamicIndexGet(objs, 0))
		return func() int {
			hx_structural_target_56 := fh_6
			hx_structural_old_57 := func(hx_value_59 any) int {
				if hx_value_59 == nil {
					var hx_zero_60 int
					return hx_zero_60
				}
				return hx_value_59.(int)
			}(hxrt.AnonymousGet(hx_structural_target_56, "n"))
			hx_structural_value_58 := int((hxrt.Int32Wrap(hx_structural_old_57) - hxrt.Int32Wrap(2)))
			hxrt.AnonymousSet(hx_structural_target_56, "n", hx_structural_value_58)
			return hx_structural_value_58
		}()
	}()), hxrt.StringFromLiteral(" ")), hxrt.StdString(func() any {
		hx_dynamic_target_61 := hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(objs, 0))
		hx_dynamic_value_63 := hxrt.StringFromLiteral("cy")
		Reflect_setField(hx_dynamic_target_61, hxrt.StringFromLiteral("name"), hx_dynamic_value_63)
		return hx_dynamic_value_63
	}())))
	hxrt.Println(v_1)
	func() any {
		hx_dynamic_target_64 := hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(objs, 0))
		hx_dynamic_value_66 := true
		Reflect_setField(hx_dynamic_target_64, hxrt.StringFromLiteral("added"), hx_dynamic_value_66)
		return hx_dynamic_value_66
	}()
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("added "), Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(objs, 0)), hxrt.StringFromLiteral("added"))), hxrt.StringFromLiteral(" ")), Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(objs, 0)), hxrt.StringFromLiteral("n"))))
	hxrt.Println(v_2)
}

func nulls() {
	var mixed any = dyn(hxrt.NewArray(nil, hxrt.StringFromLiteral("x")))
	fails(hxrt.StringFromLiteral("null field"), hxrt.NewFunction(func() any {
		return Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(mixed, 0)), hxrt.StringFromLiteral("length"))
	}))
	fails(hxrt.StringFromLiteral("null call"), hxrt.NewFunction(func() any {
		return reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(mixed, 0)), hxrt.StringFromLiteral("charAt")), []any{0})
	}))
	fails(hxrt.StringFromLiteral("null write"), hx_recover_callable_c5309b43(hxrt.NewFunction(func() int {
		return hxrt.IntFromNullableAny(func() any {
			hx_dynamic_target_67 := hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(mixed, 0))
			hx_dynamic_value_69 := 1
			Reflect_setField(hx_dynamic_target_67, hxrt.StringFromLiteral("length"), hx_dynamic_value_69)
			return hx_dynamic_value_69
		}())
	})))
	fails(hxrt.StringFromLiteral("null step"), hx_recover_callable_c5309b43(hxrt.NewFunction(func() int {
		fh := hx_structural_recover(hxrt.DynamicIndexGet(mixed, 0))
		return func() int {
			hx_structural_target_70 := fh
			hx_structural_old_71 := func(hx_value_73 any) int {
				if hx_value_73 == nil {
					var hx_zero_74 int
					return hx_zero_74
				}
				return hx_value_73.(int)
			}(hxrt.AnonymousGet(hx_structural_target_70, "n"))
			hx_structural_value_72 := int(int32((hx_structural_old_71 + 1)))
			hxrt.AnonymousSet(hx_structural_target_70, "n", hx_structural_value_72)
			return hx_structural_old_71
		}()
	})))
	fails(hxrt.StringFromLiteral("outside field"), hxrt.NewFunction(func() any {
		return Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(mixed, 5)), hxrt.StringFromLiteral("length"))
	}))
	fails(hxrt.StringFromLiteral("present field"), hxrt.NewFunction(func() any {
		return Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(mixed, 1)), hxrt.StringFromLiteral("length"))
	}))
}

var order *string = hxrt.StringFromLiteral("")

func parsed() {
	var json any = hxrt.JsonParse(hxrt.StringFromLiteral("{\"items\":[{\"name\":\"ann\",\"tags\":[\"a\",\"b\"]},{\"name\":\"bo\",\"tags\":[]}]}"))
	var v any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("json "), Reflect_field(hxrt.DynamicFieldReceiver(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(Reflect_field(json, hxrt.StringFromLiteral("items")), 0)), hxrt.StringFromLiteral("name"))), hxrt.StringFromLiteral("length"))), hxrt.StringFromLiteral(" ")), Reflect_field(hxrt.DynamicFieldReceiver(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(Reflect_field(json, hxrt.StringFromLiteral("items")), 0)), hxrt.StringFromLiteral("tags"))), hxrt.StringFromLiteral("length"))), hxrt.StringFromLiteral(" ")), Reflect_field(hxrt.DynamicFieldReceiver(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(Reflect_field(json, hxrt.StringFromLiteral("items")), 1)), hxrt.StringFromLiteral("tags"))), hxrt.StringFromLiteral("length"))))
	hxrt.Println(v)
	reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(hxrt.DynamicFieldReceiver(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(Reflect_field(json, hxrt.StringFromLiteral("items")), 0)), hxrt.StringFromLiteral("tags"))), hxrt.StringFromLiteral("push")), []any{hxrt.StringFromLiteral("c")})
	func() any {
		hx_dynamic_target_75 := hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(Reflect_field(json, hxrt.StringFromLiteral("items")), 1))
		hx_dynamic_value_77 := hxrt.StringFromLiteral("bob")
		Reflect_setField(hx_dynamic_target_75, hxrt.StringFromLiteral("name"), hx_dynamic_value_77)
		return hx_dynamic_value_77
	}()
	var v_1 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("json "), reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(hxrt.DynamicFieldReceiver(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(Reflect_field(json, hxrt.StringFromLiteral("items")), 0)), hxrt.StringFromLiteral("tags"))), hxrt.StringFromLiteral("join")), []any{hxrt.StringFromLiteral(",")})), hxrt.StringFromLiteral(" ")), Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(Reflect_field(json, hxrt.StringFromLiteral("items")), 1)), hxrt.StringFromLiteral("name"))), hxrt.StringFromLiteral(" ")), reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(hxrt.DynamicFieldReceiver(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(Reflect_field(json, hxrt.StringFromLiteral("items")), 0)), hxrt.StringFromLiteral("name"))), hxrt.StringFromLiteral("toUpperCase")), []any{})))
	hxrt.Println(v_1)
}

func step(name *string, value any) any {
	order = hxrt.StringConcatStringPtr(order, name)
	return value
}

func structures() {
	var objs any = dyn(hxrt.NewArray(func() hxrt.Anonymous {
		hx_obj_78 := map[string]any{}
		hx_obj_78["name"] = hxrt.StringFromLiteral("ann")
		hx_obj_78["n"] = 1
		hx_obj_78["f"] = hxrt.NewFunction(func(v int) int {
			return int((hxrt.Int32Wrap(v) + hxrt.Int32Wrap(1)))
		})
		hx_obj_79 := map[string]any{}
		hx_obj_79["label"] = hxrt.StringFromLiteral("deep")
		hx_obj_79["tags"] = hxrt.NewArray(hxrt.StringFromLiteral("a"), hxrt.StringFromLiteral("b"))
		hx_obj_78["inner"] = hxrt.AnonymousFromMap(hx_obj_79)
		return hxrt.AnonymousFromMap(hx_obj_78)
	}(), New_Point(4)))
	var v any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("anon "), Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(objs, 0)), hxrt.StringFromLiteral("name"))), hxrt.StringFromLiteral(" ")), Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(objs, 0)), hxrt.StringFromLiteral("n"))), hxrt.StringFromLiteral(" ")), reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(objs, 0)), hxrt.StringFromLiteral("f")), []any{3})), hxrt.StringFromLiteral(" ")), Reflect_field(hxrt.DynamicFieldReceiver(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(objs, 0)), hxrt.StringFromLiteral("inner"))), hxrt.StringFromLiteral("label"))), hxrt.StringFromLiteral(" ")), Reflect_field(hxrt.DynamicFieldReceiver(Reflect_field(hxrt.DynamicFieldReceiver(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(objs, 0)), hxrt.StringFromLiteral("inner"))), hxrt.StringFromLiteral("tags"))), hxrt.StringFromLiteral("length"))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("class "), Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(objs, 1)), hxrt.StringFromLiteral("x"))), hxrt.StringFromLiteral(" ")), reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(objs, 1)), hxrt.StringFromLiteral("twice")), []any{})))
	hxrt.Println(v_1)
	name := hxrt.StdString(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(objs, 0)), hxrt.StringFromLiteral("name")))
	n := hxrt.IntFromNullableAny(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(objs, 0)), hxrt.StringFromLiteral("n")))
	x := hxrt.DynamicToFloat(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(objs, 1)), hxrt.StringFromLiteral("x")))
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("typed "), hxrt.StringToUpperCaseStringPtr(name)), hxrt.StringFromLiteral(" ")), int((hxrt.Int32Wrap(n)+hxrt.Int32Wrap(1)))), hxrt.StringFromLiteral(" ")), (x / float64(8))))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("missing "), Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(objs, 0)), hxrt.StringFromLiteral("nope"))), hxrt.StringFromLiteral(" ")), Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(objs, 1)), hxrt.StringFromLiteral("nope"))))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("label length "), Reflect_field(hxrt.DynamicFieldReceiver(Reflect_field(hxrt.DynamicFieldReceiver(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(objs, 0)), hxrt.StringFromLiteral("inner"))), hxrt.StringFromLiteral("label"))), hxrt.StringFromLiteral("length"))), hxrt.StringFromLiteral(" ")), reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(hxrt.DynamicFieldReceiver(Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(objs, 0)), hxrt.StringFromLiteral("name"))), hxrt.StringFromLiteral("charAt")), []any{0})))
	hxrt.Println(v_4)
}

type I_Point interface {
	twice() int
}

type Point struct {
	__hx_this I_Point
	x         int
}

func (self *Point) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Point(x int) *Point {
	return hx_init_Point(hx_alloc_Point(nil), x)
}

func hx_alloc_Point(owner any) *Point {
	self := &Point{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Point)
	}
	return self
}

func hx_init_Point(self *Point, x int) *Point {
	self.x = x
	return self
}

func (self *Point) twice() int {
	return int((hxrt.Int32Wrap(self.x) * hxrt.Int32Wrap(2)))
}

func (self *Point) HxClassName() string {
	return "Point"
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *Point:
		return func(hx_structural_source_1 *Point) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
				switch name {
				case "x":
					return hxrt.AnonymousField{Value: hx_structural_source_1.x, Found: true}
				case "twice":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_4 := hx_structural_source_1
						return hxrt.BoundFunction(func() int {
							return hx_receiver_4.__hx_this.twice()
						}, func() any {
							return hx_receiver_4.__hx_this
						}, "twice")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "x":
					hx_structural_source_1.x = func(hx_value_2 any) int {
						if hx_value_2 == nil {
							var hx_zero_3 int
							return hx_zero_3
						}
						return hx_value_2.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_5 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_5, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_5.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_5.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_5.high = func(hx_value_6 any) int {
						if hx_value_6 == nil {
							var hx_zero_7 int
							return hx_zero_7
						}
						return hx_value_6.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_5.low = func(hx_value_8 any) int {
						if hx_value_8 == nil {
							var hx_zero_9 int
							return hx_zero_9
						}
						return hx_value_8.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__Exception:
		return func(hx_structural_source_10 *haxe__Exception) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_10, func(name string) hxrt.AnonymousField {
				switch name {
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_10.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_10.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_10.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_10.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_13 := hx_structural_source_10
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_13.__hx_this.get_message()
						}, func() any {
							return hx_receiver_13.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_14 := hx_structural_source_10
						return hxrt.BoundFunction(func() any {
							return hx_receiver_14.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_14.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_15 := hx_structural_source_10
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_15.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_15.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_16 := hx_structural_source_10
						return hxrt.BoundFunction(func() any {
							return hx_receiver_16.__hx_this.get_native()
						}, func() any {
							return hx_receiver_16.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_17 := hx_structural_source_10
						return hxrt.BoundFunction(func() any {
							return hx_receiver_17.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_17.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_18 := hx_structural_source_10
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_18.__hx_this.toString()
						}, func() any {
							return hx_receiver_18.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_19 := hx_structural_source_10
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_19.__hx_this.details()
						}, func() any {
							return hx_receiver_19.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_10.savedMessage = func(hx_value_11 any) *string {
						if hx_value_11 == nil {
							var hx_zero_12 *string
							return hx_zero_12
						}
						return hx_value_11.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_10.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_10.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_10.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__ValueException:
		return func(hx_structural_source_20 *haxe__ValueException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_20, func(name string) hxrt.AnonymousField {
				switch name {
				case "value":
					return hxrt.AnonymousField{Value: hx_structural_source_20.value, Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_21 := hx_structural_source_20
						return hxrt.BoundFunction(func() any {
							return hx_receiver_21.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_21.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_20.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_20.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_20.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_20.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_24 := hx_structural_source_20
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_24.__hx_this.get_message()
						}, func() any {
							return hx_receiver_24.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_25 := hx_structural_source_20
						return hxrt.BoundFunction(func() any {
							return hx_receiver_25.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_25.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_26 := hx_structural_source_20
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_26.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_26.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_27 := hx_structural_source_20
						return hxrt.BoundFunction(func() any {
							return hx_receiver_27.__hx_this.get_native()
						}, func() any {
							return hx_receiver_27.__hx_this
						}, "get_native")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_28 := hx_structural_source_20
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_28.__hx_this.toString()
						}, func() any {
							return hx_receiver_28.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_29 := hx_structural_source_20
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_29.__hx_this.details()
						}, func() any {
							return hx_receiver_29.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_20.savedMessage = func(hx_value_22 any) *string {
						if hx_value_22 == nil {
							var hx_zero_23 *string
							return hx_zero_23
						}
						return hx_value_22.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_20.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_20.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_20.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	default:
		return hxrt.AnonymousFromValue(value)
	}
}

func hxrt__generated_method_field(obj any, key string) any {
	var receiver any
	switch value := hxrt.ObjectSource(obj).(type) {
	case *Point:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *Point:
		return hxrt__generated_method_field__Point(value, key)
	case *haxe__Exception:
		return hxrt__generated_method_field__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_method_field__haxe__ValueException(value, key)
	default:
		return nil
	}
}

func hxrt__generated_method_field__Point(value *Point, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "twice":
		return func() *hxrt.Function[func() int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_1.__hx_this.twice()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "twice")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__Exception(value *haxe__Exception, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "details":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.details()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "details")
		}()
	case "get_message":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.get_message()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get_message")
		}()
	case "get_native":
		return func() *hxrt.Function[func() any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_1.__hx_this.get_native()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get_native")
		}()
	case "get_previous":
		return func() *hxrt.Function[func() *haxe__Exception] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *haxe__Exception {
				return hx_receiver_1.__hx_this.get_previous()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get_previous")
		}()
	case "get_stack":
		return func() *hxrt.Function[func() any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_1.__hx_this.get_stack()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get_stack")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.toString()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "toString")
		}()
	case "unwrap":
		return func() *hxrt.Function[func() any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_1.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "unwrap")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__ValueException(value *haxe__ValueException, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "unwrap":
		return func() *hxrt.Function[func() any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_1.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "unwrap")
		}()
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_method_field__haxe__Exception(value.haxe__Exception, key)
}

func hxrt_typeClassMetadataField(value any, key string) (any, bool) {
	classValue, ok := value.(*hxrt__TypeClassValue)
	if !ok || classValue == nil {
		return nil, false
	}
	className := *hxrt.StdString(classValue.name)
	switch className {
	default:
		return nil, false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedField(object any, field *string) any {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *Point:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *Point:
		return hxrt__generated_field_lookup__Point(value, key)
	case *haxe__Exception:
		return hxrt__generated_field_lookup__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_lookup__haxe__ValueException(value, key)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_lookup__haxe___Int64_____Int64(value, key)
	default:
		return nil
	}
}

func reflaxe__go___internal__CompilerReflect_hasGeneratedField(object any, field *string) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *Point:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *Point:
		return hxrt__generated_field_has__Point(value, key)
	case *haxe__Exception:
		return hxrt__generated_field_has__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_has__haxe__ValueException(value, key)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_has__haxe___Int64_____Int64(value, key)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_setGeneratedField(object any, field *string, incoming any) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *Point:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *Point:
		return hxrt__generated_field_set__Point(value, key, incoming)
	case *haxe__Exception:
		return hxrt__generated_field_set__haxe__Exception(value, key, incoming)
	case *haxe__ValueException:
		return hxrt__generated_field_set__haxe__ValueException(value, key, incoming)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_set__haxe___Int64_____Int64(value, key, incoming)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedFields(object any) *hxrt.Array {
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *Point:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch receiver.(type) {
	case *Point:
		return hxrt.NewArray(hxrt.StringFromLiteral("x"))
	case *haxe__Exception:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"))
	case *haxe__ValueException:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("value"))
	case *haxe___Int64_____Int64:
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	default:
		return nil
	}
}

func hxrt__generated_field_lookup__Point(value *Point, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "x":
		return hxrt.ReadStoredField(value.x)
	}
	return nil
}

func hxrt__generated_field_has__Point(value *Point, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "x":
		return true
	}
	return false
}

func hxrt__generated_field_set__Point(value *Point, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "x":
		if incoming == nil {
			var zero int
			value.x = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.x = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__Exception(value *haxe__Exception, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "savedMessage":
		return hxrt.ReadStoredField(value.savedMessage)
	case "savedNative":
		return hxrt.ReadStoredField(value.savedNative)
	case "savedPrevious":
		return hxrt.ReadStoredField(value.savedPrevious)
	case "savedStack":
		return hxrt.ReadStoredField(value.savedStack)
	}
	return nil
}

func hxrt__generated_field_has__haxe__Exception(value *haxe__Exception, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "savedMessage":
		return true
	case "savedNative":
		return true
	case "savedPrevious":
		return true
	case "savedStack":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__Exception(value *haxe__Exception, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "savedMessage":
		if incoming == nil {
			var zero *string
			value.savedMessage = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.savedMessage = typed
			return true
		default:
			return false
		}
	case "savedNative":
		if incoming == nil {
			var zero any
			value.savedNative = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.savedNative = typed
			return true
		default:
			return false
		}
	case "savedPrevious":
		if incoming == nil {
			var zero *haxe__Exception
			value.savedPrevious = zero
			return true
		}
		switch typed := incoming.(type) {
		case *haxe__Exception:
			value.savedPrevious = typed
			return true
		default:
			return false
		}
	case "savedStack":
		if incoming == nil {
			var zero any
			value.savedStack = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.savedStack = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__ValueException(value *haxe__ValueException, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "value":
		return hxrt.ReadStoredField(value.value)
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_field_lookup__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_has__haxe__ValueException(value *haxe__ValueException, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "value":
		return true
	}
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_has__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_set__haxe__ValueException(value *haxe__ValueException, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "value":
		if incoming == nil {
			var zero any
			value.value = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.value = typed
			return true
		default:
			return false
		}
	}
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_set__haxe__Exception(value.haxe__Exception, key, incoming)
}

func hxrt__generated_field_lookup__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "high":
		return hxrt.ReadStoredField(value.high)
	case "low":
		return hxrt.ReadStoredField(value.low)
	}
	return nil
}

func hxrt__generated_field_has__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "high":
		return true
	case "low":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "high":
		if incoming == nil {
			var zero int
			value.high = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.high = typed
			return true
		default:
			return false
		}
	case "low":
		if incoming == nil {
			var zero int
			value.low = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.low = typed
			return true
		default:
			return false
		}
	}
	return false
}

func reflaxe__go___internal__CompilerReflect_typeField(object any, field *string) any {
	key := *hxrt.StdString(field)
	value, found := hxrt_typeClassMetadataField(object, key)
	if !found {
		return nil
	}
	return value
}

func reflaxe__go___internal__CompilerReflect_hasTypeField(object any, field *string) bool {
	key := *hxrt.StdString(field)
	_, found := hxrt_typeClassMetadataField(object, key)
	return found
}

func reflaxe__go___internal__CompilerReflect_generatedMethod(object any, field *string) any {
	key := *hxrt.StdString(field)
	return hxrt__generated_method_field(object, key)
}

func reflaxe__go___internal__CompilerReflect_isEnumValue(value any) bool {
	switch enumValue := value.(type) {
	case *haxe__StackItem:
		return (enumValue != nil)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_callMethod(value any, arguments []any) any {
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Cannot call null"))
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() any]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() int]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(int) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_1 any) int {
			if hx_value_1 == nil {
				var hx_zero_2 int
				return hx_zero_2
			}
			return hx_value_1.(int)
		}(argument_0))
	case *hxrt.Function[func(...any) any]:
		var argument_0 []any = nil
		for index, argument := range arguments {
			if index >= 0 {
				argument_0 = append(argument_0, argument)
			}
		}
		return hxrt.FunctionCall(source)(argument_0...)
	case *hxrt.Function[func() *string]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.FunctionCall(source)()
	default:
		return hxrt.ReflectCallMethod(value, arguments)
	}
}

func hx_recover_callable_c5309b43(value any) *hxrt.Function[func() any] {
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
	case *hxrt.Function[func() any]:
		return source
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_class_haxe__Exception(value any) *haxe__Exception {
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

func hx_recover_class_haxe__ValueException(value any) *haxe__ValueException {
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

func hx_is_nominal_haxe__Exception(value any) bool {
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

func hx_is_nominal_haxe__ValueException(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}
