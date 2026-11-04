package main

import "snapshot/hxrt"

type I_haxe___Template__TokenCursor interface {
}

type haxe___Template__TokenCursor struct {
	__hx_this I_haxe___Template__TokenCursor
	tokens    *hxrt.Array
	index     int
}

func (self *haxe___Template__TokenCursor) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe___Template__TokenCursor(tokens *hxrt.Array) *haxe___Template__TokenCursor {
	return hx_init_haxe___Template__TokenCursor(hx_alloc_haxe___Template__TokenCursor(nil), tokens)
}

func hx_alloc_haxe___Template__TokenCursor(owner any) *haxe___Template__TokenCursor {
	self := &haxe___Template__TokenCursor{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_haxe___Template__TokenCursor)
	}
	return self
}

func hx_init_haxe___Template__TokenCursor(self *haxe___Template__TokenCursor, tokens *hxrt.Array) *haxe___Template__TokenCursor {
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

func (self *haxe___Template__ExprCursor) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe___Template__ExprCursor(tokens *hxrt.Array) *haxe___Template__ExprCursor {
	return hx_init_haxe___Template__ExprCursor(hx_alloc_haxe___Template__ExprCursor(nil), tokens)
}

func hx_alloc_haxe___Template__ExprCursor(owner any) *haxe___Template__ExprCursor {
	self := &haxe___Template__ExprCursor{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_haxe___Template__ExprCursor)
	}
	return self
}

func hx_init_haxe___Template__ExprCursor(self *haxe___Template__ExprCursor, tokens *hxrt.Array) *haxe___Template__ExprCursor {
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
	parseExpr(data *string) *hxrt.Function[func() any]
	makeConst(v *string) *hxrt.Function[func() any]
	makePath(e *hxrt.Function[func() any], cursor *haxe___Template__ExprCursor) *hxrt.Function[func() any]
	makeExpr(cursor *haxe___Template__ExprCursor) *hxrt.Function[func() any]
	skipSpaces(cursor *haxe___Template__ExprCursor)
	makeExpr2(cursor *haxe___Template__ExprCursor) *hxrt.Function[func() any]
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

func (self *haxe__Template) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe__Template(str *string) *haxe__Template {
	return hx_init_haxe__Template(hx_alloc_haxe__Template(nil), str)
}

func hx_alloc_haxe__Template(owner any) *haxe__Template {
	self := &haxe__Template{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_haxe__Template)
	}
	return self
}

func hx_init_haxe__Template(self *haxe__Template, str *string) *haxe__Template {
	cursor := New_haxe___Template__TokenCursor(self.__hx_this.parseTokens(str))
	self.expr = self.__hx_this.parseBlock(cursor)
	if cursor.index < cursor.tokens.Len() {
		token := hx_structural_recover(cursor.tokens.Get(cursor.index))
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Unexpected '"), hxrt.StdString(func(hx_value_1 any) bool {
			if hx_value_1 == nil {
				var hx_zero_2 bool
				return hx_zero_2
			}
			return hx_value_1.(bool)
		}(hxrt.AnonymousGet(token, "s")))), hxrt.StringFromLiteral("'"))))
	}
	return self
}

func (self *haxe__Template) execute(context any, macros any) *string {
	var hx_if_4 any
	if hxrt.AnyEqualsNull(macros) {
		hx_obj_3 := map[string]any{}
		hx_if_4 = hxrt.AnonymousFromMap(hx_obj_3)
	} else {
		hx_if_4 = macros
	}
	self.macros = hx_if_4
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
		if func(hx_value_9 any) int {
			if hx_value_9 == nil {
				var hx_zero_10 int
				return hx_zero_10
			}
			return hx_value_9.(int)
		}(hxrt.AnonymousGet(p, "pos")) > 0 {
			hx_obj_6 := map[string]any{}
			hx_obj_6["p"] = hxrt.StringSubstrStringPtr(data, 0, func(hx_value_7 any) int {
				if hx_value_7 == nil {
					var hx_zero_8 int
					return hx_zero_8
				}
				return hx_value_7.(int)
			}(hxrt.AnonymousGet(p, "pos")), true)
			hx_obj_6["s"] = true
			hx_obj_6["l"] = nil
			tokens.Push(hxrt.AnonymousFromMap(hx_obj_6))
		}
		if hxrt.StringCharCodeAtAnyStringPtr(data, func(hx_value_17 any) int {
			if hx_value_17 == nil {
				var hx_zero_18 int
				return hx_zero_18
			}
			return hx_value_17.(int)
		}(hxrt.AnonymousGet(p, "pos"))) == 58 {
			hx_obj_12 := map[string]any{}
			hx_obj_12["p"] = hxrt.StringSubstrStringPtr(data, int((hxrt.Int32Wrap(func(hx_value_13 any) int {
				if hx_value_13 == nil {
					var hx_zero_14 int
					return hx_zero_14
				}
				return hx_value_13.(int)
			}(hxrt.AnonymousGet(p, "pos"))) + hxrt.Int32Wrap(2))), int((hxrt.Int32Wrap(func(hx_value_15 any) int {
				if hx_value_15 == nil {
					var hx_zero_16 int
					return hx_zero_16
				}
				return hx_value_15.(int)
			}(hxrt.AnonymousGet(p, "len"))) - hxrt.Int32Wrap(4))), true)
			hx_obj_12["s"] = false
			hx_obj_12["l"] = nil
			tokens.Push(hxrt.AnonymousFromMap(hx_obj_12))
			data = haxe__Template_splitter.__hx_this.matchedRight()
			continue
		}
		parp := int((hxrt.Int32Wrap(func(hx_value_19 any) int {
			if hx_value_19 == nil {
				var hx_zero_20 int
				return hx_zero_20
			}
			return hx_value_19.(int)
		}(hxrt.AnonymousGet(p, "pos"))) + hxrt.Int32Wrap(func(hx_value_21 any) int {
			if hx_value_21 == nil {
				var hx_zero_22 int
				return hx_zero_22
			}
			return hx_value_21.(int)
		}(hxrt.AnonymousGet(p, "len")))))
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
						hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("Unclosed macro parenthesis")))
					}
				}
			}
			chunk := hxrt.StringSubstrStringPtr(data, int((hxrt.Int32Wrap(parp) - hxrt.Int32Wrap(1))), 1, true)
			if (c == 44) && (npar == 1) {
				params.Push(part)
				part = hxrt.StringFromLiteral("")
			} else {
				part = hxrt.StringConcatStringPtr(part, chunk)
			}
		}
		params.Push(part)
		hx_obj_26 := map[string]any{}
		hx_obj_26["p"] = haxe__Template_splitter.__hx_this.matched(2)
		hx_obj_26["s"] = false
		hx_obj_26["l"] = params
		tokens.Push(hxrt.AnonymousFromMap(hx_obj_26))
		data = hxrt.StringSubstrStringPtr(data, parp, int((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(data)) - hxrt.Int32Wrap(parp))), true)
	}
	if hxrt.StringLengthStringPtr(data) > 0 {
		hx_obj_28 := map[string]any{}
		hx_obj_28["p"] = data
		hx_obj_28["s"] = true
		hx_obj_28["l"] = nil
		tokens.Push(hxrt.AnonymousFromMap(hx_obj_28))
	}
	return tokens
}

