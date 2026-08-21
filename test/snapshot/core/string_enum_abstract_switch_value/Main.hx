enum abstract OutputMode(String) {
	final Human = "human";
	final Json = "json";
}

class Main {
	static function describe(mode:OutputMode):String {
		return switch mode {
			case Human: "human";
			case Json: "json";
		};
	}

	static function main():Void {
		if (describe(OutputMode.Human) != "human")
			throw "human enum-abstract switch mismatch";
		if (describe(OutputMode.Json) != "json")
			throw "json enum-abstract switch mismatch";

		var statementMatched = false;
		switch OutputMode.Json {
			case Human:
				throw "statement switch matched the wrong string value";
			case Json:
				statementMatched = true;
		}
		if (!statementMatched)
			throw "statement enum-abstract switch missed";
	}
}
