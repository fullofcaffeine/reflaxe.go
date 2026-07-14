package reflaxe.go.compiler;

/**
	Why
	Semantic authority must come from explicit source structure, not from a
	build-wide optimization preset.

	What
	Identifies the canonical source of portable versus Go-native semantics.

	How
	Typed `go.*` APIs and `@:goNative` modules opt into native contracts; all
	other source keeps portable Haxe semantics regardless of policy preset.
**/
enum abstract GoSemanticBoundarySource(String) from String to String {
	var TypedApiOrModule = "typed_api_or_module";

	public inline function label():String {
		return this;
	}
}
