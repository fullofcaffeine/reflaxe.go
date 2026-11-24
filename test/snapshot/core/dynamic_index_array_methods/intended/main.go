package main

import (
	"math"
	"reflect"
	"snapshot/hxrt"
)

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
		if func(hx_value any) bool {
			switch hx_value.(type) {
			case *string:
				return true
			default:
				return false
			}
		}(hx_payload_5) {
			var e *string = func(hx_value_6 any) *string {
				if hx_value_6 == nil {
					var hx_zero_7 *string
					return hx_zero_7
				}
				return hx_value_6.(*string)
			}(hx_payload_5)
			hx_try_1 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("error "), e)
		} else {
			hxrt.Throw(hx_caught_3)
		}
	})
	text := hx_try_1
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(name, hxrt.StringFromLiteral(" = ")), text)))
}

func boundTypes() {
	var a any = dyn(hxrt.NewArray(10, 20, 30))
	var words any = dyn(hxrt.NewArray(hxrt.StringFromLiteral("ab"), hxrt.StringFromLiteral("cd")))
	typedInt := hxrt.IntFromNullableAny(hxrt.DynamicIndexGet(a, 0))
	typedFloat := hxrt.DynamicToFloat(hxrt.DynamicIndexGet(a, 1))
	typedString := hxrt.StdString(hxrt.DynamicIndexGet(words, 1))
	var missing any = hxrt.DynamicIndexGet(a, 5)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("typed "), int((hxrt.Int32Wrap(typedInt)+hxrt.Int32Wrap(1)))), hxrt.StringFromLiteral(" ")), (typedFloat/float64(8))), hxrt.StringFromLiteral(" ")), hxrt.StringToUpperCaseStringPtr(typedString)), hxrt.StringFromLiteral(" ")), hxrt.StdString((missing == nil))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("returns "), int((hxrt.Int32Wrap(firstInt(a))+hxrt.Int32Wrap(1)))), hxrt.StringFromLiteral(" ")), firstText(words)))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("operators "), hxrt.DynamicToFloat(hxrt.DynamicAdd(hxrt.DynamicIndexGet(a, 0), hxrt.DynamicIndexGet(a, 1)))), hxrt.StringFromLiteral(" ")), hxrt.DynamicToFloat(hxrt.DynamicSub(hxrt.DynamicIndexGet(a, 2), 5))), hxrt.StringFromLiteral(" ")), hxrt.DynamicToFloat(hxrt.DynamicMul(hxrt.DynamicIndexGet(a, 0), 2))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.DynamicGreater(hxrt.DynamicIndexGet(a, 1), hxrt.DynamicIndexGet(a, 0)))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.HaxeEqual(hxrt.DynamicIndexGet(a, 0), 10))), hxrt.StringFromLiteral(" ")), hxrt.DynamicToFloat(hxrt.DynamicNeg(hxrt.DynamicIndexGet(a, 0)))))
	hxrt.Println(v_2)
	var flags any = dyn(hxrt.NewArray(true, false))
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("conditions "), func() *string {
		var hx_if_8 *string
		if hxrt.DynamicIsTrue(hxrt.DynamicIndexGet(flags, 0)) {
			hx_if_8 = hxrt.StringFromLiteral("yes")
		} else {
			hx_if_8 = hxrt.StringFromLiteral("no")
		}
		return hx_if_8
	}()), hxrt.StringFromLiteral(" ")), func() *string {
		var hx_if_9 *string
		if hxrt.DynamicIsTrue(hxrt.DynamicIndexGet(flags, 1)) {
			hx_if_9 = hxrt.StringFromLiteral("yes")
		} else {
			hx_if_9 = hxrt.StringFromLiteral("no")
		}
		return hx_if_9
	}()), hxrt.StringFromLiteral(" ")), hxrt.StdString((hxrt.DynamicIsTrue(hxrt.DynamicIndexGet(flags, 0)) && hxrt.DynamicNot(hxrt.DynamicIndexGet(flags, 1))))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.AnyEqualsNull(hxrt.DynamicIndexGet(a, 5)))))
	hxrt.Println(v_3)
	var big any = dyn(hxrt.NewArray(2147483647))
	var v_4 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("int "), hxrt.DynamicToFloat(hxrt.DynamicAdd(hxrt.DynamicIndexGet(big, 0), 1))), hxrt.StringFromLiteral(" ")), hxrt.MathTruncInt(hxrt.DynamicDiv(hxrt.DynamicIndexGet(a, 0), 3))), hxrt.StringFromLiteral(" ")), math.Abs(hxrt.DynamicToFloat(hxrt.DynamicIndexGet(a, 0)))))
	hxrt.Println(v_4)
	_g := hxrt.IntFromNullableAny(hxrt.DynamicIndexGet(a, 1))
	if _g == 20 {
		hxrt.Println(any(hxrt.StringFromLiteral("switch twenty")))
	} else {
		hxrt.Println(any(hxrt.StringFromLiteral("switch other")))
	}
	var mixed any = dyn(hxrt.NewArray(1, hxrt.StringFromLiteral("x"), 2.5, nil))
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("mixed "), hxrt.DynamicToFloat(hxrt.DynamicAdd(hxrt.DynamicIndexGet(mixed, 0), hxrt.DynamicIndexGet(mixed, 2)))), hxrt.StringFromLiteral(" ")), hxrt.DynamicToFloat(hxrt.DynamicMul(hxrt.DynamicIndexGet(mixed, 2), 2))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.DynamicLess(hxrt.DynamicIndexGet(mixed, 0), hxrt.DynamicIndexGet(mixed, 2)))))
	hxrt.Println(v_5)
	attempt(hxrt.StringFromLiteral("null operand"), hx_recover_callable_c5309b43(hxrt.NewFunction(func() float64 {
		return hxrt.DynamicToFloat(hxrt.DynamicMul(hxrt.DynamicIndexGet(mixed, 3), 2))
	})))
	var total any = 0
	_g_1 := 0
	_g1 := hxrt.IntFromNullableAny(Reflect_field(a, hxrt.StringFromLiteral("length")))
	for _g_1 < _g1 {
		hx_post_10 := _g_1
		_g_1 = int(int32((_g_1 + 1)))
		k := hx_post_10
		total = hxrt.DynamicAdd(total, hxrt.DynamicIndexGet(a, k))
	}
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("loop "), hxrt.StdString(total)))
	hxrt.Println(v_6)
}

func byValue(a any, b any) int {
	var hx_if_12 int
	if hxrt.DynamicToFloat(hxrt.DynamicSub(a, b)) > 0 {
		hx_if_12 = 1
	} else {
		var hx_if_11 int
		if hxrt.HaxeEqual(a, b) {
			hx_if_11 = 0
		} else {
			hx_if_11 = -1
		}
		hx_if_12 = hx_if_11
	}
	return hx_if_12
}

