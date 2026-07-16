package main

import "snapshot/hxrt"

type I_haxe__rtti__XmlParser interface {
	sort(l []*haxe__rtti__TypeTree)
	sortFields(a []map[string]any)
	process(x *Xml, platform *string)
	mergeRights(f1 map[string]any, f2 map[string]any) bool
	mergeDoc(f1 map[string]any, f2 map[string]any) bool
	mergeFields(f map[string]any, f2 map[string]any) bool
	newField(c map[string]any, f map[string]any)
	mergeClasses(c map[string]any, c2 map[string]any) bool
	mergeEnums(e map[string]any, e2 map[string]any) bool
	mergeTypedefs(t map[string]any, t2 map[string]any) bool
	mergeAbstracts(a map[string]any, a2 map[string]any) bool
	merge(t *haxe__rtti__TypeTree)
	mkPath(p *string) *string
	mkTypeParams(p *string) []*string
	mkRights(r *string) *haxe__rtti__Rights
	xroot(x *Xml)
	processElement(x *Xml) *haxe__rtti__TypeTree
	xmeta(x *Xml) []map[string]any
	xoverloads(x *Xml) []map[string]any
	xpath(x *Xml) map[string]any
	xclass(x *Xml) map[string]any
	xclassfield(x *Xml, defPublic bool) map[string]any
	xenum(x *Xml) map[string]any
	xenumfield(x *Xml) map[string]any
	xabstract(x *Xml) map[string]any
	xtypedef(x *Xml) map[string]any
	xtype(x *Xml) *haxe__rtti__CType
	xtypeparams(x *Xml) []*haxe__rtti__CType
	defplat() []*string
	joinStringArray(values []*string, separator *string) *string
	splitString(value *string, separator *string) []*string
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
	root        []*haxe__rtti__TypeTree
	curplatform *string
}

func New_haxe__rtti__XmlParser() *haxe__rtti__XmlParser {
	self := &haxe__rtti__XmlParser{}
	self.__hx_this = self
	self.root = []*haxe__rtti__TypeTree{}
	return self
}

