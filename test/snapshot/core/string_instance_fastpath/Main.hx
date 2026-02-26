class Main {
	static function main() {
		var s:String = "héllo";
		Sys.println(s.length);
		Sys.println(s.charAt(1));
		Sys.println((cast s.charCodeAt(1) : Int));
		Sys.println(s.substring(1, 4));
		Sys.println(s.substr(2, 2));
		Sys.println(s.substr(-2));
	}
}
