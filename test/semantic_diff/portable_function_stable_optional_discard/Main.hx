class Main {
	static function f(x:Int = 9):Int {
		Sys.println(x);
		return x;
	}

	static function main() {
		var callback:(?Int) -> Void = f;
		callback();
	}
}