func compound() {
	var a any = dyn(hxrt.NewArray(10, 2.5, hxrt.StringFromLiteral("s"), nil))
	func() any {
		hx_dynamic_target_13 := a
		hx_dynamic_index_14 := 0
		hx_dynamic_old_15 := hxrt.DynamicIndexGet(hx_dynamic_target_13, hx_dynamic_index_14)
		hx_dynamic_value_16 := hxrt.DynamicAdd(hx_dynamic_old_15, 1)
		hxrt.DynamicIndexSet(hx_dynamic_target_13, hx_dynamic_index_14, hx_dynamic_value_16)
		return hx_dynamic_value_16
	}()
	func() any {
		hx_dynamic_target_17 := a
		hx_dynamic_index_18 := 1
		hx_dynamic_old_19 := hxrt.DynamicIndexGet(hx_dynamic_target_17, hx_dynamic_index_18)
		hx_dynamic_value_20 := hxrt.DynamicMul(hx_dynamic_old_19, 2)
		hxrt.DynamicIndexSet(hx_dynamic_target_17, hx_dynamic_index_18, hx_dynamic_value_20)
		return hx_dynamic_value_20
	}()
	func() any {
		hx_dynamic_target_21 := a
		hx_dynamic_index_22 := 2
		hx_dynamic_old_23 := hxrt.DynamicIndexGet(hx_dynamic_target_21, hx_dynamic_index_22)
		hx_dynamic_value_24 := hxrt.StringConcatAny(hx_dynamic_old_23, hxrt.StringFromLiteral("t"))
		hxrt.DynamicIndexSet(hx_dynamic_target_21, hx_dynamic_index_22, hx_dynamic_value_24)
		return hx_dynamic_value_24
	}()
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("compound "), hxrt.StdString(a)))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("value "), hxrt.DynamicToFloat(func() any {
		hx_dynamic_target_25 := a
		hx_dynamic_index_26 := 0
		hx_dynamic_old_27 := hxrt.DynamicIndexGet(hx_dynamic_target_25, hx_dynamic_index_26)
		hx_dynamic_value_28 := hxrt.DynamicSub(hx_dynamic_old_27, 4)
		hxrt.DynamicIndexSet(hx_dynamic_target_25, hx_dynamic_index_26, hx_dynamic_value_28)
		return hx_dynamic_value_28
	}())), hxrt.StringFromLiteral(" ")), hxrt.DynamicIndexGet(a, 0)))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("postfix "), hxrt.DynamicToFloat(func() any {
		hx_dynamic_target_29 := a
		hx_dynamic_index_30 := 0
		hx_dynamic_old_31 := hxrt.DynamicIndexGet(hx_dynamic_target_29, hx_dynamic_index_30)
		hx_dynamic_value_32 := hxrt.DynamicAdd(hx_dynamic_old_31, 1)
		hxrt.DynamicIndexSet(hx_dynamic_target_29, hx_dynamic_index_30, hx_dynamic_value_32)
		return hx_dynamic_old_31
	}())), hxrt.StringFromLiteral(" ")), hxrt.DynamicIndexGet(a, 0)))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("prefix "), hxrt.DynamicToFloat(func() any {
		hx_dynamic_target_33 := a
		hx_dynamic_index_34 := 0
		hx_dynamic_old_35 := hxrt.DynamicIndexGet(hx_dynamic_target_33, hx_dynamic_index_34)
		hx_dynamic_value_36 := hxrt.DynamicAdd(hx_dynamic_old_35, 1)
		hxrt.DynamicIndexSet(hx_dynamic_target_33, hx_dynamic_index_34, hx_dynamic_value_36)
		return hx_dynamic_value_36
	}())), hxrt.StringFromLiteral(" ")), hxrt.DynamicIndexGet(a, 0)))
	hxrt.Println(v_3)
	func() any {
		hx_dynamic_target_37 := a
		hx_dynamic_index_38 := 0
		hx_dynamic_old_39 := hxrt.DynamicIndexGet(hx_dynamic_target_37, hx_dynamic_index_38)
		hx_dynamic_value_40 := hxrt.DynamicSub(hx_dynamic_old_39, 1)
		hxrt.DynamicIndexSet(hx_dynamic_target_37, hx_dynamic_index_38, hx_dynamic_value_40)
		return hx_dynamic_old_39
	}()
	func() any {
		hx_dynamic_target_41 := a
		hx_dynamic_index_42 := 0
		hx_dynamic_old_43 := hxrt.DynamicIndexGet(hx_dynamic_target_41, hx_dynamic_index_42)
		hx_dynamic_value_44 := hxrt.DynamicSub(hx_dynamic_old_43, 1)
		hxrt.DynamicIndexSet(hx_dynamic_target_41, hx_dynamic_index_42, hx_dynamic_value_44)
		return hx_dynamic_value_44
	}()
	func() any {
		hx_dynamic_target_45 := a
		hx_dynamic_index_46 := 0
		hx_dynamic_old_47 := hxrt.DynamicIndexGet(hx_dynamic_target_45, hx_dynamic_index_46)
		hx_dynamic_value_48 := hxrt.DynamicAdd(hx_dynamic_old_47, 1)
		hxrt.DynamicIndexSet(hx_dynamic_target_45, hx_dynamic_index_46, hx_dynamic_value_48)
		return hx_dynamic_old_47
	}()
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("steps "), hxrt.DynamicIndexGet(a, 0))))
	func() any {
		hx_dynamic_target_49 := a
		hx_dynamic_index_50 := 0
		hx_dynamic_old_51 := hxrt.DynamicIndexGet(hx_dynamic_target_49, hx_dynamic_index_50)
		hx_dynamic_value_52 := hxrt.DynamicAdd(hx_dynamic_old_51, hxrt.DynamicToFloat(hxrt.DynamicIndexGet(a, 0)))
		hxrt.DynamicIndexSet(hx_dynamic_target_49, hx_dynamic_index_50, hx_dynamic_value_52)
		return hx_dynamic_value_52
	}()
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("self "), hxrt.DynamicIndexGet(a, 0))))
	attempt(hxrt.StringFromLiteral("null element"), hx_recover_callable_c5309b43(hxrt.NewFunction(func() float64 {
		return hxrt.DynamicToFloat(func() any {
			hx_dynamic_target_53 := a
			hx_dynamic_index_54 := 3
			hx_dynamic_old_55 := hxrt.DynamicIndexGet(hx_dynamic_target_53, hx_dynamic_index_54)
			hx_dynamic_value_56 := hxrt.DynamicAdd(hx_dynamic_old_55, 1)
			hxrt.DynamicIndexSet(hx_dynamic_target_53, hx_dynamic_index_54, hx_dynamic_value_56)
			return hx_dynamic_value_56
		}())
	})))
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("after "), hxrt.StdString(a)))
	hxrt.Println(v_4)
	var holder any = dyn(func() hxrt.Anonymous {
		hx_obj_57 := map[string]any{}
		hx_obj_57["list"] = hxrt.NewArray(1, 2)
		return hxrt.AnonymousFromMap(hx_obj_57)
	}())
	var base any = Reflect_field(holder, hxrt.StringFromLiteral("list"))
	func() any {
		hx_dynamic_target_58 := base
		hx_dynamic_index_59 := 0
		hx_dynamic_old_60 := hxrt.DynamicIndexGet(hx_dynamic_target_58, hx_dynamic_index_59)
		hx_dynamic_value_61 := hxrt.DynamicAdd(hx_dynamic_old_60, 5)
		hxrt.DynamicIndexSet(hx_dynamic_target_58, hx_dynamic_index_59, hx_dynamic_value_61)
		return hx_dynamic_value_61
	}()
	func() any {
		hx_dynamic_target_62 := Reflect_field(holder, hxrt.StringFromLiteral("list"))
		hx_dynamic_index_63 := 1
		hx_dynamic_old_64 := hxrt.DynamicIndexGet(hx_dynamic_target_62, hx_dynamic_index_63)
		hx_dynamic_value_65 := hxrt.DynamicAdd(hx_dynamic_old_64, 1)
		hxrt.DynamicIndexSet(hx_dynamic_target_62, hx_dynamic_index_63, hx_dynamic_value_65)
		return hx_dynamic_old_64
	}()
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("field element "), hxrt.StdString(Reflect_field(holder, hxrt.StringFromLiteral("list")))))
	hxrt.Println(v_5)
}

func dyn(value any) any {
	return value
}

func evaluationOrder() {
	var a any = dyn(hxrt.NewArray(1, 2))
	func() any {
		hx_dynamic_target_66 := step(hxrt.StringFromLiteral("t"), a)
		hx_dynamic_index_67 := hxrt.IntFromNullableAny(step(hxrt.StringFromLiteral("i"), 0))
		hx_dynamic_value_69 := step(hxrt.StringFromLiteral("v"), 10)
		hxrt.DynamicIndexSet(hx_dynamic_target_66, hx_dynamic_index_67, hx_dynamic_value_69)
		return hx_dynamic_value_69
	}()
	var base any = step(hxrt.StringFromLiteral("t"), a)
	var index any = step(hxrt.StringFromLiteral("i"), 1)
	func() any {
		hx_dynamic_target_70 := base
		hx_dynamic_index_71 := hxrt.IntFromNullableAny(index)
		hx_dynamic_old_72 := hxrt.DynamicIndexGet(hx_dynamic_target_70, hx_dynamic_index_71)
		hx_dynamic_value_73 := hxrt.DynamicAdd(hx_dynamic_old_72, step(hxrt.StringFromLiteral("v"), 5))
		hxrt.DynamicIndexSet(hx_dynamic_target_70, hx_dynamic_index_71, hx_dynamic_value_73)
		return hx_dynamic_value_73
	}()
	old := hxrt.DynamicToFloat(func() any {
		hx_dynamic_target_74 := step(hxrt.StringFromLiteral("t"), a)
		hx_dynamic_index_75 := hxrt.IntFromNullableAny(step(hxrt.StringFromLiteral("i"), 0))
		hx_dynamic_old_76 := hxrt.DynamicIndexGet(hx_dynamic_target_74, hx_dynamic_index_75)
		hx_dynamic_value_77 := hxrt.DynamicAdd(hx_dynamic_old_76, 1)
		hxrt.DynamicIndexSet(hx_dynamic_target_74, hx_dynamic_index_75, hx_dynamic_value_77)
		return hx_dynamic_old_76
	}())
	var v any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("order "), order), hxrt.StringFromLiteral(" ")), hxrt.StdString(a)), hxrt.StringFromLiteral(" ")), old))
	hxrt.Println(v)
	order = hxrt.StringFromLiteral("")
	read := hxrt.StdString(hxrt.DynamicIndexGet(step(hxrt.StringFromLiteral("t"), a), hxrt.IntFromNullableAny(step(hxrt.StringFromLiteral("i"), 1))))
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("read order "), order), hxrt.StringFromLiteral(" ")), read))
	hxrt.Println(v_1)
}

func expressionMethods() {
	var a any = dyn(hxrt.NewArray(3, 1, 2))
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("push "), hxrt.StdString(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("push")), []any{4}))), hxrt.StringFromLiteral(" ")), hxrt.StdString(a)))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("pop "), hxrt.StdString(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("pop")), []any{}))), hxrt.StringFromLiteral(" ")), hxrt.StdString(a)))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("unshift "), hxrt.StdString(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("unshift")), []any{0}))), hxrt.StringFromLiteral(" ")), hxrt.StdString(a)))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("shift "), hxrt.StdString(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("shift")), []any{}))), hxrt.StringFromLiteral(" ")), hxrt.StdString(a)))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("insert "), hxrt.StdString(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("insert")), []any{1, 9}))), hxrt.StringFromLiteral(" ")), hxrt.StdString(a)))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("remove "), hxrt.StdString(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("remove")), []any{9}))), hxrt.StringFromLiteral(" ")), hxrt.StdString(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("remove")), []any{9}))), hxrt.StringFromLiteral(" ")), hxrt.StdString(a)))
	hxrt.Println(v_5)
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("reverse "), hxrt.StdString(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("reverse")), []any{}))), hxrt.StringFromLiteral(" ")), hxrt.StdString(a)))
	hxrt.Println(v_6)
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("sort "), hxrt.StdString(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("sort")), []any{hxrt.NamedFunction(byValue, "Main.byValue")}))), hxrt.StringFromLiteral(" ")), hxrt.StdString(a)))
	hxrt.Println(v_7)
	var v_8 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("resize "), hxrt.StdString(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("resize")), []any{4}))), hxrt.StringFromLiteral(" ")), hxrt.StdString(a)))
	hxrt.Println(v_8)
	var v_9 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("splice "), hxrt.StdString(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("splice")), []any{3, 1}))), hxrt.StringFromLiteral(" ")), hxrt.StdString(a)))
	hxrt.Println(v_9)
	var v_10 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("concat "), hxrt.StdString(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("concat")), []any{hxrt.NewArray(5)}))), hxrt.StringFromLiteral(" copy ")), hxrt.StdString(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("copy")), []any{}))))
	hxrt.Println(v_10)
	var v_11 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("filter "), hxrt.StdString(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("filter")), []any{hxrt.NewFunction(func(x int) bool {
		return (x != 1)
	})}))), hxrt.StringFromLiteral(" map ")), hxrt.StdString(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("map")), []any{hxrt.NewFunction(func(x int) int {
		return int((hxrt.Int32Wrap(x) * hxrt.Int32Wrap(2)))
	})}))))
	hxrt.Println(v_11)
	var v_12 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("search "), hxrt.StdString(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("indexOf")), []any{2}))), hxrt.StringFromLiteral(" ")), hxrt.StdString(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("lastIndexOf")), []any{2}))), hxrt.StringFromLiteral(" ")), hxrt.StdString(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("contains")), []any{2}))), hxrt.StringFromLiteral(" ")), hxrt.StdString(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("contains")), []any{7}))))
	hxrt.Println(v_12)
	var v_13 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("join "), hxrt.StdString(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("join")), []any{hxrt.StringFromLiteral("-")}))), hxrt.StringFromLiteral(" slice ")), hxrt.StdString(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("slice")), []any{1}))), hxrt.StringFromLiteral(" text ")), hxrt.StdString(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("toString")), []any{}))))
	hxrt.Println(v_13)
	length := hxrt.IntFromNullableAny(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("push")), []any{8}))
	popped := hxrt.IntFromNullableAny(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("pop")), []any{}))
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("typed "), int((hxrt.Int32Wrap(length) + hxrt.Int32Wrap(popped))))))
}

