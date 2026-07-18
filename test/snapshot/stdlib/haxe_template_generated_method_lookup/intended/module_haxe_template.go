package main

import "snapshot/hxrt"

type I_haxe___Template__TokenCursor interface {
}

type haxe___Template__TokenCursor struct {
	__hx_this I_haxe___Template__TokenCursor
	tokens    *hxrt.Array
	index     int
}

func New_haxe___Template__TokenCursor(tokens *hxrt.Array) *haxe___Template__TokenCursor {
	self := &haxe___Template__TokenCursor{}
	self.__hx_this = self
	self.tokens = tokens
	self.index = 0
	return self
}

type I_haxe___Template__ExprCursor interface {
}

type haxe___Template__ExprCursor struct {
	__hx_this I_haxe___Template__ExprCursor
	tokens    *hxrt.Array
	index     int
}

func New_haxe___Template__ExprCursor(tokens *hxrt.Array) *haxe___Template__ExprCursor {
	self := &haxe___Template__ExprCursor{}
	self.__hx_this = self
	self.tokens = tokens
	self.index = 0
	return self
}

type I_haxe__Template interface {
	execute(context any, macros any) *string
	resolve(v *string) any
	parseTokens(data *string) *hxrt.Array
	parseBlock(cursor *haxe___Template__TokenCursor) *haxe___Template__TemplateExpr
	parse(cursor *haxe___Template__TokenCursor) *haxe___Template__TemplateExpr
	parseExpr(data *string) func() any
	makeConst(v *string) func() any
	makePath(e func() any, cursor *haxe___Template__ExprCursor) func() any
	makeExpr(cursor *haxe___Template__ExprCursor) func() any
	skipSpaces(cursor *haxe___Template__ExprCursor)
	makeExpr2(cursor *haxe___Template__ExprCursor) func() any
	run(e *haxe___Template__TemplateExpr)
	popStackValue() any
}

type haxe__Template struct {
	__hx_this I_haxe__Template
	expr      *haxe___Template__TemplateExpr
	context   any
	macros    any
	stack     *hxrt.Array
	output    *string
}

func New_haxe__Template(str *string) *haxe__Template {
	self := &haxe__Template{}
	self.__hx_this = self
	cursor := New_haxe___Template__TokenCursor(self.__hx_this.parseTokens(str))
	self.expr = self.__hx_this.parseBlock(cursor)
	if cursor.index < cursor.tokens.Len() {
		token := func(hx_value_15 any) map[string]any {
			if hx_value_15 == nil {
				var hx_zero_16 map[string]any
				return hx_zero_16
			}
			return hx_value_15.(map[string]any)
		}(cursor.tokens.Get(cursor.index))
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Unexpected '"), hxrt.StdString(func(hx_obj_17 map[string]any) bool {
			hx_field_18 := hx_obj_17["s"]
			if hx_field_18 == nil {
				var hx_zero_19 bool
				return hx_zero_19
			}
			return hx_field_18.(bool)
		}(token))), hxrt.StringFromLiteral("'")))
	}
	return self
}

func (self *haxe__Template) execute(context any, macros any) *string {
	var hx_if_21 any
	if hxrt.AnyEqualsNull(macros) {
		hx_obj_20 := map[string]any{}
		hx_if_21 = hx_obj_20
	} else {
		hx_if_21 = macros
	}
	self.macros = hx_if_21
	self.context = context
	self.stack = hxrt.NewArray()
	self.output = hxrt.StringFromLiteral("")
	self.__hx_this.run(self.expr)
	return self.output
}

func (self *haxe__Template) resolve(v *string) any {
	if hxrt.StringEqualStringPtr(v, hxrt.StringFromLiteral("__current__")) {
		return self.context
	}
	if hxrt.TemplateIsObject(self.context) {
		var value any = Reflect_field(self.context, v)
		if !hxrt.AnyEqualsNull(value) || Reflect_hasField(self.context, v) {
			return value
		}
	}
	_g := 0
	_g1 := self.stack
	for _g < _g1.Len() {
		var ctx any = _g1.Get(_g)
		_g = int(int32((_g + 1)))
		var value_1 any = Reflect_field(ctx, v)
		if !hxrt.AnyEqualsNull(value_1) || Reflect_hasField(ctx, v) {
			return value_1
		}
	}
	return Reflect_field(haxe__Template_globals, v)
}

func (self *haxe__Template) parseTokens(data *string) *hxrt.Array {
	tokens := hxrt.NewArray()
	for haxe__Template_splitter.__hx_this.match(data) {
		p := haxe__Template_splitter.__hx_this.matchedPos()
		if func(hx_obj_27 map[string]any) int {
			hx_field_28 := hx_obj_27["pos"]
			if hx_field_28 == nil {
				var hx_zero_29 int
				return hx_zero_29
			}
			return hx_field_28.(int)
		}(p) > 0 {
			hx_obj_23 := map[string]any{}
			hx_obj_23["p"] = hxrt.StringSubstrStringPtr(data, 0, func(hx_obj_24 map[string]any) int {
				hx_field_25 := hx_obj_24["pos"]
				if hx_field_25 == nil {
					var hx_zero_26 int
					return hx_zero_26
				}
				return hx_field_25.(int)
			}(p), true)
			hx_obj_23["s"] = true
			hx_obj_23["l"] = nil
			tokens.Push(hx_obj_23)
		}
		if hxrt.StringCharCodeAtAnyStringPtr(data, func(hx_obj_38 map[string]any) int {
			hx_field_39 := hx_obj_38["pos"]
			if hx_field_39 == nil {
				var hx_zero_40 int
				return hx_zero_40
			}
			return hx_field_39.(int)
		}(p)) == 58 {
			hx_obj_31 := map[string]any{}
			hx_obj_31["p"] = hxrt.StringSubstrStringPtr(data, int(int32((hxrt.Int32Wrap(func(hx_obj_32 map[string]any) int {
				hx_field_33 := hx_obj_32["pos"]
				if hx_field_33 == nil {
					var hx_zero_34 int
					return hx_zero_34
				}
				return hx_field_33.(int)
			}(p)) + hxrt.Int32Wrap(2)))), int(int32((hxrt.Int32Wrap(func(hx_obj_35 map[string]any) int {
				hx_field_36 := hx_obj_35["len"]
				if hx_field_36 == nil {
					var hx_zero_37 int
					return hx_zero_37
				}
				return hx_field_36.(int)
			}(p)) - hxrt.Int32Wrap(4)))), true)
			hx_obj_31["s"] = false
			hx_obj_31["l"] = nil
			tokens.Push(hx_obj_31)
			data = haxe__Template_splitter.__hx_this.matchedRight()
			continue
		}
		parp := int(int32((hxrt.Int32Wrap(func(hx_obj_41 map[string]any) int {
			hx_field_42 := hx_obj_41["pos"]
			if hx_field_42 == nil {
				var hx_zero_43 int
				return hx_zero_43
			}
			return hx_field_42.(int)
		}(p)) + hxrt.Int32Wrap(func(hx_obj_44 map[string]any) int {
			hx_field_45 := hx_obj_44["len"]
			if hx_field_45 == nil {
				var hx_zero_46 int
				return hx_zero_46
			}
			return hx_field_45.(int)
		}(p)))))
		npar := 1
		params := hxrt.NewArray()
		part := hxrt.StringFromLiteral("")
		for true {
			var c any = hxrt.StringCharCodeAtAnyStringPtr(data, parp)
			parp = int(int32((parp + 1)))
			if c == 40 {
				npar = int(int32((npar + 1)))
			} else {
				if c == 41 {
					npar = int(int32((npar - 1)))
					if npar <= 0 {
						break
					}
				} else {
					if c == nil {
						hxrt.Throw(hxrt.StringFromLiteral("Unclosed macro parenthesis"))
					}
				}
			}
			chunk := hxrt.StringSubstrStringPtr(data, int(int32((hxrt.Int32Wrap(parp) - hxrt.Int32Wrap(1)))), 1, true)
			if (c == 44) && (npar == 1) {
				params.Push(part)
				part = hxrt.StringFromLiteral("")
			} else {
				part = hxrt.StringConcatStringPtr(part, chunk)
			}
		}
		params.Push(part)
		hx_obj_50 := map[string]any{}
		hx_obj_50["p"] = haxe__Template_splitter.__hx_this.matched(2)
		hx_obj_50["s"] = false
		hx_obj_50["l"] = params
		tokens.Push(hx_obj_50)
		data = hxrt.StringSubstrStringPtr(data, parp, int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(data)) - hxrt.Int32Wrap(parp)))), true)
	}
	if hxrt.StringLengthStringPtr(data) > 0 {
		hx_obj_52 := map[string]any{}
		hx_obj_52["p"] = data
		hx_obj_52["s"] = true
		hx_obj_52["l"] = nil
		tokens.Push(hx_obj_52)
	}
	return tokens
}

