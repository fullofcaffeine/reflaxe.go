package main

import "snapshot/hxrt"

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

func main() {
	basic := New_haxe__Template(hxrt.StringFromLiteral("::name::"))
	var v any = any(basic.__hx_this.execute(func() hxrt.Anonymous {
		hx_obj_1 := map[string]any{}
		hx_obj_1["name"] = hxrt.StringFromLiteral("ok")
		return hxrt.AnonymousFromMap(hx_obj_1)
	}(), nil))
	hxrt.Println(v)
	cond := New_haxe__Template(hxrt.StringFromLiteral("::if enabled::yes::else::no::end::"))
	var v_1 any = any(cond.__hx_this.execute(func() hxrt.Anonymous {
		hx_obj_2 := map[string]any{}
		hx_obj_2["enabled"] = true
		return hxrt.AnonymousFromMap(hx_obj_2)
	}(), nil))
	hxrt.Println(v_1)
	var v_2 any = any(cond.__hx_this.execute(func() hxrt.Anonymous {
		hx_obj_3 := map[string]any{}
		hx_obj_3["enabled"] = false
		return hxrt.AnonymousFromMap(hx_obj_3)
	}(), nil))
	hxrt.Println(v_2)
	loop := New_haxe__Template(hxrt.StringFromLiteral("::foreach items::::__current__::::end::"))
	var v_3 any = any(loop.__hx_this.execute(func() hxrt.Anonymous {
		hx_obj_4 := map[string]any{}
		hx_obj_4["items"] = hxrt.NewArray(hxrt.StringFromLiteral("a"), hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("c"))
		return hxrt.AnonymousFromMap(hx_obj_4)
	}(), nil))
	hxrt.Println(v_3)
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_1 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_1.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_1.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_1.high = func(hx_value_2 any) int {
						if hx_value_2 == nil {
							var hx_zero_3 int
							return hx_zero_3
						}
						return hx_value_2.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_1.low = func(hx_value_4 any) int {
						if hx_value_4 == nil {
							var hx_zero_5 int
							return hx_zero_5
						}
						return hx_value_4.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe___Template__TokenCursor:
		return func(hx_structural_source_6 *haxe___Template__TokenCursor) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_6, func(name string) hxrt.AnonymousField {
				switch name {
				case "tokens":
					return hxrt.AnonymousField{Value: hx_structural_source_6.tokens, Found: true}
				case "index":
					return hxrt.AnonymousField{Value: hx_structural_source_6.index, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "tokens":
					hx_structural_source_6.tokens = func(hx_value_7 any) *hxrt.Array {
						if hx_value_7 == nil {
							var hx_zero_8 *hxrt.Array
							return hx_zero_8
						}
						return hx_value_7.(*hxrt.Array)
					}(incoming)
					return true
				case "index":
					hx_structural_source_6.index = func(hx_value_9 any) int {
						if hx_value_9 == nil {
							var hx_zero_10 int
							return hx_zero_10
						}
						return hx_value_9.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe___Template__ExprCursor:
		return func(hx_structural_source_11 *haxe___Template__ExprCursor) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_11, func(name string) hxrt.AnonymousField {
				switch name {
				case "tokens":
					return hxrt.AnonymousField{Value: hx_structural_source_11.tokens, Found: true}
				case "index":
					return hxrt.AnonymousField{Value: hx_structural_source_11.index, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "tokens":
					hx_structural_source_11.tokens = func(hx_value_12 any) *hxrt.Array {
						if hx_value_12 == nil {
							var hx_zero_13 *hxrt.Array
							return hx_zero_13
						}
						return hx_value_12.(*hxrt.Array)
					}(incoming)
					return true
				case "index":
					hx_structural_source_11.index = func(hx_value_14 any) int {
						if hx_value_14 == nil {
							var hx_zero_15 int
							return hx_zero_15
						}
						return hx_value_14.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__Template:
		return func(hx_structural_source_16 *haxe__Template) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_16, func(name string) hxrt.AnonymousField {
				switch name {
				case "expr":
					return hxrt.AnonymousField{Value: hx_structural_source_16.expr, Found: true}
				case "context":
					return hxrt.AnonymousField{Value: hx_structural_source_16.context, Found: true}
				case "macros":
					return hxrt.AnonymousField{Value: hx_structural_source_16.macros, Found: true}
				case "stack":
					return hxrt.AnonymousField{Value: hx_structural_source_16.stack, Found: true}
				case "output":
					return hxrt.AnonymousField{Value: hx_structural_source_16.output, Found: true}
				case "execute":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any, any) *string] {
						hx_receiver_23 := hx_structural_source_16
						return hxrt.BoundFunction(func(context any, macros any) *string {
							return hx_receiver_23.__hx_this.execute(context, macros)
						}, func() any {
							return hx_receiver_23.__hx_this
						}, "execute")
					}(), Found: true}
				case "resolve":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) any] {
						hx_receiver_24 := hx_structural_source_16
						return hxrt.BoundFunction(func(v *string) any {
							return hx_receiver_24.__hx_this.resolve(v)
						}, func() any {
							return hx_receiver_24.__hx_this
						}, "resolve")
					}(), Found: true}
				case "parseTokens":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) *hxrt.Array] {
						hx_receiver_25 := hx_structural_source_16
						return hxrt.BoundFunction(func(data *string) *hxrt.Array {
							return hx_receiver_25.__hx_this.parseTokens(data)
						}, func() any {
							return hx_receiver_25.__hx_this
						}, "parseTokens")
					}(), Found: true}
				case "parseBlock":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe___Template__TokenCursor) *haxe___Template__TemplateExpr] {
						hx_receiver_26 := hx_structural_source_16
						return hxrt.BoundFunction(func(cursor *haxe___Template__TokenCursor) *haxe___Template__TemplateExpr {
							return hx_receiver_26.__hx_this.parseBlock(cursor)
						}, func() any {
							return hx_receiver_26.__hx_this
						}, "parseBlock")
					}(), Found: true}
				case "parse":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe___Template__TokenCursor) *haxe___Template__TemplateExpr] {
						hx_receiver_27 := hx_structural_source_16
						return hxrt.BoundFunction(func(cursor *haxe___Template__TokenCursor) *haxe___Template__TemplateExpr {
							return hx_receiver_27.__hx_this.parse(cursor)
						}, func() any {
							return hx_receiver_27.__hx_this
						}, "parse")
					}(), Found: true}
				case "parseExpr":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) *hxrt.Function[func() any]] {
						hx_receiver_28 := hx_structural_source_16
						return hxrt.BoundFunction(func(data *string) *hxrt.Function[func() any] {
							return hx_receiver_28.__hx_this.parseExpr(data)
						}, func() any {
							return hx_receiver_28.__hx_this
						}, "parseExpr")
					}(), Found: true}
				case "makeConst":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) *hxrt.Function[func() any]] {
						hx_receiver_29 := hx_structural_source_16
						return hxrt.BoundFunction(func(v *string) *hxrt.Function[func() any] {
							return hx_receiver_29.__hx_this.makeConst(v)
						}, func() any {
							return hx_receiver_29.__hx_this
						}, "makeConst")
					}(), Found: true}
				case "makePath":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func() any], *haxe___Template__ExprCursor) *hxrt.Function[func() any]] {
						hx_receiver_30 := hx_structural_source_16
						return hxrt.BoundFunction(func(e *hxrt.Function[func() any], cursor *haxe___Template__ExprCursor) *hxrt.Function[func() any] {
							return hx_receiver_30.__hx_this.makePath(e, cursor)
						}, func() any {
							return hx_receiver_30.__hx_this
						}, "makePath")
					}(), Found: true}
				case "makeExpr":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe___Template__ExprCursor) *hxrt.Function[func() any]] {
						hx_receiver_31 := hx_structural_source_16
						return hxrt.BoundFunction(func(cursor *haxe___Template__ExprCursor) *hxrt.Function[func() any] {
							return hx_receiver_31.__hx_this.makeExpr(cursor)
						}, func() any {
							return hx_receiver_31.__hx_this
						}, "makeExpr")
					}(), Found: true}
				case "skipSpaces":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe___Template__ExprCursor)] {
						hx_receiver_32 := hx_structural_source_16
						return hxrt.BoundFunction(func(cursor *haxe___Template__ExprCursor) {
							hx_receiver_32.__hx_this.skipSpaces(cursor)
						}, func() any {
							return hx_receiver_32.__hx_this
						}, "skipSpaces")
					}(), Found: true}
				case "makeExpr2":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe___Template__ExprCursor) *hxrt.Function[func() any]] {
						hx_receiver_33 := hx_structural_source_16
						return hxrt.BoundFunction(func(cursor *haxe___Template__ExprCursor) *hxrt.Function[func() any] {
							return hx_receiver_33.__hx_this.makeExpr2(cursor)
						}, func() any {
							return hx_receiver_33.__hx_this
						}, "makeExpr2")
					}(), Found: true}
				case "run":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe___Template__TemplateExpr)] {
						hx_receiver_34 := hx_structural_source_16
						return hxrt.BoundFunction(func(e *haxe___Template__TemplateExpr) {
							hx_receiver_34.__hx_this.run(e)
						}, func() any {
							return hx_receiver_34.__hx_this
						}, "run")
					}(), Found: true}
				case "popStackValue":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_35 := hx_structural_source_16
						return hxrt.BoundFunction(func() any {
							return hx_receiver_35.__hx_this.popStackValue()
						}, func() any {
							return hx_receiver_35.__hx_this
						}, "popStackValue")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "expr":
					hx_structural_source_16.expr = func(hx_value_17 any) *haxe___Template__TemplateExpr {
						if hx_value_17 == nil {
							var hx_zero_18 *haxe___Template__TemplateExpr
							return hx_zero_18
						}
						return hx_value_17.(*haxe___Template__TemplateExpr)
					}(incoming)
					return true
				case "context":
					hx_structural_source_16.context = incoming
					return true
				case "macros":
					hx_structural_source_16.macros = incoming
					return true
				case "stack":
					hx_structural_source_16.stack = func(hx_value_19 any) *hxrt.Array {
						if hx_value_19 == nil {
							var hx_zero_20 *hxrt.Array
							return hx_zero_20
						}
						return hx_value_19.(*hxrt.Array)
					}(incoming)
					return true
				case "output":
					hx_structural_source_16.output = func(hx_value_21 any) *string {
						if hx_value_21 == nil {
							var hx_zero_22 *string
							return hx_zero_22
						}
						return hx_value_21.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__StringKeyValueIterator:
		return func(hx_structural_source_36 *haxe__iterators__StringKeyValueIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_36, func(name string) hxrt.AnonymousField {
				switch name {
				case "offset":
					return hxrt.AnonymousField{Value: hx_structural_source_36.offset, Found: true}
				case "s":
					return hxrt.AnonymousField{Value: hx_structural_source_36.s, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_41 := hx_structural_source_36
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_41.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_41.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_42 := hx_structural_source_36
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_42.__hx_this.next()
						}, func() any {
							return hx_receiver_42.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "offset":
					hx_structural_source_36.offset = func(hx_value_37 any) int {
						if hx_value_37 == nil {
							var hx_zero_38 int
							return hx_zero_38
						}
						return hx_value_37.(int)
					}(incoming)
					return true
				case "s":
					hx_structural_source_36.s = func(hx_value_39 any) *string {
						if hx_value_39 == nil {
							var hx_zero_40 *string
							return hx_zero_40
						}
						return hx_value_39.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__StringIterator:
		return func(hx_structural_source_43 *haxe__iterators__StringIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_43, func(name string) hxrt.AnonymousField {
				switch name {
				case "offset":
					return hxrt.AnonymousField{Value: hx_structural_source_43.offset, Found: true}
				case "s":
					return hxrt.AnonymousField{Value: hx_structural_source_43.s, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_48 := hx_structural_source_43
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_48.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_48.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_49 := hx_structural_source_43
						return hxrt.BoundFunction(func() int {
							return hx_receiver_49.__hx_this.next()
						}, func() any {
							return hx_receiver_49.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "offset":
					hx_structural_source_43.offset = func(hx_value_44 any) int {
						if hx_value_44 == nil {
							var hx_zero_45 int
							return hx_zero_45
						}
						return hx_value_44.(int)
					}(incoming)
					return true
				case "s":
					hx_structural_source_43.s = func(hx_value_46 any) *string {
						if hx_value_46 == nil {
							var hx_zero_47 *string
							return hx_zero_47
						}
						return hx_value_46.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__Exception:
		return func(hx_structural_source_50 *haxe__Exception) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_50, func(name string) hxrt.AnonymousField {
				switch name {
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_50.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_50.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_50.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_50.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_53 := hx_structural_source_50
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_53.__hx_this.get_message()
						}, func() any {
							return hx_receiver_53.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_54 := hx_structural_source_50
						return hxrt.BoundFunction(func() any {
							return hx_receiver_54.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_54.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_55 := hx_structural_source_50
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_55.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_55.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_56 := hx_structural_source_50
						return hxrt.BoundFunction(func() any {
							return hx_receiver_56.__hx_this.get_native()
						}, func() any {
							return hx_receiver_56.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_57 := hx_structural_source_50
						return hxrt.BoundFunction(func() any {
							return hx_receiver_57.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_57.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_58 := hx_structural_source_50
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_58.__hx_this.toString()
						}, func() any {
							return hx_receiver_58.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_59 := hx_structural_source_50
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_59.__hx_this.details()
						}, func() any {
							return hx_receiver_59.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_50.savedMessage = func(hx_value_51 any) *string {
						if hx_value_51 == nil {
							var hx_zero_52 *string
							return hx_zero_52
						}
						return hx_value_51.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_50.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_50.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_50.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *EReg:
		return func(hx_structural_source_60 *EReg) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_60, func(name string) hxrt.AnonymousField {
				switch name {
				case "handle":
					return hxrt.AnonymousField{Value: hx_structural_source_60.handle, Found: true}
				case "global":
					return hxrt.AnonymousField{Value: hx_structural_source_60.global, Found: true}
				case "lastSource":
					return hxrt.AnonymousField{Value: hx_structural_source_60.lastSource, Found: true}
				case "lastMatch":
					return hxrt.AnonymousField{Value: hx_structural_source_60.lastMatch, Found: true}
				case "match":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) bool] {
						hx_receiver_69 := hx_structural_source_60
						return hxrt.BoundFunction(func(s *string) bool {
							return hx_receiver_69.__hx_this.match(s)
						}, func() any {
							return hx_receiver_69.__hx_this
						}, "match")
					}(), Found: true}
				case "matched":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *string] {
						hx_receiver_70 := hx_structural_source_60
						return hxrt.BoundFunction(func(n int) *string {
							return hx_receiver_70.__hx_this.matched(n)
						}, func() any {
							return hx_receiver_70.__hx_this
						}, "matched")
					}(), Found: true}
				case "matchedLeft":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_71 := hx_structural_source_60
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_71.__hx_this.matchedLeft()
						}, func() any {
							return hx_receiver_71.__hx_this
						}, "matchedLeft")
					}(), Found: true}
				case "matchedRight":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_72 := hx_structural_source_60
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_72.__hx_this.matchedRight()
						}, func() any {
							return hx_receiver_72.__hx_this
						}, "matchedRight")
					}(), Found: true}
				case "matchedPos":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_73 := hx_structural_source_60
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_73.__hx_this.matchedPos()
						}, func() any {
							return hx_receiver_73.__hx_this
						}, "matchedPos")
					}(), Found: true}
				case "matchSub":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, int, any) bool] {
						hx_receiver_74 := hx_structural_source_60
						return hxrt.BoundFunction(func(s *string, pos int, len_local any) bool {
							if len_local == nil {
								len_local = -1
							}
							return hx_receiver_74.__hx_this.matchSub(s, pos, func(hx_value_75 any) int {
								if hx_value_75 == nil {
									var hx_zero_76 int
									return hx_zero_76
								}
								return hx_value_75.(int)
							}(len_local))
						}, func() any {
							return hx_receiver_74.__hx_this
						}, "matchSub")
					}(), Found: true}
				case "split":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) *hxrt.Array] {
						hx_receiver_77 := hx_structural_source_60
						return hxrt.BoundFunction(func(s *string) *hxrt.Array {
							return hx_receiver_77.__hx_this.split(s)
						}, func() any {
							return hx_receiver_77.__hx_this
						}, "split")
					}(), Found: true}
				case "replace":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string) *string] {
						hx_receiver_78 := hx_structural_source_60
						return hxrt.BoundFunction(func(s *string, by *string) *string {
							return hx_receiver_78.__hx_this.replace(s, by)
						}, func() any {
							return hx_receiver_78.__hx_this
						}, "replace")
					}(), Found: true}
				case "map":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *hxrt.Function[func(*EReg) *string]) *string] {
						hx_receiver_79 := hx_structural_source_60
						return hxrt.BoundFunction(func(s *string, f *hxrt.Function[func(*EReg) *string]) *string {
							return hx_receiver_79.__hx_this.map_(s, f)
						}, func() any {
							return hx_receiver_79.__hx_this
						}, "map")
					}(), Found: true}
				case "remember":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *hxrt.RegexMatch)] {
						hx_receiver_80 := hx_structural_source_60
						return hxrt.BoundFunction(func(source *string, match *hxrt.RegexMatch) {
							hx_receiver_80.__hx_this.remember(source, match)
						}, func() any {
							return hx_receiver_80.__hx_this
						}, "remember")
					}(), Found: true}
				case "requireMatch":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *hxrt.RegexMatch] {
						hx_receiver_81 := hx_structural_source_60
						return hxrt.BoundFunction(func() *hxrt.RegexMatch {
							return hx_receiver_81.__hx_this.requireMatch()
						}, func() any {
							return hx_receiver_81.__hx_this
						}, "requireMatch")
					}(), Found: true}
				case "expandReplacement":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string, *hxrt.RegexMatch) *string] {
						hx_receiver_82 := hx_structural_source_60
						return hxrt.BoundFunction(func(by *string, source *string, currentMatch *hxrt.RegexMatch) *string {
							return hx_receiver_82.__hx_this.expandReplacement(by, source, currentMatch)
						}, func() any {
							return hx_receiver_82.__hx_this
						}, "expandReplacement")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "handle":
					hx_structural_source_60.handle = func(hx_value_61 any) *hxrt.RegexHandle {
						if hx_value_61 == nil {
							var hx_zero_62 *hxrt.RegexHandle
							return hx_zero_62
						}
						return hx_value_61.(*hxrt.RegexHandle)
					}(incoming)
					return true
				case "global":
					hx_structural_source_60.global = func(hx_value_63 any) bool {
						if hx_value_63 == nil {
							var hx_zero_64 bool
							return hx_zero_64
						}
						return hx_value_63.(bool)
					}(incoming)
					return true
				case "lastSource":
					hx_structural_source_60.lastSource = func(hx_value_65 any) *string {
						if hx_value_65 == nil {
							var hx_zero_66 *string
							return hx_zero_66
						}
						return hx_value_65.(*string)
					}(incoming)
					return true
				case "lastMatch":
					hx_structural_source_60.lastMatch = func(hx_value_67 any) *hxrt.RegexMatch {
						if hx_value_67 == nil {
							var hx_zero_68 *hxrt.RegexMatch
							return hx_zero_68
						}
						return hx_value_67.(*hxrt.RegexMatch)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__ValueException:
		return func(hx_structural_source_83 *haxe__ValueException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_83, func(name string) hxrt.AnonymousField {
				switch name {
				case "value":
					return hxrt.AnonymousField{Value: hx_structural_source_83.value, Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_84 := hx_structural_source_83
						return hxrt.BoundFunction(func() any {
							return hx_receiver_84.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_84.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_83.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_83.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_83.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_83.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_87 := hx_structural_source_83
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_87.__hx_this.get_message()
						}, func() any {
							return hx_receiver_87.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_88 := hx_structural_source_83
						return hxrt.BoundFunction(func() any {
							return hx_receiver_88.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_88.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_89 := hx_structural_source_83
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_89.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_89.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_90 := hx_structural_source_83
						return hxrt.BoundFunction(func() any {
							return hx_receiver_90.__hx_this.get_native()
						}, func() any {
							return hx_receiver_90.__hx_this
						}, "get_native")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_91 := hx_structural_source_83
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_91.__hx_this.toString()
						}, func() any {
							return hx_receiver_91.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_92 := hx_structural_source_83
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_92.__hx_this.details()
						}, func() any {
							return hx_receiver_92.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_83.savedMessage = func(hx_value_85 any) *string {
						if hx_value_85 == nil {
							var hx_zero_86 *string
							return hx_zero_86
						}
						return hx_value_85.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_83.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_83.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_83.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__io__Bytes:
		return func(hx_structural_source_93 *haxe__io__Bytes) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_93, func(name string) hxrt.AnonymousField {
				switch name {
				case "length":
					return hxrt.AnonymousField{Value: hx_structural_source_93.length, Found: true}
				case "b":
					return hxrt.AnonymousField{Value: hx_structural_source_93.b, Found: true}
				case "__hx_raw":
					return hxrt.AnonymousField{Value: hx_structural_source_93.__hx_raw, Found: true}
				case "__hx_rawValid":
					return hxrt.AnonymousField{Value: hx_structural_source_93.__hx_rawValid, Found: true}
				case "__hx_dataExposed":
					return hxrt.AnonymousField{Value: hx_structural_source_93.__hx_dataExposed, Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_102 := hx_structural_source_93
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_102.__hx_this.get(pos)
						}, func() any {
							return hx_receiver_102.__hx_this
						}, "get")
					}(), Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_103 := hx_structural_source_93
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_103.__hx_this.set(pos, value)
						}, func() any {
							return hx_receiver_103.__hx_this
						}, "set")
					}(), Found: true}
				case "blit":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe__io__Bytes, int, int)] {
						hx_receiver_104 := hx_structural_source_93
						return hxrt.BoundFunction(func(pos int, src *haxe__io__Bytes, srcpos int, len_local int) {
							hx_receiver_104.__hx_this.blit(pos, src, srcpos, len_local)
						}, func() any {
							return hx_receiver_104.__hx_this
						}, "blit")
					}(), Found: true}
				case "fill":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int, int)] {
						hx_receiver_105 := hx_structural_source_93
						return hxrt.BoundFunction(func(pos int, len_local int, value int) {
							hx_receiver_105.__hx_this.fill(pos, len_local, value)
						}, func() any {
							return hx_receiver_105.__hx_this
						}, "fill")
					}(), Found: true}
				case "sub":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int) *haxe__io__Bytes] {
						hx_receiver_106 := hx_structural_source_93
						return hxrt.BoundFunction(func(pos int, len_local int) *haxe__io__Bytes {
							return hx_receiver_106.__hx_this.sub(pos, len_local)
						}, func() any {
							return hx_receiver_106.__hx_this
						}, "sub")
					}(), Found: true}
				case "compare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes) int] {
						hx_receiver_107 := hx_structural_source_93
						return hxrt.BoundFunction(func(other *haxe__io__Bytes) int {
							return hx_receiver_107.__hx_this.compare(other)
						}, func() any {
							return hx_receiver_107.__hx_this
						}, "compare")
					}(), Found: true}
				case "getDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) float64] {
						hx_receiver_108 := hx_structural_source_93
						return hxrt.BoundFunction(func(pos int) float64 {
							return hx_receiver_108.__hx_this.getDouble(pos)
						}, func() any {
							return hx_receiver_108.__hx_this
						}, "getDouble")
					}(), Found: true}
				case "getFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) float64] {
						hx_receiver_109 := hx_structural_source_93
						return hxrt.BoundFunction(func(pos int) float64 {
							return hx_receiver_109.__hx_this.getFloat(pos)
						}, func() any {
							return hx_receiver_109.__hx_this
						}, "getFloat")
					}(), Found: true}
				case "setDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, float64)] {
						hx_receiver_110 := hx_structural_source_93
						return hxrt.BoundFunction(func(pos int, value float64) {
							hx_receiver_110.__hx_this.setDouble(pos, value)
						}, func() any {
							return hx_receiver_110.__hx_this
						}, "setDouble")
					}(), Found: true}
				case "setFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, float64)] {
						hx_receiver_111 := hx_structural_source_93
						return hxrt.BoundFunction(func(pos int, value float64) {
							hx_receiver_111.__hx_this.setFloat(pos, value)
						}, func() any {
							return hx_receiver_111.__hx_this
						}, "setFloat")
					}(), Found: true}
				case "getUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_112 := hx_structural_source_93
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_112.__hx_this.getUInt16(pos)
						}, func() any {
							return hx_receiver_112.__hx_this
						}, "getUInt16")
					}(), Found: true}
				case "setUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_113 := hx_structural_source_93
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_113.__hx_this.setUInt16(pos, value)
						}, func() any {
							return hx_receiver_113.__hx_this
						}, "setUInt16")
					}(), Found: true}
				case "getInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_114 := hx_structural_source_93
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_114.__hx_this.getInt32(pos)
						}, func() any {
							return hx_receiver_114.__hx_this
						}, "getInt32")
					}(), Found: true}
				case "getInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *haxe___Int64_____Int64] {
						hx_receiver_115 := hx_structural_source_93
						return hxrt.BoundFunction(func(pos int) *haxe___Int64_____Int64 {
							return hx_receiver_115.__hx_this.getInt64(pos)
						}, func() any {
							return hx_receiver_115.__hx_this
						}, "getInt64")
					}(), Found: true}
				case "setInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_116 := hx_structural_source_93
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_116.__hx_this.setInt32(pos, value)
						}, func() any {
							return hx_receiver_116.__hx_this
						}, "setInt32")
					}(), Found: true}
				case "setInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe___Int64_____Int64)] {
						hx_receiver_117 := hx_structural_source_93
						return hxrt.BoundFunction(func(pos int, value *haxe___Int64_____Int64) {
							hx_receiver_117.__hx_this.setInt64(pos, value)
						}, func() any {
							return hx_receiver_117.__hx_this
						}, "setInt64")
					}(), Found: true}
				case "getString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int, *haxe__io__Encoding) *string] {
						hx_receiver_118 := hx_structural_source_93
						return hxrt.BoundFunction(func(pos int, len_local int, encoding *haxe__io__Encoding) *string {
							return hx_receiver_118.__hx_this.getString(pos, len_local, encoding)
						}, func() any {
							return hx_receiver_118.__hx_this
						}, "getString")
					}(), Found: true}
				case "readString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int) *string] {
						hx_receiver_119 := hx_structural_source_93
						return hxrt.BoundFunction(func(pos int, len_local int) *string {
							return hx_receiver_119.__hx_this.readString(pos, len_local)
						}, func() any {
							return hx_receiver_119.__hx_this
						}, "readString")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_120 := hx_structural_source_93
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_120.__hx_this.toString()
						}, func() any {
							return hx_receiver_120.__hx_this
						}, "toString")
					}(), Found: true}
				case "toHex":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_121 := hx_structural_source_93
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_121.__hx_this.toHex()
						}, func() any {
							return hx_receiver_121.__hx_this
						}, "toHex")
					}(), Found: true}
				case "getData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() []int] {
						hx_receiver_122 := hx_structural_source_93
						return hxrt.BoundFunction(func() []int {
							return hx_receiver_122.__hx_this.getData()
						}, func() any {
							return hx_receiver_122.__hx_this
						}, "getData")
					}(), Found: true}
				case "__hx_nativeView":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *hxrt.ByteView] {
						hx_receiver_123 := hx_structural_source_93
						return hxrt.BoundFunction(func() *hxrt.ByteView {
							return hx_receiver_123.__hx_this.__hx_nativeView()
						}, func() any {
							return hx_receiver_123.__hx_this
						}, "__hx_nativeView")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "b":
					hx_structural_source_93.b = func(hx_value_94 any) []int {
						if hx_value_94 == nil {
							var hx_zero_95 []int
							return hx_zero_95
						}
						return hx_value_94.([]int)
					}(incoming)
					return true
				case "__hx_raw":
					hx_structural_source_93.__hx_raw = func(hx_value_96 any) *hxrt.ByteView {
						if hx_value_96 == nil {
							var hx_zero_97 *hxrt.ByteView
							return hx_zero_97
						}
						return hx_value_96.(*hxrt.ByteView)
					}(incoming)
					return true
				case "__hx_rawValid":
					hx_structural_source_93.__hx_rawValid = func(hx_value_98 any) bool {
						if hx_value_98 == nil {
							var hx_zero_99 bool
							return hx_zero_99
						}
						return hx_value_98.(bool)
					}(incoming)
					return true
				case "__hx_dataExposed":
					hx_structural_source_93.__hx_dataExposed = func(hx_value_100 any) bool {
						if hx_value_100 == nil {
							var hx_zero_101 bool
							return hx_zero_101
						}
						return hx_value_100.(bool)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	default:
		return hxrt.AnonymousFromValue(value)
	}
}

