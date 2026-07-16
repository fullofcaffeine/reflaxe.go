package main

import "snapshot/hxrt"

func main() {
	map_ := New_haxe__ds__StringMap()
	map_.set(hxrt.StringFromLiteral("alpha"), 2)
	map_.set(hxrt.StringFromLiteral("beta"), 5)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("map.alpha="), hxrt.StdString(func(hx_value_1 any) any {
		if hx_value_1 == nil {
			return nil
		}
		return hx_value_1.(int)
	}(map_.get(hxrt.StringFromLiteral("alpha"))))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("map.exists.beta0="), hxrt.StdString(func(hx_value_2 any) bool {
		if hx_value_2 == nil {
			var hx_zero_3 bool
			return hx_zero_3
		}
		return hx_value_2.(bool)
	}(map_.exists(hxrt.StringFromLiteral("beta"))))))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("map.remove.beta="), hxrt.StdString(func(hx_value_4 any) bool {
		if hx_value_4 == nil {
			var hx_zero_5 bool
			return hx_zero_5
		}
		return hx_value_4.(bool)
	}(map_.remove(hxrt.StringFromLiteral("beta"))))))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("map.exists.beta1="), hxrt.StdString(func(hx_value_6 any) bool {
		if hx_value_6 == nil {
			var hx_zero_7 bool
			return hx_zero_7
		}
		return hx_value_6.(bool)
	}(map_.exists(hxrt.StringFromLiteral("beta"))))))
	hxrt.Println(v_3)
}
