package main

import "snapshot/hxrt"

func haxe__rtti__Rtti_getRtti(c any) map[string]any {
	var rtti any = Reflect_field(c, hxrt.StringFromLiteral("__rtti"))
	if hxrt.AnyEqualsNull(rtti) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Class "), Type_getClassName(c)), hxrt.StringFromLiteral(" has no RTTI information, consider adding @:rtti")))
	}
	x := Xml_parse(hxrt.StdString(rtti)).firstElement()
	infos := New_haxe__rtti__XmlParser().processElement(x)
	if infos.tag == 1 {
		_g := infos.params[0].(map[string]any)
		c_1 := _g
		return c_1
	} else {
		t := infos
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Enum mismatch: expected TClassDecl but found "), hxrt.StdString(t)))
		var hx_throw_zero_671 map[string]any
		return hx_throw_zero_671
	}
}

func haxe__rtti__Rtti_hasRtti(c any) bool {
	_g := 0
	_g1 := Type_getClassFields(c)
	for _g < _g1.Len() {
		fieldName := func(hx_value_672 any) *string {
			if hx_value_672 == nil {
				var hx_zero_673 *string
				return hx_zero_673
			}
			return hx_value_672.(*string)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(fieldName, hxrt.StringFromLiteral("__rtti")) {
			return true
		}
	}
	return false
}
