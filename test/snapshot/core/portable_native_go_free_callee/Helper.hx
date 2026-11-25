import go.NativeAny;

/**
	A go-free module-level function in a `@:goNative` module.

	Its sibling `describe` has a `go.*` argument type. Callers of `plain` see only
	`():String`, so they must not be reported.
**/
@:goNative function plain():String {
	return describe(NativeAny.ofString("module"));
}

@:goNative private function describe(value:NativeAny):String {
	return value == null ? "nil" : "module value";
}
