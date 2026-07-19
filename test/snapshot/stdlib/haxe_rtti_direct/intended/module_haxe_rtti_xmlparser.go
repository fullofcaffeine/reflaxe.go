package main

import "snapshot/hxrt"

type I_haxe__rtti__XmlParser interface {
	sort(l *hxrt.Array)
	sortFields(a *hxrt.Array)
	process(x *Xml, platform *string)
	mergeRights(f1 map[string]any, f2 map[string]any) bool
	mergeDoc(f1 map[string]any, f2 map[string]any) bool
	mergeFields(f map[string]any, f2 map[string]any) bool
	mergeClasses(c map[string]any, c2 map[string]any) bool
	mergeEnums(e map[string]any, e2 map[string]any) bool
	mergeTypedefs(t map[string]any, t2 map[string]any) bool
	mergeAbstracts(a map[string]any, a2 map[string]any) bool
	merge(t *haxe__rtti__TypeTree)
	mkPath(p *string) *string
	mkTypeParams(p *string) *hxrt.Array
	mkRights(r *string) *haxe__rtti__Rights
	xroot(x *Xml)
	processElement(x *Xml) *haxe__rtti__TypeTree
	xmeta(x *Xml) *hxrt.Array
	xoverloads(x *Xml) *hxrt.Array
	xpath(x *Xml) map[string]any
	xclass(x *Xml) map[string]any
	xclassfield(x *Xml, defPublic bool) map[string]any
	xenum(x *Xml) map[string]any
	xenumfield(x *Xml) map[string]any
	xabstract(x *Xml) map[string]any
	xtypedef(x *Xml) map[string]any
	xtype(x *Xml) *haxe__rtti__CType
	xtypeparams(x *Xml) *hxrt.Array
	defplat() *hxrt.Array
	joinStringArray(values *hxrt.Array, separator *string) *string
	splitString(value *string, separator *string) *hxrt.Array
	findSeparator(value *string, separator *string, start int) int
	requireAttr(x *Xml, name *string) *string
	hasNamedElement(x *Xml, name *string) bool
	requireNamedElement(x *Xml, name *string) *Xml
	requireFirstElement(x *Xml) *Xml
	nodeDisplayName(x *Xml) *string
	elementName(x *Xml) *string
	innerData(x *Xml) *string
	innerHTML(x *Xml) *string
	parseIntString(value *string) int
}

type haxe__rtti__XmlParser struct {
	__hx_this   I_haxe__rtti__XmlParser
	root        *hxrt.Array
	curplatform *string
	newField    func(map[string]any, map[string]any)
}

func New_haxe__rtti__XmlParser() *haxe__rtti__XmlParser {
	self := &haxe__rtti__XmlParser{}
	self.__hx_this = self
	self.newField = func(c map[string]any, f map[string]any) {
	}
	self.root = hxrt.NewArray()
	return self
}

func (self *haxe__rtti__XmlParser) sort(l *hxrt.Array) {
	if l == nil {
		l = self.root
	}
	haxe__ds__ArraySort_sort(l, func(hx_cmp_left_23 any, hx_cmp_right_24 any) int {
		return func(e1 *haxe__rtti__TypeTree, e2 *haxe__rtti__TypeTree) int {
			var hx_if_18 *string
			if e1.tag == 0 {
				_g := e1.params[0].(*string)
				_g1 := e1.params[1].(*string)
				_ = _g1
				_g1_1 := e1.params[2].(*hxrt.Array)
				_ = _g1_1
				p := _g
				hx_if_18 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(" "), p)
			} else {
				hx_if_18 = func(hx_obj_15 map[string]any) *string {
					hx_field_16 := hx_obj_15["path"]
					if hx_field_16 == nil {
						var hx_zero_17 *string
						return hx_zero_17
					}
					return hx_field_16.(*string)
				}(haxe__rtti__TypeApi_typeInfos(e1))
			}
			n1 := hx_if_18
			var hx_if_22 *string
			if e2.tag == 0 {
				_g_1 := e2.params[0].(*string)
				_g1_2 := e2.params[1].(*string)
				_ = _g1_2
				_g1_3 := e2.params[2].(*hxrt.Array)
				_ = _g1_3
				p_1 := _g_1
				hx_if_22 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(" "), p_1)
			} else {
				hx_if_22 = func(hx_obj_19 map[string]any) *string {
					hx_field_20 := hx_obj_19["path"]
					if hx_field_20 == nil {
						var hx_zero_21 *string
						return hx_zero_21
					}
					return hx_field_20.(*string)
				}(haxe__rtti__TypeApi_typeInfos(e2))
			}
			n2 := hx_if_22
			return Reflect_compare(n1, n2)
		}(func(hx_value_25 any) *haxe__rtti__TypeTree {
			if hx_value_25 == nil {
				var hx_zero_26 *haxe__rtti__TypeTree
				return hx_zero_26
			}
			return hx_value_25.(*haxe__rtti__TypeTree)
		}(hx_cmp_left_23), func(hx_value_27 any) *haxe__rtti__TypeTree {
			if hx_value_27 == nil {
				var hx_zero_28 *haxe__rtti__TypeTree
				return hx_zero_28
			}
			return hx_value_27.(*haxe__rtti__TypeTree)
		}(hx_cmp_right_24))
	})
	_g := 0
	for _g < l.Len() {
		x := func(hx_value_29 any) *haxe__rtti__TypeTree {
			if hx_value_29 == nil {
				var hx_zero_30 *haxe__rtti__TypeTree
				return hx_zero_30
			}
			return hx_value_29.(*haxe__rtti__TypeTree)
		}(l.Get(_g))
		_g = int(int32((_g + 1)))
		switch x.tag {
		case 0:
			_g_1 := x.params[0].(*string)
			_ = _g_1
			_g_2 := x.params[1].(*string)
			_ = _g_2
			_g_3 := x.params[2].(*hxrt.Array)
			l_1 := _g_3
			self.__hx_this.sort(l_1)
		case 1:
			_g_4 := x.params[0].(map[string]any)
			c := _g_4
			self.__hx_this.sortFields(func(hx_obj_31 map[string]any) *hxrt.Array {
				hx_field_32 := hx_obj_31["fields"]
				if hx_field_32 == nil {
					var hx_zero_33 *hxrt.Array
					return hx_zero_33
				}
				return hx_field_32.(*hxrt.Array)
			}(c))
			self.__hx_this.sortFields(func(hx_obj_34 map[string]any) *hxrt.Array {
				hx_field_35 := hx_obj_34["statics"]
				if hx_field_35 == nil {
					var hx_zero_36 *hxrt.Array
					return hx_zero_36
				}
				return hx_field_35.(*hxrt.Array)
			}(c))
		case 2:
			_g_5 := x.params[0].(map[string]any)
			_ = _g_5
		case 3:
			_g_6 := x.params[0].(map[string]any)
			_ = _g_6
		case 4:
			_g_7 := x.params[0].(map[string]any)
			_ = _g_7
		}
	}
}

func (self *haxe__rtti__XmlParser) sortFields(a *hxrt.Array) {
	haxe__ds__ArraySort_sort(a, func(hx_cmp_left_55 any, hx_cmp_right_56 any) int {
		return func(f1 map[string]any, f2 map[string]any) int {
			v1 := haxe__rtti__TypeApi_isVar(func(hx_obj_37 map[string]any) *haxe__rtti__CType {
				hx_field_38 := hx_obj_37["type"]
				if hx_field_38 == nil {
					var hx_zero_39 *haxe__rtti__CType
					return hx_zero_39
				}
				return hx_field_38.(*haxe__rtti__CType)
			}(f1))
			v2 := haxe__rtti__TypeApi_isVar(func(hx_obj_40 map[string]any) *haxe__rtti__CType {
				hx_field_41 := hx_obj_40["type"]
				if hx_field_41 == nil {
					var hx_zero_42 *haxe__rtti__CType
					return hx_zero_42
				}
				return hx_field_41.(*haxe__rtti__CType)
			}(f2))
			if v1 && !v2 {
				return -1
			}
			if v2 && !v1 {
				return 1
			}
			if hxrt.StringEqualStringPtr(func(hx_obj_43 map[string]any) *string {
				hx_field_44 := hx_obj_43["name"]
				if hx_field_44 == nil {
					var hx_zero_45 *string
					return hx_zero_45
				}
				return hx_field_44.(*string)
			}(f1), hxrt.StringFromLiteral("new")) {
				return -1
			}
			if hxrt.StringEqualStringPtr(func(hx_obj_46 map[string]any) *string {
				hx_field_47 := hx_obj_46["name"]
				if hx_field_47 == nil {
					var hx_zero_48 *string
					return hx_zero_48
				}
				return hx_field_47.(*string)
			}(f2), hxrt.StringFromLiteral("new")) {
				return 1
			}
			return Reflect_compare(func(hx_obj_49 map[string]any) *string {
				hx_field_50 := hx_obj_49["name"]
				if hx_field_50 == nil {
					var hx_zero_51 *string
					return hx_zero_51
				}
				return hx_field_50.(*string)
			}(f1), func(hx_obj_52 map[string]any) *string {
				hx_field_53 := hx_obj_52["name"]
				if hx_field_53 == nil {
					var hx_zero_54 *string
					return hx_zero_54
				}
				return hx_field_53.(*string)
			}(f2))
		}(func(hx_value_57 any) map[string]any {
			if hx_value_57 == nil {
				var hx_zero_58 map[string]any
				return hx_zero_58
			}
			return hx_value_57.(map[string]any)
		}(hx_cmp_left_55), func(hx_value_59 any) map[string]any {
			if hx_value_59 == nil {
				var hx_zero_60 map[string]any
				return hx_zero_60
			}
			return hx_value_59.(map[string]any)
		}(hx_cmp_right_56))
	})
}

func (self *haxe__rtti__XmlParser) process(x *Xml, platform *string) {
	self.curplatform = platform
	self.__hx_this.xroot(x)
}