func (self *haxe__rtti__XmlParser) sort(l []*haxe__rtti__TypeTree) {
	if l == nil {
		l = self.root
	}
	func(hx_sort_src_22 []*haxe__rtti__TypeTree) {
		hx_sort_raw_21 := func(hx_sort_src_23 []*haxe__rtti__TypeTree) []any {
			hx_sort_out_25 := make([]any, 0, len(hx_sort_src_23))
			for _, hx_sort_item_24 := range hx_sort_src_23 {
				hx_sort_out_25 = append(hx_sort_out_25, hx_sort_item_24)
			}
			return hx_sort_out_25
		}(hx_sort_src_22)
		haxe__ds__ArraySort_sort(hx_sort_raw_21, func(hx_cmp_left_26 any, hx_cmp_right_27 any) int {
			return func(e1 *haxe__rtti__TypeTree, e2 *haxe__rtti__TypeTree) int {
				var hx_if_16 *string
				if e1.tag == 0 {
					_g := e1.params[0].(*string)
					_g1 := e1.params[1].(*string)
					_ = _g1
					_g1_1 := e1.params[2].([]*haxe__rtti__TypeTree)
					_ = _g1_1
					p := _g
					hx_if_16 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(" "), p)
				} else {
					hx_if_16 = func(hx_obj_13 map[string]any) *string {
						hx_field_14 := hx_obj_13["path"]
						if hx_field_14 == nil {
							var hx_zero_15 *string
							return hx_zero_15
						}
						return hx_field_14.(*string)
					}(haxe__rtti__TypeApi_typeInfos(e1))
				}
				n1 := hx_if_16
				var hx_if_20 *string
				if e2.tag == 0 {
					_g_1 := e2.params[0].(*string)
					_g1_2 := e2.params[1].(*string)
					_ = _g1_2
					_g1_3 := e2.params[2].([]*haxe__rtti__TypeTree)
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
			}(func(hx_value_28 any) *haxe__rtti__TypeTree {
				if hx_value_28 == nil {
					var hx_zero_29 *haxe__rtti__TypeTree
					return hx_zero_29
				}
				return hx_value_28.(*haxe__rtti__TypeTree)
			}(hx_cmp_left_26), func(hx_value_30 any) *haxe__rtti__TypeTree {
				if hx_value_30 == nil {
					var hx_zero_31 *haxe__rtti__TypeTree
					return hx_zero_31
				}
				return hx_value_30.(*haxe__rtti__TypeTree)
			}(hx_cmp_right_27))
		})
		func(hx_sort_raw_32 []any, hx_sort_dst_33 []*haxe__rtti__TypeTree) {
			for hx_sort_i_34, hx_sort_item_35 := range hx_sort_raw_32 {
				hx_sort_dst_33[hx_sort_i_34] = func(hx_value_36 any) *haxe__rtti__TypeTree {
					if hx_value_36 == nil {
						var hx_zero_37 *haxe__rtti__TypeTree
						return hx_zero_37
					}
					return hx_value_36.(*haxe__rtti__TypeTree)
				}(hx_sort_item_35)
			}
		}(hx_sort_raw_21, hx_sort_src_22)
	}(l)
	_g := 0
	for _g < len(l) {
		x := l[_g]
		_g = int(int32((_g + 1)))
		switch x.tag {
		case 0:
			_g_1 := x.params[0].(*string)
			_ = _g_1
			_g_2 := x.params[1].(*string)
			_ = _g_2
			_g_3 := x.params[2].([]*haxe__rtti__TypeTree)
			l_1 := _g_3
			self.sort(l_1)
		case 1:
			_g_4 := x.params[0].(map[string]any)
			c := _g_4
			self.sortFields(func(hx_obj_38 map[string]any) []map[string]any {
				hx_field_39 := hx_obj_38["fields"]
				if hx_field_39 == nil {
					var hx_zero_40 []map[string]any
					return hx_zero_40
				}
				return hx_field_39.([]map[string]any)
			}(c))
			self.sortFields(func(hx_obj_41 map[string]any) []map[string]any {
				hx_field_42 := hx_obj_41["statics"]
				if hx_field_42 == nil {
					var hx_zero_43 []map[string]any
					return hx_zero_43
				}
				return hx_field_42.([]map[string]any)
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

func (self *haxe__rtti__XmlParser) sortFields(a []map[string]any) {
	func(hx_sort_src_63 []map[string]any) {
		hx_sort_raw_62 := func(hx_sort_src_64 []map[string]any) []any {
			hx_sort_out_66 := make([]any, 0, len(hx_sort_src_64))
			for _, hx_sort_item_65 := range hx_sort_src_64 {
				hx_sort_out_66 = append(hx_sort_out_66, hx_sort_item_65)
			}
			return hx_sort_out_66
		}(hx_sort_src_63)
		haxe__ds__ArraySort_sort(hx_sort_raw_62, func(hx_cmp_left_67 any, hx_cmp_right_68 any) int {
			return func(f1 map[string]any, f2 map[string]any) int {
				v1 := haxe__rtti__TypeApi_isVar(func(hx_obj_44 map[string]any) *haxe__rtti__CType {
					hx_field_45 := hx_obj_44["type"]
					if hx_field_45 == nil {
						var hx_zero_46 *haxe__rtti__CType
						return hx_zero_46
					}
					return hx_field_45.(*haxe__rtti__CType)
				}(f1))
				v2 := haxe__rtti__TypeApi_isVar(func(hx_obj_47 map[string]any) *haxe__rtti__CType {
					hx_field_48 := hx_obj_47["type"]
					if hx_field_48 == nil {
						var hx_zero_49 *haxe__rtti__CType
						return hx_zero_49
					}
					return hx_field_48.(*haxe__rtti__CType)
				}(f2))
				if v1 && !v2 {
					return -1
				}
				if v2 && !v1 {
					return 1
				}
				if hxrt.StringEqualStringPtr(func(hx_obj_50 map[string]any) *string {
					hx_field_51 := hx_obj_50["name"]
					if hx_field_51 == nil {
						var hx_zero_52 *string
						return hx_zero_52
					}
					return hx_field_51.(*string)
				}(f1), hxrt.StringFromLiteral("new")) {
					return -1
				}
				if hxrt.StringEqualStringPtr(func(hx_obj_53 map[string]any) *string {
					hx_field_54 := hx_obj_53["name"]
					if hx_field_54 == nil {
						var hx_zero_55 *string
						return hx_zero_55
					}
					return hx_field_54.(*string)
				}(f2), hxrt.StringFromLiteral("new")) {
					return 1
				}
				return Reflect_compare(func(hx_obj_56 map[string]any) *string {
					hx_field_57 := hx_obj_56["name"]
					if hx_field_57 == nil {
						var hx_zero_58 *string
						return hx_zero_58
					}
					return hx_field_57.(*string)
				}(f1), func(hx_obj_59 map[string]any) *string {
					hx_field_60 := hx_obj_59["name"]
					if hx_field_60 == nil {
						var hx_zero_61 *string
						return hx_zero_61
					}
					return hx_field_60.(*string)
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
		func(hx_sort_raw_73 []any, hx_sort_dst_74 []map[string]any) {
			for hx_sort_i_75, hx_sort_item_76 := range hx_sort_raw_73 {
				hx_sort_dst_74[hx_sort_i_75] = func(hx_value_77 any) map[string]any {
					if hx_value_77 == nil {
						var hx_zero_78 map[string]any
						return hx_zero_78
					}
					return hx_value_77.(map[string]any)
				}(hx_sort_item_76)
			}
		}(hx_sort_raw_62, hx_sort_src_63)
	}(a)
}

func (self *haxe__rtti__XmlParser) process(x *Xml, platform *string) {
	self.curplatform = platform
	self.xroot(x)
}

func (self *haxe__rtti__XmlParser) mergeRights(f1 map[string]any, f2 map[string]any) bool {
	if (((func(hx_obj_79 map[string]any) *haxe__rtti__Rights {
		hx_field_80 := hx_obj_79["get"]
		if hx_field_80 == nil {
			var hx_zero_81 *haxe__rtti__Rights
			return hx_zero_81
		}
		return hx_field_80.(*haxe__rtti__Rights)
	}(f1) == haxe__rtti__Rights_RInline) && (func(hx_obj_82 map[string]any) *haxe__rtti__Rights {
		hx_field_83 := hx_obj_82["set"]
		if hx_field_83 == nil {
			var hx_zero_84 *haxe__rtti__Rights
			return hx_zero_84
		}
		return hx_field_83.(*haxe__rtti__Rights)
	}(f1) == haxe__rtti__Rights_RNo)) && (func(hx_obj_85 map[string]any) *haxe__rtti__Rights {
		hx_field_86 := hx_obj_85["get"]
		if hx_field_86 == nil {
			var hx_zero_87 *haxe__rtti__Rights
			return hx_zero_87
		}
		return hx_field_86.(*haxe__rtti__Rights)
	}(f2) == haxe__rtti__Rights_RNormal)) && (func(hx_obj_88 map[string]any) *haxe__rtti__Rights {
		hx_field_89 := hx_obj_88["set"]
		if hx_field_89 == nil {
			var hx_zero_90 *haxe__rtti__Rights
			return hx_zero_90
		}
		return hx_field_89.(*haxe__rtti__Rights)
	}(f2) == haxe__rtti__Rights_RMethod) {
		f1["get"] = haxe__rtti__Rights_RNormal
		f1["set"] = haxe__rtti__Rights_RMethod
		return true
	}
	return (Type_enumEq(func(hx_obj_91 map[string]any) *haxe__rtti__Rights {
		hx_field_92 := hx_obj_91["get"]
		if hx_field_92 == nil {
			var hx_zero_93 *haxe__rtti__Rights
			return hx_zero_93
		}
		return hx_field_92.(*haxe__rtti__Rights)
	}(f1), func(hx_obj_94 map[string]any) *haxe__rtti__Rights {
		hx_field_95 := hx_obj_94["get"]
		if hx_field_95 == nil {
			var hx_zero_96 *haxe__rtti__Rights
			return hx_zero_96
		}
		return hx_field_95.(*haxe__rtti__Rights)
	}(f2)) && Type_enumEq(func(hx_obj_97 map[string]any) *haxe__rtti__Rights {
		hx_field_98 := hx_obj_97["set"]
		if hx_field_98 == nil {
			var hx_zero_99 *haxe__rtti__Rights
			return hx_zero_99
		}
		return hx_field_98.(*haxe__rtti__Rights)
	}(f1), func(hx_obj_100 map[string]any) *haxe__rtti__Rights {
		hx_field_101 := hx_obj_100["set"]
		if hx_field_101 == nil {
			var hx_zero_102 *haxe__rtti__Rights
			return hx_zero_102
		}
		return hx_field_101.(*haxe__rtti__Rights)
	}(f2)))
}

func (self *haxe__rtti__XmlParser) mergeDoc(f1 map[string]any, f2 map[string]any) bool {
	if hxrt.StringEqualStringPtr(func(hx_obj_112 map[string]any) *string {
		hx_field_113 := hx_obj_112["doc"]
		if hx_field_113 == nil {
			var hx_zero_114 *string
			return hx_zero_114
		}
		return hx_field_113.(*string)
	}(f1), nil) {
		f1["doc"] = func(hx_obj_103 map[string]any) *string {
			hx_field_104 := hx_obj_103["doc"]
			if hx_field_104 == nil {
				var hx_zero_105 *string
				return hx_zero_105
			}
			return hx_field_104.(*string)
		}(f2)
	} else {
		if hxrt.StringEqualStringPtr(func(hx_obj_109 map[string]any) *string {
			hx_field_110 := hx_obj_109["doc"]
			if hx_field_110 == nil {
				var hx_zero_111 *string
				return hx_zero_111
			}
			return hx_field_110.(*string)
		}(f2), nil) {
			f2["doc"] = func(hx_obj_106 map[string]any) *string {
				hx_field_107 := hx_obj_106["doc"]
				if hx_field_107 == nil {
					var hx_zero_108 *string
					return hx_zero_108
				}
				return hx_field_107.(*string)
			}(f1)
		}
	}
	return true
}

func (self *haxe__rtti__XmlParser) mergeFields(f map[string]any, f2 map[string]any) bool {
	return (haxe__rtti__TypeApi_fieldEq(f, f2) || (((hxrt.StringEqualStringPtr(func(hx_obj_115 map[string]any) *string {
		hx_field_116 := hx_obj_115["name"]
		if hx_field_116 == nil {
			var hx_zero_117 *string
			return hx_zero_117
		}
		return hx_field_116.(*string)
	}(f), func(hx_obj_118 map[string]any) *string {
		hx_field_119 := hx_obj_118["name"]
		if hx_field_119 == nil {
			var hx_zero_120 *string
			return hx_zero_120
		}
		return hx_field_119.(*string)
	}(f2)) && (self.mergeRights(f, f2) || self.mergeRights(f2, f))) && self.mergeDoc(f, f2)) && haxe__rtti__TypeApi_fieldEq(f, f2)))
}

func (self *haxe__rtti__XmlParser) newField(c map[string]any, f map[string]any) {
}

func (self *haxe__rtti__XmlParser) mergeClasses(c map[string]any, c2 map[string]any) bool {
	if func(hx_obj_121 map[string]any) bool {
		hx_field_122 := hx_obj_121["isInterface"]
		if hx_field_122 == nil {
			var hx_zero_123 bool
			return hx_zero_123
		}
		return hx_field_122.(bool)
	}(c) != func(hx_obj_124 map[string]any) bool {
		hx_field_125 := hx_obj_124["isInterface"]
		if hx_field_125 == nil {
			var hx_zero_126 bool
			return hx_zero_126
		}
		return hx_field_125.(bool)
	}(c2) {
		return false
	}
	if !hxrt.StringEqualStringPtr(self.curplatform, nil) {
		hx_obj_128 := c
		hx_arr_127 := func(hx_obj_129 map[string]any) []*string {
			hx_field_130 := hx_obj_129["platforms"]
			if hx_field_130 == nil {
				var hx_zero_131 []*string
				return hx_zero_131
			}
			return hx_field_130.([]*string)
		}(c)
		hx_arr_127 = append(hx_arr_127, self.curplatform)
		hx_obj_128["platforms"] = hx_arr_127
	}
	if func(hx_obj_132 map[string]any) bool {
		hx_field_133 := hx_obj_132["isExtern"]
		if hx_field_133 == nil {
			var hx_zero_134 bool
			return hx_zero_134
		}
		return hx_field_133.(bool)
	}(c) != func(hx_obj_135 map[string]any) bool {
		hx_field_136 := hx_obj_135["isExtern"]
		if hx_field_136 == nil {
			var hx_zero_137 bool
			return hx_zero_137
		}
		return hx_field_136.(bool)
	}(c2) {
		c["isExtern"] = false
	}
	_g := 0
	_g1 := func(hx_obj_138 map[string]any) []map[string]any {
		hx_field_139 := hx_obj_138["fields"]
		if hx_field_139 == nil {
			var hx_zero_140 []map[string]any
			return hx_zero_140
		}
		return hx_field_139.([]map[string]any)
	}(c2)
	for _g < len(_g1) {
		f2 := _g1[_g]
		_g = int(int32((_g + 1)))
		var found map[string]any = nil
		_g_1 := 0
		_g1_1 := func(hx_obj_141 map[string]any) []map[string]any {
			hx_field_142 := hx_obj_141["fields"]
			if hx_field_142 == nil {
				var hx_zero_143 []map[string]any
				return hx_zero_143
			}
			return hx_field_142.([]map[string]any)
		}(c)
		for _g_1 < len(_g1_1) {
			f := _g1_1[_g_1]
			_g_1 = int(int32((_g_1 + 1)))
			if self.mergeFields(f, f2) {
				found = f
				break
			}
		}
		if found == nil {
			self.newField(c, f2)
			hx_obj_145 := c
			hx_arr_144 := func(hx_obj_146 map[string]any) []map[string]any {
				hx_field_147 := hx_obj_146["fields"]
				if hx_field_147 == nil {
					var hx_zero_148 []map[string]any
					return hx_zero_148
				}
				return hx_field_147.([]map[string]any)
			}(c)
			hx_arr_144 = append(hx_arr_144, f2)
			hx_obj_145["fields"] = hx_arr_144
		} else {
			if !hxrt.StringEqualStringPtr(self.curplatform, nil) {
				hx_obj_150 := found
				hx_arr_149 := func(hx_obj_151 map[string]any) []*string {
					hx_field_152 := hx_obj_151["platforms"]
					if hx_field_152 == nil {
						var hx_zero_153 []*string
						return hx_zero_153
					}
					return hx_field_152.([]*string)
				}(found)
				hx_arr_149 = append(hx_arr_149, self.curplatform)
				hx_obj_150["platforms"] = hx_arr_149
			}
		}
	}
	_g_2 := 0
	_g1_2 := func(hx_obj_154 map[string]any) []map[string]any {
		hx_field_155 := hx_obj_154["statics"]
		if hx_field_155 == nil {
			var hx_zero_156 []map[string]any
			return hx_zero_156
		}
		return hx_field_155.([]map[string]any)
	}(c2)
	for _g_2 < len(_g1_2) {
		f2_1 := _g1_2[_g_2]
		_g_2 = int(int32((_g_2 + 1)))
		var found_1 map[string]any = nil
		_g_3 := 0
		_g1_3 := func(hx_obj_157 map[string]any) []map[string]any {
			hx_field_158 := hx_obj_157["statics"]
			if hx_field_158 == nil {
				var hx_zero_159 []map[string]any
				return hx_zero_159
			}
			return hx_field_158.([]map[string]any)
		}(c)
		for _g_3 < len(_g1_3) {
			f_1 := _g1_3[_g_3]
			_g_3 = int(int32((_g_3 + 1)))
			if self.mergeFields(f_1, f2_1) {
				found_1 = f_1
				break
			}
		}
		if found_1 == nil {
			self.newField(c, f2_1)
			hx_obj_161 := c
			hx_arr_160 := func(hx_obj_162 map[string]any) []map[string]any {
				hx_field_163 := hx_obj_162["statics"]
				if hx_field_163 == nil {
					var hx_zero_164 []map[string]any
					return hx_zero_164
				}
				return hx_field_163.([]map[string]any)
			}(c)
			hx_arr_160 = append(hx_arr_160, f2_1)
			hx_obj_161["statics"] = hx_arr_160
		} else {
			if !hxrt.StringEqualStringPtr(self.curplatform, nil) {
				hx_obj_166 := found_1
				hx_arr_165 := func(hx_obj_167 map[string]any) []*string {
					hx_field_168 := hx_obj_167["platforms"]
					if hx_field_168 == nil {
						var hx_zero_169 []*string
						return hx_zero_169
					}
					return hx_field_168.([]*string)
				}(found_1)
				hx_arr_165 = append(hx_arr_165, self.curplatform)
				hx_obj_166["platforms"] = hx_arr_165
			}
		}
	}
	return true
}

func (self *haxe__rtti__XmlParser) mergeEnums(e map[string]any, e2 map[string]any) bool {
	if func(hx_obj_170 map[string]any) bool {
		hx_field_171 := hx_obj_170["isExtern"]
		if hx_field_171 == nil {
			var hx_zero_172 bool
			return hx_zero_172
		}
		return hx_field_171.(bool)
	}(e) != func(hx_obj_173 map[string]any) bool {
		hx_field_174 := hx_obj_173["isExtern"]
		if hx_field_174 == nil {
			var hx_zero_175 bool
			return hx_zero_175
		}
		return hx_field_174.(bool)
	}(e2) {
		return false
	}
	if !hxrt.StringEqualStringPtr(self.curplatform, nil) {
		hx_obj_177 := e
		hx_arr_176 := func(hx_obj_178 map[string]any) []*string {
			hx_field_179 := hx_obj_178["platforms"]
			if hx_field_179 == nil {
				var hx_zero_180 []*string
				return hx_zero_180
			}
			return hx_field_179.([]*string)
		}(e)
		hx_arr_176 = append(hx_arr_176, self.curplatform)
		hx_obj_177["platforms"] = hx_arr_176
	}
	_g := 0
	_g1 := func(hx_obj_181 map[string]any) []map[string]any {
		hx_field_182 := hx_obj_181["constructors"]
		if hx_field_182 == nil {
			var hx_zero_183 []map[string]any
			return hx_zero_183
		}
		return hx_field_182.([]map[string]any)
	}(e2)
	for _g < len(_g1) {
		c2 := _g1[_g]
		_g = int(int32((_g + 1)))
		var found map[string]any = nil
		_g_1 := 0
		_g1_1 := func(hx_obj_184 map[string]any) []map[string]any {
			hx_field_185 := hx_obj_184["constructors"]
			if hx_field_185 == nil {
				var hx_zero_186 []map[string]any
				return hx_zero_186
			}
			return hx_field_185.([]map[string]any)
		}(e)
		for _g_1 < len(_g1_1) {
			c := _g1_1[_g_1]
			_g_1 = int(int32((_g_1 + 1)))
			if haxe__rtti__TypeApi_constructorEq(c, c2) {
				found = c
				break
			}
		}
		if found == nil {
			hx_obj_188 := e
			hx_arr_187 := func(hx_obj_189 map[string]any) []map[string]any {
				hx_field_190 := hx_obj_189["constructors"]
				if hx_field_190 == nil {
					var hx_zero_191 []map[string]any
					return hx_zero_191
				}
				return hx_field_190.([]map[string]any)
			}(e)
			hx_arr_187 = append(hx_arr_187, c2)
			hx_obj_188["constructors"] = hx_arr_187
		} else {
			if !hxrt.StringEqualStringPtr(self.curplatform, nil) {
				hx_obj_193 := found
				hx_arr_192 := func(hx_obj_194 map[string]any) []*string {
					hx_field_195 := hx_obj_194["platforms"]
					if hx_field_195 == nil {
						var hx_zero_196 []*string
						return hx_zero_196
					}
					return hx_field_195.([]*string)
				}(found)
				hx_arr_192 = append(hx_arr_192, self.curplatform)
				hx_obj_193["platforms"] = hx_arr_192
			}
		}
	}
	return true
}

func (self *haxe__rtti__XmlParser) mergeTypedefs(t map[string]any, t2 map[string]any) bool {
	if hxrt.StringEqualStringPtr(self.curplatform, nil) {
		return false
	}
	hx_obj_198 := t
	hx_arr_197 := func(hx_obj_199 map[string]any) []*string {
		hx_field_200 := hx_obj_199["platforms"]
		if hx_field_200 == nil {
			var hx_zero_201 []*string
			return hx_zero_201
		}
		return hx_field_200.([]*string)
	}(t)
	hx_arr_197 = append(hx_arr_197, self.curplatform)
	hx_obj_198["platforms"] = hx_arr_197
	var this1 haxe__IMap = func(hx_obj_202 map[string]any) *haxe__ds__StringMap {
		hx_field_203 := hx_obj_202["types"]
		if hx_field_203 == nil {
			var hx_zero_204 *haxe__ds__StringMap
			return hx_zero_204
		}
		return hx_field_203.(*haxe__ds__StringMap)
	}(t)
	key := self.curplatform
	value := func(hx_obj_205 map[string]any) *haxe__rtti__CType {
		hx_field_206 := hx_obj_205["type"]
		if hx_field_206 == nil {
			var hx_zero_207 *haxe__rtti__CType
			return hx_zero_207
		}
		return hx_field_206.(*haxe__rtti__CType)
	}(t2)
	this1.(*haxe__ds__StringMap).set(key, value)
	return true
}

func (self *haxe__rtti__XmlParser) mergeAbstracts(a map[string]any, a2 map[string]any) bool {
	if hxrt.StringEqualStringPtr(self.curplatform, nil) {
		return false
	}
	if (len(func(hx_obj_208 map[string]any) []map[string]any {
		hx_field_209 := hx_obj_208["to"]
		if hx_field_209 == nil {
			var hx_zero_210 []map[string]any
			return hx_zero_210
		}
		return hx_field_209.([]map[string]any)
	}(a)) != len(func(hx_obj_211 map[string]any) []map[string]any {
		hx_field_212 := hx_obj_211["to"]
		if hx_field_212 == nil {
			var hx_zero_213 []map[string]any
			return hx_zero_213
		}
		return hx_field_212.([]map[string]any)
	}(a2))) || (len(func(hx_obj_214 map[string]any) []map[string]any {
		hx_field_215 := hx_obj_214["from"]
		if hx_field_215 == nil {
			var hx_zero_216 []map[string]any
			return hx_zero_216
		}
		return hx_field_215.([]map[string]any)
	}(a)) != len(func(hx_obj_217 map[string]any) []map[string]any {
		hx_field_218 := hx_obj_217["from"]
		if hx_field_218 == nil {
			var hx_zero_219 []map[string]any
			return hx_zero_219
		}
		return hx_field_218.([]map[string]any)
	}(a2))) {
		return false
	}
	_g := 0
	_g1 := len(func(hx_obj_220 map[string]any) []map[string]any {
		hx_field_221 := hx_obj_220["to"]
		if hx_field_221 == nil {
			var hx_zero_222 []map[string]any
			return hx_zero_222
		}
		return hx_field_221.([]map[string]any)
	}(a))
	for _g < _g1 {
		hx_post_223 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_223
		if !haxe__rtti__TypeApi_typeEq(func(hx_obj_227 map[string]any) *haxe__rtti__CType {
			hx_field_228 := hx_obj_227["t"]
			if hx_field_228 == nil {
				var hx_zero_229 *haxe__rtti__CType
				return hx_zero_229
			}
			return hx_field_228.(*haxe__rtti__CType)
		}(func(hx_obj_224 map[string]any) []map[string]any {
			hx_field_225 := hx_obj_224["to"]
			if hx_field_225 == nil {
				var hx_zero_226 []map[string]any
				return hx_zero_226
			}
			return hx_field_225.([]map[string]any)
		}(a)[i]), func(hx_obj_233 map[string]any) *haxe__rtti__CType {
			hx_field_234 := hx_obj_233["t"]
			if hx_field_234 == nil {
				var hx_zero_235 *haxe__rtti__CType
				return hx_zero_235
			}
			return hx_field_234.(*haxe__rtti__CType)
		}(func(hx_obj_230 map[string]any) []map[string]any {
			hx_field_231 := hx_obj_230["to"]
			if hx_field_231 == nil {
				var hx_zero_232 []map[string]any
				return hx_zero_232
			}
			return hx_field_231.([]map[string]any)
		}(a2)[i])) {
			return false
		}
	}
	_g_1 := 0
	_g1_1 := len(func(hx_obj_236 map[string]any) []map[string]any {
		hx_field_237 := hx_obj_236["from"]
		if hx_field_237 == nil {
			var hx_zero_238 []map[string]any
			return hx_zero_238
		}
		return hx_field_237.([]map[string]any)
	}(a))
	for _g_1 < _g1_1 {
		hx_post_239 := _g_1
		_g_1 = int(int32((_g_1 + 1)))
		i_1 := hx_post_239
		if !haxe__rtti__TypeApi_typeEq(func(hx_obj_243 map[string]any) *haxe__rtti__CType {
			hx_field_244 := hx_obj_243["t"]
			if hx_field_244 == nil {
				var hx_zero_245 *haxe__rtti__CType
				return hx_zero_245
			}
			return hx_field_244.(*haxe__rtti__CType)
		}(func(hx_obj_240 map[string]any) []map[string]any {
			hx_field_241 := hx_obj_240["from"]
			if hx_field_241 == nil {
				var hx_zero_242 []map[string]any
				return hx_zero_242
			}
			return hx_field_241.([]map[string]any)
		}(a)[i_1]), func(hx_obj_249 map[string]any) *haxe__rtti__CType {
			hx_field_250 := hx_obj_249["t"]
			if hx_field_250 == nil {
				var hx_zero_251 *haxe__rtti__CType
				return hx_zero_251
			}
			return hx_field_250.(*haxe__rtti__CType)
		}(func(hx_obj_246 map[string]any) []map[string]any {
			hx_field_247 := hx_obj_246["from"]
			if hx_field_247 == nil {
				var hx_zero_248 []map[string]any
				return hx_zero_248
			}
			return hx_field_247.([]map[string]any)
		}(a2)[i_1])) {
			return false
		}
	}
	if func(hx_obj_258 map[string]any) map[string]any {
		hx_field_259 := hx_obj_258["impl"]
		if hx_field_259 == nil {
			var hx_zero_260 map[string]any
			return hx_zero_260
		}
		return hx_field_259.(map[string]any)
	}(a2) != nil {
		self.mergeClasses(func(hx_obj_252 map[string]any) map[string]any {
			hx_field_253 := hx_obj_252["impl"]
			if hx_field_253 == nil {
				var hx_zero_254 map[string]any
				return hx_zero_254
			}
			return hx_field_253.(map[string]any)
		}(a), func(hx_obj_255 map[string]any) map[string]any {
			hx_field_256 := hx_obj_255["impl"]
			if hx_field_256 == nil {
				var hx_zero_257 map[string]any
				return hx_zero_257
			}
			return hx_field_256.(map[string]any)
		}(a2))
	}
	hx_obj_262 := a
	hx_arr_261 := func(hx_obj_263 map[string]any) []*string {
		hx_field_264 := hx_obj_263["platforms"]
		if hx_field_264 == nil {
			var hx_zero_265 []*string
			return hx_zero_265
		}
		return hx_field_264.([]*string)
	}(a)
	hx_arr_261 = append(hx_arr_261, self.curplatform)
	hx_obj_262["platforms"] = hx_arr_261
	return true
}

func (self *haxe__rtti__XmlParser) merge(t *haxe__rtti__TypeTree) {
	inf := haxe__rtti__TypeApi_typeInfos(t)
	pack := self.splitString(func(hx_obj_266 map[string]any) *string {
		hx_field_267 := hx_obj_266["path"]
		if hx_field_267 == nil {
			var hx_zero_268 *string
			return hx_zero_268
		}
		return hx_field_267.(*string)
	}(inf), hxrt.StringFromLiteral("."))
	cur := self.root
	curpack := []*string{}
	if len(pack) > 0 {
		pack = pack[:(len(pack) - 1)]
	}
	_g := 0
	for _g < len(pack) {
		p := pack[_g]
		_g = int(int32((_g + 1)))
		found := false
		_g_1 := 0
		for _g_1 < len(cur) {
			pk := cur[_g_1]
			_g_1 = int(int32((_g_1 + 1)))
			if pk.tag == 0 {
				_g_2 := pk.params[0].(*string)
				_g1 := pk.params[1].(*string)
				_ = _g1
				_g1_1 := pk.params[2].([]*haxe__rtti__TypeTree)
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
		curpack = append(curpack, p)
		if !found {
			pk_1 := []*haxe__rtti__TypeTree{}
			cur = append(cur, haxe__rtti__TypeTree_TPackage(p, self.joinStringArray(curpack, hxrt.StringFromLiteral(".")), pk_1))
			cur = pk_1
		}
	}
	_g_3 := 0
	for _g_3 < len(cur) {
		ct := cur[_g_3]
		_g_3 = int(int32((_g_3 + 1)))
		if func() bool {
			var hx_if_272 bool
			if ct.tag == 0 {
				_g_4 := ct.params[0].(*string)
				_ = _g_4
				_g_5 := ct.params[1].(*string)
				_ = _g_5
				_g_6 := ct.params[2].([]*haxe__rtti__TypeTree)
				_ = _g_6
				hx_if_272 = true
			} else {
				hx_if_272 = false
			}
			return hx_if_272
		}() {
			continue
		}
		tinf := haxe__rtti__TypeApi_typeInfos(ct)
		if hxrt.StringEqualStringPtr(func(hx_obj_343 map[string]any) *string {
			hx_field_344 := hx_obj_343["path"]
			if hx_field_344 == nil {
				var hx_zero_345 *string
				return hx_zero_345
			}
			return hx_field_344.(*string)
		}(tinf), func(hx_obj_346 map[string]any) *string {
			hx_field_347 := hx_obj_346["path"]
			if hx_field_347 == nil {
				var hx_zero_348 *string
				return hx_zero_348
			}
			return hx_field_347.(*string)
		}(inf)) {
			sameType := true
			if hxrt.StringEqualStringPtr(func(hx_obj_282 map[string]any) *string {
				hx_field_283 := hx_obj_282["doc"]
				if hx_field_283 == nil {
					var hx_zero_284 *string
					return hx_zero_284
				}
				return hx_field_283.(*string)
			}(tinf), nil) != hxrt.StringEqualStringPtr(func(hx_obj_285 map[string]any) *string {
				hx_field_286 := hx_obj_285["doc"]
				if hx_field_286 == nil {
					var hx_zero_287 *string
					return hx_zero_287
				}
				return hx_field_286.(*string)
			}(inf), nil) {
				if hxrt.StringEqualStringPtr(func(hx_obj_279 map[string]any) *string {
					hx_field_280 := hx_obj_279["doc"]
					if hx_field_280 == nil {
						var hx_zero_281 *string
						return hx_zero_281
					}
					return hx_field_280.(*string)
				}(inf), nil) {
					inf["doc"] = func(hx_obj_273 map[string]any) *string {
						hx_field_274 := hx_obj_273["doc"]
						if hx_field_274 == nil {
							var hx_zero_275 *string
							return hx_zero_275
						}
						return hx_field_274.(*string)
					}(tinf)
				} else {
					tinf["doc"] = func(hx_obj_276 map[string]any) *string {
						hx_field_277 := hx_obj_276["doc"]
						if hx_field_277 == nil {
							var hx_zero_278 *string
							return hx_zero_278
						}
						return hx_field_277.(*string)
					}(inf)
				}
			}
			if hxrt.StringEqualStringPtr(func(hx_obj_288 map[string]any) *string {
				hx_field_289 := hx_obj_288["path"]
				if hx_field_289 == nil {
					var hx_zero_290 *string
					return hx_zero_290
				}
				return hx_field_289.(*string)
			}(tinf), hxrt.StringFromLiteral("haxe._Int64.NativeInt64")) {
				continue
			}
			if (hxrt.StringEqualStringPtr(func(hx_obj_291 map[string]any) *string {
				hx_field_292 := hx_obj_291["module"]
				if hx_field_292 == nil {
					var hx_zero_293 *string
					return hx_zero_293
				}
				return hx_field_292.(*string)
			}(tinf), func(hx_obj_294 map[string]any) *string {
				hx_field_295 := hx_obj_294["module"]
				if hx_field_295 == nil {
					var hx_zero_296 *string
					return hx_zero_296
				}
				return hx_field_295.(*string)
			}(inf)) && hxrt.StringEqualStringPtr(func(hx_obj_297 map[string]any) *string {
				hx_field_298 := hx_obj_297["doc"]
				if hx_field_298 == nil {
					var hx_zero_299 *string
					return hx_zero_299
				}
				return hx_field_298.(*string)
			}(tinf), func(hx_obj_300 map[string]any) *string {
				hx_field_301 := hx_obj_300["doc"]
				if hx_field_301 == nil {
					var hx_zero_302 *string
					return hx_zero_302
				}
				return hx_field_301.(*string)
			}(inf))) && (func(hx_obj_303 map[string]any) bool {
				hx_field_304 := hx_obj_303["isPrivate"]
				if hx_field_304 == nil {
					var hx_zero_305 bool
					return hx_zero_305
				}
				return hx_field_304.(bool)
			}(tinf) == func(hx_obj_306 map[string]any) bool {
				hx_field_307 := hx_obj_306["isPrivate"]
				if hx_field_307 == nil {
					var hx_zero_308 bool
					return hx_zero_308
				}
				return hx_field_307.(bool)
			}(inf)) {
				switch ct.tag {
				case 0:
					_g_7 := ct.params[0].(*string)
					_ = _g_7
					_g_8 := ct.params[1].(*string)
					_ = _g_8
					_g_9 := ct.params[2].([]*haxe__rtti__TypeTree)
					_ = _g_9
					sameType = false
				case 1:
					_g_10 := ct.params[0].(map[string]any)
					c := _g_10
					if t.tag == 1 {
						_g_11 := t.params[0].(map[string]any)
						c2 := _g_11
						if self.mergeClasses(c, c2) {
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
						if self.mergeEnums(e, e2) {
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
						if self.mergeTypedefs(td, td2) {
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
						if self.mergeAbstracts(a, a2) {
							return
						}
					} else {
						sameType = false
					}
				}
			}
			var hx_if_336 *string
			if !hxrt.StringEqualStringPtr(func(hx_obj_309 map[string]any) *string {
				hx_field_310 := hx_obj_309["module"]
				if hx_field_310 == nil {
					var hx_zero_311 *string
					return hx_zero_311
				}
				return hx_field_310.(*string)
			}(tinf), func(hx_obj_312 map[string]any) *string {
				hx_field_313 := hx_obj_312["module"]
				if hx_field_313 == nil {
					var hx_zero_314 *string
					return hx_zero_314
				}
				return hx_field_313.(*string)
			}(inf)) {
				hx_if_336 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("module "), func(hx_obj_315 map[string]any) *string {
					hx_field_316 := hx_obj_315["module"]
					if hx_field_316 == nil {
						var hx_zero_317 *string
						return hx_zero_317
					}
					return hx_field_316.(*string)
				}(inf)), hxrt.StringFromLiteral(" should be ")), func(hx_obj_318 map[string]any) *string {
					hx_field_319 := hx_obj_318["module"]
					if hx_field_319 == nil {
						var hx_zero_320 *string
						return hx_zero_320
					}
					return hx_field_319.(*string)
				}(tinf))
			} else {
				var hx_if_335 *string
				if !hxrt.StringEqualStringPtr(func(hx_obj_321 map[string]any) *string {
					hx_field_322 := hx_obj_321["doc"]
					if hx_field_322 == nil {
						var hx_zero_323 *string
						return hx_zero_323
					}
					return hx_field_322.(*string)
				}(tinf), func(hx_obj_324 map[string]any) *string {
					hx_field_325 := hx_obj_324["doc"]
					if hx_field_325 == nil {
						var hx_zero_326 *string
						return hx_zero_326
					}
					return hx_field_325.(*string)
				}(inf)) {
					hx_if_335 = hxrt.StringFromLiteral("documentation is different")
				} else {
					var hx_if_334 *string
					if func(hx_obj_327 map[string]any) bool {
						hx_field_328 := hx_obj_327["isPrivate"]
						if hx_field_328 == nil {
							var hx_zero_329 bool
							return hx_zero_329
						}
						return hx_field_328.(bool)
					}(tinf) != func(hx_obj_330 map[string]any) bool {
						hx_field_331 := hx_obj_330["isPrivate"]
						if hx_field_331 == nil {
							var hx_zero_332 bool
							return hx_zero_332
						}
						return hx_field_331.(bool)
					}(inf) {
						hx_if_334 = hxrt.StringFromLiteral("private flag is different")
					} else {
						var hx_if_333 *string
						if !sameType {
							hx_if_333 = hxrt.StringFromLiteral("type kind is different")
						} else {
							hx_if_333 = hxrt.StringFromLiteral("could not merge definition")
						}
						hx_if_334 = hx_if_333
					}
					hx_if_335 = hx_if_334
				}
				hx_if_336 = hx_if_335
			}
			msg := hx_if_336
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Incompatibilities between "), func(hx_obj_337 map[string]any) *string {
				hx_field_338 := hx_obj_337["path"]
				if hx_field_338 == nil {
					var hx_zero_339 *string
					return hx_zero_339
				}
				return hx_field_338.(*string)
			}(tinf)), hxrt.StringFromLiteral(" in ")), self.joinStringArray(func(hx_obj_340 map[string]any) []*string {
				hx_field_341 := hx_obj_340["platforms"]
				if hx_field_341 == nil {
					var hx_zero_342 []*string
					return hx_zero_342
				}
				return hx_field_341.([]*string)
			}(tinf), hxrt.StringFromLiteral(","))), hxrt.StringFromLiteral(" and ")), self.curplatform), hxrt.StringFromLiteral(" (")), msg), hxrt.StringFromLiteral(")")))
		}
	}
	cur = append(cur, t)
}

func (self *haxe__rtti__XmlParser) mkPath(p *string) *string {
	return p
}

func (self *haxe__rtti__XmlParser) mkTypeParams(p *string) []*string {
	pl := self.splitString(p, hxrt.StringFromLiteral(":"))
	if hxrt.StringEqualStringPtr(pl[0], hxrt.StringFromLiteral("")) {
		return []*string{}
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
	c := x.elements()
	for func(hx_obj_350 map[string]any) func() bool {
		hx_field_351 := hx_obj_350["hasNext"]
		if hx_field_351 == nil {
			var hx_zero_352 func() bool
			return hx_zero_352
		}
		return hx_field_351.(func() bool)
	}(c)() {
		c_1 := func(hx_obj_353 map[string]any) func() *Xml {
			hx_field_354 := hx_obj_353["next"]
			if hx_field_354 == nil {
				var hx_zero_355 func() *Xml
				return hx_zero_355
			}
			return hx_field_354.(func() *Xml)
		}(c)()
		self.merge(self.processElement(c_1))
	}
}

func (self *haxe__rtti__XmlParser) processElement(x *Xml) *haxe__rtti__TypeTree {
	var hx_if_356 *string
	if x.nodeType == Xml_Document {
		hx_if_356 = hxrt.StringFromLiteral("Document")
	} else {
		hx_if_356 = x.get_nodeName()
	}
	nodeName := hx_if_356
	if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("class")) {
		return haxe__rtti__TypeTree_TClassdecl(self.xclass(x))
	}
	if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("enum")) {
		return haxe__rtti__TypeTree_TEnumdecl(self.xenum(x))
	}
	if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("typedef")) {
		return haxe__rtti__TypeTree_TTypedecl(self.xtypedef(x))
	}
	if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("abstract")) {
		return haxe__rtti__TypeTree_TAbstractdecl(self.xabstract(x))
	}
	hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Invalid "), nodeName))
	var hx_throw_zero_357 *haxe__rtti__TypeTree
	return hx_throw_zero_357
}