func hxrt__generated_method_field(obj any, key string) any {
	var receiver any
	switch value := hxrt.ObjectSource(obj).(type) {
	case *EReg:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Template:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Bytes:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *EReg:
		return hxrt__generated_method_field__EReg(value, key)
	case *haxe__Exception:
		return hxrt__generated_method_field__haxe__Exception(value, key)
	case *haxe__Template:
		return hxrt__generated_method_field__haxe__Template(value, key)
	case *haxe__ValueException:
		return hxrt__generated_method_field__haxe__ValueException(value, key)
	case *haxe__io__Bytes:
		return hxrt__generated_method_field__haxe__io__Bytes(value, key)
	case *haxe__iterators__StringIterator:
		return hxrt__generated_method_field__haxe__iterators__StringIterator(value, key)
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt__generated_method_field__haxe__iterators__StringKeyValueIterator(value, key)
	default:
		return nil
	}
}

func hxrt__generated_method_field__EReg(value *EReg, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "expandReplacement":
		return func() *hxrt.Function[func(*string, *string, *hxrt.RegexMatch) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(by *string, source *string, currentMatch *hxrt.RegexMatch) *string {
				return hx_receiver_1.__hx_this.expandReplacement(by, source, currentMatch)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "expandReplacement")
		}()
	case "map":
		return func() *hxrt.Function[func(*string, *hxrt.Function[func(*EReg) *string]) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(s *string, f *hxrt.Function[func(*EReg) *string]) *string {
				return hx_receiver_1.__hx_this.map_(s, f)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "map")
		}()
	case "match":
		return func() *hxrt.Function[func(*string) bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(s *string) bool {
				return hx_receiver_1.__hx_this.match(s)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "match")
		}()
	case "matchSub":
		return func() *hxrt.Function[func(*string, int, any) bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(s *string, pos int, len_local any) bool {
				if len_local == nil {
					len_local = -1
				}
				return hx_receiver_1.__hx_this.matchSub(s, pos, func(hx_value_2 any) int {
					if hx_value_2 == nil {
						var hx_zero_3 int
						return hx_zero_3
					}
					return hx_value_2.(int)
				}(len_local))
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "matchSub")
		}()
	case "matched":
		return func() *hxrt.Function[func(int) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(n int) *string {
				return hx_receiver_1.__hx_this.matched(n)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "matched")
		}()
	case "matchedLeft":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.matchedLeft()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "matchedLeft")
		}()
	case "matchedPos":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_1.__hx_this.matchedPos()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "matchedPos")
		}()
	case "matchedRight":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.matchedRight()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "matchedRight")
		}()
	case "remember":
		return func() *hxrt.Function[func(*string, *hxrt.RegexMatch)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(source *string, match *hxrt.RegexMatch) {
				hx_receiver_1.__hx_this.remember(source, match)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "remember")
		}()
	case "replace":
		return func() *hxrt.Function[func(*string, *string) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(s *string, by *string) *string {
				return hx_receiver_1.__hx_this.replace(s, by)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "replace")
		}()
	case "requireMatch":
		return func() *hxrt.Function[func() *hxrt.RegexMatch] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *hxrt.RegexMatch {
				return hx_receiver_1.__hx_this.requireMatch()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "requireMatch")
		}()
	case "split":
		return func() *hxrt.Function[func(*string) *hxrt.Array] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(s *string) *hxrt.Array {
				return hx_receiver_1.__hx_this.split(s)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "split")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__Exception(value *haxe__Exception, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "details":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.details()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "details")
		}()
	case "get_message":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.get_message()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get_message")
		}()
	case "get_native":
		return func() *hxrt.Function[func() any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_1.__hx_this.get_native()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get_native")
		}()
	case "get_previous":
		return func() *hxrt.Function[func() *haxe__Exception] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *haxe__Exception {
				return hx_receiver_1.__hx_this.get_previous()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get_previous")
		}()
	case "get_stack":
		return func() *hxrt.Function[func() any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_1.__hx_this.get_stack()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get_stack")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.toString()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "toString")
		}()
	case "unwrap":
		return func() *hxrt.Function[func() any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_1.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "unwrap")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__Template(value *haxe__Template, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "execute":
		return func() *hxrt.Function[func(any, any) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(context any, macros any) *string {
				return hx_receiver_1.__hx_this.execute(context, macros)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "execute")
		}()
	case "makeConst":
		return func() *hxrt.Function[func(*string) *hxrt.Function[func() any]] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(v *string) *hxrt.Function[func() any] {
				return hx_receiver_1.__hx_this.makeConst(v)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "makeConst")
		}()
	case "makeExpr":
		return func() *hxrt.Function[func(*haxe___Template__ExprCursor) *hxrt.Function[func() any]] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(cursor *haxe___Template__ExprCursor) *hxrt.Function[func() any] {
				return hx_receiver_1.__hx_this.makeExpr(cursor)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "makeExpr")
		}()
	case "makeExpr2":
		return func() *hxrt.Function[func(*haxe___Template__ExprCursor) *hxrt.Function[func() any]] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(cursor *haxe___Template__ExprCursor) *hxrt.Function[func() any] {
				return hx_receiver_1.__hx_this.makeExpr2(cursor)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "makeExpr2")
		}()
	case "makePath":
		return func() *hxrt.Function[func(*hxrt.Function[func() any], *haxe___Template__ExprCursor) *hxrt.Function[func() any]] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(e *hxrt.Function[func() any], cursor *haxe___Template__ExprCursor) *hxrt.Function[func() any] {
				return hx_receiver_1.__hx_this.makePath(e, cursor)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "makePath")
		}()
	case "parse":
		return func() *hxrt.Function[func(*haxe___Template__TokenCursor) *haxe___Template__TemplateExpr] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(cursor *haxe___Template__TokenCursor) *haxe___Template__TemplateExpr {
				return hx_receiver_1.__hx_this.parse(cursor)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "parse")
		}()
	case "parseBlock":
		return func() *hxrt.Function[func(*haxe___Template__TokenCursor) *haxe___Template__TemplateExpr] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(cursor *haxe___Template__TokenCursor) *haxe___Template__TemplateExpr {
				return hx_receiver_1.__hx_this.parseBlock(cursor)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "parseBlock")
		}()
	case "parseExpr":
		return func() *hxrt.Function[func(*string) *hxrt.Function[func() any]] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(data *string) *hxrt.Function[func() any] {
				return hx_receiver_1.__hx_this.parseExpr(data)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "parseExpr")
		}()
	case "parseTokens":
		return func() *hxrt.Function[func(*string) *hxrt.Array] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(data *string) *hxrt.Array {
				return hx_receiver_1.__hx_this.parseTokens(data)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "parseTokens")
		}()
	case "popStackValue":
		return func() *hxrt.Function[func() any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_1.__hx_this.popStackValue()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "popStackValue")
		}()
	case "resolve":
		return func() *hxrt.Function[func(*string) any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(v *string) any {
				return hx_receiver_1.__hx_this.resolve(v)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "resolve")
		}()
	case "run":
		return func() *hxrt.Function[func(*haxe___Template__TemplateExpr)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(e *haxe___Template__TemplateExpr) {
				hx_receiver_1.__hx_this.run(e)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "run")
		}()
	case "skipSpaces":
		return func() *hxrt.Function[func(*haxe___Template__ExprCursor)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(cursor *haxe___Template__ExprCursor) {
				hx_receiver_1.__hx_this.skipSpaces(cursor)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "skipSpaces")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__ValueException(value *haxe__ValueException, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "unwrap":
		return func() *hxrt.Function[func() any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_1.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "unwrap")
		}()
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_method_field__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_method_field__haxe__io__Bytes(value *haxe__io__Bytes, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "__hx_nativeView":
		return func() *hxrt.Function[func() *hxrt.ByteView] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *hxrt.ByteView {
				return hx_receiver_1.__hx_this.__hx_nativeView()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "__hx_nativeView")
		}()
	case "blit":
		return func() *hxrt.Function[func(int, *haxe__io__Bytes, int, int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, src *haxe__io__Bytes, srcpos int, len_local int) {
				hx_receiver_1.__hx_this.blit(pos, src, srcpos, len_local)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "blit")
		}()
	case "compare":
		return func() *hxrt.Function[func(*haxe__io__Bytes) int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(other *haxe__io__Bytes) int {
				return hx_receiver_1.__hx_this.compare(other)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "compare")
		}()
	case "fill":
		return func() *hxrt.Function[func(int, int, int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, len_local int, value int) {
				hx_receiver_1.__hx_this.fill(pos, len_local, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "fill")
		}()
	case "get":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int) int {
				return hx_receiver_1.__hx_this.get(pos)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get")
		}()
	case "getData":
		return func() *hxrt.Function[func() []int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() []int {
				return hx_receiver_1.__hx_this.getData()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "getData")
		}()
	case "getDouble":
		return func() *hxrt.Function[func(int) float64] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int) float64 {
				return hx_receiver_1.__hx_this.getDouble(pos)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "getDouble")
		}()
	case "getFloat":
		return func() *hxrt.Function[func(int) float64] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int) float64 {
				return hx_receiver_1.__hx_this.getFloat(pos)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "getFloat")
		}()
	case "getInt32":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int) int {
				return hx_receiver_1.__hx_this.getInt32(pos)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "getInt32")
		}()
	case "getInt64":
		return func() *hxrt.Function[func(int) *haxe___Int64_____Int64] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int) *haxe___Int64_____Int64 {
				return hx_receiver_1.__hx_this.getInt64(pos)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "getInt64")
		}()
	case "getString":
		return func() *hxrt.Function[func(int, int, *haxe__io__Encoding) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, len_local int, encoding *haxe__io__Encoding) *string {
				return hx_receiver_1.__hx_this.getString(pos, len_local, encoding)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "getString")
		}()
	case "getUInt16":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int) int {
				return hx_receiver_1.__hx_this.getUInt16(pos)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "getUInt16")
		}()
	case "readString":
		return func() *hxrt.Function[func(int, int) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, len_local int) *string {
				return hx_receiver_1.__hx_this.readString(pos, len_local)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "readString")
		}()
	case "set":
		return func() *hxrt.Function[func(int, int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, value int) {
				hx_receiver_1.__hx_this.set(pos, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "set")
		}()
	case "setDouble":
		return func() *hxrt.Function[func(int, float64)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, value float64) {
				hx_receiver_1.__hx_this.setDouble(pos, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "setDouble")
		}()
	case "setFloat":
		return func() *hxrt.Function[func(int, float64)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, value float64) {
				hx_receiver_1.__hx_this.setFloat(pos, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "setFloat")
		}()
	case "setInt32":
		return func() *hxrt.Function[func(int, int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, value int) {
				hx_receiver_1.__hx_this.setInt32(pos, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "setInt32")
		}()
	case "setInt64":
		return func() *hxrt.Function[func(int, *haxe___Int64_____Int64)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, value *haxe___Int64_____Int64) {
				hx_receiver_1.__hx_this.setInt64(pos, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "setInt64")
		}()
	case "setUInt16":
		return func() *hxrt.Function[func(int, int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, value int) {
				hx_receiver_1.__hx_this.setUInt16(pos, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "setUInt16")
		}()
	case "sub":
		return func() *hxrt.Function[func(int, int) *haxe__io__Bytes] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, len_local int) *haxe__io__Bytes {
				return hx_receiver_1.__hx_this.sub(pos, len_local)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "sub")
		}()
	case "toHex":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.toHex()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "toHex")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.toString()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "toString")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__iterators__StringIterator(value *haxe__iterators__StringIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasNext":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_1.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_1.__hx_this.next()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "next")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__iterators__StringKeyValueIterator(value *haxe__iterators__StringKeyValueIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasNext":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_1.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_1.__hx_this.next()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "next")
		}()
	}
	return nil
}

