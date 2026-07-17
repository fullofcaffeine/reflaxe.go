package main

import "snapshot/hxrt"

var haxe__Resource_content *hxrt.Array = hxrt.NewArray(map[string]any{"name": hxrt.StringFromLiteral("greet"), "data": hxrt.StringFromLiteral("aGVsbG8K"), "str": nil})

func haxe__Resource_getBytes(name *string) *haxe__io__Bytes {
	_g := 0
	_g1 := haxe__Resource_content
	for _g < _g1.Len() {
		x := func(hx_value_5 any) map[string]any {
			if hx_value_5 == nil {
				var hx_zero_6 map[string]any
				return hx_zero_6
			}
			return hx_value_5.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(func(hx_obj_16 map[string]any) *string {
			hx_field_17 := hx_obj_16["name"]
			if hx_field_17 == nil {
				var hx_zero_18 *string
				return hx_zero_18
			}
			return hx_field_17.(*string)
		}(x), name) {
			if !hxrt.StringEqualStringPtr(func(hx_obj_10 map[string]any) *string {
				hx_field_11 := hx_obj_10["str"]
				if hx_field_11 == nil {
					var hx_zero_12 *string
					return hx_zero_12
				}
				return hx_field_11.(*string)
			}(x), nil) {
				return haxe__io__Bytes_ofString(func(hx_obj_7 map[string]any) *string {
					hx_field_8 := hx_obj_7["str"]
					if hx_field_8 == nil {
						var hx_zero_9 *string
						return hx_zero_9
					}
					return hx_field_8.(*string)
				}(x))
			}
			return haxe__crypto__Base64_decode(func(hx_obj_13 map[string]any) *string {
				hx_field_14 := hx_obj_13["data"]
				if hx_field_14 == nil {
					var hx_zero_15 *string
					return hx_zero_15
				}
				return hx_field_14.(*string)
			}(x), true)
		}
	}
	return nil
}

func haxe__Resource_getString(name *string) *string {
	_g := 0
	_g1 := haxe__Resource_content
	for _g < _g1.Len() {
		x := func(hx_value_19 any) map[string]any {
			if hx_value_19 == nil {
				var hx_zero_20 map[string]any
				return hx_zero_20
			}
			return hx_value_19.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(func(hx_obj_30 map[string]any) *string {
			hx_field_31 := hx_obj_30["name"]
			if hx_field_31 == nil {
				var hx_zero_32 *string
				return hx_zero_32
			}
			return hx_field_31.(*string)
		}(x), name) {
			if !hxrt.StringEqualStringPtr(func(hx_obj_24 map[string]any) *string {
				hx_field_25 := hx_obj_24["str"]
				if hx_field_25 == nil {
					var hx_zero_26 *string
					return hx_zero_26
				}
				return hx_field_25.(*string)
			}(x), nil) {
				return func(hx_obj_21 map[string]any) *string {
					hx_field_22 := hx_obj_21["str"]
					if hx_field_22 == nil {
						var hx_zero_23 *string
						return hx_zero_23
					}
					return hx_field_22.(*string)
				}(x)
			}
			b := haxe__crypto__Base64_decode(func(hx_obj_27 map[string]any) *string {
				hx_field_28 := hx_obj_27["data"]
				if hx_field_28 == nil {
					var hx_zero_29 *string
					return hx_zero_29
				}
				return hx_field_28.(*string)
			}(x), true)
			return b.toString()
		}
	}
	return nil
}

func haxe__Resource_listNames() *hxrt.Array {
	_g := hxrt.NewArray()
	_g1 := 0
	_g2 := haxe__Resource_content
	for _g1 < _g2.Len() {
		x := func(hx_value_33 any) map[string]any {
			if hx_value_33 == nil {
				var hx_zero_34 map[string]any
				return hx_zero_34
			}
			return hx_value_33.(map[string]any)
		}(_g2.Get(_g1))
		_g1 = int(int32((_g1 + 1)))
		_g.Push(func(hx_obj_36 map[string]any) *string {
			hx_field_37 := hx_obj_36["name"]
			if hx_field_37 == nil {
				var hx_zero_38 *string
				return hx_zero_38
			}
			return hx_field_37.(*string)
		}(x))
	}
	return _g
}
