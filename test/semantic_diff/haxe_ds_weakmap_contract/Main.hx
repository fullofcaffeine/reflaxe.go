private class WeakMapKey {
	public function new() {}
}

class Main {
	static function main() {
		try {
			new haxe.ds.WeakMap<WeakMapKey, Int>();
			Sys.println("constructed");
		} catch (err:haxe.exceptions.NotImplementedException) {
			Sys.println("not_impl=" + err.message);
			Sys.println("typed=" + Std.string(Std.isOfType(err, haxe.exceptions.NotImplementedException)));
		}
	}
}
