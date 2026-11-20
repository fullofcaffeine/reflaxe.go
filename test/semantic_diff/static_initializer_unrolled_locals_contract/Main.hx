/**
	Locals declared by a field initializer keep one Go variable per Haxe variable.

	The Haxe typer unrolls a short `for` loop over an array literal into one block
	per element. Each block declares its own copy of the loop variable, with the
	same name. A static or module-level initializer must give every copy its own
	Go variable, just like a function body does. Closures created in the loop
	must still see the value of their own iteration.
**/
enum abstract FieldKind(String) to String {
	var Text = "text";
}

// The shape that first failed: a module-level comprehension over a literal.
final moduleFields = [for (tag in ["a", "b"]) {tag: tag, kind: Text}];

class Main {
	static final nestedPairs = [for (x in [1, 2]) for (y in ["p", "q"]) x + y];
	static final getters:Array<() -> String> = [for (name in ["first", "second"])() -> name];
	static final mutatedGetters:Array<() -> Int> = {
		var out:Array<() -> Int> = [];
		for (x in [1, 2]) {
			x += 10;
			out.push(() -> x);
		}
		out;
	};
	static final repeatedLoops = {
		var sum = 0;
		for (x in [1, 2, 3])
			sum += x;
		for (x in [10, 20])
			sum += x * 2;
		sum;
	};
	static final shadowedBlocks = {
		var out = [];
		{
			var x = "left";
			out.push(x);
		}
		{
			var x = "right";
			out.push(x);
		}
		out.join("+");
	};
	static final closureLocals = {
		var x = "outer";
		var inner = () -> {
			var x = "inner";
			return [for (y in ["1", "2"]) x + y].join(",");
		};
		x + "/" + inner();
	};

	final instanceTags = [for (tag in ["i", "j"]) tag.toUpperCase()];

	function new() {}

	static function main() {
		for (field in moduleFields)
			Sys.println("module=" + field.tag + ":" + field.kind);
		Sys.println("nested=" + nestedPairs.join(","));
		Sys.println("getters=" + [for (getter in getters) getter()].join(","));
		Sys.println("mutated=" + [for (getter in mutatedGetters) getter()].join(","));
		Sys.println("repeated=" + repeatedLoops);
		Sys.println("shadowed=" + shadowedBlocks);
		Sys.println("closure=" + closureLocals);
		Sys.println("instance=" + new Main().instanceTags.join(","));

		var localFields = [for (tag in ["x", "y"]) {tag: tag, kind: Text}];
		Sys.println("local=" + [for (field in localFields) field.tag + ":" + field.kind].join(","));
		var localGetters:Array<() -> String> = [];
		for (name in ["l1", "l2"])
			localGetters.push(() -> name);
		Sys.println("localGetters=" + [for (getter in localGetters) getter()].join(","));
	}
}
