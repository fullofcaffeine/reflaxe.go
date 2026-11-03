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
	haxe__ds__ArraySort_sort(l, func(hx_cmp_left_21 any, hx_cmp_right_22 any) int {
		return func(e1 *haxe__rtti__TypeTree, e2 *haxe__rtti__TypeTree) int {
			var hx_if_10 *string
			if e1.tag == 0 {
				_g := func(hx_value_1 any) *string {
					if hx_value_1 == nil {
						var hx_zero_2 *string
						return hx_zero_2
					}
					return hx_value_1.(*string)
				}(e1.params[0])
				_g1 := func(hx_value_3 any) *string {
					if hx_value_3 == nil {
						var hx_zero_4 *string
						return hx_zero_4
					}
					return hx_value_3.(*string)
				}(e1.params[1])
				_ = _g1
				_g1_1 := func(hx_value_5 any) *hxrt.Array {
					if hx_value_5 == nil {
						var hx_zero_6 *hxrt.Array
						return hx_zero_6
					}
					return hx_value_5.(*hxrt.Array)
				}(e1.params[2])
				_ = _g1_1
				p := _g
				hx_if_10 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(" "), p)
			} else {
				hx_if_10 = func(hx_obj_7 map[string]any) *string {
					hx_field_8 := hx_obj_7["path"]
					if hx_field_8 == nil {
						var hx_zero_9 *string
						return hx_zero_9
					}
					return hx_field_8.(*string)
				}(haxe__rtti__TypeApi_typeInfos(e1))
			}
			n1 := hx_if_10
			var hx_if_20 *string
			if e2.tag == 0 {
				_g_1 := func(hx_value_11 any) *string {
					if hx_value_11 == nil {
						var hx_zero_12 *string
						return hx_zero_12
					}
					return hx_value_11.(*string)
				}(e2.params[0])
				_g1_2 := func(hx_value_13 any) *string {
					if hx_value_13 == nil {
						var hx_zero_14 *string
						return hx_zero_14
					}
					return hx_value_13.(*string)
				}(e2.params[1])
				_ = _g1_2
				_g1_3 := func(hx_value_15 any) *hxrt.Array {
					if hx_value_15 == nil {
						var hx_zero_16 *hxrt.Array
						return hx_zero_16
					}
					return hx_value_15.(*hxrt.Array)
				}(e2.params[2])
				_ = _g1_3
				p_1 := _g_1
				hx_if_20 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(" "), p_1)
			} else {
				hx_if_20 = func(hx_obj_17 map[string]any) *string {
					hx_field_18 := hx_obj_17["path"]
					if hx_field_18 == nil {
						var hx_zero_19 *string
						return hx_zero_19
					}
					return hx_field_18.(*string)
				}(haxe__rtti__TypeApi_typeInfos(e2))
			}
			n2 := hx_if_20
			return Reflect_compare(n1, n2)
		}(func(hx_value_23 any) *haxe__rtti__TypeTree {
			if hx_value_23 == nil {
				var hx_zero_24 *haxe__rtti__TypeTree
				return hx_zero_24
			}
			return hx_value_23.(*haxe__rtti__TypeTree)
		}(hx_cmp_left_21), func(hx_value_25 any) *haxe__rtti__TypeTree {
			if hx_value_25 == nil {
				var hx_zero_26 *haxe__rtti__TypeTree
				return hx_zero_26
			}
			return hx_value_25.(*haxe__rtti__TypeTree)
		}(hx_cmp_right_22))
	})
	_g := 0
	for _g < l.Len() {
		x := func(hx_value_27 any) *haxe__rtti__TypeTree {
			if hx_value_27 == nil {
				var hx_zero_28 *haxe__rtti__TypeTree
				return hx_zero_28
			}
			return hx_value_27.(*haxe__rtti__TypeTree)
		}(l.Get(_g))
		_g = int(int32((_g + 1)))
		switch x.tag {
		case 0:
			_g_1 := func(hx_value_29 any) *string {
				if hx_value_29 == nil {
					var hx_zero_30 *string
					return hx_zero_30
				}
				return hx_value_29.(*string)
			}(x.params[0])
			_ = _g_1
			_g_2 := func(hx_value_31 any) *string {
				if hx_value_31 == nil {
					var hx_zero_32 *string
					return hx_zero_32
				}
				return hx_value_31.(*string)
			}(x.params[1])
			_ = _g_2
			_g_3 := func(hx_value_33 any) *hxrt.Array {
				if hx_value_33 == nil {
					var hx_zero_34 *hxrt.Array
					return hx_zero_34
				}
				return hx_value_33.(*hxrt.Array)
			}(x.params[2])
			l_1 := _g_3
			self.__hx_this.sort(l_1)
		case 1:
			_g_4 := func(hx_value_35 any) map[string]any {
				if hx_value_35 == nil {
					var hx_zero_36 map[string]any
					return hx_zero_36
				}
				return hx_value_35.(map[string]any)
			}(x.params[0])
			c := _g_4
			self.__hx_this.sortFields(func(hx_obj_37 map[string]any) *hxrt.Array {
				hx_field_38 := hx_obj_37["fields"]
				if hx_field_38 == nil {
					var hx_zero_39 *hxrt.Array
					return hx_zero_39
				}
				return hx_field_38.(*hxrt.Array)
			}(c))
			self.__hx_this.sortFields(func(hx_obj_40 map[string]any) *hxrt.Array {
				hx_field_41 := hx_obj_40["statics"]
				if hx_field_41 == nil {
					var hx_zero_42 *hxrt.Array
					return hx_zero_42
				}
				return hx_field_41.(*hxrt.Array)
			}(c))
		case 2:
			_g_5 := func(hx_value_43 any) map[string]any {
				if hx_value_43 == nil {
					var hx_zero_44 map[string]any
					return hx_zero_44
				}
				return hx_value_43.(map[string]any)
			}(x.params[0])
			_ = _g_5
		case 3:
			_g_6 := func(hx_value_45 any) map[string]any {
				if hx_value_45 == nil {
					var hx_zero_46 map[string]any
					return hx_zero_46
				}
				return hx_value_45.(map[string]any)
			}(x.params[0])
			_ = _g_6
		case 4:
			_g_7 := func(hx_value_47 any) map[string]any {
				if hx_value_47 == nil {
					var hx_zero_48 map[string]any
					return hx_zero_48
				}
				return hx_value_47.(map[string]any)
			}(x.params[0])
			_ = _g_7
		}
	}
}

func (self *haxe__rtti__XmlParser) sortFields(a *hxrt.Array) {
	haxe__ds__ArraySort_sort(a, func(hx_cmp_left_67 any, hx_cmp_right_68 any) int {
		return func(f1 map[string]any, f2 map[string]any) int {
			v1 := haxe__rtti__TypeApi_isVar(func(hx_obj_49 map[string]any) *haxe__rtti__CType {
				hx_field_50 := hx_obj_49["type"]
				if hx_field_50 == nil {
					var hx_zero_51 *haxe__rtti__CType
					return hx_zero_51
				}
				return hx_field_50.(*haxe__rtti__CType)
			}(f1))
			v2 := haxe__rtti__TypeApi_isVar(func(hx_obj_52 map[string]any) *haxe__rtti__CType {
				hx_field_53 := hx_obj_52["type"]
				if hx_field_53 == nil {
					var hx_zero_54 *haxe__rtti__CType
					return hx_zero_54
				}
				return hx_field_53.(*haxe__rtti__CType)
			}(f2))
			if v1 && !v2 {
				return -1
			}
			if v2 && !v1 {
				return 1
			}
			if hxrt.StringEqualStringPtr(func(hx_obj_55 map[string]any) *string {
				hx_field_56 := hx_obj_55["name"]
				if hx_field_56 == nil {
					var hx_zero_57 *string
					return hx_zero_57
				}
				return hx_field_56.(*string)
			}(f1), hxrt.StringFromLiteral("new")) {
				return -1
			}
			if hxrt.StringEqualStringPtr(func(hx_obj_58 map[string]any) *string {
				hx_field_59 := hx_obj_58["name"]
				if hx_field_59 == nil {
					var hx_zero_60 *string
					return hx_zero_60
				}
				return hx_field_59.(*string)
			}(f2), hxrt.StringFromLiteral("new")) {
				return 1
			}
			return Reflect_compare(func(hx_obj_61 map[string]any) *string {
				hx_field_62 := hx_obj_61["name"]
				if hx_field_62 == nil {
					var hx_zero_63 *string
					return hx_zero_63
				}
				return hx_field_62.(*string)
			}(f1), func(hx_obj_64 map[string]any) *string {
				hx_field_65 := hx_obj_64["name"]
				if hx_field_65 == nil {
					var hx_zero_66 *string
					return hx_zero_66
				}
				return hx_field_65.(*string)
			}(f2))
		}(func(hx_value_69 any) map[string]any {
			if hx_value_69 == nil {
				var hx_zero_70 map[string]any
				return hx_zero_70
			}
			return hx_value_69.(map[string]any)
		}(hx_cmp_left_67), func(hx_value_71 any) map[string]any {
			if hx_value_71 == nil {
				var hx_zero_72 map[string]any
				return hx_zero_72
			}
			return hx_value_71.(map[string]any)
		}(hx_cmp_right_68))
	})
}

func (self *haxe__rtti__XmlParser) process(x *Xml, platform *string) {
	self.curplatform = platform
	self.__hx_this.xroot(x)
}

