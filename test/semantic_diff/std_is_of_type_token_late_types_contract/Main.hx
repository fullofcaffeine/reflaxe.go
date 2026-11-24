import haxe.ds.Option;
import haxe.exceptions.NotImplementedException;
import haxe.exceptions.PosException;

/**
	What: Tests a value against a class or enum token that arrives as a runtime
	value, not as a type literal.

	Why: The compiler lowers classes in name order, so this class is lowered
	before `Main`. The standard-library types it is asked about are queued only
	while `Main` is lowered: `NotImplementedException` by its constructor call,
	its superclass `PosException` by that subclass, and `Option` by `Some`. A type
	test built while this class was lowered could not name those types.

	How: `Std.isOfType` receives `Class<Dynamic>` and `Enum<Dynamic>` parameters,
	so the generated Go must compare the runtime token against every class and
	enum of the finished program.
**/
class ClassTokenCheck {
	public static function isClass(value:Dynamic, token:Class<Dynamic>):Bool {
		return Std.isOfType(value, token);
	}

	public static function isEnum(value:Dynamic, token:Enum<Dynamic>):Bool {
		return Std.isOfType(value, token);
	}
}

class Main {
	static function fail():Void {
		throw new NotImplementedException();
	}

	static function main() {
		var error:Dynamic = new NotImplementedException();
		Sys.println("class token NotImplementedException=" + ClassTokenCheck.isClass(error, NotImplementedException));
		Sys.println("class token PosException=" + ClassTokenCheck.isClass(error, PosException));
		Sys.println("class token Exception=" + ClassTokenCheck.isClass(error, haxe.Exception));
		Sys.println("class token String=" + ClassTokenCheck.isClass(error, String));
		Sys.println("class token on string=" + ClassTokenCheck.isClass("text", PosException));

		var option:Dynamic = Some(1);
		Sys.println("enum token Option=" + ClassTokenCheck.isEnum(option, Option));
		Sys.println("enum token on error=" + ClassTokenCheck.isEnum(error, Option));

		try {
			fail();
		} catch (caught:PosException) {
			Sys.println("caught as PosException=" + Std.isOfType(caught, NotImplementedException));
		}
	}
}
