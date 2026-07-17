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
	var hx_switch_1 *string
	switch _g {
	case 0:
		hx_switch_1 = hxrt.StringFromLiteral("Element")
	case 1:
		hx_switch_1 = hxrt.StringFromLiteral("PCData")
	case 2:
		hx_switch_1 = hxrt.StringFromLiteral("CData")
	case 3:
		hx_switch_1 = hxrt.StringFromLiteral("Comment")
	case 4:
		hx_switch_1 = hxrt.StringFromLiteral("DocType")
	case 5:
		hx_switch_1 = hxrt.StringFromLiteral("ProcessingInstruction")
	case 6:
		hx_switch_1 = hxrt.StringFromLiteral("Document")
	}
	return hx_switch_1
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
	return func(hx_value_2 any) *string {
		if hx_value_2 == nil {
			var hx_zero_3 *string
			return hx_zero_3
		}
		return hx_value_2.(*string)
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
	func(hx_value_4 any) bool {
		if hx_value_4 == nil {
			var hx_zero_5 bool
			return hx_zero_5
		}
		return hx_value_4.(bool)
	}(this1.(*haxe__ds__StringMap).remove(att))
}

func (self *Xml) exists(att *string) bool {
	if !hxrt.HaxeEqual(self.nodeType, Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType))))
	}
	var this1 haxe__IMap = self.attributeMap
	return func(hx_value_6 any) bool {
		if hx_value_6 == nil {
			var hx_zero_7 bool
			return hx_zero_7
		}
		return hx_value_6.(bool)
	}(this1.(*haxe__ds__StringMap).exists(att))
}

func (self *Xml) attributes() map[string]any {
	if !hxrt.HaxeEqual(self.nodeType, Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType))))
	}
	var this1 haxe__IMap = self.attributeMap
	return func(hx_value_8 any) map[string]any {
		if hx_value_8 == nil {
			var hx_zero_9 map[string]any
			return hx_zero_9
		}
		return hx_value_8.(map[string]any)
	}(this1.(*haxe__ds__StringMap).keys())
}

func (self *Xml) iterator() map[string]any {
	self.ensureElementType()
	return func() map[string]any {
		hx_structural_array_10 := self.children
		hx_structural_array_index_11 := 0
		hx_structural_iterator_map_12 := map[string]any{}
		hx_structural_iterator_map_12["hasNext"] = func() bool {
			return (hx_structural_array_index_11 < hx_structural_array_10.Len())
		}
		hx_structural_iterator_map_12["next"] = func() *Xml {
			hx_structural_array_value_13 := hx_structural_array_10.Get(hx_structural_array_index_11)
			hx_structural_array_index_11 = (hx_structural_array_index_11 + 1)
			return func(hx_value_14 any) *Xml {
				if hx_value_14 == nil {
					var hx_zero_15 *Xml
					return hx_zero_15
				}
				return hx_value_14.(*Xml)
			}(any(hx_structural_array_value_13))
		}
		return hx_structural_iterator_map_12
	}()
}

func (self *Xml) elements() map[string]any {
	self.ensureElementType()
	_g := hxrt.NewArray()
	_g1 := 0
	_g2 := self.children
	for _g1 < _g2.Len() {
		child := func(hx_value_16 any) *Xml {
			if hx_value_16 == nil {
				var hx_zero_17 *Xml
				return hx_zero_17
			}
			return hx_value_16.(*Xml)
		}(_g2.Get(_g1))
		_g1 = int(int32((_g1 + 1)))
		if hxrt.HaxeEqual(child.nodeType, Xml_Element) {
			_g.Push(child)
		}
	}
	ret := _g
	return func() map[string]any {
		hx_structural_array_19 := ret
		hx_structural_array_index_20 := 0
		hx_structural_iterator_map_21 := map[string]any{}
		hx_structural_iterator_map_21["hasNext"] = func() bool {
			return (hx_structural_array_index_20 < hx_structural_array_19.Len())
		}
		hx_structural_iterator_map_21["next"] = func() *Xml {
			hx_structural_array_value_22 := hx_structural_array_19.Get(hx_structural_array_index_20)
			hx_structural_array_index_20 = (hx_structural_array_index_20 + 1)
			return func(hx_value_23 any) *Xml {
				if hx_value_23 == nil {
					var hx_zero_24 *Xml
					return hx_zero_24
				}
				return hx_value_23.(*Xml)
			}(any(hx_structural_array_value_22))
		}
		return hx_structural_iterator_map_21
	}()
}

