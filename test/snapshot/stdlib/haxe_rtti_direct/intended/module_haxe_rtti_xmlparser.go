package main

import "snapshot/hxrt"

type I_haxe__rtti__XmlParser interface {
	sort(l *hxrt.Array)
	sortFields(a *hxrt.Array)
	process(x *Xml, platform *string)
	mergeRights(f1 hxrt.Anonymous, f2 hxrt.Anonymous) bool
	mergeDoc(f1 hxrt.Anonymous, f2 hxrt.Anonymous) bool
	mergeFields(f hxrt.Anonymous, f2 hxrt.Anonymous) bool
	mergeClasses(c hxrt.Anonymous, c2 hxrt.Anonymous) bool
	mergeEnums(e hxrt.Anonymous, e2 hxrt.Anonymous) bool
	mergeTypedefs(t hxrt.Anonymous, t2 hxrt.Anonymous) bool
	mergeAbstracts(a hxrt.Anonymous, a2 hxrt.Anonymous) bool
	merge(t *haxe__rtti__TypeTree)
	mkPath(p *string) *string
	mkTypeParams(p *string) *hxrt.Array
	mkRights(r *string) *haxe__rtti__Rights
	xroot(x *Xml)
	processElement(x *Xml) *haxe__rtti__TypeTree
	xmeta(x *Xml) *hxrt.Array
	xoverloads(x *Xml) *hxrt.Array
	xpath(x *Xml) hxrt.Anonymous
	xclass(x *Xml) hxrt.Anonymous
	xclassfield(x *Xml, defPublic bool) hxrt.Anonymous
	xenum(x *Xml) hxrt.Anonymous
	xenumfield(x *Xml) hxrt.Anonymous
	xabstract(x *Xml) hxrt.Anonymous
	xtypedef(x *Xml) hxrt.Anonymous
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
	newField    *hxrt.Function[func(hxrt.Anonymous, hxrt.Anonymous)]
}

func (self *haxe__rtti__XmlParser) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe__rtti__XmlParser() *haxe__rtti__XmlParser {
	return hx_init_haxe__rtti__XmlParser(hx_alloc_haxe__rtti__XmlParser(nil))
}

func hx_alloc_haxe__rtti__XmlParser(owner any) *haxe__rtti__XmlParser {
	self := &haxe__rtti__XmlParser{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_haxe__rtti__XmlParser)
	}
	return self
}

func hx_init_haxe__rtti__XmlParser(self *haxe__rtti__XmlParser) *haxe__rtti__XmlParser {
	self.newField = hxrt.DefaultDynamicFunction(hxrt.NewFunction(func(c hxrt.Anonymous, f hxrt.Anonymous) {
	}), func() any {
		return self.__hx_this
	}, "newField")
	self.root = hxrt.NewArray()
	return self
}

func (self *haxe__rtti__XmlParser) sort(l *hxrt.Array) {
	if l == nil {
		l = self.root
	}
	haxe__ds__ArraySort_sort(l, func(hx_comparator_21 *hxrt.Function[func(*haxe__rtti__TypeTree, *haxe__rtti__TypeTree) int]) *hxrt.Function[func(any, any) int] {
		return hxrt.AdaptFunction(hx_comparator_21, func(hx_cmp_left_19 any, hx_cmp_right_20 any) int {
			return hxrt.FunctionCall(hx_comparator_21)(func(hx_value_22 any) *haxe__rtti__TypeTree {
				if hx_value_22 == nil {
					var hx_zero_23 *haxe__rtti__TypeTree
					return hx_zero_23
				}
				return hx_value_22.(*haxe__rtti__TypeTree)
			}(hx_cmp_left_19), func(hx_value_24 any) *haxe__rtti__TypeTree {
				if hx_value_24 == nil {
					var hx_zero_25 *haxe__rtti__TypeTree
					return hx_zero_25
				}
				return hx_value_24.(*haxe__rtti__TypeTree)
			}(hx_cmp_right_20))
		})
	}(hxrt.NewFunction(func(e1 *haxe__rtti__TypeTree, e2 *haxe__rtti__TypeTree) int {
		var hx_if_9 *string
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
			hx_if_9 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(" "), p)
		} else {
			hx_if_9 = func(hx_value_7 any) *string {
				if hx_value_7 == nil {
					var hx_zero_8 *string
					return hx_zero_8
				}
				return hx_value_7.(*string)
			}(hxrt.AnonymousGet(haxe__rtti__TypeApi_typeInfos(e1), "path"))
		}
		n1 := hx_if_9
		var hx_if_18 *string
		if e2.tag == 0 {
			_g_1 := func(hx_value_10 any) *string {
				if hx_value_10 == nil {
					var hx_zero_11 *string
					return hx_zero_11
				}
				return hx_value_10.(*string)
			}(e2.params[0])
			_g1_2 := func(hx_value_12 any) *string {
				if hx_value_12 == nil {
					var hx_zero_13 *string
					return hx_zero_13
				}
				return hx_value_12.(*string)
			}(e2.params[1])
			_ = _g1_2
			_g1_3 := func(hx_value_14 any) *hxrt.Array {
				if hx_value_14 == nil {
					var hx_zero_15 *hxrt.Array
					return hx_zero_15
				}
				return hx_value_14.(*hxrt.Array)
			}(e2.params[2])
			_ = _g1_3
			p_1 := _g_1
			hx_if_18 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(" "), p_1)
		} else {
			hx_if_18 = func(hx_value_16 any) *string {
				if hx_value_16 == nil {
					var hx_zero_17 *string
					return hx_zero_17
				}
				return hx_value_16.(*string)
			}(hxrt.AnonymousGet(haxe__rtti__TypeApi_typeInfos(e2), "path"))
		}
		n2 := hx_if_18
		return Reflect_compare(n1, n2)
	})))
	_g := 0
	for _g < l.Len() {
		x := func(hx_value_26 any) *haxe__rtti__TypeTree {
			if hx_value_26 == nil {
				var hx_zero_27 *haxe__rtti__TypeTree
				return hx_zero_27
			}
			return hx_value_26.(*haxe__rtti__TypeTree)
		}(l.Get(_g))
		_g = int(int32((_g + 1)))
		switch x.tag {
		case 0:
			_g_1 := func(hx_value_28 any) *string {
				if hx_value_28 == nil {
					var hx_zero_29 *string
					return hx_zero_29
				}
				return hx_value_28.(*string)
			}(x.params[0])
			_ = _g_1
			_g_2 := func(hx_value_30 any) *string {
				if hx_value_30 == nil {
					var hx_zero_31 *string
					return hx_zero_31
				}
				return hx_value_30.(*string)
			}(x.params[1])
			_ = _g_2
			_g_3 := func(hx_value_32 any) *hxrt.Array {
				if hx_value_32 == nil {
					var hx_zero_33 *hxrt.Array
					return hx_zero_33
				}
				return hx_value_32.(*hxrt.Array)
			}(x.params[2])
			l_1 := _g_3
			self.__hx_this.sort(l_1)
		case 1:
			_g_4 := hx_structural_recover(x.params[0])
			c := _g_4
			self.__hx_this.sortFields(func(hx_value_34 any) *hxrt.Array {
				if hx_value_34 == nil {
					var hx_zero_35 *hxrt.Array
					return hx_zero_35
				}
				return hx_value_34.(*hxrt.Array)
			}(hxrt.AnonymousGet(c, "fields")))
			self.__hx_this.sortFields(func(hx_value_36 any) *hxrt.Array {
				if hx_value_36 == nil {
					var hx_zero_37 *hxrt.Array
					return hx_zero_37
				}
				return hx_value_36.(*hxrt.Array)
			}(hxrt.AnonymousGet(c, "statics")))
		case 2:
			_g_5 := hx_structural_recover(x.params[0])
			_ = _g_5
		case 3:
			_g_6 := hx_structural_recover(x.params[0])
			_ = _g_6
		case 4:
			_g_7 := hx_structural_recover(x.params[0])
			_ = _g_7
		}
	}
}

