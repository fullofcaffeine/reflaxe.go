package main

import "snapshot/hxrt"

var haxe__Resource_content []map[string]any = []map[string]any{map[string]any{"name": hxrt.StringFromLiteral("greet"), "data": hxrt.StringFromLiteral("aGVsbG8K"), "str": nil}}

func haxe__Resource_getBytes(name *string) *haxe__io__Bytes {
	_g := 0
	_g1 := haxe__Resource_content
	for _g < len(_g1) {
		x := _g1[_g]
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(func(hx_obj_5 map[string]any) *string {
			hx_field_6 := hx_obj_5["name"]
			if hx_field_6 == nil {
				var hx_zero_7 *string
				return hx_zero_7
			}
			return hx_field_6.(*string)
		}(x), name) {
			if !hxrt.StringEqualStringPtr(func(hx_obj_8 map[string]any) *string {
				hx_field_9 := hx_obj_8["str"]
				if hx_field_9 == nil {
					var hx_zero_10 *string
					return hx_zero_10
				}
				return hx_field_9.(*string)
			}(x), nil) {
				return haxe__io__Bytes_ofString(func(hx_obj_11 map[string]any) *string {
					hx_field_12 := hx_obj_11["str"]
					if hx_field_12 == nil {
						var hx_zero_13 *string
						return hx_zero_13
					}
					return hx_field_12.(*string)
				}(x))
			}
			return haxe__crypto__Base64_decode(func(hx_obj_14 map[string]any) *string {
				hx_field_15 := hx_obj_14["data"]
				if hx_field_15 == nil {
					var hx_zero_16 *string
					return hx_zero_16
				}
				return hx_field_15.(*string)
			}(x))
		}
	}
	return nil
}

func haxe__Resource_getString(name *string) *string {
	_g := 0
	_g1 := haxe__Resource_content
	for _g < len(_g1) {
		x := _g1[_g]
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(func(hx_obj_17 map[string]any) *string {
			hx_field_18 := hx_obj_17["name"]
			if hx_field_18 == nil {
				var hx_zero_19 *string
				return hx_zero_19
			}
			return hx_field_18.(*string)
		}(x), name) {
			if !hxrt.StringEqualStringPtr(func(hx_obj_20 map[string]any) *string {
				hx_field_21 := hx_obj_20["str"]
				if hx_field_21 == nil {
					var hx_zero_22 *string
					return hx_zero_22
				}
				return hx_field_21.(*string)
			}(x), nil) {
				return func(hx_obj_23 map[string]any) *string {
					hx_field_24 := hx_obj_23["str"]
					if hx_field_24 == nil {
						var hx_zero_25 *string
						return hx_zero_25
					}
					return hx_field_24.(*string)
				}(x)
			}
			b := haxe__crypto__Base64_decode(func(hx_obj_26 map[string]any) *string {
				hx_field_27 := hx_obj_26["data"]
				if hx_field_27 == nil {
					var hx_zero_28 *string
					return hx_zero_28
				}
				return hx_field_27.(*string)
			}(x))
			return b.toString()
		}
	}
	return nil
}

func haxe__Resource_listNames() []*string {
	_g := []*string{}
	_g1 := 0
	_g2 := haxe__Resource_content
	for _g1 < len(_g2) {
		x := _g2[_g1]
		_g1 = int(int32((_g1 + 1)))
		_g = append(_g, func(hx_obj_29 map[string]any) *string {
			hx_field_30 := hx_obj_29["name"]
			if hx_field_30 == nil {
				var hx_zero_31 *string
				return hx_zero_31
			}
			return hx_field_30.(*string)
		}(x))
	}
	return _g
}
