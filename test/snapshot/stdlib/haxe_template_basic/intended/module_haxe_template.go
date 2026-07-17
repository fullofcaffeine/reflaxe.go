package main

import "snapshot/hxrt"

type I_haxe___Template__TokenCursor interface {
}

type haxe___Template__TokenCursor struct {
	__hx_this I_haxe___Template__TokenCursor
	tokens    []map[string]any
	index     int
}

func New_haxe___Template__TokenCursor(tokens []map[string]any) *haxe___Template__TokenCursor {
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
	tokens    []map[string]any
	index     int
}

func New_haxe___Template__ExprCursor(tokens []map[string]any) *haxe___Template__ExprCursor {
	self := &haxe___Template__ExprCursor{}
	self.__hx_this = self
	self.tokens = tokens
	self.index = 0
	return self
}

type I_haxe__Template interface {
	execute(context any, macros any) *string
	resolve(v *string) any
	parseTokens(data *string) []map[string]any
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
	stack     []any
	output    *string
}

func New_haxe__Template(str *string) *haxe__Template {
	self := &haxe__Template{}
	self.__hx_this = self
	cursor := New_haxe___Template__TokenCursor(self.parseTokens(str))
	self.expr = self.parseBlock(cursor)
	if cursor.index < len(cursor.tokens) {
		token := cursor.tokens[cursor.index]
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Unexpected '"), hxrt.StdString(func(hx_obj_5 map[string]any) bool {
			hx_field_6 := hx_obj_5["s"]
			if hx_field_6 == nil {
				var hx_zero_7 bool
				return hx_zero_7
			}
			return hx_field_6.(bool)
		}(token))), hxrt.StringFromLiteral("'")))
	}
	return self
}

func (self *haxe__Template) execute(context any, macros any) *string {
	var hx_if_9 any
	if hxrt.AnyEqualsNull(macros) {
		hx_obj_8 := map[string]any{}
		hx_if_9 = hx_obj_8
	} else {
		hx_if_9 = macros
	}
	self.macros = hx_if_9
	self.context = context
	self.stack = []any{}
	self.output = hxrt.StringFromLiteral("")
	self.run(self.expr)
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
	for _g < len(_g1) {
		var ctx any = _g1[_g]
		_g = int(int32((_g + 1)))
		var value_1 any = Reflect_field(ctx, v)
		if !hxrt.AnyEqualsNull(value_1) || Reflect_hasField(ctx, v) {
			return value_1
		}
	}
	return Reflect_field(haxe__Template_globals, v)
}

func (self *haxe__Template) parseTokens(data *string) []map[string]any {
	tokens := []map[string]any{}
	for haxe__Template_splitter.match(data) {
		p := haxe__Template_splitter.matchedPos()
		if func(hx_obj_15 map[string]any) int {
			hx_field_16 := hx_obj_15["pos"]
			if hx_field_16 == nil {
				var hx_zero_17 int
				return hx_zero_17
			}
			return hx_field_16.(int)
		}(p) > 0 {
			tokens = append(tokens, func() map[string]any {
				hx_obj_11 := map[string]any{}
				hx_obj_11["p"] = hxrt.StringSubstrStringPtr(data, 0, func(hx_obj_12 map[string]any) int {
					hx_field_13 := hx_obj_12["pos"]
					if hx_field_13 == nil {
						var hx_zero_14 int
						return hx_zero_14
					}
					return hx_field_13.(int)
				}(p), true)
				hx_obj_11["s"] = true
				hx_obj_11["l"] = nil
				return hx_obj_11
			}())
		}
		if hxrt.StringCharCodeAtAnyStringPtr(data, func(hx_obj_26 map[string]any) int {
			hx_field_27 := hx_obj_26["pos"]
			if hx_field_27 == nil {
				var hx_zero_28 int
				return hx_zero_28
			}
			return hx_field_27.(int)
		}(p)) == 58 {
			tokens = append(tokens, func() map[string]any {
				hx_obj_19 := map[string]any{}
				hx_obj_19["p"] = hxrt.StringSubstrStringPtr(data, int(int32((hxrt.Int32Wrap(func(hx_obj_20 map[string]any) int {
					hx_field_21 := hx_obj_20["pos"]
					if hx_field_21 == nil {
						var hx_zero_22 int
						return hx_zero_22
					}
					return hx_field_21.(int)
				}(p)) + hxrt.Int32Wrap(2)))), int(int32((hxrt.Int32Wrap(func(hx_obj_23 map[string]any) int {
					hx_field_24 := hx_obj_23["len"]
					if hx_field_24 == nil {
						var hx_zero_25 int
						return hx_zero_25
					}
					return hx_field_24.(int)
				}(p)) - hxrt.Int32Wrap(4)))), true)
				hx_obj_19["s"] = false
				hx_obj_19["l"] = nil
				return hx_obj_19
			}())
			data = haxe__Template_splitter.matchedRight()
			continue
		}
		parp := int(int32((hxrt.Int32Wrap(func(hx_obj_29 map[string]any) int {
			hx_field_30 := hx_obj_29["pos"]
			if hx_field_30 == nil {
				var hx_zero_31 int
				return hx_zero_31
			}
			return hx_field_30.(int)
		}(p)) + hxrt.Int32Wrap(func(hx_obj_32 map[string]any) int {
			hx_field_33 := hx_obj_32["len"]
			if hx_field_33 == nil {
				var hx_zero_34 int
				return hx_zero_34
			}
			return hx_field_33.(int)
		}(p)))))
		npar := 1
		params := []*string{}
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
				params = append(params, part)
				part = hxrt.StringFromLiteral("")
			} else {
				part = hxrt.StringConcatStringPtr(part, chunk)
			}
		}
		params = append(params, part)
		tokens = append(tokens, func() map[string]any {
			hx_obj_38 := map[string]any{}
			hx_obj_38["p"] = haxe__Template_splitter.matched(2)
			hx_obj_38["s"] = false
			hx_obj_38["l"] = params
			return hx_obj_38
		}())
		data = hxrt.StringSubstrStringPtr(data, parp, int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(data)) - hxrt.Int32Wrap(parp)))), true)
	}
	if hxrt.StringLengthStringPtr(data) > 0 {
		tokens = append(tokens, func() map[string]any {
			hx_obj_40 := map[string]any{}
			hx_obj_40["p"] = data
			hx_obj_40["s"] = true
			hx_obj_40["l"] = nil
			return hx_obj_40
		}())
	}
	return tokens
}

