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
		for func(hx_obj_789 map[string]any) func() bool {
			hx_field_790 := hx_obj_789["hasNext"]
			if hx_field_790 == nil {
				var hx_zero_791 func() bool
				return hx_zero_791
			}
			return hx_field_790.(func() bool)
		}(attribute)() {
			attribute_1 := func(hx_obj_792 map[string]any) func() *string {
				hx_field_793 := hx_obj_792["next"]
				if hx_field_793 == nil {
					var hx_zero_794 func() *string
					return hx_zero_794
				}
				return hx_field_793.(func() *string)
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
			for func(hx_obj_795 map[string]any) func() bool {
				hx_field_796 := hx_obj_795["hasNext"]
				if hx_field_796 == nil {
					var hx_zero_797 func() bool
					return hx_zero_797
				}
				return hx_field_796.(func() bool)
			}(child)() {
				child_1 := func(hx_obj_798 map[string]any) func() *Xml {
					hx_field_799 := hx_obj_798["next"]
					if hx_field_799 == nil {
						var hx_zero_800 func() *Xml
						return hx_zero_800
					}
					return hx_field_799.(func() *Xml)
				}(child)()
				self.writeNode(child_1, func() *string {
					var hx_if_801 *string
					if self.pretty {
						hx_if_801 = hxrt.StringConcatStringPtr(tabs, hxrt.StringFromLiteral("\t"))
					} else {
						hx_if_801 = tabs
					}
					return hx_if_801
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
		for func(hx_obj_802 map[string]any) func() bool {
			hx_field_803 := hx_obj_802["hasNext"]
			if hx_field_803 == nil {
				var hx_zero_804 func() bool
				return hx_zero_804
			}
			return hx_field_803.(func() bool)
		}(child_2)() {
			child_3 := func(hx_obj_805 map[string]any) func() *Xml {
				hx_field_806 := hx_obj_805["next"]
				if hx_field_806 == nil {
					var hx_zero_807 func() *Xml
					return hx_zero_807
				}
				return hx_field_806.(func() *Xml)
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
	for func(hx_obj_808 map[string]any) func() bool {
		hx_field_809 := hx_obj_808["hasNext"]
		if hx_field_809 == nil {
			var hx_zero_810 func() bool
			return hx_zero_810
		}
		return hx_field_809.(func() bool)
	}(child)() {
		child_1 := func(hx_obj_811 map[string]any) func() *Xml {
			hx_field_812 := hx_obj_811["next"]
			if hx_field_812 == nil {
				var hx_zero_813 func() *Xml
				return hx_zero_813
			}
			return hx_field_812.(func() *Xml)
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