func (self *haxe__rtti__XmlParser) sortFields(a *hxrt.Array) {
	haxe__ds__ArraySort_sort(a, func(hx_comparator_52 *hxrt.Function[func(hxrt.Anonymous, hxrt.Anonymous) int]) *hxrt.Function[func(any, any) int] {
		return hxrt.AdaptFunction(hx_comparator_52, func(hx_cmp_left_50 any, hx_cmp_right_51 any) int {
			return hxrt.FunctionCall(hx_comparator_52)(hx_structural_recover(hx_cmp_left_50), hx_structural_recover(hx_cmp_right_51))
		})
	}(hxrt.NewFunction(func(f1 hxrt.Anonymous, f2 hxrt.Anonymous) int {
		v1 := haxe__rtti__TypeApi_isVar(func(hx_value_38 any) *haxe__rtti__CType {
			if hx_value_38 == nil {
				var hx_zero_39 *haxe__rtti__CType
				return hx_zero_39
			}
			return hx_value_38.(*haxe__rtti__CType)
		}(hxrt.AnonymousGet(f1, "type")))
		v2 := haxe__rtti__TypeApi_isVar(func(hx_value_40 any) *haxe__rtti__CType {
			if hx_value_40 == nil {
				var hx_zero_41 *haxe__rtti__CType
				return hx_zero_41
			}
			return hx_value_40.(*haxe__rtti__CType)
		}(hxrt.AnonymousGet(f2, "type")))
		if v1 && !v2 {
			return -1
		}
		if v2 && !v1 {
			return 1
		}
		if hxrt.StringEqualStringPtr(func(hx_value_42 any) *string {
			if hx_value_42 == nil {
				var hx_zero_43 *string
				return hx_zero_43
			}
			return hx_value_42.(*string)
		}(hxrt.AnonymousGet(f1, "name")), hxrt.StringFromLiteral("new")) {
			return -1
		}
		if hxrt.StringEqualStringPtr(func(hx_value_44 any) *string {
			if hx_value_44 == nil {
				var hx_zero_45 *string
				return hx_zero_45
			}
			return hx_value_44.(*string)
		}(hxrt.AnonymousGet(f2, "name")), hxrt.StringFromLiteral("new")) {
			return 1
		}
		return Reflect_compare(func(hx_value_46 any) *string {
			if hx_value_46 == nil {
				var hx_zero_47 *string
				return hx_zero_47
			}
			return hx_value_46.(*string)
		}(hxrt.AnonymousGet(f1, "name")), func(hx_value_48 any) *string {
			if hx_value_48 == nil {
				var hx_zero_49 *string
				return hx_zero_49
			}
			return hx_value_48.(*string)
		}(hxrt.AnonymousGet(f2, "name")))
	})))
}

func (self *haxe__rtti__XmlParser) process(x *Xml, platform *string) {
	self.curplatform = platform
	self.__hx_this.xroot(x)
}

func (self *haxe__rtti__XmlParser) mergeRights(f1 hxrt.Anonymous, f2 hxrt.Anonymous) bool {
	if (((func(hx_value_59 any) *haxe__rtti__Rights {
		if hx_value_59 == nil {
			var hx_zero_60 *haxe__rtti__Rights
			return hx_zero_60
		}
		return hx_value_59.(*haxe__rtti__Rights)
	}(hxrt.AnonymousGet(f1, "get")) == haxe__rtti__Rights_RInline) && (func(hx_value_61 any) *haxe__rtti__Rights {
		if hx_value_61 == nil {
			var hx_zero_62 *haxe__rtti__Rights
			return hx_zero_62
		}
		return hx_value_61.(*haxe__rtti__Rights)
	}(hxrt.AnonymousGet(f1, "set")) == haxe__rtti__Rights_RNo)) && (func(hx_value_63 any) *haxe__rtti__Rights {
		if hx_value_63 == nil {
			var hx_zero_64 *haxe__rtti__Rights
			return hx_zero_64
		}
		return hx_value_63.(*haxe__rtti__Rights)
	}(hxrt.AnonymousGet(f2, "get")) == haxe__rtti__Rights_RNormal)) && (func(hx_value_65 any) *haxe__rtti__Rights {
		if hx_value_65 == nil {
			var hx_zero_66 *haxe__rtti__Rights
			return hx_zero_66
		}
		return hx_value_65.(*haxe__rtti__Rights)
	}(hxrt.AnonymousGet(f2, "set")) == haxe__rtti__Rights_RMethod) {
		func() *haxe__rtti__Rights {
			hx_structural_target_53 := f1
			hx_structural_value_55 := haxe__rtti__Rights_RNormal
			hxrt.AnonymousSet(hx_structural_target_53, "get", hx_structural_value_55)
			return hx_structural_value_55
		}()
		func() *haxe__rtti__Rights {
			hx_structural_target_56 := f1
			hx_structural_value_58 := haxe__rtti__Rights_RMethod
			hxrt.AnonymousSet(hx_structural_target_56, "set", hx_structural_value_58)
			return hx_structural_value_58
		}()
		return true
	}
	return (Type_enumEq(func(hx_value_67 any) *haxe__rtti__Rights {
		if hx_value_67 == nil {
			var hx_zero_68 *haxe__rtti__Rights
			return hx_zero_68
		}
		return hx_value_67.(*haxe__rtti__Rights)
	}(hxrt.AnonymousGet(f1, "get")), func(hx_value_69 any) *haxe__rtti__Rights {
		if hx_value_69 == nil {
			var hx_zero_70 *haxe__rtti__Rights
			return hx_zero_70
		}
		return hx_value_69.(*haxe__rtti__Rights)
	}(hxrt.AnonymousGet(f2, "get"))) && Type_enumEq(func(hx_value_71 any) *haxe__rtti__Rights {
		if hx_value_71 == nil {
			var hx_zero_72 *haxe__rtti__Rights
			return hx_zero_72
		}
		return hx_value_71.(*haxe__rtti__Rights)
	}(hxrt.AnonymousGet(f1, "set")), func(hx_value_73 any) *haxe__rtti__Rights {
		if hx_value_73 == nil {
			var hx_zero_74 *haxe__rtti__Rights
			return hx_zero_74
		}
		return hx_value_73.(*haxe__rtti__Rights)
	}(hxrt.AnonymousGet(f2, "set"))))
}

func (self *haxe__rtti__XmlParser) mergeDoc(f1 hxrt.Anonymous, f2 hxrt.Anonymous) bool {
	if hxrt.StringEqualStringPtr(func(hx_value_87 any) *string {
		if hx_value_87 == nil {
			var hx_zero_88 *string
			return hx_zero_88
		}
		return hx_value_87.(*string)
	}(hxrt.AnonymousGet(f1, "doc")), nil) {
		func() *string {
			hx_structural_target_75 := f1
			hx_structural_value_77 := func(hx_value_78 any) *string {
				if hx_value_78 == nil {
					var hx_zero_79 *string
					return hx_zero_79
				}
				return hx_value_78.(*string)
			}(hxrt.AnonymousGet(f2, "doc"))
			hxrt.AnonymousSet(hx_structural_target_75, "doc", hx_structural_value_77)
			return hx_structural_value_77
		}()
	} else {
		if hxrt.StringEqualStringPtr(func(hx_value_85 any) *string {
			if hx_value_85 == nil {
				var hx_zero_86 *string
				return hx_zero_86
			}
			return hx_value_85.(*string)
		}(hxrt.AnonymousGet(f2, "doc")), nil) {
			func() *string {
				hx_structural_target_80 := f2
				hx_structural_value_82 := func(hx_value_83 any) *string {
					if hx_value_83 == nil {
						var hx_zero_84 *string
						return hx_zero_84
					}
					return hx_value_83.(*string)
				}(hxrt.AnonymousGet(f1, "doc"))
				hxrt.AnonymousSet(hx_structural_target_80, "doc", hx_structural_value_82)
				return hx_structural_value_82
			}()
		}
	}
	return true
}

func (self *haxe__rtti__XmlParser) mergeFields(f hxrt.Anonymous, f2 hxrt.Anonymous) bool {
	return (haxe__rtti__TypeApi_fieldEq(f, f2) || (((hxrt.StringEqualStringPtr(func(hx_value_89 any) *string {
		if hx_value_89 == nil {
			var hx_zero_90 *string
			return hx_zero_90
		}
		return hx_value_89.(*string)
	}(hxrt.AnonymousGet(f, "name")), func(hx_value_91 any) *string {
		if hx_value_91 == nil {
			var hx_zero_92 *string
			return hx_zero_92
		}
		return hx_value_91.(*string)
	}(hxrt.AnonymousGet(f2, "name"))) && (self.__hx_this.mergeRights(f, f2) || self.__hx_this.mergeRights(f2, f))) && self.__hx_this.mergeDoc(f, f2)) && haxe__rtti__TypeApi_fieldEq(f, f2)))
}

