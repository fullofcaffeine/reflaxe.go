/**
	A checked cast `cast(value, T)` returns the value when it is null or a `T`,
	and otherwise throws Haxe's catchable class cast error. `--interp` throws
	the String "Class cast error": a `Dynamic` catch receives that String and a
	`haxe.Exception` catch receives a ValueException with it as the message.

	The failures cover a `Dynamic` value of an unrelated class, a base-class
	value that is not the subclass, an interface target, a primitive target and
	a cast used only as a statement. The successes check that a passing cast
	keeps the value's identity and lets null through.
**/
interface Named {
	function label():String;
}

class Base {
	public function new() {}

	public function name():String
		return "base";
}

class Mid extends Base {
	override public function name():String
		return "mid";
}

class Leaf extends Mid {
	override public function name():String
		return "leaf";
}

class Tagged implements Named {
	public function new() {}

	public function label():String
		return "tagged";
}

class Other {
	public function new() {}
}

class Main {
	static function main() {
		var other:Dynamic = new Other();
		try {
			var mid = cast(other, Mid);
			Sys.println("unexpected Other as Mid: " + mid.name());
		} catch (e:Dynamic) {
			Sys.println("Other as Mid, Dynamic catch: " + Std.string(e) + " isString=" + Std.isOfType(e, String));
		}
		try {
			var mid = cast(other, Mid);
			Sys.println("unexpected Other as Mid: " + mid.name());
		} catch (e:haxe.Exception) {
			Sys.println("Other as Mid, Exception catch: " + e.message);
		}
		try {
			Sys.println("unexpected Other as Mid: " + cast(other, Mid).name());
		} catch (e:String) {
			Sys.println("Other as Mid, String catch: " + e);
		}

		var base:Base = new Base();
		try {
			cast(base, Mid);
			Sys.println("unexpected Base as Mid statement");
		} catch (e:Dynamic) {
			Sys.println("Base as Mid statement: " + e);
		}

		try {
			Sys.println("unexpected Other as Named: " + cast(other, Named).label());
		} catch (e:Dynamic) {
			Sys.println("Other as Named: " + e);
		}

		var number:Dynamic = 3;
		try {
			var text = cast(number, String);
			Sys.println("unexpected Int as String: " + text);
		} catch (e:Dynamic) {
			Sys.println("Int as String: " + e);
		}
		Sys.println("Int as Int: " + (cast(number, Int) + 1));

		var leafAsBase:Base = new Leaf();
		var leafAsMid = cast(leafAsBase, Mid);
		Sys.println("Leaf as Mid: " + leafAsMid.name() + " same=" + (leafAsMid == leafAsBase));
		var leafAsDynamic:Dynamic = new Leaf();
		Sys.println("Dynamic Leaf as Mid: " + cast(leafAsDynamic, Mid).name());
		var tagged:Dynamic = new Tagged();
		Sys.println("Tagged as Named: " + cast(tagged, Named).label());

		var noBase:Base = null;
		Sys.println("null Base as Mid: " + (cast(noBase, Mid) == null));
		var noValue:Dynamic = null;
		Sys.println("null Dynamic as Mid: " + (cast(noValue, Mid) == null));
	}
}
