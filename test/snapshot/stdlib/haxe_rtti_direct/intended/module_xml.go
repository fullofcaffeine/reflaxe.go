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
	var hx_switch_714 *string
	switch _g {
	case 0:
		hx_switch_714 = hxrt.StringFromLiteral("Element")
	case 1:
		hx_switch_714 = hxrt.StringFromLiteral("PCData")
	case 2:
		hx_switch_714 = hxrt.StringFromLiteral("CData")
	case 3:
		hx_switch_714 = hxrt.StringFromLiteral("Comment")
	case 4:
		hx_switch_714 = hxrt.StringFromLiteral("DocType")
	case 5:
		hx_switch_714 = hxrt.StringFromLiteral("ProcessingInstruction")
	case 6:
		hx_switch_714 = hxrt.StringFromLiteral("Document")
	}
	return hx_switch_714
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
	children     *hxrt.Array
	attributeMap *haxe__ds__StringMap
}

func New_Xml(nodeType any) *Xml {
	self := &Xml{}
	self.__hx_this = self
	self.nodeType = nodeType
	self.children = hxrt.NewArray()
	self.attributeMap = New_haxe__ds__StringMap()
	return self
}

func (self *Xml) get_nodeName() *string {
	if !hxrt.HaxeEqual(self.nodeType, Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType))))
	}
	return self.nodeName
}

func (self *Xml) set_nodeName(v *string) *string {
	if !hxrt.HaxeEqual(self.nodeType, Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType))))
	}
	return func() *string {
		self.nodeName = v
		return self.nodeName
	}()
}

func (self *Xml) get_nodeValue() *string {
	if hxrt.HaxeEqual(self.nodeType, Xml_Document) || hxrt.HaxeEqual(self.nodeType, Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, unexpected "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType))))
	}
	return self.nodeValue
}

func (self *Xml) set_nodeValue(v *string) *string {
	if hxrt.HaxeEqual(self.nodeType, Xml_Document) || hxrt.HaxeEqual(self.nodeType, Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, unexpected "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType))))
	}
	return func() *string {
		self.nodeValue = v
		return self.nodeValue
	}()
}

func (self *Xml) get(att *string) *string {
	if !hxrt.HaxeEqual(self.nodeType, Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType))))
	}
	var this1 haxe__IMap = self.attributeMap
	return func(hx_value_715 any) *string {
		if hx_value_715 == nil {
			var hx_zero_716 *string
			return hx_zero_716
		}
		return hx_value_715.(*string)
	}(this1.(*haxe__ds__StringMap).get(att))
}

func (self *Xml) set(att *string, value *string) {
	if !hxrt.HaxeEqual(self.nodeType, Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType))))
	}
	var this1 haxe__IMap = self.attributeMap
	this1.(*haxe__ds__StringMap).set(att, value)
}

func (self *Xml) remove(att *string) {
	if !hxrt.HaxeEqual(self.nodeType, Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType))))
	}
	var this1 haxe__IMap = self.attributeMap
	func(hx_value_717 any) bool {
		if hx_value_717 == nil {
			var hx_zero_718 bool
			return hx_zero_718
		}
		return hx_value_717.(bool)
	}(this1.(*haxe__ds__StringMap).remove(att))
}

func (self *Xml) exists(att *string) bool {
	if !hxrt.HaxeEqual(self.nodeType, Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType))))
	}
	var this1 haxe__IMap = self.attributeMap
	return func(hx_value_719 any) bool {
		if hx_value_719 == nil {
			var hx_zero_720 bool
			return hx_zero_720
		}
		return hx_value_719.(bool)
	}(this1.(*haxe__ds__StringMap).exists(att))
}

func (self *Xml) attributes() map[string]any {
	if !hxrt.HaxeEqual(self.nodeType, Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType))))
	}
	var this1 haxe__IMap = self.attributeMap
	return func(hx_value_721 any) map[string]any {
		if hx_value_721 == nil {
			var hx_zero_722 map[string]any
			return hx_zero_722
		}
		return hx_value_721.(map[string]any)
	}(this1.(*haxe__ds__StringMap).keys())
}