func (self *haxe__rtti__XmlParser) mergeRights(f1 map[string]any, f2 map[string]any) bool {
	if (((func(hx_obj_73 map[string]any) *haxe__rtti__Rights {
		hx_field_74 := hx_obj_73["get"]
		if hx_field_74 == nil {
			var hx_zero_75 *haxe__rtti__Rights
			return hx_zero_75
		}
		return hx_field_74.(*haxe__rtti__Rights)
	}(f1) == haxe__rtti__Rights_RInline) && (func(hx_obj_76 map[string]any) *haxe__rtti__Rights {
		hx_field_77 := hx_obj_76["set"]
		if hx_field_77 == nil {
			var hx_zero_78 *haxe__rtti__Rights
			return hx_zero_78
		}
		return hx_field_77.(*haxe__rtti__Rights)
	}(f1) == haxe__rtti__Rights_RNo)) && (func(hx_obj_79 map[string]any) *haxe__rtti__Rights {
		hx_field_80 := hx_obj_79["get"]
		if hx_field_80 == nil {
			var hx_zero_81 *haxe__rtti__Rights
			return hx_zero_81
		}
		return hx_field_80.(*haxe__rtti__Rights)
	}(f2) == haxe__rtti__Rights_RNormal)) && (func(hx_obj_82 map[string]any) *haxe__rtti__Rights {
		hx_field_83 := hx_obj_82["set"]
		if hx_field_83 == nil {
			var hx_zero_84 *haxe__rtti__Rights
			return hx_zero_84
		}
		return hx_field_83.(*haxe__rtti__Rights)
	}(f2) == haxe__rtti__Rights_RMethod) {
		f1["get"] = haxe__rtti__Rights_RNormal
		f1["set"] = haxe__rtti__Rights_RMethod
		return true
	}
	return (Type_enumEq(func(hx_obj_85 map[string]any) *haxe__rtti__Rights {
		hx_field_86 := hx_obj_85["get"]
		if hx_field_86 == nil {
			var hx_zero_87 *haxe__rtti__Rights
			return hx_zero_87
		}
		return hx_field_86.(*haxe__rtti__Rights)
	}(f1), func(hx_obj_88 map[string]any) *haxe__rtti__Rights {
		hx_field_89 := hx_obj_88["get"]
		if hx_field_89 == nil {
			var hx_zero_90 *haxe__rtti__Rights
			return hx_zero_90
		}
		return hx_field_89.(*haxe__rtti__Rights)
	}(f2)) && Type_enumEq(func(hx_obj_91 map[string]any) *haxe__rtti__Rights {
		hx_field_92 := hx_obj_91["set"]
		if hx_field_92 == nil {
			var hx_zero_93 *haxe__rtti__Rights
			return hx_zero_93
		}
		return hx_field_92.(*haxe__rtti__Rights)
	}(f1), func(hx_obj_94 map[string]any) *haxe__rtti__Rights {
		hx_field_95 := hx_obj_94["set"]
		if hx_field_95 == nil {
			var hx_zero_96 *haxe__rtti__Rights
			return hx_zero_96
		}
		return hx_field_95.(*haxe__rtti__Rights)
	}(f2)))
}

func (self *haxe__rtti__XmlParser) mergeDoc(f1 map[string]any, f2 map[string]any) bool {
	if hxrt.StringEqualStringPtr(func(hx_obj_106 map[string]any) *string {
		hx_field_107 := hx_obj_106["doc"]
		if hx_field_107 == nil {
			var hx_zero_108 *string
			return hx_zero_108
		}
		return hx_field_107.(*string)
	}(f1), nil) {
		f1["doc"] = func(hx_obj_97 map[string]any) *string {
			hx_field_98 := hx_obj_97["doc"]
			if hx_field_98 == nil {
				var hx_zero_99 *string
				return hx_zero_99
			}
			return hx_field_98.(*string)
		}(f2)
	} else {
		if hxrt.StringEqualStringPtr(func(hx_obj_103 map[string]any) *string {
			hx_field_104 := hx_obj_103["doc"]
			if hx_field_104 == nil {
				var hx_zero_105 *string
				return hx_zero_105
			}
			return hx_field_104.(*string)
		}(f2), nil) {
			f2["doc"] = func(hx_obj_100 map[string]any) *string {
				hx_field_101 := hx_obj_100["doc"]
				if hx_field_101 == nil {
					var hx_zero_102 *string
					return hx_zero_102
				}
				return hx_field_101.(*string)
			}(f1)
		}
	}
	return true
}

func (self *haxe__rtti__XmlParser) mergeFields(f map[string]any, f2 map[string]any) bool {
	return (haxe__rtti__TypeApi_fieldEq(f, f2) || (((hxrt.StringEqualStringPtr(func(hx_obj_109 map[string]any) *string {
		hx_field_110 := hx_obj_109["name"]
		if hx_field_110 == nil {
			var hx_zero_111 *string
			return hx_zero_111
		}
		return hx_field_110.(*string)
	}(f), func(hx_obj_112 map[string]any) *string {
		hx_field_113 := hx_obj_112["name"]
		if hx_field_113 == nil {
			var hx_zero_114 *string
			return hx_zero_114
		}
		return hx_field_113.(*string)
	}(f2)) && (self.__hx_this.mergeRights(f, f2) || self.__hx_this.mergeRights(f2, f))) && self.__hx_this.mergeDoc(f, f2)) && haxe__rtti__TypeApi_fieldEq(f, f2)))
}

