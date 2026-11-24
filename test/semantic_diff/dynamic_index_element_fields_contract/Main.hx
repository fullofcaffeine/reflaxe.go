class Point {
	public var x:Int;

	public function new(x:Int) {
		this.x = x;
	}

	public function twice():Int {
		return x * 2;
	}
}

/**
	Fields and method calls on elements read from a `Dynamic` index must
	behave as in `--interp` (haxe_go-buow).

	Haxe types `d[0]` with an unknown type, and a field access on it makes
	that type a structure, whatever the element holds. `d[0].length` reads
	the length of a String or an Array element, `d[0].push(x)` and
	`d[0].charAt(i)` call its methods, and fields and methods of anonymous
	objects, class instances and parsed JSON values work as before, also
	through a field of the element (`d[0].inner.label`) and through the
	result of a method call (`d[0].toUpperCase().length`). Writes, `++` and
	compound assignment change the element, operators and conditions apply
	to the value that the field holds, the receiver runs once, before the
	arguments, and a field access or call on a null element raises an error.
	Left out: fields of an element of a typed Array that a `Dynamic` index
	read gives (`d[0][0].length`), and the text of the null-element error
	(`--interp` raises `field access on null`, haxe.go `Null Access`).
**/
class Main {
	static var order = "";

	static function dyn(value:Dynamic):Dynamic {
		return value;
	}

	static function step(name:String, value:Dynamic):Dynamic {
		order += name;
		return value;
	}

	static function attempt(name:String, run:() -> Dynamic) {
		var text = try Std.string(run()) catch (e:Dynamic) "error " + e;
		Sys.println(name + " = " + text);
	}

	static function fails(name:String, run:() -> Dynamic) {
		var text = try {
			run();
			"no error";
		} catch (e:Dynamic)
			"error";
		Sys.println(name + " = " + text);
	}

	static function main() {
		lengths();
		methods();
		structures();
		mutations();
		parsed();
		nulls();
		evaluationOrder();
	}

	static function lengths() {
		var words:Dynamic = dyn(["alpha", "be", ""]);
		var rows:Dynamic = dyn([[1, 2], [3], []]);
		Sys.println("length " + words[0].length + " " + words[2].length + " " + rows[0].length + " " + rows[2].length);
		Sys.println(words[1].length);
		var n:Int = words[0].length;
		var f:Float = rows[0].length;
		var text:String = words[0].length;
		var maybe:Null<Int> = rows[1].length;
		Sys.println("typed " + (n * 2) + " " + (f / 4) + " " + text + " " + maybe);
		Sys.println("operators " + (words[0].length + 1) + " " + (words[0].length - rows[0].length) + " " + (words[0].length > 3) + " "
			+ (rows[1].length == 1) + " " + (words[0].length / 2) + " " + (-rows[0].length));
		if (rows[1].length == 1)
			Sys.println("condition one");
		var total:Dynamic = 0;
		for (k in 0...rows.length)
			total += rows[k].length;
		Sys.println("loop " + total + ' ${words[1].length}');
		rows[0].length = 1;
		Sys.println("length write " + rows[0]);
	}

	static function methods() {
		var words:Dynamic = dyn(["alpha", "be"]);
		var rows:Dynamic = dyn([[1, 2], [3]]);
		rows[0].push(9);
		rows[1].unshift(0);
		Sys.println("statements " + rows);
		Sys.println("push " + rows[0].push(10) + " pop " + rows[0].pop() + " " + rows[0]);
		Sys.println("string " + words[1].charAt(0) + " " + words[0].indexOf("p") + " " + words[0].split("l") + " " + words[0].substr(1, 2));
		Sys.println("join " + rows[1].join("-") + " " + rows[0].indexOf(9) + " " + rows[0].contains(2));
		var c:String = words[0].charAt(1);
		var size:Int = rows[0].push(11);
		Sys.println("typed " + c.toUpperCase() + " " + (size + 1));
		Sys.println("chain " + words[0].toUpperCase().length + " " + words[0].toUpperCase().charAt(0) + " " + rows[0].concat([7]).length);
		Sys.println("call result " + (rows[1].pop() + 1) + " " + (words[1].charAt(1) + "!"));
		attempt("missing method", () -> words[0].nope());
	}

	static function structures() {
		var objs:Dynamic = dyn([
			{
				name: "ann",
				n: 1,
				f: (v:Int) -> v + 1,
				inner: {label: "deep", tags: ["a", "b"]}
			},
			new Point(4)
		]);
		Sys.println("anon " + objs[0].name + " " + objs[0].n + " " + objs[0].f(3) + " " + objs[0].inner.label + " " + objs[0].inner.tags.length);
		Sys.println("class " + objs[1].x + " " + objs[1].twice());
		var name:String = objs[0].name;
		var n:Int = objs[0].n;
		var x:Float = objs[1].x;
		Sys.println("typed " + name.toUpperCase() + " " + (n + 1) + " " + (x / 8));
		Sys.println("missing " + objs[0].nope + " " + objs[1].nope);
		Sys.println("label length " + objs[0].inner.label.length + " " + objs[0].name.charAt(0));
	}

	static function mutations() {
		var objs:Dynamic = dyn([{name: "ann", n: 1}, new Point(4)]);
		objs[0].name = "bob";
		objs[0].n++;
		objs[0].n += 5;
		objs[0].name += "!";
		objs[1].x = 9;
		objs[1].x *= 2;
		Sys.println("writes " + objs[0].name + " " + objs[0].n + " " + objs[1].x);
		Sys.println("values " + (objs[0].n++) + " " + (++objs[0].n) + " " + (objs[0].n -= 2) + " " + (objs[0].name = "cy"));
		objs[0].added = true;
		Sys.println("added " + objs[0].added + " " + objs[0].n);
	}

	static function parsed() {
		var json:Dynamic = haxe.Json.parse('{"items":[{"name":"ann","tags":["a","b"]},{"name":"bo","tags":[]}]}');
		Sys.println("json " + json.items[0].name.length + " " + json.items[0].tags.length + " " + json.items[1].tags.length);
		json.items[0].tags.push("c");
		json.items[1].name = "bob";
		Sys.println("json " + json.items[0].tags.join(",") + " " + json.items[1].name + " " + json.items[0].name.toUpperCase());
	}

	static function nulls() {
		var mixed:Dynamic = dyn([null, "x"]);
		fails("null field", () -> mixed[0].length);
		fails("null call", () -> mixed[0].charAt(0));
		fails("null write", () -> mixed[0].length = 1);
		fails("null step", () -> mixed[0].n++);
		fails("outside field", () -> mixed[5].length);
		fails("present field", () -> mixed[1].length);
	}

	static function evaluationOrder() {
		var rows:Dynamic = dyn([[1], [2]]);
		step("r", rows)[step("i", 0)].push(step("a", 5));
		Sys.println("call order " + order + " " + rows);
		order = "";
		var read = step("r", rows)[step("i", 1)].length;
		Sys.println("read order " + order + " " + read);
		order = "";
		var objs:Dynamic = dyn([{n: 1}]);
		step("r", objs)[step("i", 0)].n += step("v", 2);
		Sys.println("write order " + order + " " + objs[0].n);
	}
}
