package main

import "snapshot/hxrt"

func attempt(action *hxrt.Function[func()]) *string {
	hx_try_return_1 := false
	var hx_try_value_2 *string
	hxrt.TryCatch(func() {
		func() {
			var hx_callable_6 *hxrt.Function[func()] = action
			hxrt.FunctionCall(hx_callable_6)()
		}()
		hx_try_value_2 = hxrt.StringFromLiteral("ok")
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

func bump(start any) int {
	start = int(int32((hxrt.IntFromNullableAny(start) + 1)))
	return hxrt.IntFromNullableAny(start.(int))
}

func flip(value bool) *string {
	var hx_if_8 *string
	if value {
		hx_if_8 = hxrt.StringFromLiteral("yes")
	} else {
		hx_if_8 = hxrt.StringFromLiteral("no")
	}
	return hx_if_8
}

func keep(value any) any {
	return value
}

func main() {
	var count any = nil
	count = 5
	count = int(int32((hxrt.IntFromNullableAny(count) + 1)))
	count = int(int32((hxrt.IntFromNullableAny(count) + 1)))
	count = int(int32((hxrt.IntFromNullableAny(count) - 1)))
	count = int(int32((hxrt.IntFromNullableAny(count) - 1)))
	hx_post_9 := hxrt.IntFromNullableAny(count)
	count = int(int32((hx_post_9 + 1)))
	post := hx_post_9
	pre := func() int {
		hx_pre_10 := int(int32((hxrt.IntFromNullableAny(count) + 1)))
		count = hx_pre_10
		return hx_pre_10
	}()
	arg := note(hxrt.StringFromLiteral("arg"), func() int {
		hx_post_11 := hxrt.IntFromNullableAny(count)
		count = int(int32((hx_post_11 - 1)))
		return hx_post_11
	}())
	var kept any = keep(func() int {
		hx_pre_12 := int(int32((hxrt.IntFromNullableAny(count) - 1)))
		count = hx_pre_12
		return hx_pre_12
	}())
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("local.int="), count), hxrt.StringFromLiteral(":")), post), hxrt.StringFromLiteral(":")), pre), hxrt.StringFromLiteral(":")), arg), hxrt.StringFromLiteral(":")), kept)))
	var amount any = nil
	amount = 0.5
	amount = (func(hx_value_13 any) float64 {
		if hx_value_13 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_13.(float64)
	}(amount) + 1)
	amount = (func(hx_value_14 any) float64 {
		if hx_value_14 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_14.(float64)
	}(amount) - 1)
	hx_post_15 := func(hx_value_16 any) float64 {
		if hx_value_16 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_16.(float64)
	}(amount)
	amount = (hx_post_15 + 1)
	floatPost := hx_post_15
	floatPre := func() float64 {
		hx_pre_17 := (func(hx_value_18 any) float64 {
			if hx_value_18 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_18.(float64)
		}(amount) - 1)
		amount = hx_pre_17
		return hx_pre_17
	}()
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("local.float="), amount), hxrt.StringFromLiteral(":")), floatPost), hxrt.StringFromLiteral(":")), floatPre), hxrt.StringFromLiteral(":")), -func(hx_value_21 any) float64 {
		if hx_value_21 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_21.(float64)
	}(amount)), hxrt.StringFromLiteral(":")), pairFloat(func() float64 {
		hx_post_22 := func(hx_value_23 any) float64 {
			if hx_value_23 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_23.(float64)
		}(amount)
		amount = (hx_post_22 - 1)
		return hx_post_22
	}(), func(hx_value_24 any) float64 {
		if hx_value_24 == nil {
			var hx_zero_25 float64
			return hx_zero_25
		}
		return hx_value_24.(float64)
	}(amount))))
	hxrt.Println(v)
	meter := New_Meter()
	meter.hits = int(int32((hxrt.IntFromNullableAny(meter.hits) + 1)))
	meter.hits = int(int32((hxrt.IntFromNullableAny(meter.hits) - 1)))
	hx_post_26 := hxrt.IntFromNullableAny(meter.hits)
	meter.hits = int(int32((hx_post_26 - 1)))
	hitsPost := hx_post_26
	hitsPre := func() int {
		hx_pre_27 := int(int32((hxrt.IntFromNullableAny(meter.hits) + 1)))
		meter.hits = hx_pre_27
		return hx_pre_27
	}()
	meter.level = (func(hx_value_28 any) float64 {
		if hx_value_28 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_28.(float64)
	}(meter.level) + 1)
	hx_post_29 := func(hx_value_30 any) float64 {
		if hx_value_30 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_30.(float64)
	}(meter.level)
	meter.level = (hx_post_29 - 1)
	levelPost := hx_post_29
	var v_1 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("field="), meter.hits), hxrt.StringFromLiteral(":")), hitsPost), hxrt.StringFromLiteral(":")), hitsPre), hxrt.StringFromLiteral(":")), meter.level), hxrt.StringFromLiteral(":")), levelPost), hxrt.StringFromLiteral(":")), int(-int32(hxrt.IntFromNullableAny(meter.hits)))), hxrt.StringFromLiteral(":")), int(^int32(hxrt.IntFromNullableAny(meter.hits)))))
	hxrt.Println(v_1)
	Meter_total = int(int32((hxrt.IntFromNullableAny(Meter_total) + 1)))
	Meter_total = int(int32((hxrt.IntFromNullableAny(Meter_total) - 1)))
	hx_post_33 := hxrt.IntFromNullableAny(Meter_total)
	Meter_total = int(int32((hx_post_33 + 1)))
	totalPost := hx_post_33
	totalPre := note(hxrt.StringFromLiteral("static"), func() int {
		hx_pre_34 := int(int32((hxrt.IntFromNullableAny(Meter_total) - 1)))
		Meter_total = hx_pre_34
		return hx_pre_34
	}())
	Meter_scale = (func(hx_value_35 any) float64 {
		if hx_value_35 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_35.(float64)
	}(Meter_scale) - 1)
	scalePre := func() float64 {
		hx_pre_36 := (func(hx_value_37 any) float64 {
			if hx_value_37 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_37.(float64)
		}(Meter_scale) + 1)
		Meter_scale = hx_pre_36
		return hx_pre_36
	}()
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("static="), Meter_total), hxrt.StringFromLiteral(":")), totalPost), hxrt.StringFromLiteral(":")), totalPre), hxrt.StringFromLiteral(":")), Meter_scale), hxrt.StringFromLiteral(":")), scalePre), hxrt.StringFromLiteral(":")), -func(hx_value_40 any) float64 {
		if hx_value_40 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_40.(float64)
	}(Meter_scale)))
	hxrt.Println(v_2)
	slots := hxrt.NewArray(1, nil, 7)
	func() int {
		hx_array_target_41 := slots
		hx_array_index_42 := 0
		var hx_array_current_43 int = hxrt.IntFromNullableAny(hx_array_target_41.Get(hx_array_index_42))
		var hx_array_next_44 int = int(int32((hx_array_current_43 + 1)))
		hx_array_target_41.Set(hx_array_index_42, hx_array_next_44)
		return hx_array_current_43
	}()
	func() int {
		hx_array_target_45 := slots
		hx_array_index_46 := 2
		var hx_array_current_47 int = hxrt.IntFromNullableAny(hx_array_target_45.Get(hx_array_index_46))
		var hx_array_next_48 int = int(int32((hx_array_current_47 - 1)))
		hx_array_target_45.Set(hx_array_index_46, hx_array_next_48)
		return hx_array_next_48
	}()
	slotPost := func() int {
		hx_array_target_49 := slots
		hx_array_index_50 := 0
		var hx_array_current_51 int = hxrt.IntFromNullableAny(hx_array_target_49.Get(hx_array_index_50))
		var hx_array_next_52 int = int(int32((hx_array_current_51 - 1)))
		hx_array_target_49.Set(hx_array_index_50, hx_array_next_52)
		return hx_array_current_51
	}()
	slotPre := func() int {
		hx_array_target_53 := slots
		hx_array_index_54 := 2
		var hx_array_current_55 int = hxrt.IntFromNullableAny(hx_array_target_53.Get(hx_array_index_54))
		var hx_array_next_56 int = int(int32((hx_array_current_55 + 1)))
		hx_array_target_53.Set(hx_array_index_54, hx_array_next_56)
		return hx_array_next_56
	}()
	var position any = nil
	position = 0
	values := hxrt.NewArray(10, 20, 30)
	hx_post_57 := hxrt.IntFromNullableAny(position)
	position = int(int32((hx_post_57 + 1)))
	picked := hxrt.IntFromNullableAny(values.Get(hx_post_57))
	hx_array_target_59 := values
	hx_post_58 := hxrt.IntFromNullableAny(position)
	position = int(int32((hx_post_58 + 1)))
	hx_array_index_60 := hx_post_58
	hx_array_target_59.Set(hx_array_index_60, 99)
	func() int {
		hx_array_target_62 := slots
		hx_post_61 := hxrt.IntFromNullableAny(position)
		position = int(int32((hx_post_61 - 1)))
		hx_array_index_63 := hx_post_61
		var hx_array_current_64 int = hxrt.IntFromNullableAny(hx_array_target_62.Get(hx_array_index_63))
		var hx_array_next_65 int = int(int32((hx_array_current_64 + 1)))
		hx_array_target_62.Set(hx_array_index_63, hx_array_next_65)
		return hx_array_current_64
	}()
	var v_3 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("array="), hxrt.StdString(slots)), hxrt.StringFromLiteral(":")), slotPost), hxrt.StringFromLiteral(":")), slotPre), hxrt.StringFromLiteral(":")), picked), hxrt.StringFromLiteral(":")), hxrt.StdString(values)), hxrt.StringFromLiteral(":")), position), hxrt.StringFromLiteral(":")), int(-int32(hxrt.IntFromNullableAny(slots.Get(0))))))
	hxrt.Println(v_3)
	var tally_weight any
	var tally_seen any
	var tally_count any
	tally_count = 1
	tally_weight = 2.5
	tally_seen = false
	tally_count = int(int32((hxrt.IntFromNullableAny(tally_count) + 1)))
	hx_post_66 := hxrt.IntFromNullableAny(tally_count)
	tally_count = int(int32((hx_post_66 - 1)))
	countPost := hx_post_66
	weightPre := func() float64 {
		hx_pre_67 := (func(hx_value_68 any) float64 {
			if hx_value_68 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_68.(float64)
		}(tally_weight) + 1)
		tally_weight = hx_pre_67
		return hx_pre_67
	}()
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("struct="), tally_count), hxrt.StringFromLiteral(":")), countPost), hxrt.StringFromLiteral(":")), weightPre), hxrt.StringFromLiteral(":")), tally_weight), hxrt.StringFromLiteral(":")), hxrt.StdString(!func(hx_value_71 any) bool {
		if hx_value_71 == nil {
			var hx_zero_72 bool
			return hx_zero_72
		}
		return hx_value_71.(bool)
	}(tally_seen))))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("struct.arg="), spend(func() hxrt.Anonymous {
		hx_obj_73 := map[string]any{}
		hx_obj_73["count"] = 4
		hx_obj_73["weight"] = 0.5
		hx_obj_73["seen"] = nil
		return hxrt.AnonymousFromMap(hx_obj_73)
	}())), hxrt.StringFromLiteral(":")), attempt(hxrt.NewFunction(func() {
		spend(func() hxrt.Anonymous {
			hx_obj_74 := map[string]any{}
			hx_obj_74["count"] = nil
			hx_obj_74["weight"] = 1.5
			hx_obj_74["seen"] = true
			return hxrt.AnonymousFromMap(hx_obj_74)
		}())
	}))))
	hxrt.Println(v_5)
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("optional="), bump(4)), hxrt.StringFromLiteral(":")), attempt(hxrt.NewFunction(func() {
		bump(nil)
	}))))
	hxrt.Println(v_6)
	var flag any = nil
	flag = false
	if !func(hx_value_75 any) bool {
		if hx_value_75 == nil {
			var hx_zero_76 bool
			return hx_zero_76
		}
		return hx_value_75.(bool)
	}(flag) {
		hxrt.Println(any(hxrt.StringFromLiteral("not.local")))
	}
	rounds := 0
	var done any = nil
	done = false
	for !func(hx_value_77 any) bool {
		if hx_value_77 == nil {
			var hx_zero_78 bool
			return hx_zero_78
		}
		return hx_value_77.(bool)
	}(done) {
		rounds = int(int32((rounds + 1)))
		done = (rounds == 3)
	}
	negated := !func(hx_value_79 any) bool {
		if hx_value_79 == nil {
			var hx_zero_80 bool
			return hx_zero_80
		}
		return hx_value_79.(bool)
	}(flag)
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("bool="), hxrt.StdString(negated)), hxrt.StringFromLiteral(":")), flip(!func(hx_value_81 any) bool {
		if hx_value_81 == nil {
			var hx_zero_82 bool
			return hx_zero_82
		}
		return hx_value_81.(bool)
	}(meter.open))), hxrt.StringFromLiteral(":")), func() *string {
		var hx_if_85 *string
		if !func(hx_value_83 any) bool {
			if hx_value_83 == nil {
				var hx_zero_84 bool
				return hx_zero_84
			}
			return hx_value_83.(bool)
		}(Meter_ready) {
			hx_if_85 = hxrt.StringFromLiteral("ready")
		} else {
			hx_if_85 = hxrt.StringFromLiteral("waiting")
		}
		return hx_if_85
	}()), hxrt.StringFromLiteral(":")), rounds), hxrt.StringFromLiteral(":")), hxrt.StdString(!func(hx_value_86 any) bool {
		if hx_value_86 == nil {
			var hx_zero_87 bool
			return hx_zero_87
		}
		return hx_value_86.(bool)
	}(done))))
	hxrt.Println(v_7)
	var order any = nil
	order = 1
	sum := int((hxrt.Int32Wrap(func() int {
		hx_post_88 := hxrt.IntFromNullableAny(order)
		order = int(int32((hx_post_88 + 1)))
		return hx_post_88
	}()) + hxrt.Int32Wrap(hxrt.IntFromNullableAny(hxrt.IntFromNullableAny(order)))))
	calls := hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(pair(func() int {
		hx_post_89 := hxrt.IntFromNullableAny(order)
		order = int(int32((hx_post_89 + 1)))
		return hx_post_89
	}(), hxrt.IntFromNullableAny(order)), hxrt.StringFromLiteral("|")), pair(func() int {
		hx_pre_90 := int(int32((hxrt.IntFromNullableAny(order) + 1)))
		order = hx_pre_90
		return hx_pre_90
	}(), hxrt.IntFromNullableAny(order))), hxrt.StringFromLiteral("|")), pair(hxrt.IntFromNullableAny(order), func() int {
		hx_post_91 := hxrt.IntFromNullableAny(order)
		order = int(int32((hx_post_91 - 1)))
		return hx_post_91
	}())), hxrt.StringFromLiteral("|")), int((hxrt.Int32Wrap(hxrt.IntFromNullableAny(hxrt.IntFromNullableAny(order))) + hxrt.Int32Wrap(func() int {
		hx_post_92 := hxrt.IntFromNullableAny(order)
		order = int(int32((hx_post_92 + 1)))
		return hx_post_92
	}()))))
	var cursor any = nil
	cursor = 0
	counts := hxrt.NewArray(0, 0, 0)
	func() int {
		hx_array_target_94 := counts
		hx_post_93 := hxrt.IntFromNullableAny(cursor)
		cursor = int(int32((hx_post_93 + 1)))
		hx_array_index_95 := hx_post_93
		var hx_array_current_96 int = hxrt.IntFromNullableAny(hx_array_target_94.Get(hx_array_index_95))
		var hx_array_next_97 int = int(int32((hx_array_current_96 + 1)))
		hx_array_target_94.Set(hx_array_index_95, hx_array_next_97)
		return hx_array_current_96
	}()
	hx_post_98 := hxrt.IntFromNullableAny(cursor)
	cursor = int(int32((hx_post_98 + 1)))
	index := hx_post_98
	func() any {
		hx_array_target_99 := counts
		hx_array_index_100 := index
		var hx_array_current_101 any = hx_array_target_99.Get(hx_array_index_100)
		var hx_array_assigned_102 any = int((hxrt.Int32Wrap(hxrt.IntFromNullableAny(hx_array_current_101)) + hxrt.Int32Wrap(5)))
		hx_array_target_99.Set(hx_array_index_100, hx_array_assigned_102)
		return hx_array_assigned_102
	}()
	nested := func() int {
		hx_array_target_104 := counts
		hx_post_103 := hxrt.IntFromNullableAny(cursor)
		cursor = int(int32((hx_post_103 - 1)))
		hx_array_index_105 := hx_post_103
		var hx_array_current_106 int = hxrt.IntFromNullableAny(hx_array_target_104.Get(hx_array_index_105))
		var hx_array_next_107 int = int(int32((hx_array_current_106 + 1)))
		hx_array_target_104.Set(hx_array_index_105, hx_array_next_107)
		return hx_array_current_106
	}()
	var v_8 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("order="), sum), hxrt.StringFromLiteral(":")), calls), hxrt.StringFromLiteral(":")), order), hxrt.StringFromLiteral(":")), hxrt.StdString(counts)), hxrt.StringFromLiteral(":")), cursor), hxrt.StringFromLiteral(":")), nested))
	hxrt.Println(v_8)
	var v_9 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("trail="), hxrt.StringJoinAny(trail.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_9)
	var nothing any = nil
	nothing = nil
	var nothingFloat any = nil
	nothingFloat = nil
	var unknown any = nil
	unknown = nil
	var v_10 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("null.int="), attempt(hxrt.NewFunction(func() {
		nothing = int(int32((hxrt.IntFromNullableAny(nothing) + 1)))
	}))), hxrt.StringFromLiteral(":")), attempt(hxrt.NewFunction(func() {
		nothing = int(int32((hxrt.IntFromNullableAny(nothing) - 1)))
	}))), hxrt.StringFromLiteral(":")), attempt(hxrt.NewFunction(func() {
		note(hxrt.StringFromLiteral("null"), func() int {
			hx_post_108 := hxrt.IntFromNullableAny(nothing)
			nothing = int(int32((hx_post_108 - 1)))
			return hx_post_108
		}())
	}))), hxrt.StringFromLiteral(":")), attempt(hxrt.NewFunction(func() {
		sinkInt = int(-int32(hxrt.IntFromNullableAny(nothing)))
	}))), hxrt.StringFromLiteral(":")), attempt(hxrt.NewFunction(func() {
		sinkInt = int(^int32(hxrt.IntFromNullableAny(nothing)))
	}))), hxrt.StringFromLiteral(":")), nothing))
	hxrt.Println(v_10)
	var v_11 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("null.float="), attempt(hxrt.NewFunction(func() {
		nothingFloat = (func(hx_value_109 any) float64 {
			if hx_value_109 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_109.(float64)
		}(nothingFloat) + 1)
	}))), hxrt.StringFromLiteral(":")), attempt(hxrt.NewFunction(func() {
		nothingFloat = (func(hx_value_110 any) float64 {
			if hx_value_110 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_110.(float64)
		}(nothingFloat) + 1)
	}))), hxrt.StringFromLiteral(":")), attempt(hxrt.NewFunction(func() {
		sinkFloat = -func(hx_value_111 any) float64 {
			if hx_value_111 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_111.(float64)
		}(nothingFloat)
	}))), hxrt.StringFromLiteral(":")), nothingFloat))
	hxrt.Println(v_11)
	var v_12 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("null.field="), attempt(hxrt.NewFunction(func() {
		Meter_missing = int(int32((hxrt.IntFromNullableAny(Meter_missing) + 1)))
	}))), hxrt.StringFromLiteral(":")), attempt(hxrt.NewFunction(func() {
		meter.none = (func(hx_value_114 any) float64 {
			if hx_value_114 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_114.(float64)
		}(meter.none) + 1)
	}))), hxrt.StringFromLiteral(":")), attempt(hxrt.NewFunction(func() {
		func() int {
			hx_array_target_115 := slots
			hx_array_index_116 := 1
			var hx_array_current_117 int = hxrt.IntFromNullableAny(hx_array_target_115.Get(hx_array_index_116))
			var hx_array_next_118 int = int(int32((hx_array_current_117 - 1)))
			hx_array_target_115.Set(hx_array_index_116, hx_array_next_118)
			return hx_array_current_117
		}()
	}))), hxrt.StringFromLiteral(":")), Meter_missing), hxrt.StringFromLiteral(":")), meter.none), hxrt.StringFromLiteral(":")), slots.Get(1)))
	hxrt.Println(v_12)
	var v_13 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("null.bool="), hxrt.StdString(!func(hx_value_121 any) bool {
		if hx_value_121 == nil {
			var hx_zero_122 bool
			return hx_zero_122
		}
		return hx_value_121.(bool)
	}(unknown))), hxrt.StringFromLiteral(":")), func() *string {
		var hx_if_125 *string
		if !func(hx_value_123 any) bool {
			if hx_value_123 == nil {
				var hx_zero_124 bool
				return hx_zero_124
			}
			return hx_value_123.(bool)
		}(unknown) {
			hx_if_125 = hxrt.StringFromLiteral("taken")
		} else {
			hx_if_125 = hxrt.StringFromLiteral("skipped")
		}
		return hx_if_125
	}()), hxrt.StringFromLiteral(":")), hxrt.StdString(unknown)))
	hxrt.Println(v_13)
}

