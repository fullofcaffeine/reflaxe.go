class Main {
	static function main() {
		var s = "abc";
		Sys.println("in=" + s.charCodeAt(1));
		Sys.println("neg=" + s.charCodeAt(-1));
		Sys.println("oob=" + s.charCodeAt(99));
	}
}