func (self *haxe__Template) parseBlock(cursor *haxe___Template__TokenCursor) *haxe___Template__TemplateExpr {
	items := []*haxe___Template__TemplateExpr{}
	for cursor.index < len(cursor.tokens) {
		t := cursor.tokens[cursor.index]
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
		items = append(items, self.parse(cursor))
	}
	if len(items) == 1 {
		return items[0]
	}
	return haxe___Template__TemplateExpr_OpBlock(items)
}

func (self *haxe__Template) parse(cursor *haxe___Template__TokenCursor) *haxe___Template__TemplateExpr {
	var hx_if_55 map[string]any
	if cursor.index >= len(cursor.tokens) {
		hx_if_55 = nil
	} else {
		hx_post_54 := cursor.index
		cursor.index = int(int32((cursor.index + 1)))
		hx_if_55 = cursor.tokens[hx_post_54]
	}
	t := hx_if_55
	if t == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Unexpected <eof>"))
	}
	p := func(hx_obj_56 map[string]any) *string {
		hx_field_57 := hx_obj_56["p"]
		if hx_field_57 == nil {
			var hx_zero_58 *string
			return hx_zero_58
		}
		return hx_field_57.(*string)
	}(t)
	if func(hx_obj_59 map[string]any) bool {
		hx_field_60 := hx_obj_59["s"]
		if hx_field_60 == nil {
			var hx_zero_61 bool
			return hx_zero_61
		}
		return hx_field_60.(bool)
	}(t) {
		return haxe___Template__TemplateExpr_OpStr(p)
	}
	if func(hx_obj_66 map[string]any) []*string {
		hx_field_67 := hx_obj_66["l"]
		if hx_field_67 == nil {
			var hx_zero_68 []*string
			return hx_zero_68
		}
		return hx_field_67.([]*string)
	}(t) != nil {
		parsedParams := []*haxe___Template__TemplateExpr{}
		_g := 0
		_g1 := func(hx_obj_62 map[string]any) []*string {
			hx_field_63 := hx_obj_62["l"]
			if hx_field_63 == nil {
				var hx_zero_64 []*string
				return hx_zero_64
			}
			return hx_field_63.([]*string)
		}(t)
		for _g < len(_g1) {
			param := _g1[_g]
			_g = int(int32((_g + 1)))
			parsedParams = append(parsedParams, self.parseBlock(New_haxe___Template__TokenCursor(self.parseTokens(param))))
		}
		return haxe___Template__TemplateExpr_OpMacro(p, parsedParams)
	}
	pos := haxe__Template_kwdEnd(p, hxrt.StringFromLiteral("if"))
	if pos > 0 {
		p = hxrt.StringSubstrStringPtr(p, pos, int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(p)) - hxrt.Int32Wrap(pos)))), true)
		e := self.parseExpr(p)
		eif := self.parseBlock(cursor)
		var hx_if_69 map[string]any
		if cursor.index < len(cursor.tokens) {
			hx_if_69 = cursor.tokens[cursor.index]
		} else {
			hx_if_69 = nil
		}
		nextToken := hx_if_69
		if nextToken == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Unclosed 'if'"))
		}
		var eelse *haxe___Template__TemplateExpr = nil
		if hxrt.StringEqualStringPtr(func(hx_obj_86 map[string]any) *string {
			hx_field_87 := hx_obj_86["p"]
			if hx_field_87 == nil {
				var hx_zero_88 *string
				return hx_zero_88
			}
			return hx_field_87.(*string)
		}(nextToken), hxrt.StringFromLiteral("end")) {
			if cursor.index >= len(cursor.tokens) {
			} else {
				_ = cursor.tokens[func() int {
					hx_post_70 := cursor.index
					cursor.index = int(int32((cursor.index + 1)))
					return hx_post_70
				}()]
			}
		} else {
			if hxrt.StringEqualStringPtr(func(hx_obj_83 map[string]any) *string {
				hx_field_84 := hx_obj_83["p"]
				if hx_field_84 == nil {
					var hx_zero_85 *string
					return hx_zero_85
				}
				return hx_field_84.(*string)
			}(nextToken), hxrt.StringFromLiteral("else")) {
				if cursor.index >= len(cursor.tokens) {
				} else {
					_ = cursor.tokens[func() int {
						hx_post_71 := cursor.index
						cursor.index = int(int32((cursor.index + 1)))
						return hx_post_71
					}()]
				}
				eelse = self.parseBlock(cursor)
				var hx_if_73 map[string]any
				if cursor.index >= len(cursor.tokens) {
					hx_if_73 = nil
				} else {
					hx_post_72 := cursor.index
					cursor.index = int(int32((cursor.index + 1)))
					hx_if_73 = cursor.tokens[hx_post_72]
				}
				endToken := hx_if_73
				if (endToken == nil) || !hxrt.StringEqualStringPtr(func(hx_obj_74 map[string]any) *string {
					hx_field_75 := hx_obj_74["p"]
					if hx_field_75 == nil {
						var hx_zero_76 *string
						return hx_zero_76
					}
					return hx_field_75.(*string)
				}(endToken), hxrt.StringFromLiteral("end")) {
					hxrt.Throw(hxrt.StringFromLiteral("Unclosed 'else'"))
				}
			} else {
				nextToken["p"] = hxrt.StringSubstrStringPtr(func(hx_obj_77 map[string]any) *string {
					hx_field_78 := hx_obj_77["p"]
					if hx_field_78 == nil {
						var hx_zero_79 *string
						return hx_zero_79
					}
					return hx_field_78.(*string)
				}(nextToken), 4, int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(func(hx_obj_80 map[string]any) *string {
					hx_field_81 := hx_obj_80["p"]
					if hx_field_81 == nil {
						var hx_zero_82 *string
						return hx_zero_82
					}
					return hx_field_81.(*string)
				}(nextToken))) - hxrt.Int32Wrap(4)))), true)
				eelse = self.parse(cursor)
			}
		}
		return haxe___Template__TemplateExpr_OpIf(e, eif, eelse)
	}
	pos = haxe__Template_kwdEnd(p, hxrt.StringFromLiteral("foreach"))
	if pos >= 0 {
		p = hxrt.StringSubstrStringPtr(p, pos, int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(p)) - hxrt.Int32Wrap(pos)))), true)
		e_1 := self.parseExpr(p)
		efor := self.parseBlock(cursor)
		var hx_if_90 map[string]any
		if cursor.index >= len(cursor.tokens) {
			hx_if_90 = nil
		} else {
			hx_post_89 := cursor.index
			cursor.index = int(int32((cursor.index + 1)))
			hx_if_90 = cursor.tokens[hx_post_89]
		}
		endToken_1 := hx_if_90
		if (endToken_1 == nil) || !hxrt.StringEqualStringPtr(func(hx_obj_91 map[string]any) *string {
			hx_field_92 := hx_obj_91["p"]
			if hx_field_92 == nil {
				var hx_zero_93 *string
				return hx_zero_93
			}
			return hx_field_92.(*string)
		}(endToken_1), hxrt.StringFromLiteral("end")) {
			hxrt.Throw(hxrt.StringFromLiteral("Unclosed 'foreach'"))
		}
		return haxe___Template__TemplateExpr_OpForeach(e_1, efor)
	}
	if haxe__Template_expr_splitter.match(p) {
		return haxe___Template__TemplateExpr_OpExpr(self.parseExpr(p))
	}
	return haxe___Template__TemplateExpr_OpVar(p)
}

