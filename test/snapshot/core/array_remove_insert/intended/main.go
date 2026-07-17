package main

import "snapshot/hxrt"

var events []*string = []*string{}

var holderEvaluations int = 0

func insertGeneric(first any, second any, pos int, value any) *string {
	values := []any{first, second}
	func() {
		hx_insert_position_2 := pos
		var hx_insert_value_3 any = value
		hx_insert_length_4 := len(values)
		if hx_insert_position_2 < 0 {
			hx_insert_position_2 = (hx_insert_length_4 + hx_insert_position_2)
			if hx_insert_position_2 < 0 {
				hx_insert_position_2 = 0
			}
		}
		if hx_insert_position_2 > hx_insert_length_4 {
			hx_insert_position_2 = hx_insert_length_4
		}
		var hx_insert_zero_5 any
		values = append(values, hx_insert_zero_5)
		copy(values[(hx_insert_position_2+1):], values[hx_insert_position_2:])
		values[hx_insert_position_2] = hx_insert_value_3
	}()
	return hxrt.StringConcatStringPtr(hxrt.StringConcatAny(len(values), hxrt.StringFromLiteral(":")), hxrt.StdString(values[1]))
}

func main() {
	duplicate := []int{1, 2, 1}
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("remove.duplicate="), hxrt.StdString(func() bool {
		var hx_remove_value_7 int = 1
		for hx_remove_index_8, hx_remove_element_9 := range duplicate {
			if hx_remove_element_9 == hx_remove_value_7 {
				hx_remove_last_10 := (len(duplicate) - 1)
				copy(duplicate[hx_remove_index_8:], duplicate[(hx_remove_index_8+1):])
				var hx_remove_zero_11 int
				duplicate[hx_remove_last_10] = hx_remove_zero_11
				duplicate = duplicate[:hx_remove_last_10]
				return true
			}
		}
		return false
	}())), hxrt.StringFromLiteral(":")), hxrt.StringJoinAny(func(hx_sort_src_12 []int) []any {
		hx_sort_out_14 := make([]any, 0, len(hx_sort_src_12))
		for _, hx_sort_item_13 := range hx_sort_src_12 {
			hx_sort_out_14 = append(hx_sort_out_14, hx_sort_item_13)
		}
		return hx_sort_out_14
	}(duplicate), hxrt.StringFromLiteral(","))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("remove.missing="), hxrt.StdString(func() bool {
		var hx_remove_value_16 int = 9
		for hx_remove_index_17, hx_remove_element_18 := range duplicate {
			if hx_remove_element_18 == hx_remove_value_16 {
				hx_remove_last_19 := (len(duplicate) - 1)
				copy(duplicate[hx_remove_index_17:], duplicate[(hx_remove_index_17+1):])
				var hx_remove_zero_20 int
				duplicate[hx_remove_last_19] = hx_remove_zero_20
				duplicate = duplicate[:hx_remove_last_19]
				return true
			}
		}
		return false
	}())), hxrt.StringFromLiteral(":")), hxrt.StringJoinAny(func(hx_sort_src_21 []int) []any {
		hx_sort_out_23 := make([]any, 0, len(hx_sort_src_21))
		for _, hx_sort_item_22 := range hx_sort_src_21 {
			hx_sort_out_23 = append(hx_sort_out_23, hx_sort_item_22)
		}
		return hx_sort_out_23
	}(duplicate), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_1)
	strings := []*string{makeSame(), hxrt.StringFromLiteral("tail")}
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("remove.string="), hxrt.StdString(func() bool {
		var hx_remove_value_25 *string = makeSame()
		for hx_remove_index_26, hx_remove_element_27 := range strings {
			if hxrt.StringEqualStringPtr(hx_remove_element_27, hx_remove_value_25) {
				hx_remove_last_28 := (len(strings) - 1)
				copy(strings[hx_remove_index_26:], strings[(hx_remove_index_26+1):])
				var hx_remove_zero_29 *string
				strings[hx_remove_last_28] = hx_remove_zero_29
				strings = strings[:hx_remove_last_28]
				return true
			}
		}
		return false
	}())), hxrt.StringFromLiteral(":")), hxrt.StringJoinAny(func(hx_sort_src_30 []*string) []any {
		hx_sort_out_32 := make([]any, 0, len(hx_sort_src_30))
		for _, hx_sort_item_31 := range hx_sort_src_30 {
			hx_sort_out_32 = append(hx_sort_out_32, hx_sort_item_31)
		}
		return hx_sort_out_32
	}(strings), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_2)
	nullableInts := []any{nil, 1, nil}
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("remove.null="), hxrt.StdString(func() bool {
		var hx_remove_value_34 any = nil
		for hx_remove_index_35, hx_remove_element_36 := range nullableInts {
			if hxrt.HaxeEqual(hx_remove_element_36, hx_remove_value_34) {
				hx_remove_last_37 := (len(nullableInts) - 1)
				copy(nullableInts[hx_remove_index_35:], nullableInts[(hx_remove_index_35+1):])
				var hx_remove_zero_38 any
				nullableInts[hx_remove_last_37] = hx_remove_zero_38
				nullableInts = nullableInts[:hx_remove_last_37]
				return true
			}
		}
		return false
	}())), hxrt.StringFromLiteral(":")), showNullableInts(nullableInts)))
	hxrt.Println(v_3)
	nullableStrings := []*string{nil, hxrt.StringFromLiteral("A"), hxrt.StringFromLiteral("null"), hxrt.StringFromLiteral("B")}
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("remove.nullString.literal="), hxrt.StdString(func() bool {
		var hx_remove_value_40 *string = hxrt.StringFromLiteral("null")
		for hx_remove_index_41, hx_remove_element_42 := range nullableStrings {
			if hxrt.StringEqualStringPtr(hx_remove_element_42, hx_remove_value_40) {
				hx_remove_last_43 := (len(nullableStrings) - 1)
				copy(nullableStrings[hx_remove_index_41:], nullableStrings[(hx_remove_index_41+1):])
				var hx_remove_zero_44 *string
				nullableStrings[hx_remove_last_43] = hx_remove_zero_44
				nullableStrings = nullableStrings[:hx_remove_last_43]
				return true
			}
		}
		return false
	}())), hxrt.StringFromLiteral(":")), hxrt.StringJoinAny(func(hx_sort_src_45 []*string) []any {
		hx_sort_out_47 := make([]any, 0, len(hx_sort_src_45))
		for _, hx_sort_item_46 := range hx_sort_src_45 {
			hx_sort_out_47 = append(hx_sort_out_47, hx_sort_item_46)
		}
		return hx_sort_out_47
	}(nullableStrings), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("remove.nullString.null="), hxrt.StdString(func() bool {
		var hx_remove_value_49 *string = nil
		for hx_remove_index_50, hx_remove_element_51 := range nullableStrings {
			if hxrt.StringEqualStringPtr(hx_remove_element_51, hx_remove_value_49) {
				hx_remove_last_52 := (len(nullableStrings) - 1)
				copy(nullableStrings[hx_remove_index_50:], nullableStrings[(hx_remove_index_50+1):])
				var hx_remove_zero_53 *string
				nullableStrings[hx_remove_last_52] = hx_remove_zero_53
				nullableStrings = nullableStrings[:hx_remove_last_52]
				return true
			}
		}
		return false
	}())), hxrt.StringFromLiteral(":")), hxrt.StringJoinAny(func(hx_sort_src_54 []*string) []any {
		hx_sort_out_56 := make([]any, 0, len(hx_sort_src_54))
		for _, hx_sort_item_55 := range hx_sort_src_54 {
			hx_sort_out_56 = append(hx_sort_out_56, hx_sort_item_55)
		}
		return hx_sort_out_56
	}(nullableStrings), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_5)
	first := New_SnapshotArrayMutationBox(1)
	second := New_SnapshotArrayMutationBox(2)
	boxes := []*SnapshotArrayMutationBox{first, second}
	var v_6 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("remove.object.other="), hxrt.StdString(func() bool {
		var hx_remove_value_58 *SnapshotArrayMutationBox = New_SnapshotArrayMutationBox(1)
		for hx_remove_index_59, hx_remove_element_60 := range boxes {
			if hx_remove_element_60 == hx_remove_value_58 {
				hx_remove_last_61 := (len(boxes) - 1)
				copy(boxes[hx_remove_index_59:], boxes[(hx_remove_index_59+1):])
				var hx_remove_zero_62 *SnapshotArrayMutationBox
				boxes[hx_remove_last_61] = hx_remove_zero_62
				boxes = boxes[:hx_remove_last_61]
				return true
			}
		}
		return false
	}())), hxrt.StringFromLiteral(":")), len(boxes)))
	hxrt.Println(v_6)
	var v_7 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("remove.object.exact="), hxrt.StdString(func() bool {
		var hx_remove_value_64 *SnapshotArrayMutationBox = first
		for hx_remove_index_65, hx_remove_element_66 := range boxes {
			if hx_remove_element_66 == hx_remove_value_64 {
				hx_remove_last_67 := (len(boxes) - 1)
				copy(boxes[hx_remove_index_65:], boxes[(hx_remove_index_65+1):])
				var hx_remove_zero_68 *SnapshotArrayMutationBox
				boxes[hx_remove_last_67] = hx_remove_zero_68
				boxes = boxes[:hx_remove_last_67]
				return true
			}
		}
		return false
	}())), hxrt.StringFromLiteral(":")), len(boxes)))
	hxrt.Println(v_7)
	atStart := []int{1, 2}
	func() {
		hx_insert_position_70 := 0
		var hx_insert_value_71 int = 0
		hx_insert_length_72 := len(atStart)
		if hx_insert_position_70 < 0 {
			hx_insert_position_70 = (hx_insert_length_72 + hx_insert_position_70)
			if hx_insert_position_70 < 0 {
				hx_insert_position_70 = 0
			}
		}
		if hx_insert_position_70 > hx_insert_length_72 {
			hx_insert_position_70 = hx_insert_length_72
		}
		var hx_insert_zero_73 int
		atStart = append(atStart, hx_insert_zero_73)
		copy(atStart[(hx_insert_position_70+1):], atStart[hx_insert_position_70:])
		atStart[hx_insert_position_70] = hx_insert_value_71
	}()
	var v_8 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("insert.start="), hxrt.StringJoinAny(func(hx_sort_src_74 []int) []any {
		hx_sort_out_76 := make([]any, 0, len(hx_sort_src_74))
		for _, hx_sort_item_75 := range hx_sort_src_74 {
			hx_sort_out_76 = append(hx_sort_out_76, hx_sort_item_75)
		}
		return hx_sort_out_76
	}(atStart), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_8)
	inMiddle := []int{1, 3}
	func() {
		hx_insert_position_78 := 1
		var hx_insert_value_79 int = 2
		hx_insert_length_80 := len(inMiddle)
		if hx_insert_position_78 < 0 {
			hx_insert_position_78 = (hx_insert_length_80 + hx_insert_position_78)
			if hx_insert_position_78 < 0 {
				hx_insert_position_78 = 0
			}
		}
		if hx_insert_position_78 > hx_insert_length_80 {
			hx_insert_position_78 = hx_insert_length_80
		}
		var hx_insert_zero_81 int
		inMiddle = append(inMiddle, hx_insert_zero_81)
		copy(inMiddle[(hx_insert_position_78+1):], inMiddle[hx_insert_position_78:])
		inMiddle[hx_insert_position_78] = hx_insert_value_79
	}()
	var v_9 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("insert.middle="), hxrt.StringJoinAny(func(hx_sort_src_82 []int) []any {
		hx_sort_out_84 := make([]any, 0, len(hx_sort_src_82))
		for _, hx_sort_item_83 := range hx_sort_src_82 {
			hx_sort_out_84 = append(hx_sort_out_84, hx_sort_item_83)
		}
		return hx_sort_out_84
	}(inMiddle), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_9)
	atEnd := []int{1, 2}
	func() {
		hx_insert_position_86 := len(atEnd)
		var hx_insert_value_87 int = 3
		hx_insert_length_88 := len(atEnd)
		if hx_insert_position_86 < 0 {
			hx_insert_position_86 = (hx_insert_length_88 + hx_insert_position_86)
			if hx_insert_position_86 < 0 {
				hx_insert_position_86 = 0
			}
		}
		if hx_insert_position_86 > hx_insert_length_88 {
			hx_insert_position_86 = hx_insert_length_88
		}
		var hx_insert_zero_89 int
		atEnd = append(atEnd, hx_insert_zero_89)
		copy(atEnd[(hx_insert_position_86+1):], atEnd[hx_insert_position_86:])
		atEnd[hx_insert_position_86] = hx_insert_value_87
	}()
	var v_10 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("insert.end="), hxrt.StringJoinAny(func(hx_sort_src_90 []int) []any {
		hx_sort_out_92 := make([]any, 0, len(hx_sort_src_90))
		for _, hx_sort_item_91 := range hx_sort_src_90 {
			hx_sort_out_92 = append(hx_sort_out_92, hx_sort_item_91)
		}
		return hx_sort_out_92
	}(atEnd), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_10)
	oversized := []int{1, 2}
	func() {
		hx_insert_position_94 := 99
		var hx_insert_value_95 int = 3
		hx_insert_length_96 := len(oversized)
		if hx_insert_position_94 < 0 {
			hx_insert_position_94 = (hx_insert_length_96 + hx_insert_position_94)
			if hx_insert_position_94 < 0 {
				hx_insert_position_94 = 0
			}
		}
		if hx_insert_position_94 > hx_insert_length_96 {
			hx_insert_position_94 = hx_insert_length_96
		}
		var hx_insert_zero_97 int
		oversized = append(oversized, hx_insert_zero_97)
		copy(oversized[(hx_insert_position_94+1):], oversized[hx_insert_position_94:])
		oversized[hx_insert_position_94] = hx_insert_value_95
	}()
	var v_11 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("insert.oversized="), hxrt.StringJoinAny(func(hx_sort_src_98 []int) []any {
		hx_sort_out_100 := make([]any, 0, len(hx_sort_src_98))
		for _, hx_sort_item_99 := range hx_sort_src_98 {
			hx_sort_out_100 = append(hx_sort_out_100, hx_sort_item_99)
		}
		return hx_sort_out_100
	}(oversized), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_11)
	negative := []int{1, 3}
	func() {
		hx_insert_position_102 := -1
		var hx_insert_value_103 int = 2
		hx_insert_length_104 := len(negative)
		if hx_insert_position_102 < 0 {
			hx_insert_position_102 = (hx_insert_length_104 + hx_insert_position_102)
			if hx_insert_position_102 < 0 {
				hx_insert_position_102 = 0
			}
		}
		if hx_insert_position_102 > hx_insert_length_104 {
			hx_insert_position_102 = hx_insert_length_104
		}
		var hx_insert_zero_105 int
		negative = append(negative, hx_insert_zero_105)
		copy(negative[(hx_insert_position_102+1):], negative[hx_insert_position_102:])
		negative[hx_insert_position_102] = hx_insert_value_103
	}()
	var v_12 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("insert.negative="), hxrt.StringJoinAny(func(hx_sort_src_106 []int) []any {
		hx_sort_out_108 := make([]any, 0, len(hx_sort_src_106))
		for _, hx_sort_item_107 := range hx_sort_src_106 {
			hx_sort_out_108 = append(hx_sort_out_108, hx_sort_item_107)
		}
		return hx_sort_out_108
	}(negative), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_12)
	tooNegative := []int{2, 3}
	func() {
		hx_insert_position_110 := -99
		var hx_insert_value_111 int = 1
		hx_insert_length_112 := len(tooNegative)
		if hx_insert_position_110 < 0 {
			hx_insert_position_110 = (hx_insert_length_112 + hx_insert_position_110)
			if hx_insert_position_110 < 0 {
				hx_insert_position_110 = 0
			}
		}
		if hx_insert_position_110 > hx_insert_length_112 {
			hx_insert_position_110 = hx_insert_length_112
		}
		var hx_insert_zero_113 int
		tooNegative = append(tooNegative, hx_insert_zero_113)
		copy(tooNegative[(hx_insert_position_110+1):], tooNegative[hx_insert_position_110:])
		tooNegative[hx_insert_position_110] = hx_insert_value_111
	}()
	var v_13 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("insert.tooNegative="), hxrt.StringJoinAny(func(hx_sort_src_114 []int) []any {
		hx_sort_out_116 := make([]any, 0, len(hx_sort_src_114))
		for _, hx_sort_item_115 := range hx_sort_src_114 {
			hx_sort_out_116 = append(hx_sort_out_116, hx_sort_item_115)
		}
		return hx_sort_out_116
	}(tooNegative), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_13)
	empty := []int{}
	func() {
		hx_insert_position_118 := -1
		var hx_insert_value_119 int = 1
		hx_insert_length_120 := len(empty)
		if hx_insert_position_118 < 0 {
			hx_insert_position_118 = (hx_insert_length_120 + hx_insert_position_118)
			if hx_insert_position_118 < 0 {
				hx_insert_position_118 = 0
			}
		}
		if hx_insert_position_118 > hx_insert_length_120 {
			hx_insert_position_118 = hx_insert_length_120
		}
		var hx_insert_zero_121 int
		empty = append(empty, hx_insert_zero_121)
		copy(empty[(hx_insert_position_118+1):], empty[hx_insert_position_118:])
		empty[hx_insert_position_118] = hx_insert_value_119
	}()
	var v_14 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("insert.empty="), hxrt.StringJoinAny(func(hx_sort_src_122 []int) []any {
		hx_sort_out_124 := make([]any, 0, len(hx_sort_src_122))
		for _, hx_sort_item_123 := range hx_sort_src_122 {
			hx_sort_out_124 = append(hx_sort_out_124, hx_sort_item_123)
		}
		return hx_sort_out_124
	}(empty), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_14)
	events = []*string{}
	orderedInsert := []int{1, 3}
	func() {
		hx_insert_position_126 := markedPosition()
		var hx_insert_value_127 int = markedInsertValue()
		hx_insert_length_128 := len(orderedInsert)
		if hx_insert_position_126 < 0 {
			hx_insert_position_126 = (hx_insert_length_128 + hx_insert_position_126)
			if hx_insert_position_126 < 0 {
				hx_insert_position_126 = 0
			}
		}
		if hx_insert_position_126 > hx_insert_length_128 {
			hx_insert_position_126 = hx_insert_length_128
		}
		var hx_insert_zero_129 int
		orderedInsert = append(orderedInsert, hx_insert_zero_129)
		copy(orderedInsert[(hx_insert_position_126+1):], orderedInsert[hx_insert_position_126:])
		orderedInsert[hx_insert_position_126] = hx_insert_value_127
	}()
	var v_15 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("order.insert="), hxrt.StringJoinAny(func(hx_sort_src_130 []*string) []any {
		hx_sort_out_132 := make([]any, 0, len(hx_sort_src_130))
		for _, hx_sort_item_131 := range hx_sort_src_130 {
			hx_sort_out_132 = append(hx_sort_out_132, hx_sort_item_131)
		}
		return hx_sort_out_132
	}(events), hxrt.StringFromLiteral(","))), hxrt.StringFromLiteral(":")), hxrt.StringJoinAny(func(hx_sort_src_133 []int) []any {
		hx_sort_out_135 := make([]any, 0, len(hx_sort_src_133))
		for _, hx_sort_item_134 := range hx_sort_src_133 {
			hx_sort_out_135 = append(hx_sort_out_135, hx_sort_item_134)
		}
		return hx_sort_out_135
	}(orderedInsert), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_15)
	events = []*string{}
	orderedRemove := []int{1, 2}
	removedInOrder := func() bool {
		var hx_remove_value_137 int = markedRemoveValue()
		for hx_remove_index_138, hx_remove_element_139 := range orderedRemove {
			if hx_remove_element_139 == hx_remove_value_137 {
				hx_remove_last_140 := (len(orderedRemove) - 1)
				copy(orderedRemove[hx_remove_index_138:], orderedRemove[(hx_remove_index_138+1):])
				var hx_remove_zero_141 int
				orderedRemove[hx_remove_last_140] = hx_remove_zero_141
				orderedRemove = orderedRemove[:hx_remove_last_140]
				return true
			}
		}
		return false
	}()
	var v_16 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("order.remove="), hxrt.StringJoinAny(func(hx_sort_src_142 []*string) []any {
		hx_sort_out_144 := make([]any, 0, len(hx_sort_src_142))
		for _, hx_sort_item_143 := range hx_sort_src_142 {
			hx_sort_out_144 = append(hx_sort_out_144, hx_sort_item_143)
		}
		return hx_sort_out_144
	}(events), hxrt.StringFromLiteral(","))), hxrt.StringFromLiteral(":")), hxrt.StdString(removedInOrder)), hxrt.StringFromLiteral(":")), hxrt.StringJoinAny(func(hx_sort_src_145 []int) []any {
		hx_sort_out_147 := make([]any, 0, len(hx_sort_src_145))
		for _, hx_sort_item_146 := range hx_sort_src_145 {
			hx_sort_out_147 = append(hx_sort_out_147, hx_sort_item_146)
		}
		return hx_sort_out_147
	}(orderedRemove), hxrt.StringFromLiteral(","))))
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
	holder := makeHolder()
	var v_24 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("field.remove="), hxrt.StdString(func() bool {
		hx_obj_149 := holder
		hx_arr_148 := func(hx_obj_150 map[string]any) []int {
			hx_field_151 := hx_obj_150["values"]
			if hx_field_151 == nil {
				var hx_zero_152 []int
				return hx_zero_152
			}
			return hx_field_151.([]int)
		}(hx_obj_149)
		var hx_remove_value_153 int = 1
		for hx_remove_index_154, hx_remove_element_155 := range hx_arr_148 {
			if hx_remove_element_155 == hx_remove_value_153 {
				hx_remove_last_156 := (len(hx_arr_148) - 1)
				copy(hx_arr_148[hx_remove_index_154:], hx_arr_148[(hx_remove_index_154+1):])
				var hx_remove_zero_157 int
				hx_arr_148[hx_remove_last_156] = hx_remove_zero_157
				hx_arr_148 = hx_arr_148[:hx_remove_last_156]
				hx_obj_149["values"] = hx_arr_148
				return true
			}
		}
		return false
	}())), hxrt.StringFromLiteral(":")), hxrt.StringJoinAny(func(hx_sort_src_161 []int) []any {
		hx_sort_out_163 := make([]any, 0, len(hx_sort_src_161))
		for _, hx_sort_item_162 := range hx_sort_src_161 {
			hx_sort_out_163 = append(hx_sort_out_163, hx_sort_item_162)
		}
		return hx_sort_out_163
	}(func(hx_obj_158 map[string]any) []int {
		hx_field_159 := hx_obj_158["values"]
		if hx_field_159 == nil {
			var hx_zero_160 []int
			return hx_zero_160
		}
		return hx_field_159.([]int)
	}(holder)), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_24)
	func() {
		hx_obj_165 := holder
		hx_arr_164 := func(hx_obj_166 map[string]any) []int {
			hx_field_167 := hx_obj_166["values"]
			if hx_field_167 == nil {
				var hx_zero_168 []int
				return hx_zero_168
			}
			return hx_field_167.([]int)
		}(hx_obj_165)
		hx_insert_position_169 := 1
		var hx_insert_value_170 int = 1
		hx_insert_length_171 := len(hx_arr_164)
		if hx_insert_position_169 < 0 {
			hx_insert_position_169 = (hx_insert_length_171 + hx_insert_position_169)
			if hx_insert_position_169 < 0 {
				hx_insert_position_169 = 0
			}
		}
		if hx_insert_position_169 > hx_insert_length_171 {
			hx_insert_position_169 = hx_insert_length_171
		}
		var hx_insert_zero_172 int
		hx_arr_164 = append(hx_arr_164, hx_insert_zero_172)
		copy(hx_arr_164[(hx_insert_position_169+1):], hx_arr_164[hx_insert_position_169:])
		hx_arr_164[hx_insert_position_169] = hx_insert_value_170
		hx_obj_165["values"] = hx_arr_164
	}()
	var v_25 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("field.insert="), hxrt.StringJoinAny(func(hx_sort_src_176 []int) []any {
		hx_sort_out_178 := make([]any, 0, len(hx_sort_src_176))
		for _, hx_sort_item_177 := range hx_sort_src_176 {
			hx_sort_out_178 = append(hx_sort_out_178, hx_sort_item_177)
		}
		return hx_sort_out_178
	}(func(hx_obj_173 map[string]any) []int {
		hx_field_174 := hx_obj_173["values"]
		if hx_field_174 == nil {
			var hx_zero_175 []int
			return hx_zero_175
		}
		return hx_field_174.([]int)
	}(holder)), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_25)
	holderEvaluations = 0
	removedFromTemporary := func() bool {
		hx_obj_180 := makeCountedHolder()
		hx_arr_179 := func(hx_obj_181 map[string]any) []int {
			hx_field_182 := hx_obj_181["values"]
			if hx_field_182 == nil {
				var hx_zero_183 []int
				return hx_zero_183
			}
			return hx_field_182.([]int)
		}(hx_obj_180)
		var hx_remove_value_184 int = 1
		for hx_remove_index_185, hx_remove_element_186 := range hx_arr_179 {
			if hx_remove_element_186 == hx_remove_value_184 {
				hx_remove_last_187 := (len(hx_arr_179) - 1)
				copy(hx_arr_179[hx_remove_index_185:], hx_arr_179[(hx_remove_index_185+1):])
				var hx_remove_zero_188 int
				hx_arr_179[hx_remove_last_187] = hx_remove_zero_188
				hx_arr_179 = hx_arr_179[:hx_remove_last_187]
				hx_obj_180["values"] = hx_arr_179
				return true
			}
		}
		return false
	}()
	var v_26 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("receiver.remove="), hxrt.StdString(removedFromTemporary)), hxrt.StringFromLiteral(":")), holderEvaluations))
	hxrt.Println(v_26)
	holderEvaluations = 0
	func() {
		hx_obj_190 := makeCountedHolder()
		hx_arr_189 := func(hx_obj_191 map[string]any) []int {
			hx_field_192 := hx_obj_191["values"]
			if hx_field_192 == nil {
				var hx_zero_193 []int
				return hx_zero_193
			}
			return hx_field_192.([]int)
		}(hx_obj_190)
		hx_insert_position_194 := 1
		var hx_insert_value_195 int = 1
		hx_insert_length_196 := len(hx_arr_189)
		if hx_insert_position_194 < 0 {
			hx_insert_position_194 = (hx_insert_length_196 + hx_insert_position_194)
			if hx_insert_position_194 < 0 {
				hx_insert_position_194 = 0
			}
		}
		if hx_insert_position_194 > hx_insert_length_196 {
			hx_insert_position_194 = hx_insert_length_196
		}
		var hx_insert_zero_197 int
		hx_arr_189 = append(hx_arr_189, hx_insert_zero_197)
		copy(hx_arr_189[(hx_insert_position_194+1):], hx_arr_189[hx_insert_position_194:])
		hx_arr_189[hx_insert_position_194] = hx_insert_value_195
		hx_obj_190["values"] = hx_arr_189
	}()
	var v_27 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("receiver.insert="), holderEvaluations))
	hxrt.Println(v_27)
}