func (self *haxe__Template) parseBlock(cursor *haxe___Template__TokenCursor) *haxe___Template__TemplateExpr {
	items := hxrt.NewArray()
	for cursor.index < cursor.tokens.Len() {
		t := func(hx_value_53 any) map[string]any {
			if hx_value_53 == nil {
				var hx_zero_54 map[string]any
				return hx_zero_54
			}
			return hx_value_53.(map[string]any)
		}(cursor.tokens.Get(cursor.index))
		if !func(hx_obj_55 map[string]any) bool {
			hx_field_56 := hx_obj_55["s"]
			if hx_field_56 == nil {
				var hx_zero_57 bool
				return hx_zero_57
			}
			return hx_field_56.(bool)
		}(t) && ((hxrt.StringEqualStringPtr(func(hx_obj_58 map[string]any) *string {
			hx_field_59 := hx_obj_58["p"]
			if hx_field_59 == nil {
				var hx_zero_60 *string
				return hx_zero_60
			}
			return hx_field_59.(*string)
		}(t), hxrt.StringFromLiteral("end")) || hxrt.StringEqualStringPtr(func(hx_obj_61 map[string]any) *string {
			hx_field_62 := hx_obj_61["p"]
			if hx_field_62 == nil {
				var hx_zero_63 *string
				return hx_zero_63
			}
			return hx_field_62.(*string)
		}(t), hxrt.StringFromLiteral("else"))) || hxrt.StringEqualStringPtr(hxrt.StringSubstrStringPtr(func(hx_obj_64 map[string]any) *string {
			hx_field_65 := hx_obj_64["p"]
			if hx_field_65 == nil {
				var hx_zero_66 *string
				return hx_zero_66
			}
			return hx_field_65.(*string)
		}(t), 0, 7, true), hxrt.StringFromLiteral("elseif "))) {
			break
		}
		items.Push(self.__hx_this.parse(cursor))
	}
	if items.Len() == 1 {
		return func(hx_value_68 any) *haxe___Template__TemplateExpr {
			if hx_value_68 == nil {
				var hx_zero_69 *haxe___Template__TemplateExpr
				return hx_zero_69
			}
			return hx_value_68.(*haxe___Template__TemplateExpr)
		}(items.Get(0))
	}
	return haxe___Template__TemplateExpr_OpBlock(items)
}

