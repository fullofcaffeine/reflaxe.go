/**
	What: Exercises portable semantics of reading and writing nested Array elements.
	Why: The Go target stores Array elements as erased values and must recover an
	inner Array or an Int position before using it. That recovery must keep the
	inner Array's identity and Haxe's once-only, left-to-right evaluation order.
	How: Mutate inner rows through nested indexes and Array methods and observe
	them through aliases, count side effects in compound writes, grow an inner
	row, append to nested strings, and index with positions read from Arrays.
**/
class Main {
	static var events:Array<String> = [];

	static function makeGrid(rows:Int, cols:Int):Array<Array<Int>> {
		return [for (row in 0...rows) [for (col in 0...cols) row * 10 + col]];
	}

	static function pick(label:String, value:Int):Int {
		events.push(label);
		return value;
	}

	static function rowOf(grid:Array<Array<Int>>, label:String):Array<Array<Int>> {
		events.push(label);
		return grid;
	}

	static function render(grid:Array<Array<Int>>):String {
		return [for (row in grid) row.join(",")].join(" | ");
	}

	static function main() {
		final grid = makeGrid(3, 3);
		final middle = grid[1];

		// Writes through a nested index reach every alias of the inner row.
		grid[1][0] = 100;
		grid[1][1] += 5;
		grid[1][2]++;
		Sys.println(middle.join(","));
		Sys.println(grid[1][0] + middle[1] + grid[1][2]);

		// Array methods on a nested element mutate the same inner row.
		grid[0].push(3);
		grid[0].insert(0, -1);
		Sys.println(grid[0].length + ":" + grid[0].shift() + ":" + grid[0].pop());
		grid[2].sort((left, right) -> right - left);
		Sys.println(render(grid));
		final joined = grid[0].concat(grid[2]);
		joined[0] = 999;
		Sys.println(joined.length + ":" + grid[0][0] + ":" + grid[1].indexOf(15));

		// Receiver, row, column, and right side are each evaluated once, in order.
		events = [];
		rowOf(grid, "grid")[pick("row", 2)][pick("col", 1)] += pick("value", 7);
		Sys.println(events.join(">") + " = " + grid[2][1]);
		events = [];
		final old = rowOf(grid, "grid")[pick("row", 0)][pick("col", 0)]++;
		Sys.println(events.join(">") + " = " + old + "->" + grid[0][0]);

		// Positions read from Arrays select the right slot.
		final order = [2, 0, 1];
		final counts = [for (_ in 0...3) 0];
		for (round in 0...5)
			counts[order[round % 3]] += round;
		Sys.println(counts.join(","));
		Sys.println(grid[order[0]][order[2]] + ":" + grid[order[1]].slice(order[2]).join(","));

		// Writing past the end grows the inner row.
		grid[1][5] = 7;
		Sys.println(grid[1].length + ":" + grid[1][5] + ":" + middle.length);

		// Nested String elements: compare, append, and interpolate.
		final names = [["ada", "bob"], ["cy", "dee"]];
		names[1][0] += "!";
		if (names[1][0] == "cy!")
			Sys.println('appended ${names[1][0]} beside ${names[0][1]}');
	}
}
