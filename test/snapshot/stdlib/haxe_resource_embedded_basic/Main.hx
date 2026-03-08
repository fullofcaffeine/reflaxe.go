class Main {
	static function main() {
		var names = haxe.Resource.listNames();
		Sys.println(names.length);
		Sys.println(names.length > 0 ? names[0] : "<none>");
		Sys.println(StringTools.replace(haxe.Resource.getString("greet"), "\n", "\\n"));
		var bytes = haxe.Resource.getBytes("greet");
		Sys.println(bytes == null ? "null" : bytes.length);
		Sys.println(haxe.Resource.getString("missing") == null);
		Sys.println(haxe.Resource.getBytes("missing") == null);
	}
}