func (self *haxe__rtti__XmlParser) mergeRights(f1 map[string]any, f2 map[string]any) bool {
	if (((func(hx_obj_61 map[string]any) *haxe__rtti__Rights {
		hx_field_62 := hx_obj_61["get"]
		if hx_field_62 == nil {
			var hx_zero_63 *haxe__rtti__Rights
			return hx_zero_63
		}
		return hx_field_62.(*haxe__rtti__Rights)
	}(f1) == haxe__rtti__Rights_RInline) && (func(hx_obj_64 map[string]any) *haxe__rtti__Rights {
		hx_field_65 := hx_obj_64["set"]
		if hx_field_65 == nil {
			var hx_zero_66 *haxe__rtti__Rights
			return hx_zero_66
		}
		return hx_field_65.(*haxe__rtti__Rights)
	}(f1) == haxe__rtti__Rights_RNo)) && (func(hx_obj_67 map[string]any) *haxe__rtti__Rights {
		hx_field_68 := hx_obj_67["get"]
		if hx_field_68 == nil {
			var hx_zero_69 *haxe__rtti__Rights
			return hx_zero_69
		}
		return hx_field_68.(*haxe__rtti__Rights)
	}(f2) == haxe__rtti__Rights_RNormal)) && (func(hx_obj_70 map[string]any) *haxe__rtti__Rights {
		hx_field_71 := hx_obj_70["set"]
		if hx_field_71 == nil {
			var hx_zero_72 *haxe__rtti__Rights
			return hx_zero_72
		}
		return hx_field_71.(*haxe__rtti__Rights)
	}(f2) == haxe__rtti__Rights_RMethod) {
		f1["get"] = haxe__rtti__Rights_RNormal
		f1["set"] = haxe__rtti__Rights_RMethod
		return true
	}
	return (Type_enumEq(func(hx_obj_73 map[string]any) *haxe__rtti__Rights {
		hx_field_74 := hx_obj_73["get"]
		if hx_field_74 == nil {
			var hx_zero_75 *haxe__rtti__Rights
			return hx_zero_75
		}
		return hx_field_74.(*haxe__rtti__Rights)
	}(f1), func(hx_obj_76 map[string]any) *haxe__rtti__Rights {
		hx_field_77 := hx_obj_76["get"]
		if hx_field_77 == nil {
			var hx_zero_78 *haxe__rtti__Rights
			return hx_zero_78
		}
		return hx_field_77.(*haxe__rtti__Rights)
	}(f2)) && Type_enumEq(func(hx_obj_79 map[string]any) *haxe__rtti__Rights {
		hx_field_80 := hx_obj_79["set"]
		if hx_field_80 == nil {
			var hx_zero_81 *haxe__rtti__Rights
			return hx_zero_81
		}
		return hx_field_80.(*haxe__rtti__Rights)
	}(f1), func(hx_obj_82 map[string]any) *haxe__rtti__Rights {
		hx_field_83 := hx_obj_82["set"]
		if hx_field_83 == nil {
			var hx_zero_84 *haxe__rtti__Rights
			return hx_zero_84
		}
		return hx_field_83.(*haxe__rtti__Rights)
	}(f2)))
}

func (self *haxe__rtti__XmlParser) mergeDoc(f1 map[string]any, f2 map[string]any) bool {
	if hxrt.StringEqualStringPtr(func(hx_obj_94 map[string]any) *string {
		hx_field_95 := hx_obj_94["doc"]
		if hx_field_95 == nil {
			var hx_zero_96 *string
			return hx_zero_96
		}
		return hx_field_95.(*string)
	}(f1), nil) {
		f1["doc"] = func(hx_obj_85 map[string]any) *string {
			hx_field_86 := hx_obj_85["doc"]
			if hx_field_86 == nil {
				var hx_zero_87 *string
				return hx_zero_87
			}
			return hx_field_86.(*string)
		}(f2)
	} else {
		if hxrt.StringEqualStringPtr(func(hx_obj_91 map[string]any) *string {
			hx_field_92 := hx_obj_91["doc"]
			if hx_field_92 == nil {
				var hx_zero_93 *string
				return hx_zero_93
			}
			return hx_field_92.(*string)
		}(f2), nil) {
			f2["doc"] = func(hx_obj_88 map[string]any) *string {
				hx_field_89 := hx_obj_88["doc"]
				if hx_field_89 == nil {
					var hx_zero_90 *string
					return hx_zero_90
				}
				return hx_field_89.(*string)
			}(f1)
		}
	}
	return true
}

func (self *haxe__rtti__XmlParser) mergeFields(f map[string]any, f2 map[string]any) bool {
	return (haxe__rtti__TypeApi_fieldEq(f, f2) || (((hxrt.StringEqualStringPtr(func(hx_obj_97 map[string]any) *string {
		hx_field_98 := hx_obj_97["name"]
		if hx_field_98 == nil {
			var hx_zero_99 *string
			return hx_zero_99
		}
		return hx_field_98.(*string)
	}(f), func(hx_obj_100 map[string]any) *string {
		hx_field_101 := hx_obj_100["name"]
		if hx_field_101 == nil {
			var hx_zero_102 *string
			return hx_zero_102
		}
		return hx_field_101.(*string)
	}(f2)) && (self.__hx_this.mergeRights(f, f2) || self.__hx_this.mergeRights(f2, f))) && self.__hx_this.mergeDoc(f, f2)) && haxe__rtti__TypeApi_fieldEq(f, f2)))
}

