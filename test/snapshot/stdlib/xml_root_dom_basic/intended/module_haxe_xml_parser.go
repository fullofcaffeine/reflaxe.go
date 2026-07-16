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
		hx_post_38 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_38
		var c_1 any = hxrt.StringCharCodeAtAnyStringPtr(xml, i)
		var hx_if_39 int
		if c_1 == nil {
			hx_if_39 = -1
		} else {
			hx_if_39 = c_1.(int)
		}
		c := hx_if_39
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
		var hx_if_50 int
		if c_1 == nil {
			hx_if_50 = -1
		} else {
			hx_if_50 = c_1.(int)
		}
		c := hx_if_50
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
					var hx_if_58 int
					if c_6 == nil {
						hx_if_58 = -1
					} else {
						hx_if_58 = c_6.(int)
					}
					return hx_if_58
				}() == 91 {
					p = int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(2))))
					if !hxrt.StringEqualStringPtr(hxrt.StringToUpperCaseStringPtr(hxrt.StringSubstrStringPtr(str, p, 6, true)), hxrt.StringFromLiteral("CDATA[")) {
						hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringFromLiteral("Expected <![CDATA["), str, p))
						var hx_throw_zero_51 int
						return hx_throw_zero_51
					}
					p = int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(5))))
					state = any(17)
					start = int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1))))
				} else {
					if (func() int {
						var c_4 any = hxrt.StringCharCodeAtAnyStringPtr(str, int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1)))))
						var hx_if_56 int
						if c_4 == nil {
							hx_if_56 = -1
						} else {
							hx_if_56 = c_4.(int)
						}
						return hx_if_56
					}() == 68) || (func() int {
						var c_5 any = hxrt.StringCharCodeAtAnyStringPtr(str, int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1)))))
						var hx_if_57 int
						if c_5 == nil {
							hx_if_57 = -1
						} else {
							hx_if_57 = c_5.(int)
						}
						return hx_if_57
					}() == 100) {
						if !hxrt.StringEqualStringPtr(hxrt.StringToUpperCaseStringPtr(hxrt.StringSubstrStringPtr(str, int(int32((hxrt.Int32Wrap(p)+hxrt.Int32Wrap(2)))), 6, true)), hxrt.StringFromLiteral("OCTYPE")) {
							hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringFromLiteral("Expected <!DOCTYPE"), str, p))
							var hx_throw_zero_52 int
							return hx_throw_zero_52
						}
						p = int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(8))))
						state = any(16)
						start = int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1))))
					} else {
						if (func() int {
							var c_2 any = hxrt.StringCharCodeAtAnyStringPtr(str, int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1)))))
							var hx_if_54 int
							if c_2 == nil {
								hx_if_54 = -1
							} else {
								hx_if_54 = c_2.(int)
							}
							return hx_if_54
						}() != 45) || (func() int {
							var c_3 any = hxrt.StringCharCodeAtAnyStringPtr(str, int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(2)))))
							var hx_if_55 int
							if c_3 == nil {
								hx_if_55 = -1
							} else {
								hx_if_55 = c_3.(int)
							}
							return hx_if_55
						}() != 45) {
							hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringFromLiteral("Expected <!--"), str, p))
							var hx_throw_zero_53 int
							return hx_throw_zero_53
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
					var hx_throw_zero_59 int
					return hx_throw_zero_59
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
					var hx_throw_zero_60 int
					return hx_throw_zero_60
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
					var hx_throw_zero_61 int
					return hx_throw_zero_61
				}
				tmp = hxrt.StringSubstrStringPtr(str, start, int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start)))), true)
				aname = tmp
				if xml.exists(aname) {
					hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Duplicate attribute ["), aname), hxrt.StringFromLiteral("]")), str, p))
					var hx_throw_zero_62 int
					return hx_throw_zero_62
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
				var hx_throw_zero_63 int
				return hx_throw_zero_63
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
				var hx_throw_zero_64 int
				return hx_throw_zero_64
			}
		case 8:
			switch c {
			case 38:
				var len any = int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start))))
				var hx_if_65 *string
				if len == nil {
					hx_if_65 = hxrt.StringSubstrStringPtr(str, start, 0, false)
				} else {
					hx_if_65 = hxrt.StringSubstrStringPtr(str, start, len.(int), true)
				}
				buf.b = hxrt.StringConcatStringPtr(buf.b, hx_if_65)
				state = any(18)
				escapeNext = any(8)
				start = int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1))))
			case 60, 62:
				if strict {
					hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Invalid unescaped "), hxrt.StringFromCharCode(c)), hxrt.StringFromLiteral(" in attribute value")), str, p))
					var hx_throw_zero_66 int
					return hx_throw_zero_66
				} else {
					if c == attrValQuote {
						var len_1 any = int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start))))
						var hx_if_67 *string
						if len_1 == nil {
							hx_if_67 = hxrt.StringSubstrStringPtr(str, start, 0, false)
						} else {
							hx_if_67 = hxrt.StringSubstrStringPtr(str, start, len_1.(int), true)
						}
						buf.b = hxrt.StringConcatStringPtr(buf.b, hx_if_67)
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
					var hx_if_68 *string
					if len_2 == nil {
						hx_if_68 = hxrt.StringSubstrStringPtr(str, start, 0, false)
					} else {
						hx_if_68 = hxrt.StringSubstrStringPtr(str, start, len_2.(int), true)
					}
					buf.b = hxrt.StringConcatStringPtr(buf.b, hx_if_68)
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
					var hx_throw_zero_69 int
					return hx_throw_zero_69
				}
				v := hxrt.StringSubstrStringPtr(str, start, int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start)))), true)
				if (parent == nil) || (parent.nodeType != any(0)) {
					hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Unexpected </"), v), hxrt.StringFromLiteral(">, tag is not open")), str, p))
					var hx_throw_zero_70 int
					return hx_throw_zero_70
				}
				if !hxrt.StringEqualStringPtr(v, parent.get_nodeName()) {
					hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Expected </"), parent.get_nodeName()), hxrt.StringFromLiteral(">")), str, p))
					var hx_throw_zero_71 int
					return hx_throw_zero_71
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
				var hx_throw_zero_72 int
				return hx_throw_zero_72
			}
		case 12:
			if c == 62 {
				if nsubs == 0 {
					parent.addChild(Xml_createPCData(hxrt.StringFromLiteral("")))
				}
				return p
			} else {
				hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringFromLiteral("Expected >"), str, p))
				var hx_throw_zero_73 int
				return hx_throw_zero_73
			}
		case 13:
			if c == 60 {
				var len_3 any = int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start))))
				var hx_if_74 *string
				if len_3 == nil {
					hx_if_74 = hxrt.StringSubstrStringPtr(str, start, 0, false)
				} else {
					hx_if_74 = hxrt.StringSubstrStringPtr(str, start, len_3.(int), true)
				}
				buf.b = hxrt.StringConcatStringPtr(buf.b, hx_if_74)
				child := Xml_createPCData(buf.b)
				buf = New_StringBuf()
				parent.addChild(child)
				nsubs = int(int32((nsubs + 1)))
				state = any(0)
				next = any(2)
			} else {
				if c == 38 {
					var len_4 any = int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start))))
					var hx_if_75 *string
					if len_4 == nil {
						hx_if_75 = hxrt.StringSubstrStringPtr(str, start, 0, false)
					} else {
						hx_if_75 = hxrt.StringSubstrStringPtr(str, start, len_4.(int), true)
					}
					buf.b = hxrt.StringConcatStringPtr(buf.b, hx_if_75)
					state = any(18)
					escapeNext = any(13)
					start = int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1))))
				}
			}
		case 14:
			if (c == 63) && (func() int {
				var c_7 any = hxrt.StringCharCodeAtAnyStringPtr(str, int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1)))))
				var hx_if_76 int
				if c_7 == nil {
					hx_if_76 = -1
				} else {
					hx_if_76 = c_7.(int)
				}
				return hx_if_76
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
				var hx_if_77 int
				if c_8 == nil {
					hx_if_77 = -1
				} else {
					hx_if_77 = c_8.(int)
				}
				return hx_if_77
			}() == 45)) && (func() int {
				var c_9 any = hxrt.StringCharCodeAtAnyStringPtr(str, int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(2)))))
				var hx_if_78 int
				if c_9 == nil {
					hx_if_78 = -1
				} else {
					hx_if_78 = c_9.(int)
				}
				return hx_if_78
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
				var hx_if_79 int
				if c_10 == nil {
					hx_if_79 = -1
				} else {
					hx_if_79 = c_10.(int)
				}
				return hx_if_79
			}() == 93)) && (func() int {
				var c_11 any = hxrt.StringCharCodeAtAnyStringPtr(str, int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(2)))))
				var hx_if_80 int
				if c_11 == nil {
					hx_if_80 = -1
				} else {
					hx_if_80 = c_11.(int)
				}
				return hx_if_80
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
					var hx_if_88 int
					if c_15 == nil {
						hx_if_88 = -1
					} else {
						hx_if_88 = c_15.(int)
					}
					return hx_if_88
				}() == 35 {
					var hx_if_82 any
					if func() int {
						var c_13 any = hxrt.StringCharCodeAtAnyStringPtr(s, 1)
						var hx_if_81 int
						if c_13 == nil {
							hx_if_81 = -1
						} else {
							hx_if_81 = c_13.(int)
						}
						return hx_if_81
					}() == 120 {
						hx_if_82 = hxrt.StdParseInt(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("0"), hxrt.StringSubstrStringPtr(s, 1, int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(s))-hxrt.Int32Wrap(1)))), true)))
					} else {
						hx_if_82 = hxrt.StdParseInt(hxrt.StringSubstrStringPtr(s, 1, int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(s)) - hxrt.Int32Wrap(1)))), true))
					}
					var c_12 any = hx_if_82
					c_14 := hxrt.IntFromNullableAny(c_12)
					buf.b = hxrt.StringConcatStringPtr(buf.b, hxrt.StringFromCharCode(c_14))
				} else {
					if !func(hx_value_86 any) bool {
						if hx_value_86 == nil {
							var hx_zero_87 bool
							return hx_zero_87
						}
						return hx_value_86.(bool)
					}(haxe__xml__Parser_escapes.exists(s)) {
						if strict {
							hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Undefined entity: "), s), str, p))
							var hx_throw_zero_83 int
							return hx_throw_zero_83
						}
						buf.b = hxrt.StringConcatStringPtr(buf.b, hxrt.StdString(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("&"), s), hxrt.StringFromLiteral(";"))))
					} else {
						x := func(hx_value_84 any) *string {
							if hx_value_84 == nil {
								var hx_zero_85 *string
								return hx_zero_85
							}
							return hx_value_84.(*string)
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
						var hx_throw_zero_89 int
						return hx_throw_zero_89
					}
					buf.b = hxrt.StringConcatStringPtr(buf.b, hxrt.StringFromLiteral("&"))
					var len_5 any = int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start))))
					var hx_if_90 *string
					if len_5 == nil {
						hx_if_90 = hxrt.StringSubstrStringPtr(str, start, 0, false)
					} else {
						hx_if_90 = hxrt.StringSubstrStringPtr(str, start, len_5.(int), true)
					}
					buf.b = hxrt.StringConcatStringPtr(buf.b, hx_if_90)
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
			hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Unclosed node <"), parent.get_nodeName()), hxrt.StringFromLiteral(">")), str, p))
			var hx_throw_zero_91 int
			return hx_throw_zero_91
		}
		if (p != start) || (nsubs == 0) {
			var len_6 any = int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start))))
			var hx_if_92 *string
			if len_6 == nil {
				hx_if_92 = hxrt.StringSubstrStringPtr(str, start, 0, false)
			} else {
				hx_if_92 = hxrt.StringSubstrStringPtr(str, start, len_6.(int), true)
			}
			buf.b = hxrt.StringConcatStringPtr(buf.b, hx_if_92)
			xml_4 := Xml_createPCData(buf.b)
			parent.addChild(xml_4)
			nsubs = int(int32((nsubs + 1)))
		}
		return p
	}
	if (!strict && (state == any(18))) && (escapeNext == any(13)) {
		buf.b = hxrt.StringConcatStringPtr(buf.b, hxrt.StringFromLiteral("&"))
		var len_7 any = int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start))))
		var hx_if_93 *string
		if len_7 == nil {
			hx_if_93 = hxrt.StringSubstrStringPtr(str, start, 0, false)
		} else {
			hx_if_93 = hxrt.StringSubstrStringPtr(str, start, len_7.(int), true)
		}
		buf.b = hxrt.StringConcatStringPtr(buf.b, hx_if_93)
		xml_5 := Xml_createPCData(buf.b)
		parent.addChild(xml_5)
		nsubs = int(int32((nsubs + 1)))
		return p
	}
	hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringFromLiteral("Unexpected end"), str, p))
	var hx_throw_zero_94 int
	return hx_throw_zero_94
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