func (self *Xml) elementsNamed(name *string) map[string]any {
	self.ensureElementType()
	_g := hxrt.NewArray()
	_g1 := 0
	_g2 := self.children
	for _g1 < _g2.Len() {
		child := func(hx_value_25 any) *Xml {
			if hx_value_25 == nil {
				var hx_zero_26 *Xml
				return hx_zero_26
			}
			return hx_value_25.(*Xml)
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
		hx_structural_array_28 := ret
		hx_structural_array_index_29 := 0
		hx_structural_iterator_map_30 := map[string]any{}
		hx_structural_iterator_map_30["hasNext"] = func() bool {
			return (hx_structural_array_index_29 < hx_structural_array_28.Len())
		}
		hx_structural_iterator_map_30["next"] = func() *Xml {
			hx_structural_array_value_31 := hx_structural_array_28.Get(hx_structural_array_index_29)
			hx_structural_array_index_29 = (hx_structural_array_index_29 + 1)
			return func(hx_value_32 any) *Xml {
				if hx_value_32 == nil {
					var hx_zero_33 *Xml
					return hx_zero_33
				}
				return hx_value_32.(*Xml)
			}(any(hx_structural_array_value_31))
		}
		return hx_structural_iterator_map_30
	}()
}

func (self *Xml) firstChild() *Xml {
	self.ensureElementType()
	var hx_if_36 *Xml
	if self.children.Len() == 0 {
		hx_if_36 = nil
	} else {
		hx_if_36 = func(hx_value_34 any) *Xml {
			if hx_value_34 == nil {
				var hx_zero_35 *Xml
				return hx_zero_35
			}
			return hx_value_34.(*Xml)
		}(self.children.Get(0))
	}
	return hx_if_36
}

func (self *Xml) firstElement() *Xml {
	self.ensureElementType()
	_g := 0
	_g1 := self.children
	for _g < _g1.Len() {
		child := func(hx_value_37 any) *Xml {
			if hx_value_37 == nil {
				var hx_zero_38 *Xml
				return hx_zero_38
			}
			return hx_value_37.(*Xml)
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
	hx_arr_39 := self.children
	hx_arr_39.Push(x)
	x.parent = self
}

func (self *Xml) removeChild(x *Xml) bool {
	self.ensureElementType()
	if func() bool {
		hx_arr_40 := self.children
		var hx_remove_value_41 any = x
		hx_remove_index_42 := 0
		for hx_remove_index_42 < hx_arr_40.Len() {
			hx_remove_element_43 := hx_arr_40.Get(hx_remove_index_42)
			if hxrt.HaxeEqual(hx_remove_element_43, hx_remove_value_41) {
				hx_arr_40.RemoveAt(hx_remove_index_42)
				return true
			}
			hx_remove_index_42 = (hx_remove_index_42 + 1)
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
			hx_arr_44 := x.parent.children
			var hx_remove_value_45 any = x
			hx_remove_index_46 := 0
			for hx_remove_index_46 < hx_arr_44.Len() {
				hx_remove_element_47 := hx_arr_44.Get(hx_remove_index_46)
				if hxrt.HaxeEqual(hx_remove_element_47, hx_remove_value_45) {
					hx_arr_44.RemoveAt(hx_remove_index_46)
					return true
				}
				hx_remove_index_46 = (hx_remove_index_46 + 1)
			}
			return false
		}()
	}
	func() {
		hx_arr_48 := self.children
		hx_insert_position_49 := pos
		var hx_insert_value_50 any = x
		hx_arr_48.Insert(hx_insert_position_49, hx_insert_value_50)
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