func (self *haxe__Template) parseBlock(cursor *haxe___Template__TokenCursor) *haxe___Template__TemplateExpr {
	items := hxrt.NewArray()
	for cursor.index < cursor.tokens.Len() {
		t := hx_structural_recover(cursor.tokens.Get(cursor.index))
		if !func(hx_value_29 any) bool {
			if hx_value_29 == nil {
				var hx_zero_30 bool
				return hx_zero_30
			}
			return hx_value_29.(bool)
		}(hxrt.AnonymousGet(t, "s")) && ((hxrt.StringEqualStringPtr(func(hx_value_31 any) *string {
			if hx_value_31 == nil {
				var hx_zero_32 *string
				return hx_zero_32
			}
			return hx_value_31.(*string)
		}(hxrt.AnonymousGet(t, "p")), hxrt.StringFromLiteral("end")) || hxrt.StringEqualStringPtr(func(hx_value_33 any) *string {
			if hx_value_33 == nil {
				var hx_zero_34 *string
				return hx_zero_34
			}
			return hx_value_33.(*string)
		}(hxrt.AnonymousGet(t, "p")), hxrt.StringFromLiteral("else"))) || hxrt.StringEqualStringPtr(hxrt.StringSubstrStringPtr(func(hx_value_35 any) *string {
			if hx_value_35 == nil {
				var hx_zero_36 *string
				return hx_zero_36
			}
			return hx_value_35.(*string)
		}(hxrt.AnonymousGet(t, "p")), 0, 7, true), hxrt.StringFromLiteral("elseif "))) {
			break
		}
		items.Push(self.__hx_this.parse(cursor))
	}
	if items.Len() == 1 {
		return func(hx_value_38 any) *haxe___Template__TemplateExpr {
			if hx_value_38 == nil {
				var hx_zero_39 *haxe___Template__TemplateExpr
				return hx_zero_39
			}
			return hx_value_38.(*haxe___Template__TemplateExpr)
		}(items.Get(0))
	}
	return haxe___Template__TemplateExpr_OpBlock(items)
}

func (self *haxe__Template) parse(cursor *haxe___Template__TokenCursor) *haxe___Template__TemplateExpr {
	var hx_if_41 hxrt.Anonymous
	if cursor.index >= cursor.tokens.Len() {
		hx_if_41 = nil
	} else {
		hx_post_40 := cursor.index
		cursor.index = int(int32((cursor.index + 1)))
		hx_if_41 = hx_structural_recover(hx_structural_recover(cursor.tokens.Get(hx_post_40)))
	}
	t := hx_if_41
	if t == nil {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("Unexpected <eof>")))
	}
	p := func(hx_value_42 any) *string {
		if hx_value_42 == nil {
			var hx_zero_43 *string
			return hx_zero_43
		}
		return hx_value_42.(*string)
	}(hxrt.AnonymousGet(t, "p"))
	if func(hx_value_44 any) bool {
		if hx_value_44 == nil {
			var hx_zero_45 bool
			return hx_zero_45
		}
		return hx_value_44.(bool)
	}(hxrt.AnonymousGet(t, "s")) {
		return haxe___Template__TemplateExpr_OpStr(p)
	}
	if func(hx_value_51 any) *hxrt.Array {
		if hx_value_51 == nil {
			var hx_zero_52 *hxrt.Array
			return hx_zero_52
		}
		return hx_value_51.(*hxrt.Array)
	}(hxrt.AnonymousGet(t, "l")) != nil {
		parsedParams := hxrt.NewArray()
		_g := 0
		_g1 := func(hx_value_46 any) *hxrt.Array {
			if hx_value_46 == nil {
				var hx_zero_47 *hxrt.Array
				return hx_zero_47
			}
			return hx_value_46.(*hxrt.Array)
		}(hxrt.AnonymousGet(t, "l"))
		for _g < _g1.Len() {
			param := func(hx_value_48 any) *string {
				if hx_value_48 == nil {
					var hx_zero_49 *string
					return hx_zero_49
				}
				return hx_value_48.(*string)
			}(_g1.Get(_g))
			_g = int(int32((_g + 1)))
			parsedParams.Push(self.__hx_this.parseBlock(New_haxe___Template__TokenCursor(self.__hx_this.parseTokens(param))))
		}
		return haxe___Template__TemplateExpr_OpMacro(p, parsedParams)
	}
	pos := haxe__Template_kwdEnd(p, hxrt.StringFromLiteral("if"))
	if pos > 0 {
		p = hxrt.StringSubstrStringPtr(p, pos, int((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(p)) - hxrt.Int32Wrap(pos))), true)
		e := self.__hx_this.parseExpr(p)
		eif := self.__hx_this.parseBlock(cursor)
		var hx_if_53 hxrt.Anonymous
		if cursor.index < cursor.tokens.Len() {
			hx_if_53 = hx_structural_recover(hx_structural_recover(cursor.tokens.Get(cursor.index)))
		} else {
			hx_if_53 = nil
		}
		nextToken := hx_if_53
		if nextToken == nil {
			hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("Unclosed 'if'")))
		}
		var eelse *haxe___Template__TemplateExpr = nil
		if hxrt.StringEqualStringPtr(func(hx_value_69 any) *string {
			if hx_value_69 == nil {
				var hx_zero_70 *string
				return hx_zero_70
			}
			return hx_value_69.(*string)
		}(hxrt.AnonymousGet(nextToken, "p")), hxrt.StringFromLiteral("end")) {
			if cursor.index >= cursor.tokens.Len() {
			} else {
				cursor.tokens.Get(func() int {
					hx_post_54 := cursor.index
					cursor.index = int(int32((cursor.index + 1)))
					return hx_post_54
				}())
			}
		} else {
			if hxrt.StringEqualStringPtr(func(hx_value_67 any) *string {
				if hx_value_67 == nil {
					var hx_zero_68 *string
					return hx_zero_68
				}
				return hx_value_67.(*string)
			}(hxrt.AnonymousGet(nextToken, "p")), hxrt.StringFromLiteral("else")) {
				if cursor.index >= cursor.tokens.Len() {
				} else {
					cursor.tokens.Get(func() int {
						hx_post_55 := cursor.index
						cursor.index = int(int32((cursor.index + 1)))
						return hx_post_55
					}())
				}
				eelse = self.__hx_this.parseBlock(cursor)
				var hx_if_57 hxrt.Anonymous
				if cursor.index >= cursor.tokens.Len() {
					hx_if_57 = nil
				} else {
					hx_post_56 := cursor.index
					cursor.index = int(int32((cursor.index + 1)))
					hx_if_57 = hx_structural_recover(hx_structural_recover(cursor.tokens.Get(hx_post_56)))
				}
				endToken := hx_if_57
				if (endToken == nil) || !hxrt.StringEqualStringPtr(func(hx_value_58 any) *string {
					if hx_value_58 == nil {
						var hx_zero_59 *string
						return hx_zero_59
					}
					return hx_value_58.(*string)
				}(hxrt.AnonymousGet(endToken, "p")), hxrt.StringFromLiteral("end")) {
					hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("Unclosed 'else'")))
				}
			} else {
				func() *string {
					hx_structural_target_60 := nextToken
					hx_structural_value_62 := hxrt.StringSubstrStringPtr(func(hx_value_63 any) *string {
						if hx_value_63 == nil {
							var hx_zero_64 *string
							return hx_zero_64
						}
						return hx_value_63.(*string)
					}(hxrt.AnonymousGet(nextToken, "p")), 4, int((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(func(hx_value_65 any) *string {
						if hx_value_65 == nil {
							var hx_zero_66 *string
							return hx_zero_66
						}
						return hx_value_65.(*string)
					}(hxrt.AnonymousGet(nextToken, "p")))) - hxrt.Int32Wrap(4))), true)
					hxrt.AnonymousSet(hx_structural_target_60, "p", hx_structural_value_62)
					return hx_structural_value_62
				}()
				eelse = self.__hx_this.parse(cursor)
			}
		}
		return haxe___Template__TemplateExpr_OpIf(e, eif, eelse)
	}
	pos = haxe__Template_kwdEnd(p, hxrt.StringFromLiteral("foreach"))
	if pos >= 0 {
		p = hxrt.StringSubstrStringPtr(p, pos, int((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(p)) - hxrt.Int32Wrap(pos))), true)
		e_1 := self.__hx_this.parseExpr(p)
		efor := self.__hx_this.parseBlock(cursor)
		var hx_if_72 hxrt.Anonymous
		if cursor.index >= cursor.tokens.Len() {
			hx_if_72 = nil
		} else {
			hx_post_71 := cursor.index
			cursor.index = int(int32((cursor.index + 1)))
			hx_if_72 = hx_structural_recover(hx_structural_recover(cursor.tokens.Get(hx_post_71)))
		}
		endToken_1 := hx_if_72
		if (endToken_1 == nil) || !hxrt.StringEqualStringPtr(func(hx_value_73 any) *string {
			if hx_value_73 == nil {
				var hx_zero_74 *string
				return hx_zero_74
			}
			return hx_value_73.(*string)
		}(hxrt.AnonymousGet(endToken_1, "p")), hxrt.StringFromLiteral("end")) {
			hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("Unclosed 'foreach'")))
		}
		return haxe___Template__TemplateExpr_OpForeach(e_1, efor)
	}
	if haxe__Template_expr_splitter.__hx_this.match(p) {
		return haxe___Template__TemplateExpr_OpExpr(self.__hx_this.parseExpr(p))
	}
	return haxe___Template__TemplateExpr_OpVar(p)
}

