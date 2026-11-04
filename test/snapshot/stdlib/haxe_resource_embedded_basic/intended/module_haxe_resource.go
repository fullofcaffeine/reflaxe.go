package main

import "snapshot/hxrt"

var haxe__Resource_content *hxrt.Array = hxrt.NewArray(map[string]any{"name": hxrt.StringFromLiteral("greet"), "data": hxrt.StringFromLiteral("aGVsbG8K"), "str": nil})

func haxe__Resource_getBytes(name *string) *haxe__io__Bytes {
	_g := 0
	_g1 := haxe__Resource_content
	for _g < _g1.Len() {
		x := hx_structural_recover(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(func(hx_value_7 any) *string {
			if hx_value_7 == nil {
				var hx_zero_8 *string
				return hx_zero_8
			}
			return hx_value_7.(*string)
		}(hxrt.AnonymousGet(x, "name")), name) {
			if !hxrt.StringEqualStringPtr(func(hx_value_3 any) *string {
				if hx_value_3 == nil {
					var hx_zero_4 *string
					return hx_zero_4
				}
				return hx_value_3.(*string)
			}(hxrt.AnonymousGet(x, "str")), nil) {
				return haxe__io__Bytes_ofString(func(hx_value_1 any) *string {
					if hx_value_1 == nil {
						var hx_zero_2 *string
						return hx_zero_2
					}
					return hx_value_1.(*string)
				}(hxrt.AnonymousGet(x, "str")), nil)
			}
			return haxe__crypto__Base64_decode(func(hx_value_5 any) *string {
				if hx_value_5 == nil {
					var hx_zero_6 *string
					return hx_zero_6
				}
				return hx_value_5.(*string)
			}(hxrt.AnonymousGet(x, "data")), true)
		}
	}
	return nil
}

func haxe__Resource_getString(name *string) *string {
	_g := 0
	_g1 := haxe__Resource_content
	for _g < _g1.Len() {
		x := hx_structural_recover(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(func(hx_value_15 any) *string {
			if hx_value_15 == nil {
				var hx_zero_16 *string
				return hx_zero_16
			}
			return hx_value_15.(*string)
		}(hxrt.AnonymousGet(x, "name")), name) {
			if !hxrt.StringEqualStringPtr(func(hx_value_11 any) *string {
				if hx_value_11 == nil {
					var hx_zero_12 *string
					return hx_zero_12
				}
				return hx_value_11.(*string)
			}(hxrt.AnonymousGet(x, "str")), nil) {
				return func(hx_value_9 any) *string {
					if hx_value_9 == nil {
						var hx_zero_10 *string
						return hx_zero_10
					}
					return hx_value_9.(*string)
				}(hxrt.AnonymousGet(x, "str"))
			}
			b := haxe__crypto__Base64_decode(func(hx_value_13 any) *string {
				if hx_value_13 == nil {
					var hx_zero_14 *string
					return hx_zero_14
				}
				return hx_value_13.(*string)
			}(hxrt.AnonymousGet(x, "data")), true)
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
		x := hx_structural_recover(_g2.Get(_g1))
		_g1 = int(int32((_g1 + 1)))
		_g.Push(func(hx_value_18 any) *string {
			if hx_value_18 == nil {
				var hx_zero_19 *string
				return hx_zero_19
			}
			return hx_value_18.(*string)
		}(hxrt.AnonymousGet(x, "name")))
	}
	return _g
}
