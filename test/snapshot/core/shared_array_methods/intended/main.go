package main

import "snapshot/hxrt"

func main() {
	ints := makeInts()
	names := makeNames()
	first := New_SharedArrayItem(hxrt.StringFromLiteral("a"))
	second := New_SharedArrayItem(hxrt.StringFromLiteral("b"))
	items := hxrt.NewArray(first, second, nil)
	holder := New_SharedArrayHolder(makeInts(), makeNames(), hxrt.NewArray(second, first))
	rows := makeRows()
	var grid_1 *hxrt.Array
	var grid_0 *hxrt.Array
	grid_0 = makeNames()
	grid_1 = makeNames()
	shelves := hxrt.NewArray(hxrt.NewArray(first), hxrt.NewArray(second, first))
	var v any = any(hxrt.StdString(ints))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StdString(names))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StdString(items))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StdString(holder.ints), hxrt.StringFromLiteral(" ")), hxrt.StdString(holder.items)))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StdString(rows))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StdString(func(hx_value_1 any) *hxrt.Array {
		if hx_value_1 == nil {
			var hx_zero_2 *hxrt.Array
			return hx_zero_2
		}
		return hx_value_1.(*hxrt.Array)
	}(rows.Get(1))), hxrt.StringFromLiteral(" ")), hxrt.StdString(grid_0)), hxrt.StringFromLiteral(" ")), hxrt.StdString(func(hx_value_3 any) *hxrt.Array {
		if hx_value_3 == nil {
			var hx_zero_4 *hxrt.Array
			return hx_zero_4
		}
		return hx_value_3.(*hxrt.Array)
	}(shelves.Get(1)))))
	hxrt.Println(v_5)
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StdString(func() bool {
		hx_contains_target_5 := ints
		var hx_contains_value_6 int = 3
		hx_contains_length_7 := hx_contains_target_5.Len()
		hx_contains_index_8 := 0
		for hx_contains_index_8 < hx_contains_length_7 {
			hx_contains_element_9 := hxrt.IntFromNullableAny(hx_contains_target_5.Get(hx_contains_index_8))
			if hx_contains_element_9 == hx_contains_value_6 {
				return true
			}
			hx_contains_index_8 = (hx_contains_index_8 + 1)
		}
		return false
	}()), hxrt.StringFromLiteral(" ")), hxrt.StdString(func() bool {
		hx_contains_target_10 := ints
		var hx_contains_value_11 int = 7
		hx_contains_length_12 := hx_contains_target_10.Len()
		hx_contains_index_13 := 0
		for hx_contains_index_13 < hx_contains_length_12 {
			hx_contains_element_14 := hxrt.IntFromNullableAny(hx_contains_target_10.Get(hx_contains_index_13))
			if hx_contains_element_14 == hx_contains_value_11 {
				return true
			}
			hx_contains_index_13 = (hx_contains_index_13 + 1)
		}
		return false
	}())))
	hxrt.Println(v_6)
	wanted := hxrt.StringFromLiteral("b")
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StdString(func() bool {
		hx_contains_target_15 := names
		var hx_contains_value_16 *string = hxrt.StringConcatStringPtr(wanted, hxrt.StringFromLiteral(""))
		hx_contains_length_17 := hx_contains_target_15.Len()
		hx_contains_index_18 := 0
		for hx_contains_index_18 < hx_contains_length_17 {
			hx_contains_element_19 := func(hx_value_20 any) *string {
				if hx_value_20 == nil {
					var hx_zero_21 *string
					return hx_zero_21
				}
				return hx_value_20.(*string)
			}(hx_contains_target_15.Get(hx_contains_index_18))
			if hxrt.StringEqualStringPtr(hx_contains_element_19, hx_contains_value_16) {
				return true
			}
			hx_contains_index_18 = (hx_contains_index_18 + 1)
		}
		return false
	}()), hxrt.StringFromLiteral(" ")), hxrt.StdString(func() bool {
		hx_contains_target_22 := names
		var hx_contains_value_23 *string = nil
		hx_contains_length_24 := hx_contains_target_22.Len()
		hx_contains_index_25 := 0
		for hx_contains_index_25 < hx_contains_length_24 {
			hx_contains_element_26 := func(hx_value_27 any) *string {
				if hx_value_27 == nil {
					var hx_zero_28 *string
					return hx_zero_28
				}
				return hx_value_27.(*string)
			}(hx_contains_target_22.Get(hx_contains_index_25))
			if hxrt.StringEqualStringPtr(hx_contains_element_26, hx_contains_value_23) {
				return true
			}
			hx_contains_index_25 = (hx_contains_index_25 + 1)
		}
		return false
	}())), hxrt.StringFromLiteral(" ")), hxrt.StdString(func() bool {
		hx_contains_target_29 := names
		var hx_contains_value_30 *string = hxrt.StringFromLiteral("z")
		hx_contains_length_31 := hx_contains_target_29.Len()
		hx_contains_index_32 := 0
		for hx_contains_index_32 < hx_contains_length_31 {
			hx_contains_element_33 := func(hx_value_34 any) *string {
				if hx_value_34 == nil {
					var hx_zero_35 *string
					return hx_zero_35
				}
				return hx_value_34.(*string)
			}(hx_contains_target_29.Get(hx_contains_index_32))
			if hxrt.StringEqualStringPtr(hx_contains_element_33, hx_contains_value_30) {
				return true
			}
			hx_contains_index_32 = (hx_contains_index_32 + 1)
		}
		return false
	}())))
	hxrt.Println(v_7)
	var v_8 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StdString(func() bool {
		hx_contains_target_36 := items
		var hx_contains_value_37 *SharedArrayItem = first
		hx_contains_length_38 := hx_contains_target_36.Len()
		hx_contains_index_39 := 0
		for hx_contains_index_39 < hx_contains_length_38 {
			hx_contains_element_40 := hx_recover_class_SharedArrayItem(hx_contains_target_36.Get(hx_contains_index_39))
			if hx_contains_element_40 == hx_contains_value_37 {
				return true
			}
			hx_contains_index_39 = (hx_contains_index_39 + 1)
		}
		return false
	}()), hxrt.StringFromLiteral(" ")), hxrt.StdString(func() bool {
		hx_contains_target_41 := items
		var hx_contains_value_42 *SharedArrayItem = New_SharedArrayItem(hxrt.StringFromLiteral("a"))
		hx_contains_length_43 := hx_contains_target_41.Len()
		hx_contains_index_44 := 0
		for hx_contains_index_44 < hx_contains_length_43 {
			hx_contains_element_45 := hx_recover_class_SharedArrayItem(hx_contains_target_41.Get(hx_contains_index_44))
			if hx_contains_element_45 == hx_contains_value_42 {
				return true
			}
			hx_contains_index_44 = (hx_contains_index_44 + 1)
		}
		return false
	}())), hxrt.StringFromLiteral(" ")), hxrt.StdString(func() bool {
		hx_contains_target_46 := items
		var hx_contains_value_47 *SharedArrayItem = nil
		hx_contains_length_48 := hx_contains_target_46.Len()
		hx_contains_index_49 := 0
		for hx_contains_index_49 < hx_contains_length_48 {
			hx_contains_element_50 := hx_recover_class_SharedArrayItem(hx_contains_target_46.Get(hx_contains_index_49))
			if hx_contains_element_50 == hx_contains_value_47 {
				return true
			}
			hx_contains_index_49 = (hx_contains_index_49 + 1)
		}
		return false
	}())))
	hxrt.Println(v_8)
	var v_9 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StdString(func() bool {
		hx_contains_target_51 := holder.names
		var hx_contains_value_52 *string = hxrt.StringFromLiteral("a")
		hx_contains_length_53 := hx_contains_target_51.Len()
		hx_contains_index_54 := 0
		for hx_contains_index_54 < hx_contains_length_53 {
			hx_contains_element_55 := func(hx_value_56 any) *string {
				if hx_value_56 == nil {
					var hx_zero_57 *string
					return hx_zero_57
				}
				return hx_value_56.(*string)
			}(hx_contains_target_51.Get(hx_contains_index_54))
			if hxrt.StringEqualStringPtr(hx_contains_element_55, hx_contains_value_52) {
				return true
			}
			hx_contains_index_54 = (hx_contains_index_54 + 1)
		}
		return false
	}()), hxrt.StringFromLiteral(" ")), hxrt.StdString(func() bool {
		hx_contains_target_58 := holder.items
		var hx_contains_value_59 *SharedArrayItem = second
		hx_contains_length_60 := hx_contains_target_58.Len()
		hx_contains_index_61 := 0
		for hx_contains_index_61 < hx_contains_length_60 {
			hx_contains_element_62 := hx_recover_class_SharedArrayItem(hx_contains_target_58.Get(hx_contains_index_61))
			if hx_contains_element_62 == hx_contains_value_59 {
				return true
			}
			hx_contains_index_61 = (hx_contains_index_61 + 1)
		}
		return false
	}())))
	hxrt.Println(v_9)
	var v_10 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StdString(func() bool {
		hx_contains_target_63 := func(hx_value_68 any) *hxrt.Array {
			if hx_value_68 == nil {
				var hx_zero_69 *hxrt.Array
				return hx_zero_69
			}
			return hx_value_68.(*hxrt.Array)
		}(rows.Get(2))
		var hx_contains_value_64 int = 7
		hx_contains_length_65 := hx_contains_target_63.Len()
		hx_contains_index_66 := 0
		for hx_contains_index_66 < hx_contains_length_65 {
			hx_contains_element_67 := hxrt.IntFromNullableAny(hx_contains_target_63.Get(hx_contains_index_66))
			if hx_contains_element_67 == hx_contains_value_64 {
				return true
			}
			hx_contains_index_66 = (hx_contains_index_66 + 1)
		}
		return false
	}()), hxrt.StringFromLiteral(" ")), hxrt.StdString(func() bool {
		hx_contains_target_70 := grid_1
		var hx_contains_value_71 *string = nil
		hx_contains_length_72 := hx_contains_target_70.Len()
		hx_contains_index_73 := 0
		for hx_contains_index_73 < hx_contains_length_72 {
			hx_contains_element_74 := func(hx_value_75 any) *string {
				if hx_value_75 == nil {
					var hx_zero_76 *string
					return hx_zero_76
				}
				return hx_value_75.(*string)
			}(hx_contains_target_70.Get(hx_contains_index_73))
			if hxrt.StringEqualStringPtr(hx_contains_element_74, hx_contains_value_71) {
				return true
			}
			hx_contains_index_73 = (hx_contains_index_73 + 1)
		}
		return false
	}())), hxrt.StringFromLiteral(" ")), hxrt.StdString(func() bool {
		hx_contains_target_77 := func(hx_value_82 any) *hxrt.Array {
			if hx_value_82 == nil {
				var hx_zero_83 *hxrt.Array
				return hx_zero_83
			}
			return hx_value_82.(*hxrt.Array)
		}(shelves.Get(0))
		var hx_contains_value_78 *SharedArrayItem = second
		hx_contains_length_79 := hx_contains_target_77.Len()
		hx_contains_index_80 := 0
		for hx_contains_index_80 < hx_contains_length_79 {
			hx_contains_element_81 := hx_recover_class_SharedArrayItem(hx_contains_target_77.Get(hx_contains_index_80))
			if hx_contains_element_81 == hx_contains_value_78 {
				return true
			}
			hx_contains_index_80 = (hx_contains_index_80 + 1)
		}
		return false
	}())))
	hxrt.Println(v_10)
	var v_11 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatAny(func() int {
		hx_lastindexof_target_84 := ints
		var hx_lastindexof_value_85 int = 2
		hx_lastindexof_length_86 := hx_lastindexof_target_84.Len()
		hx_lastindexof_index_87 := (hx_lastindexof_length_86 - 1)
		for hx_lastindexof_index_87 >= 0 {
			hx_lastindexof_element_88 := hxrt.IntFromNullableAny(hx_lastindexof_target_84.Get(hx_lastindexof_index_87))
			if hx_lastindexof_element_88 == hx_lastindexof_value_85 {
				return hx_lastindexof_index_87
			}
			hx_lastindexof_index_87 = (hx_lastindexof_index_87 - 1)
		}
		return -1
	}(), hxrt.StringFromLiteral(" ")), func() int {
		hx_lastindexof_target_89 := ints
		var hx_lastindexof_value_90 int = 2
		hx_lastindexof_index_92 := 2
		hx_lastindexof_length_91 := hx_lastindexof_target_89.Len()
		if hx_lastindexof_index_92 < 0 {
			hx_lastindexof_index_92 = (hx_lastindexof_length_91 + hx_lastindexof_index_92)
		}
		if hx_lastindexof_index_92 >= hx_lastindexof_length_91 {
			hx_lastindexof_index_92 = (hx_lastindexof_length_91 - 1)
		}
		for hx_lastindexof_index_92 >= 0 {
			hx_lastindexof_element_93 := hxrt.IntFromNullableAny(hx_lastindexof_target_89.Get(hx_lastindexof_index_92))
			if hx_lastindexof_element_93 == hx_lastindexof_value_90 {
				return hx_lastindexof_index_92
			}
			hx_lastindexof_index_92 = (hx_lastindexof_index_92 - 1)
		}
		return -1
	}()), hxrt.StringFromLiteral(" ")), func() int {
		hx_lastindexof_target_94 := ints
		var hx_lastindexof_value_95 int = 2
		hx_lastindexof_index_97 := -3
		hx_lastindexof_length_96 := hx_lastindexof_target_94.Len()
		if hx_lastindexof_index_97 < 0 {
			hx_lastindexof_index_97 = (hx_lastindexof_length_96 + hx_lastindexof_index_97)
		}
		if hx_lastindexof_index_97 >= hx_lastindexof_length_96 {
			hx_lastindexof_index_97 = (hx_lastindexof_length_96 - 1)
		}
		for hx_lastindexof_index_97 >= 0 {
			hx_lastindexof_element_98 := hxrt.IntFromNullableAny(hx_lastindexof_target_94.Get(hx_lastindexof_index_97))
			if hx_lastindexof_element_98 == hx_lastindexof_value_95 {
				return hx_lastindexof_index_97
			}
			hx_lastindexof_index_97 = (hx_lastindexof_index_97 - 1)
		}
		return -1
	}()), hxrt.StringFromLiteral(" ")), func() int {
		hx_lastindexof_target_99 := ints
		var hx_lastindexof_value_100 int = 9
		hx_lastindexof_length_101 := hx_lastindexof_target_99.Len()
		hx_lastindexof_index_102 := (hx_lastindexof_length_101 - 1)
		for hx_lastindexof_index_102 >= 0 {
			hx_lastindexof_element_103 := hxrt.IntFromNullableAny(hx_lastindexof_target_99.Get(hx_lastindexof_index_102))
			if hx_lastindexof_element_103 == hx_lastindexof_value_100 {
				return hx_lastindexof_index_102
			}
			hx_lastindexof_index_102 = (hx_lastindexof_index_102 - 1)
		}
		return -1
	}()))
	hxrt.Println(v_11)
	var v_12 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatAny(func() int {
		hx_lastindexof_target_104 := names
		var hx_lastindexof_value_105 *string = hxrt.StringFromLiteral("b")
		hx_lastindexof_length_106 := hx_lastindexof_target_104.Len()
		hx_lastindexof_index_107 := (hx_lastindexof_length_106 - 1)
		for hx_lastindexof_index_107 >= 0 {
			hx_lastindexof_element_108 := func(hx_value_109 any) *string {
				if hx_value_109 == nil {
					var hx_zero_110 *string
					return hx_zero_110
				}
				return hx_value_109.(*string)
			}(hx_lastindexof_target_104.Get(hx_lastindexof_index_107))
			if hxrt.StringEqualStringPtr(hx_lastindexof_element_108, hx_lastindexof_value_105) {
				return hx_lastindexof_index_107
			}
			hx_lastindexof_index_107 = (hx_lastindexof_index_107 - 1)
		}
		return -1
	}(), hxrt.StringFromLiteral(" ")), func() int {
		hx_lastindexof_target_111 := names
		var hx_lastindexof_value_112 *string = nil
		hx_lastindexof_length_113 := hx_lastindexof_target_111.Len()
		hx_lastindexof_index_114 := (hx_lastindexof_length_113 - 1)
		for hx_lastindexof_index_114 >= 0 {
			hx_lastindexof_element_115 := func(hx_value_116 any) *string {
				if hx_value_116 == nil {
					var hx_zero_117 *string
					return hx_zero_117
				}
				return hx_value_116.(*string)
			}(hx_lastindexof_target_111.Get(hx_lastindexof_index_114))
			if hxrt.StringEqualStringPtr(hx_lastindexof_element_115, hx_lastindexof_value_112) {
				return hx_lastindexof_index_114
			}
			hx_lastindexof_index_114 = (hx_lastindexof_index_114 - 1)
		}
		return -1
	}()), hxrt.StringFromLiteral(" ")), func() int {
		hx_lastindexof_target_118 := items
		var hx_lastindexof_value_119 *SharedArrayItem = second
		hx_lastindexof_length_120 := hx_lastindexof_target_118.Len()
		hx_lastindexof_index_121 := (hx_lastindexof_length_120 - 1)
		for hx_lastindexof_index_121 >= 0 {
			hx_lastindexof_element_122 := hx_recover_class_SharedArrayItem(hx_lastindexof_target_118.Get(hx_lastindexof_index_121))
			if hx_lastindexof_element_122 == hx_lastindexof_value_119 {
				return hx_lastindexof_index_121
			}
			hx_lastindexof_index_121 = (hx_lastindexof_index_121 - 1)
		}
		return -1
	}()))
	hxrt.Println(v_12)
	var v_13 any = any(hxrt.StringConcatAny(hxrt.StringConcatAny(func() int {
		hx_lastindexof_target_123 := holder.ints
		var hx_lastindexof_value_124 int = 1
		hx_lastindexof_index_126 := 10
		hx_lastindexof_length_125 := hx_lastindexof_target_123.Len()
		if hx_lastindexof_index_126 < 0 {
			hx_lastindexof_index_126 = (hx_lastindexof_length_125 + hx_lastindexof_index_126)
		}
		if hx_lastindexof_index_126 >= hx_lastindexof_length_125 {
			hx_lastindexof_index_126 = (hx_lastindexof_length_125 - 1)
		}
		for hx_lastindexof_index_126 >= 0 {
			hx_lastindexof_element_127 := hxrt.IntFromNullableAny(hx_lastindexof_target_123.Get(hx_lastindexof_index_126))
			if hx_lastindexof_element_127 == hx_lastindexof_value_124 {
				return hx_lastindexof_index_126
			}
			hx_lastindexof_index_126 = (hx_lastindexof_index_126 - 1)
		}
		return -1
	}(), hxrt.StringFromLiteral(" ")), func() int {
		hx_lastindexof_target_128 := holder.items
		var hx_lastindexof_value_129 *SharedArrayItem = first
		hx_lastindexof_index_131 := 0
		hx_lastindexof_length_130 := hx_lastindexof_target_128.Len()
		if hx_lastindexof_index_131 < 0 {
			hx_lastindexof_index_131 = (hx_lastindexof_length_130 + hx_lastindexof_index_131)
		}
		if hx_lastindexof_index_131 >= hx_lastindexof_length_130 {
			hx_lastindexof_index_131 = (hx_lastindexof_length_130 - 1)
		}
		for hx_lastindexof_index_131 >= 0 {
			hx_lastindexof_element_132 := hx_recover_class_SharedArrayItem(hx_lastindexof_target_128.Get(hx_lastindexof_index_131))
			if hx_lastindexof_element_132 == hx_lastindexof_value_129 {
				return hx_lastindexof_index_131
			}
			hx_lastindexof_index_131 = (hx_lastindexof_index_131 - 1)
		}
		return -1
	}()))
	hxrt.Println(v_13)
	var v_14 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatAny(func() int {
		hx_lastindexof_target_133 := func(hx_value_138 any) *hxrt.Array {
			if hx_value_138 == nil {
				var hx_zero_139 *hxrt.Array
				return hx_zero_139
			}
			return hx_value_138.(*hxrt.Array)
		}(rows.Get(0))
		var hx_lastindexof_value_134 int = 2
		hx_lastindexof_length_135 := hx_lastindexof_target_133.Len()
		hx_lastindexof_index_136 := (hx_lastindexof_length_135 - 1)
		for hx_lastindexof_index_136 >= 0 {
			hx_lastindexof_element_137 := hxrt.IntFromNullableAny(hx_lastindexof_target_133.Get(hx_lastindexof_index_136))
			if hx_lastindexof_element_137 == hx_lastindexof_value_134 {
				return hx_lastindexof_index_136
			}
			hx_lastindexof_index_136 = (hx_lastindexof_index_136 - 1)
		}
		return -1
	}(), hxrt.StringFromLiteral(" ")), func() int {
		hx_lastindexof_target_140 := grid_1
		var hx_lastindexof_value_141 *string = hxrt.StringFromLiteral("b")
		hx_lastindexof_index_143 := 2
		hx_lastindexof_length_142 := hx_lastindexof_target_140.Len()
		if hx_lastindexof_index_143 < 0 {
			hx_lastindexof_index_143 = (hx_lastindexof_length_142 + hx_lastindexof_index_143)
		}
		if hx_lastindexof_index_143 >= hx_lastindexof_length_142 {
			hx_lastindexof_index_143 = (hx_lastindexof_length_142 - 1)
		}
		for hx_lastindexof_index_143 >= 0 {
			hx_lastindexof_element_144 := func(hx_value_145 any) *string {
				if hx_value_145 == nil {
					var hx_zero_146 *string
					return hx_zero_146
				}
				return hx_value_145.(*string)
			}(hx_lastindexof_target_140.Get(hx_lastindexof_index_143))
			if hxrt.StringEqualStringPtr(hx_lastindexof_element_144, hx_lastindexof_value_141) {
				return hx_lastindexof_index_143
			}
			hx_lastindexof_index_143 = (hx_lastindexof_index_143 - 1)
		}
		return -1
	}()), hxrt.StringFromLiteral(" ")), func() int {
		hx_lastindexof_target_147 := func(hx_value_152 any) *hxrt.Array {
			if hx_value_152 == nil {
				var hx_zero_153 *hxrt.Array
				return hx_zero_153
			}
			return hx_value_152.(*hxrt.Array)
		}(shelves.Get(1))
		var hx_lastindexof_value_148 *SharedArrayItem = first
		hx_lastindexof_length_149 := hx_lastindexof_target_147.Len()
		hx_lastindexof_index_150 := (hx_lastindexof_length_149 - 1)
		for hx_lastindexof_index_150 >= 0 {
			hx_lastindexof_element_151 := hx_recover_class_SharedArrayItem(hx_lastindexof_target_147.Get(hx_lastindexof_index_150))
			if hx_lastindexof_element_151 == hx_lastindexof_value_148 {
				return hx_lastindexof_index_150
			}
			hx_lastindexof_index_150 = (hx_lastindexof_index_150 - 1)
		}
		return -1
	}()))
	hxrt.Println(v_14)
	maybe := makeMaybe()
	var v_15 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatAny(func() int {
		hx_indexof_target_154 := maybe
		var hx_indexof_value_155 any = nil
		hx_indexof_length_156 := hx_indexof_target_154.Len()
		hx_indexof_index_157 := 0
		for hx_indexof_index_157 < hx_indexof_length_156 {
			hx_indexof_element_158 := hx_indexof_target_154.Get(hx_indexof_index_157)
			if hxrt.HaxeEqual(hx_indexof_element_158, hx_indexof_value_155) {
				return hx_indexof_index_157
			}
			hx_indexof_index_157 = (hx_indexof_index_157 + 1)
		}
		return -1
	}(), hxrt.StringFromLiteral(" ")), func() int {
		hx_lastindexof_target_159 := maybe
		var hx_lastindexof_value_160 any = nil
		hx_lastindexof_length_161 := hx_lastindexof_target_159.Len()
		hx_lastindexof_index_162 := (hx_lastindexof_length_161 - 1)
		for hx_lastindexof_index_162 >= 0 {
			hx_lastindexof_element_163 := hx_lastindexof_target_159.Get(hx_lastindexof_index_162)
			if hxrt.HaxeEqual(hx_lastindexof_element_163, hx_lastindexof_value_160) {
				return hx_lastindexof_index_162
			}
			hx_lastindexof_index_162 = (hx_lastindexof_index_162 - 1)
		}
		return -1
	}()), hxrt.StringFromLiteral(" ")), hxrt.StdString(func() bool {
		hx_contains_target_164 := maybe
		var hx_contains_value_165 any = 1
		hx_contains_length_166 := hx_contains_target_164.Len()
		hx_contains_index_167 := 0
		for hx_contains_index_167 < hx_contains_length_166 {
			hx_contains_element_168 := hx_contains_target_164.Get(hx_contains_index_167)
			if hxrt.HaxeEqual(hx_contains_element_168, hx_contains_value_165) {
				return true
			}
			hx_contains_index_167 = (hx_contains_index_167 + 1)
		}
		return false
	}())), hxrt.StringFromLiteral(" ")), func() int {
		hx_indexof_target_169 := names
		var hx_indexof_value_170 *string = nil
		hx_indexof_length_171 := hx_indexof_target_169.Len()
		hx_indexof_index_172 := 0
		for hx_indexof_index_172 < hx_indexof_length_171 {
			hx_indexof_element_173 := func(hx_value_174 any) *string {
				if hx_value_174 == nil {
					var hx_zero_175 *string
					return hx_zero_175
				}
				return hx_value_174.(*string)
			}(hx_indexof_target_169.Get(hx_indexof_index_172))
			if hxrt.StringEqualStringPtr(hx_indexof_element_173, hx_indexof_value_170) {
				return hx_indexof_index_172
			}
			hx_indexof_index_172 = (hx_indexof_index_172 + 1)
		}
		return -1
	}()))
	hxrt.Println(v_15)
	var v_16 any = any(hxrt.StringConcatAny(hxrt.StringConcatAny(func() int {
		hx_indexof_target_176 := ints
		var hx_indexof_value_177 int = 1
		var hx_indexof_start_181 any = noStart()
		hx_indexof_length_178 := hx_indexof_target_176.Len()
		hx_indexof_index_179 := 0
		if hx_indexof_start_181 != nil {
			hx_indexof_index_179 = hxrt.IntFromNullableAny(hx_indexof_start_181)
			if hx_indexof_index_179 < 0 {
				hx_indexof_index_179 = (hx_indexof_length_178 + hx_indexof_index_179)
			}
			if hx_indexof_index_179 < 0 {
				hx_indexof_index_179 = 0
			}
		}
		for hx_indexof_index_179 < hx_indexof_length_178 {
			hx_indexof_element_180 := hxrt.IntFromNullableAny(hx_indexof_target_176.Get(hx_indexof_index_179))
			if hx_indexof_element_180 == hx_indexof_value_177 {
				return hx_indexof_index_179
			}
			hx_indexof_index_179 = (hx_indexof_index_179 + 1)
		}
		return -1
	}(), hxrt.StringFromLiteral(" ")), func() int {
		hx_lastindexof_target_182 := ints
		var hx_lastindexof_value_183 int = 1
		var hx_lastindexof_start_187 any = noStart()
		hx_lastindexof_length_184 := hx_lastindexof_target_182.Len()
		hx_lastindexof_index_185 := (hx_lastindexof_length_184 - 1)
		if hx_lastindexof_start_187 != nil {
			hx_lastindexof_index_185 = hxrt.IntFromNullableAny(hx_lastindexof_start_187)
			if hx_lastindexof_index_185 < 0 {
				hx_lastindexof_index_185 = (hx_lastindexof_length_184 + hx_lastindexof_index_185)
			}
			if hx_lastindexof_index_185 >= hx_lastindexof_length_184 {
				hx_lastindexof_index_185 = (hx_lastindexof_length_184 - 1)
			}
		}
		for hx_lastindexof_index_185 >= 0 {
			hx_lastindexof_element_186 := hxrt.IntFromNullableAny(hx_lastindexof_target_182.Get(hx_lastindexof_index_185))
			if hx_lastindexof_element_186 == hx_lastindexof_value_183 {
				return hx_lastindexof_index_185
			}
			hx_lastindexof_index_185 = (hx_lastindexof_index_185 - 1)
		}
		return -1
	}()))
	hxrt.Println(v_16)
	var v_17 any = any(func() int {
		hx_lastindexof_target_188 := holder.ints
		var hx_lastindexof_value_189 int = 1
		hx_lastindexof_index_191 := -6
		hx_lastindexof_length_190 := hx_lastindexof_target_188.Len()
		if hx_lastindexof_index_191 < 0 {
			hx_lastindexof_index_191 = (hx_lastindexof_length_190 + hx_lastindexof_index_191)
		}
		if hx_lastindexof_index_191 >= hx_lastindexof_length_190 {
			hx_lastindexof_index_191 = (hx_lastindexof_length_190 - 1)
		}
		for hx_lastindexof_index_191 >= 0 {
			hx_lastindexof_element_192 := hxrt.IntFromNullableAny(hx_lastindexof_target_188.Get(hx_lastindexof_index_191))
			if hx_lastindexof_element_192 == hx_lastindexof_value_189 {
				return hx_lastindexof_index_191
			}
			hx_lastindexof_index_191 = (hx_lastindexof_index_191 - 1)
		}
		return -1
	}())
	hxrt.Println(v_17)
	ints.Reverse()
	names.Reverse()
	items.Reverse()
	holder.ints.Reverse()
	func(hx_value_193 any) *hxrt.Array {
		if hx_value_193 == nil {
			var hx_zero_194 *hxrt.Array
			return hx_zero_194
		}
		return hx_value_193.(*hxrt.Array)
	}(rows.Get(1)).Reverse()
	func(hx_value_195 any) *hxrt.Array {
		if hx_value_195 == nil {
			var hx_zero_196 *hxrt.Array
			return hx_zero_196
		}
		return hx_value_195.(*hxrt.Array)
	}(shelves.Get(1)).Reverse()
	var v_18 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StdString(ints), hxrt.StringFromLiteral(" ")), hxrt.StdString(names)), hxrt.StringFromLiteral(" ")), hxrt.StdString(items)))
	hxrt.Println(v_18)
	var v_19 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StdString(holder.ints), hxrt.StringFromLiteral(" ")), hxrt.StdString(func(hx_value_197 any) *hxrt.Array {
		if hx_value_197 == nil {
			var hx_zero_198 *hxrt.Array
			return hx_zero_198
		}
		return hx_value_197.(*hxrt.Array)
	}(rows.Get(1)))), hxrt.StringFromLiteral(" ")), hxrt.StdString(func(hx_value_199 any) *hxrt.Array {
		if hx_value_199 == nil {
			var hx_zero_200 *hxrt.Array
			return hx_zero_200
		}
		return hx_value_199.(*hxrt.Array)
	}(shelves.Get(1)))))
	hxrt.Println(v_19)
	ints.Insert(0, 0)
	names.Insert(0, nil)
	items.Insert(0, second)
	func() {
		hx_arr_204 := holder.names
		var hx_insert_value_205 any = hxrt.StringFromLiteral("first")
		hx_arr_204.Insert(0, hx_insert_value_205)
	}()
	func() {
		hx_arr_206 := func(hx_value_207 any) *hxrt.Array {
			if hx_value_207 == nil {
				var hx_zero_208 *hxrt.Array
				return hx_zero_208
			}
			return hx_value_207.(*hxrt.Array)
		}(rows.Get(2))
		var hx_insert_value_209 any = -1
		hx_arr_206.Insert(0, hx_insert_value_209)
	}()
	grid_0.Insert(0, hxrt.StringFromLiteral("z"))
	func() {
		hx_arr_211 := func(hx_value_212 any) *hxrt.Array {
			if hx_value_212 == nil {
				var hx_zero_213 *hxrt.Array
				return hx_zero_213
			}
			return hx_value_212.(*hxrt.Array)
		}(shelves.Get(0))
		var hx_insert_value_214 any = second
		hx_arr_211.Insert(0, hx_insert_value_214)
	}()
	var v_20 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StdString(ints), hxrt.StringFromLiteral(" ")), hxrt.StdString(names)), hxrt.StringFromLiteral(" ")), hxrt.StdString(items)))
	hxrt.Println(v_20)
	var v_21 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StdString(holder.names), hxrt.StringFromLiteral(" ")), hxrt.StdString(func(hx_value_215 any) *hxrt.Array {
		if hx_value_215 == nil {
			var hx_zero_216 *hxrt.Array
			return hx_zero_216
		}
		return hx_value_215.(*hxrt.Array)
	}(rows.Get(2)))), hxrt.StringFromLiteral(" ")), hxrt.StdString(grid_0)), hxrt.StringFromLiteral(" ")), hxrt.StdString(func(hx_value_217 any) *hxrt.Array {
		if hx_value_217 == nil {
			var hx_zero_218 *hxrt.Array
			return hx_zero_218
		}
		return hx_value_217.(*hxrt.Array)
	}(shelves.Get(0)))))
	hxrt.Println(v_21)
	removed := ints.Splice(1, 2)
	var v_22 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StdString(removed), hxrt.StringFromLiteral(" ")), hxrt.StdString(ints)))
	hxrt.Println(v_22)
	var v_23 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StdString(names.Splice(-2, 1)), hxrt.StringFromLiteral(" ")), hxrt.StdString(names)))
	hxrt.Println(v_23)
	var v_24 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StdString(items.Splice(0, 1)), hxrt.StringFromLiteral(" ")), hxrt.StdString(items)))
	hxrt.Println(v_24)
	var v_25 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StdString(holder.ints.Splice(3, 100)), hxrt.StringFromLiteral(" ")), hxrt.StdString(holder.ints)))
	hxrt.Println(v_25)
	var v_26 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StdString(func(hx_value_219 any) *hxrt.Array {
		if hx_value_219 == nil {
			var hx_zero_220 *hxrt.Array
			return hx_zero_220
		}
		return hx_value_219.(*hxrt.Array)
	}(rows.Get(0)).Splice(1, 1)), hxrt.StringFromLiteral(" ")), hxrt.StdString(rows)))
	hxrt.Println(v_26)
	var v_27 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(grid_1.Splice(-10, 2).Len(), hxrt.StringFromLiteral(" ")), hxrt.StdString(grid_1)))
	hxrt.Println(v_27)
	var v_28 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StdString(func(hx_value_221 any) *hxrt.Array {
		if hx_value_221 == nil {
			var hx_zero_222 *hxrt.Array
			return hx_zero_222
		}
		return hx_value_221.(*hxrt.Array)
	}(shelves.Get(1)).Splice(1, 5)), hxrt.StringFromLiteral(" ")), hxrt.StdString(shelves)))
	hxrt.Println(v_28)
	sized := makeInts()
	sized.SetLength(2)
	func(hx_value_223 any) *hxrt.Array {
		if hx_value_223 == nil {
			var hx_zero_224 *hxrt.Array
			return hx_zero_224
		}
		return hx_value_223.(*hxrt.Array)
	}(rows.Get(2)).SetLength(6)
	holder.names.SetLength(1)
	var v_29 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StdString(sized), hxrt.StringFromLiteral(" ")), hxrt.StdString(func(hx_value_225 any) *hxrt.Array {
		if hx_value_225 == nil {
			var hx_zero_226 *hxrt.Array
			return hx_zero_226
		}
		return hx_value_225.(*hxrt.Array)
	}(rows.Get(2)))), hxrt.StringFromLiteral(" ")), hxrt.StdString(holder.names)))
	hxrt.Println(v_29)
}

