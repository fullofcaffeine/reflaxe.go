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
		hx_post_1 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_1
		var c_1 any = hxrt.StringCharCodeAtAnyStringPtr(xml, i)
		var hx_if_2 int
		if c_1 == nil {
			hx_if_2 = -1
		} else {
			hx_if_2 = c_1.(int)
		}
		c := hx_if_2
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

func (self *haxe__xml__XmlParserException) String() string {
	return *self.__hx_this.toString()
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
		var hx_if_3 int
		if c_1 == nil {
			hx_if_3 = -1
		} else {
			hx_if_3 = c_1.(int)
		}
		c := hx_if_3
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
					var c_6 any = hxrt.StringCharCodeAtAnyStringPtr(str, int((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1))))
					var hx_if_8 int
					if c_6 == nil {
						hx_if_8 = -1
					} else {
						hx_if_8 = c_6.(int)
					}
					return hx_if_8
				}() == 91 {
					p = int((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(2)))
					if !hxrt.StringEqualStringPtr(hxrt.StringToUpperCaseStringPtr(hxrt.StringSubstrStringPtr(str, p, 6, true)), hxrt.StringFromLiteral("CDATA[")) {
						hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringFromLiteral("Expected <![CDATA["), str, p))
					}
					p = int((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(5)))
					state = any(17)
					start = int((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1)))
				} else {
					if (func() int {
						var c_4 any = hxrt.StringCharCodeAtAnyStringPtr(str, int((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1))))
						var hx_if_6 int
						if c_4 == nil {
							hx_if_6 = -1
						} else {
							hx_if_6 = c_4.(int)
						}
						return hx_if_6
					}() == 68) || (func() int {
						var c_5 any = hxrt.StringCharCodeAtAnyStringPtr(str, int((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1))))
						var hx_if_7 int
						if c_5 == nil {
							hx_if_7 = -1
						} else {
							hx_if_7 = c_5.(int)
						}
						return hx_if_7
					}() == 100) {
						if !hxrt.StringEqualStringPtr(hxrt.StringToUpperCaseStringPtr(hxrt.StringSubstrStringPtr(str, int((hxrt.Int32Wrap(p)+hxrt.Int32Wrap(2))), 6, true)), hxrt.StringFromLiteral("OCTYPE")) {
							hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringFromLiteral("Expected <!DOCTYPE"), str, p))
						}
						p = int((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(8)))
						state = any(16)
						start = int((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1)))
					} else {
						if (func() int {
							var c_2 any = hxrt.StringCharCodeAtAnyStringPtr(str, int((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1))))
							var hx_if_4 int
							if c_2 == nil {
								hx_if_4 = -1
							} else {
								hx_if_4 = c_2.(int)
							}
							return hx_if_4
						}() != 45) || (func() int {
							var c_3 any = hxrt.StringCharCodeAtAnyStringPtr(str, int((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(2))))
							var hx_if_5 int
							if c_3 == nil {
								hx_if_5 = -1
							} else {
								hx_if_5 = c_3.(int)
							}
							return hx_if_5
						}() != 45) {
							hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringFromLiteral("Expected <!--"), str, p))
						} else {
							p = int((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(2)))
							state = any(15)
							start = int((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1)))
						}
					}
				}
			case 47:
				if parent == nil {
					hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringFromLiteral("Expected node name"), str, p))
				}
				start = int((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1)))
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
				xml = Xml_createElement(hxrt.StringSubstrStringPtr(str, start, int((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start))), true))
				parent.__hx_this.addChild(xml)
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
				tmp = hxrt.StringSubstrStringPtr(str, start, int((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start))), true)
				aname = tmp
				if xml.__hx_this.exists(aname) {
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
				start = int((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1)))
				attrValQuote = c
			default:
				hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringFromLiteral("Expected \""), str, p))
			}
		case 8:
			switch c {
			case 38:
				var len any = int((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start)))
				var hx_if_9 *string
				if len == nil {
					hx_if_9 = hxrt.StringSubstrStringPtr(str, start, 0, false)
				} else {
					hx_if_9 = hxrt.StringSubstrStringPtr(str, start, len.(int), true)
				}
				buf.b = hxrt.StringConcatStringPtr(buf.b, hx_if_9)
				state = any(18)
				escapeNext = any(8)
				start = int((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1)))
			case 60, 62:
				if strict {
					hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Invalid unescaped "), hxrt.StringFromCharCode(c)), hxrt.StringFromLiteral(" in attribute value")), str, p))
				} else {
					if c == attrValQuote {
						var len_1 any = int((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start)))
						var hx_if_10 *string
						if len_1 == nil {
							hx_if_10 = hxrt.StringSubstrStringPtr(str, start, 0, false)
						} else {
							hx_if_10 = hxrt.StringSubstrStringPtr(str, start, len_1.(int), true)
						}
						buf.b = hxrt.StringConcatStringPtr(buf.b, hx_if_10)
						val := buf.b
						buf = New_StringBuf()
						xml.__hx_this.set(aname, val)
						state = any(0)
						next = any(4)
					}
				}
			default:
				if c == attrValQuote {
					var len_2 any = int((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start)))
					var hx_if_11 *string
					if len_2 == nil {
						hx_if_11 = hxrt.StringSubstrStringPtr(str, start, 0, false)
					} else {
						hx_if_11 = hxrt.StringSubstrStringPtr(str, start, len_2.(int), true)
					}
					buf.b = hxrt.StringConcatStringPtr(buf.b, hx_if_11)
					val_1 := buf.b
					buf = New_StringBuf()
					xml.__hx_this.set(aname, val_1)
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
				v := hxrt.StringSubstrStringPtr(str, start, int((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start))), true)
				if (parent == nil) || !hxrt.HaxeEqual(parent.nodeType, any(0)) {
					hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Unexpected </"), v), hxrt.StringFromLiteral(">, tag is not open")), str, p))
				}
				if !hxrt.StringEqualStringPtr(v, func() *string {
					if !hxrt.HaxeEqual(parent.nodeType, Xml_Element) {
						hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(parent.nodeType))))
					}
					return parent.nodeName
				}()) {
					hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Expected </"), func() *string {
						if !hxrt.HaxeEqual(parent.nodeType, Xml_Element) {
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
					parent.__hx_this.addChild(Xml_createPCData(hxrt.StringFromLiteral("")))
				}
				return p
			} else {
				hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringFromLiteral("Expected >"), str, p))
			}
		case 13:
			if c == 60 {
				var len_3 any = int((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start)))
				var hx_if_12 *string
				if len_3 == nil {
					hx_if_12 = hxrt.StringSubstrStringPtr(str, start, 0, false)
				} else {
					hx_if_12 = hxrt.StringSubstrStringPtr(str, start, len_3.(int), true)
				}
				buf.b = hxrt.StringConcatStringPtr(buf.b, hx_if_12)
				child := Xml_createPCData(buf.b)
				buf = New_StringBuf()
				parent.__hx_this.addChild(child)
				nsubs = int(int32((nsubs + 1)))
				state = any(0)
				next = any(2)
			} else {
				if c == 38 {
					var len_4 any = int((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start)))
					var hx_if_13 *string
					if len_4 == nil {
						hx_if_13 = hxrt.StringSubstrStringPtr(str, start, 0, false)
					} else {
						hx_if_13 = hxrt.StringSubstrStringPtr(str, start, len_4.(int), true)
					}
					buf.b = hxrt.StringConcatStringPtr(buf.b, hx_if_13)
					state = any(18)
					escapeNext = any(13)
					start = int((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1)))
				}
			}
		case 14:
			if (c == 63) && (func() int {
				var c_7 any = hxrt.StringCharCodeAtAnyStringPtr(str, int((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1))))
				var hx_if_14 int
				if c_7 == nil {
					hx_if_14 = -1
				} else {
					hx_if_14 = c_7.(int)
				}
				return hx_if_14
			}() == 62) {
				p = int(int32((p + 1)))
				str_1 := hxrt.StringSubstrStringPtr(str, int((hxrt.Int32Wrap(start) + hxrt.Int32Wrap(1))), int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start)))) - hxrt.Int32Wrap(2))), true)
				xml_1 := Xml_createProcessingInstruction(str_1)
				parent.__hx_this.addChild(xml_1)
				nsubs = int(int32((nsubs + 1)))
				state = any(1)
			}
		case 15:
			if ((c == 45) && (func() int {
				var c_8 any = hxrt.StringCharCodeAtAnyStringPtr(str, int((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1))))
				var hx_if_15 int
				if c_8 == nil {
					hx_if_15 = -1
				} else {
					hx_if_15 = c_8.(int)
				}
				return hx_if_15
			}() == 45)) && (func() int {
				var c_9 any = hxrt.StringCharCodeAtAnyStringPtr(str, int((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(2))))
				var hx_if_16 int
				if c_9 == nil {
					hx_if_16 = -1
				} else {
					hx_if_16 = c_9.(int)
				}
				return hx_if_16
			}() == 62) {
				xml_2 := Xml_createComment(hxrt.StringSubstrStringPtr(str, start, int((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start))), true))
				parent.__hx_this.addChild(xml_2)
				nsubs = int(int32((nsubs + 1)))
				p = int((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(2)))
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
						xml_3 := Xml_createDocType(hxrt.StringSubstrStringPtr(str, start, int((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start))), true))
						parent.__hx_this.addChild(xml_3)
						nsubs = int(int32((nsubs + 1)))
						state = any(1)
					}
				}
			}
		case 17:
			if ((c == 93) && (func() int {
				var c_10 any = hxrt.StringCharCodeAtAnyStringPtr(str, int((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1))))
				var hx_if_17 int
				if c_10 == nil {
					hx_if_17 = -1
				} else {
					hx_if_17 = c_10.(int)
				}
				return hx_if_17
			}() == 93)) && (func() int {
				var c_11 any = hxrt.StringCharCodeAtAnyStringPtr(str, int((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(2))))
				var hx_if_18 int
				if c_11 == nil {
					hx_if_18 = -1
				} else {
					hx_if_18 = c_11.(int)
				}
				return hx_if_18
			}() == 62) {
				child_1 := Xml_createCData(hxrt.StringSubstrStringPtr(str, start, int((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start))), true))
				parent.__hx_this.addChild(child_1)
				nsubs = int(int32((nsubs + 1)))
				p = int((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(2)))
				state = any(1)
			}
		case 18:
			if c == 59 {
				s := hxrt.StringSubstrStringPtr(str, start, int((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start))), true)
				if func() int {
					var c_15 any = hxrt.StringCharCodeAtAnyStringPtr(s, 0)
					var hx_if_25 int
					if c_15 == nil {
						hx_if_25 = -1
					} else {
						hx_if_25 = c_15.(int)
					}
					return hx_if_25
				}() == 35 {
					var hx_if_20 any
					if func() int {
						var c_13 any = hxrt.StringCharCodeAtAnyStringPtr(s, 1)
						var hx_if_19 int
						if c_13 == nil {
							hx_if_19 = -1
						} else {
							hx_if_19 = c_13.(int)
						}
						return hx_if_19
					}() == 120 {
						hx_if_20 = Std_parseInt(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("0"), hxrt.StringSubstrStringPtr(s, 1, int((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(s))-hxrt.Int32Wrap(1))), true)))
					} else {
						hx_if_20 = Std_parseInt(hxrt.StringSubstrStringPtr(s, 1, int((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(s)) - hxrt.Int32Wrap(1))), true))
					}
					var c_12 any = hx_if_20
					c_14 := hxrt.IntFromNullableAny(c_12)
					buf.b = hxrt.StringConcatStringPtr(buf.b, hxrt.StringFromCharCode(c_14))
				} else {
					if !func(hx_value_23 any) bool {
						if hx_value_23 == nil {
							var hx_zero_24 bool
							return hx_zero_24
						}
						return hx_value_23.(bool)
					}(haxe__xml__Parser_escapes.__hx_this.exists(s)) {
						if strict {
							hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Undefined entity: "), s), str, p))
						}
						buf.b = hxrt.StringConcatStringPtr(buf.b, hxrt.StdString(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("&"), s), hxrt.StringFromLiteral(";"))))
					} else {
						x := func(hx_value_21 any) *string {
							if hx_value_21 == nil {
								var hx_zero_22 *string
								return hx_zero_22
							}
							return hx_value_21.(*string)
						}(haxe__xml__Parser_escapes.__hx_this.get(s))
						buf.b = hxrt.StringConcatStringPtr(buf.b, hxrt.StdString(x))
					}
				}
				start = int((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1)))
				state = escapeNext
			} else {
				if !((((((((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || ((c >= 48) && (c <= 57))) || (c == 58)) || (c == 46)) || (c == 95)) || (c == 45)) && (c != 35) {
					if strict {
						hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Invalid character in entity: "), hxrt.StringFromCharCode(c)), str, p))
					}
					buf.b = hxrt.StringConcatStringPtr(buf.b, hxrt.StringFromLiteral("&"))
					var len_5 any = int((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start)))
					var hx_if_26 *string
					if len_5 == nil {
						hx_if_26 = hxrt.StringSubstrStringPtr(str, start, 0, false)
					} else {
						hx_if_26 = hxrt.StringSubstrStringPtr(str, start, len_5.(int), true)
					}
					buf.b = hxrt.StringConcatStringPtr(buf.b, hx_if_26)
					p = int(int32((p - 1)))
					start = int((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1)))
					state = escapeNext
				}
			}
		}
		p = int(int32((p + 1)))
	}
	if hxrt.HaxeEqual(state, any(1)) {
		start = p
		state = any(13)
	}
	if hxrt.HaxeEqual(state, any(13)) {
		if hxrt.HaxeEqual(parent.nodeType, any(0)) {
			hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Unclosed node <"), func() *string {
				if !hxrt.HaxeEqual(parent.nodeType, Xml_Element) {
					hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(parent.nodeType))))
				}
				return parent.nodeName
			}()), hxrt.StringFromLiteral(">")), str, p))
		}
		if (p != start) || (nsubs == 0) {
			var len_6 any = int((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start)))
			var hx_if_27 *string
			if len_6 == nil {
				hx_if_27 = hxrt.StringSubstrStringPtr(str, start, 0, false)
			} else {
				hx_if_27 = hxrt.StringSubstrStringPtr(str, start, len_6.(int), true)
			}
			buf.b = hxrt.StringConcatStringPtr(buf.b, hx_if_27)
			xml_4 := Xml_createPCData(buf.b)
			parent.__hx_this.addChild(xml_4)
			nsubs = int(int32((nsubs + 1)))
		}
		return p
	}
	if (!strict && hxrt.HaxeEqual(state, any(18))) && hxrt.HaxeEqual(escapeNext, any(13)) {
		buf.b = hxrt.StringConcatStringPtr(buf.b, hxrt.StringFromLiteral("&"))
		var len_7 any = int((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start)))
		var hx_if_28 *string
		if len_7 == nil {
			hx_if_28 = hxrt.StringSubstrStringPtr(str, start, 0, false)
		} else {
			hx_if_28 = hxrt.StringSubstrStringPtr(str, start, len_7.(int), true)
		}
		buf.b = hxrt.StringConcatStringPtr(buf.b, hx_if_28)
		xml_5 := Xml_createPCData(buf.b)
		parent.__hx_this.addChild(xml_5)
		nsubs = int(int32((nsubs + 1)))
		return p
	}
	hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringFromLiteral("Unexpected end"), str, p))
	var hx_throw_zero_29 int
	return hx_throw_zero_29
}

var haxe__xml__Parser_escapes *haxe__ds__StringMap = func() *haxe__ds__StringMap {
	h := New_haxe__ds__StringMap()
	h.__hx_this.set(hxrt.StringFromLiteral("lt"), hxrt.StringFromLiteral("<"))
	h.__hx_this.set(hxrt.StringFromLiteral("gt"), hxrt.StringFromLiteral(">"))
	h.__hx_this.set(hxrt.StringFromLiteral("amp"), hxrt.StringFromLiteral("&"))
	h.__hx_this.set(hxrt.StringFromLiteral("quot"), hxrt.StringFromLiteral("\""))
	h.__hx_this.set(hxrt.StringFromLiteral("apos"), hxrt.StringFromLiteral("'"))
	return h
}()

func haxe__xml__Parser_parse(str *string, strict bool) *Xml {
	doc := Xml_createDocument()
	haxe__xml__Parser_doParse(str, strict, 0, doc)
	return doc
}