func (self *haxe__Template) parse(cursor *haxe___Template__TokenCursor) *haxe___Template__TemplateExpr {
	var hx_if_73 map[string]any
	if cursor.index >= cursor.tokens.Len() {
		hx_if_73 = nil
	} else {
		hx_post_70 := cursor.index
		cursor.index = int(int32((cursor.index + 1)))
		hx_if_73 = func(hx_value_71 any) map[string]any {
			if hx_value_71 == nil {
				var hx_zero_72 map[string]any
				return hx_zero_72
			}
			return hx_value_71.(map[string]any)
		}(cursor.tokens.Get(hx_post_70))
	}
	t := hx_if_73
	if t == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Unexpected <eof>"))
	}
	p := func(hx_obj_74 map[string]any) *string {
		hx_field_75 := hx_obj_74["p"]
		if hx_field_75 == nil {
			var hx_zero_76 *string
			return hx_zero_76
		}
		return hx_field_75.(*string)
	}(t)
	if func(hx_obj_77 map[string]any) bool {
		hx_field_78 := hx_obj_77["s"]
		if hx_field_78 == nil {
			var hx_zero_79 bool
			return hx_zero_79
		}
		return hx_field_78.(bool)
	}(t) {
		return haxe___Template__TemplateExpr_OpStr(p)
	}
	if func(hx_obj_86 map[string]any) *hxrt.Array {
		hx_field_87 := hx_obj_86["l"]
		if hx_field_87 == nil {
			var hx_zero_88 *hxrt.Array
			return hx_zero_88
		}
		return hx_field_87.(*hxrt.Array)
	}(t) != nil {
		parsedParams := hxrt.NewArray()
		_g := 0
		_g1 := func(hx_obj_80 map[string]any) *hxrt.Array {
			hx_field_81 := hx_obj_80["l"]
			if hx_field_81 == nil {
				var hx_zero_82 *hxrt.Array
				return hx_zero_82
			}
			return hx_field_81.(*hxrt.Array)
		}(t)
		for _g < _g1.Len() {
			param := func(hx_value_83 any) *string {
				if hx_value_83 == nil {
					var hx_zero_84 *string
					return hx_zero_84
				}
				return hx_value_83.(*string)
			}(_g1.Get(_g))
			_g = int(int32((_g + 1)))
			parsedParams.Push(self.__hx_this.parseBlock(New_haxe___Template__TokenCursor(self.__hx_this.parseTokens(param))))
		}
		return haxe___Template__TemplateExpr_OpMacro(p, parsedParams)
	}
	pos := haxe__Template_kwdEnd(p, hxrt.StringFromLiteral("if"))
	if pos > 0 {
		p = hxrt.StringSubstrStringPtr(p, pos, int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(p)) - hxrt.Int32Wrap(pos)))), true)
		e := self.__hx_this.parseExpr(p)
		eif := self.__hx_this.parseBlock(cursor)
		var hx_if_91 map[string]any
		if cursor.index < cursor.tokens.Len() {
			hx_if_91 = func(hx_value_89 any) map[string]any {
				if hx_value_89 == nil {
					var hx_zero_90 map[string]any
					return hx_zero_90
				}
				return hx_value_89.(map[string]any)
			}(cursor.tokens.Get(cursor.index))
		} else {
			hx_if_91 = nil
		}
		nextToken := hx_if_91
		if nextToken == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Unclosed 'if'"))
		}
		var eelse *haxe___Template__TemplateExpr = nil
		if hxrt.StringEqualStringPtr(func(hx_obj_110 map[string]any) *string {
			hx_field_111 := hx_obj_110["p"]
			if hx_field_111 == nil {
				var hx_zero_112 *string
				return hx_zero_112
			}
			return hx_field_111.(*string)
		}(nextToken), hxrt.StringFromLiteral("end")) {
			if cursor.index >= cursor.tokens.Len() {
			} else {
				cursor.tokens.Get(func() int {
					hx_post_92 := cursor.index
					cursor.index = int(int32((cursor.index + 1)))
					return hx_post_92
				}())
			}
		} else {
			if hxrt.StringEqualStringPtr(func(hx_obj_107 map[string]any) *string {
				hx_field_108 := hx_obj_107["p"]
				if hx_field_108 == nil {
					var hx_zero_109 *string
					return hx_zero_109
				}
				return hx_field_108.(*string)
			}(nextToken), hxrt.StringFromLiteral("else")) {
				if cursor.index >= cursor.tokens.Len() {
				} else {
					cursor.tokens.Get(func() int {
						hx_post_93 := cursor.index
						cursor.index = int(int32((cursor.index + 1)))
						return hx_post_93
					}())
				}
				eelse = self.__hx_this.parseBlock(cursor)
				var hx_if_97 map[string]any
				if cursor.index >= cursor.tokens.Len() {
					hx_if_97 = nil
				} else {
					hx_post_94 := cursor.index
					cursor.index = int(int32((cursor.index + 1)))
					hx_if_97 = func(hx_value_95 any) map[string]any {
						if hx_value_95 == nil {
							var hx_zero_96 map[string]any
							return hx_zero_96
						}
						return hx_value_95.(map[string]any)
					}(cursor.tokens.Get(hx_post_94))
				}
				endToken := hx_if_97
				if (endToken == nil) || !hxrt.StringEqualStringPtr(func(hx_obj_98 map[string]any) *string {
					hx_field_99 := hx_obj_98["p"]
					if hx_field_99 == nil {
						var hx_zero_100 *string
						return hx_zero_100
					}
					return hx_field_99.(*string)
				}(endToken), hxrt.StringFromLiteral("end")) {
					hxrt.Throw(hxrt.StringFromLiteral("Unclosed 'else'"))
				}
			} else {
				nextToken["p"] = hxrt.StringSubstrStringPtr(func(hx_obj_101 map[string]any) *string {
					hx_field_102 := hx_obj_101["p"]
					if hx_field_102 == nil {
						var hx_zero_103 *string
						return hx_zero_103
					}
					return hx_field_102.(*string)
				}(nextToken), 4, int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(func(hx_obj_104 map[string]any) *string {
					hx_field_105 := hx_obj_104["p"]
					if hx_field_105 == nil {
						var hx_zero_106 *string
						return hx_zero_106
					}
					return hx_field_105.(*string)
				}(nextToken))) - hxrt.Int32Wrap(4)))), true)
				eelse = self.__hx_this.parse(cursor)
			}
		}
		return haxe___Template__TemplateExpr_OpIf(e, eif, eelse)
	}
	pos = haxe__Template_kwdEnd(p, hxrt.StringFromLiteral("foreach"))
	if pos >= 0 {
		p = hxrt.StringSubstrStringPtr(p, pos, int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(p)) - hxrt.Int32Wrap(pos)))), true)
		e_1 := self.__hx_this.parseExpr(p)
		efor := self.__hx_this.parseBlock(cursor)
		var hx_if_116 map[string]any
		if cursor.index >= cursor.tokens.Len() {
			hx_if_116 = nil
		} else {
			hx_post_113 := cursor.index
			cursor.index = int(int32((cursor.index + 1)))
			hx_if_116 = func(hx_value_114 any) map[string]any {
				if hx_value_114 == nil {
					var hx_zero_115 map[string]any
					return hx_zero_115
				}
				return hx_value_114.(map[string]any)
			}(cursor.tokens.Get(hx_post_113))
		}
		endToken_1 := hx_if_116
		if (endToken_1 == nil) || !hxrt.StringEqualStringPtr(func(hx_obj_117 map[string]any) *string {
			hx_field_118 := hx_obj_117["p"]
			if hx_field_118 == nil {
				var hx_zero_119 *string
				return hx_zero_119
			}
			return hx_field_118.(*string)
		}(endToken_1), hxrt.StringFromLiteral("end")) {
			hxrt.Throw(hxrt.StringFromLiteral("Unclosed 'foreach'"))
		}
		return haxe___Template__TemplateExpr_OpForeach(e_1, efor)
	}
	if haxe__Template_expr_splitter.__hx_this.match(p) {
		return haxe___Template__TemplateExpr_OpExpr(self.__hx_this.parseExpr(p))
	}
	return haxe___Template__TemplateExpr_OpVar(p)
}

func (self *haxe__Template) parseExpr(data *string) func() any {
	tokens := hxrt.NewArray()
	expr := data
	for haxe__Template_expr_splitter.__hx_this.match(data) {
		p := haxe__Template_expr_splitter.__hx_this.matchedPos()
		if func(hx_obj_125 map[string]any) int {
			hx_field_126 := hx_obj_125["pos"]
			if hx_field_126 == nil {
				var hx_zero_127 int
				return hx_zero_127
			}
			return hx_field_126.(int)
		}(p) != 0 {
			hx_obj_121 := map[string]any{}
			hx_obj_121["p"] = hxrt.StringSubstrStringPtr(data, 0, func(hx_obj_122 map[string]any) int {
				hx_field_123 := hx_obj_122["pos"]
				if hx_field_123 == nil {
					var hx_zero_124 int
					return hx_zero_124
				}
				return hx_field_123.(int)
			}(p), true)
			hx_obj_121["s"] = true
			tokens.Push(hx_obj_121)
		}
		token := haxe__Template_expr_splitter.__hx_this.matched(0)
		hx_obj_129 := map[string]any{}
		hx_obj_129["p"] = token
		hx_obj_129["s"] = StringTools_contains(token, hxrt.StringFromLiteral("\""))
		tokens.Push(hx_obj_129)
		data = haxe__Template_expr_splitter.__hx_this.matchedRight()
	}
	if hxrt.StringLengthStringPtr(data) != 0 {
		var _g_s *string
		var _g_offset int
		_g_offset = 0
		_g_s = data
		for _g_offset < hxrt.StringLengthStringPtr(_g_s) {
			var _g_value int
			var _g_key int
			current := _g_offset
			_g_offset = int(int32((_g_offset + 1)))
			code := hxrt.StringCharCodeAtStringPtr(_g_s, current)
			_g_key = current
			_g_value = code
			i := _g_key
			c := _g_value
			if c == 32 {
			} else {
				hx_obj_131 := map[string]any{}
				hx_obj_131["p"] = hxrt.StringSubstrStringPtr(data, i, 0, false)
				hx_obj_131["s"] = true
				tokens.Push(hx_obj_131)
				break
			}
		}
	}
	cursor := New_haxe___Template__ExprCursor(tokens)
	var built func() any
	hxrt.TryCatch(func() {
		built = self.__hx_this.makeExpr(cursor)
		if cursor.index < cursor.tokens.Len() {
			hxrt.Throw(func(hx_obj_136 map[string]any) *string {
				hx_field_137 := hx_obj_136["p"]
				if hx_field_137 == nil {
					var hx_zero_138 *string
					return hx_zero_138
				}
				return hx_field_137.(*string)
			}(func(hx_value_134 any) map[string]any {
				if hx_value_134 == nil {
					var hx_zero_135 map[string]any
					return hx_zero_135
				}
				return hx_value_134.(map[string]any)
			}(cursor.tokens.Get(cursor.index))))
		}
	}, func(hx_caught_132 any) {
		switch hx_typed_133 := hx_caught_132.(type) {
		case *string:
			s := hx_typed_133
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Unexpected '"), s), hxrt.StringFromLiteral("' in ")), expr))
		default:
			hxrt.Throw(hx_caught_132)
		}
	})
	me := self
	_ = me
	wrapped := func() any {
		hx_try_return_139 := false
		var hx_try_value_140 any
		hxrt.TryCatch(func() {
			hx_try_value_140 = built()
			hx_try_return_139 = true
			return
		}, func(hx_caught_141 any) {
			exc := hx_caught_141
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Error : "), hxrt.StdString(exc)), hxrt.StringFromLiteral(" in ")), expr))
		})
		if hx_try_return_139 {
			return hx_try_value_140
		}
		return nil
	}
	return wrapped
}

