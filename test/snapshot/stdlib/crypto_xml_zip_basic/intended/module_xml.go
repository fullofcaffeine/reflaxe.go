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
	var hx_switch_110 *string
	switch _g {
	case 0:
		hx_switch_110 = hxrt.StringFromLiteral("Element")
	case 1:
		hx_switch_110 = hxrt.StringFromLiteral("PCData")
	case 2:
		hx_switch_110 = hxrt.StringFromLiteral("CData")
	case 3:
		hx_switch_110 = hxrt.StringFromLiteral("Comment")
	case 4:
		hx_switch_110 = hxrt.StringFromLiteral("DocType")
	case 5:
		hx_switch_110 = hxrt.StringFromLiteral("ProcessingInstruction")
	case 6:
		hx_switch_110 = hxrt.StringFromLiteral("Document")
	}
	return hx_switch_110
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
	return func(hx_value_111 any) *string {
		if hx_value_111 == nil {
			var hx_zero_112 *string
			return hx_zero_112
		}
		return hx_value_111.(*string)
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
	func(hx_value_113 any) bool {
		if hx_value_113 == nil {
			var hx_zero_114 bool
			return hx_zero_114
		}
		return hx_value_113.(bool)
	}(this1.(*haxe__ds__StringMap).remove(att))
}

func (self *Xml) exists(att *string) bool {
	if !hxrt.HaxeEqual(self.nodeType, Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType))))
	}
	var this1 haxe__IMap = self.attributeMap
	return func(hx_value_115 any) bool {
		if hx_value_115 == nil {
			var hx_zero_116 bool
			return hx_zero_116
		}
		return hx_value_115.(bool)
	}(this1.(*haxe__ds__StringMap).exists(att))
}

func (self *Xml) attributes() map[string]any {
	if !hxrt.HaxeEqual(self.nodeType, Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType))))
	}
	var this1 haxe__IMap = self.attributeMap
	return func(hx_value_117 any) map[string]any {
		if hx_value_117 == nil {
			var hx_zero_118 map[string]any
			return hx_zero_118
		}
		return hx_value_117.(map[string]any)
	}(this1.(*haxe__ds__StringMap).keys())
}

func (self *Xml) iterator() map[string]any {
	self.ensureElementType()
	return func() map[string]any {
		hx_structural_array_119 := self.children
		hx_structural_array_index_120 := 0
		hx_structural_iterator_map_121 := map[string]any{}
		hx_structural_iterator_map_121["hasNext"] = func() bool {
			return (hx_structural_array_index_120 < hx_structural_array_119.Len())
		}
		hx_structural_iterator_map_121["next"] = func() *Xml {
			hx_structural_array_value_122 := hx_structural_array_119.Get(hx_structural_array_index_120)
			hx_structural_array_index_120 = (hx_structural_array_index_120 + 1)
			return func(hx_value_123 any) *Xml {
				if hx_value_123 == nil {
					var hx_zero_124 *Xml
					return hx_zero_124
				}
				return hx_value_123.(*Xml)
			}(any(hx_structural_array_value_122))
		}
		return hx_structural_iterator_map_121
	}()
}

func (self *Xml) elements() map[string]any {
	self.ensureElementType()
	_g := hxrt.NewArray()
	_g1 := 0
	_g2 := self.children
	for _g1 < _g2.Len() {
		child := func(hx_value_125 any) *Xml {
			if hx_value_125 == nil {
				var hx_zero_126 *Xml
				return hx_zero_126
			}
			return hx_value_125.(*Xml)
		}(_g2.Get(_g1))
		_g1 = int(int32((_g1 + 1)))
		if hxrt.HaxeEqual(child.nodeType, Xml_Element) {
			_g.Push(child)
		}
	}
	ret := _g
	return func() map[string]any {
		hx_structural_array_128 := ret
		hx_structural_array_index_129 := 0
		hx_structural_iterator_map_130 := map[string]any{}
		hx_structural_iterator_map_130["hasNext"] = func() bool {
			return (hx_structural_array_index_129 < hx_structural_array_128.Len())
		}
		hx_structural_iterator_map_130["next"] = func() *Xml {
			hx_structural_array_value_131 := hx_structural_array_128.Get(hx_structural_array_index_129)
			hx_structural_array_index_129 = (hx_structural_array_index_129 + 1)
			return func(hx_value_132 any) *Xml {
				if hx_value_132 == nil {
					var hx_zero_133 *Xml
					return hx_zero_133
				}
				return hx_value_132.(*Xml)
			}(any(hx_structural_array_value_131))
		}
		return hx_structural_iterator_map_130
	}()
}

