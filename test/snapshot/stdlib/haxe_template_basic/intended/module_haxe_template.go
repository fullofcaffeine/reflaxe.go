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
		token := func(hx_value_1 any) map[string]any {
			if hx_value_1 == nil {
				var hx_zero_2 map[string]any
				return hx_zero_2
			}
			return hx_value_1.(map[string]any)
		}(cursor.tokens.Get(cursor.index))
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Unexpected '"), hxrt.StdString(func(hx_obj_3 map[string]any) bool {
			hx_field_4 := hx_obj_3["s"]
			if hx_field_4 == nil {
				var hx_zero_5 bool
				return hx_zero_5
			}
			return hx_field_4.(bool)
		}(token))), hxrt.StringFromLiteral("'")))
	}
	return self
}

func (self *haxe__Template) execute(context any, macros any) *string {
	var hx_if_7 any
	if hxrt.AnyEqualsNull(macros) {
		hx_obj_6 := map[string]any{}
		hx_if_7 = hx_obj_6
	} else {
		hx_if_7 = macros
	}
	self.macros = hx_if_7
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
		if func(hx_obj_13 map[string]any) int {
			hx_field_14 := hx_obj_13["pos"]
			if hx_field_14 == nil {
				var hx_zero_15 int
				return hx_zero_15
			}
			return hx_field_14.(int)
		}(p) > 0 {
			hx_obj_9 := map[string]any{}
			hx_obj_9["p"] = hxrt.StringSubstrStringPtr(data, 0, func(hx_obj_10 map[string]any) int {
				hx_field_11 := hx_obj_10["pos"]
				if hx_field_11 == nil {
					var hx_zero_12 int
					return hx_zero_12
				}
				return hx_field_11.(int)
			}(p), true)
			hx_obj_9["s"] = true
			hx_obj_9["l"] = nil
			tokens.Push(hx_obj_9)
		}
		if hxrt.StringCharCodeAtAnyStringPtr(data, func(hx_obj_24 map[string]any) int {
			hx_field_25 := hx_obj_24["pos"]
			if hx_field_25 == nil {
				var hx_zero_26 int
				return hx_zero_26
			}
			return hx_field_25.(int)
		}(p)) == 58 {
			hx_obj_17 := map[string]any{}
			hx_obj_17["p"] = hxrt.StringSubstrStringPtr(data, int(int32((hxrt.Int32Wrap(func(hx_obj_18 map[string]any) int {
				hx_field_19 := hx_obj_18["pos"]
				if hx_field_19 == nil {
					var hx_zero_20 int
					return hx_zero_20
				}
				return hx_field_19.(int)
			}(p)) + hxrt.Int32Wrap(2)))), int(int32((hxrt.Int32Wrap(func(hx_obj_21 map[string]any) int {
				hx_field_22 := hx_obj_21["len"]
				if hx_field_22 == nil {
					var hx_zero_23 int
					return hx_zero_23
				}
				return hx_field_22.(int)
			}(p)) - hxrt.Int32Wrap(4)))), true)
			hx_obj_17["s"] = false
			hx_obj_17["l"] = nil
			tokens.Push(hx_obj_17)
			data = haxe__Template_splitter.__hx_this.matchedRight()
			continue
		}
		parp := int(int32((hxrt.Int32Wrap(func(hx_obj_27 map[string]any) int {
			hx_field_28 := hx_obj_27["pos"]
			if hx_field_28 == nil {
				var hx_zero_29 int
				return hx_zero_29
			}
			return hx_field_28.(int)
		}(p)) + hxrt.Int32Wrap(func(hx_obj_30 map[string]any) int {
			hx_field_31 := hx_obj_30["len"]
			if hx_field_31 == nil {
				var hx_zero_32 int
				return hx_zero_32
			}
			return hx_field_31.(int)
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
		hx_obj_36 := map[string]any{}
		hx_obj_36["p"] = haxe__Template_splitter.__hx_this.matched(2)
		hx_obj_36["s"] = false
		hx_obj_36["l"] = params
		tokens.Push(hx_obj_36)
		data = hxrt.StringSubstrStringPtr(data, parp, int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(data)) - hxrt.Int32Wrap(parp)))), true)
	}
	if hxrt.StringLengthStringPtr(data) > 0 {
		hx_obj_38 := map[string]any{}
		hx_obj_38["p"] = data
		hx_obj_38["s"] = true
		hx_obj_38["l"] = nil
		tokens.Push(hx_obj_38)
	}
	return tokens
}

