/**
	A class in the same Go package as `Main`. Its parameter and local share
	the names of `Main` statics, and it reads both through `Main.*`.
**/
class Other {
	public static function run(count:Int):Int {
		var helper = count * 2;
		return helper + Main.helper(count, Main.count);
	}
}

/**
	A constructor parameter, a method parameter and a dynamic method parameter
	named like `Main` statics. The dynamic method's body is lowered inside the
	generated constructor, not where the class declares it.
**/
class Counter {
	var base:Int;

	public function new(count:Int) {
		base = count + Main.count;
	}

	public function add(helper:Int):Int
		return base + helper + Main.helper(helper, 1);

	public dynamic function scale(count:Int):Int
		return count * 10 + Main.count;
}

/**
	The constructor never names `Main.count`, but the dynamic method does. Its
	body becomes a Go closure inside the generated constructor, where the
	constructor's `count` parameter would also be in scope.
**/
class Gauge {
	var start:Int;

	public function new(count:Int) {
		start = count;
	}

	public dynamic function read(step:Int):Int
		return start + step + Main.count;
}

/**
	A field initializer names `Main.count`. Haxe runs it at the start of the
	constructor, beside the constructor's `count` parameter.
**/
class Holder {
	var seen:Int = Main.count * 10;

	public final total:Int;

	public function new(count:Int) {
		total = count + seen;
	}
}

/**
	Main-class statics print as bare Go package-level names (`count`,
	`helper`). A Haxe local, parameter, loop variable, catch variable or local
	function with the same name would shadow that static in Go, while Haxe code
	that writes `Main.count` or `Main.helper` still means the static. A static
	function then fails `go build`, and a static var silently reads or writes
	the local instead. Each shape below is the first declaration of its name in
	its function, so no counter suffix hides the clash; the locals must take a
	`_local` suffix instead. No local may take a main-class static's Go name,
	so even a function that never names the static renames its locals
	(`unrelated`).
**/
class Main {
	public static var count:Int = 1;

	public static function helper(a:Int, b:Int):Int
		return a + b;

	/** A static initializer whose function literal has a parameter named like a static. */
	static final offset:Int->Int = (helper:Int) -> helper + Main.helper(helper, 1);

	/** Locals next to reads and writes of the statics, and a function value. */
	static function locals():String {
		var count = 3;
		var helper = 5;
		var before = count + ":" + Main.count;
		Main.count = Main.count + count;
		Main.count++;
		count++;
		final add:(Int, Int) -> Int = Main.helper;
		return before + " " + count + ":" + Main.count + " " + (helper + Main.helper(1, 2)) + " " + add(helper, count);
	}

	static function params(count:Int, helper:Int):String {
		Main.count += count;
		return count + "," + helper + "," + Main.count + "," + Main.helper(count, helper);
	}

	/** A bound the typer cannot see keeps the loop from being unrolled. */
	static function countLoop(limit:Int):Int {
		var total = 0;
		for (count in 0...limit)
			total += count * 10 + Main.count;
		return total;
	}

	static function helperLoop(values:Array<Int>):Int {
		var total = 0;
		for (helper in values)
			total += helper * 100 + Main.helper(helper, helper);
		return total;
	}

	static function countCatch():Int {
		try {
			throw 7;
		} catch (count:Int) {
			Main.count += count;
			return count * 1000 + Main.count;
		}
	}

	static function helperCatch():String {
		try {
			throw "boom";
		} catch (helper:String) {
			return helper + Main.helper(2, 3);
		}
	}

	static function localFunctions():String {
		function count():Int
			return Main.count * 2;
		function helper(a:Int):Int
			return a * 10 + Main.helper(a, 1);
		return count() + ":" + Main.count + " " + helper(2) + ":" + Main.helper(2, 2);
	}

	static function closures():String {
		var count = 40;
		var helper = 50;
		final read = () -> count + ":" + Main.count + ":" + helper + ":" + Main.helper(count, helper);
		final write = (step:Int) -> {
			Main.count += step;
			count += step;
			helper = Main.helper(helper, step);
		};
		write(3);
		return read() + " " + count + ":" + Main.count + ":" + helper;
	}

	/** Nothing here refers to `Main.count` or `Main.helper`; the locals still take the `_local` suffix. */
	static function unrelated():Int {
		var count = 9;
		var helper = 2;
		return count + helper;
	}

	static function main():Void {
		Sys.println(locals());
		Sys.println(params(10, 20));
		Sys.println(countLoop(3));
		Sys.println(helperLoop([1, 2]));
		Sys.println(countCatch());
		Sys.println(helperCatch());
		Sys.println(localFunctions());
		Sys.println(closures());
		Sys.println(Other.run(4));
		final counter = new Counter(6);
		Sys.println(counter.add(7) + " " + counter.scale(2));
		Sys.println(offset(4) + " " + unrelated());
		Sys.println(new Gauge(100).read(20) + " " + new Holder(3).total);
		Sys.println(Main.count);
	}
}
