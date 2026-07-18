package main

import "snapshot/hxrt"

var haxe__Resource_content *hxrt.Array = hxrt.NewArray(map[string]any{"name": hxrt.StringFromLiteral("greet"), "data": hxrt.StringFromLiteral("aGVsbG8K"), "str": nil})

func haxe__Resource_getBytes(name *string) *haxe__io__Bytes {
	_g := 0
	_g1 := haxe__Resource_content
	for _g < _g1.Len() {
		x := func(hx_value_15 any) map[string]any {
			if hx_value_15 == nil {
				var hx_zero_16 map[string]any
				return hx_zero_16
			}
			return hx_value_15.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(func(hx_obj_26 map[string]any) *string {
			hx_field_27 := hx_obj_26["name"]
			if hx_field_27 == nil {
				var hx_zero_28 *string
				return hx_zero_28
			}
			return hx_field_27.(*string)
		}(x), name) {
			if !hxrt.StringEqualStringPtr(func(hx_obj_20 map[string]any) *string {
				hx_field_21 := hx_obj_20["str"]
				if hx_field_21 == nil {
					var hx_zero_22 *string
					return hx_zero_22
				}
				return hx_field_21.(*string)
			}(x), nil) {
				return haxe__io__Bytes_ofString(func(hx_obj_17 map[string]any) *string {
					hx_field_18 := hx_obj_17["str"]
					if hx_field_18 == nil {
						var hx_zero_19 *string
						return hx_zero_19
					}
					return hx_field_18.(*string)
				}(x), nil)
			}
			return haxe__crypto__Base64_decode(func(hx_obj_23 map[string]any) *string {
				hx_field_24 := hx_obj_23["data"]
				if hx_field_24 == nil {
					var hx_zero_25 *string
					return hx_zero_25
				}
				return hx_field_24.(*string)
			}(x), true)
		}
	}
	return nil
}

func haxe__Resource_getString(name *string) *string {
	_g := 0
	_g1 := haxe__Resource_content
	for _g < _g1.Len() {
		x := func(hx_value_29 any) map[string]any {
			if hx_value_29 == nil {
				var hx_zero_30 map[string]any
				return hx_zero_30
			}
			return hx_value_29.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(func(hx_obj_40 map[string]any) *string {
			hx_field_41 := hx_obj_40["name"]
			if hx_field_41 == nil {
				var hx_zero_42 *string
				return hx_zero_42
			}
			return hx_field_41.(*string)
		}(x), name) {
			if !hxrt.StringEqualStringPtr(func(hx_obj_34 map[string]any) *string {
				hx_field_35 := hx_obj_34["str"]
				if hx_field_35 == nil {
					var hx_zero_36 *string
					return hx_zero_36
				}
				return hx_field_35.(*string)
			}(x), nil) {
				return func(hx_obj_31 map[string]any) *string {
					hx_field_32 := hx_obj_31["str"]
					if hx_field_32 == nil {
						var hx_zero_33 *string
						return hx_zero_33
					}
					return hx_field_32.(*string)
				}(x)
			}
			b := haxe__crypto__Base64_decode(func(hx_obj_37 map[string]any) *string {
				hx_field_38 := hx_obj_37["data"]
				if hx_field_38 == nil {
					var hx_zero_39 *string
					return hx_zero_39
				}
				return hx_field_38.(*string)
			}(x), true)
			return b.__hx_this.toString()
		}
	}
	return nil
}

func haxe__Resource_listNames() *hxrt.Array {
	_g := hxrt.NewArray()
	_g1 := 0
	_g2 := haxe__Resource_content
	for _g1 < _g2.Len() {
		x := func(hx_value_43 any) map[string]any {
			if hx_value_43 == nil {
				var hx_zero_44 map[string]any
				return hx_zero_44
			}
			return hx_value_43.(map[string]any)
		}(_g2.Get(_g1))
		_g1 = int(int32((_g1 + 1)))
		_g.Push(func(hx_obj_46 map[string]any) *string {
			hx_field_47 := hx_obj_46["name"]
			if hx_field_47 == nil {
				var hx_zero_48 *string
				return hx_zero_48
			}
			return hx_field_47.(*string)
		}(x))
	}
	return _g
}