func hxrt_typeClassMetadataField(value any, key string) (any, bool) {
	classValue, ok := value.(*hxrt__TypeClassValue)
	if !ok || classValue == nil {
		return nil, false
	}
	className := *hxrt.StdString(classValue.name)
	switch className {
	default:
		return nil, false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedField(object any, field *string) any {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *EReg:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Template:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Template__ExprCursor:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Template__TokenCursor:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Bytes:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *EReg:
		return hxrt__generated_field_lookup__EReg(value, key)
	case *haxe__Exception:
		return hxrt__generated_field_lookup__haxe__Exception(value, key)
	case *haxe__Template:
		return hxrt__generated_field_lookup__haxe__Template(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_lookup__haxe__ValueException(value, key)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_lookup__haxe___Int64_____Int64(value, key)
	case *haxe___Template__ExprCursor:
		return hxrt__generated_field_lookup__haxe___Template__ExprCursor(value, key)
	case *haxe___Template__TokenCursor:
		return hxrt__generated_field_lookup__haxe___Template__TokenCursor(value, key)
	case *haxe__io__Bytes:
		return hxrt__generated_field_lookup__haxe__io__Bytes(value, key)
	case *haxe__iterators__StringIterator:
		return hxrt__generated_field_lookup__haxe__iterators__StringIterator(value, key)
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt__generated_field_lookup__haxe__iterators__StringKeyValueIterator(value, key)
	default:
		return nil
	}
}

func reflaxe__go___internal__CompilerReflect_hasGeneratedField(object any, field *string) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *EReg:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__Template:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Template__ExprCursor:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Template__TokenCursor:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__io__Bytes:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *EReg:
		return hxrt__generated_field_has__EReg(value, key)
	case *haxe__Exception:
		return hxrt__generated_field_has__haxe__Exception(value, key)
	case *haxe__Template:
		return hxrt__generated_field_has__haxe__Template(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_has__haxe__ValueException(value, key)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_has__haxe___Int64_____Int64(value, key)
	case *haxe___Template__ExprCursor:
		return hxrt__generated_field_has__haxe___Template__ExprCursor(value, key)
	case *haxe___Template__TokenCursor:
		return hxrt__generated_field_has__haxe___Template__TokenCursor(value, key)
	case *haxe__io__Bytes:
		return hxrt__generated_field_has__haxe__io__Bytes(value, key)
	case *haxe__iterators__StringIterator:
		return hxrt__generated_field_has__haxe__iterators__StringIterator(value, key)
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt__generated_field_has__haxe__iterators__StringKeyValueIterator(value, key)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_setGeneratedField(object any, field *string, incoming any) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *EReg:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__Template:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Template__ExprCursor:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Template__TokenCursor:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__io__Bytes:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *EReg:
		return hxrt__generated_field_set__EReg(value, key, incoming)
	case *haxe__Exception:
		return hxrt__generated_field_set__haxe__Exception(value, key, incoming)
	case *haxe__Template:
		return hxrt__generated_field_set__haxe__Template(value, key, incoming)
	case *haxe__ValueException:
		return hxrt__generated_field_set__haxe__ValueException(value, key, incoming)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_set__haxe___Int64_____Int64(value, key, incoming)
	case *haxe___Template__ExprCursor:
		return hxrt__generated_field_set__haxe___Template__ExprCursor(value, key, incoming)
	case *haxe___Template__TokenCursor:
		return hxrt__generated_field_set__haxe___Template__TokenCursor(value, key, incoming)
	case *haxe__io__Bytes:
		return hxrt__generated_field_set__haxe__io__Bytes(value, key, incoming)
	case *haxe__iterators__StringIterator:
		return hxrt__generated_field_set__haxe__iterators__StringIterator(value, key, incoming)
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt__generated_field_set__haxe__iterators__StringKeyValueIterator(value, key, incoming)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedFields(object any) *hxrt.Array {
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *EReg:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Template:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Template__ExprCursor:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Template__TokenCursor:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Bytes:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch receiver.(type) {
	case *EReg:
		return hxrt.NewArray(hxrt.StringFromLiteral("handle"), hxrt.StringFromLiteral("global"), hxrt.StringFromLiteral("lastSource"), hxrt.StringFromLiteral("lastMatch"))
	case *haxe__Exception:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"))
	case *haxe__Template:
		return hxrt.NewArray(hxrt.StringFromLiteral("expr"), hxrt.StringFromLiteral("context"), hxrt.StringFromLiteral("macros"), hxrt.StringFromLiteral("stack"), hxrt.StringFromLiteral("output"))
	case *haxe__ValueException:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("value"))
	case *haxe___Int64_____Int64:
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	case *haxe___Template__ExprCursor:
		return hxrt.NewArray(hxrt.StringFromLiteral("tokens"), hxrt.StringFromLiteral("index"))
	case *haxe___Template__TokenCursor:
		return hxrt.NewArray(hxrt.StringFromLiteral("tokens"), hxrt.StringFromLiteral("index"))
	case *haxe__io__Bytes:
		return hxrt.NewArray(hxrt.StringFromLiteral("length"), hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("__hx_raw"), hxrt.StringFromLiteral("__hx_rawValid"), hxrt.StringFromLiteral("__hx_dataExposed"))
	case *haxe__iterators__StringIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("offset"), hxrt.StringFromLiteral("s"))
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("offset"), hxrt.StringFromLiteral("s"))
	default:
		return nil
	}
}

