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
	if Reflect_isObject(self.context) {
		var value any = Reflect_getProperty(self.context, v)
		if !hxrt.AnyEqualsNull(value) || Reflect_hasField(self.context, v) {
			return value
		}
	}
	_g := 0
	_g1 := self.stack
	for _g < len(_g1) {
		var ctx any = _g1[_g]
		_g = int(int32((_g + 1)))
		var value_1 any = Reflect_getProperty(ctx, v)
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
		if func(hx_obj_10 map[string]any) int {
			hx_field_11 := hx_obj_10["pos"]
			if hx_field_11 == nil {
				var hx_zero_12 int
				return hx_zero_12
			}
			return hx_field_11.(int)
		}(p) > 0 {
			tokens = append(tokens, func() map[string]any {
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
				return hx_obj_13
			}())
		}
		if hxrt.StringCharCodeAtAnyStringPtr(data, func(hx_obj_17 map[string]any) int {
			hx_field_18 := hx_obj_17["pos"]
			if hx_field_18 == nil {
				var hx_zero_19 int
				return hx_zero_19
			}
			return hx_field_18.(int)
		}(p)) == 58 {
			tokens = append(tokens, func() map[string]any {
				hx_obj_20 := map[string]any{}
				hx_obj_20["p"] = hxrt.StringSubstrStringPtr(data, int(int32((hxrt.Int32Wrap(func(hx_obj_21 map[string]any) int {
					hx_field_22 := hx_obj_21["pos"]
					if hx_field_22 == nil {
						var hx_zero_23 int
						return hx_zero_23
					}
					return hx_field_22.(int)
				}(p)) + hxrt.Int32Wrap(2)))), int(int32((hxrt.Int32Wrap(func(hx_obj_24 map[string]any) int {
					hx_field_25 := hx_obj_24["len"]
					if hx_field_25 == nil {
						var hx_zero_26 int
						return hx_zero_26
					}
					return hx_field_25.(int)
				}(p)) - hxrt.Int32Wrap(4)))), true)
				hx_obj_20["s"] = false
				hx_obj_20["l"] = nil
				return hx_obj_20
			}())
			data = haxe__Template_splitter.matchedRight()
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
		params := []*string{}
		part := hxrt.StringFromLiteral("")
		for true {
			c := hxrt.StringCharCodeAtAnyStringPtr(data, parp)
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
						var hx_throw_zero_33 []map[string]any
						return hx_throw_zero_33
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
			hx_obj_34 := map[string]any{}
			hx_obj_34["p"] = haxe__Template_splitter.matched(2)
			hx_obj_34["s"] = false
			hx_obj_34["l"] = params
			return hx_obj_34
		}())
		data = hxrt.StringSubstrStringPtr(data, parp, int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(data)) - hxrt.Int32Wrap(parp)))), true)
	}
	if hxrt.StringLengthStringPtr(data) > 0 {
		tokens = append(tokens, func() map[string]any {
			hx_obj_35 := map[string]any{}
			hx_obj_35["p"] = data
			hx_obj_35["s"] = true
			hx_obj_35["l"] = nil
			return hx_obj_35
		}())
	}
	return tokens
}

