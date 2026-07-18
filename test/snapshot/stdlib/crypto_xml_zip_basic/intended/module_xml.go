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
	var hx_switch_80 *string
	switch _g {
	case 0:
		hx_switch_80 = hxrt.StringFromLiteral("Element")
	case 1:
		hx_switch_80 = hxrt.StringFromLiteral("PCData")
	case 2:
		hx_switch_80 = hxrt.StringFromLiteral("CData")
	case 3:
		hx_switch_80 = hxrt.StringFromLiteral("Comment")
	case 4:
		hx_switch_80 = hxrt.StringFromLiteral("DocType")
	case 5:
		hx_switch_80 = hxrt.StringFromLiteral("ProcessingInstruction")
	case 6:
		hx_switch_80 = hxrt.StringFromLiteral("Document")
	}
	return hx_switch_80
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
	return func(hx_value_81 any) *string {
		if hx_value_81 == nil {
			var hx_zero_82 *string
			return hx_zero_82
		}
		return hx_value_81.(*string)
	}(this1.(*haxe__ds__StringMap).__hx_this.get(att))
}

func (self *Xml) set(att *string, value *string) {
	if !hxrt.HaxeEqual(self.nodeType, Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType))))
	}
	var this1 haxe__IMap = self.attributeMap
	this1.(*haxe__ds__StringMap).__hx_this.set(att, value)
}

func (self *Xml) remove(att *string) {
	if !hxrt.HaxeEqual(self.nodeType, Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType))))
	}
	var this1 haxe__IMap = self.attributeMap
	func(hx_value_83 any) bool {
		if hx_value_83 == nil {
			var hx_zero_84 bool
			return hx_zero_84
		}
		return hx_value_83.(bool)
	}(this1.(*haxe__ds__StringMap).__hx_this.remove(att))
}

func (self *Xml) exists(att *string) bool {
	if !hxrt.HaxeEqual(self.nodeType, Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType))))
	}
	var this1 haxe__IMap = self.attributeMap
	return func(hx_value_85 any) bool {
		if hx_value_85 == nil {
			var hx_zero_86 bool
			return hx_zero_86
		}
		return hx_value_85.(bool)
	}(this1.(*haxe__ds__StringMap).__hx_this.exists(att))
}

func (self *Xml) attributes() map[string]any {
	if !hxrt.HaxeEqual(self.nodeType, Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType))))
	}
	var this1 haxe__IMap = self.attributeMap
	return func(hx_value_87 any) map[string]any {
		if hx_value_87 == nil {
			var hx_zero_88 map[string]any
			return hx_zero_88
		}
		return hx_value_87.(map[string]any)
	}(this1.(*haxe__ds__StringMap).__hx_this.keys())
}

func (self *Xml) iterator() map[string]any {
	self.__hx_this.ensureElementType()
	return func() map[string]any {
		hx_structural_array_89 := self.children
		hx_structural_array_index_90 := 0
		hx_structural_iterator_map_91 := map[string]any{}
		hx_structural_iterator_map_91["hasNext"] = func() bool {
			return (hx_structural_array_index_90 < hx_structural_array_89.Len())
		}
		hx_structural_iterator_map_91["next"] = func() *Xml {
			hx_structural_array_value_92 := hx_structural_array_89.Get(hx_structural_array_index_90)
			hx_structural_array_index_90 = (hx_structural_array_index_90 + 1)
			return func(hx_value_93 any) *Xml {
				if hx_value_93 == nil {
					var hx_zero_94 *Xml
					return hx_zero_94
				}
				return hx_value_93.(*Xml)
			}(any(hx_structural_array_value_92))
		}
		return hx_structural_iterator_map_91
	}()
}

func (self *Xml) elements() map[string]any {
	self.__hx_this.ensureElementType()
	_g := hxrt.NewArray()
	_g1 := 0
	_g2 := self.children
	for _g1 < _g2.Len() {
		child := func(hx_value_95 any) *Xml {
			if hx_value_95 == nil {
				var hx_zero_96 *Xml
				return hx_zero_96
			}
			return hx_value_95.(*Xml)
		}(_g2.Get(_g1))
		_g1 = int(int32((_g1 + 1)))
		if hxrt.HaxeEqual(child.nodeType, Xml_Element) {
			_g.Push(child)
		}
	}
	ret := _g
	return func() map[string]any {
		hx_structural_array_98 := ret
		hx_structural_array_index_99 := 0
		hx_structural_iterator_map_100 := map[string]any{}
		hx_structural_iterator_map_100["hasNext"] = func() bool {
			return (hx_structural_array_index_99 < hx_structural_array_98.Len())
		}
		hx_structural_iterator_map_100["next"] = func() *Xml {
			hx_structural_array_value_101 := hx_structural_array_98.Get(hx_structural_array_index_99)
			hx_structural_array_index_99 = (hx_structural_array_index_99 + 1)
			return func(hx_value_102 any) *Xml {
				if hx_value_102 == nil {
					var hx_zero_103 *Xml
					return hx_zero_103
				}
				return hx_value_102.(*Xml)
			}(any(hx_structural_array_value_101))
		}
		return hx_structural_iterator_map_100
	}()
}