func (self *haxe__Template) parseExpr(data *string) func() any {
	tokens := []map[string]any{}
	expr := data
	for haxe__Template_expr_splitter.match(data) {
		p := haxe__Template_expr_splitter.matchedPos()
		if func(hx_obj_99 map[string]any) int {
			hx_field_100 := hx_obj_99["pos"]
			if hx_field_100 == nil {
				var hx_zero_101 int
				return hx_zero_101
			}
			return hx_field_100.(int)
		}(p) != 0 {
			tokens = append(tokens, func() map[string]any {
				hx_obj_95 := map[string]any{}
				hx_obj_95["p"] = hxrt.StringSubstrStringPtr(data, 0, func(hx_obj_96 map[string]any) int {
					hx_field_97 := hx_obj_96["pos"]
					if hx_field_97 == nil {
						var hx_zero_98 int
						return hx_zero_98
					}
					return hx_field_97.(int)
				}(p), true)
				hx_obj_95["s"] = true
				return hx_obj_95
			}())
		}
		token := haxe__Template_expr_splitter.matched(0)
		tokens = append(tokens, func() map[string]any {
			hx_obj_103 := map[string]any{}
			hx_obj_103["p"] = token
			hx_obj_103["s"] = StringTools_contains(token, hxrt.StringFromLiteral("\""))
			return hx_obj_103
		}())
		data = haxe__Template_expr_splitter.matchedRight()
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
				tokens = append(tokens, func() map[string]any {
					hx_obj_105 := map[string]any{}
					hx_obj_105["p"] = hxrt.StringSubstrStringPtr(data, i, 0, false)
					hx_obj_105["s"] = true
					return hx_obj_105
				}())
				break
			}
		}
	}
	cursor := New_haxe___Template__ExprCursor(tokens)
	var built func() any
	hxrt.TryCatch(func() {
		built = self.makeExpr(cursor)
		if cursor.index < len(cursor.tokens) {
			hxrt.Throw(func(hx_obj_108 map[string]any) *string {
				hx_field_109 := hx_obj_108["p"]
				if hx_field_109 == nil {
					var hx_zero_110 *string
					return hx_zero_110
				}
				return hx_field_109.(*string)
			}(cursor.tokens[cursor.index]))
		}
	}, func(hx_caught_106 any) {
		switch hx_typed_107 := hx_caught_106.(type) {
		case *string:
			s := hx_typed_107
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Unexpected '"), s), hxrt.StringFromLiteral("' in ")), expr))
		default:
			hxrt.Throw(hx_caught_106)
		}
	})
	me := self
	_ = me
	wrapped := func() any {
		hx_try_return_111 := false
		var hx_try_value_112 any
		hxrt.TryCatch(func() {
			hx_try_value_112 = built()
			hx_try_return_111 = true
			return
		}, func(hx_caught_113 any) {
			exc := hx_caught_113
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Error : "), hxrt.StdString(exc)), hxrt.StringFromLiteral(" in ")), expr))
		})
		if hx_try_return_111 {
			return hx_try_value_112
		}
		return nil
	}
	return wrapped
}

