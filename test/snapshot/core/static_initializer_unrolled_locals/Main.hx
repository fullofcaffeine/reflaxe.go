/**
	Regression: a static or module-level initializer that declares several Haxe
	locals with the same name must give each one its own Go variable.

	The Haxe typer unrolls `for (x in [literal, ...])` into one block per element,
	and each block declares its own `x`. Initializers used to emit every copy as
	`x := ...` in one Go scope, which Go rejects ("no new variables on left side
	of :="). Closures created in the loop must keep their own iteration's value.
**/
enum abstract FieldKind(String) to String {
	var Text = "text";
}

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

	static function main() {
		for (field in moduleFields)
			Sys.println("module=" + field.tag + ":" + field.kind);
		Sys.println("nested=" + nestedPairs.join(","));
		Sys.println("getters=" + [for (getter in getters) getter()].join(","));
		Sys.println("mutated=" + [for (getter in mutatedGetters) getter()].join(","));
		Sys.println("shadowed=" + shadowedBlocks);
	}
}
