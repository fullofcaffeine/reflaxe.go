package main

import "snapshot/hxrt"

type I_haxe__xml__XmlParserException interface {
	toString() *string
}

type haxe__xml__XmlParserException struct {
	__hx_this      I_haxe__xml__XmlParserException
	message        *string
	lineNumber     int
	positionAtLine int
	position       int
	xml            *string
}

func New_haxe__xml__XmlParserException(message *string, xml *string, position int) *haxe__xml__XmlParserException {
	self := &haxe__xml__XmlParserException{}
	self.__hx_this = self
	self.xml = xml
	self.message = message
	self.position = position
	self.lineNumber = 1
	self.positionAtLine = 0
	_g := 0
	_g1 := position
	for _g < _g1 {
		hx_post_15 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_15
		var c_1 any = hxrt.StringCharCodeAtAnyStringPtr(xml, i)
		var hx_if_16 int
		if c_1 == nil {
			hx_if_16 = -1
		} else {
			hx_if_16 = c_1.(int)
		}
		c := hx_if_16
		if c == 10 {
			self.lineNumber = int(int32((self.lineNumber + 1)))
			self.positionAtLine = 0
		} else {
			if c != 13 {
				self.positionAtLine = int(int32((self.positionAtLine + 1)))
			}
		}
	}
	return self
}

func (self *haxe__xml__XmlParserException) toString() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(Type_getClassName(Type_getClass(self)), hxrt.StringFromLiteral(": ")), self.message), hxrt.StringFromLiteral(" at line ")), self.lineNumber), hxrt.StringFromLiteral(" char ")), self.positionAtLine)
}