func (self *haxe__Template) makeConst(v *string) func() any {
	haxe__Template_expr_trim.match(v)
	v = haxe__Template_expr_trim.matched(1)
	if hxrt.StringCharCodeAtAnyStringPtr(v, 0) == 34 {
		str := hxrt.StringSubstrStringPtr(v, 1, int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(v)) - hxrt.Int32Wrap(2)))), true)
		literal := func() any {
			return str
		}
		return literal
	}
	if haxe__Template_expr_int.match(v) {
		i := haxe__Template_parseIntLiteral(v)
		intLiteral := func() any {
			return i
		}
		return intLiteral
	}
	if haxe__Template_expr_float.match(v) {
		f := haxe__Template_parseFloatLiteral(v)
		floatLiteral := func() any {
			return f
		}
		return floatLiteral
	}
	me := self
	resolved := func() any {
		return me.resolve(v)
	}
	return resolved
}

func (self *haxe__Template) makePath(e func() any, cursor *haxe___Template__ExprCursor) func() any {
	var hx_if_115 map[string]any
	if cursor.index < len(cursor.tokens) {
		hx_if_115 = cursor.tokens[cursor.index]
	} else {
		hx_if_115 = nil
	}
	token := hx_if_115
	if (token == nil) || !hxrt.StringEqualStringPtr(func(hx_obj_116 map[string]any) *string {
		hx_field_117 := hx_obj_116["p"]
		if hx_field_117 == nil {
			var hx_zero_118 *string
			return hx_zero_118
		}
		return hx_field_117.(*string)
	}(token), hxrt.StringFromLiteral(".")) {
		return e
	}
	if cursor.index >= len(cursor.tokens) {
	} else {
		_ = cursor.tokens[func() int {
			hx_post_119 := cursor.index
			cursor.index = int(int32((cursor.index + 1)))
			return hx_post_119
		}()]
	}
	var hx_if_121 map[string]any
	if cursor.index >= len(cursor.tokens) {
		hx_if_121 = nil
	} else {
		hx_post_120 := cursor.index
		cursor.index = int(int32((cursor.index + 1)))
		hx_if_121 = cursor.tokens[hx_post_120]
	}
	field := hx_if_121
	if (field == nil) || !func(hx_obj_126 map[string]any) bool {
		hx_field_127 := hx_obj_126["s"]
		if hx_field_127 == nil {
			var hx_zero_128 bool
			return hx_zero_128
		}
		return hx_field_127.(bool)
	}(field) {
		var hx_if_125 *string
		if field == nil {
			hx_if_125 = hxrt.StringFromLiteral("<eof>")
		} else {
			hx_if_125 = func(hx_obj_122 map[string]any) *string {
				hx_field_123 := hx_obj_122["p"]
				if hx_field_123 == nil {
					var hx_zero_124 *string
					return hx_zero_124
				}
				return hx_field_123.(*string)
			}(field)
		}
		hxrt.Throw(hx_if_125)
	}
	name := haxe__Template_trimExprToken(func(hx_obj_129 map[string]any) *string {
		hx_field_130 := hx_obj_129["p"]
		if hx_field_130 == nil {
			var hx_zero_131 *string
			return hx_zero_131
		}
		return hx_field_130.(*string)
	}(field))
	return self.makePath(func() any {
		return Reflect_field(e(), name)
	}, cursor)
}