func makeCountedHolder() map[string]any {
	holderEvaluations = int(int32((holderEvaluations + 1)))
	return makeHolder()
}

func makeHolder() map[string]any {
	hx_obj_198 := map[string]any{}
	hx_obj_198["values"] = []int{1, 2, 1}
	return hx_obj_198
}

func makeSame() *string {
	return hxrt.StringFromLiteral("same")
}

func markedInsertValue() int {
	hx_arr_199 := events
	hx_arr_199 = append(hx_arr_199, hxrt.StringFromLiteral("value"))
	events = hx_arr_199
	return 2
}

func markedPosition() int {
	hx_arr_200 := events
	hx_arr_200 = append(hx_arr_200, hxrt.StringFromLiteral("position"))
	events = hx_arr_200
	return -1
}

func markedRemoveValue() int {
	hx_arr_201 := events
	hx_arr_201 = append(hx_arr_201, hxrt.StringFromLiteral("value"))
	events = hx_arr_201
	return 1
}

func removeGeneric(first any, second any, value any) *string {
	values := []any{first, second}
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StdString(func() bool {
		var hx_remove_value_203 any = value
		for hx_remove_index_204, hx_remove_element_205 := range values {
			if hxrt.HaxeEqual(hx_remove_element_205, hx_remove_value_203) {
				hx_remove_last_206 := (len(values) - 1)
				copy(values[hx_remove_index_204:], values[(hx_remove_index_204+1):])
				var hx_remove_zero_207 any
				values[hx_remove_last_206] = hx_remove_zero_207
				values = values[:hx_remove_last_206]
				return true
			}
		}
		return false
	}()), hxrt.StringFromLiteral(":")), len(values)), hxrt.StringFromLiteral(":")), hxrt.StdString(values[0]))
}