func (self *haxe__rtti__XmlParser) mergeClasses(c map[string]any, c2 map[string]any) bool {
	if func(hx_obj_115 map[string]any) bool {
		hx_field_116 := hx_obj_115["isInterface"]
		if hx_field_116 == nil {
			var hx_zero_117 bool
			return hx_zero_117
		}
		return hx_field_116.(bool)
	}(c) != func(hx_obj_118 map[string]any) bool {
		hx_field_119 := hx_obj_118["isInterface"]
		if hx_field_119 == nil {
			var hx_zero_120 bool
			return hx_zero_120
		}
		return hx_field_119.(bool)
	}(c2) {
		return false
	}
	if !hxrt.StringEqualStringPtr(self.curplatform, nil) {
		hx_arr_121 := func(hx_obj_122 map[string]any) *hxrt.Array {
			hx_field_123 := hx_obj_122["platforms"]
			if hx_field_123 == nil {
				var hx_zero_124 *hxrt.Array
				return hx_zero_124
			}
			return hx_field_123.(*hxrt.Array)
		}(c)
		hx_arr_121.Push(self.curplatform)
	}
	if func(hx_obj_125 map[string]any) bool {
		hx_field_126 := hx_obj_125["isExtern"]
		if hx_field_126 == nil {
			var hx_zero_127 bool
			return hx_zero_127
		}
		return hx_field_126.(bool)
	}(c) != func(hx_obj_128 map[string]any) bool {
		hx_field_129 := hx_obj_128["isExtern"]
		if hx_field_129 == nil {
			var hx_zero_130 bool
			return hx_zero_130
		}
		return hx_field_129.(bool)
	}(c2) {
		c["isExtern"] = false
	}
	_g := 0
	_g1 := func(hx_obj_131 map[string]any) *hxrt.Array {
		hx_field_132 := hx_obj_131["fields"]
		if hx_field_132 == nil {
			var hx_zero_133 *hxrt.Array
			return hx_zero_133
		}
		return hx_field_132.(*hxrt.Array)
	}(c2)
	for _g < _g1.Len() {
		f2 := func(hx_value_134 any) map[string]any {
			if hx_value_134 == nil {
				var hx_zero_135 map[string]any
				return hx_zero_135
			}
			return hx_value_134.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		var found map[string]any = nil
		_g_1 := 0
		_g1_1 := func(hx_obj_136 map[string]any) *hxrt.Array {
			hx_field_137 := hx_obj_136["fields"]
			if hx_field_137 == nil {
				var hx_zero_138 *hxrt.Array
				return hx_zero_138
			}
			return hx_field_137.(*hxrt.Array)
		}(c)
		for _g_1 < _g1_1.Len() {
			f := func(hx_value_139 any) map[string]any {
				if hx_value_139 == nil {
					var hx_zero_140 map[string]any
					return hx_zero_140
				}
				return hx_value_139.(map[string]any)
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
			hx_arr_141 := func(hx_obj_142 map[string]any) *hxrt.Array {
				hx_field_143 := hx_obj_142["fields"]
				if hx_field_143 == nil {
					var hx_zero_144 *hxrt.Array
					return hx_zero_144
				}
				return hx_field_143.(*hxrt.Array)
			}(c)
			hx_arr_141.Push(f2)
		} else {
			if !hxrt.StringEqualStringPtr(self.curplatform, nil) {
				hx_arr_145 := func(hx_obj_146 map[string]any) *hxrt.Array {
					hx_field_147 := hx_obj_146["platforms"]
					if hx_field_147 == nil {
						var hx_zero_148 *hxrt.Array
						return hx_zero_148
					}
					return hx_field_147.(*hxrt.Array)
				}(found)
				hx_arr_145.Push(self.curplatform)
			}
		}
	}
	_g_2 := 0
	_g1_2 := func(hx_obj_149 map[string]any) *hxrt.Array {
		hx_field_150 := hx_obj_149["statics"]
		if hx_field_150 == nil {
			var hx_zero_151 *hxrt.Array
			return hx_zero_151
		}
		return hx_field_150.(*hxrt.Array)
	}(c2)
	for _g_2 < _g1_2.Len() {
		f2_1 := func(hx_value_152 any) map[string]any {
			if hx_value_152 == nil {
				var hx_zero_153 map[string]any
				return hx_zero_153
			}
			return hx_value_152.(map[string]any)
		}(_g1_2.Get(_g_2))
		_g_2 = int(int32((_g_2 + 1)))
		var found_1 map[string]any = nil
		_g_3 := 0
		_g1_3 := func(hx_obj_154 map[string]any) *hxrt.Array {
			hx_field_155 := hx_obj_154["statics"]
			if hx_field_155 == nil {
				var hx_zero_156 *hxrt.Array
				return hx_zero_156
			}
			return hx_field_155.(*hxrt.Array)
		}(c)
		for _g_3 < _g1_3.Len() {
			f_1 := func(hx_value_157 any) map[string]any {
				if hx_value_157 == nil {
					var hx_zero_158 map[string]any
					return hx_zero_158
				}
				return hx_value_157.(map[string]any)
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
			hx_arr_159 := func(hx_obj_160 map[string]any) *hxrt.Array {
				hx_field_161 := hx_obj_160["statics"]
				if hx_field_161 == nil {
					var hx_zero_162 *hxrt.Array
					return hx_zero_162
				}
				return hx_field_161.(*hxrt.Array)
			}(c)
			hx_arr_159.Push(f2_1)
		} else {
			if !hxrt.StringEqualStringPtr(self.curplatform, nil) {
				hx_arr_163 := func(hx_obj_164 map[string]any) *hxrt.Array {
					hx_field_165 := hx_obj_164["platforms"]
					if hx_field_165 == nil {
						var hx_zero_166 *hxrt.Array
						return hx_zero_166
					}
					return hx_field_165.(*hxrt.Array)
				}(found_1)
				hx_arr_163.Push(self.curplatform)
			}
		}
	}
	return true
}

func (self *haxe__rtti__XmlParser) mergeEnums(e map[string]any, e2 map[string]any) bool {
	if func(hx_obj_167 map[string]any) bool {
		hx_field_168 := hx_obj_167["isExtern"]
		if hx_field_168 == nil {
			var hx_zero_169 bool
			return hx_zero_169
		}
		return hx_field_168.(bool)
	}(e) != func(hx_obj_170 map[string]any) bool {
		hx_field_171 := hx_obj_170["isExtern"]
		if hx_field_171 == nil {
			var hx_zero_172 bool
			return hx_zero_172
		}
		return hx_field_171.(bool)
	}(e2) {
		return false
	}
	if !hxrt.StringEqualStringPtr(self.curplatform, nil) {
		hx_arr_173 := func(hx_obj_174 map[string]any) *hxrt.Array {
			hx_field_175 := hx_obj_174["platforms"]
			if hx_field_175 == nil {
				var hx_zero_176 *hxrt.Array
				return hx_zero_176
			}
			return hx_field_175.(*hxrt.Array)
		}(e)
		hx_arr_173.Push(self.curplatform)
	}
	_g := 0
	_g1 := func(hx_obj_177 map[string]any) *hxrt.Array {
		hx_field_178 := hx_obj_177["constructors"]
		if hx_field_178 == nil {
			var hx_zero_179 *hxrt.Array
			return hx_zero_179
		}
		return hx_field_178.(*hxrt.Array)
	}(e2)
	for _g < _g1.Len() {
		c2 := func(hx_value_180 any) map[string]any {
			if hx_value_180 == nil {
				var hx_zero_181 map[string]any
				return hx_zero_181
			}
			return hx_value_180.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		var found map[string]any = nil
		_g_1 := 0
		_g1_1 := func(hx_obj_182 map[string]any) *hxrt.Array {
			hx_field_183 := hx_obj_182["constructors"]
			if hx_field_183 == nil {
				var hx_zero_184 *hxrt.Array
				return hx_zero_184
			}
			return hx_field_183.(*hxrt.Array)
		}(e)
		for _g_1 < _g1_1.Len() {
			c := func(hx_value_185 any) map[string]any {
				if hx_value_185 == nil {
					var hx_zero_186 map[string]any
					return hx_zero_186
				}
				return hx_value_185.(map[string]any)
			}(_g1_1.Get(_g_1))
			_g_1 = int(int32((_g_1 + 1)))
			if haxe__rtti__TypeApi_constructorEq(c, c2) {
				found = c
				break
			}
		}
		if found == nil {
			hx_arr_187 := func(hx_obj_188 map[string]any) *hxrt.Array {
				hx_field_189 := hx_obj_188["constructors"]
				if hx_field_189 == nil {
					var hx_zero_190 *hxrt.Array
					return hx_zero_190
				}
				return hx_field_189.(*hxrt.Array)
			}(e)
			hx_arr_187.Push(c2)
		} else {
			if !hxrt.StringEqualStringPtr(self.curplatform, nil) {
				hx_arr_191 := func(hx_obj_192 map[string]any) *hxrt.Array {
					hx_field_193 := hx_obj_192["platforms"]
					if hx_field_193 == nil {
						var hx_zero_194 *hxrt.Array
						return hx_zero_194
					}
					return hx_field_193.(*hxrt.Array)
				}(found)
				hx_arr_191.Push(self.curplatform)
			}
		}
	}
	return true
}

func (self *haxe__rtti__XmlParser) mergeTypedefs(t map[string]any, t2 map[string]any) bool {
	if hxrt.StringEqualStringPtr(self.curplatform, nil) {
		return false
	}
	hx_arr_195 := func(hx_obj_196 map[string]any) *hxrt.Array {
		hx_field_197 := hx_obj_196["platforms"]
		if hx_field_197 == nil {
			var hx_zero_198 *hxrt.Array
			return hx_zero_198
		}
		return hx_field_197.(*hxrt.Array)
	}(t)
	hx_arr_195.Push(self.curplatform)
	var this1 haxe__IMap = func(hx_obj_199 map[string]any) *haxe__ds__StringMap {
		hx_field_200 := hx_obj_199["types"]
		if hx_field_200 == nil {
			var hx_zero_201 *haxe__ds__StringMap
			return hx_zero_201
		}
		return hx_field_200.(*haxe__ds__StringMap)
	}(t)
	key := self.curplatform
	value := func(hx_obj_202 map[string]any) *haxe__rtti__CType {
		hx_field_203 := hx_obj_202["type"]
		if hx_field_203 == nil {
			var hx_zero_204 *haxe__rtti__CType
			return hx_zero_204
		}
		return hx_field_203.(*haxe__rtti__CType)
	}(t2)
	this1.(*haxe__ds__StringMap).__hx_this.set(key, value)
	return true
}

func (self *haxe__rtti__XmlParser) mergeAbstracts(a map[string]any, a2 map[string]any) bool {
	if hxrt.StringEqualStringPtr(self.curplatform, nil) {
		return false
	}
	if (func(hx_obj_205 map[string]any) *hxrt.Array {
		hx_field_206 := hx_obj_205["to"]
		if hx_field_206 == nil {
			var hx_zero_207 *hxrt.Array
			return hx_zero_207
		}
		return hx_field_206.(*hxrt.Array)
	}(a).Len() != func(hx_obj_208 map[string]any) *hxrt.Array {
		hx_field_209 := hx_obj_208["to"]
		if hx_field_209 == nil {
			var hx_zero_210 *hxrt.Array
			return hx_zero_210
		}
		return hx_field_209.(*hxrt.Array)
	}(a2).Len()) || (func(hx_obj_211 map[string]any) *hxrt.Array {
		hx_field_212 := hx_obj_211["from"]
		if hx_field_212 == nil {
			var hx_zero_213 *hxrt.Array
			return hx_zero_213
		}
		return hx_field_212.(*hxrt.Array)
	}(a).Len() != func(hx_obj_214 map[string]any) *hxrt.Array {
		hx_field_215 := hx_obj_214["from"]
		if hx_field_215 == nil {
			var hx_zero_216 *hxrt.Array
			return hx_zero_216
		}
		return hx_field_215.(*hxrt.Array)
	}(a2).Len()) {
		return false
	}
	_g := 0
	_g1 := func(hx_obj_217 map[string]any) *hxrt.Array {
		hx_field_218 := hx_obj_217["to"]
		if hx_field_218 == nil {
			var hx_zero_219 *hxrt.Array
			return hx_zero_219
		}
		return hx_field_218.(*hxrt.Array)
	}(a).Len()
	for _g < _g1 {
		hx_post_220 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_220
		if !haxe__rtti__TypeApi_typeEq(func(hx_obj_226 map[string]any) *haxe__rtti__CType {
			hx_field_227 := hx_obj_226["t"]
			if hx_field_227 == nil {
				var hx_zero_228 *haxe__rtti__CType
				return hx_zero_228
			}
			return hx_field_227.(*haxe__rtti__CType)
		}(func(hx_value_224 any) map[string]any {
			if hx_value_224 == nil {
				var hx_zero_225 map[string]any
				return hx_zero_225
			}
			return hx_value_224.(map[string]any)
		}(func(hx_obj_221 map[string]any) *hxrt.Array {
			hx_field_222 := hx_obj_221["to"]
			if hx_field_222 == nil {
				var hx_zero_223 *hxrt.Array
				return hx_zero_223
			}
			return hx_field_222.(*hxrt.Array)
		}(a).Get(i))), func(hx_obj_234 map[string]any) *haxe__rtti__CType {
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
			hx_field_230 := hx_obj_229["to"]
			if hx_field_230 == nil {
				var hx_zero_231 *hxrt.Array
				return hx_zero_231
			}
			return hx_field_230.(*hxrt.Array)
		}(a2).Get(i)))) {
			return false
		}
	}
	_g_1 := 0
	_g1_1 := func(hx_obj_237 map[string]any) *hxrt.Array {
		hx_field_238 := hx_obj_237["from"]
		if hx_field_238 == nil {
			var hx_zero_239 *hxrt.Array
			return hx_zero_239
		}
		return hx_field_238.(*hxrt.Array)
	}(a).Len()
	for _g_1 < _g1_1 {
		hx_post_240 := _g_1
		_g_1 = int(int32((_g_1 + 1)))
		i_1 := hx_post_240
		if !haxe__rtti__TypeApi_typeEq(func(hx_obj_246 map[string]any) *haxe__rtti__CType {
			hx_field_247 := hx_obj_246["t"]
			if hx_field_247 == nil {
				var hx_zero_248 *haxe__rtti__CType
				return hx_zero_248
			}
			return hx_field_247.(*haxe__rtti__CType)
		}(func(hx_value_244 any) map[string]any {
			if hx_value_244 == nil {
				var hx_zero_245 map[string]any
				return hx_zero_245
			}
			return hx_value_244.(map[string]any)
		}(func(hx_obj_241 map[string]any) *hxrt.Array {
			hx_field_242 := hx_obj_241["from"]
			if hx_field_242 == nil {
				var hx_zero_243 *hxrt.Array
				return hx_zero_243
			}
			return hx_field_242.(*hxrt.Array)
		}(a).Get(i_1))), func(hx_obj_254 map[string]any) *haxe__rtti__CType {
			hx_field_255 := hx_obj_254["t"]
			if hx_field_255 == nil {
				var hx_zero_256 *haxe__rtti__CType
				return hx_zero_256
			}
			return hx_field_255.(*haxe__rtti__CType)
		}(func(hx_value_252 any) map[string]any {
			if hx_value_252 == nil {
				var hx_zero_253 map[string]any
				return hx_zero_253
			}
			return hx_value_252.(map[string]any)
		}(func(hx_obj_249 map[string]any) *hxrt.Array {
			hx_field_250 := hx_obj_249["from"]
			if hx_field_250 == nil {
				var hx_zero_251 *hxrt.Array
				return hx_zero_251
			}
			return hx_field_250.(*hxrt.Array)
		}(a2).Get(i_1)))) {
			return false
		}
	}
	if func(hx_obj_263 map[string]any) map[string]any {
		hx_field_264 := hx_obj_263["impl"]
		if hx_field_264 == nil {
			var hx_zero_265 map[string]any
			return hx_zero_265
		}
		return hx_field_264.(map[string]any)
	}(a2) != nil {
		self.__hx_this.mergeClasses(func(hx_obj_257 map[string]any) map[string]any {
			hx_field_258 := hx_obj_257["impl"]
			if hx_field_258 == nil {
				var hx_zero_259 map[string]any
				return hx_zero_259
			}
			return hx_field_258.(map[string]any)
		}(a), func(hx_obj_260 map[string]any) map[string]any {
			hx_field_261 := hx_obj_260["impl"]
			if hx_field_261 == nil {
				var hx_zero_262 map[string]any
				return hx_zero_262
			}
			return hx_field_261.(map[string]any)
		}(a2))
	}
	hx_arr_266 := func(hx_obj_267 map[string]any) *hxrt.Array {
		hx_field_268 := hx_obj_267["platforms"]
		if hx_field_268 == nil {
			var hx_zero_269 *hxrt.Array
			return hx_zero_269
		}
		return hx_field_268.(*hxrt.Array)
	}(a)
	hx_arr_266.Push(self.curplatform)
	return true
}

func (self *haxe__rtti__XmlParser) merge(t *haxe__rtti__TypeTree) {
	inf := haxe__rtti__TypeApi_typeInfos(t)
	pack := self.__hx_this.splitString(func(hx_obj_270 map[string]any) *string {
		hx_field_271 := hx_obj_270["path"]
		if hx_field_271 == nil {
			var hx_zero_272 *string
			return hx_zero_272
		}
		return hx_field_271.(*string)
	}(inf), hxrt.StringFromLiteral("."))
	cur := self.root
	curpack := hxrt.NewArray()
	pack.Pop()
	_g := 0
	for _g < pack.Len() {
		p := func(hx_value_274 any) *string {
			if hx_value_274 == nil {
				var hx_zero_275 *string
				return hx_zero_275
			}
			return hx_value_274.(*string)
		}(pack.Get(_g))
		_g = int(int32((_g + 1)))
		found := false
		_g_1 := 0
		for _g_1 < cur.Len() {
			pk := func(hx_value_276 any) *haxe__rtti__TypeTree {
				if hx_value_276 == nil {
					var hx_zero_277 *haxe__rtti__TypeTree
					return hx_zero_277
				}
				return hx_value_276.(*haxe__rtti__TypeTree)
			}(cur.Get(_g_1))
			_g_1 = int(int32((_g_1 + 1)))
			if pk.tag == 0 {
				_g_2 := func(hx_value_278 any) *string {
					if hx_value_278 == nil {
						var hx_zero_279 *string
						return hx_zero_279
					}
					return hx_value_278.(*string)
				}(pk.params[0])
				_g1 := func(hx_value_280 any) *string {
					if hx_value_280 == nil {
						var hx_zero_281 *string
						return hx_zero_281
					}
					return hx_value_280.(*string)
				}(pk.params[1])
				_ = _g1
				_g1_1 := func(hx_value_282 any) *hxrt.Array {
					if hx_value_282 == nil {
						var hx_zero_283 *hxrt.Array
						return hx_zero_283
					}
					return hx_value_282.(*hxrt.Array)
				}(pk.params[2])
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
		ct := func(hx_value_286 any) *haxe__rtti__TypeTree {
			if hx_value_286 == nil {
				var hx_zero_287 *haxe__rtti__TypeTree
				return hx_zero_287
			}
			return hx_value_286.(*haxe__rtti__TypeTree)
		}(cur.Get(_g_3))
		_g_3 = int(int32((_g_3 + 1)))
		if func() bool {
			var hx_if_294 bool
			if ct.tag == 0 {
				_g_4 := func(hx_value_288 any) *string {
					if hx_value_288 == nil {
						var hx_zero_289 *string
						return hx_zero_289
					}
					return hx_value_288.(*string)
				}(ct.params[0])
				_ = _g_4
				_g_5 := func(hx_value_290 any) *string {
					if hx_value_290 == nil {
						var hx_zero_291 *string
						return hx_zero_291
					}
					return hx_value_290.(*string)
				}(ct.params[1])
				_ = _g_5
				_g_6 := func(hx_value_292 any) *hxrt.Array {
					if hx_value_292 == nil {
						var hx_zero_293 *hxrt.Array
						return hx_zero_293
					}
					return hx_value_292.(*hxrt.Array)
				}(ct.params[2])
				_ = _g_6
				hx_if_294 = true
			} else {
				hx_if_294 = false
			}
			return hx_if_294
		}() {
			continue
		}
		tinf := haxe__rtti__TypeApi_typeInfos(ct)
		if hxrt.StringEqualStringPtr(func(hx_obj_387 map[string]any) *string {
			hx_field_388 := hx_obj_387["path"]
			if hx_field_388 == nil {
				var hx_zero_389 *string
				return hx_zero_389
			}
			return hx_field_388.(*string)
		}(tinf), func(hx_obj_390 map[string]any) *string {
			hx_field_391 := hx_obj_390["path"]
			if hx_field_391 == nil {
				var hx_zero_392 *string
				return hx_zero_392
			}
			return hx_field_391.(*string)
		}(inf)) {
			sameType := true
			if hxrt.StringEqualStringPtr(func(hx_obj_304 map[string]any) *string {
				hx_field_305 := hx_obj_304["doc"]
				if hx_field_305 == nil {
					var hx_zero_306 *string
					return hx_zero_306
				}
				return hx_field_305.(*string)
			}(tinf), nil) != hxrt.StringEqualStringPtr(func(hx_obj_307 map[string]any) *string {
				hx_field_308 := hx_obj_307["doc"]
				if hx_field_308 == nil {
					var hx_zero_309 *string
					return hx_zero_309
				}
				return hx_field_308.(*string)
			}(inf), nil) {
				if hxrt.StringEqualStringPtr(func(hx_obj_301 map[string]any) *string {
					hx_field_302 := hx_obj_301["doc"]
					if hx_field_302 == nil {
						var hx_zero_303 *string
						return hx_zero_303
					}
					return hx_field_302.(*string)
				}(inf), nil) {
					inf["doc"] = func(hx_obj_295 map[string]any) *string {
						hx_field_296 := hx_obj_295["doc"]
						if hx_field_296 == nil {
							var hx_zero_297 *string
							return hx_zero_297
						}
						return hx_field_296.(*string)
					}(tinf)
				} else {
					tinf["doc"] = func(hx_obj_298 map[string]any) *string {
						hx_field_299 := hx_obj_298["doc"]
						if hx_field_299 == nil {
							var hx_zero_300 *string
							return hx_zero_300
						}
						return hx_field_299.(*string)
					}(inf)
				}
			}
			if hxrt.StringEqualStringPtr(func(hx_obj_310 map[string]any) *string {
				hx_field_311 := hx_obj_310["path"]
				if hx_field_311 == nil {
					var hx_zero_312 *string
					return hx_zero_312
				}
				return hx_field_311.(*string)
			}(tinf), hxrt.StringFromLiteral("haxe._Int64.NativeInt64")) {
				continue
			}
			if (hxrt.StringEqualStringPtr(func(hx_obj_335 map[string]any) *string {
				hx_field_336 := hx_obj_335["module"]
				if hx_field_336 == nil {
					var hx_zero_337 *string
					return hx_zero_337
				}
				return hx_field_336.(*string)
			}(tinf), func(hx_obj_338 map[string]any) *string {
				hx_field_339 := hx_obj_338["module"]
				if hx_field_339 == nil {
					var hx_zero_340 *string
					return hx_zero_340
				}
				return hx_field_339.(*string)
			}(inf)) && hxrt.StringEqualStringPtr(func(hx_obj_341 map[string]any) *string {
				hx_field_342 := hx_obj_341["doc"]
				if hx_field_342 == nil {
					var hx_zero_343 *string
					return hx_zero_343
				}
				return hx_field_342.(*string)
			}(tinf), func(hx_obj_344 map[string]any) *string {
				hx_field_345 := hx_obj_344["doc"]
				if hx_field_345 == nil {
					var hx_zero_346 *string
					return hx_zero_346
				}
				return hx_field_345.(*string)
			}(inf))) && (func(hx_obj_347 map[string]any) bool {
				hx_field_348 := hx_obj_347["isPrivate"]
				if hx_field_348 == nil {
					var hx_zero_349 bool
					return hx_zero_349
				}
				return hx_field_348.(bool)
			}(tinf) == func(hx_obj_350 map[string]any) bool {
				hx_field_351 := hx_obj_350["isPrivate"]
				if hx_field_351 == nil {
					var hx_zero_352 bool
					return hx_zero_352
				}
				return hx_field_351.(bool)
			}(inf)) {
				switch ct.tag {
				case 0:
					_g_7 := func(hx_value_313 any) *string {
						if hx_value_313 == nil {
							var hx_zero_314 *string
							return hx_zero_314
						}
						return hx_value_313.(*string)
					}(ct.params[0])
					_ = _g_7
					_g_8 := func(hx_value_315 any) *string {
						if hx_value_315 == nil {
							var hx_zero_316 *string
							return hx_zero_316
						}
						return hx_value_315.(*string)
					}(ct.params[1])
					_ = _g_8
					_g_9 := func(hx_value_317 any) *hxrt.Array {
						if hx_value_317 == nil {
							var hx_zero_318 *hxrt.Array
							return hx_zero_318
						}
						return hx_value_317.(*hxrt.Array)
					}(ct.params[2])
					_ = _g_9
					sameType = false
				case 1:
					_g_10 := func(hx_value_319 any) map[string]any {
						if hx_value_319 == nil {
							var hx_zero_320 map[string]any
							return hx_zero_320
						}
						return hx_value_319.(map[string]any)
					}(ct.params[0])
					c := _g_10
					if t.tag == 1 {
						_g_11 := func(hx_value_321 any) map[string]any {
							if hx_value_321 == nil {
								var hx_zero_322 map[string]any
								return hx_zero_322
							}
							return hx_value_321.(map[string]any)
						}(t.params[0])
						c2 := _g_11
						if self.__hx_this.mergeClasses(c, c2) {
							return
						}
					} else {
						sameType = false
					}
				case 2:
					_g_12 := func(hx_value_323 any) map[string]any {
						if hx_value_323 == nil {
							var hx_zero_324 map[string]any
							return hx_zero_324
						}
						return hx_value_323.(map[string]any)
					}(ct.params[0])
					e := _g_12
					if t.tag == 2 {
						_g_13 := func(hx_value_325 any) map[string]any {
							if hx_value_325 == nil {
								var hx_zero_326 map[string]any
								return hx_zero_326
							}
							return hx_value_325.(map[string]any)
						}(t.params[0])
						e2 := _g_13
						if self.__hx_this.mergeEnums(e, e2) {
							return
						}
					} else {
						sameType = false
					}
				case 3:
					_g_14 := func(hx_value_327 any) map[string]any {
						if hx_value_327 == nil {
							var hx_zero_328 map[string]any
							return hx_zero_328
						}
						return hx_value_327.(map[string]any)
					}(ct.params[0])
					td := _g_14
					if t.tag == 3 {
						_g_15 := func(hx_value_329 any) map[string]any {
							if hx_value_329 == nil {
								var hx_zero_330 map[string]any
								return hx_zero_330
							}
							return hx_value_329.(map[string]any)
						}(t.params[0])
						td2 := _g_15
						if self.__hx_this.mergeTypedefs(td, td2) {
							return
						}
					} else {
					}
				case 4:
					_g_16 := func(hx_value_331 any) map[string]any {
						if hx_value_331 == nil {
							var hx_zero_332 map[string]any
							return hx_zero_332
						}
						return hx_value_331.(map[string]any)
					}(ct.params[0])
					a := _g_16
					if t.tag == 4 {
						_g_17 := func(hx_value_333 any) map[string]any {
							if hx_value_333 == nil {
								var hx_zero_334 map[string]any
								return hx_zero_334
							}
							return hx_value_333.(map[string]any)
						}(t.params[0])
						a2 := _g_17
						if self.__hx_this.mergeAbstracts(a, a2) {
							return
						}
					} else {
						sameType = false
					}
				}
			}
			var hx_if_380 *string
			if !hxrt.StringEqualStringPtr(func(hx_obj_353 map[string]any) *string {
				hx_field_354 := hx_obj_353["module"]
				if hx_field_354 == nil {
					var hx_zero_355 *string
					return hx_zero_355
				}
				return hx_field_354.(*string)
			}(tinf), func(hx_obj_356 map[string]any) *string {
				hx_field_357 := hx_obj_356["module"]
				if hx_field_357 == nil {
					var hx_zero_358 *string
					return hx_zero_358
				}
				return hx_field_357.(*string)
			}(inf)) {
				hx_if_380 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("module "), func(hx_obj_359 map[string]any) *string {
					hx_field_360 := hx_obj_359["module"]
					if hx_field_360 == nil {
						var hx_zero_361 *string
						return hx_zero_361
					}
					return hx_field_360.(*string)
				}(inf)), hxrt.StringFromLiteral(" should be ")), func(hx_obj_362 map[string]any) *string {
					hx_field_363 := hx_obj_362["module"]
					if hx_field_363 == nil {
						var hx_zero_364 *string
						return hx_zero_364
					}
					return hx_field_363.(*string)
				}(tinf))
			} else {
				var hx_if_379 *string
				if !hxrt.StringEqualStringPtr(func(hx_obj_365 map[string]any) *string {
					hx_field_366 := hx_obj_365["doc"]
					if hx_field_366 == nil {
						var hx_zero_367 *string
						return hx_zero_367
					}
					return hx_field_366.(*string)
				}(tinf), func(hx_obj_368 map[string]any) *string {
					hx_field_369 := hx_obj_368["doc"]
					if hx_field_369 == nil {
						var hx_zero_370 *string
						return hx_zero_370
					}
					return hx_field_369.(*string)
				}(inf)) {
					hx_if_379 = hxrt.StringFromLiteral("documentation is different")
				} else {
					var hx_if_378 *string
					if func(hx_obj_371 map[string]any) bool {
						hx_field_372 := hx_obj_371["isPrivate"]
						if hx_field_372 == nil {
							var hx_zero_373 bool
							return hx_zero_373
						}
						return hx_field_372.(bool)
					}(tinf) != func(hx_obj_374 map[string]any) bool {
						hx_field_375 := hx_obj_374["isPrivate"]
						if hx_field_375 == nil {
							var hx_zero_376 bool
							return hx_zero_376
						}
						return hx_field_375.(bool)
					}(inf) {
						hx_if_378 = hxrt.StringFromLiteral("private flag is different")
					} else {
						var hx_if_377 *string
						if !sameType {
							hx_if_377 = hxrt.StringFromLiteral("type kind is different")
						} else {
							hx_if_377 = hxrt.StringFromLiteral("could not merge definition")
						}
						hx_if_378 = hx_if_377
					}
					hx_if_379 = hx_if_378
				}
				hx_if_380 = hx_if_379
			}
			msg := hx_if_380
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Incompatibilities between "), func(hx_obj_381 map[string]any) *string {
				hx_field_382 := hx_obj_381["path"]
				if hx_field_382 == nil {
					var hx_zero_383 *string
					return hx_zero_383
				}
				return hx_field_382.(*string)
			}(tinf)), hxrt.StringFromLiteral(" in ")), self.__hx_this.joinStringArray(func(hx_obj_384 map[string]any) *hxrt.Array {
				hx_field_385 := hx_obj_384["platforms"]
				if hx_field_385 == nil {
					var hx_zero_386 *hxrt.Array
					return hx_zero_386
				}
				return hx_field_385.(*hxrt.Array)
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
	for func(hx_obj_394 map[string]any) func() bool {
		hx_field_395 := hx_obj_394["hasNext"]
		if hx_field_395 == nil {
			var hx_zero_396 func() bool
			return hx_zero_396
		}
		return hx_field_395.(func() bool)
	}(c)() {
		c_1 := func(hx_obj_397 map[string]any) func() *Xml {
			hx_field_398 := hx_obj_397["next"]
			if hx_field_398 == nil {
				var hx_zero_399 func() *Xml
				return hx_zero_399
			}
			return hx_field_398.(func() *Xml)
		}(c)()
		self.__hx_this.merge(self.__hx_this.processElement(c_1))
	}
}

func (self *haxe__rtti__XmlParser) processElement(x *Xml) *haxe__rtti__TypeTree {
	var hx_if_400 *string
	if hxrt.HaxeEqual(x.nodeType, Xml_Document) {
		hx_if_400 = hxrt.StringFromLiteral("Document")
	} else {
		if !hxrt.HaxeEqual(x.nodeType, Xml_Element) {
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(x.nodeType))))
		}
		hx_if_400 = x.nodeName
	}
	nodeName := hx_if_400
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
	var hx_throw_zero_401 *haxe__rtti__TypeTree
	return hx_throw_zero_401
}