func (self *haxe__rtti__XmlParser) xmeta(x *Xml) []map[string]any {
	ml := []map[string]any{}
	m := x.elementsNamed(hxrt.StringFromLiteral("m"))
	for func(hx_obj_358 map[string]any) func() bool {
		hx_field_359 := hx_obj_358["hasNext"]
		if hx_field_359 == nil {
			var hx_zero_360 func() bool
			return hx_zero_360
		}
		return hx_field_359.(func() bool)
	}(m)() {
		m_1 := func(hx_obj_361 map[string]any) func() *Xml {
			hx_field_362 := hx_obj_361["next"]
			if hx_field_362 == nil {
				var hx_zero_363 func() *Xml
				return hx_zero_363
			}
			return hx_field_362.(func() *Xml)
		}(m)()
		pl := []*string{}
		p := m_1.elementsNamed(hxrt.StringFromLiteral("e"))
		for func(hx_obj_364 map[string]any) func() bool {
			hx_field_365 := hx_obj_364["hasNext"]
			if hx_field_365 == nil {
				var hx_zero_366 func() bool
				return hx_zero_366
			}
			return hx_field_365.(func() bool)
		}(p)() {
			p_1 := func(hx_obj_367 map[string]any) func() *Xml {
				hx_field_368 := hx_obj_367["next"]
				if hx_field_368 == nil {
					var hx_zero_369 func() *Xml
					return hx_zero_369
				}
				return hx_field_368.(func() *Xml)
			}(p)()
			pl = append(pl, self.innerHTML(p_1))
		}
		ml = append(ml, func() map[string]any {
			hx_obj_372 := map[string]any{}
			hx_obj_372["name"] = self.requireAttr(m_1, hxrt.StringFromLiteral("n"))
			hx_obj_372["params"] = pl
			return hx_obj_372
		}())
	}
	return ml
}