func firstInt(values any) int {
	return hxrt.IntFromNullableAny(hxrt.DynamicIndexGet(values, 0))
}

func firstText(values any) *string {
	return hxrt.StdString(hxrt.DynamicIndexGet(values, 0))
}

func main() {
	reads()
	boundTypes()
	writes()
	compound()
	nonArrays()
	evaluationOrder()
	statementMethods()
	expressionMethods()
	otherReceivers()
}

func nonArrays() {
	var text any = dyn(hxrt.StringFromLiteral("hello"))
	attempt(hxrt.StringFromLiteral("string read"), hxrt.NewFunction(func() any {
		return hxrt.DynamicIndexGet(text, 0)
	}))
	attempt(hxrt.StringFromLiteral("string write"), hx_recover_callable_c5309b43(hxrt.NewFunction(func() *string {
		return hxrt.StdString(func() any {
			hx_dynamic_target_78 := text
			hx_dynamic_index_79 := 0
			hx_dynamic_value_81 := hxrt.StringFromLiteral("j")
			hxrt.DynamicIndexSet(hx_dynamic_target_78, hx_dynamic_index_79, hx_dynamic_value_81)
			return hx_dynamic_value_81
		}())
	})))
	var object any = dyn(func() hxrt.Anonymous {
		hx_obj_82 := map[string]any{}
		hx_obj_82["x"] = 1
		return hxrt.AnonymousFromMap(hx_obj_82)
	}())
	attempt(hxrt.StringFromLiteral("object read"), hxrt.NewFunction(func() any {
		return hxrt.DynamicIndexGet(object, 0)
	}))
	attempt(hxrt.StringFromLiteral("object write"), hx_recover_callable_c5309b43(hxrt.NewFunction(func() int {
		return hxrt.IntFromNullableAny(func() any {
			hx_dynamic_target_83 := object
			hx_dynamic_index_84 := 0
			hx_dynamic_value_86 := 3
			hxrt.DynamicIndexSet(hx_dynamic_target_83, hx_dynamic_index_84, hx_dynamic_value_86)
			return hx_dynamic_value_86
		}())
	})))
	attempt(hxrt.StringFromLiteral("object compound"), hx_recover_callable_c5309b43(hxrt.NewFunction(func() float64 {
		return hxrt.DynamicToFloat(func() any {
			hx_dynamic_target_87 := object
			hx_dynamic_index_88 := 0
			hx_dynamic_old_89 := hxrt.DynamicIndexGet(hx_dynamic_target_87, hx_dynamic_index_88)
			hx_dynamic_value_90 := hxrt.DynamicAdd(hx_dynamic_old_89, 1)
			hxrt.DynamicIndexSet(hx_dynamic_target_87, hx_dynamic_index_88, hx_dynamic_value_90)
			return hx_dynamic_value_90
		}())
	})))
	attempt(hxrt.StringFromLiteral("object step"), hx_recover_callable_c5309b43(hxrt.NewFunction(func() float64 {
		return hxrt.DynamicToFloat(func() any {
			hx_dynamic_target_91 := object
			hx_dynamic_index_92 := 0
			hx_dynamic_old_93 := hxrt.DynamicIndexGet(hx_dynamic_target_91, hx_dynamic_index_92)
			hx_dynamic_value_94 := hxrt.DynamicAdd(hx_dynamic_old_93, 1)
			hxrt.DynamicIndexSet(hx_dynamic_target_91, hx_dynamic_index_92, hx_dynamic_value_94)
			return hx_dynamic_old_93
		}())
	})))
	var nothing any = dyn(nil)
	attempt(hxrt.StringFromLiteral("null read"), hxrt.NewFunction(func() any {
		return hxrt.DynamicIndexGet(nothing, 0)
	}))
	attempt(hxrt.StringFromLiteral("null write"), hx_recover_callable_c5309b43(hxrt.NewFunction(func() int {
		return hxrt.IntFromNullableAny(func() any {
			hx_dynamic_target_95 := nothing
			hx_dynamic_index_96 := 0
			hx_dynamic_value_98 := 1
			hxrt.DynamicIndexSet(hx_dynamic_target_95, hx_dynamic_index_96, hx_dynamic_value_98)
			return hx_dynamic_value_98
		}())
	})))
	var number any = dyn(5)
	attempt(hxrt.StringFromLiteral("int read"), hxrt.NewFunction(func() any {
		return hxrt.DynamicIndexGet(number, 0)
	}))
	var flag any = dyn(true)
	attempt(hxrt.StringFromLiteral("bool read"), hxrt.NewFunction(func() any {
		return hxrt.DynamicIndexGet(flag, 0)
	}))
	var map_ any = dyn(func() *haxe__ds__IntMap {
		_g := New_haxe__ds__IntMap()
		_g.__hx_this.set(1, hxrt.StringFromLiteral("one"))
		return _g
	}())
	var hx_try_99 *string
	hxrt.TryCatch(func() {
		hx_try_99 = hxrt.StdString(hxrt.DynamicIndexGet(map_, 1))
	}, func(hx_caught_100 any) {
		hx_exception_101 := haxe__Exception_caught(hx_caught_100)
		hx_payload_102 := haxe__Exception_catchValue(hx_caught_100)
		_ = hx_exception_101
		if func(hx_value any) bool {
			switch hx_value.(type) {
			case *string:
				return true
			default:
				return false
			}
		}(hx_payload_102) {
			var e *string = func(hx_value_103 any) *string {
				if hx_value_103 == nil {
					var hx_zero_104 *string
					return hx_zero_104
				}
				return hx_value_103.(*string)
			}(hx_payload_102)
			_ = e
			hx_try_99 = hxrt.StringFromLiteral("error")
		} else {
			hxrt.Throw(hx_caught_100)
		}
	})
	mapText := hx_try_99
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("map read = "), mapText)))
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("unchanged "), hxrt.StdString(text)), hxrt.StringFromLiteral(" ")), hxrt.StdString(Reflect_field(object, hxrt.StringFromLiteral("x")))))
	hxrt.Println(v)
}

var order *string = hxrt.StringFromLiteral("")

func otherReceivers() {
	items := hxrt.NewArray()
	useStack(func() hxrt.Anonymous {
		hx_obj_105 := map[string]any{}
		hx_obj_105["push"] = hxrt.NewFunction(func(value int) int {
			return func() int {
				return items.Push(value)
			}()
		})
		hx_obj_105["pop"] = hxrt.NewFunction(func() any {
			return func() any {
				return items.Pop()
			}()
		})
		return hxrt.AnonymousFromMap(hx_obj_105)
	}())
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("structure "), hxrt.StdString(items)))
	hxrt.Println(v)
	var tally any = dyn(New_Tally())
	reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(tally, hxrt.StringFromLiteral("push")), []any{5})
	reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(tally, hxrt.StringFromLiteral("push")), []any{2})
	reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(tally, hxrt.StringFromLiteral("pop")), []any{})
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("class "), hxrt.StdString(Reflect_field(tally, hxrt.StringFromLiteral("total")))))
	hxrt.Println(v_1)
}