func (self *Xml) elementsNamed(name *string) map[string]any {
	self.__hx_this.ensureElementType()
	_g := hxrt.NewArray()
	_g1 := 0
	_g2 := self.children
	for _g1 < _g2.Len() {
		child := func(hx_value_104 any) *Xml {
			if hx_value_104 == nil {
				var hx_zero_105 *Xml
				return hx_zero_105
			}
			return hx_value_104.(*Xml)
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
		hx_structural_array_107 := ret
		hx_structural_array_index_108 := 0
		hx_structural_iterator_map_109 := map[string]any{}
		hx_structural_iterator_map_109["hasNext"] = func() bool {
			return (hx_structural_array_index_108 < hx_structural_array_107.Len())
		}
		hx_structural_iterator_map_109["next"] = func() *Xml {
			hx_structural_array_value_110 := hx_structural_array_107.Get(hx_structural_array_index_108)
			hx_structural_array_index_108 = (hx_structural_array_index_108 + 1)
			return func(hx_value_111 any) *Xml {
				if hx_value_111 == nil {
					var hx_zero_112 *Xml
					return hx_zero_112
				}
				return hx_value_111.(*Xml)
			}(any(hx_structural_array_value_110))
		}
		return hx_structural_iterator_map_109
	}()
}

func (self *Xml) firstChild() *Xml {
	self.__hx_this.ensureElementType()
	var hx_if_115 *Xml
	if self.children.Len() == 0 {
		hx_if_115 = nil
	} else {
		hx_if_115 = func(hx_value_113 any) *Xml {
			if hx_value_113 == nil {
				var hx_zero_114 *Xml
				return hx_zero_114
			}
			return hx_value_113.(*Xml)
		}(self.children.Get(0))
	}
	return hx_if_115
}

func (self *Xml) firstElement() *Xml {
	self.__hx_this.ensureElementType()
	_g := 0
	_g1 := self.children
	for _g < _g1.Len() {
		child := func(hx_value_116 any) *Xml {
			if hx_value_116 == nil {
				var hx_zero_117 *Xml
				return hx_zero_117
			}
			return hx_value_116.(*Xml)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		if hxrt.HaxeEqual(child.nodeType, Xml_Element) {
			return child
		}
	}
	return nil
}

func (self *Xml) addChild(x *Xml) {
	self.__hx_this.ensureElementType()
	if x.parent != nil {
		x.parent.__hx_this.removeChild(x)
	}
	hx_arr_118 := self.children
	hx_arr_118.Push(x)
	x.parent = self
}

func (self *Xml) removeChild(x *Xml) bool {
	self.__hx_this.ensureElementType()
	if func() bool {
		hx_arr_119 := self.children
		var hx_remove_value_120 any = x
		hx_remove_index_121 := 0
		for hx_remove_index_121 < hx_arr_119.Len() {
			hx_remove_element_122 := hx_arr_119.Get(hx_remove_index_121)
			if hxrt.HaxeEqual(hx_remove_element_122, hx_remove_value_120) {
				hx_arr_119.RemoveAt(hx_remove_index_121)
				return true
			}
			hx_remove_index_121 = (hx_remove_index_121 + 1)
		}
		return false
	}() {
		x.parent = nil
		return true
	}
	return false
}

func (self *Xml) insertChild(x *Xml, pos int) {
	self.__hx_this.ensureElementType()
	if x.parent != nil {
		func() bool {
			hx_arr_123 := x.parent.children
			var hx_remove_value_124 any = x
			hx_remove_index_125 := 0
			for hx_remove_index_125 < hx_arr_123.Len() {
				hx_remove_element_126 := hx_arr_123.Get(hx_remove_index_125)
				if hxrt.HaxeEqual(hx_remove_element_126, hx_remove_value_124) {
					hx_arr_123.RemoveAt(hx_remove_index_125)
					return true
				}
				hx_remove_index_125 = (hx_remove_index_125 + 1)
			}
			return false
		}()
	}
	func() {
		hx_arr_127 := self.children
		hx_insert_position_128 := pos
		var hx_insert_value_129 any = x
		hx_arr_127.Insert(hx_insert_position_128, hx_insert_value_129)
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

func (self *Xml) String() string {
	return *self.__hx_this.toString()
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