func (self *haxe__rtti__XmlParser) mergeClasses(c hxrt.Anonymous, c2 hxrt.Anonymous) bool {
	if func(hx_value_93 any) bool {
		if hx_value_93 == nil {
			var hx_zero_94 bool
			return hx_zero_94
		}
		return hx_value_93.(bool)
	}(hxrt.AnonymousGet(c, "isInterface")) != func(hx_value_95 any) bool {
		if hx_value_95 == nil {
			var hx_zero_96 bool
			return hx_zero_96
		}
		return hx_value_95.(bool)
	}(hxrt.AnonymousGet(c2, "isInterface")) {
		return false
	}
	if !hxrt.StringEqualStringPtr(self.curplatform, nil) {
		hx_arr_97 := func(hx_value_98 any) *hxrt.Array {
			if hx_value_98 == nil {
				var hx_zero_99 *hxrt.Array
				return hx_zero_99
			}
			return hx_value_98.(*hxrt.Array)
		}(hxrt.AnonymousGet(c, "platforms"))
		hx_arr_97.Push(self.curplatform)
	}
	if func(hx_value_103 any) bool {
		if hx_value_103 == nil {
			var hx_zero_104 bool
			return hx_zero_104
		}
		return hx_value_103.(bool)
	}(hxrt.AnonymousGet(c, "isExtern")) != func(hx_value_105 any) bool {
		if hx_value_105 == nil {
			var hx_zero_106 bool
			return hx_zero_106
		}
		return hx_value_105.(bool)
	}(hxrt.AnonymousGet(c2, "isExtern")) {
		func() bool {
			hx_structural_target_100 := c
			hx_structural_value_102 := false
			hxrt.AnonymousSet(hx_structural_target_100, "isExtern", hx_structural_value_102)
			return hx_structural_value_102
		}()
	}
	_g := 0
	_g1 := func(hx_value_107 any) *hxrt.Array {
		if hx_value_107 == nil {
			var hx_zero_108 *hxrt.Array
			return hx_zero_108
		}
		return hx_value_107.(*hxrt.Array)
	}(hxrt.AnonymousGet(c2, "fields"))
	for _g < _g1.Len() {
		f2 := hx_structural_recover(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		var found hxrt.Anonymous = nil
		_g_1 := 0
		_g1_1 := func(hx_value_109 any) *hxrt.Array {
			if hx_value_109 == nil {
				var hx_zero_110 *hxrt.Array
				return hx_zero_110
			}
			return hx_value_109.(*hxrt.Array)
		}(hxrt.AnonymousGet(c, "fields"))
		for _g_1 < _g1_1.Len() {
			f := hx_structural_recover(_g1_1.Get(_g_1))
			_g_1 = int(int32((_g_1 + 1)))
			if self.__hx_this.mergeFields(f, f2) {
				found = f
				break
			}
		}
		if found == nil {
			func() {
				var hx_callable_111 *hxrt.Function[func(hxrt.Anonymous, hxrt.Anonymous)] = self.newField
				var hx_argument_112 hxrt.Anonymous = c
				var hx_argument_113 hxrt.Anonymous = f2
				hxrt.FunctionCall(hx_callable_111)(hx_argument_112, hx_argument_113)
			}()
			hx_arr_114 := func(hx_value_115 any) *hxrt.Array {
				if hx_value_115 == nil {
					var hx_zero_116 *hxrt.Array
					return hx_zero_116
				}
				return hx_value_115.(*hxrt.Array)
			}(hxrt.AnonymousGet(c, "fields"))
			hx_arr_114.Push(f2)
		} else {
			if !hxrt.StringEqualStringPtr(self.curplatform, nil) {
				hx_arr_117 := func(hx_value_118 any) *hxrt.Array {
					if hx_value_118 == nil {
						var hx_zero_119 *hxrt.Array
						return hx_zero_119
					}
					return hx_value_118.(*hxrt.Array)
				}(hxrt.AnonymousGet(found, "platforms"))
				hx_arr_117.Push(self.curplatform)
			}
		}
	}
	_g_2 := 0
	_g1_2 := func(hx_value_120 any) *hxrt.Array {
		if hx_value_120 == nil {
			var hx_zero_121 *hxrt.Array
			return hx_zero_121
		}
		return hx_value_120.(*hxrt.Array)
	}(hxrt.AnonymousGet(c2, "statics"))
	for _g_2 < _g1_2.Len() {
		f2_1 := hx_structural_recover(_g1_2.Get(_g_2))
		_g_2 = int(int32((_g_2 + 1)))
		var found_1 hxrt.Anonymous = nil
		_g_3 := 0
		_g1_3 := func(hx_value_122 any) *hxrt.Array {
			if hx_value_122 == nil {
				var hx_zero_123 *hxrt.Array
				return hx_zero_123
			}
			return hx_value_122.(*hxrt.Array)
		}(hxrt.AnonymousGet(c, "statics"))
		for _g_3 < _g1_3.Len() {
			f_1 := hx_structural_recover(_g1_3.Get(_g_3))
			_g_3 = int(int32((_g_3 + 1)))
			if self.__hx_this.mergeFields(f_1, f2_1) {
				found_1 = f_1
				break
			}
		}
		if found_1 == nil {
			func() {
				var hx_callable_124 *hxrt.Function[func(hxrt.Anonymous, hxrt.Anonymous)] = self.newField
				var hx_argument_125 hxrt.Anonymous = c
				var hx_argument_126 hxrt.Anonymous = f2_1
				hxrt.FunctionCall(hx_callable_124)(hx_argument_125, hx_argument_126)
			}()
			hx_arr_127 := func(hx_value_128 any) *hxrt.Array {
				if hx_value_128 == nil {
					var hx_zero_129 *hxrt.Array
					return hx_zero_129
				}
				return hx_value_128.(*hxrt.Array)
			}(hxrt.AnonymousGet(c, "statics"))
			hx_arr_127.Push(f2_1)
		} else {
			if !hxrt.StringEqualStringPtr(self.curplatform, nil) {
				hx_arr_130 := func(hx_value_131 any) *hxrt.Array {
					if hx_value_131 == nil {
						var hx_zero_132 *hxrt.Array
						return hx_zero_132
					}
					return hx_value_131.(*hxrt.Array)
				}(hxrt.AnonymousGet(found_1, "platforms"))
				hx_arr_130.Push(self.curplatform)
			}
		}
	}
	return true
}

func (self *haxe__rtti__XmlParser) mergeEnums(e hxrt.Anonymous, e2 hxrt.Anonymous) bool {
	if func(hx_value_133 any) bool {
		if hx_value_133 == nil {
			var hx_zero_134 bool
			return hx_zero_134
		}
		return hx_value_133.(bool)
	}(hxrt.AnonymousGet(e, "isExtern")) != func(hx_value_135 any) bool {
		if hx_value_135 == nil {
			var hx_zero_136 bool
			return hx_zero_136
		}
		return hx_value_135.(bool)
	}(hxrt.AnonymousGet(e2, "isExtern")) {
		return false
	}
	if !hxrt.StringEqualStringPtr(self.curplatform, nil) {
		hx_arr_137 := func(hx_value_138 any) *hxrt.Array {
			if hx_value_138 == nil {
				var hx_zero_139 *hxrt.Array
				return hx_zero_139
			}
			return hx_value_138.(*hxrt.Array)
		}(hxrt.AnonymousGet(e, "platforms"))
		hx_arr_137.Push(self.curplatform)
	}
	_g := 0
	_g1 := func(hx_value_140 any) *hxrt.Array {
		if hx_value_140 == nil {
			var hx_zero_141 *hxrt.Array
			return hx_zero_141
		}
		return hx_value_140.(*hxrt.Array)
	}(hxrt.AnonymousGet(e2, "constructors"))
	for _g < _g1.Len() {
		c2 := hx_structural_recover(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		var found hxrt.Anonymous = nil
		_g_1 := 0
		_g1_1 := func(hx_value_142 any) *hxrt.Array {
			if hx_value_142 == nil {
				var hx_zero_143 *hxrt.Array
				return hx_zero_143
			}
			return hx_value_142.(*hxrt.Array)
		}(hxrt.AnonymousGet(e, "constructors"))
		for _g_1 < _g1_1.Len() {
			c := hx_structural_recover(_g1_1.Get(_g_1))
			_g_1 = int(int32((_g_1 + 1)))
			if haxe__rtti__TypeApi_constructorEq(c, c2) {
				found = c
				break
			}
		}
		if found == nil {
			hx_arr_144 := func(hx_value_145 any) *hxrt.Array {
				if hx_value_145 == nil {
					var hx_zero_146 *hxrt.Array
					return hx_zero_146
				}
				return hx_value_145.(*hxrt.Array)
			}(hxrt.AnonymousGet(e, "constructors"))
			hx_arr_144.Push(c2)
		} else {
			if !hxrt.StringEqualStringPtr(self.curplatform, nil) {
				hx_arr_147 := func(hx_value_148 any) *hxrt.Array {
					if hx_value_148 == nil {
						var hx_zero_149 *hxrt.Array
						return hx_zero_149
					}
					return hx_value_148.(*hxrt.Array)
				}(hxrt.AnonymousGet(found, "platforms"))
				hx_arr_147.Push(self.curplatform)
			}
		}
	}
	return true
}

func (self *haxe__rtti__XmlParser) mergeTypedefs(t hxrt.Anonymous, t2 hxrt.Anonymous) bool {
	if hxrt.StringEqualStringPtr(self.curplatform, nil) {
		return false
	}
	hx_arr_150 := func(hx_value_151 any) *hxrt.Array {
		if hx_value_151 == nil {
			var hx_zero_152 *hxrt.Array
			return hx_zero_152
		}
		return hx_value_151.(*hxrt.Array)
	}(hxrt.AnonymousGet(t, "platforms"))
	hx_arr_150.Push(self.curplatform)
	var this1 haxe__IMap = func(hx_value_153 any) *haxe__ds__StringMap {
		if hx_value_153 == nil {
			var hx_zero_154 *haxe__ds__StringMap
			return hx_zero_154
		}
		return hx_value_153.(*haxe__ds__StringMap)
	}(hxrt.AnonymousGet(t, "types"))
	key := self.curplatform
	value := func(hx_value_155 any) *haxe__rtti__CType {
		if hx_value_155 == nil {
			var hx_zero_156 *haxe__rtti__CType
			return hx_zero_156
		}
		return hx_value_155.(*haxe__rtti__CType)
	}(hxrt.AnonymousGet(t2, "type"))
	this1.(*haxe__ds__StringMap).__hx_this.set(key, value)
	return true
}

func (self *haxe__rtti__XmlParser) mergeAbstracts(a hxrt.Anonymous, a2 hxrt.Anonymous) bool {
	if hxrt.StringEqualStringPtr(self.curplatform, nil) {
		return false
	}
	if (func(hx_value_157 any) *hxrt.Array {
		if hx_value_157 == nil {
			var hx_zero_158 *hxrt.Array
			return hx_zero_158
		}
		return hx_value_157.(*hxrt.Array)
	}(hxrt.AnonymousGet(a, "to")).Len() != func(hx_value_159 any) *hxrt.Array {
		if hx_value_159 == nil {
			var hx_zero_160 *hxrt.Array
			return hx_zero_160
		}
		return hx_value_159.(*hxrt.Array)
	}(hxrt.AnonymousGet(a2, "to")).Len()) || (func(hx_value_161 any) *hxrt.Array {
		if hx_value_161 == nil {
			var hx_zero_162 *hxrt.Array
			return hx_zero_162
		}
		return hx_value_161.(*hxrt.Array)
	}(hxrt.AnonymousGet(a, "from")).Len() != func(hx_value_163 any) *hxrt.Array {
		if hx_value_163 == nil {
			var hx_zero_164 *hxrt.Array
			return hx_zero_164
		}
		return hx_value_163.(*hxrt.Array)
	}(hxrt.AnonymousGet(a2, "from")).Len()) {
		return false
	}
	_g := 0
	_g1 := func(hx_value_165 any) *hxrt.Array {
		if hx_value_165 == nil {
			var hx_zero_166 *hxrt.Array
			return hx_zero_166
		}
		return hx_value_165.(*hxrt.Array)
	}(hxrt.AnonymousGet(a, "to")).Len()
	for _g < _g1 {
		hx_post_167 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_167
		if !haxe__rtti__TypeApi_typeEq(func(hx_value_170 any) *haxe__rtti__CType {
			if hx_value_170 == nil {
				var hx_zero_171 *haxe__rtti__CType
				return hx_zero_171
			}
			return hx_value_170.(*haxe__rtti__CType)
		}(hxrt.AnonymousGet(hx_structural_recover(func(hx_value_168 any) *hxrt.Array {
			if hx_value_168 == nil {
				var hx_zero_169 *hxrt.Array
				return hx_zero_169
			}
			return hx_value_168.(*hxrt.Array)
		}(hxrt.AnonymousGet(a, "to")).Get(i)), "t")), func(hx_value_174 any) *haxe__rtti__CType {
			if hx_value_174 == nil {
				var hx_zero_175 *haxe__rtti__CType
				return hx_zero_175
			}
			return hx_value_174.(*haxe__rtti__CType)
		}(hxrt.AnonymousGet(hx_structural_recover(func(hx_value_172 any) *hxrt.Array {
			if hx_value_172 == nil {
				var hx_zero_173 *hxrt.Array
				return hx_zero_173
			}
			return hx_value_172.(*hxrt.Array)
		}(hxrt.AnonymousGet(a2, "to")).Get(i)), "t"))) {
			return false
		}
	}
	_g_1 := 0
	_g1_1 := func(hx_value_176 any) *hxrt.Array {
		if hx_value_176 == nil {
			var hx_zero_177 *hxrt.Array
			return hx_zero_177
		}
		return hx_value_176.(*hxrt.Array)
	}(hxrt.AnonymousGet(a, "from")).Len()
	for _g_1 < _g1_1 {
		hx_post_178 := _g_1
		_g_1 = int(int32((_g_1 + 1)))
		i_1 := hx_post_178
		if !haxe__rtti__TypeApi_typeEq(func(hx_value_181 any) *haxe__rtti__CType {
			if hx_value_181 == nil {
				var hx_zero_182 *haxe__rtti__CType
				return hx_zero_182
			}
			return hx_value_181.(*haxe__rtti__CType)
		}(hxrt.AnonymousGet(hx_structural_recover(func(hx_value_179 any) *hxrt.Array {
			if hx_value_179 == nil {
				var hx_zero_180 *hxrt.Array
				return hx_zero_180
			}
			return hx_value_179.(*hxrt.Array)
		}(hxrt.AnonymousGet(a, "from")).Get(i_1)), "t")), func(hx_value_185 any) *haxe__rtti__CType {
			if hx_value_185 == nil {
				var hx_zero_186 *haxe__rtti__CType
				return hx_zero_186
			}
			return hx_value_185.(*haxe__rtti__CType)
		}(hxrt.AnonymousGet(hx_structural_recover(func(hx_value_183 any) *hxrt.Array {
			if hx_value_183 == nil {
				var hx_zero_184 *hxrt.Array
				return hx_zero_184
			}
			return hx_value_183.(*hxrt.Array)
		}(hxrt.AnonymousGet(a2, "from")).Get(i_1)), "t"))) {
			return false
		}
	}
	if hx_structural_recover(hxrt.AnonymousGet(a2, "impl")) != nil {
		self.__hx_this.mergeClasses(hx_structural_recover(hxrt.AnonymousGet(a, "impl")), hx_structural_recover(hxrt.AnonymousGet(a2, "impl")))
	}
	hx_arr_187 := func(hx_value_188 any) *hxrt.Array {
		if hx_value_188 == nil {
			var hx_zero_189 *hxrt.Array
			return hx_zero_189
		}
		return hx_value_188.(*hxrt.Array)
	}(hxrt.AnonymousGet(a, "platforms"))
	hx_arr_187.Push(self.curplatform)
	return true
}

func (self *haxe__rtti__XmlParser) merge(t *haxe__rtti__TypeTree) {
	inf := haxe__rtti__TypeApi_typeInfos(t)
	pack := self.__hx_this.splitString(func(hx_value_190 any) *string {
		if hx_value_190 == nil {
			var hx_zero_191 *string
			return hx_zero_191
		}
		return hx_value_190.(*string)
	}(hxrt.AnonymousGet(inf, "path")), hxrt.StringFromLiteral("."))
	cur := self.root
	curpack := hxrt.NewArray()
	pack.Pop()
	_g := 0
	for _g < pack.Len() {
		p := func(hx_value_193 any) *string {
			if hx_value_193 == nil {
				var hx_zero_194 *string
				return hx_zero_194
			}
			return hx_value_193.(*string)
		}(pack.Get(_g))
		_g = int(int32((_g + 1)))
		found := false
		_g_1 := 0
		for _g_1 < cur.Len() {
			pk := func(hx_value_195 any) *haxe__rtti__TypeTree {
				if hx_value_195 == nil {
					var hx_zero_196 *haxe__rtti__TypeTree
					return hx_zero_196
				}
				return hx_value_195.(*haxe__rtti__TypeTree)
			}(cur.Get(_g_1))
			_g_1 = int(int32((_g_1 + 1)))
			if pk.tag == 0 {
				_g_2 := func(hx_value_197 any) *string {
					if hx_value_197 == nil {
						var hx_zero_198 *string
						return hx_zero_198
					}
					return hx_value_197.(*string)
				}(pk.params[0])
				_g1 := func(hx_value_199 any) *string {
					if hx_value_199 == nil {
						var hx_zero_200 *string
						return hx_zero_200
					}
					return hx_value_199.(*string)
				}(pk.params[1])
				_ = _g1
				_g1_1 := func(hx_value_201 any) *hxrt.Array {
					if hx_value_201 == nil {
						var hx_zero_202 *hxrt.Array
						return hx_zero_202
					}
					return hx_value_201.(*hxrt.Array)
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
		ct := func(hx_value_205 any) *haxe__rtti__TypeTree {
			if hx_value_205 == nil {
				var hx_zero_206 *haxe__rtti__TypeTree
				return hx_zero_206
			}
			return hx_value_205.(*haxe__rtti__TypeTree)
		}(cur.Get(_g_3))
		_g_3 = int(int32((_g_3 + 1)))
		if func() bool {
			var hx_if_213 bool
			if ct.tag == 0 {
				_g_4 := func(hx_value_207 any) *string {
					if hx_value_207 == nil {
						var hx_zero_208 *string
						return hx_zero_208
					}
					return hx_value_207.(*string)
				}(ct.params[0])
				_ = _g_4
				_g_5 := func(hx_value_209 any) *string {
					if hx_value_209 == nil {
						var hx_zero_210 *string
						return hx_zero_210
					}
					return hx_value_209.(*string)
				}(ct.params[1])
				_ = _g_5
				_g_6 := func(hx_value_211 any) *hxrt.Array {
					if hx_value_211 == nil {
						var hx_zero_212 *hxrt.Array
						return hx_zero_212
					}
					return hx_value_211.(*hxrt.Array)
				}(ct.params[2])
				_ = _g_6
				hx_if_213 = true
			} else {
				hx_if_213 = false
			}
			return hx_if_213
		}() {
			continue
		}
		tinf := haxe__rtti__TypeApi_typeInfos(ct)
		if hxrt.StringEqualStringPtr(func(hx_value_274 any) *string {
			if hx_value_274 == nil {
				var hx_zero_275 *string
				return hx_zero_275
			}
			return hx_value_274.(*string)
		}(hxrt.AnonymousGet(tinf, "path")), func(hx_value_276 any) *string {
			if hx_value_276 == nil {
				var hx_zero_277 *string
				return hx_zero_277
			}
			return hx_value_276.(*string)
		}(hxrt.AnonymousGet(inf, "path"))) {
			sameType := true
			if hxrt.StringEqualStringPtr(func(hx_value_226 any) *string {
				if hx_value_226 == nil {
					var hx_zero_227 *string
					return hx_zero_227
				}
				return hx_value_226.(*string)
			}(hxrt.AnonymousGet(tinf, "doc")), nil) != hxrt.StringEqualStringPtr(func(hx_value_228 any) *string {
				if hx_value_228 == nil {
					var hx_zero_229 *string
					return hx_zero_229
				}
				return hx_value_228.(*string)
			}(hxrt.AnonymousGet(inf, "doc")), nil) {
				if hxrt.StringEqualStringPtr(func(hx_value_224 any) *string {
					if hx_value_224 == nil {
						var hx_zero_225 *string
						return hx_zero_225
					}
					return hx_value_224.(*string)
				}(hxrt.AnonymousGet(inf, "doc")), nil) {
					func() *string {
						hx_structural_target_214 := inf
						hx_structural_value_216 := func(hx_value_217 any) *string {
							if hx_value_217 == nil {
								var hx_zero_218 *string
								return hx_zero_218
							}
							return hx_value_217.(*string)
						}(hxrt.AnonymousGet(tinf, "doc"))
						hxrt.AnonymousSet(hx_structural_target_214, "doc", hx_structural_value_216)
						return hx_structural_value_216
					}()
				} else {
					func() *string {
						hx_structural_target_219 := tinf
						hx_structural_value_221 := func(hx_value_222 any) *string {
							if hx_value_222 == nil {
								var hx_zero_223 *string
								return hx_zero_223
							}
							return hx_value_222.(*string)
						}(hxrt.AnonymousGet(inf, "doc"))
						hxrt.AnonymousSet(hx_structural_target_219, "doc", hx_structural_value_221)
						return hx_structural_value_221
					}()
				}
			}
			if hxrt.StringEqualStringPtr(func(hx_value_230 any) *string {
				if hx_value_230 == nil {
					var hx_zero_231 *string
					return hx_zero_231
				}
				return hx_value_230.(*string)
			}(hxrt.AnonymousGet(tinf, "path")), hxrt.StringFromLiteral("haxe._Int64.NativeInt64")) {
				continue
			}
			if (hxrt.StringEqualStringPtr(func(hx_value_238 any) *string {
				if hx_value_238 == nil {
					var hx_zero_239 *string
					return hx_zero_239
				}
				return hx_value_238.(*string)
			}(hxrt.AnonymousGet(tinf, "module")), func(hx_value_240 any) *string {
				if hx_value_240 == nil {
					var hx_zero_241 *string
					return hx_zero_241
				}
				return hx_value_240.(*string)
			}(hxrt.AnonymousGet(inf, "module"))) && hxrt.StringEqualStringPtr(func(hx_value_242 any) *string {
				if hx_value_242 == nil {
					var hx_zero_243 *string
					return hx_zero_243
				}
				return hx_value_242.(*string)
			}(hxrt.AnonymousGet(tinf, "doc")), func(hx_value_244 any) *string {
				if hx_value_244 == nil {
					var hx_zero_245 *string
					return hx_zero_245
				}
				return hx_value_244.(*string)
			}(hxrt.AnonymousGet(inf, "doc")))) && (func(hx_value_246 any) bool {
				if hx_value_246 == nil {
					var hx_zero_247 bool
					return hx_zero_247
				}
				return hx_value_246.(bool)
			}(hxrt.AnonymousGet(tinf, "isPrivate")) == func(hx_value_248 any) bool {
				if hx_value_248 == nil {
					var hx_zero_249 bool
					return hx_zero_249
				}
				return hx_value_248.(bool)
			}(hxrt.AnonymousGet(inf, "isPrivate"))) {
				switch ct.tag {
				case 0:
					_g_7 := func(hx_value_232 any) *string {
						if hx_value_232 == nil {
							var hx_zero_233 *string
							return hx_zero_233
						}
						return hx_value_232.(*string)
					}(ct.params[0])
					_ = _g_7
					_g_8 := func(hx_value_234 any) *string {
						if hx_value_234 == nil {
							var hx_zero_235 *string
							return hx_zero_235
						}
						return hx_value_234.(*string)
					}(ct.params[1])
					_ = _g_8
					_g_9 := func(hx_value_236 any) *hxrt.Array {
						if hx_value_236 == nil {
							var hx_zero_237 *hxrt.Array
							return hx_zero_237
						}
						return hx_value_236.(*hxrt.Array)
					}(ct.params[2])
					_ = _g_9
					sameType = false
				case 1:
					_g_10 := hx_structural_recover(ct.params[0])
					c := _g_10
					if t.tag == 1 {
						_g_11 := hx_structural_recover(t.params[0])
						c2 := _g_11
						if self.__hx_this.mergeClasses(c, c2) {
							return
						}
					} else {
						sameType = false
					}
				case 2:
					_g_12 := hx_structural_recover(ct.params[0])
					e := _g_12
					if t.tag == 2 {
						_g_13 := hx_structural_recover(t.params[0])
						e2 := _g_13
						if self.__hx_this.mergeEnums(e, e2) {
							return
						}
					} else {
						sameType = false
					}
				case 3:
					_g_14 := hx_structural_recover(ct.params[0])
					td := _g_14
					if t.tag == 3 {
						_g_15 := hx_structural_recover(t.params[0])
						td2 := _g_15
						if self.__hx_this.mergeTypedefs(td, td2) {
							return
						}
					} else {
					}
				case 4:
					_g_16 := hx_structural_recover(ct.params[0])
					a := _g_16
					if t.tag == 4 {
						_g_17 := hx_structural_recover(t.params[0])
						a2 := _g_17
						if self.__hx_this.mergeAbstracts(a, a2) {
							return
						}
					} else {
						sameType = false
					}
				}
			}
			var hx_if_269 *string
			if !hxrt.StringEqualStringPtr(func(hx_value_250 any) *string {
				if hx_value_250 == nil {
					var hx_zero_251 *string
					return hx_zero_251
				}
				return hx_value_250.(*string)
			}(hxrt.AnonymousGet(tinf, "module")), func(hx_value_252 any) *string {
				if hx_value_252 == nil {
					var hx_zero_253 *string
					return hx_zero_253
				}
				return hx_value_252.(*string)
			}(hxrt.AnonymousGet(inf, "module"))) {
				hx_if_269 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("module "), func(hx_value_254 any) *string {
					if hx_value_254 == nil {
						var hx_zero_255 *string
						return hx_zero_255
					}
					return hx_value_254.(*string)
				}(hxrt.AnonymousGet(inf, "module"))), hxrt.StringFromLiteral(" should be ")), func(hx_value_256 any) *string {
					if hx_value_256 == nil {
						var hx_zero_257 *string
						return hx_zero_257
					}
					return hx_value_256.(*string)
				}(hxrt.AnonymousGet(tinf, "module")))
			} else {
				var hx_if_268 *string
				if !hxrt.StringEqualStringPtr(func(hx_value_258 any) *string {
					if hx_value_258 == nil {
						var hx_zero_259 *string
						return hx_zero_259
					}
					return hx_value_258.(*string)
				}(hxrt.AnonymousGet(tinf, "doc")), func(hx_value_260 any) *string {
					if hx_value_260 == nil {
						var hx_zero_261 *string
						return hx_zero_261
					}
					return hx_value_260.(*string)
				}(hxrt.AnonymousGet(inf, "doc"))) {
					hx_if_268 = hxrt.StringFromLiteral("documentation is different")
				} else {
					var hx_if_267 *string
					if func(hx_value_262 any) bool {
						if hx_value_262 == nil {
							var hx_zero_263 bool
							return hx_zero_263
						}
						return hx_value_262.(bool)
					}(hxrt.AnonymousGet(tinf, "isPrivate")) != func(hx_value_264 any) bool {
						if hx_value_264 == nil {
							var hx_zero_265 bool
							return hx_zero_265
						}
						return hx_value_264.(bool)
					}(hxrt.AnonymousGet(inf, "isPrivate")) {
						hx_if_267 = hxrt.StringFromLiteral("private flag is different")
					} else {
						var hx_if_266 *string
						if !sameType {
							hx_if_266 = hxrt.StringFromLiteral("type kind is different")
						} else {
							hx_if_266 = hxrt.StringFromLiteral("could not merge definition")
						}
						hx_if_267 = hx_if_266
					}
					hx_if_268 = hx_if_267
				}
				hx_if_269 = hx_if_268
			}
			msg := hx_if_269
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Incompatibilities between "), func(hx_value_270 any) *string {
				if hx_value_270 == nil {
					var hx_zero_271 *string
					return hx_zero_271
				}
				return hx_value_270.(*string)
			}(hxrt.AnonymousGet(tinf, "path"))), hxrt.StringFromLiteral(" in ")), self.__hx_this.joinStringArray(func(hx_value_272 any) *hxrt.Array {
				if hx_value_272 == nil {
					var hx_zero_273 *hxrt.Array
					return hx_zero_273
				}
				return hx_value_272.(*hxrt.Array)
			}(hxrt.AnonymousGet(tinf, "platforms")), hxrt.StringFromLiteral(","))), hxrt.StringFromLiteral(" and ")), self.curplatform), hxrt.StringFromLiteral(" (")), msg), hxrt.StringFromLiteral(")")))
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
	for func() bool {
		var hx_callable_279 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(c, "hasNext"))
		return hxrt.FunctionCall(hx_callable_279)()
	}() {
		c_1 := func() *Xml {
			var hx_callable_280 *hxrt.Function[func() *Xml] = hx_recover_callable_1(hxrt.AnonymousGet(c, "next"))
			return hxrt.FunctionCall(hx_callable_280)()
		}()
		self.__hx_this.merge(self.__hx_this.processElement(c_1))
	}
}