func (self *haxe__Template) makeExpr(cursor *haxe___Template__ExprCursor) func() any {
	return self.makePath(self.makeExpr2(cursor), cursor)
}

func (self *haxe__Template) skipSpaces(cursor *haxe___Template__ExprCursor) {
	for cursor.index < len(cursor.tokens) {
		if !haxe__Template_isSpaceOnly(func(hx_obj_132 map[string]any) *string {
			hx_field_133 := hx_obj_132["p"]
			if hx_field_133 == nil {
				var hx_zero_134 *string
				return hx_zero_134
			}
			return hx_field_133.(*string)
		}(cursor.tokens[cursor.index])) {
			return
		}
		cursor.index = int(int32((cursor.index + 1)))
	}
}

func (self *haxe__Template) makeExpr2(cursor *haxe___Template__ExprCursor) func() any {
	self.skipSpaces(cursor)
	var hx_if_136 map[string]any
	if cursor.index >= len(cursor.tokens) {
		hx_if_136 = nil
	} else {
		hx_post_135 := cursor.index
		cursor.index = int(int32((cursor.index + 1)))
		hx_if_136 = cursor.tokens[hx_post_135]
	}
	token := hx_if_136
	self.skipSpaces(cursor)
	if token == nil {
		hxrt.Throw(hxrt.StringFromLiteral("<eof>"))
	}
	if func(hx_obj_140 map[string]any) bool {
		hx_field_141 := hx_obj_140["s"]
		if hx_field_141 == nil {
			var hx_zero_142 bool
			return hx_zero_142
		}
		return hx_field_141.(bool)
	}(token) {
		return self.makeConst(func(hx_obj_137 map[string]any) *string {
			hx_field_138 := hx_obj_137["p"]
			if hx_field_138 == nil {
				var hx_zero_139 *string
				return hx_zero_139
			}
			return hx_field_138.(*string)
		}(token))
	}
	_g := func(hx_obj_143 map[string]any) *string {
		hx_field_144 := hx_obj_143["p"]
		if hx_field_144 == nil {
			var hx_zero_145 *string
			return hx_zero_145
		}
		return hx_field_144.(*string)
	}(token)
	switch *hxrt.StdString(_g) {
	case *hxrt.StdString(hxrt.StringFromLiteral("!")):
		inner := self.makeExpr(cursor)
		return func() any {
			var value any = inner()
			return (hxrt.AnyEqualsNull(value) || (value == false))
		}
	case *hxrt.StdString(hxrt.StringFromLiteral("(")):
		self.skipSpaces(cursor)
		e1 := self.makeExpr(cursor)
		self.skipSpaces(cursor)
		var hx_if_147 map[string]any
		if cursor.index >= len(cursor.tokens) {
			hx_if_147 = nil
		} else {
			hx_post_146 := cursor.index
			cursor.index = int(int32((cursor.index + 1)))
			hx_if_147 = cursor.tokens[hx_post_146]
		}
		op := hx_if_147
		if (op == nil) || func(hx_obj_152 map[string]any) bool {
			hx_field_153 := hx_obj_152["s"]
			if hx_field_153 == nil {
				var hx_zero_154 bool
				return hx_zero_154
			}
			return hx_field_153.(bool)
		}(op) {
			var hx_if_151 *string
			if op == nil {
				hx_if_151 = hxrt.StringFromLiteral("<eof>")
			} else {
				hx_if_151 = func(hx_obj_148 map[string]any) *string {
					hx_field_149 := hx_obj_148["p"]
					if hx_field_149 == nil {
						var hx_zero_150 *string
						return hx_zero_150
					}
					return hx_field_149.(*string)
				}(op)
			}
			hxrt.Throw(hx_if_151)
		}
		if hxrt.StringEqualStringPtr(func(hx_obj_155 map[string]any) *string {
			hx_field_156 := hx_obj_155["p"]
			if hx_field_156 == nil {
				var hx_zero_157 *string
				return hx_zero_157
			}
			return hx_field_156.(*string)
		}(op), hxrt.StringFromLiteral(")")) {
			return e1
		}
		self.skipSpaces(cursor)
		e2 := self.makeExpr(cursor)
		self.skipSpaces(cursor)
		var hx_if_159 map[string]any
		if cursor.index >= len(cursor.tokens) {
			hx_if_159 = nil
		} else {
			hx_post_158 := cursor.index
			cursor.index = int(int32((cursor.index + 1)))
			hx_if_159 = cursor.tokens[hx_post_158]
		}
		close := hx_if_159
		self.skipSpaces(cursor)
		if (close == nil) || !hxrt.StringEqualStringPtr(func(hx_obj_164 map[string]any) *string {
			hx_field_165 := hx_obj_164["p"]
			if hx_field_165 == nil {
				var hx_zero_166 *string
				return hx_zero_166
			}
			return hx_field_165.(*string)
		}(close), hxrt.StringFromLiteral(")")) {
			var hx_if_163 *string
			if close == nil {
				hx_if_163 = hxrt.StringFromLiteral("<eof>")
			} else {
				hx_if_163 = func(hx_obj_160 map[string]any) *string {
					hx_field_161 := hx_obj_160["p"]
					if hx_field_161 == nil {
						var hx_zero_162 *string
						return hx_zero_162
					}
					return hx_field_161.(*string)
				}(close)
			}
			hxrt.Throw(hx_if_163)
		}
		_g_1 := func(hx_obj_167 map[string]any) *string {
			hx_field_168 := hx_obj_167["p"]
			if hx_field_168 == nil {
				var hx_zero_169 *string
				return hx_zero_169
			}
			return hx_field_168.(*string)
		}(op)
		var hx_switch_170 func() any
		switch *hxrt.StdString(_g_1) {
		case *hxrt.StdString(hxrt.StringFromLiteral("!=")):
			hx_switch_170 = func() any {
				return (e1() != e2())
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("&&")):
			hx_switch_170 = func() any {
				return (haxe__Template_valueAsBool(e1()) && haxe__Template_valueAsBool(e2()))
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("*")):
			hx_switch_170 = func() any {
				var left any = e1()
				var right any = e2()
				return (haxe__Template_valueAsFloat(left) * haxe__Template_valueAsFloat(right))
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("+")):
			hx_switch_170 = func() any {
				return haxe__Template_addValues(e1(), e2())
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("-")):
			hx_switch_170 = func() any {
				var left any = e1()
				var right any = e2()
				return (haxe__Template_valueAsFloat(left) - haxe__Template_valueAsFloat(right))
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("/")):
			hx_switch_170 = func() any {
				var left any = e1()
				var right any = e2()
				return (haxe__Template_valueAsFloat(left) / haxe__Template_valueAsFloat(right))
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("<")):
			hx_switch_170 = func() any {
				return (haxe__Template_compareValues(e1(), e2()) < 0)
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("<=")):
			hx_switch_170 = func() any {
				return (haxe__Template_compareValues(e1(), e2()) <= 0)
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("==")):
			hx_switch_170 = func() any {
				return (e1() == e2())
			}
		case *hxrt.StdString(hxrt.StringFromLiteral(">")):
			hx_switch_170 = func() any {
				return (haxe__Template_compareValues(e1(), e2()) > 0)
			}
		case *hxrt.StdString(hxrt.StringFromLiteral(">=")):
			hx_switch_170 = func() any {
				return (haxe__Template_compareValues(e1(), e2()) >= 0)
			}
		case *hxrt.StdString(hxrt.StringFromLiteral("||")):
			hx_switch_170 = func() any {
				return (haxe__Template_valueAsBool(e1()) || haxe__Template_valueAsBool(e2()))
			}
		default:
			hx_switch_170 = func() func() any {
				hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Unknown operation "), func(hx_obj_171 map[string]any) *string {
					hx_field_172 := hx_obj_171["p"]
					if hx_field_172 == nil {
						var hx_zero_173 *string
						return hx_zero_173
					}
					return hx_field_172.(*string)
				}(op)))
				var hx_throw_zero_174 func() any
				return hx_throw_zero_174
			}()
		}
		return hx_switch_170
	case *hxrt.StdString(hxrt.StringFromLiteral("-")):
		inner_1 := self.makeExpr(cursor)
		return func() any {
			return -haxe__Template_valueAsFloat(inner_1())
		}
	default:
		hxrt.Throw(func(hx_obj_175 map[string]any) *string {
			hx_field_176 := hx_obj_175["p"]
			if hx_field_176 == nil {
				var hx_zero_177 *string
				return hx_zero_177
			}
			return hx_field_176.(*string)
		}(token))
		var hx_throw_zero_178 func() any
		return hx_throw_zero_178
	}
}