func (self *haxe__rtti__XmlParser) xoverloads(x *Xml) []map[string]any {
	l := []map[string]any{}
	m := x.elements()
	for func(hx_obj_373 map[string]any) func() bool {
		hx_field_374 := hx_obj_373["hasNext"]
		if hx_field_374 == nil {
			var hx_zero_375 func() bool
			return hx_zero_375
		}
		return hx_field_374.(func() bool)
	}(m)() {
		m_1 := func(hx_obj_376 map[string]any) func() *Xml {
			hx_field_377 := hx_obj_376["next"]
			if hx_field_377 == nil {
				var hx_zero_378 func() *Xml
				return hx_zero_378
			}
			return hx_field_377.(func() *Xml)
		}(m)()
		l = append(l, self.xclassfield(m_1, false))
	}
	return l
}

func (self *haxe__rtti__XmlParser) xpath(x *Xml) map[string]any {
	path := self.mkPath(self.requireAttr(x, hxrt.StringFromLiteral("path")))
	params := []*haxe__rtti__CType{}
	c := x.elements()
	for func(hx_obj_380 map[string]any) func() bool {
		hx_field_381 := hx_obj_380["hasNext"]
		if hx_field_381 == nil {
			var hx_zero_382 func() bool
			return hx_zero_382
		}
		return hx_field_381.(func() bool)
	}(c)() {
		c_1 := func(hx_obj_383 map[string]any) func() *Xml {
			hx_field_384 := hx_obj_383["next"]
			if hx_field_384 == nil {
				var hx_zero_385 func() *Xml
				return hx_zero_385
			}
			return hx_field_384.(func() *Xml)
		}(c)()
		params = append(params, self.xtype(c_1))
	}
	hx_obj_387 := map[string]any{}
	hx_obj_387["path"] = path
	hx_obj_387["params"] = params
	return hx_obj_387
}

