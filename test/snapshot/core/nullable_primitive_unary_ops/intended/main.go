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
	return hxrt.IntFromNullableAny(start)
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
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("local.float="), amount), hxrt.StringFromLiteral(":")), floatPost), hxrt.StringFromLiteral(":")), floatPre), hxrt.StringFromLiteral(":")), -func(hx_value_19 any) float64 {
		if hx_value_19 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_19.(float64)
	}(amount)), hxrt.StringFromLiteral(":")), pairFloat(func() float64 {
		hx_post_20 := func(hx_value_21 any) float64 {
			if hx_value_21 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_21.(float64)
		}(amount)
		amount = (hx_post_20 - 1)
		return hx_post_20
	}(), func(hx_value_22 any) float64 {
		if hx_value_22 == nil {
			var hx_zero_23 float64
			return hx_zero_23
		}
		return hx_value_22.(float64)
	}(amount))))
	hxrt.Println(v)
	meter := New_Meter()
	meter.hits = int(int32((hxrt.IntFromNullableAny(meter.hits) + 1)))
	meter.hits = int(int32((hxrt.IntFromNullableAny(meter.hits) - 1)))
	hx_post_24 := hxrt.IntFromNullableAny(meter.hits)
	meter.hits = int(int32((hx_post_24 - 1)))
	hitsPost := hx_post_24
	hitsPre := func() int {
		hx_pre_25 := int(int32((hxrt.IntFromNullableAny(meter.hits) + 1)))
		meter.hits = hx_pre_25
		return hx_pre_25
	}()
	meter.level = (func(hx_value_26 any) float64 {
		if hx_value_26 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_26.(float64)
	}(meter.level) + 1)
	hx_post_27 := func(hx_value_28 any) float64 {
		if hx_value_28 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_28.(float64)
	}(meter.level)
	meter.level = (hx_post_27 - 1)
	levelPost := hx_post_27
	var v_1 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("field="), meter.hits), hxrt.StringFromLiteral(":")), hitsPost), hxrt.StringFromLiteral(":")), hitsPre), hxrt.StringFromLiteral(":")), meter.level), hxrt.StringFromLiteral(":")), levelPost), hxrt.StringFromLiteral(":")), int(-int32(hxrt.IntFromNullableAny(meter.hits)))), hxrt.StringFromLiteral(":")), int(^int32(hxrt.IntFromNullableAny(meter.hits)))))
	hxrt.Println(v_1)
	Meter_total = int(int32((hxrt.IntFromNullableAny(Meter_total) + 1)))
	Meter_total = int(int32((hxrt.IntFromNullableAny(Meter_total) - 1)))
	hx_post_29 := hxrt.IntFromNullableAny(Meter_total)
	Meter_total = int(int32((hx_post_29 + 1)))
	totalPost := hx_post_29
	totalPre := note(hxrt.StringFromLiteral("static"), func() int {
		hx_pre_30 := int(int32((hxrt.IntFromNullableAny(Meter_total) - 1)))
		Meter_total = hx_pre_30
		return hx_pre_30
	}())
	Meter_scale = (func(hx_value_31 any) float64 {
		if hx_value_31 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_31.(float64)
	}(Meter_scale) - 1)
	scalePre := func() float64 {
		hx_pre_32 := (func(hx_value_33 any) float64 {
			if hx_value_33 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_33.(float64)
		}(Meter_scale) + 1)
		Meter_scale = hx_pre_32
		return hx_pre_32
	}()
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("static="), Meter_total), hxrt.StringFromLiteral(":")), totalPost), hxrt.StringFromLiteral(":")), totalPre), hxrt.StringFromLiteral(":")), Meter_scale), hxrt.StringFromLiteral(":")), scalePre), hxrt.StringFromLiteral(":")), -func(hx_value_34 any) float64 {
		if hx_value_34 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_34.(float64)
	}(Meter_scale)))
	hxrt.Println(v_2)
	slots := hxrt.NewArray(1, nil, 7)
	func() int {
		hx_array_target_35 := slots
		hx_array_index_36 := 0
		var hx_array_current_37 int = hxrt.IntFromNullableAny(hx_array_target_35.Get(hx_array_index_36))
		var hx_array_next_38 int = int(int32((hx_array_current_37 + 1)))
		hx_array_target_35.Set(hx_array_index_36, hx_array_next_38)
		return hx_array_current_37
	}()
	func() int {
		hx_array_target_39 := slots
		hx_array_index_40 := 2
		var hx_array_current_41 int = hxrt.IntFromNullableAny(hx_array_target_39.Get(hx_array_index_40))
		var hx_array_next_42 int = int(int32((hx_array_current_41 - 1)))
		hx_array_target_39.Set(hx_array_index_40, hx_array_next_42)
		return hx_array_next_42
	}()
	slotPost := func() int {
		hx_array_target_43 := slots
		hx_array_index_44 := 0
		var hx_array_current_45 int = hxrt.IntFromNullableAny(hx_array_target_43.Get(hx_array_index_44))
		var hx_array_next_46 int = int(int32((hx_array_current_45 - 1)))
		hx_array_target_43.Set(hx_array_index_44, hx_array_next_46)
		return hx_array_current_45
	}()
	slotPre := func() int {
		hx_array_target_47 := slots
		hx_array_index_48 := 2
		var hx_array_current_49 int = hxrt.IntFromNullableAny(hx_array_target_47.Get(hx_array_index_48))
		var hx_array_next_50 int = int(int32((hx_array_current_49 + 1)))
		hx_array_target_47.Set(hx_array_index_48, hx_array_next_50)
		return hx_array_next_50
	}()
	var position any = nil
	position = 0
	values := hxrt.NewArray(10, 20, 30)
	hx_post_51 := hxrt.IntFromNullableAny(position)
	position = int(int32((hx_post_51 + 1)))
	picked := hxrt.IntFromNullableAny(values.Get(hx_post_51))
	hx_array_target_53 := values
	hx_post_52 := hxrt.IntFromNullableAny(position)
	position = int(int32((hx_post_52 + 1)))
	hx_array_index_54 := hx_post_52
	hx_array_target_53.Set(hx_array_index_54, 99)
	func() int {
		hx_array_target_56 := slots
		hx_post_55 := hxrt.IntFromNullableAny(position)
		position = int(int32((hx_post_55 - 1)))
		hx_array_index_57 := hx_post_55
		var hx_array_current_58 int = hxrt.IntFromNullableAny(hx_array_target_56.Get(hx_array_index_57))
		var hx_array_next_59 int = int(int32((hx_array_current_58 + 1)))
		hx_array_target_56.Set(hx_array_index_57, hx_array_next_59)
		return hx_array_current_58
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
	hx_post_60 := hxrt.IntFromNullableAny(tally_count)
	tally_count = int(int32((hx_post_60 - 1)))
	countPost := hx_post_60
	weightPre := func() float64 {
		hx_pre_61 := (func(hx_value_62 any) float64 {
			if hx_value_62 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_62.(float64)
		}(tally_weight) + 1)
		tally_weight = hx_pre_61
		return hx_pre_61
	}()
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("struct="), tally_count), hxrt.StringFromLiteral(":")), countPost), hxrt.StringFromLiteral(":")), weightPre), hxrt.StringFromLiteral(":")), tally_weight), hxrt.StringFromLiteral(":")), hxrt.StdString(!func(hx_value_63 any) bool {
		if hx_value_63 == nil {
			var hx_zero_64 bool
			return hx_zero_64
		}
		return hx_value_63.(bool)
	}(tally_seen))))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("struct.arg="), spend(func() hxrt.Anonymous {
		hx_obj_65 := map[string]any{}
		hx_obj_65["count"] = 4
		hx_obj_65["weight"] = 0.5
		hx_obj_65["seen"] = nil
		return hxrt.AnonymousFromMap(hx_obj_65)
	}())), hxrt.StringFromLiteral(":")), attempt(hxrt.NewFunction(func() {
		spend(func() hxrt.Anonymous {
			hx_obj_66 := map[string]any{}
			hx_obj_66["count"] = nil
			hx_obj_66["weight"] = 1.5
			hx_obj_66["seen"] = true
			return hxrt.AnonymousFromMap(hx_obj_66)
		}())
	}))))
	hxrt.Println(v_5)
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("optional="), bump(4)), hxrt.StringFromLiteral(":")), attempt(hxrt.NewFunction(func() {
		bump(nil)
	}))))
	hxrt.Println(v_6)
	var flag any = nil
	flag = false
	if !func(hx_value_67 any) bool {
		if hx_value_67 == nil {
			var hx_zero_68 bool
			return hx_zero_68
		}
		return hx_value_67.(bool)
	}(flag) {
		hxrt.Println(any(hxrt.StringFromLiteral("not.local")))
	}
	rounds := 0
	var done any = nil
	done = false
	for !func(hx_value_69 any) bool {
		if hx_value_69 == nil {
			var hx_zero_70 bool
			return hx_zero_70
		}
		return hx_value_69.(bool)
	}(done) {
		rounds = int(int32((rounds + 1)))
		done = (rounds == 3)
	}
	negated := !func(hx_value_71 any) bool {
		if hx_value_71 == nil {
			var hx_zero_72 bool
			return hx_zero_72
		}
		return hx_value_71.(bool)
	}(flag)
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("bool="), hxrt.StdString(negated)), hxrt.StringFromLiteral(":")), flip(!func(hx_value_73 any) bool {
		if hx_value_73 == nil {
			var hx_zero_74 bool
			return hx_zero_74
		}
		return hx_value_73.(bool)
	}(meter.open))), hxrt.StringFromLiteral(":")), func() *string {
		var hx_if_77 *string
		if !func(hx_value_75 any) bool {
			if hx_value_75 == nil {
				var hx_zero_76 bool
				return hx_zero_76
			}
			return hx_value_75.(bool)
		}(Meter_ready) {
			hx_if_77 = hxrt.StringFromLiteral("ready")
		} else {
			hx_if_77 = hxrt.StringFromLiteral("waiting")
		}
		return hx_if_77
	}()), hxrt.StringFromLiteral(":")), rounds), hxrt.StringFromLiteral(":")), hxrt.StdString(!func(hx_value_78 any) bool {
		if hx_value_78 == nil {
			var hx_zero_79 bool
			return hx_zero_79
		}
		return hx_value_78.(bool)
	}(done))))
	hxrt.Println(v_7)
	var order any = nil
	order = 1
	sum := int((hxrt.Int32Wrap(func() int {
		hx_post_80 := hxrt.IntFromNullableAny(order)
		order = int(int32((hx_post_80 + 1)))
		return hx_post_80
	}()) + hxrt.Int32Wrap(hxrt.IntFromNullableAny(hxrt.IntFromNullableAny(order)))))
	calls := hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(pair(func() int {
		hx_post_81 := hxrt.IntFromNullableAny(order)
		order = int(int32((hx_post_81 + 1)))
		return hx_post_81
	}(), hxrt.IntFromNullableAny(order)), hxrt.StringFromLiteral("|")), pair(func() int {
		hx_pre_82 := int(int32((hxrt.IntFromNullableAny(order) + 1)))
		order = hx_pre_82
		return hx_pre_82
	}(), hxrt.IntFromNullableAny(order))), hxrt.StringFromLiteral("|")), pair(hxrt.IntFromNullableAny(order), func() int {
		hx_post_83 := hxrt.IntFromNullableAny(order)
		order = int(int32((hx_post_83 - 1)))
		return hx_post_83
	}())), hxrt.StringFromLiteral("|")), int((hxrt.Int32Wrap(hxrt.IntFromNullableAny(hxrt.IntFromNullableAny(order))) + hxrt.Int32Wrap(func() int {
		hx_post_84 := hxrt.IntFromNullableAny(order)
		order = int(int32((hx_post_84 + 1)))
		return hx_post_84
	}()))))
	var cursor any = nil
	cursor = 0
	counts := hxrt.NewArray(0, 0, 0)
	func() int {
		hx_array_target_86 := counts
		hx_post_85 := hxrt.IntFromNullableAny(cursor)
		cursor = int(int32((hx_post_85 + 1)))
		hx_array_index_87 := hx_post_85
		var hx_array_current_88 int = hxrt.IntFromNullableAny(hx_array_target_86.Get(hx_array_index_87))
		var hx_array_next_89 int = int(int32((hx_array_current_88 + 1)))
		hx_array_target_86.Set(hx_array_index_87, hx_array_next_89)
		return hx_array_current_88
	}()
	hx_post_90 := hxrt.IntFromNullableAny(cursor)
	cursor = int(int32((hx_post_90 + 1)))
	index := hx_post_90
	func() any {
		hx_array_target_91 := counts
		hx_array_index_92 := index
		var hx_array_current_93 any = hx_array_target_91.Get(hx_array_index_92)
		var hx_array_assigned_94 any = int((hxrt.Int32Wrap(hxrt.IntFromNullableAny(hx_array_current_93)) + hxrt.Int32Wrap(5)))
		hx_array_target_91.Set(hx_array_index_92, hx_array_assigned_94)
		return hx_array_assigned_94
	}()
	nested := func() int {
		hx_array_target_96 := counts
		hx_post_95 := hxrt.IntFromNullableAny(cursor)
		cursor = int(int32((hx_post_95 - 1)))
		hx_array_index_97 := hx_post_95
		var hx_array_current_98 int = hxrt.IntFromNullableAny(hx_array_target_96.Get(hx_array_index_97))
		var hx_array_next_99 int = int(int32((hx_array_current_98 + 1)))
		hx_array_target_96.Set(hx_array_index_97, hx_array_next_99)
		return hx_array_current_98
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
			hx_post_100 := hxrt.IntFromNullableAny(nothing)
			nothing = int(int32((hx_post_100 - 1)))
			return hx_post_100
		}())
	}))), hxrt.StringFromLiteral(":")), attempt(hxrt.NewFunction(func() {
		sinkInt = int(-int32(hxrt.IntFromNullableAny(nothing)))
	}))), hxrt.StringFromLiteral(":")), attempt(hxrt.NewFunction(func() {
		sinkInt = int(^int32(hxrt.IntFromNullableAny(nothing)))
	}))), hxrt.StringFromLiteral(":")), nothing))
	hxrt.Println(v_10)
	var v_11 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("null.float="), attempt(hxrt.NewFunction(func() {
		nothingFloat = (func(hx_value_101 any) float64 {
			if hx_value_101 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_101.(float64)
		}(nothingFloat) + 1)
	}))), hxrt.StringFromLiteral(":")), attempt(hxrt.NewFunction(func() {
		nothingFloat = (func(hx_value_102 any) float64 {
			if hx_value_102 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_102.(float64)
		}(nothingFloat) + 1)
	}))), hxrt.StringFromLiteral(":")), attempt(hxrt.NewFunction(func() {
		sinkFloat = -func(hx_value_103 any) float64 {
			if hx_value_103 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_103.(float64)
		}(nothingFloat)
	}))), hxrt.StringFromLiteral(":")), nothingFloat))
	hxrt.Println(v_11)
	var v_12 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("null.field="), attempt(hxrt.NewFunction(func() {
		Meter_missing = int(int32((hxrt.IntFromNullableAny(Meter_missing) + 1)))
	}))), hxrt.StringFromLiteral(":")), attempt(hxrt.NewFunction(func() {
		meter.none = (func(hx_value_104 any) float64 {
			if hx_value_104 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_104.(float64)
		}(meter.none) + 1)
	}))), hxrt.StringFromLiteral(":")), attempt(hxrt.NewFunction(func() {
		func() int {
			hx_array_target_105 := slots
			hx_array_index_106 := 1
			var hx_array_current_107 int = hxrt.IntFromNullableAny(hx_array_target_105.Get(hx_array_index_106))
			var hx_array_next_108 int = int(int32((hx_array_current_107 - 1)))
			hx_array_target_105.Set(hx_array_index_106, hx_array_next_108)
			return hx_array_current_107
		}()
	}))), hxrt.StringFromLiteral(":")), Meter_missing), hxrt.StringFromLiteral(":")), meter.none), hxrt.StringFromLiteral(":")), slots.Get(1)))
	hxrt.Println(v_12)
	var v_13 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("null.bool="), hxrt.StdString(!func(hx_value_109 any) bool {
		if hx_value_109 == nil {
			var hx_zero_110 bool
			return hx_zero_110
		}
		return hx_value_109.(bool)
	}(unknown))), hxrt.StringFromLiteral(":")), func() *string {
		var hx_if_113 *string
		if !func(hx_value_111 any) bool {
			if hx_value_111 == nil {
				var hx_zero_112 bool
				return hx_zero_112
			}
			return hx_value_111.(bool)
		}(unknown) {
			hx_if_113 = hxrt.StringFromLiteral("taken")
		} else {
			hx_if_113 = hxrt.StringFromLiteral("skipped")
		}
		return hx_if_113
	}()), hxrt.StringFromLiteral(":")), hxrt.StdString(unknown)))
	hxrt.Println(v_13)
}

