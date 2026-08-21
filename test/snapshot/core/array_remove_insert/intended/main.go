package main

import "snapshot/hxrt"

var events *hxrt.Array = hxrt.NewArray()

var holderEvaluations int = 0

func insertGeneric(first any, second any, pos int, value any) *string {
	values := hxrt.NewArray(first, second)
	func() {
		hx_insert_position_2 := pos
		var hx_insert_value_3 any = value
		values.Insert(hx_insert_position_2, hx_insert_value_3)
	}()
	return hxrt.StringConcatStringPtr(hxrt.StringConcatAny(values.Len(), hxrt.StringFromLiteral(":")), hxrt.StdString(values.Get(1)))
}

func main() {
	duplicate := hxrt.NewArray(1, 2, 1)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("remove.duplicate="), hxrt.StdString(func() bool {
		var hx_remove_value_5 any = 1
		hx_remove_index_6 := 0
		for hx_remove_index_6 < duplicate.Len() {
			hx_remove_element_7 := duplicate.Get(hx_remove_index_6)
			if hxrt.HaxeEqual(hx_remove_element_7, hx_remove_value_5) {
				duplicate.RemoveAt(hx_remove_index_6)
				return true
			}
			hx_remove_index_6 = (hx_remove_index_6 + 1)
		}
		return false
	}())), hxrt.StringFromLiteral(":")), hxrt.StringJoinAny(duplicate.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("remove.missing="), hxrt.StdString(func() bool {
		var hx_remove_value_9 any = 9
		hx_remove_index_10 := 0
		for hx_remove_index_10 < duplicate.Len() {
			hx_remove_element_11 := duplicate.Get(hx_remove_index_10)
			if hxrt.HaxeEqual(hx_remove_element_11, hx_remove_value_9) {
				duplicate.RemoveAt(hx_remove_index_10)
				return true
			}
			hx_remove_index_10 = (hx_remove_index_10 + 1)
		}
		return false
	}())), hxrt.StringFromLiteral(":")), hxrt.StringJoinAny(duplicate.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_1)
	strings := hxrt.NewArray(makeSame(), hxrt.StringFromLiteral("tail"))
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("remove.string="), hxrt.StdString(func() bool {
		var hx_remove_value_13 any = makeSame()
		hx_remove_index_14 := 0
		for hx_remove_index_14 < strings.Len() {
			hx_remove_element_15 := strings.Get(hx_remove_index_14)
			if hxrt.HaxeEqual(hx_remove_element_15, hx_remove_value_13) {
				strings.RemoveAt(hx_remove_index_14)
				return true
			}
			hx_remove_index_14 = (hx_remove_index_14 + 1)
		}
		return false
	}())), hxrt.StringFromLiteral(":")), hxrt.StringJoinAny(strings.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_2)
	nullableInts := hxrt.NewArray(nil, 1, nil)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("remove.null="), hxrt.StdString(func() bool {
		var hx_remove_value_17 any = nil
		hx_remove_index_18 := 0
		for hx_remove_index_18 < nullableInts.Len() {
			hx_remove_element_19 := nullableInts.Get(hx_remove_index_18)
			if hxrt.HaxeEqual(hx_remove_element_19, hx_remove_value_17) {
				nullableInts.RemoveAt(hx_remove_index_18)
				return true
			}
			hx_remove_index_18 = (hx_remove_index_18 + 1)
		}
		return false
	}())), hxrt.StringFromLiteral(":")), showNullableInts(nullableInts)))
	hxrt.Println(v_3)
	nullableStrings := hxrt.NewArray(nil, hxrt.StringFromLiteral("A"), hxrt.StringFromLiteral("null"), hxrt.StringFromLiteral("B"))
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("remove.nullString.literal="), hxrt.StdString(func() bool {
		var hx_remove_value_21 any = hxrt.StringFromLiteral("null")
		hx_remove_index_22 := 0
		for hx_remove_index_22 < nullableStrings.Len() {
			hx_remove_element_23 := nullableStrings.Get(hx_remove_index_22)
			if hxrt.HaxeEqual(hx_remove_element_23, hx_remove_value_21) {
				nullableStrings.RemoveAt(hx_remove_index_22)
				return true
			}
			hx_remove_index_22 = (hx_remove_index_22 + 1)
		}
		return false
	}())), hxrt.StringFromLiteral(":")), hxrt.StringJoinAny(nullableStrings.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("remove.nullString.null="), hxrt.StdString(func() bool {
		var hx_remove_value_25 any = nil
		hx_remove_index_26 := 0
		for hx_remove_index_26 < nullableStrings.Len() {
			hx_remove_element_27 := nullableStrings.Get(hx_remove_index_26)
			if hxrt.HaxeEqual(hx_remove_element_27, hx_remove_value_25) {
				nullableStrings.RemoveAt(hx_remove_index_26)
				return true
			}
			hx_remove_index_26 = (hx_remove_index_26 + 1)
		}
		return false
	}())), hxrt.StringFromLiteral(":")), hxrt.StringJoinAny(nullableStrings.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_5)
	first := New_SnapshotArrayMutationBox(1)
	second := New_SnapshotArrayMutationBox(2)
	boxes := hxrt.NewArray(first, second)
	var v_6 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("remove.object.other="), hxrt.StdString(func() bool {
		var hx_remove_value_29 any = New_SnapshotArrayMutationBox(1)
		hx_remove_index_30 := 0
		for hx_remove_index_30 < boxes.Len() {
			hx_remove_element_31 := boxes.Get(hx_remove_index_30)
			if hxrt.HaxeEqual(hx_remove_element_31, hx_remove_value_29) {
				boxes.RemoveAt(hx_remove_index_30)
				return true
			}
			hx_remove_index_30 = (hx_remove_index_30 + 1)
		}
		return false
	}())), hxrt.StringFromLiteral(":")), boxes.Len()))
	hxrt.Println(v_6)
	var v_7 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("remove.object.exact="), hxrt.StdString(func() bool {
		var hx_remove_value_33 any = first
		hx_remove_index_34 := 0
		for hx_remove_index_34 < boxes.Len() {
			hx_remove_element_35 := boxes.Get(hx_remove_index_34)
			if hxrt.HaxeEqual(hx_remove_element_35, hx_remove_value_33) {
				boxes.RemoveAt(hx_remove_index_34)
				return true
			}
			hx_remove_index_34 = (hx_remove_index_34 + 1)
		}
		return false
	}())), hxrt.StringFromLiteral(":")), boxes.Len()))
	hxrt.Println(v_7)
	atStart := hxrt.NewArray(1, 2)
	func() {
		hx_insert_position_37 := 0
		var hx_insert_value_38 any = 0
		atStart.Insert(hx_insert_position_37, hx_insert_value_38)
	}()
	var v_8 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("insert.start="), hxrt.StringJoinAny(atStart.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_8)
	inMiddle := hxrt.NewArray(1, 3)
	func() {
		hx_insert_position_40 := 1
		var hx_insert_value_41 any = 2
		inMiddle.Insert(hx_insert_position_40, hx_insert_value_41)
	}()
	var v_9 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("insert.middle="), hxrt.StringJoinAny(inMiddle.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_9)
	atEnd := hxrt.NewArray(1, 2)
	func() {
		hx_insert_position_43 := atEnd.Len()
		var hx_insert_value_44 any = 3
		atEnd.Insert(hx_insert_position_43, hx_insert_value_44)
	}()
	var v_10 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("insert.end="), hxrt.StringJoinAny(atEnd.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_10)
	oversized := hxrt.NewArray(1, 2)
	func() {
		hx_insert_position_46 := 99
		var hx_insert_value_47 any = 3
		oversized.Insert(hx_insert_position_46, hx_insert_value_47)
	}()
	var v_11 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("insert.oversized="), hxrt.StringJoinAny(oversized.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_11)
	negative := hxrt.NewArray(1, 3)
	func() {
		hx_insert_position_49 := -1
		var hx_insert_value_50 any = 2
		negative.Insert(hx_insert_position_49, hx_insert_value_50)
	}()
	var v_12 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("insert.negative="), hxrt.StringJoinAny(negative.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_12)
	tooNegative := hxrt.NewArray(2, 3)
	func() {
		hx_insert_position_52 := -99
		var hx_insert_value_53 any = 1
		tooNegative.Insert(hx_insert_position_52, hx_insert_value_53)
	}()
	var v_13 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("insert.tooNegative="), hxrt.StringJoinAny(tooNegative.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_13)
	empty := hxrt.NewArray()
	func() {
		hx_insert_position_55 := -1
		var hx_insert_value_56 any = 1
		empty.Insert(hx_insert_position_55, hx_insert_value_56)
	}()
	var v_14 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("insert.empty="), hxrt.StringJoinAny(empty.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_14)
	events = hxrt.NewArray()
	orderedInsert := hxrt.NewArray(1, 3)
	func() {
		hx_insert_position_58 := markedPosition()
		var hx_insert_value_59 any = markedInsertValue()
		orderedInsert.Insert(hx_insert_position_58, hx_insert_value_59)
	}()
	var v_15 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("order.insert="), hxrt.StringJoinAny(events.Values(), hxrt.StringFromLiteral(","))), hxrt.StringFromLiteral(":")), hxrt.StringJoinAny(orderedInsert.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_15)
	events = hxrt.NewArray()
	orderedRemove := hxrt.NewArray(1, 2)
	removedInOrder := func() bool {
		var hx_remove_value_61 any = markedRemoveValue()
		hx_remove_index_62 := 0
		for hx_remove_index_62 < orderedRemove.Len() {
			hx_remove_element_63 := orderedRemove.Get(hx_remove_index_62)
			if hxrt.HaxeEqual(hx_remove_element_63, hx_remove_value_61) {
				orderedRemove.RemoveAt(hx_remove_index_62)
				return true
			}
			hx_remove_index_62 = (hx_remove_index_62 + 1)
		}
		return false
	}()
	var v_16 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("order.remove="), hxrt.StringJoinAny(events.Values(), hxrt.StringFromLiteral(","))), hxrt.StringFromLiteral(":")), hxrt.StdString(removedInOrder)), hxrt.StringFromLiteral(":")), hxrt.StringJoinAny(orderedRemove.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_16)
	var v_17 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("generic.remove.string="), removeGeneric(makeSame(), hxrt.StringFromLiteral("tail"), makeSame())))
	hxrt.Println(v_17)
	var v_18 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("generic.remove.object.other="), removeGenericCount(first, second, New_SnapshotArrayMutationBox(1))))
	hxrt.Println(v_18)
	var v_19 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("generic.remove.object.exact="), removeGenericCount(first, second, first)))
	hxrt.Println(v_19)
	var v_20 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("generic.insert.string="), insertGeneric(makeSame(), hxrt.StringFromLiteral("tail"), -1, hxrt.StringFromLiteral("middle"))))
	hxrt.Println(v_20)
	var v_21 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("generic.remove.null="), removeGeneric(nil, 2, nil)))
	hxrt.Println(v_21)
	var v_22 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("generic.remove.nullString="), removeGenericFour(nil, hxrt.StringFromLiteral("A"), hxrt.StringFromLiteral("null"), hxrt.StringFromLiteral("B"), hxrt.StringFromLiteral("null"))))
	hxrt.Println(v_22)
	var v_23 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("generic.insert.null="), insertGeneric(nil, 2, -99, nil)))
	hxrt.Println(v_23)
	var v_24 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("generic.shift.string="), shiftGeneric(makeSame(), hxrt.StringFromLiteral("tail"))))
	hxrt.Println(v_24)
	var v_25 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("generic.shift.null="), shiftGeneric(nil, 2)))
	hxrt.Println(v_25)
	shifted := hxrt.NewArray(1, 2)
	shiftedAlias := shifted
	var v_26 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("shift.value="), func() any {
		hx_value_65 := shifted.Get(0)
		shifted.RemoveAt(0)
		return hx_value_65
	}()), hxrt.StringFromLiteral(":")), hxrt.StringJoinAny(shiftedAlias.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_26)
	emptyShift := hxrt.NewArray()
	var v_27 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("shift.empty="), hxrt.StdString(func() any {
		hx_value_67 := emptyShift.Get(0)
		emptyShift.RemoveAt(0)
		return hx_value_67
	}())), hxrt.StringFromLiteral(":")), emptyShift.Len()))
	hxrt.Println(v_27)
	ignoredShift := hxrt.NewArray(1, 2)
	func() any {
		hx_value_69 := ignoredShift.Get(0)
		ignoredShift.RemoveAt(0)
		return hx_value_69
	}()
	var v_28 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("shift.ignored="), hxrt.StringJoinAny(ignoredShift.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_28)
	holder := makeHolder()
	var v_29 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("field.remove="), hxrt.StdString(func() bool {
		hx_arr_70 := func(hx_obj_71 map[string]any) *hxrt.Array {
			hx_field_72 := hx_obj_71["values"]
			if hx_field_72 == nil {
				var hx_zero_73 *hxrt.Array
				return hx_zero_73
			}
			return hx_field_72.(*hxrt.Array)
		}(holder)
		var hx_remove_value_74 any = 1
		hx_remove_index_75 := 0
		for hx_remove_index_75 < hx_arr_70.Len() {
			hx_remove_element_76 := hx_arr_70.Get(hx_remove_index_75)
			if hxrt.HaxeEqual(hx_remove_element_76, hx_remove_value_74) {
				hx_arr_70.RemoveAt(hx_remove_index_75)
				return true
			}
			hx_remove_index_75 = (hx_remove_index_75 + 1)
		}
		return false
	}())), hxrt.StringFromLiteral(":")), hxrt.StringJoinAny(func(hx_obj_77 map[string]any) *hxrt.Array {
		hx_field_78 := hx_obj_77["values"]
		if hx_field_78 == nil {
			var hx_zero_79 *hxrt.Array
			return hx_zero_79
		}
		return hx_field_78.(*hxrt.Array)
	}(holder).Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_29)
	func() {
		hx_arr_80 := func(hx_obj_81 map[string]any) *hxrt.Array {
			hx_field_82 := hx_obj_81["values"]
			if hx_field_82 == nil {
				var hx_zero_83 *hxrt.Array
				return hx_zero_83
			}
			return hx_field_82.(*hxrt.Array)
		}(holder)
		hx_insert_position_84 := 1
		var hx_insert_value_85 any = 1
		hx_arr_80.Insert(hx_insert_position_84, hx_insert_value_85)
	}()
	var v_30 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("field.insert="), hxrt.StringJoinAny(func(hx_obj_86 map[string]any) *hxrt.Array {
		hx_field_87 := hx_obj_86["values"]
		if hx_field_87 == nil {
			var hx_zero_88 *hxrt.Array
			return hx_zero_88
		}
		return hx_field_87.(*hxrt.Array)
	}(holder).Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_30)
	holderEvaluations = 0
	removedFromTemporary := func() bool {
		hx_arr_89 := func(hx_obj_90 map[string]any) *hxrt.Array {
			hx_field_91 := hx_obj_90["values"]
			if hx_field_91 == nil {
				var hx_zero_92 *hxrt.Array
				return hx_zero_92
			}
			return hx_field_91.(*hxrt.Array)
		}(makeCountedHolder())
		var hx_remove_value_93 any = 1
		hx_remove_index_94 := 0
		for hx_remove_index_94 < hx_arr_89.Len() {
			hx_remove_element_95 := hx_arr_89.Get(hx_remove_index_94)
			if hxrt.HaxeEqual(hx_remove_element_95, hx_remove_value_93) {
				hx_arr_89.RemoveAt(hx_remove_index_94)
				return true
			}
			hx_remove_index_94 = (hx_remove_index_94 + 1)
		}
		return false
	}()
	var v_31 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("receiver.remove="), hxrt.StdString(removedFromTemporary)), hxrt.StringFromLiteral(":")), holderEvaluations))
	hxrt.Println(v_31)
	holderEvaluations = 0
	func() {
		hx_arr_96 := func(hx_obj_97 map[string]any) *hxrt.Array {
			hx_field_98 := hx_obj_97["values"]
			if hx_field_98 == nil {
				var hx_zero_99 *hxrt.Array
				return hx_zero_99
			}
			return hx_field_98.(*hxrt.Array)
		}(makeCountedHolder())
		hx_insert_position_100 := 1
		var hx_insert_value_101 any = 1
		hx_arr_96.Insert(hx_insert_position_100, hx_insert_value_101)
	}()
	var v_32 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("receiver.insert="), holderEvaluations))
	hxrt.Println(v_32)
}