func (self *haxe__Template) parseBlock(cursor *haxe___Template__TokenCursor) *haxe___Template__TemplateExpr {
	items := hxrt.NewArray()
	for cursor.index < cursor.tokens.Len() {
		t := func(hx_value_39 any) map[string]any {
			if hx_value_39 == nil {
				var hx_zero_40 map[string]any
				return hx_zero_40
			}
			return hx_value_39.(map[string]any)
		}(cursor.tokens.Get(cursor.index))
		if !func(hx_obj_41 map[string]any) bool {
			hx_field_42 := hx_obj_41["s"]
			if hx_field_42 == nil {
				var hx_zero_43 bool
				return hx_zero_43
			}
			return hx_field_42.(bool)
		}(t) && ((hxrt.StringEqualStringPtr(func(hx_obj_44 map[string]any) *string {
			hx_field_45 := hx_obj_44["p"]
			if hx_field_45 == nil {
				var hx_zero_46 *string
				return hx_zero_46
			}
			return hx_field_45.(*string)
		}(t), hxrt.StringFromLiteral("end")) || hxrt.StringEqualStringPtr(func(hx_obj_47 map[string]any) *string {
			hx_field_48 := hx_obj_47["p"]
			if hx_field_48 == nil {
				var hx_zero_49 *string
				return hx_zero_49
			}
			return hx_field_48.(*string)
		}(t), hxrt.StringFromLiteral("else"))) || hxrt.StringEqualStringPtr(hxrt.StringSubstrStringPtr(func(hx_obj_50 map[string]any) *string {
			hx_field_51 := hx_obj_50["p"]
			if hx_field_51 == nil {
				var hx_zero_52 *string
				return hx_zero_52
			}
			return hx_field_51.(*string)
		}(t), 0, 7, true), hxrt.StringFromLiteral("elseif "))) {
			break
		}
		items.Push(self.__hx_this.parse(cursor))
	}
	if items.Len() == 1 {
		return func(hx_value_54 any) *haxe___Template__TemplateExpr {
			if hx_value_54 == nil {
				var hx_zero_55 *haxe___Template__TemplateExpr
				return hx_zero_55
			}
			return hx_value_54.(*haxe___Template__TemplateExpr)
		}(items.Get(0))
	}
	return haxe___Template__TemplateExpr_OpBlock(items)
}