func reads() {
	var a any = dyn(hxrt.NewArray(10, 20, 30))
	var i any = dyn(1)
	var f any = dyn(1.7)
	var n any = 2
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("read "), hxrt.DynamicIndexGet(a, 0)), hxrt.StringFromLiteral(" ")), hxrt.DynamicIndexGet(a, 2)), hxrt.StringFromLiteral(" ")), hxrt.DynamicIndexGet(a, hxrt.IntFromNullableAny(i))), hxrt.StringFromLiteral(" ")), hxrt.DynamicIndexGet(a, hxrt.IntFromNullableAny(f))), hxrt.StringFromLiteral(" ")), hxrt.DynamicIndexGet(a, hxrt.IntFromNullableAny(n)))))
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("outside "), hxrt.DynamicIndexGet(a, 3)), hxrt.StringFromLiteral(" ")), hxrt.DynamicIndexGet(a, 100)), hxrt.StringFromLiteral(" ")), hxrt.DynamicIndexGet(a, -1))))
	var nested any = dyn(hxrt.NewArray(hxrt.NewArray(1, 2), hxrt.NewArray(3, 4)))
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("nested "), func(hx_value_112 any) *hxrt.Array {
		if hx_value_112 == nil {
			var hx_zero_113 *hxrt.Array
			return hx_zero_113
		}
		return hx_value_112.(*hxrt.Array)
	}(hxrt.DynamicIndexGet(nested, 1)).Get(0)), hxrt.StringFromLiteral(" ")), func(hx_value_118 any) *hxrt.Array {
		if hx_value_118 == nil {
			var hx_zero_119 *hxrt.Array
			return hx_zero_119
		}
		return hx_value_118.(*hxrt.Array)
	}(hxrt.DynamicIndexGet(nested, 0)).Get(5))))
	var rows_1 any
	_ = rows_1
	var rows_0 any
	rows_0 = dyn(hxrt.NewArray(5, 6))
	rows_1 = hxrt.StringFromLiteral("x")
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("element "), hxrt.DynamicIndexGet(rows_0, 1))))
	var json any = hxrt.JsonParse(hxrt.StringFromLiteral("{\"items\":[1,2,3],\"grid\":[[7,8]],\"people\":[{\"name\":\"ann\"}]}"))
	var v any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("json "), hxrt.DynamicIndexGet(Reflect_field(json, hxrt.StringFromLiteral("items")), 2)), hxrt.StringFromLiteral(" ")), func(hx_value_124 any) *hxrt.Array {
		if hx_value_124 == nil {
			var hx_zero_125 *hxrt.Array
			return hx_zero_125
		}
		return hx_value_124.(*hxrt.Array)
	}(hxrt.DynamicIndexGet(Reflect_field(json, hxrt.StringFromLiteral("grid")), 0)).Get(1)), hxrt.StringFromLiteral(" ")), Reflect_field(hxrt.DynamicFieldReceiver(hxrt.DynamicIndexGet(Reflect_field(json, hxrt.StringFromLiteral("people")), 0)), hxrt.StringFromLiteral("name"))))
	hxrt.Println(v)
	var keys any = dyn(hxrt.NewArray(0, 2))
	plain := hxrt.NewArray(100, 200, 300)
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("as index "), plain.Get(hxrt.IntFromNullableAny(hxrt.DynamicIndexGet(keys, 1)))), hxrt.StringFromLiteral(" ")), hxrt.DynamicIndexGet(a, hxrt.IntFromNullableAny(hxrt.DynamicIndexGet(keys, 0))))))
	var maybe any = dyn(hxrt.NewArray(4, 5))
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("nullable "), hxrt.DynamicIndexGet(maybe, 1))))
	later := hxrt.NewFunction(func() *string {
		return hxrt.StdString(hxrt.DynamicIndexGet(a, 2))
	})
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("closure "), func() *string {
		var hx_callable_126 *hxrt.Function[func() *string] = later
		return hxrt.FunctionCall(hx_callable_126)()
	}()), hxrt.StringConcatAny(hxrt.StringFromLiteral(" "), hxrt.DynamicIndexGet(a, 1))))
	hxrt.Println(v_1)
}

func statementMethods() {
	var a any = dyn(hxrt.NewArray(3, 1, 2))
	reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("push")), []any{4})
	reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("push")), []any{5})
	reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("pop")), []any{})
	reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("unshift")), []any{0})
	reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("shift")), []any{})
	reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("insert")), []any{1, 9})
	reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("remove")), []any{9})
	reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("reverse")), []any{})
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("statements "), hxrt.StdString(a)))
	hxrt.Println(v)
	reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("sort")), []any{hxrt.NamedFunction(byValue, "Main.byValue")})
	reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("resize")), []any{5})
	reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("splice")), []any{4, 1})
	reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("concat")), []any{hxrt.NewArray(5)})
	reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("copy")), []any{})
	reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("filter")), []any{hxrt.NewFunction(func(x int) bool {
		return (x != 1)
	})})
	reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("map")), []any{hxrt.NewFunction(func(x int) int {
		return int((hxrt.Int32Wrap(x) * hxrt.Int32Wrap(2)))
	})})
	reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("indexOf")), []any{2})
	reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("lastIndexOf")), []any{2})
	reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("contains")), []any{2})
	reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("join")), []any{hxrt.StringFromLiteral(",")})
	reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("slice")), []any{1})
	reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(a, hxrt.StringFromLiteral("toString")), []any{})
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("statements "), hxrt.StdString(a)))
	hxrt.Println(v_1)
}

func step(name *string, value any) any {
	order = hxrt.StringConcatStringPtr(order, name)
	return value
}

func useStack(stack hxrt.Anonymous) {
	func() int {
		var hx_callable_127 *hxrt.Function[func(int) int] = hx_recover_callable_408c8b7a(hxrt.AnonymousGet(stack, "push"))
		var hx_argument_128 int = 1
		return hxrt.FunctionCall(hx_callable_127)(hx_argument_128)
	}()
	func() int {
		var hx_callable_129 *hxrt.Function[func(int) int] = hx_recover_callable_408c8b7a(hxrt.AnonymousGet(stack, "push"))
		var hx_argument_130 int = 2
		return hxrt.FunctionCall(hx_callable_129)(hx_argument_130)
	}()
	func() any {
		var hx_callable_131 *hxrt.Function[func() any] = hx_recover_callable_c5309b43(hxrt.AnonymousGet(stack, "pop"))
		return hxrt.FunctionCall(hx_callable_131)()
	}()
	var v any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("structure push "), func() int {
		var hx_callable_132 *hxrt.Function[func(int) int] = hx_recover_callable_408c8b7a(hxrt.AnonymousGet(stack, "push"))
		var hx_argument_133 int = 3
		return hxrt.FunctionCall(hx_callable_132)(hx_argument_133)
	}()))
	hxrt.Println(v)
}

func writes() {
	var a any = dyn(hxrt.NewArray(10, 20, 30))
	func() any {
		hx_dynamic_target_134 := a
		hx_dynamic_index_135 := 1
		hx_dynamic_value_137 := 5
		hxrt.DynamicIndexSet(hx_dynamic_target_134, hx_dynamic_index_135, hx_dynamic_value_137)
		return hx_dynamic_value_137
	}()
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("write "), hxrt.StdString(a)), hxrt.StringFromLiteral(" ")), hxrt.StdString(Reflect_field(a, hxrt.StringFromLiteral("length")))))
	hxrt.Println(v)
	func() any {
		hx_dynamic_target_138 := a
		hx_dynamic_index_139 := 5
		hx_dynamic_value_141 := 7
		hxrt.DynamicIndexSet(hx_dynamic_target_138, hx_dynamic_index_139, hx_dynamic_value_141)
		return hx_dynamic_value_141
	}()
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("grow "), hxrt.StdString(a)), hxrt.StringFromLiteral(" ")), hxrt.StdString(Reflect_field(a, hxrt.StringFromLiteral("length")))))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("value "), hxrt.IntFromNullableAny(func() any {
		hx_dynamic_target_142 := a
		hx_dynamic_index_143 := 0
		hx_dynamic_value_145 := 1
		hxrt.DynamicIndexSet(hx_dynamic_target_142, hx_dynamic_index_143, hx_dynamic_value_145)
		return hx_dynamic_value_145
	}())), hxrt.StringFromLiteral(" ")), hxrt.StdString(a)))
	hxrt.Println(v_2)
	var i any = dyn(2)
	func() any {
		hx_dynamic_target_146 := a
		hx_dynamic_index_147 := hxrt.IntFromNullableAny(i)
		hx_dynamic_value_149 := hxrt.StringFromLiteral("s")
		hxrt.DynamicIndexSet(hx_dynamic_target_146, hx_dynamic_index_147, hx_dynamic_value_149)
		return hx_dynamic_value_149
	}()
	var f any = dyn(3.9)
	func() any {
		hx_dynamic_target_150 := a
		hx_dynamic_index_151 := hxrt.IntFromNullableAny(f)
		hx_dynamic_value_153 := 3
		hxrt.DynamicIndexSet(hx_dynamic_target_150, hx_dynamic_index_151, hx_dynamic_value_153)
		return hx_dynamic_value_153
	}()
	func() any {
		hx_dynamic_target_154 := a
		hx_dynamic_index_155 := 4
		hx_dynamic_value_157 := 2.0
		hxrt.DynamicIndexSet(hx_dynamic_target_154, hx_dynamic_index_155, hx_dynamic_value_157)
		return hx_dynamic_value_157
	}()
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("dynamic index "), hxrt.StdString(a)), hxrt.StringFromLiteral(" ")), hxrt.StdString(Type_typeof(hxrt.DynamicIndexGet(a, 4)))))
	hxrt.Println(v_3)
	func() any {
		hx_dynamic_target_158 := a
		hx_dynamic_index_159 := 0
		hx_dynamic_value_161 := hxrt.DynamicIndexGet(a, 3)
		hxrt.DynamicIndexSet(hx_dynamic_target_158, hx_dynamic_index_159, hx_dynamic_value_161)
		return hx_dynamic_value_161
	}()
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("copy "), hxrt.StdString(a)))
	hxrt.Println(v_4)
	typed := hxrt.NewArray(1, 2)
	var alias any = typed
	func() any {
		hx_dynamic_target_162 := alias
		hx_dynamic_index_163 := 0
		hx_dynamic_value_165 := 100
		hxrt.DynamicIndexSet(hx_dynamic_target_162, hx_dynamic_index_163, hx_dynamic_value_165)
		return hx_dynamic_value_165
	}()
	func() any {
		hx_dynamic_target_166 := alias
		hx_dynamic_index_167 := 3
		hx_dynamic_value_169 := 300
		hxrt.DynamicIndexSet(hx_dynamic_target_166, hx_dynamic_index_167, hx_dynamic_value_169)
		return hx_dynamic_value_169
	}()
	var v_5 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("alias "), hxrt.StdString(typed)), hxrt.StringFromLiteral(" ")), typed.Len()))
	hxrt.Println(v_5)
	var nested any = dyn(hxrt.NewArray(hxrt.NewArray(1, 2), hxrt.NewArray(3, 4)))
	hx_array_target_172 := func(hx_value_170 any) *hxrt.Array {
		if hx_value_170 == nil {
			var hx_zero_171 *hxrt.Array
			return hx_zero_171
		}
		return hx_value_170.(*hxrt.Array)
	}(hxrt.DynamicIndexGet(nested, 1))
	hx_array_index_173 := 0
	hx_array_target_172.Set(hx_array_index_173, 9)
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("nested "), hxrt.StdString(nested)))
	hxrt.Println(v_6)
	var json any = hxrt.JsonParse(hxrt.StringFromLiteral("{\"items\":[1,2,3]}"))
	func() any {
		hx_dynamic_target_174 := Reflect_field(json, hxrt.StringFromLiteral("items"))
		hx_dynamic_index_175 := 0
		hx_dynamic_value_177 := 7
		hxrt.DynamicIndexSet(hx_dynamic_target_174, hx_dynamic_index_175, hx_dynamic_value_177)
		return hx_dynamic_value_177
	}()
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("json "), func() *string {
		var space *string = nil
		return hxrt.StdString(hxrt.JsonStringify(json, space))
	}()))
	hxrt.Println(v_7)
	rows := hxrt.NewArray(dyn(hxrt.NewArray(5, 6)))
	func() any {
		hx_dynamic_target_178 := rows.Get(0)
		hx_dynamic_index_179 := 1
		hx_dynamic_value_181 := 60
		hxrt.DynamicIndexSet(hx_dynamic_target_178, hx_dynamic_index_179, hx_dynamic_value_181)
		return hx_dynamic_value_181
	}()
	var v_8 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("element "), hxrt.StdString(rows)))
	hxrt.Println(v_8)
	target := hxrt.NewArray(0, 0)
	hx_array_target_182 := target
	hx_array_index_183 := 1
	hx_array_target_182.Set(hx_array_index_183, hxrt.IntFromNullableAny(hxrt.DynamicIndexGet(a, 3)))
	var v_9 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("typed target "), hxrt.StdString(target)))
	hxrt.Println(v_9)
}

