class Main {
	static function main() {
		var left:String = null;
		var right:String = "value";

		var a:String = left + "x";
		var b:String = right + null;
		var c:String = "p" + 12;

		Sys.println(a);
		Sys.println(b);
		Sys.println(c);
	}
}
