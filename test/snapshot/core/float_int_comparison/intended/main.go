package main

import "snapshot/hxrt"

var HALF float64 = 0.5

var MAX_INT int = 2147483647

func all(f float64, i int) *string {
	return hxrt.StringJoinAny(func() *hxrt.Array {
		f1 := hxrt.NamedFunction(flag, "Main.flag")
		_g := hxrt.NewArray()
		_g1 := 0
		_g2 := hxrt.NewArray((f < float64(i)), (f <= float64(i)), (f > float64(i)), (f >= float64(i)), (f == float64(i)), (f != float64(i)), (float64(i) < f), (float64(i) <= f), (float64(i) > f), (float64(i) >= f), (float64(i) == f), (float64(i) != f))
		for _g1 < _g2.Len() {
			v := func(hx_value_1 any) bool {
				if hx_value_1 == nil {
					var hx_zero_2 bool
					return hx_zero_2
				}
				return hx_value_1.(bool)
			}(_g2.Get(_g1))
			_g1 = int(int32((_g1 + 1)))
			_g.Push(func() *string {
				var hx_callable_4 *hxrt.Function[func(bool) *string] = f1
				var hx_argument_5 bool = v
				return hxrt.FunctionCall(hx_callable_4)(hx_argument_5)
			}())
		}
		return _g
	}().Values(), hxrt.StringFromLiteral(""))
}

func flag(value bool) *string {
	var hx_if_6 *string
	if value {
		hx_if_6 = hxrt.StringFromLiteral("t")
	} else {
		hx_if_6 = hxrt.StringFromLiteral("f")
	}
	return hx_if_6
}