func (self *haxe__Template) parseExpr(data *string) *hxrt.Function[func() any] {
	tokens := hxrt.NewArray()
	expr := data
	for haxe__Template_expr_splitter.__hx_this.match(data) {
		p := haxe__Template_expr_splitter.__hx_this.matchedPos()
		if func(hx_value_79 any) int {
			if hx_value_79 == nil {
				var hx_zero_80 int
				return hx_zero_80
			}
			return hx_value_79.(int)
		}(hxrt.AnonymousGet(p, "pos")) != 0 {
			hx_obj_76 := map[string]any{}
			hx_obj_76["p"] = hxrt.StringSubstrStringPtr(data, 0, func(hx_value_77 any) int {
				if hx_value_77 == nil {
					var hx_zero_78 int
					return hx_zero_78
				}
				return hx_value_77.(int)
			}(hxrt.AnonymousGet(p, "pos")), true)
			hx_obj_76["s"] = true
			tokens.Push(hxrt.AnonymousFromMap(hx_obj_76))
		}
		token := haxe__Template_expr_splitter.__hx_this.matched(0)
		hx_obj_82 := map[string]any{}
		hx_obj_82["p"] = token
		hx_obj_82["s"] = StringTools_contains(token, hxrt.StringFromLiteral("\""))
		tokens.Push(hxrt.AnonymousFromMap(hx_obj_82))
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
				hx_obj_84 := map[string]any{}
				hx_obj_84["p"] = hxrt.StringSubstrStringPtr(data, i, 0, false)
				hx_obj_84["s"] = true
				tokens.Push(hxrt.AnonymousFromMap(hx_obj_84))
				break
			}
		}
	}
	cursor := New_haxe___Template__ExprCursor(tokens)
	var built *hxrt.Function[func() any]
	hxrt.TryCatch(func() {
		built = self.__hx_this.makeExpr(cursor)
		if cursor.index < cursor.tokens.Len() {
			hxrt.Throw(haxe__Exception_thrown(func(hx_value_90 any) *string {
				if hx_value_90 == nil {
					var hx_zero_91 *string
					return hx_zero_91
				}
				return hx_value_90.(*string)
			}(hxrt.AnonymousGet(hx_structural_recover(cursor.tokens.Get(cursor.index)), "p"))))
		}
	}, func(hx_caught_85 any) {
		hx_exception_86 := haxe__Exception_caught(hx_caught_85)
		hx_payload_87 := haxe__Exception_catchValue(hx_caught_85)
		_ = hx_exception_86
		if func(hx_value any) bool {
			switch hx_value.(type) {
			case *string:
				return true
			default:
				return false
			}
		}(hx_payload_87) {
			var s *string = func(hx_value_88 any) *string {
				if hx_value_88 == nil {
					var hx_zero_89 *string
					return hx_zero_89
				}
				return hx_value_88.(*string)
			}(hx_payload_87)
			hxrt.Throw(haxe__Exception_thrown(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Unexpected '"), s), hxrt.StringFromLiteral("' in ")), expr)))
		} else {
			hxrt.Throw(hx_caught_85)
		}
	})
	me := self
	_ = me
	wrapped := hxrt.NewFunction(func() any {
		hx_try_return_92 := false
		var hx_try_value_93 any
		hxrt.TryCatch(func() {
			hx_try_value_93 = func() any {
				var hx_callable_97 *hxrt.Function[func() any] = built
				return hxrt.FunctionCall(hx_callable_97)()
			}()
			hx_try_return_92 = true
			return
		}, func(hx_caught_94 any) {
			hx_exception_95 := haxe__Exception_caught(hx_caught_94)
			hx_payload_96 := haxe__Exception_catchValue(hx_caught_94)
			_ = hx_exception_95
			var exc any = hx_payload_96
			hxrt.Throw(haxe__Exception_thrown(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Error : "), hxrt.StdString(exc)), hxrt.StringFromLiteral(" in ")), expr)))
		})
		if hx_try_return_92 {
			return hx_try_value_93
		}
		return nil
	})
	return wrapped
}

func (self *haxe__Template) makeConst(v *string) *hxrt.Function[func() any] {
	haxe__Template_expr_trim.__hx_this.match(v)
	v = haxe__Template_expr_trim.__hx_this.matched(1)
	if hxrt.StringCharCodeAtAnyStringPtr(v, 0) == 34 {
		str := hxrt.StringSubstrStringPtr(v, 1, int((hxrt.Int32Wrap(hxrt.StringLengthStringPtr(v)) - hxrt.Int32Wrap(2))), true)
		literal := hxrt.NewFunction(func() any {
			return str
		})
		return literal
	}
	if haxe__Template_expr_int.__hx_this.match(v) {
		i := haxe__Template_parseIntLiteral(v)
		intLiteral := hxrt.NewFunction(func() any {
			return i
		})
		return intLiteral
	}
	if haxe__Template_expr_float.__hx_this.match(v) {
		f := haxe__Template_parseFloatLiteral(v)
		floatLiteral := hxrt.NewFunction(func() any {
			return f
		})
		return floatLiteral
	}
	me := self
	resolved := hxrt.NewFunction(func() any {
		return me.__hx_this.resolve(v)
	})
	return resolved
}

func (self *haxe__Template) makePath(e *hxrt.Function[func() any], cursor *haxe___Template__ExprCursor) *hxrt.Function[func() any] {
	var hx_if_98 hxrt.Anonymous
	if cursor.index < cursor.tokens.Len() {
		hx_if_98 = hx_structural_recover(hx_structural_recover(cursor.tokens.Get(cursor.index)))
	} else {
		hx_if_98 = nil
	}
	token := hx_if_98
	if (token == nil) || !hxrt.StringEqualStringPtr(func(hx_value_99 any) *string {
		if hx_value_99 == nil {
			var hx_zero_100 *string
			return hx_zero_100
		}
		return hx_value_99.(*string)
	}(hxrt.AnonymousGet(token, "p")), hxrt.StringFromLiteral(".")) {
		return e
	}
	if cursor.index >= cursor.tokens.Len() {
	} else {
		cursor.tokens.Get(func() int {
			hx_post_101 := cursor.index
			cursor.index = int(int32((cursor.index + 1)))
			return hx_post_101
		}())
	}
	var hx_if_103 hxrt.Anonymous
	if cursor.index >= cursor.tokens.Len() {
		hx_if_103 = nil
	} else {
		hx_post_102 := cursor.index
		cursor.index = int(int32((cursor.index + 1)))
		hx_if_103 = hx_structural_recover(hx_structural_recover(cursor.tokens.Get(hx_post_102)))
	}
	field := hx_if_103
	if (field == nil) || !func(hx_value_107 any) bool {
		if hx_value_107 == nil {
			var hx_zero_108 bool
			return hx_zero_108
		}
		return hx_value_107.(bool)
	}(hxrt.AnonymousGet(field, "s")) {
		var hx_if_106 *string
		if field == nil {
			hx_if_106 = hxrt.StringFromLiteral("<eof>")
		} else {
			hx_if_106 = func(hx_value_104 any) *string {
				if hx_value_104 == nil {
					var hx_zero_105 *string
					return hx_zero_105
				}
				return hx_value_104.(*string)
			}(hxrt.AnonymousGet(field, "p"))
		}
		hxrt.Throw(haxe__Exception_thrown(hx_if_106))
	}
	name := haxe__Template_trimExprToken(func(hx_value_109 any) *string {
		if hx_value_109 == nil {
			var hx_zero_110 *string
			return hx_zero_110
		}
		return hx_value_109.(*string)
	}(hxrt.AnonymousGet(field, "p")))
	return self.__hx_this.makePath(hxrt.NewFunction(func() any {
		return Reflect_field(func() any {
			var hx_callable_111 *hxrt.Function[func() any] = e
			return hxrt.FunctionCall(hx_callable_111)()
		}(), name)
	}), cursor)
}

func (self *haxe__Template) makeExpr(cursor *haxe___Template__ExprCursor) *hxrt.Function[func() any] {
	return self.__hx_this.makePath(self.__hx_this.makeExpr2(cursor), cursor)
}

func (self *haxe__Template) skipSpaces(cursor *haxe___Template__ExprCursor) {
	for cursor.index < cursor.tokens.Len() {
		if !haxe__Template_isSpaceOnly(func(hx_value_112 any) *string {
			if hx_value_112 == nil {
				var hx_zero_113 *string
				return hx_zero_113
			}
			return hx_value_112.(*string)
		}(hxrt.AnonymousGet(hx_structural_recover(cursor.tokens.Get(cursor.index)), "p"))) {
			return
		}
		cursor.index = int(int32((cursor.index + 1)))
	}
}

func (self *haxe__Template) makeExpr2(cursor *haxe___Template__ExprCursor) *hxrt.Function[func() any] {
	self.__hx_this.skipSpaces(cursor)
	var hx_if_115 hxrt.Anonymous
	if cursor.index >= cursor.tokens.Len() {
		hx_if_115 = nil
	} else {
		hx_post_114 := cursor.index
		cursor.index = int(int32((cursor.index + 1)))
		hx_if_115 = hx_structural_recover(hx_structural_recover(cursor.tokens.Get(hx_post_114)))
	}
	token := hx_if_115
	self.__hx_this.skipSpaces(cursor)
	if token == nil {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("<eof>")))
	}
	if func(hx_value_118 any) bool {
		if hx_value_118 == nil {
			var hx_zero_119 bool
			return hx_zero_119
		}
		return hx_value_118.(bool)
	}(hxrt.AnonymousGet(token, "s")) {
		return self.__hx_this.makeConst(func(hx_value_116 any) *string {
			if hx_value_116 == nil {
				var hx_zero_117 *string
				return hx_zero_117
			}
			return hx_value_116.(*string)
		}(hxrt.AnonymousGet(token, "p")))
	}
	_g := func(hx_value_120 any) *string {
		if hx_value_120 == nil {
			var hx_zero_121 *string
			return hx_zero_121
		}
		return hx_value_120.(*string)
	}(hxrt.AnonymousGet(token, "p"))
	switch *hxrt.StdString(_g) {
	case *hxrt.StdString(hxrt.StringFromLiteral("!")):
		inner := self.__hx_this.makeExpr(cursor)
		return hxrt.NewFunction(func() any {
			var value any = func() any {
				var hx_callable_122 *hxrt.Function[func() any] = inner
				return hxrt.FunctionCall(hx_callable_122)()
			}()
			return (hxrt.AnyEqualsNull(value) || hxrt.HaxeEqual(value, false))
		})
	case *hxrt.StdString(hxrt.StringFromLiteral("(")):
		self.__hx_this.skipSpaces(cursor)
		e1 := self.__hx_this.makeExpr(cursor)
		self.__hx_this.skipSpaces(cursor)
		var hx_if_124 hxrt.Anonymous
		if cursor.index >= cursor.tokens.Len() {
			hx_if_124 = nil
		} else {
			hx_post_123 := cursor.index
			cursor.index = int(int32((cursor.index + 1)))
			hx_if_124 = hx_structural_recover(hx_structural_recover(cursor.tokens.Get(hx_post_123)))
		}
		op := hx_if_124
		if (op == nil) || func(hx_value_128 any) bool {
			if hx_value_128 == nil {
				var hx_zero_129 bool
				return hx_zero_129
			}
			return hx_value_128.(bool)
		}(hxrt.AnonymousGet(op, "s")) {
			var hx_if_127 *string
			if op == nil {
				hx_if_127 = hxrt.StringFromLiteral("<eof>")
			} else {
				hx_if_127 = func(hx_value_125 any) *string {
					if hx_value_125 == nil {
						var hx_zero_126 *string
						return hx_zero_126
					}
					return hx_value_125.(*string)
				}(hxrt.AnonymousGet(op, "p"))
			}
			hxrt.Throw(haxe__Exception_thrown(hx_if_127))
		}
		if hxrt.StringEqualStringPtr(func(hx_value_130 any) *string {
			if hx_value_130 == nil {
				var hx_zero_131 *string
				return hx_zero_131
			}
			return hx_value_130.(*string)
		}(hxrt.AnonymousGet(op, "p")), hxrt.StringFromLiteral(")")) {
			return e1
		}
		self.__hx_this.skipSpaces(cursor)
		e2 := self.__hx_this.makeExpr(cursor)
		self.__hx_this.skipSpaces(cursor)
		var hx_if_133 hxrt.Anonymous
		if cursor.index >= cursor.tokens.Len() {
			hx_if_133 = nil
		} else {
			hx_post_132 := cursor.index
			cursor.index = int(int32((cursor.index + 1)))
			hx_if_133 = hx_structural_recover(hx_structural_recover(cursor.tokens.Get(hx_post_132)))
		}
		close := hx_if_133
		self.__hx_this.skipSpaces(cursor)
		if (close == nil) || !hxrt.StringEqualStringPtr(func(hx_value_137 any) *string {
			if hx_value_137 == nil {
				var hx_zero_138 *string
				return hx_zero_138
			}
			return hx_value_137.(*string)
		}(hxrt.AnonymousGet(close, "p")), hxrt.StringFromLiteral(")")) {
			var hx_if_136 *string
			if close == nil {
				hx_if_136 = hxrt.StringFromLiteral("<eof>")
			} else {
				hx_if_136 = func(hx_value_134 any) *string {
					if hx_value_134 == nil {
						var hx_zero_135 *string
						return hx_zero_135
					}
					return hx_value_134.(*string)
				}(hxrt.AnonymousGet(close, "p"))
			}
			hxrt.Throw(haxe__Exception_thrown(hx_if_136))
		}
		_g_1 := func(hx_value_139 any) *string {
			if hx_value_139 == nil {
				var hx_zero_140 *string
				return hx_zero_140
			}
			return hx_value_139.(*string)
		}(hxrt.AnonymousGet(op, "p"))
		var hx_switch_141 *hxrt.Function[func() any]
		switch *hxrt.StdString(_g_1) {
		case *hxrt.StdString(hxrt.StringFromLiteral("!=")):
			hx_switch_141 = hxrt.NewFunction(func() any {
				return !hxrt.HaxeEqual(func() any {
					var hx_callable_142 *hxrt.Function[func() any] = e1
					return hxrt.FunctionCall(hx_callable_142)()
				}(), func() any {
					var hx_callable_143 *hxrt.Function[func() any] = e2
					return hxrt.FunctionCall(hx_callable_143)()
				}())
			})
		case *hxrt.StdString(hxrt.StringFromLiteral("&&")):
			hx_switch_141 = hxrt.NewFunction(func() any {
				return (haxe__Template_valueAsBool(func() any {
					var hx_callable_144 *hxrt.Function[func() any] = e1
					return hxrt.FunctionCall(hx_callable_144)()
				}()) && haxe__Template_valueAsBool(func() any {
					var hx_callable_145 *hxrt.Function[func() any] = e2
					return hxrt.FunctionCall(hx_callable_145)()
				}()))
			})
		case *hxrt.StdString(hxrt.StringFromLiteral("*")):
			hx_switch_141 = hxrt.NewFunction(func() any {
				var left any = func() any {
					var hx_callable_146 *hxrt.Function[func() any] = e1
					return hxrt.FunctionCall(hx_callable_146)()
				}()
				var right any = func() any {
					var hx_callable_147 *hxrt.Function[func() any] = e2
					return hxrt.FunctionCall(hx_callable_147)()
				}()
				return (haxe__Template_valueAsFloat(left) * haxe__Template_valueAsFloat(right))
			})
		case *hxrt.StdString(hxrt.StringFromLiteral("+")):
			hx_switch_141 = hxrt.NewFunction(func() any {
				return haxe__Template_addValues(func() any {
					var hx_callable_148 *hxrt.Function[func() any] = e1
					return hxrt.FunctionCall(hx_callable_148)()
				}(), func() any {
					var hx_callable_149 *hxrt.Function[func() any] = e2
					return hxrt.FunctionCall(hx_callable_149)()
				}())
			})
		case *hxrt.StdString(hxrt.StringFromLiteral("-")):
			hx_switch_141 = hxrt.NewFunction(func() any {
				var left any = func() any {
					var hx_callable_150 *hxrt.Function[func() any] = e1
					return hxrt.FunctionCall(hx_callable_150)()
				}()
				var right any = func() any {
					var hx_callable_151 *hxrt.Function[func() any] = e2
					return hxrt.FunctionCall(hx_callable_151)()
				}()
				return (haxe__Template_valueAsFloat(left) - haxe__Template_valueAsFloat(right))
			})
		case *hxrt.StdString(hxrt.StringFromLiteral("/")):
			hx_switch_141 = hxrt.NewFunction(func() any {
				var left any = func() any {
					var hx_callable_152 *hxrt.Function[func() any] = e1
					return hxrt.FunctionCall(hx_callable_152)()
				}()
				var right any = func() any {
					var hx_callable_153 *hxrt.Function[func() any] = e2
					return hxrt.FunctionCall(hx_callable_153)()
				}()
				return (haxe__Template_valueAsFloat(left) / haxe__Template_valueAsFloat(right))
			})
		case *hxrt.StdString(hxrt.StringFromLiteral("<")):
			hx_switch_141 = hxrt.NewFunction(func() any {
				return (haxe__Template_compareValues(func() any {
					var hx_callable_154 *hxrt.Function[func() any] = e1
					return hxrt.FunctionCall(hx_callable_154)()
				}(), func() any {
					var hx_callable_155 *hxrt.Function[func() any] = e2
					return hxrt.FunctionCall(hx_callable_155)()
				}()) < 0)
			})
		case *hxrt.StdString(hxrt.StringFromLiteral("<=")):
			hx_switch_141 = hxrt.NewFunction(func() any {
				return (haxe__Template_compareValues(func() any {
					var hx_callable_156 *hxrt.Function[func() any] = e1
					return hxrt.FunctionCall(hx_callable_156)()
				}(), func() any {
					var hx_callable_157 *hxrt.Function[func() any] = e2
					return hxrt.FunctionCall(hx_callable_157)()
				}()) <= 0)
			})
		case *hxrt.StdString(hxrt.StringFromLiteral("==")):
			hx_switch_141 = hxrt.NewFunction(func() any {
				return hxrt.HaxeEqual(func() any {
					var hx_callable_158 *hxrt.Function[func() any] = e1
					return hxrt.FunctionCall(hx_callable_158)()
				}(), func() any {
					var hx_callable_159 *hxrt.Function[func() any] = e2
					return hxrt.FunctionCall(hx_callable_159)()
				}())
			})
		case *hxrt.StdString(hxrt.StringFromLiteral(">")):
			hx_switch_141 = hxrt.NewFunction(func() any {
				return (haxe__Template_compareValues(func() any {
					var hx_callable_160 *hxrt.Function[func() any] = e1
					return hxrt.FunctionCall(hx_callable_160)()
				}(), func() any {
					var hx_callable_161 *hxrt.Function[func() any] = e2
					return hxrt.FunctionCall(hx_callable_161)()
				}()) > 0)
			})
		case *hxrt.StdString(hxrt.StringFromLiteral(">=")):
			hx_switch_141 = hxrt.NewFunction(func() any {
				return (haxe__Template_compareValues(func() any {
					var hx_callable_162 *hxrt.Function[func() any] = e1
					return hxrt.FunctionCall(hx_callable_162)()
				}(), func() any {
					var hx_callable_163 *hxrt.Function[func() any] = e2
					return hxrt.FunctionCall(hx_callable_163)()
				}()) >= 0)
			})
		case *hxrt.StdString(hxrt.StringFromLiteral("||")):
			hx_switch_141 = hxrt.NewFunction(func() any {
				return (haxe__Template_valueAsBool(func() any {
					var hx_callable_164 *hxrt.Function[func() any] = e1
					return hxrt.FunctionCall(hx_callable_164)()
				}()) || haxe__Template_valueAsBool(func() any {
					var hx_callable_165 *hxrt.Function[func() any] = e2
					return hxrt.FunctionCall(hx_callable_165)()
				}()))
			})
		default:
			hx_switch_141 = func() *hxrt.Function[func() any] {
				hxrt.Throw(haxe__Exception_thrown(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Unknown operation "), func(hx_value_166 any) *string {
					if hx_value_166 == nil {
						var hx_zero_167 *string
						return hx_zero_167
					}
					return hx_value_166.(*string)
				}(hxrt.AnonymousGet(op, "p")))))
				var hx_throw_zero_168 *hxrt.Function[func() any]
				return hx_throw_zero_168
			}()
		}
		return hx_switch_141
	case *hxrt.StdString(hxrt.StringFromLiteral("-")):
		inner_1 := self.__hx_this.makeExpr(cursor)
		return hxrt.NewFunction(func() any {
			return -haxe__Template_valueAsFloat(func() any {
				var hx_callable_169 *hxrt.Function[func() any] = inner_1
				return hxrt.FunctionCall(hx_callable_169)()
			}())
		})
	default:
		hxrt.Throw(haxe__Exception_thrown(func(hx_value_170 any) *string {
			if hx_value_170 == nil {
				var hx_zero_171 *string
				return hx_zero_171
			}
			return hx_value_170.(*string)
		}(hxrt.AnonymousGet(token, "p"))))
		var hx_throw_zero_172 *hxrt.Function[func() any]
		return hx_throw_zero_172
	}
}