func haxe__xml__Parser_doParse(str *string, strict bool, p int, parent *Xml) int {
	var xml *Xml = nil
	var state any = any(1)
	var next any = any(1)
	var aname *string = nil
	start := 0
	nsubs := 0
	nbrackets := 0
	buf := New_StringBuf()
	var escapeNext any = any(1)
	attrValQuote := -1
	for p < hxrt.StringLengthStringPtr(str) {
		var c_1 any = hxrt.StringCharCodeAtAnyStringPtr(str, p)
		var hx_if_27 int
		if c_1 == nil {
			hx_if_27 = -1
		} else {
			hx_if_27 = c_1.(int)
		}
		c := hx_if_27
		switch state {
		case 0:
			switch c {
			case 9, 10, 13, 32:
			default:
				state = next
				continue
			}
		case 1:
			if c == 60 {
				state = any(0)
				next = any(2)
			} else {
				start = p
				state = any(13)
				continue
			}
		case 2:
			switch c {
			case 33:
				if func() int {
					var c_6 any = hxrt.StringCharCodeAtAnyStringPtr(str, int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1)))))
					var hx_if_32 int
					if c_6 == nil {
						hx_if_32 = -1
					} else {
						hx_if_32 = c_6.(int)
					}
					return hx_if_32
				}() == 91 {
					p = int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(2))))
					if !hxrt.StringEqualStringPtr(hxrt.StringToUpperCaseStringPtr(hxrt.StringSubstrStringPtr(str, p, 6, true)), hxrt.StringFromLiteral("CDATA[")) {
						hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringFromLiteral("Expected <![CDATA["), str, p))
					}
					p = int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(5))))
					state = any(17)
					start = int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1))))
				} else {
					if (func() int {
						var c_4 any = hxrt.StringCharCodeAtAnyStringPtr(str, int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1)))))
						var hx_if_30 int
						if c_4 == nil {
							hx_if_30 = -1
						} else {
							hx_if_30 = c_4.(int)
						}
						return hx_if_30
					}() == 68) || (func() int {
						var c_5 any = hxrt.StringCharCodeAtAnyStringPtr(str, int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1)))))
						var hx_if_31 int
						if c_5 == nil {
							hx_if_31 = -1
						} else {
							hx_if_31 = c_5.(int)
						}
						return hx_if_31
					}() == 100) {
						if !hxrt.StringEqualStringPtr(hxrt.StringToUpperCaseStringPtr(hxrt.StringSubstrStringPtr(str, int(int32((hxrt.Int32Wrap(p)+hxrt.Int32Wrap(2)))), 6, true)), hxrt.StringFromLiteral("OCTYPE")) {
							hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringFromLiteral("Expected <!DOCTYPE"), str, p))
						}
						p = int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(8))))
						state = any(16)
						start = int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1))))
					} else {
						if (func() int {
							var c_2 any = hxrt.StringCharCodeAtAnyStringPtr(str, int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1)))))
							var hx_if_28 int
							if c_2 == nil {
								hx_if_28 = -1
							} else {
								hx_if_28 = c_2.(int)
							}
							return hx_if_28
						}() != 45) || (func() int {
							var c_3 any = hxrt.StringCharCodeAtAnyStringPtr(str, int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(2)))))
							var hx_if_29 int
							if c_3 == nil {
								hx_if_29 = -1
							} else {
								hx_if_29 = c_3.(int)
							}
							return hx_if_29
						}() != 45) {
							hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringFromLiteral("Expected <!--"), str, p))
						} else {
							p = int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(2))))
							state = any(15)
							start = int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1))))
						}
					}
				}
			case 47:
				if parent == nil {
					hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringFromLiteral("Expected node name"), str, p))
				}
				start = int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1))))
				state = any(0)
				next = any(10)
			case 63:
				state = any(14)
				start = p
			default:
				state = any(3)
				start = p
				continue
			}
		case 3:
			if !((((((((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || ((c >= 48) && (c <= 57))) || (c == 58)) || (c == 46)) || (c == 95)) || (c == 45)) {
				if p == start {
					hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringFromLiteral("Expected node name"), str, p))
				}
				xml = Xml_createElement(hxrt.StringSubstrStringPtr(str, start, int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start)))), true))
				parent.addChild(xml)
				nsubs = int(int32((nsubs + 1)))
				state = any(0)
				next = any(4)
				continue
			}
		case 4:
			switch c {
			case 47:
				state = any(11)
			case 62:
				state = any(9)
			default:
				state = any(5)
				start = p
				continue
			}
		case 5:
			if !((((((((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || ((c >= 48) && (c <= 57))) || (c == 58)) || (c == 46)) || (c == 95)) || (c == 45)) {
				var tmp *string
				if start == p {
					hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringFromLiteral("Expected attribute name"), str, p))
				}
				tmp = hxrt.StringSubstrStringPtr(str, start, int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start)))), true)
				aname = tmp
				if xml.exists(aname) {
					hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Duplicate attribute ["), aname), hxrt.StringFromLiteral("]")), str, p))
				}
				state = any(0)
				next = any(6)
				continue
			}
		case 6:
			if c == 61 {
				state = any(0)
				next = any(7)
			} else {
				hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringFromLiteral("Expected ="), str, p))
			}
		case 7:
			switch c {
			case 34, 39:
				buf = New_StringBuf()
				state = any(8)
				start = int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1))))
				attrValQuote = c
			default:
				hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringFromLiteral("Expected \""), str, p))
			}
		case 8:
			switch c {
			case 38:
				var len any = int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start))))
				var hx_if_33 *string
				if len == nil {
					hx_if_33 = hxrt.StringSubstrStringPtr(str, start, 0, false)
				} else {
					hx_if_33 = hxrt.StringSubstrStringPtr(str, start, len.(int), true)
				}
				buf.b = hxrt.StringConcatStringPtr(buf.b, hx_if_33)
				state = any(18)
				escapeNext = any(8)
				start = int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1))))
			case 60, 62:
				if strict {
					hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Invalid unescaped "), hxrt.StringFromCharCode(c)), hxrt.StringFromLiteral(" in attribute value")), str, p))
				} else {
					if c == attrValQuote {
						var len_1 any = int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start))))
						var hx_if_34 *string
						if len_1 == nil {
							hx_if_34 = hxrt.StringSubstrStringPtr(str, start, 0, false)
						} else {
							hx_if_34 = hxrt.StringSubstrStringPtr(str, start, len_1.(int), true)
						}
						buf.b = hxrt.StringConcatStringPtr(buf.b, hx_if_34)
						val := buf.b
						buf = New_StringBuf()
						xml.set(aname, val)
						state = any(0)
						next = any(4)
					}
				}
			default:
				if c == attrValQuote {
					var len_2 any = int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start))))
					var hx_if_35 *string
					if len_2 == nil {
						hx_if_35 = hxrt.StringSubstrStringPtr(str, start, 0, false)
					} else {
						hx_if_35 = hxrt.StringSubstrStringPtr(str, start, len_2.(int), true)
					}
					buf.b = hxrt.StringConcatStringPtr(buf.b, hx_if_35)
					val_1 := buf.b
					buf = New_StringBuf()
					xml.set(aname, val_1)
					state = any(0)
					next = any(4)
				}
			}
		case 9:
			p = haxe__xml__Parser_doParse(str, strict, p, xml)
			start = p
			state = any(1)
		case 10:
			if !((((((((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || ((c >= 48) && (c <= 57))) || (c == 58)) || (c == 46)) || (c == 95)) || (c == 45)) {
				if start == p {
					hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringFromLiteral("Expected node name"), str, p))
				}
				v := hxrt.StringSubstrStringPtr(str, start, int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start)))), true)
				if (parent == nil) || (parent.nodeType != any(0)) {
					hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Unexpected </"), v), hxrt.StringFromLiteral(">, tag is not open")), str, p))
				}
				if !hxrt.StringEqualStringPtr(v, func() *string {
					if parent.nodeType != Xml_Element {
						hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(parent.nodeType))))
					}
					return parent.nodeName
				}()) {
					hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Expected </"), func() *string {
						if parent.nodeType != Xml_Element {
							hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(parent.nodeType))))
						}
						return parent.nodeName
					}()), hxrt.StringFromLiteral(">")), str, p))
				}
				state = any(0)
				next = any(12)
				continue
			}
		case 11:
			if c == 62 {
				state = any(1)
			} else {
				hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringFromLiteral("Expected >"), str, p))
			}
		case 12:
			if c == 62 {
				if nsubs == 0 {
					parent.addChild(Xml_createPCData(hxrt.StringFromLiteral("")))
				}
				return p
			} else {
				hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringFromLiteral("Expected >"), str, p))
			}
		case 13:
			if c == 60 {
				var len_3 any = int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start))))
				var hx_if_36 *string
				if len_3 == nil {
					hx_if_36 = hxrt.StringSubstrStringPtr(str, start, 0, false)
				} else {
					hx_if_36 = hxrt.StringSubstrStringPtr(str, start, len_3.(int), true)
				}
				buf.b = hxrt.StringConcatStringPtr(buf.b, hx_if_36)
				child := Xml_createPCData(buf.b)
				buf = New_StringBuf()
				parent.addChild(child)
				nsubs = int(int32((nsubs + 1)))
				state = any(0)
				next = any(2)
			} else {
				if c == 38 {
					var len_4 any = int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start))))
					var hx_if_37 *string
					if len_4 == nil {
						hx_if_37 = hxrt.StringSubstrStringPtr(str, start, 0, false)
					} else {
						hx_if_37 = hxrt.StringSubstrStringPtr(str, start, len_4.(int), true)
					}
					buf.b = hxrt.StringConcatStringPtr(buf.b, hx_if_37)
					state = any(18)
					escapeNext = any(13)
					start = int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1))))
				}
			}
		case 14:
			if (c == 63) && (func() int {
				var c_7 any = hxrt.StringCharCodeAtAnyStringPtr(str, int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1)))))
				var hx_if_38 int
				if c_7 == nil {
					hx_if_38 = -1
				} else {
					hx_if_38 = c_7.(int)
				}
				return hx_if_38
			}() == 62) {
				p = int(int32((p + 1)))
				str_1 := hxrt.StringSubstrStringPtr(str, int(int32((hxrt.Int32Wrap(start) + hxrt.Int32Wrap(1)))), int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start))))) - hxrt.Int32Wrap(2)))), true)
				xml_1 := Xml_createProcessingInstruction(str_1)
				parent.addChild(xml_1)
				nsubs = int(int32((nsubs + 1)))
				state = any(1)
			}
		case 15:
			if ((c == 45) && (func() int {
				var c_8 any = hxrt.StringCharCodeAtAnyStringPtr(str, int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1)))))
				var hx_if_39 int
				if c_8 == nil {
					hx_if_39 = -1
				} else {
					hx_if_39 = c_8.(int)
				}
				return hx_if_39
			}() == 45)) && (func() int {
				var c_9 any = hxrt.StringCharCodeAtAnyStringPtr(str, int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(2)))))
				var hx_if_40 int
				if c_9 == nil {
					hx_if_40 = -1
				} else {
					hx_if_40 = c_9.(int)
				}
				return hx_if_40
			}() == 62) {
				xml_2 := Xml_createComment(hxrt.StringSubstrStringPtr(str, start, int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start)))), true))
				parent.addChild(xml_2)
				nsubs = int(int32((nsubs + 1)))
				p = int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(2))))
				state = any(1)
			}
		case 16:
			if c == 91 {
				nbrackets = int(int32((nbrackets + 1)))
			} else {
				if c == 93 {
					nbrackets = int(int32((nbrackets - 1)))
				} else {
					if (c == 62) && (nbrackets == 0) {
						xml_3 := Xml_createDocType(hxrt.StringSubstrStringPtr(str, start, int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start)))), true))
						parent.addChild(xml_3)
						nsubs = int(int32((nsubs + 1)))
						state = any(1)
					}
				}
			}
		case 17:
			if ((c == 93) && (func() int {
				var c_10 any = hxrt.StringCharCodeAtAnyStringPtr(str, int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1)))))
				var hx_if_41 int
				if c_10 == nil {
					hx_if_41 = -1
				} else {
					hx_if_41 = c_10.(int)
				}
				return hx_if_41
			}() == 93)) && (func() int {
				var c_11 any = hxrt.StringCharCodeAtAnyStringPtr(str, int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(2)))))
				var hx_if_42 int
				if c_11 == nil {
					hx_if_42 = -1
				} else {
					hx_if_42 = c_11.(int)
				}
				return hx_if_42
			}() == 62) {
				child_1 := Xml_createCData(hxrt.StringSubstrStringPtr(str, start, int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start)))), true))
				parent.addChild(child_1)
				nsubs = int(int32((nsubs + 1)))
				p = int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(2))))
				state = any(1)
			}
		case 18:
			if c == 59 {
				s := hxrt.StringSubstrStringPtr(str, start, int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start)))), true)
				if func() int {
					var c_15 any = hxrt.StringCharCodeAtAnyStringPtr(s, 0)
					var hx_if_49 int
					if c_15 == nil {
						hx_if_49 = -1
					} else {
						hx_if_49 = c_15.(int)
					}
					return hx_if_49
				}() == 35 {
					var hx_if_44 any
					if func() int {
						var c_13 any = hxrt.StringCharCodeAtAnyStringPtr(s, 1)
						var hx_if_43 int
						if c_13 == nil {
							hx_if_43 = -1
						} else {
							hx_if_43 = c_13.(int)
						}
						return hx_if_43
					}() == 120 {
						hx_if_44 = hxrt.StdParseInt(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("0"), hxrt.StringSubstrStringPtr(s, 1, int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(s))-hxrt.Int32Wrap(1)))), true)))
					} else {
						hx_if_44 = hxrt.StdParseInt(hxrt.StringSubstrStringPtr(s, 1, int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(s)) - hxrt.Int32Wrap(1)))), true))
					}
					var c_12 any = hx_if_44
					c_14 := hxrt.IntFromNullableAny(c_12)
					buf.b = hxrt.StringConcatStringPtr(buf.b, hxrt.StringFromCharCode(c_14))
				} else {
					if !func(hx_value_47 any) bool {
						if hx_value_47 == nil {
							var hx_zero_48 bool
							return hx_zero_48
						}
						return hx_value_47.(bool)
					}(haxe__xml__Parser_escapes.exists(s)) {
						if strict {
							hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Undefined entity: "), s), str, p))
						}
						buf.b = hxrt.StringConcatStringPtr(buf.b, hxrt.StdString(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("&"), s), hxrt.StringFromLiteral(";"))))
					} else {
						x := func(hx_value_45 any) *string {
							if hx_value_45 == nil {
								var hx_zero_46 *string
								return hx_zero_46
							}
							return hx_value_45.(*string)
						}(haxe__xml__Parser_escapes.get(s))
						buf.b = hxrt.StringConcatStringPtr(buf.b, hxrt.StdString(x))
					}
				}
				start = int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1))))
				state = escapeNext
			} else {
				if !((((((((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || ((c >= 48) && (c <= 57))) || (c == 58)) || (c == 46)) || (c == 95)) || (c == 45)) && (c != 35) {
					if strict {
						hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Invalid character in entity: "), hxrt.StringFromCharCode(c)), str, p))
					}
					buf.b = hxrt.StringConcatStringPtr(buf.b, hxrt.StringFromLiteral("&"))
					var len_5 any = int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start))))
					var hx_if_50 *string
					if len_5 == nil {
						hx_if_50 = hxrt.StringSubstrStringPtr(str, start, 0, false)
					} else {
						hx_if_50 = hxrt.StringSubstrStringPtr(str, start, len_5.(int), true)
					}
					buf.b = hxrt.StringConcatStringPtr(buf.b, hx_if_50)
					p = int(int32((p - 1)))
					start = int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1))))
					state = escapeNext
				}
			}
		}
		p = int(int32((p + 1)))
	}
	if state == any(1) {
		start = p
		state = any(13)
	}
	if state == any(13) {
		if parent.nodeType == any(0) {
			hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Unclosed node <"), func() *string {
				if parent.nodeType != Xml_Element {
					hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(parent.nodeType))))
				}
				return parent.nodeName
			}()), hxrt.StringFromLiteral(">")), str, p))
		}
		if (p != start) || (nsubs == 0) {
			var len_6 any = int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start))))
			var hx_if_51 *string
			if len_6 == nil {
				hx_if_51 = hxrt.StringSubstrStringPtr(str, start, 0, false)
			} else {
				hx_if_51 = hxrt.StringSubstrStringPtr(str, start, len_6.(int), true)
			}
			buf.b = hxrt.StringConcatStringPtr(buf.b, hx_if_51)
			xml_4 := Xml_createPCData(buf.b)
			parent.addChild(xml_4)
			nsubs = int(int32((nsubs + 1)))
		}
		return p
	}
	if (!strict && (state == any(18))) && (escapeNext == any(13)) {
		buf.b = hxrt.StringConcatStringPtr(buf.b, hxrt.StringFromLiteral("&"))
		var len_7 any = int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start))))
		var hx_if_52 *string
		if len_7 == nil {
			hx_if_52 = hxrt.StringSubstrStringPtr(str, start, 0, false)
		} else {
			hx_if_52 = hxrt.StringSubstrStringPtr(str, start, len_7.(int), true)
		}
		buf.b = hxrt.StringConcatStringPtr(buf.b, hx_if_52)
		xml_5 := Xml_createPCData(buf.b)
		parent.addChild(xml_5)
		nsubs = int(int32((nsubs + 1)))
		return p
	}
	hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringFromLiteral("Unexpected end"), str, p))
	var hx_throw_zero_53 int
	return hx_throw_zero_53
}

var haxe__xml__Parser_escapes *haxe__ds__StringMap = func() *haxe__ds__StringMap {
	h := New_haxe__ds__StringMap()
	h.set(hxrt.StringFromLiteral("lt"), hxrt.StringFromLiteral("<"))
	h.set(hxrt.StringFromLiteral("gt"), hxrt.StringFromLiteral(">"))
	h.set(hxrt.StringFromLiteral("amp"), hxrt.StringFromLiteral("&"))
	h.set(hxrt.StringFromLiteral("quot"), hxrt.StringFromLiteral("\""))
	h.set(hxrt.StringFromLiteral("apos"), hxrt.StringFromLiteral("'"))
	return h
}()

func haxe__xml__Parser_parse(str *string, strict bool) *Xml {
	doc := Xml_createDocument()
	haxe__xml__Parser_doParse(str, strict, 0, doc)
	return doc
}
