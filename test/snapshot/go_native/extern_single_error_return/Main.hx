@:go.import("errors")
extern class ErrorsPkg {
	@:go.name("New")
	public static function newError(message:String):go.Error;
}

@:go.import("context")
extern interface GoContext {
	@:go.name("Err")
	public function error():go.Error;
}

@:go.import("context")
extern class ContextPkg {
	@:go.name("Background")
	public static function background():GoContext;
}

class Main {
	static function main():Void {
		final noError = ContextPkg.background().error();
		final error = ErrorsPkg.newError("broken");
		ErrorsPkg.newError("ignored");

		Sys.println('nil=${noError == null}');
		Sys.println('message=${error.toString()}');
	}
}