func (self *haxe__Template) makeConst(v *string) func() any {
	haxe__Template_expr_trim.__hx_this.match(v)
	v = haxe__Template_expr_trim.__hx_this.matched(1)
	if hxrt.StringCharCodeAtAnyStringPtr(v, 0) == 34 {
		str := hxrt.StringSubstrStringPtr(v, 1, int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(v)) - hxrt.Int32Wrap(2)))), true)
		literal := func() any {
			return str
		}
		return literal
	}
	if haxe__Template_expr_int.__hx_this.match(v) {
		i := haxe__Template_parseIntLiteral(v)
		intLiteral := func() any {
			return i
		}
		return intLiteral
	}
	if haxe__Template_expr_float.__hx_this.match(v) {
		f := haxe__Template_parseFloatLiteral(v)
		floatLiteral := func() any {
			return f
		}
		return floatLiteral
	}
	me := self
	resolved := func() any {
		return me.__hx_this.resolve(v)
	}
	return resolved
}

func (self *haxe__Template) makePath(e func() any, cursor *haxe___Template__ExprCursor) func() any {
	var hx_if_145 map[string]any
	if cursor.index < cursor.tokens.Len() {
		hx_if_145 = func(hx_value_143 any) map[string]any {
			if hx_value_143 == nil {
				var hx_zero_144 map[string]any
				return hx_zero_144
			}
			return hx_value_143.(map[string]any)
		}(cursor.tokens.Get(cursor.index))
	} else {
		hx_if_145 = nil
	}
	token := hx_if_145
	if (token == nil) || !hxrt.StringEqualStringPtr(func(hx_obj_146 map[string]any) *string {
		hx_field_147 := hx_obj_146["p"]
		if hx_field_147 == nil {
			var hx_zero_148 *string
			return hx_zero_148
		}
		return hx_field_147.(*string)
	}(token), hxrt.StringFromLiteral(".")) {
		return e
	}
	if cursor.index >= cursor.tokens.Len() {
	} else {
		cursor.tokens.Get(func() int {
			hx_post_149 := cursor.index
			cursor.index = int(int32((cursor.index + 1)))
			return hx_post_149
		}())
	}
	var hx_if_153 map[string]any
	if cursor.index >= cursor.tokens.Len() {
		hx_if_153 = nil
	} else {
		hx_post_150 := cursor.index
		cursor.index = int(int32((cursor.index + 1)))
		hx_if_153 = func(hx_value_151 any) map[string]any {
			if hx_value_151 == nil {
				var hx_zero_152 map[string]any
				return hx_zero_152
			}
			return hx_value_151.(map[string]any)
		}(cursor.tokens.Get(hx_post_150))
	}
	field := hx_if_153
	if (field == nil) || !func(hx_obj_158 map[string]any) bool {
		hx_field_159 := hx_obj_158["s"]
		if hx_field_159 == nil {
			var hx_zero_160 bool
			return hx_zero_160
		}
		return hx_field_159.(bool)
	}(field) {
		var hx_if_157 *string
		if field == nil {
			hx_if_157 = hxrt.StringFromLiteral("<eof>")
		} else {
			hx_if_157 = func(hx_obj_154 map[string]any) *string {
				hx_field_155 := hx_obj_154["p"]
				if hx_field_155 == nil {
					var hx_zero_156 *string
					return hx_zero_156
				}
				return hx_field_155.(*string)
			}(field)
		}
		hxrt.Throw(hx_if_157)
	}
	name := haxe__Template_trimExprToken(func(hx_obj_161 map[string]any) *string {
		hx_field_162 := hx_obj_161["p"]
		if hx_field_162 == nil {
			var hx_zero_163 *string
			return hx_zero_163
		}
		return hx_field_162.(*string)
	}(field))
	return self.__hx_this.makePath(func() any {
		return Reflect_field(e(), name)
	}, cursor)
}

func (self *haxe__Template) makeExpr(cursor *haxe___Template__ExprCursor) func() any {
	return self.__hx_this.makePath(self.__hx_this.makeExpr2(cursor), cursor)
}

func (self *haxe__Template) skipSpaces(cursor *haxe___Template__ExprCursor) {
	for cursor.index < cursor.tokens.Len() {
		if !haxe__Template_isSpaceOnly(func(hx_obj_166 map[string]any) *string {
			hx_field_167 := hx_obj_166["p"]
			if hx_field_167 == nil {
				var hx_zero_168 *string
				return hx_zero_168
			}
			return hx_field_167.(*string)
		}(func(hx_value_164 any) map[string]any {
			if hx_value_164 == nil {
				var hx_zero_165 map[string]any
				return hx_zero_165
			}
			return hx_value_164.(map[string]any)
		}(cursor.tokens.Get(cursor.index)))) {
			return
		}
		cursor.index = int(int32((cursor.index + 1)))
	}
}