func hxrt__generated_field_lookup__EReg(value *EReg, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "global":
		return hxrt.ReadStoredField(value.global)
	case "handle":
		return hxrt.ReadStoredField(value.handle)
	case "lastMatch":
		return hxrt.ReadStoredField(value.lastMatch)
	case "lastSource":
		return hxrt.ReadStoredField(value.lastSource)
	}
	return nil
}

func hxrt__generated_field_has__EReg(value *EReg, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "global":
		return true
	case "handle":
		return true
	case "lastMatch":
		return true
	case "lastSource":
		return true
	}
	return false
}

func hxrt__generated_field_set__EReg(value *EReg, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "global":
		if incoming == nil {
			var zero bool
			value.global = zero
			return true
		}
		switch typed := incoming.(type) {
		case bool:
			value.global = typed
			return true
		default:
			return false
		}
	case "handle":
		if incoming == nil {
			var zero *hxrt.RegexHandle
			value.handle = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.RegexHandle:
			value.handle = typed
			return true
		default:
			return false
		}
	case "lastMatch":
		if incoming == nil {
			var zero *hxrt.RegexMatch
			value.lastMatch = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.RegexMatch:
			value.lastMatch = typed
			return true
		default:
			return false
		}
	case "lastSource":
		if incoming == nil {
			var zero *string
			value.lastSource = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.lastSource = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__Exception(value *haxe__Exception, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "savedMessage":
		return hxrt.ReadStoredField(value.savedMessage)
	case "savedNative":
		return hxrt.ReadStoredField(value.savedNative)
	case "savedPrevious":
		return hxrt.ReadStoredField(value.savedPrevious)
	case "savedStack":
		return hxrt.ReadStoredField(value.savedStack)
	}
	return nil
}

func hxrt__generated_field_has__haxe__Exception(value *haxe__Exception, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "savedMessage":
		return true
	case "savedNative":
		return true
	case "savedPrevious":
		return true
	case "savedStack":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__Exception(value *haxe__Exception, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "savedMessage":
		if incoming == nil {
			var zero *string
			value.savedMessage = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.savedMessage = typed
			return true
		default:
			return false
		}
	case "savedNative":
		if incoming == nil {
			var zero any
			value.savedNative = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.savedNative = typed
			return true
		default:
			return false
		}
	case "savedPrevious":
		if incoming == nil {
			var zero *haxe__Exception
			value.savedPrevious = zero
			return true
		}
		switch typed := incoming.(type) {
		case *haxe__Exception:
			value.savedPrevious = typed
			return true
		default:
			return false
		}
	case "savedStack":
		if incoming == nil {
			var zero any
			value.savedStack = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.savedStack = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__Template(value *haxe__Template, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "context":
		return hxrt.ReadStoredField(value.context)
	case "expr":
		return hxrt.ReadStoredField(value.expr)
	case "macros":
		return hxrt.ReadStoredField(value.macros)
	case "output":
		return hxrt.ReadStoredField(value.output)
	case "stack":
		return hxrt.ReadStoredField(value.stack)
	}
	return nil
}

func hxrt__generated_field_has__haxe__Template(value *haxe__Template, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "context":
		return true
	case "expr":
		return true
	case "macros":
		return true
	case "output":
		return true
	case "stack":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__Template(value *haxe__Template, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "context":
		if incoming == nil {
			var zero any
			value.context = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.context = typed
			return true
		default:
			return false
		}
	case "expr":
		if incoming == nil {
			var zero *haxe___Template__TemplateExpr
			value.expr = zero
			return true
		}
		switch typed := incoming.(type) {
		case *haxe___Template__TemplateExpr:
			value.expr = typed
			return true
		default:
			return false
		}
	case "macros":
		if incoming == nil {
			var zero any
			value.macros = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.macros = typed
			return true
		default:
			return false
		}
	case "output":
		if incoming == nil {
			var zero *string
			value.output = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.output = typed
			return true
		default:
			return false
		}
	case "stack":
		if incoming == nil {
			var zero *hxrt.Array
			value.stack = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.Array:
			value.stack = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__ValueException(value *haxe__ValueException, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "value":
		return hxrt.ReadStoredField(value.value)
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_field_lookup__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_has__haxe__ValueException(value *haxe__ValueException, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "value":
		return true
	}
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_has__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_set__haxe__ValueException(value *haxe__ValueException, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "value":
		if incoming == nil {
			var zero any
			value.value = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.value = typed
			return true
		default:
			return false
		}
	}
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_set__haxe__Exception(value.haxe__Exception, key, incoming)
}

func hxrt__generated_field_lookup__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "high":
		return hxrt.ReadStoredField(value.high)
	case "low":
		return hxrt.ReadStoredField(value.low)
	}
	return nil
}

func hxrt__generated_field_has__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "high":
		return true
	case "low":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "high":
		if incoming == nil {
			var zero int
			value.high = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.high = typed
			return true
		default:
			return false
		}
	case "low":
		if incoming == nil {
			var zero int
			value.low = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.low = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe___Template__ExprCursor(value *haxe___Template__ExprCursor, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "index":
		return hxrt.ReadStoredField(value.index)
	case "tokens":
		return hxrt.ReadStoredField(value.tokens)
	}
	return nil
}

func hxrt__generated_field_has__haxe___Template__ExprCursor(value *haxe___Template__ExprCursor, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "index":
		return true
	case "tokens":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe___Template__ExprCursor(value *haxe___Template__ExprCursor, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "index":
		if incoming == nil {
			var zero int
			value.index = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.index = typed
			return true
		default:
			return false
		}
	case "tokens":
		if incoming == nil {
			var zero *hxrt.Array
			value.tokens = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.Array:
			value.tokens = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe___Template__TokenCursor(value *haxe___Template__TokenCursor, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "index":
		return hxrt.ReadStoredField(value.index)
	case "tokens":
		return hxrt.ReadStoredField(value.tokens)
	}
	return nil
}

func hxrt__generated_field_has__haxe___Template__TokenCursor(value *haxe___Template__TokenCursor, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "index":
		return true
	case "tokens":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe___Template__TokenCursor(value *haxe___Template__TokenCursor, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "index":
		if incoming == nil {
			var zero int
			value.index = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.index = typed
			return true
		default:
			return false
		}
	case "tokens":
		if incoming == nil {
			var zero *hxrt.Array
			value.tokens = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.Array:
			value.tokens = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__io__Bytes(value *haxe__io__Bytes, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "__hx_dataExposed":
		return hxrt.ReadStoredField(value.__hx_dataExposed)
	case "__hx_raw":
		return hxrt.ReadStoredField(value.__hx_raw)
	case "__hx_rawValid":
		return hxrt.ReadStoredField(value.__hx_rawValid)
	case "b":
		return hxrt.ReadStoredField(value.b)
	case "length":
		return hxrt.ReadStoredField(value.length)
	}
	return nil
}

func hxrt__generated_field_has__haxe__io__Bytes(value *haxe__io__Bytes, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "__hx_dataExposed":
		return true
	case "__hx_raw":
		return true
	case "__hx_rawValid":
		return true
	case "b":
		return true
	case "length":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__io__Bytes(value *haxe__io__Bytes, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "__hx_dataExposed":
		if incoming == nil {
			var zero bool
			value.__hx_dataExposed = zero
			return true
		}
		switch typed := incoming.(type) {
		case bool:
			value.__hx_dataExposed = typed
			return true
		default:
			return false
		}
	case "__hx_raw":
		if incoming == nil {
			var zero *hxrt.ByteView
			value.__hx_raw = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.ByteView:
			value.__hx_raw = typed
			return true
		default:
			return false
		}
	case "__hx_rawValid":
		if incoming == nil {
			var zero bool
			value.__hx_rawValid = zero
			return true
		}
		switch typed := incoming.(type) {
		case bool:
			value.__hx_rawValid = typed
			return true
		default:
			return false
		}
	case "b":
		if incoming == nil {
			var zero []int
			value.b = zero
			return true
		}
		switch typed := incoming.(type) {
		case []int:
			value.b = typed
			return true
		default:
			return false
		}
	case "length":
		if incoming == nil {
			var zero int
			value.length = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.length = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__iterators__StringIterator(value *haxe__iterators__StringIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "offset":
		return hxrt.ReadStoredField(value.offset)
	case "s":
		return hxrt.ReadStoredField(value.s)
	}
	return nil
}

func hxrt__generated_field_has__haxe__iterators__StringIterator(value *haxe__iterators__StringIterator, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "offset":
		return true
	case "s":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__iterators__StringIterator(value *haxe__iterators__StringIterator, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "offset":
		if incoming == nil {
			var zero int
			value.offset = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.offset = typed
			return true
		default:
			return false
		}
	case "s":
		if incoming == nil {
			var zero *string
			value.s = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.s = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__iterators__StringKeyValueIterator(value *haxe__iterators__StringKeyValueIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "offset":
		return hxrt.ReadStoredField(value.offset)
	case "s":
		return hxrt.ReadStoredField(value.s)
	}
	return nil
}

func hxrt__generated_field_has__haxe__iterators__StringKeyValueIterator(value *haxe__iterators__StringKeyValueIterator, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "offset":
		return true
	case "s":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__iterators__StringKeyValueIterator(value *haxe__iterators__StringKeyValueIterator, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "offset":
		if incoming == nil {
			var zero int
			value.offset = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.offset = typed
			return true
		default:
			return false
		}
	case "s":
		if incoming == nil {
			var zero *string
			value.s = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.s = typed
			return true
		default:
			return false
		}
	}
	return false
}

func reflaxe__go___internal__CompilerReflect_typeField(object any, field *string) any {
	key := *hxrt.StdString(field)
	value, found := hxrt_typeClassMetadataField(object, key)
	if !found {
		return nil
	}
	return value
}

func reflaxe__go___internal__CompilerReflect_hasTypeField(object any, field *string) bool {
	key := *hxrt.StdString(field)
	_, found := hxrt_typeClassMetadataField(object, key)
	return found
}

func reflaxe__go___internal__CompilerReflect_generatedMethod(object any, field *string) any {
	key := *hxrt.StdString(field)
	return hxrt__generated_method_field(object, key)
}

func reflaxe__go___internal__CompilerReflect_isEnumValue(value any) bool {
	switch enumValue := value.(type) {
	case *haxe__StackItem:
		return (enumValue != nil)
	case *haxe___Template__TemplateExpr:
		return (enumValue != nil)
	case *haxe__io__Encoding:
		return (enumValue != nil)
	case *haxe__io__Error:
		return (enumValue != nil)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_callMethod(value any, arguments []any) any {
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Cannot call null"))
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() any]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(*string) any]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_1 any) *string {
			if hx_value_1 == nil {
				var hx_zero_2 *string
				return hx_zero_2
			}
			return hx_value_1.(*string)
		}(argument_0))
	case *hxrt.Function[func(...any) any]:
		var argument_0 []any = nil
		for index, argument := range arguments {
			if index >= 0 {
				argument_0 = append(argument_0, argument)
			}
		}
		return hxrt.FunctionCall(source)(argument_0...)
	case *hxrt.Function[func(any, any) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(argument_0, argument_1)
	case *hxrt.Function[func(*string) *hxrt.Array]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_3 any) *string {
			if hx_value_3 == nil {
				var hx_zero_4 *string
				return hx_zero_4
			}
			return hx_value_3.(*string)
		}(argument_0))
	case *hxrt.Function[func(*haxe___Template__TokenCursor) *haxe___Template__TemplateExpr]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_haxe___Template__TokenCursor(argument_0))
	case *hxrt.Function[func(*string) *hxrt.Function[func() any]]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_5 any) *string {
			if hx_value_5 == nil {
				var hx_zero_6 *string
				return hx_zero_6
			}
			return hx_value_5.(*string)
		}(argument_0))
	case *hxrt.Function[func(*hxrt.Function[func() any], *haxe___Template__ExprCursor) *hxrt.Function[func() any]]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(hx_recover_callable_c5309b43(argument_0), hx_recover_class_haxe___Template__ExprCursor(argument_1))
	case *hxrt.Function[func(*haxe___Template__ExprCursor) *hxrt.Function[func() any]]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_haxe___Template__ExprCursor(argument_0))
	case *hxrt.Function[func(*haxe___Template__ExprCursor)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(hx_recover_class_haxe___Template__ExprCursor(argument_0))
		return nil
	case *hxrt.Function[func(*haxe___Template__TemplateExpr)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(func(hx_value_7 any) *haxe___Template__TemplateExpr {
			if hx_value_7 == nil {
				var hx_zero_8 *haxe___Template__TemplateExpr
				return hx_zero_8
			}
			return hx_value_7.(*haxe___Template__TemplateExpr)
		}(argument_0))
		return nil
	case *hxrt.Function[func() bool]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() int]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *string]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(*string) bool]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_9 any) *string {
			if hx_value_9 == nil {
				var hx_zero_10 *string
				return hx_zero_10
			}
			return hx_value_9.(*string)
		}(argument_0))
	case *hxrt.Function[func(int) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_11 any) int {
			if hx_value_11 == nil {
				var hx_zero_12 int
				return hx_zero_12
			}
			return hx_value_11.(int)
		}(argument_0))
	case *hxrt.Function[func(*string, int, any) bool]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		var argument_2 any = nil
		if 2 < len(arguments) {
			argument_2 = arguments[2]
		}
		return hxrt.FunctionCall(source)(func(hx_value_13 any) *string {
			if hx_value_13 == nil {
				var hx_zero_14 *string
				return hx_zero_14
			}
			return hx_value_13.(*string)
		}(argument_0), func(hx_value_15 any) int {
			if hx_value_15 == nil {
				var hx_zero_16 int
				return hx_zero_16
			}
			return hx_value_15.(int)
		}(argument_1), argument_2)
	case *hxrt.Function[func(*string, *string) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(func(hx_value_17 any) *string {
			if hx_value_17 == nil {
				var hx_zero_18 *string
				return hx_zero_18
			}
			return hx_value_17.(*string)
		}(argument_0), func(hx_value_19 any) *string {
			if hx_value_19 == nil {
				var hx_zero_20 *string
				return hx_zero_20
			}
			return hx_value_19.(*string)
		}(argument_1))
	case *hxrt.Function[func(*string, *hxrt.Function[func(*EReg) *string]) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(func(hx_value_21 any) *string {
			if hx_value_21 == nil {
				var hx_zero_22 *string
				return hx_zero_22
			}
			return hx_value_21.(*string)
		}(argument_0), hx_recover_callable_97658df1(argument_1))
	case *hxrt.Function[func(*string, *hxrt.RegexMatch)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_23 any) *string {
			if hx_value_23 == nil {
				var hx_zero_24 *string
				return hx_zero_24
			}
			return hx_value_23.(*string)
		}(argument_0), func(hx_value_25 any) *hxrt.RegexMatch {
			if hx_value_25 == nil {
				var hx_zero_26 *hxrt.RegexMatch
				return hx_zero_26
			}
			return hx_value_25.(*hxrt.RegexMatch)
		}(argument_1))
		return nil
	case *hxrt.Function[func() *hxrt.RegexMatch]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(*string, *string, *hxrt.RegexMatch) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		var argument_2 any = nil
		if 2 < len(arguments) {
			argument_2 = arguments[2]
		}
		return hxrt.FunctionCall(source)(func(hx_value_27 any) *string {
			if hx_value_27 == nil {
				var hx_zero_28 *string
				return hx_zero_28
			}
			return hx_value_27.(*string)
		}(argument_0), func(hx_value_29 any) *string {
			if hx_value_29 == nil {
				var hx_zero_30 *string
				return hx_zero_30
			}
			return hx_value_29.(*string)
		}(argument_1), func(hx_value_31 any) *hxrt.RegexMatch {
			if hx_value_31 == nil {
				var hx_zero_32 *hxrt.RegexMatch
				return hx_zero_32
			}
			return hx_value_31.(*hxrt.RegexMatch)
		}(argument_2))
	case *hxrt.Function[func(int) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_33 any) int {
			if hx_value_33 == nil {
				var hx_zero_34 int
				return hx_zero_34
			}
			return hx_value_33.(int)
		}(argument_0))
	case *hxrt.Function[func(int, int)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_35 any) int {
			if hx_value_35 == nil {
				var hx_zero_36 int
				return hx_zero_36
			}
			return hx_value_35.(int)
		}(argument_0), func(hx_value_37 any) int {
			if hx_value_37 == nil {
				var hx_zero_38 int
				return hx_zero_38
			}
			return hx_value_37.(int)
		}(argument_1))
		return nil
	case *hxrt.Function[func(int, *haxe__io__Bytes, int, int)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		var argument_2 any = nil
		if 2 < len(arguments) {
			argument_2 = arguments[2]
		}
		var argument_3 any = nil
		if 3 < len(arguments) {
			argument_3 = arguments[3]
		}
		hxrt.FunctionCall(source)(func(hx_value_39 any) int {
			if hx_value_39 == nil {
				var hx_zero_40 int
				return hx_zero_40
			}
			return hx_value_39.(int)
		}(argument_0), hx_recover_class_haxe__io__Bytes(argument_1), func(hx_value_41 any) int {
			if hx_value_41 == nil {
				var hx_zero_42 int
				return hx_zero_42
			}
			return hx_value_41.(int)
		}(argument_2), func(hx_value_43 any) int {
			if hx_value_43 == nil {
				var hx_zero_44 int
				return hx_zero_44
			}
			return hx_value_43.(int)
		}(argument_3))
		return nil
	case *hxrt.Function[func(int, int, int)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		var argument_2 any = nil
		if 2 < len(arguments) {
			argument_2 = arguments[2]
		}
		hxrt.FunctionCall(source)(func(hx_value_45 any) int {
			if hx_value_45 == nil {
				var hx_zero_46 int
				return hx_zero_46
			}
			return hx_value_45.(int)
		}(argument_0), func(hx_value_47 any) int {
			if hx_value_47 == nil {
				var hx_zero_48 int
				return hx_zero_48
			}
			return hx_value_47.(int)
		}(argument_1), func(hx_value_49 any) int {
			if hx_value_49 == nil {
				var hx_zero_50 int
				return hx_zero_50
			}
			return hx_value_49.(int)
		}(argument_2))
		return nil
	case *hxrt.Function[func(int, int) *haxe__io__Bytes]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(func(hx_value_51 any) int {
			if hx_value_51 == nil {
				var hx_zero_52 int
				return hx_zero_52
			}
			return hx_value_51.(int)
		}(argument_0), func(hx_value_53 any) int {
			if hx_value_53 == nil {
				var hx_zero_54 int
				return hx_zero_54
			}
			return hx_value_53.(int)
		}(argument_1))
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_haxe__io__Bytes(argument_0))
	case *hxrt.Function[func(int) float64]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_55 any) int {
			if hx_value_55 == nil {
				var hx_zero_56 int
				return hx_zero_56
			}
			return hx_value_55.(int)
		}(argument_0))
	case *hxrt.Function[func(int, float64)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_57 any) int {
			if hx_value_57 == nil {
				var hx_zero_58 int
				return hx_zero_58
			}
			return hx_value_57.(int)
		}(argument_0), func(hx_value_59 any) float64 {
			if hx_value_59 == nil {
				var hx_zero_60 float64
				return hx_zero_60
			}
			return hx_value_59.(float64)
		}(argument_1))
		return nil
	case *hxrt.Function[func(int) *haxe___Int64_____Int64]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_61 any) int {
			if hx_value_61 == nil {
				var hx_zero_62 int
				return hx_zero_62
			}
			return hx_value_61.(int)
		}(argument_0))
	case *hxrt.Function[func(int, *haxe___Int64_____Int64)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_63 any) int {
			if hx_value_63 == nil {
				var hx_zero_64 int
				return hx_zero_64
			}
			return hx_value_63.(int)
		}(argument_0), hx_recover_class_haxe___Int64_____Int64(argument_1))
		return nil
	case *hxrt.Function[func(int, int, *haxe__io__Encoding) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		var argument_2 any = nil
		if 2 < len(arguments) {
			argument_2 = arguments[2]
		}
		return hxrt.FunctionCall(source)(func(hx_value_65 any) int {
			if hx_value_65 == nil {
				var hx_zero_66 int
				return hx_zero_66
			}
			return hx_value_65.(int)
		}(argument_0), func(hx_value_67 any) int {
			if hx_value_67 == nil {
				var hx_zero_68 int
				return hx_zero_68
			}
			return hx_value_67.(int)
		}(argument_1), func(hx_value_69 any) *haxe__io__Encoding {
			if hx_value_69 == nil {
				var hx_zero_70 *haxe__io__Encoding
				return hx_zero_70
			}
			return hx_value_69.(*haxe__io__Encoding)
		}(argument_2))
	case *hxrt.Function[func(int, int) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(func(hx_value_71 any) int {
			if hx_value_71 == nil {
				var hx_zero_72 int
				return hx_zero_72
			}
			return hx_value_71.(int)
		}(argument_0), func(hx_value_73 any) int {
			if hx_value_73 == nil {
				var hx_zero_74 int
				return hx_zero_74
			}
			return hx_value_73.(int)
		}(argument_1))
	case *hxrt.Function[func() []int]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.FunctionCall(source)()
	default:
		return hxrt.ReflectCallMethod(value, arguments)
	}
}

