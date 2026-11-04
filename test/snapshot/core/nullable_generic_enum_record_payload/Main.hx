typedef Item = {
	final name:String;
}

interface Named {
	public function name():String;
}

class NamedItem implements Named {
	public function new() {}

	public function name():String
		return "kept";
}

enum Outcome<T> {
	Success(value:T);
	Failure(message:String);
}

class Main {
	static function describeInterface(outcome:Outcome<Null<Named>>):String {
		return switch outcome {
			case Failure(message): 'failure=$message';
			case Success(value): value == null ? "missing" : 'name=${value.name()}';
		};
	}

	static function describe(outcome:Outcome<Null<Item>>):String {
		return switch outcome {
			case Failure(message): 'failure=$message';
			case Success(value): value == null ? "missing" : 'name=${value.name}';
		};
	}

	static function main():Void {
		if (describe(Success(null)) != "missing")
			throw "null payload changed";
		if (describe(Success({name: "kept"})) != "name=kept")
			throw "record payload changed";
		if (describeInterface(Success(null)) != "missing")
			throw "null interface payload changed";
		if (describeInterface(Success(new NamedItem())) != "name=kept")
			throw "interface payload changed";
	}
}
