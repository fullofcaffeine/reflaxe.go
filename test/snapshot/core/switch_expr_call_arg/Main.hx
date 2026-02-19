enum Kind {
	A;
	B(value:Int);
}

class Main {
	static function id(value:Int):Int {
		return value;
	}

	static function asInt(kind:Kind):Int {
		return id(switch (kind) {
			case A:
				1;
			case B(value):
				value + 1;
		});
	}

	static function main() {
		Sys.println(Std.string(asInt(A)));
		Sys.println(Std.string(asInt(B(6))));
	}
}