func (self *haxe__rtti__XmlParser) xclass(x *Xml) map[string]any {
	var csuper map[string]any = nil
	var doc *string = nil
	var tdynamic *haxe__rtti__CType = nil
	interfaces := []map[string]any{}
	fields := []map[string]any{}
	statics := []map[string]any{}
	meta := []map[string]any{}
	isInterface := x.exists(hxrt.StringFromLiteral("interface"))
	c := x.elements()
	for func(hx_obj_388 map[string]any) func() bool {
		hx_field_389 := hx_obj_388["hasNext"]
		if hx_field_389 == nil {
			var hx_zero_390 func() bool
			return hx_zero_390
		}
		return hx_field_389.(func() bool)
	}(c)() {
		c_1 := func(hx_obj_391 map[string]any) func() *Xml {
			hx_field_392 := hx_obj_391["next"]
			if hx_field_392 == nil {
				var hx_zero_393 func() *Xml
				return hx_zero_393
			}
			return hx_field_392.(func() *Xml)
		}(c)()
		nodeName := self.elementName(c_1)
		if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("haxe_doc")) {
			doc = self.innerData(c_1)
		} else {
			if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("extends")) {
				if isInterface {
					interfaces = append(interfaces, self.xpath(c_1))
				} else {
					csuper = self.xpath(c_1)
				}
			} else {
				if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("implements")) {
					interfaces = append(interfaces, self.xpath(c_1))
				} else {
					if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("haxe_dynamic")) {
						tdynamic = self.xtype(self.requireFirstElement(c_1))
					} else {
						if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("meta")) {
							meta = self.xmeta(c_1)
						} else {
							if c_1.exists(hxrt.StringFromLiteral("static")) {
								statics = append(statics, self.xclassfield(c_1, false))
							} else {
								fields = append(fields, self.xclassfield(c_1, false))
							}
						}
					}
				}
			}
		}
	}
	hx_obj_398 := map[string]any{}
	hx_obj_398["file"] = x.get(hxrt.StringFromLiteral("file"))
	hx_obj_398["path"] = self.mkPath(self.requireAttr(x, hxrt.StringFromLiteral("path")))
	var hx_if_399 *string
	if x.exists(hxrt.StringFromLiteral("module")) {
		hx_if_399 = self.mkPath(self.requireAttr(x, hxrt.StringFromLiteral("module")))
	} else {
		hx_if_399 = nil
	}
	hx_obj_398["module"] = hx_if_399
	hx_obj_398["doc"] = doc
	hx_obj_398["isPrivate"] = x.exists(hxrt.StringFromLiteral("private"))
	hx_obj_398["isExtern"] = x.exists(hxrt.StringFromLiteral("extern"))
	hx_obj_398["isFinal"] = x.exists(hxrt.StringFromLiteral("final"))
	hx_obj_398["isInterface"] = isInterface
	hx_obj_398["params"] = self.mkTypeParams(self.requireAttr(x, hxrt.StringFromLiteral("params")))
	hx_obj_398["superClass"] = csuper
	hx_obj_398["interfaces"] = interfaces
	hx_obj_398["fields"] = fields
	hx_obj_398["statics"] = statics
	hx_obj_398["tdynamic"] = tdynamic
	hx_obj_398["platforms"] = self.defplat()
	hx_obj_398["meta"] = meta
	return hx_obj_398
}