func (self *haxe__rtti__XmlParser) processElement(x *Xml) *haxe__rtti__TypeTree {
	var hx_if_281 *string
	if hxrt.HaxeEqual(x.nodeType, Xml_Document) {
		hx_if_281 = hxrt.StringFromLiteral("Document")
	} else {
		if !hxrt.HaxeEqual(x.nodeType, Xml_Element) {
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(x.nodeType))))
		}
		hx_if_281 = x.nodeName
	}
	nodeName := hx_if_281
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
	var hx_throw_zero_282 *haxe__rtti__TypeTree
	return hx_throw_zero_282
}

func (self *haxe__rtti__XmlParser) xmeta(x *Xml) *hxrt.Array {
	ml := hxrt.NewArray()
	m := x.__hx_this.elementsNamed(hxrt.StringFromLiteral("m"))
	for func() bool {
		var hx_callable_283 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(m, "hasNext"))
		return hxrt.FunctionCall(hx_callable_283)()
	}() {
		m_1 := func() *Xml {
			var hx_callable_284 *hxrt.Function[func() *Xml] = hx_recover_callable_1(hxrt.AnonymousGet(m, "next"))
			return hxrt.FunctionCall(hx_callable_284)()
		}()
		pl := hxrt.NewArray()
		p := m_1.__hx_this.elementsNamed(hxrt.StringFromLiteral("e"))
		for func() bool {
			var hx_callable_285 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(p, "hasNext"))
			return hxrt.FunctionCall(hx_callable_285)()
		}() {
			p_1 := func() *Xml {
				var hx_callable_286 *hxrt.Function[func() *Xml] = hx_recover_callable_1(hxrt.AnonymousGet(p, "next"))
				return hxrt.FunctionCall(hx_callable_286)()
			}()
			pl.Push(self.__hx_this.innerHTML(p_1))
		}
		hx_obj_289 := map[string]any{}
		hx_obj_289["name"] = self.__hx_this.requireAttr(m_1, hxrt.StringFromLiteral("n"))
		hx_obj_289["params"] = pl
		ml.Push(hxrt.AnonymousFromMap(hx_obj_289))
	}
	return ml
}