func (self *haxe__rtti__XmlParser) mergeClasses(c map[string]any, c2 map[string]any) bool {
	if func(hx_obj_103 map[string]any) bool {
		hx_field_104 := hx_obj_103["isInterface"]
		if hx_field_104 == nil {
			var hx_zero_105 bool
			return hx_zero_105
		}
		return hx_field_104.(bool)
	}(c) != func(hx_obj_106 map[string]any) bool {
		hx_field_107 := hx_obj_106["isInterface"]
		if hx_field_107 == nil {
			var hx_zero_108 bool
			return hx_zero_108
		}
		return hx_field_107.(bool)
	}(c2) {
		return false
	}
	if !hxrt.StringEqualStringPtr(self.curplatform, nil) {
		hx_arr_109 := func(hx_obj_110 map[string]any) *hxrt.Array {
			hx_field_111 := hx_obj_110["platforms"]
			if hx_field_111 == nil {
				var hx_zero_112 *hxrt.Array
				return hx_zero_112
			}
			return hx_field_111.(*hxrt.Array)
		}(c)
		hx_arr_109.Push(self.curplatform)
	}
	if func(hx_obj_113 map[string]any) bool {
		hx_field_114 := hx_obj_113["isExtern"]
		if hx_field_114 == nil {
			var hx_zero_115 bool
			return hx_zero_115
		}
		return hx_field_114.(bool)
	}(c) != func(hx_obj_116 map[string]any) bool {
		hx_field_117 := hx_obj_116["isExtern"]
		if hx_field_117 == nil {
			var hx_zero_118 bool
			return hx_zero_118
		}
		return hx_field_117.(bool)
	}(c2) {
		c["isExtern"] = false
	}
	_g := 0
	_g1 := func(hx_obj_119 map[string]any) *hxrt.Array {
		hx_field_120 := hx_obj_119["fields"]
		if hx_field_120 == nil {
			var hx_zero_121 *hxrt.Array
			return hx_zero_121
		}
		return hx_field_120.(*hxrt.Array)
	}(c2)
	for _g < _g1.Len() {
		f2 := func(hx_value_122 any) map[string]any {
			if hx_value_122 == nil {
				var hx_zero_123 map[string]any
				return hx_zero_123
			}
			return hx_value_122.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		var found map[string]any = nil
		_g_1 := 0
		_g1_1 := func(hx_obj_124 map[string]any) *hxrt.Array {
			hx_field_125 := hx_obj_124["fields"]
			if hx_field_125 == nil {
				var hx_zero_126 *hxrt.Array
				return hx_zero_126
			}
			return hx_field_125.(*hxrt.Array)
		}(c)
		for _g_1 < _g1_1.Len() {
			f := func(hx_value_127 any) map[string]any {
				if hx_value_127 == nil {
					var hx_zero_128 map[string]any
					return hx_zero_128
				}
				return hx_value_127.(map[string]any)
			}(_g1_1.Get(_g_1))
			_g_1 = int(int32((_g_1 + 1)))
			if self.__hx_this.mergeFields(f, f2) {
				found = f
				break
			}
		}
		if found == nil {
			func(hx_fn func(map[string]any, map[string]any), hx_arg_0 map[string]any, hx_arg_1 map[string]any) {
				if hx_fn == nil {
					hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
					return
				}
				hx_fn(hx_arg_0, hx_arg_1)
			}(self.newField, c, f2)
			hx_arr_129 := func(hx_obj_130 map[string]any) *hxrt.Array {
				hx_field_131 := hx_obj_130["fields"]
				if hx_field_131 == nil {
					var hx_zero_132 *hxrt.Array
					return hx_zero_132
				}
				return hx_field_131.(*hxrt.Array)
			}(c)
			hx_arr_129.Push(f2)
		} else {
			if !hxrt.StringEqualStringPtr(self.curplatform, nil) {
				hx_arr_133 := func(hx_obj_134 map[string]any) *hxrt.Array {
					hx_field_135 := hx_obj_134["platforms"]
					if hx_field_135 == nil {
						var hx_zero_136 *hxrt.Array
						return hx_zero_136
					}
					return hx_field_135.(*hxrt.Array)
				}(found)
				hx_arr_133.Push(self.curplatform)
			}
		}
	}
	_g_2 := 0
	_g1_2 := func(hx_obj_137 map[string]any) *hxrt.Array {
		hx_field_138 := hx_obj_137["statics"]
		if hx_field_138 == nil {
			var hx_zero_139 *hxrt.Array
			return hx_zero_139
		}
		return hx_field_138.(*hxrt.Array)
	}(c2)
	for _g_2 < _g1_2.Len() {
		f2_1 := func(hx_value_140 any) map[string]any {
			if hx_value_140 == nil {
				var hx_zero_141 map[string]any
				return hx_zero_141
			}
			return hx_value_140.(map[string]any)
		}(_g1_2.Get(_g_2))
		_g_2 = int(int32((_g_2 + 1)))
		var found_1 map[string]any = nil
		_g_3 := 0
		_g1_3 := func(hx_obj_142 map[string]any) *hxrt.Array {
			hx_field_143 := hx_obj_142["statics"]
			if hx_field_143 == nil {
				var hx_zero_144 *hxrt.Array
				return hx_zero_144
			}
			return hx_field_143.(*hxrt.Array)
		}(c)
		for _g_3 < _g1_3.Len() {
			f_1 := func(hx_value_145 any) map[string]any {
				if hx_value_145 == nil {
					var hx_zero_146 map[string]any
					return hx_zero_146
				}
				return hx_value_145.(map[string]any)
			}(_g1_3.Get(_g_3))
			_g_3 = int(int32((_g_3 + 1)))
			if self.__hx_this.mergeFields(f_1, f2_1) {
				found_1 = f_1
				break
			}
		}
		if found_1 == nil {
			func(hx_fn func(map[string]any, map[string]any), hx_arg_0 map[string]any, hx_arg_1 map[string]any) {
				if hx_fn == nil {
					hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
					return
				}
				hx_fn(hx_arg_0, hx_arg_1)
			}(self.newField, c, f2_1)
			hx_arr_147 := func(hx_obj_148 map[string]any) *hxrt.Array {
				hx_field_149 := hx_obj_148["statics"]
				if hx_field_149 == nil {
					var hx_zero_150 *hxrt.Array
					return hx_zero_150
				}
				return hx_field_149.(*hxrt.Array)
			}(c)
			hx_arr_147.Push(f2_1)
		} else {
			if !hxrt.StringEqualStringPtr(self.curplatform, nil) {
				hx_arr_151 := func(hx_obj_152 map[string]any) *hxrt.Array {
					hx_field_153 := hx_obj_152["platforms"]
					if hx_field_153 == nil {
						var hx_zero_154 *hxrt.Array
						return hx_zero_154
					}
					return hx_field_153.(*hxrt.Array)
				}(found_1)
				hx_arr_151.Push(self.curplatform)
			}
		}
	}
	return true
}

func (self *haxe__rtti__XmlParser) mergeEnums(e map[string]any, e2 map[string]any) bool {
	if func(hx_obj_155 map[string]any) bool {
		hx_field_156 := hx_obj_155["isExtern"]
		if hx_field_156 == nil {
			var hx_zero_157 bool
			return hx_zero_157
		}
		return hx_field_156.(bool)
	}(e) != func(hx_obj_158 map[string]any) bool {
		hx_field_159 := hx_obj_158["isExtern"]
		if hx_field_159 == nil {
			var hx_zero_160 bool
			return hx_zero_160
		}
		return hx_field_159.(bool)
	}(e2) {
		return false
	}
	if !hxrt.StringEqualStringPtr(self.curplatform, nil) {
		hx_arr_161 := func(hx_obj_162 map[string]any) *hxrt.Array {
			hx_field_163 := hx_obj_162["platforms"]
			if hx_field_163 == nil {
				var hx_zero_164 *hxrt.Array
				return hx_zero_164
			}
			return hx_field_163.(*hxrt.Array)
		}(e)
		hx_arr_161.Push(self.curplatform)
	}
	_g := 0
	_g1 := func(hx_obj_165 map[string]any) *hxrt.Array {
		hx_field_166 := hx_obj_165["constructors"]
		if hx_field_166 == nil {
			var hx_zero_167 *hxrt.Array
			return hx_zero_167
		}
		return hx_field_166.(*hxrt.Array)
	}(e2)
	for _g < _g1.Len() {
		c2 := func(hx_value_168 any) map[string]any {
			if hx_value_168 == nil {
				var hx_zero_169 map[string]any
				return hx_zero_169
			}
			return hx_value_168.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		var found map[string]any = nil
		_g_1 := 0
		_g1_1 := func(hx_obj_170 map[string]any) *hxrt.Array {
			hx_field_171 := hx_obj_170["constructors"]
			if hx_field_171 == nil {
				var hx_zero_172 *hxrt.Array
				return hx_zero_172
			}
			return hx_field_171.(*hxrt.Array)
		}(e)
		for _g_1 < _g1_1.Len() {
			c := func(hx_value_173 any) map[string]any {
				if hx_value_173 == nil {
					var hx_zero_174 map[string]any
					return hx_zero_174
				}
				return hx_value_173.(map[string]any)
			}(_g1_1.Get(_g_1))
			_g_1 = int(int32((_g_1 + 1)))
			if haxe__rtti__TypeApi_constructorEq(c, c2) {
				found = c
				break
			}
		}
		if found == nil {
			hx_arr_175 := func(hx_obj_176 map[string]any) *hxrt.Array {
				hx_field_177 := hx_obj_176["constructors"]
				if hx_field_177 == nil {
					var hx_zero_178 *hxrt.Array
					return hx_zero_178
				}
				return hx_field_177.(*hxrt.Array)
			}(e)
			hx_arr_175.Push(c2)
		} else {
			if !hxrt.StringEqualStringPtr(self.curplatform, nil) {
				hx_arr_179 := func(hx_obj_180 map[string]any) *hxrt.Array {
					hx_field_181 := hx_obj_180["platforms"]
					if hx_field_181 == nil {
						var hx_zero_182 *hxrt.Array
						return hx_zero_182
					}
					return hx_field_181.(*hxrt.Array)
				}(found)
				hx_arr_179.Push(self.curplatform)
			}
		}
	}
	return true
}

func (self *haxe__rtti__XmlParser) mergeTypedefs(t map[string]any, t2 map[string]any) bool {
	if hxrt.StringEqualStringPtr(self.curplatform, nil) {
		return false
	}
	hx_arr_183 := func(hx_obj_184 map[string]any) *hxrt.Array {
		hx_field_185 := hx_obj_184["platforms"]
		if hx_field_185 == nil {
			var hx_zero_186 *hxrt.Array
			return hx_zero_186
		}
		return hx_field_185.(*hxrt.Array)
	}(t)
	hx_arr_183.Push(self.curplatform)
	var this1 haxe__IMap = func(hx_obj_187 map[string]any) *haxe__ds__StringMap {
		hx_field_188 := hx_obj_187["types"]
		if hx_field_188 == nil {
			var hx_zero_189 *haxe__ds__StringMap
			return hx_zero_189
		}
		return hx_field_188.(*haxe__ds__StringMap)
	}(t)
	key := self.curplatform
	value := func(hx_obj_190 map[string]any) *haxe__rtti__CType {
		hx_field_191 := hx_obj_190["type"]
		if hx_field_191 == nil {
			var hx_zero_192 *haxe__rtti__CType
			return hx_zero_192
		}
		return hx_field_191.(*haxe__rtti__CType)
	}(t2)
	this1.(*haxe__ds__StringMap).__hx_this.set(key, value)
	return true
}

func (self *haxe__rtti__XmlParser) mergeAbstracts(a map[string]any, a2 map[string]any) bool {
	if hxrt.StringEqualStringPtr(self.curplatform, nil) {
		return false
	}
	if (func(hx_obj_193 map[string]any) *hxrt.Array {
		hx_field_194 := hx_obj_193["to"]
		if hx_field_194 == nil {
			var hx_zero_195 *hxrt.Array
			return hx_zero_195
		}
		return hx_field_194.(*hxrt.Array)
	}(a).Len() != func(hx_obj_196 map[string]any) *hxrt.Array {
		hx_field_197 := hx_obj_196["to"]
		if hx_field_197 == nil {
			var hx_zero_198 *hxrt.Array
			return hx_zero_198
		}
		return hx_field_197.(*hxrt.Array)
	}(a2).Len()) || (func(hx_obj_199 map[string]any) *hxrt.Array {
		hx_field_200 := hx_obj_199["from"]
		if hx_field_200 == nil {
			var hx_zero_201 *hxrt.Array
			return hx_zero_201
		}
		return hx_field_200.(*hxrt.Array)
	}(a).Len() != func(hx_obj_202 map[string]any) *hxrt.Array {
		hx_field_203 := hx_obj_202["from"]
		if hx_field_203 == nil {
			var hx_zero_204 *hxrt.Array
			return hx_zero_204
		}
		return hx_field_203.(*hxrt.Array)
	}(a2).Len()) {
		return false
	}
	_g := 0
	_g1 := func(hx_obj_205 map[string]any) *hxrt.Array {
		hx_field_206 := hx_obj_205["to"]
		if hx_field_206 == nil {
			var hx_zero_207 *hxrt.Array
			return hx_zero_207
		}
		return hx_field_206.(*hxrt.Array)
	}(a).Len()
	for _g < _g1 {
		hx_post_208 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_208
		if !haxe__rtti__TypeApi_typeEq(func(hx_obj_214 map[string]any) *haxe__rtti__CType {
			hx_field_215 := hx_obj_214["t"]
			if hx_field_215 == nil {
				var hx_zero_216 *haxe__rtti__CType
				return hx_zero_216
			}
			return hx_field_215.(*haxe__rtti__CType)
		}(func(hx_value_212 any) map[string]any {
			if hx_value_212 == nil {
				var hx_zero_213 map[string]any
				return hx_zero_213
			}
			return hx_value_212.(map[string]any)
		}(func(hx_obj_209 map[string]any) *hxrt.Array {
			hx_field_210 := hx_obj_209["to"]
			if hx_field_210 == nil {
				var hx_zero_211 *hxrt.Array
				return hx_zero_211
			}
			return hx_field_210.(*hxrt.Array)
		}(a).Get(i))), func(hx_obj_222 map[string]any) *haxe__rtti__CType {
			hx_field_223 := hx_obj_222["t"]
			if hx_field_223 == nil {
				var hx_zero_224 *haxe__rtti__CType
				return hx_zero_224
			}
			return hx_field_223.(*haxe__rtti__CType)
		}(func(hx_value_220 any) map[string]any {
			if hx_value_220 == nil {
				var hx_zero_221 map[string]any
				return hx_zero_221
			}
			return hx_value_220.(map[string]any)
		}(func(hx_obj_217 map[string]any) *hxrt.Array {
			hx_field_218 := hx_obj_217["to"]
			if hx_field_218 == nil {
				var hx_zero_219 *hxrt.Array
				return hx_zero_219
			}
			return hx_field_218.(*hxrt.Array)
		}(a2).Get(i)))) {
			return false
		}
	}
	_g_1 := 0
	_g1_1 := func(hx_obj_225 map[string]any) *hxrt.Array {
		hx_field_226 := hx_obj_225["from"]
		if hx_field_226 == nil {
			var hx_zero_227 *hxrt.Array
			return hx_zero_227
		}
		return hx_field_226.(*hxrt.Array)
	}(a).Len()
	for _g_1 < _g1_1 {
		hx_post_228 := _g_1
		_g_1 = int(int32((_g_1 + 1)))
		i_1 := hx_post_228
		if !haxe__rtti__TypeApi_typeEq(func(hx_obj_234 map[string]any) *haxe__rtti__CType {
			hx_field_235 := hx_obj_234["t"]
			if hx_field_235 == nil {
				var hx_zero_236 *haxe__rtti__CType
				return hx_zero_236
			}
			return hx_field_235.(*haxe__rtti__CType)
		}(func(hx_value_232 any) map[string]any {
			if hx_value_232 == nil {
				var hx_zero_233 map[string]any
				return hx_zero_233
			}
			return hx_value_232.(map[string]any)
		}(func(hx_obj_229 map[string]any) *hxrt.Array {
			hx_field_230 := hx_obj_229["from"]
			if hx_field_230 == nil {
				var hx_zero_231 *hxrt.Array
				return hx_zero_231
			}
			return hx_field_230.(*hxrt.Array)
		}(a).Get(i_1))), func(hx_obj_242 map[string]any) *haxe__rtti__CType {
			hx_field_243 := hx_obj_242["t"]
			if hx_field_243 == nil {
				var hx_zero_244 *haxe__rtti__CType
				return hx_zero_244
			}
			return hx_field_243.(*haxe__rtti__CType)
		}(func(hx_value_240 any) map[string]any {
			if hx_value_240 == nil {
				var hx_zero_241 map[string]any
				return hx_zero_241
			}
			return hx_value_240.(map[string]any)
		}(func(hx_obj_237 map[string]any) *hxrt.Array {
			hx_field_238 := hx_obj_237["from"]
			if hx_field_238 == nil {
				var hx_zero_239 *hxrt.Array
				return hx_zero_239
			}
			return hx_field_238.(*hxrt.Array)
		}(a2).Get(i_1)))) {
			return false
		}
	}
	if func(hx_obj_251 map[string]any) map[string]any {
		hx_field_252 := hx_obj_251["impl"]
		if hx_field_252 == nil {
			var hx_zero_253 map[string]any
			return hx_zero_253
		}
		return hx_field_252.(map[string]any)
	}(a2) != nil {
		self.__hx_this.mergeClasses(func(hx_obj_245 map[string]any) map[string]any {
			hx_field_246 := hx_obj_245["impl"]
			if hx_field_246 == nil {
				var hx_zero_247 map[string]any
				return hx_zero_247
			}
			return hx_field_246.(map[string]any)
		}(a), func(hx_obj_248 map[string]any) map[string]any {
			hx_field_249 := hx_obj_248["impl"]
			if hx_field_249 == nil {
				var hx_zero_250 map[string]any
				return hx_zero_250
			}
			return hx_field_249.(map[string]any)
		}(a2))
	}
	hx_arr_254 := func(hx_obj_255 map[string]any) *hxrt.Array {
		hx_field_256 := hx_obj_255["platforms"]
		if hx_field_256 == nil {
			var hx_zero_257 *hxrt.Array
			return hx_zero_257
		}
		return hx_field_256.(*hxrt.Array)
	}(a)
	hx_arr_254.Push(self.curplatform)
	return true
}

func (self *haxe__rtti__XmlParser) merge(t *haxe__rtti__TypeTree) {
	inf := haxe__rtti__TypeApi_typeInfos(t)
	pack := self.__hx_this.splitString(func(hx_obj_258 map[string]any) *string {
		hx_field_259 := hx_obj_258["path"]
		if hx_field_259 == nil {
			var hx_zero_260 *string
			return hx_zero_260
		}
		return hx_field_259.(*string)
	}(inf), hxrt.StringFromLiteral("."))
	cur := self.root
	curpack := hxrt.NewArray()
	pack.Pop()
	_g := 0
	for _g < pack.Len() {
		p := func(hx_value_262 any) *string {
			if hx_value_262 == nil {
				var hx_zero_263 *string
				return hx_zero_263
			}
			return hx_value_262.(*string)
		}(pack.Get(_g))
		_g = int(int32((_g + 1)))
		found := false
		_g_1 := 0
		for _g_1 < cur.Len() {
			pk := func(hx_value_264 any) *haxe__rtti__TypeTree {
				if hx_value_264 == nil {
					var hx_zero_265 *haxe__rtti__TypeTree
					return hx_zero_265
				}
				return hx_value_264.(*haxe__rtti__TypeTree)
			}(cur.Get(_g_1))
			_g_1 = int(int32((_g_1 + 1)))
			if pk.tag == 0 {
				_g_2 := pk.params[0].(*string)
				_g1 := pk.params[1].(*string)
				_ = _g1
				_g1_1 := pk.params[2].(*hxrt.Array)
				pname := _g_2
				subs := _g1_1
				if hxrt.StringEqualStringPtr(pname, p) {
					found = true
					cur = subs
					break
				}
			} else {
			}
		}
		curpack.Push(p)
		if !found {
			pk_1 := hxrt.NewArray()
			cur.Push(haxe__rtti__TypeTree_TPackage(p, self.__hx_this.joinStringArray(curpack, hxrt.StringFromLiteral(".")), pk_1))
			cur = pk_1
		}
	}
	_g_3 := 0
	for _g_3 < cur.Len() {
		ct := func(hx_value_268 any) *haxe__rtti__TypeTree {
			if hx_value_268 == nil {
				var hx_zero_269 *haxe__rtti__TypeTree
				return hx_zero_269
			}
			return hx_value_268.(*haxe__rtti__TypeTree)
		}(cur.Get(_g_3))
		_g_3 = int(int32((_g_3 + 1)))
		if func() bool {
			var hx_if_270 bool
			if ct.tag == 0 {
				_g_4 := ct.params[0].(*string)
				_ = _g_4
				_g_5 := ct.params[1].(*string)
				_ = _g_5
				_g_6 := ct.params[2].(*hxrt.Array)
				_ = _g_6
				hx_if_270 = true
			} else {
				hx_if_270 = false
			}
			return hx_if_270
		}() {
			continue
		}
		tinf := haxe__rtti__TypeApi_typeInfos(ct)
		if hxrt.StringEqualStringPtr(func(hx_obj_341 map[string]any) *string {
			hx_field_342 := hx_obj_341["path"]
			if hx_field_342 == nil {
				var hx_zero_343 *string
				return hx_zero_343
			}
			return hx_field_342.(*string)
		}(tinf), func(hx_obj_344 map[string]any) *string {
			hx_field_345 := hx_obj_344["path"]
			if hx_field_345 == nil {
				var hx_zero_346 *string
				return hx_zero_346
			}
			return hx_field_345.(*string)
		}(inf)) {
			sameType := true
			if hxrt.StringEqualStringPtr(func(hx_obj_280 map[string]any) *string {
				hx_field_281 := hx_obj_280["doc"]
				if hx_field_281 == nil {
					var hx_zero_282 *string
					return hx_zero_282
				}
				return hx_field_281.(*string)
			}(tinf), nil) != hxrt.StringEqualStringPtr(func(hx_obj_283 map[string]any) *string {
				hx_field_284 := hx_obj_283["doc"]
				if hx_field_284 == nil {
					var hx_zero_285 *string
					return hx_zero_285
				}
				return hx_field_284.(*string)
			}(inf), nil) {
				if hxrt.StringEqualStringPtr(func(hx_obj_277 map[string]any) *string {
					hx_field_278 := hx_obj_277["doc"]
					if hx_field_278 == nil {
						var hx_zero_279 *string
						return hx_zero_279
					}
					return hx_field_278.(*string)
				}(inf), nil) {
					inf["doc"] = func(hx_obj_271 map[string]any) *string {
						hx_field_272 := hx_obj_271["doc"]
						if hx_field_272 == nil {
							var hx_zero_273 *string
							return hx_zero_273
						}
						return hx_field_272.(*string)
					}(tinf)
				} else {
					tinf["doc"] = func(hx_obj_274 map[string]any) *string {
						hx_field_275 := hx_obj_274["doc"]
						if hx_field_275 == nil {
							var hx_zero_276 *string
							return hx_zero_276
						}
						return hx_field_275.(*string)
					}(inf)
				}
			}
			if hxrt.StringEqualStringPtr(func(hx_obj_286 map[string]any) *string {
				hx_field_287 := hx_obj_286["path"]
				if hx_field_287 == nil {
					var hx_zero_288 *string
					return hx_zero_288
				}
				return hx_field_287.(*string)
			}(tinf), hxrt.StringFromLiteral("haxe._Int64.NativeInt64")) {
				continue
			}
			if (hxrt.StringEqualStringPtr(func(hx_obj_289 map[string]any) *string {
				hx_field_290 := hx_obj_289["module"]
				if hx_field_290 == nil {
					var hx_zero_291 *string
					return hx_zero_291
				}
				return hx_field_290.(*string)
			}(tinf), func(hx_obj_292 map[string]any) *string {
				hx_field_293 := hx_obj_292["module"]
				if hx_field_293 == nil {
					var hx_zero_294 *string
					return hx_zero_294
				}
				return hx_field_293.(*string)
			}(inf)) && hxrt.StringEqualStringPtr(func(hx_obj_295 map[string]any) *string {
				hx_field_296 := hx_obj_295["doc"]
				if hx_field_296 == nil {
					var hx_zero_297 *string
					return hx_zero_297
				}
				return hx_field_296.(*string)
			}(tinf), func(hx_obj_298 map[string]any) *string {
				hx_field_299 := hx_obj_298["doc"]
				if hx_field_299 == nil {
					var hx_zero_300 *string
					return hx_zero_300
				}
				return hx_field_299.(*string)
			}(inf))) && (func(hx_obj_301 map[string]any) bool {
				hx_field_302 := hx_obj_301["isPrivate"]
				if hx_field_302 == nil {
					var hx_zero_303 bool
					return hx_zero_303
				}
				return hx_field_302.(bool)
			}(tinf) == func(hx_obj_304 map[string]any) bool {
				hx_field_305 := hx_obj_304["isPrivate"]
				if hx_field_305 == nil {
					var hx_zero_306 bool
					return hx_zero_306
				}
				return hx_field_305.(bool)
			}(inf)) {
				switch ct.tag {
				case 0:
					_g_7 := ct.params[0].(*string)
					_ = _g_7
					_g_8 := ct.params[1].(*string)
					_ = _g_8
					_g_9 := ct.params[2].(*hxrt.Array)
					_ = _g_9
					sameType = false
				case 1:
					_g_10 := ct.params[0].(map[string]any)
					c := _g_10
					if t.tag == 1 {
						_g_11 := t.params[0].(map[string]any)
						c2 := _g_11
						if self.__hx_this.mergeClasses(c, c2) {
							return
						}
					} else {
						sameType = false
					}
				case 2:
					_g_12 := ct.params[0].(map[string]any)
					e := _g_12
					if t.tag == 2 {
						_g_13 := t.params[0].(map[string]any)
						e2 := _g_13
						if self.__hx_this.mergeEnums(e, e2) {
							return
						}
					} else {
						sameType = false
					}
				case 3:
					_g_14 := ct.params[0].(map[string]any)
					td := _g_14
					if t.tag == 3 {
						_g_15 := t.params[0].(map[string]any)
						td2 := _g_15
						if self.__hx_this.mergeTypedefs(td, td2) {
							return
						}
					} else {
					}
				case 4:
					_g_16 := ct.params[0].(map[string]any)
					a := _g_16
					if t.tag == 4 {
						_g_17 := t.params[0].(map[string]any)
						a2 := _g_17
						if self.__hx_this.mergeAbstracts(a, a2) {
							return
						}
					} else {
						sameType = false
					}
				}
			}
			var hx_if_334 *string
			if !hxrt.StringEqualStringPtr(func(hx_obj_307 map[string]any) *string {
				hx_field_308 := hx_obj_307["module"]
				if hx_field_308 == nil {
					var hx_zero_309 *string
					return hx_zero_309
				}
				return hx_field_308.(*string)
			}(tinf), func(hx_obj_310 map[string]any) *string {
				hx_field_311 := hx_obj_310["module"]
				if hx_field_311 == nil {
					var hx_zero_312 *string
					return hx_zero_312
				}
				return hx_field_311.(*string)
			}(inf)) {
				hx_if_334 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("module "), func(hx_obj_313 map[string]any) *string {
					hx_field_314 := hx_obj_313["module"]
					if hx_field_314 == nil {
						var hx_zero_315 *string
						return hx_zero_315
					}
					return hx_field_314.(*string)
				}(inf)), hxrt.StringFromLiteral(" should be ")), func(hx_obj_316 map[string]any) *string {
					hx_field_317 := hx_obj_316["module"]
					if hx_field_317 == nil {
						var hx_zero_318 *string
						return hx_zero_318
					}
					return hx_field_317.(*string)
				}(tinf))
			} else {
				var hx_if_333 *string
				if !hxrt.StringEqualStringPtr(func(hx_obj_319 map[string]any) *string {
					hx_field_320 := hx_obj_319["doc"]
					if hx_field_320 == nil {
						var hx_zero_321 *string
						return hx_zero_321
					}
					return hx_field_320.(*string)
				}(tinf), func(hx_obj_322 map[string]any) *string {
					hx_field_323 := hx_obj_322["doc"]
					if hx_field_323 == nil {
						var hx_zero_324 *string
						return hx_zero_324
					}
					return hx_field_323.(*string)
				}(inf)) {
					hx_if_333 = hxrt.StringFromLiteral("documentation is different")
				} else {
					var hx_if_332 *string
					if func(hx_obj_325 map[string]any) bool {
						hx_field_326 := hx_obj_325["isPrivate"]
						if hx_field_326 == nil {
							var hx_zero_327 bool
							return hx_zero_327
						}
						return hx_field_326.(bool)
					}(tinf) != func(hx_obj_328 map[string]any) bool {
						hx_field_329 := hx_obj_328["isPrivate"]
						if hx_field_329 == nil {
							var hx_zero_330 bool
							return hx_zero_330
						}
						return hx_field_329.(bool)
					}(inf) {
						hx_if_332 = hxrt.StringFromLiteral("private flag is different")
					} else {
						var hx_if_331 *string
						if !sameType {
							hx_if_331 = hxrt.StringFromLiteral("type kind is different")
						} else {
							hx_if_331 = hxrt.StringFromLiteral("could not merge definition")
						}
						hx_if_332 = hx_if_331
					}
					hx_if_333 = hx_if_332
				}
				hx_if_334 = hx_if_333
			}
			msg := hx_if_334
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Incompatibilities between "), func(hx_obj_335 map[string]any) *string {
				hx_field_336 := hx_obj_335["path"]
				if hx_field_336 == nil {
					var hx_zero_337 *string
					return hx_zero_337
				}
				return hx_field_336.(*string)
			}(tinf)), hxrt.StringFromLiteral(" in ")), self.__hx_this.joinStringArray(func(hx_obj_338 map[string]any) *hxrt.Array {
				hx_field_339 := hx_obj_338["platforms"]
				if hx_field_339 == nil {
					var hx_zero_340 *hxrt.Array
					return hx_zero_340
				}
				return hx_field_339.(*hxrt.Array)
			}(tinf), hxrt.StringFromLiteral(","))), hxrt.StringFromLiteral(" and ")), self.curplatform), hxrt.StringFromLiteral(" (")), msg), hxrt.StringFromLiteral(")")))
		}
	}
	cur.Push(t)
}