func (self *haxe__Template) run(e *haxe___Template__TemplateExpr) {
	switch e.tag {
	case 0:
		_g := func(hx_value_173 any) *string {
			if hx_value_173 == nil {
				var hx_zero_174 *string
				return hx_zero_174
			}
			return hx_value_173.(*string)
		}(e.params[0])
		v := _g
		self.output = hxrt.StringConcatStringPtr(self.output, hxrt.StdString(self.__hx_this.resolve(v)))
	case 1:
		_g_1 := hx_recover_callable_0(e.params[0])
		expr := _g_1
		self.output = hxrt.StringConcatStringPtr(self.output, hxrt.StdString(func() any {
			var hx_callable_175 *hxrt.Function[func() any] = expr
			return hxrt.FunctionCall(hx_callable_175)()
		}()))
	case 2:
		_g_2 := hx_recover_callable_0(e.params[0])
		_g1 := func(hx_value_176 any) *haxe___Template__TemplateExpr {
			if hx_value_176 == nil {
				var hx_zero_177 *haxe___Template__TemplateExpr
				return hx_zero_177
			}
			return hx_value_176.(*haxe___Template__TemplateExpr)
		}(e.params[1])
		_g2 := func(hx_value_178 any) *haxe___Template__TemplateExpr {
			if hx_value_178 == nil {
				var hx_zero_179 *haxe___Template__TemplateExpr
				return hx_zero_179
			}
			return hx_value_178.(*haxe___Template__TemplateExpr)
		}(e.params[2])
		expr_1 := _g_2
		ifExpr := _g1
		elseExpr := _g2
		var value any = func() any {
			var hx_callable_180 *hxrt.Function[func() any] = expr_1
			return hxrt.FunctionCall(hx_callable_180)()
		}()
		if hxrt.AnyEqualsNull(value) || hxrt.HaxeEqual(value, false) {
			if elseExpr != nil {
				self.__hx_this.run(elseExpr)
			}
		} else {
			self.__hx_this.run(ifExpr)
		}
	case 3:
		_g_3 := func(hx_value_181 any) *string {
			if hx_value_181 == nil {
				var hx_zero_182 *string
				return hx_zero_182
			}
			return hx_value_181.(*string)
		}(e.params[0])
		str := _g_3
		self.output = hxrt.StringConcatStringPtr(self.output, str)
	case 4:
		_g_4 := func(hx_value_183 any) *hxrt.Array {
			if hx_value_183 == nil {
				var hx_zero_184 *hxrt.Array
				return hx_zero_184
			}
			return hx_value_183.(*hxrt.Array)
		}(e.params[0])
		items := _g_4
		_g_5 := 0
		for _g_5 < items.Len() {
			item := func(hx_value_185 any) *haxe___Template__TemplateExpr {
				if hx_value_185 == nil {
					var hx_zero_186 *haxe___Template__TemplateExpr
					return hx_zero_186
				}
				return hx_value_185.(*haxe___Template__TemplateExpr)
			}(items.Get(_g_5))
			_g_5 = int(int32((_g_5 + 1)))
			self.__hx_this.run(item)
		}
	case 5:
		_g_6 := hx_recover_callable_0(e.params[0])
		_g1_1 := func(hx_value_187 any) *haxe___Template__TemplateExpr {
			if hx_value_187 == nil {
				var hx_zero_188 *haxe___Template__TemplateExpr
				return hx_zero_188
			}
			return hx_value_187.(*haxe___Template__TemplateExpr)
		}(e.params[1])
		expr_2 := _g_6
		loop := _g1_1
		var value_1 any = func() any {
			var hx_callable_189 *hxrt.Function[func() any] = expr_2
			return hxrt.FunctionCall(hx_callable_189)()
		}()
		arrayValues := hxrt.TemplateArrayValues(value_1)
		if arrayValues != nil {
			hx_arr_190 := self.stack
			hx_arr_190.Push(self.context)
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
				hxrt.Throw(haxe__Exception_thrown(nil))
			}
			var candidate any = hxrt.TemplateCall(iteratorField, []any{})
			if !Reflect_hasField(candidate, hxrt.StringFromLiteral("hasNext")) {
				hxrt.Throw(haxe__Exception_thrown(nil))
			}
			iterator = candidate
		}, func(hx_caught_191 any) {
			hx_exception_195 := haxe__Exception_caught(hx_caught_191)
			hx_payload_196 := haxe__Exception_catchValue(hx_caught_191)
			_ = hx_exception_195
			var hx_tmp_1 any = hx_payload_196
			_ = hx_tmp_1
			hxrt.TryCatch(func() {
				if hxrt.AnyEqualsNull(value_1) || !Reflect_hasField(value_1, hxrt.StringFromLiteral("hasNext")) {
					hxrt.Throw(haxe__Exception_thrown(nil))
				}
				iterator = value_1
			}, func(hx_caught_192 any) {
				hx_exception_193 := haxe__Exception_caught(hx_caught_192)
				hx_payload_194 := haxe__Exception_catchValue(hx_caught_192)
				_ = hx_exception_193
				var hx_tmp any = hx_payload_194
				_ = hx_tmp
				hxrt.Throw(haxe__Exception_thrown(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Cannot iter on "), hxrt.StdString(value_1))))
			})
		})
		var hasNext any = Reflect_field(iterator, hxrt.StringFromLiteral("hasNext"))
		var next any = Reflect_field(iterator, hxrt.StringFromLiteral("next"))
		if hxrt.AnyEqualsNull(hasNext) || hxrt.AnyEqualsNull(next) {
			hxrt.Throw(haxe__Exception_thrown(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Cannot iter on "), hxrt.StdString(value_1))))
		}
		hx_arr_197 := self.stack
		hx_arr_197.Push(self.context)
		for hxrt.HaxeEqual(hxrt.TemplateCall(hasNext, []any{}), true) {
			self.context = hxrt.TemplateCall(next, []any{})
			self.__hx_this.run(loop)
		}
		self.context = self.__hx_this.popStackValue()
	case 6:
		_g_8 := func(hx_value_198 any) *string {
			if hx_value_198 == nil {
				var hx_zero_199 *string
				return hx_zero_199
			}
			return hx_value_198.(*string)
		}(e.params[0])
		_g1_2 := func(hx_value_200 any) *hxrt.Array {
			if hx_value_200 == nil {
				var hx_zero_201 *hxrt.Array
				return hx_zero_201
			}
			return hx_value_200.(*hxrt.Array)
		}(e.params[1])
		name := _g_8
		params := _g1_2
		var fn any = Reflect_field(self.macros, name)
		callArgs := hxrt.NewArray()
		callArgs.Push(func() *hxrt.Function[func(*string) any] {
			hx_receiver_203 := self
			return hxrt.BoundFunction(func(v *string) any {
				return hx_receiver_203.__hx_this.resolve(v)
			}, func() any {
				return hx_receiver_203.__hx_this
			}, "resolve")
		}())
		_g_9 := 0
		for _g_9 < params.Len() {
			param := func(hx_value_204 any) *haxe___Template__TemplateExpr {
				if hx_value_204 == nil {
					var hx_zero_205 *haxe___Template__TemplateExpr
					return hx_zero_205
				}
				return hx_value_204.(*haxe___Template__TemplateExpr)
			}(params.Get(_g_9))
			_g_9 = int(int32((_g_9 + 1)))
			if param.tag == 0 {
				_g_10 := func(hx_value_206 any) *string {
					if hx_value_206 == nil {
						var hx_zero_207 *string
						return hx_zero_207
					}
					return hx_value_206.(*string)
				}(param.params[0])
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
		}, func(hx_caught_210 any) {
			hx_exception_215 := haxe__Exception_caught(hx_caught_210)
			hx_payload_216 := haxe__Exception_catchValue(hx_caught_210)
			_ = hx_exception_215
			var err any = hx_payload_216
			var hx_try_211 *string
			hxrt.TryCatch(func() {
				hx_try_211 = haxe__Template_joinDynamicArgs(callArgs)
			}, func(hx_caught_212 any) {
				hx_exception_213 := haxe__Exception_caught(hx_caught_212)
				hx_payload_214 := haxe__Exception_catchValue(hx_caught_212)
				_ = hx_exception_213
				var hx_tmp_2 any = hx_payload_214
				_ = hx_tmp_2
				hx_try_211 = hxrt.StringFromLiteral("???")
			})
			argsText := hx_try_211
			hxrt.Throw(haxe__Exception_thrown(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Macro call "), name), hxrt.StringFromLiteral("(")), argsText), hxrt.StringFromLiteral(") failed (")), hxrt.StdString(err)), hxrt.StringFromLiteral(")"))))
		})
	}
}

func (self *haxe__Template) popStackValue() any {
	lastIndex := int((hxrt.Int32Wrap(self.stack.Len()) - hxrt.Int32Wrap(1)))
	var value any = self.stack.Get(lastIndex)
	remaining := hxrt.NewArray()
	_g := 0
	_g1 := lastIndex
	for _g < _g1 {
		hx_post_217 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_217
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
	}(left) || func(hx_value any) bool {
		switch hx_value.(type) {
		case *string:
			return true
		case string:
			return true
		default:
			return false
		}
	}(right) {
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
	}(left) || func(hx_value any) bool {
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
	}(left))
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
	}(right) || func(hx_value any) bool {
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
	}(right))
	if leftNumeric && rightNumeric {
		leftFloat := haxe__Template_valueAsFloat(left)
		rightFloat := haxe__Template_valueAsFloat(right)
		var hx_if_220 int
		if leftFloat < rightFloat {
			hx_if_220 = -1
		} else {
			var hx_if_219 int
			if leftFloat > rightFloat {
				hx_if_219 = 1
			} else {
				hx_if_219 = 0
			}
			hx_if_220 = hx_if_219
		}
		return hx_if_220
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

var haxe__Template_globals any = any(func() hxrt.Anonymous {
	hx_obj_221 := map[string]any{}
	return hxrt.AnonymousFromMap(hx_obj_221)
}())

func haxe__Template_isSpaceOnly(value *string) bool {
	var _g_s *string
	var _g_offset int
	_g_offset = 0
	_g_s = value
	for _g_offset < hxrt.StringLengthStringPtr(_g_s) {
		code := hxrt.StringCharCodeAtStringPtr(_g_s, func() int {
			hx_post_222 := _g_offset
			_g_offset = int(int32((_g_offset + 1)))
			return hx_post_222
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
		hx_post_223 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_223
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
				hx_post_224 := _g_offset
				_g_offset = int(int32((_g_offset + 1)))
				return hx_post_224
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
		intPart = ((intPart * 10.0) + float64(int((hxrt.Int32Wrap(code) - hxrt.Int32Wrap(48)))))
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
			fracPart = ((fracPart * 10.0) + float64(int((hxrt.Int32Wrap(code_1) - hxrt.Int32Wrap(48)))))
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
				exponent = int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(exponent) * hxrt.Int32Wrap(10)))) + hxrt.Int32Wrap(int((hxrt.Int32Wrap(code_2) - hxrt.Int32Wrap(48))))))
				index = int(int32((index + 1)))
			}
			for exponent > 0 {
				var hx_if_225 float64
				if exponentSign < 0 {
					hx_if_225 = (result / 10.0)
				} else {
					hx_if_225 = (result * 10.0)
				}
				result = hx_if_225
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
		out = int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(out) * hxrt.Int32Wrap(10)))) + hxrt.Int32Wrap(int((hxrt.Int32Wrap(code) - hxrt.Int32Wrap(48))))))
		index = int(int32((index + 1)))
	}
	return out
}