func makeCountedHolder() map[string]any {
	holderEvaluations = int(int32((holderEvaluations + 1)))
	return makeHolder()
}

func makeHolder() map[string]any {
	hx_obj_102 := map[string]any{}
	hx_obj_102["values"] = hxrt.NewArray(1, 2, 1)
	return hx_obj_102
}

func makeSame() *string {
	return hxrt.StringFromLiteral("same")
}

func markedInsertValue() int {
	hx_arr_103 := events
	hx_arr_103.Push(hxrt.StringFromLiteral("value"))
	return 2
}

func markedPosition() int {
	hx_arr_104 := events
	hx_arr_104.Push(hxrt.StringFromLiteral("position"))
	return -1
}

func markedRemoveValue() int {
	hx_arr_105 := events
	hx_arr_105.Push(hxrt.StringFromLiteral("value"))
	return 1
}

func removeGeneric(first any, second any, value any) *string {
	values := hxrt.NewArray(first, second)
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StdString(func() bool {
		var hx_remove_value_107 any = value
		hx_remove_index_108 := 0
		for hx_remove_index_108 < values.Len() {
			hx_remove_element_109 := values.Get(hx_remove_index_108)
			if hxrt.HaxeEqual(hx_remove_element_109, hx_remove_value_107) {
				values.RemoveAt(hx_remove_index_108)
				return true
			}
			hx_remove_index_108 = (hx_remove_index_108 + 1)
		}
		return false
	}()), hxrt.StringFromLiteral(":")), values.Len()), hxrt.StringFromLiteral(":")), hxrt.StdString(values.Get(0)))
}

