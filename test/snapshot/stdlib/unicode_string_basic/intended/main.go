package main

import "snapshot/hxrt"

func main() {
	s := hxrt.StringFromLiteral("a😀bé")
	var v any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("length="), hxrt.StringLengthStringPtr(s)))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("char1="), _UnicodeString__UnicodeString_Impl__charAt(s, 1)))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("code1="), _UnicodeString__UnicodeString_Impl__charCodeAt(s, 1)))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("substr="), _UnicodeString__UnicodeString_Impl__substr(s, 1, 2)))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("substring="), _UnicodeString__UnicodeString_Impl__substring(s, 1, 3)))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("substring.swap="), _UnicodeString__UnicodeString_Impl__substring(s, 3, 1)))
	hxrt.Println(v_5)
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("substring.neg="), _UnicodeString__UnicodeString_Impl__substring(s, -2, 2)))
	hxrt.Println(v_6)
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("substring.omit="), _UnicodeString__UnicodeString_Impl__substring(s, 2, nil)))
	hxrt.Println(v_7)
	var v_8 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("substr.neglen="), _UnicodeString__UnicodeString_Impl__substr(s, 1, -1)))
	hxrt.Println(v_8)
	var v_9 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("substr.negpos="), _UnicodeString__UnicodeString_Impl__substr(s, -2, 2)))
	hxrt.Println(v_9)
	var v_10 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("index="), _UnicodeString__UnicodeString_Impl__indexOf(s, hxrt.StringFromLiteral("bé"), nil)))
	hxrt.Println(v_10)
	var v_11 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("index.empty="), _UnicodeString__UnicodeString_Impl__indexOf(s, hxrt.StringFromLiteral(""), nil)))
	hxrt.Println(v_11)
	var v_12 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("index.startNeg="), _UnicodeString__UnicodeString_Impl__indexOf(s, hxrt.StringFromLiteral("bé"), -2)))
	hxrt.Println(v_12)
	var v_13 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("last="), _UnicodeString__UnicodeString_Impl__lastIndexOf(s, hxrt.StringFromLiteral("a"), nil)))
	hxrt.Println(v_13)
	var v_14 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("last.empty="), _UnicodeString__UnicodeString_Impl__lastIndexOf(s, hxrt.StringFromLiteral(""), nil)))
	hxrt.Println(v_14)
	var v_15 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("last.start="), _UnicodeString__UnicodeString_Impl__lastIndexOf(s, hxrt.StringFromLiteral("bé"), 2)))
	hxrt.Println(v_15)
	left := hxrt.StringFromLiteral("a😀")
	right := hxrt.StringFromLiteral("bé")
	var v_16 any = any(func() *string {
		a := hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("operators="), hxrt.StdString(_UnicodeString__UnicodeString_Impl__lessThan(left, right))), hxrt.StringFromLiteral("|")), hxrt.StdString(_UnicodeString__UnicodeString_Impl__lessThanOrEqual(left, right))), hxrt.StringFromLiteral("|")), hxrt.StdString(_UnicodeString__UnicodeString_Impl__greaterThan(right, left))), hxrt.StringFromLiteral("|")), hxrt.StdString(_UnicodeString__UnicodeString_Impl__greaterThanOrEqual(right, left))), hxrt.StringFromLiteral("|")), hxrt.StdString(hxrt.StringEqualStringPtr(left, left))), hxrt.StringFromLiteral("|")), hxrt.StdString(!hxrt.StringEqualStringPtr(left, right))), hxrt.StringFromLiteral("|"))
		return hxrt.StringConcatStringPtr(a, hxrt.StringConcatStringPtr(left, right))
	}())
	hxrt.Println(v_16)
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("mixed="), hxrt.StringConcatStringPtr(left, hxrt.StringFromLiteral("x"))), hxrt.StringFromLiteral("|")), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("x"), left))))
	assigned := left
	assigned = hxrt.StringConcatStringPtr(assigned, right)
	assigned = hxrt.StringConcatStringPtr(assigned, hxrt.StringFromLiteral("x"))
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("assigned="), assigned)))
	var v_17 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("valid.utf8="), hxrt.StdString(_UnicodeString__UnicodeString_Impl__validate(haxe__io__Bytes_ofString(hxrt.StringFromLiteral("ok"), nil), haxe__io__Encoding_UTF8))))
	hxrt.Println(v_17)
	var v_18 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("valid.invalid="), hxrt.StdString(_UnicodeString__UnicodeString_Impl__validate(haxe__io__Bytes_ofHex(hxrt.StringFromLiteral("ff")), haxe__io__Encoding_UTF8))))
	hxrt.Println(v_18)
	hxrt.TryCatch(func() {
		_UnicodeString__UnicodeString_Impl__validate(haxe__io__Bytes_ofString(hxrt.StringFromLiteral("ok"), nil), haxe__io__Encoding_RawNative)
		hxrt.Println(any(hxrt.StringFromLiteral("valid.raw=ok")))
	}, func(hx_caught_1 any) {
		hx_exception_2 := haxe__Exception_caught(hx_caught_1)
		hx_payload_3 := haxe__Exception_catchValue(hx_caught_1)
		_ = hx_exception_2
		var error any = hx_payload_3
		var v_19 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("valid.raw="), hxrt.StdString(error)))
		hxrt.Println(v_19)
	})
}

func hx_recover_class_haxe__Exception(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__Exception:
		return source
	case *haxe__ValueException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}

func hx_recover_class_haxe__ValueException(value any) *haxe__ValueException {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ValueException:
		return source
	default:
		return value.(*haxe__ValueException)
	}
}

func hx_is_nominal_haxe__Exception(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__Exception:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__ValueException(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}
