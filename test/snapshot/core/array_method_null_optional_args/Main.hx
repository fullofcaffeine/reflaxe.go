typedef Window = {
	var maybeFrom:Null<Int>;
}

class Bounds {
	public static var shared:Null<Int> = null;

	public var maybeFrom:Null<Int>;

	public function new(maybeFrom:Null<Int>) {
		this.maybeFrom = maybeFrom;
	}
}

/**
	A null `Null<Int>` passed as the optional argument of `Array.indexOf`,
	`lastIndexOf` or `slice` (haxe_go-b9sn).

	Haxe reads a null optional argument as omitted: `a.indexOf(x, n)` with a
	null `n` is `a.indexOf(x)`, and `a.slice(1, n)` is `a.slice(1)`, in
	`--interp`, neko and python. haxe.go raised an uncaught Go panic for a
	null `?from:Int` or `Null<Int>` parameter, also when a caller left the
	optional argument out, and did not compile a `Dynamic` call result as
	the start of a search.

	`three` prints `indexOf("b", v)|lastIndexOf("b", v)|slice(1, v)|
	ints.indexOf(2, v)` for one value `v`. (A `lastIndexOf` start that stays
	below 0 after adding the length is left out: haxe.go finds nothing there,
	as the Haxe API documents, while `--interp` searches from 0.)
	The `order` line checks that the receiver and the arguments run once, in
	source order.
**/
class Main {
	static var trail:Array<String> = [];

	static function three(a:Array<String>, ints:Array<Int>, value:Null<Int>):String {
		return [
			Std.string(a.indexOf("b", value)),
			Std.string(a.lastIndexOf("b", value)),
			a.slice(1, value).join(","),
			Std.string(ints.indexOf(2, value))
		].join("|");
	}

	static function optional(a:Array<String>, ?from:Int):String {
		return [
			Std.string(a.indexOf("b", from)),
			Std.string(a.lastIndexOf("b", from)),
			a.slice(1, from).join(",")
		].join("|");
	}

	// Inside the `if`, `from` is known to be non-null.
	static function narrowed(a:Array<String>, ?from:Int):String {
		if (from != null) {
			return a.indexOf("b", from) + "|" + a.lastIndexOf("b", from) + "|" + a.slice(1, from).join(",");
		}
		return "missing";
	}

	static function noInt():Null<Int> {
		return null;
	}

	static function someInt(value:Int):Null<Int> {
		return value;
	}

	static function dynamicValue(value:Null<Int>):Dynamic {
		return value;
	}

	static function tick(label:String, value:Null<Int>):Null<Int> {
		trail.push(label);
		return value;
	}

	static function tickInt(label:String, value:Int):Int {
		trail.push(label);
		return value;
	}

	static function tickText(label:String, value:String):String {
		trail.push(label);
		return value;
	}

	static function tickArray(label:String, value:Array<String>):Array<String> {
		trail.push(label);
		return value;
	}