func note(label *string, value int) int {
	hx_arr_114 := trail
	hx_arr_114.Push(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(label, hxrt.StringFromLiteral("=")), value))
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
		hx_structural_target_115 := tally
		hx_structural_old_116 := hxrt.IntFromNullableAny(hxrt.AnonymousGet(hx_structural_target_115, "count"))
		hx_structural_value_117 := int(int32((hx_structural_old_116 + 1)))
		hxrt.AnonymousSet(hx_structural_target_115, "count", hx_structural_value_117)
		return hx_structural_old_116
	}()
	post := func() int {
		hx_structural_target_118 := tally
		hx_structural_old_119 := hxrt.IntFromNullableAny(hxrt.AnonymousGet(hx_structural_target_118, "count"))
		hx_structural_value_120 := int(int32((hx_structural_old_119 - 1)))
		hxrt.AnonymousSet(hx_structural_target_118, "count", hx_structural_value_120)
		return hx_structural_old_119
	}()
	pre := func() float64 {
		hx_structural_target_121 := tally
		hx_structural_old_122 := func(hx_value_124 any) float64 {
			if hx_value_124 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_124.(float64)
		}(hxrt.AnonymousGet(hx_structural_target_121, "weight"))
		hx_structural_value_123 := (hx_structural_old_122 + 1)
		hxrt.AnonymousSet(hx_structural_target_121, "weight", hx_structural_value_123)
		return hx_structural_value_123
	}()
	func() float64 {
		hx_structural_target_125 := tally
		hx_structural_old_126 := func(hx_value_128 any) float64 {
			if hx_value_128 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_128.(float64)
		}(hxrt.AnonymousGet(hx_structural_target_125, "weight"))
		hx_structural_value_127 := (hx_structural_old_126 - 1)
		hxrt.AnonymousSet(hx_structural_target_125, "weight", hx_structural_value_127)
		return hx_structural_value_127
	}()
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatAny(func(hx_value_129 any) any {
		if hx_value_129 == nil {
			return nil
		}
		return hx_value_129.(int)
	}(hxrt.AnonymousGet(tally, "count")), hxrt.StringFromLiteral(":")), post), hxrt.StringFromLiteral(":")), pre), hxrt.StringFromLiteral(":")), func(hx_value_130 any) any {
		if hx_value_130 == nil {
			return nil
		}
		return hx_value_130.(float64)
	}(hxrt.AnonymousGet(tally, "weight"))), hxrt.StringFromLiteral(":")), hxrt.StdString(!func(hx_value_132 any) bool {
		if hx_value_132 == nil {
			var hx_zero_133 bool
			return hx_zero_133
		}
		return hx_value_132.(bool)
	}(func(hx_value_131 any) any {
		if hx_value_131 == nil {
			return nil
		}
		return hx_value_131.(bool)
	}(hxrt.AnonymousGet(tally, "seen"))))), hxrt.StringFromLiteral(":")), int(-int32(hxrt.IntFromNullableAny(func(hx_value_134 any) any {
		if hx_value_134 == nil {
			return nil
		}
		return hx_value_134.(int)
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
