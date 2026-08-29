class Main {
	static function readValue(fail:Bool):String {
		try {
			if (fail) {
				throw "failed";
			}
			return "value";
		} catch (_:Dynamic) {
			return "fallback";
		}
	}

	static function main():Void {
		if (readValue(false) != "value" || readValue(true) != "fallback") {
			throw "unexpected result";
		}
	}
}