func (self *haxe__Template) run(e *haxe___Template__TemplateExpr) {
	switch e.tag {
	case 0:
		_g := e.params[0].(*string)
		v := _g
		self.output = hxrt.StringConcatStringPtr(self.output, hxrt.StdString(self.resolve(v)))
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
		if hxrt.AnyEqualsNull(value) || (value == false) {
			if elseExpr != nil {
				self.run(elseExpr)
			}
		} else {
			self.run(ifExpr)
		}
	case 3:
		_g_3 := e.params[0].(*string)
		str := _g_3
		self.output = hxrt.StringConcatStringPtr(self.output, str)
	case 4:
		_g_4 := e.params[0].([]*haxe___Template__TemplateExpr)
		items := _g_4
		_g_5 := 0
		for _g_5 < len(items) {
			item := items[_g_5]
			_g_5 = int(int32((_g_5 + 1)))
			self.run(item)
		}
	case 5:
		_g_6 := e.params[0].(func() any)
		_g1_1 := e.params[1].(*haxe___Template__TemplateExpr)
		expr_2 := _g_6
		loop := _g1_1
		var value_1 any = expr_2()
		arrayValues := hxrt.TemplateArrayValues(value_1)
		if arrayValues != nil {
			hx_arr_179 := self.stack
			hx_arr_179 = append(hx_arr_179, self.context)
			self.stack = hx_arr_179
			_g_7 := 0
			for _g_7 < len(arrayValues) {
				var ctx any = arrayValues[_g_7]
				_g_7 = int(int32((_g_7 + 1)))
				self.context = ctx
				self.run(loop)
			}
			self.context = self.popStackValue()
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
		}, func(hx_caught_180 any) {
			hx_tmp := hx_caught_180
			_ = hx_tmp
			hxrt.TryCatch(func() {
				if hxrt.AnyEqualsNull(value_1) || !Reflect_hasField(value_1, hxrt.StringFromLiteral("hasNext")) {
					hxrt.Throw(nil)
				}
				iterator = value_1
			}, func(hx_caught_182 any) {
				hx_tmp_1 := hx_caught_182
				_ = hx_tmp_1
				hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Cannot iter on "), hxrt.StdString(value_1)))
			})
		})
		hx_arr_184 := self.stack
		hx_arr_184 = append(hx_arr_184, self.context)
		self.stack = hx_arr_184
		iterable := func(hx_value_185 any) map[string]any {
			if hx_value_185 == nil {
				var hx_zero_186 map[string]any
				return hx_zero_186
			}
			return hx_value_185.(map[string]any)
		}(iterator)
		ctx_1 := iterable
		for func(hx_obj_187 map[string]any) func() bool {
			hx_field_188 := hx_obj_187["hasNext"]
			if hx_field_188 == nil {
				var hx_zero_189 func() bool
				return hx_zero_189
			}
			return hx_field_188.(func() bool)
		}(ctx_1)() {
			var ctx_2 any = func(hx_obj_190 map[string]any) func() any {
				hx_field_191 := hx_obj_190["next"]
				if hx_field_191 == nil {
					var hx_zero_192 func() any
					return hx_zero_192
				}
				return hx_field_191.(func() any)
			}(ctx_1)()
			self.context = ctx_2
			self.run(loop)
		}
		self.context = self.popStackValue()
	case 6:
		_g_8 := e.params[0].(*string)
		_g1_2 := e.params[1].([]*haxe___Template__TemplateExpr)
		name := _g_8
		params := _g1_2
		var fn any = Reflect_field(self.macros, name)
		callArgs := []any{}
		callArgs = append(callArgs, self.resolve)
		_g_9 := 0
		for _g_9 < len(params) {
			param := params[_g_9]
			_g_9 = int(int32((_g_9 + 1)))
			if param.tag == 0 {
				_g_10 := param.params[0].(*string)
				value_2 := _g_10
				callArgs = append(callArgs, self.resolve(value_2))
			} else {
				previous := self.output
				self.output = hxrt.StringFromLiteral("")
				self.run(param)
				callArgs = append(callArgs, self.output)
				self.output = previous
			}
		}
		hxrt.TryCatch(func() {
			self.output = hxrt.StringConcatStringPtr(self.output, hxrt.StdString(hxrt.TemplateCall(fn, callArgs)))
		}, func(hx_caught_196 any) {
			err := hx_caught_196
			var hx_try_198 *string
			hxrt.TryCatch(func() {
				hx_try_198 = haxe__Template_joinDynamicArgs(callArgs)
			}, func(hx_caught_199 any) {
				hx_tmp_2 := hx_caught_199
				_ = hx_tmp_2
				hx_try_198 = hxrt.StringFromLiteral("???")
			})
			argsText := hx_try_198
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Macro call "), name), hxrt.StringFromLiteral("(")), argsText), hxrt.StringFromLiteral(") failed (")), hxrt.StdString(err)), hxrt.StringFromLiteral(")")))
		})
	}
}

