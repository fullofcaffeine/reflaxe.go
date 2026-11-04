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
	attributes() hxrt.Anonymous
	iterator() hxrt.Anonymous
	elements() hxrt.Anonymous
	elementsNamed(name *string) hxrt.Anonymous
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

func (self *Xml) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Xml(nodeType any) *Xml {
	return hx_init_Xml(hx_alloc_Xml(nil), nodeType)
}

func hx_alloc_Xml(owner any) *Xml {
	self := &Xml{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Xml)
	}
	return self
}

func hx_init_Xml(self *Xml, nodeType any) *Xml {
	self.nodeType = nodeType
	self.children = hxrt.NewArray()
	self.attributeMap = New_haxe__ds__StringMap()
	return self
}

func (self *Xml) get_nodeName() *string {
	if !hxrt.HaxeEqual(self.nodeType, Xml_Element) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType)))))
	}
	return self.nodeName
}

func (self *Xml) set_nodeName(v *string) *string {
	if !hxrt.HaxeEqual(self.nodeType, Xml_Element) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType)))))
	}
	return func() *string {
		self.nodeName = v
		return self.nodeName
	}()
}

func (self *Xml) get_nodeValue() *string {
	if hxrt.HaxeEqual(self.nodeType, Xml_Document) || hxrt.HaxeEqual(self.nodeType, Xml_Element) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, unexpected "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType)))))
	}
	return self.nodeValue
}

func (self *Xml) set_nodeValue(v *string) *string {
	if hxrt.HaxeEqual(self.nodeType, Xml_Document) || hxrt.HaxeEqual(self.nodeType, Xml_Element) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, unexpected "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType)))))
	}
	return func() *string {
		self.nodeValue = v
		return self.nodeValue
	}()
}

func (self *Xml) get(att *string) *string {
	if !hxrt.HaxeEqual(self.nodeType, Xml_Element) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType)))))
	}
	var this1 haxe__IMap = self.attributeMap
	return func(hx_value_2 any) *string {
		if hx_value_2 == nil {
			var hx_zero_3 *string
			return hx_zero_3
		}
		return hx_value_2.(*string)
	}(this1.(*haxe__ds__StringMap).__hx_this.get(att))
}

func (self *Xml) set(att *string, value *string) {
	if !hxrt.HaxeEqual(self.nodeType, Xml_Element) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType)))))
	}
	var this1 haxe__IMap = self.attributeMap
	this1.(*haxe__ds__StringMap).__hx_this.set(att, value)
}

func (self *Xml) remove(att *string) {
	if !hxrt.HaxeEqual(self.nodeType, Xml_Element) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType)))))
	}
	var this1 haxe__IMap = self.attributeMap
	func(hx_value_4 any) bool {
		if hx_value_4 == nil {
			var hx_zero_5 bool
			return hx_zero_5
		}
		return hx_value_4.(bool)
	}(this1.(*haxe__ds__StringMap).__hx_this.remove(att))
}

func (self *Xml) exists(att *string) bool {
	if !hxrt.HaxeEqual(self.nodeType, Xml_Element) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType)))))
	}
	var this1 haxe__IMap = self.attributeMap
	return func(hx_value_6 any) bool {
		if hx_value_6 == nil {
			var hx_zero_7 bool
			return hx_zero_7
		}
		return hx_value_6.(bool)
	}(this1.(*haxe__ds__StringMap).__hx_this.exists(att))
}

func (self *Xml) attributes() hxrt.Anonymous {
	if !hxrt.HaxeEqual(self.nodeType, Xml_Element) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType)))))
	}
	var this1 haxe__IMap = self.attributeMap
	return this1.(*haxe__ds__StringMap).__hx_this.keys()
}

