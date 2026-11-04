package main

import "snapshot/hxrt"

func invoke(callback *hxrt.Function[func(any) int]) int {
	return func() int {
		var hx_callable_1 *hxrt.Function[func(any) int] = callback
		var hx_argument_2 any = nil
		return hxrt.FunctionCall(hx_callable_1)(hx_argument_2)
	}()
}

func main() {
	func() {
		var hx_callable_7 *hxrt.Function[func(*hxrt.Function[func()], any)] = schedule
		var hx_argument_8 *hxrt.Function[func()] = hxrt.NewFunction(func() {
			func() {
				var hx_callable_4 *hxrt.Function[func(any, hxrt.Anonymous)] = haxe__Log_trace
				var hx_argument_5 any = hxrt.StringFromLiteral("scheduled")
				var hx_argument_6 hxrt.Anonymous = func() hxrt.Anonymous {
					hx_obj_3 := map[string]any{}
					hx_obj_3["fileName"] = hxrt.StringFromLiteral("Main.hx")
					hx_obj_3["lineNumber"] = 26
					hx_obj_3["className"] = hxrt.StringFromLiteral("Main")
					hx_obj_3["methodName"] = hxrt.StringFromLiteral("main")
					return hxrt.AnonymousFromMap(hx_obj_3)
				}()
				hxrt.FunctionCall(hx_callable_4)(hx_argument_5, hx_argument_6)
			}()
		})
		var hx_argument_9 any = nil
		hxrt.FunctionCall(hx_callable_7)(hx_argument_8, hx_argument_9)
	}()
	func() {
		var hx_callable_10 *hxrt.Function[func(*hxrt.Function[func()], any)] = schedule
		var hx_argument_11 *hxrt.Function[func()] = hxrt.NewFunction(func() {
		})
		var hx_argument_12 any = 10
		hxrt.FunctionCall(hx_callable_10)(hx_argument_11, hx_argument_12)
	}()
	func() {
		var hx_callable_18 *hxrt.Function[func(any, hxrt.Anonymous)] = haxe__Log_trace
		var hx_argument_19 any = func() float64 {
			var hx_callable_13 *hxrt.Function[func(any, any, any) float64] = summarize
			var hx_argument_14 any = nil
			var hx_argument_15 any = nil
			var hx_argument_16 any = nil
			return hxrt.FunctionCall(hx_callable_13)(hx_argument_14, hx_argument_15, hx_argument_16)
		}()
		var hx_argument_20 hxrt.Anonymous = func() hxrt.Anonymous {
			hx_obj_17 := map[string]any{}
			hx_obj_17["fileName"] = hxrt.StringFromLiteral("Main.hx")
			hx_obj_17["lineNumber"] = 28
			hx_obj_17["className"] = hxrt.StringFromLiteral("Main")
			hx_obj_17["methodName"] = hxrt.StringFromLiteral("main")
			return hxrt.AnonymousFromMap(hx_obj_17)
		}()
		hxrt.FunctionCall(hx_callable_18)(hx_argument_19, hx_argument_20)
	}()
	func() {
		var hx_callable_26 *hxrt.Function[func(any, hxrt.Anonymous)] = haxe__Log_trace
		var hx_argument_27 any = func() float64 {
			var hx_callable_21 *hxrt.Function[func(any, any, any) float64] = summarize
			var hx_argument_22 any = false
			var hx_argument_23 any = 9
			var hx_argument_24 any = 2.5
			return hxrt.FunctionCall(hx_callable_21)(hx_argument_22, hx_argument_23, hx_argument_24)
		}()
		var hx_argument_28 hxrt.Anonymous = func() hxrt.Anonymous {
			hx_obj_25 := map[string]any{}
			hx_obj_25["fileName"] = hxrt.StringFromLiteral("Main.hx")
			hx_obj_25["lineNumber"] = 29
			hx_obj_25["className"] = hxrt.StringFromLiteral("Main")
			hx_obj_25["methodName"] = hxrt.StringFromLiteral("main")
			return hxrt.AnonymousFromMap(hx_obj_25)
		}()
		hxrt.FunctionCall(hx_callable_26)(hx_argument_27, hx_argument_28)
	}()
	summarize = hxrt.NewFunction(func(enabled any, count any, ratio any) float64 {
		if enabled == nil {
			enabled = false
		}
		if count == nil {
			count = 4
		}
		if ratio == nil {
			ratio = 2.0
		}
		var hx_if_29 float64
		if enabled.(bool) {
			hx_if_29 = (float64(count.(int)) + ratio.(float64))
		} else {
			hx_if_29 = ratio.(float64)
		}
		return hx_if_29
	})
	func() {
		var hx_callable_35 *hxrt.Function[func(any, hxrt.Anonymous)] = haxe__Log_trace
		var hx_argument_36 any = func() float64 {
			var hx_callable_30 *hxrt.Function[func(any, any, any) float64] = summarize
			var hx_argument_31 any = false
			var hx_argument_32 any = 8
			var hx_argument_33 any = 3.5
			return hxrt.FunctionCall(hx_callable_30)(hx_argument_31, hx_argument_32, hx_argument_33)
		}()
		var hx_argument_37 hxrt.Anonymous = func() hxrt.Anonymous {
			hx_obj_34 := map[string]any{}
			hx_obj_34["fileName"] = hxrt.StringFromLiteral("Main.hx")
			hx_obj_34["lineNumber"] = 34
			hx_obj_34["className"] = hxrt.StringFromLiteral("Main")
			hx_obj_34["methodName"] = hxrt.StringFromLiteral("main")
			return hxrt.AnonymousFromMap(hx_obj_34)
		}()
		hxrt.FunctionCall(hx_callable_35)(hx_argument_36, hx_argument_37)
	}()
	var v any = any(func() float64 {
		var hx_callable_38 *hxrt.Function[func(any, any, any) float64] = summarize
		var hx_argument_39 any = nil
		var hx_argument_40 any = nil
		var hx_argument_41 any = nil
		return hxrt.FunctionCall(hx_callable_38)(hx_argument_39, hx_argument_40, hx_argument_41)
	}())
	hxrt.Println(v)
	var local *hxrt.Function[func(any, any) int] = nil
	local = hxrt.NewFunction(func(enabled any, count any) int {
		if enabled == nil {
			enabled = true
		}
		if count == nil {
			count = 6
		}
		var hx_if_42 int
		if enabled.(bool) {
			hx_if_42 = count.(int)
		} else {
			hx_if_42 = int(-int32(count.(int)))
		}
		return hx_if_42
	})
	func() {
		var hx_callable_47 *hxrt.Function[func(any, hxrt.Anonymous)] = haxe__Log_trace
		var hx_argument_48 any = func() int {
			var hx_callable_43 *hxrt.Function[func(any, any) int] = local
			var hx_argument_44 any = false
			var hx_argument_45 any = 11
			return hxrt.FunctionCall(hx_callable_43)(hx_argument_44, hx_argument_45)
		}()
		var hx_argument_49 hxrt.Anonymous = func() hxrt.Anonymous {
			hx_obj_46 := map[string]any{}
			hx_obj_46["fileName"] = hxrt.StringFromLiteral("Main.hx")
			hx_obj_46["lineNumber"] = 41
			hx_obj_46["className"] = hxrt.StringFromLiteral("Main")
			hx_obj_46["methodName"] = hxrt.StringFromLiteral("main")
			return hxrt.AnonymousFromMap(hx_obj_46)
		}()
		hxrt.FunctionCall(hx_callable_47)(hx_argument_48, hx_argument_49)
	}()
	var initialized *hxrt.Function[func(any, any) int] = hxrt.NewFunction(func(enabled any, count any) int {
		if enabled == nil {
			enabled = true
		}
		if count == nil {
			count = 6
		}
		var hx_if_50 int
		if enabled.(bool) {
			hx_if_50 = count.(int)
		} else {
			hx_if_50 = int(-int32(count.(int)))
		}
		return hx_if_50
	})
	var v_1 any = any(func() int {
		var hx_callable_51 *hxrt.Function[func(any, any) int] = initialized
		var hx_argument_52 any = nil
		var hx_argument_53 any = nil
		return hxrt.FunctionCall(hx_callable_51)(hx_argument_52, hx_argument_53)
	}())
	hxrt.Println(v_1)
	initialized = hxrt.NewFunction(func(enabled any, count any) int {
		if enabled == nil {
			enabled = false
		}
		if count == nil {
			count = 4
		}
		var hx_if_54 int
		if enabled.(bool) {
			hx_if_54 = count.(int)
		} else {
			hx_if_54 = int(-int32(count.(int)))
		}
		return hx_if_54
	})
	var v_2 any = any(func() int {
		var hx_callable_55 *hxrt.Function[func(any, any) int] = initialized
		var hx_argument_56 any = nil
		var hx_argument_57 any = nil
		return hxrt.FunctionCall(hx_callable_55)(hx_argument_56, hx_argument_57)
	}())
	hxrt.Println(v_2)
	var casted *hxrt.Function[func(any) int] = hxrt.NewFunction(func(count any) int {
		if count == nil {
			count = 2
		}
		return count.(int)
	})
	var v_3 any = any(func() int {
		var hx_callable_58 *hxrt.Function[func(any) int] = casted
		var hx_argument_59 any = 2
		return hxrt.FunctionCall(hx_callable_58)(hx_argument_59)
	}())
	hxrt.Println(v_3)
	hx_callable_60 := hxrt.NamedFunction(namedDefault, "Main.namedDefault")
	hx_normalized_callable_61 := hxrt.AdaptFunction(hx_callable_60, func(count any) int {
		if count == nil {
			count = 8
		}
		return hxrt.FunctionCall(hx_callable_60)(hxrt.IntFromNullableAny(count))
	})
	var namedCarrier *hxrt.Function[func(any) int] = hx_normalized_callable_61
	var v_4 any = any(func() int {
		var hx_callable_62 *hxrt.Function[func(any) int] = namedCarrier
		var hx_argument_63 any = nil
		return hxrt.FunctionCall(hx_callable_62)(hx_argument_63)
	}())
	hxrt.Println(v_4)
	var localSource *hxrt.Function[func(any) int] = hxrt.NewFunction(func(count any) int {
		if count == nil {
			count = 5
		}
		return count.(int)
	})
	hx_callable_64 := localSource
	hx_normalized_callable_65 := hxrt.AdaptFunction(hx_callable_64, func(count any) int {
		if count == nil {
			count = 5
		}
		return hxrt.FunctionCall(hx_callable_64)(hxrt.IntFromNullableAny(count))
	})
	var localAlias *hxrt.Function[func(any) int] = hx_normalized_callable_65
	var v_5 any = any(func() int {
		var hx_callable_66 *hxrt.Function[func(any) int] = localAlias
		var hx_argument_67 any = nil
		return hxrt.FunctionCall(hx_callable_66)(hx_argument_67)
	}())
	hxrt.Println(v_5)
	var v_6 any = any(invoke(func() *hxrt.Function[func(any) int] {
		hx_callable_68 := localSource
		hx_normalized_callable_69 := hxrt.AdaptFunction(hx_callable_68, func(count any) int {
			if count == nil {
				count = 5
			}
			return hxrt.FunctionCall(hx_callable_68)(hxrt.IntFromNullableAny(count))
		})
		return hx_normalized_callable_69
	}()))
	hxrt.Println(v_6)
	var v_7 any = any(invoke(hxrt.NewFunction(func(count any) int {
		if count == nil {
			count = 12
		}
		return count.(int)
	})))
	hxrt.Println(v_7)
	var v_8 any = any(func() int {
		var hx_callable_70 *hxrt.Function[func(any) int] = makeLocalCarrier()
		var hx_argument_71 any = nil
		return hxrt.FunctionCall(hx_callable_70)(hx_argument_71)
	}())
	hxrt.Println(v_8)
	var localOther *hxrt.Function[func(any) int] = hxrt.NewFunction(func(count any) int {
		if count == nil {
			count = 6
		}
		return count.(int)
	})
	var hx_if_78 *hxrt.Function[func(any) int]
	if func() int {
		var hx_callable_72 *hxrt.Function[func(any) int] = localAlias
		var hx_argument_73 any = nil
		return hxrt.FunctionCall(hx_callable_72)(hx_argument_73)
	}() == 5 {
		hx_callable_74 := localSource
		hx_normalized_callable_75 := hxrt.AdaptFunction(hx_callable_74, func(count any) int {
			if count == nil {
				count = 5
			}
			return hxrt.FunctionCall(hx_callable_74)(hxrt.IntFromNullableAny(count))
		})
		hx_if_78 = hx_normalized_callable_75
	} else {
		hx_callable_76 := localOther
		hx_normalized_callable_77 := hxrt.AdaptFunction(hx_callable_76, func(count any) int {
			if count == nil {
				count = 6
			}
			return hxrt.FunctionCall(hx_callable_76)(hxrt.IntFromNullableAny(count))
		})
		hx_if_78 = hx_normalized_callable_77
	}
	var conditional *hxrt.Function[func(any) int] = hx_if_78
	var v_9 any = any(func() int {
		var hx_callable_79 *hxrt.Function[func(any) int] = conditional
		var hx_argument_80 any = nil
		return hxrt.FunctionCall(hx_callable_79)(hx_argument_80)
	}())
	hxrt.Println(v_9)
	_g := func() int {
		var hx_callable_81 *hxrt.Function[func(any) int] = localAlias
		var hx_argument_82 any = nil
		return hxrt.FunctionCall(hx_callable_81)(hx_argument_82)
	}()
	var hx_if_87 *hxrt.Function[func(any) int]
	if _g == 5 {
		hx_callable_83 := localOther
		hx_normalized_callable_84 := hxrt.AdaptFunction(hx_callable_83, func(count any) int {
			if count == nil {
				count = 6
			}
			return hxrt.FunctionCall(hx_callable_83)(hxrt.IntFromNullableAny(count))
		})
		hx_if_87 = hx_normalized_callable_84
	} else {
		hx_callable_85 := localSource
		hx_normalized_callable_86 := hxrt.AdaptFunction(hx_callable_85, func(count any) int {
			if count == nil {
				count = 5
			}
			return hxrt.FunctionCall(hx_callable_85)(hxrt.IntFromNullableAny(count))
		})
		hx_if_87 = hx_normalized_callable_86
	}
	var switched *hxrt.Function[func(any) int] = hx_if_87
	var v_10 any = any(func() int {
		var hx_callable_88 *hxrt.Function[func(any) int] = switched
		var hx_argument_89 any = nil
		return hxrt.FunctionCall(hx_callable_88)(hx_argument_89)
	}())
	hxrt.Println(v_10)
	observed := func() int {
		var hx_callable_90 *hxrt.Function[func(any) int] = localAlias
		var hx_argument_91 any = nil
		return hxrt.FunctionCall(hx_callable_90)(hx_argument_91)
	}()
	if observed != 5 {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("unexpected")))
	}
	hx_callable_92 := localSource
	hx_normalized_callable_93 := hxrt.AdaptFunction(hx_callable_92, func(count any) int {
		if count == nil {
			count = 5
		}
		return hxrt.FunctionCall(hx_callable_92)(hxrt.IntFromNullableAny(count))
	})
	var blocked *hxrt.Function[func(any) int] = hx_normalized_callable_93
	var v_11 any = any(func() int {
		var hx_callable_94 *hxrt.Function[func(any) int] = blocked
		var hx_argument_95 any = nil
		return hxrt.FunctionCall(hx_callable_94)(hx_argument_95)
	}())
	hxrt.Println(v_11)
	var hx_try_96 *hxrt.Function[func(any) int]
	hxrt.TryCatch(func() {
		hx_callable_97 := localOther
		hx_normalized_callable_98 := hxrt.AdaptFunction(hx_callable_97, func(count any) int {
			if count == nil {
				count = 6
			}
			return hxrt.FunctionCall(hx_callable_97)(hxrt.IntFromNullableAny(count))
		})
		hx_try_96 = hx_normalized_callable_98
	}, func(hx_caught_99 any) {
		hx_exception_102 := haxe__Exception_caught(hx_caught_99)
		hx_payload_103 := haxe__Exception_catchValue(hx_caught_99)
		_ = hx_exception_102
		var hx_tmp any = hx_payload_103
		_ = hx_tmp
		hx_callable_100 := localSource
		hx_normalized_callable_101 := hxrt.AdaptFunction(hx_callable_100, func(count any) int {
			if count == nil {
				count = 5
			}
			return hxrt.FunctionCall(hx_callable_100)(hxrt.IntFromNullableAny(count))
		})
		hx_try_96 = hx_normalized_callable_101
	})
	var caught *hxrt.Function[func(any) int] = hx_try_96
	var v_12 any = any(func() int {
		var hx_callable_104 *hxrt.Function[func(any) int] = caught
		var hx_argument_105 any = nil
		return hxrt.FunctionCall(hx_callable_104)(hx_argument_105)
	}())
	hxrt.Println(v_12)
	var callbacks_0 *hxrt.Function[func(any) int]
	hx_callable_106 := localSource
	hx_normalized_callable_107 := hxrt.AdaptFunction(hx_callable_106, func(count any) int {
		if count == nil {
			count = 5
		}
		return hxrt.FunctionCall(hx_callable_106)(hxrt.IntFromNullableAny(count))
	})
	callbacks_0 = hx_normalized_callable_107
	var v_13 any = any(func() int {
		var hx_callable_108 *hxrt.Function[func(any) int] = callbacks_0
		var hx_argument_109 any = nil
		return hxrt.FunctionCall(hx_callable_108)(hx_argument_109)
	}())
	hxrt.Println(v_13)
	hx_callable_110 := localOther
	hx_normalized_callable_111 := hxrt.AdaptFunction(hx_callable_110, func(count any) int {
		if count == nil {
			count = 6
		}
		return hxrt.FunctionCall(hx_callable_110)(hxrt.IntFromNullableAny(count))
	})
	callbacks_0 = hx_normalized_callable_111
	var v_14 any = any(func() int {
		var hx_callable_112 *hxrt.Function[func(any) int] = callbacks_0
		var hx_argument_113 any = nil
		return hxrt.FunctionCall(hx_callable_112)(hx_argument_113)
	}())
	hxrt.Println(v_14)
	var holder_callback *hxrt.Function[func(any) int]
	hx_callable_114 := localSource
	hx_normalized_callable_115 := hxrt.AdaptFunction(hx_callable_114, func(count any) int {
		if count == nil {
			count = 5
		}
		return hxrt.FunctionCall(hx_callable_114)(hxrt.IntFromNullableAny(count))
	})
	holder_callback = hx_normalized_callable_115
	var v_15 any = any(func() int {
		var hx_callable_116 *hxrt.Function[func(any) int] = holder_callback
		var hx_argument_117 any = nil
		return hxrt.FunctionCall(hx_callable_116)(hx_argument_117)
	}())
	hxrt.Println(v_15)
}