	static function main() {
		var a = ["a", "b", "c", "a", "b", "c"];
		var ints = [1, 2, 3, 1, 2, 3];

		// Omitted, a null constant, and Int constants.
		Sys.println("omitted=" + [Std.string(a.indexOf("b")), Std.string(a.lastIndexOf("b")), a.slice(1).join(",")].join("|"));
		Sys.println("literal.null=" + [
			Std.string(a.indexOf("b", null)),
			Std.string(a.lastIndexOf("b", null)),
			a.slice(1, null).join(",")
		].join("|"));
		Sys.println("constant=" + [
			Std.string(a.indexOf("b", 2)),
			Std.string(a.lastIndexOf("b", 2)),
			a.slice(1, 3).join(","),
			Std.string(a.indexOf("b", -2)),
			Std.string(a.lastIndexOf("b", -2)),
			a.slice(1, -2).join(",")
		].join("|"));

		// A `Null<Int>` parameter: null, zero, non-null, negative, out of range.
		for (value in [null, 0, 2, -2, 5, 99]) {
			Sys.println("param." + value + "=" + three(a, ints, value));
		}

		// A `?from:Int` parameter: omitted, null and a value.
		Sys.println("optional.omitted=" + optional(a));
		Sys.println("optional.null=" + optional(a, null));
		Sys.println("optional.value=" + optional(a, 3));
		Sys.println("narrowed=" + narrowed(a, 3) + "," + narrowed(a));

		// Locals.
		var none = noInt();
		var some = someInt(2);
		Sys.println("local=" + [
			Std.string(a.indexOf("b", none)),
			Std.string(a.lastIndexOf("b", none)),
			a.slice(1, none).join(","),
			Std.string(a.indexOf("b", some)),
			Std.string(a.lastIndexOf("b", some)),
			a.slice(1, some).join(",")
		].join("|"));

		// Instance, static and structure fields.
		var bounds = new Bounds(null);
		var window:Window = {maybeFrom: null};
		Sys.println("null.field=" + [
			Std.string(a.indexOf("b", bounds.maybeFrom)),
			Std.string(a.lastIndexOf("b", Bounds.shared)),
			a.slice(1, window.maybeFrom).join(",")
		].join("|"));
		bounds.maybeFrom = 2;
		Bounds.shared = 3;
		window.maybeFrom = 3;
		Sys.println("value.field=" + [
			Std.string(a.indexOf("b", bounds.maybeFrom)),
			Std.string(a.lastIndexOf("b", Bounds.shared)),
			a.slice(1, window.maybeFrom).join(",")
		].join("|"));

		// Array<Null<Int>> elements, and an Array<Int> element read past the end.
		var starts:Array<Null<Int>> = [null, 2];
		var i = 0;
		Sys.println("array=" + [
			Std.string(a.indexOf("b", starts[0])),
			Std.string(a.lastIndexOf("b", starts[i])),
			a.slice(1, starts[i]).join(","),
			Std.string(a.indexOf("b", starts[i + 1])),
			a.slice(1, starts[1]).join(","),
			Std.string(a.lastIndexOf("b", ints[1])),
			Std.string(a.indexOf("b", ints[10]))
		].join("|"));

		// Call results, a ternary and Dynamic values.
		var flag = a.length > 3;
		var dynamicNull:Dynamic = noInt();
		var dynamicSome:Dynamic = someInt(3);
		Sys.println("call=" + [
			Std.string(a.indexOf("b", noInt())),
			Std.string(a.lastIndexOf("b", noInt())),
			a.slice(1, noInt()).join(","),
			Std.string(a.indexOf("b", someInt(2))),
			Std.string(a.lastIndexOf("b", flag ? null : 1)),
			a.slice(1, flag ? null : 1).join(",")
		].join("|"));
		Sys.println("dynamic=" + [
			Std.string(a.indexOf("b", dynamicNull)),
			Std.string(a.lastIndexOf("b", dynamicSome)),
			a.slice(1, dynamicNull).join(","),
			Std.string(a.indexOf("b", dynamicValue(null))),
			Std.string(a.lastIndexOf("b", dynamicValue(2))),
			a.slice(1, dynamicValue(3)).join(",")
		].join("|"));

		// The receiver and the arguments run once, in source order.
		var order = [
			Std.string(tickArray("a", a).indexOf(tickText("b", "b"), tick("c", null))),
			Std.string(tickArray("d", a).lastIndexOf(tickText("e", "b"), tick("f", null))),
			tickArray("g", a).slice(tickInt("h", 1), tick("i", null)).join(","),
			Std.string(tickArray("j", a).indexOf(tickText("k", "b"), tick("l", 2))),
			tickArray("m", a).slice(tickInt("n", 1), tick("o", 3)).join(",")
		];
		Sys.println("order=" + order.join("|") + ":" + trail.join(","));
	}
}