type I_Tally interface {
	push(value int) int
	pop() int
}

type Tally struct {
	__hx_this I_Tally
	total     int
}

func (self *Tally) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Tally() *Tally {
	return hx_init_Tally(hx_alloc_Tally(nil))
}

func hx_alloc_Tally(owner any) *Tally {
	self := &Tally{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Tally)
	}
	return self
}

func hx_init_Tally(self *Tally) *Tally {
	self.total = 0
	return self
}

func (self *Tally) push(value int) int {
	self.total = int((hxrt.Int32Wrap(self.total) + hxrt.Int32Wrap(value)))
	return self.total
}

func (self *Tally) pop() int {
	self.total = int(int32((self.total - 1)))
	return self.total
}

func (self *Tally) HxClassName() string {
	return "Tally"
}

func hxrt__generated_method_field(obj any, key string) any {
	var receiver any
	switch value := hxrt.ObjectSource(obj).(type) {
	case *Tally:
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
	case *haxe__ds__IntMap:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *Tally:
		return hxrt__generated_method_field__Tally(value, key)
	case *haxe__Exception:
		return hxrt__generated_method_field__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_method_field__haxe__ValueException(value, key)
	case *haxe__ds__IntMap:
		return hxrt__generated_method_field__haxe__ds__IntMap(value, key)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_method_field__haxe__iterators__MapKeyValueIterator(value, key)
	default:
		return nil
	}
}

func hxrt__generated_method_field__Tally(value *Tally, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "pop":
		return func() *hxrt.Function[func() int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_1.__hx_this.pop()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "pop")
		}()
	case "push":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(value int) int {
				return hx_receiver_1.__hx_this.push(value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "push")
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

func hxrt__generated_method_field__haxe__ds__IntMap(value *haxe__ds__IntMap, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "clear":
		return func() *hxrt.Function[func()] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_1.__hx_this.clear()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "clear")
		}()
	case "copy":
		return func() *hxrt.Function[func() *haxe__ds__IntMap] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *haxe__ds__IntMap {
				return hx_receiver_1.__hx_this.copy()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "copy")
		}()
	case "copyIMap":
		return func() *hxrt.Function[func() haxe__IMap] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() haxe__IMap {
				return hx_receiver_1.__hx_this.copyIMap()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "copyIMap")
		}()
	case "exists":
		return func() *hxrt.Function[func(int) bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(key int) bool {
				return hx_receiver_1.__hx_this.exists(key)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "exists")
		}()
	case "existsIMap":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_1.__hx_this.existsIMap(key)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "existsIMap")
		}()
	case "get":
		return func() *hxrt.Function[func(int) any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(key int) any {
				return hx_receiver_1.__hx_this.get(key)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get")
		}()
	case "getIMap":
		return func() *hxrt.Function[func(any) any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(key any) any {
				return hx_receiver_1.__hx_this.getIMap(key)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "getIMap")
		}()
	case "iterator":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_1.__hx_this.iterator()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "iterator")
		}()
	case "keyValueIterator":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_1.__hx_this.keyValueIterator()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "keyValueIterator")
		}()
	case "keys":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_1.__hx_this.keys()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "keys")
		}()
	case "remove":
		return func() *hxrt.Function[func(int) bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(key int) bool {
				return hx_receiver_1.__hx_this.remove(key)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "remove")
		}()
	case "removeIMap":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_1.__hx_this.removeIMap(key)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "removeIMap")
		}()
	case "set":
		return func() *hxrt.Function[func(int, any)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(key int, value any) {
				hx_receiver_1.__hx_this.set(key, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "set")
		}()
	case "setIMap":
		return func() *hxrt.Function[func(any, any)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(key any, value any) {
				hx_receiver_1.__hx_this.setIMap(key, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "setIMap")
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
	}
	return nil
}

func hxrt__generated_method_field__haxe__iterators__MapKeyValueIterator(value *haxe__iterators__MapKeyValueIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasNext":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_1.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_1.__hx_this.next()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "next")
		}()
	}
	return nil
}

type Type struct {
}

func hxrt_typeCreateEmpty__Tally() *Tally {
	instance := &Tally{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__Exception() *haxe__Exception {
	instance := &haxe__Exception{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__ValueException() *haxe__ValueException {
	instance := &haxe__ValueException{}
	instance.haxe__Exception = &haxe__Exception{}
	instance.haxe__Exception.__hx_this = instance
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe___Int64_____Int64() *haxe___Int64_____Int64 {
	instance := &haxe___Int64_____Int64{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__ds__IntMap() *haxe__ds__IntMap {
	instance := &haxe__ds__IntMap{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__iterators__MapKeyValueIterator() *haxe__iterators__MapKeyValueIterator {
	instance := &haxe__iterators__MapKeyValueIterator{}
	instance.__hx_this = instance
	return instance
}

type ValueType struct {
	tag    int
	params []any
}

func (self *ValueType) HxEnumValue() (string, []any) {
	return [9]string{"TNull", "TInt", "TFloat", "TBool", "TObject", "TFunction", "TClass", "TEnum", "TUnknown"}[self.tag], self.params
}

var ValueType_TNull *ValueType = &ValueType{tag: 0, params: []any{}}

var ValueType_TInt *ValueType = &ValueType{tag: 1, params: []any{}}

var ValueType_TFloat *ValueType = &ValueType{tag: 2, params: []any{}}

var ValueType_TBool *ValueType = &ValueType{tag: 3, params: []any{}}

var ValueType_TObject *ValueType = &ValueType{tag: 4, params: []any{}}

var ValueType_TFunction *ValueType = &ValueType{tag: 5, params: []any{}}

var ValueType_TUnknown *ValueType = &ValueType{tag: 8, params: []any{}}

func ValueType_TClass(c any) *ValueType {
	return &ValueType{tag: 6, params: []any{c}}
}

func ValueType_TEnum(e any) *ValueType {
	return &ValueType{tag: 7, params: []any{e}}
}

func hxrt_typeCallAny(callable any, args []any) (any, bool) {
	result := any(nil)
	ok := false
	defer func() {
		if recover() != nil {
			result = nil
			ok = false
		}
	}()
	if callable == nil {
		return nil, false
	}
	fn := reflect.ValueOf(callable)
	if !fn.IsValid() || fn.Kind() != reflect.Func {
		return nil, false
	}
	fnType := fn.Type()
	if fnType.NumIn() != len(args) {
		return nil, false
	}
	in := make([]reflect.Value, len(args))
	for i := 0; i < len(args); i++ {
		paramType := fnType.In(i)
		arg := args[i]
		if arg == nil {
			in[i] = reflect.Zero(paramType)
			continue
		}
		v := reflect.ValueOf(arg)
		if v.IsValid() && v.Type().AssignableTo(paramType) {
			in[i] = v
			continue
		}
		if v.IsValid() && v.Type().ConvertibleTo(paramType) {
			in[i] = v.Convert(paramType)
			continue
		}
		if paramType.Kind() == reflect.Interface && v.IsValid() {
			in[i] = v
			continue
		}
		return nil, false
	}
	out := fn.Call(in)
	if len(out) == 0 {
		return nil, true
	}
	first := out[0]
	if !first.IsValid() {
		return nil, true
	}
	result = first.Interface()
	ok = true
	return result, ok
}

func hxrt_typeArrayValues(value *hxrt.Array) []any {
	if value == nil {
		return []any{}
	}
	return value.Values()
}

func hxrt_typeResolvedClassName(value any) (string, bool) {
	switch current := value.(type) {
	case *hxrt__TypeClassValue:
		if current == nil || current.name == nil {
			return "", false
		}
		return *current.name, true
	case hxrt__TypeClassValue:
		if current.name == nil {
			return "", false
		}
		return *current.name, true
	case string:
		return current, true
	case *string:
		if current == nil {
			return "", false
		}
		return *current, true
	default:
		return "", false
	}
}

func hxrt_typeResolvedEnumName(value any) (string, bool) {
	switch current := value.(type) {
	case *hxrt__TypeEnumValue:
		if current == nil || current.name == nil {
			return "", false
		}
		return *current.name, true
	case hxrt__TypeEnumValue:
		if current.name == nil {
			return "", false
		}
		return *current.name, true
	case string:
		return current, true
	case *string:
		if current == nil {
			return "", false
		}
		return *current, true
	default:
		return "", false
	}
}

func hxrt_typeCreateClassInstance(className string, args []any) (any, bool) {
	switch className {
	case "Main":
		return nil, false
	case "Reflect":
		return nil, false
	case "Tally":
		return hxrt_typeCallAny(New_Tally, args)
	case "haxe.Exception":
		return hxrt_typeCallAny(New_haxe__Exception, args)
	case "haxe.Int64Helper":
		return nil, false
	case "haxe.Json":
		return nil, false
	case "haxe.ValueException":
		return hxrt_typeCallAny(New_haxe__ValueException, args)
	case "haxe._CallStack.CallStack_Impl_":
		return nil, false
	case "haxe._Int32.Int32_Impl_":
		return nil, false
	case "haxe._Int64.Int64_Impl_":
		return nil, false
	case "haxe._Int64.___Int64":
		return hxrt_typeCallAny(New_haxe___Int64_____Int64, args)
	case "haxe.ds.IntMap":
		return hxrt_typeCallAny(New_haxe__ds__IntMap, args)
	case "haxe.iterators.MapKeyValueIterator":
		return hxrt_typeCallAny(New_haxe__iterators__MapKeyValueIterator, args)
	default:
		return nil, false
	}
}

func hxrt_typeCreateClassEmptyInstance(className string) (any, bool) {
	switch className {
	case "Tally":
		return hxrt_typeCreateEmpty__Tally(), true
	case "haxe.Exception":
		return hxrt_typeCreateEmpty__haxe__Exception(), true
	case "haxe.ValueException":
		return hxrt_typeCreateEmpty__haxe__ValueException(), true
	case "haxe._Int64.___Int64":
		return hxrt_typeCreateEmpty__haxe___Int64_____Int64(), true
	case "haxe.ds.IntMap":
		return hxrt_typeCreateEmpty__haxe__ds__IntMap(), true
	case "haxe.iterators.MapKeyValueIterator":
		return hxrt_typeCreateEmpty__haxe__iterators__MapKeyValueIterator(), true
	default:
		return nil, false
	}
}

func hxrt_typeCreateEnumInstance(enumName string, constructorName string, constructorIndex int, useIndex bool, args []any) (any, bool) {
	switch enumName {
	case "ValueType":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TNull, true
			case 1:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TInt, true
			case 2:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TFloat, true
			case 3:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TBool, true
			case 4:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TObject, true
			case 5:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TFunction, true
			case 6:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(ValueType_TClass, args)
			case 7:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(ValueType_TEnum, args)
			case 8:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TUnknown, true
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "TNull":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TNull, true
		case "TInt":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TInt, true
		case "TFloat":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TFloat, true
		case "TBool":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TBool, true
		case "TObject":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TObject, true
		case "TFunction":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TFunction, true
		case "TClass":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(ValueType_TClass, args)
		case "TEnum":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(ValueType_TEnum, args)
		case "TUnknown":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TUnknown, true
		default:
			return nil, false
		}
	case "haxe.StackItem":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__StackItem_CFunction, true
			case 1:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__StackItem_Module, args)
			case 2:
				if len(args) != 4 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__StackItem_FilePos, args)
			case 3:
				if len(args) != 2 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__StackItem_Method, args)
			case 4:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__StackItem_LocalFunction, args)
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "CFunction":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__StackItem_CFunction, true
		case "Module":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__StackItem_Module, args)
		case "FilePos":
			if len(args) != 4 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__StackItem_FilePos, args)
		case "Method":
			if len(args) != 2 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__StackItem_Method, args)
		case "LocalFunction":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__StackItem_LocalFunction, args)
		default:
			return nil, false
		}
	default:
		return nil, false
	}
}