func (self *haxe__Template) popStackValue() any {
	lastIndex := int(int32((hxrt.Int32Wrap(len(self.stack)) - hxrt.Int32Wrap(1))))
	var value any = self.stack[lastIndex]
	remaining := []any{}
	_g := 0
	_g1 := lastIndex
	for _g < _g1 {
		hx_post_201 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_201
		remaining = append(remaining, self.stack[index])
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
		var hx_if_204 int
		if leftFloat < rightFloat {
			hx_if_204 = -1
		} else {
			var hx_if_203 int
			if leftFloat > rightFloat {
				hx_if_203 = 1
			} else {
				hx_if_203 = 0
			}
			hx_if_204 = hx_if_203
		}
		return hx_if_204
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
	hx_obj_205 := map[string]any{}
	return hx_obj_205
}())

func haxe__Template_isSpaceOnly(value *string) bool {
	var _g_s *string
	var _g_offset int
	_g_offset = 0
	_g_s = value
	for _g_offset < hxrt.StringLengthStringPtr(_g_s) {
		code := hxrt.StringCharCodeAtStringPtr(_g_s, func() int {
			hx_post_206 := _g_offset
			_g_offset = int(int32((_g_offset + 1)))
			return hx_post_206
		}())
		if code != 32 {
			return false
		}
	}
	return true
}