func main() {
	minInt := -2147483648
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("fraction="), all(2.5, 2)))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("equal="), all(2.0, 2)))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("below.zero="), all(-0.5, 0)))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("negative="), all(-3.0, -3)))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("negative.fraction="), all(-2.5, -3)))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("max="), all(2147483647.0, 2147483647)))
	hxrt.Println(v_5)
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("max.fraction="), all(2147483646.5, 2147483647)))
	hxrt.Println(v_6)
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("min="), all(-2147483648.0, minInt)))
	hxrt.Println(v_7)
	var v_8 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("float32.gap="), all(16777216.0, 16777217)))
	hxrt.Println(v_8)
	var v_9 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("nan="), all(Math_NaN, 0)))
	hxrt.Println(v_9)
	var v_10 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("inf="), all(Math_POSITIVE_INFINITY, 2147483647)))
	hxrt.Println(v_10)
	var v_11 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("neg.inf="), all(Math_NEGATIVE_INFINITY, minInt)))
	hxrt.Println(v_11)
	top := 2147483647
	topReal := 2147483647.0
	row(hxrt.StringFromLiteral("wrapped"), hxrt.NewArray((topReal < float64(int((hxrt.Int32Wrap(top)+hxrt.Int32Wrap(1))))), (float64(int((hxrt.Int32Wrap(top)+hxrt.Int32Wrap(1)))) < topReal), (topReal == float64(int((hxrt.Int32Wrap(top)+hxrt.Int32Wrap(1))))), (float64(int((hxrt.Int32Wrap(top)*hxrt.Int32Wrap(2)))) == -2.0)))
	i := 2
	f := 2.5
	row(hxrt.StringFromLiteral("constant"), hxrt.NewArray((float64(i) < 2.5), (2.5 > float64(i)), (float64(i) <= 2.0), (float64(i) == 2.0), (2.0 == float64(i)), (float64(i) != 2.5), (f < 3), (3 > f), (f == 2), (2 != f), (f >= -1), (-1 <= f), (float64(i) < 0.5), (0.5 < float64(i)), true, (float64(i) < Math_POSITIVE_INFINITY), (Math_NaN != float64(i)), (Math_NEGATIVE_INFINITY < float64(i))))
	meter := New_Meter(2, 2.5)
	var sample_whole int
	var sample_real float64
	var sample_maybeWhole any
	var sample_maybeReal any
	sample_whole = 3
	sample_real = 3.0
	sample_maybeWhole = nil
	sample_maybeReal = nil
	row(hxrt.StringFromLiteral("field"), hxrt.NewArray((float64(meter.count) < meter.ratio), (meter.ratio >= float64(meter.count)), (float64(meter.count) == meter.ratio), (float64(Meter_limit) > Meter_scale), (Meter_scale == float64(Meter_limit)), (Meter_scale <= float64(Meter_limit)), (float64(sample_whole) <= sample_real), (sample_real != float64(sample_whole)), (float64(sample_whole) == sample_real), (float64(meter.count) < Meter_scale), (float64(Meter_limit) == sample_real), (float64(sample_whole) > meter.ratio)))
	row(hxrt.StringFromLiteral("call"), hxrt.NewArray((float64(whole(2)) < real(2.5)), (real(2.0) == float64(whole(2))), (float64(whole(3)) > 2.5), (2.5 < float64(whole(3))), (real(-1.5) > float64(whole(-2))), (float64(whole(2147483647)) != real(2147483647.0)), (real(f) >= float64(whole(i)))))
	ints := hxrt.NewArray(1, 2, 3)
	reals := hxrt.NewArray(1.5, 2.0, 2.5)
	row(hxrt.StringFromLiteral("array"), hxrt.NewArray((float64(hxrt.IntFromNullableAny(ints.Get(1))) < func(hx_value_7 any) float64 {
		if hx_value_7 == nil {
			var hx_zero_8 float64
			return hx_zero_8
		}
		return hx_value_7.(float64)
	}(reals.Get(2))), hxrt.HaxeEqual(reals.Get(1), ints.Get(1)), (float64(hxrt.IntFromNullableAny(ints.Get(2))) >= func(hx_value_9 any) float64 {
		if hx_value_9 == nil {
			var hx_zero_10 float64
			return hx_zero_10
		}
		return hx_value_9.(float64)
	}(reals.Get(2))), (func(hx_value_11 any) float64 {
		if hx_value_11 == nil {
			var hx_zero_12 float64
			return hx_zero_12
		}
		return hx_value_11.(float64)
	}(reals.Get(0)) > float64(hxrt.IntFromNullableAny(ints.Get(0)))), !hxrt.HaxeEqual(ints.Get(0), reals.Get(0)), (float64(hxrt.IntFromNullableAny(ints.Get(2))) > f), (f < float64(hxrt.IntFromNullableAny(ints.Get(2)))), hxrt.HaxeEqual(reals.Get(1), i), (float64(i) < func(hx_value_13 any) float64 {
		if hx_value_13 == nil {
			var hx_zero_14 float64
			return hx_zero_14
		}
		return hx_value_13.(float64)
	}(reals.Get(2))), hxrt.HaxeEqual(ints.Get(i), 3.0), (func(hx_value_15 any) float64 {
		if hx_value_15 == nil {
			var hx_zero_16 float64
			return hx_zero_16
		}
		return hx_value_15.(float64)
	}(reals.Get(int((hxrt.Int32Wrap(i)-hxrt.Int32Wrap(1))))) <= float64(hxrt.IntFromNullableAny(ints.Get(1))))))
	var noInt any = noWhole()
	var two any = someWhole(2)
	var noFloat any = noReal()
	var twoReal any = someReal(2.0)
	var half any = someReal(2.5)
	row(hxrt.StringFromLiteral("null.ordered"), hxrt.NewArray((func(hx_value_17 any) float64 {
		if hx_value_17 == nil {
			var hx_zero_18 float64
			return (0 / hx_zero_18)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_17))
	}(noInt) < f), (f > func(hx_value_19 any) float64 {
		if hx_value_19 == nil {
			var hx_zero_20 float64
			return (0 / hx_zero_20)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_19))
	}(noInt)), (func(hx_value_21 any) float64 {
		if hx_value_21 == nil {
			var hx_zero_22 float64
			return (0 / hx_zero_22)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_21))
	}(noInt) <= 2.5), (2.5 >= func(hx_value_23 any) float64 {
		if hx_value_23 == nil {
			var hx_zero_24 float64
			return (0 / hx_zero_24)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_23))
	}(noInt)), (func(hx_value_25 any) float64 {
		if hx_value_25 == nil {
			var hx_zero_26 float64
			return (0 / hx_zero_26)
		}
		return hx_value_25.(float64)
	}(noFloat) < float64(i)), (float64(i) > func(hx_value_27 any) float64 {
		if hx_value_27 == nil {
			var hx_zero_28 float64
			return (0 / hx_zero_28)
		}
		return hx_value_27.(float64)
	}(noFloat)), (func(hx_value_29 any) float64 {
		if hx_value_29 == nil {
			var hx_zero_30 float64
			return (0 / hx_zero_30)
		}
		return hx_value_29.(float64)
	}(noFloat) < func(hx_value_31 any) float64 {
		if hx_value_31 == nil {
			var hx_zero_32 float64
			return (0 / hx_zero_32)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_31))
	}(noInt)), (func(hx_value_33 any) float64 {
		if hx_value_33 == nil {
			var hx_zero_34 float64
			return (0 / hx_zero_34)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_33))
	}(noInt) > func(hx_value_35 any) float64 {
		if hx_value_35 == nil {
			var hx_zero_36 float64
			return (0 / hx_zero_36)
		}
		return hx_value_35.(float64)
	}(noFloat)), (func(hx_value_37 any) float64 {
		if hx_value_37 == nil {
			var hx_zero_38 float64
			return (0 / hx_zero_38)
		}
		return hx_value_37.(float64)
	}(noFloat) <= 2), !(func(hx_value_39 any) float64 {
		if hx_value_39 == nil {
			var hx_zero_40 float64
			return (0 / hx_zero_40)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_39))
	}(noInt) < f)))
	row(hxrt.StringFromLiteral("value.ordered"), hxrt.NewArray((func(hx_value_41 any) float64 {
		if hx_value_41 == nil {
			var hx_zero_42 float64
			return (0 / hx_zero_42)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_41))
	}(two) < f), (f > func(hx_value_43 any) float64 {
		if hx_value_43 == nil {
			var hx_zero_44 float64
			return (0 / hx_zero_44)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_43))
	}(two)), (func(hx_value_45 any) float64 {
		if hx_value_45 == nil {
			var hx_zero_46 float64
			return (0 / hx_zero_46)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_45))
	}(two) >= func(hx_value_47 any) float64 {
		if hx_value_47 == nil {
			var hx_zero_48 float64
			return (0 / hx_zero_48)
		}
		return hx_value_47.(float64)
	}(twoReal)), (func(hx_value_49 any) float64 {
		if hx_value_49 == nil {
			var hx_zero_50 float64
			return (0 / hx_zero_50)
		}
		return hx_value_49.(float64)
	}(half) > func(hx_value_51 any) float64 {
		if hx_value_51 == nil {
			var hx_zero_52 float64
			return (0 / hx_zero_52)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_51))
	}(two)), (func(hx_value_53 any) float64 {
		if hx_value_53 == nil {
			var hx_zero_54 float64
			return (0 / hx_zero_54)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_53))
	}(two) <= 2.5), (func(hx_value_55 any) float64 {
		if hx_value_55 == nil {
			var hx_zero_56 float64
			return (0 / hx_zero_56)
		}
		return hx_value_55.(float64)
	}(twoReal) <= float64(i)), (float64(i) < func(hx_value_57 any) float64 {
		if hx_value_57 == nil {
			var hx_zero_58 float64
			return (0 / hx_zero_58)
		}
		return hx_value_57.(float64)
	}(half)), (func(hx_value_59 any) float64 {
		if hx_value_59 == nil {
			var hx_zero_60 float64
			return (0 / hx_zero_60)
		}
		return hx_value_59.(float64)
	}(half) >= 2), (func(hx_value_61 any) float64 {
		if hx_value_61 == nil {
			var hx_zero_62 float64
			return (0 / hx_zero_62)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_61))
	}(two) < func(hx_value_63 any) float64 {
		if hx_value_63 == nil {
			var hx_zero_64 float64
			return (0 / hx_zero_64)
		}
		return hx_value_63.(float64)
	}(half)), (func(hx_value_65 any) float64 {
		if hx_value_65 == nil {
			var hx_zero_66 float64
			return (0 / hx_zero_66)
		}
		return hx_value_65.(float64)
	}(half) < func(hx_value_67 any) float64 {
		if hx_value_67 == nil {
			var hx_zero_68 float64
			return (0 / hx_zero_68)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_67))
	}(two))))
	row(hxrt.StringFromLiteral("null.equal"), hxrt.NewArray((noFloat == func(hx_value_69 any) any {
		if hx_value_69 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_69))
	}(noInt)), (func(hx_value_70 any) any {
		if hx_value_70 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_70))
	}(noInt) == noFloat), (noFloat != func(hx_value_71 any) any {
		if hx_value_71 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_71))
	}(noInt)), (noFloat == func(hx_value_72 any) any {
		if hx_value_72 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_72))
	}(two)), (func(hx_value_73 any) any {
		if hx_value_73 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_73))
	}(two) != noFloat), (func(hx_value_74 any) any {
		if hx_value_74 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_74))
	}(noInt) == f), (f != func(hx_value_75 any) any {
		if hx_value_75 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_75))
	}(noInt)), (noFloat == float64(i)), (float64(i) != noFloat), (noFloat == 2.0), (func(hx_value_76 any) any {
		if hx_value_76 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_76))
	}(noInt) == 2.0)))
	row(hxrt.StringFromLiteral("value.equal"), hxrt.NewArray((twoReal == func(hx_value_77 any) any {
		if hx_value_77 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_77))
	}(two)), (func(hx_value_78 any) any {
		if hx_value_78 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_78))
	}(two) == twoReal), (half != func(hx_value_79 any) any {
		if hx_value_79 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_79))
	}(two)), (twoReal == float64(i)), (float64(i) == twoReal), (func(hx_value_80 any) any {
		if hx_value_80 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_80))
	}(two) == 2.0), (2.0 == func(hx_value_81 any) any {
		if hx_value_81 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_81))
	}(two)), (func(hx_value_82 any) any {
		if hx_value_82 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_82))
	}(two) != 2.5), (twoReal == 2.0), (2.0 == twoReal), (half != 2.0), (func(hx_value_83 any) any {
		if hx_value_83 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_83))
	}(two) == f)))
	var maybeInts_1 any
	var maybeInts_0 any
	maybeInts_0 = nil
	maybeInts_1 = 2
	var maybeReals_1 any
	var maybeReals_0 any
	maybeReals_0 = nil
	maybeReals_1 = 2.0
	row(hxrt.StringFromLiteral("null.field"), hxrt.NewArray((func(hx_value_84 any) float64 {
		if hx_value_84 == nil {
			var hx_zero_85 float64
			return (0 / hx_zero_85)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_84))
	}(meter.maybeCount) < meter.ratio), (meter.maybeRatio == float64(meter.count)), (meter.maybeRatio != float64(meter.count)), (func(hx_value_86 any) float64 {
		if hx_value_86 == nil {
			var hx_zero_87 float64
			return (0 / hx_zero_87)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_86))
	}(Meter_spare) <= Meter_scale), (func(hx_value_88 any) any {
		if hx_value_88 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_88))
	}(Meter_spare) == Meter_scale), (func(hx_value_89 any) float64 {
		if hx_value_89 == nil {
			var hx_zero_90 float64
			return (0 / hx_zero_90)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_89))
	}(sample_maybeWhole) < sample_real), (sample_maybeReal == float64(sample_whole)), (func(hx_value_91 any) any {
		if hx_value_91 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_91))
	}(sample_maybeWhole) == sample_maybeReal)))
	meter.maybeCount = 2
	meter.maybeRatio = 2.0
	Meter_spare = 3
	sample_maybeWhole = 3
	sample_maybeReal = 2.5
	row(hxrt.StringFromLiteral("value.field"), hxrt.NewArray((func(hx_value_92 any) float64 {
		if hx_value_92 == nil {
			var hx_zero_93 float64
			return (0 / hx_zero_93)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_92))
	}(meter.maybeCount) < meter.ratio), (meter.maybeRatio == float64(meter.count)), (func(hx_value_94 any) any {
		if hx_value_94 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_94))
	}(meter.maybeCount) == meter.maybeRatio), (func(hx_value_95 any) float64 {
		if hx_value_95 == nil {
			var hx_zero_96 float64
			return (0 / hx_zero_96)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_95))
	}(Meter_spare) > Meter_scale), (func(hx_value_97 any) any {
		if hx_value_97 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_97))
	}(Meter_spare) == Meter_scale), (func(hx_value_98 any) float64 {
		if hx_value_98 == nil {
			var hx_zero_99 float64
			return (0 / hx_zero_99)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_98))
	}(sample_maybeWhole) > sample_real), (sample_maybeReal != float64(sample_whole)), (func(hx_value_100 any) float64 {
		if hx_value_100 == nil {
			var hx_zero_101 float64
			return (0 / hx_zero_101)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_100))
	}(sample_maybeWhole) >= func(hx_value_102 any) float64 {
		if hx_value_102 == nil {
			var hx_zero_103 float64
			return (0 / hx_zero_103)
		}
		return hx_value_102.(float64)
	}(sample_maybeReal))))
	row(hxrt.StringFromLiteral("null.array"), hxrt.NewArray((func(hx_value_104 any) float64 {
		if hx_value_104 == nil {
			var hx_zero_105 float64
			return (0 / hx_zero_105)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_104))
	}(maybeInts_0) < f), (func(hx_value_106 any) any {
		if hx_value_106 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_106))
	}(maybeInts_0) == f), (func(hx_value_107 any) any {
		if hx_value_107 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_107))
	}(maybeInts_1) == 2.0), (func(hx_value_110 any) float64 {
		if hx_value_110 == nil {
			var hx_zero_111 float64
			return (0 / hx_zero_111)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_110))
	}(maybeInts_1) < func(hx_value_108 any) float64 {
		if hx_value_108 == nil {
			var hx_zero_109 float64
			return hx_zero_109
		}
		return hx_value_108.(float64)
	}(reals.Get(2))), (func(hx_value_112 any) float64 {
		if hx_value_112 == nil {
			var hx_zero_113 float64
			return (0 / hx_zero_113)
		}
		return hx_value_112.(float64)
	}(maybeReals_0) > float64(i)), (maybeReals_1 == float64(i)), (func(hx_value_114 any) float64 {
		if hx_value_114 == nil {
			var hx_zero_115 float64
			return (0 / hx_zero_115)
		}
		return hx_value_114.(float64)
	}(maybeReals_1) >= float64(hxrt.IntFromNullableAny(ints.Get(1)))), (func(hx_value_116 any) any {
		if hx_value_116 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_116))
	}(maybeInts_1) == maybeReals_1), (func(hx_value_117 any) any {
		if hx_value_117 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_117))
	}(maybeInts_0) == maybeReals_0)))
	row(hxrt.StringFromLiteral("null.call"), hxrt.NewArray((func(hx_value_118 any) float64 {
		if hx_value_118 == nil {
			var hx_zero_119 float64
			return (0 / hx_zero_119)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_118))
	}(noWhole()) < 1.5), (func(hx_value_120 any) any {
		if hx_value_120 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_120))
	}(someWhole(2)) == 2.0), (someReal(2.0) == 2.0), (func(hx_value_121 any) float64 {
		if hx_value_121 == nil {
			var hx_zero_122 float64
			return (0 / hx_zero_122)
		}
		return hx_value_121.(float64)
	}(noReal()) >= 0), (func(hx_value_123 any) float64 {
		if hx_value_123 == nil {
			var hx_zero_124 float64
			return (0 / hx_zero_124)
		}
		return hx_value_123.(float64)
	}(someReal(2.5)) > func(hx_value_125 any) float64 {
		if hx_value_125 == nil {
			var hx_zero_126 float64
			return (0 / hx_zero_126)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_125))
	}(someWhole(2))), (func(hx_value_127 any) any {
		if hx_value_127 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_127))
	}(noWhole()) == noReal()), (func(hx_value_128 any) any {
		if hx_value_128 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_128))
	}(someWhole(2)) != someReal(2.0)), (noReal() == float64(i))))
	var maybeTwo any = someWhole(2)
	var maybeHalf any = someReal(2.5)
	row(hxrt.StringFromLiteral("narrowed"), hxrt.NewArray(((maybeTwo != nil) && (func(hx_value_129 any) float64 {
		if hx_value_129 == nil {
			var hx_zero_130 float64
			return (0 / hx_zero_130)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_129))
	}(maybeTwo) < f)), ((maybeTwo != nil) && (func(hx_value_131 any) any {
		if hx_value_131 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_131))
	}(maybeTwo) == 2.0)), ((maybeHalf != nil) && (func(hx_value_132 any) float64 {
		if hx_value_132 == nil {
			var hx_zero_133 float64
			return (0 / hx_zero_133)
		}
		return hx_value_132.(float64)
	}(maybeHalf) > float64(i))), ((maybeHalf != nil) && (maybeHalf != float64(i))), (((maybeTwo != nil) && (maybeHalf != nil)) && (func(hx_value_134 any) float64 {
		if hx_value_134 == nil {
			var hx_zero_135 float64
			return (0 / hx_zero_135)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_134))
	}(maybeTwo) < func(hx_value_136 any) float64 {
		if hx_value_136 == nil {
			var hx_zero_137 float64
			return (0 / hx_zero_137)
		}
		return hx_value_136.(float64)
	}(maybeHalf)))))
	taken := hxrt.NewArray()
	if f > float64(i) {
		taken.Push(hxrt.StringFromLiteral("if"))
	}
	if float64(i) >= f {
		taken.Push(hxrt.StringFromLiteral("never"))
	} else {
		taken.Push(hxrt.StringFromLiteral("else"))
	}
	var hx_if_142 *string
	if float64(i) < f {
		hx_if_142 = hxrt.StringFromLiteral("ternary")
	} else {
		hx_if_142 = hxrt.StringFromLiteral("never")
	}
	taken.Push(hx_if_142)
	var hx_if_146 *string
	if func(hx_value_144 any) float64 {
		if hx_value_144 == nil {
			var hx_zero_145 float64
			return (0 / hx_zero_145)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_144))
	}(noInt) < f {
		hx_if_146 = hxrt.StringFromLiteral("never")
	} else {
		hx_if_146 = hxrt.StringFromLiteral("null.ternary")
	}
	taken.Push(hx_if_146)
	var hx_if_148 *string
	if (f < float64(i)) && (i < 3) {
		hx_if_148 = hxrt.StringFromLiteral("never")
	} else {
		hx_if_148 = hxrt.StringFromLiteral("and")
	}
	taken.Push(hx_if_148)
	var hx_if_152 *string
	if (f > float64(i)) || (func(hx_value_150 any) float64 {
		if hx_value_150 == nil {
			var hx_zero_151 float64
			return (0 / hx_zero_151)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_150))
	}(noInt) < f) {
		hx_if_152 = hxrt.StringFromLiteral("or")
	} else {
		hx_if_152 = hxrt.StringFromLiteral("never")
	}
	taken.Push(hx_if_152)
	steps := 0
	for float64(steps) < f {
		steps = int(int32((steps + 1)))
	}
	level := 10.0
	drops := 0
	for level > float64(i) {
		level = (level - 1.5)
		drops = int(int32((drops + 1)))
	}
	rounds := 0
	hx_do_first_153 := true
	for hx_do_first_153 || (float64(rounds) <= 2.5) {
		hx_do_first_153 = false
		rounds = int(int32((rounds + 1)))
	}
	var counter any = 0
	for func(hx_value_154 any) float64 {
		if hx_value_154 == nil {
			var hx_zero_155 float64
			return (0 / hx_zero_155)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_154))
	}(counter) < 1.5 {
		counter = int((hxrt.Int32Wrap(hxrt.IntFromNullableAny(hxrt.IntFromNullableAny(counter))) + hxrt.Int32Wrap(1)))
	}
	value := i
	var hx_if_157 *string
	if float64(value) > f {
		hx_if_157 = hxrt.StringFromLiteral("above")
	} else {
		value_1 := i
		var hx_if_156 *string
		if float64(value_1) == 2.0 {
			hx_if_156 = hxrt.StringFromLiteral("two")
		} else {
			hx_if_156 = hxrt.StringFromLiteral("other")
		}
		hx_if_157 = hx_if_156
	}
	guard := hx_if_157
	stored := (f >= float64(i))
	var v_12 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("flow="), hxrt.StringJoinAny(taken.Values(), hxrt.StringFromLiteral(","))), hxrt.StringFromLiteral(":")), steps), hxrt.StringFromLiteral(":")), drops), hxrt.StringFromLiteral(":")), rounds), hxrt.StringFromLiteral(":")), counter), hxrt.StringFromLiteral(":")), guard), hxrt.StringFromLiteral(":")), flag(stored)))
	hxrt.Println(v_12)
	order := hxrt.NewArray((tickReal(hxrt.StringFromLiteral("a"), 1.5) < float64(tickWhole(hxrt.StringFromLiteral("b"), 2))), (float64(tickWhole(hxrt.StringFromLiteral("c"), 2)) >= tickReal(hxrt.StringFromLiteral("d"), 2.0)), (float64(tickWhole(hxrt.StringFromLiteral("e"), 2)) == tickReal(hxrt.StringFromLiteral("f"), 2.0)), (tickReal(hxrt.StringFromLiteral("g"), 1.0) != float64(tickWhole(hxrt.StringFromLiteral("h"), 1))), (func(hx_value_158 any) float64 {
		if hx_value_158 == nil {
			var hx_zero_159 float64
			return (0 / hx_zero_159)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_158))
	}(tickMaybeWhole(hxrt.StringFromLiteral("i"), nil)) < tickReal(hxrt.StringFromLiteral("j"), 1.0)), (tickReal(hxrt.StringFromLiteral("k"), 2.0) == func(hx_value_160 any) any {
		if hx_value_160 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_160))
	}(tickMaybeWhole(hxrt.StringFromLiteral("l"), 2))))
	var v_13 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("order="), hxrt.StringJoinAny(func() *hxrt.Array {
		f_1 := hxrt.NamedFunction(flag, "Main.flag")
		_g := hxrt.NewArray()
		_g1 := 0
		_g2 := order
		for _g1 < _g2.Len() {
			v_14 := func(hx_value_161 any) bool {
				if hx_value_161 == nil {
					var hx_zero_162 bool
					return hx_zero_162
				}
				return hx_value_161.(bool)
			}(_g2.Get(_g1))
			_g1 = int(int32((_g1 + 1)))
			_g.Push(func() *string {
				var hx_callable_164 *hxrt.Function[func(bool) *string] = f_1
				var hx_argument_165 bool = v_14
				return hxrt.FunctionCall(hx_callable_164)(hx_argument_165)
			}())
		}
		return _g
	}().Values(), hxrt.StringFromLiteral(""))), hxrt.StringFromLiteral(":")), hxrt.StringJoinAny(trail.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_13)
	small := 3
	huge := -1
	narrow := 5
	row(hxrt.StringFromLiteral("uint"), hxrt.NewArray((func() float64 {
		int := small
		var hx_if_166 float64
		if int < 0 {
			hx_if_166 = (4294967296.0 + float64(int))
		} else {
			hx_if_166 = (float64(int) + 0.0)
		}
		return hx_if_166
	}() > f), (f < func() float64 {
		int_1 := small
		var hx_if_167 float64
		if int_1 < 0 {
			hx_if_167 = (4294967296.0 + float64(int_1))
		} else {
			hx_if_167 = (float64(int_1) + 0.0)
		}
		return hx_if_167
	}()), (func() float64 {
		int_2 := small
		var hx_if_168 float64
		if int_2 < 0 {
			hx_if_168 = (4294967296.0 + float64(int_2))
		} else {
			hx_if_168 = (float64(int_2) + 0.0)
		}
		return hx_if_168
	}() == 3.0), (func() float64 {
		int_3 := huge
		var hx_if_169 float64
		if int_3 < 0 {
			hx_if_169 = (4294967296.0 + float64(int_3))
		} else {
			hx_if_169 = (float64(int_3) + 0.0)
		}
		return hx_if_169
	}() > f), (func() float64 {
		int_4 := huge
		var hx_if_170 float64
		if int_4 < 0 {
			hx_if_170 = (4294967296.0 + float64(int_4))
		} else {
			hx_if_170 = (float64(int_4) + 0.0)
		}
		return hx_if_170
	}() > 4000000000.0), (func() float64 {
		int_5 := huge
		var hx_if_171 float64
		if int_5 < 0 {
			hx_if_171 = (4294967296.0 + float64(int_5))
		} else {
			hx_if_171 = (float64(int_5) + 0.0)
		}
		return hx_if_171
	}() == 4294967295.0)))
	row(hxrt.StringFromLiteral("int32"), hxrt.NewArray((float64(narrow) > f), (f < float64(narrow)), (float64(narrow) == 5.0), (float64(narrow) != f), (float64(narrow) <= 5.0)))
}