func (self *haxe__rtti__XmlParser) mkPath(p *string) *string {
	return p
}

func (self *haxe__rtti__XmlParser) mkTypeParams(p *string) *hxrt.Array {
	pl := self.__hx_this.splitString(p, hxrt.StringFromLiteral(":"))
	if hxrt.StringEqualAny(pl.Get(0), hxrt.StringFromLiteral("")) {
		return hxrt.NewArray()
	}
	return pl
}

func (self *haxe__rtti__XmlParser) mkRights(r *string) *haxe__rtti__Rights {
	if hxrt.StringEqualStringPtr(r, hxrt.StringFromLiteral("null")) {
		return haxe__rtti__Rights_RNo
	}
	if hxrt.StringEqualStringPtr(r, hxrt.StringFromLiteral("method")) {
		return haxe__rtti__Rights_RMethod
	}
	if hxrt.StringEqualStringPtr(r, hxrt.StringFromLiteral("dynamic")) {
		return haxe__rtti__Rights_RDynamic
	}
	if hxrt.StringEqualStringPtr(r, hxrt.StringFromLiteral("inline")) {
		return haxe__rtti__Rights_RInline
	}
	return haxe__rtti__Rights_RCall(r)
}

func (self *haxe__rtti__XmlParser) xroot(x *Xml) {
	c := x.__hx_this.elements()
	for func(hx_obj_348 map[string]any) func() bool {
		hx_field_349 := hx_obj_348["hasNext"]
		if hx_field_349 == nil {
			var hx_zero_350 func() bool
			return hx_zero_350
		}
		return hx_field_349.(func() bool)
	}(c)() {
		c_1 := func(hx_obj_351 map[string]any) func() *Xml {
			hx_field_352 := hx_obj_351["next"]
			if hx_field_352 == nil {
				var hx_zero_353 func() *Xml
				return hx_zero_353
			}
			return hx_field_352.(func() *Xml)
		}(c)()
		self.__hx_this.merge(self.__hx_this.processElement(c_1))
	}
}