func note(label *string, value int) int {
	hx_arr_126 := trail
	hx_arr_126.Push(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(label, hxrt.StringFromLiteral("=")), value))
	return value
}

func pair(first int, second int) *string {
	return hxrt.StringConcatAny(hxrt.StringConcatAny(first, hxrt.StringFromLiteral(",")), second)
}

func pairFloat(first float64, second float64) *string {
	return hxrt.StringConcatAny(hxrt.StringConcatAny(first, hxrt.StringFromLiteral(",")), second)
}

var sinkFloat float64 = 0

var sinkInt int = 0

func spend(tally hxrt.Anonymous) *string {
	func() int {
		hx_structural_target_127 := tally
		hx_structural_old_128 := hxrt.IntFromNullableAny(hxrt.AnonymousGet(hx_structural_target_127, "count"))
		hx_structural_value_129 := int(int32((hx_structural_old_128 + 1)))
		hxrt.AnonymousSet(hx_structural_target_127, "count", hx_structural_value_129)
		return hx_structural_old_128
	}()
	post := func() int {
		hx_structural_target_130 := tally
		hx_structural_old_131 := hxrt.IntFromNullableAny(hxrt.AnonymousGet(hx_structural_target_130, "count"))
		hx_structural_value_132 := int(int32((hx_structural_old_131 - 1)))
		hxrt.AnonymousSet(hx_structural_target_130, "count", hx_structural_value_132)
		return hx_structural_old_131
	}()
	pre := func() float64 {
		hx_structural_target_133 := tally
		hx_structural_old_134 := func(hx_value_136 any) float64 {
			if hx_value_136 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_136.(float64)
		}(hxrt.AnonymousGet(hx_structural_target_133, "weight"))
		hx_structural_value_135 := (hx_structural_old_134 + 1)
		hxrt.AnonymousSet(hx_structural_target_133, "weight", hx_structural_value_135)
		return hx_structural_value_135
	}()
	func() float64 {
		hx_structural_target_137 := tally
		hx_structural_old_138 := func(hx_value_140 any) float64 {
			if hx_value_140 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_140.(float64)
		}(hxrt.AnonymousGet(hx_structural_target_137, "weight"))
		hx_structural_value_139 := (hx_structural_old_138 - 1)
		hxrt.AnonymousSet(hx_structural_target_137, "weight", hx_structural_value_139)
		return hx_structural_value_139
	}()
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatAny(func(hx_value_141 any) any {
		if hx_value_141 == nil {
			return nil
		}
		return hx_value_141.(int)
	}(hxrt.AnonymousGet(tally, "count")), hxrt.StringFromLiteral(":")), post), hxrt.StringFromLiteral(":")), pre), hxrt.StringFromLiteral(":")), func(hx_value_142 any) any {
		if hx_value_142 == nil {
			return nil
		}
		return hx_value_142.(float64)
	}(hxrt.AnonymousGet(tally, "weight"))), hxrt.StringFromLiteral(":")), hxrt.StdString(!func(hx_value_146 any) bool {
		if hx_value_146 == nil {
			var hx_zero_147 bool
			return hx_zero_147
		}
		return hx_value_146.(bool)
	}(func(hx_value_145 any) any {
		if hx_value_145 == nil {
			return nil
		}
		return hx_value_145.(bool)
	}(hxrt.AnonymousGet(tally, "seen"))))), hxrt.StringFromLiteral(":")), int(-int32(hxrt.IntFromNullableAny(func(hx_value_148 any) any {
		if hx_value_148 == nil {
			return nil
		}
		return hx_value_148.(int)
	}(hxrt.AnonymousGet(tally, "count"))))))
}

var trail *hxrt.Array = hxrt.NewArray()

type I_Meter interface {
}

type Meter struct {
	__hx_this I_Meter
	hits      any
	level     any
	open      any
	none      any
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
	self.none = nil
	self.open = true
	self.level = 0.25
	self.hits = 3
	return self
}

func (self *Meter) HxClassName() string {
	return "Meter"
}

var Meter_missing any

var Meter_ready any = false

var Meter_scale any = 1.5

var Meter_total any = 10

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
