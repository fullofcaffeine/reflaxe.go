enum MaybeFloat {
	Some(value:Null<Float>);
}

class Main {
	static function describe(value:MaybeFloat):String {
		return switch (value) {
			case Some(payload):
				payload == null ? "missing" : "value=" + Std.string(payload);
		};
	}

	static function main():Void {
		Sys.println(describe(Some(null)));
		Sys.println(describe(Some(1.5)));
	}
}
