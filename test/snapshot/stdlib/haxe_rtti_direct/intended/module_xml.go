package main

import "snapshot/hxrt"

var _Xml__XmlType_Impl__CData any = any(2)

var _Xml__XmlType_Impl__Comment any = any(3)

var _Xml__XmlType_Impl__DocType any = any(4)

var _Xml__XmlType_Impl__Document any = any(6)

var _Xml__XmlType_Impl__Element any = any(0)

var _Xml__XmlType_Impl__PCData any = any(1)

var _Xml__XmlType_Impl__ProcessingInstruction any = any(5)

func _Xml__XmlType_Impl__toString(this1 int) *string {
	var _g any = any(this1)
	var hx_switch_680 *string
	switch _g {
	case 0:
		hx_switch_680 = hxrt.StringFromLiteral("Element")
	case 1:
		hx_switch_680 = hxrt.StringFromLiteral("PCData")
	case 2:
		hx_switch_680 = hxrt.StringFromLiteral("CData")
	case 3:
		hx_switch_680 = hxrt.StringFromLiteral("Comment")
	case 4:
		hx_switch_680 = hxrt.StringFromLiteral("DocType")
	case 5:
		hx_switch_680 = hxrt.StringFromLiteral("ProcessingInstruction")
	case 6:
		hx_switch_680 = hxrt.StringFromLiteral("Document")
	}
	return hx_switch_680
}

type I_Xml interface {
	get_nodeName() *string
	set_nodeName(v *string) *string
	get_nodeValue() *string
	set_nodeValue(v *string) *string
	get(att *string) *string
	set(att *string, value *string)
	remove(att *string)
	exists(att *string) bool
	attributes() map[string]any
	iterator() map[string]any
	elements() map[string]any
	elementsNamed(name *string) map[string]any
	firstChild() *Xml
	firstElement() *Xml
	addChild(x *Xml)
	removeChild(x *Xml) bool
	insertChild(x *Xml, pos int)
	toString() *string
	ensureElementType()
}

type Xml struct {
	__hx_this    I_Xml
	nodeType     any
	nodeName     *string
	nodeValue    *string
	parent       *Xml
	children     []*Xml
	attributeMap *haxe__ds__StringMap
}

func New_Xml(nodeType any) *Xml {
	self := &Xml{}
	self.__hx_this = self
	self.nodeType = nodeType
	self.children = []*Xml{}
	self.attributeMap = New_haxe__ds__StringMap()
	return self
}

func (self *Xml) get_nodeName() *string {
	if self.nodeType != Xml_Element {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType))))
	}
	return self.nodeName
}

func (self *Xml) set_nodeName(v *string) *string {
	if self.nodeType != Xml_Element {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType))))
	}
	return func() *string {
		self.nodeName = v
		return self.nodeName
	}()
}

func (self *Xml) get_nodeValue() *string {
	if (self.nodeType == Xml_Document) || (self.nodeType == Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, unexpected "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType))))
	}
	return self.nodeValue
}

func (self *Xml) set_nodeValue(v *string) *string {
	if (self.nodeType == Xml_Document) || (self.nodeType == Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, unexpected "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType))))
	}
	return func() *string {
		self.nodeValue = v
		return self.nodeValue
	}()
}

func (self *Xml) get(att *string) *string {
	if self.nodeType != Xml_Element {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType))))
	}
	var this1 haxe__IMap = self.attributeMap
	return func(hx_value_681 any) *string {
		if hx_value_681 == nil {
			var hx_zero_682 *string
			return hx_zero_682
		}
		return hx_value_681.(*string)
	}(this1.(*haxe__ds__StringMap).get(att))
}

func (self *Xml) set(att *string, value *string) {
	if self.nodeType != Xml_Element {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType))))
	}
	var this1 haxe__IMap = self.attributeMap
	this1.(*haxe__ds__StringMap).set(att, value)
}

func (self *Xml) remove(att *string) {
	if self.nodeType != Xml_Element {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType))))
	}
	var this1 haxe__IMap = self.attributeMap
	func(hx_value_683 any) bool {
		if hx_value_683 == nil {
			var hx_zero_684 bool
			return hx_zero_684
		}
		return hx_value_683.(bool)
	}(this1.(*haxe__ds__StringMap).remove(att))
}

func (self *Xml) exists(att *string) bool {
	if self.nodeType != Xml_Element {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType))))
	}
	var this1 haxe__IMap = self.attributeMap
	return func(hx_value_685 any) bool {
		if hx_value_685 == nil {
			var hx_zero_686 bool
			return hx_zero_686
		}
		return hx_value_685.(bool)
	}(this1.(*haxe__ds__StringMap).exists(att))
}