func (self *haxe__Template) makeExpr2(cursor *haxe___Template__ExprCursor) func() any {
	self.__hx_this.skipSpaces(cursor)
	var hx_if_172 map[string]any
	if cursor.index >= cursor.tokens.Len() {
		hx_if_172 = nil
	} else {
		hx_post_169 := cursor.index
		cursor.index = int(int32((cursor.index + 1)))
		hx_if_172 = func(hx_value_170 any) map[string]any {
			if hx_value_170 == nil {
				var hx_zero_171 map[string]any
				return hx_zero_171
			}
			return hx_value_170.(map[string]any)
		}(cursor.tokens.Get(hx_post_169))
	}
	token := hx_if_172
	self.__hx_this.skipSpaces(cursor)
	if token == nil {
		hxrt.Throw(hxrt.StringFromLiteral("<eof>"))
	}
	if func(hx_obj_176 map[string]any) bool {
		hx_field_177 := hx_obj_176["s"]
		if hx_field_177 == nil {
			var hx_zero_178 bool
			return hx_zero_178
		}
		return hx_field_177.(bool)
	}(token) {
		return self.__hx_this.makeConst(func(hx_obj_173 map[string]any) *string {
			hx_field_174 := hx_obj_173["p"]
			if hx_field_174 == nil {
				var hx_zero_175 *string
				return hx_zero_175
			}
			return hx_field_174.(*string)
		}(token))
	}
	_g := func(hx_obj_179 map[string]any) *string {
		hx_field_180 := hx_obj_179["p"]
		if hx_field_180 == nil {
			var hx_zero_181 *string
			return hx_zero_181
		}
		return hx_field_180.(*string)
	}(token)
	switch *hxrt.StdString(_g) {
	case *hxrt.StdString(hxrt.StringFromLiteral("!")):
		inner := self.__hx_this.makeExpr(cursor)
		return func() any {
			var value any = inner()
			return (hxrt.AnyEqualsNull(value) || hxrt.HaxeEqual(value, false))
		}
	case *hxrt.StdString(hxrt.StringFromLiteral("(")):
		self.__hx_this.skipSpaces(cursor)
		e1 := self.__hx_this.makeExpr(cursor)
		self.__hx_this.skipSpaces(cursor)
		var hx_if_185 map[string]any
		if cursor.index >= cursor.tokens.Len() {
			hx_if_185 = nil
		} else {
			hx_post_182 := cursor.index
			cursor.index = int(int32((cursor.index + 1)))
			hx_if_185 = func(hx_value_183 any) map[string]any {
				if hx_value_183 == nil {
					var hx_zero_184 map[string]any
					return hx_zero_184
				}
				return hx_value_183.(map[string]any)
			}(cursor.tokens.Get(hx_post_182))
		}
		op := hx_if_185
		if (op == nil) || func(hx_obj_190 map[string]any) bool {
			hx_field_191 := hx_obj_190["s"]
			if hx_field_191 == nil {
				var hx_zero_192 bool
				return hx_zero_192
			}
			return hx_field_191.(bool)
		}(op) {
			var hx_if_189 *string
			if op == nil {
				hx_if_189 = hxrt.StringFromLiteral("<eof>")
			} else {
				hx_if_189 = func(hx_obj_186 map[string]any) *string {
					hx_field_187 := hx_obj_186["p"]
					if hx_field_187 == nil {
						var hx_zero_188 *string
						return hx_zero_188
					}
					return hx_field_187.(*string)
				}(op)
			}
			hxrt.Throw(hx_if_189)
		}
		if hxrt.StringEqualStringPtr(func(hx_obj_193 map[string]any) *string {
			hx_field_194 := hx_obj_193["p"]
			if hx_field_194 == nil {
				var hx_zero_195 *string
				return hx_zero_195
			}
			return hx_field_194.(*string)
		}(op), hxrt.StringFromLiteral(")")) {
			return e1
		}
		self.__hx_this.skipSpaces(cursor)
		e2 := self.__hx_this.makeExpr(cursor)
		self.__hx_this.skipSpaces(cursor)
		var hx_if_199 map[string]any
		if cursor.index >= cursor.tokens.Len() {
			hx_if_199 = nil
		} else {
			hx_post_196 := cursor.index
			cursor.index = int(int32((cursor.index + 1)))
			hx_if_199 = func(hx_value_197 any) map[string]any {
				if hx_value_197 == nil {
					var hx_zero_198 map[string]any
					return hx_zero_198
				}
				return hx_value_197.(map[string]any)
			}(cursor.tokens.Get(hx_post_196))
		}
		close := hx_if_199
		self.__hx_this.skipSpaces(cursor)
		if (close == nil) || !hxrt.StringEqualStringPtr(func(hx_obj_204 map[string]any) *string {
			hx_field_205 := hx_obj_204["p"]
			if hx_field_205 == nil {
				var hx_zero_206 *string
				return hx_zero_206
			}
			return hx_field_205.(*string)
		}(close), hxrt.StringFromLiteral(")")) {
			var hx_if_203 *string
			if close == nil {
				hx_if_203 = hxrt.StringFromLiteral("<eof>")
			} else {
				hx_if_203 = func(hx_obj_200 map[string]any) *string {
					hx_field_201 := hx_obj_200["p"]
					if hx_field_201 == nil {
						var hx_zero_202 *string
						return hx_zero_202
					}
					return hx_field_201.(*string)
				}(close)
			}
			hxrt.Throw(hx_if_203)
		}
		_g_1 := func(hx_obj_207 map[string]any) *string {
			hx_field_208 := hx_obj_207["p"]
			if hx_field_208 == nil {
				var hx_zero_209 *string
				return hx_zero_209
			}
			return hx_field_208.(*string)
		}(op)
		var hx_switch_210 func() any
		switch *hxrt.StdString(_g_1) {
		case *hxrt.StdString(hxrt.StringFromLiteral("!=")):
			hx_switch_210 = func() any {
				return !hxrt.HaxeEqual(e1(), e2())
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("&&")):
			hx_switch_210 = func() any {
				return (haxe__Template_valueAsBool(e1()) && haxe__Template_valueAsBool(e2()))
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("*")):
			hx_switch_210 = func() any {
				var left any = e1()
				var right any = e2()
				return (haxe__Template_valueAsFloat(left) * haxe__Template_valueAsFloat(right))
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("+")):
			hx_switch_210 = func() any {
				return haxe__Template_addValues(e1(), e2())
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("-")):
			hx_switch_210 = func() any {
				var left any = e1()
				var right any = e2()
				return (haxe__Template_valueAsFloat(left) - haxe__Template_valueAsFloat(right))
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("/")):
			hx_switch_210 = func() any {
				var left any = e1()
				var right any = e2()
				return (haxe__Template_valueAsFloat(left) / haxe__Template_valueAsFloat(right))
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("<")):
			hx_switch_210 = func() any {
				return (haxe__Template_compareValues(e1(), e2()) < 0)
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("<=")):
			hx_switch_210 = func() any {
				return (haxe__Template_compareValues(e1(), e2()) <= 0)
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("==")):
			hx_switch_210 = func() any {
				return hxrt.HaxeEqual(e1(), e2())
			}
		case *hxrt.StdString(hxrt.StringFromLiteral(">")):
			hx_switch_210 = func() any {
				return (haxe__Template_compareValues(e1(), e2()) > 0)
			}
		case *hxrt.StdString(hxrt.StringFromLiteral(">=")):
			hx_switch_210 = func() any {
				return (haxe__Template_compareValues(e1(), e2()) >= 0)
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("||")):
			hx_switch_210 = func() any {
				return (haxe__Template_valueAsBool(e1()) || haxe__Template_valueAsBool(e2()))
			}
		default:
			hx_switch_210 = func() func() any {
				hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Unknown operation "), func(hx_obj_211 map[string]any) *string {
					hx_field_212 := hx_obj_211["p"]
					if hx_field_212 == nil {
						var hx_zero_213 *string
						return hx_zero_213
					}
					return hx_field_212.(*string)
				}(op)))
				var hx_throw_zero_214 func() any
				return hx_throw_zero_214
			}()
		}
		return hx_switch_210
	case *hxrt.StdString(hxrt.StringFromLiteral("-")):
		inner_1 := self.__hx_this.makeExpr(cursor)
		return func() any {
			return -haxe__Template_valueAsFloat(inner_1())
		}
	default:
		hxrt.Throw(func(hx_obj_215 map[string]any) *string {
			hx_field_216 := hx_obj_215["p"]
			if hx_field_216 == nil {
				var hx_zero_217 *string
				return hx_zero_217
			}
			return hx_field_216.(*string)
		}(token))
		var hx_throw_zero_218 func() any
		return hx_throw_zero_218
	}
}