func (self *haxe__rtti__XmlParser) xoverloads(x *Xml) *hxrt.Array {
	l := hxrt.NewArray()
	m := x.__hx_this.elements()
	for func() bool {
		var hx_callable_290 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(m, "hasNext"))
		return hxrt.FunctionCall(hx_callable_290)()
	}() {
		m_1 := func() *Xml {
			var hx_callable_291 *hxrt.Function[func() *Xml] = hx_recover_callable_1(hxrt.AnonymousGet(m, "next"))
			return hxrt.FunctionCall(hx_callable_291)()
		}()
		l.Push(self.__hx_this.xclassfield(m_1, false))
	}
	return l
}

func (self *haxe__rtti__XmlParser) xpath(x *Xml) hxrt.Anonymous {
	path := self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("path")))
	params := hxrt.NewArray()
	c := x.__hx_this.elements()
	for func() bool {
		var hx_callable_293 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(c, "hasNext"))
		return hxrt.FunctionCall(hx_callable_293)()
	}() {
		c_1 := func() *Xml {
			var hx_callable_294 *hxrt.Function[func() *Xml] = hx_recover_callable_1(hxrt.AnonymousGet(c, "next"))
			return hxrt.FunctionCall(hx_callable_294)()
		}()
		params.Push(self.__hx_this.xtype(c_1))
	}
	hx_obj_296 := map[string]any{}
	hx_obj_296["path"] = path
	hx_obj_296["params"] = params
	return hxrt.AnonymousFromMap(hx_obj_296)
}