func Type_getClass(o any) any {
	o = hxrt.ObjectSource(o)
	if hxrt.AnyEqualsNull(o) {
		return nil
	}
	switch value := o.(type) {
	case *hxrt__TypeClassValue:
		if value == nil {
			return nil
		}
		return value
	case hxrt__TypeClassValue:
		copyValue := value
		return &copyValue
	case *hxrt.Array:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Array")}
	case *Tally:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Tally")}
	case *haxe__Exception:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.Exception")}
	case *haxe__ValueException:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ValueException")}
	case *haxe___Int64_____Int64:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe._Int64.___Int64")}
	case *haxe__ds__IntMap:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ds.IntMap")}
	case *haxe__iterators__MapKeyValueIterator:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.iterators.MapKeyValueIterator")}
	default:
		return nil
	}
}

func Type_getEnum(o any) any {
	if hxrt.AnyEqualsNull(o) {
		return nil
	}
	switch value := o.(type) {
	case *hxrt__TypeEnumValue:
		if value == nil {
			return nil
		}
		return value
	case hxrt__TypeEnumValue:
		copyValue := value
		return &copyValue
	case *ValueType:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("ValueType")}
	case *haxe__StackItem:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("haxe.StackItem")}
	default:
		return nil
	}
}

func Type_getSuperClass(c any) any {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return nil
	}
	switch className {
	case "Main":
		return nil
	case "Reflect":
		return nil
	case "Tally":
		return nil
	case "haxe.Exception":
		return nil
	case "haxe.Int64Helper":
		return nil
	case "haxe.Json":
		return nil
	case "haxe.ValueException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.Exception")}
	case "haxe._CallStack.CallStack_Impl_":
		return nil
	case "haxe._Int32.Int32_Impl_":
		return nil
	case "haxe._Int64.Int64_Impl_":
		return nil
	case "haxe._Int64.___Int64":
		return nil
	case "haxe.ds.IntMap":
		return nil
	case "haxe.iterators.MapKeyValueIterator":
		return nil
	default:
		return nil
	}
}

func Type_getClassName(c any) *string {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return nil
	}
	return hxrt.StringFromLiteral(className)
}

func Type_getClassFields(c any) *hxrt.Array {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return hxrt.NewArray()
	}
	switch className {
	case "Main":
		return hxrt.NewArray(hxrt.StringFromLiteral("attempt"), hxrt.StringFromLiteral("boundTypes"), hxrt.StringFromLiteral("byValue"), hxrt.StringFromLiteral("compound"), hxrt.StringFromLiteral("dyn"), hxrt.StringFromLiteral("evaluationOrder"), hxrt.StringFromLiteral("expressionMethods"), hxrt.StringFromLiteral("firstInt"), hxrt.StringFromLiteral("firstText"), hxrt.StringFromLiteral("main"), hxrt.StringFromLiteral("nonArrays"), hxrt.StringFromLiteral("order"), hxrt.StringFromLiteral("otherReceivers"), hxrt.StringFromLiteral("reads"), hxrt.StringFromLiteral("statementMethods"), hxrt.StringFromLiteral("step"), hxrt.StringFromLiteral("useStack"), hxrt.StringFromLiteral("writes"))
	case "Reflect":
		return hxrt.NewArray(hxrt.StringFromLiteral("callMethod"), hxrt.StringFromLiteral("compare"), hxrt.StringFromLiteral("compareMethods"), hxrt.StringFromLiteral("copy"), hxrt.StringFromLiteral("deleteField"), hxrt.StringFromLiteral("field"), hxrt.StringFromLiteral("fields"), hxrt.StringFromLiteral("getProperty"), hxrt.StringFromLiteral("hasField"), hxrt.StringFromLiteral("isEnumValue"), hxrt.StringFromLiteral("isFunction"), hxrt.StringFromLiteral("isObject"), hxrt.StringFromLiteral("makeVarArgs"), hxrt.StringFromLiteral("setField"), hxrt.StringFromLiteral("setProperty"))
	case "Tally":
		return hxrt.NewArray()
	case "haxe.Exception":
		return hxrt.NewArray(hxrt.StringFromLiteral("catchValue"), hxrt.StringFromLiteral("caught"), hxrt.StringFromLiteral("thrown"))
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe.Json":
		return hxrt.NewArray(hxrt.StringFromLiteral("parse"), hxrt.StringFromLiteral("stringify"))
	case "haxe.ValueException":
		return hxrt.NewArray()
	case "haxe._CallStack.CallStack_Impl_":
		return hxrt.NewArray(hxrt.StringFromLiteral("asArray"), hxrt.StringFromLiteral("callStack"), hxrt.StringFromLiteral("copy"), hxrt.StringFromLiteral("equalItems"), hxrt.StringFromLiteral("exceptionStack"), hxrt.StringFromLiteral("exceptionToString"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("get_length"), hxrt.StringFromLiteral("itemToString"), hxrt.StringFromLiteral("length"), hxrt.StringFromLiteral("subtract"), hxrt.StringFromLiteral("toString"))
	case "haxe._Int32.Int32_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.Int64_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.___Int64":
		return hxrt.NewArray()
	case "haxe.ds.IntMap":
		return hxrt.NewArray()
	case "haxe.iterators.MapKeyValueIterator":
		return hxrt.NewArray()
	default:
		return hxrt.NewArray()
	}
}

