package main

import "snapshot/hxrt"

type I_Bounds interface {
}

type Bounds struct {
	__hx_this I_Bounds
	maybeFrom any
}

func (self *Bounds) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Bounds(maybeFrom any) *Bounds {
	return hx_init_Bounds(hx_alloc_Bounds(nil), maybeFrom)
}

func hx_alloc_Bounds(owner any) *Bounds {
	self := &Bounds{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Bounds)
	}
	return self
}

func hx_init_Bounds(self *Bounds, maybeFrom any) *Bounds {
	self.maybeFrom = maybeFrom
	return self
}

func (self *Bounds) HxClassName() string {
	return "Bounds"
}

var Bounds_shared any = nil

func dynamicValue(value any) any {
	return value
}

func main() {
	a := hxrt.NewArray(hxrt.StringFromLiteral("a"), hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("c"), hxrt.StringFromLiteral("a"), hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("c"))
	ints := hxrt.NewArray(1, 2, 3, 1, 2, 3)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("omitted="), hxrt.StringJoinAny(hxrt.NewArray(hxrt.StdString(func() int {
		hx_indexof_target_1 := a
		var hx_indexof_value_2 *string = hxrt.StringFromLiteral("b")
		hx_indexof_length_3 := hx_indexof_target_1.Len()
		hx_indexof_index_4 := 0
		for hx_indexof_index_4 < hx_indexof_length_3 {
			hx_indexof_element_5 := func(hx_value_6 any) *string {
				if hx_value_6 == nil {
					var hx_zero_7 *string
					return hx_zero_7
				}
				return hx_value_6.(*string)
			}(hx_indexof_target_1.Get(hx_indexof_index_4))
			if hxrt.StringEqualStringPtr(hx_indexof_element_5, hx_indexof_value_2) {
				return hx_indexof_index_4
			}
			hx_indexof_index_4 = (hx_indexof_index_4 + 1)
		}
		return -1
	}()), hxrt.StdString(func() int {
		hx_lastindexof_target_8 := a
		var hx_lastindexof_value_9 *string = hxrt.StringFromLiteral("b")
		hx_lastindexof_length_10 := hx_lastindexof_target_8.Len()
		hx_lastindexof_index_11 := (hx_lastindexof_length_10 - 1)
		for hx_lastindexof_index_11 >= 0 {
			hx_lastindexof_element_12 := func(hx_value_13 any) *string {
				if hx_value_13 == nil {
					var hx_zero_14 *string
					return hx_zero_14
				}
				return hx_value_13.(*string)
			}(hx_lastindexof_target_8.Get(hx_lastindexof_index_11))
			if hxrt.StringEqualStringPtr(hx_lastindexof_element_12, hx_lastindexof_value_9) {
				return hx_lastindexof_index_11
			}
			hx_lastindexof_index_11 = (hx_lastindexof_index_11 - 1)
		}
		return -1
	}()), hxrt.StringJoinAny(a.SliceFrom(1).Values(), hxrt.StringFromLiteral(","))).Values(), hxrt.StringFromLiteral("|"))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("literal.null="), hxrt.StringJoinAny(hxrt.NewArray(hxrt.StdString(func() int {
		hx_indexof_target_15 := a
		var hx_indexof_value_16 *string = hxrt.StringFromLiteral("b")
		hx_indexof_length_17 := hx_indexof_target_15.Len()
		hx_indexof_index_18 := 0
		for hx_indexof_index_18 < hx_indexof_length_17 {
			hx_indexof_element_19 := func(hx_value_20 any) *string {
				if hx_value_20 == nil {
					var hx_zero_21 *string
					return hx_zero_21
				}
				return hx_value_20.(*string)
			}(hx_indexof_target_15.Get(hx_indexof_index_18))
			if hxrt.StringEqualStringPtr(hx_indexof_element_19, hx_indexof_value_16) {
				return hx_indexof_index_18
			}
			hx_indexof_index_18 = (hx_indexof_index_18 + 1)
		}
		return -1
	}()), hxrt.StdString(func() int {
		hx_lastindexof_target_22 := a
		var hx_lastindexof_value_23 *string = hxrt.StringFromLiteral("b")
		hx_lastindexof_length_24 := hx_lastindexof_target_22.Len()
		hx_lastindexof_index_25 := (hx_lastindexof_length_24 - 1)
		for hx_lastindexof_index_25 >= 0 {
			hx_lastindexof_element_26 := func(hx_value_27 any) *string {
				if hx_value_27 == nil {
					var hx_zero_28 *string
					return hx_zero_28
				}
				return hx_value_27.(*string)
			}(hx_lastindexof_target_22.Get(hx_lastindexof_index_25))
			if hxrt.StringEqualStringPtr(hx_lastindexof_element_26, hx_lastindexof_value_23) {
				return hx_lastindexof_index_25
			}
			hx_lastindexof_index_25 = (hx_lastindexof_index_25 - 1)
		}
		return -1
	}()), hxrt.StringJoinAny(a.SliceOptional(1, nil).Values(), hxrt.StringFromLiteral(","))).Values(), hxrt.StringFromLiteral("|"))))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("constant="), hxrt.StringJoinAny(hxrt.NewArray(hxrt.StdString(func() int {
		hx_indexof_target_29 := a
		var hx_indexof_value_30 *string = hxrt.StringFromLiteral("b")
		hx_indexof_index_32 := 2
		hx_indexof_length_31 := hx_indexof_target_29.Len()
		if hx_indexof_index_32 < 0 {
			hx_indexof_index_32 = (hx_indexof_length_31 + hx_indexof_index_32)
		}
		if hx_indexof_index_32 < 0 {
			hx_indexof_index_32 = 0
		}
		for hx_indexof_index_32 < hx_indexof_length_31 {
			hx_indexof_element_33 := func(hx_value_34 any) *string {
				if hx_value_34 == nil {
					var hx_zero_35 *string
					return hx_zero_35
				}
				return hx_value_34.(*string)
			}(hx_indexof_target_29.Get(hx_indexof_index_32))
			if hxrt.StringEqualStringPtr(hx_indexof_element_33, hx_indexof_value_30) {
				return hx_indexof_index_32
			}
			hx_indexof_index_32 = (hx_indexof_index_32 + 1)
		}
		return -1
	}()), hxrt.StdString(func() int {
		hx_lastindexof_target_36 := a
		var hx_lastindexof_value_37 *string = hxrt.StringFromLiteral("b")
		hx_lastindexof_index_39 := 2
		hx_lastindexof_length_38 := hx_lastindexof_target_36.Len()
		if hx_lastindexof_index_39 < 0 {
			hx_lastindexof_index_39 = (hx_lastindexof_length_38 + hx_lastindexof_index_39)
		}
		if hx_lastindexof_index_39 >= hx_lastindexof_length_38 {
			hx_lastindexof_index_39 = (hx_lastindexof_length_38 - 1)
		}
		for hx_lastindexof_index_39 >= 0 {
			hx_lastindexof_element_40 := func(hx_value_41 any) *string {
				if hx_value_41 == nil {
					var hx_zero_42 *string
					return hx_zero_42
				}
				return hx_value_41.(*string)
			}(hx_lastindexof_target_36.Get(hx_lastindexof_index_39))
			if hxrt.StringEqualStringPtr(hx_lastindexof_element_40, hx_lastindexof_value_37) {
				return hx_lastindexof_index_39
			}
			hx_lastindexof_index_39 = (hx_lastindexof_index_39 - 1)
		}
		return -1
	}()), hxrt.StringJoinAny(a.SliceOptional(1, 3).Values(), hxrt.StringFromLiteral(",")), hxrt.StdString(func() int {
		hx_indexof_target_43 := a
		var hx_indexof_value_44 *string = hxrt.StringFromLiteral("b")
		hx_indexof_index_46 := -2
		hx_indexof_length_45 := hx_indexof_target_43.Len()
		if hx_indexof_index_46 < 0 {
			hx_indexof_index_46 = (hx_indexof_length_45 + hx_indexof_index_46)
		}
		if hx_indexof_index_46 < 0 {
			hx_indexof_index_46 = 0
		}
		for hx_indexof_index_46 < hx_indexof_length_45 {
			hx_indexof_element_47 := func(hx_value_48 any) *string {
				if hx_value_48 == nil {
					var hx_zero_49 *string
					return hx_zero_49
				}
				return hx_value_48.(*string)
			}(hx_indexof_target_43.Get(hx_indexof_index_46))
			if hxrt.StringEqualStringPtr(hx_indexof_element_47, hx_indexof_value_44) {
				return hx_indexof_index_46
			}
			hx_indexof_index_46 = (hx_indexof_index_46 + 1)
		}
		return -1
	}()), hxrt.StdString(func() int {
		hx_lastindexof_target_50 := a
		var hx_lastindexof_value_51 *string = hxrt.StringFromLiteral("b")
		hx_lastindexof_index_53 := -2
		hx_lastindexof_length_52 := hx_lastindexof_target_50.Len()
		if hx_lastindexof_index_53 < 0 {
			hx_lastindexof_index_53 = (hx_lastindexof_length_52 + hx_lastindexof_index_53)
		}
		if hx_lastindexof_index_53 >= hx_lastindexof_length_52 {
			hx_lastindexof_index_53 = (hx_lastindexof_length_52 - 1)
		}
		for hx_lastindexof_index_53 >= 0 {
			hx_lastindexof_element_54 := func(hx_value_55 any) *string {
				if hx_value_55 == nil {
					var hx_zero_56 *string
					return hx_zero_56
				}
				return hx_value_55.(*string)
			}(hx_lastindexof_target_50.Get(hx_lastindexof_index_53))
			if hxrt.StringEqualStringPtr(hx_lastindexof_element_54, hx_lastindexof_value_51) {
				return hx_lastindexof_index_53
			}
			hx_lastindexof_index_53 = (hx_lastindexof_index_53 - 1)
		}
		return -1
	}()), hxrt.StringJoinAny(a.SliceOptional(1, -2).Values(), hxrt.StringFromLiteral(","))).Values(), hxrt.StringFromLiteral("|"))))
	hxrt.Println(v_2)
	var value any = nil
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("param."), value), hxrt.StringFromLiteral("=")), three(a, ints, value)))
	hxrt.Println(v_3)
	var value_1 any = 0
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("param."), value_1), hxrt.StringFromLiteral("=")), three(a, ints, value_1)))
	hxrt.Println(v_4)
	var value_2 any = 2
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("param."), value_2), hxrt.StringFromLiteral("=")), three(a, ints, value_2)))
	hxrt.Println(v_5)
	var value_3 any = -2
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("param."), value_3), hxrt.StringFromLiteral("=")), three(a, ints, value_3)))
	hxrt.Println(v_6)
	var value_4 any = 5
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("param."), value_4), hxrt.StringFromLiteral("=")), three(a, ints, value_4)))
	hxrt.Println(v_7)
	var value_5 any = 99
	var v_8 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("param."), value_5), hxrt.StringFromLiteral("=")), three(a, ints, value_5)))
	hxrt.Println(v_8)
	var v_9 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("optional.omitted="), optional(a, nil)))
	hxrt.Println(v_9)
	var v_10 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("optional.null="), optional(a, nil)))
	hxrt.Println(v_10)
	var v_11 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("optional.value="), optional(a, 3)))
	hxrt.Println(v_11)
	var v_12 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("narrowed="), narrowed(a, 3)), hxrt.StringFromLiteral(",")), narrowed(a, nil)))
	hxrt.Println(v_12)
	var none any = noInt()
	var some any = someInt(2)
	var v_13 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("local="), hxrt.StringJoinAny(hxrt.NewArray(hxrt.StdString(func() int {
		hx_indexof_target_57 := a
		var hx_indexof_value_58 *string = hxrt.StringFromLiteral("b")
		var hx_indexof_start_62 any = none
		hx_indexof_length_59 := hx_indexof_target_57.Len()
		hx_indexof_index_60 := 0
		if hx_indexof_start_62 != nil {
			hx_indexof_index_60 = hxrt.IntFromNullableAny(hx_indexof_start_62)
			if hx_indexof_index_60 < 0 {
				hx_indexof_index_60 = (hx_indexof_length_59 + hx_indexof_index_60)
			}
			if hx_indexof_index_60 < 0 {
				hx_indexof_index_60 = 0
			}
		}
		for hx_indexof_index_60 < hx_indexof_length_59 {
			hx_indexof_element_61 := func(hx_value_63 any) *string {
				if hx_value_63 == nil {
					var hx_zero_64 *string
					return hx_zero_64
				}
				return hx_value_63.(*string)
			}(hx_indexof_target_57.Get(hx_indexof_index_60))
			if hxrt.StringEqualStringPtr(hx_indexof_element_61, hx_indexof_value_58) {
				return hx_indexof_index_60
			}
			hx_indexof_index_60 = (hx_indexof_index_60 + 1)
		}
		return -1
	}()), hxrt.StdString(func() int {
		hx_lastindexof_target_65 := a
		var hx_lastindexof_value_66 *string = hxrt.StringFromLiteral("b")
		var hx_lastindexof_start_70 any = none
		hx_lastindexof_length_67 := hx_lastindexof_target_65.Len()
		hx_lastindexof_index_68 := (hx_lastindexof_length_67 - 1)
		if hx_lastindexof_start_70 != nil {
			hx_lastindexof_index_68 = hxrt.IntFromNullableAny(hx_lastindexof_start_70)
			if hx_lastindexof_index_68 < 0 {
				hx_lastindexof_index_68 = (hx_lastindexof_length_67 + hx_lastindexof_index_68)
			}
			if hx_lastindexof_index_68 >= hx_lastindexof_length_67 {
				hx_lastindexof_index_68 = (hx_lastindexof_length_67 - 1)
			}
		}
		for hx_lastindexof_index_68 >= 0 {
			hx_lastindexof_element_69 := func(hx_value_71 any) *string {
				if hx_value_71 == nil {
					var hx_zero_72 *string
					return hx_zero_72
				}
				return hx_value_71.(*string)
			}(hx_lastindexof_target_65.Get(hx_lastindexof_index_68))
			if hxrt.StringEqualStringPtr(hx_lastindexof_element_69, hx_lastindexof_value_66) {
				return hx_lastindexof_index_68
			}
			hx_lastindexof_index_68 = (hx_lastindexof_index_68 - 1)
		}
		return -1
	}()), hxrt.StringJoinAny(a.SliceOptional(1, none).Values(), hxrt.StringFromLiteral(",")), hxrt.StdString(func() int {
		hx_indexof_target_73 := a
		var hx_indexof_value_74 *string = hxrt.StringFromLiteral("b")
		var hx_indexof_start_78 any = some
		hx_indexof_length_75 := hx_indexof_target_73.Len()
		hx_indexof_index_76 := 0
		if hx_indexof_start_78 != nil {
			hx_indexof_index_76 = hxrt.IntFromNullableAny(hx_indexof_start_78)
			if hx_indexof_index_76 < 0 {
				hx_indexof_index_76 = (hx_indexof_length_75 + hx_indexof_index_76)
			}
			if hx_indexof_index_76 < 0 {
				hx_indexof_index_76 = 0
			}
		}
		for hx_indexof_index_76 < hx_indexof_length_75 {
			hx_indexof_element_77 := func(hx_value_79 any) *string {
				if hx_value_79 == nil {
					var hx_zero_80 *string
					return hx_zero_80
				}
				return hx_value_79.(*string)
			}(hx_indexof_target_73.Get(hx_indexof_index_76))
			if hxrt.StringEqualStringPtr(hx_indexof_element_77, hx_indexof_value_74) {
				return hx_indexof_index_76
			}
			hx_indexof_index_76 = (hx_indexof_index_76 + 1)
		}
		return -1
	}()), hxrt.StdString(func() int {
		hx_lastindexof_target_81 := a
		var hx_lastindexof_value_82 *string = hxrt.StringFromLiteral("b")
		var hx_lastindexof_start_86 any = some
		hx_lastindexof_length_83 := hx_lastindexof_target_81.Len()
		hx_lastindexof_index_84 := (hx_lastindexof_length_83 - 1)
		if hx_lastindexof_start_86 != nil {
			hx_lastindexof_index_84 = hxrt.IntFromNullableAny(hx_lastindexof_start_86)
			if hx_lastindexof_index_84 < 0 {
				hx_lastindexof_index_84 = (hx_lastindexof_length_83 + hx_lastindexof_index_84)
			}
			if hx_lastindexof_index_84 >= hx_lastindexof_length_83 {
				hx_lastindexof_index_84 = (hx_lastindexof_length_83 - 1)
			}
		}
		for hx_lastindexof_index_84 >= 0 {
			hx_lastindexof_element_85 := func(hx_value_87 any) *string {
				if hx_value_87 == nil {
					var hx_zero_88 *string
					return hx_zero_88
				}
				return hx_value_87.(*string)
			}(hx_lastindexof_target_81.Get(hx_lastindexof_index_84))
			if hxrt.StringEqualStringPtr(hx_lastindexof_element_85, hx_lastindexof_value_82) {
				return hx_lastindexof_index_84
			}
			hx_lastindexof_index_84 = (hx_lastindexof_index_84 - 1)
		}
		return -1
	}()), hxrt.StringJoinAny(a.SliceOptional(1, some).Values(), hxrt.StringFromLiteral(","))).Values(), hxrt.StringFromLiteral("|"))))
	hxrt.Println(v_13)
	bounds := New_Bounds(nil)
	var window_maybeFrom any
	window_maybeFrom = nil
	var v_14 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("null.field="), hxrt.StringJoinAny(hxrt.NewArray(hxrt.StdString(func() int {
		hx_indexof_target_89 := a
		var hx_indexof_value_90 *string = hxrt.StringFromLiteral("b")
		var hx_indexof_start_94 any = bounds.maybeFrom
		hx_indexof_length_91 := hx_indexof_target_89.Len()
		hx_indexof_index_92 := 0
		if hx_indexof_start_94 != nil {
			hx_indexof_index_92 = hxrt.IntFromNullableAny(hx_indexof_start_94)
			if hx_indexof_index_92 < 0 {
				hx_indexof_index_92 = (hx_indexof_length_91 + hx_indexof_index_92)
			}
			if hx_indexof_index_92 < 0 {
				hx_indexof_index_92 = 0
			}
		}
		for hx_indexof_index_92 < hx_indexof_length_91 {
			hx_indexof_element_93 := func(hx_value_95 any) *string {
				if hx_value_95 == nil {
					var hx_zero_96 *string
					return hx_zero_96
				}
				return hx_value_95.(*string)
			}(hx_indexof_target_89.Get(hx_indexof_index_92))
			if hxrt.StringEqualStringPtr(hx_indexof_element_93, hx_indexof_value_90) {
				return hx_indexof_index_92
			}
			hx_indexof_index_92 = (hx_indexof_index_92 + 1)
		}
		return -1
	}()), hxrt.StdString(func() int {
		hx_lastindexof_target_97 := a
		var hx_lastindexof_value_98 *string = hxrt.StringFromLiteral("b")
		var hx_lastindexof_start_102 any = Bounds_shared
		hx_lastindexof_length_99 := hx_lastindexof_target_97.Len()
		hx_lastindexof_index_100 := (hx_lastindexof_length_99 - 1)
		if hx_lastindexof_start_102 != nil {
			hx_lastindexof_index_100 = hxrt.IntFromNullableAny(hx_lastindexof_start_102)
			if hx_lastindexof_index_100 < 0 {
				hx_lastindexof_index_100 = (hx_lastindexof_length_99 + hx_lastindexof_index_100)
			}
			if hx_lastindexof_index_100 >= hx_lastindexof_length_99 {
				hx_lastindexof_index_100 = (hx_lastindexof_length_99 - 1)
			}
		}
		for hx_lastindexof_index_100 >= 0 {
			hx_lastindexof_element_101 := func(hx_value_103 any) *string {
				if hx_value_103 == nil {
					var hx_zero_104 *string
					return hx_zero_104
				}
				return hx_value_103.(*string)
			}(hx_lastindexof_target_97.Get(hx_lastindexof_index_100))
			if hxrt.StringEqualStringPtr(hx_lastindexof_element_101, hx_lastindexof_value_98) {
				return hx_lastindexof_index_100
			}
			hx_lastindexof_index_100 = (hx_lastindexof_index_100 - 1)
		}
		return -1
	}()), hxrt.StringJoinAny(a.SliceOptional(1, window_maybeFrom).Values(), hxrt.StringFromLiteral(","))).Values(), hxrt.StringFromLiteral("|"))))
	hxrt.Println(v_14)
	bounds.maybeFrom = 2
	Bounds_shared = 3
	window_maybeFrom = 3
	var v_15 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("value.field="), hxrt.StringJoinAny(hxrt.NewArray(hxrt.StdString(func() int {
		hx_indexof_target_105 := a
		var hx_indexof_value_106 *string = hxrt.StringFromLiteral("b")
		var hx_indexof_start_110 any = bounds.maybeFrom
		hx_indexof_length_107 := hx_indexof_target_105.Len()
		hx_indexof_index_108 := 0
		if hx_indexof_start_110 != nil {
			hx_indexof_index_108 = hxrt.IntFromNullableAny(hx_indexof_start_110)
			if hx_indexof_index_108 < 0 {
				hx_indexof_index_108 = (hx_indexof_length_107 + hx_indexof_index_108)
			}
			if hx_indexof_index_108 < 0 {
				hx_indexof_index_108 = 0
			}
		}
		for hx_indexof_index_108 < hx_indexof_length_107 {
			hx_indexof_element_109 := func(hx_value_111 any) *string {
				if hx_value_111 == nil {
					var hx_zero_112 *string
					return hx_zero_112
				}
				return hx_value_111.(*string)
			}(hx_indexof_target_105.Get(hx_indexof_index_108))
			if hxrt.StringEqualStringPtr(hx_indexof_element_109, hx_indexof_value_106) {
				return hx_indexof_index_108
			}
			hx_indexof_index_108 = (hx_indexof_index_108 + 1)
		}
		return -1
	}()), hxrt.StdString(func() int {
		hx_lastindexof_target_113 := a
		var hx_lastindexof_value_114 *string = hxrt.StringFromLiteral("b")
		var hx_lastindexof_start_118 any = Bounds_shared
		hx_lastindexof_length_115 := hx_lastindexof_target_113.Len()
		hx_lastindexof_index_116 := (hx_lastindexof_length_115 - 1)
		if hx_lastindexof_start_118 != nil {
			hx_lastindexof_index_116 = hxrt.IntFromNullableAny(hx_lastindexof_start_118)
			if hx_lastindexof_index_116 < 0 {
				hx_lastindexof_index_116 = (hx_lastindexof_length_115 + hx_lastindexof_index_116)
			}
			if hx_lastindexof_index_116 >= hx_lastindexof_length_115 {
				hx_lastindexof_index_116 = (hx_lastindexof_length_115 - 1)
			}
		}
		for hx_lastindexof_index_116 >= 0 {
			hx_lastindexof_element_117 := func(hx_value_119 any) *string {
				if hx_value_119 == nil {
					var hx_zero_120 *string
					return hx_zero_120
				}
				return hx_value_119.(*string)
			}(hx_lastindexof_target_113.Get(hx_lastindexof_index_116))
			if hxrt.StringEqualStringPtr(hx_lastindexof_element_117, hx_lastindexof_value_114) {
				return hx_lastindexof_index_116
			}
			hx_lastindexof_index_116 = (hx_lastindexof_index_116 - 1)
		}
		return -1
	}()), hxrt.StringJoinAny(a.SliceOptional(1, window_maybeFrom).Values(), hxrt.StringFromLiteral(","))).Values(), hxrt.StringFromLiteral("|"))))
	hxrt.Println(v_15)
	starts := hxrt.NewArray(nil, 2)
	i := 0
	var v_16 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("array="), hxrt.StringJoinAny(hxrt.NewArray(hxrt.StdString(func() int {
		hx_indexof_target_121 := a
		var hx_indexof_value_122 *string = hxrt.StringFromLiteral("b")
		var hx_indexof_start_126 any = starts.Get(0)
		hx_indexof_length_123 := hx_indexof_target_121.Len()
		hx_indexof_index_124 := 0
		if hx_indexof_start_126 != nil {
			hx_indexof_index_124 = hxrt.IntFromNullableAny(hx_indexof_start_126)
			if hx_indexof_index_124 < 0 {
				hx_indexof_index_124 = (hx_indexof_length_123 + hx_indexof_index_124)
			}
			if hx_indexof_index_124 < 0 {
				hx_indexof_index_124 = 0
			}
		}
		for hx_indexof_index_124 < hx_indexof_length_123 {
			hx_indexof_element_125 := func(hx_value_127 any) *string {
				if hx_value_127 == nil {
					var hx_zero_128 *string
					return hx_zero_128
				}
				return hx_value_127.(*string)
			}(hx_indexof_target_121.Get(hx_indexof_index_124))
			if hxrt.StringEqualStringPtr(hx_indexof_element_125, hx_indexof_value_122) {
				return hx_indexof_index_124
			}
			hx_indexof_index_124 = (hx_indexof_index_124 + 1)
		}
		return -1
	}()), hxrt.StdString(func() int {
		hx_lastindexof_target_129 := a
		var hx_lastindexof_value_130 *string = hxrt.StringFromLiteral("b")
		var hx_lastindexof_start_134 any = starts.Get(i)
		hx_lastindexof_length_131 := hx_lastindexof_target_129.Len()
		hx_lastindexof_index_132 := (hx_lastindexof_length_131 - 1)
		if hx_lastindexof_start_134 != nil {
			hx_lastindexof_index_132 = hxrt.IntFromNullableAny(hx_lastindexof_start_134)
			if hx_lastindexof_index_132 < 0 {
				hx_lastindexof_index_132 = (hx_lastindexof_length_131 + hx_lastindexof_index_132)
			}
			if hx_lastindexof_index_132 >= hx_lastindexof_length_131 {
				hx_lastindexof_index_132 = (hx_lastindexof_length_131 - 1)
			}
		}
		for hx_lastindexof_index_132 >= 0 {
			hx_lastindexof_element_133 := func(hx_value_135 any) *string {
				if hx_value_135 == nil {
					var hx_zero_136 *string
					return hx_zero_136
				}
				return hx_value_135.(*string)
			}(hx_lastindexof_target_129.Get(hx_lastindexof_index_132))
			if hxrt.StringEqualStringPtr(hx_lastindexof_element_133, hx_lastindexof_value_130) {
				return hx_lastindexof_index_132
			}
			hx_lastindexof_index_132 = (hx_lastindexof_index_132 - 1)
		}
		return -1
	}()), hxrt.StringJoinAny(a.SliceOptional(1, starts.Get(i)).Values(), hxrt.StringFromLiteral(",")), hxrt.StdString(func() int {
		hx_indexof_target_137 := a
		var hx_indexof_value_138 *string = hxrt.StringFromLiteral("b")
		var hx_indexof_start_142 any = starts.Get(int((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(1))))
		hx_indexof_length_139 := hx_indexof_target_137.Len()
		hx_indexof_index_140 := 0
		if hx_indexof_start_142 != nil {
			hx_indexof_index_140 = hxrt.IntFromNullableAny(hx_indexof_start_142)
			if hx_indexof_index_140 < 0 {
				hx_indexof_index_140 = (hx_indexof_length_139 + hx_indexof_index_140)
			}
			if hx_indexof_index_140 < 0 {
				hx_indexof_index_140 = 0
			}
		}
		for hx_indexof_index_140 < hx_indexof_length_139 {
			hx_indexof_element_141 := func(hx_value_143 any) *string {
				if hx_value_143 == nil {
					var hx_zero_144 *string
					return hx_zero_144
				}
				return hx_value_143.(*string)
			}(hx_indexof_target_137.Get(hx_indexof_index_140))
			if hxrt.StringEqualStringPtr(hx_indexof_element_141, hx_indexof_value_138) {
				return hx_indexof_index_140
			}
			hx_indexof_index_140 = (hx_indexof_index_140 + 1)
		}
		return -1
	}()), hxrt.StringJoinAny(a.SliceOptional(1, starts.Get(1)).Values(), hxrt.StringFromLiteral(",")), hxrt.StdString(func() int {
		hx_lastindexof_target_145 := a
		var hx_lastindexof_value_146 *string = hxrt.StringFromLiteral("b")
		var hx_lastindexof_start_150 any = ints.Get(1)
		hx_lastindexof_length_147 := hx_lastindexof_target_145.Len()
		hx_lastindexof_index_148 := (hx_lastindexof_length_147 - 1)
		if hx_lastindexof_start_150 != nil {
			hx_lastindexof_index_148 = hxrt.IntFromNullableAny(hx_lastindexof_start_150)
			if hx_lastindexof_index_148 < 0 {
				hx_lastindexof_index_148 = (hx_lastindexof_length_147 + hx_lastindexof_index_148)
			}
			if hx_lastindexof_index_148 >= hx_lastindexof_length_147 {
				hx_lastindexof_index_148 = (hx_lastindexof_length_147 - 1)
			}
		}
		for hx_lastindexof_index_148 >= 0 {
			hx_lastindexof_element_149 := func(hx_value_151 any) *string {
				if hx_value_151 == nil {
					var hx_zero_152 *string
					return hx_zero_152
				}
				return hx_value_151.(*string)
			}(hx_lastindexof_target_145.Get(hx_lastindexof_index_148))
			if hxrt.StringEqualStringPtr(hx_lastindexof_element_149, hx_lastindexof_value_146) {
				return hx_lastindexof_index_148
			}
			hx_lastindexof_index_148 = (hx_lastindexof_index_148 - 1)
		}
		return -1
	}()), hxrt.StdString(func() int {
		hx_indexof_target_153 := a
		var hx_indexof_value_154 *string = hxrt.StringFromLiteral("b")
		var hx_indexof_start_158 any = ints.Get(10)
		hx_indexof_length_155 := hx_indexof_target_153.Len()
		hx_indexof_index_156 := 0
		if hx_indexof_start_158 != nil {
			hx_indexof_index_156 = hxrt.IntFromNullableAny(hx_indexof_start_158)
			if hx_indexof_index_156 < 0 {
				hx_indexof_index_156 = (hx_indexof_length_155 + hx_indexof_index_156)
			}
			if hx_indexof_index_156 < 0 {
				hx_indexof_index_156 = 0
			}
		}
		for hx_indexof_index_156 < hx_indexof_length_155 {
			hx_indexof_element_157 := func(hx_value_159 any) *string {
				if hx_value_159 == nil {
					var hx_zero_160 *string
					return hx_zero_160
				}
				return hx_value_159.(*string)
			}(hx_indexof_target_153.Get(hx_indexof_index_156))
			if hxrt.StringEqualStringPtr(hx_indexof_element_157, hx_indexof_value_154) {
				return hx_indexof_index_156
			}
			hx_indexof_index_156 = (hx_indexof_index_156 + 1)
		}
		return -1
	}())).Values(), hxrt.StringFromLiteral("|"))))
	hxrt.Println(v_16)
	flag := (a.Len() > 3)
	var dynamicNull any = noInt()
	var dynamicSome any = someInt(3)
	var v_17 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("call="), hxrt.StringJoinAny(hxrt.NewArray(hxrt.StdString(func() int {
		hx_indexof_target_161 := a
		var hx_indexof_value_162 *string = hxrt.StringFromLiteral("b")
		var hx_indexof_start_166 any = noInt()
		hx_indexof_length_163 := hx_indexof_target_161.Len()
		hx_indexof_index_164 := 0
		if hx_indexof_start_166 != nil {
			hx_indexof_index_164 = hxrt.IntFromNullableAny(hx_indexof_start_166)
			if hx_indexof_index_164 < 0 {
				hx_indexof_index_164 = (hx_indexof_length_163 + hx_indexof_index_164)
			}
			if hx_indexof_index_164 < 0 {
				hx_indexof_index_164 = 0
			}
		}
		for hx_indexof_index_164 < hx_indexof_length_163 {
			hx_indexof_element_165 := func(hx_value_167 any) *string {
				if hx_value_167 == nil {
					var hx_zero_168 *string
					return hx_zero_168
				}
				return hx_value_167.(*string)
			}(hx_indexof_target_161.Get(hx_indexof_index_164))
			if hxrt.StringEqualStringPtr(hx_indexof_element_165, hx_indexof_value_162) {
				return hx_indexof_index_164
			}
			hx_indexof_index_164 = (hx_indexof_index_164 + 1)
		}
		return -1
	}()), hxrt.StdString(func() int {
		hx_lastindexof_target_169 := a
		var hx_lastindexof_value_170 *string = hxrt.StringFromLiteral("b")
		var hx_lastindexof_start_174 any = noInt()
		hx_lastindexof_length_171 := hx_lastindexof_target_169.Len()
		hx_lastindexof_index_172 := (hx_lastindexof_length_171 - 1)
		if hx_lastindexof_start_174 != nil {
			hx_lastindexof_index_172 = hxrt.IntFromNullableAny(hx_lastindexof_start_174)
			if hx_lastindexof_index_172 < 0 {
				hx_lastindexof_index_172 = (hx_lastindexof_length_171 + hx_lastindexof_index_172)
			}
			if hx_lastindexof_index_172 >= hx_lastindexof_length_171 {
				hx_lastindexof_index_172 = (hx_lastindexof_length_171 - 1)
			}
		}
		for hx_lastindexof_index_172 >= 0 {
			hx_lastindexof_element_173 := func(hx_value_175 any) *string {
				if hx_value_175 == nil {
					var hx_zero_176 *string
					return hx_zero_176
				}
				return hx_value_175.(*string)
			}(hx_lastindexof_target_169.Get(hx_lastindexof_index_172))
			if hxrt.StringEqualStringPtr(hx_lastindexof_element_173, hx_lastindexof_value_170) {
				return hx_lastindexof_index_172
			}
			hx_lastindexof_index_172 = (hx_lastindexof_index_172 - 1)
		}
		return -1
	}()), hxrt.StringJoinAny(a.SliceOptional(1, noInt()).Values(), hxrt.StringFromLiteral(",")), hxrt.StdString(func() int {
		hx_indexof_target_177 := a
		var hx_indexof_value_178 *string = hxrt.StringFromLiteral("b")
		var hx_indexof_start_182 any = someInt(2)
		hx_indexof_length_179 := hx_indexof_target_177.Len()
		hx_indexof_index_180 := 0
		if hx_indexof_start_182 != nil {
			hx_indexof_index_180 = hxrt.IntFromNullableAny(hx_indexof_start_182)
			if hx_indexof_index_180 < 0 {
				hx_indexof_index_180 = (hx_indexof_length_179 + hx_indexof_index_180)
			}
			if hx_indexof_index_180 < 0 {
				hx_indexof_index_180 = 0
			}
		}
		for hx_indexof_index_180 < hx_indexof_length_179 {
			hx_indexof_element_181 := func(hx_value_183 any) *string {
				if hx_value_183 == nil {
					var hx_zero_184 *string
					return hx_zero_184
				}
				return hx_value_183.(*string)
			}(hx_indexof_target_177.Get(hx_indexof_index_180))
			if hxrt.StringEqualStringPtr(hx_indexof_element_181, hx_indexof_value_178) {
				return hx_indexof_index_180
			}
			hx_indexof_index_180 = (hx_indexof_index_180 + 1)
		}
		return -1
	}()), hxrt.StdString(func() int {
		hx_lastindexof_target_186 := a
		var hx_lastindexof_value_187 *string = hxrt.StringFromLiteral("b")
		var hx_if_185 any
		if flag {
			hx_if_185 = nil
		} else {
			hx_if_185 = 1
		}
		var hx_lastindexof_start_191 any = hx_if_185
		hx_lastindexof_length_188 := hx_lastindexof_target_186.Len()
		hx_lastindexof_index_189 := (hx_lastindexof_length_188 - 1)
		if hx_lastindexof_start_191 != nil {
			hx_lastindexof_index_189 = hxrt.IntFromNullableAny(hx_lastindexof_start_191)
			if hx_lastindexof_index_189 < 0 {
				hx_lastindexof_index_189 = (hx_lastindexof_length_188 + hx_lastindexof_index_189)
			}
			if hx_lastindexof_index_189 >= hx_lastindexof_length_188 {
				hx_lastindexof_index_189 = (hx_lastindexof_length_188 - 1)
			}
		}
		for hx_lastindexof_index_189 >= 0 {
			hx_lastindexof_element_190 := func(hx_value_192 any) *string {
				if hx_value_192 == nil {
					var hx_zero_193 *string
					return hx_zero_193
				}
				return hx_value_192.(*string)
			}(hx_lastindexof_target_186.Get(hx_lastindexof_index_189))
			if hxrt.StringEqualStringPtr(hx_lastindexof_element_190, hx_lastindexof_value_187) {
				return hx_lastindexof_index_189
			}
			hx_lastindexof_index_189 = (hx_lastindexof_index_189 - 1)
		}
		return -1
	}()), hxrt.StringJoinAny(a.SliceOptional(1, func() any {
		var hx_if_194 any
		if flag {
			hx_if_194 = nil
		} else {
			hx_if_194 = 1
		}
		return hx_if_194
	}()).Values(), hxrt.StringFromLiteral(","))).Values(), hxrt.StringFromLiteral("|"))))
	hxrt.Println(v_17)
	var v_18 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("dynamic="), hxrt.StringJoinAny(hxrt.NewArray(hxrt.StdString(func() int {
		hx_indexof_target_195 := a
		var hx_indexof_value_196 *string = hxrt.StringFromLiteral("b")
		var hx_indexof_start_200 any = dynamicNull
		hx_indexof_length_197 := hx_indexof_target_195.Len()
		hx_indexof_index_198 := 0
		if hx_indexof_start_200 != nil {
			hx_indexof_index_198 = hxrt.IntFromNullableAny(hx_indexof_start_200)
			if hx_indexof_index_198 < 0 {
				hx_indexof_index_198 = (hx_indexof_length_197 + hx_indexof_index_198)
			}
			if hx_indexof_index_198 < 0 {
				hx_indexof_index_198 = 0
			}
		}
		for hx_indexof_index_198 < hx_indexof_length_197 {
			hx_indexof_element_199 := func(hx_value_201 any) *string {
				if hx_value_201 == nil {
					var hx_zero_202 *string
					return hx_zero_202
				}
				return hx_value_201.(*string)
			}(hx_indexof_target_195.Get(hx_indexof_index_198))
			if hxrt.StringEqualStringPtr(hx_indexof_element_199, hx_indexof_value_196) {
				return hx_indexof_index_198
			}
			hx_indexof_index_198 = (hx_indexof_index_198 + 1)
		}
		return -1
	}()), hxrt.StdString(func() int {
		hx_lastindexof_target_203 := a
		var hx_lastindexof_value_204 *string = hxrt.StringFromLiteral("b")
		var hx_lastindexof_start_208 any = dynamicSome
		hx_lastindexof_length_205 := hx_lastindexof_target_203.Len()
		hx_lastindexof_index_206 := (hx_lastindexof_length_205 - 1)
		if hx_lastindexof_start_208 != nil {
			hx_lastindexof_index_206 = hxrt.IntFromNullableAny(hx_lastindexof_start_208)
			if hx_lastindexof_index_206 < 0 {
				hx_lastindexof_index_206 = (hx_lastindexof_length_205 + hx_lastindexof_index_206)
			}
			if hx_lastindexof_index_206 >= hx_lastindexof_length_205 {
				hx_lastindexof_index_206 = (hx_lastindexof_length_205 - 1)
			}
		}
		for hx_lastindexof_index_206 >= 0 {
			hx_lastindexof_element_207 := func(hx_value_209 any) *string {
				if hx_value_209 == nil {
					var hx_zero_210 *string
					return hx_zero_210
				}
				return hx_value_209.(*string)
			}(hx_lastindexof_target_203.Get(hx_lastindexof_index_206))
			if hxrt.StringEqualStringPtr(hx_lastindexof_element_207, hx_lastindexof_value_204) {
				return hx_lastindexof_index_206
			}
			hx_lastindexof_index_206 = (hx_lastindexof_index_206 - 1)
		}
		return -1
	}()), hxrt.StringJoinAny(a.SliceOptional(1, dynamicNull).Values(), hxrt.StringFromLiteral(",")), hxrt.StdString(func() int {
		hx_indexof_target_211 := a
		var hx_indexof_value_212 *string = hxrt.StringFromLiteral("b")
		var hx_indexof_start_216 any = dynamicValue(nil)
		hx_indexof_length_213 := hx_indexof_target_211.Len()
		hx_indexof_index_214 := 0
		if hx_indexof_start_216 != nil {
			hx_indexof_index_214 = hxrt.IntFromNullableAny(hx_indexof_start_216)
			if hx_indexof_index_214 < 0 {
				hx_indexof_index_214 = (hx_indexof_length_213 + hx_indexof_index_214)
			}
			if hx_indexof_index_214 < 0 {
				hx_indexof_index_214 = 0
			}
		}
		for hx_indexof_index_214 < hx_indexof_length_213 {
			hx_indexof_element_215 := func(hx_value_217 any) *string {
				if hx_value_217 == nil {
					var hx_zero_218 *string
					return hx_zero_218
				}
				return hx_value_217.(*string)
			}(hx_indexof_target_211.Get(hx_indexof_index_214))
			if hxrt.StringEqualStringPtr(hx_indexof_element_215, hx_indexof_value_212) {
				return hx_indexof_index_214
			}
			hx_indexof_index_214 = (hx_indexof_index_214 + 1)
		}
		return -1
	}()), hxrt.StdString(func() int {
		hx_lastindexof_target_219 := a
		var hx_lastindexof_value_220 *string = hxrt.StringFromLiteral("b")
		var hx_lastindexof_start_224 any = dynamicValue(2)
		hx_lastindexof_length_221 := hx_lastindexof_target_219.Len()
		hx_lastindexof_index_222 := (hx_lastindexof_length_221 - 1)
		if hx_lastindexof_start_224 != nil {
			hx_lastindexof_index_222 = hxrt.IntFromNullableAny(hx_lastindexof_start_224)
			if hx_lastindexof_index_222 < 0 {
				hx_lastindexof_index_222 = (hx_lastindexof_length_221 + hx_lastindexof_index_222)
			}
			if hx_lastindexof_index_222 >= hx_lastindexof_length_221 {
				hx_lastindexof_index_222 = (hx_lastindexof_length_221 - 1)
			}
		}
		for hx_lastindexof_index_222 >= 0 {
			hx_lastindexof_element_223 := func(hx_value_225 any) *string {
				if hx_value_225 == nil {
					var hx_zero_226 *string
					return hx_zero_226
				}
				return hx_value_225.(*string)
			}(hx_lastindexof_target_219.Get(hx_lastindexof_index_222))
			if hxrt.StringEqualStringPtr(hx_lastindexof_element_223, hx_lastindexof_value_220) {
				return hx_lastindexof_index_222
			}
			hx_lastindexof_index_222 = (hx_lastindexof_index_222 - 1)
		}
		return -1
	}()), hxrt.StringJoinAny(a.SliceOptional(1, dynamicValue(3)).Values(), hxrt.StringFromLiteral(","))).Values(), hxrt.StringFromLiteral("|"))))
	hxrt.Println(v_18)
	order := hxrt.NewArray(hxrt.StdString(func() int {
		hx_indexof_target_227 := tickArray(hxrt.StringFromLiteral("a"), a)
		var hx_indexof_value_228 *string = tickText(hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("b"))
		var hx_indexof_start_232 any = tick(hxrt.StringFromLiteral("c"), nil)
		hx_indexof_length_229 := hx_indexof_target_227.Len()
		hx_indexof_index_230 := 0
		if hx_indexof_start_232 != nil {
			hx_indexof_index_230 = hxrt.IntFromNullableAny(hx_indexof_start_232)
			if hx_indexof_index_230 < 0 {
				hx_indexof_index_230 = (hx_indexof_length_229 + hx_indexof_index_230)
			}
			if hx_indexof_index_230 < 0 {
				hx_indexof_index_230 = 0
			}
		}
		for hx_indexof_index_230 < hx_indexof_length_229 {
			hx_indexof_element_231 := func(hx_value_233 any) *string {
				if hx_value_233 == nil {
					var hx_zero_234 *string
					return hx_zero_234
				}
				return hx_value_233.(*string)
			}(hx_indexof_target_227.Get(hx_indexof_index_230))
			if hxrt.StringEqualStringPtr(hx_indexof_element_231, hx_indexof_value_228) {
				return hx_indexof_index_230
			}
			hx_indexof_index_230 = (hx_indexof_index_230 + 1)
		}
		return -1
	}()), hxrt.StdString(func() int {
		hx_lastindexof_target_235 := tickArray(hxrt.StringFromLiteral("d"), a)
		var hx_lastindexof_value_236 *string = tickText(hxrt.StringFromLiteral("e"), hxrt.StringFromLiteral("b"))
		var hx_lastindexof_start_240 any = tick(hxrt.StringFromLiteral("f"), nil)
		hx_lastindexof_length_237 := hx_lastindexof_target_235.Len()
		hx_lastindexof_index_238 := (hx_lastindexof_length_237 - 1)
		if hx_lastindexof_start_240 != nil {
			hx_lastindexof_index_238 = hxrt.IntFromNullableAny(hx_lastindexof_start_240)
			if hx_lastindexof_index_238 < 0 {
				hx_lastindexof_index_238 = (hx_lastindexof_length_237 + hx_lastindexof_index_238)
			}
			if hx_lastindexof_index_238 >= hx_lastindexof_length_237 {
				hx_lastindexof_index_238 = (hx_lastindexof_length_237 - 1)
			}
		}
		for hx_lastindexof_index_238 >= 0 {
			hx_lastindexof_element_239 := func(hx_value_241 any) *string {
				if hx_value_241 == nil {
					var hx_zero_242 *string
					return hx_zero_242
				}
				return hx_value_241.(*string)
			}(hx_lastindexof_target_235.Get(hx_lastindexof_index_238))
			if hxrt.StringEqualStringPtr(hx_lastindexof_element_239, hx_lastindexof_value_236) {
				return hx_lastindexof_index_238
			}
			hx_lastindexof_index_238 = (hx_lastindexof_index_238 - 1)
		}
		return -1
	}()), hxrt.StringJoinAny(tickArray(hxrt.StringFromLiteral("g"), a).SliceOptional(tickInt(hxrt.StringFromLiteral("h"), 1), tick(hxrt.StringFromLiteral("i"), nil)).Values(), hxrt.StringFromLiteral(",")), hxrt.StdString(func() int {
		hx_indexof_target_243 := tickArray(hxrt.StringFromLiteral("j"), a)
		var hx_indexof_value_244 *string = tickText(hxrt.StringFromLiteral("k"), hxrt.StringFromLiteral("b"))
		var hx_indexof_start_248 any = tick(hxrt.StringFromLiteral("l"), 2)
		hx_indexof_length_245 := hx_indexof_target_243.Len()
		hx_indexof_index_246 := 0
		if hx_indexof_start_248 != nil {
			hx_indexof_index_246 = hxrt.IntFromNullableAny(hx_indexof_start_248)
			if hx_indexof_index_246 < 0 {
				hx_indexof_index_246 = (hx_indexof_length_245 + hx_indexof_index_246)
			}
			if hx_indexof_index_246 < 0 {
				hx_indexof_index_246 = 0
			}
		}
		for hx_indexof_index_246 < hx_indexof_length_245 {
			hx_indexof_element_247 := func(hx_value_249 any) *string {
				if hx_value_249 == nil {
					var hx_zero_250 *string
					return hx_zero_250
				}
				return hx_value_249.(*string)
			}(hx_indexof_target_243.Get(hx_indexof_index_246))
			if hxrt.StringEqualStringPtr(hx_indexof_element_247, hx_indexof_value_244) {
				return hx_indexof_index_246
			}
			hx_indexof_index_246 = (hx_indexof_index_246 + 1)
		}
		return -1
	}()), hxrt.StringJoinAny(tickArray(hxrt.StringFromLiteral("m"), a).SliceOptional(tickInt(hxrt.StringFromLiteral("n"), 1), tick(hxrt.StringFromLiteral("o"), 3)).Values(), hxrt.StringFromLiteral(",")))
	var v_19 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("order="), hxrt.StringJoinAny(order.Values(), hxrt.StringFromLiteral("|"))), hxrt.StringFromLiteral(":")), hxrt.StringJoinAny(trail.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_19)
}

