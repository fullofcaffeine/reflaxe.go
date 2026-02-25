package go;

/**
	Package-level helpers for Go `context`.
**/
@:go.import("context")
extern class ContextPkg {
	@:go.name("Background")
	public static function background():Context;
}