func (self *haxe__rtti__XmlParser) xmeta(x *Xml) *hxrt.Array {
	ml := hxrt.NewArray()
	m := x.__hx_this.elementsNamed(hxrt.StringFromLiteral("m"))
	for func(hx_obj_402 map[string]any) func() bool {
		hx_field_403 := hx_obj_402["hasNext"]
		if hx_field_403 == nil {
			var hx_zero_404 func() bool
			return hx_zero_404
		}
		return hx_field_403.(func() bool)
	}(m)() {
		m_1 := func(hx_obj_405 map[string]any) func() *Xml {
			hx_field_406 := hx_obj_405["next"]
			if hx_field_406 == nil {
				var hx_zero_407 func() *Xml
				return hx_zero_407
			}
			return hx_field_406.(func() *Xml)
		}(m)()
		pl := hxrt.NewArray()
		p := m_1.__hx_this.elementsNamed(hxrt.StringFromLiteral("e"))
		for func(hx_obj_408 map[string]any) func() bool {
			hx_field_409 := hx_obj_408["hasNext"]
			if hx_field_409 == nil {
				var hx_zero_410 func() bool
				return hx_zero_410
			}
			return hx_field_409.(func() bool)
		}(p)() {
			p_1 := func(hx_obj_411 map[string]any) func() *Xml {
				hx_field_412 := hx_obj_411["next"]
				if hx_field_412 == nil {
					var hx_zero_413 func() *Xml
					return hx_zero_413
				}
				return hx_field_412.(func() *Xml)
			}(p)()
			pl.Push(self.__hx_this.innerHTML(p_1))
		}
		hx_obj_416 := map[string]any{}
		hx_obj_416["name"] = self.__hx_this.requireAttr(m_1, hxrt.StringFromLiteral("n"))
		hx_obj_416["params"] = pl
		ml.Push(hx_obj_416)
	}
	return ml
}

