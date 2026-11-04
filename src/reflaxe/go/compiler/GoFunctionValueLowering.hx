package reflaxe.go.compiler;

#if macro
import reflaxe.go.ast.GoAST.GoExpr;

/**
	What: Builds the shared runtime operations for first-class Haxe callables.
	Why: Function identity belongs to creation and transport, not executable Go
	addresses. Internal Go functions used for sequencing must remain unwrapped.
	How: Keep the concrete function signature in the generic runtime carrier;
	creation, adaptation and native export use separate explicit operations.
**/
function fresh(callable:GoExpr):GoExpr {
	return GoExpr.GoCall(GoExpr.GoIdent("hxrt.NewFunction"), [callable]);
}

/** Guard a stored Haxe value after its call arguments have been evaluated. */
function callable(value:GoExpr):GoExpr {
	return GoExpr.GoCall(GoExpr.GoIdent("hxrt.FunctionCall"), [value]);
}

/** Export the exact payload only at an explicit native callback boundary. */
function native(value:GoExpr):GoExpr {
	return GoExpr.GoCall(GoExpr.GoIdent("hxrt.FunctionNative"), [value]);
}

/** Retain source identity and null state while changing invocation shape. */
function adapt(source:GoExpr, callable:GoExpr):GoExpr {
	return GoExpr.GoCall(GoExpr.GoIdent("hxrt.AdaptFunction"), [source, callable]);
}

/** Preserve normalized source behavior when exposing another typed signature. */
function project(source:GoExpr, callable:GoExpr):GoExpr {
	return GoExpr.GoCall(GoExpr.GoIdent("hxrt.ProjectFunction"), [source, callable]);
}
#end
