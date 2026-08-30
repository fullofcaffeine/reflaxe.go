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
	haxe__ds__ArraySort_sort(l, func(hx_cmp_left_9 any, hx_cmp_right_10 any) int {
		return func(e1 *haxe__rtti__TypeTree, e2 *haxe__rtti__TypeTree) int {
			var hx_if_4 *string
			if e1.tag == 0 {
				_g := e1.params[0].(*string)
				_g1 := e1.params[1].(*string)
				_ = _g1
				_g1_1 := e1.params[2].(*hxrt.Array)
				_ = _g1_1
				p := _g
				hx_if_4 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(" "), p)
			} else {
				hx_if_4 = func(hx_obj_1 map[string]any) *string {
					hx_field_2 := hx_obj_1["path"]
					if hx_field_2 == nil {
						var hx_zero_3 *string
						return hx_zero_3
					}
					return hx_field_2.(*string)
				}(haxe__rtti__TypeApi_typeInfos(e1))
			}
			n1 := hx_if_4
			var hx_if_8 *string
			if e2.tag == 0 {
				_g_1 := e2.params[0].(*string)
				_g1_2 := e2.params[1].(*string)
				_ = _g1_2
				_g1_3 := e2.params[2].(*hxrt.Array)
				_ = _g1_3
				p_1 := _g_1
				hx_if_8 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(" "), p_1)
			} else {
				hx_if_8 = func(hx_obj_5 map[string]any) *string {
					hx_field_6 := hx_obj_5["path"]
					if hx_field_6 == nil {
						var hx_zero_7 *string
						return hx_zero_7
					}
					return hx_field_6.(*string)
				}(haxe__rtti__TypeApi_typeInfos(e2))
			}
			n2 := hx_if_8
			return Reflect_compare(n1, n2)
		}(func(hx_value_11 any) *haxe__rtti__TypeTree {
			if hx_value_11 == nil {
				var hx_zero_12 *haxe__rtti__TypeTree
				return hx_zero_12
			}
			return hx_value_11.(*haxe__rtti__TypeTree)
		}(hx_cmp_left_9), func(hx_value_13 any) *haxe__rtti__TypeTree {
			if hx_value_13 == nil {
				var hx_zero_14 *haxe__rtti__TypeTree
				return hx_zero_14
			}
			return hx_value_13.(*haxe__rtti__TypeTree)
		}(hx_cmp_right_10))
	})
	_g := 0
	for _g < l.Len() {
		x := func(hx_value_15 any) *haxe__rtti__TypeTree {
			if hx_value_15 == nil {
				var hx_zero_16 *haxe__rtti__TypeTree
				return hx_zero_16
			}
			return hx_value_15.(*haxe__rtti__TypeTree)
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
			self.__hx_this.sortFields(func(hx_obj_17 map[string]any) *hxrt.Array {
				hx_field_18 := hx_obj_17["fields"]
				if hx_field_18 == nil {
					var hx_zero_19 *hxrt.Array
					return hx_zero_19
				}
				return hx_field_18.(*hxrt.Array)
			}(c))
			self.__hx_this.sortFields(func(hx_obj_20 map[string]any) *hxrt.Array {
				hx_field_21 := hx_obj_20["statics"]
				if hx_field_21 == nil {
					var hx_zero_22 *hxrt.Array
					return hx_zero_22
				}
				return hx_field_21.(*hxrt.Array)
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
	haxe__ds__ArraySort_sort(a, func(hx_cmp_left_41 any, hx_cmp_right_42 any) int {
		return func(f1 map[string]any, f2 map[string]any) int {
			v1 := haxe__rtti__TypeApi_isVar(func(hx_obj_23 map[string]any) *haxe__rtti__CType {
				hx_field_24 := hx_obj_23["type"]
				if hx_field_24 == nil {
					var hx_zero_25 *haxe__rtti__CType
					return hx_zero_25
				}
				return hx_field_24.(*haxe__rtti__CType)
			}(f1))
			v2 := haxe__rtti__TypeApi_isVar(func(hx_obj_26 map[string]any) *haxe__rtti__CType {
				hx_field_27 := hx_obj_26["type"]
				if hx_field_27 == nil {
					var hx_zero_28 *haxe__rtti__CType
					return hx_zero_28
				}
				return hx_field_27.(*haxe__rtti__CType)
			}(f2))
			if v1 && !v2 {
				return -1
			}
			if v2 && !v1 {
				return 1
			}
			if hxrt.StringEqualStringPtr(func(hx_obj_29 map[string]any) *string {
				hx_field_30 := hx_obj_29["name"]
				if hx_field_30 == nil {
					var hx_zero_31 *string
					return hx_zero_31
				}
				return hx_field_30.(*string)
			}(f1), hxrt.StringFromLiteral("new")) {
				return -1
			}
			if hxrt.StringEqualStringPtr(func(hx_obj_32 map[string]any) *string {
				hx_field_33 := hx_obj_32["name"]
				if hx_field_33 == nil {
					var hx_zero_34 *string
					return hx_zero_34
				}
				return hx_field_33.(*string)
			}(f2), hxrt.StringFromLiteral("new")) {
				return 1
			}
			return Reflect_compare(func(hx_obj_35 map[string]any) *string {
				hx_field_36 := hx_obj_35["name"]
				if hx_field_36 == nil {
					var hx_zero_37 *string
					return hx_zero_37
				}
				return hx_field_36.(*string)
			}(f1), func(hx_obj_38 map[string]any) *string {
				hx_field_39 := hx_obj_38["name"]
				if hx_field_39 == nil {
					var hx_zero_40 *string
					return hx_zero_40
				}
				return hx_field_39.(*string)
			}(f2))
		}(func(hx_value_43 any) map[string]any {
			if hx_value_43 == nil {
				var hx_zero_44 map[string]any
				return hx_zero_44
			}
			return hx_value_43.(map[string]any)
		}(hx_cmp_left_41), func(hx_value_45 any) map[string]any {
			if hx_value_45 == nil {
				var hx_zero_46 map[string]any
				return hx_zero_46
			}
			return hx_value_45.(map[string]any)
		}(hx_cmp_right_42))
	})
}

func (self *haxe__rtti__XmlParser) process(x *Xml, platform *string) {
	self.curplatform = platform
	self.__hx_this.xroot(x)
}

func (self *haxe__rtti__XmlParser) mergeRights(f1 map[string]any, f2 map[string]any) bool {
	if (((func(hx_obj_47 map[string]any) *haxe__rtti__Rights {
		hx_field_48 := hx_obj_47["get"]
		if hx_field_48 == nil {
			var hx_zero_49 *haxe__rtti__Rights
			return hx_zero_49
		}
		return hx_field_48.(*haxe__rtti__Rights)
	}(f1) == haxe__rtti__Rights_RInline) && (func(hx_obj_50 map[string]any) *haxe__rtti__Rights {
		hx_field_51 := hx_obj_50["set"]
		if hx_field_51 == nil {
			var hx_zero_52 *haxe__rtti__Rights
			return hx_zero_52
		}
		return hx_field_51.(*haxe__rtti__Rights)
	}(f1) == haxe__rtti__Rights_RNo)) && (func(hx_obj_53 map[string]any) *haxe__rtti__Rights {
		hx_field_54 := hx_obj_53["get"]
		if hx_field_54 == nil {
			var hx_zero_55 *haxe__rtti__Rights
			return hx_zero_55
		}
		return hx_field_54.(*haxe__rtti__Rights)
	}(f2) == haxe__rtti__Rights_RNormal)) && (func(hx_obj_56 map[string]any) *haxe__rtti__Rights {
		hx_field_57 := hx_obj_56["set"]
		if hx_field_57 == nil {
			var hx_zero_58 *haxe__rtti__Rights
			return hx_zero_58
		}
		return hx_field_57.(*haxe__rtti__Rights)
	}(f2) == haxe__rtti__Rights_RMethod) {
		f1["get"] = haxe__rtti__Rights_RNormal
		f1["set"] = haxe__rtti__Rights_RMethod
		return true
	}
	return (Type_enumEq(func(hx_obj_59 map[string]any) *haxe__rtti__Rights {
		hx_field_60 := hx_obj_59["get"]
		if hx_field_60 == nil {
			var hx_zero_61 *haxe__rtti__Rights
			return hx_zero_61
		}
		return hx_field_60.(*haxe__rtti__Rights)
	}(f1), func(hx_obj_62 map[string]any) *haxe__rtti__Rights {
		hx_field_63 := hx_obj_62["get"]
		if hx_field_63 == nil {
			var hx_zero_64 *haxe__rtti__Rights
			return hx_zero_64
		}
		return hx_field_63.(*haxe__rtti__Rights)
	}(f2)) && Type_enumEq(func(hx_obj_65 map[string]any) *haxe__rtti__Rights {
		hx_field_66 := hx_obj_65["set"]
		if hx_field_66 == nil {
			var hx_zero_67 *haxe__rtti__Rights
			return hx_zero_67
		}
		return hx_field_66.(*haxe__rtti__Rights)
	}(f1), func(hx_obj_68 map[string]any) *haxe__rtti__Rights {
		hx_field_69 := hx_obj_68["set"]
		if hx_field_69 == nil {
			var hx_zero_70 *haxe__rtti__Rights
			return hx_zero_70
		}
		return hx_field_69.(*haxe__rtti__Rights)
	}(f2)))
}

func (self *haxe__rtti__XmlParser) mergeDoc(f1 map[string]any, f2 map[string]any) bool {
	if hxrt.StringEqualStringPtr(func(hx_obj_80 map[string]any) *string {
		hx_field_81 := hx_obj_80["doc"]
		if hx_field_81 == nil {
			var hx_zero_82 *string
			return hx_zero_82
		}
		return hx_field_81.(*string)
	}(f1), nil) {
		f1["doc"] = func(hx_obj_71 map[string]any) *string {
			hx_field_72 := hx_obj_71["doc"]
			if hx_field_72 == nil {
				var hx_zero_73 *string
				return hx_zero_73
			}
			return hx_field_72.(*string)
		}(f2)
	} else {
		if hxrt.StringEqualStringPtr(func(hx_obj_77 map[string]any) *string {
			hx_field_78 := hx_obj_77["doc"]
			if hx_field_78 == nil {
				var hx_zero_79 *string
				return hx_zero_79
			}
			return hx_field_78.(*string)
		}(f2), nil) {
			f2["doc"] = func(hx_obj_74 map[string]any) *string {
				hx_field_75 := hx_obj_74["doc"]
				if hx_field_75 == nil {
					var hx_zero_76 *string
					return hx_zero_76
				}
				return hx_field_75.(*string)
			}(f1)
		}
	}
	return true
}

func (self *haxe__rtti__XmlParser) mergeFields(f map[string]any, f2 map[string]any) bool {
	return (haxe__rtti__TypeApi_fieldEq(f, f2) || (((hxrt.StringEqualStringPtr(func(hx_obj_83 map[string]any) *string {
		hx_field_84 := hx_obj_83["name"]
		if hx_field_84 == nil {
			var hx_zero_85 *string
			return hx_zero_85
		}
		return hx_field_84.(*string)
	}(f), func(hx_obj_86 map[string]any) *string {
		hx_field_87 := hx_obj_86["name"]
		if hx_field_87 == nil {
			var hx_zero_88 *string
			return hx_zero_88
		}
		return hx_field_87.(*string)
	}(f2)) && (self.__hx_this.mergeRights(f, f2) || self.__hx_this.mergeRights(f2, f))) && self.__hx_this.mergeDoc(f, f2)) && haxe__rtti__TypeApi_fieldEq(f, f2)))
}

func (self *haxe__rtti__XmlParser) mergeClasses(c map[string]any, c2 map[string]any) bool {
	if func(hx_obj_89 map[string]any) bool {
		hx_field_90 := hx_obj_89["isInterface"]
		if hx_field_90 == nil {
			var hx_zero_91 bool
			return hx_zero_91
		}
		return hx_field_90.(bool)
	}(c) != func(hx_obj_92 map[string]any) bool {
		hx_field_93 := hx_obj_92["isInterface"]
		if hx_field_93 == nil {
			var hx_zero_94 bool
			return hx_zero_94
		}
		return hx_field_93.(bool)
	}(c2) {
		return false
	}
	if !hxrt.StringEqualStringPtr(self.curplatform, nil) {
		hx_arr_95 := func(hx_obj_96 map[string]any) *hxrt.Array {
			hx_field_97 := hx_obj_96["platforms"]
			if hx_field_97 == nil {
				var hx_zero_98 *hxrt.Array
				return hx_zero_98
			}
			return hx_field_97.(*hxrt.Array)
		}(c)
		hx_arr_95.Push(self.curplatform)
	}
	if func(hx_obj_99 map[string]any) bool {
		hx_field_100 := hx_obj_99["isExtern"]
		if hx_field_100 == nil {
			var hx_zero_101 bool
			return hx_zero_101
		}
		return hx_field_100.(bool)
	}(c) != func(hx_obj_102 map[string]any) bool {
		hx_field_103 := hx_obj_102["isExtern"]
		if hx_field_103 == nil {
			var hx_zero_104 bool
			return hx_zero_104
		}
		return hx_field_103.(bool)
	}(c2) {
		c["isExtern"] = false
	}
	_g := 0
	_g1 := func(hx_obj_105 map[string]any) *hxrt.Array {
		hx_field_106 := hx_obj_105["fields"]
		if hx_field_106 == nil {
			var hx_zero_107 *hxrt.Array
			return hx_zero_107
		}
		return hx_field_106.(*hxrt.Array)
	}(c2)
	for _g < _g1.Len() {
		f2 := func(hx_value_108 any) map[string]any {
			if hx_value_108 == nil {
				var hx_zero_109 map[string]any
				return hx_zero_109
			}
			return hx_value_108.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		var found map[string]any = nil
		_g_1 := 0
		_g1_1 := func(hx_obj_110 map[string]any) *hxrt.Array {
			hx_field_111 := hx_obj_110["fields"]
			if hx_field_111 == nil {
				var hx_zero_112 *hxrt.Array
				return hx_zero_112
			}
			return hx_field_111.(*hxrt.Array)
		}(c)
		for _g_1 < _g1_1.Len() {
			f := func(hx_value_113 any) map[string]any {
				if hx_value_113 == nil {
					var hx_zero_114 map[string]any
					return hx_zero_114
				}
				return hx_value_113.(map[string]any)
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
			hx_arr_115 := func(hx_obj_116 map[string]any) *hxrt.Array {
				hx_field_117 := hx_obj_116["fields"]
				if hx_field_117 == nil {
					var hx_zero_118 *hxrt.Array
					return hx_zero_118
				}
				return hx_field_117.(*hxrt.Array)
			}(c)
			hx_arr_115.Push(f2)
		} else {
			if !hxrt.StringEqualStringPtr(self.curplatform, nil) {
				hx_arr_119 := func(hx_obj_120 map[string]any) *hxrt.Array {
					hx_field_121 := hx_obj_120["platforms"]
					if hx_field_121 == nil {
						var hx_zero_122 *hxrt.Array
						return hx_zero_122
					}
					return hx_field_121.(*hxrt.Array)
				}(found)
				hx_arr_119.Push(self.curplatform)
			}
		}
	}
	_g_2 := 0
	_g1_2 := func(hx_obj_123 map[string]any) *hxrt.Array {
		hx_field_124 := hx_obj_123["statics"]
		if hx_field_124 == nil {
			var hx_zero_125 *hxrt.Array
			return hx_zero_125
		}
		return hx_field_124.(*hxrt.Array)
	}(c2)
	for _g_2 < _g1_2.Len() {
		f2_1 := func(hx_value_126 any) map[string]any {
			if hx_value_126 == nil {
				var hx_zero_127 map[string]any
				return hx_zero_127
			}
			return hx_value_126.(map[string]any)
		}(_g1_2.Get(_g_2))
		_g_2 = int(int32((_g_2 + 1)))
		var found_1 map[string]any = nil
		_g_3 := 0
		_g1_3 := func(hx_obj_128 map[string]any) *hxrt.Array {
			hx_field_129 := hx_obj_128["statics"]
			if hx_field_129 == nil {
				var hx_zero_130 *hxrt.Array
				return hx_zero_130
			}
			return hx_field_129.(*hxrt.Array)
		}(c)
		for _g_3 < _g1_3.Len() {
			f_1 := func(hx_value_131 any) map[string]any {
				if hx_value_131 == nil {
					var hx_zero_132 map[string]any
					return hx_zero_132
				}
				return hx_value_131.(map[string]any)
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
			hx_arr_133 := func(hx_obj_134 map[string]any) *hxrt.Array {
				hx_field_135 := hx_obj_134["statics"]
				if hx_field_135 == nil {
					var hx_zero_136 *hxrt.Array
					return hx_zero_136
				}
				return hx_field_135.(*hxrt.Array)
			}(c)
			hx_arr_133.Push(f2_1)
		} else {
			if !hxrt.StringEqualStringPtr(self.curplatform, nil) {
				hx_arr_137 := func(hx_obj_138 map[string]any) *hxrt.Array {
					hx_field_139 := hx_obj_138["platforms"]
					if hx_field_139 == nil {
						var hx_zero_140 *hxrt.Array
						return hx_zero_140
					}
					return hx_field_139.(*hxrt.Array)
				}(found_1)
				hx_arr_137.Push(self.curplatform)
			}
		}
	}
	return true
}

func (self *haxe__rtti__XmlParser) mergeEnums(e map[string]any, e2 map[string]any) bool {
	if func(hx_obj_141 map[string]any) bool {
		hx_field_142 := hx_obj_141["isExtern"]
		if hx_field_142 == nil {
			var hx_zero_143 bool
			return hx_zero_143
		}
		return hx_field_142.(bool)
	}(e) != func(hx_obj_144 map[string]any) bool {
		hx_field_145 := hx_obj_144["isExtern"]
		if hx_field_145 == nil {
			var hx_zero_146 bool
			return hx_zero_146
		}
		return hx_field_145.(bool)
	}(e2) {
		return false
	}
	if !hxrt.StringEqualStringPtr(self.curplatform, nil) {
		hx_arr_147 := func(hx_obj_148 map[string]any) *hxrt.Array {
			hx_field_149 := hx_obj_148["platforms"]
			if hx_field_149 == nil {
				var hx_zero_150 *hxrt.Array
				return hx_zero_150
			}
			return hx_field_149.(*hxrt.Array)
		}(e)
		hx_arr_147.Push(self.curplatform)
	}
	_g := 0
	_g1 := func(hx_obj_151 map[string]any) *hxrt.Array {
		hx_field_152 := hx_obj_151["constructors"]
		if hx_field_152 == nil {
			var hx_zero_153 *hxrt.Array
			return hx_zero_153
		}
		return hx_field_152.(*hxrt.Array)
	}(e2)
	for _g < _g1.Len() {
		c2 := func(hx_value_154 any) map[string]any {
			if hx_value_154 == nil {
				var hx_zero_155 map[string]any
				return hx_zero_155
			}
			return hx_value_154.(map[string]any)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		var found map[string]any = nil
		_g_1 := 0
		_g1_1 := func(hx_obj_156 map[string]any) *hxrt.Array {
			hx_field_157 := hx_obj_156["constructors"]
			if hx_field_157 == nil {
				var hx_zero_158 *hxrt.Array
				return hx_zero_158
			}
			return hx_field_157.(*hxrt.Array)
		}(e)
		for _g_1 < _g1_1.Len() {
			c := func(hx_value_159 any) map[string]any {
				if hx_value_159 == nil {
					var hx_zero_160 map[string]any
					return hx_zero_160
				}
				return hx_value_159.(map[string]any)
			}(_g1_1.Get(_g_1))
			_g_1 = int(int32((_g_1 + 1)))
			if haxe__rtti__TypeApi_constructorEq(c, c2) {
				found = c
				break
			}
		}
		if found == nil {
			hx_arr_161 := func(hx_obj_162 map[string]any) *hxrt.Array {
				hx_field_163 := hx_obj_162["constructors"]
				if hx_field_163 == nil {
					var hx_zero_164 *hxrt.Array
					return hx_zero_164
				}
				return hx_field_163.(*hxrt.Array)
			}(e)
			hx_arr_161.Push(c2)
		} else {
			if !hxrt.StringEqualStringPtr(self.curplatform, nil) {
				hx_arr_165 := func(hx_obj_166 map[string]any) *hxrt.Array {
					hx_field_167 := hx_obj_166["platforms"]
					if hx_field_167 == nil {
						var hx_zero_168 *hxrt.Array
						return hx_zero_168
					}
					return hx_field_167.(*hxrt.Array)
				}(found)
				hx_arr_165.Push(self.curplatform)
			}
		}
	}
	return true
}

func (self *haxe__rtti__XmlParser) mergeTypedefs(t map[string]any, t2 map[string]any) bool {
	if hxrt.StringEqualStringPtr(self.curplatform, nil) {
		return false
	}
	hx_arr_169 := func(hx_obj_170 map[string]any) *hxrt.Array {
		hx_field_171 := hx_obj_170["platforms"]
		if hx_field_171 == nil {
			var hx_zero_172 *hxrt.Array
			return hx_zero_172
		}
		return hx_field_171.(*hxrt.Array)
	}(t)
	hx_arr_169.Push(self.curplatform)
	var this1 haxe__IMap = func(hx_obj_173 map[string]any) *haxe__ds__StringMap {
		hx_field_174 := hx_obj_173["types"]
		if hx_field_174 == nil {
			var hx_zero_175 *haxe__ds__StringMap
			return hx_zero_175
		}
		return hx_field_174.(*haxe__ds__StringMap)
	}(t)
	key := self.curplatform
	value := func(hx_obj_176 map[string]any) *haxe__rtti__CType {
		hx_field_177 := hx_obj_176["type"]
		if hx_field_177 == nil {
			var hx_zero_178 *haxe__rtti__CType
			return hx_zero_178
		}
		return hx_field_177.(*haxe__rtti__CType)
	}(t2)
	this1.(*haxe__ds__StringMap).__hx_this.set(key, value)
	return true
}

func (self *haxe__rtti__XmlParser) mergeAbstracts(a map[string]any, a2 map[string]any) bool {
	if hxrt.StringEqualStringPtr(self.curplatform, nil) {
		return false
	}
	if (func(hx_obj_179 map[string]any) *hxrt.Array {
		hx_field_180 := hx_obj_179["to"]
		if hx_field_180 == nil {
			var hx_zero_181 *hxrt.Array
			return hx_zero_181
		}
		return hx_field_180.(*hxrt.Array)
	}(a).Len() != func(hx_obj_182 map[string]any) *hxrt.Array {
		hx_field_183 := hx_obj_182["to"]
		if hx_field_183 == nil {
			var hx_zero_184 *hxrt.Array
			return hx_zero_184
		}
		return hx_field_183.(*hxrt.Array)
	}(a2).Len()) || (func(hx_obj_185 map[string]any) *hxrt.Array {
		hx_field_186 := hx_obj_185["from"]
		if hx_field_186 == nil {
			var hx_zero_187 *hxrt.Array
			return hx_zero_187
		}
		return hx_field_186.(*hxrt.Array)
	}(a).Len() != func(hx_obj_188 map[string]any) *hxrt.Array {
		hx_field_189 := hx_obj_188["from"]
		if hx_field_189 == nil {
			var hx_zero_190 *hxrt.Array
			return hx_zero_190
		}
		return hx_field_189.(*hxrt.Array)
	}(a2).Len()) {
		return false
	}
	_g := 0
	_g1 := func(hx_obj_191 map[string]any) *hxrt.Array {
		hx_field_192 := hx_obj_191["to"]
		if hx_field_192 == nil {
			var hx_zero_193 *hxrt.Array
			return hx_zero_193
		}
		return hx_field_192.(*hxrt.Array)
	}(a).Len()
	for _g < _g1 {
		hx_post_194 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_194
		if !haxe__rtti__TypeApi_typeEq(func(hx_obj_200 map[string]any) *haxe__rtti__CType {
			hx_field_201 := hx_obj_200["t"]
			if hx_field_201 == nil {
				var hx_zero_202 *haxe__rtti__CType
				return hx_zero_202
			}
			return hx_field_201.(*haxe__rtti__CType)
		}(func(hx_value_198 any) map[string]any {
			if hx_value_198 == nil {
				var hx_zero_199 map[string]any
				return hx_zero_199
			}
			return hx_value_198.(map[string]any)
		}(func(hx_obj_195 map[string]any) *hxrt.Array {
			hx_field_196 := hx_obj_195["to"]
			if hx_field_196 == nil {
				var hx_zero_197 *hxrt.Array
				return hx_zero_197
			}
			return hx_field_196.(*hxrt.Array)
		}(a).Get(i))), func(hx_obj_208 map[string]any) *haxe__rtti__CType {
			hx_field_209 := hx_obj_208["t"]
			if hx_field_209 == nil {
				var hx_zero_210 *haxe__rtti__CType
				return hx_zero_210
			}
			return hx_field_209.(*haxe__rtti__CType)
		}(func(hx_value_206 any) map[string]any {
			if hx_value_206 == nil {
				var hx_zero_207 map[string]any
				return hx_zero_207
			}
			return hx_value_206.(map[string]any)
		}(func(hx_obj_203 map[string]any) *hxrt.Array {
			hx_field_204 := hx_obj_203["to"]
			if hx_field_204 == nil {
				var hx_zero_205 *hxrt.Array
				return hx_zero_205
			}
			return hx_field_204.(*hxrt.Array)
		}(a2).Get(i)))) {
			return false
		}
	}
	_g_1 := 0
	_g1_1 := func(hx_obj_211 map[string]any) *hxrt.Array {
		hx_field_212 := hx_obj_211["from"]
		if hx_field_212 == nil {
			var hx_zero_213 *hxrt.Array
			return hx_zero_213
		}
		return hx_field_212.(*hxrt.Array)
	}(a).Len()
	for _g_1 < _g1_1 {
		hx_post_214 := _g_1
		_g_1 = int(int32((_g_1 + 1)))
		i_1 := hx_post_214
		if !haxe__rtti__TypeApi_typeEq(func(hx_obj_220 map[string]any) *haxe__rtti__CType {
			hx_field_221 := hx_obj_220["t"]
			if hx_field_221 == nil {
				var hx_zero_222 *haxe__rtti__CType
				return hx_zero_222
			}
			return hx_field_221.(*haxe__rtti__CType)
		}(func(hx_value_218 any) map[string]any {
			if hx_value_218 == nil {
				var hx_zero_219 map[string]any
				return hx_zero_219
			}
			return hx_value_218.(map[string]any)
		}(func(hx_obj_215 map[string]any) *hxrt.Array {
			hx_field_216 := hx_obj_215["from"]
			if hx_field_216 == nil {
				var hx_zero_217 *hxrt.Array
				return hx_zero_217
			}
			return hx_field_216.(*hxrt.Array)
		}(a).Get(i_1))), func(hx_obj_228 map[string]any) *haxe__rtti__CType {
			hx_field_229 := hx_obj_228["t"]
			if hx_field_229 == nil {
				var hx_zero_230 *haxe__rtti__CType
				return hx_zero_230
			}
			return hx_field_229.(*haxe__rtti__CType)
		}(func(hx_value_226 any) map[string]any {
			if hx_value_226 == nil {
				var hx_zero_227 map[string]any
				return hx_zero_227
			}
			return hx_value_226.(map[string]any)
		}(func(hx_obj_223 map[string]any) *hxrt.Array {
			hx_field_224 := hx_obj_223["from"]
			if hx_field_224 == nil {
				var hx_zero_225 *hxrt.Array
				return hx_zero_225
			}
			return hx_field_224.(*hxrt.Array)
		}(a2).Get(i_1)))) {
			return false
		}
	}
	if func(hx_obj_237 map[string]any) map[string]any {
		hx_field_238 := hx_obj_237["impl"]
		if hx_field_238 == nil {
			var hx_zero_239 map[string]any
			return hx_zero_239
		}
		return hx_field_238.(map[string]any)
	}(a2) != nil {
		self.__hx_this.mergeClasses(func(hx_obj_231 map[string]any) map[string]any {
			hx_field_232 := hx_obj_231["impl"]
			if hx_field_232 == nil {
				var hx_zero_233 map[string]any
				return hx_zero_233
			}
			return hx_field_232.(map[string]any)
		}(a), func(hx_obj_234 map[string]any) map[string]any {
			hx_field_235 := hx_obj_234["impl"]
			if hx_field_235 == nil {
				var hx_zero_236 map[string]any
				return hx_zero_236
			}
			return hx_field_235.(map[string]any)
		}(a2))
	}
	hx_arr_240 := func(hx_obj_241 map[string]any) *hxrt.Array {
		hx_field_242 := hx_obj_241["platforms"]
		if hx_field_242 == nil {
			var hx_zero_243 *hxrt.Array
			return hx_zero_243
		}
		return hx_field_242.(*hxrt.Array)
	}(a)
	hx_arr_240.Push(self.curplatform)
	return true
}

func (self *haxe__rtti__XmlParser) merge(t *haxe__rtti__TypeTree) {
	inf := haxe__rtti__TypeApi_typeInfos(t)
	pack := self.__hx_this.splitString(func(hx_obj_244 map[string]any) *string {
		hx_field_245 := hx_obj_244["path"]
		if hx_field_245 == nil {
			var hx_zero_246 *string
			return hx_zero_246
		}
		return hx_field_245.(*string)
	}(inf), hxrt.StringFromLiteral("."))
	cur := self.root
	curpack := hxrt.NewArray()
	pack.Pop()
	_g := 0
	for _g < pack.Len() {
		p := func(hx_value_248 any) *string {
			if hx_value_248 == nil {
				var hx_zero_249 *string
				return hx_zero_249
			}
			return hx_value_248.(*string)
		}(pack.Get(_g))
		_g = int(int32((_g + 1)))
		found := false
		_g_1 := 0
		for _g_1 < cur.Len() {
			pk := func(hx_value_250 any) *haxe__rtti__TypeTree {
				if hx_value_250 == nil {
					var hx_zero_251 *haxe__rtti__TypeTree
					return hx_zero_251
				}
				return hx_value_250.(*haxe__rtti__TypeTree)
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
		ct := func(hx_value_254 any) *haxe__rtti__TypeTree {
			if hx_value_254 == nil {
				var hx_zero_255 *haxe__rtti__TypeTree
				return hx_zero_255
			}
			return hx_value_254.(*haxe__rtti__TypeTree)
		}(cur.Get(_g_3))
		_g_3 = int(int32((_g_3 + 1)))
		if func() bool {
			var hx_if_256 bool
			if ct.tag == 0 {
				_g_4 := ct.params[0].(*string)
				_ = _g_4
				_g_5 := ct.params[1].(*string)
				_ = _g_5
				_g_6 := ct.params[2].(*hxrt.Array)
				_ = _g_6
				hx_if_256 = true
			} else {
				hx_if_256 = false
			}
			return hx_if_256
		}() {
			continue
		}
		tinf := haxe__rtti__TypeApi_typeInfos(ct)
		if hxrt.StringEqualStringPtr(func(hx_obj_327 map[string]any) *string {
			hx_field_328 := hx_obj_327["path"]
			if hx_field_328 == nil {
				var hx_zero_329 *string
				return hx_zero_329
			}
			return hx_field_328.(*string)
		}(tinf), func(hx_obj_330 map[string]any) *string {
			hx_field_331 := hx_obj_330["path"]
			if hx_field_331 == nil {
				var hx_zero_332 *string
				return hx_zero_332
			}
			return hx_field_331.(*string)
		}(inf)) {
			sameType := true
			if hxrt.StringEqualStringPtr(func(hx_obj_266 map[string]any) *string {
				hx_field_267 := hx_obj_266["doc"]
				if hx_field_267 == nil {
					var hx_zero_268 *string
					return hx_zero_268
				}
				return hx_field_267.(*string)
			}(tinf), nil) != hxrt.StringEqualStringPtr(func(hx_obj_269 map[string]any) *string {
				hx_field_270 := hx_obj_269["doc"]
				if hx_field_270 == nil {
					var hx_zero_271 *string
					return hx_zero_271
				}
				return hx_field_270.(*string)
			}(inf), nil) {
				if hxrt.StringEqualStringPtr(func(hx_obj_263 map[string]any) *string {
					hx_field_264 := hx_obj_263["doc"]
					if hx_field_264 == nil {
						var hx_zero_265 *string
						return hx_zero_265
					}
					return hx_field_264.(*string)
				}(inf), nil) {
					inf["doc"] = func(hx_obj_257 map[string]any) *string {
						hx_field_258 := hx_obj_257["doc"]
						if hx_field_258 == nil {
							var hx_zero_259 *string
							return hx_zero_259
						}
						return hx_field_258.(*string)
					}(tinf)
				} else {
					tinf["doc"] = func(hx_obj_260 map[string]any) *string {
						hx_field_261 := hx_obj_260["doc"]
						if hx_field_261 == nil {
							var hx_zero_262 *string
							return hx_zero_262
						}
						return hx_field_261.(*string)
					}(inf)
				}
			}
			if hxrt.StringEqualStringPtr(func(hx_obj_272 map[string]any) *string {
				hx_field_273 := hx_obj_272["path"]
				if hx_field_273 == nil {
					var hx_zero_274 *string
					return hx_zero_274
				}
				return hx_field_273.(*string)
			}(tinf), hxrt.StringFromLiteral("haxe._Int64.NativeInt64")) {
				continue
			}
			if (hxrt.StringEqualStringPtr(func(hx_obj_275 map[string]any) *string {
				hx_field_276 := hx_obj_275["module"]
				if hx_field_276 == nil {
					var hx_zero_277 *string
					return hx_zero_277
				}
				return hx_field_276.(*string)
			}(tinf), func(hx_obj_278 map[string]any) *string {
				hx_field_279 := hx_obj_278["module"]
				if hx_field_279 == nil {
					var hx_zero_280 *string
					return hx_zero_280
				}
				return hx_field_279.(*string)
			}(inf)) && hxrt.StringEqualStringPtr(func(hx_obj_281 map[string]any) *string {
				hx_field_282 := hx_obj_281["doc"]
				if hx_field_282 == nil {
					var hx_zero_283 *string
					return hx_zero_283
				}
				return hx_field_282.(*string)
			}(tinf), func(hx_obj_284 map[string]any) *string {
				hx_field_285 := hx_obj_284["doc"]
				if hx_field_285 == nil {
					var hx_zero_286 *string
					return hx_zero_286
				}
				return hx_field_285.(*string)
			}(inf))) && (func(hx_obj_287 map[string]any) bool {
				hx_field_288 := hx_obj_287["isPrivate"]
				if hx_field_288 == nil {
					var hx_zero_289 bool
					return hx_zero_289
				}
				return hx_field_288.(bool)
			}(tinf) == func(hx_obj_290 map[string]any) bool {
				hx_field_291 := hx_obj_290["isPrivate"]
				if hx_field_291 == nil {
					var hx_zero_292 bool
					return hx_zero_292
				}
				return hx_field_291.(bool)
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
			var hx_if_320 *string
			if !hxrt.StringEqualStringPtr(func(hx_obj_293 map[string]any) *string {
				hx_field_294 := hx_obj_293["module"]
				if hx_field_294 == nil {
					var hx_zero_295 *string
					return hx_zero_295
				}
				return hx_field_294.(*string)
			}(tinf), func(hx_obj_296 map[string]any) *string {
				hx_field_297 := hx_obj_296["module"]
				if hx_field_297 == nil {
					var hx_zero_298 *string
					return hx_zero_298
				}
				return hx_field_297.(*string)
			}(inf)) {
				hx_if_320 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("module "), func(hx_obj_299 map[string]any) *string {
					hx_field_300 := hx_obj_299["module"]
					if hx_field_300 == nil {
						var hx_zero_301 *string
						return hx_zero_301
					}
					return hx_field_300.(*string)
				}(inf)), hxrt.StringFromLiteral(" should be ")), func(hx_obj_302 map[string]any) *string {
					hx_field_303 := hx_obj_302["module"]
					if hx_field_303 == nil {
						var hx_zero_304 *string
						return hx_zero_304
					}
					return hx_field_303.(*string)
				}(tinf))
			} else {
				var hx_if_319 *string
				if !hxrt.StringEqualStringPtr(func(hx_obj_305 map[string]any) *string {
					hx_field_306 := hx_obj_305["doc"]
					if hx_field_306 == nil {
						var hx_zero_307 *string
						return hx_zero_307
					}
					return hx_field_306.(*string)
				}(tinf), func(hx_obj_308 map[string]any) *string {
					hx_field_309 := hx_obj_308["doc"]
					if hx_field_309 == nil {
						var hx_zero_310 *string
						return hx_zero_310
					}
					return hx_field_309.(*string)
				}(inf)) {
					hx_if_319 = hxrt.StringFromLiteral("documentation is different")
				} else {
					var hx_if_318 *string
					if func(hx_obj_311 map[string]any) bool {
						hx_field_312 := hx_obj_311["isPrivate"]
						if hx_field_312 == nil {
							var hx_zero_313 bool
							return hx_zero_313
						}
						return hx_field_312.(bool)
					}(tinf) != func(hx_obj_314 map[string]any) bool {
						hx_field_315 := hx_obj_314["isPrivate"]
						if hx_field_315 == nil {
							var hx_zero_316 bool
							return hx_zero_316
						}
						return hx_field_315.(bool)
					}(inf) {
						hx_if_318 = hxrt.StringFromLiteral("private flag is different")
					} else {
						var hx_if_317 *string
						if !sameType {
							hx_if_317 = hxrt.StringFromLiteral("type kind is different")
						} else {
							hx_if_317 = hxrt.StringFromLiteral("could not merge definition")
						}
						hx_if_318 = hx_if_317
					}
					hx_if_319 = hx_if_318
				}
				hx_if_320 = hx_if_319
			}
			msg := hx_if_320
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Incompatibilities between "), func(hx_obj_321 map[string]any) *string {
				hx_field_322 := hx_obj_321["path"]
				if hx_field_322 == nil {
					var hx_zero_323 *string
					return hx_zero_323
				}
				return hx_field_322.(*string)
			}(tinf)), hxrt.StringFromLiteral(" in ")), self.__hx_this.joinStringArray(func(hx_obj_324 map[string]any) *hxrt.Array {
				hx_field_325 := hx_obj_324["platforms"]
				if hx_field_325 == nil {
					var hx_zero_326 *hxrt.Array
					return hx_zero_326
				}
				return hx_field_325.(*hxrt.Array)
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
	for func(hx_obj_334 map[string]any) func() bool {
		hx_field_335 := hx_obj_334["hasNext"]
		if hx_field_335 == nil {
			var hx_zero_336 func() bool
			return hx_zero_336
		}
		return hx_field_335.(func() bool)
	}(c)() {
		c_1 := func(hx_obj_337 map[string]any) func() *Xml {
			hx_field_338 := hx_obj_337["next"]
			if hx_field_338 == nil {
				var hx_zero_339 func() *Xml
				return hx_zero_339
			}
			return hx_field_338.(func() *Xml)
		}(c)()
		self.__hx_this.merge(self.__hx_this.processElement(c_1))
	}
}

func (self *haxe__rtti__XmlParser) processElement(x *Xml) *haxe__rtti__TypeTree {
	var hx_if_340 *string
	if hxrt.HaxeEqual(x.nodeType, Xml_Document) {
		hx_if_340 = hxrt.StringFromLiteral("Document")
	} else {
		if !hxrt.HaxeEqual(x.nodeType, Xml_Element) {
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(x.nodeType))))
		}
		hx_if_340 = x.nodeName
	}
	nodeName := hx_if_340
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
	var hx_throw_zero_341 *haxe__rtti__TypeTree
	return hx_throw_zero_341
}

func (self *haxe__rtti__XmlParser) xmeta(x *Xml) *hxrt.Array {
	ml := hxrt.NewArray()
	m := x.__hx_this.elementsNamed(hxrt.StringFromLiteral("m"))
	for func(hx_obj_342 map[string]any) func() bool {
		hx_field_343 := hx_obj_342["hasNext"]
		if hx_field_343 == nil {
			var hx_zero_344 func() bool
			return hx_zero_344
		}
		return hx_field_343.(func() bool)
	}(m)() {
		m_1 := func(hx_obj_345 map[string]any) func() *Xml {
			hx_field_346 := hx_obj_345["next"]
			if hx_field_346 == nil {
				var hx_zero_347 func() *Xml
				return hx_zero_347
			}
			return hx_field_346.(func() *Xml)
		}(m)()
		pl := hxrt.NewArray()
		p := m_1.__hx_this.elementsNamed(hxrt.StringFromLiteral("e"))
		for func(hx_obj_348 map[string]any) func() bool {
			hx_field_349 := hx_obj_348["hasNext"]
			if hx_field_349 == nil {
				var hx_zero_350 func() bool
				return hx_zero_350
			}
			return hx_field_349.(func() bool)
		}(p)() {
			p_1 := func(hx_obj_351 map[string]any) func() *Xml {
				hx_field_352 := hx_obj_351["next"]
				if hx_field_352 == nil {
					var hx_zero_353 func() *Xml
					return hx_zero_353
				}
				return hx_field_352.(func() *Xml)
			}(p)()
			pl.Push(self.__hx_this.innerHTML(p_1))
		}
		hx_obj_356 := map[string]any{}
		hx_obj_356["name"] = self.__hx_this.requireAttr(m_1, hxrt.StringFromLiteral("n"))
		hx_obj_356["params"] = pl
		ml.Push(hx_obj_356)
	}
	return ml
}

func (self *haxe__rtti__XmlParser) xoverloads(x *Xml) *hxrt.Array {
	l := hxrt.NewArray()
	m := x.__hx_this.elements()
	for func(hx_obj_357 map[string]any) func() bool {
		hx_field_358 := hx_obj_357["hasNext"]
		if hx_field_358 == nil {
			var hx_zero_359 func() bool
			return hx_zero_359
		}
		return hx_field_358.(func() bool)
	}(m)() {
		m_1 := func(hx_obj_360 map[string]any) func() *Xml {
			hx_field_361 := hx_obj_360["next"]
			if hx_field_361 == nil {
				var hx_zero_362 func() *Xml
				return hx_zero_362
			}
			return hx_field_361.(func() *Xml)
		}(m)()
		l.Push(self.__hx_this.xclassfield(m_1, false))
	}
	return l
}

func (self *haxe__rtti__XmlParser) xpath(x *Xml) map[string]any {
	path := self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("path")))
	params := hxrt.NewArray()
	c := x.__hx_this.elements()
	for func(hx_obj_364 map[string]any) func() bool {
		hx_field_365 := hx_obj_364["hasNext"]
		if hx_field_365 == nil {
			var hx_zero_366 func() bool
			return hx_zero_366
		}
		return hx_field_365.(func() bool)
	}(c)() {
		c_1 := func(hx_obj_367 map[string]any) func() *Xml {
			hx_field_368 := hx_obj_367["next"]
			if hx_field_368 == nil {
				var hx_zero_369 func() *Xml
				return hx_zero_369
			}
			return hx_field_368.(func() *Xml)
		}(c)()
		params.Push(self.__hx_this.xtype(c_1))
	}
	hx_obj_371 := map[string]any{}
	hx_obj_371["path"] = path
	hx_obj_371["params"] = params
	return hx_obj_371
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
	for func(hx_obj_372 map[string]any) func() bool {
		hx_field_373 := hx_obj_372["hasNext"]
		if hx_field_373 == nil {
			var hx_zero_374 func() bool
			return hx_zero_374
		}
		return hx_field_373.(func() bool)
	}(c)() {
		c_1 := func(hx_obj_375 map[string]any) func() *Xml {
			hx_field_376 := hx_obj_375["next"]
			if hx_field_376 == nil {
				var hx_zero_377 func() *Xml
				return hx_zero_377
			}
			return hx_field_376.(func() *Xml)
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
	hx_obj_382 := map[string]any{}
	hx_obj_382["file"] = x.__hx_this.get(hxrt.StringFromLiteral("file"))
	hx_obj_382["path"] = self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("path")))
	var hx_if_383 *string
	if x.__hx_this.exists(hxrt.StringFromLiteral("module")) {
		hx_if_383 = self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("module")))
	} else {
		hx_if_383 = nil
	}
	hx_obj_382["module"] = hx_if_383
	hx_obj_382["doc"] = doc
	hx_obj_382["isPrivate"] = x.__hx_this.exists(hxrt.StringFromLiteral("private"))
	hx_obj_382["isExtern"] = x.__hx_this.exists(hxrt.StringFromLiteral("extern"))
	hx_obj_382["isFinal"] = x.__hx_this.exists(hxrt.StringFromLiteral("final"))
	hx_obj_382["isInterface"] = isInterface
	hx_obj_382["params"] = self.__hx_this.mkTypeParams(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("params")))
	hx_obj_382["superClass"] = csuper
	hx_obj_382["interfaces"] = interfaces
	hx_obj_382["fields"] = fields
	hx_obj_382["statics"] = statics
	hx_obj_382["tdynamic"] = tdynamic
	hx_obj_382["platforms"] = self.__hx_this.defplat()
	hx_obj_382["meta"] = meta
	return hx_obj_382
}

func (self *haxe__rtti__XmlParser) xclassfield(x *Xml, defPublic bool) map[string]any {
	e := x.__hx_this.elements()
	t := self.__hx_this.xtype(func(hx_obj_384 map[string]any) func() *Xml {
		hx_field_385 := hx_obj_384["next"]
		if hx_field_385 == nil {
			var hx_zero_386 func() *Xml
			return hx_zero_386
		}
		return hx_field_385.(func() *Xml)
	}(e)())
	var doc *string = nil
	meta := hxrt.NewArray()
	var overloads *hxrt.Array = nil
	var line any = nil
	if x.__hx_this.exists(hxrt.StringFromLiteral("line")) {
		line = self.__hx_this.parseIntString(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("line")))
	}
	c := e
	for func(hx_obj_387 map[string]any) func() bool {
		hx_field_388 := hx_obj_387["hasNext"]
		if hx_field_388 == nil {
			var hx_zero_389 func() bool
			return hx_zero_389
		}
		return hx_field_388.(func() bool)
	}(c)() {
		c_1 := func(hx_obj_390 map[string]any) func() *Xml {
			hx_field_391 := hx_obj_390["next"]
			if hx_field_391 == nil {
				var hx_zero_392 func() *Xml
				return hx_zero_392
			}
			return hx_field_391.(func() *Xml)
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
	hx_obj_393 := map[string]any{}
	hx_obj_393["name"] = self.__hx_this.elementName(x)
	hx_obj_393["type"] = t
	hx_obj_393["isPublic"] = (x.__hx_this.exists(hxrt.StringFromLiteral("public")) || func(hx_value_394 any) bool {
		if hx_value_394 == nil {
			var hx_zero_395 bool
			return hx_zero_395
		}
		return hx_value_394.(bool)
	}(defPublic))
	hx_obj_393["isFinal"] = x.__hx_this.exists(hxrt.StringFromLiteral("final"))
	hx_obj_393["isOverride"] = x.__hx_this.exists(hxrt.StringFromLiteral("override"))
	hx_obj_393["line"] = func(hx_value_396 any) any {
		if hx_value_396 == nil {
			return nil
		}
		return hx_value_396.(int)
	}(line)
	hx_obj_393["doc"] = doc
	var hx_if_397 *haxe__rtti__Rights
	if x.__hx_this.exists(hxrt.StringFromLiteral("get")) {
		hx_if_397 = self.__hx_this.mkRights(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("get")))
	} else {
		hx_if_397 = haxe__rtti__Rights_RNormal
	}
	hx_obj_393["get"] = hx_if_397
	var hx_if_398 *haxe__rtti__Rights
	if x.__hx_this.exists(hxrt.StringFromLiteral("set")) {
		hx_if_398 = self.__hx_this.mkRights(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("set")))
	} else {
		hx_if_398 = haxe__rtti__Rights_RNormal
	}
	hx_obj_393["set"] = hx_if_398
	var hx_if_399 *hxrt.Array
	if x.__hx_this.exists(hxrt.StringFromLiteral("params")) {
		hx_if_399 = self.__hx_this.mkTypeParams(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("params")))
	} else {
		hx_if_399 = hxrt.NewArray()
	}
	hx_obj_393["params"] = hx_if_399
	hx_obj_393["platforms"] = self.__hx_this.defplat()
	hx_obj_393["meta"] = meta
	hx_obj_393["overloads"] = overloads
	var hx_if_400 *string
	if x.__hx_this.exists(hxrt.StringFromLiteral("expr")) {
		hx_if_400 = self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("expr"))
	} else {
		hx_if_400 = nil
	}
	hx_obj_393["expr"] = hx_if_400
	return hx_obj_393
}