func (self *haxe__rtti__XmlParser) xoverloads(x *Xml) *hxrt.Array {
	l := hxrt.NewArray()
	m := x.__hx_this.elements()
	for func(hx_obj_417 map[string]any) func() bool {
		hx_field_418 := hx_obj_417["hasNext"]
		if hx_field_418 == nil {
			var hx_zero_419 func() bool
			return hx_zero_419
		}
		return hx_field_418.(func() bool)
	}(m)() {
		m_1 := func(hx_obj_420 map[string]any) func() *Xml {
			hx_field_421 := hx_obj_420["next"]
			if hx_field_421 == nil {
				var hx_zero_422 func() *Xml
				return hx_zero_422
			}
			return hx_field_421.(func() *Xml)
		}(m)()
		l.Push(self.__hx_this.xclassfield(m_1, false))
	}
	return l
}

func (self *haxe__rtti__XmlParser) xpath(x *Xml) map[string]any {
	path := self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("path")))
	params := hxrt.NewArray()
	c := x.__hx_this.elements()
	for func(hx_obj_424 map[string]any) func() bool {
		hx_field_425 := hx_obj_424["hasNext"]
		if hx_field_425 == nil {
			var hx_zero_426 func() bool
			return hx_zero_426
		}
		return hx_field_425.(func() bool)
	}(c)() {
		c_1 := func(hx_obj_427 map[string]any) func() *Xml {
			hx_field_428 := hx_obj_427["next"]
			if hx_field_428 == nil {
				var hx_zero_429 func() *Xml
				return hx_zero_429
			}
			return hx_field_428.(func() *Xml)
		}(c)()
		params.Push(self.__hx_this.xtype(c_1))
	}
	hx_obj_431 := map[string]any{}
	hx_obj_431["path"] = path
	hx_obj_431["params"] = params
	return hx_obj_431
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
	for func(hx_obj_432 map[string]any) func() bool {
		hx_field_433 := hx_obj_432["hasNext"]
		if hx_field_433 == nil {
			var hx_zero_434 func() bool
			return hx_zero_434
		}
		return hx_field_433.(func() bool)
	}(c)() {
		c_1 := func(hx_obj_435 map[string]any) func() *Xml {
			hx_field_436 := hx_obj_435["next"]
			if hx_field_436 == nil {
				var hx_zero_437 func() *Xml
				return hx_zero_437
			}
			return hx_field_436.(func() *Xml)
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
	hx_obj_442 := map[string]any{}
	hx_obj_442["file"] = x.__hx_this.get(hxrt.StringFromLiteral("file"))
	hx_obj_442["path"] = self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("path")))
	var hx_if_443 *string
	if x.__hx_this.exists(hxrt.StringFromLiteral("module")) {
		hx_if_443 = self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("module")))
	} else {
		hx_if_443 = nil
	}
	hx_obj_442["module"] = hx_if_443
	hx_obj_442["doc"] = doc
	hx_obj_442["isPrivate"] = x.__hx_this.exists(hxrt.StringFromLiteral("private"))
	hx_obj_442["isExtern"] = x.__hx_this.exists(hxrt.StringFromLiteral("extern"))
	hx_obj_442["isFinal"] = x.__hx_this.exists(hxrt.StringFromLiteral("final"))
	hx_obj_442["isInterface"] = isInterface
	hx_obj_442["params"] = self.__hx_this.mkTypeParams(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("params")))
	hx_obj_442["superClass"] = csuper
	hx_obj_442["interfaces"] = interfaces
	hx_obj_442["fields"] = fields
	hx_obj_442["statics"] = statics
	hx_obj_442["tdynamic"] = tdynamic
	hx_obj_442["platforms"] = self.__hx_this.defplat()
	hx_obj_442["meta"] = meta
	return hx_obj_442
}