func (self *haxe__Template) parseBlock(cursor *haxe___Template__TokenCursor) *haxe___Template__TemplateExpr {
	items := []*haxe___Template__TemplateExpr{}
	for cursor.index < len(cursor.tokens) {
		t := cursor.tokens[cursor.index]
		if !func(hx_obj_36 map[string]any) bool {
			hx_field_37 := hx_obj_36["s"]
			if hx_field_37 == nil {
				var hx_zero_38 bool
				return hx_zero_38
			}
			return hx_field_37.(bool)
		}(t) && ((hxrt.StringEqualStringPtr(func(hx_obj_39 map[string]any) *string {
			hx_field_40 := hx_obj_39["p"]
			if hx_field_40 == nil {
				var hx_zero_41 *string
				return hx_zero_41
			}
			return hx_field_40.(*string)
		}(t), hxrt.StringFromLiteral("end")) || hxrt.StringEqualStringPtr(func(hx_obj_42 map[string]any) *string {
			hx_field_43 := hx_obj_42["p"]
			if hx_field_43 == nil {
				var hx_zero_44 *string
				return hx_zero_44
			}
			return hx_field_43.(*string)
		}(t), hxrt.StringFromLiteral("else"))) || hxrt.StringEqualStringPtr(hxrt.StringSubstrStringPtr(func(hx_obj_45 map[string]any) *string {
			hx_field_46 := hx_obj_45["p"]
			if hx_field_46 == nil {
				var hx_zero_47 *string
				return hx_zero_47
			}
			return hx_field_46.(*string)
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
	var hx_if_49 map[string]any
	if cursor.index >= len(cursor.tokens) {
		hx_if_49 = nil
	} else {
		hx_post_48 := cursor.index
		cursor.index = int(int32((cursor.index + 1)))
		hx_if_49 = cursor.tokens[hx_post_48]
	}
	t := hx_if_49
	if t == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Unexpected <eof>"))
		var hx_throw_zero_50 *haxe___Template__TemplateExpr
		return hx_throw_zero_50
	}
	p := func(hx_obj_51 map[string]any) *string {
		hx_field_52 := hx_obj_51["p"]
		if hx_field_52 == nil {
			var hx_zero_53 *string
			return hx_zero_53
		}
		return hx_field_52.(*string)
	}(t)
	if func(hx_obj_54 map[string]any) bool {
		hx_field_55 := hx_obj_54["s"]
		if hx_field_55 == nil {
			var hx_zero_56 bool
			return hx_zero_56
		}
		return hx_field_55.(bool)
	}(t) {
		return haxe___Template__TemplateExpr_OpStr(p)
	}
	if func(hx_obj_57 map[string]any) []*string {
		hx_field_58 := hx_obj_57["l"]
		if hx_field_58 == nil {
			var hx_zero_59 []*string
			return hx_zero_59
		}
		return hx_field_58.([]*string)
	}(t) != nil {
		parsedParams := []*haxe___Template__TemplateExpr{}
		_g := 0
		_g1 := func(hx_obj_60 map[string]any) []*string {
			hx_field_61 := hx_obj_60["l"]
			if hx_field_61 == nil {
				var hx_zero_62 []*string
				return hx_zero_62
			}
			return hx_field_61.([]*string)
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
		var hx_if_63 map[string]any
		if cursor.index < len(cursor.tokens) {
			hx_if_63 = cursor.tokens[cursor.index]
		} else {
			hx_if_63 = nil
		}
		nextToken := hx_if_63
		if nextToken == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Unclosed 'if'"))
			var hx_throw_zero_64 *haxe___Template__TemplateExpr
			return hx_throw_zero_64
		}
		var eelse *haxe___Template__TemplateExpr = nil
		if hxrt.StringEqualStringPtr(func(hx_obj_65 map[string]any) *string {
			hx_field_66 := hx_obj_65["p"]
			if hx_field_66 == nil {
				var hx_zero_67 *string
				return hx_zero_67
			}
			return hx_field_66.(*string)
		}(nextToken), hxrt.StringFromLiteral("end")) {
			if cursor.index >= len(cursor.tokens) {
			} else {
				_ = cursor.tokens[func() int {
					hx_post_68 := cursor.index
					cursor.index = int(int32((cursor.index + 1)))
					return hx_post_68
				}()]
			}
		} else {
			if hxrt.StringEqualStringPtr(func(hx_obj_69 map[string]any) *string {
				hx_field_70 := hx_obj_69["p"]
				if hx_field_70 == nil {
					var hx_zero_71 *string
					return hx_zero_71
				}
				return hx_field_70.(*string)
			}(nextToken), hxrt.StringFromLiteral("else")) {
				if cursor.index >= len(cursor.tokens) {
				} else {
					_ = cursor.tokens[func() int {
						hx_post_72 := cursor.index
						cursor.index = int(int32((cursor.index + 1)))
						return hx_post_72
					}()]
				}
				eelse = self.parseBlock(cursor)
				var hx_if_74 map[string]any
				if cursor.index >= len(cursor.tokens) {
					hx_if_74 = nil
				} else {
					hx_post_73 := cursor.index
					cursor.index = int(int32((cursor.index + 1)))
					hx_if_74 = cursor.tokens[hx_post_73]
				}
				endToken := hx_if_74
				if (endToken == nil) || !hxrt.StringEqualStringPtr(func(hx_obj_75 map[string]any) *string {
					hx_field_76 := hx_obj_75["p"]
					if hx_field_76 == nil {
						var hx_zero_77 *string
						return hx_zero_77
					}
					return hx_field_76.(*string)
				}(endToken), hxrt.StringFromLiteral("end")) {
					hxrt.Throw(hxrt.StringFromLiteral("Unclosed 'else'"))
					var hx_throw_zero_78 *haxe___Template__TemplateExpr
					return hx_throw_zero_78
				}
			} else {
				nextToken["p"] = hxrt.StringSubstrStringPtr(func(hx_obj_79 map[string]any) *string {
					hx_field_80 := hx_obj_79["p"]
					if hx_field_80 == nil {
						var hx_zero_81 *string
						return hx_zero_81
					}
					return hx_field_80.(*string)
				}(nextToken), 4, int(int32((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(func(hx_obj_82 map[string]any) *string {
					hx_field_83 := hx_obj_82["p"]
					if hx_field_83 == nil {
						var hx_zero_84 *string
						return hx_zero_84
					}
					return hx_field_83.(*string)
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
		var hx_if_86 map[string]any
		if cursor.index >= len(cursor.tokens) {
			hx_if_86 = nil
		} else {
			hx_post_85 := cursor.index
			cursor.index = int(int32((cursor.index + 1)))
			hx_if_86 = cursor.tokens[hx_post_85]
		}
		endToken_1 := hx_if_86
		if (endToken_1 == nil) || !hxrt.StringEqualStringPtr(func(hx_obj_87 map[string]any) *string {
			hx_field_88 := hx_obj_87["p"]
			if hx_field_88 == nil {
				var hx_zero_89 *string
				return hx_zero_89
			}
			return hx_field_88.(*string)
		}(endToken_1), hxrt.StringFromLiteral("end")) {
			hxrt.Throw(hxrt.StringFromLiteral("Unclosed 'foreach'"))
			var hx_throw_zero_90 *haxe___Template__TemplateExpr
			return hx_throw_zero_90
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
		if func(hx_obj_91 map[string]any) int {
			hx_field_92 := hx_obj_91["pos"]
			if hx_field_92 == nil {
				var hx_zero_93 int
				return hx_zero_93
			}
			return hx_field_92.(int)
		}(p) != 0 {
			tokens = append(tokens, func() map[string]any {
				hx_obj_94 := map[string]any{}
				hx_obj_94["p"] = hxrt.StringSubstrStringPtr(data, 0, func(hx_obj_95 map[string]any) int {
					hx_field_96 := hx_obj_95["pos"]
					if hx_field_96 == nil {
						var hx_zero_97 int
						return hx_zero_97
					}
					return hx_field_96.(int)
				}(p), true)
				hx_obj_94["s"] = true
				return hx_obj_94
			}())
		}
		token := haxe__Template_expr_splitter.matched(0)
		tokens = append(tokens, func() map[string]any {
			hx_obj_98 := map[string]any{}
			hx_obj_98["p"] = token
			hx_obj_98["s"] = StringTools_contains(token, hxrt.StringFromLiteral("\""))
			return hx_obj_98
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
					hx_obj_99 := map[string]any{}
					hx_obj_99["p"] = hxrt.StringSubstrStringPtr(data, i, 0, false)
					hx_obj_99["s"] = true
					return hx_obj_99
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
			hxrt.Throw(func(hx_obj_102 map[string]any) *string {
				hx_field_103 := hx_obj_102["p"]
				if hx_field_103 == nil {
					var hx_zero_104 *string
					return hx_zero_104
				}
				return hx_field_103.(*string)
			}(cursor.tokens[cursor.index]))
		}
	}, func(hx_caught_100 any) {
		switch hx_typed_101 := hx_caught_100.(type) {
		case *string:
			s := hx_typed_101
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Unexpected '"), s), hxrt.StringFromLiteral("' in ")), expr))
		default:
			hxrt.Throw(hx_caught_100)
		}
	})
	me := self
	_ = me
	wrapped := func() any {
		hx_try_return_105 := false
		var hx_try_value_106 any
		hxrt.TryCatch(func() {
			hx_try_value_106 = built()
			hx_try_return_105 = true
			return
		}, func(hx_caught_107 any) {
			exc := hx_caught_107
			hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Error : "), hxrt.StdString(exc)), hxrt.StringFromLiteral(" in ")), expr))
		})
		if hx_try_return_105 {
			return hx_try_value_106
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
	var hx_if_109 map[string]any
	if cursor.index < len(cursor.tokens) {
		hx_if_109 = cursor.tokens[cursor.index]
	} else {
		hx_if_109 = nil
	}
	token := hx_if_109
	if (token == nil) || !hxrt.StringEqualStringPtr(func(hx_obj_110 map[string]any) *string {
		hx_field_111 := hx_obj_110["p"]
		if hx_field_111 == nil {
			var hx_zero_112 *string
			return hx_zero_112
		}
		return hx_field_111.(*string)
	}(token), hxrt.StringFromLiteral(".")) {
		return e
	}
	if cursor.index >= len(cursor.tokens) {
	} else {
		_ = cursor.tokens[func() int {
			hx_post_113 := cursor.index
			cursor.index = int(int32((cursor.index + 1)))
			return hx_post_113
		}()]
	}
	var hx_if_115 map[string]any
	if cursor.index >= len(cursor.tokens) {
		hx_if_115 = nil
	} else {
		hx_post_114 := cursor.index
		cursor.index = int(int32((cursor.index + 1)))
		hx_if_115 = cursor.tokens[hx_post_114]
	}
	field := hx_if_115
	if (field == nil) || !func(hx_obj_116 map[string]any) bool {
		hx_field_117 := hx_obj_116["s"]
		if hx_field_117 == nil {
			var hx_zero_118 bool
			return hx_zero_118
		}
		return hx_field_117.(bool)
	}(field) {
		var hx_if_122 *string
		if field == nil {
			hx_if_122 = hxrt.StringFromLiteral("<eof>")
		} else {
			hx_if_122 = func(hx_obj_119 map[string]any) *string {
				hx_field_120 := hx_obj_119["p"]
				if hx_field_120 == nil {
					var hx_zero_121 *string
					return hx_zero_121
				}
				return hx_field_120.(*string)
			}(field)
		}
		hxrt.Throw(hx_if_122)
		var hx_throw_zero_123 func() any
		return hx_throw_zero_123
	}
	name := haxe__Template_trimExprToken(func(hx_obj_124 map[string]any) *string {
		hx_field_125 := hx_obj_124["p"]
		if hx_field_125 == nil {
			var hx_zero_126 *string
			return hx_zero_126
		}
		return hx_field_125.(*string)
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
		if !haxe__Template_isSpaceOnly(func(hx_obj_127 map[string]any) *string {
			hx_field_128 := hx_obj_127["p"]
			if hx_field_128 == nil {
				var hx_zero_129 *string
				return hx_zero_129
			}
			return hx_field_128.(*string)
		}(cursor.tokens[cursor.index])) {
			return
		}
		cursor.index = int(int32((cursor.index + 1)))
	}
}

func (self *haxe__Template) makeExpr2(cursor *haxe___Template__ExprCursor) func() any {
	self.skipSpaces(cursor)
	var hx_if_131 map[string]any
	if cursor.index >= len(cursor.tokens) {
		hx_if_131 = nil
	} else {
		hx_post_130 := cursor.index
		cursor.index = int(int32((cursor.index + 1)))
		hx_if_131 = cursor.tokens[hx_post_130]
	}
	token := hx_if_131
	self.skipSpaces(cursor)
	if token == nil {
		hxrt.Throw(hxrt.StringFromLiteral("<eof>"))
		var hx_throw_zero_132 func() any
		return hx_throw_zero_132
	}
	if func(hx_obj_133 map[string]any) bool {
		hx_field_134 := hx_obj_133["s"]
		if hx_field_134 == nil {
			var hx_zero_135 bool
			return hx_zero_135
		}
		return hx_field_134.(bool)
	}(token) {
		return self.makeConst(func(hx_obj_136 map[string]any) *string {
			hx_field_137 := hx_obj_136["p"]
			if hx_field_137 == nil {
				var hx_zero_138 *string
				return hx_zero_138
			}
			return hx_field_137.(*string)
		}(token))
	}
	_g := func(hx_obj_139 map[string]any) *string {
		hx_field_140 := hx_obj_139["p"]
		if hx_field_140 == nil {
			var hx_zero_141 *string
			return hx_zero_141
		}
		return hx_field_140.(*string)
	}(token)
	switch _g {
	case hxrt.StringFromLiteral("!"):
		inner := self.makeExpr(cursor)
		return func() any {
			var value any = inner()
			return (hxrt.AnyEqualsNull(value) || (value == false))
		}
	case hxrt.StringFromLiteral("("):
		self.skipSpaces(cursor)
		e1 := self.makeExpr(cursor)
		self.skipSpaces(cursor)
		var hx_if_143 map[string]any
		if cursor.index >= len(cursor.tokens) {
			hx_if_143 = nil
		} else {
			hx_post_142 := cursor.index
			cursor.index = int(int32((cursor.index + 1)))
			hx_if_143 = cursor.tokens[hx_post_142]
		}
		op := hx_if_143
		if (op == nil) || func(hx_obj_144 map[string]any) bool {
			hx_field_145 := hx_obj_144["s"]
			if hx_field_145 == nil {
				var hx_zero_146 bool
				return hx_zero_146
			}
			return hx_field_145.(bool)
		}(op) {
			var hx_if_150 *string
			if op == nil {
				hx_if_150 = hxrt.StringFromLiteral("<eof>")
			} else {
				hx_if_150 = func(hx_obj_147 map[string]any) *string {
					hx_field_148 := hx_obj_147["p"]
					if hx_field_148 == nil {
						var hx_zero_149 *string
						return hx_zero_149
					}
					return hx_field_148.(*string)
				}(op)
			}
			hxrt.Throw(hx_if_150)
			var hx_throw_zero_151 func() any
			return hx_throw_zero_151
		}
		if hxrt.StringEqualStringPtr(func(hx_obj_152 map[string]any) *string {
			hx_field_153 := hx_obj_152["p"]
			if hx_field_153 == nil {
				var hx_zero_154 *string
				return hx_zero_154
			}
			return hx_field_153.(*string)
		}(op), hxrt.StringFromLiteral(")")) {
			return e1
		}
		self.skipSpaces(cursor)
		e2 := self.makeExpr(cursor)
		self.skipSpaces(cursor)
		var hx_if_156 map[string]any
		if cursor.index >= len(cursor.tokens) {
			hx_if_156 = nil
		} else {
			hx_post_155 := cursor.index
			cursor.index = int(int32((cursor.index + 1)))
			hx_if_156 = cursor.tokens[hx_post_155]
		}
		close := hx_if_156
		self.skipSpaces(cursor)
		if (close == nil) || !hxrt.StringEqualStringPtr(func(hx_obj_157 map[string]any) *string {
			hx_field_158 := hx_obj_157["p"]
			if hx_field_158 == nil {
				var hx_zero_159 *string
				return hx_zero_159
			}
			return hx_field_158.(*string)
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
			var hx_throw_zero_164 func() any
			return hx_throw_zero_164
		}
		_g_1 := func(hx_obj_165 map[string]any) *string {
			hx_field_166 := hx_obj_165["p"]
			if hx_field_166 == nil {
				var hx_zero_167 *string
				return hx_zero_167
			}
			return hx_field_166.(*string)
		}(op)
		var hx_switch_168 func() any
		switch _g_1 {
		case hxrt.StringFromLiteral("!="):
			hx_switch_168 = func() any {
				return (e1() != e2())
			}
		case hxrt.StringFromLiteral("&&"):
			hx_switch_168 = func() any {
				return (haxe__Template_valueAsBool(e1()) && haxe__Template_valueAsBool(e2()))
			}
		case hxrt.StringFromLiteral("*"):
			hx_switch_168 = func() any {
				var left any = e1()
				var right any = e2()
				return (haxe__Template_valueAsFloat(left) * haxe__Template_valueAsFloat(right))
			}
		case hxrt.StringFromLiteral("+"):
			hx_switch_168 = func() any {
				return haxe__Template_addValues(e1(), e2())
			}
		case hxrt.StringFromLiteral("-"):
			hx_switch_168 = func() any {
				var left any = e1()
				var right any = e2()
				return (haxe__Template_valueAsFloat(left) - haxe__Template_valueAsFloat(right))
			}
		case hxrt.StringFromLiteral("/"):
			hx_switch_168 = func() any {
				var left any = e1()
				var right any = e2()
				return (haxe__Template_valueAsFloat(left) / haxe__Template_valueAsFloat(right))
			}
		case hxrt.StringFromLiteral("<"):
			hx_switch_168 = func() any {
				return (haxe__Template_compareValues(e1(), e2()) < 0)
			}
		case hxrt.StringFromLiteral("<="):
			hx_switch_168 = func() any {
				return (haxe__Template_compareValues(e1(), e2()) <= 0)
			}
		case hxrt.StringFromLiteral("=="):
			hx_switch_168 = func() any {
				return (e1() == e2())
			}
		case hxrt.StringFromLiteral(">"):
			hx_switch_168 = func() any {
				return (haxe__Template_compareValues(e1(), e2()) > 0)
			}
		case hxrt.StringFromLiteral(">="):
			hx_switch_168 = func() any {
				return (haxe__Template_compareValues(e1(), e2()) >= 0)
			}
		case hxrt.StringFromLiteral("||"):
			hx_switch_168 = func() any {
				return (haxe__Template_valueAsBool(e1()) || haxe__Template_valueAsBool(e2()))
			}
		default:
			hx_switch_168 = func() func() any {
				hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Unknown operation "), func(hx_obj_169 map[string]any) *string {
					hx_field_170 := hx_obj_169["p"]
					if hx_field_170 == nil {
						var hx_zero_171 *string
						return hx_zero_171
					}
					return hx_field_170.(*string)
				}(op)))
				var hx_throw_zero_172 func() any
				return hx_throw_zero_172
			}()
		}
		return hx_switch_168
	case hxrt.StringFromLiteral("-"):
		inner_1 := self.makeExpr(cursor)
		return func() any {
			return -haxe__Template_valueAsFloat(inner_1())
		}
	default:
		hxrt.Throw(func(hx_obj_173 map[string]any) *string {
			hx_field_174 := hx_obj_173["p"]
			if hx_field_174 == nil {
				var hx_zero_175 *string
				return hx_zero_175
			}
			return hx_field_174.(*string)
		}(token))
		var hx_throw_zero_176 func() any
		return hx_throw_zero_176
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
		arrayValues := haxe__Template_anyArrayToSlice(value_1)
		if arrayValues != nil {
			self.stack = append(self.stack, self.context)
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
			var candidate any = Reflect_callMethod(value_1, iteratorField, []any{})
			if !Reflect_hasField(candidate, hxrt.StringFromLiteral("hasNext")) {
				hxrt.Throw(nil)
			}
			iterator = candidate
		}, func(hx_caught_177 any) {
			hx_tmp := hx_caught_177
			_ = hx_tmp
			hxrt.TryCatch(func() {
				if hxrt.AnyEqualsNull(value_1) || !Reflect_hasField(value_1, hxrt.StringFromLiteral("hasNext")) {
					hxrt.Throw(nil)
				}
				iterator = value_1
			}, func(hx_caught_179 any) {
				hx_tmp_1 := hx_caught_179
				_ = hx_tmp_1
				hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Cannot iter on "), hxrt.StdString(value_1)))
			})
		})
		self.stack = append(self.stack, self.context)
		iterable := func(hx_value_181 any) map[string]any {
			if hx_value_181 == nil {
				var hx_zero_182 map[string]any
				return hx_zero_182
			}
			return hx_value_181.(map[string]any)
		}(iterator)
		ctx_1 := iterable
		for func(hx_obj_183 map[string]any) func() bool {
			hx_field_184 := hx_obj_183["hasNext"]
			if hx_field_184 == nil {
				var hx_zero_185 func() bool
				return hx_zero_185
			}
			return hx_field_184.(func() bool)
		}(ctx_1)() {
			var ctx_2 any = func(hx_obj_186 map[string]any) func() any {
				hx_field_187 := hx_obj_186["next"]
				if hx_field_187 == nil {
					var hx_zero_188 func() any
					return hx_zero_188
				}
				return hx_field_187.(func() any)
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
			self.output = hxrt.StringConcatStringPtr(self.output, hxrt.StdString(Reflect_callMethod(self.macros, fn, callArgs)))
		}, func(hx_caught_189 any) {
			err := hx_caught_189
			var hx_try_191 *string
			hxrt.TryCatch(func() {
				hx_try_191 = haxe__Template_joinDynamicArgs(callArgs)
			}, func(hx_caught_192 any) {
				hx_tmp_2 := hx_caught_192
				_ = hx_tmp_2
				hx_try_191 = hxrt.StringFromLiteral("???")
			})
			argsText := hx_try_191
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
		hx_post_194 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_194
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

func haxe__Template_anyArrayToSlice(value any) []any {
	return haxe__Template_anyArrayToSlice_runtime(value)
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
		var hx_if_196 int
		if leftFloat < rightFloat {
			hx_if_196 = -1
		} else {
			var hx_if_195 int
			if leftFloat > rightFloat {
				hx_if_195 = 1
			} else {
				hx_if_195 = 0
			}
			hx_if_196 = hx_if_195
		}
		return hx_if_196
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
	hx_obj_197 := map[string]any{}
	return hx_obj_197
}())

func haxe__Template_isSpaceOnly(value *string) bool {
	var _g_s *string
	var _g_offset int
	_g_offset = 0
	_g_s = value
	for _g_offset < hxrt.StringLengthStringPtr(_g_s) {
		code := hxrt.StringCharCodeAtStringPtr(_g_s, func() int {
			hx_post_198 := _g_offset
			_g_offset = int(int32((_g_offset + 1)))
			return hx_post_198
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
		hx_post_199 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_199
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
				hx_post_200 := _g_offset
				_g_offset = int(int32((_g_offset + 1)))
				return hx_post_200
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
		exponentCode := hxrt.StringCharCodeAtAnyStringPtr(normalized, index)
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
				var hx_if_201 float64
				if exponentSign < 0 {
					hx_if_201 = (result / 10.0)
				} else {
					hx_if_201 = (result * 10.0)
				}
				result = hx_if_201
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
	var hx_if_202 map[string]any
	if cursor.index < len(cursor.tokens) {
		hx_if_202 = cursor.tokens[cursor.index]
	} else {
		hx_if_202 = nil
	}
	return hx_if_202
}

func haxe__Template_peekToken(cursor *haxe___Template__TokenCursor) map[string]any {
	var hx_if_203 map[string]any
	if cursor.index < len(cursor.tokens) {
		hx_if_203 = cursor.tokens[cursor.index]
	} else {
		hx_if_203 = nil
	}
	return hx_if_203
}

func haxe__Template_popExprToken(cursor *haxe___Template__ExprCursor) map[string]any {
	if cursor.index >= len(cursor.tokens) {
		return nil
	}
	hx_post_204 := cursor.index
	cursor.index = int(int32((cursor.index + 1)))
	return cursor.tokens[hx_post_204]
}

func haxe__Template_popToken(cursor *haxe___Template__TokenCursor) map[string]any {
	if cursor.index >= len(cursor.tokens) {
		return nil
	}
	hx_post_205 := cursor.index
	cursor.index = int(int32((cursor.index + 1)))
	return cursor.tokens[hx_post_205]
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
		return func(hx_value_206 any) float64 {
			if hx_value_206 == nil {
				var hx_zero_207 float64
				return hx_zero_207
			}
			return hx_value_206.(float64)
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
		return haxe__Template_parseFloatLiteral(func(hx_value_208 any) *string {
			if hx_value_208 == nil {
				var hx_zero_209 *string
				return hx_zero_209
			}
			return hx_value_208.(*string)
		}(value))
	}
	hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Expected numeric expression value, got "), hxrt.StdString(value)))
	var hx_throw_zero_210 float64
	return hx_throw_zero_210
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
