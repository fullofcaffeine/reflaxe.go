package main

import "snapshot/hxrt"

var argumentEffectCount int = 0

func checkedArgumentIterator(values []int) map[string]any {
	argumentEffectCount = int(int32((argumentEffectCount + 1)))
	return func() map[string]any {
		hx_structural_array_1 := values
		hx_structural_array_index_2 := 0
		hx_structural_iterator_map_3 := map[string]any{}
		hx_structural_iterator_map_3["hasNext"] = func() bool {
			return (hx_structural_array_index_2 < len(hx_structural_array_1))
		}
		hx_structural_iterator_map_3["next"] = func() int {
			hx_structural_array_value_4 := hx_structural_array_1[hx_structural_array_index_2]
			hx_structural_array_index_2 = (hx_structural_array_index_2 + 1)
			return hx_structural_array_value_4
		}
		return hx_structural_iterator_map_3
	}()
}

func checkedIterator(values []int) map[string]any {
	effectCount = int(int32((effectCount + 1)))
	return func() map[string]any {
		hx_structural_array_5 := values
		hx_structural_array_index_6 := 0
		hx_structural_iterator_map_7 := map[string]any{}
		hx_structural_iterator_map_7["hasNext"] = func() bool {
			return (hx_structural_array_index_6 < len(hx_structural_array_5))
		}
		hx_structural_iterator_map_7["next"] = func() int {
			hx_structural_array_value_8 := hx_structural_array_5[hx_structural_array_index_6]
			hx_structural_array_index_6 = (hx_structural_array_index_6 + 1)
			return hx_structural_array_value_8
		}
		return hx_structural_iterator_map_7
	}()
}

func collect(iterator map[string]any) *string {
	values := []int{}
	for func(hx_obj_9 map[string]any) func() bool {
		hx_field_10 := hx_obj_9["hasNext"]
		if hx_field_10 == nil {
			var hx_zero_11 func() bool
			return hx_zero_11
		}
		return hx_field_10.(func() bool)
	}(iterator)() {
		values = append(values, func(hx_obj_13 map[string]any) func() int {
			hx_field_14 := hx_obj_13["next"]
			if hx_field_14 == nil {
				var hx_zero_15 func() int
				return hx_zero_15
			}
			return hx_field_14.(func() int)
		}(iterator)())
	}
	return hxrt.StringJoinAny(func(hx_sort_src_16 []int) []any {
		hx_sort_out_18 := make([]any, 0, len(hx_sort_src_16))
		for _, hx_sort_item_17 := range hx_sort_src_16 {
			hx_sort_out_18 = append(hx_sort_out_18, hx_sort_item_17)
		}
		return hx_sort_out_18
	}(values), hxrt.StringFromLiteral(","))
}

var effectCount int = 0

func main() {
	values := []int{1, 2}
	effectCount = int(int32((effectCount + 1)))
	iterator := func() map[string]any {
		hx_structural_array_19 := values
		hx_structural_array_index_20 := 0
		hx_structural_iterator_map_21 := map[string]any{}
		hx_structural_iterator_map_21["hasNext"] = func() bool {
			return (hx_structural_array_index_20 < len(hx_structural_array_19))
		}
		hx_structural_iterator_map_21["next"] = func() int {
			hx_structural_array_value_22 := hx_structural_array_19[hx_structural_array_index_20]
			hx_structural_array_index_20 = (hx_structural_array_index_20 + 1)
			return hx_structural_array_value_22
		}
		return hx_structural_iterator_map_21
	}()
	values[0] = 9
	var v any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("effect-before="), effectCount))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("values="), collect(iterator)))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("effect-after="), effectCount))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("argument-values="), collect(func() map[string]any {
		argumentEffectCount = int(int32((argumentEffectCount + 1)))
		return func() map[string]any {
			hx_structural_array_23 := []int{3, 4}
			hx_structural_array_index_24 := 0
			hx_structural_iterator_map_25 := map[string]any{}
			hx_structural_iterator_map_25["hasNext"] = func() bool {
				return (hx_structural_array_index_24 < len(hx_structural_array_23))
			}
			hx_structural_iterator_map_25["next"] = func() int {
				hx_structural_array_value_26 := hx_structural_array_23[hx_structural_array_index_24]
				hx_structural_array_index_24 = (hx_structural_array_index_24 + 1)
				return hx_structural_array_value_26
			}
			return hx_structural_iterator_map_25
		}()
	}())))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("argument-effects="), argumentEffectCount))
	hxrt.Println(v_4)
}