func (self *haxe__rtti__XmlParser) processElement(x *Xml) *haxe__rtti__TypeTree {
	var hx_if_354 *string
	if hxrt.HaxeEqual(x.nodeType, Xml_Document) {
		hx_if_354 = hxrt.StringFromLiteral("Document")
	} else {
		if !hxrt.HaxeEqual(x.nodeType, Xml_Element) {
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(x.nodeType))))
		}
		hx_if_354 = x.nodeName
	}
	nodeName := hx_if_354
	if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("class")) {
		return haxe__rtti__TypeTree_TClassdecl(self.__hx_this.xclass(x))
	}
	if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("enum")) {
		return haxe__rtti__TypeTree_TEnumdecl(self.__hx_this.xenum(x))
	}
	if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("typedef")) {
		return haxe__rtti__TypeTree_TTypedecl(self.__hx_this.xtypedef(x))
	}
	if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("abstract")) {
		return haxe__rtti__TypeTree_TAbstractdecl(self.__hx_this.xabstract(x))
	}
	hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Invalid "), nodeName))
	var hx_throw_zero_355 *haxe__rtti__TypeTree
	return hx_throw_zero_355
}

func (self *haxe__rtti__XmlParser) xmeta(x *Xml) *hxrt.Array {
	ml := hxrt.NewArray()
	m := x.__hx_this.elementsNamed(hxrt.StringFromLiteral("m"))
	for func(hx_obj_356 map[string]any) func() bool {
		hx_field_357 := hx_obj_356["hasNext"]
		if hx_field_357 == nil {
			var hx_zero_358 func() bool
			return hx_zero_358
		}
		return hx_field_357.(func() bool)
	}(m)() {
		m_1 := func(hx_obj_359 map[string]any) func() *Xml {
			hx_field_360 := hx_obj_359["next"]
			if hx_field_360 == nil {
				var hx_zero_361 func() *Xml
				return hx_zero_361
			}
			return hx_field_360.(func() *Xml)
		}(m)()
		pl := hxrt.NewArray()
		p := m_1.__hx_this.elementsNamed(hxrt.StringFromLiteral("e"))
		for func(hx_obj_362 map[string]any) func() bool {
			hx_field_363 := hx_obj_362["hasNext"]
			if hx_field_363 == nil {
				var hx_zero_364 func() bool
				return hx_zero_364
			}
			return hx_field_363.(func() bool)
		}(p)() {
			p_1 := func(hx_obj_365 map[string]any) func() *Xml {
				hx_field_366 := hx_obj_365["next"]
				if hx_field_366 == nil {
					var hx_zero_367 func() *Xml
					return hx_zero_367
				}
				return hx_field_366.(func() *Xml)
			}(p)()
			pl.Push(self.__hx_this.innerHTML(p_1))
		}
		hx_obj_370 := map[string]any{}
		hx_obj_370["name"] = self.__hx_this.requireAttr(m_1, hxrt.StringFromLiteral("n"))
		hx_obj_370["params"] = pl
		ml.Push(hx_obj_370)
	}
	return ml
}

func (self *haxe__rtti__XmlParser) xoverloads(x *Xml) *hxrt.Array {
	l := hxrt.NewArray()
	m := x.__hx_this.elements()
	for func(hx_obj_371 map[string]any) func() bool {
		hx_field_372 := hx_obj_371["hasNext"]
		if hx_field_372 == nil {
			var hx_zero_373 func() bool
			return hx_zero_373
		}
		return hx_field_372.(func() bool)
	}(m)() {
		m_1 := func(hx_obj_374 map[string]any) func() *Xml {
			hx_field_375 := hx_obj_374["next"]
			if hx_field_375 == nil {
				var hx_zero_376 func() *Xml
				return hx_zero_376
			}
			return hx_field_375.(func() *Xml)
		}(m)()
		l.Push(self.__hx_this.xclassfield(m_1, false))
	}
	return l
}

func (self *haxe__rtti__XmlParser) xpath(x *Xml) map[string]any {
	path := self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("path")))
	params := hxrt.NewArray()
	c := x.__hx_this.elements()
	for func(hx_obj_378 map[string]any) func() bool {
		hx_field_379 := hx_obj_378["hasNext"]
		if hx_field_379 == nil {
			var hx_zero_380 func() bool
			return hx_zero_380
		}
		return hx_field_379.(func() bool)
	}(c)() {
		c_1 := func(hx_obj_381 map[string]any) func() *Xml {
			hx_field_382 := hx_obj_381["next"]
			if hx_field_382 == nil {
				var hx_zero_383 func() *Xml
				return hx_zero_383
			}
			return hx_field_382.(func() *Xml)
		}(c)()
		params.Push(self.__hx_this.xtype(c_1))
	}
	hx_obj_385 := map[string]any{}
	hx_obj_385["path"] = path
	hx_obj_385["params"] = params
	return hx_obj_385
}

func (self *haxe__rtti__XmlParser) xclass(x *Xml) map[string]any {
	var csuper map[string]any = nil
	var doc *string = nil
	var tdynamic *haxe__rtti__CType = nil
	interfaces := hxrt.NewArray()
	fields := hxrt.NewArray()
	statics := hxrt.NewArray()
	meta := hxrt.NewArray()
	isInterface := x.__hx_this.exists(hxrt.StringFromLiteral("interface"))
	c := x.__hx_this.elements()
	for func(hx_obj_386 map[string]any) func() bool {
		hx_field_387 := hx_obj_386["hasNext"]
		if hx_field_387 == nil {
			var hx_zero_388 func() bool
			return hx_zero_388
		}
		return hx_field_387.(func() bool)
	}(c)() {
		c_1 := func(hx_obj_389 map[string]any) func() *Xml {
			hx_field_390 := hx_obj_389["next"]
			if hx_field_390 == nil {
				var hx_zero_391 func() *Xml
				return hx_zero_391
			}
			return hx_field_390.(func() *Xml)
		}(c)()
		nodeName := self.__hx_this.elementName(c_1)
		if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("haxe_doc")) {
			doc = self.__hx_this.innerData(c_1)
		} else {
			if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("extends")) {
				if isInterface {
					interfaces.Push(self.__hx_this.xpath(c_1))
				} else {
					csuper = self.__hx_this.xpath(c_1)
				}
			} else {
				if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("implements")) {
					interfaces.Push(self.__hx_this.xpath(c_1))
				} else {
					if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("haxe_dynamic")) {
						tdynamic = self.__hx_this.xtype(self.__hx_this.requireFirstElement(c_1))
					} else {
						if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("meta")) {
							meta = self.__hx_this.xmeta(c_1)
						} else {
							if c_1.__hx_this.exists(hxrt.StringFromLiteral("static")) {
								statics.Push(self.__hx_this.xclassfield(c_1, false))
							} else {
								fields.Push(self.__hx_this.xclassfield(c_1, false))
							}
						}
					}
				}
			}
		}
	}
	hx_obj_396 := map[string]any{}
	hx_obj_396["file"] = x.__hx_this.get(hxrt.StringFromLiteral("file"))
	hx_obj_396["path"] = self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("path")))
	var hx_if_397 *string
	if x.__hx_this.exists(hxrt.StringFromLiteral("module")) {
		hx_if_397 = self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("module")))
	} else {
		hx_if_397 = nil
	}
	hx_obj_396["module"] = hx_if_397
	hx_obj_396["doc"] = doc
	hx_obj_396["isPrivate"] = x.__hx_this.exists(hxrt.StringFromLiteral("private"))
	hx_obj_396["isExtern"] = x.__hx_this.exists(hxrt.StringFromLiteral("extern"))
	hx_obj_396["isFinal"] = x.__hx_this.exists(hxrt.StringFromLiteral("final"))
	hx_obj_396["isInterface"] = isInterface
	hx_obj_396["params"] = self.__hx_this.mkTypeParams(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("params")))
	hx_obj_396["superClass"] = csuper
	hx_obj_396["interfaces"] = interfaces
	hx_obj_396["fields"] = fields
	hx_obj_396["statics"] = statics
	hx_obj_396["tdynamic"] = tdynamic
	hx_obj_396["platforms"] = self.__hx_this.defplat()
	hx_obj_396["meta"] = meta
	return hx_obj_396
}

