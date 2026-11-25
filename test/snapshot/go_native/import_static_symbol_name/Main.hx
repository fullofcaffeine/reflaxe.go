@:go.import("strings") extern class NativeStrings {
	@:go.name("ToUpper") static function toUpper(value:String):String;
}

/**
	Main-class statics print as bare Go package-level names. A name that equals
	an imported Go package (`strings` from the extern above, `hxrt` from the
	runtime) must get the qualified `Main_*` form instead, or Go rejects the
	package as "already declared through import of package".
**/
class Main {
	static var hxrt:String = "runtime";

	static function strings(value:String):String
		return NativeStrings.toUpper(value);

	static function plain(value:String):String
		return value + "!";

	static function main():Void {
		final convert:String->String = strings;
		Sys.println(strings("x") + convert("y") + plain(hxrt));
	}
}