func (self *haxe__rtti__XmlParser) xenum(x *Xml) map[string]any {
	cl := hxrt.NewArray()
	var doc *string = nil
	meta := hxrt.NewArray()
	c := x.__hx_this.elements()
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
	hx_obj_408 := map[string]any{}
	hx_obj_408["file"] = x.__hx_this.get(hxrt.StringFromLiteral("file"))
	hx_obj_408["path"] = self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("path")))
	var hx_if_409 *string
	if x.__hx_this.exists(hxrt.StringFromLiteral("module")) {
		hx_if_409 = self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("module")))
	} else {
		hx_if_409 = nil
	}
	hx_obj_408["module"] = hx_if_409
	hx_obj_408["doc"] = doc
	hx_obj_408["isPrivate"] = x.__hx_this.exists(hxrt.StringFromLiteral("private"))
	hx_obj_408["isExtern"] = x.__hx_this.exists(hxrt.StringFromLiteral("extern"))
	hx_obj_408["params"] = self.__hx_this.mkTypeParams(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("params")))
	hx_obj_408["constructors"] = cl
	hx_obj_408["platforms"] = self.__hx_this.defplat()
	hx_obj_408["meta"] = meta
	return hx_obj_408
}

func (self *haxe__rtti__XmlParser) xenumfield(x *Xml) map[string]any {
	var args *hxrt.Array = nil
	docElements := x.__hx_this.elementsNamed(hxrt.StringFromLiteral("haxe_doc"))
	var hx_if_416 *Xml
	if func(hx_obj_410 map[string]any) func() bool {
		hx_field_411 := hx_obj_410["hasNext"]
		if hx_field_411 == nil {
			var hx_zero_412 func() bool
			return hx_zero_412
		}
		return hx_field_411.(func() bool)
	}(docElements)() {
		hx_if_416 = func(hx_obj_413 map[string]any) func() *Xml {
			hx_field_414 := hx_obj_413["next"]
			if hx_field_414 == nil {
				var hx_zero_415 func() *Xml
				return hx_zero_415
			}
			return hx_field_414.(func() *Xml)
		}(docElements)()
	} else {
		hx_if_416 = nil
	}
	xdoc := hx_if_416
	var hx_if_417 *hxrt.Array
	if self.__hx_this.hasNamedElement(x, hxrt.StringFromLiteral("meta")) {
		hx_if_417 = self.__hx_this.xmeta(self.__hx_this.requireNamedElement(x, hxrt.StringFromLiteral("meta")))
	} else {
		hx_if_417 = hxrt.NewArray()
	}
	meta := hx_if_417
	if x.__hx_this.exists(hxrt.StringFromLiteral("a")) {
		names := self.__hx_this.splitString(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("a")), hxrt.StringFromLiteral(":"))
		elts := x.__hx_this.elements()
		args = hxrt.NewArray()
		_g := 0
		for _g < names.Len() {
			c := func(hx_value_418 any) *string {
				if hx_value_418 == nil {
					var hx_zero_419 *string
					return hx_zero_419
				}
				return hx_value_418.(*string)
			}(names.Get(_g))
			_g = int(int32((_g + 1)))
			opt := false
			if hxrt.StringEqualStringPtr(hxrt.StringCharAtStringPtr(c, 0), hxrt.StringFromLiteral("?")) {
				opt = true
				c = hxrt.StringSubstrStringPtr(c, 1, 0, false)
			}
			hx_obj_421 := map[string]any{}
			hx_obj_421["name"] = c
			hx_obj_421["opt"] = opt
			hx_obj_421["t"] = self.__hx_this.xtype(func(hx_obj_422 map[string]any) func() *Xml {
				hx_field_423 := hx_obj_422["next"]
				if hx_field_423 == nil {
					var hx_zero_424 func() *Xml
					return hx_zero_424
				}
				return hx_field_423.(func() *Xml)
			}(elts)())
			args.Push(hx_obj_421)
		}
	}
	hx_obj_425 := map[string]any{}
	hx_obj_425["name"] = self.__hx_this.elementName(x)
	hx_obj_425["args"] = args
	var hx_if_426 *string
	if xdoc == nil {
		hx_if_426 = nil
	} else {
		hx_if_426 = self.__hx_this.innerData(xdoc)
	}
	hx_obj_425["doc"] = hx_if_426
	hx_obj_425["meta"] = meta
	hx_obj_425["platforms"] = self.__hx_this.defplat()
	return hx_obj_425
}