func removeGenericCount(first any, second any, value any) *string {
	values := hxrt.NewArray(first, second)
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StdString(func() bool {
		var hx_remove_value_111 any = value
		hx_remove_index_112 := 0
		for hx_remove_index_112 < values.Len() {
			hx_remove_element_113 := values.Get(hx_remove_index_112)
			if hxrt.HaxeEqual(hx_remove_element_113, hx_remove_value_111) {
				values.RemoveAt(hx_remove_index_112)
				return true
			}
			hx_remove_index_112 = (hx_remove_index_112 + 1)
		}
		return false
	}()), hxrt.StringFromLiteral(":")), values.Len())
}

func removeGenericFour(first any, second any, third any, fourth any, value any) *string {
	values := hxrt.NewArray(first, second, third, fourth)
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StdString(func() bool {
		var hx_remove_value_115 any = value
		hx_remove_index_116 := 0
		for hx_remove_index_116 < values.Len() {
			hx_remove_element_117 := values.Get(hx_remove_index_116)
			if hxrt.HaxeEqual(hx_remove_element_117, hx_remove_value_115) {
				values.RemoveAt(hx_remove_index_116)
				return true
			}
			hx_remove_index_116 = (hx_remove_index_116 + 1)
		}
		return false
	}()), hxrt.StringFromLiteral(":")), hxrt.StringJoinAny(func() *hxrt.Array {
		_g := hxrt.NewArray()
		_g1 := 0
		for _g1 < values.Len() {
			var item any = values.Get(_g1)
			_g1 = int(int32((_g1 + 1)))
			_g.Push(hxrt.StdString(item))
		}
		return _g
	}().Values(), hxrt.StringFromLiteral(",")))
}

func shiftGeneric(first any, second any) *string {
	values := hxrt.NewArray(first, second)
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StdString(func() any {
		hx_value_120 := values.Get(0)
		values.RemoveAt(0)
		return hx_value_120
	}()), hxrt.StringFromLiteral(":")), values.Len()), hxrt.StringFromLiteral(":")), hxrt.StdString(values.Get(0)))
}

func showNullableInts(values *hxrt.Array) *string {
	return hxrt.StringJoinAny(func() *hxrt.Array {
		_g := hxrt.NewArray()
		_g1 := 0
		for _g1 < values.Len() {
			var value any = values.Get(_g1)
			_g1 = int(int32((_g1 + 1)))
			_g.Push(hxrt.StdString(value))
		}
		return _g
	}().Values(), hxrt.StringFromLiteral(","))
}

type I_SnapshotArrayMutationBox interface {
}

type SnapshotArrayMutationBox struct {
	__hx_this I_SnapshotArrayMutationBox
	id        int
}

func New_SnapshotArrayMutationBox(id int) *SnapshotArrayMutationBox {
	self := &SnapshotArrayMutationBox{}
	self.__hx_this = self
	self.id = id
	return self
}
