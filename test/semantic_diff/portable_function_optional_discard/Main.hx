class Main {
	static function make():(?Null<Int>) -> Int {
		return function(value:Null<Int> = 9):Int {
			final result:Int = value == null ? 9 : value;
			Sys.println(result);
			return result;
		};
	}

	static function invoke(callback:(?Null<Int>) -> Void) {
		callback();
		callback(null);
		callback(0);
		callback(4);
	}

	static function main() {
		invoke(make());
	}
}
