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
		token := func(hx_value_5 any) map[string]any {
			if hx_value_5 == nil {
				var hx_zero_6 map[string]any
				return hx_zero_6
			}
			return hx_value_5.(map[string]any)
		}(cursor.tokens.Get(cursor.index))
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Unexpected '"), hxrt.StdString(func(hx_obj_7 map[string]any) bool {
			hx_field_8 := hx_obj_7["s"]
			if hx_field_8 == nil {
				var hx_zero_9 bool
				return hx_zero_9
			}
			return hx_field_8.(bool)
		}(token))), hxrt.StringFromLiteral("'")))
	}
	return self
}

func (self *haxe__Template) execute(context any, macros any) *string {
	var hx_if_11 any
	if hxrt.AnyEqualsNull(macros) {
		hx_obj_10 := map[string]any{}
		hx_if_11 = hx_obj_10
	} else {
		hx_if_11 = macros
	}
	self.macros = hx_if_11
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
		if func(hx_obj_17 map[string]any) int {
			hx_field_18 := hx_obj_17["pos"]
			if hx_field_18 == nil {
				var hx_zero_19 int
				return hx_zero_19
			}
			return hx_field_18.(int)
		}(p) > 0 {
			hx_obj_13 := map[string]any{}
			hx_obj_13["p"] = hxrt.StringSubstrStringPtr(data, 0, func(hx_obj_14 map[string]any) int {
				hx_field_15 := hx_obj_14["pos"]
				if hx_field_15 == nil {
					var hx_zero_16 int
					return hx_zero_16
				}
				return hx_field_15.(int)
			}(p), true)
			hx_obj_13["s"] = true
			hx_obj_13["l"] = nil
			tokens.Push(hx_obj_13)
		}
		if hxrt.StringCharCodeAtAnyStringPtr(data, func(hx_obj_28 map[string]any) int {
			hx_field_29 := hx_obj_28["pos"]
			if hx_field_29 == nil {
				var hx_zero_30 int
				return hx_zero_30
			}
			return hx_field_29.(int)
		}(p)) == 58 {
			hx_obj_21 := map[string]any{}
			hx_obj_21["p"] = hxrt.StringSubstrStringPtr(data, int(int32((hxrt.Int32Wrap(func(hx_obj_22 map[string]any) int {
				hx_field_23 := hx_obj_22["pos"]
				if hx_field_23 == nil {
					var hx_zero_24 int
					return hx_zero_24
				}
				return hx_field_23.(int)
			}(p)) + hxrt.Int32Wrap(2)))), int(int32((hxrt.Int32Wrap(func(hx_obj_25 map[string]any) int {
				hx_field_26 := hx_obj_25["len"]
				if hx_field_26 == nil {
					var hx_zero_27 int
					return hx_zero_27
				}
				return hx_field_26.(int)
			}(p)) - hxrt.Int32Wrap(4)))), true)
			hx_obj_21["s"] = false
			hx_obj_21["l"] = nil
			tokens.Push(hx_obj_21)
			data = haxe__Template_splitter.__hx_this.matchedRight()
			continue
		}
		parp := int(int32((hxrt.Int32Wrap(func(hx_obj_31 map[string]any) int {
			hx_field_32 := hx_obj_31["pos"]
			if hx_field_32 == nil {
				var hx_zero_33 int
				return hx_zero_33
			}
			return hx_field_32.(int)
		}(p)) + hxrt.Int32Wrap(func(hx_obj_34 map[string]any) int {
			hx_field_35 := hx_obj_34["len"]
			if hx_field_35 == nil {
				var hx_zero_36 int
				return hx_zero_36
			}
			return hx_field_35.(int)
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
		hx_obj_40 := map[string]any{}
		hx_obj_40["p"] = haxe__Template_splitter.__hx_this.matched(2)
		hx_obj_40["s"] = false
		hx_obj_40["l"] = params
		tokens.Push(hx_obj_40)
		data = hxrt.StringSubstrStringPtr(data, parp, int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(data)) - hxrt.Int32Wrap(parp)))), true)
	}
	if hxrt.StringLengthStringPtr(data) > 0 {
		hx_obj_42 := map[string]any{}
		hx_obj_42["p"] = data
		hx_obj_42["s"] = true
		hx_obj_42["l"] = nil
		tokens.Push(hx_obj_42)
	}
	return tokens
}

