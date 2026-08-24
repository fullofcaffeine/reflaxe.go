package reflaxe.go.compiler;

/**
	Closed names for metadata that changes haxe.go lowering or native binding.

	Why
	Metadata arrives from Haxe as plain text and may include a leading colon.
	Scattered string comparisons can silently disagree about either spelling.

	What
	The values cover the source-boundary and typed-extern metadata understood by
	the compiler. Compatibility aliases remain explicit members of the same type.

	How
	Use `matches(...)` when reading a `MetadataEntry`. The abstract converts
	outward to `String`, but arbitrary strings cannot become trusted metadata names.
**/
enum abstract GoMetadataName(String) to String {
	var GoNative = "goNative";
	var GoMetal = "goMetal";
	var RemovedHaxeMetal = "haxeMetal";
	var GoAllowRaw = "goAllowRaw";
	var GoImport = "go.import";
	var GoPackage = "go.package";
	var GoPackageAlias = "go.pkg";
	var GoName = "go.name";
	var GoStruct = "go.struct";
	var NativeName = "native";
	var GoReceiver = "go.receiver";
	var GoValueError = "go.valueError";
	var GoValueErrorAlias = "go.value_error";
	var GoTupleReturn = "go.tupleReturn";
	var GoTupleReturnAlias = "go.tuple_return";

	/** Returns whether Haxe's normalized or colon-prefixed metadata name matches. */
	public inline function matches(actual:String):Bool {
		return actual == this || actual == ":" + this;
	}
}
