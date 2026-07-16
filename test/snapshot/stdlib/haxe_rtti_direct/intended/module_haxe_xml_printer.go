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
		input := value.get_nodeName()
		_this_1 := self.output
		_this_1.b = hxrt.StringConcatStringPtr(_this_1.b, hxrt.StdString(input))
		attribute := value.attributes()
		for func(hx_obj_799 map[string]any) func() bool {
			hx_field_800 := hx_obj_799["hasNext"]
			if hx_field_800 == nil {
				var hx_zero_801 func() bool
				return hx_zero_801
			}
			return hx_field_800.(func() bool)
		}(attribute)() {
			attribute_1 := func(hx_obj_802 map[string]any) func() *string {
				hx_field_803 := hx_obj_802["next"]
				if hx_field_803 == nil {
					var hx_zero_804 func() *string
					return hx_zero_804
				}
				return hx_field_803.(func() *string)
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
			child := value.iterator()
			for func(hx_obj_805 map[string]any) func() bool {
				hx_field_806 := hx_obj_805["hasNext"]
				if hx_field_806 == nil {
					var hx_zero_807 func() bool
					return hx_zero_807
				}
				return hx_field_806.(func() bool)
			}(child)() {
				child_1 := func(hx_obj_808 map[string]any) func() *Xml {
					hx_field_809 := hx_obj_808["next"]
					if hx_field_809 == nil {
						var hx_zero_810 func() *Xml
						return hx_zero_810
					}
					return hx_field_809.(func() *Xml)
				}(child)()
				self.writeNode(child_1, func() *string {
					var hx_if_811 *string
					if self.pretty {
						hx_if_811 = hxrt.StringConcatStringPtr(tabs, hxrt.StringFromLiteral("\t"))
					} else {
						hx_if_811 = tabs
					}
					return hx_if_811
				}())
			}
			_this_7 := self.output
			_this_7.b = hxrt.StringConcatStringPtr(_this_7.b, hxrt.StdString(hxrt.StringConcatStringPtr(tabs, hxrt.StringFromLiteral("</"))))
			input_2 := value.get_nodeName()
			_this_8 := self.output
			_this_8.b = hxrt.StringConcatStringPtr(_this_8.b, hxrt.StdString(input_2))
			_this_9 := self.output
			_this_9.b = hxrt.StringConcatStringPtr(_this_9.b, hxrt.StringFromLiteral(">"))
			if self.pretty {
				_this_10 := self.output
				_this_10.b = hxrt.StringConcatStringPtr(_this_10.b, hxrt.StringFromLiteral("\n"))
			}
		} else {
			_this_11 := self.output
			_this_11.b = hxrt.StringConcatStringPtr(_this_11.b, hxrt.StringFromLiteral("/>"))
			if self.pretty {
				_this_12 := self.output
				_this_12.b = hxrt.StringConcatStringPtr(_this_12.b, hxrt.StringFromLiteral("\n"))
			}
		}
	case 1:
		nodeValue := value.get_nodeValue()
		if hxrt.StringLengthStringPtr(nodeValue) != 0 {
			input_3 := hxrt.StringConcatStringPtr(tabs, StringTools_htmlEscape(nodeValue, nil))
			_this_13 := self.output
			_this_13.b = hxrt.StringConcatStringPtr(_this_13.b, hxrt.StdString(input_3))
			if self.pretty {
				_this_14 := self.output
				_this_14.b = hxrt.StringConcatStringPtr(_this_14.b, hxrt.StringFromLiteral("\n"))
			}
		}
	case 2:
		_this_15 := self.output
		_this_15.b = hxrt.StringConcatStringPtr(_this_15.b, hxrt.StdString(hxrt.StringConcatStringPtr(tabs, hxrt.StringFromLiteral("<![CDATA["))))
		input_4 := value.get_nodeValue()
		_this_16 := self.output
		_this_16.b = hxrt.StringConcatStringPtr(_this_16.b, hxrt.StdString(input_4))
		_this_17 := self.output
		_this_17.b = hxrt.StringConcatStringPtr(_this_17.b, hxrt.StringFromLiteral("]]>"))
		if self.pretty {
			_this_18 := self.output
			_this_18.b = hxrt.StringConcatStringPtr(_this_18.b, hxrt.StringFromLiteral("\n"))
		}
	case 3:
		commentContent := value.get_nodeValue()
		commentContent = StringTools_replace(commentContent, hxrt.StringFromLiteral("\n"), hxrt.StringFromLiteral(""))
		commentContent = StringTools_replace(commentContent, hxrt.StringFromLiteral("\r"), hxrt.StringFromLiteral(""))
		commentContent = StringTools_replace(commentContent, hxrt.StringFromLiteral("\t"), hxrt.StringFromLiteral(""))
		commentContent = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("<!--"), commentContent), hxrt.StringFromLiteral("-->"))
		_this_19 := self.output
		_this_19.b = hxrt.StringConcatStringPtr(_this_19.b, hxrt.StdString(tabs))
		input_5 := StringTools_trim(commentContent)
		_this_20 := self.output
		_this_20.b = hxrt.StringConcatStringPtr(_this_20.b, hxrt.StdString(input_5))
		if self.pretty {
			_this_21 := self.output
			_this_21.b = hxrt.StringConcatStringPtr(_this_21.b, hxrt.StringFromLiteral("\n"))
		}
	case 4:
		input_6 := hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("<!DOCTYPE "), value.get_nodeValue()), hxrt.StringFromLiteral(">"))
		_this_22 := self.output
		_this_22.b = hxrt.StringConcatStringPtr(_this_22.b, hxrt.StdString(input_6))
		if self.pretty {
			_this_23 := self.output
			_this_23.b = hxrt.StringConcatStringPtr(_this_23.b, hxrt.StringFromLiteral("\n"))
		}
	case 5:
		input_7 := hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("<?"), value.get_nodeValue()), hxrt.StringFromLiteral("?>"))
		_this_24 := self.output
		_this_24.b = hxrt.StringConcatStringPtr(_this_24.b, hxrt.StdString(input_7))
		if self.pretty {
			_this_25 := self.output
			_this_25.b = hxrt.StringConcatStringPtr(_this_25.b, hxrt.StringFromLiteral("\n"))
		}
	case 6:
		child_2 := value.iterator()
		for func(hx_obj_812 map[string]any) func() bool {
			hx_field_813 := hx_obj_812["hasNext"]
			if hx_field_813 == nil {
				var hx_zero_814 func() bool
				return hx_zero_814
			}
			return hx_field_813.(func() bool)
		}(child_2)() {
			child_3 := func(hx_obj_815 map[string]any) func() *Xml {
				hx_field_816 := hx_obj_815["next"]
				if hx_field_816 == nil {
					var hx_zero_817 func() *Xml
					return hx_zero_817
				}
				return hx_field_816.(func() *Xml)
			}(child_2)()
			self.writeNode(child_3, tabs)
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
	child := value.iterator()
	for func(hx_obj_818 map[string]any) func() bool {
		hx_field_819 := hx_obj_818["hasNext"]
		if hx_field_819 == nil {
			var hx_zero_820 func() bool
			return hx_zero_820
		}
		return hx_field_819.(func() bool)
	}(child)() {
		child_1 := func(hx_obj_821 map[string]any) func() *Xml {
			hx_field_822 := hx_obj_821["next"]
			if hx_field_822 == nil {
				var hx_zero_823 func() *Xml
				return hx_zero_823
			}
			return hx_field_822.(func() *Xml)
		}(child)()
		var _g any = child_1.nodeType
		switch _g {
		case 0, 1:
			return true
		case 2, 3:
			if hxrt.StringLengthStringPtr(StringTools_ltrim(child_1.get_nodeValue())) != 0 {
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
