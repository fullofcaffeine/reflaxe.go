class Labeled {
	public final label:String;

	public function new(label:String) {
		this.label = label;
	}

	public function toString():String {
		return "<" + label + ">";
	}
}

/**
	What: Exercises the portable semantics of `reverse`, `contains`,
	`lastIndexOf`, `unshift`, `splice`, `resize`, and `toString` on Arrays.
	Why: The Go target stores Arrays in one shared runtime carrier. These methods
	must change that one identity (so every alias sees the change), follow Haxe's
	rules for negative and out-of-range positions, compare with Haxe equality,
	and print with Haxe's Array text.
	How: Mutate Arrays through aliases and nested rows, count receiver
	evaluations, splice and search with edge-case positions, and print nested,
	null-containing, and mixed Arrays through `toString`, `Std.string`, string
	concatenation, interpolation, and `Sys.println`.
**/
class Main {
	static var events:Array<String> = [];

	static function numbers():Array<Int> {
		return [0, 1, 2, 3, 4];
	}

	static function source(label:String, values:Array<Int>):Array<Int> {
		events.push(label);
		return values;
	}

	static function noStart():Null<Int> {
		return null;
	}

	static function spliceCase(pos:Int, len:Int):String {
		final values = numbers();
		final alias = values;
		final removed = values.splice(pos, len);
		return 'splice($pos,$len) -> ${removed.toString()} leaves ${alias.toString()}';
	}

	static function main() {
		// splice: Haxe's position and length rules.
		for (pair in [[1, 2], [-2, 1], [-10, 2], [3, 100], [5, 1], [9, 1], [2, 0], [1, -3], [0, 5]])
			Sys.println(spliceCase(pair[0], pair[1]));
		final emptySplice:Array<Int> = [];
		Sys.println(emptySplice.splice(0, 1).length + " " + emptySplice.length);
		final spliced = numbers();
		final removed = spliced.splice(1, 2);
		removed.push(99);
		Sys.println(spliced.toString() + " " + removed.toString());

		// lastIndexOf and indexOf: start positions inside, before, and past the Array.
		final repeated = [1, 2, 3, 2, 1];
		Sys.println(repeated.lastIndexOf(2)
			+ " "
			+ repeated.lastIndexOf(2, 2)
			+ " "
			+ repeated.lastIndexOf(2, 1)
			+ " "
			+ repeated.lastIndexOf(2, -3));
		Sys.println(repeated.lastIndexOf(1, -1)
			+ " "
			+ repeated.lastIndexOf(1, 5)
			+ " "
			+ repeated.lastIndexOf(1, 99)
			+ " "
			+ repeated.lastIndexOf(9));
		Sys.println(repeated.lastIndexOf(1, 0) + " " + repeated.lastIndexOf(1, -5) + " " + repeated.lastIndexOf(2, 0));
		Sys.println(repeated.indexOf(1, noStart()) + " " + repeated.lastIndexOf(1, noStart()) + " " + repeated.indexOf(2, -2));
		final none:Array<Int> = [];
		Sys.println(none.lastIndexOf(1) + " " + none.lastIndexOf(1, 0) + " " + none.contains(1));

		// contains: value equality for strings and numbers, identity for objects.
		final tag = "b";
		final names = ["a", "bb", null];
		Sys.println(names.contains(tag + tag) + " " + names.contains(null) + " " + names.contains("b") + " " + names.lastIndexOf(null));
		final first = new Labeled("first");
		final objects = [first, null];
		Sys.println(objects.contains(first) + " " + objects.contains(new Labeled("first")) + " " + objects.contains(null));
		final mixed:Array<Dynamic> = [1, "two", null, 2.5];
		Sys.println(mixed.contains(1.0) + " " + mixed.contains("two") + " " + mixed.lastIndexOf(null) + " " + mixed.indexOf(2.5));
		final floats = [1.5, 3.0];
		Sys.println(floats.contains(3) + " " + floats.lastIndexOf(1.5));
		final rows = [[1], [2]];
		final second = rows[1];
		Sys.println(rows.contains(second) + " " + rows.contains([2]) + " " + rows.lastIndexOf(second));
		final maybe:Array<Null<Int>> = [null, 0, null];
		Sys.println(maybe.contains(null) + " " + maybe.contains(0) + " " + maybe.indexOf(null) + " " + maybe.lastIndexOf(null));

		// reverse and unshift change the one Array every alias sees.
		final shared = numbers();
		final alias = shared;
		shared.reverse();
		alias.unshift(9);
		Sys.println(shared.toString() + " " + (shared == alias));
		final grid = [numbers(), numbers()];
		final row = grid[1];
		grid[1].reverse();
		grid[1].unshift(-1);
		grid[0].splice(1, 3);
		Sys.println(row.toString() + " " + grid.toString());
		final emptyRow:Array<String> = [];
		emptyRow.reverse();
		emptyRow.unshift(null);
		Sys.println(emptyRow.toString() + " " + emptyRow.length);

		// Each receiver is evaluated once, before the arguments.
		events = [];
		final evaluated = numbers();
		source("reverse", evaluated).reverse();
		source("unshift", evaluated).unshift(7);
		final cut = source("splice", evaluated).splice(1, 1);
		final found = source("lastIndexOf", evaluated).lastIndexOf(0, -1);
		final present = source("contains", evaluated).contains(4);
		Sys.println(events.join(">") + " " + evaluated.toString() + " " + cut.toString() + " " + found + " " + present);

		// resize truncates, or extends with null.
		final sized = numbers();
		final sizedAlias = sized;
		sized.resize(2);
		Sys.println(sizedAlias.toString() + " " + sizedAlias.length);
		sized.resize(4);
		Sys.println(sized.toString() + " " + sized.length);
		grid[0].resize(0);
		Sys.println(grid.toString());

		// toString and every other conversion use Haxe's Array text.
		final nested = [[1, 2], [], [3]];
		final strings = ["x", null, ""];
		final labeled = [first, null, new Labeled("x")];
		final decimals = [1.5, 2.0, 0.1, -3.25];
		final flags = [true, false];
		final dynamicMix:Array<Dynamic> = [1, "two", [3, [4]], null, 4.5, false];
		final empty:Array<Int> = [];
		Sys.println(nested.toString() + " " + strings.toString() + " " + labeled.toString());
		Sys.println(decimals.toString() + " " + flags.toString() + " " + dynamicMix.toString() + " " + empty.toString());
		Sys.println(Std.string(nested) + " " + Std.string(maybe));
		Sys.println("joined " + nested + " " + strings);
		Sys.println('interpolated ${grid[1]} $labeled');
		Sys.println(nested);
		Sys.println(nested.join(";") + " " + [[["deep"]]].toString());
	}
}