func (self *haxe__rtti__XmlParser) xclass(x *Xml) hxrt.Anonymous {
	var csuper hxrt.Anonymous = nil
	var doc *string = nil
	var tdynamic *haxe__rtti__CType = nil
	interfaces := hxrt.NewArray()
	fields := hxrt.NewArray()
	statics := hxrt.NewArray()
	meta := hxrt.NewArray()
	isInterface := x.__hx_this.exists(hxrt.StringFromLiteral("interface"))
	c := x.__hx_this.elements()
	for func() bool {
		var hx_callable_297 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(c, "hasNext"))
		return hxrt.FunctionCall(hx_callable_297)()
	}() {
		c_1 := func() *Xml {
			var hx_callable_298 *hxrt.Function[func() *Xml] = hx_recover_callable_1(hxrt.AnonymousGet(c, "next"))
			return hxrt.FunctionCall(hx_callable_298)()
		}()
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
	hx_obj_303 := map[string]any{}
	hx_obj_303["file"] = x.__hx_this.get(hxrt.StringFromLiteral("file"))
	hx_obj_303["path"] = self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("path")))
	var hx_if_304 *string
	if x.__hx_this.exists(hxrt.StringFromLiteral("module")) {
		hx_if_304 = self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("module")))
	} else {
		hx_if_304 = nil
	}
	hx_obj_303["module"] = hx_if_304
	hx_obj_303["doc"] = doc
	hx_obj_303["isPrivate"] = x.__hx_this.exists(hxrt.StringFromLiteral("private"))
	hx_obj_303["isExtern"] = x.__hx_this.exists(hxrt.StringFromLiteral("extern"))
	hx_obj_303["isFinal"] = x.__hx_this.exists(hxrt.StringFromLiteral("final"))
	hx_obj_303["isInterface"] = isInterface
	hx_obj_303["params"] = self.__hx_this.mkTypeParams(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("params")))
	hx_obj_303["superClass"] = csuper
	hx_obj_303["interfaces"] = interfaces
	hx_obj_303["fields"] = fields
	hx_obj_303["statics"] = statics
	hx_obj_303["tdynamic"] = tdynamic
	hx_obj_303["platforms"] = self.__hx_this.defplat()
	hx_obj_303["meta"] = meta
	return hxrt.AnonymousFromMap(hx_obj_303)
}

func (self *haxe__rtti__XmlParser) xclassfield(x *Xml, defPublic bool) hxrt.Anonymous {
	e := x.__hx_this.elements()
	t := self.__hx_this.xtype(func() *Xml {
		var hx_callable_305 *hxrt.Function[func() *Xml] = hx_recover_callable_1(hxrt.AnonymousGet(e, "next"))
		return hxrt.FunctionCall(hx_callable_305)()
	}())
	var doc *string = nil
	meta := hxrt.NewArray()
	var overloads *hxrt.Array = nil
	var line any = nil
	if x.__hx_this.exists(hxrt.StringFromLiteral("line")) {
		line = self.__hx_this.parseIntString(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("line")))
	}
	c := e
	for func() bool {
		var hx_callable_306 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(c, "hasNext"))
		return hxrt.FunctionCall(hx_callable_306)()
	}() {
		c_1 := func() *Xml {
			var hx_callable_307 *hxrt.Function[func() *Xml] = hx_recover_callable_1(hxrt.AnonymousGet(c, "next"))
			return hxrt.FunctionCall(hx_callable_307)()
		}()
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
	hx_obj_308 := map[string]any{}
	hx_obj_308["name"] = self.__hx_this.elementName(x)
	hx_obj_308["type"] = t
	hx_obj_308["isPublic"] = (x.__hx_this.exists(hxrt.StringFromLiteral("public")) || func(hx_value_309 any) bool {
		if hx_value_309 == nil {
			var hx_zero_310 bool
			return hx_zero_310
		}
		return hx_value_309.(bool)
	}(defPublic))
	hx_obj_308["isFinal"] = x.__hx_this.exists(hxrt.StringFromLiteral("final"))
	hx_obj_308["isOverride"] = x.__hx_this.exists(hxrt.StringFromLiteral("override"))
	hx_obj_308["line"] = func(hx_value_311 any) any {
		if hx_value_311 == nil {
			return nil
		}
		return hx_value_311.(int)
	}(line)
	hx_obj_308["doc"] = doc
	var hx_if_312 *haxe__rtti__Rights
	if x.__hx_this.exists(hxrt.StringFromLiteral("get")) {
		hx_if_312 = self.__hx_this.mkRights(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("get")))
	} else {
		hx_if_312 = haxe__rtti__Rights_RNormal
	}
	hx_obj_308["get"] = hx_if_312
	var hx_if_313 *haxe__rtti__Rights
	if x.__hx_this.exists(hxrt.StringFromLiteral("set")) {
		hx_if_313 = self.__hx_this.mkRights(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("set")))
	} else {
		hx_if_313 = haxe__rtti__Rights_RNormal
	}
	hx_obj_308["set"] = hx_if_313
	var hx_if_314 *hxrt.Array
	if x.__hx_this.exists(hxrt.StringFromLiteral("params")) {
		hx_if_314 = self.__hx_this.mkTypeParams(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("params")))
	} else {
		hx_if_314 = hxrt.NewArray()
	}
	hx_obj_308["params"] = hx_if_314
	hx_obj_308["platforms"] = self.__hx_this.defplat()
	hx_obj_308["meta"] = meta
	hx_obj_308["overloads"] = overloads
	var hx_if_315 *string
	if x.__hx_this.exists(hxrt.StringFromLiteral("expr")) {
		hx_if_315 = self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("expr"))
	} else {
		hx_if_315 = nil
	}
	hx_obj_308["expr"] = hx_if_315
	return hxrt.AnonymousFromMap(hx_obj_308)
}

