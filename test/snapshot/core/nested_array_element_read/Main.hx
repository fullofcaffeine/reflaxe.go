typedef NestedArrayLayout = {
	final layers:Array<Array<String>>;
}

class NestedArrayBoard {
	public final cells:Array<Array<Int>>;

	public function new(cells:Array<Array<Int>>) {
		this.cells = cells;
	}
}

/**
	What: Pins generated Go for elements of nested portable Arrays used directly.
	Why: The shared Array carrier stores elements as erased `any`. An element that
	is itself an Array (`grid[row][col]`, `rows[i].push(x)`) or an Int used as an
	index (`counts[order[i]]`) must regain its static type before Go can call a
	carrier method on it or index with it; it once produced Go that did not compile.
	How: Read, write, and mutate nested elements through locals, a class field, a
	structure field, three levels of nesting, Array methods, and indexes that are
	themselves Array elements or a `Null<Int>` local. Arrays come from functions
	so the analyzer cannot fold them into scalar locals.
**/
class Main {
	static function makeGrid():Array<Array<Int>> {
		return [for (row in 0...3) [for (col in 0...3) row * 3 + col]];
	}

	static function makeNames():Array<Array<String>> {
		return [for (row in 0...2) [for (col in 0...2) String.fromCharCode(97 + row * 2 + col)]];
	}

	static function makeMaybe():Array<Null<Int>> {
		return [2, null];
	}

	static function label(value:Int):String {
		return "<" + value + ">";
	}

	static function main() {
		final grid = makeGrid();
		final names = makeNames();

		// Reads used directly in expressions.
		Sys.println(grid[1][2] + 1);
		if (names[1][0] == "c")
			Sys.println("found c");
		Sys.println(label(grid[2][0]));
		Sys.println('cell ${grid[0][1]} name ${names[0][1]}');
		Sys.println(names[0][0] + names[1][1]);
		Sys.println(grid[2][1] > grid[1][2]);
		Sys.println(grid[1][1] == grid[1][1]);

		// Writes and mutations of one nested slot.
		grid[0][0] += 10;
		grid[2][2]++;
		Sys.println(grid[0][0] + ":" + (++grid[2][2]));
		grid[1][1] = 40;
		final assigned = (grid[1][0] = 30);
		Sys.println(assigned + grid[1][1]);

		// Reads through a structure field and a class field.
		final layout:NestedArrayLayout = {layers: names};
		Sys.println(layout.layers[1][1] + "!");
		final board = new NestedArrayBoard(grid);
		Sys.println(board.cells[2][1] * 2);
		board.cells[2][1] -= 1;
		Sys.println(grid[2][1]);

		// Three levels of nesting.
		final cube = [for (a in 0...2) [for (b in 0...2) [for (c in 0...2) a * 4 + b * 2 + c]]];
		cube[1][0][1] += 100;
		Sys.println(cube[1][0][1] + cube[0][1][0]);

		// Indexes that are themselves Array elements.
		final order = [for (index in 0...3) if (index != 1) index];
		final counts = [for (_ in 0...3) 0];
		counts[order[1]] = 7;
		counts[order[0]] += 2;
		Sys.println(counts[order[1]] + counts[order[0]]);
		Sys.println(grid[order[0]][order[1]]);
		final maybe = makeMaybe();
		Sys.println(grid[0][maybe[0]]);
		final nullableIndex:Null<Int> = maybe[0];
		Sys.println(grid[1][nullableIndex]);
		final vector = new haxe.ds.Vector<Int>(3);
		vector[order[1]] = 5;
		Sys.println(vector[order[1]]);

		// Array methods called on a nested element.
		grid[0].push(9);
		Sys.println(grid[0].length + ":" + grid[0].pop());
		Sys.println(grid[1].join("-") + " " + grid[1].copy().length);
		Sys.println(grid[1].indexOf(40) + ":" + grid[0].concat(grid[1]).length);
		grid[2].insert(order[0], -1);
		Sys.println(grid[2].remove(-1) + ":" + grid[2].shift());
		Sys.println(grid[2].slice(order[0]).length);
		grid[2].sort((left, right) -> right - left);
		Sys.println(grid[2][0]);
		for (value in grid[1])
			Sys.println(value);

		// A Map element and a local function that redraws one nested cell.
		final seen = [for (_ in 0...2) new Map<String, Bool>()];
		if (!seen[1].exists("k"))
			seen[1].set("k", true);
		Sys.println(seen[1].exists("k"));
		final canvas = [for (_ in 0...2) [for (_ in 0...3) " "]];
		function draw(y:Int, x:Int, incoming:String):Void {
			canvas[y][x] = canvas[y][x] == " " ? incoming : "+";
		}
		draw(1, 2, "-");
		draw(1, 2, "|");
		draw(0, 0, "*");
		Sys.println([for (line in canvas) line.join("")].join("/"));
	}
}