func makeInts() *hxrt.Array {
	return hxrt.NewArray(1, 2, 3, 2, 1)
}

func makeMaybe() *hxrt.Array {
	return hxrt.NewArray(nil, 1, nil)
}

func makeNames() *hxrt.Array {
	return hxrt.NewArray(hxrt.StringFromLiteral("a"), hxrt.StringFromLiteral("b"), nil, hxrt.StringFromLiteral("b"))
}

func makeRows() *hxrt.Array {
	_g := hxrt.NewArray()
	_g_1 := hxrt.NewArray()
	_g_1.Push(0)
	_g_1.Push(1)
	_g_1.Push(2)
	_g.Push(_g_1)
	_g_2 := hxrt.NewArray()
	_g_2.Push(3)
	_g_2.Push(4)
	_g_2.Push(5)
	_g.Push(_g_2)
	_g_3 := hxrt.NewArray()
	_g_3.Push(6)
	_g_3.Push(7)
	_g_3.Push(8)
	_g.Push(_g_3)
	return _g
}

func noStart() any {
	return nil
}

type I_SharedArrayHolder interface {
}

type SharedArrayHolder struct {
	__hx_this I_SharedArrayHolder
	ints      *hxrt.Array
	names     *hxrt.Array
	items     *hxrt.Array
}