func narrowed(a *hxrt.Array, from any) *string {
	if from != nil {
		return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatAny(func() int {
			hx_indexof_target_251 := a
			var hx_indexof_value_252 *string = hxrt.StringFromLiteral("b")
			hx_indexof_index_254 := from.(int)
			hx_indexof_length_253 := hx_indexof_target_251.Len()
			if hx_indexof_index_254 < 0 {
				hx_indexof_index_254 = (hx_indexof_length_253 + hx_indexof_index_254)
			}
			if hx_indexof_index_254 < 0 {
				hx_indexof_index_254 = 0
			}
			for hx_indexof_index_254 < hx_indexof_length_253 {
				hx_indexof_element_255 := func(hx_value_256 any) *string {
					if hx_value_256 == nil {
						var hx_zero_257 *string
						return hx_zero_257
					}
					return hx_value_256.(*string)
				}(hx_indexof_target_251.Get(hx_indexof_index_254))
				if hxrt.StringEqualStringPtr(hx_indexof_element_255, hx_indexof_value_252) {
					return hx_indexof_index_254
				}
				hx_indexof_index_254 = (hx_indexof_index_254 + 1)
			}
			return -1
		}(), hxrt.StringFromLiteral("|")), func() int {
			hx_lastindexof_target_258 := a
			var hx_lastindexof_value_259 *string = hxrt.StringFromLiteral("b")
			hx_lastindexof_index_261 := from.(int)
			hx_lastindexof_length_260 := hx_lastindexof_target_258.Len()
			if hx_lastindexof_index_261 < 0 {
				hx_lastindexof_index_261 = (hx_lastindexof_length_260 + hx_lastindexof_index_261)
			}
			if hx_lastindexof_index_261 >= hx_lastindexof_length_260 {
				hx_lastindexof_index_261 = (hx_lastindexof_length_260 - 1)
			}
			for hx_lastindexof_index_261 >= 0 {
				hx_lastindexof_element_262 := func(hx_value_263 any) *string {
					if hx_value_263 == nil {
						var hx_zero_264 *string
						return hx_zero_264
					}
					return hx_value_263.(*string)
				}(hx_lastindexof_target_258.Get(hx_lastindexof_index_261))
				if hxrt.StringEqualStringPtr(hx_lastindexof_element_262, hx_lastindexof_value_259) {
					return hx_lastindexof_index_261
				}
				hx_lastindexof_index_261 = (hx_lastindexof_index_261 - 1)
			}
			return -1
		}()), hxrt.StringFromLiteral("|")), hxrt.StringJoinAny(a.SliceOptional(1, from.(int)).Values(), hxrt.StringFromLiteral(",")))
	}
	return hxrt.StringFromLiteral("missing")
}