func (self *haxe__rtti__XmlParser) xclassfield(x *Xml, defPublic bool) map[string]any {
	e := x.elements()
	t := self.xtype(func(hx_obj_400 map[string]any) func() *Xml {
		hx_field_401 := hx_obj_400["next"]
		if hx_field_401 == nil {
			var hx_zero_402 func() *Xml
			return hx_zero_402
		}
		return hx_field_401.(func() *Xml)
	}(e)())
	var doc *string = nil
	meta := []map[string]any{}
	var overloads []map[string]any = nil
	var line any = nil
	if x.exists(hxrt.StringFromLiteral("line")) {
		line = self.parseIntString(self.requireAttr(x, hxrt.StringFromLiteral("line")))
	}
	c := e
	for func(hx_obj_403 map[string]any) func() bool {
		hx_field_404 := hx_obj_403["hasNext"]
		if hx_field_404 == nil {
			var hx_zero_405 func() bool
			return hx_zero_405
		}
		return hx_field_404.(func() bool)
	}(c)() {
		c_1 := func(hx_obj_406 map[string]any) func() *Xml {
			hx_field_407 := hx_obj_406["next"]
			if hx_field_407 == nil {
				var hx_zero_408 func() *Xml
				return hx_zero_408
			}
			return hx_field_407.(func() *Xml)
		}(c)()
		nodeName := self.elementName(c_1)
		if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("haxe_doc")) {
			doc = self.innerData(c_1)
		} else {
			if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("meta")) {
				meta = self.xmeta(c_1)
			} else {
				if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("overloads")) {
					overloads = self.xoverloads(c_1)
				} else {
					hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Invalid "), nodeName))
					var hx_throw_zero_409 map[string]any
					return hx_throw_zero_409
				}
			}
		}
	}
	hx_obj_410 := map[string]any{}
	hx_obj_410["name"] = self.elementName(x)
	hx_obj_410["type"] = t
	hx_obj_410["isPublic"] = (x.exists(hxrt.StringFromLiteral("public")) || func(hx_value_411 any) bool {
		if hx_value_411 == nil {
			var hx_zero_412 bool
			return hx_zero_412
		}
		return hx_value_411.(bool)
	}(defPublic))
	hx_obj_410["isFinal"] = x.exists(hxrt.StringFromLiteral("final"))
	hx_obj_410["isOverride"] = x.exists(hxrt.StringFromLiteral("override"))
	hx_obj_410["line"] = func(hx_value_413 any) any {
		if hx_value_413 == nil {
			return nil
		}
		return hx_value_413.(int)
	}(line)
	hx_obj_410["doc"] = doc
	var hx_if_414 *haxe__rtti__Rights
	if x.exists(hxrt.StringFromLiteral("get")) {
		hx_if_414 = self.mkRights(self.requireAttr(x, hxrt.StringFromLiteral("get")))
	} else {
		hx_if_414 = haxe__rtti__Rights_RNormal
	}
	hx_obj_410["get"] = hx_if_414
	var hx_if_415 *haxe__rtti__Rights
	if x.exists(hxrt.StringFromLiteral("set")) {
		hx_if_415 = self.mkRights(self.requireAttr(x, hxrt.StringFromLiteral("set")))
	} else {
		hx_if_415 = haxe__rtti__Rights_RNormal
	}
	hx_obj_410["set"] = hx_if_415
	var hx_if_416 []*string
	if x.exists(hxrt.StringFromLiteral("params")) {
		hx_if_416 = self.mkTypeParams(self.requireAttr(x, hxrt.StringFromLiteral("params")))
	} else {
		hx_if_416 = []*string{}
	}
	hx_obj_410["params"] = hx_if_416
	hx_obj_410["platforms"] = self.defplat()
	hx_obj_410["meta"] = meta
	hx_obj_410["overloads"] = overloads
	var hx_if_417 *string
	if x.exists(hxrt.StringFromLiteral("expr")) {
		hx_if_417 = self.requireAttr(x, hxrt.StringFromLiteral("expr"))
	} else {
		hx_if_417 = nil
	}
	hx_obj_410["expr"] = hx_if_417
	return hx_obj_410
}

func (self *haxe__rtti__XmlParser) xenum(x *Xml) map[string]any {
	cl := []map[string]any{}
	var doc *string = nil
	meta := []map[string]any{}
	c := x.elements()
	for func(hx_obj_418 map[string]any) func() bool {
		hx_field_419 := hx_obj_418["hasNext"]
		if hx_field_419 == nil {
			var hx_zero_420 func() bool
			return hx_zero_420
		}
		return hx_field_419.(func() bool)
	}(c)() {
		c_1 := func(hx_obj_421 map[string]any) func() *Xml {
			hx_field_422 := hx_obj_421["next"]
			if hx_field_422 == nil {
				var hx_zero_423 func() *Xml
				return hx_zero_423
			}
			return hx_field_422.(func() *Xml)
		}(c)()
		if hxrt.StringEqualStringPtr(self.elementName(c_1), hxrt.StringFromLiteral("haxe_doc")) {
			doc = self.innerData(c_1)
		} else {
			if hxrt.StringEqualStringPtr(self.elementName(c_1), hxrt.StringFromLiteral("meta")) {
				meta = self.xmeta(c_1)
			} else {
				cl = append(cl, self.xenumfield(c_1))
			}
		}
	}
	hx_obj_425 := map[string]any{}
	hx_obj_425["file"] = x.get(hxrt.StringFromLiteral("file"))
	hx_obj_425["path"] = self.mkPath(self.requireAttr(x, hxrt.StringFromLiteral("path")))
	var hx_if_426 *string
	if x.exists(hxrt.StringFromLiteral("module")) {
		hx_if_426 = self.mkPath(self.requireAttr(x, hxrt.StringFromLiteral("module")))
	} else {
		hx_if_426 = nil
	}
	hx_obj_425["module"] = hx_if_426
	hx_obj_425["doc"] = doc
	hx_obj_425["isPrivate"] = x.exists(hxrt.StringFromLiteral("private"))
	hx_obj_425["isExtern"] = x.exists(hxrt.StringFromLiteral("extern"))
	hx_obj_425["params"] = self.mkTypeParams(self.requireAttr(x, hxrt.StringFromLiteral("params")))
	hx_obj_425["constructors"] = cl
	hx_obj_425["platforms"] = self.defplat()
	hx_obj_425["meta"] = meta
	return hx_obj_425
}

func (self *haxe__rtti__XmlParser) xenumfield(x *Xml) map[string]any {
	var args []map[string]any = nil
	docElements := x.elementsNamed(hxrt.StringFromLiteral("haxe_doc"))
	var hx_if_433 *Xml
	if func(hx_obj_427 map[string]any) func() bool {
		hx_field_428 := hx_obj_427["hasNext"]
		if hx_field_428 == nil {
			var hx_zero_429 func() bool
			return hx_zero_429
		}
		return hx_field_428.(func() bool)
	}(docElements)() {
		hx_if_433 = func(hx_obj_430 map[string]any) func() *Xml {
			hx_field_431 := hx_obj_430["next"]
			if hx_field_431 == nil {
				var hx_zero_432 func() *Xml
				return hx_zero_432
			}
			return hx_field_431.(func() *Xml)
		}(docElements)()
	} else {
		hx_if_433 = nil
	}
	xdoc := hx_if_433
	var hx_if_434 []map[string]any
	if self.hasNamedElement(x, hxrt.StringFromLiteral("meta")) {
		hx_if_434 = self.xmeta(self.requireNamedElement(x, hxrt.StringFromLiteral("meta")))
	} else {
		hx_if_434 = []map[string]any{}
	}
	meta := hx_if_434
	if x.exists(hxrt.StringFromLiteral("a")) {
		names := self.splitString(self.requireAttr(x, hxrt.StringFromLiteral("a")), hxrt.StringFromLiteral(":"))
		elts := x.elements()
		args = []map[string]any{}
		_g := 0
		for _g < len(names) {
			c := names[_g]
			_g = int(int32((_g + 1)))
			opt := false
			if hxrt.StringEqualStringPtr(hxrt.StringCharAtStringPtr(c, 0), hxrt.StringFromLiteral("?")) {
				opt = true
				c = hxrt.StringSubstrStringPtr(c, 1, 0, false)
			}
			args = append(args, func() map[string]any {
				hx_obj_436 := map[string]any{}
				hx_obj_436["name"] = c
				hx_obj_436["opt"] = opt
				hx_obj_436["t"] = self.xtype(func(hx_obj_437 map[string]any) func() *Xml {
					hx_field_438 := hx_obj_437["next"]
					if hx_field_438 == nil {
						var hx_zero_439 func() *Xml
						return hx_zero_439
					}
					return hx_field_438.(func() *Xml)
				}(elts)())
				return hx_obj_436
			}())
		}
	}
	hx_obj_440 := map[string]any{}
	hx_obj_440["name"] = self.elementName(x)
	hx_obj_440["args"] = args
	var hx_if_441 *string
	if xdoc == nil {
		hx_if_441 = nil
	} else {
		hx_if_441 = self.innerData(xdoc)
	}
	hx_obj_440["doc"] = hx_if_441
	hx_obj_440["meta"] = meta
	hx_obj_440["platforms"] = self.defplat()
	return hx_obj_440
}