func Type_getInstanceFields(c any) *hxrt.Array {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return hxrt.NewArray()
	}
	switch className {
	case "Main":
		return hxrt.NewArray()
	case "Reflect":
		return hxrt.NewArray()
	case "Tally":
		return hxrt.NewArray(hxrt.StringFromLiteral("pop"), hxrt.StringFromLiteral("push"), hxrt.StringFromLiteral("total"))
	case "haxe.Exception":
		return hxrt.NewArray(hxrt.StringFromLiteral("details"), hxrt.StringFromLiteral("get_message"), hxrt.StringFromLiteral("get_native"), hxrt.StringFromLiteral("get_previous"), hxrt.StringFromLiteral("get_stack"), hxrt.StringFromLiteral("message"), hxrt.StringFromLiteral("native"), hxrt.StringFromLiteral("previous"), hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("stack"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("unwrap"))
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe.Json":
		return hxrt.NewArray()
	case "haxe.ValueException":
		return hxrt.NewArray(hxrt.StringFromLiteral("details"), hxrt.StringFromLiteral("get_message"), hxrt.StringFromLiteral("get_native"), hxrt.StringFromLiteral("get_previous"), hxrt.StringFromLiteral("get_stack"), hxrt.StringFromLiteral("message"), hxrt.StringFromLiteral("native"), hxrt.StringFromLiteral("previous"), hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("stack"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("unwrap"), hxrt.StringFromLiteral("value"))
	case "haxe._CallStack.CallStack_Impl_":
		return hxrt.NewArray()
	case "haxe._Int32.Int32_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.Int64_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.___Int64":
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	case "haxe.ds.IntMap":
		return hxrt.NewArray(hxrt.StringFromLiteral("clear"), hxrt.StringFromLiteral("copy"), hxrt.StringFromLiteral("copyIMap"), hxrt.StringFromLiteral("exists"), hxrt.StringFromLiteral("existsIMap"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("getIMap"), hxrt.StringFromLiteral("h"), hxrt.StringFromLiteral("iterator"), hxrt.StringFromLiteral("keyValueIterator"), hxrt.StringFromLiteral("keys"), hxrt.StringFromLiteral("remove"), hxrt.StringFromLiteral("removeIMap"), hxrt.StringFromLiteral("set"), hxrt.StringFromLiteral("setIMap"), hxrt.StringFromLiteral("toString"))
	case "haxe.iterators.MapKeyValueIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("keys"), hxrt.StringFromLiteral("map"), hxrt.StringFromLiteral("next"))
	default:
		return hxrt.NewArray()
	}
}

func Type_getEnumName(e any) *string {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	return hxrt.StringFromLiteral(enumName)
}

func Type_resolveClass(name *string) any {
	if name == nil {
		return nil
	}
	rawName := *hxrt.StdString(name)
	switch rawName {
	case "Main":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "Reflect":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "Tally":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Exception":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Int64Helper":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Json":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ValueException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._CallStack.CallStack_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int32.Int32_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int64.Int64_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int64.___Int64":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds.IntMap":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.iterators.MapKeyValueIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	default:
		return nil
	}
}

func Type_resolveEnum(name *string) any {
	if name == nil {
		return nil
	}
	rawName := *hxrt.StdString(name)
	switch rawName {
	case "ValueType":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.StackItem":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	default:
		return nil
	}
}

func Type_createInstance(cl any, args *hxrt.Array) any {
	className, ok := hxrt_typeResolvedClassName(cl)
	if !ok {
		return nil
	}
	instance, ok := hxrt_typeCreateClassInstance(className, hxrt_typeArrayValues(args))
	if !ok {
		return nil
	}
	return instance
}

func Type_createEmptyInstance(cl any) any {
	className, ok := hxrt_typeResolvedClassName(cl)
	if !ok {
		return nil
	}
	instance, ok := hxrt_typeCreateClassEmptyInstance(className)
	if !ok {
		return nil
	}
	return instance
}

func Type_createEnum(e any, constr *string, params *hxrt.Array) any {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	constructorName := ""
	if constr != nil {
		constructorName = *hxrt.StdString(constr)
	}
	enumValue, ok := hxrt_typeCreateEnumInstance(enumName, constructorName, 0, false, hxrt_typeArrayValues(params))
	if !ok {
		return nil
	}
	return enumValue
}

func Type_createEnumIndex(e any, index int, params *hxrt.Array) any {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	enumValue, ok := hxrt_typeCreateEnumInstance(enumName, "", index, true, hxrt_typeArrayValues(params))
	if !ok {
		return nil
	}
	return enumValue
}

func Type_enumConstructor(e any) *string {
	if hxrt.AnyEqualsNull(e) {
		return nil
	}
	switch value := e.(type) {
	case *ValueType:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("TNull")
		case 1:
			return hxrt.StringFromLiteral("TInt")
		case 2:
			return hxrt.StringFromLiteral("TFloat")
		case 3:
			return hxrt.StringFromLiteral("TBool")
		case 4:
			return hxrt.StringFromLiteral("TObject")
		case 5:
			return hxrt.StringFromLiteral("TFunction")
		case 6:
			return hxrt.StringFromLiteral("TClass")
		case 7:
			return hxrt.StringFromLiteral("TEnum")
		case 8:
			return hxrt.StringFromLiteral("TUnknown")
		default:
			return nil
		}
	case *haxe__StackItem:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("CFunction")
		case 1:
			return hxrt.StringFromLiteral("Module")
		case 2:
			return hxrt.StringFromLiteral("FilePos")
		case 3:
			return hxrt.StringFromLiteral("Method")
		case 4:
			return hxrt.StringFromLiteral("LocalFunction")
		default:
			return nil
		}
	default:
		return nil
	}
}

func Type_enumIndex(e any) int {
	if hxrt.AnyEqualsNull(e) {
		return -1
	}
	switch value := e.(type) {
	case *ValueType:
		if value == nil {
			return -1
		}
		return value.tag
	case *haxe__StackItem:
		if value == nil {
			return -1
		}
		return value.tag
	default:
		return -1
	}
}

func Type_getEnumConstructs(e any) *hxrt.Array {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return hxrt.NewArray()
	}
	switch enumName {
	case "ValueType":
		return hxrt.NewArray(hxrt.StringFromLiteral("TNull"), hxrt.StringFromLiteral("TInt"), hxrt.StringFromLiteral("TFloat"), hxrt.StringFromLiteral("TBool"), hxrt.StringFromLiteral("TObject"), hxrt.StringFromLiteral("TFunction"), hxrt.StringFromLiteral("TClass"), hxrt.StringFromLiteral("TEnum"), hxrt.StringFromLiteral("TUnknown"))
	case "haxe.StackItem":
		return hxrt.NewArray(hxrt.StringFromLiteral("CFunction"), hxrt.StringFromLiteral("Module"), hxrt.StringFromLiteral("FilePos"), hxrt.StringFromLiteral("Method"), hxrt.StringFromLiteral("LocalFunction"))
	default:
		return hxrt.NewArray()
	}
}

func Type_enumParameters(e any) *hxrt.Array {
	if hxrt.AnyEqualsNull(e) {
		return hxrt.NewArray()
	}
	switch value := e.(type) {
	case *ValueType:
		if value == nil || value.params == nil {
			return hxrt.NewArray()
		}
		return hxrt.NewArray(value.params...)
	case *haxe__StackItem:
		if value == nil || value.params == nil {
			return hxrt.NewArray()
		}
		return hxrt.NewArray(value.params...)
	default:
		return hxrt.NewArray()
	}
}

func Type_allEnums(e any) *hxrt.Array {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return hxrt.NewArray()
	}
	switch enumName {
	case "ValueType":
		return hxrt.NewArray(ValueType_TNull, ValueType_TInt, ValueType_TFloat, ValueType_TBool, ValueType_TObject, ValueType_TFunction, ValueType_TUnknown)
	case "haxe.StackItem":
		return hxrt.NewArray(haxe__StackItem_CFunction)
	default:
		return hxrt.NewArray()
	}
}

func Type_typeof(v any) *ValueType {
	v = hxrt.ObjectSource(v)
	if hxrt.ReflectIsFunction(v) {
		return ValueType_TFunction
	}
	if hxrt.AnyEqualsNull(v) {
		return ValueType_TNull
	}
	switch v.(type) {
	case *hxrt__TypeClassValue, hxrt__TypeClassValue, *hxrt__TypeEnumValue, hxrt__TypeEnumValue:
		return ValueType_TObject
	}
	if enumValue := Type_getEnum(v); enumValue != nil {
		return ValueType_TEnum(enumValue)
	}
	if classValue := Type_getClass(v); classValue != nil {
		return ValueType_TClass(classValue)
	}
	switch v.(type) {
	case bool:
		return ValueType_TBool
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr:
		return ValueType_TInt
	case float32, float64:
		return ValueType_TFloat
	case string, *string:
		return ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("String")})
	case *hxrt.Array:
		return ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Array")})
	}
	ref := reflect.ValueOf(v)
	if !ref.IsValid() {
		return ValueType_TNull
	}
	switch ref.Kind() {
	case reflect.Slice, reflect.Array:
		return ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Array")})
	case reflect.Map, reflect.Struct, reflect.Interface, reflect.Pointer:
		return ValueType_TObject
	default:
		return ValueType_TUnknown
	}
}

func Type_enumEq(a any, b any) bool {
	return reflect.DeepEqual(a, b)
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
	case *Tally:
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
	case *haxe__ds__IntMap:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *Tally:
		return hxrt__generated_field_lookup__Tally(value, key)
	case *haxe__Exception:
		return hxrt__generated_field_lookup__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_lookup__haxe__ValueException(value, key)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_lookup__haxe___Int64_____Int64(value, key)
	case *haxe__ds__IntMap:
		return hxrt__generated_field_lookup__haxe__ds__IntMap(value, key)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_field_lookup__haxe__iterators__MapKeyValueIterator(value, key)
	default:
		return nil
	}
}

