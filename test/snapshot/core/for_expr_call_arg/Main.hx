class Main {
	static function consume(v:Void):Void {}

	static function main() {
		consume(for (i in 0...3) {
			var x = i;
			Sys.println(Std.string(x));
		});
		Sys.println("ok");
	}
}