func (self *haxe__Template) parseBlock(cursor *haxe___Template__TokenCursor) *haxe___Template__TemplateExpr {
	items := hxrt.NewArray()
	for cursor.index < cursor.tokens.Len() {
		t := func(hx_value_43 any) map[string]any {
			if hx_value_43 == nil {
				var hx_zero_44 map[string]any
				return hx_zero_44
			}
			return hx_value_43.(map[string]any)
		}(cursor.tokens.Get(cursor.index))
		if !func(hx_obj_45 map[string]any) bool {
			hx_field_46 := hx_obj_45["s"]
			if hx_field_46 == nil {
				var hx_zero_47 bool
				return hx_zero_47
			}
			return hx_field_46.(bool)
		}(t) && ((hxrt.StringEqualStringPtr(func(hx_obj_48 map[string]any) *string {
			hx_field_49 := hx_obj_48["p"]
			if hx_field_49 == nil {
				var hx_zero_50 *string
				return hx_zero_50
			}
			return hx_field_49.(*string)
		}(t), hxrt.StringFromLiteral("end")) || hxrt.StringEqualStringPtr(func(hx_obj_51 map[string]any) *string {
			hx_field_52 := hx_obj_51["p"]
			if hx_field_52 == nil {
				var hx_zero_53 *string
				return hx_zero_53
			}
			return hx_field_52.(*string)
		}(t), hxrt.StringFromLiteral("else"))) || hxrt.StringEqualStringPtr(hxrt.StringSubstrStringPtr(func(hx_obj_54 map[string]any) *string {
			hx_field_55 := hx_obj_54["p"]
			if hx_field_55 == nil {
				var hx_zero_56 *string
				return hx_zero_56
			}
			return hx_field_55.(*string)
		}(t), 0, 7, true), hxrt.StringFromLiteral("elseif "))) {
			break
		}
		items.Push(self.__hx_this.parse(cursor))
	}
	if items.Len() == 1 {
		return func(hx_value_58 any) *haxe___Template__TemplateExpr {
			if hx_value_58 == nil {
				var hx_zero_59 *haxe___Template__TemplateExpr
				return hx_zero_59
			}
			return hx_value_58.(*haxe___Template__TemplateExpr)
		}(items.Get(0))
	}
	return haxe___Template__TemplateExpr_OpBlock(items)
}

