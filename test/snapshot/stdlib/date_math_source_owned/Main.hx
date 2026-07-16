class Main {
	static function invalidDateThrows():Bool {
		try {
			Date.fromString("not-a-date");
		} catch (_:Dynamic) {
			return true;
		}
		return false;
	}

	static function main() {
		var value = new Date(2024, 1, 29, 12, 34, 56);
		Sys.println(value.toString());
		Sys.println(Math.isFinite(Math.sqrt(4.0)) && Math.min(Math.PI, 4.0) == Math.PI);
		Sys.println(invalidDateThrows());
	}
}
