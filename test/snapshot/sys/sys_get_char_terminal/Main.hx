class Main {
	static function main() {
		var arguments = Sys.args();
		var echo = arguments.length > 0 && arguments[0] == "echo";
		Sys.print("ready|");
		try {
			var value = Sys.getChar(echo);
			Sys.println("|" + value + "|");
		} catch (_:haxe.io.Eof) {
			Sys.println("eof|");
		}
	}
}