func (self *haxe__Template) run(e *haxe___Template__TemplateExpr) {
	switch e.tag {
	case 0:
		_g := e.params[0].(*string)
		v := _g
		self.output = hxrt.StringConcatStringPtr(self.output, hxrt.StdString(self.__hx_this.resolve(v)))
	case 1:
		_g_1 := e.params[0].(func() any)
		expr := _g_1
		self.output = hxrt.StringConcatStringPtr(self.output, hxrt.StdString(expr()))
	case 2:
		_g_2 := e.params[0].(func() any)
		_g1 := e.params[1].(*haxe___Template__TemplateExpr)
		_g2 := e.params[2].(*haxe___Template__TemplateExpr)
		expr_1 := _g_2
		ifExpr := _g1
		elseExpr := _g2
		var value any = expr_1()
		if hxrt.AnyEqualsNull(value) || hxrt.HaxeEqual(value, false) {
			if elseExpr != nil {
				self.__hx_this.run(elseExpr)
			}
		} else {
			self.__hx_this.run(ifExpr)
		}
	case 3:
		_g_3 := e.params[0].(*string)
		str := _g_3
		self.output = hxrt.StringConcatStringPtr(self.output, str)
	case 4:
		_g_4 := e.params[0].(*hxrt.Array)
		items := _g_4
		_g_5 := 0
		for _g_5 < items.Len() {
			item := func(hx_value_219 any) *haxe___Template__TemplateExpr {
				if hx_value_219 == nil {
					var hx_zero_220 *haxe___Template__TemplateExpr
					return hx_zero_220
				}
				return hx_value_219.(*haxe___Template__TemplateExpr)
			}(items.Get(_g_5))
			_g_5 = int(int32((_g_5 + 1)))
			self.__hx_this.run(item)
		}
	case 5:
		_g_6 := e.params[0].(func() any)
		_g1_1 := e.params[1].(*haxe___Template__TemplateExpr)
		expr_2 := _g_6
		loop := _g1_1
		var value_1 any = expr_2()
		arrayValues := hxrt.TemplateArrayValues(value_1)
		if arrayValues != nil {
			hx_arr_221 := self.stack
			hx_arr_221.Push(self.context)
			_g_7 := 0
			for _g_7 < len(arrayValues) {
				var ctx any = arrayValues[_g_7]
				_g_7 = int(int32((_g_7 + 1)))
				self.context = ctx
				self.__hx_this.run(loop)
			}
			self.context = self.__hx_this.popStackValue()
			return
		}
		var iterator any = nil
		hxrt.TryCatch(func() {
			var iteratorField any = Reflect_field(value_1, hxrt.StringFromLiteral("iterator"))
			if hxrt.AnyEqualsNull(iteratorField) {
				hxrt.Throw(nil)
			}
			var candidate any = hxrt.TemplateCall(iteratorField, []any{})
			if !Reflect_hasField(candidate, hxrt.StringFromLiteral("hasNext")) {
				hxrt.Throw(nil)
			}
			iterator = candidate
		}, func(hx_caught_222 any) {
			hx_tmp := hx_caught_222
			_ = hx_tmp
			hxrt.TryCatch(func() {
				if hxrt.AnyEqualsNull(value_1) || !Reflect_hasField(value_1, hxrt.StringFromLiteral("hasNext")) {
					hxrt.Throw(nil)
				}
				iterator = value_1
			}, func(hx_caught_224 any) {
				hx_tmp_1 := hx_caught_224
				_ = hx_tmp_1
				hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Cannot iter on "), hxrt.StdString(value_1)))
			})
		})
		var hasNext any = Reflect_field(iterator, hxrt.StringFromLiteral("hasNext"))
		var next any = Reflect_field(iterator, hxrt.StringFromLiteral("next"))
		if hxrt.AnyEqualsNull(hasNext) || hxrt.AnyEqualsNull(next) {
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Cannot iter on "), hxrt.StdString(value_1)))
		}
		hx_arr_226 := self.stack
		hx_arr_226.Push(self.context)
		for hxrt.HaxeEqual(hxrt.TemplateCall(hasNext, []any{}), true) {
			self.context = hxrt.TemplateCall(next, []any{})
			self.__hx_this.run(loop)
		}
		self.context = self.__hx_this.popStackValue()
	case 6:
		_g_8 := e.params[0].(*string)
		_g1_2 := e.params[1].(*hxrt.Array)
		name := _g_8
		params := _g1_2
		var fn any = Reflect_field(self.macros, name)
		callArgs := hxrt.NewArray()
		callArgs.Push(self.resolve)
		_g_9 := 0
		for _g_9 < params.Len() {
			param := func(hx_value_228 any) *haxe___Template__TemplateExpr {
				if hx_value_228 == nil {
					var hx_zero_229 *haxe___Template__TemplateExpr
					return hx_zero_229
				}
				return hx_value_228.(*haxe___Template__TemplateExpr)
			}(params.Get(_g_9))
			_g_9 = int(int32((_g_9 + 1)))
			if param.tag == 0 {
				_g_10 := param.params[0].(*string)
				value_2 := _g_10
				callArgs.Push(self.__hx_this.resolve(value_2))
			} else {
				previous := self.output
				self.output = hxrt.StringFromLiteral("")
				self.__hx_this.run(param)
				callArgs.Push(self.output)
				self.output = previous
			}
		}
		hxrt.TryCatch(func() {
			self.output = hxrt.StringConcatStringPtr(self.output, hxrt.StdString(hxrt.TemplateCall(fn, callArgs.ValuesCopy())))
		}, func(hx_caught_232 any) {
			err := hx_caught_232
			var hx_try_234 *string
			hxrt.TryCatch(func() {
				hx_try_234 = haxe__Template_joinDynamicArgs(callArgs)
			}, func(hx_caught_235 any) {
				hx_tmp_2 := hx_caught_235
				_ = hx_tmp_2
				hx_try_234 = hxrt.StringFromLiteral("???")
			})
			argsText := hx_try_234
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Macro call "), name), hxrt.StringFromLiteral("(")), argsText), hxrt.StringFromLiteral(") failed (")), hxrt.StdString(err)), hxrt.StringFromLiteral(")")))
		})
	}
}

func (self *haxe__Template) popStackValue() any {
	lastIndex := int(int32((hxrt.Int32Wrap(self.stack.Len()) - hxrt.Int32Wrap(1))))
	var value any = self.stack.Get(lastIndex)
	remaining := hxrt.NewArray()
	_g := 0
	_g1 := lastIndex
	for _g < _g1 {
		hx_post_237 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_237
		remaining.Push(self.stack.Get(index))
	}
	self.stack = remaining
	return value
}

func haxe__Template_addValues(left any, right any) any {
	if func(hx_value any) bool {
		switch hx_value.(type) {
		case *string:
			return true
		case string:
			return true
		default:
			return false
		}
	}(any(left)) || func(hx_value any) bool {
		switch hx_value.(type) {
		case *string:
			return true
		case string:
			return true
		default:
			return false
		}
	}(any(right)) {
		return hxrt.StringConcatStringPtr(hxrt.StdString(left), hxrt.StdString(right))
	}
	return (haxe__Template_valueAsFloat(left) + haxe__Template_valueAsFloat(right))
}

