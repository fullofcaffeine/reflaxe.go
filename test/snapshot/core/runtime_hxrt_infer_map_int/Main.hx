class Main {
	static function main() {
		var values = new haxe.ds.IntMap<String>();
		values.set(1, "one");
		values.exists(1);
	}
}
