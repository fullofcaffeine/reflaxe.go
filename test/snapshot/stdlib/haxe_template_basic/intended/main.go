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
		return func(hx_structural_source_303 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_303, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_303.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_303.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_303.high = func(hx_value_304 any) int {
						if hx_value_304 == nil {
							var hx_zero_305 int
							return hx_zero_305
						}
						return hx_value_304.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_303.low = func(hx_value_306 any) int {
						if hx_value_306 == nil {
							var hx_zero_307 int
							return hx_zero_307
						}
						return hx_value_306.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe___Template__TokenCursor:
		return func(hx_structural_source_308 *haxe___Template__TokenCursor) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_308, func(name string) hxrt.AnonymousField {
				switch name {
				case "tokens":
					return hxrt.AnonymousField{Value: hx_structural_source_308.tokens, Found: true}
				case "index":
					return hxrt.AnonymousField{Value: hx_structural_source_308.index, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "tokens":
					hx_structural_source_308.tokens = func(hx_value_309 any) *hxrt.Array {
						if hx_value_309 == nil {
							var hx_zero_310 *hxrt.Array
							return hx_zero_310
						}
						return hx_value_309.(*hxrt.Array)
					}(incoming)
					return true
				case "index":
					hx_structural_source_308.index = func(hx_value_311 any) int {
						if hx_value_311 == nil {
							var hx_zero_312 int
							return hx_zero_312
						}
						return hx_value_311.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe___Template__ExprCursor:
		return func(hx_structural_source_313 *haxe___Template__ExprCursor) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_313, func(name string) hxrt.AnonymousField {
				switch name {
				case "tokens":
					return hxrt.AnonymousField{Value: hx_structural_source_313.tokens, Found: true}
				case "index":
					return hxrt.AnonymousField{Value: hx_structural_source_313.index, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "tokens":
					hx_structural_source_313.tokens = func(hx_value_314 any) *hxrt.Array {
						if hx_value_314 == nil {
							var hx_zero_315 *hxrt.Array
							return hx_zero_315
						}
						return hx_value_314.(*hxrt.Array)
					}(incoming)
					return true
				case "index":
					hx_structural_source_313.index = func(hx_value_316 any) int {
						if hx_value_316 == nil {
							var hx_zero_317 int
							return hx_zero_317
						}
						return hx_value_316.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__Template:
		return func(hx_structural_source_318 *haxe__Template) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_318, func(name string) hxrt.AnonymousField {
				switch name {
				case "expr":
					return hxrt.AnonymousField{Value: hx_structural_source_318.expr, Found: true}
				case "context":
					return hxrt.AnonymousField{Value: hx_structural_source_318.context, Found: true}
				case "macros":
					return hxrt.AnonymousField{Value: hx_structural_source_318.macros, Found: true}
				case "stack":
					return hxrt.AnonymousField{Value: hx_structural_source_318.stack, Found: true}
				case "output":
					return hxrt.AnonymousField{Value: hx_structural_source_318.output, Found: true}
				case "execute":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any, any) *string] {
						hx_receiver_325 := hx_structural_source_318
						return hxrt.BoundFunction(func(context any, macros any) *string {
							return hx_receiver_325.__hx_this.execute(context, macros)
						}, func() any {
							return hx_receiver_325.__hx_this
						}, "execute")
					}(), Found: true}
				case "resolve":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) any] {
						hx_receiver_326 := hx_structural_source_318
						return hxrt.BoundFunction(func(v *string) any {
							return hx_receiver_326.__hx_this.resolve(v)
						}, func() any {
							return hx_receiver_326.__hx_this
						}, "resolve")
					}(), Found: true}
				case "parseTokens":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) *hxrt.Array] {
						hx_receiver_327 := hx_structural_source_318
						return hxrt.BoundFunction(func(data *string) *hxrt.Array {
							return hx_receiver_327.__hx_this.parseTokens(data)
						}, func() any {
							return hx_receiver_327.__hx_this
						}, "parseTokens")
					}(), Found: true}
				case "parseBlock":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe___Template__TokenCursor) *haxe___Template__TemplateExpr] {
						hx_receiver_328 := hx_structural_source_318
						return hxrt.BoundFunction(func(cursor *haxe___Template__TokenCursor) *haxe___Template__TemplateExpr {
							return hx_receiver_328.__hx_this.parseBlock(cursor)
						}, func() any {
							return hx_receiver_328.__hx_this
						}, "parseBlock")
					}(), Found: true}
				case "parse":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe___Template__TokenCursor) *haxe___Template__TemplateExpr] {
						hx_receiver_329 := hx_structural_source_318
						return hxrt.BoundFunction(func(cursor *haxe___Template__TokenCursor) *haxe___Template__TemplateExpr {
							return hx_receiver_329.__hx_this.parse(cursor)
						}, func() any {
							return hx_receiver_329.__hx_this
						}, "parse")
					}(), Found: true}
				case "parseExpr":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) *hxrt.Function[func() any]] {
						hx_receiver_330 := hx_structural_source_318
						return hxrt.BoundFunction(func(data *string) *hxrt.Function[func() any] {
							return hx_receiver_330.__hx_this.parseExpr(data)
						}, func() any {
							return hx_receiver_330.__hx_this
						}, "parseExpr")
					}(), Found: true}
				case "makeConst":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) *hxrt.Function[func() any]] {
						hx_receiver_331 := hx_structural_source_318
						return hxrt.BoundFunction(func(v *string) *hxrt.Function[func() any] {
							return hx_receiver_331.__hx_this.makeConst(v)
						}, func() any {
							return hx_receiver_331.__hx_this
						}, "makeConst")
					}(), Found: true}
				case "makePath":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func() any], *haxe___Template__ExprCursor) *hxrt.Function[func() any]] {
						hx_receiver_332 := hx_structural_source_318
						return hxrt.BoundFunction(func(e *hxrt.Function[func() any], cursor *haxe___Template__ExprCursor) *hxrt.Function[func() any] {
							return hx_receiver_332.__hx_this.makePath(e, cursor)
						}, func() any {
							return hx_receiver_332.__hx_this
						}, "makePath")
					}(), Found: true}
				case "makeExpr":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe___Template__ExprCursor) *hxrt.Function[func() any]] {
						hx_receiver_333 := hx_structural_source_318
						return hxrt.BoundFunction(func(cursor *haxe___Template__ExprCursor) *hxrt.Function[func() any] {
							return hx_receiver_333.__hx_this.makeExpr(cursor)
						}, func() any {
							return hx_receiver_333.__hx_this
						}, "makeExpr")
					}(), Found: true}
				case "skipSpaces":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe___Template__ExprCursor)] {
						hx_receiver_334 := hx_structural_source_318
						return hxrt.BoundFunction(func(cursor *haxe___Template__ExprCursor) {
							hx_receiver_334.__hx_this.skipSpaces(cursor)
						}, func() any {
							return hx_receiver_334.__hx_this
						}, "skipSpaces")
					}(), Found: true}
				case "makeExpr2":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe___Template__ExprCursor) *hxrt.Function[func() any]] {
						hx_receiver_335 := hx_structural_source_318
						return hxrt.BoundFunction(func(cursor *haxe___Template__ExprCursor) *hxrt.Function[func() any] {
							return hx_receiver_335.__hx_this.makeExpr2(cursor)
						}, func() any {
							return hx_receiver_335.__hx_this
						}, "makeExpr2")
					}(), Found: true}
				case "run":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe___Template__TemplateExpr)] {
						hx_receiver_336 := hx_structural_source_318
						return hxrt.BoundFunction(func(e *haxe___Template__TemplateExpr) {
							hx_receiver_336.__hx_this.run(e)
						}, func() any {
							return hx_receiver_336.__hx_this
						}, "run")
					}(), Found: true}
				case "popStackValue":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_337 := hx_structural_source_318
						return hxrt.BoundFunction(func() any {
							return hx_receiver_337.__hx_this.popStackValue()
						}, func() any {
							return hx_receiver_337.__hx_this
						}, "popStackValue")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "expr":
					hx_structural_source_318.expr = func(hx_value_319 any) *haxe___Template__TemplateExpr {
						if hx_value_319 == nil {
							var hx_zero_320 *haxe___Template__TemplateExpr
							return hx_zero_320
						}
						return hx_value_319.(*haxe___Template__TemplateExpr)
					}(incoming)
					return true
				case "context":
					hx_structural_source_318.context = incoming
					return true
				case "macros":
					hx_structural_source_318.macros = incoming
					return true
				case "stack":
					hx_structural_source_318.stack = func(hx_value_321 any) *hxrt.Array {
						if hx_value_321 == nil {
							var hx_zero_322 *hxrt.Array
							return hx_zero_322
						}
						return hx_value_321.(*hxrt.Array)
					}(incoming)
					return true
				case "output":
					hx_structural_source_318.output = func(hx_value_323 any) *string {
						if hx_value_323 == nil {
							var hx_zero_324 *string
							return hx_zero_324
						}
						return hx_value_323.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__StringKeyValueIterator:
		return func(hx_structural_source_338 *haxe__iterators__StringKeyValueIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_338, func(name string) hxrt.AnonymousField {
				switch name {
				case "offset":
					return hxrt.AnonymousField{Value: hx_structural_source_338.offset, Found: true}
				case "s":
					return hxrt.AnonymousField{Value: hx_structural_source_338.s, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_343 := hx_structural_source_338
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_343.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_343.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_344 := hx_structural_source_338
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_344.__hx_this.next()
						}, func() any {
							return hx_receiver_344.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "offset":
					hx_structural_source_338.offset = func(hx_value_339 any) int {
						if hx_value_339 == nil {
							var hx_zero_340 int
							return hx_zero_340
						}
						return hx_value_339.(int)
					}(incoming)
					return true
				case "s":
					hx_structural_source_338.s = func(hx_value_341 any) *string {
						if hx_value_341 == nil {
							var hx_zero_342 *string
							return hx_zero_342
						}
						return hx_value_341.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__StringIterator:
		return func(hx_structural_source_345 *haxe__iterators__StringIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_345, func(name string) hxrt.AnonymousField {
				switch name {
				case "offset":
					return hxrt.AnonymousField{Value: hx_structural_source_345.offset, Found: true}
				case "s":
					return hxrt.AnonymousField{Value: hx_structural_source_345.s, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_350 := hx_structural_source_345
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_350.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_350.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_351 := hx_structural_source_345
						return hxrt.BoundFunction(func() int {
							return hx_receiver_351.__hx_this.next()
						}, func() any {
							return hx_receiver_351.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "offset":
					hx_structural_source_345.offset = func(hx_value_346 any) int {
						if hx_value_346 == nil {
							var hx_zero_347 int
							return hx_zero_347
						}
						return hx_value_346.(int)
					}(incoming)
					return true
				case "s":
					hx_structural_source_345.s = func(hx_value_348 any) *string {
						if hx_value_348 == nil {
							var hx_zero_349 *string
							return hx_zero_349
						}
						return hx_value_348.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__Exception:
		return func(hx_structural_source_352 *haxe__Exception) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_352, func(name string) hxrt.AnonymousField {
				switch name {
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_352.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_352.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_352.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_352.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_355 := hx_structural_source_352
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_355.__hx_this.get_message()
						}, func() any {
							return hx_receiver_355.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_356 := hx_structural_source_352
						return hxrt.BoundFunction(func() any {
							return hx_receiver_356.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_356.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_357 := hx_structural_source_352
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_357.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_357.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_358 := hx_structural_source_352
						return hxrt.BoundFunction(func() any {
							return hx_receiver_358.__hx_this.get_native()
						}, func() any {
							return hx_receiver_358.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_359 := hx_structural_source_352
						return hxrt.BoundFunction(func() any {
							return hx_receiver_359.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_359.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_360 := hx_structural_source_352
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_360.__hx_this.toString()
						}, func() any {
							return hx_receiver_360.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_361 := hx_structural_source_352
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_361.__hx_this.details()
						}, func() any {
							return hx_receiver_361.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_352.savedMessage = func(hx_value_353 any) *string {
						if hx_value_353 == nil {
							var hx_zero_354 *string
							return hx_zero_354
						}
						return hx_value_353.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_352.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_352.savedPrevious = hx_recover_class_1(incoming)
					return true
				case "savedNative":
					hx_structural_source_352.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *EReg:
		return func(hx_structural_source_362 *EReg) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_362, func(name string) hxrt.AnonymousField {
				switch name {
				case "handle":
					return hxrt.AnonymousField{Value: hx_structural_source_362.handle, Found: true}
				case "global":
					return hxrt.AnonymousField{Value: hx_structural_source_362.global, Found: true}
				case "lastSource":
					return hxrt.AnonymousField{Value: hx_structural_source_362.lastSource, Found: true}
				case "lastMatch":
					return hxrt.AnonymousField{Value: hx_structural_source_362.lastMatch, Found: true}
				case "match":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) bool] {
						hx_receiver_371 := hx_structural_source_362
						return hxrt.BoundFunction(func(s *string) bool {
							return hx_receiver_371.__hx_this.match(s)
						}, func() any {
							return hx_receiver_371.__hx_this
						}, "match")
					}(), Found: true}
				case "matched":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *string] {
						hx_receiver_372 := hx_structural_source_362
						return hxrt.BoundFunction(func(n int) *string {
							return hx_receiver_372.__hx_this.matched(n)
						}, func() any {
							return hx_receiver_372.__hx_this
						}, "matched")
					}(), Found: true}
				case "matchedLeft":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_373 := hx_structural_source_362
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_373.__hx_this.matchedLeft()
						}, func() any {
							return hx_receiver_373.__hx_this
						}, "matchedLeft")
					}(), Found: true}
				case "matchedRight":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_374 := hx_structural_source_362
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_374.__hx_this.matchedRight()
						}, func() any {
							return hx_receiver_374.__hx_this
						}, "matchedRight")
					}(), Found: true}
				case "matchedPos":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_375 := hx_structural_source_362
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_375.__hx_this.matchedPos()
						}, func() any {
							return hx_receiver_375.__hx_this
						}, "matchedPos")
					}(), Found: true}
				case "matchSub":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, int, any) bool] {
						hx_receiver_376 := hx_structural_source_362
						return hxrt.BoundFunction(func(s *string, pos int, len any) bool {
							if len == nil {
								len = -1
							}
							return hx_receiver_376.__hx_this.matchSub(s, pos, func(hx_value_377 any) int {
								if hx_value_377 == nil {
									var hx_zero_378 int
									return hx_zero_378
								}
								return hx_value_377.(int)
							}(len))
						}, func() any {
							return hx_receiver_376.__hx_this
						}, "matchSub")
					}(), Found: true}
				case "split":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) *hxrt.Array] {
						hx_receiver_379 := hx_structural_source_362
						return hxrt.BoundFunction(func(s *string) *hxrt.Array {
							return hx_receiver_379.__hx_this.split(s)
						}, func() any {
							return hx_receiver_379.__hx_this
						}, "split")
					}(), Found: true}
				case "replace":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string) *string] {
						hx_receiver_380 := hx_structural_source_362
						return hxrt.BoundFunction(func(s *string, by *string) *string {
							return hx_receiver_380.__hx_this.replace(s, by)
						}, func() any {
							return hx_receiver_380.__hx_this
						}, "replace")
					}(), Found: true}
				case "map":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *hxrt.Function[func(*EReg) *string]) *string] {
						hx_receiver_381 := hx_structural_source_362
						return hxrt.BoundFunction(func(s *string, f *hxrt.Function[func(*EReg) *string]) *string {
							return hx_receiver_381.__hx_this.map_(s, f)
						}, func() any {
							return hx_receiver_381.__hx_this
						}, "map")
					}(), Found: true}
				case "remember":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *hxrt.RegexMatch)] {
						hx_receiver_382 := hx_structural_source_362
						return hxrt.BoundFunction(func(source *string, match *hxrt.RegexMatch) {
							hx_receiver_382.__hx_this.remember(source, match)
						}, func() any {
							return hx_receiver_382.__hx_this
						}, "remember")
					}(), Found: true}
				case "requireMatch":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *hxrt.RegexMatch] {
						hx_receiver_383 := hx_structural_source_362
						return hxrt.BoundFunction(func() *hxrt.RegexMatch {
							return hx_receiver_383.__hx_this.requireMatch()
						}, func() any {
							return hx_receiver_383.__hx_this
						}, "requireMatch")
					}(), Found: true}
				case "expandReplacement":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string, *hxrt.RegexMatch) *string] {
						hx_receiver_384 := hx_structural_source_362
						return hxrt.BoundFunction(func(by *string, source *string, currentMatch *hxrt.RegexMatch) *string {
							return hx_receiver_384.__hx_this.expandReplacement(by, source, currentMatch)
						}, func() any {
							return hx_receiver_384.__hx_this
						}, "expandReplacement")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "handle":
					hx_structural_source_362.handle = func(hx_value_363 any) *hxrt.RegexHandle {
						if hx_value_363 == nil {
							var hx_zero_364 *hxrt.RegexHandle
							return hx_zero_364
						}
						return hx_value_363.(*hxrt.RegexHandle)
					}(incoming)
					return true
				case "global":
					hx_structural_source_362.global = func(hx_value_365 any) bool {
						if hx_value_365 == nil {
							var hx_zero_366 bool
							return hx_zero_366
						}
						return hx_value_365.(bool)
					}(incoming)
					return true
				case "lastSource":
					hx_structural_source_362.lastSource = func(hx_value_367 any) *string {
						if hx_value_367 == nil {
							var hx_zero_368 *string
							return hx_zero_368
						}
						return hx_value_367.(*string)
					}(incoming)
					return true
				case "lastMatch":
					hx_structural_source_362.lastMatch = func(hx_value_369 any) *hxrt.RegexMatch {
						if hx_value_369 == nil {
							var hx_zero_370 *hxrt.RegexMatch
							return hx_zero_370
						}
						return hx_value_369.(*hxrt.RegexMatch)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__ValueException:
		return func(hx_structural_source_385 *haxe__ValueException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_385, func(name string) hxrt.AnonymousField {
				switch name {
				case "value":
					return hxrt.AnonymousField{Value: hx_structural_source_385.value, Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_386 := hx_structural_source_385
						return hxrt.BoundFunction(func() any {
							return hx_receiver_386.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_386.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_385.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_385.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_385.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_385.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_389 := hx_structural_source_385
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_389.__hx_this.get_message()
						}, func() any {
							return hx_receiver_389.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_390 := hx_structural_source_385
						return hxrt.BoundFunction(func() any {
							return hx_receiver_390.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_390.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_391 := hx_structural_source_385
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_391.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_391.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_392 := hx_structural_source_385
						return hxrt.BoundFunction(func() any {
							return hx_receiver_392.__hx_this.get_native()
						}, func() any {
							return hx_receiver_392.__hx_this
						}, "get_native")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_393 := hx_structural_source_385
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_393.__hx_this.toString()
						}, func() any {
							return hx_receiver_393.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_394 := hx_structural_source_385
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_394.__hx_this.details()
						}, func() any {
							return hx_receiver_394.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_385.savedMessage = func(hx_value_387 any) *string {
						if hx_value_387 == nil {
							var hx_zero_388 *string
							return hx_zero_388
						}
						return hx_value_387.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_385.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_385.savedPrevious = hx_recover_class_1(incoming)
					return true
				case "savedNative":
					hx_structural_source_385.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__io__Bytes:
		return func(hx_structural_source_395 *haxe__io__Bytes) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_395, func(name string) hxrt.AnonymousField {
				switch name {
				case "length":
					return hxrt.AnonymousField{Value: hx_structural_source_395.length, Found: true}
				case "b":
					return hxrt.AnonymousField{Value: hx_structural_source_395.b, Found: true}
				case "__hx_raw":
					return hxrt.AnonymousField{Value: hx_structural_source_395.__hx_raw, Found: true}
				case "__hx_rawValid":
					return hxrt.AnonymousField{Value: hx_structural_source_395.__hx_rawValid, Found: true}
				case "__hx_dataExposed":
					return hxrt.AnonymousField{Value: hx_structural_source_395.__hx_dataExposed, Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_404 := hx_structural_source_395
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_404.__hx_this.get(pos)
						}, func() any {
							return hx_receiver_404.__hx_this
						}, "get")
					}(), Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_405 := hx_structural_source_395
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_405.__hx_this.set(pos, value)
						}, func() any {
							return hx_receiver_405.__hx_this
						}, "set")
					}(), Found: true}
				case "blit":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe__io__Bytes, int, int)] {
						hx_receiver_406 := hx_structural_source_395
						return hxrt.BoundFunction(func(pos int, src *haxe__io__Bytes, srcpos int, len int) {
							hx_receiver_406.__hx_this.blit(pos, src, srcpos, len)
						}, func() any {
							return hx_receiver_406.__hx_this
						}, "blit")
					}(), Found: true}
				case "fill":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int, int)] {
						hx_receiver_407 := hx_structural_source_395
						return hxrt.BoundFunction(func(pos int, len int, value int) {
							hx_receiver_407.__hx_this.fill(pos, len, value)
						}, func() any {
							return hx_receiver_407.__hx_this
						}, "fill")
					}(), Found: true}
				case "sub":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int) *haxe__io__Bytes] {
						hx_receiver_408 := hx_structural_source_395
						return hxrt.BoundFunction(func(pos int, len int) *haxe__io__Bytes {
							return hx_receiver_408.__hx_this.sub(pos, len)
						}, func() any {
							return hx_receiver_408.__hx_this
						}, "sub")
					}(), Found: true}
				case "compare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes) int] {
						hx_receiver_409 := hx_structural_source_395
						return hxrt.BoundFunction(func(other *haxe__io__Bytes) int {
							return hx_receiver_409.__hx_this.compare(other)
						}, func() any {
							return hx_receiver_409.__hx_this
						}, "compare")
					}(), Found: true}
				case "getDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) float64] {
						hx_receiver_410 := hx_structural_source_395
						return hxrt.BoundFunction(func(pos int) float64 {
							return hx_receiver_410.__hx_this.getDouble(pos)
						}, func() any {
							return hx_receiver_410.__hx_this
						}, "getDouble")
					}(), Found: true}
				case "getFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) float64] {
						hx_receiver_411 := hx_structural_source_395
						return hxrt.BoundFunction(func(pos int) float64 {
							return hx_receiver_411.__hx_this.getFloat(pos)
						}, func() any {
							return hx_receiver_411.__hx_this
						}, "getFloat")
					}(), Found: true}
				case "setDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, float64)] {
						hx_receiver_412 := hx_structural_source_395
						return hxrt.BoundFunction(func(pos int, value float64) {
							hx_receiver_412.__hx_this.setDouble(pos, value)
						}, func() any {
							return hx_receiver_412.__hx_this
						}, "setDouble")
					}(), Found: true}
				case "setFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, float64)] {
						hx_receiver_413 := hx_structural_source_395
						return hxrt.BoundFunction(func(pos int, value float64) {
							hx_receiver_413.__hx_this.setFloat(pos, value)
						}, func() any {
							return hx_receiver_413.__hx_this
						}, "setFloat")
					}(), Found: true}
				case "getUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_414 := hx_structural_source_395
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_414.__hx_this.getUInt16(pos)
						}, func() any {
							return hx_receiver_414.__hx_this
						}, "getUInt16")
					}(), Found: true}
				case "setUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_415 := hx_structural_source_395
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_415.__hx_this.setUInt16(pos, value)
						}, func() any {
							return hx_receiver_415.__hx_this
						}, "setUInt16")
					}(), Found: true}
				case "getInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_416 := hx_structural_source_395
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_416.__hx_this.getInt32(pos)
						}, func() any {
							return hx_receiver_416.__hx_this
						}, "getInt32")
					}(), Found: true}
				case "getInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *haxe___Int64_____Int64] {
						hx_receiver_417 := hx_structural_source_395
						return hxrt.BoundFunction(func(pos int) *haxe___Int64_____Int64 {
							return hx_receiver_417.__hx_this.getInt64(pos)
						}, func() any {
							return hx_receiver_417.__hx_this
						}, "getInt64")
					}(), Found: true}
				case "setInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_418 := hx_structural_source_395
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_418.__hx_this.setInt32(pos, value)
						}, func() any {
							return hx_receiver_418.__hx_this
						}, "setInt32")
					}(), Found: true}
				case "setInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe___Int64_____Int64)] {
						hx_receiver_419 := hx_structural_source_395
						return hxrt.BoundFunction(func(pos int, value *haxe___Int64_____Int64) {
							hx_receiver_419.__hx_this.setInt64(pos, value)
						}, func() any {
							return hx_receiver_419.__hx_this
						}, "setInt64")
					}(), Found: true}
				case "getString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int, *haxe__io__Encoding) *string] {
						hx_receiver_420 := hx_structural_source_395
						return hxrt.BoundFunction(func(pos int, len int, encoding *haxe__io__Encoding) *string {
							return hx_receiver_420.__hx_this.getString(pos, len, encoding)
						}, func() any {
							return hx_receiver_420.__hx_this
						}, "getString")
					}(), Found: true}
				case "readString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int) *string] {
						hx_receiver_421 := hx_structural_source_395
						return hxrt.BoundFunction(func(pos int, len int) *string {
							return hx_receiver_421.__hx_this.readString(pos, len)
						}, func() any {
							return hx_receiver_421.__hx_this
						}, "readString")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_422 := hx_structural_source_395
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_422.__hx_this.toString()
						}, func() any {
							return hx_receiver_422.__hx_this
						}, "toString")
					}(), Found: true}
				case "toHex":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_423 := hx_structural_source_395
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_423.__hx_this.toHex()
						}, func() any {
							return hx_receiver_423.__hx_this
						}, "toHex")
					}(), Found: true}
				case "getData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() []int] {
						hx_receiver_424 := hx_structural_source_395
						return hxrt.BoundFunction(func() []int {
							return hx_receiver_424.__hx_this.getData()
						}, func() any {
							return hx_receiver_424.__hx_this
						}, "getData")
					}(), Found: true}
				case "__hx_nativeView":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *hxrt.ByteView] {
						hx_receiver_425 := hx_structural_source_395
						return hxrt.BoundFunction(func() *hxrt.ByteView {
							return hx_receiver_425.__hx_this.__hx_nativeView()
						}, func() any {
							return hx_receiver_425.__hx_this
						}, "__hx_nativeView")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "b":
					hx_structural_source_395.b = func(hx_value_396 any) []int {
						if hx_value_396 == nil {
							var hx_zero_397 []int
							return hx_zero_397
						}
						return hx_value_396.([]int)
					}(incoming)
					return true
				case "__hx_raw":
					hx_structural_source_395.__hx_raw = func(hx_value_398 any) *hxrt.ByteView {
						if hx_value_398 == nil {
							var hx_zero_399 *hxrt.ByteView
							return hx_zero_399
						}
						return hx_value_398.(*hxrt.ByteView)
					}(incoming)
					return true
				case "__hx_rawValid":
					hx_structural_source_395.__hx_rawValid = func(hx_value_400 any) bool {
						if hx_value_400 == nil {
							var hx_zero_401 bool
							return hx_zero_401
						}
						return hx_value_400.(bool)
					}(incoming)
					return true
				case "__hx_dataExposed":
					hx_structural_source_395.__hx_dataExposed = func(hx_value_402 any) bool {
						if hx_value_402 == nil {
							var hx_zero_403 bool
							return hx_zero_403
						}
						return hx_value_402.(bool)
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
			hx_receiver_426 := value
			return hxrt.BoundFunction(func(by *string, source *string, currentMatch *hxrt.RegexMatch) *string {
				return hx_receiver_426.__hx_this.expandReplacement(by, source, currentMatch)
			}, func() any {
				return hx_receiver_426.__hx_this
			}, "expandReplacement")
		}()
	case "map":
		return func() *hxrt.Function[func(*string, *hxrt.Function[func(*EReg) *string]) *string] {
			hx_receiver_427 := value
			return hxrt.BoundFunction(func(s *string, f *hxrt.Function[func(*EReg) *string]) *string {
				return hx_receiver_427.__hx_this.map_(s, f)
			}, func() any {
				return hx_receiver_427.__hx_this
			}, "map")
		}()
	case "match":
		return func() *hxrt.Function[func(*string) bool] {
			hx_receiver_428 := value
			return hxrt.BoundFunction(func(s *string) bool {
				return hx_receiver_428.__hx_this.match(s)
			}, func() any {
				return hx_receiver_428.__hx_this
			}, "match")
		}()
	case "matchSub":
		return func() *hxrt.Function[func(*string, int, any) bool] {
			hx_receiver_429 := value
			return hxrt.BoundFunction(func(s *string, pos int, len any) bool {
				if len == nil {
					len = -1
				}
				return hx_receiver_429.__hx_this.matchSub(s, pos, func(hx_value_430 any) int {
					if hx_value_430 == nil {
						var hx_zero_431 int
						return hx_zero_431
					}
					return hx_value_430.(int)
				}(len))
			}, func() any {
				return hx_receiver_429.__hx_this
			}, "matchSub")
		}()
	case "matched":
		return func() *hxrt.Function[func(int) *string] {
			hx_receiver_432 := value
			return hxrt.BoundFunction(func(n int) *string {
				return hx_receiver_432.__hx_this.matched(n)
			}, func() any {
				return hx_receiver_432.__hx_this
			}, "matched")
		}()
	case "matchedLeft":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_433 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_433.__hx_this.matchedLeft()
			}, func() any {
				return hx_receiver_433.__hx_this
			}, "matchedLeft")
		}()
	case "matchedPos":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_434 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_434.__hx_this.matchedPos()
			}, func() any {
				return hx_receiver_434.__hx_this
			}, "matchedPos")
		}()
	case "matchedRight":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_435 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_435.__hx_this.matchedRight()
			}, func() any {
				return hx_receiver_435.__hx_this
			}, "matchedRight")
		}()
	case "remember":
		return func() *hxrt.Function[func(*string, *hxrt.RegexMatch)] {
			hx_receiver_436 := value
			return hxrt.BoundFunction(func(source *string, match *hxrt.RegexMatch) {
				hx_receiver_436.__hx_this.remember(source, match)
			}, func() any {
				return hx_receiver_436.__hx_this
			}, "remember")
		}()
	case "replace":
		return func() *hxrt.Function[func(*string, *string) *string] {
			hx_receiver_437 := value
			return hxrt.BoundFunction(func(s *string, by *string) *string {
				return hx_receiver_437.__hx_this.replace(s, by)
			}, func() any {
				return hx_receiver_437.__hx_this
			}, "replace")
		}()
	case "requireMatch":
		return func() *hxrt.Function[func() *hxrt.RegexMatch] {
			hx_receiver_438 := value
			return hxrt.BoundFunction(func() *hxrt.RegexMatch {
				return hx_receiver_438.__hx_this.requireMatch()
			}, func() any {
				return hx_receiver_438.__hx_this
			}, "requireMatch")
		}()
	case "split":
		return func() *hxrt.Function[func(*string) *hxrt.Array] {
			hx_receiver_439 := value
			return hxrt.BoundFunction(func(s *string) *hxrt.Array {
				return hx_receiver_439.__hx_this.split(s)
			}, func() any {
				return hx_receiver_439.__hx_this
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
			hx_receiver_440 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_440.__hx_this.details()
			}, func() any {
				return hx_receiver_440.__hx_this
			}, "details")
		}()
	case "get_message":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_441 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_441.__hx_this.get_message()
			}, func() any {
				return hx_receiver_441.__hx_this
			}, "get_message")
		}()
	case "get_native":
		return func() *hxrt.Function[func() any] {
			hx_receiver_442 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_442.__hx_this.get_native()
			}, func() any {
				return hx_receiver_442.__hx_this
			}, "get_native")
		}()
	case "get_previous":
		return func() *hxrt.Function[func() *haxe__Exception] {
			hx_receiver_443 := value
			return hxrt.BoundFunction(func() *haxe__Exception {
				return hx_receiver_443.__hx_this.get_previous()
			}, func() any {
				return hx_receiver_443.__hx_this
			}, "get_previous")
		}()
	case "get_stack":
		return func() *hxrt.Function[func() any] {
			hx_receiver_444 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_444.__hx_this.get_stack()
			}, func() any {
				return hx_receiver_444.__hx_this
			}, "get_stack")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_445 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_445.__hx_this.toString()
			}, func() any {
				return hx_receiver_445.__hx_this
			}, "toString")
		}()
	case "unwrap":
		return func() *hxrt.Function[func() any] {
			hx_receiver_446 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_446.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_446.__hx_this
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
			hx_receiver_447 := value
			return hxrt.BoundFunction(func(context any, macros any) *string {
				return hx_receiver_447.__hx_this.execute(context, macros)
			}, func() any {
				return hx_receiver_447.__hx_this
			}, "execute")
		}()
	case "makeConst":
		return func() *hxrt.Function[func(*string) *hxrt.Function[func() any]] {
			hx_receiver_448 := value
			return hxrt.BoundFunction(func(v *string) *hxrt.Function[func() any] {
				return hx_receiver_448.__hx_this.makeConst(v)
			}, func() any {
				return hx_receiver_448.__hx_this
			}, "makeConst")
		}()
	case "makeExpr":
		return func() *hxrt.Function[func(*haxe___Template__ExprCursor) *hxrt.Function[func() any]] {
			hx_receiver_449 := value
			return hxrt.BoundFunction(func(cursor *haxe___Template__ExprCursor) *hxrt.Function[func() any] {
				return hx_receiver_449.__hx_this.makeExpr(cursor)
			}, func() any {
				return hx_receiver_449.__hx_this
			}, "makeExpr")
		}()
	case "makeExpr2":
		return func() *hxrt.Function[func(*haxe___Template__ExprCursor) *hxrt.Function[func() any]] {
			hx_receiver_450 := value
			return hxrt.BoundFunction(func(cursor *haxe___Template__ExprCursor) *hxrt.Function[func() any] {
				return hx_receiver_450.__hx_this.makeExpr2(cursor)
			}, func() any {
				return hx_receiver_450.__hx_this
			}, "makeExpr2")
		}()
	case "makePath":
		return func() *hxrt.Function[func(*hxrt.Function[func() any], *haxe___Template__ExprCursor) *hxrt.Function[func() any]] {
			hx_receiver_451 := value
			return hxrt.BoundFunction(func(e *hxrt.Function[func() any], cursor *haxe___Template__ExprCursor) *hxrt.Function[func() any] {
				return hx_receiver_451.__hx_this.makePath(e, cursor)
			}, func() any {
				return hx_receiver_451.__hx_this
			}, "makePath")
		}()
	case "parse":
		return func() *hxrt.Function[func(*haxe___Template__TokenCursor) *haxe___Template__TemplateExpr] {
			hx_receiver_452 := value
			return hxrt.BoundFunction(func(cursor *haxe___Template__TokenCursor) *haxe___Template__TemplateExpr {
				return hx_receiver_452.__hx_this.parse(cursor)
			}, func() any {
				return hx_receiver_452.__hx_this
			}, "parse")
		}()
	case "parseBlock":
		return func() *hxrt.Function[func(*haxe___Template__TokenCursor) *haxe___Template__TemplateExpr] {
			hx_receiver_453 := value
			return hxrt.BoundFunction(func(cursor *haxe___Template__TokenCursor) *haxe___Template__TemplateExpr {
				return hx_receiver_453.__hx_this.parseBlock(cursor)
			}, func() any {
				return hx_receiver_453.__hx_this
			}, "parseBlock")
		}()
	case "parseExpr":
		return func() *hxrt.Function[func(*string) *hxrt.Function[func() any]] {
			hx_receiver_454 := value
			return hxrt.BoundFunction(func(data *string) *hxrt.Function[func() any] {
				return hx_receiver_454.__hx_this.parseExpr(data)
			}, func() any {
				return hx_receiver_454.__hx_this
			}, "parseExpr")
		}()
	case "parseTokens":
		return func() *hxrt.Function[func(*string) *hxrt.Array] {
			hx_receiver_455 := value
			return hxrt.BoundFunction(func(data *string) *hxrt.Array {
				return hx_receiver_455.__hx_this.parseTokens(data)
			}, func() any {
				return hx_receiver_455.__hx_this
			}, "parseTokens")
		}()
	case "popStackValue":
		return func() *hxrt.Function[func() any] {
			hx_receiver_456 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_456.__hx_this.popStackValue()
			}, func() any {
				return hx_receiver_456.__hx_this
			}, "popStackValue")
		}()
	case "resolve":
		return func() *hxrt.Function[func(*string) any] {
			hx_receiver_457 := value
			return hxrt.BoundFunction(func(v *string) any {
				return hx_receiver_457.__hx_this.resolve(v)
			}, func() any {
				return hx_receiver_457.__hx_this
			}, "resolve")
		}()
	case "run":
		return func() *hxrt.Function[func(*haxe___Template__TemplateExpr)] {
			hx_receiver_458 := value
			return hxrt.BoundFunction(func(e *haxe___Template__TemplateExpr) {
				hx_receiver_458.__hx_this.run(e)
			}, func() any {
				return hx_receiver_458.__hx_this
			}, "run")
		}()
	case "skipSpaces":
		return func() *hxrt.Function[func(*haxe___Template__ExprCursor)] {
			hx_receiver_459 := value
			return hxrt.BoundFunction(func(cursor *haxe___Template__ExprCursor) {
				hx_receiver_459.__hx_this.skipSpaces(cursor)
			}, func() any {
				return hx_receiver_459.__hx_this
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
			hx_receiver_460 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_460.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_460.__hx_this
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
			hx_receiver_461 := value
			return hxrt.BoundFunction(func() *hxrt.ByteView {
				return hx_receiver_461.__hx_this.__hx_nativeView()
			}, func() any {
				return hx_receiver_461.__hx_this
			}, "__hx_nativeView")
		}()
	case "blit":
		return func() *hxrt.Function[func(int, *haxe__io__Bytes, int, int)] {
			hx_receiver_462 := value
			return hxrt.BoundFunction(func(pos int, src *haxe__io__Bytes, srcpos int, len int) {
				hx_receiver_462.__hx_this.blit(pos, src, srcpos, len)
			}, func() any {
				return hx_receiver_462.__hx_this
			}, "blit")
		}()
	case "compare":
		return func() *hxrt.Function[func(*haxe__io__Bytes) int] {
			hx_receiver_463 := value
			return hxrt.BoundFunction(func(other *haxe__io__Bytes) int {
				return hx_receiver_463.__hx_this.compare(other)
			}, func() any {
				return hx_receiver_463.__hx_this
			}, "compare")
		}()
	case "fill":
		return func() *hxrt.Function[func(int, int, int)] {
			hx_receiver_464 := value
			return hxrt.BoundFunction(func(pos int, len int, value int) {
				hx_receiver_464.__hx_this.fill(pos, len, value)
			}, func() any {
				return hx_receiver_464.__hx_this
			}, "fill")
		}()
	case "get":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_465 := value
			return hxrt.BoundFunction(func(pos int) int {
				return hx_receiver_465.__hx_this.get(pos)
			}, func() any {
				return hx_receiver_465.__hx_this
			}, "get")
		}()
	case "getData":
		return func() *hxrt.Function[func() []int] {
			hx_receiver_466 := value
			return hxrt.BoundFunction(func() []int {
				return hx_receiver_466.__hx_this.getData()
			}, func() any {
				return hx_receiver_466.__hx_this
			}, "getData")
		}()
	case "getDouble":
		return func() *hxrt.Function[func(int) float64] {
			hx_receiver_467 := value
			return hxrt.BoundFunction(func(pos int) float64 {
				return hx_receiver_467.__hx_this.getDouble(pos)
			}, func() any {
				return hx_receiver_467.__hx_this
			}, "getDouble")
		}()
	case "getFloat":
		return func() *hxrt.Function[func(int) float64] {
			hx_receiver_468 := value
			return hxrt.BoundFunction(func(pos int) float64 {
				return hx_receiver_468.__hx_this.getFloat(pos)
			}, func() any {
				return hx_receiver_468.__hx_this
			}, "getFloat")
		}()
	case "getInt32":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_469 := value
			return hxrt.BoundFunction(func(pos int) int {
				return hx_receiver_469.__hx_this.getInt32(pos)
			}, func() any {
				return hx_receiver_469.__hx_this
			}, "getInt32")
		}()
	case "getInt64":
		return func() *hxrt.Function[func(int) *haxe___Int64_____Int64] {
			hx_receiver_470 := value
			return hxrt.BoundFunction(func(pos int) *haxe___Int64_____Int64 {
				return hx_receiver_470.__hx_this.getInt64(pos)
			}, func() any {
				return hx_receiver_470.__hx_this
			}, "getInt64")
		}()
	case "getString":
		return func() *hxrt.Function[func(int, int, *haxe__io__Encoding) *string] {
			hx_receiver_471 := value
			return hxrt.BoundFunction(func(pos int, len int, encoding *haxe__io__Encoding) *string {
				return hx_receiver_471.__hx_this.getString(pos, len, encoding)
			}, func() any {
				return hx_receiver_471.__hx_this
			}, "getString")
		}()
	case "getUInt16":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_472 := value
			return hxrt.BoundFunction(func(pos int) int {
				return hx_receiver_472.__hx_this.getUInt16(pos)
			}, func() any {
				return hx_receiver_472.__hx_this
			}, "getUInt16")
		}()
	case "readString":
		return func() *hxrt.Function[func(int, int) *string] {
			hx_receiver_473 := value
			return hxrt.BoundFunction(func(pos int, len int) *string {
				return hx_receiver_473.__hx_this.readString(pos, len)
			}, func() any {
				return hx_receiver_473.__hx_this
			}, "readString")
		}()
	case "set":
		return func() *hxrt.Function[func(int, int)] {
			hx_receiver_474 := value
			return hxrt.BoundFunction(func(pos int, value int) {
				hx_receiver_474.__hx_this.set(pos, value)
			}, func() any {
				return hx_receiver_474.__hx_this
			}, "set")
		}()
	case "setDouble":
		return func() *hxrt.Function[func(int, float64)] {
			hx_receiver_475 := value
			return hxrt.BoundFunction(func(pos int, value float64) {
				hx_receiver_475.__hx_this.setDouble(pos, value)
			}, func() any {
				return hx_receiver_475.__hx_this
			}, "setDouble")
		}()
	case "setFloat":
		return func() *hxrt.Function[func(int, float64)] {
			hx_receiver_476 := value
			return hxrt.BoundFunction(func(pos int, value float64) {
				hx_receiver_476.__hx_this.setFloat(pos, value)
			}, func() any {
				return hx_receiver_476.__hx_this
			}, "setFloat")
		}()
	case "setInt32":
		return func() *hxrt.Function[func(int, int)] {
			hx_receiver_477 := value
			return hxrt.BoundFunction(func(pos int, value int) {
				hx_receiver_477.__hx_this.setInt32(pos, value)
			}, func() any {
				return hx_receiver_477.__hx_this
			}, "setInt32")
		}()
	case "setInt64":
		return func() *hxrt.Function[func(int, *haxe___Int64_____Int64)] {
			hx_receiver_478 := value
			return hxrt.BoundFunction(func(pos int, value *haxe___Int64_____Int64) {
				hx_receiver_478.__hx_this.setInt64(pos, value)
			}, func() any {
				return hx_receiver_478.__hx_this
			}, "setInt64")
		}()
	case "setUInt16":
		return func() *hxrt.Function[func(int, int)] {
			hx_receiver_479 := value
			return hxrt.BoundFunction(func(pos int, value int) {
				hx_receiver_479.__hx_this.setUInt16(pos, value)
			}, func() any {
				return hx_receiver_479.__hx_this
			}, "setUInt16")
		}()
	case "sub":
		return func() *hxrt.Function[func(int, int) *haxe__io__Bytes] {
			hx_receiver_480 := value
			return hxrt.BoundFunction(func(pos int, len int) *haxe__io__Bytes {
				return hx_receiver_480.__hx_this.sub(pos, len)
			}, func() any {
				return hx_receiver_480.__hx_this
			}, "sub")
		}()
	case "toHex":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_481 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_481.__hx_this.toHex()
			}, func() any {
				return hx_receiver_481.__hx_this
			}, "toHex")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_482 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_482.__hx_this.toString()
			}, func() any {
				return hx_receiver_482.__hx_this
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
			hx_receiver_483 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_483.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_483.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() int] {
			hx_receiver_484 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_484.__hx_this.next()
			}, func() any {
				return hx_receiver_484.__hx_this
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
			hx_receiver_485 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_485.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_485.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_486 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_486.__hx_this.next()
			}, func() any {
				return hx_receiver_486.__hx_this
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
		return hxrt.FunctionCall(source)(func(hx_value_487 any) *string {
			if hx_value_487 == nil {
				var hx_zero_488 *string
				return hx_zero_488
			}
			return hx_value_487.(*string)
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
		return hxrt.FunctionCall(source)(func(hx_value_489 any) *string {
			if hx_value_489 == nil {
				var hx_zero_490 *string
				return hx_zero_490
			}
			return hx_value_489.(*string)
		}(argument_0))
	case *hxrt.Function[func(*haxe___Template__TokenCursor) *haxe___Template__TemplateExpr]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_2(argument_0))
	case *hxrt.Function[func(*string) *hxrt.Function[func() any]]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_491 any) *string {
			if hx_value_491 == nil {
				var hx_zero_492 *string
				return hx_zero_492
			}
			return hx_value_491.(*string)
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
		return hxrt.FunctionCall(source)(hx_recover_callable_0(argument_0), hx_recover_class_3(argument_1))
	case *hxrt.Function[func(*haxe___Template__ExprCursor) *hxrt.Function[func() any]]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_3(argument_0))
	case *hxrt.Function[func(*haxe___Template__ExprCursor)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(hx_recover_class_3(argument_0))
		return nil
	case *hxrt.Function[func(*haxe___Template__TemplateExpr)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(func(hx_value_493 any) *haxe___Template__TemplateExpr {
			if hx_value_493 == nil {
				var hx_zero_494 *haxe___Template__TemplateExpr
				return hx_zero_494
			}
			return hx_value_493.(*haxe___Template__TemplateExpr)
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
		return hxrt.FunctionCall(source)(func(hx_value_495 any) *string {
			if hx_value_495 == nil {
				var hx_zero_496 *string
				return hx_zero_496
			}
			return hx_value_495.(*string)
		}(argument_0))
	case *hxrt.Function[func(int) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_497 any) int {
			if hx_value_497 == nil {
				var hx_zero_498 int
				return hx_zero_498
			}
			return hx_value_497.(int)
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
		return hxrt.FunctionCall(source)(func(hx_value_499 any) *string {
			if hx_value_499 == nil {
				var hx_zero_500 *string
				return hx_zero_500
			}
			return hx_value_499.(*string)
		}(argument_0), func(hx_value_501 any) int {
			if hx_value_501 == nil {
				var hx_zero_502 int
				return hx_zero_502
			}
			return hx_value_501.(int)
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
		return hxrt.FunctionCall(source)(func(hx_value_503 any) *string {
			if hx_value_503 == nil {
				var hx_zero_504 *string
				return hx_zero_504
			}
			return hx_value_503.(*string)
		}(argument_0), func(hx_value_505 any) *string {
			if hx_value_505 == nil {
				var hx_zero_506 *string
				return hx_zero_506
			}
			return hx_value_505.(*string)
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
		return hxrt.FunctionCall(source)(func(hx_value_507 any) *string {
			if hx_value_507 == nil {
				var hx_zero_508 *string
				return hx_zero_508
			}
			return hx_value_507.(*string)
		}(argument_0), hx_recover_callable_1(argument_1))
	case *hxrt.Function[func(*string, *hxrt.RegexMatch)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_509 any) *string {
			if hx_value_509 == nil {
				var hx_zero_510 *string
				return hx_zero_510
			}
			return hx_value_509.(*string)
		}(argument_0), func(hx_value_511 any) *hxrt.RegexMatch {
			if hx_value_511 == nil {
				var hx_zero_512 *hxrt.RegexMatch
				return hx_zero_512
			}
			return hx_value_511.(*hxrt.RegexMatch)
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
		return hxrt.FunctionCall(source)(func(hx_value_513 any) *string {
			if hx_value_513 == nil {
				var hx_zero_514 *string
				return hx_zero_514
			}
			return hx_value_513.(*string)
		}(argument_0), func(hx_value_515 any) *string {
			if hx_value_515 == nil {
				var hx_zero_516 *string
				return hx_zero_516
			}
			return hx_value_515.(*string)
		}(argument_1), func(hx_value_517 any) *hxrt.RegexMatch {
			if hx_value_517 == nil {
				var hx_zero_518 *hxrt.RegexMatch
				return hx_zero_518
			}
			return hx_value_517.(*hxrt.RegexMatch)
		}(argument_2))
	case *hxrt.Function[func(int) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_519 any) int {
			if hx_value_519 == nil {
				var hx_zero_520 int
				return hx_zero_520
			}
			return hx_value_519.(int)
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
		hxrt.FunctionCall(source)(func(hx_value_521 any) int {
			if hx_value_521 == nil {
				var hx_zero_522 int
				return hx_zero_522
			}
			return hx_value_521.(int)
		}(argument_0), func(hx_value_523 any) int {
			if hx_value_523 == nil {
				var hx_zero_524 int
				return hx_zero_524
			}
			return hx_value_523.(int)
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
		hxrt.FunctionCall(source)(func(hx_value_525 any) int {
			if hx_value_525 == nil {
				var hx_zero_526 int
				return hx_zero_526
			}
			return hx_value_525.(int)
		}(argument_0), hx_recover_class_4(argument_1), func(hx_value_527 any) int {
			if hx_value_527 == nil {
				var hx_zero_528 int
				return hx_zero_528
			}
			return hx_value_527.(int)
		}(argument_2), func(hx_value_529 any) int {
			if hx_value_529 == nil {
				var hx_zero_530 int
				return hx_zero_530
			}
			return hx_value_529.(int)
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
		hxrt.FunctionCall(source)(func(hx_value_531 any) int {
			if hx_value_531 == nil {
				var hx_zero_532 int
				return hx_zero_532
			}
			return hx_value_531.(int)
		}(argument_0), func(hx_value_533 any) int {
			if hx_value_533 == nil {
				var hx_zero_534 int
				return hx_zero_534
			}
			return hx_value_533.(int)
		}(argument_1), func(hx_value_535 any) int {
			if hx_value_535 == nil {
				var hx_zero_536 int
				return hx_zero_536
			}
			return hx_value_535.(int)
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
		return hxrt.FunctionCall(source)(func(hx_value_537 any) int {
			if hx_value_537 == nil {
				var hx_zero_538 int
				return hx_zero_538
			}
			return hx_value_537.(int)
		}(argument_0), func(hx_value_539 any) int {
			if hx_value_539 == nil {
				var hx_zero_540 int
				return hx_zero_540
			}
			return hx_value_539.(int)
		}(argument_1))
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_4(argument_0))
	case *hxrt.Function[func(int) float64]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_541 any) int {
			if hx_value_541 == nil {
				var hx_zero_542 int
				return hx_zero_542
			}
			return hx_value_541.(int)
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
		hxrt.FunctionCall(source)(func(hx_value_543 any) int {
			if hx_value_543 == nil {
				var hx_zero_544 int
				return hx_zero_544
			}
			return hx_value_543.(int)
		}(argument_0), func(hx_value_545 any) float64 {
			if hx_value_545 == nil {
				var hx_zero_546 float64
				return hx_zero_546
			}
			return hx_value_545.(float64)
		}(argument_1))
		return nil
	case *hxrt.Function[func(int) *haxe___Int64_____Int64]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_547 any) int {
			if hx_value_547 == nil {
				var hx_zero_548 int
				return hx_zero_548
			}
			return hx_value_547.(int)
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
		hxrt.FunctionCall(source)(func(hx_value_549 any) int {
			if hx_value_549 == nil {
				var hx_zero_550 int
				return hx_zero_550
			}
			return hx_value_549.(int)
		}(argument_0), hx_recover_class_5(argument_1))
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
		return hxrt.FunctionCall(source)(func(hx_value_551 any) int {
			if hx_value_551 == nil {
				var hx_zero_552 int
				return hx_zero_552
			}
			return hx_value_551.(int)
		}(argument_0), func(hx_value_553 any) int {
			if hx_value_553 == nil {
				var hx_zero_554 int
				return hx_zero_554
			}
			return hx_value_553.(int)
		}(argument_1), func(hx_value_555 any) *haxe__io__Encoding {
			if hx_value_555 == nil {
				var hx_zero_556 *haxe__io__Encoding
				return hx_zero_556
			}
			return hx_value_555.(*haxe__io__Encoding)
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
		return hxrt.FunctionCall(source)(func(hx_value_557 any) int {
			if hx_value_557 == nil {
				var hx_zero_558 int
				return hx_zero_558
			}
			return hx_value_557.(int)
		}(argument_0), func(hx_value_559 any) int {
			if hx_value_559 == nil {
				var hx_zero_560 int
				return hx_zero_560
			}
			return hx_value_559.(int)
		}(argument_1))
	case *hxrt.Function[func() []int]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.FunctionCall(source)()
	default:
		return hxrt.ReflectCallMethod(value, arguments)
	}
}

func hx_recover_callable_0(value any) *hxrt.Function[func() any] {
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

func hx_recover_callable_1(value any) *hxrt.Function[func(*EReg) *string] {
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
			return func(hx_value_563 any) *string {
				if hx_value_563 == nil {
					var hx_zero_564 *string
					return hx_zero_564
				}
				return hx_value_563.(*string)
			}(hxrt.FunctionCall(source)(func(hx_value_561 any) *string {
				if hx_value_561 == nil {
					var hx_zero_562 *string
					return hx_zero_562
				}
				return hx_value_561.(*string)
			}(arg_0)))
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_565 any) *string {
				if hx_value_565 == nil {
					var hx_zero_566 *string
					return hx_zero_566
				}
				return hx_value_565.(*string)
			}(hxrt.FunctionCall(source)(arg_0))
		})
	case *hxrt.Function[func(*string) *hxrt.Array]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_569 any) *string {
				if hx_value_569 == nil {
					var hx_zero_570 *string
					return hx_zero_570
				}
				return hx_value_569.(*string)
			}(hxrt.FunctionCall(source)(func(hx_value_567 any) *string {
				if hx_value_567 == nil {
					var hx_zero_568 *string
					return hx_zero_568
				}
				return hx_value_567.(*string)
			}(arg_0)))
		})
	case *hxrt.Function[func(*haxe___Template__TokenCursor) *haxe___Template__TemplateExpr]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_571 any) *string {
				if hx_value_571 == nil {
					var hx_zero_572 *string
					return hx_zero_572
				}
				return hx_value_571.(*string)
			}(hxrt.FunctionCall(source)(hx_recover_class_2(arg_0)))
		})
	case *hxrt.Function[func(*string) *hxrt.Function[func() any]]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_575 any) *string {
				if hx_value_575 == nil {
					var hx_zero_576 *string
					return hx_zero_576
				}
				return hx_value_575.(*string)
			}(hxrt.FunctionCall(source)(func(hx_value_573 any) *string {
				if hx_value_573 == nil {
					var hx_zero_574 *string
					return hx_zero_574
				}
				return hx_value_573.(*string)
			}(arg_0)))
		})
	case *hxrt.Function[func(*haxe___Template__ExprCursor) *hxrt.Function[func() any]]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_577 any) *string {
				if hx_value_577 == nil {
					var hx_zero_578 *string
					return hx_zero_578
				}
				return hx_value_577.(*string)
			}(hxrt.FunctionCall(source)(hx_recover_class_3(arg_0)))
		})
	case *hxrt.Function[func(*haxe___Template__ExprCursor)]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			hxrt.FunctionCall(source)(hx_recover_class_3(arg_0))
			return func(hx_value_579 any) *string {
				if hx_value_579 == nil {
					var hx_zero_580 *string
					return hx_zero_580
				}
				return hx_value_579.(*string)
			}(nil)
		})
	case *hxrt.Function[func(*haxe___Template__TemplateExpr)]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			hxrt.FunctionCall(source)(func(hx_value_581 any) *haxe___Template__TemplateExpr {
				if hx_value_581 == nil {
					var hx_zero_582 *haxe___Template__TemplateExpr
					return hx_zero_582
				}
				return hx_value_581.(*haxe___Template__TemplateExpr)
			}(arg_0))
			return func(hx_value_583 any) *string {
				if hx_value_583 == nil {
					var hx_zero_584 *string
					return hx_zero_584
				}
				return hx_value_583.(*string)
			}(nil)
		})
	case *hxrt.Function[func(*string) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_587 any) *string {
				if hx_value_587 == nil {
					var hx_zero_588 *string
					return hx_zero_588
				}
				return hx_value_587.(*string)
			}(hxrt.FunctionCall(source)(func(hx_value_585 any) *string {
				if hx_value_585 == nil {
					var hx_zero_586 *string
					return hx_zero_586
				}
				return hx_value_585.(*string)
			}(arg_0)))
		})
	case *hxrt.Function[func(int) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return hxrt.FunctionCall(source)(func(hx_value_589 any) int {
				if hx_value_589 == nil {
					var hx_zero_590 int
					return hx_zero_590
				}
				return hx_value_589.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_593 any) *string {
				if hx_value_593 == nil {
					var hx_zero_594 *string
					return hx_zero_594
				}
				return hx_value_593.(*string)
			}(hxrt.FunctionCall(source)(func(hx_value_591 any) int {
				if hx_value_591 == nil {
					var hx_zero_592 int
					return hx_zero_592
				}
				return hx_value_591.(int)
			}(arg_0)))
		})
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_595 any) *string {
				if hx_value_595 == nil {
					var hx_zero_596 *string
					return hx_zero_596
				}
				return hx_value_595.(*string)
			}(hxrt.FunctionCall(source)(hx_recover_class_4(arg_0)))
		})
	case *hxrt.Function[func(int) float64]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_599 any) *string {
				if hx_value_599 == nil {
					var hx_zero_600 *string
					return hx_zero_600
				}
				return hx_value_599.(*string)
			}(hxrt.FunctionCall(source)(func(hx_value_597 any) int {
				if hx_value_597 == nil {
					var hx_zero_598 int
					return hx_zero_598
				}
				return hx_value_597.(int)
			}(arg_0)))
		})
	case *hxrt.Function[func(int) *haxe___Int64_____Int64]:
		return hxrt.ProjectFunction(source, func(arg_0 *EReg) *string {
			return func(hx_value_603 any) *string {
				if hx_value_603 == nil {
					var hx_zero_604 *string
					return hx_zero_604
				}
				return hx_value_603.(*string)
			}(hxrt.FunctionCall(source)(func(hx_value_601 any) int {
				if hx_value_601 == nil {
					var hx_zero_602 int
					return hx_zero_602
				}
				return hx_value_601.(int)
			}(arg_0)))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_class_0(value any) *haxe__ValueException {
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

func hx_recover_class_1(value any) *haxe__Exception {
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

func hx_recover_class_2(value any) *haxe___Template__TokenCursor {
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

func hx_recover_class_3(value any) *haxe___Template__ExprCursor {
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

func hx_recover_class_4(value any) *haxe__io__Bytes {
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

func hx_recover_class_5(value any) *haxe___Int64_____Int64 {
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

func hx_is_nominal_0(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_1(value any) bool {
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
