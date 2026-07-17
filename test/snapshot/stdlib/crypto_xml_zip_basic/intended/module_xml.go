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
	var hx_switch_69 *string
	switch _g {
	case 0:
		hx_switch_69 = hxrt.StringFromLiteral("Element")
	case 1:
		hx_switch_69 = hxrt.StringFromLiteral("PCData")
	case 2:
		hx_switch_69 = hxrt.StringFromLiteral("CData")
	case 3:
		hx_switch_69 = hxrt.StringFromLiteral("Comment")
	case 4:
		hx_switch_69 = hxrt.StringFromLiteral("DocType")
	case 5:
		hx_switch_69 = hxrt.StringFromLiteral("ProcessingInstruction")
	case 6:
		hx_switch_69 = hxrt.StringFromLiteral("Document")
	}
	return hx_switch_69
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
	return func(hx_value_70 any) *string {
		if hx_value_70 == nil {
			var hx_zero_71 *string
			return hx_zero_71
		}
		return hx_value_70.(*string)
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
	func(hx_value_72 any) bool {
		if hx_value_72 == nil {
			var hx_zero_73 bool
			return hx_zero_73
		}
		return hx_value_72.(bool)
	}(this1.(*haxe__ds__StringMap).remove(att))
}

func (self *Xml) exists(att *string) bool {
	if self.nodeType != Xml_Element {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType))))
	}
	var this1 haxe__IMap = self.attributeMap
	return func(hx_value_74 any) bool {
		if hx_value_74 == nil {
			var hx_zero_75 bool
			return hx_zero_75
		}
		return hx_value_74.(bool)
	}(this1.(*haxe__ds__StringMap).exists(att))
}

func (self *Xml) attributes() map[string]any {
	if self.nodeType != Xml_Element {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType))))
	}
	var this1 haxe__IMap = self.attributeMap
	return func(hx_value_76 any) map[string]any {
		if hx_value_76 == nil {
			var hx_zero_77 map[string]any
			return hx_zero_77
		}
		return hx_value_76.(map[string]any)
	}(this1.(*haxe__ds__StringMap).keys())
}

func (self *Xml) iterator() map[string]any {
	self.ensureElementType()
	return func() map[string]any {
		hx_structural_array_78 := self.children
		hx_structural_array_index_79 := 0
		hx_structural_iterator_map_80 := map[string]any{}
		hx_structural_iterator_map_80["hasNext"] = func() bool {
			return (hx_structural_array_index_79 < len(hx_structural_array_78))
		}
		hx_structural_iterator_map_80["next"] = func() *Xml {
			hx_structural_array_value_81 := hx_structural_array_78[hx_structural_array_index_79]
			hx_structural_array_index_79 = (hx_structural_array_index_79 + 1)
			return hx_structural_array_value_81
		}
		return hx_structural_iterator_map_80
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
		hx_structural_array_83 := ret
		hx_structural_array_index_84 := 0
		hx_structural_iterator_map_85 := map[string]any{}
		hx_structural_iterator_map_85["hasNext"] = func() bool {
			return (hx_structural_array_index_84 < len(hx_structural_array_83))
		}
		hx_structural_iterator_map_85["next"] = func() *Xml {
			hx_structural_array_value_86 := hx_structural_array_83[hx_structural_array_index_84]
			hx_structural_array_index_84 = (hx_structural_array_index_84 + 1)
			return hx_structural_array_value_86
		}
		return hx_structural_iterator_map_85
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
		hx_structural_array_88 := ret
		hx_structural_array_index_89 := 0
		hx_structural_iterator_map_90 := map[string]any{}
		hx_structural_iterator_map_90["hasNext"] = func() bool {
			return (hx_structural_array_index_89 < len(hx_structural_array_88))
		}
		hx_structural_iterator_map_90["next"] = func() *Xml {
			hx_structural_array_value_91 := hx_structural_array_88[hx_structural_array_index_89]
			hx_structural_array_index_89 = (hx_structural_array_index_89 + 1)
			return hx_structural_array_value_91
		}
		return hx_structural_iterator_map_90
	}()
}

func (self *Xml) firstChild() *Xml {
	self.ensureElementType()
	var hx_if_92 *Xml
	if len(self.children) == 0 {
		hx_if_92 = nil
	} else {
		hx_if_92 = self.children[0]
	}
	return hx_if_92
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
	hx_arr_93 := self.children
	hx_arr_93 = append(hx_arr_93, x)
	self.children = hx_arr_93
	x.parent = self
}

func (self *Xml) removeChild(x *Xml) bool {
	self.ensureElementType()
	remaining := []*Xml{}
	removed := false
	_g := 0
	_g1 := self.children
	for _g < len(_g1) {
		child := _g1[_g]
		_g = int(int32((_g + 1)))
		if !removed && (child == x) {
			removed = true
		} else {
			remaining = append(remaining, child)
		}
	}
	if removed {
		self.children = remaining
		x.parent = nil
		return true
	}
	return false
}

func (self *Xml) insertChild(x *Xml, pos int) {
	self.ensureElementType()
	if x.parent != nil {
		x.parent.removeChild(x)
	}
	length := len(self.children)
	if pos < 0 {
		pos = int(int32((hxrt.Int32Wrap(length) + hxrt.Int32Wrap(pos))))
		if pos < 0 {
			pos = 0
		}
	}
	if pos > length {
		pos = length
	}
	inserted := []*Xml{}
	_g := 0
	_g1 := length
	for _g < _g1 {
		hx_post_95 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_95
		if index == pos {
			inserted = append(inserted, x)
		}
		inserted = append(inserted, self.children[index])
	}
	if pos == length {
		inserted = append(inserted, x)
	}
	self.children = inserted
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
