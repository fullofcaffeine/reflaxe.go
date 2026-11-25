class Box {
	public function new() {}

	public function add(a:Int, b:Int):Int
		return a + b;
}

/**
	Main-class statics print as bare Go package-level names. A name that equals
	a Go predeclared identifier (`len`, `append`, `string`, `any`, `nil`, ...)
	would shadow that builtin for the whole Go package, so generated support
	code that calls the builtin stops compiling. Those statics must get the
	qualified `Main_*` form instead. Locals and parameters with such names
	shadow the builtin only in their own scope, but generated code in that
	scope still uses it (`any`, `string`, `int`), so they take a `_local`
	suffix.
**/
class Main {
	static var string:String = "s";
	static var any:Int = 1;
	static var nil:Int = 2;
	static var error:String = "e";
	static var iota:Int = 3;
	static var bool:Bool = true;
	static var byte:Int = 4;
	static var rune:Int = 5;
	static var float64:Float = 6.5;
	static var uintptr:Int = 7;
	static var comparable:Int = 8;

	static function len(value:String):Int
		return value.length;

	static function append(a:String, b:String):String
		return a + b;

	static function make(value:Int):Int
		return value * 2;

	static function copy(value:String):String
		return value;

	static function cap(value:Int):Int
		return value + 1;

	static function min(a:Int, b:Int):Int
		return a < b ? a : b;

	static function max(a:Int, b:Int):Int
		return a > b ? a : b;

	static function clear():String
		return "cleared";

	static function delete(value:String):String
		return "-" + value;

	static function panic(value:String):String
		return "!" + value;

	static function recover():String
		return "ok";

	static function print(value:String):String
		return "<" + value + ">";

	static function println(value:String):String
		return value + ";";

	static function close():Int
		return 9;

	static function complex():Int
		return 10;

	static function real():Int
		return 11;

	static function imag():Int
		return 12;

	static function plain(value:String):String
		return value + "?";

	static function locals(string:String):String {
		var len = 3;
		var any = 1;
		var int = 4;
		var nil = 5;
		var error = "e";
		var copy = [1, 2];
		var append = [3];
		var values = [1, 2, 3];
		var boxed:Dynamic = {x: 1};
		values.push(len);
		var joined = values.concat(append).concat(copy);
		return (values.length + len + any + int + nil + joined.length) + string + error + Std.string(boxed.x);
	}

	/** Iterating a rest argument and reading its length use Go's `len`. */
	static function count(len:Int, values:haxe.Rest<Int>):Int {
		var total = len;
		for (value in values)
			total += value;
		return total + values.length;
	}

	static function main():Void {
		var box:Dynamic = new Box();
		Sys.println(Reflect.callMethod(box, Reflect.field(box, "add"), [1, 2]) + len("ab"));
		final measure:String->Int = len;
		Sys.println(measure("abc"));
		Sys.println(append(string, error)
			+ copy("c")
			+ delete("d")
			+ panic("p")
			+ recover()
			+ print("q")
			+ println("r")
			+ clear()
			+ plain("z"));
		Sys.println([
			any,
			nil,
			iota,
			byte,
			rune,
			uintptr,
			comparable,
			make(3),
			cap(4),
			min(5, 6),
			max(7, 8),
			close(),
			complex(),
			real(),
			imag()
		].join(","));
		Sys.println(bool + " " + float64);
		Sys.println(locals("x"));
		Sys.println(count(1, 2, 3));
	}
}
