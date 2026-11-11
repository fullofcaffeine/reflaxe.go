/** Optional boolean fields retain Haxe condition semantics through structural storage. */
class Main {
	static var record:{?enabled:Bool} = {enabled: true};

	static function selected():String {
		if (record.enabled)
			return "on";
		return "off";
	}

	static function main():Void {
		Sys.println(selected());
		record.enabled = false;
		Sys.println(selected());
		record = {};
		Sys.println(selected());
		var visits = 0;
		var checks = 0;
		record = {enabled: true};
		while ({
			checks++;
			record.enabled;
		}) {
			visits++;
			record.enabled = false;
		}
		Sys.println(visits + ":" + checks);
		record = {};
		do {
			visits++;
		} while ({
			checks++;
			record.enabled;
			});
		Sys.println(visits + ":" + checks);
		record = {enabled: true};
		Sys.println(({checks++; record.enabled;}) ? "yes" : "no");
		Sys.println(checks);
	}
}