func (self *Xml) attributes() map[string]any {
	if self.nodeType != Xml_Element {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType))))
	}
	var this1 haxe__IMap = self.attributeMap
	return func(hx_value_687 any) map[string]any {
		if hx_value_687 == nil {
			var hx_zero_688 map[string]any
			return hx_zero_688
		}
		return hx_value_687.(map[string]any)
	}(this1.(*haxe__ds__StringMap).keys())
}

func (self *Xml) iterator() map[string]any {
	self.ensureElementType()
	return func() map[string]any {
		hx_structural_array_689 := self.children
		hx_structural_array_index_690 := 0
		hx_structural_iterator_map_691 := map[string]any{}
		hx_structural_iterator_map_691["hasNext"] = func() bool {
			return (hx_structural_array_index_690 < len(hx_structural_array_689))
		}
		hx_structural_iterator_map_691["next"] = func() *Xml {
			hx_structural_array_value_692 := hx_structural_array_689[hx_structural_array_index_690]
			hx_structural_array_index_690 = (hx_structural_array_index_690 + 1)
			return hx_structural_array_value_692
		}
		return hx_structural_iterator_map_691
	}()
}

func (self *Xml) elements() map[string]any {
	self.ensureElementType()
	_g := []*Xml{}
	_g1 := 0
	_g2 := self.children
	for _g1 < len(_g2) {
		child := _g2[_g1]
		_g1 = int(int32((_g1 + 1)))
		if child.nodeType == Xml_Element {
			_g = append(_g, child)
		}
	}
	ret := _g
	return func() map[string]any {
		hx_structural_array_694 := ret
		hx_structural_array_index_695 := 0
		hx_structural_iterator_map_696 := map[string]any{}
		hx_structural_iterator_map_696["hasNext"] = func() bool {
			return (hx_structural_array_index_695 < len(hx_structural_array_694))
		}
		hx_structural_iterator_map_696["next"] = func() *Xml {
			hx_structural_array_value_697 := hx_structural_array_694[hx_structural_array_index_695]
			hx_structural_array_index_695 = (hx_structural_array_index_695 + 1)
			return hx_structural_array_value_697
		}
		return hx_structural_iterator_map_696
	}()
}

func (self *Xml) elementsNamed(name *string) map[string]any {
	self.ensureElementType()
	_g := []*Xml{}
	_g1 := 0
	_g2 := self.children
	for _g1 < len(_g2) {
		child := _g2[_g1]
		_g1 = int(int32((_g1 + 1)))
		if (child.nodeType == Xml_Element) && hxrt.StringEqualStringPtr(func() *string {
			if child.nodeType != Xml_Element {
				hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(child.nodeType))))
			}
			return child.nodeName
		}(), name) {
			_g = append(_g, child)
		}
	}
	ret := _g
	return func() map[string]any {
		hx_structural_array_699 := ret
		hx_structural_array_index_700 := 0
		hx_structural_iterator_map_701 := map[string]any{}
		hx_structural_iterator_map_701["hasNext"] = func() bool {
			return (hx_structural_array_index_700 < len(hx_structural_array_699))
		}
		hx_structural_iterator_map_701["next"] = func() *Xml {
			hx_structural_array_value_702 := hx_structural_array_699[hx_structural_array_index_700]
			hx_structural_array_index_700 = (hx_structural_array_index_700 + 1)
			return hx_structural_array_value_702
		}
		return hx_structural_iterator_map_701
	}()
}

func (self *Xml) firstChild() *Xml {
	self.ensureElementType()
	var hx_if_703 *Xml
	if len(self.children) == 0 {
		hx_if_703 = nil
	} else {
		hx_if_703 = self.children[0]
	}
	return hx_if_703
}

func (self *Xml) firstElement() *Xml {
	self.ensureElementType()
	_g := 0
	_g1 := self.children
	for _g < len(_g1) {
		child := _g1[_g]
		_g = int(int32((_g + 1)))
		if child.nodeType == Xml_Element {
			return child
		}
	}
	return nil
}

func (self *Xml) addChild(x *Xml) {
	self.ensureElementType()
	if x.parent != nil {
		x.parent.removeChild(x)
	}
	hx_arr_704 := self.children
	hx_arr_704 = append(hx_arr_704, x)
	self.children = hx_arr_704
	x.parent = self
}