func hx_recover_callable_97658df1(value any) *hxrt.Function[func(*EReg) *string] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(*EReg) *string]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(*string) any]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_3 any) *string {
				if hx_value_3 == nil {
					var hx_zero_4 *string
					return hx_zero_4
				}
				return hx_value_3.(*string)
			}(hxrt.FunctionCall(source)(func(hx_value_1 any) *string {
				if hx_value_1 == nil {
					var hx_zero_2 *string
					return hx_zero_2
				}
				return hx_value_1.(*string)
			}(arg_0)))
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_5 any) *string {
				if hx_value_5 == nil {
					var hx_zero_6 *string
					return hx_zero_6
				}
				return hx_value_5.(*string)
			}(hxrt.FunctionCall(source)(arg_0))
		})
	case *hxrt.Function[func(*string) *hxrt.Array]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_9 any) *string {
				if hx_value_9 == nil {
					var hx_zero_10 *string
					return hx_zero_10
				}
				return hx_value_9.(*string)
			}(hxrt.FunctionCall(source)(func(hx_value_7 any) *string {
				if hx_value_7 == nil {
					var hx_zero_8 *string
					return hx_zero_8
				}
				return hx_value_7.(*string)
			}(arg_0)))
		})
	case *hxrt.Function[func(*haxe___Template__TokenCursor) *haxe___Template__TemplateExpr]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_11 any) *string {
				if hx_value_11 == nil {
					var hx_zero_12 *string
					return hx_zero_12
				}
				return hx_value_11.(*string)
			}(hxrt.FunctionCall(source)(hx_recover_class_haxe___Template__TokenCursor(arg_0)))
		})
	case *hxrt.Function[func(*string) *hxrt.Function[func() any]]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_15 any) *string {
				if hx_value_15 == nil {
					var hx_zero_16 *string
					return hx_zero_16
				}
				return hx_value_15.(*string)
			}(hxrt.FunctionCall(source)(func(hx_value_13 any) *string {
				if hx_value_13 == nil {
					var hx_zero_14 *string
					return hx_zero_14
				}
				return hx_value_13.(*string)
			}(arg_0)))
		})
	case *hxrt.Function[func(*haxe___Template__ExprCursor) *hxrt.Function[func() any]]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_17 any) *string {
				if hx_value_17 == nil {
					var hx_zero_18 *string
					return hx_zero_18
				}
				return hx_value_17.(*string)
			}(hxrt.FunctionCall(source)(hx_recover_class_haxe___Template__ExprCursor(arg_0)))
		})
	case *hxrt.Function[func(*haxe___Template__ExprCursor)]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			hxrt.FunctionCall(source)(hx_recover_class_haxe___Template__ExprCursor(arg_0))
			return func(hx_value_19 any) *string {
				if hx_value_19 == nil {
					var hx_zero_20 *string
					return hx_zero_20
				}
				return hx_value_19.(*string)
			}(nil)
		})
	case *hxrt.Function[func(*haxe___Template__TemplateExpr)]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			hxrt.FunctionCall(source)(func(hx_value_21 any) *haxe___Template__TemplateExpr {
				if hx_value_21 == nil {
					var hx_zero_22 *haxe___Template__TemplateExpr
					return hx_zero_22
				}
				return hx_value_21.(*haxe___Template__TemplateExpr)
			}(arg_0))
			return func(hx_value_23 any) *string {
				if hx_value_23 == nil {
					var hx_zero_24 *string
					return hx_zero_24
				}
				return hx_value_23.(*string)
			}(nil)
		})
	case *hxrt.Function[func(*string) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_27 any) *string {
				if hx_value_27 == nil {
					var hx_zero_28 *string
					return hx_zero_28
				}
				return hx_value_27.(*string)
			}(hxrt.FunctionCall(source)(func(hx_value_25 any) *string {
				if hx_value_25 == nil {
					var hx_zero_26 *string
					return hx_zero_26
				}
				return hx_value_25.(*string)
			}(arg_0)))
		})
	case *hxrt.Function[func(int) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return hxrt.FunctionCall(source)(func(hx_value_29 any) int {
				if hx_value_29 == nil {
					var hx_zero_30 int
					return hx_zero_30
				}
				return hx_value_29.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_33 any) *string {
				if hx_value_33 == nil {
					var hx_zero_34 *string
					return hx_zero_34
				}
				return hx_value_33.(*string)
			}(hxrt.FunctionCall(source)(func(hx_value_31 any) int {
				if hx_value_31 == nil {
					var hx_zero_32 int
					return hx_zero_32
				}
				return hx_value_31.(int)
			}(arg_0)))
		})
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_35 any) *string {
				if hx_value_35 == nil {
					var hx_zero_36 *string
					return hx_zero_36
				}
				return hx_value_35.(*string)
			}(hxrt.FunctionCall(source)(hx_recover_class_haxe__io__Bytes(arg_0)))
		})
	case *hxrt.Function[func(int) float64]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_39 any) *string {
				if hx_value_39 == nil {
					var hx_zero_40 *string
					return hx_zero_40
				}
				return hx_value_39.(*string)
			}(hxrt.FunctionCall(source)(func(hx_value_37 any) int {
				if hx_value_37 == nil {
					var hx_zero_38 int
					return hx_zero_38
				}
				return hx_value_37.(int)
			}(arg_0)))
		})
	case *hxrt.Function[func(int) *haxe___Int64_____Int64]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_43 any) *string {
				if hx_value_43 == nil {
					var hx_zero_44 *string
					return hx_zero_44
				}
				return hx_value_43.(*string)
			}(hxrt.FunctionCall(source)(func(hx_value_41 any) int {
				if hx_value_41 == nil {
					var hx_zero_42 int
					return hx_zero_42
				}
				return hx_value_41.(int)
			}(arg_0)))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_c5309b43(value any) *hxrt.Function[func() any] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() any]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() any]:
		return source
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *hxrt.RegexMatch]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() []int]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_class_haxe__Exception(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__Exception:
		return source
	case *haxe__ValueException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}

func hx_recover_class_haxe__ValueException(value any) *haxe__ValueException {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ValueException:
		return source
	default:
		return value.(*haxe__ValueException)
	}
}

func hx_recover_class_haxe___Int64_____Int64(value any) *haxe___Int64_____Int64 {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe___Int64_____Int64:
		return source
	default:
		return value.(*haxe___Int64_____Int64)
	}
}

func hx_recover_class_haxe___Template__ExprCursor(value any) *haxe___Template__ExprCursor {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe___Template__ExprCursor:
		return source
	default:
		return value.(*haxe___Template__ExprCursor)
	}
}

func hx_recover_class_haxe___Template__TokenCursor(value any) *haxe___Template__TokenCursor {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe___Template__TokenCursor:
		return source
	default:
		return value.(*haxe___Template__TokenCursor)
	}
}

func hx_recover_class_haxe__io__Bytes(value any) *haxe__io__Bytes {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__io__Bytes:
		return source
	default:
		return value.(*haxe__io__Bytes)
	}
}

func hx_is_nominal_haxe__Exception(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__Exception:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__ValueException(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}
