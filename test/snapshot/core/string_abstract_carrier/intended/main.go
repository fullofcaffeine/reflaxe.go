package main

import "snapshot/hxrt"

func main() {
	token := _Main__Token_Impl__make(hxrt.StringFromLiteral("exact bytes"))
	var this1 *string
	this1 = _Main__Token_Impl__make(hxrt.StringFromLiteral("nested"))
	nested := this1
	var absent *string = nil
	nullBacked := _Main__Token_Impl__make(nil)
	var this1_1 int
	this1_1 = 7
	var count any = any(this1_1)
	var this1_2 *string
	this1_2 = hxrt.StringFromLiteral("applied")
	boxed := this1_2
	var this1_3 *string
	this1_3 = hxrt.StringFromLiteral("repeated")
	value := this1_3
	var this1_4 *string
	this1_4 = value
	repeated := this1_4
	if !hxrt.StringEqualStringPtr(repeated, hxrt.StringFromLiteral("repeated")) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("repeated applied backing changed")))
	}
	var value_1 *string = nil
	var this1_5 *string
	this1_5 = value_1
	nullableBox := this1_5
	if !hxrt.StringEqualStringPtr(nullableBox, nil) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("nullable alias changed")))
	}
	var this1_6 int
	this1_6 = 9
	var boxedInt any = any(this1_6)
	if !hxrt.StringEqualStringPtr(boxed, hxrt.StringFromLiteral("applied")) || (func(hx_value_1 any) int {
		if hx_value_1 == nil {
			var hx_zero_2 int
			return hx_zero_2
		}
		return hx_value_1.(int)
	}(boxedInt) != 9) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("applied backing changed")))
	}
	if !hxrt.StringEqualStringPtr(token, hxrt.StringFromLiteral("exact bytes")) || !hxrt.StringEqualStringPtr(nested, hxrt.StringFromLiteral("nested")) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("string carrier changed")))
	}
	if (absent != nil) || !hxrt.StringEqualStringPtr(nullBacked, nil) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("nil string changed")))
	}
	if func(hx_value_3 any) int {
		if hx_value_3 == nil {
			var hx_zero_4 int
			return hx_zero_4
		}
		return hx_value_3.(int)
	}(count) != 7 {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("non-string fallback changed")))
	}
	hxrt.Println(any(hxrt.StringFromLiteral("string abstract carrier: PASS")))
}

func _Main__Box_Impl___new(value any) any {
	var this1 any
	this1 = value
	return this1
}

func _Main__Box_Impl__value(this1 any) any {
	return this1
}

func _Main__Count_Impl___new(raw int) any {
	var this1 int
	this1 = raw
	return any(this1)
}

func _Main__Count_Impl__value(this1 int) int {
	return this1
}

func _Main__Tagged_Impl___new(raw *string) *string {
	var this1 *string
	this1 = _Main__Token_Impl__make(raw)
	return this1
}

func _Main__Tagged_Impl__text(this1 *string) *string {
	return this1
}

func _Main__Token_Impl___new(raw *string) *string {
	var this1 *string
	this1 = raw
	return this1
}

func _Main__Token_Impl__make(raw *string) *string {
	var this1 *string
	this1 = raw
	return this1
}

func _Main__Token_Impl__text(this1 *string) *string {
	return this1
}

func hx_recover_class_0(value any) *haxe__ValueException {
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

func hx_recover_class_1(value any) *haxe__Exception {
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

func hx_is_nominal_0(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_1(value any) bool {
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
