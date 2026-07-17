package main

import "snapshot/hxrt"

type I_haxe__xml__Printer interface {
	writeNode(value *Xml, tabs *string)
	write(input *string)
	newline()
	hasChildren(value *Xml) bool
}

type haxe__xml__Printer struct {
	__hx_this I_haxe__xml__Printer
	output    *StringBuf
	pretty    bool
}

func New_haxe__xml__Printer(pretty bool) *haxe__xml__Printer {
	self := &haxe__xml__Printer{}
	self.__hx_this = self
	self.output = New_StringBuf()
	self.pretty = pretty
	return self
}

func (self *haxe__xml__Printer) writeNode(value *Xml, tabs *string) {
	var _g any = value.nodeType
	switch _g {
	case 0:
		_this := self.output
		_this.b = hxrt.StringConcatStringPtr(_this.b, hxrt.StdString(hxrt.StringConcatStringPtr(tabs, hxrt.StringFromLiteral("<"))))
		if !hxrt.HaxeEqual(value.nodeType, Xml_Element) {
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(value.nodeType))))
		}
		input := value.nodeName
		_this_1 := self.output
		_this_1.b = hxrt.StringConcatStringPtr(_this_1.b, hxrt.StdString(input))
		attribute := value.attributes()
		for func(hx_obj_832 map[string]any) func() bool {
			hx_field_833 := hx_obj_832["hasNext"]
			if hx_field_833 == nil {
				var hx_zero_834 func() bool
				return hx_zero_834
			}
			return hx_field_833.(func() bool)
		}(attribute)() {
			attribute_1 := func(hx_obj_835 map[string]any) func() *string {
				hx_field_836 := hx_obj_835["next"]
				if hx_field_836 == nil {
					var hx_zero_837 func() *string
					return hx_zero_837
				}
				return hx_field_836.(func() *string)
			}(attribute)()
			_this_2 := self.output
			_this_2.b = hxrt.StringConcatStringPtr(_this_2.b, hxrt.StdString(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(" "), attribute_1), hxrt.StringFromLiteral("=\""))))
			input_1 := StringTools_htmlEscape(value.get(attribute_1), true)
			_this_3 := self.output
			_this_3.b = hxrt.StringConcatStringPtr(_this_3.b, hxrt.StdString(input_1))
			_this_4 := self.output
			_this_4.b = hxrt.StringConcatStringPtr(_this_4.b, hxrt.StringFromLiteral("\""))
		}
		if self.hasChildren(value) {
			_this_5 := self.output
			_this_5.b = hxrt.StringConcatStringPtr(_this_5.b, hxrt.StringFromLiteral(">"))
			if self.pretty {
				_this_6 := self.output
				_this_6.b = hxrt.StringConcatStringPtr(_this_6.b, hxrt.StringFromLiteral("\n"))
			}
			var _g_current int
			var _g_array *hxrt.Array
			value.ensureElementType()
			_this_7 := value.children
			_g_current = 0
			_g_array = _this_7
			for _g_current < _g_array.Len() {
				hx_post_838 := _g_current
				_g_current = int(int32((_g_current + 1)))
				child := func(hx_value_839 any) *Xml {
					if hx_value_839 == nil {
						var hx_zero_840 *Xml
						return hx_zero_840
					}
					return hx_value_839.(*Xml)
				}(_g_array.Get(hx_post_838))
				self.writeNode(child, func() *string {
					var hx_if_841 *string
					if self.pretty {
						hx_if_841 = hxrt.StringConcatStringPtr(tabs, hxrt.StringFromLiteral("\t"))
					} else {
						hx_if_841 = tabs
					}
					return hx_if_841
				}())
			}
			_this_8 := self.output
			_this_8.b = hxrt.StringConcatStringPtr(_this_8.b, hxrt.StdString(hxrt.StringConcatStringPtr(tabs, hxrt.StringFromLiteral("</"))))
			if !hxrt.HaxeEqual(value.nodeType, Xml_Element) {
				hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(value.nodeType))))
			}
			input_2 := value.nodeName
			_this_9 := self.output
			_this_9.b = hxrt.StringConcatStringPtr(_this_9.b, hxrt.StdString(input_2))
			_this_10 := self.output
			_this_10.b = hxrt.StringConcatStringPtr(_this_10.b, hxrt.StringFromLiteral(">"))
			if self.pretty {
				_this_11 := self.output
				_this_11.b = hxrt.StringConcatStringPtr(_this_11.b, hxrt.StringFromLiteral("\n"))
			}
		} else {
			_this_12 := self.output
			_this_12.b = hxrt.StringConcatStringPtr(_this_12.b, hxrt.StringFromLiteral("/>"))
			if self.pretty {
				_this_13 := self.output
				_this_13.b = hxrt.StringConcatStringPtr(_this_13.b, hxrt.StringFromLiteral("\n"))
			}
		}
	case 1:
		if hxrt.HaxeEqual(value.nodeType, Xml_Document) || hxrt.HaxeEqual(value.nodeType, Xml_Element) {
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, unexpected "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(value.nodeType))))
		}
		nodeValue := value.nodeValue
		if hxrt.StringLengthStringPtr(nodeValue) != 0 {
			input_3 := hxrt.StringConcatStringPtr(tabs, StringTools_htmlEscape(nodeValue, nil))
			_this_14 := self.output
			_this_14.b = hxrt.StringConcatStringPtr(_this_14.b, hxrt.StdString(input_3))
			if self.pretty {
				_this_15 := self.output
				_this_15.b = hxrt.StringConcatStringPtr(_this_15.b, hxrt.StringFromLiteral("\n"))
			}
		}
	case 2:
		_this_16 := self.output
		_this_16.b = hxrt.StringConcatStringPtr(_this_16.b, hxrt.StdString(hxrt.StringConcatStringPtr(tabs, hxrt.StringFromLiteral("<![CDATA["))))
		if hxrt.HaxeEqual(value.nodeType, Xml_Document) || hxrt.HaxeEqual(value.nodeType, Xml_Element) {
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, unexpected "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(value.nodeType))))
		}
		input_4 := value.nodeValue
		_this_17 := self.output
		_this_17.b = hxrt.StringConcatStringPtr(_this_17.b, hxrt.StdString(input_4))
		_this_18 := self.output
		_this_18.b = hxrt.StringConcatStringPtr(_this_18.b, hxrt.StringFromLiteral("]]>"))
		if self.pretty {
			_this_19 := self.output
			_this_19.b = hxrt.StringConcatStringPtr(_this_19.b, hxrt.StringFromLiteral("\n"))
		}
	case 3:
		if hxrt.HaxeEqual(value.nodeType, Xml_Document) || hxrt.HaxeEqual(value.nodeType, Xml_Element) {
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, unexpected "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(value.nodeType))))
		}
		commentContent := value.nodeValue
		commentContent = StringTools_replace(commentContent, hxrt.StringFromLiteral("\n"), hxrt.StringFromLiteral(""))
		commentContent = StringTools_replace(commentContent, hxrt.StringFromLiteral("\r"), hxrt.StringFromLiteral(""))
		commentContent = StringTools_replace(commentContent, hxrt.StringFromLiteral("\t"), hxrt.StringFromLiteral(""))
		commentContent = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("<!--"), commentContent), hxrt.StringFromLiteral("-->"))
		_this_20 := self.output
		_this_20.b = hxrt.StringConcatStringPtr(_this_20.b, hxrt.StdString(tabs))
		input_5 := StringTools_trim(commentContent)
		_this_21 := self.output
		_this_21.b = hxrt.StringConcatStringPtr(_this_21.b, hxrt.StdString(input_5))
		if self.pretty {
			_this_22 := self.output
			_this_22.b = hxrt.StringConcatStringPtr(_this_22.b, hxrt.StringFromLiteral("\n"))
		}
	case 4:
		input_6 := hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("<!DOCTYPE "), func() *string {
			if hxrt.HaxeEqual(value.nodeType, Xml_Document) || hxrt.HaxeEqual(value.nodeType, Xml_Element) {
				hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, unexpected "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(value.nodeType))))
			}
			return value.nodeValue
		}()), hxrt.StringFromLiteral(">"))
		_this_23 := self.output
		_this_23.b = hxrt.StringConcatStringPtr(_this_23.b, hxrt.StdString(input_6))
		if self.pretty {
			_this_24 := self.output
			_this_24.b = hxrt.StringConcatStringPtr(_this_24.b, hxrt.StringFromLiteral("\n"))
		}
	case 5:
		input_7 := hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("<?"), func() *string {
			if hxrt.HaxeEqual(value.nodeType, Xml_Document) || hxrt.HaxeEqual(value.nodeType, Xml_Element) {
				hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, unexpected "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(value.nodeType))))
			}
			return value.nodeValue
		}()), hxrt.StringFromLiteral("?>"))
		_this_25 := self.output
		_this_25.b = hxrt.StringConcatStringPtr(_this_25.b, hxrt.StdString(input_7))
		if self.pretty {
			_this_26 := self.output
			_this_26.b = hxrt.StringConcatStringPtr(_this_26.b, hxrt.StringFromLiteral("\n"))
		}
	case 6:
		var _g_current_1 int
		var _g_array_1 *hxrt.Array
		value.ensureElementType()
		_this_27 := value.children
		_g_current_1 = 0
		_g_array_1 = _this_27
		for _g_current_1 < _g_array_1.Len() {
			hx_post_842 := _g_current_1
			_g_current_1 = int(int32((_g_current_1 + 1)))
			child_1 := func(hx_value_843 any) *Xml {
				if hx_value_843 == nil {
					var hx_zero_844 *Xml
					return hx_zero_844
				}
				return hx_value_843.(*Xml)
			}(_g_array_1.Get(hx_post_842))
			self.writeNode(child_1, tabs)
		}
	}
}

