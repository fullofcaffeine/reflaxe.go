class Main {
	static function main() {
		var names = haxe.Resource.listNames();
		Sys.println(names.length);
		Sys.println(haxe.Resource.getString("missing") == null);
		Sys.println(haxe.Resource.getBytes("missing") == null);
	}
}
