package main

import "snapshot/hxrt"

var argumentEffectCount int = 0

func checkedArgumentIterator(values *hxrt.Array) map[string]any {
	argumentEffectCount = int(int32((argumentEffectCount + 1)))
	return func() map[string]any {
		hx_structural_array_1 := values
		hx_structural_array_index_2 := 0
		hx_structural_iterator_map_3 := map[string]any{}
		hx_structural_iterator_map_3["hasNext"] = func() bool {
			return (hx_structural_array_index_2 < hx_structural_array_1.Len())
		}
		hx_structural_iterator_map_3["next"] = func() int {
			hx_structural_array_value_4 := hx_structural_array_1.Get(hx_structural_array_index_2)
			hx_structural_array_index_2 = (hx_structural_array_index_2 + 1)
			return func(hx_value_5 any) int {
				if hx_value_5 == nil {
					var hx_zero_6 int
					return hx_zero_6
				}
				return hx_value_5.(int)
			}(any(hx_structural_array_value_4))
		}
		return hx_structural_iterator_map_3
	}()
}

func checkedIterator(values *hxrt.Array) map[string]any {
	effectCount = int(int32((effectCount + 1)))
	return func() map[string]any {
		hx_structural_array_7 := values
		hx_structural_array_index_8 := 0
		hx_structural_iterator_map_9 := map[string]any{}
		hx_structural_iterator_map_9["hasNext"] = func() bool {
			return (hx_structural_array_index_8 < hx_structural_array_7.Len())
		}
		hx_structural_iterator_map_9["next"] = func() int {
			hx_structural_array_value_10 := hx_structural_array_7.Get(hx_structural_array_index_8)
			hx_structural_array_index_8 = (hx_structural_array_index_8 + 1)
			return func(hx_value_11 any) int {
				if hx_value_11 == nil {
					var hx_zero_12 int
					return hx_zero_12
				}
				return hx_value_11.(int)
			}(any(hx_structural_array_value_10))
		}
		return hx_structural_iterator_map_9
	}()
}

func collect(iterator map[string]any) *string {
	values := hxrt.NewArray()
	for func(hx_obj_13 map[string]any) func() bool {
		hx_field_14 := hx_obj_13["hasNext"]
		if hx_field_14 == nil {
			var hx_zero_15 func() bool
			return hx_zero_15
		}
		return hx_field_14.(func() bool)
	}(iterator)() {
		values.Push(func(hx_obj_17 map[string]any) func() int {
			hx_field_18 := hx_obj_17["next"]
			if hx_field_18 == nil {
				var hx_zero_19 func() int
				return hx_zero_19
			}
			return hx_field_18.(func() int)
		}(iterator)())
	}
	return hxrt.StringJoinAny(values.Values(), hxrt.StringFromLiteral(","))
}

var effectCount int = 0

func main() {
	values := hxrt.NewArray(1, 2)
	effectCount = int(int32((effectCount + 1)))
	iterator := func() map[string]any {
		hx_structural_array_20 := values
		hx_structural_array_index_21 := 0
		hx_structural_iterator_map_22 := map[string]any{}
		hx_structural_iterator_map_22["hasNext"] = func() bool {
			return (hx_structural_array_index_21 < hx_structural_array_20.Len())
		}
		hx_structural_iterator_map_22["next"] = func() int {
			hx_structural_array_value_23 := hx_structural_array_20.Get(hx_structural_array_index_21)
			hx_structural_array_index_21 = (hx_structural_array_index_21 + 1)
			return func(hx_value_24 any) int {
				if hx_value_24 == nil {
					var hx_zero_25 int
					return hx_zero_25
				}
				return hx_value_24.(int)
			}(any(hx_structural_array_value_23))
		}
		return hx_structural_iterator_map_22
	}()
	hx_array_target_26 := values
	hx_array_index_27 := 0
	hx_array_target_26.Set(hx_array_index_27, 9)
	var v any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("effect-before="), effectCount))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("values="), collect(iterator)))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("effect-after="), effectCount))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("argument-values="), collect(func() map[string]any {
		argumentEffectCount = int(int32((argumentEffectCount + 1)))
		return func() map[string]any {
			hx_structural_array_28 := hxrt.NewArray(3, 4)
			hx_structural_array_index_29 := 0
			hx_structural_iterator_map_30 := map[string]any{}
			hx_structural_iterator_map_30["hasNext"] = func() bool {
				return (hx_structural_array_index_29 < hx_structural_array_28.Len())
			}
			hx_structural_iterator_map_30["next"] = func() int {
				hx_structural_array_value_31 := hx_structural_array_28.Get(hx_structural_array_index_29)
				hx_structural_array_index_29 = (hx_structural_array_index_29 + 1)
				return func(hx_value_32 any) int {
					if hx_value_32 == nil {
						var hx_zero_33 int
						return hx_zero_33
					}
					return hx_value_32.(int)
				}(any(hx_structural_array_value_31))
			}
			return hx_structural_iterator_map_30
		}()
	}())))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("argument-effects="), argumentEffectCount))
	hxrt.Println(v_4)
}