func noInt() any {
	return nil
}

func optional(a *hxrt.Array, from any) *string {
	return hxrt.StringJoinAny(hxrt.NewArray(hxrt.StdString(func() int {
		hx_indexof_target_265 := a
		var hx_indexof_value_266 *string = hxrt.StringFromLiteral("b")
		var hx_indexof_start_270 any = from
		hx_indexof_length_267 := hx_indexof_target_265.Len()
		hx_indexof_index_268 := 0
		if hx_indexof_start_270 != nil {
			hx_indexof_index_268 = hxrt.IntFromNullableAny(hx_indexof_start_270)
			if hx_indexof_index_268 < 0 {
				hx_indexof_index_268 = (hx_indexof_length_267 + hx_indexof_index_268)
			}
			if hx_indexof_index_268 < 0 {
				hx_indexof_index_268 = 0
			}
		}
		for hx_indexof_index_268 < hx_indexof_length_267 {
			hx_indexof_element_269 := func(hx_value_271 any) *string {
				if hx_value_271 == nil {
					var hx_zero_272 *string
					return hx_zero_272
				}
				return hx_value_271.(*string)
			}(hx_indexof_target_265.Get(hx_indexof_index_268))
			if hxrt.StringEqualStringPtr(hx_indexof_element_269, hx_indexof_value_266) {
				return hx_indexof_index_268
			}
			hx_indexof_index_268 = (hx_indexof_index_268 + 1)
		}
		return -1
	}()), hxrt.StdString(func() int {
		hx_lastindexof_target_273 := a
		var hx_lastindexof_value_274 *string = hxrt.StringFromLiteral("b")
		var hx_lastindexof_start_278 any = from
		hx_lastindexof_length_275 := hx_lastindexof_target_273.Len()
		hx_lastindexof_index_276 := (hx_lastindexof_length_275 - 1)
		if hx_lastindexof_start_278 != nil {
			hx_lastindexof_index_276 = hxrt.IntFromNullableAny(hx_lastindexof_start_278)
			if hx_lastindexof_index_276 < 0 {
				hx_lastindexof_index_276 = (hx_lastindexof_length_275 + hx_lastindexof_index_276)
			}
			if hx_lastindexof_index_276 >= hx_lastindexof_length_275 {
				hx_lastindexof_index_276 = (hx_lastindexof_length_275 - 1)
			}
		}
		for hx_lastindexof_index_276 >= 0 {
			hx_lastindexof_element_277 := func(hx_value_279 any) *string {
				if hx_value_279 == nil {
					var hx_zero_280 *string
					return hx_zero_280
				}
				return hx_value_279.(*string)
			}(hx_lastindexof_target_273.Get(hx_lastindexof_index_276))
			if hxrt.StringEqualStringPtr(hx_lastindexof_element_277, hx_lastindexof_value_274) {
				return hx_lastindexof_index_276
			}
			hx_lastindexof_index_276 = (hx_lastindexof_index_276 - 1)
		}
		return -1
	}()), hxrt.StringJoinAny(a.SliceOptional(1, from).Values(), hxrt.StringFromLiteral(","))).Values(), hxrt.StringFromLiteral("|"))
}

