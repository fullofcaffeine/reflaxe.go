package main

import "snapshot/hxrt"

func label(value int) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("<"), value), hxrt.StringFromLiteral(">"))
}

func main() {
	grid := makeGrid()
	names := makeNames()
	hxrt.Println(any(int((hxrt.Int32Wrap(hxrt.IntFromNullableAny(func(hx_value_1 any) *hxrt.Array {
		if hx_value_1 == nil {
			var hx_zero_2 *hxrt.Array
			return hx_zero_2
		}
		return hx_value_1.(*hxrt.Array)
	}(grid.Get(1)).Get(2))) + hxrt.Int32Wrap(1)))))
	if hxrt.StringEqualAny(func(hx_value_3 any) *hxrt.Array {
		if hx_value_3 == nil {
			var hx_zero_4 *hxrt.Array
			return hx_zero_4
		}
		return hx_value_3.(*hxrt.Array)
	}(names.Get(1)).Get(0), hxrt.StringFromLiteral("c")) {
		hxrt.Println(any(hxrt.StringFromLiteral("found c")))
	}
	var v any = any(label(hxrt.IntFromNullableAny(func(hx_value_5 any) *hxrt.Array {
		if hx_value_5 == nil {
			var hx_zero_6 *hxrt.Array
			return hx_zero_6
		}
		return hx_value_5.(*hxrt.Array)
	}(grid.Get(2)).Get(0))))
	hxrt.Println(v)
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("cell "), func(hx_value_9 any) *hxrt.Array {
		if hx_value_9 == nil {
			var hx_zero_10 *hxrt.Array
			return hx_zero_10
		}
		return hx_value_9.(*hxrt.Array)
	}(grid.Get(0)).Get(1)), hxrt.StringFromLiteral(" name ")), func(hx_value_15 any) *hxrt.Array {
		if hx_value_15 == nil {
			var hx_zero_16 *hxrt.Array
			return hx_zero_16
		}
		return hx_value_15.(*hxrt.Array)
	}(names.Get(0)).Get(1))))
	hxrt.Println(any(hxrt.StringConcatAny(func(hx_value_25 any) *hxrt.Array {
		if hx_value_25 == nil {
			var hx_zero_26 *hxrt.Array
			return hx_zero_26
		}
		return hx_value_25.(*hxrt.Array)
	}(names.Get(0)).Get(0), func(hx_value_27 any) *hxrt.Array {
		if hx_value_27 == nil {
			var hx_zero_28 *hxrt.Array
			return hx_zero_28
		}
		return hx_value_27.(*hxrt.Array)
	}(names.Get(1)).Get(1))))
	hxrt.Println(any((hxrt.IntFromNullableAny(func(hx_value_29 any) *hxrt.Array {
		if hx_value_29 == nil {
			var hx_zero_30 *hxrt.Array
			return hx_zero_30
		}
		return hx_value_29.(*hxrt.Array)
	}(grid.Get(2)).Get(1)) > hxrt.IntFromNullableAny(func(hx_value_31 any) *hxrt.Array {
		if hx_value_31 == nil {
			var hx_zero_32 *hxrt.Array
			return hx_zero_32
		}
		return hx_value_31.(*hxrt.Array)
	}(grid.Get(1)).Get(2)))))
	hxrt.Println(any(hxrt.HaxeEqual(func(hx_value_33 any) *hxrt.Array {
		if hx_value_33 == nil {
			var hx_zero_34 *hxrt.Array
			return hx_zero_34
		}
		return hx_value_33.(*hxrt.Array)
	}(grid.Get(1)).Get(1), func(hx_value_35 any) *hxrt.Array {
		if hx_value_35 == nil {
			var hx_zero_36 *hxrt.Array
			return hx_zero_36
		}
		return hx_value_35.(*hxrt.Array)
	}(grid.Get(1)).Get(1))))
	base := func(hx_value_37 any) *hxrt.Array {
		if hx_value_37 == nil {
			var hx_zero_38 *hxrt.Array
			return hx_zero_38
		}
		return hx_value_37.(*hxrt.Array)
	}(grid.Get(0))
	func() int {
		hx_array_target_39 := base
		hx_array_index_40 := 0
		var hx_array_current_41 int = hxrt.IntFromNullableAny(hx_array_target_39.Get(hx_array_index_40))
		var hx_array_assigned_42 int = int((hxrt.Int32Wrap(hx_array_current_41) + hxrt.Int32Wrap(10)))
		hx_array_target_39.Set(hx_array_index_40, hx_array_assigned_42)
		return hx_array_assigned_42
	}()
	func() int {
		hx_array_target_43 := func(hx_value_47 any) *hxrt.Array {
			if hx_value_47 == nil {
				var hx_zero_48 *hxrt.Array
				return hx_zero_48
			}
			return hx_value_47.(*hxrt.Array)
		}(grid.Get(2))
		hx_array_index_44 := 2
		var hx_array_current_45 int = hxrt.IntFromNullableAny(hx_array_target_43.Get(hx_array_index_44))
		var hx_array_next_46 int = int(int32((hx_array_current_45 + 1)))
		hx_array_target_43.Set(hx_array_index_44, hx_array_next_46)
		return hx_array_current_45
	}()
	var v_1 any = any(hxrt.StringConcatAny(hxrt.StringConcatAny(func(hx_value_51 any) *hxrt.Array {
		if hx_value_51 == nil {
			var hx_zero_52 *hxrt.Array
			return hx_zero_52
		}
		return hx_value_51.(*hxrt.Array)
	}(grid.Get(0)).Get(0), hxrt.StringFromLiteral(":")), func() int {
		hx_array_target_53 := func(hx_value_57 any) *hxrt.Array {
			if hx_value_57 == nil {
				var hx_zero_58 *hxrt.Array
				return hx_zero_58
			}
			return hx_value_57.(*hxrt.Array)
		}(grid.Get(2))
		hx_array_index_54 := 2
		var hx_array_current_55 int = hxrt.IntFromNullableAny(hx_array_target_53.Get(hx_array_index_54))
		var hx_array_next_56 int = int(int32((hx_array_current_55 + 1)))
		hx_array_target_53.Set(hx_array_index_54, hx_array_next_56)
		return hx_array_next_56
	}()))
	hxrt.Println(v_1)
	hx_array_target_59 := func(hx_value_61 any) *hxrt.Array {
		if hx_value_61 == nil {
			var hx_zero_62 *hxrt.Array
			return hx_zero_62
		}
		return hx_value_61.(*hxrt.Array)
	}(grid.Get(1))
	hx_array_index_60 := 1
	hx_array_target_59.Set(hx_array_index_60, 40)
	assigned := func() int {
		hx_array_target_63 := func(hx_value_65 any) *hxrt.Array {
			if hx_value_65 == nil {
				var hx_zero_66 *hxrt.Array
				return hx_zero_66
			}
			return hx_value_65.(*hxrt.Array)
		}(grid.Get(1))
		hx_array_index_64 := 0
		return hxrt.IntFromNullableAny(hx_array_target_63.Set(hx_array_index_64, 30))
	}()
	hxrt.Println(any(int((hxrt.Int32Wrap(assigned) + hxrt.Int32Wrap(hxrt.IntFromNullableAny(func(hx_value_67 any) *hxrt.Array {
		if hx_value_67 == nil {
			var hx_zero_68 *hxrt.Array
			return hx_zero_68
		}
		return hx_value_67.(*hxrt.Array)
	}(grid.Get(1)).Get(1)))))))
	var layout_layers *hxrt.Array
	layout_layers = names
	var v_2 any = any(hxrt.StringConcatAny(func(hx_value_73 any) *hxrt.Array {
		if hx_value_73 == nil {
			var hx_zero_74 *hxrt.Array
			return hx_zero_74
		}
		return hx_value_73.(*hxrt.Array)
	}(layout_layers.Get(1)).Get(1), hxrt.StringFromLiteral("!")))
	hxrt.Println(v_2)
	board := New_NestedArrayBoard(grid)
	var v_3 any = any(int((hxrt.Int32Wrap(hxrt.IntFromNullableAny(func(hx_value_75 any) *hxrt.Array {
		if hx_value_75 == nil {
			var hx_zero_76 *hxrt.Array
			return hx_zero_76
		}
		return hx_value_75.(*hxrt.Array)
	}(board.cells.Get(2)).Get(1))) * hxrt.Int32Wrap(2))))
	hxrt.Println(v_3)
	base_1 := func(hx_value_77 any) *hxrt.Array {
		if hx_value_77 == nil {
			var hx_zero_78 *hxrt.Array
			return hx_zero_78
		}
		return hx_value_77.(*hxrt.Array)
	}(board.cells.Get(2))
	func() int {
		hx_array_target_79 := base_1
		hx_array_index_80 := 1
		var hx_array_current_81 int = hxrt.IntFromNullableAny(hx_array_target_79.Get(hx_array_index_80))
		var hx_array_assigned_82 int = int((hxrt.Int32Wrap(hx_array_current_81) - hxrt.Int32Wrap(1)))
		hx_array_target_79.Set(hx_array_index_80, hx_array_assigned_82)
		return hx_array_assigned_82
	}()
	hxrt.Println(func(hx_value_83 any) *hxrt.Array {
		if hx_value_83 == nil {
			var hx_zero_84 *hxrt.Array
			return hx_zero_84
		}
		return hx_value_83.(*hxrt.Array)
	}(grid.Get(2)).Get(1))
	_g := hxrt.NewArray()
	_g_1 := hxrt.NewArray()
	_g_2 := hxrt.NewArray()
	_g_2.Push(0)
	_g_2.Push(1)
	_g_1.Push(_g_2)
	_g_3 := hxrt.NewArray()
	_g_3.Push(2)
	_g_3.Push(3)
	_g_1.Push(_g_3)
	_g.Push(_g_1)
	_g_4 := hxrt.NewArray()
	_g_5 := hxrt.NewArray()
	_g_5.Push(4)
	_g_5.Push(5)
	_g_4.Push(_g_5)
	_g_6 := hxrt.NewArray()
	_g_6.Push(6)
	_g_6.Push(7)
	_g_4.Push(_g_6)
	_g.Push(_g_4)
	cube := _g
	base_2 := func(hx_value_101 any) *hxrt.Array {
		if hx_value_101 == nil {
			var hx_zero_102 *hxrt.Array
			return hx_zero_102
		}
		return hx_value_101.(*hxrt.Array)
	}(func(hx_value_99 any) *hxrt.Array {
		if hx_value_99 == nil {
			var hx_zero_100 *hxrt.Array
			return hx_zero_100
		}
		return hx_value_99.(*hxrt.Array)
	}(cube.Get(1)).Get(0))
	func() int {
		hx_array_target_103 := base_2
		hx_array_index_104 := 1
		var hx_array_current_105 int = hxrt.IntFromNullableAny(hx_array_target_103.Get(hx_array_index_104))
		var hx_array_assigned_106 int = int((hxrt.Int32Wrap(hx_array_current_105) + hxrt.Int32Wrap(100)))
		hx_array_target_103.Set(hx_array_index_104, hx_array_assigned_106)
		return hx_array_assigned_106
	}()
	hxrt.Println(any(int((hxrt.Int32Wrap(hxrt.IntFromNullableAny(func(hx_value_109 any) *hxrt.Array {
		if hx_value_109 == nil {
			var hx_zero_110 *hxrt.Array
			return hx_zero_110
		}
		return hx_value_109.(*hxrt.Array)
	}(func(hx_value_107 any) *hxrt.Array {
		if hx_value_107 == nil {
			var hx_zero_108 *hxrt.Array
			return hx_zero_108
		}
		return hx_value_107.(*hxrt.Array)
	}(cube.Get(1)).Get(0)).Get(1))) + hxrt.Int32Wrap(hxrt.IntFromNullableAny(func(hx_value_113 any) *hxrt.Array {
		if hx_value_113 == nil {
			var hx_zero_114 *hxrt.Array
			return hx_zero_114
		}
		return hx_value_113.(*hxrt.Array)
	}(func(hx_value_111 any) *hxrt.Array {
		if hx_value_111 == nil {
			var hx_zero_112 *hxrt.Array
			return hx_zero_112
		}
		return hx_value_111.(*hxrt.Array)
	}(cube.Get(0)).Get(1)).Get(0)))))))
	_g_7 := hxrt.NewArray()
	_g_7.Push(0)
	_g_7.Push(2)
	order := _g_7
	_g_8 := hxrt.NewArray()
	_g_8.Push(0)
	_g_8.Push(0)
	_g_8.Push(0)
	counts := _g_8
	hx_array_target_120 := counts
	hx_array_index_121 := hxrt.IntFromNullableAny(order.Get(1))
	hx_array_target_120.Set(hx_array_index_121, 7)
	index := hxrt.IntFromNullableAny(order.Get(0))
	func() int {
		hx_array_target_122 := counts
		hx_array_index_123 := index
		var hx_array_current_124 int = hxrt.IntFromNullableAny(hx_array_target_122.Get(hx_array_index_123))
		var hx_array_assigned_125 int = int((hxrt.Int32Wrap(hx_array_current_124) + hxrt.Int32Wrap(2)))
		hx_array_target_122.Set(hx_array_index_123, hx_array_assigned_125)
		return hx_array_assigned_125
	}()
	hxrt.Println(any(int((hxrt.Int32Wrap(hxrt.IntFromNullableAny(counts.Get(hxrt.IntFromNullableAny(order.Get(1))))) + hxrt.Int32Wrap(hxrt.IntFromNullableAny(counts.Get(hxrt.IntFromNullableAny(order.Get(0)))))))))
	hxrt.Println(func(hx_value_126 any) *hxrt.Array {
		if hx_value_126 == nil {
			var hx_zero_127 *hxrt.Array
			return hx_zero_127
		}
		return hx_value_126.(*hxrt.Array)
	}(grid.Get(hxrt.IntFromNullableAny(order.Get(0)))).Get(hxrt.IntFromNullableAny(order.Get(1))))
	maybe := makeMaybe()
	hxrt.Println(func(hx_value_128 any) *hxrt.Array {
		if hx_value_128 == nil {
			var hx_zero_129 *hxrt.Array
			return hx_zero_129
		}
		return hx_value_128.(*hxrt.Array)
	}(grid.Get(0)).Get(hxrt.IntFromNullableAny(maybe.Get(0))))
	var nullableIndex any = maybe.Get(0)
	hxrt.Println(func(hx_value_130 any) *hxrt.Array {
		if hx_value_130 == nil {
			var hx_zero_131 *hxrt.Array
			return hx_zero_131
		}
		return hx_value_130.(*hxrt.Array)
	}(grid.Get(1)).Get(hxrt.IntFromNullableAny(nullableIndex)))
	var this1 []int
	this1 = []int{}
	hx_len_132 := 3
	if hx_len_132 < 0 {
		hx_len_132 = 0
	}
	if hx_len_132 <= len(this1) {
		this1 = this1[:hx_len_132]
	} else {
		var hx_zero_133 int
		for len(this1) < hx_len_132 {
			this1 = append(this1, hx_zero_133)
		}
	}
	vector := this1
	vector[hxrt.IntFromNullableAny(order.Get(1))] = 5
	hxrt.Println(any(vector[hxrt.IntFromNullableAny(order.Get(1))]))
	hx_arr_134 := func(hx_value_135 any) *hxrt.Array {
		if hx_value_135 == nil {
			var hx_zero_136 *hxrt.Array
			return hx_zero_136
		}
		return hx_value_135.(*hxrt.Array)
	}(grid.Get(0))
	hx_arr_134.Push(9)
	var v_4 any = any(hxrt.StringConcatAny(hxrt.StringConcatAny(func(hx_value_137 any) *hxrt.Array {
		if hx_value_137 == nil {
			var hx_zero_138 *hxrt.Array
			return hx_zero_138
		}
		return hx_value_137.(*hxrt.Array)
	}(grid.Get(0)).Len(), hxrt.StringFromLiteral(":")), func() any {
		hx_arr_139 := func(hx_value_140 any) *hxrt.Array {
			if hx_value_140 == nil {
				var hx_zero_141 *hxrt.Array
				return hx_zero_141
			}
			return hx_value_140.(*hxrt.Array)
		}(grid.Get(0))
		return hx_arr_139.Pop()
	}()))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringJoinAny(func(hx_value_142 any) *hxrt.Array {
		if hx_value_142 == nil {
			var hx_zero_143 *hxrt.Array
			return hx_zero_143
		}
		return hx_value_142.(*hxrt.Array)
	}(grid.Get(1)).Values(), hxrt.StringFromLiteral("-")), hxrt.StringFromLiteral(" ")), func(hx_value_144 any) *hxrt.Array {
		if hx_value_144 == nil {
			var hx_zero_145 *hxrt.Array
			return hx_zero_145
		}
		return hx_value_144.(*hxrt.Array)
	}(grid.Get(1)).Copy().Len()))
	hxrt.Println(v_5)
	var v_6 any = any(hxrt.StringConcatAny(hxrt.StringConcatAny(func() int {
		hx_indexof_target_146 := func(hx_value_151 any) *hxrt.Array {
			if hx_value_151 == nil {
				var hx_zero_152 *hxrt.Array
				return hx_zero_152
			}
			return hx_value_151.(*hxrt.Array)
		}(grid.Get(1))
		var hx_indexof_value_147 int = 40
		hx_indexof_length_148 := hx_indexof_target_146.Len()
		hx_indexof_index_149 := 0
		for hx_indexof_index_149 < hx_indexof_length_148 {
			hx_indexof_element_150 := hxrt.IntFromNullableAny(hx_indexof_target_146.Get(hx_indexof_index_149))
			if hx_indexof_element_150 == hx_indexof_value_147 {
				return hx_indexof_index_149
			}
			hx_indexof_index_149 = (hx_indexof_index_149 + 1)
		}
		return -1
	}(), hxrt.StringFromLiteral(":")), func() *hxrt.Array {
		hx_concat_target_153 := func(hx_value_157 any) *hxrt.Array {
			if hx_value_157 == nil {
				var hx_zero_158 *hxrt.Array
				return hx_zero_158
			}
			return hx_value_157.(*hxrt.Array)
		}(grid.Get(0))
		hx_concat_other_154 := func(hx_value_159 any) *hxrt.Array {
			if hx_value_159 == nil {
				var hx_zero_160 *hxrt.Array
				return hx_zero_160
			}
			return hx_value_159.(*hxrt.Array)
		}(grid.Get(1))
		hx_concat_result_155 := hx_concat_target_153.Copy()
		for _, hx_concat_item_156 := range hx_concat_other_154.Values() {
			hx_concat_result_155.Push(hx_concat_item_156)
		}
		return hx_concat_result_155
	}().Len()))
	hxrt.Println(v_6)
	func() {
		hx_arr_161 := func(hx_value_162 any) *hxrt.Array {
			if hx_value_162 == nil {
				var hx_zero_163 *hxrt.Array
				return hx_zero_163
			}
			return hx_value_162.(*hxrt.Array)
		}(grid.Get(2))
		hx_insert_position_164 := hxrt.IntFromNullableAny(order.Get(0))
		var hx_insert_value_165 any = -1
		hx_arr_161.Insert(hx_insert_position_164, hx_insert_value_165)
	}()
	var v_7 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StdString(func() bool {
		hx_arr_166 := func(hx_value_167 any) *hxrt.Array {
			if hx_value_167 == nil {
				var hx_zero_168 *hxrt.Array
				return hx_zero_168
			}
			return hx_value_167.(*hxrt.Array)
		}(grid.Get(2))
		var hx_remove_value_169 any = -1
		hx_remove_index_170 := 0
		for hx_remove_index_170 < hx_arr_166.Len() {
			hx_remove_element_171 := hx_arr_166.Get(hx_remove_index_170)
			if hxrt.HaxeEqual(hx_remove_element_171, hx_remove_value_169) {
				hx_arr_166.RemoveAt(hx_remove_index_170)
				return true
			}
			hx_remove_index_170 = (hx_remove_index_170 + 1)
		}
		return false
	}()), hxrt.StringFromLiteral(":")), func() any {
		hx_arr_172 := func(hx_value_173 any) *hxrt.Array {
			if hx_value_173 == nil {
				var hx_zero_174 *hxrt.Array
				return hx_zero_174
			}
			return hx_value_173.(*hxrt.Array)
		}(grid.Get(2))
		hx_value_175 := hx_arr_172.Get(0)
		hx_arr_172.RemoveAt(0)
		return hx_value_175
	}()))
	hxrt.Println(v_7)
	var v_8 any = any(func(hx_value_176 any) *hxrt.Array {
		if hx_value_176 == nil {
			var hx_zero_177 *hxrt.Array
			return hx_zero_177
		}
		return hx_value_176.(*hxrt.Array)
	}(grid.Get(2)).SliceFrom(hxrt.IntFromNullableAny(order.Get(0))).Len())
	hxrt.Println(v_8)
	hxrt.ArraySort(func(hx_value_178 any) *hxrt.Array {
		if hx_value_178 == nil {
			var hx_zero_179 *hxrt.Array
			return hx_zero_179
		}
		return hx_value_178.(*hxrt.Array)
	}(grid.Get(2)), func(hx_comparator_182 *hxrt.Function[func(int, int) int]) *hxrt.Function[func(any, any) int] {
		return hxrt.AdaptFunction(hx_comparator_182, func(hx_cmp_left_180 any, hx_cmp_right_181 any) int {
			return hxrt.FunctionCall(hx_comparator_182)(func(hx_value_183 any) int {
				if hx_value_183 == nil {
					var hx_zero_184 int
					return hx_zero_184
				}
				return hx_value_183.(int)
			}(hx_cmp_left_180), func(hx_value_185 any) int {
				if hx_value_185 == nil {
					var hx_zero_186 int
					return hx_zero_186
				}
				return hx_value_185.(int)
			}(hx_cmp_right_181))
		})
	}(hxrt.NewFunction(func(left int, right int) int {
		return int((hxrt.Int32Wrap(right) - hxrt.Int32Wrap(left)))
	})))
	hxrt.Println(func(hx_value_187 any) *hxrt.Array {
		if hx_value_187 == nil {
			var hx_zero_188 *hxrt.Array
			return hx_zero_188
		}
		return hx_value_187.(*hxrt.Array)
	}(grid.Get(2)).Get(0))
	_g_9 := 0
	_g1 := func(hx_value_189 any) *hxrt.Array {
		if hx_value_189 == nil {
			var hx_zero_190 *hxrt.Array
			return hx_zero_190
		}
		return hx_value_189.(*hxrt.Array)
	}(grid.Get(1))
	for _g_9 < _g1.Len() {
		value := hxrt.IntFromNullableAny(_g1.Get(_g_9))
		_g_9 = int(int32((_g_9 + 1)))
		hxrt.Println(any(value))
	}
	_g_10 := hxrt.NewArray()
	_g_10.Push(New_haxe__ds__StringMap())
	_g_10.Push(New_haxe__ds__StringMap())
	seen := _g_10
	if !func(hx_value_193 any) bool {
		if hx_value_193 == nil {
			var hx_zero_194 bool
			return hx_zero_194
		}
		return hx_value_193.(bool)
	}(hx_recover_class_0(seen.Get(1)).__hx_this.exists(hxrt.StringFromLiteral("k"))) {
		hx_recover_class_0(seen.Get(1)).__hx_this.set(hxrt.StringFromLiteral("k"), true)
	}
	var v_9 any = any(func(hx_value_195 any) bool {
		if hx_value_195 == nil {
			var hx_zero_196 bool
			return hx_zero_196
		}
		return hx_value_195.(bool)
	}(hx_recover_class_0(seen.Get(1)).__hx_this.exists(hxrt.StringFromLiteral("k"))))
	hxrt.Println(v_9)
	_g_11 := hxrt.NewArray()
	_g_12 := hxrt.NewArray()
	_g_12.Push(hxrt.StringFromLiteral(" "))
	_g_12.Push(hxrt.StringFromLiteral(" "))
	_g_12.Push(hxrt.StringFromLiteral(" "))
	_g_11.Push(_g_12)
	_g_13 := hxrt.NewArray()
	_g_13.Push(hxrt.StringFromLiteral(" "))
	_g_13.Push(hxrt.StringFromLiteral(" "))
	_g_13.Push(hxrt.StringFromLiteral(" "))
	_g_11.Push(_g_13)
	canvas := _g_11
	draw := hxrt.NewFunction(func(y int, x int, incoming *string) {
		hx_array_target_208 := func(hx_value_210 any) *hxrt.Array {
			if hx_value_210 == nil {
				var hx_zero_211 *hxrt.Array
				return hx_zero_211
			}
			return hx_value_210.(*hxrt.Array)
		}(canvas.Get(y))
		hx_array_index_209 := x
		var hx_if_207 *string
		if hxrt.StringEqualAny(func(hx_value_205 any) *hxrt.Array {
			if hx_value_205 == nil {
				var hx_zero_206 *hxrt.Array
				return hx_zero_206
			}
			return hx_value_205.(*hxrt.Array)
		}(canvas.Get(y)).Get(x), hxrt.StringFromLiteral(" ")) {
			hx_if_207 = incoming
		} else {
			hx_if_207 = hxrt.StringFromLiteral("+")
		}
		hx_array_target_208.Set(hx_array_index_209, hx_if_207)
	})
	func() {
		var hx_callable_212 *hxrt.Function[func(int, int, *string)] = draw
		var hx_argument_213 int = 1
		var hx_argument_214 int = 2
		var hx_argument_215 *string = hxrt.StringFromLiteral("-")
		hxrt.FunctionCall(hx_callable_212)(hx_argument_213, hx_argument_214, hx_argument_215)
	}()
	func() {
		var hx_callable_216 *hxrt.Function[func(int, int, *string)] = draw
		var hx_argument_217 int = 1
		var hx_argument_218 int = 2
		var hx_argument_219 *string = hxrt.StringFromLiteral("|")
		hxrt.FunctionCall(hx_callable_216)(hx_argument_217, hx_argument_218, hx_argument_219)
	}()
	func() {
		var hx_callable_220 *hxrt.Function[func(int, int, *string)] = draw
		var hx_argument_221 int = 0
		var hx_argument_222 int = 0
		var hx_argument_223 *string = hxrt.StringFromLiteral("*")
		hxrt.FunctionCall(hx_callable_220)(hx_argument_221, hx_argument_222, hx_argument_223)
	}()
	var v_10 any = any(hxrt.StringJoinAny(func() *hxrt.Array {
		_g_14 := hxrt.NewArray()
		_g1_1 := 0
		for _g1_1 < canvas.Len() {
			line := func(hx_value_224 any) *hxrt.Array {
				if hx_value_224 == nil {
					var hx_zero_225 *hxrt.Array
					return hx_zero_225
				}
				return hx_value_224.(*hxrt.Array)
			}(canvas.Get(_g1_1))
			_g1_1 = int(int32((_g1_1 + 1)))
			_g_14.Push(hxrt.StringJoinAny(line.Values(), hxrt.StringFromLiteral("")))
		}
		return _g_14
	}().Values(), hxrt.StringFromLiteral("/")))
	hxrt.Println(v_10)
}

