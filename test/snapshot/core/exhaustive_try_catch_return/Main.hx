class Main {
	static function readValue(skip:Bool, fail:Bool):String {
		if (skip) {
			return "";
		}
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
		if (readValue(true, false) != "" || readValue(false, false) != "value" || readValue(false, true) != "fallback") {
			throw "unexpected result";
		}
	}
}