func (self *Xml) iterator() hxrt.Anonymous {
	self.__hx_this.ensureElementType()
	_this := self.children
	return func(hx_structural_source_8 *haxe__iterators__ArrayIterator) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_8, func(name string) hxrt.AnonymousField {
			switch name {
			case "array":
				return hxrt.AnonymousField{Value: hx_structural_source_8.array, Found: true}
			case "current":
				return hxrt.AnonymousField{Value: hx_structural_source_8.current, Found: true}
			case "hasNext":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
					hx_receiver_11 := hx_structural_source_8
					return hxrt.BoundFunction(func() bool {
						return hx_receiver_11.__hx_this.hasNext()
					}, func() any {
						return hx_receiver_11.__hx_this
					}, "hasNext")
				}(), Found: true}
			case "next":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
					hx_receiver_12 := hx_structural_source_8
					return hxrt.BoundFunction(func() any {
						return hx_receiver_12.__hx_this.next()
					}, func() any {
						return hx_receiver_12.__hx_this
					}, "next")
				}(), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			case "current":
				hx_structural_source_8.current = func(hx_value_9 any) int {
					if hx_value_9 == nil {
						var hx_zero_10 int
						return hx_zero_10
					}
					return hx_value_9.(int)
				}(incoming)
				return true
			}
			return false
		})
	}(New_haxe__iterators__ArrayIterator(_this))
}

func (self *Xml) elements() hxrt.Anonymous {
	self.__hx_this.ensureElementType()
	_g := hxrt.NewArray()
	_g1 := 0
	_g2 := self.children
	for _g1 < _g2.Len() {
		child := hx_recover_class_0(_g2.Get(_g1))
		_g1 = int(int32((_g1 + 1)))
		if hxrt.HaxeEqual(child.nodeType, Xml_Element) {
			_g.Push(child)
		}
	}
	ret := _g
	return func(hx_structural_source_14 *haxe__iterators__ArrayIterator) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_14, func(name string) hxrt.AnonymousField {
			switch name {
			case "array":
				return hxrt.AnonymousField{Value: hx_structural_source_14.array, Found: true}
			case "current":
				return hxrt.AnonymousField{Value: hx_structural_source_14.current, Found: true}
			case "hasNext":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
					hx_receiver_17 := hx_structural_source_14
					return hxrt.BoundFunction(func() bool {
						return hx_receiver_17.__hx_this.hasNext()
					}, func() any {
						return hx_receiver_17.__hx_this
					}, "hasNext")
				}(), Found: true}
			case "next":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
					hx_receiver_18 := hx_structural_source_14
					return hxrt.BoundFunction(func() any {
						return hx_receiver_18.__hx_this.next()
					}, func() any {
						return hx_receiver_18.__hx_this
					}, "next")
				}(), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			case "current":
				hx_structural_source_14.current = func(hx_value_15 any) int {
					if hx_value_15 == nil {
						var hx_zero_16 int
						return hx_zero_16
					}
					return hx_value_15.(int)
				}(incoming)
				return true
			}
			return false
		})
	}(New_haxe__iterators__ArrayIterator(ret))
}

func (self *Xml) elementsNamed(name *string) hxrt.Anonymous {
	self.__hx_this.ensureElementType()
	_g := hxrt.NewArray()
	_g1 := 0
	_g2 := self.children
	for _g1 < _g2.Len() {
		child := hx_recover_class_0(_g2.Get(_g1))
		_g1 = int(int32((_g1 + 1)))
		if hxrt.HaxeEqual(child.nodeType, Xml_Element) && hxrt.StringEqualStringPtr(func() *string {
			if !hxrt.HaxeEqual(child.nodeType, Xml_Element) {
				hxrt.Throw(haxe__Exception_thrown(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(child.nodeType)))))
			}
			return child.nodeName
		}(), name) {
			_g.Push(child)
		}
	}
	ret := _g
	return func(hx_structural_source_20 *haxe__iterators__ArrayIterator) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_20, func(name string) hxrt.AnonymousField {
			switch name {
			case "array":
				return hxrt.AnonymousField{Value: hx_structural_source_20.array, Found: true}
			case "current":
				return hxrt.AnonymousField{Value: hx_structural_source_20.current, Found: true}
			case "hasNext":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
					hx_receiver_23 := hx_structural_source_20
					return hxrt.BoundFunction(func() bool {
						return hx_receiver_23.__hx_this.hasNext()
					}, func() any {
						return hx_receiver_23.__hx_this
					}, "hasNext")
				}(), Found: true}
			case "next":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
					hx_receiver_24 := hx_structural_source_20
					return hxrt.BoundFunction(func() any {
						return hx_receiver_24.__hx_this.next()
					}, func() any {
						return hx_receiver_24.__hx_this
					}, "next")
				}(), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			case "current":
				hx_structural_source_20.current = func(hx_value_21 any) int {
					if hx_value_21 == nil {
						var hx_zero_22 int
						return hx_zero_22
					}
					return hx_value_21.(int)
				}(incoming)
				return true
			}
			return false
		})
	}(New_haxe__iterators__ArrayIterator(ret))
}