func makeGrid() *hxrt.Array {
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

func makeMaybe() *hxrt.Array {
	return hxrt.NewArray(2, nil)
}

func makeNames() *hxrt.Array {
	_g := hxrt.NewArray()
	_g_1 := hxrt.NewArray()
	_g_1.Push(hxrt.StringFromLiteral("a"))
	_g_1.Push(hxrt.StringFromLiteral("b"))
	_g.Push(_g_1)
	_g_2 := hxrt.NewArray()
	_g_2.Push(hxrt.StringFromLiteral("c"))
	_g_2.Push(hxrt.StringFromLiteral("d"))
	_g.Push(_g_2)
	return _g
}

type I_NestedArrayBoard interface {
}

type NestedArrayBoard struct {
	__hx_this I_NestedArrayBoard
	cells     *hxrt.Array
}

func (self *NestedArrayBoard) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_NestedArrayBoard(cells *hxrt.Array) *NestedArrayBoard {
	return hx_init_NestedArrayBoard(hx_alloc_NestedArrayBoard(nil), cells)
}

func hx_alloc_NestedArrayBoard(owner any) *NestedArrayBoard {
	self := &NestedArrayBoard{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_NestedArrayBoard)
	}
	return self
}

func hx_init_NestedArrayBoard(self *NestedArrayBoard, cells *hxrt.Array) *NestedArrayBoard {
	self.cells = cells
	return self
}

func hx_recover_callable_0(value any) *hxrt.Function[func() bool] {
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
	case *hxrt.Function[func() bool]:
		return source
	case *hxrt.Function[func() *string]:
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
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_5 any) bool {
				if hx_value_5 == nil {
					var hx_zero_6 bool
					return hx_zero_6
				}
				return hx_value_5.(bool)
			}(hxrt.FunctionCall(source)())
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_1(value any) *hxrt.Function[func() any] {
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
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() any]:
		return source
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_2(value any) *hxrt.Function[func() *string] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() *string]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_7 any) *string {
				if hx_value_7 == nil {
					var hx_zero_8 *string
					return hx_zero_8
				}
				return hx_value_7.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *string]:
		return source
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_9 any) *string {
				if hx_value_9 == nil {
					var hx_zero_10 *string
					return hx_zero_10
				}
				return hx_value_9.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_11 any) *string {
				if hx_value_11 == nil {
					var hx_zero_12 *string
					return hx_zero_12
				}
				return hx_value_11.(*string)
			}(hxrt.FunctionCall(source)())
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_class_0(value any) *haxe__ds__StringMap {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ds__StringMap:
		return source
	default:
		return value.(*haxe__ds__StringMap)
	}
}