func someInt(value int) any {
	return value
}

func three(a *hxrt.Array, ints *hxrt.Array, value any) *string {
	return hxrt.StringJoinAny(hxrt.NewArray(hxrt.StdString(func() int {
		hx_indexof_target_281 := a
		var hx_indexof_value_282 *string = hxrt.StringFromLiteral("b")
		var hx_indexof_start_286 any = value
		hx_indexof_length_283 := hx_indexof_target_281.Len()
		hx_indexof_index_284 := 0
		if hx_indexof_start_286 != nil {
			hx_indexof_index_284 = hxrt.IntFromNullableAny(hx_indexof_start_286)
			if hx_indexof_index_284 < 0 {
				hx_indexof_index_284 = (hx_indexof_length_283 + hx_indexof_index_284)
			}
			if hx_indexof_index_284 < 0 {
				hx_indexof_index_284 = 0
			}
		}
		for hx_indexof_index_284 < hx_indexof_length_283 {
			hx_indexof_element_285 := func(hx_value_287 any) *string {
				if hx_value_287 == nil {
					var hx_zero_288 *string
					return hx_zero_288
				}
				return hx_value_287.(*string)
			}(hx_indexof_target_281.Get(hx_indexof_index_284))
			if hxrt.StringEqualStringPtr(hx_indexof_element_285, hx_indexof_value_282) {
				return hx_indexof_index_284
			}
			hx_indexof_index_284 = (hx_indexof_index_284 + 1)
		}
		return -1
	}()), hxrt.StdString(func() int {
		hx_lastindexof_target_289 := a
		var hx_lastindexof_value_290 *string = hxrt.StringFromLiteral("b")
		var hx_lastindexof_start_294 any = value
		hx_lastindexof_length_291 := hx_lastindexof_target_289.Len()
		hx_lastindexof_index_292 := (hx_lastindexof_length_291 - 1)
		if hx_lastindexof_start_294 != nil {
			hx_lastindexof_index_292 = hxrt.IntFromNullableAny(hx_lastindexof_start_294)
			if hx_lastindexof_index_292 < 0 {
				hx_lastindexof_index_292 = (hx_lastindexof_length_291 + hx_lastindexof_index_292)
			}
			if hx_lastindexof_index_292 >= hx_lastindexof_length_291 {
				hx_lastindexof_index_292 = (hx_lastindexof_length_291 - 1)
			}
		}
		for hx_lastindexof_index_292 >= 0 {
			hx_lastindexof_element_293 := func(hx_value_295 any) *string {
				if hx_value_295 == nil {
					var hx_zero_296 *string
					return hx_zero_296
				}
				return hx_value_295.(*string)
			}(hx_lastindexof_target_289.Get(hx_lastindexof_index_292))
			if hxrt.StringEqualStringPtr(hx_lastindexof_element_293, hx_lastindexof_value_290) {
				return hx_lastindexof_index_292
			}
			hx_lastindexof_index_292 = (hx_lastindexof_index_292 - 1)
		}
		return -1
	}()), hxrt.StringJoinAny(a.SliceOptional(1, value).Values(), hxrt.StringFromLiteral(",")), hxrt.StdString(func() int {
		hx_indexof_target_297 := ints
		var hx_indexof_value_298 int = 2
		var hx_indexof_start_302 any = value
		hx_indexof_length_299 := hx_indexof_target_297.Len()
		hx_indexof_index_300 := 0
		if hx_indexof_start_302 != nil {
			hx_indexof_index_300 = hxrt.IntFromNullableAny(hx_indexof_start_302)
			if hx_indexof_index_300 < 0 {
				hx_indexof_index_300 = (hx_indexof_length_299 + hx_indexof_index_300)
			}
			if hx_indexof_index_300 < 0 {
				hx_indexof_index_300 = 0
			}
		}
		for hx_indexof_index_300 < hx_indexof_length_299 {
			hx_indexof_element_301 := hxrt.IntFromNullableAny(hx_indexof_target_297.Get(hx_indexof_index_300))
			if hx_indexof_element_301 == hx_indexof_value_298 {
				return hx_indexof_index_300
			}
			hx_indexof_index_300 = (hx_indexof_index_300 + 1)
		}
		return -1
	}())).Values(), hxrt.StringFromLiteral("|"))
}

func tick(label *string, value any) any {
	hx_arr_303 := trail
	hx_arr_303.Push(label)
	return value
}

func tickArray(label *string, value *hxrt.Array) *hxrt.Array {
	hx_arr_304 := trail
	hx_arr_304.Push(label)
	return value
}

func tickInt(label *string, value int) int {
	hx_arr_305 := trail
	hx_arr_305.Push(label)
	return value
}

func tickText(label *string, value *string) *string {
	hx_arr_306 := trail
	hx_arr_306.Push(label)
	return value
}

var trail *hxrt.Array = hxrt.NewArray()
