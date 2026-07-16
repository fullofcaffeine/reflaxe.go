class Main {
	static function main() {
		var values = new haxe.ds.StringMap<Int>();
		values.set("one", 1);
		values.exists("one");
	}
}
