package main

import "snapshot/hxrt"

func main() {
	var map_ haxe__IMap = New_haxe__ds__StringMap()
	map_.setIMap(hxrt.StringFromLiteral("alpha"), 1)
	var copied haxe__IMap = func(hx_value_1 any) haxe__IMap {
		if hx_value_1 == nil {
			var hx_zero_2 haxe__IMap
			return hx_zero_2
		}
		return hx_value_1.(haxe__IMap)
	}(map_.copyIMap())
	copied.setIMap(hxrt.StringFromLiteral("copied"), 7)
	var v any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("imap="), hxrt.StdString(func(hx_value_3 any) bool {
		if hx_value_3 == nil {
			var hx_zero_4 bool
			return hx_zero_4
		}
		return hx_value_3.(bool)
	}(map_.existsIMap(hxrt.StringFromLiteral("copied"))))), hxrt.StringFromLiteral(":")), hxrt.StdString(func(hx_value_5 any) bool {
		if hx_value_5 == nil {
			var hx_zero_6 bool
			return hx_zero_6
		}
		return hx_value_5.(bool)
	}(copied.existsIMap(hxrt.StringFromLiteral("copied"))))), hxrt.StringFromLiteral(":")), func(hx_value_7 any) any {
		if hx_value_7 == nil {
			return nil
		}
		return hx_value_7.(int)
	}(copied.getIMap(hxrt.StringFromLiteral("copied")))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("rest="), restDigest([]int{3, 1, 4})))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("rest.empty="), restDigest([]int{})))
	hxrt.Println(v_2)
}

func restDigest(args []int) *string {
	rest := args
	copied := hxrt.ArrayFromValues(func(hx_sort_src_8 []int) []any {
		hx_sort_out_10 := make([]any, 0, len(hx_sort_src_8))
		for _, hx_sort_item_9 := range hx_sort_src_8 {
			hx_sort_out_10 = append(hx_sort_out_10, hx_sort_item_9)
		}
		return hx_sort_out_10
	}(func(src []int) []int {
		out := append([]int{}, src...)
		return out
	}(rest)))
	appended := func(src []int, value int) []int {
		out := append([]int{}, src...)
		out = append(out, value)
		return out
	}(rest, 9)
	prepended := func(src []int, value int) []int {
		return append([]int{value}, src...)
	}(rest, -1)
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatAny(copied.Len(), hxrt.StringFromLiteral(":")), func() int {
		var hx_if_11 int
		if copied.Len() > 0 {
			hx_if_11 = hxrt.IntFromNullableAny(hxrt.IntFromNullableAny(copied.Get(0)))
		} else {
			hx_if_11 = -99
		}
		return hx_if_11
	}()), hxrt.StringFromLiteral(":")), func() int {
		var hx_if_12 int
		if copied.Len() > 0 {
			hx_if_12 = hxrt.IntFromNullableAny(hxrt.IntFromNullableAny(copied.Get(int(int32((hxrt.Int32Wrap(copied.Len()) - hxrt.Int32Wrap(1)))))))
		} else {
			hx_if_12 = -99
		}
		return hx_if_12
	}()), hxrt.StringFromLiteral("|append=")), appended[int(int32((hxrt.Int32Wrap(len(appended))-hxrt.Int32Wrap(1))))]), hxrt.StringFromLiteral("|prepend=")), prepended[0])
}