func (self *haxe__rtti__XmlParser) xabstract(x *Xml) map[string]any {
	var doc *string = nil
	var impl map[string]any = nil
	var athis *haxe__rtti__CType = nil
	meta := hxrt.NewArray()
	to := hxrt.NewArray()
	from := hxrt.NewArray()
	c := x.__hx_this.elements()
	for func(hx_obj_427 map[string]any) func() bool {
		hx_field_428 := hx_obj_427["hasNext"]
		if hx_field_428 == nil {
			var hx_zero_429 func() bool
			return hx_zero_429
		}
		return hx_field_428.(func() bool)
	}(c)() {
		c_1 := func(hx_obj_430 map[string]any) func() *Xml {
			hx_field_431 := hx_obj_430["next"]
			if hx_field_431 == nil {
				var hx_zero_432 func() *Xml
				return hx_zero_432
			}
			return hx_field_431.(func() *Xml)
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
					for func(hx_obj_433 map[string]any) func() bool {
						hx_field_434 := hx_obj_433["hasNext"]
						if hx_field_434 == nil {
							var hx_zero_435 func() bool
							return hx_zero_435
						}
						return hx_field_434.(func() bool)
					}(t)() {
						t_1 := func(hx_obj_436 map[string]any) func() *Xml {
							hx_field_437 := hx_obj_436["next"]
							if hx_field_437 == nil {
								var hx_zero_438 func() *Xml
								return hx_zero_438
							}
							return hx_field_437.(func() *Xml)
						}(t)()
						hx_obj_440 := map[string]any{}
						hx_obj_440["t"] = self.__hx_this.xtype(self.__hx_this.requireFirstElement(t_1))
						hx_obj_440["field"] = t_1.__hx_this.get(hxrt.StringFromLiteral("field"))
						to.Push(hx_obj_440)
					}
				} else {
					if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("from")) {
						t_2 := c_1.__hx_this.elements()
						for func(hx_obj_441 map[string]any) func() bool {
							hx_field_442 := hx_obj_441["hasNext"]
							if hx_field_442 == nil {
								var hx_zero_443 func() bool
								return hx_zero_443
							}
							return hx_field_442.(func() bool)
						}(t_2)() {
							t_3 := func(hx_obj_444 map[string]any) func() *Xml {
								hx_field_445 := hx_obj_444["next"]
								if hx_field_445 == nil {
									var hx_zero_446 func() *Xml
									return hx_zero_446
								}
								return hx_field_445.(func() *Xml)
							}(t_2)()
							hx_obj_448 := map[string]any{}
							hx_obj_448["t"] = self.__hx_this.xtype(self.__hx_this.requireFirstElement(t_3))
							hx_obj_448["field"] = t_3.__hx_this.get(hxrt.StringFromLiteral("field"))
							from.Push(hx_obj_448)
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
	hx_obj_449 := map[string]any{}
	hx_obj_449["file"] = x.__hx_this.get(hxrt.StringFromLiteral("file"))
	hx_obj_449["path"] = self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("path")))
	var hx_if_450 *string
	if x.__hx_this.exists(hxrt.StringFromLiteral("module")) {
		hx_if_450 = self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("module")))
	} else {
		hx_if_450 = nil
	}
	hx_obj_449["module"] = hx_if_450
	hx_obj_449["doc"] = doc
	hx_obj_449["isPrivate"] = x.__hx_this.exists(hxrt.StringFromLiteral("private"))
	hx_obj_449["params"] = self.__hx_this.mkTypeParams(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("params")))
	hx_obj_449["platforms"] = self.__hx_this.defplat()
	hx_obj_449["meta"] = meta
	hx_obj_449["athis"] = athis
	hx_obj_449["to"] = to
	hx_obj_449["from"] = from
	hx_obj_449["impl"] = impl
	return hx_obj_449
}