func (self *haxe__rtti__XmlParser) xclassfield(x *Xml, defPublic bool) map[string]any {
	e := x.__hx_this.elements()
	t := self.__hx_this.xtype(func(hx_obj_398 map[string]any) func() *Xml {
		hx_field_399 := hx_obj_398["next"]
		if hx_field_399 == nil {
			var hx_zero_400 func() *Xml
			return hx_zero_400
		}
		return hx_field_399.(func() *Xml)
	}(e)())
	var doc *string = nil
	meta := hxrt.NewArray()
	var overloads *hxrt.Array = nil
	var line any = nil
	if x.__hx_this.exists(hxrt.StringFromLiteral("line")) {
		line = self.__hx_this.parseIntString(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("line")))
	}
	c := e
	for func(hx_obj_401 map[string]any) func() bool {
		hx_field_402 := hx_obj_401["hasNext"]
		if hx_field_402 == nil {
			var hx_zero_403 func() bool
			return hx_zero_403
		}
		return hx_field_402.(func() bool)
	}(c)() {
		c_1 := func(hx_obj_404 map[string]any) func() *Xml {
			hx_field_405 := hx_obj_404["next"]
			if hx_field_405 == nil {
				var hx_zero_406 func() *Xml
				return hx_zero_406
			}
			return hx_field_405.(func() *Xml)
		}(c)()
		nodeName := self.__hx_this.elementName(c_1)
		if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("haxe_doc")) {
			doc = self.__hx_this.innerData(c_1)
		} else {
			if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("meta")) {
				meta = self.__hx_this.xmeta(c_1)
			} else {
				if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("overloads")) {
					overloads = self.__hx_this.xoverloads(c_1)
				} else {
					hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Invalid "), nodeName))
				}
			}
		}
	}
	hx_obj_407 := map[string]any{}
	hx_obj_407["name"] = self.__hx_this.elementName(x)
	hx_obj_407["type"] = t
	hx_obj_407["isPublic"] = (x.__hx_this.exists(hxrt.StringFromLiteral("public")) || func(hx_value_408 any) bool {
		if hx_value_408 == nil {
			var hx_zero_409 bool
			return hx_zero_409
		}
		return hx_value_408.(bool)
	}(defPublic))
	hx_obj_407["isFinal"] = x.__hx_this.exists(hxrt.StringFromLiteral("final"))
	hx_obj_407["isOverride"] = x.__hx_this.exists(hxrt.StringFromLiteral("override"))
	hx_obj_407["line"] = func(hx_value_410 any) any {
		if hx_value_410 == nil {
			return nil
		}
		return hx_value_410.(int)
	}(line)
	hx_obj_407["doc"] = doc
	var hx_if_411 *haxe__rtti__Rights
	if x.__hx_this.exists(hxrt.StringFromLiteral("get")) {
		hx_if_411 = self.__hx_this.mkRights(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("get")))
	} else {
		hx_if_411 = haxe__rtti__Rights_RNormal
	}
	hx_obj_407["get"] = hx_if_411
	var hx_if_412 *haxe__rtti__Rights
	if x.__hx_this.exists(hxrt.StringFromLiteral("set")) {
		hx_if_412 = self.__hx_this.mkRights(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("set")))
	} else {
		hx_if_412 = haxe__rtti__Rights_RNormal
	}
	hx_obj_407["set"] = hx_if_412
	var hx_if_413 *hxrt.Array
	if x.__hx_this.exists(hxrt.StringFromLiteral("params")) {
		hx_if_413 = self.__hx_this.mkTypeParams(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("params")))
	} else {
		hx_if_413 = hxrt.NewArray()
	}
	hx_obj_407["params"] = hx_if_413
	hx_obj_407["platforms"] = self.__hx_this.defplat()
	hx_obj_407["meta"] = meta
	hx_obj_407["overloads"] = overloads
	var hx_if_414 *string
	if x.__hx_this.exists(hxrt.StringFromLiteral("expr")) {
		hx_if_414 = self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("expr"))
	} else {
		hx_if_414 = nil
	}
	hx_obj_407["expr"] = hx_if_414
	return hx_obj_407
}

func (self *haxe__rtti__XmlParser) xenum(x *Xml) map[string]any {
	cl := hxrt.NewArray()
	var doc *string = nil
	meta := hxrt.NewArray()
	c := x.__hx_this.elements()
	for func(hx_obj_415 map[string]any) func() bool {
		hx_field_416 := hx_obj_415["hasNext"]
		if hx_field_416 == nil {
			var hx_zero_417 func() bool
			return hx_zero_417
		}
		return hx_field_416.(func() bool)
	}(c)() {
		c_1 := func(hx_obj_418 map[string]any) func() *Xml {
			hx_field_419 := hx_obj_418["next"]
			if hx_field_419 == nil {
				var hx_zero_420 func() *Xml
				return hx_zero_420
			}
			return hx_field_419.(func() *Xml)
		}(c)()
		if hxrt.StringEqualStringPtr(self.__hx_this.elementName(c_1), hxrt.StringFromLiteral("haxe_doc")) {
			doc = self.__hx_this.innerData(c_1)
		} else {
			if hxrt.StringEqualStringPtr(self.__hx_this.elementName(c_1), hxrt.StringFromLiteral("meta")) {
				meta = self.__hx_this.xmeta(c_1)
			} else {
				cl.Push(self.__hx_this.xenumfield(c_1))
			}
		}
	}
	hx_obj_422 := map[string]any{}
	hx_obj_422["file"] = x.__hx_this.get(hxrt.StringFromLiteral("file"))
	hx_obj_422["path"] = self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("path")))
	var hx_if_423 *string
	if x.__hx_this.exists(hxrt.StringFromLiteral("module")) {
		hx_if_423 = self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("module")))
	} else {
		hx_if_423 = nil
	}
	hx_obj_422["module"] = hx_if_423
	hx_obj_422["doc"] = doc
	hx_obj_422["isPrivate"] = x.__hx_this.exists(hxrt.StringFromLiteral("private"))
	hx_obj_422["isExtern"] = x.__hx_this.exists(hxrt.StringFromLiteral("extern"))
	hx_obj_422["params"] = self.__hx_this.mkTypeParams(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("params")))
	hx_obj_422["constructors"] = cl
	hx_obj_422["platforms"] = self.__hx_this.defplat()
	hx_obj_422["meta"] = meta
	return hx_obj_422
}

func (self *haxe__rtti__XmlParser) xenumfield(x *Xml) map[string]any {
	var args *hxrt.Array = nil
	docElements := x.__hx_this.elementsNamed(hxrt.StringFromLiteral("haxe_doc"))
	var hx_if_430 *Xml
	if func(hx_obj_424 map[string]any) func() bool {
		hx_field_425 := hx_obj_424["hasNext"]
		if hx_field_425 == nil {
			var hx_zero_426 func() bool
			return hx_zero_426
		}
		return hx_field_425.(func() bool)
	}(docElements)() {
		hx_if_430 = func(hx_obj_427 map[string]any) func() *Xml {
			hx_field_428 := hx_obj_427["next"]
			if hx_field_428 == nil {
				var hx_zero_429 func() *Xml
				return hx_zero_429
			}
			return hx_field_428.(func() *Xml)
		}(docElements)()
	} else {
		hx_if_430 = nil
	}
	xdoc := hx_if_430
	var hx_if_431 *hxrt.Array
	if self.__hx_this.hasNamedElement(x, hxrt.StringFromLiteral("meta")) {
		hx_if_431 = self.__hx_this.xmeta(self.__hx_this.requireNamedElement(x, hxrt.StringFromLiteral("meta")))
	} else {
		hx_if_431 = hxrt.NewArray()
	}
	meta := hx_if_431
	if x.__hx_this.exists(hxrt.StringFromLiteral("a")) {
		names := self.__hx_this.splitString(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("a")), hxrt.StringFromLiteral(":"))
		elts := x.__hx_this.elements()
		args = hxrt.NewArray()
		_g := 0
		for _g < names.Len() {
			c := func(hx_value_432 any) *string {
				if hx_value_432 == nil {
					var hx_zero_433 *string
					return hx_zero_433
				}
				return hx_value_432.(*string)
			}(names.Get(_g))
			_g = int(int32((_g + 1)))
			opt := false
			if hxrt.StringEqualStringPtr(hxrt.StringCharAtStringPtr(c, 0), hxrt.StringFromLiteral("?")) {
				opt = true
				c = hxrt.StringSubstrStringPtr(c, 1, 0, false)
			}
			hx_obj_435 := map[string]any{}
			hx_obj_435["name"] = c
			hx_obj_435["opt"] = opt
			hx_obj_435["t"] = self.__hx_this.xtype(func(hx_obj_436 map[string]any) func() *Xml {
				hx_field_437 := hx_obj_436["next"]
				if hx_field_437 == nil {
					var hx_zero_438 func() *Xml
					return hx_zero_438
				}
				return hx_field_437.(func() *Xml)
			}(elts)())
			args.Push(hx_obj_435)
		}
	}
	hx_obj_439 := map[string]any{}
	hx_obj_439["name"] = self.__hx_this.elementName(x)
	hx_obj_439["args"] = args
	var hx_if_440 *string
	if xdoc == nil {
		hx_if_440 = nil
	} else {
		hx_if_440 = self.__hx_this.innerData(xdoc)
	}
	hx_obj_439["doc"] = hx_if_440
	hx_obj_439["meta"] = meta
	hx_obj_439["platforms"] = self.__hx_this.defplat()
	return hx_obj_439
}

