package reflaxe.go.ast;

import reflaxe.go.ast.GoAST.GoExpr;

/**
	What: One positional, expression-keyed, or struct-field element in a Go
	composite literal.
	Why: Struct, map, slice, and array initialization need different key meanings,
	which were previously flattened into raw target text.
	How: Values and expression keys stay as `GoExpr`; struct field names use their
	own variant so transforms can traverse values without guessing syntax.
**/
enum GoCompositeElement {
	GoCompositeValue(value:GoExpr);
	GoCompositeKeyValue(key:GoExpr, value:GoExpr);
	GoCompositeField(fieldName:String, value:GoExpr);
}