func (self *haxe__rtti__XmlParser) xenum(x *Xml) hxrt.Anonymous {
	cl := hxrt.NewArray()
	var doc *string = nil
	meta := hxrt.NewArray()
	c := x.__hx_this.elements()
	for func() bool {
		var hx_callable_316 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(c, "hasNext"))
		return hxrt.FunctionCall(hx_callable_316)()
	}() {
		c_1 := func() *Xml {
			var hx_callable_317 *hxrt.Function[func() *Xml] = hx_recover_callable_1(hxrt.AnonymousGet(c, "next"))
			return hxrt.FunctionCall(hx_callable_317)()
		}()
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
	hx_obj_319 := map[string]any{}
	hx_obj_319["file"] = x.__hx_this.get(hxrt.StringFromLiteral("file"))
	hx_obj_319["path"] = self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("path")))
	var hx_if_320 *string
	if x.__hx_this.exists(hxrt.StringFromLiteral("module")) {
		hx_if_320 = self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("module")))
	} else {
		hx_if_320 = nil
	}
	hx_obj_319["module"] = hx_if_320
	hx_obj_319["doc"] = doc
	hx_obj_319["isPrivate"] = x.__hx_this.exists(hxrt.StringFromLiteral("private"))
	hx_obj_319["isExtern"] = x.__hx_this.exists(hxrt.StringFromLiteral("extern"))
	hx_obj_319["params"] = self.__hx_this.mkTypeParams(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("params")))
	hx_obj_319["constructors"] = cl
	hx_obj_319["platforms"] = self.__hx_this.defplat()
	hx_obj_319["meta"] = meta
	return hxrt.AnonymousFromMap(hx_obj_319)
}

func (self *haxe__rtti__XmlParser) xenumfield(x *Xml) hxrt.Anonymous {
	var args *hxrt.Array = nil
	docElements := x.__hx_this.elementsNamed(hxrt.StringFromLiteral("haxe_doc"))
	var hx_if_323 *Xml
	if func() bool {
		var hx_callable_321 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(docElements, "hasNext"))
		return hxrt.FunctionCall(hx_callable_321)()
	}() {
		hx_if_323 = func() *Xml {
			var hx_callable_322 *hxrt.Function[func() *Xml] = hx_recover_callable_1(hxrt.AnonymousGet(docElements, "next"))
			return hxrt.FunctionCall(hx_callable_322)()
		}()
	} else {
		hx_if_323 = nil
	}
	xdoc := hx_if_323
	var hx_if_324 *hxrt.Array
	if self.__hx_this.hasNamedElement(x, hxrt.StringFromLiteral("meta")) {
		hx_if_324 = self.__hx_this.xmeta(self.__hx_this.requireNamedElement(x, hxrt.StringFromLiteral("meta")))
	} else {
		hx_if_324 = hxrt.NewArray()
	}
	meta := hx_if_324
	if x.__hx_this.exists(hxrt.StringFromLiteral("a")) {
		names := self.__hx_this.splitString(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("a")), hxrt.StringFromLiteral(":"))
		elts := x.__hx_this.elements()
		args = hxrt.NewArray()
		_g := 0
		for _g < names.Len() {
			c := func(hx_value_325 any) *string {
				if hx_value_325 == nil {
					var hx_zero_326 *string
					return hx_zero_326
				}
				return hx_value_325.(*string)
			}(names.Get(_g))
			_g = int(int32((_g + 1)))
			opt := false
			if hxrt.StringEqualStringPtr(hxrt.StringCharAtStringPtr(c, 0), hxrt.StringFromLiteral("?")) {
				opt = true
				c = hxrt.StringSubstrStringPtr(c, 1, 0, false)
			}
			hx_obj_328 := map[string]any{}
			hx_obj_328["name"] = c
			hx_obj_328["opt"] = opt
			hx_obj_328["t"] = self.__hx_this.xtype(func() *Xml {
				var hx_callable_329 *hxrt.Function[func() *Xml] = hx_recover_callable_1(hxrt.AnonymousGet(elts, "next"))
				return hxrt.FunctionCall(hx_callable_329)()
			}())
			args.Push(hxrt.AnonymousFromMap(hx_obj_328))
		}
	}
	hx_obj_330 := map[string]any{}
	hx_obj_330["name"] = self.__hx_this.elementName(x)
	hx_obj_330["args"] = args
	var hx_if_331 *string
	if xdoc == nil {
		hx_if_331 = nil
	} else {
		hx_if_331 = self.__hx_this.innerData(xdoc)
	}
	hx_obj_330["doc"] = hx_if_331
	hx_obj_330["meta"] = meta
	hx_obj_330["platforms"] = self.__hx_this.defplat()
	return hxrt.AnonymousFromMap(hx_obj_330)
}

func (self *haxe__rtti__XmlParser) xabstract(x *Xml) hxrt.Anonymous {
	var doc *string = nil
	var impl hxrt.Anonymous = nil
	var athis *haxe__rtti__CType = nil
	meta := hxrt.NewArray()
	to := hxrt.NewArray()
	from := hxrt.NewArray()
	c := x.__hx_this.elements()
	for func() bool {
		var hx_callable_332 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(c, "hasNext"))
		return hxrt.FunctionCall(hx_callable_332)()
	}() {
		c_1 := func() *Xml {
			var hx_callable_333 *hxrt.Function[func() *Xml] = hx_recover_callable_1(hxrt.AnonymousGet(c, "next"))
			return hxrt.FunctionCall(hx_callable_333)()
		}()
		nodeName := self.__hx_this.elementName(c_1)
		if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("haxe_doc")) {
			doc = self.__hx_this.innerData(c_1)
		} else {
			if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("meta")) {
				meta = self.__hx_this.xmeta(c_1)
			} else {
				if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("to")) {
					t := c_1.__hx_this.elements()
					for func() bool {
						var hx_callable_334 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(t, "hasNext"))
						return hxrt.FunctionCall(hx_callable_334)()
					}() {
						t_1 := func() *Xml {
							var hx_callable_335 *hxrt.Function[func() *Xml] = hx_recover_callable_1(hxrt.AnonymousGet(t, "next"))
							return hxrt.FunctionCall(hx_callable_335)()
						}()
						hx_obj_337 := map[string]any{}
						hx_obj_337["t"] = self.__hx_this.xtype(self.__hx_this.requireFirstElement(t_1))
						hx_obj_337["field"] = t_1.__hx_this.get(hxrt.StringFromLiteral("field"))
						to.Push(hxrt.AnonymousFromMap(hx_obj_337))
					}
				} else {
					if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("from")) {
						t_2 := c_1.__hx_this.elements()
						for func() bool {
							var hx_callable_338 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(t_2, "hasNext"))
							return hxrt.FunctionCall(hx_callable_338)()
						}() {
							t_3 := func() *Xml {
								var hx_callable_339 *hxrt.Function[func() *Xml] = hx_recover_callable_1(hxrt.AnonymousGet(t_2, "next"))
								return hxrt.FunctionCall(hx_callable_339)()
							}()
							hx_obj_341 := map[string]any{}
							hx_obj_341["t"] = self.__hx_this.xtype(self.__hx_this.requireFirstElement(t_3))
							hx_obj_341["field"] = t_3.__hx_this.get(hxrt.StringFromLiteral("field"))
							from.Push(hxrt.AnonymousFromMap(hx_obj_341))
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
	hx_obj_342 := map[string]any{}
	hx_obj_342["file"] = x.__hx_this.get(hxrt.StringFromLiteral("file"))
	hx_obj_342["path"] = self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("path")))
	var hx_if_343 *string
	if x.__hx_this.exists(hxrt.StringFromLiteral("module")) {
		hx_if_343 = self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("module")))
	} else {
		hx_if_343 = nil
	}
	hx_obj_342["module"] = hx_if_343
	hx_obj_342["doc"] = doc
	hx_obj_342["isPrivate"] = x.__hx_this.exists(hxrt.StringFromLiteral("private"))
	hx_obj_342["params"] = self.__hx_this.mkTypeParams(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("params")))
	hx_obj_342["platforms"] = self.__hx_this.defplat()
	hx_obj_342["meta"] = meta
	hx_obj_342["athis"] = athis
	hx_obj_342["to"] = to
	hx_obj_342["from"] = from
	hx_obj_342["impl"] = impl
	return hxrt.AnonymousFromMap(hx_obj_342)
}

