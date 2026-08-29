class Main {
	static function main():Void {
		if (GuardedTry.readValue(true, false) != ""
			|| GuardedTry.readValue(false, false) != "value"
			|| GuardedTry.readValue(false, true) != "fallback") {
			throw "unexpected result";
		}
	}
}