func removeGenericCount(first any, second any, value any) *string {
	values := []any{first, second}
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StdString(func() bool {
		var hx_remove_value_209 any = value
		for hx_remove_index_210, hx_remove_element_211 := range values {
			if hxrt.HaxeEqual(hx_remove_element_211, hx_remove_value_209) {
				hx_remove_last_212 := (len(values) - 1)
				copy(values[hx_remove_index_210:], values[(hx_remove_index_210+1):])
				var hx_remove_zero_213 any
				values[hx_remove_last_212] = hx_remove_zero_213
				values = values[:hx_remove_last_212]
				return true
			}
		}
		return false
	}()), hxrt.StringFromLiteral(":")), len(values))
}

func removeGenericFour(first any, second any, third any, fourth any, value any) *string {
	values := []any{first, second, third, fourth}
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StdString(func() bool {
		var hx_remove_value_215 any = value
		for hx_remove_index_216, hx_remove_element_217 := range values {
			if hxrt.HaxeEqual(hx_remove_element_217, hx_remove_value_215) {
				hx_remove_last_218 := (len(values) - 1)
				copy(values[hx_remove_index_216:], values[(hx_remove_index_216+1):])
				var hx_remove_zero_219 any
				values[hx_remove_last_218] = hx_remove_zero_219
				values = values[:hx_remove_last_218]
				return true
			}
		}
		return false
	}()), hxrt.StringFromLiteral(":")), hxrt.StringJoinAny(func(hx_sort_src_221 []*string) []any {
		hx_sort_out_223 := make([]any, 0, len(hx_sort_src_221))
		for _, hx_sort_item_222 := range hx_sort_src_221 {
			hx_sort_out_223 = append(hx_sort_out_223, hx_sort_item_222)
		}
		return hx_sort_out_223
	}(func() []*string {
		_g := []*string{}
		_g1 := 0
		for _g1 < len(values) {
			var item any = values[_g1]
			_g1 = int(int32((_g1 + 1)))
			_g = append(_g, hxrt.StdString(item))
		}
		return _g
	}()), hxrt.StringFromLiteral(",")))
}

func showNullableInts(values []any) *string {
	return hxrt.StringJoinAny(func(hx_sort_src_225 []*string) []any {
		hx_sort_out_227 := make([]any, 0, len(hx_sort_src_225))
		for _, hx_sort_item_226 := range hx_sort_src_225 {
			hx_sort_out_227 = append(hx_sort_out_227, hx_sort_item_226)
		}
		return hx_sort_out_227
	}(func() []*string {
		_g := []*string{}
		_g1 := 0
		for _g1 < len(values) {
			var value any = values[_g1]
			_g1 = int(int32((_g1 + 1)))
			_g = append(_g, hxrt.StdString(value))
		}
		return _g
	}()), hxrt.StringFromLiteral(","))
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