func (self *haxe__Template) parse(cursor *haxe___Template__TokenCursor) *haxe___Template__TemplateExpr {
	var hx_if_59 map[string]any
	if cursor.index >= cursor.tokens.Len() {
		hx_if_59 = nil
	} else {
		hx_post_56 := cursor.index
		cursor.index = int(int32((cursor.index + 1)))
		hx_if_59 = func(hx_value_57 any) map[string]any {
			if hx_value_57 == nil {
				var hx_zero_58 map[string]any
				return hx_zero_58
			}
			return hx_value_57.(map[string]any)
		}(cursor.tokens.Get(hx_post_56))
	}
	t := hx_if_59
	if t == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Unexpected <eof>"))
	}
	p := func(hx_obj_60 map[string]any) *string {
		hx_field_61 := hx_obj_60["p"]
		if hx_field_61 == nil {
			var hx_zero_62 *string
			return hx_zero_62
		}
		return hx_field_61.(*string)
	}(t)
	if func(hx_obj_63 map[string]any) bool {
		hx_field_64 := hx_obj_63["s"]
		if hx_field_64 == nil {
			var hx_zero_65 bool
			return hx_zero_65
		}
		return hx_field_64.(bool)
	}(t) {
		return haxe___Template__TemplateExpr_OpStr(p)
	}
	if func(hx_obj_72 map[string]any) *hxrt.Array {
		hx_field_73 := hx_obj_72["l"]
		if hx_field_73 == nil {
			var hx_zero_74 *hxrt.Array
			return hx_zero_74
		}
		return hx_field_73.(*hxrt.Array)
	}(t) != nil {
		parsedParams := hxrt.NewArray()
		_g := 0
		_g1 := func(hx_obj_66 map[string]any) *hxrt.Array {
			hx_field_67 := hx_obj_66["l"]
			if hx_field_67 == nil {
				var hx_zero_68 *hxrt.Array
				return hx_zero_68
			}
			return hx_field_67.(*hxrt.Array)
		}(t)
		for _g < _g1.Len() {
			param := func(hx_value_69 any) *string {
				if hx_value_69 == nil {
					var hx_zero_70 *string
					return hx_zero_70
				}
				return hx_value_69.(*string)
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
		var hx_if_77 map[string]any
		if cursor.index < cursor.tokens.Len() {
			hx_if_77 = func(hx_value_75 any) map[string]any {
				if hx_value_75 == nil {
					var hx_zero_76 map[string]any
					return hx_zero_76
				}
				return hx_value_75.(map[string]any)
			}(cursor.tokens.Get(cursor.index))
		} else {
			hx_if_77 = nil
		}
		nextToken := hx_if_77
		if nextToken == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Unclosed 'if'"))
		}
		var eelse *haxe___Template__TemplateExpr = nil
		if hxrt.StringEqualStringPtr(func(hx_obj_96 map[string]any) *string {
			hx_field_97 := hx_obj_96["p"]
			if hx_field_97 == nil {
				var hx_zero_98 *string
				return hx_zero_98
			}
			return hx_field_97.(*string)
		}(nextToken), hxrt.StringFromLiteral("end")) {
			if cursor.index >= cursor.tokens.Len() {
			} else {
				cursor.tokens.Get(func() int {
					hx_post_78 := cursor.index
					cursor.index = int(int32((cursor.index + 1)))
					return hx_post_78
				}())
			}
		} else {
			if hxrt.StringEqualStringPtr(func(hx_obj_93 map[string]any) *string {
				hx_field_94 := hx_obj_93["p"]
				if hx_field_94 == nil {
					var hx_zero_95 *string
					return hx_zero_95
				}
				return hx_field_94.(*string)
			}(nextToken), hxrt.StringFromLiteral("else")) {
				if cursor.index >= cursor.tokens.Len() {
				} else {
					cursor.tokens.Get(func() int {
						hx_post_79 := cursor.index
						cursor.index = int(int32((cursor.index + 1)))
						return hx_post_79
					}())
				}
				eelse = self.__hx_this.parseBlock(cursor)
				var hx_if_83 map[string]any
				if cursor.index >= cursor.tokens.Len() {
					hx_if_83 = nil
				} else {
					hx_post_80 := cursor.index
					cursor.index = int(int32((cursor.index + 1)))
					hx_if_83 = func(hx_value_81 any) map[string]any {
						if hx_value_81 == nil {
							var hx_zero_82 map[string]any
							return hx_zero_82
						}
						return hx_value_81.(map[string]any)
					}(cursor.tokens.Get(hx_post_80))
				}
				endToken := hx_if_83
				if (endToken == nil) || !hxrt.StringEqualStringPtr(func(hx_obj_84 map[string]any) *string {
					hx_field_85 := hx_obj_84["p"]
					if hx_field_85 == nil {
						var hx_zero_86 *string
						return hx_zero_86
					}
					return hx_field_85.(*string)
				}(endToken), hxrt.StringFromLiteral("end")) {
					hxrt.Throw(hxrt.StringFromLiteral("Unclosed 'else'"))
				}
			} else {
				nextToken["p"] = hxrt.StringSubstrStringPtr(func(hx_obj_87 map[string]any) *string {
					hx_field_88 := hx_obj_87["p"]
					if hx_field_88 == nil {
						var hx_zero_89 *string
						return hx_zero_89
					}
					return hx_field_88.(*string)
				}(nextToken), 4, int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(func(hx_obj_90 map[string]any) *string {
					hx_field_91 := hx_obj_90["p"]
					if hx_field_91 == nil {
						var hx_zero_92 *string
						return hx_zero_92
					}
					return hx_field_91.(*string)
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
		var hx_if_102 map[string]any
		if cursor.index >= cursor.tokens.Len() {
			hx_if_102 = nil
		} else {
			hx_post_99 := cursor.index
			cursor.index = int(int32((cursor.index + 1)))
			hx_if_102 = func(hx_value_100 any) map[string]any {
				if hx_value_100 == nil {
					var hx_zero_101 map[string]any
					return hx_zero_101
				}
				return hx_value_100.(map[string]any)
			}(cursor.tokens.Get(hx_post_99))
		}
		endToken_1 := hx_if_102
		if (endToken_1 == nil) || !hxrt.StringEqualStringPtr(func(hx_obj_103 map[string]any) *string {
			hx_field_104 := hx_obj_103["p"]
			if hx_field_104 == nil {
				var hx_zero_105 *string
				return hx_zero_105
			}
			return hx_field_104.(*string)
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
		if func(hx_obj_111 map[string]any) int {
			hx_field_112 := hx_obj_111["pos"]
			if hx_field_112 == nil {
				var hx_zero_113 int
				return hx_zero_113
			}
			return hx_field_112.(int)
		}(p) != 0 {
			hx_obj_107 := map[string]any{}
			hx_obj_107["p"] = hxrt.StringSubstrStringPtr(data, 0, func(hx_obj_108 map[string]any) int {
				hx_field_109 := hx_obj_108["pos"]
				if hx_field_109 == nil {
					var hx_zero_110 int
					return hx_zero_110
				}
				return hx_field_109.(int)
			}(p), true)
			hx_obj_107["s"] = true
			tokens.Push(hx_obj_107)
		}
		token := haxe__Template_expr_splitter.__hx_this.matched(0)
		hx_obj_115 := map[string]any{}
		hx_obj_115["p"] = token
		hx_obj_115["s"] = StringTools_contains(token, hxrt.StringFromLiteral("\""))
		tokens.Push(hx_obj_115)
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
				hx_obj_117 := map[string]any{}
				hx_obj_117["p"] = hxrt.StringSubstrStringPtr(data, i, 0, false)
				hx_obj_117["s"] = true
				tokens.Push(hx_obj_117)
				break
			}
		}
	}
	cursor := New_haxe___Template__ExprCursor(tokens)
	var built func() any
	hxrt.TryCatch(func() {
		built = self.__hx_this.makeExpr(cursor)
		if cursor.index < cursor.tokens.Len() {
			hxrt.Throw(func(hx_obj_122 map[string]any) *string {
				hx_field_123 := hx_obj_122["p"]
				if hx_field_123 == nil {
					var hx_zero_124 *string
					return hx_zero_124
				}
				return hx_field_123.(*string)
			}(func(hx_value_120 any) map[string]any {
				if hx_value_120 == nil {
					var hx_zero_121 map[string]any
					return hx_zero_121
				}
				return hx_value_120.(map[string]any)
			}(cursor.tokens.Get(cursor.index))))
		}
	}, func(hx_caught_118 any) {
		switch hx_typed_119 := hx_caught_118.(type) {
		case *string:
			s := hx_typed_119
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Unexpected '"), s), hxrt.StringFromLiteral("' in ")), expr))
		default:
			hxrt.Throw(hx_caught_118)
		}
	})
	me := self
	_ = me
	wrapped := func() any {
		hx_try_return_125 := false
		var hx_try_value_126 any
		hxrt.TryCatch(func() {
			hx_try_value_126 = built()
			hx_try_return_125 = true
			return
		}, func(hx_caught_127 any) {
			exc := hx_caught_127
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Error : "), hxrt.StdString(exc)), hxrt.StringFromLiteral(" in ")), expr))
		})
		if hx_try_return_125 {
			return hx_try_value_126
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
	var hx_if_131 map[string]any
	if cursor.index < cursor.tokens.Len() {
		hx_if_131 = func(hx_value_129 any) map[string]any {
			if hx_value_129 == nil {
				var hx_zero_130 map[string]any
				return hx_zero_130
			}
			return hx_value_129.(map[string]any)
		}(cursor.tokens.Get(cursor.index))
	} else {
		hx_if_131 = nil
	}
	token := hx_if_131
	if (token == nil) || !hxrt.StringEqualStringPtr(func(hx_obj_132 map[string]any) *string {
		hx_field_133 := hx_obj_132["p"]
		if hx_field_133 == nil {
			var hx_zero_134 *string
			return hx_zero_134
		}
		return hx_field_133.(*string)
	}(token), hxrt.StringFromLiteral(".")) {
		return e
	}
	if cursor.index >= cursor.tokens.Len() {
	} else {
		cursor.tokens.Get(func() int {
			hx_post_135 := cursor.index
			cursor.index = int(int32((cursor.index + 1)))
			return hx_post_135
		}())
	}
	var hx_if_139 map[string]any
	if cursor.index >= cursor.tokens.Len() {
		hx_if_139 = nil
	} else {
		hx_post_136 := cursor.index
		cursor.index = int(int32((cursor.index + 1)))
		hx_if_139 = func(hx_value_137 any) map[string]any {
			if hx_value_137 == nil {
				var hx_zero_138 map[string]any
				return hx_zero_138
			}
			return hx_value_137.(map[string]any)
		}(cursor.tokens.Get(hx_post_136))
	}
	field := hx_if_139
	if (field == nil) || !func(hx_obj_144 map[string]any) bool {
		hx_field_145 := hx_obj_144["s"]
		if hx_field_145 == nil {
			var hx_zero_146 bool
			return hx_zero_146
		}
		return hx_field_145.(bool)
	}(field) {
		var hx_if_143 *string
		if field == nil {
			hx_if_143 = hxrt.StringFromLiteral("<eof>")
		} else {
			hx_if_143 = func(hx_obj_140 map[string]any) *string {
				hx_field_141 := hx_obj_140["p"]
				if hx_field_141 == nil {
					var hx_zero_142 *string
					return hx_zero_142
				}
				return hx_field_141.(*string)
			}(field)
		}
		hxrt.Throw(hx_if_143)
	}
	name := haxe__Template_trimExprToken(func(hx_obj_147 map[string]any) *string {
		hx_field_148 := hx_obj_147["p"]
		if hx_field_148 == nil {
			var hx_zero_149 *string
			return hx_zero_149
		}
		return hx_field_148.(*string)
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
		if !haxe__Template_isSpaceOnly(func(hx_obj_152 map[string]any) *string {
			hx_field_153 := hx_obj_152["p"]
			if hx_field_153 == nil {
				var hx_zero_154 *string
				return hx_zero_154
			}
			return hx_field_153.(*string)
		}(func(hx_value_150 any) map[string]any {
			if hx_value_150 == nil {
				var hx_zero_151 map[string]any
				return hx_zero_151
			}
			return hx_value_150.(map[string]any)
		}(cursor.tokens.Get(cursor.index)))) {
			return
		}
		cursor.index = int(int32((cursor.index + 1)))
	}
}

func (self *haxe__Template) makeExpr2(cursor *haxe___Template__ExprCursor) func() any {
	self.__hx_this.skipSpaces(cursor)
	var hx_if_158 map[string]any
	if cursor.index >= cursor.tokens.Len() {
		hx_if_158 = nil
	} else {
		hx_post_155 := cursor.index
		cursor.index = int(int32((cursor.index + 1)))
		hx_if_158 = func(hx_value_156 any) map[string]any {
			if hx_value_156 == nil {
				var hx_zero_157 map[string]any
				return hx_zero_157
			}
			return hx_value_156.(map[string]any)
		}(cursor.tokens.Get(hx_post_155))
	}
	token := hx_if_158
	self.__hx_this.skipSpaces(cursor)
	if token == nil {
		hxrt.Throw(hxrt.StringFromLiteral("<eof>"))
	}
	if func(hx_obj_162 map[string]any) bool {
		hx_field_163 := hx_obj_162["s"]
		if hx_field_163 == nil {
			var hx_zero_164 bool
			return hx_zero_164
		}
		return hx_field_163.(bool)
	}(token) {
		return self.__hx_this.makeConst(func(hx_obj_159 map[string]any) *string {
			hx_field_160 := hx_obj_159["p"]
			if hx_field_160 == nil {
				var hx_zero_161 *string
				return hx_zero_161
			}
			return hx_field_160.(*string)
		}(token))
	}
	_g := func(hx_obj_165 map[string]any) *string {
		hx_field_166 := hx_obj_165["p"]
		if hx_field_166 == nil {
			var hx_zero_167 *string
			return hx_zero_167
		}
		return hx_field_166.(*string)
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
		var hx_if_171 map[string]any
		if cursor.index >= cursor.tokens.Len() {
			hx_if_171 = nil
		} else {
			hx_post_168 := cursor.index
			cursor.index = int(int32((cursor.index + 1)))
			hx_if_171 = func(hx_value_169 any) map[string]any {
				if hx_value_169 == nil {
					var hx_zero_170 map[string]any
					return hx_zero_170
				}
				return hx_value_169.(map[string]any)
			}(cursor.tokens.Get(hx_post_168))
		}
		op := hx_if_171
		if (op == nil) || func(hx_obj_176 map[string]any) bool {
			hx_field_177 := hx_obj_176["s"]
			if hx_field_177 == nil {
				var hx_zero_178 bool
				return hx_zero_178
			}
			return hx_field_177.(bool)
		}(op) {
			var hx_if_175 *string
			if op == nil {
				hx_if_175 = hxrt.StringFromLiteral("<eof>")
			} else {
				hx_if_175 = func(hx_obj_172 map[string]any) *string {
					hx_field_173 := hx_obj_172["p"]
					if hx_field_173 == nil {
						var hx_zero_174 *string
						return hx_zero_174
					}
					return hx_field_173.(*string)
				}(op)
			}
			hxrt.Throw(hx_if_175)
		}
		if hxrt.StringEqualStringPtr(func(hx_obj_179 map[string]any) *string {
			hx_field_180 := hx_obj_179["p"]
			if hx_field_180 == nil {
				var hx_zero_181 *string
				return hx_zero_181
			}
			return hx_field_180.(*string)
		}(op), hxrt.StringFromLiteral(")")) {
			return e1
		}
		self.__hx_this.skipSpaces(cursor)
		e2 := self.__hx_this.makeExpr(cursor)
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
		close := hx_if_185
		self.__hx_this.skipSpaces(cursor)
		if (close == nil) || !hxrt.StringEqualStringPtr(func(hx_obj_190 map[string]any) *string {
			hx_field_191 := hx_obj_190["p"]
			if hx_field_191 == nil {
				var hx_zero_192 *string
				return hx_zero_192
			}
			return hx_field_191.(*string)
		}(close), hxrt.StringFromLiteral(")")) {
			var hx_if_189 *string
			if close == nil {
				hx_if_189 = hxrt.StringFromLiteral("<eof>")
			} else {
				hx_if_189 = func(hx_obj_186 map[string]any) *string {
					hx_field_187 := hx_obj_186["p"]
					if hx_field_187 == nil {
						var hx_zero_188 *string
						return hx_zero_188
					}
					return hx_field_187.(*string)
				}(close)
			}
			hxrt.Throw(hx_if_189)
		}
		_g_1 := func(hx_obj_193 map[string]any) *string {
			hx_field_194 := hx_obj_193["p"]
			if hx_field_194 == nil {
				var hx_zero_195 *string
				return hx_zero_195
			}
			return hx_field_194.(*string)
		}(op)
		var hx_switch_196 func() any
		switch *hxrt.StdString(_g_1) {
		case *hxrt.StdString(hxrt.StringFromLiteral("!=")):
			hx_switch_196 = func() any {
				return !hxrt.HaxeEqual(e1(), e2())
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("&&")):
			hx_switch_196 = func() any {
				return (haxe__Template_valueAsBool(e1()) && haxe__Template_valueAsBool(e2()))
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("*")):
			hx_switch_196 = func() any {
				var left any = e1()
				var right any = e2()
				return (haxe__Template_valueAsFloat(left) * haxe__Template_valueAsFloat(right))
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("+")):
			hx_switch_196 = func() any {
				return haxe__Template_addValues(e1(), e2())
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("-")):
			hx_switch_196 = func() any {
				var left any = e1()
				var right any = e2()
				return (haxe__Template_valueAsFloat(left) - haxe__Template_valueAsFloat(right))
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("/")):
			hx_switch_196 = func() any {
				var left any = e1()
				var right any = e2()
				return (haxe__Template_valueAsFloat(left) / haxe__Template_valueAsFloat(right))
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("<")):
			hx_switch_196 = func() any {
				return (haxe__Template_compareValues(e1(), e2()) < 0)
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("<=")):
			hx_switch_196 = func() any {
				return (haxe__Template_compareValues(e1(), e2()) <= 0)
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("==")):
			hx_switch_196 = func() any {
				return hxrt.HaxeEqual(e1(), e2())
			}
		case *hxrt.StdString(hxrt.StringFromLiteral(">")):
			hx_switch_196 = func() any {
				return (haxe__Template_compareValues(e1(), e2()) > 0)
			}
		case *hxrt.StdString(hxrt.StringFromLiteral(">=")):
			hx_switch_196 = func() any {
				return (haxe__Template_compareValues(e1(), e2()) >= 0)
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("||")):
			hx_switch_196 = func() any {
				return (haxe__Template_valueAsBool(e1()) || haxe__Template_valueAsBool(e2()))
			}
		default:
			hx_switch_196 = func() func() any {
				hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Unknown operation "), func(hx_obj_197 map[string]any) *string {
					hx_field_198 := hx_obj_197["p"]
					if hx_field_198 == nil {
						var hx_zero_199 *string
						return hx_zero_199
					}
					return hx_field_198.(*string)
				}(op)))
				var hx_throw_zero_200 func() any
				return hx_throw_zero_200
			}()
		}
		return hx_switch_196
	case *hxrt.StdString(hxrt.StringFromLiteral("-")):
		inner_1 := self.__hx_this.makeExpr(cursor)
		return func() any {
			return -haxe__Template_valueAsFloat(inner_1())
		}
	default:
		hxrt.Throw(func(hx_obj_201 map[string]any) *string {
			hx_field_202 := hx_obj_201["p"]
			if hx_field_202 == nil {
				var hx_zero_203 *string
				return hx_zero_203
			}
			return hx_field_202.(*string)
		}(token))
		var hx_throw_zero_204 func() any
		return hx_throw_zero_204
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
			item := func(hx_value_205 any) *haxe___Template__TemplateExpr {
				if hx_value_205 == nil {
					var hx_zero_206 *haxe___Template__TemplateExpr
					return hx_zero_206
				}
				return hx_value_205.(*haxe___Template__TemplateExpr)
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
			hx_arr_207 := self.stack
			hx_arr_207.Push(self.context)
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
		}, func(hx_caught_208 any) {
			hx_tmp := hx_caught_208
			_ = hx_tmp
			hxrt.TryCatch(func() {
				if hxrt.AnyEqualsNull(value_1) || !Reflect_hasField(value_1, hxrt.StringFromLiteral("hasNext")) {
					hxrt.Throw(nil)
				}
				iterator = value_1
			}, func(hx_caught_210 any) {
				hx_tmp_1 := hx_caught_210
				_ = hx_tmp_1
				hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Cannot iter on "), hxrt.StdString(value_1)))
			})
		})
		var hasNext any = Reflect_field(iterator, hxrt.StringFromLiteral("hasNext"))
		var next any = Reflect_field(iterator, hxrt.StringFromLiteral("next"))
		if hxrt.AnyEqualsNull(hasNext) || hxrt.AnyEqualsNull(next) {
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Cannot iter on "), hxrt.StdString(value_1)))
		}
		hx_arr_212 := self.stack
		hx_arr_212.Push(self.context)
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
			param := func(hx_value_214 any) *haxe___Template__TemplateExpr {
				if hx_value_214 == nil {
					var hx_zero_215 *haxe___Template__TemplateExpr
					return hx_zero_215
				}
				return hx_value_214.(*haxe___Template__TemplateExpr)
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
		}, func(hx_caught_218 any) {
			err := hx_caught_218
			var hx_try_220 *string
			hxrt.TryCatch(func() {
				hx_try_220 = haxe__Template_joinDynamicArgs(callArgs)
			}, func(hx_caught_221 any) {
				hx_tmp_2 := hx_caught_221
				_ = hx_tmp_2
				hx_try_220 = hxrt.StringFromLiteral("???")
			})
			argsText := hx_try_220
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
		hx_post_223 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_223
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
		var hx_if_226 int
		if leftFloat < rightFloat {
			hx_if_226 = -1
		} else {
			var hx_if_225 int
			if leftFloat > rightFloat {
				hx_if_225 = 1
			} else {
				hx_if_225 = 0
			}
			hx_if_226 = hx_if_225
		}
		return hx_if_226
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
	hx_obj_227 := map[string]any{}
	return hx_obj_227
}())

func haxe__Template_isSpaceOnly(value *string) bool {
	var _g_s *string
	var _g_offset int
	_g_offset = 0
	_g_s = value
	for _g_offset < hxrt.StringLengthStringPtr(_g_s) {
		code := hxrt.StringCharCodeAtStringPtr(_g_s, func() int {
			hx_post_228 := _g_offset
			_g_offset = int(int32((_g_offset + 1)))
			return hx_post_228
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
		hx_post_229 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_229
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
				hx_post_230 := _g_offset
				_g_offset = int(int32((_g_offset + 1)))
				return hx_post_230
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
				var hx_if_231 float64
				if exponentSign < 0 {
					hx_if_231 = (result / 10.0)
				} else {
					hx_if_231 = (result * 10.0)
				}
				result = hx_if_231
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
	var hx_if_234 map[string]any
	if cursor.index < cursor.tokens.Len() {
		hx_if_234 = func(hx_value_232 any) map[string]any {
			if hx_value_232 == nil {
				var hx_zero_233 map[string]any
				return hx_zero_233
			}
			return hx_value_232.(map[string]any)
		}(cursor.tokens.Get(cursor.index))
	} else {
		hx_if_234 = nil
	}
	return hx_if_234
}

func haxe__Template_peekToken(cursor *haxe___Template__TokenCursor) map[string]any {
	var hx_if_237 map[string]any
	if cursor.index < cursor.tokens.Len() {
		hx_if_237 = func(hx_value_235 any) map[string]any {
			if hx_value_235 == nil {
				var hx_zero_236 map[string]any
				return hx_zero_236
			}
			return hx_value_235.(map[string]any)
		}(cursor.tokens.Get(cursor.index))
	} else {
		hx_if_237 = nil
	}
	return hx_if_237
}

func haxe__Template_popExprToken(cursor *haxe___Template__ExprCursor) map[string]any {
	if cursor.index >= cursor.tokens.Len() {
		return nil
	}
	hx_post_238 := cursor.index
	cursor.index = int(int32((cursor.index + 1)))
	return func(hx_value_239 any) map[string]any {
		if hx_value_239 == nil {
			var hx_zero_240 map[string]any
			return hx_zero_240
		}
		return hx_value_239.(map[string]any)
	}(cursor.tokens.Get(hx_post_238))
}

func haxe__Template_popToken(cursor *haxe___Template__TokenCursor) map[string]any {
	if cursor.index >= cursor.tokens.Len() {
		return nil
	}
	hx_post_241 := cursor.index
	cursor.index = int(int32((cursor.index + 1)))
	return func(hx_value_242 any) map[string]any {
		if hx_value_242 == nil {
			var hx_zero_243 map[string]any
			return hx_zero_243
		}
		return hx_value_242.(map[string]any)
	}(cursor.tokens.Get(hx_post_241))
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
		return func(hx_value_244 any) float64 {
			if hx_value_244 == nil {
				var hx_zero_245 float64
				return hx_zero_245
			}
			return hx_value_244.(float64)
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
		return haxe__Template_parseFloatLiteral(hxrt.StdString(func(hx_value_246 any) *string {
			if hx_value_246 == nil {
				var hx_zero_247 *string
				return hx_zero_247
			}
			return hx_value_246.(*string)
		}(value)))
	}
	hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Expected numeric expression value, got "), hxrt.StdString(value)))
	var hx_throw_zero_248 float64
	return hx_throw_zero_248
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