func makeLocalCarrier() *hxrt.Function[func(any) int] {
	var source *hxrt.Function[func(any) int] = hxrt.NewFunction(func(count any) int {
		if count == nil {
			count = 10
		}
		return count.(int)
	})
	hx_callable_118 := source
	hx_normalized_callable_119 := hxrt.AdaptFunction(hx_callable_118, func(count any) int {
		if count == nil {
			count = 10
		}
		return hxrt.FunctionCall(hx_callable_118)(hxrt.IntFromNullableAny(count))
	})
	return hx_normalized_callable_119
}

func namedDefault(count int) int {
	return count
}

var schedule *hxrt.Function[func(*hxrt.Function[func()], any)] = hxrt.NewFunction(func(callback *hxrt.Function[func()], timeout any) {
	func() {
		var hx_callable_120 *hxrt.Function[func()] = callback
		hxrt.FunctionCall(hx_callable_120)()
	}()
})

var summarize *hxrt.Function[func(any, any, any) float64] = hxrt.NewFunction(func(enabled any, count any, ratio any) float64 {
	if enabled == nil {
		enabled = true
	}
	if count == nil {
		count = 3
	}
	if ratio == nil {
		ratio = 1.5
	}
	var hx_if_121 float64
	if enabled.(bool) {
		hx_if_121 = (float64(count.(int)) + ratio.(float64))
	} else {
		hx_if_121 = ratio.(float64)
	}
	return hx_if_121
})

func hx_recover_class_0(value any) *haxe__ValueException {
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

func hx_recover_class_1(value any) *haxe__Exception {
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

func hx_is_nominal_0(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_1(value any) bool {
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