func (self *Xml) removeChild(x *Xml) bool {
	self.ensureElementType()
	if func() bool {
		hx_arr_705 := self.children
		var hx_remove_value_706 *Xml = x
		for hx_remove_index_707, hx_remove_element_708 := range hx_arr_705 {
			if hx_remove_element_708 == hx_remove_value_706 {
				hx_remove_last_709 := (len(hx_arr_705) - 1)
				copy(hx_arr_705[hx_remove_index_707:], hx_arr_705[(hx_remove_index_707+1):])
				var hx_remove_zero_710 *Xml
				hx_arr_705[hx_remove_last_709] = hx_remove_zero_710
				hx_arr_705 = hx_arr_705[:hx_remove_last_709]
				self.children = hx_arr_705
				return true
			}
		}
		return false
	}() {
		x.parent = nil
		return true
	}
	return false
}

func (self *Xml) insertChild(x *Xml, pos int) {
	self.ensureElementType()
	if x.parent != nil {
		func() bool {
			hx_arr_711 := x.parent.children
			var hx_remove_value_712 *Xml = x
			for hx_remove_index_713, hx_remove_element_714 := range hx_arr_711 {
				if hx_remove_element_714 == hx_remove_value_712 {
					hx_remove_last_715 := (len(hx_arr_711) - 1)
					copy(hx_arr_711[hx_remove_index_713:], hx_arr_711[(hx_remove_index_713+1):])
					var hx_remove_zero_716 *Xml
					hx_arr_711[hx_remove_last_715] = hx_remove_zero_716
					hx_arr_711 = hx_arr_711[:hx_remove_last_715]
					x.parent.children = hx_arr_711
					return true
				}
			}
			return false
		}()
	}
	func() {
		hx_arr_717 := self.children
		hx_insert_position_718 := pos
		var hx_insert_value_719 *Xml = x
		hx_insert_length_720 := len(hx_arr_717)
		if hx_insert_position_718 < 0 {
			hx_insert_position_718 = (hx_insert_length_720 + hx_insert_position_718)
			if hx_insert_position_718 < 0 {
				hx_insert_position_718 = 0
			}
		}
		if hx_insert_position_718 > hx_insert_length_720 {
			hx_insert_position_718 = hx_insert_length_720
		}
		var hx_insert_zero_721 *Xml
		hx_arr_717 = append(hx_arr_717, hx_insert_zero_721)
		copy(hx_arr_717[(hx_insert_position_718+1):], hx_arr_717[hx_insert_position_718:])
		hx_arr_717[hx_insert_position_718] = hx_insert_value_719
		self.children = hx_arr_717
	}()
	x.parent = self
}

func (self *Xml) toString() *string {
	return haxe__xml__Printer_print(self, false)
}

func (self *Xml) ensureElementType() {
	if (self.nodeType != Xml_Document) && (self.nodeType != Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element or Document but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType))))
	}
}

var Xml_CData any = any(2)

var Xml_Comment any = any(3)

var Xml_DocType any = any(4)

var Xml_Document any = any(6)

var Xml_Element any = any(0)

var Xml_PCData any = any(1)

var Xml_ProcessingInstruction any = any(5)

func Xml_createCData(data *string) *Xml {
	xml := New_Xml(Xml_CData)
	if (xml.nodeType == Xml_Document) || (xml.nodeType == Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, unexpected "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(xml.nodeType))))
	}
	xml.nodeValue = data
	return xml
}

func Xml_createComment(data *string) *Xml {
	xml := New_Xml(Xml_Comment)
	if (xml.nodeType == Xml_Document) || (xml.nodeType == Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, unexpected "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(xml.nodeType))))
	}
	xml.nodeValue = data
	return xml
}

func Xml_createDocType(data *string) *Xml {
	xml := New_Xml(Xml_DocType)
	if (xml.nodeType == Xml_Document) || (xml.nodeType == Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, unexpected "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(xml.nodeType))))
	}
	xml.nodeValue = data
	return xml
}

func Xml_createDocument() *Xml {
	return New_Xml(Xml_Document)
}

func Xml_createElement(name *string) *Xml {
	xml := New_Xml(Xml_Element)
	if xml.nodeType != Xml_Element {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(xml.nodeType))))
	}
	xml.nodeName = name
	return xml
}

func Xml_createPCData(data *string) *Xml {
	xml := New_Xml(Xml_PCData)
	if (xml.nodeType == Xml_Document) || (xml.nodeType == Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, unexpected "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(xml.nodeType))))
	}
	xml.nodeValue = data
	return xml
}

func Xml_createProcessingInstruction(data *string) *Xml {
	xml := New_Xml(Xml_ProcessingInstruction)
	if (xml.nodeType == Xml_Document) || (xml.nodeType == Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, unexpected "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(xml.nodeType))))
	}
	xml.nodeValue = data
	return xml
}

func Xml_parse(str *string) *Xml {
	return haxe__xml__Parser_parse(str, false)
}