func (self *haxe__rtti__XmlParser) xabstract(x *Xml) map[string]any {
	var doc *string = nil
	var impl map[string]any = nil
	var athis *haxe__rtti__CType = nil
	meta := []map[string]any{}
	to := []map[string]any{}
	from := []map[string]any{}
	c := x.elements()
	for func(hx_obj_442 map[string]any) func() bool {
		hx_field_443 := hx_obj_442["hasNext"]
		if hx_field_443 == nil {
			var hx_zero_444 func() bool
			return hx_zero_444
		}
		return hx_field_443.(func() bool)
	}(c)() {
		c_1 := func(hx_obj_445 map[string]any) func() *Xml {
			hx_field_446 := hx_obj_445["next"]
			if hx_field_446 == nil {
				var hx_zero_447 func() *Xml
				return hx_zero_447
			}
			return hx_field_446.(func() *Xml)
		}(c)()
		nodeName := self.elementName(c_1)
		if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("haxe_doc")) {
			doc = self.innerData(c_1)
		} else {
			if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("meta")) {
				meta = self.xmeta(c_1)
			} else {
				if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("to")) {
					t := c_1.elements()
					for func(hx_obj_448 map[string]any) func() bool {
						hx_field_449 := hx_obj_448["hasNext"]
						if hx_field_449 == nil {
							var hx_zero_450 func() bool
							return hx_zero_450
						}
						return hx_field_449.(func() bool)
					}(t)() {
						t_1 := func(hx_obj_451 map[string]any) func() *Xml {
							hx_field_452 := hx_obj_451["next"]
							if hx_field_452 == nil {
								var hx_zero_453 func() *Xml
								return hx_zero_453
							}
							return hx_field_452.(func() *Xml)
						}(t)()
						to = append(to, func() map[string]any {
							hx_obj_455 := map[string]any{}
							hx_obj_455["t"] = self.xtype(self.requireFirstElement(t_1))
							hx_obj_455["field"] = t_1.get(hxrt.StringFromLiteral("field"))
							return hx_obj_455
						}())
					}
				} else {
					if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("from")) {
						t_2 := c_1.elements()
						for func(hx_obj_456 map[string]any) func() bool {
							hx_field_457 := hx_obj_456["hasNext"]
							if hx_field_457 == nil {
								var hx_zero_458 func() bool
								return hx_zero_458
							}
							return hx_field_457.(func() bool)
						}(t_2)() {
							t_3 := func(hx_obj_459 map[string]any) func() *Xml {
								hx_field_460 := hx_obj_459["next"]
								if hx_field_460 == nil {
									var hx_zero_461 func() *Xml
									return hx_zero_461
								}
								return hx_field_460.(func() *Xml)
							}(t_2)()
							from = append(from, func() map[string]any {
								hx_obj_463 := map[string]any{}
								hx_obj_463["t"] = self.xtype(self.requireFirstElement(t_3))
								hx_obj_463["field"] = t_3.get(hxrt.StringFromLiteral("field"))
								return hx_obj_463
							}())
						}
					} else {
						if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("impl")) {
							impl = self.xclass(self.requireNamedElement(c_1, hxrt.StringFromLiteral("class")))
						} else {
							if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("this")) {
								athis = self.xtype(self.requireFirstElement(c_1))
							} else {
								hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Invalid "), nodeName))
								var hx_throw_zero_464 map[string]any
								return hx_throw_zero_464
							}
						}
					}
				}
			}
		}
	}
	hx_obj_465 := map[string]any{}
	hx_obj_465["file"] = x.get(hxrt.StringFromLiteral("file"))
	hx_obj_465["path"] = self.mkPath(self.requireAttr(x, hxrt.StringFromLiteral("path")))
	var hx_if_466 *string
	if x.exists(hxrt.StringFromLiteral("module")) {
		hx_if_466 = self.mkPath(self.requireAttr(x, hxrt.StringFromLiteral("module")))
	} else {
		hx_if_466 = nil
	}
	hx_obj_465["module"] = hx_if_466
	hx_obj_465["doc"] = doc
	hx_obj_465["isPrivate"] = x.exists(hxrt.StringFromLiteral("private"))
	hx_obj_465["params"] = self.mkTypeParams(self.requireAttr(x, hxrt.StringFromLiteral("params")))
	hx_obj_465["platforms"] = self.defplat()
	hx_obj_465["meta"] = meta
	hx_obj_465["athis"] = athis
	hx_obj_465["to"] = to
	hx_obj_465["from"] = from
	hx_obj_465["impl"] = impl
	return hx_obj_465
}

func (self *haxe__rtti__XmlParser) xtypedef(x *Xml) map[string]any {
	var doc *string = nil
	var t *haxe__rtti__CType = nil
	meta := []map[string]any{}
	c := x.elements()
	for func(hx_obj_467 map[string]any) func() bool {
		hx_field_468 := hx_obj_467["hasNext"]
		if hx_field_468 == nil {
			var hx_zero_469 func() bool
			return hx_zero_469
		}
		return hx_field_468.(func() bool)
	}(c)() {
		c_1 := func(hx_obj_470 map[string]any) func() *Xml {
			hx_field_471 := hx_obj_470["next"]
			if hx_field_471 == nil {
				var hx_zero_472 func() *Xml
				return hx_zero_472
			}
			return hx_field_471.(func() *Xml)
		}(c)()
		if hxrt.StringEqualStringPtr(self.elementName(c_1), hxrt.StringFromLiteral("haxe_doc")) {
			doc = self.innerData(c_1)
		} else {
			if hxrt.StringEqualStringPtr(self.elementName(c_1), hxrt.StringFromLiteral("meta")) {
				meta = self.xmeta(c_1)
			} else {
				t = self.xtype(c_1)
			}
		}
	}
	types := New_haxe__ds__StringMap()
	if !hxrt.StringEqualStringPtr(self.curplatform, nil) {
		types.set(self.curplatform, t)
	}
	hx_obj_473 := map[string]any{}
	hx_obj_473["file"] = x.get(hxrt.StringFromLiteral("file"))
	hx_obj_473["path"] = self.mkPath(self.requireAttr(x, hxrt.StringFromLiteral("path")))
	var hx_if_474 *string
	if x.exists(hxrt.StringFromLiteral("module")) {
		hx_if_474 = self.mkPath(self.requireAttr(x, hxrt.StringFromLiteral("module")))
	} else {
		hx_if_474 = nil
	}
	hx_obj_473["module"] = hx_if_474
	hx_obj_473["doc"] = doc
	hx_obj_473["isPrivate"] = x.exists(hxrt.StringFromLiteral("private"))
	hx_obj_473["params"] = self.mkTypeParams(self.requireAttr(x, hxrt.StringFromLiteral("params")))
	hx_obj_473["type"] = t
	hx_obj_473["types"] = types
	hx_obj_473["platforms"] = self.defplat()
	hx_obj_473["meta"] = meta
	return hx_obj_473
}

func (self *haxe__rtti__XmlParser) xtype(x *Xml) *haxe__rtti__CType {
	nodeName := self.elementName(x)
	if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("unknown")) {
		return haxe__rtti__CType_CUnknown
	}
	if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("e")) {
		return haxe__rtti__CType_CEnum(self.mkPath(self.requireAttr(x, hxrt.StringFromLiteral("path"))), self.xtypeparams(x))
	}
	if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("c")) {
		return haxe__rtti__CType_CClass(self.mkPath(self.requireAttr(x, hxrt.StringFromLiteral("path"))), self.xtypeparams(x))
	}
	if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("t")) {
		return haxe__rtti__CType_CTypedef(self.mkPath(self.requireAttr(x, hxrt.StringFromLiteral("path"))), self.xtypeparams(x))
	}
	if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("x")) {
		return haxe__rtti__CType_CAbstract(self.mkPath(self.requireAttr(x, hxrt.StringFromLiteral("path"))), self.xtypeparams(x))
	}
	if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("f")) {
		args := []map[string]any{}
		aname := self.splitString(self.requireAttr(x, hxrt.StringFromLiteral("a")), hxrt.StringFromLiteral(":"))
		argIndex := 0
		var hx_if_475 []*string
		if x.exists(hxrt.StringFromLiteral("v")) {
			hx_if_475 = self.splitString(self.requireAttr(x, hxrt.StringFromLiteral("v")), hxrt.StringFromLiteral(":"))
		} else {
			hx_if_475 = nil
		}
		evalues := hx_if_475
		valueIndex := 0
		e := x.elements()
		for func(hx_obj_476 map[string]any) func() bool {
			hx_field_477 := hx_obj_476["hasNext"]
			if hx_field_477 == nil {
				var hx_zero_478 func() bool
				return hx_zero_478
			}
			return hx_field_477.(func() bool)
		}(e)() {
			e_1 := func(hx_obj_479 map[string]any) func() *Xml {
				hx_field_480 := hx_obj_479["next"]
				if hx_field_480 == nil {
					var hx_zero_481 func() *Xml
					return hx_zero_481
				}
				return hx_field_480.(func() *Xml)
			}(e)()
			opt := false
			var hx_if_482 *string
			if argIndex < len(aname) {
				hx_if_482 = aname[argIndex]
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
			var hx_if_484 *string
			if (evalues == nil) || (valueIndex >= len(evalues)) {
				hx_if_484 = nil
			} else {
				hx_post_483 := valueIndex
				valueIndex = int(int32((valueIndex + 1)))
				hx_if_484 = evalues[hx_post_483]
			}
			v := hx_if_484
			args = append(args, func() map[string]any {
				hx_obj_486 := map[string]any{}
				hx_obj_486["name"] = a
				hx_obj_486["opt"] = opt
				hx_obj_486["t"] = self.xtype(e_1)
				var hx_if_487 *string
				if hxrt.StringEqualStringPtr(v, hxrt.StringFromLiteral("")) {
					hx_if_487 = nil
				} else {
					hx_if_487 = v
				}
				hx_obj_486["value"] = hx_if_487
				return hx_obj_486
			}())
		}
		ret := args[int(int32((hxrt.Int32Wrap(len(args)) - hxrt.Int32Wrap(1))))]
		callArgs := []map[string]any{}
		_g := 0
		_g1 := int(int32((hxrt.Int32Wrap(len(args)) - hxrt.Int32Wrap(1))))
		for _g < _g1 {
			hx_post_488 := _g
			_g = int(int32((_g + 1)))
			i := hx_post_488
			callArgs = append(callArgs, args[i])
		}
		return haxe__rtti__CType_CFunction(callArgs, func(hx_obj_490 map[string]any) *haxe__rtti__CType {
			hx_field_491 := hx_obj_490["t"]
			if hx_field_491 == nil {
				var hx_zero_492 *haxe__rtti__CType
				return hx_zero_492
			}
			return hx_field_491.(*haxe__rtti__CType)
		}(ret))
	}
	if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("a")) {
		fields := []map[string]any{}
		f := x.elements()
		for func(hx_obj_493 map[string]any) func() bool {
			hx_field_494 := hx_obj_493["hasNext"]
			if hx_field_494 == nil {
				var hx_zero_495 func() bool
				return hx_zero_495
			}
			return hx_field_494.(func() bool)
		}(f)() {
			f_1 := func(hx_obj_496 map[string]any) func() *Xml {
				hx_field_497 := hx_obj_496["next"]
				if hx_field_497 == nil {
					var hx_zero_498 func() *Xml
					return hx_zero_498
				}
				return hx_field_497.(func() *Xml)
			}(f)()
			f_2 := self.xclassfield(f_1, true)
			f_2["platforms"] = []*string{}
			fields = append(fields, f_2)
		}
		return haxe__rtti__CType_CAnonymous(fields)
	}
	if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("d")) {
		var t *haxe__rtti__CType = nil
		tx := x.firstElement()
		if tx != nil {
			t = self.xtype(tx)
		}
		return haxe__rtti__CType_CDynamic(t)
	}
	hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Invalid "), nodeName))
	var hx_throw_zero_500 *haxe__rtti__CType
	return hx_throw_zero_500
}