func (self *haxe__rtti__XmlParser) xclassfield(x *Xml, defPublic bool) map[string]any {
	e := x.__hx_this.elements()
	t := self.__hx_this.xtype(func(hx_obj_444 map[string]any) func() *Xml {
		hx_field_445 := hx_obj_444["next"]
		if hx_field_445 == nil {
			var hx_zero_446 func() *Xml
			return hx_zero_446
		}
		return hx_field_445.(func() *Xml)
	}(e)())
	var doc *string = nil
	meta := hxrt.NewArray()
	var overloads *hxrt.Array = nil
	var line any = nil
	if x.__hx_this.exists(hxrt.StringFromLiteral("line")) {
		line = self.__hx_this.parseIntString(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("line")))
	}
	c := e
	for func(hx_obj_447 map[string]any) func() bool {
		hx_field_448 := hx_obj_447["hasNext"]
		if hx_field_448 == nil {
			var hx_zero_449 func() bool
			return hx_zero_449
		}
		return hx_field_448.(func() bool)
	}(c)() {
		c_1 := func(hx_obj_450 map[string]any) func() *Xml {
			hx_field_451 := hx_obj_450["next"]
			if hx_field_451 == nil {
				var hx_zero_452 func() *Xml
				return hx_zero_452
			}
			return hx_field_451.(func() *Xml)
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
	hx_obj_453 := map[string]any{}
	hx_obj_453["name"] = self.__hx_this.elementName(x)
	hx_obj_453["type"] = t
	hx_obj_453["isPublic"] = (x.__hx_this.exists(hxrt.StringFromLiteral("public")) || func(hx_value_454 any) bool {
		if hx_value_454 == nil {
			var hx_zero_455 bool
			return hx_zero_455
		}
		return hx_value_454.(bool)
	}(defPublic))
	hx_obj_453["isFinal"] = x.__hx_this.exists(hxrt.StringFromLiteral("final"))
	hx_obj_453["isOverride"] = x.__hx_this.exists(hxrt.StringFromLiteral("override"))
	hx_obj_453["line"] = func(hx_value_456 any) any {
		if hx_value_456 == nil {
			return nil
		}
		return hx_value_456.(int)
	}(line)
	hx_obj_453["doc"] = doc
	var hx_if_457 *haxe__rtti__Rights
	if x.__hx_this.exists(hxrt.StringFromLiteral("get")) {
		hx_if_457 = self.__hx_this.mkRights(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("get")))
	} else {
		hx_if_457 = haxe__rtti__Rights_RNormal
	}
	hx_obj_453["get"] = hx_if_457
	var hx_if_458 *haxe__rtti__Rights
	if x.__hx_this.exists(hxrt.StringFromLiteral("set")) {
		hx_if_458 = self.__hx_this.mkRights(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("set")))
	} else {
		hx_if_458 = haxe__rtti__Rights_RNormal
	}
	hx_obj_453["set"] = hx_if_458
	var hx_if_459 *hxrt.Array
	if x.__hx_this.exists(hxrt.StringFromLiteral("params")) {
		hx_if_459 = self.__hx_this.mkTypeParams(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("params")))
	} else {
		hx_if_459 = hxrt.NewArray()
	}
	hx_obj_453["params"] = hx_if_459
	hx_obj_453["platforms"] = self.__hx_this.defplat()
	hx_obj_453["meta"] = meta
	hx_obj_453["overloads"] = overloads
	var hx_if_460 *string
	if x.__hx_this.exists(hxrt.StringFromLiteral("expr")) {
		hx_if_460 = self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("expr"))
	} else {
		hx_if_460 = nil
	}
	hx_obj_453["expr"] = hx_if_460
	return hx_obj_453
}

