typedef Window = {
	var maybeEnd:Null<Int>;
}

class Bounds {
	public static var shared:Null<Int> = null;

	public var maybeEnd:Null<Int>;

	public function new(maybeEnd:Null<Int>) {
		this.maybeEnd = maybeEnd;
	}
}

/**
	A null `Null<Int>` passed as the optional argument of `String.substr`,
	`substring`, `indexOf` or `lastIndexOf` must behave as in `--interp`
	(haxe_go-dumy).

	A null optional argument reads as omitted: `s.substr(1, n)` with a null
	`n` is `s.substr(1)`. The null comes from a constant, `Null<Int>` and
	`?x:Int` parameters, locals, instance, static and structure fields,
	`Array<Null<Int>>` elements, call results, a ternary and `Dynamic`
	values. Non-null, negative, out-of-range and zero values keep their
	ordinary results.

	`four` prints `substr(1, v)|substring(0, v)|indexOf("b", v)|
	lastIndexOf("b", v)|indexOf("", v)` for one value `v`. (`substring` with an
	end below its start and `lastIndexOf("", v)` with a negative `v` are
	left out: haxe.go does not match `--interp` there for any Int.)
	The `order` line checks that the receiver and the arguments run once, in
	source order.
**/
class Main {
	static var trail:Array<String> = [];

	static function four(s:String, value:Null<Int>):String {
		return [
			s.substr(1, value),
			s.substring(0, value),
			Std.string(s.indexOf("b", value)),
			Std.string(s.lastIndexOf("b", value)),
			Std.string(s.indexOf("", value))
		].join("|");
	}

	static function optional(s:String, ?value:Int):String {
		return [
			s.substr(1, value),
			s.substring(1, value),
			Std.string(s.indexOf("b", value)),
			Std.string(s.lastIndexOf("b", value)),
			Std.string(s.lastIndexOf("", value))
		].join("|");
	}

	// Inside the `if`, `value` is known to be non-null.
	static function narrowed(s:String, ?value:Int):String {
		if (value != null) {
			return s.substr(1, value) + "|" + s.substring(1, value) + "|" + s.indexOf("b", value) + "|" + s.lastIndexOf("b", value);
		}
		return "missing";
	}

	static function noInt():Null<Int> {
		return null;
	}

	static function someInt(value:Int):Null<Int> {
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

	static function main() {
		var s = "abcabc";

		// Omitted, a null constant, and Int constants.
		Sys.println("omitted=" + [
			s.substr(1),
			s.substring(1),
			s.substring(4),
			Std.string(s.indexOf("b")),
			Std.string(s.lastIndexOf("b"))
		].join("|"));
		Sys.println("literal.null=" + [
			s.substr(1, null),
			s.substring(1, null),
			Std.string(s.indexOf("b", null)),
			Std.string(s.lastIndexOf("b", null))
		].join("|"));
		Sys.println("constant=" + [
			s.substr(1, 2),
			s.substring(1, 2),
			Std.string(s.indexOf("b", 2)),
			Std.string(s.lastIndexOf("b", 2))
		].join("|"));
		Sys.println("constant.negative=" + [
			s.substr(1, -2),
			s.substring(0, -2),
			Std.string(s.indexOf("b", -2)),
			Std.string(s.lastIndexOf("b", -2))
		].join("|"));

		// A `Null<Int>` parameter: null, non-null, negative, out of range, zero.
		for (value in [null, 2, -2, 99, -99, 0]) {
			Sys.println("param." + value + "=" + four(s, value));
		}

		// A `?value:Int` parameter: omitted, null and a value.
		Sys.println("optional.omitted=" + optional(s));
		Sys.println("optional.null=" + optional(s, null));
		Sys.println("optional.value=" + optional(s, 3));
		Sys.println("narrowed=" + narrowed(s, 3) + "," + narrowed(s));

		// Locals.
		var none = noInt();
		var some = someInt(2);
		Sys.println("local=" + [
			s.substr(1, none),
			s.substring(1, none),
			Std.string(s.indexOf("b", none)),
			Std.string(s.lastIndexOf("b", none)),
			s.substr(1, some),
			Std.string(s.lastIndexOf("b", some))
		].join("|"));

		// Instance, static and structure fields.
		var bounds = new Bounds(null);
		var window:Window = {maybeEnd: null};
		Sys.println("null.field=" + [
			s.substr(1, bounds.maybeEnd),
			s.substring(1, Bounds.shared),
			Std.string(s.indexOf("b", window.maybeEnd)),
			Std.string(s.lastIndexOf("b", bounds.maybeEnd))
		].join("|"));
		bounds.maybeEnd = 3;
		Bounds.shared = 2;
		window.maybeEnd = 2;
		Sys.println("value.field=" + [
			s.substr(1, bounds.maybeEnd),
			s.substring(1, Bounds.shared),
			Std.string(s.indexOf("b", window.maybeEnd)),
			Std.string(s.lastIndexOf("b", bounds.maybeEnd))
		].join("|"));

		// Array<Null<Int>> elements.
		var ends:Array<Null<Int>> = [null, 3];
		var i = 0;
		Sys.println("array=" + [
			s.substr(1, ends[0]),
			s.substring(1, ends[i]),
			Std.string(s.indexOf("b", ends[i])),
			Std.string(s.lastIndexOf("b", ends[0])),
			s.substr(1, ends[i + 1]),
			Std.string(s.indexOf("b", ends[1]))
		].join("|"));

		// Call results, a ternary and Dynamic values.
		var flag = s.length > 3;
		var dynamicNull:Dynamic = noInt();
		var dynamicValue:Dynamic = someInt(2);
		Sys.println("call=" + [
			s.substr(1, noInt()),
			s.substring(1, noInt()),
			Std.string(s.indexOf("b", noInt())),
			Std.string(s.lastIndexOf("b", noInt())),
			s.substring(1, someInt(3)),
			s.substr(1, flag ? null : 1),
			s.substr(1, dynamicNull),
			Std.string(s.lastIndexOf("b", dynamicValue))
		].join("|"));

		// Multi-byte text keeps code-point positions.
		var accented = "héllo";
		Sys.println("unicode=" + [
			accented.substr(1, none),
			accented.substring(1, none),
			Std.string(accented.indexOf("l", none)),
			Std.string(accented.lastIndexOf("l", none))
		].join("|"));

		// The receiver and the arguments run once, in source order.
		var order = [
			tickText("a", s).substr(tickInt("b", 1), tick("c", null)),
			tickText("d", s).substring(tickInt("e", 1), tick("f", null)),
			Std.string(tickText("g", s).indexOf(tickText("h", "b"), tick("i", null))),
			Std.string(tickText("j", s).lastIndexOf(tickText("k", "b"), tick("l", null))),
			tickText("m", s).substr(tickInt("n", 1), tick("o", 2)),
			tickText("p", s).substring(tickInt("q", 1))
		];
		Sys.println("order=" + order.join("|") + ":" + trail.join(","));
	}
}
