package main

import "snapshot/hxrt"

func attempt(action *hxrt.Function[func() *string]) *string {
	hx_try_return_1 := false
	var hx_try_value_2 *string
	hxrt.TryCatch(func() {
		hx_try_value_2 = func() *string {
			var hx_callable_6 *hxrt.Function[func() *string] = action
			return hxrt.FunctionCall(hx_callable_6)()
		}()
		hx_try_return_1 = true
		return
	}, func(hx_caught_3 any) {
		hx_exception_4 := haxe__Exception_caught(hx_caught_3)
		hx_payload_5 := haxe__Exception_catchValue(hx_caught_3)
		_ = hx_exception_4
		var hx_tmp any = hx_payload_5
		_ = hx_tmp
		hx_try_value_2 = hxrt.StringFromLiteral("threw")
		hx_try_return_1 = true
		return
	})
	if hx_try_return_1 {
		return hx_try_value_2
	}
	var hx_try_zero_7 *string
	return hx_try_zero_7
}

func cells() *hxrt.Array {
	return hxrt.NewArray(nil, 1.5)
}

func main() {
	var empty any = none()
	var half any = some(0.5)
	two := 2.0
	three := 3
	row(hxrt.StringFromLiteral("add"), hxrt.NewArray(attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (func(hx_value_8 any) float64 {
			if hx_value_8 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_8.(float64)
		}(empty) + float64(1)))
	})), attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (float64(1) + func(hx_value_9 any) float64 {
			if hx_value_9 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_9.(float64)
		}(empty)))
	})), attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (func(hx_value_10 any) float64 {
			if hx_value_10 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_10.(float64)
		}(empty) + func(hx_value_11 any) float64 {
			if hx_value_11 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_11.(float64)
		}(empty)))
	})), attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (func(hx_value_12 any) float64 {
			if hx_value_12 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_12.(float64)
		}(empty) + two))
	})), attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (float64(three) + func(hx_value_13 any) float64 {
			if hx_value_13 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_13.(float64)
		}(empty)))
	})), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (func(hx_value_14 any) float64 {
		if hx_value_14 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_14.(float64)
	}(half)+float64(1))), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (float64(1)+func(hx_value_15 any) float64 {
		if hx_value_15 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_15.(float64)
	}(half))), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (func(hx_value_16 any) float64 {
		if hx_value_16 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_16.(float64)
	}(half)+func(hx_value_17 any) float64 {
		if hx_value_17 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_17.(float64)
	}(half))), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (float64(three)+func(hx_value_18 any) float64 {
		if hx_value_18 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_18.(float64)
	}(half)))))
	row(hxrt.StringFromLiteral("sub"), hxrt.NewArray(attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (func(hx_value_19 any) float64 {
			if hx_value_19 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_19.(float64)
		}(empty) - float64(1)))
	})), attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (two - func(hx_value_20 any) float64 {
			if hx_value_20 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_20.(float64)
		}(empty)))
	})), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (func(hx_value_21 any) float64 {
		if hx_value_21 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_21.(float64)
	}(half)-float64(1))), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (two-func(hx_value_22 any) float64 {
		if hx_value_22 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_22.(float64)
	}(half)))))
	row(hxrt.StringFromLiteral("mul"), hxrt.NewArray(attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (func(hx_value_23 any) float64 {
			if hx_value_23 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_23.(float64)
		}(empty) * float64(2)))
	})), attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (two * func(hx_value_24 any) float64 {
			if hx_value_24 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_24.(float64)
		}(empty)))
	})), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (func(hx_value_25 any) float64 {
		if hx_value_25 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_25.(float64)
	}(half)*float64(4))), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (two*func(hx_value_26 any) float64 {
		if hx_value_26 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_26.(float64)
	}(half)))))
	row(hxrt.StringFromLiteral("div"), hxrt.NewArray(attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (func(hx_value_27 any) float64 {
			if hx_value_27 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_27.(float64)
		}(empty) / float64(2)))
	})), attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (two / func(hx_value_28 any) float64 {
			if hx_value_28 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_28.(float64)
		}(empty)))
	})), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (func(hx_value_29 any) float64 {
		if hx_value_29 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_29.(float64)
	}(half)/float64(2))), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (two/func(hx_value_30 any) float64 {
		if hx_value_30 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_30.(float64)
	}(half)))))
	row(hxrt.StringFromLiteral("mod"), hxrt.NewArray(attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), hxrt.FloatMod(func(hx_value_31 any) float64 {
			if hx_value_31 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_31.(float64)
		}(empty), float64(2)))
	})), attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), hxrt.FloatMod(two, func(hx_value_32 any) float64 {
			if hx_value_32 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_32.(float64)
		}(empty)))
	})), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), hxrt.FloatMod(func(hx_value_33 any) float64 {
		if hx_value_33 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_33.(float64)
	}(half), float64(2))), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), hxrt.FloatMod(two, func(hx_value_34 any) float64 {
		if hx_value_34 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_34.(float64)
	}(half)))))
	row(hxrt.StringFromLiteral("compare.null"), hxrt.NewArray(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString((func(hx_value_35 any) float64 {
		if hx_value_35 == nil {
			var hx_zero_36 float64
			return (0 / hx_zero_36)
		}
		return hx_value_35.(float64)
	}(empty) < 1))), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString((1 < func(hx_value_37 any) float64 {
		if hx_value_37 == nil {
			var hx_zero_38 float64
			return (0 / hx_zero_38)
		}
		return hx_value_37.(float64)
	}(empty)))), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString((func(hx_value_39 any) float64 {
		if hx_value_39 == nil {
			var hx_zero_40 float64
			return (0 / hx_zero_40)
		}
		return hx_value_39.(float64)
	}(empty) <= 1))), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString((1 <= func(hx_value_41 any) float64 {
		if hx_value_41 == nil {
			var hx_zero_42 float64
			return (0 / hx_zero_42)
		}
		return hx_value_41.(float64)
	}(empty)))), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString((func(hx_value_43 any) float64 {
		if hx_value_43 == nil {
			var hx_zero_44 float64
			return (0 / hx_zero_44)
		}
		return hx_value_43.(float64)
	}(empty) > 1))), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString((1 > func(hx_value_45 any) float64 {
		if hx_value_45 == nil {
			var hx_zero_46 float64
			return (0 / hx_zero_46)
		}
		return hx_value_45.(float64)
	}(empty)))), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString((func(hx_value_47 any) float64 {
		if hx_value_47 == nil {
			var hx_zero_48 float64
			return (0 / hx_zero_48)
		}
		return hx_value_47.(float64)
	}(empty) >= 1))), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString((1 >= func(hx_value_49 any) float64 {
		if hx_value_49 == nil {
			var hx_zero_50 float64
			return (0 / hx_zero_50)
		}
		return hx_value_49.(float64)
	}(empty)))), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString(hxrt.NullableFloatsLess(empty, empty))), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString((func(hx_value_51 any) float64 {
		if hx_value_51 == nil {
			var hx_zero_52 float64
			return (0 / hx_zero_52)
		}
		return hx_value_51.(float64)
	}(empty) >= two))), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString((two <= func(hx_value_53 any) float64 {
		if hx_value_53 == nil {
			var hx_zero_54 float64
			return (0 / hx_zero_54)
		}
		return hx_value_53.(float64)
	}(empty)))), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString(!(func(hx_value_55 any) float64 {
		if hx_value_55 == nil {
			var hx_zero_56 float64
			return (0 / hx_zero_56)
		}
		return hx_value_55.(float64)
	}(empty) < 1)))))
	row(hxrt.StringFromLiteral("compare.value"), hxrt.NewArray(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString((func(hx_value_57 any) float64 {
		if hx_value_57 == nil {
			var hx_zero_58 float64
			return (0 / hx_zero_58)
		}
		return hx_value_57.(float64)
	}(half) < 1))), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString((1 < func(hx_value_59 any) float64 {
		if hx_value_59 == nil {
			var hx_zero_60 float64
			return (0 / hx_zero_60)
		}
		return hx_value_59.(float64)
	}(half)))), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString((func(hx_value_61 any) float64 {
		if hx_value_61 == nil {
			var hx_zero_62 float64
			return (0 / hx_zero_62)
		}
		return hx_value_61.(float64)
	}(half) <= 0.5))), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString((func(hx_value_63 any) float64 {
		if hx_value_63 == nil {
			var hx_zero_64 float64
			return (0 / hx_zero_64)
		}
		return hx_value_63.(float64)
	}(half) >= two))), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString((two > func(hx_value_65 any) float64 {
		if hx_value_65 == nil {
			var hx_zero_66 float64
			return (0 / hx_zero_66)
		}
		return hx_value_65.(float64)
	}(half)))), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString(hxrt.NullableFloatsGreater(half, half)))))
	row(hxrt.StringFromLiteral("condition"), hxrt.NewArray(func() *string {
		var hx_if_69 *string
		if func(hx_value_67 any) float64 {
			if hx_value_67 == nil {
				var hx_zero_68 float64
				return (0 / hx_zero_68)
			}
			return hx_value_67.(float64)
		}(empty) < 1 {
			hx_if_69 = hxrt.StringFromLiteral("taken")
		} else {
			hx_if_69 = hxrt.StringFromLiteral("skipped")
		}
		return hx_if_69
	}(), func() *string {
		var hx_if_72 *string
		if func(hx_value_70 any) float64 {
			if hx_value_70 == nil {
				var hx_zero_71 float64
				return (0 / hx_zero_71)
			}
			return hx_value_70.(float64)
		}(empty) >= 0 {
			hx_if_72 = hxrt.StringFromLiteral("taken")
		} else {
			hx_if_72 = hxrt.StringFromLiteral("skipped")
		}
		return hx_if_72
	}(), func() *string {
		var hx_if_75 *string
		if !(func(hx_value_73 any) float64 {
			if hx_value_73 == nil {
				var hx_zero_74 float64
				return (0 / hx_zero_74)
			}
			return hx_value_73.(float64)
		}(empty) > 0) {
			hx_if_75 = hxrt.StringFromLiteral("negated")
		} else {
			hx_if_75 = hxrt.StringFromLiteral("kept")
		}
		return hx_if_75
	}(), func() *string {
		var hx_if_78 *string
		if func(hx_value_76 any) float64 {
			if hx_value_76 == nil {
				var hx_zero_77 float64
				return (0 / hx_zero_77)
			}
			return hx_value_76.(float64)
		}(half) < 1 {
			hx_if_78 = hxrt.StringFromLiteral("taken")
		} else {
			hx_if_78 = hxrt.StringFromLiteral("skipped")
		}
		return hx_if_78
	}()))
	row(hxrt.StringFromLiteral("equal"), hxrt.NewArray(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString((empty == nil))), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString((empty != nil))), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString((empty == 1.0))), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString((empty != 1.0))), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString((half == nil))), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString((half == 0.5))), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString((half != 0.5)))))
	var sum any = none()
	var diff any = none()
	var product any = none()
	var quotient any = none()
	var rest any = none()
	row(hxrt.StringFromLiteral("compound.null"), hxrt.NewArray(attempt(hxrt.NewFunction(func() *string {
		sum = (func(hx_value_79 any) float64 {
			if hx_value_79 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_79.(float64)
		}(sum) + float64(1))
		return hxrt.StringFromLiteral("stepped")
	})), attempt(hxrt.NewFunction(func() *string {
		diff = (func(hx_value_80 any) float64 {
			if hx_value_80 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_80.(float64)
		}(diff) - float64(1))
		return hxrt.StringFromLiteral("stepped")
	})), attempt(hxrt.NewFunction(func() *string {
		product = (func(hx_value_81 any) float64 {
			if hx_value_81 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_81.(float64)
		}(product) * float64(2))
		return hxrt.StringFromLiteral("stepped")
	})), attempt(hxrt.NewFunction(func() *string {
		quotient = (func(hx_value_82 any) float64 {
			if hx_value_82 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_82.(float64)
		}(quotient) / float64(2))
		return hxrt.StringFromLiteral("stepped")
	})), attempt(hxrt.NewFunction(func() *string {
		rest = hxrt.FloatMod(func(hx_value_83 any) float64 {
			if hx_value_83 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_83.(float64)
		}(rest), float64(2))
		return hxrt.StringFromLiteral("stepped")
	})), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), sum), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), diff), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), product), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), quotient), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), rest)))
	plain := 1.0
	row(hxrt.StringFromLiteral("compound.rhs"), hxrt.NewArray(attempt(hxrt.NewFunction(func() *string {
		plain = (plain + func(hx_value_84 any) float64 {
			if hx_value_84 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_84.(float64)
		}(empty))
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), plain)
	})), attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), func() float64 {
			plain = (plain * func(hx_value_85 any) float64 {
				if hx_value_85 == nil {
					hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
				}
				return hx_value_85.(float64)
			}(empty))
			return plain
		}())
	})), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), plain)))
	var level any = some(1.5)
	level = (func(hx_value_86 any) float64 {
		if hx_value_86 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_86.(float64)
	}(level) + float64(1))
	level = (func(hx_value_87 any) float64 {
		if hx_value_87 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_87.(float64)
	}(level) - 0.5)
	level = (func(hx_value_88 any) float64 {
		if hx_value_88 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_88.(float64)
	}(level) * float64(3))
	level = (func(hx_value_89 any) float64 {
		if hx_value_89 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_89.(float64)
	}(level) / float64(4))
	level = hxrt.FloatMod(func(hx_value_90 any) float64 {
		if hx_value_90 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_90.(float64)
	}(level), float64(1))
	row(hxrt.StringFromLiteral("compound.value"), hxrt.NewArray(hxrt.StringConcatAny(hxrt.StringFromLiteral(""), level)))
	meter := New_Meter()
	row(hxrt.StringFromLiteral("field"), hxrt.NewArray(attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (func(hx_value_91 any) float64 {
			if hx_value_91 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_91.(float64)
		}(meter.reading) + float64(1)))
	})), attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (float64(2) * func(hx_value_92 any) float64 {
			if hx_value_92 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_92.(float64)
		}(meter.reading)))
	})), attempt(hxrt.NewFunction(func() *string {
		meter.reading = (func(hx_value_93 any) float64 {
			if hx_value_93 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_93.(float64)
		}(meter.reading) + float64(1))
		return hxrt.StringFromLiteral("stepped")
	})), attempt(hxrt.NewFunction(func() *string {
		meter.reading = (func(hx_value_94 any) float64 {
			if hx_value_94 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_94.(float64)
		}(meter.reading) - float64(1))
		return hxrt.StringFromLiteral("stepped")
	})), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString((func(hx_value_95 any) float64 {
		if hx_value_95 == nil {
			var hx_zero_96 float64
			return (0 / hx_zero_96)
		}
		return hx_value_95.(float64)
	}(meter.reading) < 1))), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString((func(hx_value_97 any) float64 {
		if hx_value_97 == nil {
			var hx_zero_98 float64
			return (0 / hx_zero_98)
		}
		return hx_value_97.(float64)
	}(meter.reading) >= 0))), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), meter.reading)))
	row(hxrt.StringFromLiteral("static"), hxrt.NewArray(attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (func(hx_value_99 any) float64 {
			if hx_value_99 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_99.(float64)
		}(Meter_shared) - float64(1)))
	})), attempt(hxrt.NewFunction(func() *string {
		Meter_shared = (func(hx_value_100 any) float64 {
			if hx_value_100 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_100.(float64)
		}(Meter_shared) * float64(2))
		return hxrt.StringFromLiteral("stepped")
	})), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString((func(hx_value_101 any) float64 {
		if hx_value_101 == nil {
			var hx_zero_102 float64
			return (0 / hx_zero_102)
		}
		return hx_value_101.(float64)
	}(Meter_shared) > 0))), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), Meter_shared)))
	meter.reading = 2.5
	Meter_shared = 4.0
	meter.reading = (func(hx_value_103 any) float64 {
		if hx_value_103 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_103.(float64)
	}(meter.reading) + func(hx_value_104 any) float64 {
		if hx_value_104 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_104.(float64)
	}(Meter_shared))
	row(hxrt.StringFromLiteral("field.value"), hxrt.NewArray(hxrt.StringConcatAny(hxrt.StringFromLiteral(""), meter.reading), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (func(hx_value_105 any) float64 {
		if hx_value_105 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_105.(float64)
	}(Meter_shared)*func(hx_value_106 any) float64 {
		if hx_value_106 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_106.(float64)
	}(meter.reading))), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString(hxrt.NullableFloatsGreater(meter.reading, Meter_shared)))))
	var gauge_level any
	gauge_level = nil
	row(hxrt.StringFromLiteral("struct"), hxrt.NewArray(attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (func(hx_value_107 any) float64 {
			if hx_value_107 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_107.(float64)
		}(gauge_level) + float64(1)))
	})), attempt(hxrt.NewFunction(func() *string {
		gauge_level = (func(hx_value_108 any) float64 {
			if hx_value_108 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_108.(float64)
		}(gauge_level) / float64(2))
		return hxrt.StringFromLiteral("stepped")
	})), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString((func(hx_value_109 any) float64 {
		if hx_value_109 == nil {
			var hx_zero_110 float64
			return (0 / hx_zero_110)
		}
		return hx_value_109.(float64)
	}(gauge_level) <= 0))), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), gauge_level)))
	gauge_level = 3.0
	gauge_level = (func(hx_value_111 any) float64 {
		if hx_value_111 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_111.(float64)
	}(gauge_level) - 0.5)
	row(hxrt.StringFromLiteral("struct.value"), hxrt.NewArray(hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (func(hx_value_112 any) float64 {
		if hx_value_112 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_112.(float64)
	}(gauge_level)*float64(2))), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString((func(hx_value_113 any) float64 {
		if hx_value_113 == nil {
			var hx_zero_114 float64
			return (0 / hx_zero_114)
		}
		return hx_value_113.(float64)
	}(gauge_level) > 2)))))
	slots := cells()
	row(hxrt.StringFromLiteral("array"), hxrt.NewArray(attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (func(hx_value_115 any) float64 {
			if hx_value_115 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_115.(float64)
		}(hxrt.DynamicToNullableFloat(slots.Get(0))) + float64(1)))
	})), attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (float64(3) - func(hx_value_116 any) float64 {
			if hx_value_116 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_116.(float64)
		}(hxrt.DynamicToNullableFloat(slots.Get(0)))))
	})), attempt(hxrt.NewFunction(func() *string {
		func() any {
			hx_array_target_117 := slots
			hx_array_index_118 := 0
			var hx_array_current_119 any = hxrt.DynamicToNullableFloat(hx_array_target_117.Get(hx_array_index_118))
			var hx_array_assigned_120 any = (func(hx_value_121 any) float64 {
				if hx_value_121 == nil {
					hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
				}
				return hx_value_121.(float64)
			}(hx_array_current_119) + float64(1))
			hx_array_target_117.Set(hx_array_index_118, hx_array_assigned_120)
			return hx_array_assigned_120
		}()
		return hxrt.StringFromLiteral("stepped")
	})), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString((func(hx_value_122 any) float64 {
		if hx_value_122 == nil {
			var hx_zero_123 float64
			return (0 / hx_zero_123)
		}
		return hx_value_122.(float64)
	}(hxrt.DynamicToNullableFloat(slots.Get(0))) < 2))), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString((2 > func(hx_value_124 any) float64 {
		if hx_value_124 == nil {
			var hx_zero_125 float64
			return (0 / hx_zero_125)
		}
		return hx_value_124.(float64)
	}(hxrt.DynamicToNullableFloat(slots.Get(0)))))), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), slots.Get(0)), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (func(hx_value_126 any) float64 {
		if hx_value_126 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_126.(float64)
	}(hxrt.DynamicToNullableFloat(slots.Get(1)))+float64(1))), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString((func(hx_value_127 any) float64 {
		if hx_value_127 == nil {
			var hx_zero_128 float64
			return (0 / hx_zero_128)
		}
		return hx_value_127.(float64)
	}(hxrt.DynamicToNullableFloat(slots.Get(1))) < 2))), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (func(hx_value_129 any) float64 {
		if hx_value_129 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_129.(float64)
	}(hxrt.DynamicToNullableFloat(slots.Get(1)))*func(hx_value_130 any) float64 {
		if hx_value_130 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_130.(float64)
	}(hxrt.DynamicToNullableFloat(slots.Get(1)))))))
	row(hxrt.StringFromLiteral("call"), hxrt.NewArray(attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (func(hx_value_131 any) float64 {
			if hx_value_131 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_131.(float64)
		}(none()) + float64(1)))
	})), attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (float64(1) - func(hx_value_132 any) float64 {
			if hx_value_132 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_132.(float64)
		}(none())))
	})), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString((func(hx_value_133 any) float64 {
		if hx_value_133 == nil {
			var hx_zero_134 float64
			return (0 / hx_zero_134)
		}
		return hx_value_133.(float64)
	}(none()) < 1))), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (func(hx_value_135 any) float64 {
		if hx_value_135 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_135.(float64)
	}(some(2))*float64(2)))))
	var count any = nil
	count = 2
	row(hxrt.StringFromLiteral("mixed"), hxrt.NewArray(attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (float64(hxrt.IntFromNullableAny(count)) + func(hx_value_136 any) float64 {
			if hx_value_136 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_136.(float64)
		}(empty)))
	})), attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (func(hx_value_137 any) float64 {
			if hx_value_137 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_137.(float64)
		}(empty) * float64(hxrt.IntFromNullableAny(count))))
	})), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (float64(hxrt.IntFromNullableAny(count))+func(hx_value_138 any) float64 {
		if hx_value_138 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_138.(float64)
	}(half))), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (func(hx_value_139 any) float64 {
		if hx_value_139 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_139.(float64)
	}(half)*float64(hxrt.IntFromNullableAny(count))))))
	text := hxrt.StringFromLiteral("s")
	text = hxrt.StringConcatAny(text, empty)
	row(hxrt.StringFromLiteral("text"), hxrt.NewArray(hxrt.StringConcatAny(hxrt.StringFromLiteral("x"), empty), hxrt.StringConcatAny(empty, hxrt.StringFromLiteral("x")), text, hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringFromLiteral("a"), empty), 1), hxrt.StringConcatAny(half, hxrt.StringFromLiteral("!")), hxrt.StringConcatAny(hxrt.StringFromLiteral("v"), meter.reading)))
	var maybe any = none()
	var known any = some(1.25)
	row(hxrt.StringFromLiteral("narrowed"), hxrt.NewArray(func() *string {
		var hx_if_140 *string
		if maybe != nil {
			hx_if_140 = hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (maybe.(float64) + float64(1)))
		} else {
			hx_if_140 = hxrt.StringFromLiteral("none")
		}
		return hx_if_140
	}(), func() *string {
		var hx_if_141 *string
		if known != nil {
			hx_if_141 = hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (known.(float64) + float64(1)))
		} else {
			hx_if_141 = hxrt.StringFromLiteral("none")
		}
		return hx_if_141
	}(), func() *string {
		var hx_if_144 *string
		if (known != nil) && (func(hx_value_142 any) float64 {
			if hx_value_142 == nil {
				var hx_zero_143 float64
				return (0 / hx_zero_143)
			}
			return hx_value_142.(float64)
		}(known) < 2) {
			hx_if_144 = hxrt.StringFromLiteral("small")
		} else {
			hx_if_144 = hxrt.StringFromLiteral("large")
		}
		return hx_if_144
	}()))
	order := hxrt.NewArray(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString((func(hx_value_145 any) float64 {
		if hx_value_145 == nil {
			var hx_zero_146 float64
			return (0 / hx_zero_146)
		}
		return hx_value_145.(float64)
	}(empty) < tick(hxrt.StringFromLiteral("compare"), 1)))), attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (tick(hxrt.StringFromLiteral("left"), 1) + func(hx_value_147 any) float64 {
			if hx_value_147 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_147.(float64)
		}(empty)))
	})), attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (tick(hxrt.StringFromLiteral("scaled"), 2) * func(hx_value_148 any) float64 {
			if hx_value_148 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_148.(float64)
		}(empty)))
	})))
	row(hxrt.StringFromLiteral("order"), func() *hxrt.Array {
		hx_concat_target_149 := order
		hx_concat_other_150 := hxrt.NewArray(hxrt.StringJoinAny(trail.Values(), hxrt.StringFromLiteral(",")))
		hx_concat_result_151 := hx_concat_target_149.Copy()
		for _, hx_concat_item_152 := range hx_concat_other_150.Values() {
			hx_concat_result_151.Push(hx_concat_item_152)
		}
		return hx_concat_result_151
	}())
}

func none() any {
	return nil
}

func row(label *string, values *hxrt.Array) {
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(label, hxrt.StringFromLiteral("=")), hxrt.StringJoinAny(values.Values(), hxrt.StringFromLiteral(":"))))
	hxrt.Println(v)
}

func some(value float64) any {
	return value
}

func tick(label *string, value float64) float64 {
	hx_arr_153 := trail
	hx_arr_153.Push(label)
	return value
}

var trail *hxrt.Array = hxrt.NewArray()

type I_Meter interface {
}

type Meter struct {
	__hx_this I_Meter
	reading   any
}

func (self *Meter) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Meter() *Meter {
	return hx_init_Meter(hx_alloc_Meter(nil))
}

func hx_alloc_Meter(owner any) *Meter {
	self := &Meter{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Meter)
	}
	return self
}

func hx_init_Meter(self *Meter) *Meter {
	self.reading = nil
	return self
}

func (self *Meter) HxClassName() string {
	return "Meter"
}

var Meter_shared any = nil

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