func (self *haxe__rtti__XmlParser) xabstract(x *Xml) map[string]any {
	var doc *string = nil
	var impl map[string]any = nil
	var athis *haxe__rtti__CType = nil
	meta := hxrt.NewArray()
	to := hxrt.NewArray()
	from := hxrt.NewArray()
	c := x.__hx_this.elements()
	for func(hx_obj_441 map[string]any) func() bool {
		hx_field_442 := hx_obj_441["hasNext"]
		if hx_field_442 == nil {
			var hx_zero_443 func() bool
			return hx_zero_443
		}
		return hx_field_442.(func() bool)
	}(c)() {
		c_1 := func(hx_obj_444 map[string]any) func() *Xml {
			hx_field_445 := hx_obj_444["next"]
			if hx_field_445 == nil {
				var hx_zero_446 func() *Xml
				return hx_zero_446
			}
			return hx_field_445.(func() *Xml)
		}(c)()
		nodeName := self.__hx_this.elementName(c_1)
		if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("haxe_doc")) {
			doc = self.__hx_this.innerData(c_1)
		} else {
			if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("meta")) {
				meta = self.__hx_this.xmeta(c_1)
			} else {
				if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("to")) {
					t := c_1.__hx_this.elements()
					for func(hx_obj_447 map[string]any) func() bool {
						hx_field_448 := hx_obj_447["hasNext"]
						if hx_field_448 == nil {
							var hx_zero_449 func() bool
							return hx_zero_449
						}
						return hx_field_448.(func() bool)
					}(t)() {
						t_1 := func(hx_obj_450 map[string]any) func() *Xml {
							hx_field_451 := hx_obj_450["next"]
							if hx_field_451 == nil {
								var hx_zero_452 func() *Xml
								return hx_zero_452
							}
							return hx_field_451.(func() *Xml)
						}(t)()
						hx_obj_454 := map[string]any{}
						hx_obj_454["t"] = self.__hx_this.xtype(self.__hx_this.requireFirstElement(t_1))
						hx_obj_454["field"] = t_1.__hx_this.get(hxrt.StringFromLiteral("field"))
						to.Push(hx_obj_454)
					}
				} else {
					if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("from")) {
						t_2 := c_1.__hx_this.elements()
						for func(hx_obj_455 map[string]any) func() bool {
							hx_field_456 := hx_obj_455["hasNext"]
							if hx_field_456 == nil {
								var hx_zero_457 func() bool
								return hx_zero_457
							}
							return hx_field_456.(func() bool)
						}(t_2)() {
							t_3 := func(hx_obj_458 map[string]any) func() *Xml {
								hx_field_459 := hx_obj_458["next"]
								if hx_field_459 == nil {
									var hx_zero_460 func() *Xml
									return hx_zero_460
								}
								return hx_field_459.(func() *Xml)
							}(t_2)()
							hx_obj_462 := map[string]any{}
							hx_obj_462["t"] = self.__hx_this.xtype(self.__hx_this.requireFirstElement(t_3))
							hx_obj_462["field"] = t_3.__hx_this.get(hxrt.StringFromLiteral("field"))
							from.Push(hx_obj_462)
						}
					} else {
						if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("impl")) {
							impl = self.__hx_this.xclass(self.__hx_this.requireNamedElement(c_1, hxrt.StringFromLiteral("class")))
						} else {
							if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("this")) {
								athis = self.__hx_this.xtype(self.__hx_this.requireFirstElement(c_1))
							} else {
								hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Invalid "), nodeName))
							}
						}
					}
				}
			}
		}
	}
	hx_obj_463 := map[string]any{}
	hx_obj_463["file"] = x.__hx_this.get(hxrt.StringFromLiteral("file"))
	hx_obj_463["path"] = self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("path")))
	var hx_if_464 *string
	if x.__hx_this.exists(hxrt.StringFromLiteral("module")) {
		hx_if_464 = self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("module")))
	} else {
		hx_if_464 = nil
	}
	hx_obj_463["module"] = hx_if_464
	hx_obj_463["doc"] = doc
	hx_obj_463["isPrivate"] = x.__hx_this.exists(hxrt.StringFromLiteral("private"))
	hx_obj_463["params"] = self.__hx_this.mkTypeParams(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("params")))
	hx_obj_463["platforms"] = self.__hx_this.defplat()
	hx_obj_463["meta"] = meta
	hx_obj_463["athis"] = athis
	hx_obj_463["to"] = to
	hx_obj_463["from"] = from
	hx_obj_463["impl"] = impl
	return hx_obj_463
}

func (self *haxe__rtti__XmlParser) xtypedef(x *Xml) map[string]any {
	var doc *string = nil
	var t *haxe__rtti__CType = nil
	meta := hxrt.NewArray()
	c := x.__hx_this.elements()
	for func(hx_obj_465 map[string]any) func() bool {
		hx_field_466 := hx_obj_465["hasNext"]
		if hx_field_466 == nil {
			var hx_zero_467 func() bool
			return hx_zero_467
		}
		return hx_field_466.(func() bool)
	}(c)() {
		c_1 := func(hx_obj_468 map[string]any) func() *Xml {
			hx_field_469 := hx_obj_468["next"]
			if hx_field_469 == nil {
				var hx_zero_470 func() *Xml
				return hx_zero_470
			}
			return hx_field_469.(func() *Xml)
		}(c)()
		if hxrt.StringEqualStringPtr(self.__hx_this.elementName(c_1), hxrt.StringFromLiteral("haxe_doc")) {
			doc = self.__hx_this.innerData(c_1)
		} else {
			if hxrt.StringEqualStringPtr(self.__hx_this.elementName(c_1), hxrt.StringFromLiteral("meta")) {
				meta = self.__hx_this.xmeta(c_1)
			} else {
				t = self.__hx_this.xtype(c_1)
			}
		}
	}
	types := New_haxe__ds__StringMap()
	if !hxrt.StringEqualStringPtr(self.curplatform, nil) {
		types.__hx_this.set(self.curplatform, t)
	}
	hx_obj_471 := map[string]any{}
	hx_obj_471["file"] = x.__hx_this.get(hxrt.StringFromLiteral("file"))
	hx_obj_471["path"] = self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("path")))
	var hx_if_472 *string
	if x.__hx_this.exists(hxrt.StringFromLiteral("module")) {
		hx_if_472 = self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("module")))
	} else {
		hx_if_472 = nil
	}
	hx_obj_471["module"] = hx_if_472
	hx_obj_471["doc"] = doc
	hx_obj_471["isPrivate"] = x.__hx_this.exists(hxrt.StringFromLiteral("private"))
	hx_obj_471["params"] = self.__hx_this.mkTypeParams(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("params")))
	hx_obj_471["type"] = t
	hx_obj_471["types"] = types
	hx_obj_471["platforms"] = self.__hx_this.defplat()
	hx_obj_471["meta"] = meta
	return hx_obj_471
}

func (self *haxe__rtti__XmlParser) xtype(x *Xml) *haxe__rtti__CType {
	nodeName := self.__hx_this.elementName(x)
	if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("unknown")) {
		return haxe__rtti__CType_CUnknown
	}
	if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("e")) {
		return haxe__rtti__CType_CEnum(self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("path"))), self.__hx_this.xtypeparams(x))
	}
	if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("c")) {
		return haxe__rtti__CType_CClass(self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("path"))), self.__hx_this.xtypeparams(x))
	}
	if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("t")) {
		return haxe__rtti__CType_CTypedef(self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("path"))), self.__hx_this.xtypeparams(x))
	}
	if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("x")) {
		return haxe__rtti__CType_CAbstract(self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("path"))), self.__hx_this.xtypeparams(x))
	}
	if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("f")) {
		args := hxrt.NewArray()
		aname := self.__hx_this.splitString(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("a")), hxrt.StringFromLiteral(":"))
		argIndex := 0
		var hx_if_473 *hxrt.Array
		if x.__hx_this.exists(hxrt.StringFromLiteral("v")) {
			hx_if_473 = self.__hx_this.splitString(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("v")), hxrt.StringFromLiteral(":"))
		} else {
			hx_if_473 = nil
		}
		evalues := hx_if_473
		valueIndex := 0
		e := x.__hx_this.elements()
		for func(hx_obj_474 map[string]any) func() bool {
			hx_field_475 := hx_obj_474["hasNext"]
			if hx_field_475 == nil {
				var hx_zero_476 func() bool
				return hx_zero_476
			}
			return hx_field_475.(func() bool)
		}(e)() {
			e_1 := func(hx_obj_477 map[string]any) func() *Xml {
				hx_field_478 := hx_obj_477["next"]
				if hx_field_478 == nil {
					var hx_zero_479 func() *Xml
					return hx_zero_479
				}
				return hx_field_478.(func() *Xml)
			}(e)()
			opt := false
			var hx_if_482 *string
			if argIndex < aname.Len() {
				hx_if_482 = hxrt.StdString(func(hx_value_480 any) *string {
					if hx_value_480 == nil {
						var hx_zero_481 *string
						return hx_zero_481
					}
					return hx_value_480.(*string)
				}(aname.Get(argIndex)))
			} else {
				hx_if_482 = nil
			}
			a := hx_if_482
			argIndex = int(int32((argIndex + 1)))
			if hxrt.StringEqualStringPtr(a, nil) {
				a = hxrt.StringFromLiteral("")
			}
			if hxrt.StringEqualStringPtr(hxrt.StringCharAtStringPtr(a, 0), hxrt.StringFromLiteral("?")) {
				opt = true
				a = hxrt.StringSubstrStringPtr(a, 1, 0, false)
			}
			var hx_if_486 *string
			if (evalues == nil) || (valueIndex >= evalues.Len()) {
				hx_if_486 = nil
			} else {
				hx_post_483 := valueIndex
				valueIndex = int(int32((valueIndex + 1)))
				hx_if_486 = hxrt.StdString(func(hx_value_484 any) *string {
					if hx_value_484 == nil {
						var hx_zero_485 *string
						return hx_zero_485
					}
					return hx_value_484.(*string)
				}(evalues.Get(hx_post_483)))
			}
			v := hx_if_486
			hx_obj_488 := map[string]any{}
			hx_obj_488["name"] = a
			hx_obj_488["opt"] = opt
			hx_obj_488["t"] = self.__hx_this.xtype(e_1)
			var hx_if_489 *string
			if hxrt.StringEqualStringPtr(v, hxrt.StringFromLiteral("")) {
				hx_if_489 = nil
			} else {
				hx_if_489 = v
			}
			hx_obj_488["value"] = hx_if_489
			args.Push(hx_obj_488)
		}
		ret := func(hx_value_490 any) map[string]any {
			if hx_value_490 == nil {
				var hx_zero_491 map[string]any
				return hx_zero_491
			}
			return hx_value_490.(map[string]any)
		}(args.Get(int(int32((hxrt.Int32Wrap(args.Len()) - hxrt.Int32Wrap(1))))))
		callArgs := hxrt.NewArray()
		_g := 0
		_g1 := int(int32((hxrt.Int32Wrap(args.Len()) - hxrt.Int32Wrap(1))))
		for _g < _g1 {
			hx_post_492 := _g
			_g = int(int32((_g + 1)))
			i := hx_post_492
			callArgs.Push(args.Get(i))
		}
		return haxe__rtti__CType_CFunction(callArgs, func(hx_obj_494 map[string]any) *haxe__rtti__CType {
			hx_field_495 := hx_obj_494["t"]
			if hx_field_495 == nil {
				var hx_zero_496 *haxe__rtti__CType
				return hx_zero_496
			}
			return hx_field_495.(*haxe__rtti__CType)
		}(ret))
	}
	if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("a")) {
		fields := hxrt.NewArray()
		f := x.__hx_this.elements()
		for func(hx_obj_497 map[string]any) func() bool {
			hx_field_498 := hx_obj_497["hasNext"]
			if hx_field_498 == nil {
				var hx_zero_499 func() bool
				return hx_zero_499
			}
			return hx_field_498.(func() bool)
		}(f)() {
			f_1 := func(hx_obj_500 map[string]any) func() *Xml {
				hx_field_501 := hx_obj_500["next"]
				if hx_field_501 == nil {
					var hx_zero_502 func() *Xml
					return hx_zero_502
				}
				return hx_field_501.(func() *Xml)
			}(f)()
			f_2 := self.__hx_this.xclassfield(f_1, true)
			f_2["platforms"] = hxrt.NewArray()
			fields.Push(f_2)
		}
		return haxe__rtti__CType_CAnonymous(fields)
	}
	if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("d")) {
		var t *haxe__rtti__CType = nil
		tx := x.__hx_this.firstElement()
		if tx != nil {
			t = self.__hx_this.xtype(tx)
		}
		return haxe__rtti__CType_CDynamic(t)
	}
	hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Invalid "), nodeName))
	var hx_throw_zero_504 *haxe__rtti__CType
	return hx_throw_zero_504
}