func haxe__Template_joinDynamicArgs(values []any) *string {
	out := hxrt.StringFromLiteral("")
	_g := 0
	_g1 := len(values)
	for _g < _g1 {
		hx_post_207 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_207
		if index > 0 {
			out = hxrt.StringConcatStringPtr(out, hxrt.StringFromLiteral(","))
		}
		out = hxrt.StringConcatStringPtr(out, hxrt.StdString(values[index]))
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
				hx_post_208 := _g_offset
				_g_offset = int(int32((_g_offset + 1)))
				return hx_post_208
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
				var hx_if_209 float64
				if exponentSign < 0 {
					hx_if_209 = (result / 10.0)
				} else {
					hx_if_209 = (result * 10.0)
				}
				result = hx_if_209
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
	var hx_if_210 map[string]any
	if cursor.index < len(cursor.tokens) {
		hx_if_210 = cursor.tokens[cursor.index]
	} else {
		hx_if_210 = nil
	}
	return hx_if_210
}

func haxe__Template_peekToken(cursor *haxe___Template__TokenCursor) map[string]any {
	var hx_if_211 map[string]any
	if cursor.index < len(cursor.tokens) {
		hx_if_211 = cursor.tokens[cursor.index]
	} else {
		hx_if_211 = nil
	}
	return hx_if_211
}

func haxe__Template_popExprToken(cursor *haxe___Template__ExprCursor) map[string]any {
	if cursor.index >= len(cursor.tokens) {
		return nil
	}
	hx_post_212 := cursor.index
	cursor.index = int(int32((cursor.index + 1)))
	return cursor.tokens[hx_post_212]
}

func haxe__Template_popToken(cursor *haxe___Template__TokenCursor) map[string]any {
	if cursor.index >= len(cursor.tokens) {
		return nil
	}
	hx_post_213 := cursor.index
	cursor.index = int(int32((cursor.index + 1)))
	return cursor.tokens[hx_post_213]
}

var haxe__Template_splitter *EReg = New_EReg(hxrt.StringFromLiteral("(::[A-Za-z0-9_ ()&|!+=/><*.\"-]+::|\\$\\$([A-Za-z0-9_-]+)\\()"), hxrt.StringFromLiteral(""))

func haxe__Template_subtractValues(left any, right any) float64 {
	return (haxe__Template_valueAsFloat(left) - haxe__Template_valueAsFloat(right))
}

func haxe__Template_trimExprToken(value *string) *string {
	haxe__Template_expr_trim.match(value)
	return haxe__Template_expr_trim.matched(1)
}

func haxe__Template_valueAsBool(value any) bool {
	return !(hxrt.AnyEqualsNull(value) || (value == false))
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
		return func(hx_value_214 any) float64 {
			if hx_value_214 == nil {
				var hx_zero_215 float64
				return hx_zero_215
			}
			return hx_value_214.(float64)
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
		return haxe__Template_parseFloatLiteral(hxrt.StdString(func(hx_value_216 any) *string {
			if hx_value_216 == nil {
				var hx_zero_217 *string
				return hx_zero_217
			}
			return hx_value_216.(*string)
		}(value)))
	}
	hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Expected numeric expression value, got "), hxrt.StdString(value)))
	var hx_throw_zero_218 float64
	return hx_throw_zero_218
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

func haxe___Template__TemplateExpr_OpBlock(items []*haxe___Template__TemplateExpr) *haxe___Template__TemplateExpr {
	enumValue := &haxe___Template__TemplateExpr{tag: 4}
	enumValue.params = []any{items}
	return enumValue
}

func haxe___Template__TemplateExpr_OpForeach(expr func() any, loop *haxe___Template__TemplateExpr) *haxe___Template__TemplateExpr {
	enumValue := &haxe___Template__TemplateExpr{tag: 5}
	enumValue.params = []any{expr, loop}
	return enumValue
}

func haxe___Template__TemplateExpr_OpMacro(name *string, params []*haxe___Template__TemplateExpr) *haxe___Template__TemplateExpr {
	enumValue := &haxe___Template__TemplateExpr{tag: 6}
	enumValue.params = []any{name, params}
	return enumValue
}
