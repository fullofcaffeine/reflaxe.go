class Main {
	static function main() {
		var expression = new EReg("a", "g");
		Sys.println(expression.replace("a-a", "b"));
	}
}