func (self *haxe__rtti__XmlParser) xtypeparams(x *Xml) []*haxe__rtti__CType {
	p := []*haxe__rtti__CType{}
	c := x.elements()
	for func(hx_obj_501 map[string]any) func() bool {
		hx_field_502 := hx_obj_501["hasNext"]
		if hx_field_502 == nil {
			var hx_zero_503 func() bool
			return hx_zero_503
		}
		return hx_field_502.(func() bool)
	}(c)() {
		c_1 := func(hx_obj_504 map[string]any) func() *Xml {
			hx_field_505 := hx_obj_504["next"]
			if hx_field_505 == nil {
				var hx_zero_506 func() *Xml
				return hx_zero_506
			}
			return hx_field_505.(func() *Xml)
		}(c)()
		p = append(p, self.xtype(c_1))
	}
	return p
}

func (self *haxe__rtti__XmlParser) defplat() []*string {
	l := []*string{}
	if !hxrt.StringEqualStringPtr(self.curplatform, nil) {
		l = append(l, self.curplatform)
	}
	return l
}

func (self *haxe__rtti__XmlParser) joinStringArray(values []*string, separator *string) *string {
	var buf_b *string
	buf_b = hxrt.StringFromLiteral("")
	_g := 0
	_g1 := len(values)
	for _g < _g1 {
		hx_post_509 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_509
		if i > 0 {
			buf_b = hxrt.StringConcatStringPtr(buf_b, hxrt.StdString(separator))
		}
		buf_b = hxrt.StringConcatStringPtr(buf_b, hxrt.StdString(values[i]))
	}
	return buf_b
}

func (self *haxe__rtti__XmlParser) splitString(value *string, separator *string) []*string {
	if hxrt.StringEqualStringPtr(separator, hxrt.StringFromLiteral("")) {
		return []*string{value}
	}
	parts := []*string{}
	start := 0
	for true {
		index := self.findSeparator(value, separator, start)
		if index == -1 {
			parts = append(parts, hxrt.StringSubstrStringPtr(value, start, 0, false))
			break
		}
		parts = append(parts, hxrt.StringSubstrStringPtr(value, start, int(int32((hxrt.Int32Wrap(index)-hxrt.Int32Wrap(start)))), true))
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
	value := x.get(name)
	var hx_if_512 *string
	if hxrt.StringEqualStringPtr(value, nil) {
		hx_if_512 = hxrt.StringFromLiteral("")
	} else {
		hx_if_512 = value
	}
	return hx_if_512
}

func (self *haxe__rtti__XmlParser) hasNamedElement(x *Xml, name *string) bool {
	return func(hx_obj_513 map[string]any) func() bool {
		hx_field_514 := hx_obj_513["hasNext"]
		if hx_field_514 == nil {
			var hx_zero_515 func() bool
			return hx_zero_515
		}
		return hx_field_514.(func() bool)
	}(x.elementsNamed(name))()
}

func (self *haxe__rtti__XmlParser) requireNamedElement(x *Xml, name *string) *Xml {
	elements := x.elementsNamed(name)
	if !func(hx_obj_517 map[string]any) func() bool {
		hx_field_518 := hx_obj_517["hasNext"]
		if hx_field_518 == nil {
			var hx_zero_519 func() bool
			return hx_zero_519
		}
		return hx_field_518.(func() bool)
	}(elements)() {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(self.nodeDisplayName(x), hxrt.StringFromLiteral(" is missing element ")), name))
		var hx_throw_zero_516 *Xml
		return hx_throw_zero_516
	}
	return func(hx_obj_520 map[string]any) func() *Xml {
		hx_field_521 := hx_obj_520["next"]
		if hx_field_521 == nil {
			var hx_zero_522 func() *Xml
			return hx_zero_522
		}
		return hx_field_521.(func() *Xml)
	}(elements)()
}

func (self *haxe__rtti__XmlParser) requireFirstElement(x *Xml) *Xml {
	first := x.firstElement()
	if first == nil {
		hxrt.Throw(hxrt.StringConcatStringPtr(self.nodeDisplayName(x), hxrt.StringFromLiteral(" is missing first element")))
		var hx_throw_zero_523 *Xml
		return hx_throw_zero_523
	}
	return first
}

func (self *haxe__rtti__XmlParser) nodeDisplayName(x *Xml) *string {
	var hx_if_524 *string
	if x.nodeType == Xml_Document {
		hx_if_524 = hxrt.StringFromLiteral("Document")
	} else {
		hx_if_524 = self.elementName(x)
	}
	return hx_if_524
}

func (self *haxe__rtti__XmlParser) elementName(x *Xml) *string {
	name := x.get_nodeName()
	var hx_if_525 *string
	if hxrt.StringEqualStringPtr(name, nil) {
		hx_if_525 = hxrt.StringFromLiteral("")
	} else {
		hx_if_525 = name
	}
	return hx_if_525
}

func (self *haxe__rtti__XmlParser) innerData(x *Xml) *string {
	it := x.iterator()
	if !func(hx_obj_527 map[string]any) func() bool {
		hx_field_528 := hx_obj_527["hasNext"]
		if hx_field_528 == nil {
			var hx_zero_529 func() bool
			return hx_zero_529
		}
		return hx_field_528.(func() bool)
	}(it)() {
		hxrt.Throw(hxrt.StringConcatStringPtr(self.nodeDisplayName(x), hxrt.StringFromLiteral(" does not have data")))
		var hx_throw_zero_526 *string
		return hx_throw_zero_526
	}
	value := func(hx_obj_530 map[string]any) func() *Xml {
		hx_field_531 := hx_obj_530["next"]
		if hx_field_531 == nil {
			var hx_zero_532 func() *Xml
			return hx_zero_532
		}
		return hx_field_531.(func() *Xml)
	}(it)()
	if func(hx_obj_534 map[string]any) func() bool {
		hx_field_535 := hx_obj_534["hasNext"]
		if hx_field_535 == nil {
			var hx_zero_536 func() bool
			return hx_zero_536
		}
		return hx_field_535.(func() bool)
	}(it)() {
		hxrt.Throw(hxrt.StringConcatStringPtr(self.nodeDisplayName(x), hxrt.StringFromLiteral(" does not only have data")))
		var hx_throw_zero_533 *string
		return hx_throw_zero_533
	}
	if (value.nodeType != Xml_PCData) && (value.nodeType != Xml_CData) {
		hxrt.Throw(hxrt.StringConcatStringPtr(self.nodeDisplayName(x), hxrt.StringFromLiteral(" does not have data")))
		var hx_throw_zero_537 *string
		return hx_throw_zero_537
	}
	return value.get_nodeValue()
}

func (self *haxe__rtti__XmlParser) innerHTML(x *Xml) *string {
	var buf_b *string
	buf_b = hxrt.StringFromLiteral("")
	child := x.iterator()
	for func(hx_obj_538 map[string]any) func() bool {
		hx_field_539 := hx_obj_538["hasNext"]
		if hx_field_539 == nil {
			var hx_zero_540 func() bool
			return hx_zero_540
		}
		return hx_field_539.(func() bool)
	}(child)() {
		child_1 := func(hx_obj_541 map[string]any) func() *Xml {
			hx_field_542 := hx_obj_541["next"]
			if hx_field_542 == nil {
				var hx_zero_543 func() *Xml
				return hx_zero_543
			}
			return hx_field_542.(func() *Xml)
		}(child)()
		x_1 := child_1.toString()
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
	var hx_if_544 int
	if negative {
		hx_if_544 = int(int32(-int32(result)))
	} else {
		hx_if_544 = result
	}
	return hx_if_544
}