func (self *haxe__xml__Printer) write(input *string) {
	_this := self.output
	_this.b = hxrt.StringConcatStringPtr(_this.b, hxrt.StdString(input))
}

func (self *haxe__xml__Printer) newline() {
	if self.pretty {
		_this := self.output
		_this.b = hxrt.StringConcatStringPtr(_this.b, hxrt.StringFromLiteral("\n"))
	}
}

func (self *haxe__xml__Printer) hasChildren(value *Xml) bool {
	var _g_current int
	var _g_array *hxrt.Array
	value.ensureElementType()
	_this := value.children
	_g_current = 0
	_g_array = _this
	for _g_current < _g_array.Len() {
		hx_post_845 := _g_current
		_g_current = int(int32((_g_current + 1)))
		child := func(hx_value_846 any) *Xml {
			if hx_value_846 == nil {
				var hx_zero_847 *Xml
				return hx_zero_847
			}
			return hx_value_846.(*Xml)
		}(_g_array.Get(hx_post_845))
		var _g any = child.nodeType
		switch _g {
		case 0, 1:
			return true
		case 2, 3:
			if hxrt.StringLengthStringPtr(StringTools_ltrim(func() *string {
				if hxrt.HaxeEqual(child.nodeType, Xml_Document) || hxrt.HaxeEqual(child.nodeType, Xml_Element) {
					hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, unexpected "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(child.nodeType))))
				}
				return child.nodeValue
			}())) != 0 {
				return true
			}
		default:
		}
	}
	return false
}

func haxe__xml__Printer_print(xml *Xml, pretty bool) *string {
	printer := New_haxe__xml__Printer(pretty)
	printer.writeNode(xml, hxrt.StringFromLiteral(""))
	return printer.output.b
}
