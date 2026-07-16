class Main {
	static function localParts(value:Date):String {
		return [
			value.getFullYear(),
			value.getMonth(),
			value.getDate(),
			value.getDay(),
			value.getHours(),
			value.getMinutes(),
			value.getSeconds()
		].join(":");
	}

	static function utcParts(value:Date):String {
		return [
			value.getUTCFullYear(),
			value.getUTCMonth(),
			value.getUTCDate(),
			value.getUTCDay(),
			value.getUTCHours(),
			value.getUTCMinutes(),
			value.getUTCSeconds()
		].join(":");
	}

	static function invalidDateThrows():Bool {
		try {
			Date.fromString("invalid");
		} catch (_:Dynamic) {
			return true;
		}
		return false;
	}

	static function main() {
		var constructed = new Date(2024, 1, 29, 15, 4, 5);
		Sys.println("constructed=" + constructed.toString());
		Sys.println("local=" + localParts(constructed));
		Sys.println("utc=" + utcParts(constructed));
		Sys.println("offset=" + constructed.getTimezoneOffset());

		var full = Date.fromString("2024-02-29 15:04:05");
		var dateOnly = Date.fromString("2024-02-29");
		var timeOnly = Date.fromString("03:04:05");
		Sys.println("parse=" + full.toString() + "|" + dateOnly.toString() + "|" + (timeOnly.getTime() == 11045000.0));
		Sys.println("constructorMatches=" + (constructed.getTime() == full.getTime()));

		var epoch = Date.fromTime(0);
		Sys.println("epochUtc=" + utcParts(epoch));
		Sys.println("fractional=" + (Date.fromTime(1234.5).getTime() == 1234.5));
		var future = new Date(2500, 0, 2, 3, 4, 5);
		Sys.println("wideRange=" + (localParts(future) == "2500:0:2:6:3:4:5"));

		var before = Sys.time() * 1000.0;
		var current = Date.now().getTime();
		var after = Sys.time() * 1000.0;
		Sys.println("nowBounded=" + (current >= before - 2000.0 && current <= after + 2000.0));
		Sys.println("invalid=" + invalidDateThrows());
	}
}