func haxe__Template_peekExprToken(cursor *haxe___Template__ExprCursor) hxrt.Anonymous {
	var hx_if_226 hxrt.Anonymous
	if cursor.index < cursor.tokens.Len() {
		hx_if_226 = hx_structural_recover(hx_structural_recover(cursor.tokens.Get(cursor.index)))
	} else {
		hx_if_226 = nil
	}
	return hx_if_226
}

func haxe__Template_peekToken(cursor *haxe___Template__TokenCursor) hxrt.Anonymous {
	var hx_if_227 hxrt.Anonymous
	if cursor.index < cursor.tokens.Len() {
		hx_if_227 = hx_structural_recover(hx_structural_recover(cursor.tokens.Get(cursor.index)))
	} else {
		hx_if_227 = nil
	}
	return hx_if_227
}

func haxe__Template_popExprToken(cursor *haxe___Template__ExprCursor) hxrt.Anonymous {
	if cursor.index >= cursor.tokens.Len() {
		return nil
	}
	hx_post_228 := cursor.index
	cursor.index = int(int32((cursor.index + 1)))
	return hx_structural_recover(cursor.tokens.Get(hx_post_228))
}

func haxe__Template_popToken(cursor *haxe___Template__TokenCursor) hxrt.Anonymous {
	if cursor.index >= cursor.tokens.Len() {
		return nil
	}
	hx_post_229 := cursor.index
	cursor.index = int(int32((cursor.index + 1)))
	return hx_structural_recover(cursor.tokens.Get(hx_post_229))
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
	}(value) || func(hx_value any) bool {
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
	}(value) {
		return func(hx_value_232 any) float64 {
			if hx_value_232 == nil {
				var hx_zero_233 float64
				return hx_zero_233
			}
			return hx_value_232.(float64)
		}(func(hx_value_230 any) float64 {
			if hx_value_230 == nil {
				var hx_zero_231 float64
				return hx_zero_231
			}
			return hx_value_230.(float64)
		}(value))
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
	}(value) {
		return haxe__Template_parseFloatLiteral(hxrt.StdString(func(hx_value_234 any) *string {
			if hx_value_234 == nil {
				var hx_zero_235 *string
				return hx_zero_235
			}
			return hx_value_234.(*string)
		}(value)))
	}
	hxrt.Throw(haxe__Exception_thrown(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Expected numeric expression value, got "), hxrt.StdString(value))))
	var hx_throw_zero_236 float64
	return hx_throw_zero_236
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

func haxe___Template__TemplateExpr_OpExpr(expr *hxrt.Function[func() any]) *haxe___Template__TemplateExpr {
	enumValue := &haxe___Template__TemplateExpr{tag: 1}
	enumValue.params = []any{expr}
	return enumValue
}

func haxe___Template__TemplateExpr_OpIf(expr *hxrt.Function[func() any], eif *haxe___Template__TemplateExpr, eelse *haxe___Template__TemplateExpr) *haxe___Template__TemplateExpr {
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

func haxe___Template__TemplateExpr_OpForeach(expr *hxrt.Function[func() any], loop *haxe___Template__TemplateExpr) *haxe___Template__TemplateExpr {
	enumValue := &haxe___Template__TemplateExpr{tag: 5}
	enumValue.params = []any{expr, loop}
	return enumValue
}

func haxe___Template__TemplateExpr_OpMacro(name *string, params *hxrt.Array) *haxe___Template__TemplateExpr {
	enumValue := &haxe___Template__TemplateExpr{tag: 6}
	enumValue.params = []any{name, params}
	return enumValue
}