func (self *haxe__rtti__XmlParser) xtypedef(x *Xml) map[string]any {
	var doc *string = nil
	var t *haxe__rtti__CType = nil
	meta := hxrt.NewArray()
	c := x.__hx_this.elements()
	for func(hx_obj_451 map[string]any) func() bool {
		hx_field_452 := hx_obj_451["hasNext"]
		if hx_field_452 == nil {
			var hx_zero_453 func() bool
			return hx_zero_453
		}
		return hx_field_452.(func() bool)
	}(c)() {
		c_1 := func(hx_obj_454 map[string]any) func() *Xml {
			hx_field_455 := hx_obj_454["next"]
			if hx_field_455 == nil {
				var hx_zero_456 func() *Xml
				return hx_zero_456
			}
			return hx_field_455.(func() *Xml)
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
	hx_obj_457 := map[string]any{}
	hx_obj_457["file"] = x.__hx_this.get(hxrt.StringFromLiteral("file"))
	hx_obj_457["path"] = self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("path")))
	var hx_if_458 *string
	if x.__hx_this.exists(hxrt.StringFromLiteral("module")) {
		hx_if_458 = self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("module")))
	} else {
		hx_if_458 = nil
	}
	hx_obj_457["module"] = hx_if_458
	hx_obj_457["doc"] = doc
	hx_obj_457["isPrivate"] = x.__hx_this.exists(hxrt.StringFromLiteral("private"))
	hx_obj_457["params"] = self.__hx_this.mkTypeParams(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("params")))
	hx_obj_457["type"] = t
	hx_obj_457["types"] = types
	hx_obj_457["platforms"] = self.__hx_this.defplat()
	hx_obj_457["meta"] = meta
	return hx_obj_457
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
		var hx_if_459 *hxrt.Array
		if x.__hx_this.exists(hxrt.StringFromLiteral("v")) {
			hx_if_459 = self.__hx_this.splitString(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("v")), hxrt.StringFromLiteral(":"))
		} else {
			hx_if_459 = nil
		}
		evalues := hx_if_459
		valueIndex := 0
		e := x.__hx_this.elements()
		for func(hx_obj_460 map[string]any) func() bool {
			hx_field_461 := hx_obj_460["hasNext"]
			if hx_field_461 == nil {
				var hx_zero_462 func() bool
				return hx_zero_462
			}
			return hx_field_461.(func() bool)
		}(e)() {
			e_1 := func(hx_obj_463 map[string]any) func() *Xml {
				hx_field_464 := hx_obj_463["next"]
				if hx_field_464 == nil {
					var hx_zero_465 func() *Xml
					return hx_zero_465
				}
				return hx_field_464.(func() *Xml)
			}(e)()
			opt := false
			var hx_if_468 *string
			if argIndex < aname.Len() {
				hx_if_468 = hxrt.StdString(func(hx_value_466 any) *string {
					if hx_value_466 == nil {
						var hx_zero_467 *string
						return hx_zero_467
					}
					return hx_value_466.(*string)
				}(aname.Get(argIndex)))
			} else {
				hx_if_468 = nil
			}
			a := hx_if_468
			argIndex = int(int32((argIndex + 1)))
			if hxrt.StringEqualStringPtr(a, nil) {
				a = hxrt.StringFromLiteral("")
			}
			if hxrt.StringEqualStringPtr(hxrt.StringCharAtStringPtr(a, 0), hxrt.StringFromLiteral("?")) {
				opt = true
				a = hxrt.StringSubstrStringPtr(a, 1, 0, false)
			}
			var hx_if_472 *string
			if (evalues == nil) || (valueIndex >= evalues.Len()) {
				hx_if_472 = nil
			} else {
				hx_post_469 := valueIndex
				valueIndex = int(int32((valueIndex + 1)))
				hx_if_472 = hxrt.StdString(func(hx_value_470 any) *string {
					if hx_value_470 == nil {
						var hx_zero_471 *string
						return hx_zero_471
					}
					return hx_value_470.(*string)
				}(evalues.Get(hx_post_469)))
			}
			v := hx_if_472
			hx_obj_474 := map[string]any{}
			hx_obj_474["name"] = a
			hx_obj_474["opt"] = opt
			hx_obj_474["t"] = self.__hx_this.xtype(e_1)
			var hx_if_475 *string
			if hxrt.StringEqualStringPtr(v, hxrt.StringFromLiteral("")) {
				hx_if_475 = nil
			} else {
				hx_if_475 = v
			}
			hx_obj_474["value"] = hx_if_475
			args.Push(hx_obj_474)
		}
		ret := func(hx_value_476 any) map[string]any {
			if hx_value_476 == nil {
				var hx_zero_477 map[string]any
				return hx_zero_477
			}
			return hx_value_476.(map[string]any)
		}(args.Get(int((hxrt.Int32Wrap(args.Len()) - hxrt.Int32Wrap(1)))))
		callArgs := hxrt.NewArray()
		_g := 0
		_g1 := int((hxrt.Int32Wrap(args.Len()) - hxrt.Int32Wrap(1)))
		for _g < _g1 {
			hx_post_478 := _g
			_g = int(int32((_g + 1)))
			i := hx_post_478
			callArgs.Push(args.Get(i))
		}
		return haxe__rtti__CType_CFunction(callArgs, func(hx_obj_480 map[string]any) *haxe__rtti__CType {
			hx_field_481 := hx_obj_480["t"]
			if hx_field_481 == nil {
				var hx_zero_482 *haxe__rtti__CType
				return hx_zero_482
			}
			return hx_field_481.(*haxe__rtti__CType)
		}(ret))
	}
	if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("a")) {
		fields := hxrt.NewArray()
		f := x.__hx_this.elements()
		for func(hx_obj_483 map[string]any) func() bool {
			hx_field_484 := hx_obj_483["hasNext"]
			if hx_field_484 == nil {
				var hx_zero_485 func() bool
				return hx_zero_485
			}
			return hx_field_484.(func() bool)
		}(f)() {
			f_1 := func(hx_obj_486 map[string]any) func() *Xml {
				hx_field_487 := hx_obj_486["next"]
				if hx_field_487 == nil {
					var hx_zero_488 func() *Xml
					return hx_zero_488
				}
				return hx_field_487.(func() *Xml)
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
	var hx_throw_zero_490 *haxe__rtti__CType
	return hx_throw_zero_490
}

func (self *haxe__rtti__XmlParser) xtypeparams(x *Xml) *hxrt.Array {
	p := hxrt.NewArray()
	c := x.__hx_this.elements()
	for func(hx_obj_491 map[string]any) func() bool {
		hx_field_492 := hx_obj_491["hasNext"]
		if hx_field_492 == nil {
			var hx_zero_493 func() bool
			return hx_zero_493
		}
		return hx_field_492.(func() bool)
	}(c)() {
		c_1 := func(hx_obj_494 map[string]any) func() *Xml {
			hx_field_495 := hx_obj_494["next"]
			if hx_field_495 == nil {
				var hx_zero_496 func() *Xml
				return hx_zero_496
			}
			return hx_field_495.(func() *Xml)
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
		hx_post_499 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_499
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
	var hx_if_502 *string
	if hxrt.StringEqualStringPtr(value, nil) {
		hx_if_502 = hxrt.StringFromLiteral("")
	} else {
		hx_if_502 = value
	}
	return hx_if_502
}

func (self *haxe__rtti__XmlParser) hasNamedElement(x *Xml, name *string) bool {
	return func(hx_obj_503 map[string]any) func() bool {
		hx_field_504 := hx_obj_503["hasNext"]
		if hx_field_504 == nil {
			var hx_zero_505 func() bool
			return hx_zero_505
		}
		return hx_field_504.(func() bool)
	}(x.__hx_this.elementsNamed(name))()
}

func (self *haxe__rtti__XmlParser) requireNamedElement(x *Xml, name *string) *Xml {
	elements := x.__hx_this.elementsNamed(name)
	if !func(hx_obj_506 map[string]any) func() bool {
		hx_field_507 := hx_obj_506["hasNext"]
		if hx_field_507 == nil {
			var hx_zero_508 func() bool
			return hx_zero_508
		}
		return hx_field_507.(func() bool)
	}(elements)() {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(self.__hx_this.nodeDisplayName(x), hxrt.StringFromLiteral(" is missing element ")), name))
	}
	return func(hx_obj_509 map[string]any) func() *Xml {
		hx_field_510 := hx_obj_509["next"]
		if hx_field_510 == nil {
			var hx_zero_511 func() *Xml
			return hx_zero_511
		}
		return hx_field_510.(func() *Xml)
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
	var hx_if_512 *string
	if hxrt.HaxeEqual(x.nodeType, Xml_Document) {
		hx_if_512 = hxrt.StringFromLiteral("Document")
	} else {
		hx_if_512 = self.__hx_this.elementName(x)
	}
	return hx_if_512
}

func (self *haxe__rtti__XmlParser) elementName(x *Xml) *string {
	if !hxrt.HaxeEqual(x.nodeType, Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(x.nodeType))))
	}
	name := x.nodeName
	var hx_if_513 *string
	if hxrt.StringEqualStringPtr(name, nil) {
		hx_if_513 = hxrt.StringFromLiteral("")
	} else {
		hx_if_513 = name
	}
	return hx_if_513
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
	value := func(hx_value_515 any) *Xml {
		if hx_value_515 == nil {
			var hx_zero_516 *Xml
			return hx_zero_516
		}
		return hx_value_515.(*Xml)
	}(it_array.Get(func() int {
		hx_post_514 := it_current
		it_current = int(int32((it_current + 1)))
		return hx_post_514
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
		hx_post_517 := _g_current
		_g_current = int(int32((_g_current + 1)))
		child := func(hx_value_518 any) *Xml {
			if hx_value_518 == nil {
				var hx_zero_519 *Xml
				return hx_zero_519
			}
			return hx_value_518.(*Xml)
		}(_g_array.Get(hx_post_517))
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
	var hx_if_520 int
	if negative {
		hx_if_520 = int(-int32(result))
	} else {
		hx_if_520 = result
	}
	return hx_if_520
}
