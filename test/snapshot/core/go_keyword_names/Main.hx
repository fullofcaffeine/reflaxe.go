/** Instance fields named like Go keywords that Haxe accepts as identifiers. **/
class Fields {
	public var struct:Int;
	public var const:Int;
	public var goto:Int;

	public function new(struct:Int, const:Int, goto:Int) {
		this.struct = struct;
		this.const = const;
		this.goto = goto;
	}
}

/** Instance methods named like the same keywords. **/
class Methods {
	public function new() {}

	public function struct(value:Int):Int
		return value + 1;

	public function const(value:Int):Int
		return value * 2;

	public function goto(value:Int):Int
		return value - 3;
}

/** Statics of a class that is not the selected main class. **/
class Statics {
	public static var goto:Int = 40;

	public static function struct():Int
		return 10;

	public static function const():Int
		return 20;
}

enum Shape {
	Box(struct:Int, const:Int, goto:Int);
}

typedef Record = {
	var struct:Int;
	var const:Int;
	var goto:Int;
}

/**
	Go reserves 25 keywords. Haxe reserves most of them too, but `struct`,
	`const` and `goto` are ordinary Haxe identifiers, so a Haxe local,
	parameter, field, static or method can take one of these names. Printed
	unchanged, the name is a Go keyword and the generated package does not
	parse. Each name must print with the same escape the other Go keywords
	get, while reflection keeps the Haxe name.
**/
class Main {
	static var struct:Int = 1;
	static var const:Int = 2;

	static function goto(struct:Int, const:Int, ?goto:Int):Int {
		return struct + const + (goto == null ? 0 : goto);
	}

	/** Main-class statics print as bare Go names, so no local here may share one. **/
	static function statics():Void {
		Sys.println(Main.struct + Main.const + Main.goto(1, 2) + Main.goto(1, 2, 3));
		Sys.println(Statics.struct() + Statics.const() + Statics.goto);
	}

	static function locals():Void {
		var struct = 3;
		var const = 4;
		var goto = 5;
		Sys.println(struct + const + goto);

		final twice = function(struct:Int):Int return struct * 2;
		Sys.println(twice(21));

		function const(goto:Int):Int
			return goto + 100;
		Sys.println(const(1));

		var total = 0;
		for (goto in 0...4)
			total += goto;
		Sys.println(total);

		try {
			throw "boom";
		} catch (struct:String) {
			Sys.println(struct);
		}

		switch (Box(1, 2, 3)) {
			case Box(struct, const, goto):
				Sys.println(struct * 100 + const * 10 + goto);
		}
	}

	static function members():Void {
		final fields = new Fields(6, 7, 8);
		fields.goto += 1;
		Sys.println(fields.struct + fields.const + fields.goto);
		Sys.println(Reflect.field(fields, "struct") + "," + Reflect.field(fields, "const") + "," + Reflect.field(fields, "goto"));

		final methods = new Methods();
		Sys.println(methods.struct(1) + methods.const(2) + methods.goto(10));
		final bound:Int->Int = methods.const;
		Sys.println(bound(5));
		final dynamicMethods:Dynamic = methods;
		Sys.println(Reflect.callMethod(dynamicMethods, Reflect.field(dynamicMethods, "goto"), [20]));

		final record:Record = {struct: 7, const: 8, goto: 9};
		Sys.println(record.struct + record.const + record.goto);

		final quoted:Dynamic = {
			"break": 1,
			"continue": 2,
			"import": 3,
			"interface": 4
		};
		Sys.println(Reflect.field(quoted, "break")
			+ Reflect.field(quoted, "continue")
			+ Reflect.field(quoted, "import")
			+ Reflect.field(quoted, "interface"));
	}

	static function main():Void {
		statics();
		locals();
		members();
	}
}
