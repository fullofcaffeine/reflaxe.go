@:runtimeValue
@:coreType
abstract MyCore {}

class Main {
	static function main() {
		var d:Dynamic = 1;
		Sys.println("d.core=" + Std.isOfType(d, MyCore));

		var core:MyCore = cast 1;
		Sys.println("typed.core=" + Std.isOfType(core, MyCore));

		var dynCore:Dynamic = core;
		Sys.println("dyn.core=" + Std.isOfType(dynCore, MyCore));

		var dynNull:Dynamic = null;
		Sys.println("null.core=" + Std.isOfType(dynNull, MyCore));
	}
}
