class Main {
	static function noInt():Null<Int> {
		return null;
	}

	static function main() {
		var s:String = "héllo";
		Sys.println(s.length);
		Sys.println(s.charAt(1));
		Sys.println((cast s.charCodeAt(1) : Int));
		Sys.println(s.substring(1, 4));
		Sys.println(s.substr(2, 2));
		Sys.println(s.substr(-2));

		// A null optional argument reads as omitted (haxe_go-dumy).
		var none = noInt();
		Sys.println(s.substr(1, none));
		Sys.println(s.substring(1, none));
		Sys.println(s.indexOf("l", none));
		Sys.println(s.lastIndexOf("l", none));
	}
}