func (self *haxe__rtti__XmlParser) xtypeparams(x *Xml) *hxrt.Array {
	p := hxrt.NewArray()
	c := x.__hx_this.elements()
	for func(hx_obj_505 map[string]any) func() bool {
		hx_field_506 := hx_obj_505["hasNext"]
		if hx_field_506 == nil {
			var hx_zero_507 func() bool
			return hx_zero_507
		}
		return hx_field_506.(func() bool)
	}(c)() {
		c_1 := func(hx_obj_508 map[string]any) func() *Xml {
			hx_field_509 := hx_obj_508["next"]
			if hx_field_509 == nil {
				var hx_zero_510 func() *Xml
				return hx_zero_510
			}
			return hx_field_509.(func() *Xml)
		}(c)()
		p.Push(self.__hx_this.xtype(c_1))
	}
	return p
}

func (self *haxe__rtti__XmlParser) defplat() *hxrt.Array {
	l := hxrt.NewArray()
	if !hxrt.StringEqualStringPtr(self.curplatform, nil) {
		l.Push(self.curplatform)
	}
	return l
}

func (self *haxe__rtti__XmlParser) joinStringArray(values *hxrt.Array, separator *string) *string {
	var buf_b *string
	buf_b = hxrt.StringFromLiteral("")
	_g := 0
	_g1 := values.Len()
	for _g < _g1 {
		hx_post_513 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_513
		if i > 0 {
			buf_b = hxrt.StringConcatStringPtr(buf_b, hxrt.StdString(separator))
		}
		buf_b = hxrt.StringConcatStringPtr(buf_b, hxrt.StdString(values.Get(i)))
	}
	return buf_b
}

func (self *haxe__rtti__XmlParser) splitString(value *string, separator *string) *hxrt.Array {
	if hxrt.StringEqualStringPtr(separator, hxrt.StringFromLiteral("")) {
		return hxrt.NewArray(value)
	}
	parts := hxrt.NewArray()
	start := 0
	for true {
		index := self.__hx_this.findSeparator(value, separator, start)
		if index == -1 {
			parts.Push(hxrt.StringSubstrStringPtr(value, start, 0, false))
			break
		}
		parts.Push(hxrt.StringSubstrStringPtr(value, start, int(int32((hxrt.Int32Wrap(index) - hxrt.Int32Wrap(start)))), true))
		start = int(int32((hxrt.Int32Wrap(index) + hxrt.Int32Wrap(hxrt.StringLengthStringPtr(separator)))))
	}
	return parts
}

func (self *haxe__rtti__XmlParser) findSeparator(value *string, separator *string, start int) int {
	limit := int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(value)) - hxrt.Int32Wrap(hxrt.StringLengthStringPtr(separator)))))
	index := start
	for index <= limit {
		if hxrt.StringEqualStringPtr(hxrt.StringSubstrStringPtr(value, index, hxrt.StringLengthStringPtr(separator), true), separator) {
			return index
		}
		index = int(int32((index + 1)))
	}
	return -1
}

func (self *haxe__rtti__XmlParser) requireAttr(x *Xml, name *string) *string {
	value := x.__hx_this.get(name)
	var hx_if_516 *string
	if hxrt.StringEqualStringPtr(value, nil) {
		hx_if_516 = hxrt.StringFromLiteral("")
	} else {
		hx_if_516 = value
	}
	return hx_if_516
}

func (self *haxe__rtti__XmlParser) hasNamedElement(x *Xml, name *string) bool {
	return func(hx_obj_517 map[string]any) func() bool {
		hx_field_518 := hx_obj_517["hasNext"]
		if hx_field_518 == nil {
			var hx_zero_519 func() bool
			return hx_zero_519
		}
		return hx_field_518.(func() bool)
	}(x.__hx_this.elementsNamed(name))()
}

func (self *haxe__rtti__XmlParser) requireNamedElement(x *Xml, name *string) *Xml {
	elements := x.__hx_this.elementsNamed(name)
	if !func(hx_obj_520 map[string]any) func() bool {
		hx_field_521 := hx_obj_520["hasNext"]
		if hx_field_521 == nil {
			var hx_zero_522 func() bool
			return hx_zero_522
		}
		return hx_field_521.(func() bool)
	}(elements)() {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(self.__hx_this.nodeDisplayName(x), hxrt.StringFromLiteral(" is missing element ")), name))
	}
	return func(hx_obj_523 map[string]any) func() *Xml {
		hx_field_524 := hx_obj_523["next"]
		if hx_field_524 == nil {
			var hx_zero_525 func() *Xml
			return hx_zero_525
		}
		return hx_field_524.(func() *Xml)
	}(elements)()
}

func (self *haxe__rtti__XmlParser) requireFirstElement(x *Xml) *Xml {
	first := x.__hx_this.firstElement()
	if first == nil {
		hxrt.Throw(hxrt.StringConcatStringPtr(self.__hx_this.nodeDisplayName(x), hxrt.StringFromLiteral(" is missing first element")))
	}
	return first
}

func (self *haxe__rtti__XmlParser) nodeDisplayName(x *Xml) *string {
	var hx_if_526 *string
	if hxrt.HaxeEqual(x.nodeType, Xml_Document) {
		hx_if_526 = hxrt.StringFromLiteral("Document")
	} else {
		hx_if_526 = self.__hx_this.elementName(x)
	}
	return hx_if_526
}

func (self *haxe__rtti__XmlParser) elementName(x *Xml) *string {
	if !hxrt.HaxeEqual(x.nodeType, Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(x.nodeType))))
	}
	name := x.nodeName
	var hx_if_527 *string
	if hxrt.StringEqualStringPtr(name, nil) {
		hx_if_527 = hxrt.StringFromLiteral("")
	} else {
		hx_if_527 = name
	}
	return hx_if_527
}

func (self *haxe__rtti__XmlParser) innerData(x *Xml) *string {
	var it_current int
	var it_array *hxrt.Array
	x.__hx_this.ensureElementType()
	_this := x.children
	it_current = 0
	it_array = _this
	if !(it_current < it_array.Len()) {
		hxrt.Throw(hxrt.StringConcatStringPtr(self.__hx_this.nodeDisplayName(x), hxrt.StringFromLiteral(" does not have data")))
	}
	value := func(hx_value_529 any) *Xml {
		if hx_value_529 == nil {
			var hx_zero_530 *Xml
			return hx_zero_530
		}
		return hx_value_529.(*Xml)
	}(it_array.Get(func() int {
		hx_post_528 := it_current
		it_current = int(int32((it_current + 1)))
		return hx_post_528
	}()))
	if it_current < it_array.Len() {
		hxrt.Throw(hxrt.StringConcatStringPtr(self.__hx_this.nodeDisplayName(x), hxrt.StringFromLiteral(" does not only have data")))
	}
	if !hxrt.HaxeEqual(value.nodeType, Xml_PCData) && !hxrt.HaxeEqual(value.nodeType, Xml_CData) {
		hxrt.Throw(hxrt.StringConcatStringPtr(self.__hx_this.nodeDisplayName(x), hxrt.StringFromLiteral(" does not have data")))
	}
	if hxrt.HaxeEqual(value.nodeType, Xml_Document) || hxrt.HaxeEqual(value.nodeType, Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, unexpected "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(value.nodeType))))
	}
	return value.nodeValue
}

func (self *haxe__rtti__XmlParser) innerHTML(x *Xml) *string {
	var buf_b *string
	buf_b = hxrt.StringFromLiteral("")
	var _g_current int
	var _g_array *hxrt.Array
	x.__hx_this.ensureElementType()
	_this := x.children
	_g_current = 0
	_g_array = _this
	for _g_current < _g_array.Len() {
		hx_post_531 := _g_current
		_g_current = int(int32((_g_current + 1)))
		child := func(hx_value_532 any) *Xml {
			if hx_value_532 == nil {
				var hx_zero_533 *Xml
				return hx_zero_533
			}
			return hx_value_532.(*Xml)
		}(_g_array.Get(hx_post_531))
		x_1 := child.__hx_this.toString()
		buf_b = hxrt.StringConcatStringPtr(buf_b, hxrt.StdString(x_1))
	}
	return buf_b
}

func (self *haxe__rtti__XmlParser) parseIntString(value *string) int {
	if hxrt.StringEqualStringPtr(value, nil) || hxrt.StringEqualStringPtr(value, hxrt.StringFromLiteral("")) {
		return 0
	}
	zeroCode := 48
	nineCode := 57
	negative := false
	index := 0
	if hxrt.StringEqualStringPtr(hxrt.StringCharAtStringPtr(value, 0), hxrt.StringFromLiteral("-")) {
		negative = true
		index = 1
	}
	result := 0
	for index < hxrt.StringLengthStringPtr(value) {
		code := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(value, index))
		if (code < zeroCode) || (code > nineCode) {
			return 0
		}
		result = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(result) * hxrt.Int32Wrap(10))))) + hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(code) - hxrt.Int32Wrap(zeroCode))))))))
		index = int(int32((index + 1)))
	}
	var hx_if_534 int
	if negative {
		hx_if_534 = int(int32(-int32(result)))
	} else {
		hx_if_534 = result
	}
	return hx_if_534
}
