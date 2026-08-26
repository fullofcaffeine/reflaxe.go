typedef Entry = {
	final name:String;
	final count:Int;
}

class Main {
	static function main():Void {
		final entries = new Map<String, Entry>();
		entries.set("first", {name: "alpha", count: 3});

		final copied = [for (entry in entries) entry];
		final entry = copied[0];
		Sys.println('${entry.name}:${entry.count}');
	}
}