func (self *haxe__Template) parse(cursor *haxe___Template__TokenCursor) *haxe___Template__TemplateExpr {
	var hx_if_63 map[string]any
	if cursor.index >= cursor.tokens.Len() {
		hx_if_63 = nil
	} else {
		hx_post_60 := cursor.index
		cursor.index = int(int32((cursor.index + 1)))
		hx_if_63 = func(hx_value_61 any) map[string]any {
			if hx_value_61 == nil {
				var hx_zero_62 map[string]any
				return hx_zero_62
			}
			return hx_value_61.(map[string]any)
		}(cursor.tokens.Get(hx_post_60))
	}
	t := hx_if_63
	if t == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Unexpected <eof>"))
	}
	p := func(hx_obj_64 map[string]any) *string {
		hx_field_65 := hx_obj_64["p"]
		if hx_field_65 == nil {
			var hx_zero_66 *string
			return hx_zero_66
		}
		return hx_field_65.(*string)
	}(t)
	if func(hx_obj_67 map[string]any) bool {
		hx_field_68 := hx_obj_67["s"]
		if hx_field_68 == nil {
			var hx_zero_69 bool
			return hx_zero_69
		}
		return hx_field_68.(bool)
	}(t) {
		return haxe___Template__TemplateExpr_OpStr(p)
	}
	if func(hx_obj_76 map[string]any) *hxrt.Array {
		hx_field_77 := hx_obj_76["l"]
		if hx_field_77 == nil {
			var hx_zero_78 *hxrt.Array
			return hx_zero_78
		}
		return hx_field_77.(*hxrt.Array)
	}(t) != nil {
		parsedParams := hxrt.NewArray()
		_g := 0
		_g1 := func(hx_obj_70 map[string]any) *hxrt.Array {
			hx_field_71 := hx_obj_70["l"]
			if hx_field_71 == nil {
				var hx_zero_72 *hxrt.Array
				return hx_zero_72
			}
			return hx_field_71.(*hxrt.Array)
		}(t)
		for _g < _g1.Len() {
			param := func(hx_value_73 any) *string {
				if hx_value_73 == nil {
					var hx_zero_74 *string
					return hx_zero_74
				}
				return hx_value_73.(*string)
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
		var hx_if_81 map[string]any
		if cursor.index < cursor.tokens.Len() {
			hx_if_81 = func(hx_value_79 any) map[string]any {
				if hx_value_79 == nil {
					var hx_zero_80 map[string]any
					return hx_zero_80
				}
				return hx_value_79.(map[string]any)
			}(cursor.tokens.Get(cursor.index))
		} else {
			hx_if_81 = nil
		}
		nextToken := hx_if_81
		if nextToken == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Unclosed 'if'"))
		}
		var eelse *haxe___Template__TemplateExpr = nil
		if hxrt.StringEqualStringPtr(func(hx_obj_100 map[string]any) *string {
			hx_field_101 := hx_obj_100["p"]
			if hx_field_101 == nil {
				var hx_zero_102 *string
				return hx_zero_102
			}
			return hx_field_101.(*string)
		}(nextToken), hxrt.StringFromLiteral("end")) {
			if cursor.index >= cursor.tokens.Len() {
			} else {
				cursor.tokens.Get(func() int {
					hx_post_82 := cursor.index
					cursor.index = int(int32((cursor.index + 1)))
					return hx_post_82
				}())
			}
		} else {
			if hxrt.StringEqualStringPtr(func(hx_obj_97 map[string]any) *string {
				hx_field_98 := hx_obj_97["p"]
				if hx_field_98 == nil {
					var hx_zero_99 *string
					return hx_zero_99
				}
				return hx_field_98.(*string)
			}(nextToken), hxrt.StringFromLiteral("else")) {
				if cursor.index >= cursor.tokens.Len() {
				} else {
					cursor.tokens.Get(func() int {
						hx_post_83 := cursor.index
						cursor.index = int(int32((cursor.index + 1)))
						return hx_post_83
					}())
				}
				eelse = self.__hx_this.parseBlock(cursor)
				var hx_if_87 map[string]any
				if cursor.index >= cursor.tokens.Len() {
					hx_if_87 = nil
				} else {
					hx_post_84 := cursor.index
					cursor.index = int(int32((cursor.index + 1)))
					hx_if_87 = func(hx_value_85 any) map[string]any {
						if hx_value_85 == nil {
							var hx_zero_86 map[string]any
							return hx_zero_86
						}
						return hx_value_85.(map[string]any)
					}(cursor.tokens.Get(hx_post_84))
				}
				endToken := hx_if_87
				if (endToken == nil) || !hxrt.StringEqualStringPtr(func(hx_obj_88 map[string]any) *string {
					hx_field_89 := hx_obj_88["p"]
					if hx_field_89 == nil {
						var hx_zero_90 *string
						return hx_zero_90
					}
					return hx_field_89.(*string)
				}(endToken), hxrt.StringFromLiteral("end")) {
					hxrt.Throw(hxrt.StringFromLiteral("Unclosed 'else'"))
				}
			} else {
				nextToken["p"] = hxrt.StringSubstrStringPtr(func(hx_obj_91 map[string]any) *string {
					hx_field_92 := hx_obj_91["p"]
					if hx_field_92 == nil {
						var hx_zero_93 *string
						return hx_zero_93
					}
					return hx_field_92.(*string)
				}(nextToken), 4, int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(func(hx_obj_94 map[string]any) *string {
					hx_field_95 := hx_obj_94["p"]
					if hx_field_95 == nil {
						var hx_zero_96 *string
						return hx_zero_96
					}
					return hx_field_95.(*string)
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
		var hx_if_106 map[string]any
		if cursor.index >= cursor.tokens.Len() {
			hx_if_106 = nil
		} else {
			hx_post_103 := cursor.index
			cursor.index = int(int32((cursor.index + 1)))
			hx_if_106 = func(hx_value_104 any) map[string]any {
				if hx_value_104 == nil {
					var hx_zero_105 map[string]any
					return hx_zero_105
				}
				return hx_value_104.(map[string]any)
			}(cursor.tokens.Get(hx_post_103))
		}
		endToken_1 := hx_if_106
		if (endToken_1 == nil) || !hxrt.StringEqualStringPtr(func(hx_obj_107 map[string]any) *string {
			hx_field_108 := hx_obj_107["p"]
			if hx_field_108 == nil {
				var hx_zero_109 *string
				return hx_zero_109
			}
			return hx_field_108.(*string)
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
		if func(hx_obj_115 map[string]any) int {
			hx_field_116 := hx_obj_115["pos"]
			if hx_field_116 == nil {
				var hx_zero_117 int
				return hx_zero_117
			}
			return hx_field_116.(int)
		}(p) != 0 {
			hx_obj_111 := map[string]any{}
			hx_obj_111["p"] = hxrt.StringSubstrStringPtr(data, 0, func(hx_obj_112 map[string]any) int {
				hx_field_113 := hx_obj_112["pos"]
				if hx_field_113 == nil {
					var hx_zero_114 int
					return hx_zero_114
				}
				return hx_field_113.(int)
			}(p), true)
			hx_obj_111["s"] = true
			tokens.Push(hx_obj_111)
		}
		token := haxe__Template_expr_splitter.__hx_this.matched(0)
		hx_obj_119 := map[string]any{}
		hx_obj_119["p"] = token
		hx_obj_119["s"] = StringTools_contains(token, hxrt.StringFromLiteral("\""))
		tokens.Push(hx_obj_119)
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
				hx_obj_121 := map[string]any{}
				hx_obj_121["p"] = hxrt.StringSubstrStringPtr(data, i, 0, false)
				hx_obj_121["s"] = true
				tokens.Push(hx_obj_121)
				break
			}
		}
	}
	cursor := New_haxe___Template__ExprCursor(tokens)
	var built func() any
	hxrt.TryCatch(func() {
		built = self.__hx_this.makeExpr(cursor)
		if cursor.index < cursor.tokens.Len() {
			hxrt.Throw(func(hx_obj_126 map[string]any) *string {
				hx_field_127 := hx_obj_126["p"]
				if hx_field_127 == nil {
					var hx_zero_128 *string
					return hx_zero_128
				}
				return hx_field_127.(*string)
			}(func(hx_value_124 any) map[string]any {
				if hx_value_124 == nil {
					var hx_zero_125 map[string]any
					return hx_zero_125
				}
				return hx_value_124.(map[string]any)
			}(cursor.tokens.Get(cursor.index))))
		}
	}, func(hx_caught_122 any) {
		switch hx_typed_123 := hx_caught_122.(type) {
		case *string:
			s := hx_typed_123
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Unexpected '"), s), hxrt.StringFromLiteral("' in ")), expr))
		default:
			hxrt.Throw(hx_caught_122)
		}
	})
	me := self
	_ = me
	wrapped := func() any {
		hx_try_return_129 := false
		var hx_try_value_130 any
		hxrt.TryCatch(func() {
			hx_try_value_130 = built()
			hx_try_return_129 = true
			return
		}, func(hx_caught_131 any) {
			exc := hx_caught_131
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Error : "), hxrt.StdString(exc)), hxrt.StringFromLiteral(" in ")), expr))
		})
		if hx_try_return_129 {
			return hx_try_value_130
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
	var hx_if_135 map[string]any
	if cursor.index < cursor.tokens.Len() {
		hx_if_135 = func(hx_value_133 any) map[string]any {
			if hx_value_133 == nil {
				var hx_zero_134 map[string]any
				return hx_zero_134
			}
			return hx_value_133.(map[string]any)
		}(cursor.tokens.Get(cursor.index))
	} else {
		hx_if_135 = nil
	}
	token := hx_if_135
	if (token == nil) || !hxrt.StringEqualStringPtr(func(hx_obj_136 map[string]any) *string {
		hx_field_137 := hx_obj_136["p"]
		if hx_field_137 == nil {
			var hx_zero_138 *string
			return hx_zero_138
		}
		return hx_field_137.(*string)
	}(token), hxrt.StringFromLiteral(".")) {
		return e
	}
	if cursor.index >= cursor.tokens.Len() {
	} else {
		cursor.tokens.Get(func() int {
			hx_post_139 := cursor.index
			cursor.index = int(int32((cursor.index + 1)))
			return hx_post_139
		}())
	}
	var hx_if_143 map[string]any
	if cursor.index >= cursor.tokens.Len() {
		hx_if_143 = nil
	} else {
		hx_post_140 := cursor.index
		cursor.index = int(int32((cursor.index + 1)))
		hx_if_143 = func(hx_value_141 any) map[string]any {
			if hx_value_141 == nil {
				var hx_zero_142 map[string]any
				return hx_zero_142
			}
			return hx_value_141.(map[string]any)
		}(cursor.tokens.Get(hx_post_140))
	}
	field := hx_if_143
	if (field == nil) || !func(hx_obj_148 map[string]any) bool {
		hx_field_149 := hx_obj_148["s"]
		if hx_field_149 == nil {
			var hx_zero_150 bool
			return hx_zero_150
		}
		return hx_field_149.(bool)
	}(field) {
		var hx_if_147 *string
		if field == nil {
			hx_if_147 = hxrt.StringFromLiteral("<eof>")
		} else {
			hx_if_147 = func(hx_obj_144 map[string]any) *string {
				hx_field_145 := hx_obj_144["p"]
				if hx_field_145 == nil {
					var hx_zero_146 *string
					return hx_zero_146
				}
				return hx_field_145.(*string)
			}(field)
		}
		hxrt.Throw(hx_if_147)
	}
	name := haxe__Template_trimExprToken(func(hx_obj_151 map[string]any) *string {
		hx_field_152 := hx_obj_151["p"]
		if hx_field_152 == nil {
			var hx_zero_153 *string
			return hx_zero_153
		}
		return hx_field_152.(*string)
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
		if !haxe__Template_isSpaceOnly(func(hx_obj_156 map[string]any) *string {
			hx_field_157 := hx_obj_156["p"]
			if hx_field_157 == nil {
				var hx_zero_158 *string
				return hx_zero_158
			}
			return hx_field_157.(*string)
		}(func(hx_value_154 any) map[string]any {
			if hx_value_154 == nil {
				var hx_zero_155 map[string]any
				return hx_zero_155
			}
			return hx_value_154.(map[string]any)
		}(cursor.tokens.Get(cursor.index)))) {
			return
		}
		cursor.index = int(int32((cursor.index + 1)))
	}
}

func (self *haxe__Template) makeExpr2(cursor *haxe___Template__ExprCursor) func() any {
	self.__hx_this.skipSpaces(cursor)
	var hx_if_162 map[string]any
	if cursor.index >= cursor.tokens.Len() {
		hx_if_162 = nil
	} else {
		hx_post_159 := cursor.index
		cursor.index = int(int32((cursor.index + 1)))
		hx_if_162 = func(hx_value_160 any) map[string]any {
			if hx_value_160 == nil {
				var hx_zero_161 map[string]any
				return hx_zero_161
			}
			return hx_value_160.(map[string]any)
		}(cursor.tokens.Get(hx_post_159))
	}
	token := hx_if_162
	self.__hx_this.skipSpaces(cursor)
	if token == nil {
		hxrt.Throw(hxrt.StringFromLiteral("<eof>"))
	}
	if func(hx_obj_166 map[string]any) bool {
		hx_field_167 := hx_obj_166["s"]
		if hx_field_167 == nil {
			var hx_zero_168 bool
			return hx_zero_168
		}
		return hx_field_167.(bool)
	}(token) {
		return self.__hx_this.makeConst(func(hx_obj_163 map[string]any) *string {
			hx_field_164 := hx_obj_163["p"]
			if hx_field_164 == nil {
				var hx_zero_165 *string
				return hx_zero_165
			}
			return hx_field_164.(*string)
		}(token))
	}
	_g := func(hx_obj_169 map[string]any) *string {
		hx_field_170 := hx_obj_169["p"]
		if hx_field_170 == nil {
			var hx_zero_171 *string
			return hx_zero_171
		}
		return hx_field_170.(*string)
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
		var hx_if_175 map[string]any
		if cursor.index >= cursor.tokens.Len() {
			hx_if_175 = nil
		} else {
			hx_post_172 := cursor.index
			cursor.index = int(int32((cursor.index + 1)))
			hx_if_175 = func(hx_value_173 any) map[string]any {
				if hx_value_173 == nil {
					var hx_zero_174 map[string]any
					return hx_zero_174
				}
				return hx_value_173.(map[string]any)
			}(cursor.tokens.Get(hx_post_172))
		}
		op := hx_if_175
		if (op == nil) || func(hx_obj_180 map[string]any) bool {
			hx_field_181 := hx_obj_180["s"]
			if hx_field_181 == nil {
				var hx_zero_182 bool
				return hx_zero_182
			}
			return hx_field_181.(bool)
		}(op) {
			var hx_if_179 *string
			if op == nil {
				hx_if_179 = hxrt.StringFromLiteral("<eof>")
			} else {
				hx_if_179 = func(hx_obj_176 map[string]any) *string {
					hx_field_177 := hx_obj_176["p"]
					if hx_field_177 == nil {
						var hx_zero_178 *string
						return hx_zero_178
					}
					return hx_field_177.(*string)
				}(op)
			}
			hxrt.Throw(hx_if_179)
		}
		if hxrt.StringEqualStringPtr(func(hx_obj_183 map[string]any) *string {
			hx_field_184 := hx_obj_183["p"]
			if hx_field_184 == nil {
				var hx_zero_185 *string
				return hx_zero_185
			}
			return hx_field_184.(*string)
		}(op), hxrt.StringFromLiteral(")")) {
			return e1
		}
		self.__hx_this.skipSpaces(cursor)
		e2 := self.__hx_this.makeExpr(cursor)
		self.__hx_this.skipSpaces(cursor)
		var hx_if_189 map[string]any
		if cursor.index >= cursor.tokens.Len() {
			hx_if_189 = nil
		} else {
			hx_post_186 := cursor.index
			cursor.index = int(int32((cursor.index + 1)))
			hx_if_189 = func(hx_value_187 any) map[string]any {
				if hx_value_187 == nil {
					var hx_zero_188 map[string]any
					return hx_zero_188
				}
				return hx_value_187.(map[string]any)
			}(cursor.tokens.Get(hx_post_186))
		}
		close := hx_if_189
		self.__hx_this.skipSpaces(cursor)
		if (close == nil) || !hxrt.StringEqualStringPtr(func(hx_obj_194 map[string]any) *string {
			hx_field_195 := hx_obj_194["p"]
			if hx_field_195 == nil {
				var hx_zero_196 *string
				return hx_zero_196
			}
			return hx_field_195.(*string)
		}(close), hxrt.StringFromLiteral(")")) {
			var hx_if_193 *string
			if close == nil {
				hx_if_193 = hxrt.StringFromLiteral("<eof>")
			} else {
				hx_if_193 = func(hx_obj_190 map[string]any) *string {
					hx_field_191 := hx_obj_190["p"]
					if hx_field_191 == nil {
						var hx_zero_192 *string
						return hx_zero_192
					}
					return hx_field_191.(*string)
				}(close)
			}
			hxrt.Throw(hx_if_193)
		}
		_g_1 := func(hx_obj_197 map[string]any) *string {
			hx_field_198 := hx_obj_197["p"]
			if hx_field_198 == nil {
				var hx_zero_199 *string
				return hx_zero_199
			}
			return hx_field_198.(*string)
		}(op)
		var hx_switch_200 func() any
		switch *hxrt.StdString(_g_1) {
		case *hxrt.StdString(hxrt.StringFromLiteral("!=")):
			hx_switch_200 = func() any {
				return !hxrt.HaxeEqual(e1(), e2())
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("&&")):
			hx_switch_200 = func() any {
				return (haxe__Template_valueAsBool(e1()) && haxe__Template_valueAsBool(e2()))
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("*")):
			hx_switch_200 = func() any {
				var left any = e1()
				var right any = e2()
				return (haxe__Template_valueAsFloat(left) * haxe__Template_valueAsFloat(right))
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("+")):
			hx_switch_200 = func() any {
				return haxe__Template_addValues(e1(), e2())
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("-")):
			hx_switch_200 = func() any {
				var left any = e1()
				var right any = e2()
				return (haxe__Template_valueAsFloat(left) - haxe__Template_valueAsFloat(right))
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("/")):
			hx_switch_200 = func() any {
				var left any = e1()
				var right any = e2()
				return (haxe__Template_valueAsFloat(left) / haxe__Template_valueAsFloat(right))
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("<")):
			hx_switch_200 = func() any {
				return (haxe__Template_compareValues(e1(), e2()) < 0)
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("<=")):
			hx_switch_200 = func() any {
				return (haxe__Template_compareValues(e1(), e2()) <= 0)
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("==")):
			hx_switch_200 = func() any {
				return hxrt.HaxeEqual(e1(), e2())
			}
		case *hxrt.StdString(hxrt.StringFromLiteral(">")):
			hx_switch_200 = func() any {
				return (haxe__Template_compareValues(e1(), e2()) > 0)
			}
		case *hxrt.StdString(hxrt.StringFromLiteral(">=")):
			hx_switch_200 = func() any {
				return (haxe__Template_compareValues(e1(), e2()) >= 0)
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("||")):
			hx_switch_200 = func() any {
				return (haxe__Template_valueAsBool(e1()) || haxe__Template_valueAsBool(e2()))
			}
		default:
			hx_switch_200 = func() func() any {
				hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Unknown operation "), func(hx_obj_201 map[string]any) *string {
					hx_field_202 := hx_obj_201["p"]
					if hx_field_202 == nil {
						var hx_zero_203 *string
						return hx_zero_203
					}
					return hx_field_202.(*string)
				}(op)))
				var hx_throw_zero_204 func() any
				return hx_throw_zero_204
			}()
		}
		return hx_switch_200
	case *hxrt.StdString(hxrt.StringFromLiteral("-")):
		inner_1 := self.__hx_this.makeExpr(cursor)
		return func() any {
			return -haxe__Template_valueAsFloat(inner_1())
		}
	default:
		hxrt.Throw(func(hx_obj_205 map[string]any) *string {
			hx_field_206 := hx_obj_205["p"]
			if hx_field_206 == nil {
				var hx_zero_207 *string
				return hx_zero_207
			}
			return hx_field_206.(*string)
		}(token))
		var hx_throw_zero_208 func() any
		return hx_throw_zero_208
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
			item := func(hx_value_209 any) *haxe___Template__TemplateExpr {
				if hx_value_209 == nil {
					var hx_zero_210 *haxe___Template__TemplateExpr
					return hx_zero_210
				}
				return hx_value_209.(*haxe___Template__TemplateExpr)
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
			hx_arr_211 := self.stack
			hx_arr_211.Push(self.context)
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
		}, func(hx_caught_212 any) {
			hx_tmp := hx_caught_212
			_ = hx_tmp
			hxrt.TryCatch(func() {
				if hxrt.AnyEqualsNull(value_1) || !Reflect_hasField(value_1, hxrt.StringFromLiteral("hasNext")) {
					hxrt.Throw(nil)
				}
				iterator = value_1
			}, func(hx_caught_214 any) {
				hx_tmp_1 := hx_caught_214
				_ = hx_tmp_1
				hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Cannot iter on "), hxrt.StdString(value_1)))
			})
		})
		hx_arr_216 := self.stack
		hx_arr_216.Push(self.context)
		iterable := func(hx_value_217 any) map[string]any {
			if hx_value_217 == nil {
				var hx_zero_218 map[string]any
				return hx_zero_218
			}
			return hx_value_217.(map[string]any)
		}(iterator)
		ctx_1 := iterable
		for func(hx_obj_219 map[string]any) func() bool {
			hx_field_220 := hx_obj_219["hasNext"]
			if hx_field_220 == nil {
				var hx_zero_221 func() bool
				return hx_zero_221
			}
			return hx_field_220.(func() bool)
		}(ctx_1)() {
			var ctx_2 any = func(hx_obj_222 map[string]any) func() any {
				hx_field_223 := hx_obj_222["next"]
				if hx_field_223 == nil {
					var hx_zero_224 func() any
					return hx_zero_224
				}
				return hx_field_223.(func() any)
			}(ctx_1)()
			self.context = ctx_2
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
			param := func(hx_value_226 any) *haxe___Template__TemplateExpr {
				if hx_value_226 == nil {
					var hx_zero_227 *haxe___Template__TemplateExpr
					return hx_zero_227
				}
				return hx_value_226.(*haxe___Template__TemplateExpr)
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
		}, func(hx_caught_230 any) {
			err := hx_caught_230
			var hx_try_232 *string
			hxrt.TryCatch(func() {
				hx_try_232 = haxe__Template_joinDynamicArgs(callArgs)
			}, func(hx_caught_233 any) {
				hx_tmp_2 := hx_caught_233
				_ = hx_tmp_2
				hx_try_232 = hxrt.StringFromLiteral("???")
			})
			argsText := hx_try_232
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
		hx_post_235 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_235
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
		var hx_if_238 int
		if leftFloat < rightFloat {
			hx_if_238 = -1
		} else {
			var hx_if_237 int
			if leftFloat > rightFloat {
				hx_if_237 = 1
			} else {
				hx_if_237 = 0
			}
			hx_if_238 = hx_if_237
		}
		return hx_if_238
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
	hx_obj_239 := map[string]any{}
	return hx_obj_239
}())

func haxe__Template_isSpaceOnly(value *string) bool {
	var _g_s *string
	var _g_offset int
	_g_offset = 0
	_g_s = value
	for _g_offset < hxrt.StringLengthStringPtr(_g_s) {
		code := hxrt.StringCharCodeAtStringPtr(_g_s, func() int {
			hx_post_240 := _g_offset
			_g_offset = int(int32((_g_offset + 1)))
			return hx_post_240
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
		hx_post_241 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_241
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
				hx_post_242 := _g_offset
				_g_offset = int(int32((_g_offset + 1)))
				return hx_post_242
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
				var hx_if_243 float64
				if exponentSign < 0 {
					hx_if_243 = (result / 10.0)
				} else {
					hx_if_243 = (result * 10.0)
				}
				result = hx_if_243
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
	var hx_if_246 map[string]any
	if cursor.index < cursor.tokens.Len() {
		hx_if_246 = func(hx_value_244 any) map[string]any {
			if hx_value_244 == nil {
				var hx_zero_245 map[string]any
				return hx_zero_245
			}
			return hx_value_244.(map[string]any)
		}(cursor.tokens.Get(cursor.index))
	} else {
		hx_if_246 = nil
	}
	return hx_if_246
}

func haxe__Template_peekToken(cursor *haxe___Template__TokenCursor) map[string]any {
	var hx_if_249 map[string]any
	if cursor.index < cursor.tokens.Len() {
		hx_if_249 = func(hx_value_247 any) map[string]any {
			if hx_value_247 == nil {
				var hx_zero_248 map[string]any
				return hx_zero_248
			}
			return hx_value_247.(map[string]any)
		}(cursor.tokens.Get(cursor.index))
	} else {
		hx_if_249 = nil
	}
	return hx_if_249
}

func haxe__Template_popExprToken(cursor *haxe___Template__ExprCursor) map[string]any {
	if cursor.index >= cursor.tokens.Len() {
		return nil
	}
	hx_post_250 := cursor.index
	cursor.index = int(int32((cursor.index + 1)))
	return func(hx_value_251 any) map[string]any {
		if hx_value_251 == nil {
			var hx_zero_252 map[string]any
			return hx_zero_252
		}
		return hx_value_251.(map[string]any)
	}(cursor.tokens.Get(hx_post_250))
}

func haxe__Template_popToken(cursor *haxe___Template__TokenCursor) map[string]any {
	if cursor.index >= cursor.tokens.Len() {
		return nil
	}
	hx_post_253 := cursor.index
	cursor.index = int(int32((cursor.index + 1)))
	return func(hx_value_254 any) map[string]any {
		if hx_value_254 == nil {
			var hx_zero_255 map[string]any
			return hx_zero_255
		}
		return hx_value_254.(map[string]any)
	}(cursor.tokens.Get(hx_post_253))
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
		return func(hx_value_256 any) float64 {
			if hx_value_256 == nil {
				var hx_zero_257 float64
				return hx_zero_257
			}
			return hx_value_256.(float64)
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
		return haxe__Template_parseFloatLiteral(hxrt.StdString(func(hx_value_258 any) *string {
			if hx_value_258 == nil {
				var hx_zero_259 *string
				return hx_zero_259
			}
			return hx_value_258.(*string)
		}(value)))
	}
	hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Expected numeric expression value, got "), hxrt.StdString(value)))
	var hx_throw_zero_260 float64
	return hx_throw_zero_260
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