func (self *Xml) iterator() map[string]any {
	self.ensureElementType()
	return func() map[string]any {
		hx_structural_array_723 := self.children
		hx_structural_array_index_724 := 0
		hx_structural_iterator_map_725 := map[string]any{}
		hx_structural_iterator_map_725["hasNext"] = func() bool {
			return (hx_structural_array_index_724 < hx_structural_array_723.Len())
		}
		hx_structural_iterator_map_725["next"] = func() *Xml {
			hx_structural_array_value_726 := hx_structural_array_723.Get(hx_structural_array_index_724)
			hx_structural_array_index_724 = (hx_structural_array_index_724 + 1)
			return func(hx_value_727 any) *Xml {
				if hx_value_727 == nil {
					var hx_zero_728 *Xml
					return hx_zero_728
				}
				return hx_value_727.(*Xml)
			}(any(hx_structural_array_value_726))
		}
		return hx_structural_iterator_map_725
	}()
}

func (self *Xml) elements() map[string]any {
	self.ensureElementType()
	_g := hxrt.NewArray()
	_g1 := 0
	_g2 := self.children
	for _g1 < _g2.Len() {
		child := func(hx_value_729 any) *Xml {
			if hx_value_729 == nil {
				var hx_zero_730 *Xml
				return hx_zero_730
			}
			return hx_value_729.(*Xml)
		}(_g2.Get(_g1))
		_g1 = int(int32((_g1 + 1)))
		if hxrt.HaxeEqual(child.nodeType, Xml_Element) {
			_g.Push(child)
		}
	}
	ret := _g
	return func() map[string]any {
		hx_structural_array_732 := ret
		hx_structural_array_index_733 := 0
		hx_structural_iterator_map_734 := map[string]any{}
		hx_structural_iterator_map_734["hasNext"] = func() bool {
			return (hx_structural_array_index_733 < hx_structural_array_732.Len())
		}
		hx_structural_iterator_map_734["next"] = func() *Xml {
			hx_structural_array_value_735 := hx_structural_array_732.Get(hx_structural_array_index_733)
			hx_structural_array_index_733 = (hx_structural_array_index_733 + 1)
			return func(hx_value_736 any) *Xml {
				if hx_value_736 == nil {
					var hx_zero_737 *Xml
					return hx_zero_737
				}
				return hx_value_736.(*Xml)
			}(any(hx_structural_array_value_735))
		}
		return hx_structural_iterator_map_734
	}()
}

func (self *Xml) elementsNamed(name *string) map[string]any {
	self.ensureElementType()
	_g := hxrt.NewArray()
	_g1 := 0
	_g2 := self.children
	for _g1 < _g2.Len() {
		child := func(hx_value_738 any) *Xml {
			if hx_value_738 == nil {
				var hx_zero_739 *Xml
				return hx_zero_739
			}
			return hx_value_738.(*Xml)
		}(_g2.Get(_g1))
		_g1 = int(int32((_g1 + 1)))
		if hxrt.HaxeEqual(child.nodeType, Xml_Element) && hxrt.StringEqualStringPtr(func() *string {
			if !hxrt.HaxeEqual(child.nodeType, Xml_Element) {
				hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(child.nodeType))))
			}
			return child.nodeName
		}(), name) {
			_g.Push(child)
		}
	}
	ret := _g
	return func() map[string]any {
		hx_structural_array_741 := ret
		hx_structural_array_index_742 := 0
		hx_structural_iterator_map_743 := map[string]any{}
		hx_structural_iterator_map_743["hasNext"] = func() bool {
			return (hx_structural_array_index_742 < hx_structural_array_741.Len())
		}
		hx_structural_iterator_map_743["next"] = func() *Xml {
			hx_structural_array_value_744 := hx_structural_array_741.Get(hx_structural_array_index_742)
			hx_structural_array_index_742 = (hx_structural_array_index_742 + 1)
			return func(hx_value_745 any) *Xml {
				if hx_value_745 == nil {
					var hx_zero_746 *Xml
					return hx_zero_746
				}
				return hx_value_745.(*Xml)
			}(any(hx_structural_array_value_744))
		}
		return hx_structural_iterator_map_743
	}()
}

