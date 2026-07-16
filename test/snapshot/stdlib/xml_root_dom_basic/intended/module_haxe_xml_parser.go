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
		hx_post_28 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_28
		var c_1 any = hxrt.StringCharCodeAtAnyStringPtr(xml, i)
		var hx_if_29 int
		if c_1 == nil {
			hx_if_29 = -1
		} else {
			hx_if_29 = c_1.(int)
		}
		c := hx_if_29
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
		var hx_if_55 int
		if c_1 == nil {
			hx_if_55 = -1
		} else {
			hx_if_55 = c_1.(int)
		}
		c := hx_if_55
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
					var hx_if_63 int
					if c_6 == nil {
						hx_if_63 = -1
					} else {
						hx_if_63 = c_6.(int)
					}
					return hx_if_63
				}() == 91 {
					p = int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(2))))
					if !hxrt.StringEqualStringPtr(hxrt.StringToUpperCaseStringPtr(hxrt.StringSubstrStringPtr(str, p, 6, true)), hxrt.StringFromLiteral("CDATA[")) {
						hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringFromLiteral("Expected <![CDATA["), str, p))
						var hx_throw_zero_56 int
						return hx_throw_zero_56
					}
					p = int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(5))))
					state = any(17)
					start = int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1))))
				} else {
					if (func() int {
						var c_4 any = hxrt.StringCharCodeAtAnyStringPtr(str, int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1)))))
						var hx_if_61 int
						if c_4 == nil {
							hx_if_61 = -1
						} else {
							hx_if_61 = c_4.(int)
						}
						return hx_if_61
					}() == 68) || (func() int {
						var c_5 any = hxrt.StringCharCodeAtAnyStringPtr(str, int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1)))))
						var hx_if_62 int
						if c_5 == nil {
							hx_if_62 = -1
						} else {
							hx_if_62 = c_5.(int)
						}
						return hx_if_62
					}() == 100) {
						if !hxrt.StringEqualStringPtr(hxrt.StringToUpperCaseStringPtr(hxrt.StringSubstrStringPtr(str, int(int32((hxrt.Int32Wrap(p)+hxrt.Int32Wrap(2)))), 6, true)), hxrt.StringFromLiteral("OCTYPE")) {
							hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringFromLiteral("Expected <!DOCTYPE"), str, p))
							var hx_throw_zero_57 int
							return hx_throw_zero_57
						}
						p = int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(8))))
						state = any(16)
						start = int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1))))
					} else {
						if (func() int {
							var c_2 any = hxrt.StringCharCodeAtAnyStringPtr(str, int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1)))))
							var hx_if_59 int
							if c_2 == nil {
								hx_if_59 = -1
							} else {
								hx_if_59 = c_2.(int)
							}
							return hx_if_59
						}() != 45) || (func() int {
							var c_3 any = hxrt.StringCharCodeAtAnyStringPtr(str, int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(2)))))
							var hx_if_60 int
							if c_3 == nil {
								hx_if_60 = -1
							} else {
								hx_if_60 = c_3.(int)
							}
							return hx_if_60
						}() != 45) {
							hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringFromLiteral("Expected <!--"), str, p))
							var hx_throw_zero_58 int
							return hx_throw_zero_58
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
					var hx_throw_zero_64 int
					return hx_throw_zero_64
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
					var hx_throw_zero_65 int
					return hx_throw_zero_65
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
					var hx_throw_zero_66 int
					return hx_throw_zero_66
				}
				tmp = hxrt.StringSubstrStringPtr(str, start, int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start)))), true)
				aname = tmp
				if xml.exists(aname) {
					hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Duplicate attribute ["), aname), hxrt.StringFromLiteral("]")), str, p))
					var hx_throw_zero_67 int
					return hx_throw_zero_67
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
				var hx_throw_zero_68 int
				return hx_throw_zero_68
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
				var hx_throw_zero_69 int
				return hx_throw_zero_69
			}
		case 8:
			switch c {
			case 38:
				var len any = int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start))))
				var hx_if_70 *string
				if len == nil {
					hx_if_70 = hxrt.StringSubstrStringPtr(str, start, 0, false)
				} else {
					hx_if_70 = hxrt.StringSubstrStringPtr(str, start, len.(int), true)
				}
				buf.b = hxrt.StringConcatStringPtr(buf.b, hx_if_70)
				state = any(18)
				escapeNext = any(8)
				start = int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1))))
			case 60, 62:
				if strict {
					hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Invalid unescaped "), hxrt.StringFromCharCode(c)), hxrt.StringFromLiteral(" in attribute value")), str, p))
					var hx_throw_zero_71 int
					return hx_throw_zero_71
				} else {
					if c == attrValQuote {
						var len_1 any = int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start))))
						var hx_if_72 *string
						if len_1 == nil {
							hx_if_72 = hxrt.StringSubstrStringPtr(str, start, 0, false)
						} else {
							hx_if_72 = hxrt.StringSubstrStringPtr(str, start, len_1.(int), true)
						}
						buf.b = hxrt.StringConcatStringPtr(buf.b, hx_if_72)
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
					var hx_if_73 *string
					if len_2 == nil {
						hx_if_73 = hxrt.StringSubstrStringPtr(str, start, 0, false)
					} else {
						hx_if_73 = hxrt.StringSubstrStringPtr(str, start, len_2.(int), true)
					}
					buf.b = hxrt.StringConcatStringPtr(buf.b, hx_if_73)
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
					var hx_throw_zero_74 int
					return hx_throw_zero_74
				}
				v := hxrt.StringSubstrStringPtr(str, start, int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start)))), true)
				if (parent == nil) || (parent.nodeType != any(0)) {
					hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Unexpected </"), v), hxrt.StringFromLiteral(">, tag is not open")), str, p))
					var hx_throw_zero_75 int
					return hx_throw_zero_75
				}
				if !hxrt.StringEqualStringPtr(v, parent.get_nodeName()) {
					hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Expected </"), parent.get_nodeName()), hxrt.StringFromLiteral(">")), str, p))
					var hx_throw_zero_76 int
					return hx_throw_zero_76
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
				var hx_throw_zero_77 int
				return hx_throw_zero_77
			}
		case 12:
			if c == 62 {
				if nsubs == 0 {
					parent.addChild(Xml_createPCData(hxrt.StringFromLiteral("")))
				}
				return p
			} else {
				hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringFromLiteral("Expected >"), str, p))
				var hx_throw_zero_78 int
				return hx_throw_zero_78
			}
		case 13:
			if c == 60 {
				var len_3 any = int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start))))
				var hx_if_79 *string
				if len_3 == nil {
					hx_if_79 = hxrt.StringSubstrStringPtr(str, start, 0, false)
				} else {
					hx_if_79 = hxrt.StringSubstrStringPtr(str, start, len_3.(int), true)
				}
				buf.b = hxrt.StringConcatStringPtr(buf.b, hx_if_79)
				child := Xml_createPCData(buf.b)
				buf = New_StringBuf()
				parent.addChild(child)
				nsubs = int(int32((nsubs + 1)))
				state = any(0)
				next = any(2)
			} else {
				if c == 38 {
					var len_4 any = int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start))))
					var hx_if_80 *string
					if len_4 == nil {
						hx_if_80 = hxrt.StringSubstrStringPtr(str, start, 0, false)
					} else {
						hx_if_80 = hxrt.StringSubstrStringPtr(str, start, len_4.(int), true)
					}
					buf.b = hxrt.StringConcatStringPtr(buf.b, hx_if_80)
					state = any(18)
					escapeNext = any(13)
					start = int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1))))
				}
			}
		case 14:
			if (c == 63) && (func() int {
				var c_7 any = hxrt.StringCharCodeAtAnyStringPtr(str, int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(1)))))
				var hx_if_81 int
				if c_7 == nil {
					hx_if_81 = -1
				} else {
					hx_if_81 = c_7.(int)
				}
				return hx_if_81
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
				var hx_if_82 int
				if c_8 == nil {
					hx_if_82 = -1
				} else {
					hx_if_82 = c_8.(int)
				}
				return hx_if_82
			}() == 45)) && (func() int {
				var c_9 any = hxrt.StringCharCodeAtAnyStringPtr(str, int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(2)))))
				var hx_if_83 int
				if c_9 == nil {
					hx_if_83 = -1
				} else {
					hx_if_83 = c_9.(int)
				}
				return hx_if_83
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
				var hx_if_84 int
				if c_10 == nil {
					hx_if_84 = -1
				} else {
					hx_if_84 = c_10.(int)
				}
				return hx_if_84
			}() == 93)) && (func() int {
				var c_11 any = hxrt.StringCharCodeAtAnyStringPtr(str, int(int32((hxrt.Int32Wrap(p) + hxrt.Int32Wrap(2)))))
				var hx_if_85 int
				if c_11 == nil {
					hx_if_85 = -1
				} else {
					hx_if_85 = c_11.(int)
				}
				return hx_if_85
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
					var hx_if_93 int
					if c_15 == nil {
						hx_if_93 = -1
					} else {
						hx_if_93 = c_15.(int)
					}
					return hx_if_93
				}() == 35 {
					var hx_if_87 any
					if func() int {
						var c_13 any = hxrt.StringCharCodeAtAnyStringPtr(s, 1)
						var hx_if_86 int
						if c_13 == nil {
							hx_if_86 = -1
						} else {
							hx_if_86 = c_13.(int)
						}
						return hx_if_86
					}() == 120 {
						hx_if_87 = hxrt.StdParseInt(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("0"), hxrt.StringSubstrStringPtr(s, 1, int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(s))-hxrt.Int32Wrap(1)))), true)))
					} else {
						hx_if_87 = hxrt.StdParseInt(hxrt.StringSubstrStringPtr(s, 1, int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(s)) - hxrt.Int32Wrap(1)))), true))
					}
					var c_12 any = hx_if_87
					c_14 := hxrt.IntFromNullableAny(c_12)
					buf.b = hxrt.StringConcatStringPtr(buf.b, hxrt.StringFromCharCode(c_14))
				} else {
					if !func(hx_value_91 any) bool {
						if hx_value_91 == nil {
							var hx_zero_92 bool
							return hx_zero_92
						}
						return hx_value_91.(bool)
					}(haxe__xml__Parser_escapes.exists(s)) {
						if strict {
							hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Undefined entity: "), s), str, p))
							var hx_throw_zero_88 int
							return hx_throw_zero_88
						}
						buf.b = hxrt.StringConcatStringPtr(buf.b, hxrt.StdString(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("&"), s), hxrt.StringFromLiteral(";"))))
					} else {
						x := func(hx_value_89 any) *string {
							if hx_value_89 == nil {
								var hx_zero_90 *string
								return hx_zero_90
							}
							return hx_value_89.(*string)
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
						var hx_throw_zero_94 int
						return hx_throw_zero_94
					}
					buf.b = hxrt.StringConcatStringPtr(buf.b, hxrt.StringFromLiteral("&"))
					var len_5 any = int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start))))
					var hx_if_95 *string
					if len_5 == nil {
						hx_if_95 = hxrt.StringSubstrStringPtr(str, start, 0, false)
					} else {
						hx_if_95 = hxrt.StringSubstrStringPtr(str, start, len_5.(int), true)
					}
					buf.b = hxrt.StringConcatStringPtr(buf.b, hx_if_95)
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
			var hx_throw_zero_96 int
			return hx_throw_zero_96
		}
		if (p != start) || (nsubs == 0) {
			var len_6 any = int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start))))
			var hx_if_97 *string
			if len_6 == nil {
				hx_if_97 = hxrt.StringSubstrStringPtr(str, start, 0, false)
			} else {
				hx_if_97 = hxrt.StringSubstrStringPtr(str, start, len_6.(int), true)
			}
			buf.b = hxrt.StringConcatStringPtr(buf.b, hx_if_97)
			xml_4 := Xml_createPCData(buf.b)
			parent.addChild(xml_4)
			nsubs = int(int32((nsubs + 1)))
		}
		return p
	}
	if (!strict && (state == any(18))) && (escapeNext == any(13)) {
		buf.b = hxrt.StringConcatStringPtr(buf.b, hxrt.StringFromLiteral("&"))
		var len_7 any = int(int32((hxrt.Int32Wrap(p) - hxrt.Int32Wrap(start))))
		var hx_if_98 *string
		if len_7 == nil {
			hx_if_98 = hxrt.StringSubstrStringPtr(str, start, 0, false)
		} else {
			hx_if_98 = hxrt.StringSubstrStringPtr(str, start, len_7.(int), true)
		}
		buf.b = hxrt.StringConcatStringPtr(buf.b, hx_if_98)
		xml_5 := Xml_createPCData(buf.b)
		parent.addChild(xml_5)
		nsubs = int(int32((nsubs + 1)))
		return p
	}
	hxrt.Throw(New_haxe__xml__XmlParserException(hxrt.StringFromLiteral("Unexpected end"), str, p))
	var hx_throw_zero_99 int
	return hx_throw_zero_99
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