func haxe__Template_compareValues(left any, right any) int {
	leftNumeric := (func(hx_value any) bool {
		switch hx_value.(type) {
		case int:
			return true
		case int8:
			return true
		case int16:
			return true
		case int32:
			return true
		case int64:
			return true
		case uint:
			return true
		case uint8:
			return true
		case uint16:
			return true
		case uint32:
			return true
		case uint64:
			return true
		case uintptr:
			return true
		default:
			return false
		}
	}(any(left)) || func(hx_value any) bool {
		switch hx_value.(type) {
		case int:
			return true
		case int8:
			return true
		case int16:
			return true
		case int32:
			return true
		case int64:
			return true
		case uint:
			return true
		case uint8:
			return true
		case uint16:
			return true
		case uint32:
			return true
		case uint64:
			return true
		case uintptr:
			return true
		case float32:
			return true
		case float64:
			return true
		default:
			return false
		}
	}(any(left)))
	rightNumeric := (func(hx_value any) bool {
		switch hx_value.(type) {
		case int:
			return true
		case int8:
			return true
		case int16:
			return true
		case int32:
			return true
		case int64:
			return true
		case uint:
			return true
		case uint8:
			return true
		case uint16:
			return true
		case uint32:
			return true
		case uint64:
			return true
		case uintptr:
			return true
		default:
			return false
		}
	}(any(right)) || func(hx_value any) bool {
		switch hx_value.(type) {
		case int:
			return true
		case int8:
			return true
		case int16:
			return true
		case int32:
			return true
		case int64:
			return true
		case uint:
			return true
		case uint8:
			return true
		case uint16:
			return true
		case uint32:
			return true
		case uint64:
			return true
		case uintptr:
			return true
		case float32:
			return true
		case float64:
			return true
		default:
			return false
		}
	}(any(right)))
	if leftNumeric && rightNumeric {
		leftFloat := haxe__Template_valueAsFloat(left)
		rightFloat := haxe__Template_valueAsFloat(right)
		var hx_if_240 int
		if leftFloat < rightFloat {
			hx_if_240 = -1
		} else {
			var hx_if_239 int
			if leftFloat > rightFloat {
				hx_if_239 = 1
			} else {
				hx_if_239 = 0
			}
			hx_if_240 = hx_if_239
		}
		return hx_if_240
	}
	return Reflect_compare(hxrt.StdString(left), hxrt.StdString(right))
}

func haxe__Template_divideValues(left any, right any) float64 {
	return (haxe__Template_valueAsFloat(left) / haxe__Template_valueAsFloat(right))
}

var haxe__Template_expr_float *EReg = New_EReg(hxrt.StringFromLiteral("^[+-]?([0-9]+(,[0-9]*)?|,[0-9]+)([Ee][+-]?[0-9]+)?$"), hxrt.StringFromLiteral(""))

var haxe__Template_expr_int *EReg = New_EReg(hxrt.StringFromLiteral("^[0-9]+$"), hxrt.StringFromLiteral(""))

var haxe__Template_expr_splitter *EReg = New_EReg(hxrt.StringFromLiteral("(\\(|\\)|[ \r\n\t]*\"[^\"]*\"[ \r\n\t]*|[!+=/><*.&|-]+)"), hxrt.StringFromLiteral(""))

var haxe__Template_expr_trim *EReg = New_EReg(hxrt.StringFromLiteral("^[ ]*([^ ]+)[ ]*$"), hxrt.StringFromLiteral(""))

var haxe__Template_globals any = any(func() map[string]any {
	hx_obj_241 := map[string]any{}
	return hx_obj_241
}())

func haxe__Template_isSpaceOnly(value *string) bool {
	var _g_s *string
	var _g_offset int
	_g_offset = 0
	_g_s = value
	for _g_offset < hxrt.StringLengthStringPtr(_g_s) {
		code := hxrt.StringCharCodeAtStringPtr(_g_s, func() int {
			hx_post_242 := _g_offset
			_g_offset = int(int32((_g_offset + 1)))
			return hx_post_242
		}())
		if code != 32 {
			return false
		}
	}
	return true
}

func haxe__Template_joinDynamicArgs(values *hxrt.Array) *string {
	out := hxrt.StringFromLiteral("")
	_g := 0
	_g1 := values.Len()
	for _g < _g1 {
		hx_post_243 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_243
		if index > 0 {
			out = hxrt.StringConcatStringPtr(out, hxrt.StringFromLiteral(","))
		}
		out = hxrt.StringConcatStringPtr(out, hxrt.StdString(values.Get(index)))
	}
	return out
}

func haxe__Template_kwdEnd(value *string, keyword *string) int {
	pos := -1
	length := hxrt.StringLengthStringPtr(keyword)
	if hxrt.StringEqualStringPtr(hxrt.StringSubstrStringPtr(value, 0, length, true), keyword) {
		pos = length
		var _g_s *string
		var _g_offset int
		s := hxrt.StringSubstrStringPtr(value, length, 0, false)
		_g_offset = 0
		_g_s = s
		for _g_offset < hxrt.StringLengthStringPtr(_g_s) {
			code := hxrt.StringCharCodeAtStringPtr(_g_s, func() int {
				hx_post_244 := _g_offset
				_g_offset = int(int32((_g_offset + 1)))
				return hx_post_244
			}())
			if code == 32 {
				pos = int(int32((pos + 1)))
			} else {
				break
			}
		}
	}
	return pos
}

func haxe__Template_multiplyValues(left any, right any) float64 {
	return (haxe__Template_valueAsFloat(left) * haxe__Template_valueAsFloat(right))
}

func haxe__Template_parseFloatLiteral(value *string) float64 {
	normalized := StringTools_replace(value, hxrt.StringFromLiteral(","), hxrt.StringFromLiteral("."))
	index := 0
	sign := 1.0
	if hxrt.StringCharCodeAtAnyStringPtr(normalized, index) == 45 {
		sign = -1.0
		index = int(int32((index + 1)))
	} else {
		if hxrt.StringCharCodeAtAnyStringPtr(normalized, index) == 43 {
			index = int(int32((index + 1)))
		}
	}
	intPart := 0.0
	for index < hxrt.StringLengthStringPtr(normalized) {
		code := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(normalized, index))
		if (code < 48) || (code > 57) {
			break
		}
		intPart = ((intPart * 10.0) + float64(int(int32((hxrt.Int32Wrap(code) - hxrt.Int32Wrap(48))))))
		index = int(int32((index + 1)))
	}
	fracPart := 0.0
	divisor := 1.0
	if (index < hxrt.StringLengthStringPtr(normalized)) && (hxrt.StringCharCodeAtAnyStringPtr(normalized, index) == 46) {
		index = int(int32((index + 1)))
		for index < hxrt.StringLengthStringPtr(normalized) {
			code_1 := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(normalized, index))
			if (code_1 < 48) || (code_1 > 57) {
				break
			}
			fracPart = ((fracPart * 10.0) + float64(int(int32((hxrt.Int32Wrap(code_1) - hxrt.Int32Wrap(48))))))
			divisor = (divisor * 10.0)
			index = int(int32((index + 1)))
		}
	}
	result := (intPart + (fracPart / divisor))
	if index < hxrt.StringLengthStringPtr(normalized) {
		var exponentCode any = hxrt.StringCharCodeAtAnyStringPtr(normalized, index)
		if (exponentCode == 101) || (exponentCode == 69) {
			index = int(int32((index + 1)))
			exponentSign := 1
			if hxrt.StringCharCodeAtAnyStringPtr(normalized, index) == 45 {
				exponentSign = -1
				index = int(int32((index + 1)))
			} else {
				if hxrt.StringCharCodeAtAnyStringPtr(normalized, index) == 43 {
					index = int(int32((index + 1)))
				}
			}
			exponent := 0
			for index < hxrt.StringLengthStringPtr(normalized) {
				code_2 := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(normalized, index))
				if (code_2 < 48) || (code_2 > 57) {
					break
				}
				exponent = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(exponent) * hxrt.Int32Wrap(10))))) + hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(code_2) - hxrt.Int32Wrap(48))))))))
				index = int(int32((index + 1)))
			}
			for exponent > 0 {
				var hx_if_245 float64
				if exponentSign < 0 {
					hx_if_245 = (result / 10.0)
				} else {
					hx_if_245 = (result * 10.0)
				}
				result = hx_if_245
				exponent = int(int32((exponent - 1)))
			}
		}
	}
	return (sign * result)
}