func (self *Xml) firstChild() *Xml {
	self.ensureElementType()
	var hx_if_749 *Xml
	if self.children.Len() == 0 {
		hx_if_749 = nil
	} else {
		hx_if_749 = func(hx_value_747 any) *Xml {
			if hx_value_747 == nil {
				var hx_zero_748 *Xml
				return hx_zero_748
			}
			return hx_value_747.(*Xml)
		}(self.children.Get(0))
	}
	return hx_if_749
}

func (self *Xml) firstElement() *Xml {
	self.ensureElementType()
	_g := 0
	_g1 := self.children
	for _g < _g1.Len() {
		child := func(hx_value_750 any) *Xml {
			if hx_value_750 == nil {
				var hx_zero_751 *Xml
				return hx_zero_751
			}
			return hx_value_750.(*Xml)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		if hxrt.HaxeEqual(child.nodeType, Xml_Element) {
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
	hx_arr_752 := self.children
	hx_arr_752.Push(x)
	x.parent = self
}

func (self *Xml) removeChild(x *Xml) bool {
	self.ensureElementType()
	if func() bool {
		hx_arr_753 := self.children
		var hx_remove_value_754 any = x
		hx_remove_index_755 := 0
		for hx_remove_index_755 < hx_arr_753.Len() {
			hx_remove_element_756 := hx_arr_753.Get(hx_remove_index_755)
			if hxrt.HaxeEqual(hx_remove_element_756, hx_remove_value_754) {
				hx_arr_753.RemoveAt(hx_remove_index_755)
				return true
			}
			hx_remove_index_755 = (hx_remove_index_755 + 1)
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
			hx_arr_757 := x.parent.children
			var hx_remove_value_758 any = x
			hx_remove_index_759 := 0
			for hx_remove_index_759 < hx_arr_757.Len() {
				hx_remove_element_760 := hx_arr_757.Get(hx_remove_index_759)
				if hxrt.HaxeEqual(hx_remove_element_760, hx_remove_value_758) {
					hx_arr_757.RemoveAt(hx_remove_index_759)
					return true
				}
				hx_remove_index_759 = (hx_remove_index_759 + 1)
			}
			return false
		}()
	}
	func() {
		hx_arr_761 := self.children
		hx_insert_position_762 := pos
		var hx_insert_value_763 any = x
		hx_arr_761.Insert(hx_insert_position_762, hx_insert_value_763)
	}()
	x.parent = self
}

func (self *Xml) toString() *string {
	return haxe__xml__Printer_print(self, false)
}

func (self *Xml) ensureElementType() {
	if !hxrt.HaxeEqual(self.nodeType, Xml_Document) && !hxrt.HaxeEqual(self.nodeType, Xml_Element) {
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
	if hxrt.HaxeEqual(xml.nodeType, Xml_Document) || hxrt.HaxeEqual(xml.nodeType, Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, unexpected "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(xml.nodeType))))
	}
	xml.nodeValue = data
	return xml
}

func Xml_createComment(data *string) *Xml {
	xml := New_Xml(Xml_Comment)
	if hxrt.HaxeEqual(xml.nodeType, Xml_Document) || hxrt.HaxeEqual(xml.nodeType, Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, unexpected "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(xml.nodeType))))
	}
	xml.nodeValue = data
	return xml
}

func Xml_createDocType(data *string) *Xml {
	xml := New_Xml(Xml_DocType)
	if hxrt.HaxeEqual(xml.nodeType, Xml_Document) || hxrt.HaxeEqual(xml.nodeType, Xml_Element) {
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
	if !hxrt.HaxeEqual(xml.nodeType, Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(xml.nodeType))))
	}
	xml.nodeName = name
	return xml
}

func Xml_createPCData(data *string) *Xml {
	xml := New_Xml(Xml_PCData)
	if hxrt.HaxeEqual(xml.nodeType, Xml_Document) || hxrt.HaxeEqual(xml.nodeType, Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, unexpected "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(xml.nodeType))))
	}
	xml.nodeValue = data
	return xml
}

func Xml_createProcessingInstruction(data *string) *Xml {
	xml := New_Xml(Xml_ProcessingInstruction)
	if hxrt.HaxeEqual(xml.nodeType, Xml_Document) || hxrt.HaxeEqual(xml.nodeType, Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, unexpected "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(xml.nodeType))))
	}
	xml.nodeValue = data
	return xml
}

func Xml_parse(str *string) *Xml {
	return haxe__xml__Parser_parse(str, false)
}