func reflaxe__go___internal__CompilerReflect_hasGeneratedField(object any, field *string) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *Tally:
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
	case *haxe__ds__IntMap:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *Tally:
		return hxrt__generated_field_has__Tally(value, key)
	case *haxe__Exception:
		return hxrt__generated_field_has__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_has__haxe__ValueException(value, key)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_has__haxe___Int64_____Int64(value, key)
	case *haxe__ds__IntMap:
		return hxrt__generated_field_has__haxe__ds__IntMap(value, key)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_field_has__haxe__iterators__MapKeyValueIterator(value, key)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_setGeneratedField(object any, field *string, incoming any) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *Tally:
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
	case *haxe__ds__IntMap:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *Tally:
		return hxrt__generated_field_set__Tally(value, key, incoming)
	case *haxe__Exception:
		return hxrt__generated_field_set__haxe__Exception(value, key, incoming)
	case *haxe__ValueException:
		return hxrt__generated_field_set__haxe__ValueException(value, key, incoming)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_set__haxe___Int64_____Int64(value, key, incoming)
	case *haxe__ds__IntMap:
		return hxrt__generated_field_set__haxe__ds__IntMap(value, key, incoming)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_field_set__haxe__iterators__MapKeyValueIterator(value, key, incoming)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedFields(object any) *hxrt.Array {
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *Tally:
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
	case *haxe__ds__IntMap:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch receiver.(type) {
	case *Tally:
		return hxrt.NewArray(hxrt.StringFromLiteral("total"))
	case *haxe__Exception:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"))
	case *haxe__ValueException:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("value"))
	case *haxe___Int64_____Int64:
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	case *haxe__ds__IntMap:
		return hxrt.NewArray(hxrt.StringFromLiteral("h"))
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("map"), hxrt.StringFromLiteral("keys"))
	default:
		return nil
	}
}

func hxrt__generated_field_lookup__Tally(value *Tally, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "total":
		return hxrt.ReadStoredField(value.total)
	}
	return nil
}

func hxrt__generated_field_has__Tally(value *Tally, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "total":
		return true
	}
	return false
}

func hxrt__generated_field_set__Tally(value *Tally, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "total":
		if incoming == nil {
			var zero int
			value.total = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.total = typed
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

func hxrt__generated_field_lookup__haxe__ds__IntMap(value *haxe__ds__IntMap, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "h":
		return hxrt.ReadStoredField(value.h)
	}
	return nil
}

func hxrt__generated_field_has__haxe__ds__IntMap(value *haxe__ds__IntMap, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "h":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__ds__IntMap(value *haxe__ds__IntMap, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "h":
		if incoming == nil {
			var zero *hxrt.IntMapCell
			value.h = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.IntMapCell:
			value.h = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__iterators__MapKeyValueIterator(value *haxe__iterators__MapKeyValueIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "keys":
		return hxrt.ReadStoredField(value.keys)
	case "map":
		return hxrt.ReadStoredField(value.map_)
	}
	return nil
}

func hxrt__generated_field_has__haxe__iterators__MapKeyValueIterator(value *haxe__iterators__MapKeyValueIterator, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "keys":
		return true
	case "map":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__iterators__MapKeyValueIterator(value *haxe__iterators__MapKeyValueIterator, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "keys":
		if incoming == nil {
			var zero hxrt.Anonymous
			value.keys = zero
			return true
		}
		switch typed := incoming.(type) {
		case hxrt.Anonymous:
			value.keys = typed
			return true
		default:
			return false
		}
	case "map":
		if incoming == nil {
			var zero haxe__IMap
			value.map_ = zero
			return true
		}
		switch typed := incoming.(type) {
		case haxe__IMap:
			value.map_ = typed
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
	case *hxrt.Function[func() float64]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(int) bool]:
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
	case *hxrt.Function[func(int) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_3 any) int {
			if hx_value_3 == nil {
				var hx_zero_4 int
				return hx_zero_4
			}
			return hx_value_3.(int)
		}(argument_0))
	case *hxrt.Function[func() any]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *string]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() int]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(any, any) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(argument_0, argument_1)
	case *hxrt.Function[func() bool]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(...any) any]:
		var argument_0 []any = nil
		for index, argument := range arguments {
			if index >= 0 {
				argument_0 = append(argument_0, argument)
			}
		}
		return hxrt.FunctionCall(source)(argument_0...)
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func()]:
		hxrt.FunctionCall(source)()
		return nil
	case *hxrt.Function[func() *haxe__ds__IntMap]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(any) bool]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(argument_0)
	case *hxrt.Function[func(int) any]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_5 any) int {
			if hx_value_5 == nil {
				var hx_zero_6 int
				return hx_zero_6
			}
			return hx_value_5.(int)
		}(argument_0))
	case *hxrt.Function[func(any) any]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(argument_0)
	case *hxrt.Function[func(int, any)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_7 any) int {
			if hx_value_7 == nil {
				var hx_zero_8 int
				return hx_zero_8
			}
			return hx_value_7.(int)
		}(argument_0), argument_1)
		return nil
	case *hxrt.Function[func(any, any)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(argument_0, argument_1)
		return nil
	default:
		return hxrt.ReflectCallMethod(value, arguments)
	}
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
	case *hxrt.Function[func(int) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 int) int {
			return func(hx_value_1 any) int {
				if hx_value_1 == nil {
					var hx_zero_2 int
					return hx_zero_2
				}
				return hx_value_1.(int)
			}(hxrt.FunctionCall(source)(arg_0))
		})
	case *hxrt.Function[func(int) int]:
		return source
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 int) int {
			return func(hx_value_3 any) int {
				if hx_value_3 == nil {
					var hx_zero_4 int
					return hx_zero_4
				}
				return hx_value_3.(int)
			}(hxrt.FunctionCall(source)(arg_0))
		})
	case *hxrt.Function[func(any) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 int) int {
			return func(hx_value_5 any) int {
				if hx_value_5 == nil {
					var hx_zero_6 int
					return hx_zero_6
				}
				return hx_value_5.(int)
			}(hxrt.FunctionCall(source)(arg_0))
		})
	case *hxrt.Function[func(int) any]:
		return hxrt.ProjectFunction(source, func(arg_0 int) int {
			return func(hx_value_7 any) int {
				if hx_value_7 == nil {
					var hx_zero_8 int
					return hx_zero_8
				}
				return hx_value_7.(int)
			}(hxrt.FunctionCall(source)(arg_0))
		})
	case *hxrt.Function[func(any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 int) int {
			return func(hx_value_9 any) int {
				if hx_value_9 == nil {
					var hx_zero_10 int
					return hx_zero_10
				}
				return hx_value_9.(int)
			}(hxrt.FunctionCall(source)(arg_0))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
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
	case *hxrt.Function[func() float64]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() any]:
		return source
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() any {
			hxrt.FunctionCall(source)()
			return nil
		})
	case *hxrt.Function[func() *haxe__ds__IntMap]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_cf6928b1(value any) *hxrt.Function[func() bool] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() bool]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() float64]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_1 any) bool {
				if hx_value_1 == nil {
					var hx_zero_2 bool
					return hx_zero_2
				}
				return hx_value_1.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_3 any) bool {
				if hx_value_3 == nil {
					var hx_zero_4 bool
					return hx_zero_4
				}
				return hx_value_3.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_5 any) bool {
				if hx_value_5 == nil {
					var hx_zero_6 bool
					return hx_zero_6
				}
				return hx_value_5.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_7 any) bool {
				if hx_value_7 == nil {
					var hx_zero_8 bool
					return hx_zero_8
				}
				return hx_value_7.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() bool]:
		return source
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_9 any) bool {
				if hx_value_9 == nil {
					var hx_zero_10 bool
					return hx_zero_10
				}
				return hx_value_9.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_11 any) bool {
				if hx_value_11 == nil {
					var hx_zero_12 bool
					return hx_zero_12
				}
				return hx_value_11.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_13 any) bool {
				if hx_value_13 == nil {
					var hx_zero_14 bool
					return hx_zero_14
				}
				return hx_value_13.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() bool {
			hxrt.FunctionCall(source)()
			return func(hx_value_15 any) bool {
				if hx_value_15 == nil {
					var hx_zero_16 bool
					return hx_zero_16
				}
				return hx_value_15.(bool)
			}(nil)
		})
	case *hxrt.Function[func() *haxe__ds__IntMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_17 any) bool {
				if hx_value_17 == nil {
					var hx_zero_18 bool
					return hx_zero_18
				}
				return hx_value_17.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_19 any) bool {
				if hx_value_19 == nil {
					var hx_zero_20 bool
					return hx_zero_20
				}
				return hx_value_19.(bool)
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
	case *hxrt.Function[func() float64]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_1 any) int {
				if hx_value_1 == nil {
					var hx_zero_2 int
					return hx_zero_2
				}
				return hx_value_1.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_3 any) int {
				if hx_value_3 == nil {
					var hx_zero_4 int
					return hx_zero_4
				}
				return hx_value_3.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_5 any) int {
				if hx_value_5 == nil {
					var hx_zero_6 int
					return hx_zero_6
				}
				return hx_value_5.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() int]:
		return source
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_7 any) int {
				if hx_value_7 == nil {
					var hx_zero_8 int
					return hx_zero_8
				}
				return hx_value_7.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_9 any) int {
				if hx_value_9 == nil {
					var hx_zero_10 int
					return hx_zero_10
				}
				return hx_value_9.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_11 any) int {
				if hx_value_11 == nil {
					var hx_zero_12 int
					return hx_zero_12
				}
				return hx_value_11.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_13 any) int {
				if hx_value_13 == nil {
					var hx_zero_14 int
					return hx_zero_14
				}
				return hx_value_13.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() int {
			hxrt.FunctionCall(source)()
			return func(hx_value_15 any) int {
				if hx_value_15 == nil {
					var hx_zero_16 int
					return hx_zero_16
				}
				return hx_value_15.(int)
			}(nil)
		})
	case *hxrt.Function[func() *haxe__ds__IntMap]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_17 any) int {
				if hx_value_17 == nil {
					var hx_zero_18 int
					return hx_zero_18
				}
				return hx_value_17.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_19 any) int {
				if hx_value_19 == nil {
					var hx_zero_20 int
					return hx_zero_20
				}
				return hx_value_19.(int)
			}(hxrt.FunctionCall(source)())
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

func hx_recover_class_haxe__ds__IntMap(value any) *haxe__ds__IntMap {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ds__IntMap:
		return source
	default:
		return value.(*haxe__ds__IntMap)
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
