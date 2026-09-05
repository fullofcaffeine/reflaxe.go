typedef Item = {
	final name:String;
}

enum Outcome<T> {
	Success(value:T);
	Failure(message:String);
}

class Main {
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
	}
}