func (self *SharedArrayHolder) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_SharedArrayHolder(ints *hxrt.Array, names *hxrt.Array, items *hxrt.Array) *SharedArrayHolder {
	return hx_init_SharedArrayHolder(hx_alloc_SharedArrayHolder(nil), ints, names, items)
}

func hx_alloc_SharedArrayHolder(owner any) *SharedArrayHolder {
	self := &SharedArrayHolder{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_SharedArrayHolder)
	}
	return self
}

func hx_init_SharedArrayHolder(self *SharedArrayHolder, ints *hxrt.Array, names *hxrt.Array, items *hxrt.Array) *SharedArrayHolder {
	self.ints = ints
	self.names = names
	self.items = items
	return self
}

func (self *SharedArrayHolder) HxClassName() string {
	return "SharedArrayHolder"
}

type I_SharedArrayItem interface {
	toString() *string
}

type SharedArrayItem struct {
	__hx_this I_SharedArrayItem
	name      *string
}

func (self *SharedArrayItem) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_SharedArrayItem(name *string) *SharedArrayItem {
	return hx_init_SharedArrayItem(hx_alloc_SharedArrayItem(nil), name)
}

func hx_alloc_SharedArrayItem(owner any) *SharedArrayItem {
	self := &SharedArrayItem{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_SharedArrayItem)
	}
	return self
}

func hx_init_SharedArrayItem(self *SharedArrayItem, name *string) *SharedArrayItem {
	self.name = name
	return self
}

func (self *SharedArrayItem) toString() *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Item("), self.name), hxrt.StringFromLiteral(")"))
}

func (self *SharedArrayItem) String() string {
	return *self.toString()
}

func hx_recover_class_SharedArrayItem(value any) *SharedArrayItem {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *SharedArrayItem:
		return source
	default:
		return value.(*SharedArrayItem)
	}
}
