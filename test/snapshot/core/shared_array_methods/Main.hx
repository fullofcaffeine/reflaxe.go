class SharedArrayItem {
	public final name:String;

	public function new(name:String) {
		this.name = name;
	}

	public function toString():String {
		return "Item(" + name + ")";
	}
}

class SharedArrayHolder {
	public final ints:Array<Int>;
	public final names:Array<String>;
	public final items:Array<SharedArrayItem>;

	public function new(ints:Array<Int>, names:Array<String>, items:Array<SharedArrayItem>) {
		this.ints = ints;
		this.names = names;
		this.items = items;
	}
}

/**
	What: Pins generated Go for `reverse`, `contains`, `lastIndexOf`, `unshift`,
	`splice`, `resize`, and `toString` on portable Arrays.
	Why: These Array methods had no Go lowering, so a call on any portable Array
	(a local, a field, or a nested element) produced Go that did not compile.
	`indexOf(null)` also did not compile, and `indexOf` with a null start threw.
	How: Call each method on a local Array, a class field Array, and a nested
	element (`rows[i]`) with Int, String, and class elements. Arrays come from
	functions so the analyzer cannot fold them into scalar locals.
**/
class Main {
	static function makeInts():Array<Int> {
		return [1, 2, 3, 2, 1];
	}

	static function makeNames():Array<String> {
		return ["a", "b", null, "b"];
	}

	static function makeMaybe():Array<Null<Int>> {
		return [null, 1, null];
	}

	static function noStart():Null<Int> {
		return null;
	}

	static function makeRows():Array<Array<Int>> {
		return [for (row in 0...3) [for (col in 0...3) row * 3 + col]];
	}

	static function main() {
		final ints = makeInts();
		final names = makeNames();
		final first = new SharedArrayItem("a");
		final second = new SharedArrayItem("b");
		final items = [first, second, null];
		final holder = new SharedArrayHolder(makeInts(), makeNames(), [second, first]);
		final rows = makeRows();
		final grid = [makeNames(), makeNames()];
		final shelves = [[first], [second, first]];

		// toString on locals, fields, nested elements, nested Arrays, and nulls.
		Sys.println(ints.toString());
		Sys.println(names.toString());
		Sys.println(items.toString());
		Sys.println(holder.ints.toString() + " " + holder.items.toString());
		Sys.println(rows.toString());
		Sys.println(rows[1].toString() + " " + grid[0].toString() + " " + shelves[1].toString());

		// contains with values, null, and object identity.
		Sys.println(ints.contains(3) + " " + ints.contains(7));
		final wanted = "b";
		Sys.println(names.contains(wanted + "") + " " + names.contains(null) + " " + names.contains("z"));
		Sys.println(items.contains(first) + " " + items.contains(new SharedArrayItem("a")) + " " + items.contains(null));
		Sys.println(holder.names.contains("a") + " " + holder.items.contains(second));
		Sys.println(rows[2].contains(7) + " " + grid[1].contains(null) + " " + shelves[0].contains(second));

		// lastIndexOf with and without a starting position.
		Sys.println(ints.lastIndexOf(2) + " " + ints.lastIndexOf(2, 2) + " " + ints.lastIndexOf(2, -3) + " " + ints.lastIndexOf(9));
		Sys.println(names.lastIndexOf("b") + " " + names.lastIndexOf(null) + " " + items.lastIndexOf(second));
		Sys.println(holder.ints.lastIndexOf(1, 10) + " " + holder.items.lastIndexOf(first, 0));
		Sys.println(rows[0].lastIndexOf(2) + " " + grid[1].lastIndexOf("b", 2) + " " + shelves[1].lastIndexOf(first));

		// A null value and a null start (a null start means "omitted").
		final maybe = makeMaybe();
		Sys.println(maybe.indexOf(null) + " " + maybe.lastIndexOf(null) + " " + maybe.contains(1) + " " + names.indexOf(null));
		Sys.println(ints.indexOf(1, noStart()) + " " + ints.lastIndexOf(1, noStart()));

		// The Haxe API documents -1 when a negative start stays below 0 after
		// counting from the end; JS, C++, and Python agree. The `--interp`
		// evaluator clamps it to 0 and prints 0 here, so semantic-diff omits it.
		Sys.println(holder.ints.lastIndexOf(1, -6));

		// reverse in place.
		ints.reverse();
		names.reverse();
		items.reverse();
		holder.ints.reverse();
		rows[1].reverse();
		shelves[1].reverse();
		Sys.println(ints.toString() + " " + names.toString() + " " + items.toString());
		Sys.println(holder.ints.toString() + " " + rows[1].toString() + " " + shelves[1].toString());

		// unshift adds at the front.
		ints.unshift(0);
		names.unshift(null);
		items.unshift(second);
		holder.names.unshift("first");
		rows[2].unshift(-1);
		grid[0].unshift("z");
		shelves[0].unshift(second);
		Sys.println(ints.toString() + " " + names.toString() + " " + items.toString());
		Sys.println(holder.names.toString() + " " + rows[2].toString() + " " + grid[0].toString() + " " + shelves[0].toString());

		// splice removes and returns a range.
		final removed = ints.splice(1, 2);
		Sys.println(removed.toString() + " " + ints.toString());
		Sys.println(names.splice(-2, 1).toString() + " " + names.toString());
		Sys.println(items.splice(0, 1).toString() + " " + items.toString());
		Sys.println(holder.ints.splice(3, 100).toString() + " " + holder.ints.toString());
		Sys.println(rows[0].splice(1, 1).toString() + " " + rows.toString());
		Sys.println(grid[1].splice(-10, 2).length + " " + grid[1].toString());
		Sys.println(shelves[1].splice(1, 5).toString() + " " + shelves.toString());

		// resize truncates, or extends with null.
		final sized = makeInts();
		sized.resize(2);
		rows[2].resize(6);
		holder.names.resize(1);
		Sys.println(sized.toString() + " " + rows[2].toString() + " " + holder.names.toString());
	}
}
