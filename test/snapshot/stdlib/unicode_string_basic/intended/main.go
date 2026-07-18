package main

import "snapshot/hxrt"

func main() {
	var s any = hxrt.StringFromLiteral("a😀bé")
	var v any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("length="), hxrt.StringLengthStringPtr(hxrt.StdString(func(hx_value_1 any) *string {
		if hx_value_1 == nil {
			var hx_zero_2 *string
			return hx_zero_2
		}
		return hx_value_1.(*string)
	}(s)))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("char1="), _UnicodeString__UnicodeString_Impl__charAt(hxrt.StdString(s), 1)))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("code1="), _UnicodeString__UnicodeString_Impl__charCodeAt(hxrt.StdString(s), 1)))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("substr="), _UnicodeString__UnicodeString_Impl__substr(hxrt.StdString(s), 1, 2)))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("substring="), _UnicodeString__UnicodeString_Impl__substring(hxrt.StdString(s), 1, 3)))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("substring.swap="), _UnicodeString__UnicodeString_Impl__substring(hxrt.StdString(s), 3, 1)))
	hxrt.Println(v_5)
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("substring.neg="), _UnicodeString__UnicodeString_Impl__substring(hxrt.StdString(s), -2, 2)))
	hxrt.Println(v_6)
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("substring.omit="), _UnicodeString__UnicodeString_Impl__substring(hxrt.StdString(s), 2, nil)))
	hxrt.Println(v_7)
	var v_8 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("substr.neglen="), _UnicodeString__UnicodeString_Impl__substr(hxrt.StdString(s), 1, -1)))
	hxrt.Println(v_8)
	var v_9 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("substr.negpos="), _UnicodeString__UnicodeString_Impl__substr(hxrt.StdString(s), -2, 2)))
	hxrt.Println(v_9)
	var v_10 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("index="), _UnicodeString__UnicodeString_Impl__indexOf(hxrt.StdString(s), hxrt.StringFromLiteral("bé"), nil)))
	hxrt.Println(v_10)
	var v_11 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("index.empty="), _UnicodeString__UnicodeString_Impl__indexOf(hxrt.StdString(s), hxrt.StringFromLiteral(""), nil)))
	hxrt.Println(v_11)
	var v_12 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("index.startNeg="), _UnicodeString__UnicodeString_Impl__indexOf(hxrt.StdString(s), hxrt.StringFromLiteral("bé"), -2)))
	hxrt.Println(v_12)
	var v_13 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("last="), _UnicodeString__UnicodeString_Impl__lastIndexOf(hxrt.StdString(s), hxrt.StringFromLiteral("a"), nil)))
	hxrt.Println(v_13)
	var v_14 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("last.empty="), _UnicodeString__UnicodeString_Impl__lastIndexOf(hxrt.StdString(s), hxrt.StringFromLiteral(""), nil)))
	hxrt.Println(v_14)
	var v_15 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("last.start="), _UnicodeString__UnicodeString_Impl__lastIndexOf(hxrt.StdString(s), hxrt.StringFromLiteral("bé"), 2)))
	hxrt.Println(v_15)
	var left any = hxrt.StringFromLiteral("a😀")
	var right any = hxrt.StringFromLiteral("bé")
	var v_16 any = func() any {
		var a any = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("operators="), hxrt.StdString(_UnicodeString__UnicodeString_Impl__lessThan(left, right))), hxrt.StringFromLiteral("|")), hxrt.StdString(_UnicodeString__UnicodeString_Impl__lessThanOrEqual(left, right))), hxrt.StringFromLiteral("|")), hxrt.StdString(_UnicodeString__UnicodeString_Impl__greaterThan(right, left))), hxrt.StringFromLiteral("|")), hxrt.StdString(_UnicodeString__UnicodeString_Impl__greaterThanOrEqual(right, left))), hxrt.StringFromLiteral("|")), hxrt.StdString(hxrt.StringEqualStringPtr(func(hx_value_3 any) *string {
			if hx_value_3 == nil {
				var hx_zero_4 *string
				return hx_zero_4
			}
			return hx_value_3.(*string)
		}(left), func(hx_value_5 any) *string {
			if hx_value_5 == nil {
				var hx_zero_6 *string
				return hx_zero_6
			}
			return hx_value_5.(*string)
		}(left)))), hxrt.StringFromLiteral("|")), hxrt.StdString(!hxrt.StringEqualStringPtr(func(hx_value_7 any) *string {
			if hx_value_7 == nil {
				var hx_zero_8 *string
				return hx_zero_8
			}
			return hx_value_7.(*string)
		}(left), func(hx_value_9 any) *string {
			if hx_value_9 == nil {
				var hx_zero_10 *string
				return hx_zero_10
			}
			return hx_value_9.(*string)
		}(right)))), hxrt.StringFromLiteral("|"))
		return hxrt.StringConcatStringPtr(func(hx_value_11 any) *string {
			if hx_value_11 == nil {
				var hx_zero_12 *string
				return hx_zero_12
			}
			return hx_value_11.(*string)
		}(a), func(hx_value_17 any) *string {
			if hx_value_17 == nil {
				var hx_zero_18 *string
				return hx_zero_18
			}
			return hx_value_17.(*string)
		}(any(hxrt.StringConcatStringPtr(func(hx_value_13 any) *string {
			if hx_value_13 == nil {
				var hx_zero_14 *string
				return hx_zero_14
			}
			return hx_value_13.(*string)
		}(left), func(hx_value_15 any) *string {
			if hx_value_15 == nil {
				var hx_zero_16 *string
				return hx_zero_16
			}
			return hx_value_15.(*string)
		}(right)))))
	}()
	hxrt.Println(v_16)
	hxrt.Println(any(hxrt.StringConcatStringPtr(func(hx_value_25 any) *string {
		if hx_value_25 == nil {
			var hx_zero_26 *string
			return hx_zero_26
		}
		return hx_value_25.(*string)
	}(any(hxrt.StringConcatStringPtr(func(hx_value_23 any) *string {
		if hx_value_23 == nil {
			var hx_zero_24 *string
			return hx_zero_24
		}
		return hx_value_23.(*string)
	}(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("mixed="), func(hx_value_21 any) *string {
		if hx_value_21 == nil {
			var hx_zero_22 *string
			return hx_zero_22
		}
		return hx_value_21.(*string)
	}(any(hxrt.StringConcatStringPtr(func(hx_value_19 any) *string {
		if hx_value_19 == nil {
			var hx_zero_20 *string
			return hx_zero_20
		}
		return hx_value_19.(*string)
	}(left), hxrt.StringFromLiteral("x"))))))), hxrt.StringFromLiteral("|")))), func(hx_value_29 any) *string {
		if hx_value_29 == nil {
			var hx_zero_30 *string
			return hx_zero_30
		}
		return hx_value_29.(*string)
	}(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("x"), func(hx_value_27 any) *string {
		if hx_value_27 == nil {
			var hx_zero_28 *string
			return hx_zero_28
		}
		return hx_value_27.(*string)
	}(left)))))))
	var assigned any = left
	assigned = any(hxrt.StringConcatStringPtr(func(hx_value_31 any) *string {
		if hx_value_31 == nil {
			var hx_zero_32 *string
			return hx_zero_32
		}
		return hx_value_31.(*string)
	}(assigned), func(hx_value_33 any) *string {
		if hx_value_33 == nil {
			var hx_zero_34 *string
			return hx_zero_34
		}
		return hx_value_33.(*string)
	}(right)))
	assigned = any(hxrt.StringConcatStringPtr(func(hx_value_35 any) *string {
		if hx_value_35 == nil {
			var hx_zero_36 *string
			return hx_zero_36
		}
		return hx_value_35.(*string)
	}(assigned), hxrt.StringFromLiteral("x")))
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("assigned="), func(hx_value_37 any) *string {
		if hx_value_37 == nil {
			var hx_zero_38 *string
			return hx_zero_38
		}
		return hx_value_37.(*string)
	}(assigned))))
	var v_17 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("valid.utf8="), hxrt.StdString(_UnicodeString__UnicodeString_Impl__validate(haxe__io__Bytes_ofString(hxrt.StringFromLiteral("ok"), nil), haxe__io__Encoding_UTF8))))
	hxrt.Println(v_17)
	var v_18 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("valid.invalid="), hxrt.StdString(_UnicodeString__UnicodeString_Impl__validate(haxe__io__Bytes_ofHex(hxrt.StringFromLiteral("ff")), haxe__io__Encoding_UTF8))))
	hxrt.Println(v_18)
	hxrt.TryCatch(func() {
		_UnicodeString__UnicodeString_Impl__validate(haxe__io__Bytes_ofString(hxrt.StringFromLiteral("ok"), nil), haxe__io__Encoding_RawNative)
		hxrt.Println(any(hxrt.StringFromLiteral("valid.raw=ok")))
	}, func(hx_caught_39 any) {
		error := hx_caught_39
		var v_19 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("valid.raw="), hxrt.StdString(error)))
		hxrt.Println(v_19)
	})
}