func haxe__Template_parseIntLiteral(value *string) int {
	out := 0
	index := 0
	for index < hxrt.StringLengthStringPtr(value) {
		code := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(value, index))
		out = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(out) * hxrt.Int32Wrap(10))))) + hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(code) - hxrt.Int32Wrap(48))))))))
		index = int(int32((index + 1)))
	}
	return out
}

func haxe__Template_peekExprToken(cursor *haxe___Template__ExprCursor) map[string]any {
	var hx_if_248 map[string]any
	if cursor.index < cursor.tokens.Len() {
		hx_if_248 = func(hx_value_246 any) map[string]any {
			if hx_value_246 == nil {
				var hx_zero_247 map[string]any
				return hx_zero_247
			}
			return hx_value_246.(map[string]any)
		}(cursor.tokens.Get(cursor.index))
	} else {
		hx_if_248 = nil
	}
	return hx_if_248
}

func haxe__Template_peekToken(cursor *haxe___Template__TokenCursor) map[string]any {
	var hx_if_251 map[string]any
	if cursor.index < cursor.tokens.Len() {
		hx_if_251 = func(hx_value_249 any) map[string]any {
			if hx_value_249 == nil {
				var hx_zero_250 map[string]any
				return hx_zero_250
			}
			return hx_value_249.(map[string]any)
		}(cursor.tokens.Get(cursor.index))
	} else {
		hx_if_251 = nil
	}
	return hx_if_251
}

func haxe__Template_popExprToken(cursor *haxe___Template__ExprCursor) map[string]any {
	if cursor.index >= cursor.tokens.Len() {
		return nil
	}
	hx_post_252 := cursor.index
	cursor.index = int(int32((cursor.index + 1)))
	return func(hx_value_253 any) map[string]any {
		if hx_value_253 == nil {
			var hx_zero_254 map[string]any
			return hx_zero_254
		}
		return hx_value_253.(map[string]any)
	}(cursor.tokens.Get(hx_post_252))
}

func haxe__Template_popToken(cursor *haxe___Template__TokenCursor) map[string]any {
	if cursor.index >= cursor.tokens.Len() {
		return nil
	}
	hx_post_255 := cursor.index
	cursor.index = int(int32((cursor.index + 1)))
	return func(hx_value_256 any) map[string]any {
		if hx_value_256 == nil {
			var hx_zero_257 map[string]any
			return hx_zero_257
		}
		return hx_value_256.(map[string]any)
	}(cursor.tokens.Get(hx_post_255))
}

var haxe__Template_splitter *EReg = New_EReg(hxrt.StringFromLiteral("(::[A-Za-z0-9_ ()&|!+=/><*.\"-]+::|\\$\\$([A-Za-z0-9_-]+)\\()"), hxrt.StringFromLiteral(""))

func haxe__Template_subtractValues(left any, right any) float64 {
	return (haxe__Template_valueAsFloat(left) - haxe__Template_valueAsFloat(right))
}

func haxe__Template_trimExprToken(value *string) *string {
	haxe__Template_expr_trim.__hx_this.match(value)
	return haxe__Template_expr_trim.__hx_this.matched(1)
}

func haxe__Template_valueAsBool(value any) bool {
	return !(hxrt.AnyEqualsNull(value) || hxrt.HaxeEqual(value, false))
}

func haxe__Template_valueAsFloat(value any) float64 {
	if func(hx_value any) bool {
		switch hx_value.(type) {
		case int:
			return true
		case int8:
			return true
		case int16:
			return true
		case int32:
			return true
		case int64:
			return true
		case uint:
			return true
		case uint8:
			return true
		case uint16:
			return true
		case uint32:
			return true
		case uint64:
			return true
		case uintptr:
			return true
		default:
			return false
		}
	}(any(value)) || func(hx_value any) bool {
		switch hx_value.(type) {
		case int:
			return true
		case int8:
			return true
		case int16:
			return true
		case int32:
			return true
		case int64:
			return true
		case uint:
			return true
		case uint8:
			return true
		case uint16:
			return true
		case uint32:
			return true
		case uint64:
			return true
		case uintptr:
			return true
		case float32:
			return true
		case float64:
			return true
		default:
			return false
		}
	}(any(value)) {
		return func(hx_value_258 any) float64 {
			if hx_value_258 == nil {
				var hx_zero_259 float64
				return hx_zero_259
			}
			return hx_value_258.(float64)
		}(value)
	}
	if func(hx_value any) bool {
		switch hx_value.(type) {
		case *string:
			return true
		case string:
			return true
		default:
			return false
		}
	}(any(value)) {
		return haxe__Template_parseFloatLiteral(hxrt.StdString(func(hx_value_260 any) *string {
			if hx_value_260 == nil {
				var hx_zero_261 *string
				return hx_zero_261
			}
			return hx_value_260.(*string)
		}(value)))
	}
	hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Expected numeric expression value, got "), hxrt.StdString(value)))
	var hx_throw_zero_262 float64
	return hx_throw_zero_262
}

type haxe___Template__TemplateExpr struct {
	tag    int
	params []any
}

func haxe___Template__TemplateExpr_OpVar(v *string) *haxe___Template__TemplateExpr {
	enumValue := &haxe___Template__TemplateExpr{tag: 0}
	enumValue.params = []any{v}
	return enumValue
}

func haxe___Template__TemplateExpr_OpExpr(expr func() any) *haxe___Template__TemplateExpr {
	enumValue := &haxe___Template__TemplateExpr{tag: 1}
	enumValue.params = []any{expr}
	return enumValue
}

func haxe___Template__TemplateExpr_OpIf(expr func() any, eif *haxe___Template__TemplateExpr, eelse *haxe___Template__TemplateExpr) *haxe___Template__TemplateExpr {
	enumValue := &haxe___Template__TemplateExpr{tag: 2}
	enumValue.params = []any{expr, eif, eelse}
	return enumValue
}

func haxe___Template__TemplateExpr_OpStr(str *string) *haxe___Template__TemplateExpr {
	enumValue := &haxe___Template__TemplateExpr{tag: 3}
	enumValue.params = []any{str}
	return enumValue
}

func haxe___Template__TemplateExpr_OpBlock(items *hxrt.Array) *haxe___Template__TemplateExpr {
	enumValue := &haxe___Template__TemplateExpr{tag: 4}
	enumValue.params = []any{items}
	return enumValue
}

func haxe___Template__TemplateExpr_OpForeach(expr func() any, loop *haxe___Template__TemplateExpr) *haxe___Template__TemplateExpr {
	enumValue := &haxe___Template__TemplateExpr{tag: 5}
	enumValue.params = []any{expr, loop}
	return enumValue
}

func haxe___Template__TemplateExpr_OpMacro(name *string, params *hxrt.Array) *haxe___Template__TemplateExpr {
	enumValue := &haxe___Template__TemplateExpr{tag: 6}
	enumValue.params = []any{name, params}
	return enumValue
}