func (self *haxe__rtti__XmlParser) xenum(x *Xml) map[string]any {
	cl := hxrt.NewArray()
	var doc *string = nil
	meta := hxrt.NewArray()
	c := x.__hx_this.elements()
	for func(hx_obj_461 map[string]any) func() bool {
		hx_field_462 := hx_obj_461["hasNext"]
		if hx_field_462 == nil {
			var hx_zero_463 func() bool
			return hx_zero_463
		}
		return hx_field_462.(func() bool)
	}(c)() {
		c_1 := func(hx_obj_464 map[string]any) func() *Xml {
			hx_field_465 := hx_obj_464["next"]
			if hx_field_465 == nil {
				var hx_zero_466 func() *Xml
				return hx_zero_466
			}
			return hx_field_465.(func() *Xml)
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
	hx_obj_468 := map[string]any{}
	hx_obj_468["file"] = x.__hx_this.get(hxrt.StringFromLiteral("file"))
	hx_obj_468["path"] = self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("path")))
	var hx_if_469 *string
	if x.__hx_this.exists(hxrt.StringFromLiteral("module")) {
		hx_if_469 = self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("module")))
	} else {
		hx_if_469 = nil
	}
	hx_obj_468["module"] = hx_if_469
	hx_obj_468["doc"] = doc
	hx_obj_468["isPrivate"] = x.__hx_this.exists(hxrt.StringFromLiteral("private"))
	hx_obj_468["isExtern"] = x.__hx_this.exists(hxrt.StringFromLiteral("extern"))
	hx_obj_468["params"] = self.__hx_this.mkTypeParams(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("params")))
	hx_obj_468["constructors"] = cl
	hx_obj_468["platforms"] = self.__hx_this.defplat()
	hx_obj_468["meta"] = meta
	return hx_obj_468
}

func (self *haxe__rtti__XmlParser) xenumfield(x *Xml) map[string]any {
	var args *hxrt.Array = nil
	docElements := x.__hx_this.elementsNamed(hxrt.StringFromLiteral("haxe_doc"))
	var hx_if_476 *Xml
	if func(hx_obj_470 map[string]any) func() bool {
		hx_field_471 := hx_obj_470["hasNext"]
		if hx_field_471 == nil {
			var hx_zero_472 func() bool
			return hx_zero_472
		}
		return hx_field_471.(func() bool)
	}(docElements)() {
		hx_if_476 = func(hx_obj_473 map[string]any) func() *Xml {
			hx_field_474 := hx_obj_473["next"]
			if hx_field_474 == nil {
				var hx_zero_475 func() *Xml
				return hx_zero_475
			}
			return hx_field_474.(func() *Xml)
		}(docElements)()
	} else {
		hx_if_476 = nil
	}
	xdoc := hx_if_476
	var hx_if_477 *hxrt.Array
	if self.__hx_this.hasNamedElement(x, hxrt.StringFromLiteral("meta")) {
		hx_if_477 = self.__hx_this.xmeta(self.__hx_this.requireNamedElement(x, hxrt.StringFromLiteral("meta")))
	} else {
		hx_if_477 = hxrt.NewArray()
	}
	meta := hx_if_477
	if x.__hx_this.exists(hxrt.StringFromLiteral("a")) {
		names := self.__hx_this.splitString(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("a")), hxrt.StringFromLiteral(":"))
		elts := x.__hx_this.elements()
		args = hxrt.NewArray()
		_g := 0
		for _g < names.Len() {
			c := func(hx_value_478 any) *string {
				if hx_value_478 == nil {
					var hx_zero_479 *string
					return hx_zero_479
				}
				return hx_value_478.(*string)
			}(names.Get(_g))
			_g = int(int32((_g + 1)))
			opt := false
			if hxrt.StringEqualStringPtr(hxrt.StringCharAtStringPtr(c, 0), hxrt.StringFromLiteral("?")) {
				opt = true
				c = hxrt.StringSubstrStringPtr(c, 1, 0, false)
			}
			hx_obj_481 := map[string]any{}
			hx_obj_481["name"] = c
			hx_obj_481["opt"] = opt
			hx_obj_481["t"] = self.__hx_this.xtype(func(hx_obj_482 map[string]any) func() *Xml {
				hx_field_483 := hx_obj_482["next"]
				if hx_field_483 == nil {
					var hx_zero_484 func() *Xml
					return hx_zero_484
				}
				return hx_field_483.(func() *Xml)
			}(elts)())
			args.Push(hx_obj_481)
		}
	}
	hx_obj_485 := map[string]any{}
	hx_obj_485["name"] = self.__hx_this.elementName(x)
	hx_obj_485["args"] = args
	var hx_if_486 *string
	if xdoc == nil {
		hx_if_486 = nil
	} else {
		hx_if_486 = self.__hx_this.innerData(xdoc)
	}
	hx_obj_485["doc"] = hx_if_486
	hx_obj_485["meta"] = meta
	hx_obj_485["platforms"] = self.__hx_this.defplat()
	return hx_obj_485
}

func (self *haxe__rtti__XmlParser) xabstract(x *Xml) map[string]any {
	var doc *string = nil
	var impl map[string]any = nil
	var athis *haxe__rtti__CType = nil
	meta := hxrt.NewArray()
	to := hxrt.NewArray()
	from := hxrt.NewArray()
	c := x.__hx_this.elements()
	for func(hx_obj_487 map[string]any) func() bool {
		hx_field_488 := hx_obj_487["hasNext"]
		if hx_field_488 == nil {
			var hx_zero_489 func() bool
			return hx_zero_489
		}
		return hx_field_488.(func() bool)
	}(c)() {
		c_1 := func(hx_obj_490 map[string]any) func() *Xml {
			hx_field_491 := hx_obj_490["next"]
			if hx_field_491 == nil {
				var hx_zero_492 func() *Xml
				return hx_zero_492
			}
			return hx_field_491.(func() *Xml)
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
					for func(hx_obj_493 map[string]any) func() bool {
						hx_field_494 := hx_obj_493["hasNext"]
						if hx_field_494 == nil {
							var hx_zero_495 func() bool
							return hx_zero_495
						}
						return hx_field_494.(func() bool)
					}(t)() {
						t_1 := func(hx_obj_496 map[string]any) func() *Xml {
							hx_field_497 := hx_obj_496["next"]
							if hx_field_497 == nil {
								var hx_zero_498 func() *Xml
								return hx_zero_498
							}
							return hx_field_497.(func() *Xml)
						}(t)()
						hx_obj_500 := map[string]any{}
						hx_obj_500["t"] = self.__hx_this.xtype(self.__hx_this.requireFirstElement(t_1))
						hx_obj_500["field"] = t_1.__hx_this.get(hxrt.StringFromLiteral("field"))
						to.Push(hx_obj_500)
					}
				} else {
					if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("from")) {
						t_2 := c_1.__hx_this.elements()
						for func(hx_obj_501 map[string]any) func() bool {
							hx_field_502 := hx_obj_501["hasNext"]
							if hx_field_502 == nil {
								var hx_zero_503 func() bool
								return hx_zero_503
							}
							return hx_field_502.(func() bool)
						}(t_2)() {
							t_3 := func(hx_obj_504 map[string]any) func() *Xml {
								hx_field_505 := hx_obj_504["next"]
								if hx_field_505 == nil {
									var hx_zero_506 func() *Xml
									return hx_zero_506
								}
								return hx_field_505.(func() *Xml)
							}(t_2)()
							hx_obj_508 := map[string]any{}
							hx_obj_508["t"] = self.__hx_this.xtype(self.__hx_this.requireFirstElement(t_3))
							hx_obj_508["field"] = t_3.__hx_this.get(hxrt.StringFromLiteral("field"))
							from.Push(hx_obj_508)
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
	hx_obj_509 := map[string]any{}
	hx_obj_509["file"] = x.__hx_this.get(hxrt.StringFromLiteral("file"))
	hx_obj_509["path"] = self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("path")))
	var hx_if_510 *string
	if x.__hx_this.exists(hxrt.StringFromLiteral("module")) {
		hx_if_510 = self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("module")))
	} else {
		hx_if_510 = nil
	}
	hx_obj_509["module"] = hx_if_510
	hx_obj_509["doc"] = doc
	hx_obj_509["isPrivate"] = x.__hx_this.exists(hxrt.StringFromLiteral("private"))
	hx_obj_509["params"] = self.__hx_this.mkTypeParams(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("params")))
	hx_obj_509["platforms"] = self.__hx_this.defplat()
	hx_obj_509["meta"] = meta
	hx_obj_509["athis"] = athis
	hx_obj_509["to"] = to
	hx_obj_509["from"] = from
	hx_obj_509["impl"] = impl
	return hx_obj_509
}

