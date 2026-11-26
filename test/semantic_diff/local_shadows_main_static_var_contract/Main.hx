/**
	A local named like a static var of the main class must not change what
	`Main.total` means. Haxe code that writes `Main.total` reads and writes the
	static even where a local, parameter, loop variable, catch variable or
	closure capture of the same name is in scope.

	Main-class statics print as bare Go package-level names, so a Go local with
	the same name would shadow the static and every `Main.total` in its scope
	would silently use the local. Each shape is the first declaration of its
	name in its function, so the printed values tell the static and the local
	apart. A local function named like the static does not compile in that
	case; `test/snapshot/core/local_shadows_main_static` covers it.
**/
class Other {
	public static function run(total:Int):Int {
		Main.total += total;
		return total * 1000 + Main.total;
	}
}

class Main {
	public static var total:Int = 1;

	static function local():String {
		var total = 3;
		var before = total + ":" + Main.total;
		Main.total = Main.total + total;
		Main.total++;
		total++;
		return before + " " + total + ":" + Main.total;
	}

	static function param(total:Int):String {
		Main.total += total;
		return total + ":" + Main.total;
	}

	/** A bound the typer cannot see keeps the loop from being unrolled. */
	static function loop(limit:Int):Int {
		var sum = 0;
		for (total in 0...limit)
			sum += total * 100 + Main.total;
		return sum;
	}

	static function caught():Int {
		try {
			throw 7;
		} catch (total:Int) {
			Main.total += total;
			return total * 1000 + Main.total;
		}
	}

	static function closure():String {
		var total = 40;
		final read = () -> total + ":" + Main.total;
		final write = (step:Int) -> {
			Main.total += step;
			total += step * 10;
		};
		write(3);
		return read() + " " + total + ":" + Main.total;
	}

	static function main():Void {
		Sys.println("local " + local());
		Sys.println("param " + param(10));
		Sys.println("loop " + loop(3));
		Sys.println("catch " + caught());
		Sys.println("closure " + closure());
		Sys.println("other class " + Other.run(4));
		Sys.println("final " + Main.total);
	}
}
