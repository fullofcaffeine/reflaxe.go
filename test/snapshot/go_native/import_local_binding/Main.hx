@:go.import("bytes") @:go.name("Buffer") extern class Buffer {
	@:go.name("Len") function length():Int;
}

@:go.import("bytes") extern class Buffers {
	@:go.name("NewBufferString") static function create(value:String):Buffer;
}

@:go.import("strings") extern class Strings {
	@:go.name("Contains") static function contains(value:String, part:String):Bool;
}

class Box<T> {
	public var value:T;

	public function new(value:T)
		this.value = value;
}

class Main {
	static function main() {
		final bytes = new Box(Buffers.create("abc"));
		if (bytes.value.length() != 3)
			throw "generic native recovery";
		if (!Strings.contains("abc", "b"))
			throw "real second import";
		if (Local.read() != 7 || Local.argument(new Local()) != 7 || Local.loops() != 6 || Local.catchValue() != "caught")
			throw "local values";
		if (Local.closure() != 73 || Local.runtimeAlias() != 10)
			throw "nested or runtime binding";
		if (Local.suffix() != 15)
			throw "suffix collision";
		Sys.println("import binding: PASS");
	}
}