func noReal() any {
	return nil
}

func noWhole() any {
	return nil
}

func real(value float64) float64 {
	return value
}

func row(label *string, values *hxrt.Array) {
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(label, hxrt.StringFromLiteral("=")), hxrt.StringJoinAny(func() *hxrt.Array {
		f := hxrt.NamedFunction(flag, "Main.flag")
		_g := hxrt.NewArray()
		_g1 := 0
		_g2 := values
		for _g1 < _g2.Len() {
			v_1 := func(hx_value_172 any) bool {
				if hx_value_172 == nil {
					var hx_zero_173 bool
					return hx_zero_173
				}
				return hx_value_172.(bool)
			}(_g2.Get(_g1))
			_g1 = int(int32((_g1 + 1)))
			_g.Push(func() *string {
				var hx_callable_175 *hxrt.Function[func(bool) *string] = f
				var hx_argument_176 bool = v_1
				return hxrt.FunctionCall(hx_callable_175)(hx_argument_176)
			}())
		}
		return _g
	}().Values(), hxrt.StringFromLiteral(""))))
	hxrt.Println(v)
}

func someReal(value float64) any {
	return value
}

func someWhole(value int) any {
	return value
}

func tickMaybeWhole(label *string, value any) any {
	hx_arr_177 := trail
	hx_arr_177.Push(label)
	return value
}

func tickReal(label *string, value float64) float64 {
	hx_arr_178 := trail
	hx_arr_178.Push(label)
	return value
}

func tickWhole(label *string, value int) int {
	hx_arr_179 := trail
	hx_arr_179.Push(label)
	return value
}

var trail *hxrt.Array = hxrt.NewArray()

func whole(value int) int {
	return value
}

type I_Meter interface {
}

type Meter struct {
	__hx_this  I_Meter
	count      int
	ratio      float64
	maybeCount any
	maybeRatio any
}

func (self *Meter) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Meter(count int, ratio float64) *Meter {
	return hx_init_Meter(hx_alloc_Meter(nil), count, ratio)
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

func hx_init_Meter(self *Meter, count int, ratio float64) *Meter {
	self.maybeRatio = nil
	self.maybeCount = nil
	self.count = count
	self.ratio = ratio
	return self
}

func (self *Meter) HxClassName() string {
	return "Meter"
}

var Meter_limit int = 3

var Meter_scale float64 = 2.5

var Meter_spare any = nil
