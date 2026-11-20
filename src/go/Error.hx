package go;

class Error {
	public var message(default, null):String;

	public function new(message:String) {
		this.message = message;
	}

	/**
		What: Returns the message; generated Go exposes this method as `String()`.
		Why: The compiler passes a go.Error to a native Go `error` parameter through
		that `String()` method (see `hxrt.ExportLegacyError`). Haxe dead-code
		elimination cannot see that use in generated Go. Without `@:keep`,
		`-dce full` removes the method when no Haxe code calls it, and the Go build
		fails.
		How: `@:keep` retains the method even when no Haxe code calls it.
	**/
	@:keep
	public function toString():String {
		return message;
	}
}
