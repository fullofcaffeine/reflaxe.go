package main

import "snapshot/hxrt"

func haxe__rtti__Rtti_getRtti(c any) hxrt.Anonymous {
	var rtti any = Reflect_field(c, hxrt.StringFromLiteral("__rtti"))
	if hxrt.AnyEqualsNull(rtti) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Class "), Type_getClassName(c)), hxrt.StringFromLiteral(" has no RTTI information, consider adding @:rtti"))))
	}
	x := Xml_parse(hxrt.StdString(rtti)).__hx_this.firstElement()
	infos := New_haxe__rtti__XmlParser().__hx_this.processElement(x)
	if infos.tag == 1 {
		_g := hx_structural_recover(infos.params[0])
		c_1 := _g
		return c_1
	} else {
		t := infos
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Enum mismatch: expected TClassDecl but found "), hxrt.StdString(t))))
		var hx_throw_zero_1 hxrt.Anonymous
		return hx_throw_zero_1
	}
}

func haxe__rtti__Rtti_hasRtti(c any) bool {
	_g := 0
	_g1 := Type_getClassFields(c)
	for _g < _g1.Len() {
		fieldName := func(hx_value_2 any) *string {
			if hx_value_2 == nil {
				var hx_zero_3 *string
				return hx_zero_3
			}
			return hx_value_2.(*string)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(fieldName, hxrt.StringFromLiteral("__rtti")) {
			return true
		}
	}
	return false
}
