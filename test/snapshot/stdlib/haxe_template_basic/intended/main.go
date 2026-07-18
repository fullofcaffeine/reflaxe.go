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
	var v any = any(basic.__hx_this.execute(func() map[string]any {
		hx_obj_1 := map[string]any{}
		hx_obj_1["name"] = hxrt.StringFromLiteral("ok")
		return hx_obj_1
	}(), nil))
	hxrt.Println(v)
	cond := New_haxe__Template(hxrt.StringFromLiteral("::if enabled::yes::else::no::end::"))
	var v_1 any = any(cond.__hx_this.execute(func() map[string]any {
		hx_obj_2 := map[string]any{}
		hx_obj_2["enabled"] = true
		return hx_obj_2
	}(), nil))
	hxrt.Println(v_1)
	var v_2 any = any(cond.__hx_this.execute(func() map[string]any {
		hx_obj_3 := map[string]any{}
		hx_obj_3["enabled"] = false
		return hx_obj_3
	}(), nil))
	hxrt.Println(v_2)
	loop := New_haxe__Template(hxrt.StringFromLiteral("::foreach items::::__current__::::end::"))
	var v_3 any = any(loop.__hx_this.execute(func() map[string]any {
		hx_obj_4 := map[string]any{}
		hx_obj_4["items"] = hxrt.NewArray(hxrt.StringFromLiteral("a"), hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("c"))
		return hx_obj_4
	}(), nil))
	hxrt.Println(v_3)
}

func hxrt__generated_method_field(obj any, key string) any {
	var receiver any
	switch value := obj.(type) {
	case *EReg:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Template:
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
	case *haxe__Template:
		return hxrt__generated_method_field__haxe__Template(value, key)
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
		return value.expandReplacement
	case "map":
		return value.map_
	case "match":
		return value.match
	case "matchSub":
		return value.matchSub
	case "matched":
		return value.matched
	case "matchedLeft":
		return value.matchedLeft
	case "matchedPos":
		return value.matchedPos
	case "matchedRight":
		return value.matchedRight
	case "remember":
		return value.remember
	case "replace":
		return value.replace
	case "requireMatch":
		return value.requireMatch
	case "split":
		return value.split
	}
	return nil
}

func hxrt__generated_method_field__haxe__Template(value *haxe__Template, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "execute":
		return value.execute
	case "makeConst":
		return value.makeConst
	case "makeExpr":
		return value.makeExpr
	case "makeExpr2":
		return value.makeExpr2
	case "makePath":
		return value.makePath
	case "parse":
		return value.parse
	case "parseBlock":
		return value.parseBlock
	case "parseExpr":
		return value.parseExpr
	case "parseTokens":
		return value.parseTokens
	case "popStackValue":
		return value.popStackValue
	case "resolve":
		return value.resolve
	case "run":
		return value.run
	case "skipSpaces":
		return value.skipSpaces
	}
	return nil
}

func hxrt__generated_method_field__haxe__io__Bytes(value *haxe__io__Bytes, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "__hx_nativeView":
		return value.__hx_nativeView
	case "blit":
		return value.blit
	case "compare":
		return value.compare
	case "fill":
		return value.fill
	case "get":
		return value.get
	case "getData":
		return value.getData
	case "getDouble":
		return value.getDouble
	case "getFloat":
		return value.getFloat
	case "getInt32":
		return value.getInt32
	case "getInt64":
		return value.getInt64
	case "getString":
		return value.getString
	case "getUInt16":
		return value.getUInt16
	case "readString":
		return value.readString
	case "set":
		return value.set
	case "setDouble":
		return value.setDouble
	case "setFloat":
		return value.setFloat
	case "setInt32":
		return value.setInt32
	case "setInt64":
		return value.setInt64
	case "setUInt16":
		return value.setUInt16
	case "sub":
		return value.sub
	case "toHex":
		return value.toHex
	case "toString":
		return value.toString
	}
	return nil
}

func hxrt__generated_method_field__haxe__iterators__StringIterator(value *haxe__iterators__StringIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasNext":
		return value.hasNext
	case "next":
		return value.next
	}
	return nil
}

func hxrt__generated_method_field__haxe__iterators__StringKeyValueIterator(value *haxe__iterators__StringKeyValueIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasNext":
		return value.hasNext
	case "next":
		return value.next
	}
	return nil
}

type Std struct {
}

type haxe__ds__Option struct {
	tag    int
	params []any
}

var haxe__ds__Option_None *haxe__ds__Option = &haxe__ds__Option{tag: 1, params: []any{}}

func haxe__ds__Option_Some(value any) *haxe__ds__Option {
	return &haxe__ds__Option{tag: 0, params: []any{value}}
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
	switch value := object.(type) {
	case *EReg:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Template:
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
	case *haxe__Template:
		return hxrt__generated_field_lookup__haxe__Template(value, key)
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
	switch value := object.(type) {
	case *EReg:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__Template:
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
	case *haxe__Template:
		return hxrt__generated_field_has__haxe__Template(value, key)
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
	switch value := object.(type) {
	case *EReg:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__Template:
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
	case *haxe__Template:
		return hxrt__generated_field_set__haxe__Template(value, key, incoming)
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
	switch value := object.(type) {
	case *EReg:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Template:
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
		return hxrt.NewArray(hxrt.StringFromLiteral("global"), hxrt.StringFromLiteral("handle"), hxrt.StringFromLiteral("lastMatch"), hxrt.StringFromLiteral("lastSource"))
	case *haxe__Template:
		return hxrt.NewArray(hxrt.StringFromLiteral("context"), hxrt.StringFromLiteral("expr"), hxrt.StringFromLiteral("macros"), hxrt.StringFromLiteral("output"), hxrt.StringFromLiteral("stack"))
	case *haxe___Int64_____Int64:
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	case *haxe___Template__ExprCursor:
		return hxrt.NewArray(hxrt.StringFromLiteral("index"), hxrt.StringFromLiteral("tokens"))
	case *haxe___Template__TokenCursor:
		return hxrt.NewArray(hxrt.StringFromLiteral("index"), hxrt.StringFromLiteral("tokens"))
	case *haxe__io__Bytes:
		return hxrt.NewArray(hxrt.StringFromLiteral("__hx_dataExposed"), hxrt.StringFromLiteral("__hx_raw"), hxrt.StringFromLiteral("__hx_rawValid"), hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("length"))
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
		return value.global
	case "handle":
		return value.handle
	case "lastMatch":
		return value.lastMatch
	case "lastSource":
		return value.lastSource
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

func hxrt__generated_field_lookup__haxe__Template(value *haxe__Template, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "context":
		return value.context
	case "expr":
		return value.expr
	case "macros":
		return value.macros
	case "output":
		return value.output
	case "stack":
		return value.stack
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

func hxrt__generated_field_lookup__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "high":
		return value.high
	case "low":
		return value.low
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
		return value.index
	case "tokens":
		return value.tokens
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
		return value.index
	case "tokens":
		return value.tokens
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
		return value.__hx_dataExposed
	case "__hx_raw":
		return value.__hx_raw
	case "__hx_rawValid":
		return value.__hx_rawValid
	case "b":
		return value.b
	case "length":
		return value.length
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
		return value.offset
	case "s":
		return value.s
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
		return value.offset
	case "s":
		return value.s
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