func (self *Xml) firstChild() *Xml {
	self.__hx_this.ensureElementType()
	var hx_if_25 *Xml
	if self.children.Len() == 0 {
		hx_if_25 = nil
	} else {
		hx_if_25 = hx_recover_class_0(hx_recover_class_0(self.children.Get(0)))
	}
	return hx_if_25
}

func (self *Xml) firstElement() *Xml {
	self.__hx_this.ensureElementType()
	_g := 0
	_g1 := self.children
	for _g < _g1.Len() {
		child := hx_recover_class_0(_g1.Get(_g))
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
	hx_arr_26 := self.children
	hx_arr_26.Push(x)
	x.parent = self
}

func (self *Xml) removeChild(x *Xml) bool {
	self.__hx_this.ensureElementType()
	if func() bool {
		hx_arr_27 := self.children
		var hx_remove_value_28 any = x
		hx_remove_index_29 := 0
		for hx_remove_index_29 < hx_arr_27.Len() {
			hx_remove_element_30 := hx_arr_27.Get(hx_remove_index_29)
			if hxrt.HaxeEqual(hx_remove_element_30, hx_remove_value_28) {
				hx_arr_27.RemoveAt(hx_remove_index_29)
				return true
			}
			hx_remove_index_29 = (hx_remove_index_29 + 1)
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
			hx_arr_31 := x.parent.children
			var hx_remove_value_32 any = x
			hx_remove_index_33 := 0
			for hx_remove_index_33 < hx_arr_31.Len() {
				hx_remove_element_34 := hx_arr_31.Get(hx_remove_index_33)
				if hxrt.HaxeEqual(hx_remove_element_34, hx_remove_value_32) {
					hx_arr_31.RemoveAt(hx_remove_index_33)
					return true
				}
				hx_remove_index_33 = (hx_remove_index_33 + 1)
			}
			return false
		}()
	}
	func() {
		hx_arr_35 := self.children
		hx_insert_position_36 := pos
		var hx_insert_value_37 any = x
		hx_arr_35.Insert(hx_insert_position_36, hx_insert_value_37)
	}()
	x.parent = self
}

func (self *Xml) toString() *string {
	return haxe__xml__Printer_print(self, false)
}

func (self *Xml) ensureElementType() {
	if !hxrt.HaxeEqual(self.nodeType, Xml_Document) && !hxrt.HaxeEqual(self.nodeType, Xml_Element) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element or Document but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(self.nodeType)))))
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
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, unexpected "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(xml.nodeType)))))
	}
	xml.nodeValue = data
	return xml
}

func Xml_createComment(data *string) *Xml {
	xml := New_Xml(Xml_Comment)
	if hxrt.HaxeEqual(xml.nodeType, Xml_Document) || hxrt.HaxeEqual(xml.nodeType, Xml_Element) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, unexpected "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(xml.nodeType)))))
	}
	xml.nodeValue = data
	return xml
}

func Xml_createDocType(data *string) *Xml {
	xml := New_Xml(Xml_DocType)
	if hxrt.HaxeEqual(xml.nodeType, Xml_Document) || hxrt.HaxeEqual(xml.nodeType, Xml_Element) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, unexpected "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(xml.nodeType)))))
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
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(xml.nodeType)))))
	}
	xml.nodeName = name
	return xml
}

func Xml_createPCData(data *string) *Xml {
	xml := New_Xml(Xml_PCData)
	if hxrt.HaxeEqual(xml.nodeType, Xml_Document) || hxrt.HaxeEqual(xml.nodeType, Xml_Element) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, unexpected "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(xml.nodeType)))))
	}
	xml.nodeValue = data
	return xml
}

func Xml_createProcessingInstruction(data *string) *Xml {
	xml := New_Xml(Xml_ProcessingInstruction)
	if hxrt.HaxeEqual(xml.nodeType, Xml_Document) || hxrt.HaxeEqual(xml.nodeType, Xml_Element) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, unexpected "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(xml.nodeType)))))
	}
	xml.nodeValue = data
	return xml
}

func Xml_parse(str *string) *Xml {
	return haxe__xml__Parser_parse(str, false)
}