func (self *haxe__rtti__XmlParser) xtypedef(x *Xml) hxrt.Anonymous {
	var doc *string = nil
	var t *haxe__rtti__CType = nil
	meta := hxrt.NewArray()
	c := x.__hx_this.elements()
	for func() bool {
		var hx_callable_344 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(c, "hasNext"))
		return hxrt.FunctionCall(hx_callable_344)()
	}() {
		c_1 := func() *Xml {
			var hx_callable_345 *hxrt.Function[func() *Xml] = hx_recover_callable_1(hxrt.AnonymousGet(c, "next"))
			return hxrt.FunctionCall(hx_callable_345)()
		}()
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
	hx_obj_346 := map[string]any{}
	hx_obj_346["file"] = x.__hx_this.get(hxrt.StringFromLiteral("file"))
	hx_obj_346["path"] = self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("path")))
	var hx_if_347 *string
	if x.__hx_this.exists(hxrt.StringFromLiteral("module")) {
		hx_if_347 = self.__hx_this.mkPath(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("module")))
	} else {
		hx_if_347 = nil
	}
	hx_obj_346["module"] = hx_if_347
	hx_obj_346["doc"] = doc
	hx_obj_346["isPrivate"] = x.__hx_this.exists(hxrt.StringFromLiteral("private"))
	hx_obj_346["params"] = self.__hx_this.mkTypeParams(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("params")))
	hx_obj_346["type"] = t
	hx_obj_346["types"] = types
	hx_obj_346["platforms"] = self.__hx_this.defplat()
	hx_obj_346["meta"] = meta
	return hxrt.AnonymousFromMap(hx_obj_346)
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
		var hx_if_348 *hxrt.Array
		if x.__hx_this.exists(hxrt.StringFromLiteral("v")) {
			hx_if_348 = self.__hx_this.splitString(self.__hx_this.requireAttr(x, hxrt.StringFromLiteral("v")), hxrt.StringFromLiteral(":"))
		} else {
			hx_if_348 = nil
		}
		evalues := hx_if_348
		valueIndex := 0
		e := x.__hx_this.elements()
		for func() bool {
			var hx_callable_349 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(e, "hasNext"))
			return hxrt.FunctionCall(hx_callable_349)()
		}() {
			e_1 := func() *Xml {
				var hx_callable_350 *hxrt.Function[func() *Xml] = hx_recover_callable_1(hxrt.AnonymousGet(e, "next"))
				return hxrt.FunctionCall(hx_callable_350)()
			}()
			opt := false
			var hx_if_353 *string
			if argIndex < aname.Len() {
				hx_if_353 = hxrt.StdString(func(hx_value_351 any) *string {
					if hx_value_351 == nil {
						var hx_zero_352 *string
						return hx_zero_352
					}
					return hx_value_351.(*string)
				}(aname.Get(argIndex)))
			} else {
				hx_if_353 = nil
			}
			a := hx_if_353
			argIndex = int(int32((argIndex + 1)))
			if hxrt.StringEqualStringPtr(a, nil) {
				a = hxrt.StringFromLiteral("")
			}
			if hxrt.StringEqualStringPtr(hxrt.StringCharAtStringPtr(a, 0), hxrt.StringFromLiteral("?")) {
				opt = true
				a = hxrt.StringSubstrStringPtr(a, 1, 0, false)
			}
			var hx_if_357 *string
			if (evalues == nil) || (valueIndex >= evalues.Len()) {
				hx_if_357 = nil
			} else {
				hx_post_354 := valueIndex
				valueIndex = int(int32((valueIndex + 1)))
				hx_if_357 = hxrt.StdString(func(hx_value_355 any) *string {
					if hx_value_355 == nil {
						var hx_zero_356 *string
						return hx_zero_356
					}
					return hx_value_355.(*string)
				}(evalues.Get(hx_post_354)))
			}
			v := hx_if_357
			hx_obj_359 := map[string]any{}
			hx_obj_359["name"] = a
			hx_obj_359["opt"] = opt
			hx_obj_359["t"] = self.__hx_this.xtype(e_1)
			var hx_if_360 *string
			if hxrt.StringEqualStringPtr(v, hxrt.StringFromLiteral("")) {
				hx_if_360 = nil
			} else {
				hx_if_360 = v
			}
			hx_obj_359["value"] = hx_if_360
			args.Push(hxrt.AnonymousFromMap(hx_obj_359))
		}
		ret := hx_structural_recover(args.Get(int((hxrt.Int32Wrap(args.Len()) - hxrt.Int32Wrap(1)))))
		callArgs := hxrt.NewArray()
		_g := 0
		_g1 := int((hxrt.Int32Wrap(args.Len()) - hxrt.Int32Wrap(1)))
		for _g < _g1 {
			hx_post_361 := _g
			_g = int(int32((_g + 1)))
			i := hx_post_361
			callArgs.Push(args.Get(i))
		}
		return haxe__rtti__CType_CFunction(callArgs, func(hx_value_363 any) *haxe__rtti__CType {
			if hx_value_363 == nil {
				var hx_zero_364 *haxe__rtti__CType
				return hx_zero_364
			}
			return hx_value_363.(*haxe__rtti__CType)
		}(hxrt.AnonymousGet(ret, "t")))
	}
	if hxrt.StringEqualStringPtr(nodeName, hxrt.StringFromLiteral("a")) {
		fields := hxrt.NewArray()
		f := x.__hx_this.elements()
		for func() bool {
			var hx_callable_365 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(f, "hasNext"))
			return hxrt.FunctionCall(hx_callable_365)()
		}() {
			f_1 := func() *Xml {
				var hx_callable_366 *hxrt.Function[func() *Xml] = hx_recover_callable_1(hxrt.AnonymousGet(f, "next"))
				return hxrt.FunctionCall(hx_callable_366)()
			}()
			f_2 := self.__hx_this.xclassfield(f_1, true)
			func() *hxrt.Array {
				hx_structural_target_367 := f_2
				hx_structural_value_369 := hxrt.NewArray()
				hxrt.AnonymousSet(hx_structural_target_367, "platforms", hx_structural_value_369)
				return hx_structural_value_369
			}()
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
	var hx_throw_zero_371 *haxe__rtti__CType
	return hx_throw_zero_371
}

func (self *haxe__rtti__XmlParser) xtypeparams(x *Xml) *hxrt.Array {
	p := hxrt.NewArray()
	c := x.__hx_this.elements()
	for func() bool {
		var hx_callable_372 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(c, "hasNext"))
		return hxrt.FunctionCall(hx_callable_372)()
	}() {
		c_1 := func() *Xml {
			var hx_callable_373 *hxrt.Function[func() *Xml] = hx_recover_callable_1(hxrt.AnonymousGet(c, "next"))
			return hxrt.FunctionCall(hx_callable_373)()
		}()
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
		hx_post_376 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_376
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
	var hx_if_379 *string
	if hxrt.StringEqualStringPtr(value, nil) {
		hx_if_379 = hxrt.StringFromLiteral("")
	} else {
		hx_if_379 = value
	}
	return hx_if_379
}

func (self *haxe__rtti__XmlParser) hasNamedElement(x *Xml, name *string) bool {
	return func() bool {
		var hx_callable_380 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(x.__hx_this.elementsNamed(name), "hasNext"))
		return hxrt.FunctionCall(hx_callable_380)()
	}()
}

func (self *haxe__rtti__XmlParser) requireNamedElement(x *Xml, name *string) *Xml {
	elements := x.__hx_this.elementsNamed(name)
	if !func() bool {
		var hx_callable_381 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(elements, "hasNext"))
		return hxrt.FunctionCall(hx_callable_381)()
	}() {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(self.__hx_this.nodeDisplayName(x), hxrt.StringFromLiteral(" is missing element ")), name))
	}
	return func() *Xml {
		var hx_callable_382 *hxrt.Function[func() *Xml] = hx_recover_callable_1(hxrt.AnonymousGet(elements, "next"))
		return hxrt.FunctionCall(hx_callable_382)()
	}()
}

func (self *haxe__rtti__XmlParser) requireFirstElement(x *Xml) *Xml {
	first := x.__hx_this.firstElement()
	if first == nil {
		hxrt.Throw(hxrt.StringConcatStringPtr(self.__hx_this.nodeDisplayName(x), hxrt.StringFromLiteral(" is missing first element")))
	}
	return first
}

func (self *haxe__rtti__XmlParser) nodeDisplayName(x *Xml) *string {
	var hx_if_383 *string
	if hxrt.HaxeEqual(x.nodeType, Xml_Document) {
		hx_if_383 = hxrt.StringFromLiteral("Document")
	} else {
		hx_if_383 = self.__hx_this.elementName(x)
	}
	return hx_if_383
}

func (self *haxe__rtti__XmlParser) elementName(x *Xml) *string {
	if !hxrt.HaxeEqual(x.nodeType, Xml_Element) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(hxrt.IntFromNullableAny(x.nodeType))))
	}
	name := x.nodeName
	var hx_if_384 *string
	if hxrt.StringEqualStringPtr(name, nil) {
		hx_if_384 = hxrt.StringFromLiteral("")
	} else {
		hx_if_384 = name
	}
	return hx_if_384
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
	value := hx_recover_class_0(it_array.Get(func() int {
		hx_post_385 := it_current
		it_current = int(int32((it_current + 1)))
		return hx_post_385
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
		hx_post_386 := _g_current
		_g_current = int(int32((_g_current + 1)))
		child := hx_recover_class_0(_g_array.Get(hx_post_386))
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
	var hx_if_387 int
	if negative {
		hx_if_387 = int(-int32(result))
	} else {
		hx_if_387 = result
	}
	return hx_if_387
}