func (self *Xml) elementsNamed(name *string) map[string]any {
	self.ensureElementType()
	_g := hxrt.NewArray()
	_g1 := 0
	_g2 := self.children
	for _g1 < _g2.Len() {
		child := func(hx_value_134 any) *Xml {
			if hx_value_134 == nil {
				var hx_zero_135 *Xml
				return hx_zero_135
			}
			return hx_value_134.(*Xml)
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
		hx_structural_array_137 := ret
		hx_structural_array_index_138 := 0
		hx_structural_iterator_map_139 := map[string]any{}
		hx_structural_iterator_map_139["hasNext"] = func() bool {
			return (hx_structural_array_index_138 < hx_structural_array_137.Len())
		}
		hx_structural_iterator_map_139["next"] = func() *Xml {
			hx_structural_array_value_140 := hx_structural_array_137.Get(hx_structural_array_index_138)
			hx_structural_array_index_138 = (hx_structural_array_index_138 + 1)
			return func(hx_value_141 any) *Xml {
				if hx_value_141 == nil {
					var hx_zero_142 *Xml
					return hx_zero_142
				}
				return hx_value_141.(*Xml)
			}(any(hx_structural_array_value_140))
		}
		return hx_structural_iterator_map_139
	}()
}

func (self *Xml) firstChild() *Xml {
	self.ensureElementType()
	var hx_if_145 *Xml
	if self.children.Len() == 0 {
		hx_if_145 = nil
	} else {
		hx_if_145 = func(hx_value_143 any) *Xml {
			if hx_value_143 == nil {
				var hx_zero_144 *Xml
				return hx_zero_144
			}
			return hx_value_143.(*Xml)
		}(self.children.Get(0))
	}
	return hx_if_145
}

func (self *Xml) firstElement() *Xml {
	self.ensureElementType()
	_g := 0
	_g1 := self.children
	for _g < _g1.Len() {
		child := func(hx_value_146 any) *Xml {
			if hx_value_146 == nil {
				var hx_zero_147 *Xml
				return hx_zero_147
			}
			return hx_value_146.(*Xml)
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
	hx_arr_148 := self.children
	hx_arr_148.Push(x)
	x.parent = self
}

func (self *Xml) removeChild(x *Xml) bool {
	self.ensureElementType()
	if func() bool {
		hx_arr_149 := self.children
		var hx_remove_value_150 any = x
		hx_remove_index_151 := 0
		for hx_remove_index_151 < hx_arr_149.Len() {
			hx_remove_element_152 := hx_arr_149.Get(hx_remove_index_151)
			if hxrt.HaxeEqual(hx_remove_element_152, hx_remove_value_150) {
				hx_arr_149.RemoveAt(hx_remove_index_151)
				return true
			}
			hx_remove_index_151 = (hx_remove_index_151 + 1)
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
			hx_arr_153 := x.parent.children
			var hx_remove_value_154 any = x
			hx_remove_index_155 := 0
			for hx_remove_index_155 < hx_arr_153.Len() {
				hx_remove_element_156 := hx_arr_153.Get(hx_remove_index_155)
				if hxrt.HaxeEqual(hx_remove_element_156, hx_remove_value_154) {
					hx_arr_153.RemoveAt(hx_remove_index_155)
					return true
				}
				hx_remove_index_155 = (hx_remove_index_155 + 1)
			}
			return false
		}()
	}
	func() {
		hx_arr_157 := self.children
		hx_insert_position_158 := pos
		var hx_insert_value_159 any = x
		hx_arr_157.Insert(hx_insert_position_158, hx_insert_value_159)
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