func (self *haxe__rtti__XmlParser) xtypedef(x *Xml) map[string]any {
	var doc *string = nil
	var t *haxe__rtti__CType = nil
	meta := hxrt.NewArray()
	c := x.__hx_this.elements()
	for func(hx_obj_511 map[string]any) func() bool {
		hx_field_512 := hx_obj_511["hasNext"]
		if hx_field_512 == nil {
			var hx_zero_513 func() bool
			return hx_zero_513
		}
		return hx_field_512.(func() bool)
	}(c)() {
		c_1 := func(hx_obj_514 map[string]any) func() *Xml {
			hx_field_515 := hx_obj_514["next"]
			if hx_field_515 == nil {
				var hx_zero_516 func() *Xml
				return hx_zero_516
			}
			return hx_field_515.(func() *Xml)
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
	hx_obj_517 := map[string]any{}
	hx_obj_517["file"] = x.__hx_this.get(hxrt.StringFromLiteral("file"))
	hx_obj_517["path"] = self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("path")))
	var hx_if_518 *string
	if x.__hx_this.exists(hxrt.StringFromLiteral("module")) {
		hx_if_518 = self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("module")))
	} else {
		hx_if_518 = nil
	}
	hx_obj_517["module"] = hx_if_518
	hx_obj_517["doc"] = doc
	hx_obj_517["isPrivate"] = x.__hx_this.exists(hxrt.StringFromLiteral("private"))
	hx_obj_517["params"] = self.__hx_this.mkTypeParams(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("params")))
	hx_obj_517["type"] = t
	hx_obj_517["types"] = types
	hx_obj_517["platforms"] = self.__hx_this.defplat()
	hx_obj_517["meta"] = meta
	return hx_obj_517
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
		var hx_if_519 *hxrt.Array
		if x.__hx_this.exists(hxrt.StringFromLiteral("v")) {
			hx_if_519 = self.__hx_this.splitString(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("v")), hxrt.StringFromLiteral(":"))
		} else {
			hx_if_519 = nil
		}
		evalues := hx_if_519
		valueIndex := 0
		e := x.__hx_this.elements()
		for func(hx_obj_520 map[string]any) func() bool {
			hx_field_521 := hx_obj_520["hasNext"]
			if hx_field_521 == nil {
				var hx_zero_522 func() bool
				return hx_zero_522
			}
			return hx_field_521.(func() bool)
		}(e)() {
			e_1 := func(hx_obj_523 map[string]any) func() *Xml {
				hx_field_524 := hx_obj_523["next"]
				if hx_field_524 == nil {
					var hx_zero_525 func() *Xml
					return hx_zero_525
				}
				return hx_field_524.(func() *Xml)
			}(e)()
			opt := false
			var hx_if_528 *string
			if argIndex < aname.Len() {
				hx_if_528 = hxrt.StdString(func(hx_value_526 any) *string {
					if hx_value_526 == nil {
						var hx_zero_527 *string
						return hx_zero_527
					}
					return hx_value_526.(*string)
				}(aname.Get(argIndex)))
			} else {
				hx_if_528 = nil
			}
			a := hx_if_528
			argIndex = int(int32((argIndex + 1)))
			if hxrt.StringEqualStringPtr(a, nil) {
				a = hxrt.StringFromLiteral("")
			}
			if hxrt.StringEqualStringPtr(hxrt.StringCharAtStringPtr(a, 0), hxrt.StringFromLiteral("?")) {
				opt = true
				a = hxrt.StringSubstrStringPtr(a, 1, 0, false)
			}
			var hx_if_532 *string
			if (evalues == nil) || (valueIndex >= evalues.Len()) {
				hx_if_532 = nil
			} else {
				hx_post_529 := valueIndex
				valueIndex = int(int32((valueIndex + 1)))
				hx_if_532 = hxrt.StdString(func(hx_value_530 any) *string {
					if hx_value_530 == nil {
						var hx_zero_531 *string
						return hx_zero_531
					}
					return hx_value_530.(*string)
				}(evalues.Get(hx_post_529)))
			}
			v := hx_if_532
			hx_obj_534 := map[string]any{}
			hx_obj_534["name"] = a
			hx_obj_534["opt"] = opt
			hx_obj_534["t"] = self.__hx_this.xtype(e_1)
			var hx_if_535 *string
			if hxrt.StringEqualStringPtr(v, hxrt.StringFromLiteral("")) {
				hx_if_535 = nil
			} else {
				hx_if_535 = v
			}
			hx_obj_534["value"] = hx_if_535
			args.Push(hx_obj_534)
		}
		ret := func(hx_value_536 any) map[string]any {
			if hx_value_536 == nil {
				var hx_zero_537 map[string]any
				return hx_zero_537
			}
			return hx_value_536.(map[string]any)
		}(args.Get(int((hxrt.Int32Wrap(args.Len()) - hxrt.Int32Wrap(1)))))
		callArgs := hxrt.NewArray()
		_g := 0
		_g1 := int((hxrt.Int32Wrap(args.Len()) - hxrt.Int32Wrap(1)))
		for _g < _g1 {
			hx_post_538 := _g
			_g = int(int32((_g + 1)))
			i := hx_post_538
			callArgs.Push(args.Get(i))
		}
		return haxe__rtti__CType_CFunction(callArgs, func(hx_obj_540 map[string]any) *haxe__rtti__CType {
			hx_field_541 := hx_obj_540["t"]
			if hx_field_541 == nil {
				var hx_zero_542 *haxe__rtti__CType
				return hx_zero_542
			}
			return hx_field_541.(*haxe__rtti__CType)
		}(ret))
	}
	if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("a")) {
		fields := hxrt.NewArray()
		f := x.__hx_this.elements()
		for func(hx_obj_543 map[string]any) func() bool {
			hx_field_544 := hx_obj_543["hasNext"]
			if hx_field_544 == nil {
				var hx_zero_545 func() bool
				return hx_zero_545
			}
			return hx_field_544.(func() bool)
		}(f)() {
			f_1 := func(hx_obj_546 map[string]any) func() *Xml {
				hx_field_547 := hx_obj_546["next"]
				if hx_field_547 == nil {
					var hx_zero_548 func() *Xml
					return hx_zero_548
				}
				return hx_field_547.(func() *Xml)
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
	var hx_throw_zero_550 *haxe__rtti__CType
	return hx_throw_zero_550
}

func (self *haxe__rtti__XmlParser) xtypeparams(x *Xml) *hxrt.Array {
	p := hxrt.NewArray()
	c := x.__hx_this.elements()
	for func(hx_obj_551 map[string]any) func() bool {
		hx_field_552 := hx_obj_551["hasNext"]
		if hx_field_552 == nil {
			var hx_zero_553 func() bool
			return hx_zero_553
		}
		return hx_field_552.(func() bool)
	}(c)() {
		c_1 := func(hx_obj_554 map[string]any) func() *Xml {
			hx_field_555 := hx_obj_554["next"]
			if hx_field_555 == nil {
				var hx_zero_556 func() *Xml
				return hx_zero_556
			}
			return hx_field_555.(func() *Xml)
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
		hx_post_559 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_559
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
		parts.Push(hxrt.StringSubstrStringPtr(value, start, int((hxrt.Int32Wrap(index) - hxrt.Int32Wrap(start))), true))
		start = int((hxrt.Int32Wrap(index) + hxrt.Int32Wrap(hxrt.StringLengthStringPtr(separator))))
	}
	return parts
}

func (self *haxe__rtti__XmlParser) findSeparator(value *string, separator *string, start int) int {
	limit := int((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(value)) - hxrt.Int32Wrap(hxrt.StringLengthStringPtr(separator))))
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
	var hx_if_562 *string
	if hxrt.StringEqualStringPtr(value, nil) {
		hx_if_562 = hxrt.StringFromLiteral("")
	} else {
		hx_if_562 = value
	}
	return hx_if_562
}

func (self *haxe__rtti__XmlParser) hasNamedElement(x *Xml, name *string) bool {
	return func(hx_obj_563 map[string]any) func() bool {
		hx_field_564 := hx_obj_563["hasNext"]
		if hx_field_564 == nil {
			var hx_zero_565 func() bool
			return hx_zero_565
		}
		return hx_field_564.(func() bool)
	}(x.__hx_this.elementsNamed(name))()
}

func (self *haxe__rtti__XmlParser) requireNamedElement(x *Xml, name *string) *Xml {
	elements := x.__hx_this.elementsNamed(name)
	if !func(hx_obj_566 map[string]any) func() bool {
		hx_field_567 := hx_obj_566["hasNext"]
		if hx_field_567 == nil {
			var hx_zero_568 func() bool
			return hx_zero_568
		}
		return hx_field_567.(func() bool)
	}(elements)() {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(self.__hx_this.nodeDisplayName(x), hxrt.StringFromLiteral(" is missing element ")), name))
	}
	return func(hx_obj_569 map[string]any) func() *Xml {
		hx_field_570 := hx_obj_569["next"]
		if hx_field_570 == nil {
			var hx_zero_571 func() *Xml
			return hx_zero_571
		}
		return hx_field_570.(func() *Xml)
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
	var hx_if_572 *string
	if hxrt.HaxeEqual(x.nodeType, Xml_Document) {
		hx_if_572 = hxrt.StringFromLiteral("Document")
	} else {
		hx_if_572 = self.__hx_this.elementName(x)
	}
	return hx_if_572
}

func (self *haxe__rtti__XmlParser) elementName(x *Xml) *string {
	if !hxrt.HaxeEqual(x.nodeType, Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(x.nodeType))))
	}
	name := x.nodeName
	var hx_if_573 *string
	if hxrt.StringEqualStringPtr(name, nil) {
		hx_if_573 = hxrt.StringFromLiteral("")
	} else {
		hx_if_573 = name
	}
	return hx_if_573
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
	value := func(hx_value_575 any) *Xml {
		if hx_value_575 == nil {
			var hx_zero_576 *Xml
			return hx_zero_576
		}
		return hx_value_575.(*Xml)
	}(it_array.Get(func() int {
		hx_post_574 := it_current
		it_current = int(int32((it_current + 1)))
		return hx_post_574
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
		hx_post_577 := _g_current
		_g_current = int(int32((_g_current + 1)))
		child := func(hx_value_578 any) *Xml {
			if hx_value_578 == nil {
				var hx_zero_579 *Xml
				return hx_zero_579
			}
			return hx_value_578.(*Xml)
		}(_g_array.Get(hx_post_577))
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
		result = int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(result) * hxrt.Int32Wrap(10)))) + hxrt.Int32Wrap(int((hxrt.Int32Wrap(code) - hxrt.Int32Wrap(zeroCode))))))
		index = int(int32((index + 1)))
	}
	var hx_if_580 int
	if negative {
		hx_if_580 = int(-int32(result))
	} else {
		hx_if_580 = result
	}
	return hx_if_580
}
